package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/saisesai/www.sxu.ink/config"
	"github.com/saisesai/www.sxu.ink/controller"
	"github.com/saisesai/www.sxu.ink/middleware/gin-logrus"
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

	router.POST("/api/artwork/new", controller.ArtworkAddHandler)
	router.POST("/api/artwork/del/:uuid", controller.ArtworkDeleteHandler)
	router.POST("/api/artwork/get/:uuid", controller.ArtworkGetHandler)
	router.POST("/api/artwork/set/:uuid", controller.ArtworkSetHandler)

	router.POST("/api/image/new", controller.ImageAddHandler)
	router.POST("/api/image/del/:uuid", controller.ImageDeleteHandler)
	router.POST("/api/image/get/:uuid", controller.ImageGetHandler)
	router.POST("/api/image/set/:uuid", controller.ImageSetHandler)

	router.POST("/api/char/new", controller.CharAddHandler)
	router.POST("/api/char/del/:uuid", controller.CharDeleteHandler)
	router.POST("/api/char/get/:uuid", controller.CharGetHandler)
	router.POST("/api/char/set/:uuid", controller.CharSetHandler)

	err = router.Run(*config.Listen)
	if err != nil {
		logrus.WithError(err).Fatalln("service failed to startup!")
	}
}
