package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type User struct {
	Username string `json:"username"`
}

func main() {
	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		startTime := time.Now()

		c.JSON(http.StatusOK, gin.H{
			"method":         http.MethodGet,
			"elapsedTime/ms": time.Since(startTime).Milliseconds(),
		})
	})

	engine.POST("/", func(c *gin.Context) {
		startTime := time.Now()

		var u User
		err := c.BindJSON(&u)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"method":         http.MethodPost,
			"elapsedTime/ms": time.Since(startTime).Milliseconds(),
			"username":       u.Username,
		})
	})

	//
	//// 增加一个redis访问
	//rdb := redis.NewClient(&redis.Options{
	//	Addr:     "redis:6379",
	//	Password: "", // no password set
	//	DB:       0,  // use default DB
	//})
	//ctx:= context.Background()
	//pong, err := rdb.Ping(ctx).Result()
	//fmt.Println(pong, err)
	//if err!=nil{
	//	os.Exit(-1)
	//}
	//_,err = rdb.Get(ctx,"key").Result()
	//// 不存在key则创建
	//if err==redis.Nil{
	//	_,err = rdb.Set(ctx,"key","1",0).Result()
	//	if err!=nil{
	//		fmt.Println(err)
	//		os.Exit(-1)
	//	}
	//}
	//
	//engine.GET("/redis", func(c *gin.Context) {
	//	startTime := time.Now()
	//	val ,err:= rdb.Get(ctx,"key").Result()
	//	if err!=nil{
	//		c.JSON(http.StatusOK, gin.H{
	//			"method":         http.MethodGet,
	//			"elapsedTime/ms": time.Since(startTime).Milliseconds(),
	//			"err":"Redis Get Key Error",
	//		})
	//		return
	//	}
	//	current,_ := strconv.Atoi(val)
	//	_,err = rdb.Set(ctx,"key",current+1,0).Result()
	//	if err!=nil{
	//		c.JSON(http.StatusOK, gin.H{
	//			"method":         http.MethodGet,
	//			"elapsedTime/ms": time.Since(startTime).Milliseconds(),
	//			"err":"Redis Set Key Error",
	//		})
	//		return
	//	}
	//	c.JSON(http.StatusOK, gin.H{
	//		"method":         http.MethodGet,
	//		"elapsedTime/ms": time.Since(startTime).Milliseconds(),
	//		"redisKey":current+1,
	//	})
	//})

    // 增加打印hostname功能
	engine.GET("/name", func(c *gin.Context) {
		startTime := time.Now()

		name,_ := os.Hostname()
		c.JSON(http.StatusOK, gin.H{
			"method":         http.MethodGet,
			"elapsedTime/ms": time.Since(startTime).Milliseconds(),
			"hostname": name,
		})
	})

	engine.Run(":8081")
}