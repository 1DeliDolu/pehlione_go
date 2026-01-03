
## 🚫 Varsayılan Olarak Middleware Olmadan

```go
r := gin.New()
```

Yerine:

```go
// Default: Logger ve Recovery middleware’ları zaten bağlı gelir
r := gin.Default()
```

