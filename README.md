# geeWeb

## 1 net/http库

### 1.1 url

是用于标识互联网上某个资源的地址。它通常由多个部分组成，例如：

```
https://www.example.com:8080/path/to/myfile.html?key1=valu
```

### 1.2 接口

#### 1.2.1 `http.HandleFunc()` 用于将HTTP请求路由到指定的处理器函数。该函数接受两个参数：一个字符串类型的路径和一个 `func(http.ResponseWriter, *http.Request)` 类型的处理器函数。

`http.ResponseWriter` 接口用于构建 HTTP 响应。

`http.Request` 结构体则表示一个 HTTP 请求。它包含了如下字段：

- `Method`：请求的方法，如 GET、POST 等。
- `URL`：请求的 URL。
- `Header`：请求头。
- `Body`：请求体（如果有）。

#### 1.2.2 `http.ListenAndServe()` 启动HTTP服务器并监听来自客户端的请求。

`http.ListenAndServe()`函数的第二个参数是一个`Handler` 接口类型，它定义了对于任何进入的 HTTP 请求应该执行哪些操作，通常使用`nil`值表示采用默认的处理器。如果需要自定义处理器，则可以传递实现了 `Handler` 接口的结构体指针，并在该结构体中实现所需的逻辑。

## 2 geeWeb

### 2.1 gee.go

```go
type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}
```

`router`为路由表，用于将字符串类型的 `URL `匹配到相应的处理函数上。

```go
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}
```

调用了 `make` 函数来创建一个空的映射，然后将它作为参数赋值给 `router` 属性。

```go
func (engine *Engine) addRoute(method string, pattern string,handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}
```

这个函数有三个参数：`mothod`、`pattern` 和 `handler`。其中，`mothod` 参数表示 HTTP 方法，例如 GET、POST 等；`pattern` 参数表示 URL 匹配模式，例如 `/user/:name`、`/article/:id` 等；`handler` 参数表示处理函数，它的类型是 `HandlerFunc`。

在上面的代码中，我们将 `method` 和 `pattern` 拼接起来作为映射的键名，然后将 `handler` 作为映射的键值存储到 `router` 属性中。这样，在请求到达时，只需要从 `router` 中查找与请求 URL 匹配的处理函数，就可以完成请求的处理了。
