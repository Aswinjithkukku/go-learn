package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/aswinjithkukku/jwt-auth/initializer"
	"github.com/aswinjithkukku/jwt-auth/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	var body struct {
		Email    string
		Name     string
		Password string
	}

	err := c.Bind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var existingUser models.User

	// check email already exist
	initializer.DB.First(&existingUser, "email = ? ", body.Email)

	if existingUser.Email == body.Email {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User found with same Email",
		})
		return
	}

	// bcrypt the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user := models.User{Email: body.Email, Name: body.Name, Password: string(hash)}

	result := initializer.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"user": gin.H{
			"name":      user.Name,
			"email":     user.Email,
			"createdAt": user.CreatedAt,
		},
	})
}

func SignIn(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	err := c.Bind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// find existing user
	var existingUser models.User

	initializer.DB.First(&existingUser, "email = ?", body.Email)

	if existingUser.Email != body.Email {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid User",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": existingUser.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"user": gin.H{
			"name":  existingUser.Name,
			"email": existingUser.Email,
			"token": tokenString,
		},
	})

}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	var userResponse = struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		ID:    user.(models.User).ID,
		Name:  user.(models.User).Name,
		Email: user.(models.User).Email,
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"user":    userResponse,
	})
}
