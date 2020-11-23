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
		ttfb, _ := time.ParseDuration(c.DefaultQuery("ttfb", "0.5s"))
		fcp, _ := time.ParseDuration(c.DefaultQuery("fcp", "1.0s"))
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
			"lcp":       lcp,
			"cls":       cls,
			"timestamp": timestamp,
			"seq":       nextSeq,
			"fcpMillis": fcpMillis,
		})
	})

	r.GET("/images/:filename", func(c *gin.Context) {
		filename, _ := c.Params.Get("filename")
		ttfb, _ := time.ParseDuration(c.DefaultQuery("ttfb", "0"))
		time.Sleep(ttfb)

		c.FileFromFS(filename, gin.Dir("./static", false))
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "webvitals-server")
	})

	fmt.Println("listening at :8080")
	r.Run(":8080")
}
