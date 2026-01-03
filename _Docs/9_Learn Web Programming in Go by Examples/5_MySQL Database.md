
## 🗄️ MySQL Veritabanı

## 🚀 Giriş

Bir noktada, web uygulamanızın bir veritabanına veri kaydetmesi ve veritabanından veri okuması gerekir. Bu durum, dinamik içerikle çalışırken, kullanıcıların veri girmesi için formlar sunarken veya kullanıcıların kimlik doğrulaması için giriş ve parola bilgilerini saklarken neredeyse her zaman geçerlidir. Bu amaçla veritabanları vardır.

Veritabanları birçok farklı biçimde karşımıza çıkar. Web genelinde yaygın olarak kullanılan veritabanlarından biri de MySQL veritabanıdır. Uzun zamandır vardır ve konumunu ve kararlılığını defalarca kanıtlamıştır.

Bu örnekte, Go’da veritabanına erişimin temellerine ineceğiz; veritabanı tabloları oluşturacak, veri saklayacak ve tekrar geri okuyacağız.

---

## 📦 go-sql-driver/mysql Paketini Yükleme

Go programlama dili, her türden SQL veritabanını sorgulamak için kullanışlı bir paket olan `database/sql` ile birlikte gelir. Bu, ortak SQL özelliklerini tek bir API altında soyutladığı için faydalıdır. Ancak Go’nun dahil etmediği şey veritabanı sürücüleridir. Go’da veritabanı sürücüsü, belirli bir veritabanının (bizim durumumuzda MySQL) düşük seviye detaylarını uygulayan bir pakettir. Tahmin edebileceğiniz gibi bu yaklaşım ileriye dönük uyumluluk için faydalıdır. Çünkü tüm Go paketleri oluşturulurken yazarlar gelecekte ortaya çıkabilecek her veritabanını önceden bilemez; üstelik mümkün olan her veritabanını desteklemek ciddi bir bakım yükü olurdu.

MySQL veritabanı sürücüsünü yüklemek için terminalinizde şu komutu çalıştırın:

```bash
go get -u github.com/go-sql-driver/mysql
```

---

## 🔌 MySQL Veritabanına Bağlanma

Gerekli tüm paketleri kurduktan sonra kontrol etmemiz gereken ilk şey, MySQL veritabanımıza başarıyla bağlanıp bağlanamadığımızdır. Eğer hali hazırda çalışan bir MySQL veritabanı sunucunuz yoksa, Docker ile kolayca yeni bir örnek başlatabilirsiniz. Docker MySQL imajının resmi dokümanları burada: [https://hub.docker.com/_/mysql](https://hub.docker.com/_/mysql)

Veritabanımıza bağlanabildiğimizi kontrol etmek için `database/sql` ve `go-sql-driver/mysql` paketlerini import edin ve şu şekilde bir bağlantı açın:

```go
import "database/sql"
import _ "go-sql-driver/mysql"


// Configure the database connection (always check errors)
db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")



// Initialize the first connection to the database, to see if everything works correctly.
// Make sure to check the error.
err := db.Ping()
```

---

## 🧱 İlk Veritabanı Tablomuzu Oluşturma

Veritabanımızdaki her veri girişi belirli bir tabloda saklanır. Bir veritabanı tablosu sütunlardan ve satırlardan oluşur. Sütunlar, her veri girişine bir etiket verir ve türünü belirtir. Satırlar ise eklenen veri değerleridir. İlk örneğimizde şöyle bir tablo oluşturmak istiyoruz:

id | username | password | created_at
1 | johndoe | secret | 2019-08-10 12:30:00

SQL’e çevrildiğinde, bu tabloyu oluşturma komutu şöyle görünür:

```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at DATETIME,
    PRIMARY KEY (id)
);
```

Artık SQL komutumuz olduğuna göre, `database/sql` paketini kullanarak tabloyu MySQL veritabanımızda oluşturabiliriz:

```go
query := `
    CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`

// Executes the SQL query in our database. Check err to ensure there was no error.
_, err := db.Exec(query)
```

---

## 👤 İlk Kullanıcımızı Ekleme

Eğer SQL’e aşinaysanız, tablomuza yeni veri eklemek tablo oluşturmak kadar kolaydır. Dikkat edilmesi gereken bir nokta: Varsayılan olarak Go, SQL sorgularına dinamik veri eklemek için *prepared statement* kullanır; bu, kullanıcıdan gelen verileri veritabanına güvenli biçimde aktarmanın ve zarar verme riskini ortadan kaldırmanın bir yoludur. Web programlamanın ilk dönemlerinde programcılar veriyi sorguyla birlikte doğrudan veritabanına gönderirdi; bu yaklaşım büyük güvenlik açıklarına yol açar ve tüm uygulamayı bozabilirdi. Lütfen bunu yapmayın. Doğru yapmak kolaydır.

İlk kullanıcıyı eklemek için aşağıdaki gibi bir SQL sorgusu oluştururuz. Gördüğünüz gibi `id` sütununu atlıyoruz; çünkü MySQL bunu otomatik olarak ayarlar. Soru işaretleri, SQL sürücüsüne gerçek veriler için yer tutucu olduklarını söyler. Hazırlanmış ifadeleri (*prepared statements*) burada görebilirsiniz.

```sql
INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)
```

Bu SQL sorgusunu Go’da kullanıp tablomuzda yeni bir satır ekleyebiliriz:

```go
import "time"

username := "johndoe"
password := "secret"
createdAt := time.Now()

// Inserts our data into the users table and returns with the result and a possible error.
// The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.
result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
```

Yeni oluşturulan kullanıcı id’sini almak için şöyle yapın:

```go
userID, err := result.LastInsertId()
```

---

## 🔎 users Tablosunu Sorgulama

Artık tabloda bir kullanıcı olduğuna göre, onu sorgulayıp tüm bilgilerini geri almak istiyoruz. Go’da tabloları sorgulamak için iki seçeneğimiz vardır. `db.Query` birden fazla satırı sorgular ve üzerinde dolaşmamıza izin verir; `db.QueryRow` ise yalnızca belirli bir satırı sorgulamak istediğimiz durumlar içindir.

Belirli bir satırı sorgulamak, daha önce ele aldığımız diğer SQL komutlarıyla temelde aynı şekilde çalışır. Tek bir kullanıcıyı ID’sine göre sorgulamak için SQL komutumuz şöyle görünür:

```sql
SELECT id, username, password, created_at FROM users WHERE id = ?
```

Go’da önce veriyi saklamak için bazı değişkenler tanımlarız ve ardından tek bir veritabanı satırını şöyle sorgularız:

```go
var (
    id        int
    username  string
    password  string
    createdAt time.Time
)

// Query the database and scan the values into out variables. Don't forget to check for errors.
query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt)
```

---

## 👥 Tüm Kullanıcıları Sorgulama

Önceki bölümde tek bir kullanıcı satırını sorgulamayı ele aldık. Birçok uygulamada tüm mevcut kullanıcıları sorgulamak isteyeceğiniz kullanım senaryoları vardır. Bu, yukarıdaki örneğe benzer şekilde çalışır ancak biraz daha fazla kod gerektirir.

Yukarıdaki örnekteki SQL komutunu kullanıp `WHERE` kısmını çıkarabiliriz. Böylece tüm kullanıcıları sorgularız.

```sql
SELECT id, username, password, created_at FROM users
```

Go’da önce veriyi saklamak için bazı değişkenler tanımlarız ve ardından veritabanı satırlarını şöyle sorgularız:

```go
type user struct {
    id        int
    username  string
    password  string
    createdAt time.Time
}

rows, err := db.Query(`SELECT id, username, password, created_at FROM users`) // check err
defer rows.Close()

var users []user
for rows.Next() {
    var u user
    err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt) // check err
    users = append(users, u)
}
err := rows.Err() // check err
```

`users` dilimi artık şuna benzer bir içerik barındırabilir:

```go
users {
    user {
        id:        1,
        username:  "johndoe",
        password:  "secret",
        createdAt: time.Time{wall: 0x0, ext: 63701044325, loc: (*time.Location)(nil)},
    },
    user {
        id:        2,
        username:  "alice",
        password:  "bob",
        createdAt: time.Time{wall: 0x0, ext: 63701044622, loc: (*time.Location)(nil)},
    },
}
```

---

## 🗑️ Tablodan Kullanıcı Silme

Son olarak, tablodan kullanıcı silmek, önceki bölümlerdeki `.Exec` kullanımı kadar doğrudandır:

```go
_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1) // check err
```

---

## 📄 Kod (Kopyala/Yapıştır İçin)

Bu, bu örnekte öğrendiklerinizi denemek için kullanabileceğiniz tam koddur.

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/root?parseTime=true")
    if err != nil {
        log.Fatal(err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    { // Create a new table
        query := `
            CREATE TABLE users (
                id INT AUTO_INCREMENT,
                username TEXT NOT NULL,
                password TEXT NOT NULL,
                created_at DATETIME,
                PRIMARY KEY (id)
            );`

        if _, err := db.Exec(query); err != nil {
            log.Fatal(err)
        }
    }

    { // Insert a new user
        username := "johndoe"
        password := "secret"
        createdAt := time.Now()

        result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
        if err != nil {
            log.Fatal(err)
        }

        id, err := result.LastInsertId()
        fmt.Println(id)
    }

    { // Query a single user
        var (
            id        int
            username  string
            password  string
            createdAt time.Time
        )

        query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
        if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
            log.Fatal(err)
        }

        fmt.Println(id, username, password, createdAt)
    }

    { // Query all users
        type user struct {
            id        int
            username  string
            password  string
            createdAt time.Time
        }

        rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
        if err != nil {
            log.Fatal(err)
        }
        defer rows.Close()

        var users []user
        for rows.Next() {
            var u user

            err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
            if err != nil {
                log.Fatal(err)
            }
            users = append(users, u)
        }
        if err := rows.Err(); err != nil {
            log.Fatal(err)
        }

        fmt.Printf("%#v", users)
    }

    {
        _, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
        if err != nil {
            log.Fatal(err)
        }
    }
}
```

