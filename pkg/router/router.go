package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"tasktracker-api/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Router struct {
	services *service.Service
}

func NewRouter(services *service.Service) *Router {
	return &Router{services: services}
}

func (r *Router) InitRoutes() *gin.Engine {

	router := gin.New()
	router.Use(CORSMiddleware())
	router.POST("/login", r.Login)
	router.POST("/logup", r.Logup)
	router.POST("/auth", r.Auth)
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
				tasks.GET("/ws", func(c *gin.Context) {
					r.WSHandler(c.Writer, c.Request)
				})
			}
			groups := v1.Group("/groups")
			{
				groups.GET("/", r.GetAllGroupes)
				groups.GET("/:id", r.GetGroupById)
				groups.GET("/:id/tasks", r.GetTasksByGroupId)
				groups.POST("/", r.CreateGroup)
				// groups.PUT("/:id", r.UpdateGroup)
				groups.DELETE("/:id", r.DeleteGroup)
			}
		}
	}
	return router
}

type userCtx string

const ctxKeyUser userCtx = "user"

func AuthMiddleware(r *Router) gin.HandlerFunc {
	fmt.Print("ssssss")
	return func(ctx *gin.Context) {
		token := extractTokenFromHeader(ctx)
		if len(token) == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		fmt.Printf("token is %v", token)
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
		return []byte("secret"), nil
	})
	if err != nil {
		log.Printf("get claims error: %v", err)
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "http://localhost:3001")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Access-Control-Allow-Origin")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT,DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
