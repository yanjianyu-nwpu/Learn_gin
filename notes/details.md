# tips

## 0 写日志文件

```
f, _ := os.Create("gin.log")
gin.DefaultWriter = io.MultiWriter(f)
```

放到文件里面



## 1 日志格式

```

```
