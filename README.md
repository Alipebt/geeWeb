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

