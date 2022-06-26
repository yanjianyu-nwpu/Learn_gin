# Gin

## 0 Instruction

- based on [GitHub - julienschmidt/httprouter: A high performance HTTP request router that scales well](https://github.com/julienschmidt/httprouter)

- examples [GitHub - gin-gonic/examples: A repository to host examples and tutorials for Gin.](https://github.com/gin-gonic/examples)

## 1 gin.Engine

gin-gonic gin.go New

```
// Engine is the framework's instance, it contains the muxer, middleware and configuration settings.
// Create an instance of Engine, by using New() or Default()
// Engine 是框架的实例,包含muxer中间件和默认设置
type Engine struct {
    RouterGroup

    // RedirectTrailingSlash enables automatic redirection if the current route can't be matched but a
    // handler for the path with (without) the trailing slash exists.
    // For example if /foo/ is requested but a route only exists for /foo, the
    // client is redirected to /foo with http status code 301 for GET requests
    // and 307 for all other request methods.
    // trailing 尾随的 slash 斜线
    // RedirectTrailingSlash 开 自动重定向 如果没有匹配的带斜线。
    // 举个例子  如果请求 /foo/ 但是只有 /foo 会自动给301 或者3-07重定向
    RedirectTrailingSlash bool

    // RedirectFixedPath if enabled, the router tries to fix the current request path, if no
    // handle is registered for it.
    // RedirectFixedPath 如果开启 router会尝试修复当前请求path 如果没有handle
    // First superfluous path elements like ../ or // are removed.
    // Afterwards the router does a case-insensitive lookup of the cleaned path.
    // If a handle can be found for this route, the router makes a redirection
    // to the corrected path with status code 301 for GET requests and 307 for
    // all other request methods.
    // 比如 /foo 那 /..//foO不能被重定向
    // For example /FOO and /..//Foo could be redirected to /foo.
    // RedirectTrailingSlash is independent of this option.
    RedirectFixedPath bool

    // HandleMethodNotAllowed if enabled, the router checks if another method is allowed for the
    // current route, if the current request can not be routed.
    // HandleMethodNotAllowed 如果开启，router 检查anther method 是否也符合这个route 如果当前请求
    // If this is the case, the request is answered with 'Method Not Allowed'
    // and HTTP status code 405.
    // If no other Method is allowed, the request is delegated to the NotFound
    // handler.
    HandleMethodNotAllowed bool

    // ForwardedByClientIP if enabled, client IP will be parsed from the request's headers that
    // match those stored at `(*gin.Engine).RemoteIPHeaders`. If no IP was
    // fetched, it falls back to the IP obtained from
    // `(*gin.Context).Request.RemoteAddr`.
    // 如果ForwardedByClient 如果开启那么 客户端IP 挥会被从请求头 解析(*gin.Engine).RemoteIPHeaders 
    // 如果没有IP返回 会返回获得从从 `(*gin.Context).Request.RemoteAddr`.获得的IP
    ForwardedByClientIP bool

    // AppEngine was deprecated.
    // Deprecated: USE `TrustedPlatform` WITH VALUE `gin.PlatformGoogleAppEngine` INSTEAD
    // #726 #755 If enabled, it will trust some headers starting with
    // 'X-AppEngine...' for better integration with that PaaS.
    // AppEngine 被抛弃 
    AppEngine bool

    // UseRawPath if enabled, the url.RawPath will be used to find parameters.
    // UseRawPath 如果开启 url.RawPath会用于找到参数
    UseRawPath bool

    // UnescapePathValues if true, the path value will be unescaped.
    // If UseRawPath is false (by default), the UnescapePathValues effectively is true,
    // as url.Path gonna be used, which is already unescaped.
    // UnesapgePathValues 如果ture 如果userawpath是false unescapePathvalues 效率是ture
    // url Path 会被使用 
    UnescapePathValues bool

    // RemoveExtraSlash a parameter can be parsed from the URL even with extra slashes.
    // See the PR #1817 and issue #1644
    // RemoveExtraSlash 参数能否被解析url参数 和额外的斜线
    RemoveExtraSlash bool

    // RemoteIPHeaders list of headers used to obtain the client IP when
    // `(*gin.Engine).ForwardedByClientIP` is `true` and
    // `(*gin.Context).Request.RemoteAddr` is matched by at least one of the
    // network origins of list defined by `(*gin.Engine).SetTrustedProxies()`.
    // 上面那个参数的配套
    RemoteIPHeaders []string

    // TrustedPlatform if set to a constant of value gin.Platform*, trusts the headers set by
    // that platform, for example to determine the client IP
    // TrustedPlatform 设置
    TrustedPlatform string

    // MaxMultipartMemory value of 'maxMemory' param that is given to http.Request's ParseMultipartForm
    // method call. 
    // MaxMultiPartMemor 最大内存 
    MaxMultipartMemory int64

    // UseH2C enable h2c support.
    // http/2 定义了两个版本的https 一个是基于 tls上的http 2 协议。  
    UseH2C bool

    // ContextWithFallback enable fallback Context.Deadline(), Context.Done(), Context.Err() and Context.Value() when Context.Request.Context() is not nil.
    // ContextWithFallback 开启了 fallback 退路 Context.Deadline()
    ContextWithFallback bool
    // Delims表示用于HTML模板呈现的一组左右分隔符。
    delims           render.Delims
    // 安全的string
    secureJSONPrefix string
    // HTML 渲染器
    HTMLRender       render.HTMLRender
    // map[string]interface{}
    FuncMap          template.FuncMap
    // []FUNC 所有没有路由的
    allNoRoute       HandlersChain
    //  all没有method的
    allNoMethod      HandlersChain
    noRoute          HandlersChain
    noMethod         HandlersChain 
    // sync pool池
    pool             sync.Pool
    // 每个method 对应一个前缀树
    trees            methodTrees
    // 最大的参数
    maxParams        uint16
    maxSections      uint16
    trustedProxies   []string
    trustedCIDRs     []*net.IPNet
}
```

## 1.1 Default 返回一个engin指针

   默认的情况会带上   Logger 和 Recovery 两个中间件

```
func New() *Engine {
    debugPrintWARNINGNew()
    engine := &Engine{
        RouterGroup: RouterGroup{
            Handlers: nil,
            basePath: "/",
            root:     true,
        },
        FuncMap:                template.FuncMap{},
        RedirectTrailingSlash:  true,
        RedirectFixedPath:      false,
        HandleMethodNotAllowed: false,
        ForwardedByClientIP:    true,

        RemoteIPHeaders:        []string{"X-Forwarded-For", "X-Real-IP"},
        TrustedPlatform:        defaultPlatform,
        UseRawPath:             false,
        RemoveExtraSlash:       false,
        UnescapePathValues:     true,
        MaxMultipartMemory:     defaultMultipartMemory,
        trees:                  make(methodTrees, 0, 9),
        delims:                 render.Delims{Left: "{{", Right: "}}"},
        secureJSONPrefix:       "while(1);",
        trustedProxies:         []string{"0.0.0.0/0", "::/0"},
        trustedCIDRs:           defaultTrustedCIDRs,
    }
    engine.RouterGroup.engine = engine
    engine.pool.New = func() any {
        return engine.allocateContext()
    }
    return engine
}
```

这里New 函数默认的结构体

  

# 
