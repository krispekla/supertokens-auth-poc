
	"github.com/joho/godotenv"

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
