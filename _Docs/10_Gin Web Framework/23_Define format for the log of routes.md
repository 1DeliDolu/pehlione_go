
## 🧾 Route Log Formatını Tanımlama

Route’ların varsayılan log çıktısı şöyledir:

```text
[GIN-debug] POST   /foo                      --> main.main.func1 (3 handlers)
[GIN-debug] GET    /bar                      --> main.main.func2 (3 handlers)
[GIN-debug] GET    /status                   --> main.main.func3 (3 handlers)
```

Bu bilgiyi belirli bir formatta (ör. **JSON**, **key-value** veya başka bir format) loglamak istiyorsanız, bu formatı `gin.DebugPrintRouteFunc` ile tanımlayabilirsiniz. Aşağıdaki örnekte tüm route’lar standart `log` paketi ile loglanır; ancak ihtiyaçlarınıza uygun başka log araçları da kullanabilirsiniz.

```go
import (
  "log"
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
    log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
  }

  router.POST("/foo", func(c *gin.Context) {
    c.JSON(http.StatusOK, "foo")
  })

  router.GET("/bar", func(c *gin.Context) {
    c.JSON(http.StatusOK, "bar")
  })

  router.GET("/status", func(c *gin.Context) {
    c.JSON(http.StatusOK, "ok")
  })

  // http://0.0.0.0:8080 üzerinde dinle ve sun
  router.Run()
}
```

