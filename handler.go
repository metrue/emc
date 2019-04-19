package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	gocache "github.com/patrickmn/go-cache"
)

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fibNumbers(n int) []int {
	arr := []int{}
	for i := 0; i < n; i++ {
		arr = append(arr, fib(i))
	}
	return arr
}

func fibonacci(c *gin.Context) {
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

	cache := GetCacheClient()
	val, ok := cache.Get(num)
	if !ok {
		val = fibNumbers(int(n))
		cache.Set(num, val, gocache.NoExpiration)
	} else {
		fmt.Println("read from cache")
	}
	c.JSON(200, gin.H{
		"message": "ok",
		"data":    val,
	})
}
