
# 🧾 Hazırlanmış İfadeleri Kullanma

## 📚 İçindekiler

* Hazırlanmış ifade nedir?
* Hazırlanmış ifadeleri nasıl kullanırsınız
* Hazırlanmış ifadenin davranışı

Tekrar tekrar kullanılmak üzere bir hazırlanmış ifade tanımlayabilirsiniz. Bu, kodunuzun veritabanı işlemini her gerçekleştirdiğinde ifadeyi yeniden oluşturmanın ek yükünden kaçınarak, kodunuzun biraz daha hızlı çalışmasına yardımcı olabilir.

> Not: Hazırlanmış ifadelerdeki parametre yer tutucuları, kullandığınız DBMS ve sürücüye göre değişir. Örneğin Postgres için `pq` sürücüsü, `?` yerine `$1` gibi bir yer tutucu gerektirir.

---

## 🧩 Hazırlanmış ifade nedir?

Hazırlanmış ifade, DBMS tarafından ayrıştırılan (*parsed*) ve kaydedilen SQL’dir; tipik olarak yer tutucular içerir, ancak gerçek parametre değerlerini içermez. Daha sonra bu ifade, bir parametre değerleri kümesiyle çalıştırılabilir.

---

## 🛠️ Hazırlanmış ifadeleri nasıl kullanırsınız

Aynı SQL’i tekrar tekrar çalıştırmayı beklediğinizde, SQL ifadesini önceden hazırlamak ve gerektiğinde çalıştırmak için bir `sql.Stmt` kullanabilirsiniz.

Aşağıdaki örnek, veritabanından belirli bir albümü seçen bir hazırlanmış ifade oluşturur. `DB.Prepare`, verilen bir SQL metni için hazırlanmış bir ifadeyi temsil eden bir `sql.Stmt` döndürür. İfadeyi çalıştırmak için SQL’in parametrelerini `Stmt.Exec`, `Stmt.QueryRow` veya `Stmt.Query` metodlarına geçebilirsiniz.

```go
// AlbumByID belirtilen albümü getirir.
func AlbumByID(id int) (Album, error) {
    // Hazırlanmış bir ifade tanımlayın. Tipik olarak ifadeyi
    // başka bir yerde tanımlar ve bunun gibi fonksiyonlarda kullanmak üzere saklarsınız.
    stmt, err := db.Prepare("SELECT * FROM album WHERE id = ?")
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()

    var album Album

    // Hazırlanmış ifadeyi çalıştırın; ? yer tutucusuna karşılık gelen
    // parametre için bir id değeri geçin.
    err := stmt.QueryRow(id).Scan(&album.ID, &album.Title, &album.Artist, &album.Price, &album.Quantity)
    if err != nil {
        if err == sql.ErrNoRows {
            // Hiç satır dönmemesi durumunu ele alın.
        }
        return album, err
    }
    return album, nil
}
```

---

## ⚙️ Hazırlanmış ifadenin davranışı

Hazırlanmış bir `sql.Stmt`, ifadeyi çağırmak için alışıldık `Exec`, `QueryRow` ve `Query` metodlarını sağlar. Bu metodların kullanımı hakkında daha fazlası için **Querying for data** ve **Executing SQL statements that don’t return data** bölümlerine bakın.

Bununla birlikte, bir `sql.Stmt` zaten önceden belirlenmiş bir SQL ifadesini temsil ettiğinden, onun `Exec`, `QueryRow` ve `Query` metodları, SQL metnini atlayarak yalnızca yer tutuculara karşılık gelen SQL parametre değerlerini alır.

Yeni bir `sql.Stmt`’yi, onu nasıl kullanacağınıza bağlı olarak farklı yollarla tanımlayabilirsiniz.

* `DB.Prepare` ve `DB.PrepareContext`, tıpkı `DB.Exec` ve `DB.Query` gibi, bir transaction dışında tek başına, bağımsız olarak çalıştırılabilen bir hazırlanmış ifade oluşturur.
* `Tx.Prepare`, `Tx.PrepareContext`, `Tx.Stmt` ve `Tx.StmtContext`, belirli bir transaction içinde kullanım için hazırlanmış bir ifade oluşturur. `Prepare` ve `PrepareContext`, ifadeyi tanımlamak için SQL metni kullanır. `Stmt` ve `StmtContext` ise `DB.Prepare` veya `DB.PrepareContext` sonucunu kullanır. Yani, transaction için olmayan bir `sql.Stmt`’yi “bu transaction için” olan bir `sql.Stmt`’ye dönüştürürler.
* `Conn.PrepareContext`, rezerve edilmiş bir bağlantıyı temsil eden `sql.Conn` üzerinden bir hazırlanmış ifade oluşturur.

Kodunuz bir ifadeyi kullanmayı bitirdiğinde `stmt.Close` çağrıldığından emin olun. Bu, onunla ilişkilendirilmiş olabilecek veritabanı kaynaklarını (örneğin alttaki bağlantılar gibi) serbest bırakır. Yalnızca bir fonksiyon içinde yerel değişken olan ifadeler için `defer stmt.Close()` yeterlidir.

---

## 🧩 Hazırlanmış ifade oluşturmaya yönelik fonksiyonlar

| Fonksiyon             | Açıklama                                                                                                   |                                                                                                                                                                 |                  |                                                                                                                               |
| --------------------- | ---------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------- | ----------------------------------------------------------------------------------------------------------------------------- |
| `DB.Prepare`          | `DB.PrepareContext`                                                                                        | Bağımsız yürütme için bir ifade hazırlar veya `Tx.Stmt` kullanılarak transaction içinde kullanılacak bir hazırlanmış ifadeye dönüştürülecek bir ifade hazırlar. |                  |                                                                                                                               |
| `Tx.Prepare`          | `Tx.PrepareContext`                                                                                        | `Tx.Stmt`                                                                                                                                                       | `Tx.StmtContext` | Belirli bir transaction içinde kullanım için bir ifade hazırlar. Daha fazlası için **Executing transactions** konusuna bakın. |
| `Conn.PrepareContext` | Rezerve edilmiş bağlantılarla kullanım içindir. Daha fazlası için **Managing connections** konusuna bakın. |                                                                                                                                                                 |                  |                                                                                                                               |

