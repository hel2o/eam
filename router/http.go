package router

import (
	"crypto/tls"
	"io"
	"net/http"
	"os"

	"github.com/op/go-logging"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var log = logging.MustGetLogger("router")

func Start() {
	ginLogFile, _ := os.OpenFile("var/gin.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	gin.DefaultWriter = io.MultiWriter(ginLogFile)
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.LoadHTMLGlob("templates/*")
	router.Use(static.Serve("/", static.LocalFile("./public", false)))

	router.GET("/", indexPage)
	router.GET("/print", printPage)
	router.GET("/look", lookPage)
	router.POST("/getEAM", getEAMRouter)
	router.POST("/add", addEAMRouter)
	router.POST("/modify", modifyEAMRouter)
	router.POST("/del", delEAMRouter)

	var addr = ":9443"
	s := &http.Server{
		Addr:           addr,
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}
	err := s.ListenAndServeTLS("public/ssl/cert.pem", "public/ssl/key.pem")

	if err != nil {
		log.Fatal(err)
	}
	log.Info("Listening and serving HTTPS on", addr)
}
