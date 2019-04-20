package handlers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/metrue/emc/cache"
	"github.com/metrue/emc/utils"
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
	if err != nil {
		msg := fmt.Sprintf("number (%s) is incorrect: %v", num, err)
		c.JSON(400, gin.H{
			"message": msg,
		})
		return
	}

	res, ok := cache.Get(num)
	if !ok {
		res = utils.FibNumbers(int64(n))
		cache.Set(num, res)
	}
	c.JSON(200, gin.H{
		"message": "ok",
		"data":    res,
	})
}
