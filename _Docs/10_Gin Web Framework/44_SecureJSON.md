
## 🛡️ SecureJSON

*JSON hijacking’i önlemek için SecureJSON kullanma.* Varsayılan olarak, verilen struct **dizi (array) değerleri** ise yanıt gövdesinin başına `"while(1);"` ekler.

```go
func main() {
  router := gin.Default()

  // Kendi secure json prefix’inizi de kullanabilirsiniz
  // router.SecureJsonPrefix(")]}',\n")

  router.GET("/someJSON", func(c *gin.Context) {
    names := []string{"lena", "austin", "foo"}

    // Çıktı şöyle olacaktır  :   while(1);["lena","austin","foo"]
    c.SecureJSON(http.StatusOK, names)
  })

  // 0.0.0.0:8080 üzerinde dinle ve servis et
  router.Run(":8080")
}
```

