package main
import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {

	r :=gin.Default()

	// 协程 测试
	r.GET("/long_async", func(c *gin.Context) {

		// 创建在协程中使用的副本
		//cp := c.Copy()

		go func() {
			time.Sleep(5*time.Second)
			log.Println("Done ! in path" + c.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5*time.Second)
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	r.Run(":9999")
}