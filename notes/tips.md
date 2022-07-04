# tips

## 0 写日志文件

```
f, _ := os.Create("gin.log")
gin.DefaultWriter = io.MultiWriter(f)
```

放到文件里面

## 1 绑定HTML的复选框

这里哦

```
type myForm struct {
    Colors []string `form:"colors[]"`
}

...

func formHandler(c *gin.Context) {
    var fakeForm myForm
    c.ShouldBind(&fakeForm)
    c.JSON(200, gin.H{"color": fakeForm.Colors})
}
```

## 2 静态文件

```

```

## 3 第三方文件

## 4 ginH

其实就是 H map[string]interface{}

gin.H{"test": gin.H{"data": ""}}

## 5 第三方数据

```
func main() {
    router := gin.Default()
    router.GET("/someDataFromReader", func(c *gin.Context) {
        response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
        if err != nil || response.StatusCode != http.StatusOK {
            c.Status(http.StatusServiceUnavailable)
            return
        }

        reader := response.Body
        contentLength := response.ContentLength
        contentType := response.Header.Get("Content-Type")

        extraHeaders := map[string]string{
            "Content-Disposition": `attachment; filename="gopher.png"`,
        }

        c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
    })
    router.Run(":8080")
}
```

这里清理

## 6 HTML 渲然

返回指定的html

```
func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*")
    //router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
    router.GET("/index", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title": "Main website",
        })
    })
    router.Run(":8080")
}
```

分为氛围

```
<html>
    <h1>
        {{ .title }}
    </h1>
</html>
```

html 模板

## 7 重新定向

很简单就是 很简单的 

```
c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
```

## 8 # BasicAuth()（验证）中间件

这里没搞明白

## 9 HTTP 协议参数设置

可以内嵌在net/http的服务上干

```
func main() {
    router := gin.Default()

    s := &http.Server{
        Addr:           ":8080",
        Handler:        router,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe()
}
```

## 10 HTTPS 证书设置

这里 Let's encrypt 证书

```
package main

import (
    "log"

    "github.com/gin-gonic/autotls"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/acme/autocert"
)

func main() {
    r := gin.Default()

    // Ping handler
    r.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

    m := autocert.Manager{
        Prompt:     autocert.AcceptTOS,
        HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
        Cache:      autocert.DirCache("/var/www/.cache"),
    }

    log.Fatal(autotls.RunWithManager(r, &m))
}
```

## 11 优雅启停

要想优雅的重启或者停止 web 服务器 

可以使用 [fvbock/endless](https://github.com/fvbock/endless) 替代listenAndServer

这里就是 服务器是一个g 来跑，主 协程用来等信号

```
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
```



## 12 二进制文件

## 13 绑定自定义结构体

```
type StructA struct {
    FieldA string `form:"field_a"`
}

type StructB struct {
    NestedStruct StructA
    FieldB string `form:"field_b"`
}

type StructC struct {
    NestedStructPointer *StructA
    FieldC string `form:"field_c"`
}

type StructD struct {
    NestedAnonyStruct struct {
        FieldX string `form:"field_x"`
    }
    FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
    var b StructB
    c.Bind(&b)
    c.JSON(200, gin.H{
        "a": b.NestedStruct,
        "b": b.FieldB,
    })
}

func GetDataC(c *gin.Context) {
    var b StructC
    c.Bind(&b)
    c.JSON(200, gin.H{
        "a": b.NestedStructPointer,
        "c": b.FieldC,
    })
}

func GetDataD(c *gin.Context) {
    var b StructD
    c.Bind(&b)
    c.JSON(200, gin.H{
        "x": b.NestedAnonyStruct,
        "d": b.FieldD,
    })
}

func main() {
    r := gin.Default()
    r.GET("/getb", GetDataB)
    r.GET("/getc", GetDataC)
    r.GET("/getd", GetDataD)

    r.Run()
}
```

这里结果可以用 getc/这样的结果

```
$ curl "http://localhost:8080/getb?field_a=hello&field_b=world"
{"a":{"FieldA":"hello"},"b":"world"}
$ curl "http://localhost:8080/getc?field_a=hello&field_c=world"
{"a":{"FieldA":"hello"},"c":"world"}
$ curl "http://localhost:8080/getd?field_x=hello&field_d=world"
{"d":"world","x":{"FieldX":"hello"}}
```





## 14 设置并获取cookie

比较简单  c。Cokkie（“gin_cookie”）  这里和c.String 是一样的

c.setCookie 设置cookie

## 15 多绑定

还可以 支持很多格式

`JSON`, `XML`, `MsgPack`, `ProtoBuf`。对于其他格式，`Query`, `Form`, `FormPost`, `FormMultipart`, 可以被`c.ShouldBind()`多次调用而不影响性能

## http2推送
