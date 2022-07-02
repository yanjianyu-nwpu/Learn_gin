# Engine

## 0 RouterGroup

```
type RouterGroup struct {
    Handlers HandlersChain
    basePath string
    engine   *Engine
    root     bool
}
```

bool 是否是根

用于Handlers 【】HandlerFunc

这里engine 是engine 的结构体函数

## 1 Use函数

```
// Use attaches a global middleware to the router. i.e. the middleware attached through Use() will be
// included in the handlers chain for every single request. Even 404, 405, static files...
// For example, this is the right place for a logger or error management middleware.
func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes {
    engine.RouterGroup.Use(middleware...)
    engine.rebuild404Handlers()
    engine.rebuild405Handlers()
    return engine
}
```

这个engine use函数 这里是routerGroup结构体的Use

函数如下

```
// Use adds middleware to the group, see example code in GitHub.
func (group *RouterGroup) Use(middleware ...HandlerFunc) IRoutes {
    group.Handlers = append(group.Handlers, middleware...)
    return group.returnObj()
}
```

基本就是吧handler 加进去

然后回到engine的use 这里 rebuild404Handlers

## 2 ServeHTTP

```
// ServeHTTP conforms to the http.Handler interface.
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    c := engine.pool.Get().(*Context)
    c.writermem.reset(w)
    c.Request = req
    c.reset()

    engine.handleHTTPRequest(c)

    engine.pool.Put(c)
}
```

大概就是

牛逼

### 2.1 GETvALUE

```
// Returns the handle registered with the given path (key). The values of
// wildcards are saved to a map.
// If no handle can be found, a TSR (trailing slash redirect) recommendation is
// made if a handle exists with an extra (without the) trailing slash for the
// given path.
// 返回注册的node 然后会被，计入
func (n *node) getValue(path string, params *Params, skippedNodes *[]skippedNode, unescape bool) (value nodeValue) {
    var globalParamsCount int16

walk: // Outer loop for walking the tree
    for {
        prefix := n.path
        if len(path) > len(prefix) {
            if path[:len(prefix)] == prefix {
                path = path[len(prefix):]

                // Try all the non-wildcard children first by matching the indices
                idxc := path[0]
                for i, c := range []byte(n.indices) {
                    if c == idxc {
                        //  strings.HasPrefix(n.children[len(n.children)-1].path, ":") == n.wildChild
                        if n.wildChild {
                            index := len(*skippedNodes)
                            *skippedNodes = (*skippedNodes)[:index+1]
                            (*skippedNodes)[index] = skippedNode{
                                path: prefix + path,
                                node: &node{
                                    path:      n.path,
                                    wildChild: n.wildChild,
                                    nType:     n.nType,
                                    priority:  n.priority,
                                    children:  n.children,
                                    handlers:  n.handlers,
                                    fullPath:  n.fullPath,
                                },
                                paramsCount: globalParamsCount,
                            }
                        }

                        n = n.children[i]
                        continue walk
                    }
                }

                if !n.wildChild {
                    // If the path at the end of the loop is not equal to '/' and the current node has no child nodes
                    // the current node needs to roll back to last vaild skippedNode
                    if path != "/" {
                        for l := len(*skippedNodes); l > 0; {
                            skippedNode := (*skippedNodes)[l-1]
                            *skippedNodes = (*skippedNodes)[:l-1]
                            if strings.HasSuffix(skippedNode.path, path) {
                                path = skippedNode.path
                                n = skippedNode.node
                                if value.params != nil {
                                    *value.params = (*value.params)[:skippedNode.paramsCount]
                                }
                                globalParamsCount = skippedNode.paramsCount
                                continue walk
                            }
                        }
                    }

                    // Nothing found.
                    // We can recommend to redirect to the same URL without a
                    // trailing slash if a leaf exists for that path.
                    value.tsr = path == "/" && n.handlers != nil
                    return
                }

                // Handle wildcard child, which is always at the end of the array
                n = n.children[len(n.children)-1]
                globalParamsCount++

                switch n.nType {
                case param:
                    // fix truncate the parameter
                    // tree_test.go  line: 204

                    // Find param end (either '/' or path end)
                    end := 0
                    for end < len(path) && path[end] != '/' {
                        end++
                    }

                    // Save param value
                    if params != nil && cap(*params) > 0 {
                        if value.params == nil {
                            value.params = params
                        }
                        // Expand slice within preallocated capacity
                        i := len(*value.params)
                        *value.params = (*value.params)[:i+1]
                        val := path[:end]
                        if unescape {
                            if v, err := url.QueryUnescape(val); err == nil {
                                val = v
                            }
                        }
                        (*value.params)[i] = Param{
                            Key:   n.path[1:],
                            Value: val,
                        }
                    }

                    // we need to go deeper!
                    if end < len(path) {
                        if len(n.children) > 0 {
                            path = path[end:]
                            n = n.children[0]
                            continue walk
                        }

                        // ... but we can't
                        value.tsr = len(path) == end+1
                        return
                    }

                    if value.handlers = n.handlers; value.handlers != nil {
                        value.fullPath = n.fullPath
                        return
                    }
                    if len(n.children) == 1 {
                        // No handle found. Check if a handle for this path + a
                        // trailing slash exists for TSR recommendation
                        n = n.children[0]
                        value.tsr = (n.path == "/" && n.handlers != nil) || (n.path == "" && n.indices == "/")
                    }
                    return

                case catchAll:
                    // Save param value
                    if params != nil {
                        if value.params == nil {
                            value.params = params
                        }
                        // Expand slice within preallocated capacity
                        i := len(*value.params)
                        *value.params = (*value.params)[:i+1]
                        val := path
                        if unescape {
                            if v, err := url.QueryUnescape(path); err == nil {
                                val = v
                            }
                        }
                        (*value.params)[i] = Param{
                            Key:   n.path[2:],
                            Value: val,
                        }
                    }

                    value.handlers = n.handlers
                    value.fullPath = n.fullPath
                    return

                default:
                    panic("invalid node type")
                }
            }
        }

        if path == prefix {
            // If the current path does not equal '/' and the node does not have a registered handle and the most recently matched node has a child node
            // the current node needs to roll back to last vaild skippedNode
            if n.handlers == nil && path != "/" {
                for l := len(*skippedNodes); l > 0; {
                    skippedNode := (*skippedNodes)[l-1]
                    *skippedNodes = (*skippedNodes)[:l-1]
                    if strings.HasSuffix(skippedNode.path, path) {
                        path = skippedNode.path
                        n = skippedNode.node
                        if value.params != nil {
                            *value.params = (*value.params)[:skippedNode.paramsCount]
                        }
                        globalParamsCount = skippedNode.paramsCount
                        continue walk
                    }
                }
                //    n = latestNode.children[len(latestNode.children)-1]
            }
            // We should have reached the node containing the handle.
            // Check if this node has a handle registered.
            if value.handlers = n.handlers; value.handlers != nil {
                value.fullPath = n.fullPath
                return
            }

            // If there is no handle for this route, but this route has a
            // wildcard child, there must be a handle for this path with an
            // additional trailing slash
            if path == "/" && n.wildChild && n.nType != root {
                value.tsr = true
                return
            }

            // No handle found. Check if a handle for this path + a
            // trailing slash exists for trailing slash recommendation
            for i, c := range []byte(n.indices) {
                if c == '/' {
                    n = n.children[i]
                    value.tsr = (len(n.path) == 1 && n.handlers != nil) ||
                        (n.nType == catchAll && n.children[0].handlers != nil)
                    return
                }
            }

            return
        }

        // Nothing found. We can recommend to redirect to the same URL with an
        // extra trailing slash if a leaf exists for that path
        value.tsr = path == "/" ||
            (len(prefix) == len(path)+1 && prefix[len(path)] == '/' &&
                path == prefix[:len(prefix)-1] && n.handlers != nil)

        // roll back to last valid skippedNode
        if !value.tsr && path != "/" {
            for l := len(*skippedNodes); l > 0; {
                skippedNode := (*skippedNodes)[l-1]
                *skippedNodes = (*skippedNodes)[:l-1]
                if strings.HasSuffix(skippedNode.path, path) {
                    path = skippedNode.path
                    n = skippedNode.node
                    if value.params != nil {
                        *value.params = (*value.params)[:skippedNode.paramsCount]
                    }
                    globalParamsCount = skippedNode.paramsCount
                    continue walk
                }
            }
        }

        return
    }
}
```

有个跳出多重循环的 break label

 walk 这样
