
# 🔁 Transaction Çalıştırma

## 📚 İçindekiler

* En iyi uygulamalar
* Örnek

Bir transaction’ı temsil eden `sql.Tx` kullanarak veritabanı transaction’larını çalıştırabilirsiniz. Transaction’a özgü semantiği temsil eden `Commit` ve `Rollback` metodlarına ek olarak, `sql.Tx`, yaygın veritabanı işlemlerini gerçekleştirmek için kullandığınız metodların tamamına sahiptir. `sql.Tx` almak için `DB.Begin` veya `DB.BeginTx` çağırırsınız.

Bir veritabanı transaction’ı, birden fazla işlemi daha büyük bir hedefin parçası olarak gruplar. Tüm işlemler başarılı olmak zorundadır ya da hiçbiri başarılı olamaz; her iki durumda da verinin bütünlüğü korunur. Tipik olarak bir transaction iş akışı şunları içerir:

* Transaction’ı başlatma.
* Bir dizi veritabanı işlemi gerçekleştirme.
* Hata oluşmazsa, veritabanı değişikliklerini uygulamak için transaction’ı *commit* etme.
* Hata oluşursa, veritabanını değişmeden bırakmak için transaction’ı *rollback* etme.

`sql` paketi; bir transaction’ı başlatmak ve sonuçlandırmak için metodlar ve aradaki veritabanı işlemlerini gerçekleştirmek için metodlar sağlar. Bu metodlar, yukarıdaki iş akışındaki dört adıma karşılık gelir.

---

## ▶️ Transaction başlatma

`DB.Begin` veya `DB.BeginTx`, yeni bir veritabanı transaction’ı başlatır ve onu temsil eden bir `sql.Tx` döndürür.

---

## 🛠️ Veritabanı işlemlerini gerçekleştirme

Bir `sql.Tx` kullanarak, tek bir bağlantıyı kullanan bir dizi işlemde veritabanını sorgulayabilir veya güncelleyebilirsiniz. Bunu desteklemek için `Tx`, aşağıdaki metodları dışa aktarır:

* `INSERT`, `UPDATE` ve `DELETE` gibi SQL ifadeleriyle veritabanı değişiklikleri yapmak için `Exec` ve `ExecContext`.
  Daha fazlası için **Executing SQL statements that don’t return data** konusuna bakın.

* Satır döndüren işlemler için `Query`, `QueryContext`, `QueryRow` ve `QueryRowContext`.
  Daha fazlası için **Querying for data** konusuna bakın.

* Hazırlanmış ifadeleri önceden tanımlamak için `Prepare`, `PrepareContext`, `Stmt` ve `StmtContext`.
  Daha fazlası için **Using prepared statements** konusuna bakın.

---

## ✅ Transaction’ı bitirme

Transaction’ı aşağıdakilerden biriyle sonlandırın:

* `Tx.Commit` kullanarak transaction’ı *commit* edin.
  `Commit` başarılı olursa (`nil` hata döndürürse), tüm sorgu sonuçları geçerli olarak doğrulanır ve yürütülen tüm güncellemeler veritabanına tek bir atomik değişiklik olarak uygulanır. `Commit` başarısız olursa, `Tx` üzerindeki `Query` ve `Exec` sonuçlarının tamamı geçersiz olarak atılmalıdır.

* `Tx.Rollback` kullanarak transaction’ı *rollback* edin.
  `Tx.Rollback` başarısız olsa bile, transaction artık geçerli olmayacaktır ve veritabanına *commit* edilmiş de olmayacaktır.

---

## ✅ En iyi uygulamalar

Transaction’ların bazen gerektirdiği karmaşık semantik ve bağlantı yönetiminde daha iyi yol almak için aşağıdaki en iyi uygulamaları izleyin.

* Bu bölümde açıklanan API’leri kullanarak transaction’ları yönetin. `BEGIN` ve `COMMIT` gibi transaction ile ilgili SQL ifadelerini doğrudan kullanmayın — bunu yapmak, özellikle eşzamanlı (*concurrent*) programlarda veritabanınızı öngörülemez bir duruma sokabilir.
* Bir transaction kullanırken, transaction dışı `sql.DB` metodlarını doğrudan çağırmamaya dikkat edin; çünkü bunlar transaction dışında çalışır ve kodunuza veritabanı durumunun tutarsız bir görünümünü verebilir veya hatta kilitlenmelere (*deadlocks*) neden olabilir.

---

## 🧪 Örnek

Aşağıdaki örnekteki kod, bir albüm için yeni bir müşteri siparişi oluşturmak üzere bir transaction kullanır. Bu sırada kod şunları yapar:

* Bir transaction başlatır.
* Transaction’ın rollback’ini `defer` eder. Transaction başarılı olursa, fonksiyon çıkmadan önce commit edileceği için ertelenmiş rollback çağrısı etkisiz (*no-op*) olur. Transaction başarısız olursa commit edilmeyeceğinden, fonksiyon çıkarken rollback çağrılır.
* Müşterinin sipariş ettiği albüm için yeterli envanter olup olmadığını doğrular.
* Yeterliyse, envanter sayısını günceller ve sipariş edilen albüm adedi kadar azaltır.
* Yeni bir sipariş oluşturur ve istemci için yeni siparişin oluşturulan ID’sini alır.
* Transaction’ı commit eder ve ID’yi döndürür.

Bu örnek, `context.Context` argümanı alan `Tx` metodlarını kullanır. Bu, fonksiyonun yürütülmesinin — veritabanı işlemleri dahil — çok uzun sürerse veya istemci bağlantısı kapanırsa iptal edilebilmesini sağlar. Daha fazlası için **Canceling in-progress operations** konusuna bakın.

```go
// CreateOrder bir albüm için sipariş oluşturur ve yeni sipariş ID'sini döndürür.
func CreateOrder(ctx context.Context, albumID, quantity, custID int) (orderID int64, err error) {

    // Başarısızlık sonuçlarını hazırlamak için bir yardımcı fonksiyon oluşturun.
    fail := func(err error) (int64, error) {
        return 0, fmt.Errorf("CreateOrder: %v", err)
    }

    // Transaction istekleri yapmak için bir Tx alın.
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return fail(err)
    }
    // Herhangi bir şey başarısız olursa diye bir rollback'i defer edin.
    defer tx.Rollback()

    // Sipariş için albüm envanterinin yeterli olduğunu doğrulayın.
    var enough bool
    if err = tx.QueryRowContext(ctx, "SELECT (quantity >= ?) from album where id = ?",
        quantity, albumID).Scan(&enough); err != nil {
        if err == sql.ErrNoRows {
            return fail(fmt.Errorf("no such album"))
        }
        return fail(err)
    }
    if !enough {
        return fail(fmt.Errorf("not enough inventory"))
    }

    // Siparişteki miktarı düşmek için albüm envanterini güncelleyin.
    _, err = tx.ExecContext(ctx, "UPDATE album SET quantity = quantity - ? WHERE id = ?",
        quantity, albumID)
    if err != nil {
        return fail(err)
    }

    // album_order tablosunda yeni bir satır oluşturun.
    result, err := tx.ExecContext(ctx, "INSERT INTO album_order (album_id, cust_id, quantity, date) VALUES (?, ?, ?, ?)",
        albumID, custID, quantity, time.Now())
    if err != nil {
        return fail(err)
    }
    // Az önce oluşturulan sipariş öğesinin ID'sini alın.
    orderID, err = result.LastInsertId()
    if err != nil {
        return fail(err)
    }

    // Transaction'ı commit edin.
    if err = tx.Commit(); err != nil {
        return fail(err)
    }

    // Sipariş ID'sini döndürün.
    return orderID, nil
}
```

