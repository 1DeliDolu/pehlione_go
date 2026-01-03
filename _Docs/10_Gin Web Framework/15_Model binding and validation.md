
## 🧩 Model Bağlama ve Doğrulama

Bir istek gövdesini bir tipe bağlamak için *model binding* kullanın. Şu anda **JSON**, **XML**, **YAML** ve standart form değerleri (**foo=bar&boo=baz**) bağlamayı destekliyoruz.

Gin, doğrulama için **go-playground/validator/v10** kullanır. Etiket (tag) kullanımına dair tam dokümantasyonu buradan inceleyin.

Bağlamak istediğiniz tüm alanlarda ilgili *binding* etiketini ayarlamanız gerektiğini unutmayın. Örneğin JSON’dan bağlarken `json:"fieldname"` etiketini ayarlayın.

Ayrıca Gin, bağlama için iki yöntem seti sağlar:

## ✅ Must bind

* **Metotlar:** `Bind`, `BindJSON`, `BindXML`, `BindQuery`, `BindYAML`
* **Davranış:** Bu metotlar arka planda `MustBindWith` kullanır. Bir bağlama hatası olursa istek `c.AbortWithError(400, err).SetType(ErrorTypeBind)` ile durdurulur. Bu, yanıt durum kodunu **400** yapar ve `Content-Type` başlığı `text/plain; charset=utf-8` olarak ayarlanır. Bundan sonra yanıt kodunu değiştirmeye çalışırsanız şu uyarı ile sonuçlanır:
  `[GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422.`
  Davranışı daha fazla kontrol etmek istiyorsanız `ShouldBind` karşılıklarını kullanmayı düşünün.

## ✅ Should bind

* **Metotlar:** `ShouldBind`, `ShouldBindJSON`, `ShouldBindXML`, `ShouldBindQuery`, `ShouldBindYAML`
* **Davranış:** Bu metotlar arka planda `ShouldBindWith` kullanır. Bir bağlama hatası olursa hata döndürülür ve isteği/hata yönetimini uygun şekilde ele almak geliştiricinin sorumluluğundadır.

`Bind` metodunu kullandığınızda Gin, `Content-Type` başlığına göre uygun *binder*’ı çıkarsamaya çalışır. Ne bağladığınızdan eminseniz `MustBindWith` veya `ShouldBindWith` kullanabilirsiniz.

Ayrıca belirli alanların zorunlu olduğunu belirtebilirsiniz. Bir alan `binding:"required"` ile işaretlenmişse ve bağlama sırasında boş bir değer alırsa hata döndürülür.

Struct alanlarından biri başka bir struct ise (*nested struct*), doğru doğrulama için o struct’ın alanlarının da `binding:"required"` ile işaretlenmesi gerekir.

```go
// JSON'dan bağlama
type Login struct {
  User     string `form:"user" json:"user" xml:"user"  binding:"required"`
  Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
  router := gin.Default()

  // JSON bağlama örneği ({"user": "manu", "password": "123"})
  router.POST("/loginJSON", func(c *gin.Context) {
    var json Login
    if err := c.ShouldBindJSON(&json); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    if json.User != "manu" || json.Password != "123" {
      c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
      return
    }

    c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
  })

  // XML bağlama örneği (
  //  <?xml version="1.0" encoding="UTF-8"?>
  //  <root>
  //    <user>manu</user>
  //    <password>123</password>
  //  </root>)
  router.POST("/loginXML", func(c *gin.Context) {
    var xml Login
    if err := c.ShouldBindXML(&xml); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    if xml.User != "manu" || xml.Password != "123" {
      c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
      return
    }

    c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
  })

  // HTML form bağlama örneği (user=manu&password=123)
  router.POST("/loginForm", func(c *gin.Context) {
    var form Login
    // Bu, content-type başlığına göre hangi binder'ın kullanılacağını çıkarımlar.
    if err := c.ShouldBind(&form); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    if form.User != "manu" || form.Password != "123" {
      c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
      return
    }

    c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
  })

  // 0.0.0.0:8080 üzerinde dinle ve sun
  router.Run(":8080")
}
```

## 🧪 Örnek İstek

Terminal penceresi

```bash
$ curl -v -X POST \
  http://localhost:8080/loginJSON \
  -H 'content-type: application/json' \
  -d '{ "user": "manu" }'
> POST /loginJSON HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.51.0
> Accept: */*
> content-type: application/json
> Content-Length: 18
>
* upload completely sent off: 18 out of 18 bytes
< HTTP/1.1 400 Bad Request
< Content-Type: application/json; charset=utf-8
< Date: Fri, 04 Aug 2017 03:51:31 GMT
< Content-Length: 100
<
{"error":"Key: 'Login.Password' Error:Field validation for 'Password' failed on the 'required' tag"}
```

## 🚫 Doğrulamayı Atlamak

Yukarıdaki örneği yukarıdaki `curl` komutuyla çalıştırırken hata döner. Çünkü örnekte `Password` için `binding:"required"` kullanılmıştır. `Password` için `binding:"-"` kullanılırsa, aynı örnek tekrar çalıştırıldığında hata dönmez.

