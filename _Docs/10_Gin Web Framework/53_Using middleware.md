
## 🧩 Middleware Kullanımı

```go
func main() {
  // Varsayılan olarak hiçbir middleware olmadan bir router oluşturur
  router := gin.New()

  // Global middleware
  // Logger middleware, GIN_MODE=release ile ayarlasanız bile logları gin.DefaultWriter'a yazar.
  // Varsayılan olarak gin.DefaultWriter = os.Stdout
  router.Use(gin.Logger())

  // Recovery middleware, herhangi bir panic durumunda toparlar ve varsa 500 yazar.
  router.Use(gin.Recovery())

  // Route bazlı middleware, istediğiniz kadar ekleyebilirsiniz.
  router.GET("/benchmark", MyBenchLogger(), benchEndpoint)

  // Authorization grubu
  // authorized := router.Group("/", AuthRequired())
  // birebir aynı olan:
  authorized := router.Group("/")
  // grup bazlı middleware! bu örnekte özel oluşturulan
  // AuthRequired() middleware sadece "authorized" grubunda kullanılır.
  authorized.Use(AuthRequired())
  {
    authorized.POST("/login", loginEndpoint)
    authorized.POST("/submit", submitEndpoint)
    authorized.POST("/read", readEndpoint)

    // iç içe (nested) grup
    testing := authorized.Group("testing")
    testing.GET("/analytics", analyticsEndpoint)
  }

  // 0.0.0.0:8080 üzerinde dinle ve servis et
  router.Run(":8080")
}
```

