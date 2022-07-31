package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/saisesai/www.sxu.ink/config"
	"github.com/saisesai/www.sxu.ink/controller"
	"github.com/saisesai/www.sxu.ink/middleware/auth"
	gin_logrus "github.com/saisesai/www.sxu.ink/middleware/gin-logrus"
	"github.com/sirupsen/logrus"
)

func init() {
	flag.Parse()
	if !*config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	var err error
	router := gin.New()

	err = router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		logrus.WithError(err).Fatalln("failed to set trusted proxy!")
	}

	router.Use(gin_logrus.Logger(logrus.StandardLogger()), gin.Recovery())

	router.POST("/api/user/register", controller.UserRegisterHandler)
	router.POST("/api/user/login", controller.UserLoginHandler)

	var jwtAuth *gin.RouterGroup
	if *config.Debug {
		jwtAuth = router.Group("/")
	} else {
		jwtAuth = router.Group("/", auth.JWTAuth())
	}
	{
		jwtAuth.POST("/api/hello", controller.Hello)

		jwtAuth.POST("/api/user/password/set", controller.UserSetPassword)
		jwtAuth.POST("/api/user/nickname/set", controller.UserSetNickName)
		jwtAuth.POST("/api/user/group/set", controller.UserSetGroup)

		jwtAuth.POST("/api/artwork/new", controller.ArtworkAddHandler)
		jwtAuth.POST("/api/artwork/search", controller.ArtworkSearchHandler)
		jwtAuth.POST("/api/artwork/del/:uuid", controller.ArtworkDeleteHandler)
		jwtAuth.POST("/api/artwork/get/:uuid", controller.ArtworkGetHandler)
		jwtAuth.POST("/api/artwork/set/:uuid", controller.ArtworkSetHandler)

		jwtAuth.POST("/api/image/new", controller.ImageAddHandler)
		jwtAuth.POST("/api/image/del/:uuid", controller.ImageDeleteHandler)
		jwtAuth.POST("/api/image/get/:uuid", controller.ImageGetHandler)
		jwtAuth.POST("/api/image/set/:uuid", controller.ImageSetHandler)

		jwtAuth.POST("/api/char/new", controller.CharAddHandler)
		jwtAuth.POST("/api/char/search", controller.CharSearchHandler)
		jwtAuth.POST("/api/char/del/:uuid", controller.CharDeleteHandler)
		jwtAuth.POST("/api/char/get/:uuid", controller.CharGetHandler)
		jwtAuth.POST("/api/char/set/:uuid", controller.CharSetHandler)

	}
	err = router.Run(*config.Listen)
	if err != nil {
		logrus.WithError(err).Fatalln("service failed to startup!")
	}
}
