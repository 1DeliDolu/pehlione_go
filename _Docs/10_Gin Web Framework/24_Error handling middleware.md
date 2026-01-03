
## 🧯 Hata Yönetimi Middleware’i

Tipik bir RESTful uygulamada, herhangi bir route’ta şu tür hatalarla karşılaşabilirsiniz:

* Kullanıcıdan gelen geçersiz giriş
* Veritabanı hataları
* Yetkisiz erişim
* Sunucu içi (internal) bug’lar

Varsayılan olarak Gin, her route içinde `c.Error(err)` kullanarak hataları manuel yönetmenize izin verir. Ancak bu yaklaşım hızlıca tekrarlı ve tutarsız hale gelebilir.

Bunu çözmek için tüm hataları tek bir yerde ele alan özel bir middleware kullanabiliriz. Bu middleware her isteğin ardından çalışır ve Gin context’e eklenen hataları (`c.Errors`) kontrol eder. Bir hata bulursa, uygun bir status code ile yapılandırılmış bir JSON yanıtı döner.

## 🧪 Örnek

```go
import (
  "errors"
  "net/http"
  "github.com/gin-gonic/gin"
)

// ErrorHandler hataları yakalar ve tutarlı bir JSON hata yanıtı döndürür
func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next() // Adım 1: Önce isteği işle.

        // Adım 2: Context'e herhangi bir hata eklenip eklenmediğini kontrol et
        if len(c.Errors) > 0 {
            // Adım 3: Son hatayı kullan
            err := c.Errors.Last().Err

            // Adım 4: Genel bir hata mesajı ile yanıt ver
            c.JSON(http.StatusInternalServerError, map[string]any{
                "success": false,
                "message": err.Error(),
            })
        }

        // Hata bulunmazsa uygulanacak diğer adımlar (varsa)
    }
}

func main() {
    r := gin.Default()

    // Hata yönetimi middleware’ini ekle
    r.Use(ErrorHandler())

    r.GET("/ok", func(c *gin.Context) {
        somethingWentWrong := false

        if somethingWentWrong {
            c.Error(errors.New("something went wrong"))
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "success": true,
            "message": "Everything is fine!",
        })
    })

    r.GET("/error", func(c *gin.Context) {
        somethingWentWrong := true

        if somethingWentWrong {
            c.Error(errors.New("something went wrong"))
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "success": true,
            "message": "Everything is fine!",
        })
    })

    r.Run()
}
```

## 🧩 Genişletmeler

* Hataları status code’lara eşleme
* Hata kodlarına göre farklı hata yanıtları üretme
* Şunu kullanarak hataları loglama

## ✅ Hata Yönetimi Middleware’inin Faydaları

* **Tutarlılık:** Tüm hatalar aynı formatı takip eder
* **Temiz route’lar:** İş mantığı hata formatlamasından ayrılır
* **Daha az tekrar:** Her handler’da hata yönetimi mantığını tekrar etmeye gerek kalmaz

