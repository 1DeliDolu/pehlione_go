
# 🗃️ İlişkisel Veritabanlarına Erişim

## 📚 İçindekiler

* Desteklenen veritabanı yönetim sistemleri
* Sorgu çalıştırma veya veritabanı değişiklikleri yapma fonksiyonları
* İşlemler (*transactions*)
* Sorgu iptali
* Yönetilen bağlantı havuzu

Go kullanarak, uygulamalarınıza çok çeşitli veritabanlarını ve veri erişim yaklaşımlarını entegre edebilirsiniz. Bu bölümdeki konular, standart kütüphanedeki *database/sql* paketini kullanarak ilişkisel veritabanlarına nasıl erişeceğinizi açıklar.

Go ile veri erişimine yönelik giriş niteliğinde bir eğitim için lütfen **Tutorial: Accessing a relational database** konusuna bakın.

Go, daha yüksek seviyede ilişkisel veritabanı erişimi için ORM kütüphaneleri ve ilişkisel olmayan NoSQL veri depoları dahil olmak üzere başka veri erişim teknolojilerini de destekler.

* *Object-relational mapping (ORM) libraries.* *database/sql* paketi daha düşük seviyeli veri erişim mantığı için fonksiyonlar içerirken, Go’yu veri depolarına daha yüksek bir soyutlama seviyesinde erişmek için de kullanabilirsiniz. Go için popüler iki nesne-ilişkisel eşleme (ORM) kütüphanesi hakkında daha fazla bilgi için **GORM (package reference)** ve **ent (package reference)** bölümüne bakın.
* *NoSQL data stores.* Go topluluğu; MongoDB ve Couchbase dahil, NoSQL veri depolarının çoğunluğu için sürücüler geliştirmiştir. Daha fazlası için pkg.go.dev üzerinde arama yapabilirsiniz.

---

## 🧩 Desteklenen veritabanı yönetim sistemleri

Go; MySQL, Oracle, Postgres, SQL Server, SQLite ve daha fazlası dahil olmak üzere en yaygın ilişkisel veritabanı yönetim sistemlerinin tamamını destekler.

Sürücülerin eksiksiz bir listesini **SQLDrivers** sayfasında bulabilirsiniz.

---

## 🛠️ Sorgu çalıştırma veya veritabanı değişiklikleri yapma fonksiyonları

*database/sql* paketi, yürüttüğünüz veritabanı işlemi türü için özel olarak tasarlanmış fonksiyonlar içerir. Örneğin sorguları çalıştırmak için `Query` veya `QueryRow` kullanabilseniz de, `QueryRow` yalnızca tek bir satır beklediğiniz durum için tasarlanmıştır; yalnızca bir satır içeren bir `sql.Rows` döndürmenin ek yükünü çıkarır. `INSERT`, `UPDATE` veya `DELETE` gibi SQL ifadeleriyle veritabanı değişiklikleri yapmak için `Exec` fonksiyonunu kullanabilirsiniz.

Daha fazlası için aşağıdakilere bakın:

* Veri döndürmeyen SQL ifadelerini çalıştırma
* Veri sorgulama

---

## 🔁 İşlemler (*transactions*)

`sql.Tx` aracılığıyla, işlemleri bir transaction içinde yürüten kod yazabilirsiniz. Bir transaction içinde, birden fazla işlem birlikte gerçekleştirilebilir ve tüm değişiklikleri tek bir atomik adımda uygulamak için son aşamada bir *commit* ile tamamlanabilir ya da değişiklikleri atmak için bir *rollback* yapılabilir.

Transaction’lar hakkında daha fazla bilgi için **Executing transactions** konusuna bakın.

---

## ⛔ Sorgu iptali

Bir istemcinin bağlantısı kapandığında veya işlem istediğinizden uzun sürdüğünde olduğu gibi bir veritabanı işlemini iptal edebilmek istediğinizde `context.Context` kullanabilirsiniz.

Herhangi bir veritabanı işlemi için, argüman olarak `Context` alan bir *database/sql* paketi fonksiyonunu kullanabilirsiniz. `Context` kullanarak işlem için bir zaman aşımı (*timeout*) veya son tarih (*deadline*) belirleyebilirsiniz. Ayrıca `Context` ile, uygulamanız boyunca bir iptal isteğini SQL ifadesini yürüten fonksiyona kadar iletebilir; artık ihtiyaç duyulmadığında kaynakların serbest bırakılmasını sağlayabilirsiniz.

Daha fazlası için **Canceling in-progress operations** konusuna bakın.

---

## 🧵 Yönetilen bağlantı havuzu

`sql.DB` veritabanı tanıtıcısını (*handle*) kullandığınızda, kodunuzun ihtiyaçlarına göre bağlantıları oluşturan ve sonlandıran yerleşik bir bağlantı havuzu (*connection pool*) ile bağlanırsınız. `sql.DB` üzerinden bir tanıtıcı, Go’da veritabanı erişimi için en yaygın yoldur. Daha fazlası için **Opening a database handle** konusuna bakın.

*database/sql* paketi bağlantı havuzunu sizin için yönetir. Ancak daha gelişmiş ihtiyaçlarda, **Setting connection pool properties** bölümünde açıklandığı gibi bağlantı havuzu özelliklerini ayarlayabilirsiniz.

Tek bir ayrılmış (*reserved*) bağlantıya ihtiyaç duyduğunuz işlemler için, *database/sql* paketi `sql.Conn` sağlar. `Conn`, `sql.Tx` ile bir transaction’ın kötü bir seçim olacağı durumlarda özellikle kullanışlıdır.

Örneğin kodunuz şunlara ihtiyaç duyabilir:

* Kendi transaction semantiklerini içeren mantık dahil olmak üzere, bir DDL üzerinden şema değişiklikleri yapmak. SQL transaction ifadeleriyle `sql` paketinin transaction fonksiyonlarını karıştırmak, **Executing transactions** bölümünde açıklandığı gibi kötü bir pratiktir.
* Geçici tablolar oluşturan sorgu kilitleme (*query locking*) işlemleri gerçekleştirmek.

Daha fazlası için **Using dedicated connections** konusuna bakın.

