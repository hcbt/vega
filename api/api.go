package api

import (
	"github.com/gin-gonic/gin"
	//"github.com/gin-contrib/static"
	"github.com/thinkerou/favicon"

	//"net/http"
	//"fmt"
	//"os"

	"github.com/hcbt/vega/backend"
)

//RunServer runs api server, receives videoID as input and returns song genre
func RunServer() {
	server := gin.Default()

	//server.Use(static.Serve("/files", static.LocalFile("../cmd/", false)))
	server.Use(favicon.New("../static/favicon.ico"))

	server.GET("/video/:videoid", func(c *gin.Context) {
		videoID := c.Params.ByName("videoid")

		backend.Process(videoID)

		//c.Redirect(http.StatusFound, "/files/" + videoID + ".png")
	})

	server.Run()
}
