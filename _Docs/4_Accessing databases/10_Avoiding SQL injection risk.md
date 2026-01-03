
# 🛡️ SQL Enjeksiyonu Riskinden Kaçınma

SQL enjeksiyonu riskinden, SQL parametre değerlerini `sql` paketi fonksiyonlarının argümanları olarak sağlayarak kaçınabilirsiniz. `sql` paketindeki birçok fonksiyon, SQL ifadesi için bir parametre ve bu ifadenin parametrelerinde kullanılacak değerler için ayrıca parametreler sağlar (diğerleri hazırlanmış bir ifade için bir parametre ve parametreler sağlar).

Aşağıdaki örnekteki kod, `id` parametresi için bir yer tutucu olarak `?` sembolünü kullanır; `id` ise fonksiyon argümanı olarak sağlanır:

```go
// Parametreli bir SQL ifadesini çalıştırmak için doğru format.
rows, err := db.Query("SELECT * FROM user WHERE id = ?", id)
```

Veritabanı işlemleri yapan `sql` paketi fonksiyonları, sağladığınız argümanlardan hazırlanmış ifadeler (*prepared statements*) oluşturur. Çalışma zamanında `sql` paketi SQL ifadesini hazırlanmış bir ifadeye dönüştürür ve parametreyi ayrı olarak onunla birlikte gönderir.

> Not: Parametre yer tutucuları, kullandığınız DBMS ve sürücüye göre değişir. Örneğin Postgres için `pq` sürücüsü, `?` yerine `$1` gibi bir yer tutucu biçimini kabul eder.

SQL ifadesini parametreler dahil olacak şekilde bir dizge olarak birleştirmek için `fmt` paketinden bir fonksiyon kullanmaya yönelebilirsiniz — örneğin şöyle:

```go
// GÜVENLİK RİSKİ!
rows, err := db.Query(fmt.Sprintf("SELECT * FROM user WHERE id = %s", id))
```

Bu güvenli değildir! Bunu yaptığınızda Go, `%s` biçim belirtecini parametre değeriyle değiştirerek tüm SQL ifadesini birleştirir ve tam ifadeyi DBMS’e göndermeden önce oluşturur. Bu, SQL enjeksiyonu riski doğurur; çünkü kodun çağıranı `id` argümanı olarak beklenmedik bir SQL parçacığı gönderebilir. Bu parça, SQL ifadesini uygulamanız için tehlikeli olabilecek öngörülemez şekillerde tamamlayabilir.

Örneğin belirli bir `%s` değeri geçirerek, veritabanınızdaki tüm kullanıcı kayıtlarını döndürebilecek aşağıdakine benzer bir sonuç elde edebilirsiniz:

```sql
SELECT * FROM user WHERE id = 1 OR 1=1;
```

