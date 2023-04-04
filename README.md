# geeWeb

## 1 http协议

HTTP（Hypertext Transfer Protocol）是一种用于传输超媒体文件（例如HTML，图片等）的应用层协议。它基于请求/响应模式，客户端向服务器发送一个HTTP请求并接收服务器返回的HTTP响应。

HTTP请求由三个部分组成：请求行、消息报头和请求正文。以下是HTTP请求的结构：

```
[请求行]
[消息报头]
[请求正文]
```

- **请求行** 包含请求方法、URL和HTTP协议版本。

  ```
  GET /index.html HTTP/1.1
  ```

- **消息报头** 包括若干行属性信息，每行由属性名和属性值组成，用冒号分隔。消息报头也可以为空。

  ```
  Host: www.example.com
  User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3
  Accept-Encoding: gzip, deflate, sdch
  ```

- **请求正文** 可选，包含任意有效载荷数据（通常在POST请求中使用）。

同样，HTTP响应也由三个部分组成：状态行、消息报头和响应正文。以下是HTTP响应的结构：

```
[状态行]
[消息报头]
[响应正文]
```

- **状态行** 包含HTTP协议版本、状态码和状态描述。

  ```
  HTTP/1.1 200 OK
  ```

- **消息报头** 包括若干行属性信息，每行由属性名和属性值组成，用冒号分隔。消息报头也可以为空。

  ```
  Content-Type: text/html;charset=utf-8
  Server: Apache/2.4.23 (Win32)
  ```

- **响应正文** 可选，包含任意有效载荷数据（通常是HTML页面或其他类型文件的内容）。

## 2 net/http库

HTTP（Hypertext Transfer Protocol，超文本传输协议）是一个用于传输超媒体文档的应用层协议。一个 HTTP 请求-响应事务通常由以下几部分组成：

1. 起始行（Start Line）：

起始行包括请求行和响应行两种格式。

- 请求行：包括 HTTP 方法、请求 URL 和协议版本。例如："GET /index.html HTTP/1.1"
- 响应行：包括协议版本、状态码和状态短语。例如："HTTP/1.1 200 OK"

2. 头部字段（Header Fields）：

头部字段包含了若干个键值对，每个键值对之间通过冒号分隔，多个键值对以换行符隔开。常见的头部字段有：

- `Host`：指示要访问的主机名；
- `User-Agent`：指示客户端类型，如浏览器、爬虫等；
- `Content-Type`：指示请求或响应消息体的类型；
- `Accept`：指示客户端能够接受的内容类型；
- `Set-Cookie`：用于在客户端设置 Cookie 等。

3. 空行（空白行）：

空行是 HTTP 头部与消息体之间必须加入的隔行符，表示已经结束了头部字段的传输。

4. 消息体（Message Body）：

消息体是可选的，它包含了由上述头部字段描述的内容信息。HTTP 消息体可包含文本、HTML、XML、JSON、二进制等数据，取决于 Content-Type 头部字段。

以上是 HTTP 请求和响应的基本组成部分。需要注意的是，HTTP 是一个无状态协议，每个请求都在一个独立的事务里进行处理，服务器并不会保留之前请求的任何信息。为了解决这个问题，常用的方式是使用 Cookie 或者 Token 等机制记录用户状态信息。

另外，HTTPS（Hypertext Transfer Protocol Secure）是一个基于 HTTP 协议的加密传输协议，其结构与 HTTP 协议类似，但具有安全性更高的特点。

### 2.1 url

是用于标识互联网上某个资源的地址。它通常由多个部分组成，例如：

```
https://www.example.com:8080/path/to/myfile.html?key1=valu
```

### 2.2 接口

#### 2.2.1 `http.HandleFunc()` 用于将HTTP请求路由到指定的处理器函数。该函数接受两个参数：一个字符串类型的路径和一个 `func(http.ResponseWriter, *http.Request)` 类型的处理器函数。

`http.ResponseWriter` 接口用于构建 HTTP 响应。

`http.Request` 结构体则表示一个 HTTP 请求。它包含了如下字段：

- `Method`：请求的方法，如 GET、POST 等。
- `URL`：请求的 URL。
- `Header`：请求头。
- `Body`：请求体（如果有）。

#### 2.2.2 `http.ListenAndServe()` 启动HTTP服务器并监听来自客户端的请求。

`http.ListenAndServe()`函数的第二个参数是一个`Handler` 接口类型，它定义了对于任何进入的 HTTP 请求应该执行哪些操作，通常使用`nil`值表示采用默认的处理器。如果需要自定义处理器，则可以传递实现了 `Handler` 接口的结构体指针，并在该结构体中实现所需的逻辑。

#### 2.2.3 http.Request 中的FormValue()函数

`FormValue()`函数会返回一个字符串类型的值，该值为参数key对应的表单数据的第一个值。

#### 2.2.4 http.Request 中的FormQuery()函数

`Query()`函数会返回一个`url.Values`类型的值，该类型实际上是一个`map[string][]string`类型的别名，可以方便地获取和处理URL的查询参数。Get方法则可以从这个map中获取指定key对应的第一个值，如果没有找到指定的key，则会返回空字符串。

#### 2.2.5 http.POST()

`http.Post()` 函数发送的数据将会放入 HTTP 请求的消息体中。

## 3 geeWeb

### 3.1 gee.go

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

### 3.2 context.go

#### 3.2.1 Context结构体 

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

#### 3.2.2 Context方法实现

```go
func (c *Context) PostFrom(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}
```

- `req.FormValue()` 用于获取表单（POST 或 PUT）提交的参数值。在 HTTP 请求中，表单参数通常被编码为 x-www-form-urlencoded 格式，可以使用 `req.FormValue()` 方法获取这些参数值。如果对应的参数不存在或为空，则该方法返回空字符串。
- `req.URL.Query()` 用于获取 URL 查询参数。在 HTTP 请求中，URL 可以包含查询字符串，类似于 `http://example.com/path?name=value&foo=bar` 。我们可以使用 `req.URL.Query()` 方法来获取这个查询字符串中所有参数的值，并以一个 map 的形式进行返回。

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

### 3.3 trie.go

#### 3.3.1 前缀树

```shell
/
├── /:lang
│   ├── /intro
│   ├── /tutorial
│   └── /doc
├── /about
└── /p
     ├── /blog
	└── /related
```



```go
type node struct {
	pattern string // 待匹配路由，如 /p/:lang
	part	string // 路由中的一部分，如 :lang
	child	[]*node // 子节点，如[doc,tutorial,intro]
	isWild	bool // 是否精确匹配，part 含有 ：或 * 时为true
}
```

为了实现动态路由匹配，加上了`isWild`这个参数。即当我们匹配 `/p/go/doc/`这个路由时，第一层节点，`p`精准匹配到了`p`，第二层节点，`go`模糊匹配到`:lang`，那么将会把`lang`这个参数赋值为`go`，继续下一层匹配。

```go
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}
```

该代码段接收三个参数，`pattern`(路由路径)，`parts`(通过 `/` 分隔开的路径片段)，`height`(当前的高度，与parts的个数作比较)。

首先判断当前`parts`是否已经被完全遍历过了，如果是，那么就将 `pattern` 的值赋给当前节点 `n` 的 `pattern` 字段，并返回。当 `parts` 切片被遍历完时，我们就已经将 `pattern` 插入到了 Trie 树中。

如果parts还没有被完全遍历，则取出当前`height`的`part`作为被对比的字符串，并在当前节点遍历查找子节点是否为`part`或者为模糊匹配`:`和`*`。如果子节点不存在，则创建一个新的子节点。如果当前子节点为模糊匹配，则直接设置`isWild`为true，以便能进入该`part`的下一级。最后将这个节点加入到当前节点的`children`数组中。

接着，函数对该子节点进行递归调用，同时传入 `pattern`、`parts` 和 `height+1` 作为参数。递归调用的目的是将 `pattern` 插入到以该子节点为根节点的子树中。

```go
func (n *node) metchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.metchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
```

函数首先判断当前函数是否已经遍历到了 `parts` 切片的末尾或者当前节点的键名以 `*` 开头。如果是，那么就判断当前节点 `n` 的 `pattern` 字段是否为空。如果为空，则说明该节点并没有被插入任何真实的模式字符串，返回 `nil` 

如果当前节点不满足终止条件，函数会取出 `parts` 中第 `height` 个元素作为键名 `part`，然后在当前节点 `n` 的子节点数组中查找所有键名与 `part` 匹配的子节点，返回一个子节点数组 `children`。

接着，函数对 `children` 数组中的每一个子节点进行递归调用

如果递归调用 `result` 返回的结果不为空，说明已经找到了匹配的节点，直接返回该结果即可，否则继续循环查找下一个节点。

最终，如果整个 `children` 数组都被遍历过了，那么说明 Trie 树中并没有匹配的节点，函数返回 `nil` 表示未查询到结果。

```go
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchPaths := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	
	if !ok {
		return nil, nil
	}

	n :=root.search(searchPaths, 0)

	if n != nil {
		parts :=parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchPaths[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchPaths[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}
```

这段代码是一个`router`类型的方法，接受两个字符串参数`method`和`path`，返回两个值：第一个值是一个指向节点的指针，第二个值是一个命名参数的映射表。具体流程如下：

1. 将传递进来的path进行解析，解析成一个路径数组searchPaths。
2. 创建一个空的命名参数映射表params。
3. 从路由树中查找对应的method的根节点，如果没有找到则返回nil，nil。
4. 调用根节点的search方法，在树中搜索满足给定路径的最后一个节点，如果找到了，则将其pattern按照解析路径之后的格式解析成一个parts数组。
5. 对于parts数组中每一个以":"开头的部分而言，将其添加到params命名参数映射表中。
6. 对于parts数组中每一个以"*"开头的部分而言，将其余部分拼接起来，并将其添加到params命名参数映射表中作为该参数的值。
7. 返回找到的节点和参数映射表。若未找到，则返回nil，nil。

`params`是一个命名参数的映射表，其中key表示命名参数的名称，value表示命名参数的值。例如当路由中包含像`:name`或者`*file`这样的占位符时，对应的参数会被提取出来，并加到映射表`params`中。

`pattern`是路由规则的字符串表示形式，而`path`是请求URL的实际路径。它们之间的关系是，路由规则中定义的占位符（如`:name`或`*file`）可以匹配到请求URL中相应的部分，从而提取出相应的参数。

以下是一个示例：

假设我们有一个路径规则 `/user/:id/*path`， 它允许任何以 `/user/` 开头的请求通过，并提取 URL 中的 `id` 和 `path` 两个参数。 如果路径为 `/user/123/file/filename.zip`，则路由器根据 `/user/:id/*path`找到节点并解析出命名参数：

```go
params = map[string]string{
    "id":   "123",
    "path": "file/filename.zip",
}
```

可以看到，路由规则中的 `:id` 和 `*path` 已经被匹配了相应的参数值。

在路由匹配中，`pattern` 和 `path` 的关系是一个匹配的过程。

`pattern` 是指声明路由规则时的模式，定义了一个 URL 匹配的模板。比如 `/users/:id` 就是一个模式，其中 `:id` 表示一个路径参数，它将匹配任意字符，直到遇到下一个斜杠为止。

`path` 是指浏览器地址栏中实际输入的 URL 路径部分。当请求到来时，会将 `path` 与路由模式中的 `pattern` 进行匹配，以确定该请求应该由哪个路由处理。

如果 `path` 成功匹配到一个适合的 `pattern`，则路由程序将按照该 `pattern` 的规则处理匹配到的路径参数，并执行该路由的回调函数或加载对应的组件。如果没有找到匹配的路由，则可以显示 404 页面或者跳转到默认页面。
