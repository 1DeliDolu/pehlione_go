
# 🗄️ Eğitim: İlişkisel Bir Veritabanına Erişim 

## 🧭 İçindekiler

* Önkoşullar
* Kodunuz için bir klasör oluşturun
* Bir veritabanı kurun
* Bir veritabanı sürücüsü bulun ve içe aktarın
* Bir veritabanı tanıtıcısı alın ve bağlanın
* Birden çok satır için sorgu çalıştırın
* Tek bir satır için sorgu çalıştırın
* Veri ekleyin
* Sonuç
* Tamamlanmış kod

Bu eğitim, Go’da ilişkisel bir veritabanına erişmenin temellerini ve standart kütüphanedeki *database/sql* paketini tanıtır.

Bu eğitimden en iyi şekilde yararlanmak için Go ve araçlarına dair temel bir aşinalığınız olması faydalıdır. Go ile ilk kez karşılaşıyorsanız, hızlı bir giriş için **Tutorial: Get started with Go** eğitimine bakın.

Kullanacağınız *database/sql* paketi; veritabanlarına bağlanma, işlemleri (*transactions*) yürütme, devam eden bir işlemi iptal etme ve daha fazlası için türler ve fonksiyonlar içerir. Paketin kullanımı hakkında daha fazla ayrıntı için **Accessing databases** konusuna bakın.

Bu eğitimde bir veritabanı oluşturacak, ardından bu veritabanına erişen kodu yazacaksınız. Örnek projeniz, vintage caz plak kayıtlarına ilişkin verilerin bulunduğu bir depo (*repository*) olacaktır.

Bu eğitimde aşağıdaki bölümlerden geçeceksiniz:

* Kodunuz için bir klasör oluşturun.
* Bir veritabanı kurun.
* Veritabanı sürücüsünü içe aktarın.
* Bir veritabanı tanıtıcısı alın ve bağlanın.
* Birden çok satır için sorgu çalıştırın.
* Tek bir satır için sorgu çalıştırın.
* Veri ekleyin.

> Not: Diğer eğitimler için **Tutorials** bölümüne bakın.

---

## ✅ Önkoşullar

* MySQL ilişkisel veritabanı yönetim sistemi (DBMS) kurulu olmalı.
* Go kurulu olmalı. Kurulum talimatları için **Installing Go** bölümüne bakın.
* Kod düzenlemek için bir araç. Elinizdeki herhangi bir metin düzenleyici iş görür.
* Bir komut satırı terminali. Go; Linux ve Mac’te herhangi bir terminalde, Windows’ta PowerShell veya cmd üzerinde iyi çalışır.

---

## 📁 Kodunuz için bir klasör oluşturun

Başlamak için yazacağınız kod için bir klasör oluşturun.

Bir komut istemi açın ve ana dizininize geçin.

Linux veya Mac’te:

```bash
$ cd
```

Windows’ta:

```bat
C:\> cd %HOMEPATH%
```

Eğitimin geri kalanında, istemi `$` olarak göstereceğiz. Kullandığımız komutlar Windows’ta da çalışacaktır.

Komut isteminde, kodunuz için **data-access** adlı bir dizin oluşturun.

```bash
$ mkdir data-access
$ cd data-access
```

Bu eğitim boyunca ekleyeceğiniz bağımlılıkları yönetebileceğiniz bir modül oluşturun.

Yeni kodunuzun modül yolunu vererek `go mod init` komutunu çalıştırın.

```bash
$ go mod init example/data-access
```

```text
go: creating new go.mod: module example/data-access
```

Bu komut, eklediğiniz bağımlılıkların izlenmesi için listeleneceği bir `go.mod` dosyası oluşturur. Daha fazlası için **Managing dependencies** konusuna bakın.

> Not: Gerçek geliştirmede, kendi ihtiyaçlarınıza daha uygun, daha spesifik bir modül yolu belirlersiniz. Daha fazlası için **Managing dependencies** konusuna bakın.

Sonraki adımda bir veritabanı oluşturacaksınız.

---

## 🧱 Bir veritabanı kurun

Bu adımda üzerinde çalışacağınız veritabanını oluşturacaksınız. Veritabanını ve tabloyu oluşturmak ve veri eklemek için DBMS’in kendi CLI’ını kullanacaksınız.

Vinyl üzerindeki vintage caz kayıtlarına ilişkin veriler içeren bir veritabanı oluşturacaksınız.

Buradaki kod MySQL CLI’ını kullanır; ancak çoğu DBMS’in benzer özelliklere sahip kendi CLI’ı vardır.

Yeni bir komut istemi açın.

Komut satırında, aşağıdaki MySQL örneğinde olduğu gibi DBMS’inize giriş yapın.

```bash
$ mysql -u root -p
```

```text
Enter password:
```

```text
mysql>
```

`mysql` komut isteminde bir veritabanı oluşturun.

```sql
mysql> create database recordings;
```

Az önce oluşturduğunuz veritabanına geçin; böylece tablolar ekleyebilirsiniz.

```sql
mysql> use recordings;
```

```text
Database changed
```

Metin düzenleyicinizde, **data-access** klasörünün içinde tablo eklemek için SQL betiğini tutacak **create-tables.sql** adlı bir dosya oluşturun.

Dosyanın içine aşağıdaki SQL kodunu yapıştırın, ardından kaydedin.

```sql
DROP TABLE IF EXISTS album;
CREATE TABLE album (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  artist     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO album
  (title, artist, price)
VALUES
  ('Blue Train', 'John Coltrane', 56.99),
  ('Giant Steps', 'John Coltrane', 63.99),
  ('Jeru', 'Gerry Mulligan', 17.99),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98);
```

Bu SQL kodunda şunları yaparsınız:

* `album` adlı bir tabloyu silersiniz (*drop*). Bu komutu önce çalıştırmak, daha sonra tabloyla yeniden başlamak isterseniz betiği tekrar çalıştırmanızı kolaylaştırır.
* Dört sütunlu bir `album` tablosu oluşturursunuz: `title`, `artist` ve `price`. Her satırın `id` değeri DBMS tarafından otomatik olarak oluşturulur.
* Dört satır veri eklersiniz.

`mysql` komut isteminden, az önce oluşturduğunuz betiği çalıştırın.

Aşağıdaki biçimde `source` komutunu kullanacaksınız:

```sql
mysql> source /path/to/create-tables.sql
```

DBMS komut isteminizde, tabloyu verilerle birlikte başarıyla oluşturduğunuzu doğrulamak için bir `SELECT` ifadesi kullanın.

```sql
mysql> select * from album;
```

```text
+----+---------------+----------------+-------+
| id | title         | artist         | price |
+----+---------------+----------------+-------+
|  1 | Blue Train    | John Coltrane  | 56.99 |
|  2 | Giant Steps   | John Coltrane  | 63.99 |
|  3 | Jeru          | Gerry Mulligan | 17.99 |
|  4 | Sarah Vaughan | Sarah Vaughan  | 34.98 |
+----+---------------+----------------+-------+
4 rows in set (0.00 sec)
```

Sırada, sorgulayabilmek için bağlanacak Go kodunu yazmak var.

---

## 🔌 Bir veritabanı sürücüsü bulun ve içe aktarın

Artık içinde veri olan bir veritabanınız olduğuna göre, Go kodunuzu başlatın.

*database/sql* paketindeki fonksiyonlar üzerinden yaptığınız istekleri, veritabanının anlayacağı isteklere çevirecek bir veritabanı sürücüsü bulun ve içe aktarın.

Tarayıcınızda, kullanabileceğiniz bir sürücüyü belirlemek için **SQLDrivers** wiki sayfasını ziyaret edin.

Sayfadaki listeyi kullanarak hangi sürücüyü kullanacağınızı belirleyin. Bu eğitimde MySQL’e erişmek için **Go-MySQL-Driver** kullanacaksınız.

Sürücü için paket adını not edin — burada: `github.com/go-sql-driver/mysql`.

Metin düzenleyicinizi kullanarak Go kodunuzu yazacağınız bir dosya oluşturun ve daha önce oluşturduğunuz **data-access** dizininde **main.go** olarak kaydedin.

`main.go` dosyasına, sürücü paketini içe aktarmak için aşağıdaki kodu yapıştırın.

```go
package main

import "github.com/go-sql-driver/mysql"
```

Bu kodda:

* Kodunuzu bağımsız çalıştırabilmek için `main` paketine eklersiniz.
* MySQL sürücüsü `github.com/go-sql-driver/mysql` paketini içe aktarırsınız.

Sürücüyü içe aktardıktan sonra, veritabanına erişmek için kod yazmaya başlayacaksınız.

---

## 🧩 Bir veritabanı tanıtıcısı alın ve bağlanın

Şimdi, bir veritabanı tanıtıcısı (*handle*) ile veritabanına erişim sağlayan Go kodunu yazın.

Belirli bir veritabanına erişimi temsil eden bir *sql.DB* yapısına işaretçi (*pointer*) kullanacaksınız.

### ✍️ Kodu yazın

`main.go` dosyasında, az önce eklediğiniz `import` kodunun altına, veritabanı tanıtıcısı oluşturmak için aşağıdaki Go kodunu yapıştırın.

```go
var db *sql.DB

func main() {
    // Bağlantı özelliklerini yakalayın.
    cfg := mysql.NewConfig()
    cfg.User = os.Getenv("DBUSER")
    cfg.Passwd = os.Getenv("DBPASS")
    cfg.Net = "tcp"
    cfg.Addr = "127.0.0.1:3306"
    cfg.DBName = "recordings"

    // Bir veritabanı tanıtıcısı alın.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}
```

Bu kodda:

* `*sql.DB` türünde bir `db` değişkeni tanımlarsınız. Bu, veritabanı tanıtıcınızdır.
* `db` değişkenini global yapmak bu örneği basitleştirir. Üretimde, fonksiyonlara parametre geçmek veya bir struct içinde sarmalamak gibi yollarla global değişkenden kaçınırsınız.
* MySQL sürücüsünün `Config` yapısını ve türün `FormatDSN` metodunu kullanarak bağlantı özelliklerini toplar ve bunları bir bağlantı dizesi için bir *DSN* biçimine dönüştürürsünüz.
* `Config` yapısı, bağlantı dizesine kıyasla daha okunabilir bir kod sağlar.
* `sql.Open` çağrısı ile `db` değişkenini başlatır, `FormatDSN` dönüş değerini geçirirsiniz.
* `sql.Open`’dan gelen hatayı kontrol edersiniz. Örneğin bağlantı ayrıntıları doğru biçimlenmemişse başarısız olabilir.
* Kodu basitleştirmek için hatada `log.Fatal` çağırarak çalışmayı sonlandırır ve hatayı konsola yazdırırsınız. Üretimde hataları daha zarif ele almak istersiniz.
* Çalışma zamanında `sql.Open` sürücüye göre anında bağlanmayabilir. *database/sql* paketinin gerektiğinde bağlanabildiğini doğrulamak için `DB.Ping` kullanırsınız.
* Bağlantı başarısız olmuş olabileceği için `Ping` hatasını kontrol edersiniz.
* `Ping` başarılı olursa bir mesaj yazdırırsınız.

`main.go` dosyasının en üstüne, `package` bildiriminden hemen sonra, az önce yazdığınız kodu desteklemek için gereken paketleri içe aktarın.

Dosyanın üst kısmı artık şöyle görünmelidir:

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/go-sql-driver/mysql"
)
```

`main.go` dosyasını kaydedin.

### ▶️ Kodu çalıştırın

MySQL sürücü modülünü bir bağımlılık olarak izlemeye başlayın.

Mevcut dizindeki kodun bağımlılıklarını “al” anlamına gelen nokta argümanı ile `go get` kullanarak `github.com/go-sql-driver/mysql` modülünü kendi modülünüze bağımlılık olarak ekleyin.

```bash
$ go get .
```

```text
go: added filippo.io/edwards25519 v1.1.0
go: added github.com/go-sql-driver/mysql v1.8.1
```

Go, önceki adımda `import` bildirimine eklediğiniz için bu bağımlılığı indirdi. Bağımlılık izlemesi hakkında daha fazla bilgi için **Adding a dependency** konusuna bakın.

Go programının kullanması için komut isteminden `DBUSER` ve `DBPASS` ortam değişkenlerini ayarlayın.

Linux veya Mac’te:

```bash
$ export DBUSER=username
$ export DBPASS=password
```

Windows’ta:

```bat
C:\Users\you\data-access> set DBUSER=username
C:\Users\you\data-access> set DBPASS=password
```

`main.go` dosyasının bulunduğu dizinde komut satırından, “mevcut dizindeki paketi çalıştır” anlamına gelen nokta argümanı ile `go run` yazarak kodu çalıştırın.

```bash
$ go run .
```

```text
Connected!
```

Bağlanabiliyorsunuz. Sırada veri sorgulamak var.

---

## 🧾 Birden çok satır için sorgu çalıştırın

Bu bölümde, birden çok satır döndürmek üzere tasarlanmış bir SQL sorgusunu Go ile çalıştıracaksınız.

Birden çok satır döndürebilecek SQL ifadeleri için, *database/sql* paketindeki `Query` metodunu kullanır, ardından dönen satırlar üzerinde döngü kurarsınız. (Tek bir satır sorgulamayı daha sonra **Tek bir satır için sorgu çalıştırın** bölümünde öğreneceksiniz.)

### ✍️ Kodu yazın

`main.go` dosyasında, `func main` tanımının hemen üstüne aşağıdaki `Album` struct tanımını yapıştırın. Sorgudan dönen satır verisini tutmak için bunu kullanacaksınız.

```go
type Album struct {
    ID     int64
    Title  string
    Artist string
    Price  float32
}
```

`func main`’in altına, veritabanını sorgulayan aşağıdaki `albumsByArtist` fonksiyonunu yapıştırın.

```go
// albumsByArtist, belirtilen sanatçı adına sahip albümler için sorgu yapar.
func albumsByArtist(name string) ([]Album, error) {
    // Dönen satırlardan gelen veriyi tutacak bir albums slice'ı.
    var albums []Album

    rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    defer rows.Close()
    // Satırlar üzerinde dolaşın, sütun verisini struct alanlarına atamak için Scan kullanın.
    for rows.Next() {
        var alb Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
        }
        albums = append(albums, alb)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    return albums, nil
}
```

Bu kodda:

* Tanımladığınız `Album` türünde bir `albums` slice’ı tanımlarsınız. Bu, dönen satırlardaki veriyi tutar. Struct alan adları ve türleri, veritabanı sütun adları ve türleriyle eşleşir.
* Belirtilen sanatçı adına sahip albümleri sorgulamak için `DB.Query` ile bir `SELECT` ifadesi çalıştırırsınız.
* `Query`’nin ilk parametresi SQL ifadesidir. Sonrasında sıfır veya daha fazla parametre değeri geçebilirsiniz. Bunlar, SQL ifadenizdeki parametreler için değerleri belirtmenize yarar. SQL’i ve parametre değerlerini ayırarak (ör. `fmt.Sprintf` ile birleştirmek yerine) *database/sql* paketinin değerleri SQL metninden ayrı göndermesini sağlarsınız; böylece SQL enjeksiyonu riski ortadan kalkar.
* `rows` kapatmayı `defer` edersiniz; böylece fonksiyon biterken tuttuğu kaynaklar serbest bırakılır.
* Dönen satırlar üzerinde döngü kurar, `Rows.Scan` ile her satırın sütun değerlerini `Album` struct alanlarına atarsınız.
* `Scan`, sütun değerlerinin yazılacağı Go değerlerine işaretçilerin listesini alır. Burada `&` operatörüyle oluşturulan `alb` değişkeninin alanlarına işaretçiler geçirirsiniz. `Scan`, struct alanlarını güncellemek için işaretçiler üzerinden yazar.
* Döngü içinde tarama (*scan*) hatasını kontrol edersiniz.
* Döngü içinde yeni `alb` değerini `albums` slice’ına eklersiniz.
* Döngü sonrasında, `rows.Err` ile sorgunun genelinde bir hata olup olmadığını kontrol edersiniz. Sorgunun kendisi başarısız olursa sonuçların eksik olduğunu anlamanın tek yolu burada hatayı kontrol etmektir.

`main` fonksiyonunuzu `albumsByArtist` çağıracak şekilde güncelleyin.

`func main`’in sonuna aşağıdaki kodu ekleyin.

```go
albums, err := albumsByArtist("John Coltrane")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Albums found: %v\n", albums)
```

Yeni kodda artık:

* Eklediğiniz `albumsByArtist` fonksiyonunu çağırır, dönüş değerini yeni bir `albums` değişkenine atarsınız.
* Sonucu yazdırırsınız.

### ▶️ Kodu çalıştırın

`main.go` dosyasının bulunduğu dizinde komut satırından kodu çalıştırın.

```bash
$ go run .
```

```text
Connected!
Albums found: [{1 Blue Train John Coltrane 56.99} {2 Giant Steps John Coltrane 63.99}]
```

Sırada tek bir satır sorgulamak var.

---

## 🎯 Tek bir satır için sorgu çalıştırın

Bu bölümde, veritabanında tek bir satırı Go ile sorgulayacaksınız.

En fazla tek bir satır döneceğini bildiğiniz SQL ifadeleri için, `Query` döngüsü kullanmaktan daha basit olan `QueryRow`’u kullanabilirsiniz.

### ✍️ Kodu yazın

`albumsByArtist` fonksiyonunun altına aşağıdaki `albumByID` fonksiyonunu yapıştırın.

```go
// albumByID, belirtilen ID'ye sahip albüm için sorgu yapar.
func albumByID(id int64) (Album, error) {
    // Dönen satırdan gelen veriyi tutacak bir albüm.
    var alb Album

    row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
    if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
        if err == sql.ErrNoRows {
            return alb, fmt.Errorf("albumsById %d: böyle bir albüm yok", id)
        }
        return alb, fmt.Errorf("albumsById %d: %v", id, err)
    }
    return alb, nil
}
```

Bu kodda:

* Belirtilen ID’ye sahip albümü sorgulamak için `DB.QueryRow` ile bir `SELECT` ifadesi çalıştırırsınız.
* Bu çağrı bir `sql.Row` döndürür. Çağıran kodu (sizin kodunuzu) basitleştirmek için `QueryRow` hata döndürmez; bunun yerine herhangi bir sorgu hatasını (örn. `sql.ErrNoRows`) daha sonra `Rows.Scan` üzerinden döndürmek üzere düzenler.
* `Row.Scan` ile sütun değerlerini struct alanlarına kopyalarsınız.
* `Scan` hatasını kontrol edersiniz.
* Özel hata `sql.ErrNoRows`, sorgunun hiç satır döndürmediğini belirtir. Tipik olarak bu hatayı burada olduğu gibi “böyle bir albüm yok” gibi daha spesifik bir metinle değiştirmek yararlıdır.

`main` fonksiyonunu `albumByID` çağıracak şekilde güncelleyin.

`func main`’in sonuna aşağıdaki kodu ekleyin.

```go
// Sorguyu test etmek için burada ID 2'yi sabit kodlayın.
alb, err := albumByID(2)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Album found: %v\n", alb)
```

Yeni kodda artık:

* Eklediğiniz `albumByID` fonksiyonunu çağırırsınız.
* Dönen albüm ID’sini yazdırırsınız.

### ▶️ Kodu çalıştırın

`main.go` dosyasının bulunduğu dizinde komut satırından kodu çalıştırın.

```bash
$ go run .
```

```text
Connected!
Albums found: [{1 Blue Train John Coltrane 56.99} {2 Giant Steps John Coltrane 63.99}]
Album found: {2 Giant Steps John Coltrane 63.99}
```

Sırada veritabanına bir albüm eklemek var.

---

## ➕ Veri ekleyin

Bu bölümde, veritabanına yeni bir satır eklemek için bir SQL `INSERT` ifadesini Go ile çalıştıracaksınız.

Veri döndüren SQL ifadeleri ile `Query` ve `QueryRow` kullanımını gördünüz. Veri döndürmeyen SQL ifadelerini çalıştırmak için `Exec` kullanırsınız.

### ✍️ Kodu yazın

`albumByID` fonksiyonunun altına, veritabanına yeni bir albüm ekleyen aşağıdaki `addAlbum` fonksiyonunu yapıştırın, ardından `main.go` dosyasını kaydedin.

```go
// addAlbum, belirtilen albümü veritabanına ekler,
// yeni kaydın albüm ID'sini döndürür
func addAlbum(alb Album) (int64, error) {
    result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    return id, nil
}
```

Bu kodda:

* Bir `INSERT` ifadesi çalıştırmak için `DB.Exec` kullanırsınız.
* `Query` gibi, `Exec` de bir SQL ifadesi ve ardından SQL ifadesi için parametre değerleri alır.
* `INSERT` denemesinden gelen hatayı kontrol edersiniz.
* Eklenen veritabanı satırının ID’sini `Result.LastInsertId` ile alırsınız.
* ID’yi alma denemesinden gelen hatayı kontrol edersiniz.

`main` fonksiyonunu yeni `addAlbum` fonksiyonunu çağıracak şekilde güncelleyin.

`func main`’in sonuna aşağıdaki kodu ekleyin.

```go
albID, err := addAlbum(Album{
    Title:  "The Modern Sound of Betty Carter",
    Artist: "Betty Carter",
    Price:  49.99,
})
if err != nil {
    log.Fatal(err)
}
fmt.Printf("ID of added album: %v\n", albID)
```

Yeni kodda artık:

* Yeni bir albümle `addAlbum` çağırır, eklediğiniz albümün ID’sini `albID` değişkenine atarsınız.

### ▶️ Kodu çalıştırın

`main.go` dosyasının bulunduğu dizinde komut satırından kodu çalıştırın.

```bash
$ go run .
```

```text
Connected!
Albums found: [{1 Blue Train John Coltrane 56.99} {2 Giant Steps John Coltrane 63.99}]
Album found: {2 Giant Steps John Coltrane 63.99}
ID of added album: 5
```

---

## 🏁 Sonuç

Tebrikler! Go kullanarak ilişkisel bir veritabanıyla basit işlemler gerçekleştirdiniz.

Önerilen sonraki konular:

* Yalnızca değinilen konular hakkında daha fazla bilgi içeren veri erişim kılavuzuna göz atın.
* Go’da yeniyseniz, **Effective Go** ve **How to write Go code** içinde faydalı en iyi uygulamalar bulursunuz.
* **The Go Tour**, Go’nun temellerine adım adım harika bir giriş sunar.

---

## 🧾 Tamamlanmış kod

Bu bölüm, bu eğitimle oluşturduğunuz uygulamanın kodunu içerir.

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
    ID     int64
    Title  string
    Artist string
    Price  float32
}

func main() {
    // Bağlantı özelliklerini yakalayın.
    cfg := mysql.NewConfig()
    cfg.User = os.Getenv("DBUSER")
    cfg.Passwd = os.Getenv("DBPASS")
    cfg.Net = "tcp"
    cfg.Addr = "127.0.0.1:3306"
    cfg.DBName = "recordings"

    // Bir veritabanı tanıtıcısı alın.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")

    albums, err := albumsByArtist("John Coltrane")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Albums found: %v\n", albums)

    // Sorguyu test etmek için burada ID 2'yi sabit kodlayın.
    alb, err := albumByID(2)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Album found: %v\n", alb)

    albID, err := addAlbum(Album{
        Title:  "The Modern Sound of Betty Carter",
        Artist: "Betty Carter",
        Price:  49.99,
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("ID of added album: %v\n", albID)
}

// albumsByArtist, belirtilen sanatçı adına sahip albümler için sorgu yapar.
func albumsByArtist(name string) ([]Album, error) {
    // Dönen satırlardan gelen veriyi tutacak bir albums slice'ı.
    var albums []Album

    rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    defer rows.Close()
    // Satırlar üzerinde dolaşın, sütun verisini struct alanlarına atamak için Scan kullanın.
    for rows.Next() {
        var alb Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
        }
        albums = append(albums, alb)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    return albums, nil
}

// albumByID, belirtilen ID'ye sahip albüm için sorgu yapar.
func albumByID(id int64) (Album, error) {
    // Dönen satırdan gelen veriyi tutacak bir albüm.
    var alb Album

    row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
    if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
        if err == sql.ErrNoRows {
            return alb, fmt.Errorf("albumsById %d: böyle bir albüm yok", id)
        }
        return alb, fmt.Errorf("albumsById %d: %v", id, err)
    }
    return alb, nil
}

// addAlbum, belirtilen albümü veritabanına ekler,
// yeni kaydın albüm ID'sini döndürür
func addAlbum(alb Album) (int64, error) {
    result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    return id, nil
}
```

