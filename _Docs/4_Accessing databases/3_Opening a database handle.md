
# 🧰 Bir Veritabanı Tanıtıcısı Açma

## 📚 İçindekiler

* Bir veritabanı sürücüsünü bulma ve içe aktarma
* Bir veritabanı tanıtıcısı açma
* Bağlantıyı doğrulama
* Veritabanı kimlik bilgilerini saklama
* Kaynakları serbest bırakma

*database/sql* paketi, bağlantıları yönetme ihtiyacınızı azaltarak veritabanı erişimini basitleştirir. Birçok veri erişim API’sinin aksine, *database/sql* ile bir bağlantıyı açıkça açıp, iş yapıp, sonra bağlantıyı kapatmazsınız. Bunun yerine kodunuz, bir bağlantı havuzunu (*connection pool*) temsil eden bir veritabanı tanıtıcısı (*database handle*) açar; ardından tanıtıcıyı kullanarak veri erişim işlemlerini yürütür ve yalnızca ihtiyaç duyulduğunda — örneğin getirilen satırların veya hazırlanmış bir ifadenin (*prepared statement*) tuttuğu kaynakları serbest bırakmak için — `Close` metodunu çağırır.

Başka bir deyişle, bağlantıları sizin kodunuz adına açıp kapatan şey, `sql.DB` ile temsil edilen veritabanı tanıtıcısıdır. Kodunuz tanıtıcıyı veritabanı işlemlerini yürütmek için kullandıkça, bu işlemler veritabanına eşzamanlı (*concurrent*) erişime sahip olur. Daha fazlası için **Managing connections** konusuna bakın.

> Not: Bir veritabanı bağlantısını ayrıca rezerve edebilirsiniz. Daha fazla bilgi için **Using dedicated connections** konusuna bakın.

*database/sql* paketinde bulunan API’lere ek olarak, Go topluluğu en yaygın (ve daha az yaygın pek çok) veritabanı yönetim sistemi (DBMS) için sürücüler geliştirmiştir.

Bir veritabanı tanıtıcısı açarken, şu üst seviye adımları izlersiniz:

* Bir sürücü bulun.
  Bir sürücü, Go kodunuz ile veritabanı arasında istekleri ve yanıtları çevirir. Daha fazlası için **Locating and importing a database driver** konusuna bakın.

* Bir veritabanı tanıtıcısı açın.
  Sürücüyü içe aktardıktan sonra, belirli bir veritabanı için bir tanıtıcı açabilirsiniz. Daha fazlası için **Opening a database handle** konusuna bakın.

* Bağlantıyı doğrulayın.
  Bir veritabanı tanıtıcısı açtıktan sonra, kodunuz bir bağlantının kullanılabilir olduğunu kontrol edebilir. Daha fazlası için **Confirming a connection** konusuna bakın.

Kodunuz genellikle veritabanı bağlantılarını açıkça açıp kapatmaz — bunu veritabanı tanıtıcısı yapar. Ancak kodunuz, sorgu sonuçlarını içeren bir `sql.Rows` gibi süreç boyunca elde ettiği kaynakları serbest bırakmalıdır. Daha fazlası için **Freeing resources** konusuna bakın.

---

## 🔎 Bir veritabanı sürücüsünü bulma ve içe aktarma

Kullandığınız DBMS’i destekleyen bir veritabanı sürücüsüne ihtiyacınız olacak. Veritabanınız için bir sürücü bulmak üzere **SQLDrivers** sayfasına bakın.

Sürücüyü kodunuza kullanılabilir hale getirmek için, onu başka bir Go paketi gibi içe aktarırsınız. Örneğin:

```go
import "github.com/go-sql-driver/mysql"
```

Sürücü paketinden doğrudan hiçbir fonksiyon çağırmıyorsanız — örneğin `sql` paketi tarafından örtük (*implicitly*) kullanılıyorsa — içe aktarma yolunun başına alt çizgi koyan bir *blank import* kullanmanız gerekir:

```go
import _ "github.com/go-sql-driver/mysql"
```

> Not: En iyi uygulama olarak, veritabanı işlemleri için sürücünün kendi API’sini kullanmaktan kaçının. Bunun yerine *database/sql* paketindeki fonksiyonları kullanın. Bu, kodunuzun DBMS’e gevşek bağlı (*loosely coupled*) kalmasına yardımcı olur ve gerekirse farklı bir DBMS’e geçişi kolaylaştırır.

---

## 🧩 Bir veritabanı tanıtıcısı açma

Bir `sql.DB` veritabanı tanıtıcısı, veritabanından okuma ve veritabanına yazma yeteneği sağlar; bunu tekil işlemler olarak ya da bir transaction içinde yapabilirsiniz.

Bir veritabanı tanıtıcısını, ya bir bağlantı dizesi alan `sql.Open` ile ya da bir `driver.Connector` alan `sql.OpenDB` ile alabilirsiniz. Her ikisi de bir `*sql.DB` döndürür.

> Not: Veritabanı kimlik bilgilerinizi Go kaynak kodunuzun dışında tuttuğunuzdan emin olun. Daha fazlası için **Storing database credentials** konusuna bakın.

### 🔗 Bağlantı dizesi ile açma

Bir bağlantı dizesi ile bağlanmak istediğinizde `sql.Open` fonksiyonunu kullanın. Dizgenin biçimi, kullandığınız sürücüye göre değişir.

MySQL için bir örnek:

```go
db, err = sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/jazzrecords")
if err != nil {
    log.Fatal(err)
}
```

Ancak bağlantı özelliklerini daha yapılandırılmış bir şekilde yakalamanın kodu daha okunur hale getirdiğini çoğu zaman görürsünüz. Ayrıntılar sürücüye göre değişir.

Örneğin, MySQL sürücüsünün `Config` yapısını özellikleri belirtmek için ve `FormatDSN` metodunu bir bağlantı dizesi oluşturmak için kullanarak, önceki örneği aşağıdaki ile değiştirebilirsiniz.

```go
// Bağlantı özelliklerini belirtin.
cfg := mysql.NewConfig()
cfg.User = username
cfg.Passwd = password
cfg.Net = "tcp"
cfg.Addr = "127.0.0.1:3306"
cfg.DBName = "jazzrecords"

// Bir veritabanı tanıtıcısı alın.
db, err = sql.Open("mysql", cfg.FormatDSN())
if err != nil {
    log.Fatal(err)
}
```

### 🧷 Connector ile açma

Bağlantı dizesinde kullanılamayan, sürücüye özgü bağlantı özelliklerinden yararlanmak istediğinizde `sql.OpenDB` fonksiyonunu kullanın. Her sürücü kendi bağlantı özellikleri kümesini destekler ve çoğu zaman DBMS’e özgü bağlantı isteğini özelleştirmek için yollar sunar.

Önceki `sql.Open` örneğini `sql.OpenDB` kullanacak şekilde uyarlayarak, aşağıdaki gibi bir kodla tanıtıcı oluşturabilirsiniz:

```go
// Bağlantı özelliklerini belirtin.
cfg := mysql.NewConfig()
cfg.User = username
cfg.Passwd = password
cfg.Net = "tcp"
cfg.Addr = "127.0.0.1:3306"
cfg.DBName = "jazzrecords"

// Sürücüye özgü bir connector alın.
connector, err := mysql.NewConnector(&cfg)
if err != nil {
    log.Fatal(err)
}

// Bir veritabanı tanıtıcısı alın.
db = sql.OpenDB(connector)
```

### ⚠️ Hataları ele alma

Kodunuz, `sql.Open` gibi bir tanıtıcı oluşturma girişiminden dönen hatayı kontrol etmelidir. Bu bir bağlantı hatası olmayacaktır. Bunun yerine, `sql.Open` tanıtıcıyı başlatamadığında hata alırsınız. Bu, örneğin belirttiğiniz DSN’yi ayrıştıramıyorsa (*parse edemiyorsa*) meydana gelebilir.

---

## ✅ Bağlantıyı doğrulama

Bir veritabanı tanıtıcısı açtığınızda, `sql` paketi veritabanı bağlantısını hemen oluşturmayabilir. Bunun yerine, kodunuz ihtiyaç duyduğunda bağlantıyı oluşturabilir. Veritabanını hemen kullanmayacaksanız ve bir bağlantı kurulabileceğini doğrulamak istiyorsanız, `Ping` veya `PingContext` çağırın.

Aşağıdaki örnekteki kod, bağlantıyı doğrulamak için veritabanına ping atar.

```go
db, err = sql.Open("mysql", connString)

// Başarılı bir bağlantıyı doğrulayın.
if err := db.Ping(); err != nil {
    log.Fatal(err)
}
```

---

## 🔐 Veritabanı kimlik bilgilerini saklama

Veritabanı kimlik bilgilerini Go kaynak kodunuzda saklamaktan kaçının; bu, veritabanınızın içeriğini başkalarına açığa çıkarabilir. Bunun yerine, bunları kodunuzun dışında ama kodunuz tarafından erişilebilir bir yerde saklamanın bir yolunu bulun. Örneğin, kimlik bilgilerini depolayan ve kodunuzun DBMS’e kimlik doğrulaması için kullanabileceği bir API sağlayan bir *secret keeper* uygulaması düşünebilirsiniz.

Popüler yaklaşımlardan biri, program başlamadan önce sırları ortamda (*environment*) saklamaktır; belki bir *secret manager* üzerinden yüklenir ve ardından Go programınız bunları `os.Getenv` ile okuyabilir:

```go
username := os.Getenv("DB_USER")
password := os.Getenv("DB_PASS")
```

Bu yaklaşım ayrıca, yerel test için ortam değişkenlerini kendiniz ayarlamanıza da olanak tanır.

---

## 🧹 Kaynakları serbest bırakma

*database/sql* paketi ile bağlantıları açıkça yönetmeseniz veya kapatmasanız da, kodunuz artık ihtiyaç duyulmadığında elde ettiği kaynakları serbest bırakmalıdır. Bunlara, bir sorgudan dönen veriyi temsil eden `sql.Rows`’un veya hazırlanmış bir ifadeyi temsil eden `sql.Stmt`’nin tuttuğu kaynaklar dahil olabilir.

Tipik olarak, kaynakları `Close` fonksiyonuna yapılan bir çağrıyı `defer` ederek kapatırsınız; böylece kaynaklar, saran (*enclosing*) fonksiyon çıkmadan önce serbest bırakılır.

Aşağıdaki örnek, `sql.Rows` tarafından tutulan kaynağı serbest bırakmak için `Close` çağrısını `defer` eder.

```go
rows, err := db.Query("SELECT * FROM album WHERE artist = ?", artist)
if err != nil {
    log.Fatal(err)
}
defer rows.Close()

// Dönen satırlar üzerinde döngü kurun.
```

