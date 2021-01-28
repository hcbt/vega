package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
	"github.com/thinkerou/favicon"

	//"path/filepath"
	"net/http"
	"io/ioutil"
	"log"
	"os"

	"github.com/hcbt/vega/backend"
)

//RunServer runs api server, receives videoID as input and returns song genre
func RunServer() {
	server := gin.Default()

	//create tempdir here
	tempdir, err := ioutil.TempDir("", "vega")
	if err != nil {
		log.Fatal(err)
	}

	server.Use(static.Serve("/files", static.LocalFile(tempdir, false)))
	server.Use(favicon.New("../static/favicon.ico"))

	server.GET("/video/:videoid", func(c *gin.Context) {
		cwd, _ := os.Getwd()
	
		videoID := c.Params.ByName("videoid")

		backend.Process(tempdir, videoID)

		c.Redirect(http.StatusFound, "/files/" + videoID + ".png")

		os.Chdir(cwd)
	})

	server.Run()
}
