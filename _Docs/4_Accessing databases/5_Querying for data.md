
# 🔎 Veri Sorgulama

## 📚 İçindekiler

* Tek bir satır için sorgulama
* Birden çok satır için sorgulama
* Null olabilir sütun değerlerini ele alma
* Sütunlardan veri alma
* Birden çok sonuç kümesini ele alma

Veri döndüren bir SQL ifadesi çalıştırırken, *database/sql* paketinde sağlanan `Query` metodlarından birini kullanın. Bunların her biri, `Scan` metodu ile verilerini değişkenlere kopyalayabileceğiniz bir `Row` veya `Rows` döndürür. Örneğin, `SELECT` ifadelerini çalıştırmak için bu metodları kullanırsınız.

Veri döndürmeyen bir ifade çalıştırırken bunun yerine `Exec` veya `ExecContext` kullanabilirsiniz. Daha fazlası için **Executing statements that don’t return data** konusuna bakın.

*database/sql* paketi, bir sorguyu sonuçlar için çalıştırmanın iki yolunu sağlar.

* Tek bir satır için sorgulama – `QueryRow` veritabanından en fazla tek bir `Row` döndürür. Daha fazlası için **Querying for a single row** konusuna bakın.
* Birden çok satır için sorgulama – `Query` eşleşen tüm satırları, kodunuzun üzerinde döngü kurabileceği bir `Rows` struct’ı olarak döndürür. Daha fazlası için **Querying for multiple rows** konusuna bakın.

Kodunuz aynı SQL ifadesini tekrar tekrar çalıştıracaksa, hazırlanmış ifade (*prepared statement*) kullanmayı değerlendirin. Daha fazlası için **Using prepared statements** konusuna bakın.

> Dikkat: Bir SQL ifadesini birleştirmek için `fmt.Sprintf` gibi dizge biçimlendirme fonksiyonlarını kullanmayın! SQL enjeksiyonu riski oluşturabilirsiniz. Daha fazlası için **Avoiding SQL injection risk** konusuna bakın.

---

## 🎯 Tek bir satır için sorgulama

`QueryRow`, tek bir veritabanı satırını en fazla bir tane olacak şekilde getirir; örneğin benzersiz bir ID üzerinden veri aramak istediğinizde. Sorgu birden fazla satır döndürürse, `Scan` metodu ilk satır dışındakilerin hepsini atar.

`QueryRowContext`, `QueryRow` gibi çalışır; ancak bir `context.Context` argümanı alır. Daha fazlası için **Canceling in-progress operations** konusuna bakın.

Aşağıdaki örnek, bir satın almayı destekleyecek yeterli envanter olup olmadığını anlamak için bir sorgu kullanır. SQL ifadesi, yeterliyse `true`, değilse `false` döndürür. `Row.Scan`, boolean dönüş değerini bir işaretçi (*pointer*) aracılığıyla `enough` değişkenine kopyalar.

```go
func canPurchase(id int, quantity int) (bool, error) {
    var enough bool
    // Tek bir satıra dayalı olarak bir değer sorgulayın.
    if err := db.QueryRow("SELECT (quantity >= ?) from album where id = ?",
        quantity, id).Scan(&enough); err != nil {
        if err == sql.ErrNoRows {
            return false, fmt.Errorf("canPurchase %d: unknown album", id)
        }
        return false, fmt.Errorf("canPurchase %d: %v", id, err)
    }
    return enough, nil
}
```

> Not: Hazırlanmış ifadelerdeki (*prepared statements*) parametre yer tutucuları, kullandığınız DBMS ve sürücüye göre değişir. Örneğin Postgres için `pq` sürücüsü, `?` yerine `$1` gibi bir yer tutucu gerektirir.

### ⚠️ Hataları ele alma

`QueryRow` kendi başına hata döndürmez. Bunun yerine `Scan`, arama (*lookup*) ve tarama (*scan*) işlemlerinin birleşiminden kaynaklanan her hatayı raporlar. Sorgu hiç satır bulmazsa `sql.ErrNoRows` döndürür.

### 🧩 Tek satır döndürmeye yönelik fonksiyonlar

| Fonksiyon              | Açıklama                                                                                                   |                                                                                                                                                 |
| ---------------------- | ---------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| `DB.QueryRow`          | `DB.QueryRowContext`                                                                                       | Tek satırlık bir sorguyu bağımsız olarak çalıştırır.                                                                                            |
| `Tx.QueryRow`          | `Tx.QueryRowContext`                                                                                       | Daha büyük bir transaction içinde tek satırlık bir sorgu çalıştırır. Daha fazlası için **Executing transactions** konusuna bakın.               |
| `Stmt.QueryRow`        | `Stmt.QueryRowContext`                                                                                     | Daha önce hazırlanmış bir ifadeyi kullanarak tek satırlık bir sorgu çalıştırır. Daha fazlası için **Using prepared statements** konusuna bakın. |
| `Conn.QueryRowContext` | Rezerve edilmiş bağlantılarla kullanım içindir. Daha fazlası için **Managing connections** konusuna bakın. |                                                                                                                                                 |

---

## 🧾 Birden çok satır için sorgulama

Birden çok satır için `Query` veya `QueryContext` kullanarak sorgu yapabilirsiniz; bunlar sorgu sonuçlarını temsil eden bir `Rows` döndürür. Kodunuz, dönen satırlar üzerinde `Rows.Next` kullanarak yineleme yapar. Her yinelemede `Scan`, sütun değerlerini değişkenlere kopyalar.

`QueryContext`, `Query` gibi çalışır; ancak bir `context.Context` argümanı alır. Daha fazlası için **Canceling in-progress operations** konusuna bakın.

Aşağıdaki örnek, belirli bir sanatçıya ait albümleri döndürmek için bir sorgu çalıştırır. Albümler bir `sql.Rows` içinde döndürülür. Kod, sütun değerlerini işaretçilerle temsil edilen değişkenlere kopyalamak için `Rows.Scan` kullanır.

```go
func albumsByArtist(artist string) ([]Album, error) {
    rows, err := db.Query("SELECT * FROM album WHERE artist = ?", artist)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    // Dönen satırlardan gelen veriyi tutacak bir album slice'ı.
    var albums []Album

    // Satırlar üzerinde dolaşın, sütun verisini struct alanlarına atamak için Scan kullanın.
    for rows.Next() {
        var alb Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist,
            &alb.Price, &alb.Quantity); err != nil {
            return albums, err
        }
        albums = append(albums, alb)
    }
    if err = rows.Err(); err != nil {
        return albums, err
    }
    return albums, nil
}
```

`rows.Close` için ertelenmiş (*deferred*) çağrıya dikkat edin. Bu, fonksiyon nasıl dönerse dönsün satırların tuttuğu kaynakları serbest bırakır. Satırların tamamı boyunca döngü kurmak onu örtük olarak kapatır; ancak `rows`’un ne olursa olsun kapatıldığından emin olmak için `defer` kullanmak daha iyidir.

> Not: Hazırlanmış ifadelerdeki (*prepared statements*) parametre yer tutucuları, kullandığınız DBMS ve sürücüye göre değişir. Örneğin Postgres için `pq` sürücüsü, `?` yerine `$1` gibi bir yer tutucu gerektirir.

### ⚠️ Hataları ele alma

Sorgu sonuçları üzerinde döngü kurduktan sonra `sql.Rows` üzerinden hata kontrolü yaptığınızdan emin olun. Sorgu başarısız olduysa, kodunuz bunu bu şekilde öğrenir.

### 🧩 Birden çok satır döndürmeye yönelik fonksiyonlar

| Fonksiyon           | Açıklama                                                                                                   |                                                                                                                                |
| ------------------- | ---------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------ |
| `DB.Query`          | `DB.QueryContext`                                                                                          | Bir sorguyu bağımsız olarak çalıştırır.                                                                                        |
| `Tx.Query`          | `Tx.QueryContext`                                                                                          | Bir transaction içinde bir sorgu çalıştırır. Daha fazlası için **Executing transactions** konusuna bakın.                      |
| `Stmt.Query`        | `Stmt.QueryContext`                                                                                        | Daha önce hazırlanmış bir ifadeyi kullanarak sorgu çalıştırır. Daha fazlası için **Using prepared statements** konusuna bakın. |
| `Conn.QueryContext` | Rezerve edilmiş bağlantılarla kullanım içindir. Daha fazlası için **Managing connections** konusuna bakın. |                                                                                                                                |

---

## 🕳️ Null olabilir sütun değerlerini ele alma

*database/sql* paketi, bir sütunun değeri null olabileceğinde `Scan` fonksiyonu için argüman olarak kullanabileceğiniz birkaç özel tür sağlar. Her birinde, değerin null olmadığını raporlayan bir `Valid` alanı ve eğer öyleyse değeri tutan bir alan bulunur.

Aşağıdaki örnekteki kod, bir müşteri adını sorgular. `name` değeri null ise, kod uygulamada kullanılmak üzere başka bir değeri yerine koyar.

```go
var s sql.NullString
err := db.QueryRow("SELECT name FROM customer WHERE id = ?", id).Scan(&s)
if err != nil {
    log.Fatal(err)
}

// Müşteri adını bulun; yoksa yer tutucu kullanın.
name := "Valued Customer"
if s.Valid {
    name = s.String
}
```

Her bir tür hakkında daha fazla bilgi için `sql` paket referansına bakın:

* `NullBool`
* `NullFloat64`
* `NullInt32`
* `NullInt64`
* `NullString`
* `NullTime`

---

## 🧱 Sütunlardan veri alma

Bir sorgunun döndürdüğü satırlar üzerinde döngü kurarken, `Rows.Scan` referansında açıklandığı gibi bir satırın sütun değerlerini Go değerlerine kopyalamak için `Scan` kullanırsınız.

Tüm sürücülerin desteklediği temel bir veri dönüşümü kümesi vardır; örneğin SQL `INT`’in Go `int`’e dönüştürülmesi gibi. Bazı sürücüler bu dönüşüm kümesini genişletir; ayrıntılar için her bir sürücünün belgelerine bakın.

Beklediğiniz gibi `Scan`, sütun türlerinden benzer Go türlerine dönüştürür. Örneğin `Scan`, SQL `CHAR`, `VARCHAR` ve `TEXT` türlerini Go `string` türüne dönüştürür. Bununla birlikte `Scan`, sütun değeri için uygun olan başka bir Go türüne de dönüşüm yapabilir. Örneğin sütun, her zaman bir sayı içerecek bir `VARCHAR` ise, değeri almak için `int` gibi sayısal bir Go türü belirtebilir ve `Scan` bunu sizin için `strconv.Atoi` kullanarak dönüştürür.

`Scan` fonksiyonu tarafından yapılan dönüşümler hakkında daha fazla ayrıntı için **Rows.Scan** referansına bakın.

---

## 🧾 Birden çok sonuç kümesini ele alma

Veritabanı işleminiz birden çok sonuç kümesi döndürebilecekse, bunları `Rows.NextResultSet` kullanarak alabilirsiniz. Bu, örneğin birden fazla tabloyu ayrı ayrı sorgulayan ve her biri için bir sonuç kümesi döndüren SQL gönderdiğinizde faydalı olabilir.

`Rows.NextResultSet`, bir sonraki sonuç kümesini hazırlar; böylece `Rows.Next` çağrısı o sonraki kümedeki ilk satırı getirir. Ayrıca, gerçekten bir sonraki sonuç kümesinin olup olmadığını belirten bir boolean döndürür.

Aşağıdaki örnekteki kod, iki SQL ifadesini yürütmek için `DB.Query` kullanır. İlk sonuç kümesi, prosedürdeki ilk sorgudan gelir ve `album` tablosundaki tüm satırları getirir. Sonraki sonuç kümesi ise ikinci sorgudan gelir ve `song` tablosundan satırları getirir.

```go
rows, err := db.Query("SELECT * from album; SELECT * from song;")
if err != nil {
    log.Fatal(err)
}
defer rows.Close()

// İlk sonuç kümesi üzerinde döngü kurun.
for rows.Next() {
    // Sonuç kümesini ele alın.
}

// Sonraki sonuç kümesine ilerleyin.
rows.NextResultSet()

// İkinci sonuç kümesi üzerinde döngü kurun.
for rows.Next() {
    // İkinci kümeyi ele alın.
}

// Herhangi bir sonuç kümesinde hata olup olmadığını kontrol edin.
if err := rows.Err(); err != nil {
    log.Fatal(err)
}
```

