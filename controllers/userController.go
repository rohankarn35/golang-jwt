package controllers

import (
	"auth/initializers"
	"auth/models"
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var user models.User
	if c.Bind(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the body",
		})
		return
	}
	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No body",
		})
		return

	}
	if len(user.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password must be atleast 8 characters",
		})
		return

	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed hashing password",
		})
		return
	}
	user.Password = string(hashPassword)
	connection := initializers.Client.Database("Users").Collection("auth")
	filler := bson.M{"email": user.Email}

	var existingUser models.User

	err = connection.FindOne(context.TODO(), filler).Decode(&existingUser)
	if err == mongo.ErrNoDocuments {
		_, err = connection.InsertOne(context.TODO(), user)
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to save user to database",
			})
			return

		}

	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot connect to the host",
		})
		return
	} else {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Email Already Exists",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})

}

func Login(c *gin.Context) {
	var login models.Login // Use a separate struct for login credentials
	if c.BindJSON(&login) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the body",
		})
		return
	}

	// Find the user by email
	filter := bson.M{"email": login.Email}
	var existingUser models.User
	err := initializers.Client.Database("Users").Collection("auth").FindOne(context.TODO(), filter).Decode(&existingUser)

	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No email",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to check credentials",
		})
		return
	}

	// Compare hashed passwords (corrected issue)
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(login.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Generate and return JWT token
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": existingUser.Email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := tokens.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenString,
	})
}
