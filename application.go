package main

import (
    "os"
    "log"
	"time"
	"strings"
    "github.com/gin-gonic/gin"
    "github.com/appleboy/gin-jwt"
	"github.com/Htech/Htech-api/db"
	"github.com/Htech/Htech-api/models"
	"github.com/Htech/Htech-api/utils"
	"github.com/Htech/Htech-api/helpers"
	"github.com/Htech/Htech-api/httpmodels"
	"github.com/Htech/Htech-api/controllers"
)

var identityKey = "email"

func main() {
	router := gin.Default()
	db.InitDB()
    Migrate()
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.Users); ok {
				return jwt.MapClaims{
					identityKey:    v.Email,
					"id":    v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			log.Printf("claims = %+v\n", claims)
			return &models.Users{
				Email: claims["email"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			loginVals := &httpmodels.Login{}
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			user := models.Users{}
			db.DB.Where("email = ? AND status = ? ", loginVals.Email, helpers.STATUS_ENABLED).First(&user)
			log.Printf("usuario = %+v\n", &user)

			err := utils.ComparePassword(user.Password, loginVals.Password)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			return &user, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
            path := c.Request.URL.Path
            for _, param := range c.Params {
                path = strings.Replace(path, param.Value, ":"+param.Key, -1)
            }
		    //log.Printf("Path: ", path)
			// TODO Aqui hay que agregar validaciones, tal vez roles?
			if v, ok := data.(*models.Users); ok && v.Email != "" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header: Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	router.POST("/login/", authMiddleware.LoginHandler)
	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"status": false, "message": "Page not found"})
	})
	auth := router.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)

	auth.Use(authMiddleware.MiddlewareFunc())
	{
	    auth.POST("/clients/getClients", controllers.GetClients)
	    auth.POST("/clients/createClient", controllers.CreateClient)
	    auth.POST("/clients/updateClient", controllers.UpdateClient)
	    auth.POST("/cards/getCards", controllers.GetCards)
	    auth.POST("/cards/createCards", controllers.CreateCards)
	    auth.POST("/cards/updateCards", controllers.UpdateCards)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	erro := router.Run(":" + port)
	log.Panic(erro)
}

func Migrate(){
    db.DB.AutoMigrate(&models.TypeDocuments{})
    db.DB.AutoMigrate(&models.Users{})
    db.DB.AutoMigrate(&models.Clients{})
    db.DB.AutoMigrate(&models.Cards{})
    db.DB.AutoMigrate(&models.Users{}).AddForeignKey("type_documents_id", "type_documents(id)", "RESTRICT", "RESTRICT")
    db.DB.AutoMigrate(&models.Clients{}).AddForeignKey("type_documents_id", "type_documents(id)", "RESTRICT", "RESTRICT").AddForeignKey("users_id", "users(id)", "RESTRICT", "RESTRICT")
    db.DB.AutoMigrate(&models.Cards{}).AddForeignKey("clients_id", "users(id)", "RESTRICT", "RESTRICT")
}
