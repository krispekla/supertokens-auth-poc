package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mailgun/mailgun-go/v4"

	"github.com/supertokens/supertokens-golang/ingredients/emaildelivery"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/emailverification"
	"github.com/supertokens/supertokens-golang/recipe/emailverification/evmodels"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/recipe/userroles"
	"github.com/supertokens/supertokens-golang/supertokens"
)

var db *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	fmt.Printf("DATABASE_URL: %s\n", os.Getenv("DATABASE_URL"))

	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Failed to open database:", err)
		return
	}

	apiBasePath := "/auth"
	websiteBasePath := "/auth"
	err = supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: os.Getenv("SUPERTOKEN_CORE_SVC_URL"),
		},
		AppInfo: supertokens.AppInfo{
			AppName:         "test",
			APIDomain:       "http://localhost:3003",
			WebsiteDomain:   "http://localhost:8084",
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			thirdparty.Init(&tpmodels.TypeInput{ /*TODO: See next step*/ }),
			emailpassword.Init(nil),
			emailverification.Init(evmodels.TypeInput{
				Mode: evmodels.ModeRequired,
				EmailDelivery: &emaildelivery.TypeInput{
					Override: func(originalImplementation emaildelivery.EmailDeliveryInterface) emaildelivery.EmailDeliveryInterface {
						originalSendEmail := *originalImplementation.SendEmail

						(*originalImplementation.SendEmail) = func(input emaildelivery.EmailType, userContext supertokens.UserContext) error {
							// TODO: create and email verification email
							(*originalImplementation.SendEmail) = CustomSendEmail

							// Or use the original implementation which calls the default service,
							// or a service that you may have specified in the EmailDelivery object.
							return originalSendEmail(input, userContext)
						}

						return originalImplementation
					},
				},
			}),
			session.Init(nil), // initializes session features
			userroles.Init(nil),
		},
	})

	if err != nil {
		panic(err.Error())
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:8084"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: append([]string{"Content-Type"},
			supertokens.GetAllCORSHeaders()...),
		AllowCredentials: true,
	}))

	r.Use(supertokens.Middleware)
	// Public routes
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Live"))
	})

	r.Get("/mail-test", func(w http.ResponseWriter, r *http.Request) {
		msg := r.URL.Query().Get("msg")
		emailAdr := r.URL.Query().Get("email")
		id, err := SendSimpleMessage(r.Context(), os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_API_KEY"), msg, emailAdr, os.Getenv("MAILGUN_SENDER_EMAIL"))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte(fmt.Sprint("Message sent: ", id)))
	})

	// Private routes
	r.Group(func(pr chi.Router) {
		pr.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, World!"))
		})

		pr.Get("/protected", session.VerifySession(&sessmodels.VerifySessionOptions{},
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("This is another route"))
			}))
	})

	http.ListenAndServe(":3003", r)
}

// TODO: Remove and refactor this with func below for sending email
func SendSimpleMessage(ctx context.Context, domain, apiKey, msg, emailAdd, sender string) (string, error) {
	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(
		fmt.Sprintf("Supertoken Generic <%s>", sender),
		"Testing sp",
		msg,
		emailAdd,
	)
	_, id, err := mg.Send(ctx, m)
	return id, err
}

// CustomSendEmail implements the SendEmail method of emaildelivery.EmailDeliveryInterface
func CustomSendEmail(input emaildelivery.EmailType, userContext supertokens.UserContext) error {
	// TODO: create and email verification email
	_, err := SendSimpleMessage(
		context.Background(),
		os.Getenv("MAILGUN_DOMAIN"),
		os.Getenv("MAILGUN_API_KEY"),
		fmt.Sprintf("Verify email: %s", input.EmailVerification.EmailVerifyLink),
		input.EmailVerification.User.Email,
		os.Getenv("MAILGUN_SENDER_EMAIL"),
	)
	if err != nil {
		return err
	}
	return nil
}
