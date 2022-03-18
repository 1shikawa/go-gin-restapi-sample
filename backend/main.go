package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
)

func main() {
	router := gin.Default()

	 // ここからCorsの設定
    router.Use(cors.New(cors.Config{
    // アクセスを許可したいアクセス元
    AllowOrigins: []string{
        "http://127.0.0.1",
    },
    // アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
    AllowMethods: []string{
        "POST",
        "GET",
        "OPTIONS",
    },
    // 許可したいHTTPリクエストヘッダ
    AllowHeaders: []string{
        "Access-Control-Allow-Credentials",
        "Access-Control-Allow-Headers",
        "Content-Type",
        "Content-Length",
        "Accept-Encoding",
        "Authorization",
    },
    // cookieなどの情報を必要とするかどうか
    AllowCredentials: false,
    // preflightリクエストの結果をキャッシュする時間
    MaxAge: 24 * time.Hour,
}))

	router.GET("/hello", func(c *gin.Context){
		c.JSON(200, gin.H{
			"massage": "Hello World",
		})
	})

	router.Run(":3000" )
}

func hello() string {
	return "Hello Golang"
}
