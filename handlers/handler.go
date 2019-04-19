package handlers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/metrue/emc/cache"
	"github.com/metrue/emc/utils"
	gocache "github.com/patrickmn/go-cache"
)

func Fibonacci(c *gin.Context) {
	num := c.Query("number")
	if num == "" {
		c.JSON(400, gin.H{
			"message": "number is required",
		})
		return
	}
	n, err := strconv.ParseInt(num, 10, 64)
	fmt.Println(err)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "number is incorrect",
		})
		return
	}

	kache := cache.GetCacheClient()
	val, ok := kache.Get(num)
	if !ok {
		val = utils.FibNumbers(int(n))
		kache.Set(num, val, gocache.NoExpiration)
	} else {
		fmt.Println("read from cache")
	}
	c.JSON(200, gin.H{
		"message": "ok",
		"data":    val,
	})
}
