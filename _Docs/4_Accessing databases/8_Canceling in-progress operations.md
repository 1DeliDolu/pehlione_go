
# ⛔ Devam Eden İşlemleri İptal Etme

Go `context.Context` kullanarak devam eden işlemleri yönetebilirsiniz. Bir `Context`, temsil ettiği genel işlemin iptal edilip edilmediğini ve artık ihtiyaç duyulup duyulmadığını raporlayabilen standart bir Go veri değeridir. Uygulamanızdaki fonksiyon çağrıları ve servisler boyunca bir `context.Context` geçirerek, bunlar işlemeye erken son verebilir ve artık gerekli olmadığında bir hata döndürebilir. `Context` hakkında daha fazla bilgi için **Go Concurrency Patterns: Context** konusuna bakın.

Örneğin şunları yapmak isteyebilirsiniz:

* Çok uzun süren işlemleri — çok uzun süren veritabanı işlemleri dahil — sonlandırmak.
* Başka bir yerden gelen iptal isteklerini yaymak (*propagate*); örneğin bir istemci bağlantıyı kapattığında.

Go geliştiricileri için birçok API, `Context` argümanı alan metodlar içerir; bu da uygulamanız genelinde `Context` kullanmanızı kolaylaştırır.

---

## ⏱️ Zaman aşımından sonra veritabanı işlemlerini iptal etme

Bir `Context` kullanarak, bir işlemden sonra işlemin iptal edileceği bir zaman aşımı (*timeout*) veya son tarih (*deadline*) belirleyebilirsiniz. Zaman aşımı veya son tarih içeren bir `Context` türetmek için `context.WithTimeout` veya `context.WithDeadline` çağırın.

Aşağıdaki zaman aşımı örneğindeki kod, bir `Context` türetir ve bunu `sql.DB`’nin `QueryContext` metoduna geçirir.

```go
func QueryWithTimeout(ctx context.Context) {
    // Zaman aşımı olan bir Context oluşturun.
    queryCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    // Zaman aşımı Context'ini bir sorguyla birlikte geçin.
    rows, err := db.QueryContext(queryCtx, "SELECT * FROM album")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // Dönen satırları ele alın.
}
```

Bu örnekte `queryCtx`’nin `ctx`’den türetildiği gibi, bir context başka bir dış context’ten türetildiğinde; dış context iptal edilirse, türetilen context de otomatik olarak iptal edilir. Örneğin HTTP sunucularında `http.Request.Context` metodu, isteğe bağlı bir context döndürür. HTTP istemcisi bağlantıyı keserse veya HTTP isteğini iptal ederse (HTTP/2’de mümkün), bu context iptal edilir. Yukarıdaki `QueryWithTimeout` fonksiyonuna bir HTTP isteğinin context’ini geçirmek, veritabanı sorgusunun ya genel HTTP isteği iptal edildiğinde ya da sorgu beş saniyeden uzun sürdüğünde erken durmasına neden olur.

> Not: Bir zaman aşımı veya son tarih ile yeni bir `Context` oluşturduğunuzda dönen `cancel` fonksiyonunu her zaman `defer` edin. Bu, kapsayan fonksiyon çıktığında yeni `Context` tarafından tutulan kaynakları serbest bırakır. Ayrıca `queryCtx`’yi de iptal eder; ancak fonksiyon döndüğünde, artık hiçbir şeyin `queryCtx`’yi kullanmıyor olması gerekir.

