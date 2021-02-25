package api

import (
	"github.com/gin-gonic/gin"
	//"github.com/gin-contrib/static"
	"github.com/thinkerou/favicon"

	//"path/filepath"
	//"net/http"
	//"io/ioutil"
	//"log"
	//"os"

	//"github.com/hcbt/vega/backend"
	"github.com/hcbt/vega/db"
)

//RunServer runs api server, receives videoID as input and returns song genre
func RunServer() {
	server := gin.Default()

	//server.Use(static.Serve("/files", static.LocalFile(tempdir, false)))
	server.Use(favicon.New("../static/favicon.ico"))

	server.GET("/video/:videoid", func(c *gin.Context) {
		//cwd, _ := os.Getwd()

		videoID := c.Params.ByName("videoid")

		//json := "ayy"
		result, err := db.FindEntry(videoID)
		if err == nil {
			c.JSON(200, result)
		} else if err != nil {
			db.AddEntry(videoID)
			
			json, err := db.FindEntry(videoID)

			if err == nil {
				c.JSON(200, json)
			}
		}
	})

	server.Run()
}
