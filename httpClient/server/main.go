package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(350)))
		c.String(http.StatusOK, strconv.Itoa(rand.Int()))
	})

	fmt.Println("start")
	r.Run(":8888")
}
