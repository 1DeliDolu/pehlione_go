## 🗃️ Bir Go Modülünü Organize Etme

### 📑 İçindekiler

* Temel paket
* Temel komut
* Destekleyici paketlerle paket veya komut
* Birden çok paket
* Birden çok komut
* Aynı depoda paketler ve komutlar
* Sunucu projesi

Go’ya yeni başlayan geliştiricilerin sık sorduğu bir soru şudur: Dosya ve klasör yerleşimi açısından “Go projemi nasıl organize etmeliyim?” Bu belgenin amacı, bu soruyu yanıtlamaya yardımcı olacak bazı yönergeler sağlamaktır. Bu belgeden en iyi şekilde yararlanmak için, öğreticiyi okuyarak ve modül kaynağını yönetme konusuna bakarak Go modüllerinin temellerine aşina olduğunuzdan emin olun.

Go projeleri paketler, komut satırı programları veya ikisinin bir kombinasyonunu içerebilir. Bu kılavuz, proje türüne göre düzenlenmiştir.

---

## 📦 Temel Paket

Basit bir Go paketinde tüm kod, projenin kök dizininde bulunur. Proje tek bir modülden oluşur ve bu modül tek bir paketten oluşur. Paket adı, modül adının yolundaki son bileşenle eşleşir. Tek bir Go dosyası gerektiren çok basit bir paket için proje yapısı şöyledir:

```text
project-root-directory/
  go.mod
  modname.go
  modname_test.go
```

[bu belge boyunca dosya/paket adları tamamen keyfidir]

Bu dizinin GitHub’da `github.com/someuser/modname` konumuna yüklendiğini varsayarsak, **`go.mod`** dosyasındaki **`module`** satırı şöyle olmalıdır: `module github.com/someuser/modname`.

**`modname.go`** içindeki kod paketi şu şekilde bildirir:

```go
package modname

// ... package code here
```

Kullanıcılar daha sonra bu pakete, Go kodlarında şunu import ederek bağımlı olabilir:

```go
import "github.com/someuser/modname"
```

Bir Go paketi aynı dizinde kalan birden çok dosyaya bölünebilir; örneğin:

```text
project-root-directory/
  go.mod
  modname.go
  modname_test.go
  auth.go
  auth_test.go
  hash.go
  hash_test.go
```

Dizindeki tüm dosyalar `package modname` bildirir.

---

## 🛠️ Temel Komut

Temel bir çalıştırılabilir program (veya komut satırı aracı), karmaşıklığına ve kod boyutuna göre yapılandırılır. En basit program,  **`func main`** ’in tanımlandığı tek bir Go dosyasından oluşabilir. Daha büyük programlarda kod, tamamı `package main` bildiren birden çok dosyaya bölünebilir:

```text
project-root-directory/
  go.mod
  auth.go
  auth_test.go
  client.go
  main.go
```

Burada **`main.go`** dosyası **`func main`** içerir, ancak bu yalnızca bir konvansiyondur. “main” dosyası **`modname.go`** (uygun bir `modname` değeri için) veya başka herhangi bir ad da olabilir.

Bu dizinin GitHub’da `github.com/someuser/modname` konumuna yüklendiğini varsayarsak, **`go.mod`** dosyasındaki **`module`** satırı şöyle olmalıdır:

```text
module github.com/someuser/modname
```

Ve bir kullanıcı bunu makinesine şu şekilde kurabilmelidir:

```bash
$ go install github.com/someuser/modname@latest
```

---

## 🧩 Destekleyici Paketlerle Paket veya Komut

Daha büyük paketler veya komutlar, bazı işlevselliği destekleyici paketlere ayırmaktan fayda görebilir. Başlangıçta, bu tür paketleri **`internal`** adlı bir dizine koymak önerilir; bu, diğer modüllerin mutlaka dışa açmak ve harici kullanımlar için desteklemek istemediğimiz paketlere bağımlı olmasını engeller. Diğer projeler **`internal`** dizinimizden kod import edemediği için, dış kullanıcıları bozmadan API’sini refactor edebilir ve genel olarak yapıyı değiştirebiliriz. Bir paket için proje yapısı şu şekildedir:

```text
project-root-directory/
  internal/
    auth/
      auth.go
      auth_test.go
    hash/
      hash.go
      hash_test.go
  go.mod
  modname.go
  modname_test.go
```

**`modname.go`** dosyası `package modname` bildirir, **`auth.go`** `package auth` bildirir vb.  **`modname.go`** , **`auth`** paketini şu şekilde import edebilir:

```go
import "github.com/someuser/modname/internal/auth"
```

**`internal`** dizininde destekleyici paketleri olan bir komutun yerleşimi çok benzerdir; tek fark, kök dizindeki dosya(lar)ın `package main` bildirmesidir.

---

## 🧱 Birden Çok Paket

Bir modül birden çok import edilebilir paketten oluşabilir; her paketin kendi dizini vardır ve hiyerarşik olarak yapılandırılabilir. İşte örnek bir proje yapısı:

```text
project-root-directory/
  go.mod
  modname.go
  modname_test.go
  auth/
    auth.go
    auth_test.go
    token/
      token.go
      token_test.go
  hash/
    hash.go
  internal/
    trace/
      trace.go
```

Hatırlatma olarak, **`go.mod`** içindeki **`module`** satırının şöyle olduğunu varsayıyoruz:

```text
module github.com/someuser/modname
```

**`modname`** paketi kök dizinde bulunur, `package modname` bildirir ve kullanıcılar tarafından şu şekilde import edilebilir:

```go
import "github.com/someuser/modname"
```

Alt paketler kullanıcılar tarafından şu şekilde import edilebilir:

```go
import "github.com/someuser/modname/auth"
import "github.com/someuser/modname/auth/token"
import "github.com/someuser/modname/hash"
```

`internal/trace` içinde bulunan `trace` paketi bu modülün dışından import edilemez. Paketleri mümkün olduğunca **`internal`** içinde tutmak önerilir.

---

## 🧰 Birden Çok Komut

Aynı depodaki birden fazla program genellikle ayrı dizinlere sahip olur:

```text
project-root-directory/
  go.mod
  internal/
    ... shared internal packages
  prog1/
    main.go
  prog2/
    main.go
```

Her dizinde, programın Go dosyaları `package main` bildirir. Üst düzey bir **`internal`** dizini, depodaki tüm komutlar tarafından kullanılan paylaşılan paketleri içerebilir.

Kullanıcılar bu programları şu şekilde kurabilir:

```bash
$ go install github.com/someuser/modname/prog1@latest
$ go install github.com/someuser/modname/prog2@latest
```

Yaygın bir konvansiyon, bir depodaki tüm komutları **`cmd`** adlı bir dizine koymaktır; bu, yalnızca komutlardan oluşan bir depoda zorunlu değildir, ancak hem komutlar hem de import edilebilir paketler içeren karma depolarda çok faydalıdır; bunu bir sonraki bölümde ele alacağız.

---

## 🧱🛠️ Aynı Depoda Paketler ve Komutlar

Bazen bir depo, ilgili işlevselliğe sahip hem import edilebilir paketler hem de kurulabilir komutlar sağlayacaktır. İşte böyle bir depo için örnek bir proje yapısı:

```text
project-root-directory/
  go.mod
  modname.go
  modname_test.go
  auth/
    auth.go
    auth_test.go
  internal/
    ... internal packages
  cmd/
    prog1/
      main.go
    prog2/
      main.go
```

Bu modülün adının `github.com/someuser/modname` olduğunu varsayarsak, kullanıcılar artık hem ondan paket import edebilir:

```go
import "github.com/someuser/modname"
import "github.com/someuser/modname/auth"
```

Hem de ondan program kurabilir:

```bash
$ go install github.com/someuser/modname/cmd/prog1@latest
$ go install github.com/someuser/modname/cmd/prog2@latest
```

---

## 🖥️ Sunucu Projesi

Go, sunucu uygulamak için yaygın bir dil tercihidir. Sunucu geliştirmesinin birçok yönü (protokoller (REST? gRPC?), dağıtımlar, ön yüz dosyaları, containerization, script’ler vb.) olduğundan, bu tür projelerin yapısında çok büyük bir çeşitlilik vardır. Buradaki rehberliğimizi, projenin Go ile yazılmış kısımlarına odaklayacağız.

Sunucu projelerinde genellikle dışa aktarım için paketler bulunmaz; çünkü bir sunucu genellikle kendi kendine yeterli bir ikili dosyadır (veya bir grup ikili dosya). Bu nedenle, sunucu mantığını uygulayan Go paketlerini **`internal`** dizininde tutmak önerilir. Dahası, projenin Go olmayan dosyalar içeren başka birçok dizini olacağı muhtemel olduğundan, tüm Go komutlarını **`cmd`** dizininde birlikte tutmak iyi bir fikirdir:

```text
project-root-directory/
  go.mod
  internal/
    auth/
      ...
    metrics/
      ...
    model/
      ...
  cmd/
    api-server/
      main.go
    metrics-analyzer/
      main.go
    ...
  ... the project's other directories with non-Go code
```

Sunucu deposu, diğer projelerle paylaşmak için faydalı hâle gelen paketler büyütürse, bunları ayrı modüllere ayırmak en iyisidir.
