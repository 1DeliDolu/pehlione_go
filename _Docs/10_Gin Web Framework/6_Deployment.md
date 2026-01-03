
## 🚀 Deployment

Gin projeleri, herhangi bir bulut sağlayıcısında kolayca deploy edilebilir.

---

## 🚄 Railway

Railway, uygulama ve servisleri deploy etmek, yönetmek ve ölçeklemek için geliştirilmiş, modern bir bulut geliştirme platformudur. Sunuculardan gözlemlenebilirliğe (*observability*) kadar altyapı yığınınızı tek bir ölçeklenebilir, kullanımı kolay platformla sadeleştirir.

Gin projelerinizi deploy etmek için Railway rehberini takip edin.

---

## 🌿 Seenode

Seenode, uygulamaları hızlı ve verimli şekilde deploy etmek isteyen geliştiriciler için özel olarak tasarlanmış modern bir bulut platformudur. Git tabanlı deploy, otomatik SSL sertifikaları, yerleşik veritabanları ve Gin uygulamalarınızı dakikalar içinde yayına almanızı sağlayan sade bir arayüz sunar.

Gin projelerinizi deploy etmek için Seenode rehberini takip edin.

---

## 🌍 Koyeb

Koyeb, uygulamaları global olarak deploy etmek için geliştirici dostu bir *serverless* platformdur. Git tabanlı deploy, TLS şifreleme, yerel otomatik ölçekleme (*native autoscaling*), global edge ağı ve yerleşik servis mesh & discovery sunar.

Gin projelerinizi deploy etmek için Koyeb rehberini takip edin.

---

## ☁️ Qovery

Qovery, veritabanları, SSL, global CDN ve Git ile otomatik deploy’lar içeren ücretsiz bulut hosting sağlar.

Daha fazla bilgi için Qovery’ye bakın.

---

## 🧩 Render

Render, Go için yerel destek, tamamen yönetilen SSL, veritabanları, *zero-downtime* deploy’lar, HTTP/2 ve websocket desteği sunan modern bir bulut platformudur.

Gin projelerini deploy etmek için Render rehberini takip edin.

---

## 🟦 Google App Engine

GAE, Go uygulamalarını deploy etmek için iki yol sunar. *Standard environment* kullanımı daha kolaydır; ancak daha az özelleştirilebilir ve güvenlik nedeniyle *syscall*’lara izin vermez. *Flexible environment* ise herhangi bir framework veya kütüphaneyi çalıştırabilir.

Daha fazla bilgi edinmek ve tercih ettiğiniz ortamı seçmek için Go on Google App Engine bölümüne bakın.

---

## 🏠 Self Hosted

Gin projeleri, *self-hosted* şekilde de deploy edilebilir. Deployment mimarisi ve güvenlik değerlendirmeleri hedef ortama göre değişir. Aşağıdaki bölüm, deployment planlarken dikkate alınacak yapılandırma seçeneklerine dair yalnızca üst seviye bir genel bakış sunar.

---

## ⚙️ Yapılandırma Seçenekleri

Gin proje deploy’ları, ortam değişkenleri (*environment variables*) veya doğrudan kod içinde ayarlanarak ince ayar yapılabilir.

Gin’i yapılandırmak için kullanılabilen ortam değişkenleri şunlardır:

| Ortam Değişkeni | Açıklama                                                                                                                                                                                                                                               |
| --------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `PORT`          | `router.Run()` ile (yani argüman vermeden) Gin sunucusu başlatılırken dinlenecek TCP portu.                                                                                                                                                            |
| `GIN_MODE`      | `debug`, `release` veya `test` değerlerinden birine ayarlanır. Gin modlarının yönetimini sağlar; örneğin ne zaman debug çıktıları üretileceği. Ayrıca kod içinde `gin.SetMode(gin.ReleaseMode)` veya `gin.SetMode(gin.TestMode)` ile de ayarlanabilir. |

Aşağıdaki kod, Gin’i yapılandırmak için kullanılabilir.

```go
// Gin için bind adresi veya port belirtmeyin. Varsayılan olarak tüm arayüzlerde 8080 portuna bind eder.
// Argümansız `Run()` kullanırken dinleme portunu değiştirmek için `PORT` ortam değişkeni kullanılabilir.
router := gin.Default()
router.Run()

// Gin için bind adresi ve portu belirtin.
router := gin.Default()
router.Run("192.168.1.100:8080")

// Sadece dinlenecek portu belirtin. Tüm arayüzlere bind eder.
router := gin.Default()
router.Run(":8080")

// Hangi IP adreslerinin veya CIDR’ların, gerçek istemci IP adreslerini belirlemek için header’ları ayarlamada güvenilir sayılacağını belirleyin.
// Ek detaylar için dokümantasyona bakın.
router := gin.Default()
router.SetTrustedProxies([]string{"192.168.1.2"})
```

---

## 🛑 Tüm Proxy’lere Güvenmeyin

Gin, gerçek istemci IP’sini (varsa) tutacak header’ları belirtmenize ve ayrıca bu header’lardan birini belirtmesine izin verilen proxy’leri (veya doğrudan istemcileri) güvenilir olarak tanımlamanıza olanak sağlar.

`gin.Engine` üzerinde `SetTrustedProxies()` fonksiyonunu kullanarak, istemci IP’siyle ilişkili request header’larının güvenilir sayılacağı ağ adreslerini veya ağ CIDR’larını belirleyin. Bunlar IPv4 adresleri, IPv4 CIDR’ları, IPv6 adresleri veya IPv6 CIDR’ları olabilir.

Dikkat: Yukarıdaki fonksiyonla güvenilir bir proxy belirtmezseniz, Gin varsayılan olarak tüm proxy’lere güvenir; bu **güvenli değildir**. Aynı zamanda, hiç proxy kullanmıyorsanız bu özelliği `Engine.SetTrustedProxies(nil)` ile devre dışı bırakabilirsiniz; böylece `Context.ClientIP()` gereksiz hesaplamalardan kaçınmak için doğrudan uzak adresi döndürür.

```go
import (
  "fmt"

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  router.SetTrustedProxies([]string{"192.168.1.2"})

  router.GET("/", func(c *gin.Context) {
    // Eğer istemci 192.168.1.2 ise, X-Forwarded-For
    // header’ını kullanarak, o header’ın güvenilir
    // kısımlarından orijinal istemci IP’sini çıkar.
    // Aksi halde, doğrudan istemci IP’sini döndür.
    fmt.Printf("ClientIP: %s\n", c.ClientIP())
  })
  router.Run()
}
```

---

## 🧾 Not

Eğer bir CDN servisi kullanıyorsanız, `TrustedProxies` kontrolünü atlamak için `Engine.TrustedPlatform` ayarlayabilirsiniz; bu, `TrustedProxies`’den daha yüksek önceliğe sahiptir. Aşağıdaki örneğe bakın:

```go
import (
  "fmt"

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  // Önceden tanımlı header: gin.PlatformXXX
  // Google App Engine
  router.TrustedPlatform = gin.PlatformGoogleAppEngine
  // Cloudflare
  router.TrustedPlatform = gin.PlatformCloudflare
  // Fly.io
  router.TrustedPlatform = gin.PlatformFlyIO
  // Ya da kendi güvenilir request header’ınızı ayarlayabilirsiniz. Ancak CDN’iniz
  // kullanıcıların bu header’ı geçmesini engelliyor olmalıdır!
  // Örneğin CDN’iniz istemci IP’sini X-CDN-Client-IP içinde koyuyorsa:
  router.TrustedPlatform = "X-CDN-Client-IP"

  router.GET("/", func(c *gin.Context) {
    // TrustedPlatform ayarlarsanız, ClientIP() ilgili
    // header’ı çözümler ve IP’yi doğrudan döndürür
    fmt.Printf("ClientIP: %s\n", c.ClientIP())
  })
  router.Run()
}
```

