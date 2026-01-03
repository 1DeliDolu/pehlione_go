
## 🗄️ MySQL Veritabanı

## 🚀 Giriş

Bir noktada web uygulamanızın bir veritabanından veri saklaması ve geri alması gerekir. Bu durum, dinamik içerikle uğraştığınızda, kullanıcıların veri girmesi için formlar sunduğunuzda veya kullanıcılarınızın kimlik doğrulaması için giriş ve parola bilgilerini sakladığınızda neredeyse her zaman geçerlidir. Bu amaçla veritabanları vardır.

Veritabanları her tür ve biçimde gelir. Web genelinde yaygın olarak kullanılan veritabanlarından biri MySQL veritabanıdır. Uzun zamandır vardır ve konumunu ve kararlılığını sayısız kez kanıtlamıştır.

Bu örnekte, Go’da veritabanı erişiminin temellerine ineceğiz, veritabanı tabloları oluşturacağız, veri saklayacağız ve tekrar geri alacağız.

---

## 📦 go-sql-driver/mysql Paketini Kurma

Go programlama dili, her türden SQL veritabanını sorgulamak için `database/sql` adlı kullanışlı bir paketle gelir. Bu, yaygın SQL özelliklerini sizin kullanmanız için tek bir API altında soyutladığı için faydalıdır. Ancak Go’nun dahil etmediği şey veritabanı *driver*’larıdır. Go’da veritabanı driver’ı, belirli bir veritabanının (bizim durumumuzda MySQL) düşük seviye detaylarını uygulayan bir pakettir. Muhtemelen tahmin ettiğiniz gibi, bu ileriye dönük uyumluluğu korumak için faydalıdır. Çünkü tüm Go paketleri yazılırken, yazarlar gelecekte ortaya çıkacak her bir veritabanını öngöremez; piyasadaki her olası veritabanını desteklemek ise büyük bir bakım yükü olurdu.

MySQL veritabanı driver’ını kurmak için terminalinize gidin ve şunu çalıştırın:

```bash
go get -u github.com/go-sql-driver/mysql
```

---

## 🔌 MySQL Veritabanına Bağlanma

Gerekli paketleri kurduktan sonra kontrol etmemiz gereken ilk şey, MySQL veritabanımıza başarıyla bağlanıp bağlanamadığımızdır. Eğer halihazırda çalışan bir MySQL veritabanı sunucunuz yoksa, Docker ile kolayca yeni bir instance başlatabilirsiniz. Docker MySQL imajı için resmi dokümanlar burada: [https://hub.docker.com/_/mysql](https://hub.docker.com/_/mysql)

Veritabanımıza bağlanabildiğimizi kontrol etmek için `database/sql` ve `go-sql-driver/mysql` paketlerini içe aktarın ve şu şekilde bir bağlantı açın:

```go
import "database/sql"
import _ "go-sql-driver/mysql"
```

```go
// Veritabanı bağlantısını yapılandırın (her zaman hataları kontrol edin)
db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")
```

```go
// Her şeyin doğru çalıştığını görmek için veritabanına ilk bağlantıyı başlatın.
// Hatanın kontrol edildiğinden emin olun.
err := db.Ping()
```

