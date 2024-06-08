package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
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
			session.Init(nil), // initializes session features
			dashboard.Init(nil),
		},
	})

	if err != nil {
		panic(err.Error())
	}

	r := chi.NewRouter()

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:8084"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: append([]string{"Content-Type"},
			supertokens.GetAllCORSHeaders()...),
		AllowCredentials: true,
	}))

	r.Use(supertokens.Middleware)

	// Routes

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	r.Get("/another-route", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is another route"))
	})

	http.ListenAndServe(":3003", r)
}
