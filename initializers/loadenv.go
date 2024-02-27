package initializers

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnvInitializers() {
	err := godotenv.Load("initializers/.env")
	if err != nil {
		fmt.Println(err)
	}
}
