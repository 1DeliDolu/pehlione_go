
## 🔀 Yönlendirmeler

HTTP yönlendirmesi (*redirect*) yapmak kolaydır. Hem dahili (*internal*) hem de harici (*external*) konumlar desteklenir.

## 🌍 Harici Yönlendirme

```go
router.GET("/test", func(c *gin.Context) {
  c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
})
```

## 📮 POST’tan HTTP Redirect

POST’tan HTTP yönlendirmesi yapmak. Issue: **#444**’e bakın.

```go
router.POST("/test", func(c *gin.Context) {
  c.Redirect(http.StatusFound, "/foo")
})
```

## 🧭 Router Redirect

Router yönlendirmesi için aşağıdaki gibi `HandleContext` kullanın.

```go
router.GET("/test", func(c *gin.Context) {
    c.Request.URL.Path = "/test2"
    router.HandleContext(c)
})
router.GET("/test2", func(c *gin.Context) {
    c.JSON(200, gin.H{"hello": "world"})
})
```

