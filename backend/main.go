package main

import (
	"net/http"
	"time"
    "log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

// ================== Rest API ==================
	router.GET("/fuga", func(c *gin.Context){
		c.JSON(200, gin.H{
			"massage": "Hello fuga",
		})
	})

	router.GET("/hoge", func(c *gin.Context){
		c.JSON(200, gin.H{
			"massage": "Hello hoge",
		})
	})


// ================== Web APP ==================
    // 自動的にファイルを返す
    router.StaticFS("/static", http.Dir("static"))

    // ルートならリダイレクト
    router.GET("/", func(ctx *gin.Context){
        ctx.Redirect(302, "/static/index.html")
    })

    // フォームの内容を受け取って返す
    router.GET("/hello", func(ctx *gin.Context){
        name := ctx.Query("name")
        ctx.Header("Content-Type", "text/html; charset=UTF-8")
        ctx.String(200, "<h1>Hello, " + name + "</h1>")
    })

	err := router.Run(":3000" )
    if err != nil {
        log.Fatal("サーバ起動に失敗", err)
    }
}

func hello() string {
	return "Hello Golang"
}
