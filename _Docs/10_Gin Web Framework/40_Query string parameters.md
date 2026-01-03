
## 🔎 Query String Parametreleri

```go
func main() {
  router := gin.Default()

  // Query string parametreleri, mevcut alttaki request nesnesi kullanılarak parse edilir.
  // Request şu URL ile eşleşen bir isteğe yanıt verir: /welcome?firstname=Jane&lastname=Doe
  router.GET("/welcome", func(c *gin.Context) {
    firstname := c.DefaultQuery("firstname", "Guest")
    lastname := c.Query("lastname") // c.Request.URL.Query().Get("lastname") için kısayol

    c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
  })
  router.Run(":8080")
}
```

