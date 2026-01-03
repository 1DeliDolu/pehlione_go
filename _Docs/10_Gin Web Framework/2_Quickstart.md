
## 🚀 Quickstart

Gin quickstart’a hoş geldiniz! Bu rehber, Gin’i kurma, bir proje oluşturma ve ilk API’nizi çalıştırma adımlarında size yol gösterir—böylece güvenle web servisleri geliştirmeye başlayabilirsiniz.

---

## 🧩 Ön Koşullar

### 🟦 Go

Go: **1.23 veya daha yüksek** bir sürüm kurulu olmalıdır.

Go’nun `PATH` içinde olduğundan ve terminalinizden kullanılabildiğinden emin olun. Go kurulum yardımı için resmi dokümantasyona bakın.

---

## 🛠️ Adım 1: Gin’i Kurun ve Projenizi Başlatın

Önce yeni bir proje klasörü oluşturup bir Go modülü başlatın:

```bash
mkdir gin-quickstart && cd gin-quickstart
go mod init gin-quickstart
```

Gin’i bağımlılık olarak ekleyin:

```bash
go get -u github.com/gin-gonic/gin
```

---

## 🧪 Adım 2: İlk Gin Uygulamanızı Oluşturun

`main.go` adlı bir dosya oluşturun:

```bash
touch main.go
```

`main.go` dosyasını açın ve aşağıdaki kodu ekleyin:

```go
package main

import "github.com/gin-gonic/gin"

func main() {
  router := gin.Default()
  router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
  router.Run() // varsayılan olarak 0.0.0.0:8080 üzerinde dinler
}
```

---

## ▶️ Adım 3: API Sunucunuzu Çalıştırın

Sunucunuzu şu komutla başlatın:

```bash
go run main.go
```

Tarayıcıda `http://localhost:8080/ping` adresine gidin; şunu görmelisiniz:

```json
{"message":"pong"}
```

---

## ➕ Ek Örnek: Gin ile net/http Kullanımı

Yanıt kodları için `net/http` sabitlerini kullanmak isterseniz, onu da import edin:

```go
package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func main() {
  router := gin.Default()
  router.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  router.Run()
}
```

---

## 📚 İpuçları ve Kaynaklar

Go’da yeni misiniz? Resmi Go dokümantasyonunda Go kodu yazmayı ve çalıştırmayı öğrenin.

Gin kavramlarını uygulamalı öğrenmek mi istiyorsunuz? Etkileşimli alıştırmalar ve eğitimler için *Learning Resources* bölümüne göz atın.

Tam özellikli bir örneğe mi ihtiyacınız var? Şununla scaffold oluşturmayı deneyin:

```bash
curl https://raw.githubusercontent.com/gin-gonic/examples/master/basic/main.go > main.go
```

Daha ayrıntılı dokümantasyon için Gin kaynak kodu dokümanlarını ziyaret edin.

