# 新笔记
## 1 html渲染
```
func main() {
    router := gin.Default()
    //加载模板
    router.LoadHTMLGlob("templates/*")
    //router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
    //定义路由
    router.GET("/index", func(c *gin.Context) {
        //根据完整文件名渲染模板，并传递参数
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title": "Main website",
        })
    })
    router.Run(":8080")
}
```
HTMl
```
<html>
    <h1>
        {{ .title }}
    </h1>
</html>

```
总共分两步
- LoadHTMLGLob 解析导入模板
- 就是在gin.H{"title":"FEWFW"}加上参数，html中用{{.title}}替代；返回时候用c.HTML
  
# 1 静态路径
```
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    router := gin.Default()
    // 下面测试静态文件服务
    // 显示当前文件夹下的所有文件/或者指定文件
    router.StaticFS("/showDir", http.Dir("."))
    router.StaticFS("/files", http.Dir("/bin"))
    //Static提供给定文件系统根目录中的文件。
    //router.Static("/files", "/bin")
    router.StaticFile("/image", "./assets/miao.jpg")

    router.Run(":8080")

```
静态文件文件还是需要的，就是要记录路径，前面是url的路径，后面是文件路径

# 3 中间件
- 全局中间件可以USe 注册，在请求前执行c.Next之前的部分
- 中间件用作打日志，做接口鉴权，做校验是否登录

# 4 Session
用github.com/gin-contrib/sessions包
```
package main

import (
        // 导入session包
	"github.com/gin-contrib/sessions"
       // 导入session存储引擎
	"github.com/gin-contrib/sessions/cookie"
        // 导入gin框架包
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

        // 创建基于cookie的存储引擎，shuiche 参数是用于加密的密钥，可以随便填写
	store := cookie.NewStore([]byte("shuiche"))

        // 设置session中间件，参数mysession，指的是session的名字，也是cookie的名字
       // store是前面创建的存储引擎
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/test", func(c *gin.Context) {
                // 初始化session对象
		session := sessions.Default(c)
                
                // 通过session.Get读取session值
                // session是键值对格式数据，因此需要通过key查询数据

		if session.Get("hello") != "world" {
                        // 设置session数据,()
			session.Set("hello", "world")
                        // 删除session数据
                        session.Delete("tizi365")
                        // 保存session数据
			session.Save()
                        // 删除整个session
                        // session.Clear()
		}
                
		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})
	r.Run(":8000")
}
```
这里store中放到cookie中的store 然后填入一个密钥，啥都可以
然后可以记录到session里面