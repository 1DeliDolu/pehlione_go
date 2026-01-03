
## 🧾 XML/JSON/YAML/ProtoBuf Çıktısı (Rendering)

```go
func main() {
  router := gin.Default()

  // gin.H, map[string]interface{} için bir kısayoldur
  router.GET("/someJSON", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
  })

  router.GET("/moreJSON", func(c *gin.Context) {
    // Ayrıca bir struct da kullanabilirsiniz
    var msg struct {
      Name    string `json:"user"`
      Message string
      Number  int
    }
    msg.Name = "Lena"
    msg.Message = "hey"
    msg.Number = 123
    // msg.Name alanının JSON içinde "user" olduğuna dikkat edin
    // Çıktı şöyle olacaktır  :   {"user": "Lena", "Message": "hey", "Number": 123}
    c.JSON(http.StatusOK, msg)
  })

  router.GET("/someXML", func(c *gin.Context) {
    c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
  })

  router.GET("/someYAML", func(c *gin.Context) {
    c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
  })

  router.GET("/someProtoBuf", func(c *gin.Context) {
    reps := []int64{int64(1), int64(2)}
    label := "test"
    // Protobuf’un spesifik tanımı testdata/protoexample dosyasında yazılıdır.
    data := &protoexample.Test{
      Label: &label,
      Reps:  reps,
    }
    // data'nın yanıtta binary veriye dönüştüğüne dikkat edin
    // Çıktı protoexample.Test protobuf serialize edilmiş veri olacaktır
    c.ProtoBuf(http.StatusOK, data)
  })

  // 0.0.0.0:8080 üzerinde dinle ve servis et
  router.Run(":8080")
}
```

