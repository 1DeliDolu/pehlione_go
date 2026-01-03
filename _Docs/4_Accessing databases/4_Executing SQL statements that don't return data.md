
# 🧾 Veri Döndürmeyen SQL İfadelerini Çalıştırma

Veri döndürmeyen veritabanı işlemleri yaptığınızda, *database/sql* paketindeki `Exec` veya `ExecContext` metodunu kullanın. Bu şekilde çalıştıracağınız SQL ifadelerine `INSERT`, `DELETE` ve `UPDATE` dahildir.

Sorgunuzun satır döndürebileceği durumlarda bunun yerine `Query` veya `QueryContext` metodunu kullanın. Daha fazlası için **Querying a database** konusuna bakın.

`ExecContext` metodu, `Exec` metodunun yaptığı gibi çalışır; ancak **Canceling in-progress operations** bölümünde açıklandığı üzere ek bir `context.Context` argümanı alır.

Aşağıdaki örnekteki kod, `album` tablosuna yeni bir `album` kaydı eklemek için bir ifade çalıştırmak üzere `DB.Exec` kullanır.

```go
func AddAlbum(alb Album) (int64, error) {
    result, err := db.Exec("INSERT INTO album (title, artist) VALUES (?, ?)", alb.Title, alb.Artist)
    if err != nil {
        return 0, fmt.Errorf("AddAlbum: %v", err)
    }

    // İstemci için yeni albümün oluşturulan ID'sini alın.
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("AddAlbum: %v", err)
    }
    // Yeni albümün ID'sini döndürün.
    return id, nil
}
```

`DB.Exec` iki değer döndürür: bir `sql.Result` ve bir `error`. Hata `nil` olduğunda, `Result` üzerinden son eklenen öğenin ID’sini (örnekteki gibi) alabilir veya işlemden etkilenen satır sayısını getirebilirsiniz.

> Not: Hazırlanmış ifadelerdeki (*prepared statements*) parametre yer tutucuları, kullandığınız DBMS ve sürücüye göre değişir. Örneğin Postgres için `pq` sürücüsü, `?` yerine `$1` gibi bir yer tutucu gerektirir.

Kodunuz aynı SQL ifadesini tekrar tekrar çalıştıracaksa, SQL ifadesinden yeniden kullanılabilir bir hazırlanmış ifade oluşturmak için `sql.Stmt` kullanmayı değerlendirin. Daha fazlası için **Using prepared statements** konusuna bakın.

> Dikkat: Bir SQL ifadesini birleştirmek için `fmt.Sprintf` gibi dizge biçimlendirme fonksiyonlarını kullanmayın! SQL enjeksiyonu riski oluşturabilirsiniz. Daha fazlası için **Avoiding SQL injection risk** konusuna bakın.

---

## 🧩 Satır döndürmeyen SQL ifadelerini çalıştırmaya yönelik fonksiyonlar

| Fonksiyon          | Açıklama                                                                                                   |                                                                                                                            |
| ------------------ | ---------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `DB.Exec`          | `DB.ExecContext`                                                                                           | Tek bir SQL ifadesini bağımsız olarak çalıştırır.                                                                          |
| `Tx.Exec`          | `Tx.ExecContext`                                                                                           | Daha büyük bir transaction içinde bir SQL ifadesi çalıştırır. Daha fazlası için **Executing transactions** konusuna bakın. |
| `Stmt.Exec`        | `Stmt.ExecContext`                                                                                         | Daha önce hazırlanmış bir SQL ifadesini çalıştırır. Daha fazlası için **Using prepared statements** konusuna bakın.        |
| `Conn.ExecContext` | Rezerve edilmiş bağlantılarla kullanım içindir. Daha fazlası için **Managing connections** konusuna bakın. |                                                                                                                            |

