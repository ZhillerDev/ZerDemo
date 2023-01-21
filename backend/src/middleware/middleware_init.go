package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func m1(c *gin.Context) {
	c.Set("name", "tom")
}
func m2(c *gin.Context) {
	name, _ := c.Get("name")
	fmt.Println(name)
}

func m3(c *gin.Context) {
	// 携程内必须使用复制过的上下文
	ccopy := c.Copy()
	go func() {
		time.Sleep(time.Millisecond * 500)
		fmt.Println(ccopy.Request.RequestURI)
	}()
}
