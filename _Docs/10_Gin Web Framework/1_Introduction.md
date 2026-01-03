
## 🚀 Giriş

Gin, Go (Golang) ile yazılmış bir web çatısıdır (*web framework*). Martini benzeri bir API sunar; ancak `httprouter` sayesinde çok daha iyi performans sağlar ve **40 kata kadar daha hızlı** olabilir. Hem performansa hem de yüksek üretkenliğe ihtiyacınız varsa, Gin’i seveceksiniz.

Bu bölümde Gin’in ne olduğunu, hangi problemleri çözdüğünü ve projenize nasıl yardımcı olabileceğini adım adım ele alacağız.

Ya da Gin’i projenizde hemen kullanmaya hazırsanız, *Quickstart* bölümünü ziyaret edin.

---

## ✨ Özellikler

### ⚡ Hızlı

Radix-tree tabanlı yönlendirme (*routing*), düşük bellek ayak izi (*memory footprint*). *Reflection* yok. Öngörülebilir API performansı.

---

### 🧩 Middleware Desteği

Gelen bir HTTP isteği, bir *middleware* zinciri ve en sondaki nihai aksiyon tarafından işlenebilir. Örneğin: *Logger*, *Authorization*, *GZIP* ve son olarak DB’ye bir mesaj yazma.

---

### 🛡️ Çökmesiz

Gin, bir HTTP isteği sırasında oluşan `panic` durumunu yakalayıp toparlayabilir (*recover*). Bu sayede sunucunuz her zaman erişilebilir olur. Örnek olarak, bu `panic` durumunu Sentry’ye raporlamak da mümkündür!

---

### ✅ JSON Doğrulama

Gin, bir isteğin JSON’unu ayrıştırıp doğrulayabilir; örneğin zorunlu değerlerin varlığını kontrol edebilir.

---

### 🧱 Route Gruplama

Route’larınızı daha iyi organize edin. Yetkilendirme gerektirenler vs gerektirmeyenler, farklı API sürümleri… Ayrıca gruplar, performansı düşürmeden sınırsız şekilde iç içe (*nested*) kullanılabilir.

---

### 🧯 Hata Yönetimi

Gin, bir HTTP isteği sırasında oluşan tüm hataları toplamak için kullanışlı bir yöntem sunar. Sonuçta bir *middleware*, bu hataları bir log dosyasına, bir veritabanına yazabilir ve ağ üzerinden gönderebilir.

---

### 🖼️ Dahili Render Desteği

Gin, JSON, XML ve HTML render etmek için kullanımı kolay bir API sağlar.

---

### 🧰 Genişletilebilir

Yeni bir *middleware* oluşturmak çok kolaydır; sadece örnek koda göz atın.

