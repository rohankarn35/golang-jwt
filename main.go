package main

import (
	"encoding/json"
	"fmt"
	"golang-auth/config"
	"golang-auth/models"
	"time"
)

func main() {
	// Generate a new JWT
	userId := "60d5ec49f1a4c2d2d8e8b456" // Example user ID
	role := "admin"                      // Example role
	expiration := time.Minute * 2        // Token expiration time

	tokens, err := config.GenerateJWT(userId, role, expiration)
	if err != nil {
		fmt.Printf("Error generating token: %v\n", err)
		return
	}

	// fmt.Printf("Generated Token: %s\n", tokens)

	// Validate the generated JWT
	validatedToken, err := config.ValidateToken(tokens)
	if err != nil {
		fmt.Printf("Error validating token: %v\n", err)
		return
	}

	// fmt.Println("Validated token", validatedToken)

	claims, ok := validatedToken.Claims.(interface{})
	if !ok {
		fmt.Println("Error extracting claims from token")
		return
	}

	val, _ := json.Marshal(claims)

	strval := string(val)

	var jwtClaims models.JWTClaims
	err = json.Unmarshal([]byte(strval), &jwtClaims)
	if err != nil {
		fmt.Printf("Error unmarshalling claims: %v\n", err)
		return
	}
	exp := int64(jwtClaims.Expiration)

	expTime := time.Unix(exp, 0)

	fmt.Println(expTime)
	fmt.Printf("")

}
