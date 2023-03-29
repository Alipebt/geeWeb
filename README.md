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

#### 1.2.3 http.Request 中的FormValue()函数

`FormValue()`函数会返回一个字符串类型的值，该值为参数key对应的表单数据的第一个值。

#### 1.2.4 http.Request 中的FormValue()函数

`Query()`函数会返回一个`url.Values`类型的值，该类型实际上是一个`map[string][]string`类型的别名，可以方便地获取和处理URL的查询参数。Get方法则可以从这个map中获取指定key对应的第一个值，如果没有找到指定的key，则会返回空字符串。

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

```go
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
```

实现了`ServeHTTP`方法，用来传入到`http.ListenAndServe()`

该方法首先从请求（`req`）中获取1）HTTP 请求方法（`Method`）和URI 路径（`URL.Path`），然后将二者组合成一个 `key` 字符串，用于在存储路由处理程序的映射表中查找对应的处理函数。

### 2.2 context.go

#### 2.2.1 Context结构体 

在`context.go`中设计一个`Context`结构体，其中封装了`http.ResponseWriter`和`http.Request`等。`Context` 随着每一个请求的出现而产生，请求的结束而销毁，和当前请求强相关的信息都应由 `Context` 承载。

```go
type Context struct {
	// origin object
	Writer http.ResponseWriter
	Req *http.Request
	// request info
	Path string
	Method string
	// response info
	StatusCode int
}
```

设计Context结构体的好处主要有以下几点：

1. 避免使用全局变量

在处理Web请求时，往往需要处理很多与请求相关的信息（如请求路径、HTTP方法、响应状态码等）。如果不使用Context对象，我们通常需要将这些信息存储在全局变量中，而全局变量会带来一系列问题，如难以追踪修改、存在线程安全性问题等。通过引入Context对象，我们可以避免使用全局变量，从而更好地管理和处理请求相关的信息。

2. 方便扩展和维护

使用Context对象可以使代码更加规范化和模块化，方便后续的扩展和维护。比如，当需要添加新的功能（如中间件支持）时，只需要在Context对象中添加对应的字段和方法即可，而不需要修改原有的代码。

3. 支持多种并发模型

通过Context对象，我们可以轻松地实现多种并发模型，如Goroutine之间共享数据、异步调用等。同时，在并发场景下，Context对象还能起到传递请求相关信息的作用，比如在一个请求被多个Goroutine同时处理时，我们可以通过Context来确保所有协程使用的是同一个请求对象。

#### 2.2.2 Context方法实现

```go
func (c *Context) PostFrom(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}
```

HTTP中的 POST 和 GET 方法都是用于客户端和服务器之间交换数据的方法。

HTTP POST 请求通常用于向服务器提交数据。在 Web 应用程序中，经常使用 POST 请求来提交表单数据，上传文件，执行数据库操作等。POST 请求可以将请求体中携带的数据传递给服务器处理，并根据需要作出响应。

而 HTTP GET 请求通常用于从服务器请求数据。在 Web 应用程序中，GET 请求可以在 URL 中附加查询字符串参数，以获取服务器上特定资源的数据。GET 请求在请求到达服务器时会被解析，然后服务器可以将请求结果返回给客户端。

```go
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}
```

该方法用于设置 HTTP 响应的状态码。

- 1xx (信息类状态码)：指示已经接收到请求，并且正在进一步处理中。

- 2xx (成功状态码)：代表请求已经被成功地接收、理解和处理。

- 3xx (重定向状态码)：需要客户端采取进一步的操作才能完成请求。

- 4xx (客户端错误状态码)：代表请求出现错误或者请求无法被执行。

- 5xx (服务器错误状态码)：代表服务器执行请求时发生了错误。

```go
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}
```

`Set(key, value)` 方法设置头部信息的值。其中，`key` 代表头部信息的名称，比如 `Content-Type`，`value` 代表头部信息的具体内容，比如 `application/json`。

```go
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

```

`...interface{}` 表示可变参数列表，它允许函数接受任意数量的参数。在 Go 函数中，使用 `...` 语法来表示一个可变参数。

例如，在上面的例子中，`values ...interface{}` 表示可以传递任意数量的参数到 `values` 中，并把它们存储为 `interface{}` 类型的值。

`fmt.Sprintf(format, values...)` 是一个字符串格式化函数，它将 format 和 values 作为参数传入，生成一个根据 format 格式化过的字符串。其中 `format` 是一个字符串格式化模板，`values...` 是一个 interface{} 类型的可变参数列表，代表需要填充到格式化字符串中的值。

例如：

```go
s1 := fmt.Sprintf("Hello %s!", "world") // s1 = "Hello world!"
s2 := fmt.Sprintf("Value: %d", 42) // s2 = "Value: 42"
```

第一个例子中，`"Hello %s!"` 是字符串格式化模板，`"world"` 是要填充到模板中的参数之一。

第二个例子中，`"Value: %d"` 是字符串格式化模板，数字 `42` 是要填充到模板中的参数之一。

注意到在 `String` 方法中，最后使用了 `values...` 语法，将该可变参数 `values` 展开为多个参数，这是因为 `Sprintf` 函数所需的是一个不定数量的 `interface{}` 参数而不是一个切片，因此需要使用 `...` 语法展开参数。