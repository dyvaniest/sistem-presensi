package api

import (
	"net/http"
	"sistem-presensi/models"
	"sistem-presensi/services"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserAPI interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type userAPI struct {
	userService services.UserService
}

func NewUserAPI(userService services.UserService) *userAPI {
	return &userAPI{userService}
}

func (u *userAPI) Register(c *gin.Context) {
	var user models.Dosen

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse("invalid decode json"))
		return
	}

	if user.User.Username == "" || user.User.Password == "" || user.User.Email == "" {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse("register data is empty"))
		return
	}

	var recordUser = models.User{
		Username: user.User.Username,
		Email:    user.User.Email,
		Password: user.User.Password,
	}

	recordUser, err := u.userService.Register(&recordUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusCreated, models.NewSuccessResponse("register success"))
}

func (u *userAPI) Login(c *gin.Context) {
	var loginReq models.UserLogin
	if err := c.BindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse("invalid decode json"))
		return
	}

	if loginReq.Username == "" || loginReq.Password == "" {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse("email or password is empty"))
		return
	}

	user := &models.User{
		Username: loginReq.Username,
		Password: loginReq.Password,
	}

	tokenStr, err := u.userService.Login(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error internal server"})
		return
	}

	expirationTime := time.Now().Add(10 * time.Minute)
	token, err := jwt.ParseWithClaims(*tokenStr, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return models.JwtKey, nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error internal server"})
		return
	}

	claims, ok := token.Claims.(*models.Claims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, models.NewErrorResponse("invalid token"))
		return
	}
	claims.ExpiresAt = expirationTime.Unix()
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "session_token",
		Value:   *tokenStr,
		Expires: expirationTime,
	})
	c.JSON(http.StatusOK, models.NewSuccessResponse("login success"))
}

func (u *userAPI) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")

	user, err := u.userService.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
