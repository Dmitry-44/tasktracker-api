package router

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"tasktracker-api/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type Router struct {
	services *service.Service
}

func NewRouter(services *service.Service) *Router {
	return &Router{services: services}
}

func (r *Router) InitRoutes() *gin.Engine {

	router := gin.New()
	router.POST("/login", r.Login)
	router.POST("/logup", r.Logup)
	api := router.Group("/api")
	api.Use(AuthMiddleware(r))
	{
		v1 := api.Group("/v1")
		{
			tasks := v1.Group(("/tasks"))
			{
				tasks.GET("/", r.GetAllTasks)
				tasks.GET("/:id", r.GetTaskById)
				tasks.POST("/", r.CreateTask)
				tasks.PUT("/:id", r.UpdateTask)
				tasks.DELETE("/:id", r.DeleteTask)

			}
			// groups := v1.Group("/groups")
			{
				// groups.GET("/", r.GetAllGroupes)
				// groups.GET("/:id", r.GetGroupById)
				// groups.POST("/", r.CreateGroup)
				// groups.PUT("/:id", r.UpdateGroup)
				// groups.DELETE("/:id", r.DeleteGroup)
			}
		}
	}
	return router
}

type userCtx string

const ctxKeyUser userCtx = "user"

func AuthMiddleware(r *Router) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := extractTokenFromHeader(ctx)
		if len(token) == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, ok := GetClaimsFromToken(token)
		if ok != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		userIDString := claims["sub"].(string)
		userId, err := strconv.Atoi(userIDString)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		user, err := r.services.Auth.GetUserById(userId)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), ctxKeyUser, user))
		ctx.Next()
	}
}

func extractTokenFromHeader(c *gin.Context) string {
	bearToken := c.GetHeader("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func GetClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("jwtSignedKey")), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
