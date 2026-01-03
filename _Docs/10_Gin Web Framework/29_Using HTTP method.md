
## 🌐 HTTP Metotlarını Kullanma

```go
func main() {
  // Varsayılan middleware ile bir gin router oluşturur:
  // logger ve recovery (çökmesiz) middleware
  router := gin.Default()

  router.GET("/someGet", getting)
  router.POST("/somePost", posting)
  router.PUT("/somePut", putting)
  router.DELETE("/someDelete", deleting)
  router.PATCH("/somePatch", patching)
  router.HEAD("/someHead", head)
  router.OPTIONS("/someOptions", options)

  // Varsayılan olarak :8080 üzerinde sunar, ancak bir
  // PORT ortam değişkeni tanımlandıysa onu kullanır.
  router.Run()
  // Sabit (hard coded) port için router.Run(":3000")
}
```

