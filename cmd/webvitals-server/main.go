package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./views/*")
	r.GET("/default", func(c *gin.Context) {
		ttfb, _ := time.ParseDuration(c.DefaultQuery("ttfb", "0s"))
		fcp, _ := time.ParseDuration(c.DefaultQuery("fcp", "1.0s"))
		dom, _ := time.ParseDuration(c.DefaultQuery("dom", "1.5s"))
		lcp, _ := time.ParseDuration(c.DefaultQuery("lcp", "2.0s"))
		cls, _ := time.ParseDuration(c.DefaultQuery("cls", "2.5s"))

		timestamp := c.DefaultQuery("timestamp", strconv.FormatInt(time.Now().UTC().UnixNano(), 10))
		seq, _ := strconv.Atoi(c.DefaultQuery("seq", "1"))
		nextSeq := seq + 1
		fcpMillis := int32(fcp) / 1000000

		time.Sleep(ttfb)
		c.HTML(http.StatusOK, "default.tmpl", gin.H{
			"ttfb":      ttfb,
			"fcp":       fcp,
			"dom":       dom,
			"lcp":       lcp,
			"cls":       cls,
			"timestamp": timestamp,
			"seq":       nextSeq,
			"fcpMillis": fcpMillis,
		})
	})

	r.GET("/assets/:filename", func(c *gin.Context) {
		filename, _ := c.Params.Get("filename")
		ttfb, _ := time.ParseDuration(c.DefaultQuery("ttfb", "0"))
		time.Sleep(ttfb)

		c.FileFromFS(filename, gin.Dir("./assets", false))
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	fmt.Println("listening at :8080")
	r.Run(":8080")
}
