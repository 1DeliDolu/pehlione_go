
## 🧭 Path İçindeki Parametreler

```go
func main() {
  router := gin.Default()

  // Bu handler /user/john ile eşleşir, ancak /user/ veya /user ile eşleşmez
  router.GET("/user/:name", func(c *gin.Context) {
    name := c.Param("name")
    c.String(http.StatusOK, "Hello %s", name)
  })

  // Ancak bu, /user/john/ ve ayrıca /user/john/send ile de eşleşir
  // /user/john ile eşleşen başka bir route yoksa, /user/john/ adresine yönlendirir
  router.GET("/user/:name/*action", func(c *gin.Context) {
    name := c.Param("name")
    action := c.Param("action")
    message := name + " is " + action
    c.String(http.StatusOK, message)
  })

  router.Run(":8080")
}
```

