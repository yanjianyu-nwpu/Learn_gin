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

## 8 自定义中间件
