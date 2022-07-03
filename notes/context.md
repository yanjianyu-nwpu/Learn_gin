# Context

## 0 inclusion

context is import partition in golang project

context 是重要概念，中间件中间传递数据

         // Context is the most important part of gin. It allows us to pass variables between middleware,
    // manage the flow, validate the JSON of a request and render a JSON response for example.
    type Context struct {
        // http responsewRITER
    
        writermem responseWriter
        Request   *http.Request
        Writer    ResponseWriter
    
        Params   Params
        handlers HandlersChain
        index    int8
        fullPath string
    
        engine       *Engine
        params       *Params
        skippedNodes *[]skippedNode
    
        // This mutex protects Keys map.
        mu sync.RWMutex
    
        // Keys is a key/value pair exclusively for the context of each request.
        Keys map[string]any
    
        // Errors is a list of errors attached to all the handlers/middlewares who used this context.
        Errors errorMsgs
    
        // Accepted defines a list of manually accepted formats for content negotiation.
        Accepted []string
    
        // queryCache caches the query result from c.Request.URL.Query().
        queryCache url.Values
    
        // formCache caches c.Request.PostForm, which contains the parsed form data from POST, PATCH,
        // or PUT body parameters.
        formCache url.Values
    
        // SameSite allows a server to define a cookie attribute making it impossible for
        // the browser to send this cookie along with cross-site requests.
        sameSite http.SameSite
    }                                                                                                                        

## 详解

这里Default 函数,这里router 记录所有的routergroup 整个router group 文件

```
type RouterGroup struct {
    Handlers HandlersChain
    basePath string
    engine   *Engine
    root     bool
}
```

```

```

这里是 use 中间件是再router group 上添加的

这里handlers 就是整个group的handle 

## 整体研究一下所有的功能就好了

### 1.0  获取路径中的参数

 这里有获取 真实的 name 这里接如                                                                       

```
// Param returns the value of the URL param.
// It is a shortcut for c.Params.ByName(key)
//     router.GET("/user/:id", func(c *gin.Context) {
//         // a GET request to /user/john
//         id := c.Param("id") // id == "john"
//     })
func (c *Context) Param(key string) string {
    return c.Params.ByName(key)
}
```

这里获得 key value 的结果

在context 中会有专门的字段params  param 是kv

是一个list 会for 循环get的

gin 路径中的参数

```
/user/:name
这种只能匹配 
```

这种结果/user/bill 能 但 ./USER/bill/ /user/bill/不 能匹配

但是/user/*t

可以匹配 上面

### DefaultQuery

GetQuery 用于获取参数

比如这种 /welcome 可以拿到 

GetQuery 这种

```
/welcome?firstname=Jane&lastname=Doe
```

这样的结果

这里Post 有PotForm 结果

## 路由分组

然后为了方便管理

```

```
