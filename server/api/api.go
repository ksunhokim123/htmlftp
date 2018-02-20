package api

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	mouse "github.com/sunho/mouse-hosting/server/mouse"
)

var Service *mouse.Service
var Routes *gin.Engine

func Start(allowedDomains []string, service *mouse.Service) {
	Service = service
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	api := r.Group("/api")
	{
		api.POST("/users", addUserEndpoint)
		api.GET("/keys/:key", getKeyEndpoint)
		admin := api.Group("/admin", gin.BasicAuth(gin.Accounts{
			Service.Config.Username: Service.Config.Password,
		}))
		{
			admin.POST("/keygen", keyGenEndpoint)
			admin.GET("/keys", retriveKeysEndpoint)
		}
	}
	//TODO seperate this
	r.Use(static.Serve("/", static.LocalFile("home", true)))
	go r.Run(Service.Config.Address.String())
}

func getKeyEndpoint(ctx *gin.Context) {
	key := ctx.Param("key")
	if Service.KeyContainer.Exist(key) == -1 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "no such key"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"error": "success"})
	}
}

type addUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Key      string `json:"key"`
}

func addUserEndpoint(ctx *gin.Context) {
	var request addUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else if keyindex := Service.KeyContainer.Exist(request.Key); keyindex == -1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no such key"})
	} else if err = Service.UserContainer.AddUser(request.Username, request.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		Service.KeyContainer.Remove(keyindex)
		ctx.JSON(http.StatusOK, gin.H{"error": "success"})
	}
}

func keyGenEndpoint(ctx *gin.Context) {
	key := Service.KeyContainer.Generate()
	ctx.JSON(http.StatusCreated, gin.H{"error": "success", "key": key})
}

func retriveKeysEndpoint(ctx *gin.Context) {
	keys := []string{}
	for _, key := range *Service.KeyContainer {
		keys = append(keys, key)
	}
	ctx.JSON(http.StatusOK, gin.H{"error": "success", "keys": keys})
}
