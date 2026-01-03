## 🚀 Tutorial: Go ile Başlayın

### 📌 İçindekiler

* Prerequisites
* Install Go
* Write some code
* Call code in an external package
* Write more code

Bu öğreticide Go programlamaya kısa bir giriş yapacaksınız. Bu süreçte şunları yapacaksınız:

* Go’yu kurmak (henüz kurmadıysanız).
* Basit bir **"Hello, world"** kodu yazmak.
* `go` komutunu kullanarak kodunuzu çalıştırmak.
* Kendi kodunuzda kullanabileceğiniz paketleri bulmak için Go paket keşif aracını kullanmak.
* Harici bir modülün fonksiyonlarını çağırmak.

> **Not:** Diğer öğreticiler için bkz.  *Tutorials* .

---

## ✅ Prerequisites

* Bir miktar programlama deneyimi. Buradaki kod oldukça basit, ancak fonksiyonlar hakkında biraz bilgi sahibi olmak faydalıdır.
* Kodunuzu düzenlemek için bir araç. Herhangi bir metin düzenleyici işinizi görür. Çoğu metin düzenleyici Go için iyi destek sunar. En popüler olanlar: **VSCode** (ücretsiz), **GoLand** (ücretli) ve **Vim** (ücretsiz).
* Bir komut terminali. Go; Linux ve Mac’te herhangi bir terminalde, Windows’ta ise PowerShell veya `cmd` üzerinde iyi çalışır.

---

## 🧰 Install Go

Sadece **Download and install** adımlarını kullanın.

---

## ✍️ Write some code

**Hello, World** ile başlayın.

Bir komut istemi açın ve ev dizininize (`home directory`) gidin.

Linux veya Mac’te:

```bash
cd
```

Windows’ta:

```bash
cd %HOMEPATH%
```

İlk Go kaynak kodunuz için bir `hello` dizini oluşturun. Örneğin aşağıdaki komutları kullanın:

```bash
mkdir hello
cd hello
```

### 📦 Kodunuz için bağımlılık takibini etkinleştirin

Kodunuz başka modüllerde bulunan paketleri içe aktardığında, bu bağımlılıkları kodunuzun kendi modülü üzerinden yönetirsiniz. Bu modül, paketleri sağlayan modülleri izleyen bir `go.mod` dosyası ile tanımlanır. Bu `go.mod` dosyası, kaynak kod deponuz dahil olmak üzere kodunuzla birlikte kalır.

Kodunuz için bir `go.mod` dosyası oluşturarak bağımlılık takibini etkinleştirmek için `go mod init` komutunu çalıştırın ve kodunuzun içinde bulunacağı modülün adını verin. Bu ad, modülün  *module path* ’idir.

Gerçek geliştirmede, *module path* genellikle kaynak kodunuzun tutulacağı depo konumu olur. Örneğin, `module path` **github.com/mymodule** olabilir. Modülünüzü başkalarının kullanması için yayımlamayı planlıyorsanız, *module path* Go araçlarının modülünüzü indirebileceği bir konum olmalıdır. Bir modülü *module path* ile adlandırma hakkında daha fazla bilgi için bkz.  *Managing dependencies* .

Bu öğretici için sadece `example/hello` kullanın.

```bash
$ go mod init example/hello
go: creating new go.mod: module example/hello
```

Metin düzenleyicinizde, kodunuzu yazmak için `hello.go` adında bir dosya oluşturun.

Aşağıdaki kodu `hello.go` dosyanıza yapıştırın ve kaydedin:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Bu sizin Go kodunuzdur. Bu kodda şunları yaparsınız:

* **`main` paketini** tanımlarsınız ( *package* , fonksiyonları gruplamanın bir yoludur ve aynı dizindeki tüm dosyalardan oluşur).
* Konsola yazdırma dahil metni biçimlendirmek için fonksiyonlar içeren popüler **`fmt` paketini** içe aktarırsınız. Bu paket, Go’yu kurduğunuzda gelen standart kütüphane paketlerinden biridir.
* Konsola bir mesaj yazdırmak için bir **`main` fonksiyonu** uygularsınız. `main` fonksiyonu, `main` paketini çalıştırdığınızda varsayılan olarak yürütülür.

Kodunuzu çalıştırarak selamlamayı görün:

```bash
$ go run .
Hello, World!
```

`go run` komutu, Go ile işlerinizi yapmak için kullanacağınız birçok `go` komutundan biridir. Diğerlerinin listesini almak için şu komutu kullanın:

```bash
$ go help
```

---

## 🌐 Call code in an external package

Kodunuzun, başkası tarafından zaten uygulanmış olabilecek bir işi yapması gerektiğinde, kodunuzda kullanabileceğiniz fonksiyonlara sahip bir paket arayabilirsiniz.

Harici bir modülden bir fonksiyon kullanarak yazdırdığınız mesajı biraz daha ilginç hale getirin.

* `pkg.go.dev` sitesini ziyaret edin ve bir **"quote"** paketi arayın.
* Arama sonuçlarında `rsc.io/quote` paketinin **v1** sürümünü bulun ve tıklayın (bu, `rsc.io/quote/v4`’ün **"Other major versions"** bölümünde listelenmiş olmalıdır).
* **Documentation** bölümünde, **Index** altında kodunuzdan çağırabileceğiniz fonksiyonların listesini not edin. **Go** fonksiyonunu kullanacaksınız.
* Bu sayfanın üst kısmında, `quote` paketinin `rsc.io/quote` modülüne dahil olduğunu not edin.

`pkg.go.dev` sitesini, kendi kodunuzda kullanabileceğiniz fonksiyonlara sahip paketleri barındıran yayımlanmış modülleri bulmak için kullanabilirsiniz. Paketler, başkalarının kullanabileceği modüller içinde yayımlanır — `rsc.io/quote` gibi. Modüller zamanla yeni sürümlerle geliştirilir ve siz de bu geliştirilmiş sürümleri kullanmak için kodunuzu yükseltebilirsiniz.

### 🧩 Go kodunuzda harici paketi içe aktarın ve fonksiyonu çağırın

Go kodunuzda `rsc.io/quote` paketini içe aktarın ve onun `Go` fonksiyonuna bir çağrı ekleyin.

Vurgulanan satırları ekledikten sonra kodunuz aşağıdaki gibi olmalıdır:

```go
package main

import "fmt"

import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}
```

### 📥 Yeni modül gereksinimlerini ve özetleri ekleyin

Go, `quote` modülünü bir gereksinim olarak ekleyecek ve ayrıca modülü doğrulamak için kullanılacak bir `go.sum` dosyası oluşturacaktır. Daha fazlası için bkz.  *Authenticating modules in the Go Modules Reference* .

```bash
$ go mod tidy
go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.2
```

Kodunuzu çalıştırarak çağırdığınız fonksiyonun ürettiği mesajı görün:

```bash
$ go run .
Don't communicate by sharing memory, share memory by communicating.
```

Kodunuzun `Go` fonksiyonunu çağırdığını ve iletişim hakkında zekice bir mesaj yazdırdığını fark edin.

`go mod tidy` çalıştırdığınızda, içe aktardığınız paketi içeren `rsc.io/quote` modülünü buldu ve indirdi. Varsayılan olarak en son sürümü indirdi —  **v1.5.2** .

---

## 🧱 Write more code

Bu hızlı girişle Go’yu kurdunuz ve temel konulardan bazılarını öğrendiniz. Başka bir öğreticiyle biraz daha kod yazmak için *Create a Go module* öğreticisine göz atın.
