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

这里有获取
