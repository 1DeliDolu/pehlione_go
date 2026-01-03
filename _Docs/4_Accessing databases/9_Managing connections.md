
# 🔌 Bağlantıları Yönetme

## 📚 İçindekiler

* Bağlantı havuzu özelliklerini ayarlama
* Ayrılmış (*dedicated*) bağlantıları kullanma

Programların çok büyük bir çoğunluğu için `sql.DB` bağlantı havuzu varsayılanlarını ayarlamanız gerekmez. Ancak bazı ileri seviye programlarda, bağlantı havuzu parametrelerini ayarlamanız veya bağlantılarla açıkça çalışmanız gerekebilir. Bu konu bunu açıklar.

`sql.DB` veritabanı tanıtıcısı (*database handle*), birden fazla goroutine tarafından eşzamanlı kullanım için güvenlidir (yani bu tanıtıcı, diğer dillerin “thread-safe” dediği şeye karşılık gelir). Bazı diğer veritabanı erişim kütüphaneleri, aynı anda yalnızca tek bir işlem için kullanılabilen bağlantılar üzerine kuruludur. Bu boşluğu kapatmak için her `sql.DB`, alttaki veritabanına yönelik etkin (*active*) bağlantıların bir havuzunu yönetir ve Go programınızdaki paralellik (*parallelism*) ihtiyacına göre gerektiğinde yenilerini oluşturur.

Bağlantı havuzu, çoğu veri erişim ihtiyacı için uygundur. Bir `sql.DB` üzerinde `Query` veya `Exec` metodu çağırdığınızda, `sql.DB` uygulaması havuzdan kullanılabilir bir bağlantı alır veya gerekirse bir tane oluşturur. Bağlantı artık gerekli olmadığında paket, bağlantıyı havuza geri döndürür. Bu, veritabanı erişimi için yüksek bir paralellik seviyesini destekler.

---

## ⚙️ Bağlantı havuzu özelliklerini ayarlama

`sql` paketinin bağlantı havuzunu nasıl yönettiğine rehberlik eden özellikleri ayarlayabilirsiniz. Bu özelliklerin etkilerine ilişkin istatistik almak için `DB.Stats` kullanın.

---

## 🔓 Açık bağlantıların azami sayısını ayarlama

`DB.SetMaxOpenConns`, açık bağlantı sayısına bir üst sınır getirir. Bu sınır aşıldığında, yeni veritabanı işlemleri mevcut bir işlemin bitmesini bekler; bu noktada `sql.DB` başka bir bağlantı oluşturur. Varsayılan olarak `sql.DB`, bir bağlantıya ihtiyaç duyulduğunda mevcut bağlantıların tamamı kullanımda ise her seferinde yeni bir bağlantı oluşturur.

Bir sınır belirlemenin, veritabanı kullanımını bir kilit (*lock*) veya semafor (*semaphore*) edinmeye benzer hale getirdiğini unutmayın; bunun sonucu olarak uygulamanız yeni bir veritabanı bağlantısı beklerken kilitlenebilir (*deadlock*).

---

## 💤 Boşta (*idle*) bağlantıların azami sayısını ayarlama

`DB.SetMaxIdleConns`, `sql.DB`’nin sürdürdüğü boşta bağlantıların azami sayısı üzerindeki sınırı değiştirir.

Belirli bir veritabanı bağlantısı üzerindeki bir SQL işlemi bittiğinde, bağlantı tipik olarak hemen kapatılmaz: uygulamanın yakında yeniden bir bağlantıya ihtiyacı olabilir ve açık bağlantıyı tutmak, bir sonraki işlem için veritabanına yeniden bağlanma ihtiyacını ortadan kaldırır. Varsayılan olarak bir `sql.DB`, herhangi bir anda iki boşta bağlantı tutar. Sınırı artırmak, önemli paralelliğe sahip programlarda sık yeniden bağlanmaları önleyebilir.

---

## ⏳ Bir bağlantının ne kadar süre boşta kalabileceğini ayarlama

`DB.SetConnMaxIdleTime`, bir bağlantının kapatılmadan önce boşta kalabileceği azami süreyi belirler. Bu, `sql.DB`’nin belirtilen süreden daha uzun süre boşta kalan bağlantıları kapatmasına neden olur.

Varsayılan olarak, boşta bir bağlantı bağlantı havuzuna eklendiğinde, tekrar ihtiyaç duyulana kadar orada kalır. Paralel etkinlik patlamaları (*bursts*) sırasında izin verilen boşta bağlantı sayısını artırmak için `DB.SetMaxIdleConns` kullanırken, ek olarak `DB.SetConnMaxIdleTime` kullanmak; sistem sakinleştiğinde bu bağlantıların daha sonra serbest bırakılmasını sağlayabilir.

---

## 🕰️ Bağlantıların azami yaşam süresini ayarlama

`DB.SetConnMaxLifetime` kullanımı, bir bağlantının kapatılmadan önce ne kadar süre açık tutulabileceğinin azami süresini belirler.

Varsayılan olarak bir bağlantı, yukarıda açıklanan sınırlar dahilinde, keyfi olarak uzun bir süre kullanılabilir ve yeniden kullanılabilir. Yük dengeli (*load-balanced*) bir veritabanı sunucusu kullananlar gibi bazı sistemlerde, uygulamanın yeniden bağlanmadan belirli bir bağlantıyı çok uzun süre kullanmamasını sağlamak yararlı olabilir.

---

## 🧷 Ayrılmış (*dedicated*) bağlantıları kullanma

*database/sql* paketi, bir veritabanının belirli bir bağlantı üzerinde yürütülen bir işlem dizisine örtük anlam (*implicit meaning*) atayabileceği durumlarda kullanabileceğiniz fonksiyonlar içerir.

En yaygın örnek transaction’lardır; bunlar tipik olarak bir `BEGIN` komutuyla başlar, bir `COMMIT` veya `ROLLBACK` komutuyla biter ve bu komutlar arasında bağlantı üzerinde yürütülen tüm komutları genel transaction’ın parçası olarak içerir. Bu kullanım durumu için `sql` paketinin transaction desteğini kullanın. **Executing transactions** konusuna bakın.

Bunun dışındaki kullanım durumlarında, bireysel işlemler dizisinin tamamının aynı bağlantı üzerinde yürütülmesi gerektiğinde, `sql` paketi ayrılmış bağlantılar sağlar. `DB.Conn`, ayrılmış bir bağlantı olan `sql.Conn` elde eder. `sql.Conn`, `DB` üzerindeki eşdeğer metodlar gibi davranan ancak yalnızca ayrılmış bağlantıyı kullanan `BeginTx`, `ExecContext`, `PingContext`, `PrepareContext`, `QueryContext` ve `QueryRowContext` metodlarına sahiptir. Ayrılmış bağlantı ile işiniz bittiğinde, kodunuz `Conn.Close` kullanarak onu serbest bırakmalıdır.

