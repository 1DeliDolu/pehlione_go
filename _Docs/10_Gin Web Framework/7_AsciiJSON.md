
## 🔤 AsciiJSON

### 🧾 AsciiJSON Kullanımı

AsciiJSON kullanarak, ASCII olmayan karakterleri escape ederek yalnızca ASCII içeren JSON üretirsiniz.

```go
func main() {
  router := gin.Default()

  router.GET("/someJSON", func(c *gin.Context) {
    data := map[string]interface{}{
      "lang": "GO语言",
      "tag":  "<br>",
    }

    // şu çıktıyı üretir : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
    c.AsciiJSON(http.StatusOK, data)
  })

  // 0.0.0.0:8080 üzerinde dinle ve servis et
  router.Run(":8080")
}
```

