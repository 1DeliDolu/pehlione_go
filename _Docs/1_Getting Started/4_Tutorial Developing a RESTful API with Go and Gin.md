## 📘 Eğitim: Go ve Gin ile RESTful API Geliştirme

## 📑 İçindekiler

* Ön koşullar
* API uç noktalarını tasarlayın
* Kodunuz için bir klasör oluşturun
* Veriyi oluşturun
* Tüm öğeleri döndüren bir handler yazın
* Yeni bir öğe ekleyen bir handler yazın
* Belirli bir öğeyi döndüren bir handler yazın
* Sonuç
* Tamamlanmış kod

Bu eğitim, Go ile Gin Web Framework (Gin) kullanarak RESTful bir web servis API’si yazmanın temellerini tanıtır.

Bu eğitimden en iyi şekilde yararlanmak için Go ve araçları hakkında temel bir aşinalığa sahip olmanız faydalı olur. Eğer Go ile ilk kez karşılaşıyorsanız, hızlı bir giriş için **Tutorial: Get started with Go** bölümüne bakın.

Gin; web servisleri de dahil olmak üzere web uygulamaları geliştirirken ilgili birçok kodlama görevini basitleştirir. Bu eğitimde Gin’i; istekleri yönlendirmek (*route* etmek), istek ayrıntılarını almak ve yanıtlar için JSON’u *marshal* etmek için kullanacaksınız.

Bu eğitimde iki uç noktası olan bir RESTful API sunucusu oluşturacaksınız. Örnek projeniz, eski ( *vintage* ) caz plakları hakkında veri barındıran bir depo olacaktır.

Eğitim aşağıdaki bölümleri içerir:

* API uç noktalarını tasarlayın.
* Kodunuz için bir klasör oluşturun.
* Veriyi oluşturun.
* Tüm öğeleri döndüren bir handler yazın.
* Yeni bir öğe ekleyen bir handler yazın.
* Belirli bir öğeyi döndüren bir handler yazın.

Not: Diğer eğitimler için bkz. Tutorials.

Bunu Google Cloud Shell’de tamamlayacağınız etkileşimli bir eğitim olarak denemek için aşağıdaki düğmeye tıklayın.

Open in Cloud Shell

---

## ✅ Ön koşullar¶

* Go 1.16 veya daha yenisinin yüklü olması. Kurulum talimatları için bkz. Installing Go.
* Kodunuzu düzenlemek için bir araç. Elinizdeki herhangi bir metin düzenleyici işinizi görür.
* Bir komut terminali. Go; Linux ve Mac’te herhangi bir terminalde, Windows’ta PowerShell veya cmd üzerinde iyi çalışır.
* `curl` aracı. Linux ve Mac’te genellikle zaten yüklüdür. Windows’ta Windows 10 Insider build 17063 ve daha yenisinde dahildir. Daha eski Windows sürümleri için yüklemeniz gerekebilir. Daha fazlası için bkz. Tar and Curl Come to Windows.

---

## 🧭 API uç noktalarını tasarlayın¶

Vinyl üzerinde eski kayıtlar satan bir mağazaya erişim sağlayan bir API oluşturacaksınız. Bu nedenle, istemcinin kullanıcılar için albümleri alıp ekleyebilmesi adına uç noktalar sağlamanız gerekir.

Bir API geliştirirken genellikle uç noktaları tasarlayarak başlarsınız. Uç noktalar anlaşılması kolay olduğunda, API kullanıcıları daha başarılı olur.

Bu eğitimde oluşturacağınız uç noktalar şunlardır.

### `/albums`

* **GET** – Tüm albümlerin listesini JSON olarak döndürür.
* **POST** – JSON olarak gönderilen istek verisinden yeni bir albüm ekler.

### `/albums/:id`

* **GET** – ID’sine göre bir albümü alır ve albüm verisini JSON olarak döndürür.

Sonraki adımda, kodunuz için bir klasör oluşturacaksınız.

---

## 📁 Kodunuz için bir klasör oluşturun¶

Başlamak için, yazacağınız kod için bir proje oluşturun.

Bir komut istemi açın ve ev dizininize geçin.

Linux veya Mac’te:

```bash
$ cd
```

Windows’ta:

```cmd
C:\> cd %HOMEPATH%
```

Komut istemini kullanarak, kodunuz için `web-service-gin` adlı bir dizin oluşturun.

```bash
$ mkdir web-service-gin
$ cd web-service-gin
```

Bağımlılıkları yönetebileceğiniz bir modül oluşturun.

Kodunuzun içinde olacağı modül yolunu vererek `go mod init` komutunu çalıştırın.

```bash
$ go mod init example/web-service-gin
go: creating new go.mod: module example/web-service-gin
```

Bu komut, eklediğiniz bağımlılıkların izlenmesi için listeleneceği bir `go.mod` dosyası oluşturur. Bir modülü modül yolu ile adlandırma hakkında daha fazla bilgi için bkz. Managing dependencies.

Sonraki adımda, veriyi işlemek için veri yapılarını tasarlayacaksınız.

---

## 🗃️ Veriyi oluşturun¶

Eğitimi basit tutmak için veriyi bellekte saklayacaksınız. Daha tipik bir API, bir veritabanıyla etkileşir.

Veriyi bellekte saklamanın, sunucuyu her durdurduğunuzda albüm kümesinin kaybolacağı ve sunucuyu başlattığınızda yeniden oluşturulacağı anlamına geldiğini unutmayın.

### Kodu yazın¶

Metin düzenleyicinizi kullanarak `web-service` dizininde `main.go` adlı bir dosya oluşturun. Go kodunuzu bu dosyada yazacaksınız.

`main.go` içine, dosyanın en üstüne aşağıdaki `package` bildirimini yapıştırın.

```go
package main
```

Tek başına çalışan bir program (bir kütüphanenin aksine) her zaman `package main` içindedir.

`package` bildiriminden sonra, albüm verisini bellekte saklamak için kullanacağınız aşağıdaki `album` struct bildirimini yapıştırın.

`json:"artist"` gibi struct etiketleri, struct içeriği JSON’a serileştirildiğinde bir alanın adının ne olması gerektiğini belirtir. Etiketler olmadan JSON, struct’ın büyük harfli alan adlarını kullanırdı; bu da JSON’da çok yaygın bir stil değildir.

```go
// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}
```

Az önce eklediğiniz struct bildiriminden sonra, başlangıç verisi olarak kullanacağınız `album` struct’larından oluşan aşağıdaki slice’ı yapıştırın.

```go
// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
```

Sonraki adımda, ilk uç noktanızı uygulamak için kod yazacaksınız.

---

## 📤 Tüm öğeleri döndüren bir handler yazın¶

İstemci `GET /albums` isteği yaptığında, tüm albümleri JSON olarak döndürmek istersiniz.

Bunu yapmak için şunları yazacaksınız:

* Yanıt hazırlama mantığı
* İstek yolunu mantığınıza eşleyen kod

Bunun, çalışma zamanında ( *runtime* ) yürütülecek sıranın tersi olduğuna dikkat edin; önce bağımlılıkları, sonra onlara bağlı olan kodu ekliyorsunuz.

### Kodu yazın¶

Bir önceki bölümde eklediğiniz struct kodunun altına, albüm listesini döndürmek için aşağıdaki kodu yapıştırın.

Bu `getAlbums` fonksiyonu, `album` struct slice’ından JSON oluşturur ve JSON’u yanıta yazar.

```go
// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}
```

Bu kodda şunları yaparsınız:

* `gin.Context` parametresi alan bir `getAlbums` fonksiyonu yazarsınız. Bu fonksiyona herhangi bir ad verebilirdiniz; ne Gin ne de Go belirli bir fonksiyon adı formatı gerektirir.
* `gin.Context`, Gin’in en önemli parçasıdır. İstek ayrıntılarını taşır, JSON’u doğrular ve serileştirir, ve daha fazlasını yapar. (Benzer adına rağmen, bu Go’nun yerleşik `context` paketiyle farklıdır.)
* Struct’ı JSON’a serileştirmek ve yanıta eklemek için `Context.IndentedJSON` çağırırsınız.
* Fonksiyonun ilk argümanı, istemciye göndermek istediğiniz HTTP durum kodudur. Burada `net/http` paketindeki `StatusOK` sabitini geçirerek **200 OK** belirtiyorsunuz.
* Daha kompakt JSON göndermek için `Context.IndentedJSON` yerine `Context.JSON` çağrısı kullanabileceğinizi unutmayın. Pratikte, girintili form hata ayıklarken çok daha kolaydır ve boyut farkı genellikle küçüktür.

`main.go` dosyasının üst kısmına, `albums` slice bildiriminden hemen sonra, handler fonksiyonunu bir uç nokta yoluna atamak için aşağıdaki kodu yapıştırın.

Bu, `getAlbums` fonksiyonunun `/albums` uç nokta yoluna gelen istekleri ele alacağı bir ilişkilendirme kurar.

```go
func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)

    router.Run("localhost:8080")
}
```

Bu kodda şunları yaparsınız:

* `Default` kullanarak bir Gin router başlatırsınız.
* `GET` fonksiyonunu kullanarak **GET** HTTP metodunu ve `/albums` yolunu bir handler fonksiyonuyla ilişkilendirirsiniz.
* `getAlbums` fonksiyonunun adını geçtiğinize dikkat edin. Bu, `getAlbums()` (parantezlere dikkat) geçirerek fonksiyonun sonucunu geçirmekten farklıdır.
* `Run` fonksiyonunu kullanarak router’ı bir `http.Server`’a bağlar ve sunucuyu başlatırsınız.

`main.go` dosyasının üst tarafında, `package` bildiriminden hemen sonra, az önce yazdığınız kodu desteklemek için gereken paketleri içe aktarın.

İlk satırlar şu şekilde görünmelidir:

```go
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)
```

`main.go` dosyasını kaydedin.

### Kodu çalıştırın¶

Gin modülünü bağımlılık olarak izlemeye başlayın.

Komut satırında `go get` kullanarak `github.com/gin-gonic/gin` modülünü modülünüz için bağımlılık olarak ekleyin. Nokta argümanı, “geçerli dizindeki kod için bağımlılıkları al” anlamına gelir.

```bash
$ go get .
go get: added github.com/gin-gonic/gin v1.7.2
```

Go, bir önceki adımda eklediğiniz import bildirimini karşılamak için bu bağımlılığı çözüp indirdi.

`main.go` dosyasını içeren dizinde komut satırından kodu çalıştırın. Nokta argümanı, “geçerli dizindeki kodu çalıştır” anlamına gelir.

```bash
$ go run .
```

Kod çalışmaya başladığında, istek gönderebileceğiniz çalışan bir HTTP sunucunuz olur.

Yeni bir komut satırı penceresinden `curl` ile çalışan web servisinize bir istek yapın.

```bash
$ curl http://localhost:8080/albums
```

Komut, servis için tohumladığınız ( *seed* ) veriyi göstermelidir.

```json
[
        {
                "id": "1",
                "title": "Blue Train",
                "artist": "John Coltrane",
                "price": 56.99
        },
        {
                "id": "2",
                "title": "Jeru",
                "artist": "Gerry Mulligan",
                "price": 17.99
        },
        {
                "id": "3",
                "title": "Sarah Vaughan and Clifford Brown",
                "artist": "Sarah Vaughan",
                "price": 39.99
        }
]
```

Bir API başlattınız! Sonraki bölümde, bir öğe eklemek için **POST** isteğini ele alan kodla başka bir uç nokta oluşturacaksınız.

---

## ➕ Yeni bir öğe ekleyen bir handler yazın¶

İstemci `/albums` üzerinde bir **POST** isteği yaptığında, istek gövdesinde tanımlanan albümü mevcut albüm verisine eklemek istersiniz.

Bunu yapmak için şunları yazacaksınız:

* Yeni albümü mevcut listeye ekleme mantığı
* POST isteğini mantığınıza yönlendirecek biraz routing kodu

### Kodu yazın¶

Albümleri listeye eklemek için kod ekleyin.

Import ifadelerinden sonra bir yere aşağıdaki kodu yapıştırın. (Dosyanın sonu bu kod için iyi bir yerdir; ancak Go fonksiyonları hangi sırayla bildirdiğinizi zorunlu kılmaz.)

```go
// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}
```

Bu kodda şunları yaparsınız:

* İstek gövdesini `newAlbum`’a bağlamak için `Context.BindJSON` kullanırsınız.
* JSON’dan başlatılan `album` struct’ını `albums` slice’ına eklersiniz.
* Yanıta, eklediğiniz albümü temsil eden JSON ile birlikte **201** durum kodu eklersiniz.

`main` fonksiyonunuzu aşağıdaki gibi `router.POST` çağrısını da içerecek şekilde değiştirin.

```go
func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}
```

Bu kodda şunları yaparsınız:

* `/albums` yolundaki **POST** metodunu `postAlbums` fonksiyonuyla ilişkilendirirsiniz.

Gin ile, bir handler’ı HTTP metodu-yol kombinasyonu ile ilişkilendirebilirsiniz. Bu şekilde, istemcinin kullandığı metoda bağlı olarak aynı yola gönderilen istekleri ayrı ayrı yönlendirebilirsiniz.

### Kodu çalıştırın¶

Sunucu önceki bölümden beri hâlâ çalışıyorsa, durdurun.

`main.go` dosyasını içeren dizinde komut satırından kodu çalıştırın.

```bash
$ go run .
```

Farklı bir komut satırı penceresinden `curl` kullanarak çalışan web servisinize bir istek yapın.

```bash
$ curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

Komut, eklenen albüm için header’ları ve JSON’u göstermelidir.

```http
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Wed, 02 Jun 2021 00:34:12 GMT
Content-Length: 116

{
    "id": "4",
    "title": "The Modern Sound of Betty Carter",
    "artist": "Betty Carter",
    "price": 49.99
}
```

Önceki bölümde olduğu gibi, yeni albümün eklendiğini doğrulamak için `curl` ile albümlerin tam listesini alın.

```bash
$ curl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "GET"
```

Komut albüm listesini göstermelidir.

```json
[
        {
                "id": "1",
                "title": "Blue Train",
                "artist": "John Coltrane",
                "price": 56.99
        },
        {
                "id": "2",
                "title": "Jeru",
                "artist": "Gerry Mulligan",
                "price": 17.99
        },
        {
                "id": "3",
                "title": "Sarah Vaughan and Clifford Brown",
                "artist": "Sarah Vaughan",
                "price": 39.99
        },
        {
                "id": "4",
                "title": "The Modern Sound of Betty Carter",
                "artist": "Betty Carter",
                "price": 49.99
        }
]
```

Sonraki bölümde, belirli bir öğe için GET isteğini ele alacak kod ekleyeceksiniz.

---

## 🎯 Belirli bir öğeyi döndüren bir handler yazın¶

İstemci `GET /albums/[id]` isteği yaptığında, ID’si `id` yol parametresiyle eşleşen albümü döndürmek istersiniz.

Bunu yapmak için şunları yapacaksınız:

* İstenen albümü almak için mantık eklemek
* Yolu mantığa eşlemek

### Kodu yazın¶

Bir önceki bölümde eklediğiniz `postAlbums` fonksiyonunun altına, belirli bir albümü almak için aşağıdaki kodu yapıştırın.

Bu `getAlbumByID` fonksiyonu, istek yolundaki ID’yi çıkarır ve eşleşen bir albümü bulur.

```go
// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
```

Bu kodda şunları yaparsınız:

* URL’den `id` yol parametresini almak için `Context.Param` kullanırsınız. Bu handler’ı bir yola eşlerken, yolda parametre için bir yer tutucu kullanacaksınız.
* Slice içindeki `album` struct’ları üzerinde dönerek ID alanı `id` parametresiyle eşleşen bir albüm ararsınız. Bulunursa, albümü JSON’a serileştirir ve **200 OK** HTTP koduyla yanıt olarak döndürürsünüz.
* Yukarıda bahsedildiği gibi, gerçek dünyadaki bir servis büyük olasılıkla bu aramayı yapmak için bir veritabanı sorgusu kullanırdı.
* Albüm bulunamazsa `http.StatusNotFound` ile bir HTTP **404** hatası döndürürsünüz.

Son olarak `main` fonksiyonunuzu, yolu artık `/albums/:id` olan yeni bir `router.GET` çağrısını içerecek şekilde aşağıdaki örnekteki gibi değiştirin.

```go
func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}
```

Bu kodda şunları yaparsınız:

* `/albums/:id` yolunu `getAlbumByID` fonksiyonuyla ilişkilendirirsiniz. Gin’de, yolda bir öğenin önündeki iki nokta üst üste ( *colon* ), o öğenin bir yol parametresi olduğunu belirtir.

### Kodu çalıştırın¶

Sunucu önceki bölümden beri hâlâ çalışıyorsa, durdurun.

`main.go` dosyasını içeren dizinde komut satırından kodu çalıştırarak sunucuyu başlatın.

```bash
$ go run .
```

Farklı bir komut satırı penceresinden `curl` ile çalışan web servisinize bir istek yapın.

```bash
$ curl http://localhost:8080/albums/2
```

Komut, kullandığınız ID’ye sahip albümün JSON’unu göstermelidir. Albüm bulunamazsa, bir hata mesajı içeren JSON alırsınız.

```json
{
        "id": "2",
        "title": "Jeru",
        "artist": "Gerry Mulligan",
        "price": 17.99
}
```

---

## 🏁 Sonuç¶

Tebrikler! Go ve Gin kullanarak basit bir RESTful web servisi yazdınız.

Önerilen sonraki konular:

* Go’ya yeniyseniz, Effective Go ve How to write Go code içinde açıklanan faydalı en iyi uygulamaları bulabilirsiniz.
* Go Tour, Go temellerine adım adım bir giriş için harikadır.
* Gin hakkında daha fazla bilgi için Gin Web Framework paket dokümantasyonuna veya Gin Web Framework dokümanlarına bakın.

---

## 🧾 Tamamlanmış kod¶

Bu bölüm, bu eğitimle oluşturduğunuz uygulamanın kodunu içerir.

```go
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
```
