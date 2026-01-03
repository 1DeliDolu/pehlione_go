## 🧩 Tutorial: Bir Go Modülü Oluşturma

### 📌 İçindekiler

* Prerequisites
* Start a module that others can use

Bu, Go dilinin birkaç temel özelliğini tanıtan bir öğreticinin ilk bölümüdür. Go’ya yeni başlıyorsanız, `go` komutunu, Go modüllerini ve çok basit Go kodunu tanıtan **Tutorial: Get started with Go** öğreticisine mutlaka göz atın.

Bu öğreticide iki modül oluşturacaksınız. İlki, diğer kütüphaneler veya uygulamalar tarafından içe aktarılması amaçlanan bir **kütüphane** olacaktır. İkincisi ise birincisini kullanacak bir **çağıran uygulama** olacaktır.

Bu öğreticinin sırası, dilin farklı bir parçasını gösteren yedi kısa konudan oluşur:

* **Create a module** — Başka bir modülden çağırabileceğiniz fonksiyonlara sahip küçük bir modül yazın.
* **Call your code from another module** — Yeni modülünüzü içe aktarın ve kullanın.
* **Return and handle an error** — Basit hata yönetimi ekleyin.
* **Return a random greeting** — Dilimlerde (Go’nun dinamik boyutlu dizileri) veri işleyin.
* **Return greetings for multiple people** — Bir map içinde anahtar/değer çiftlerini saklayın.
* **Add a test** — Go’nun yerleşik birim test özelliklerini kullanarak kodunuzu test edin.
* **Compile and install the application** — Kodunuzu yerelde derleyin ve kurun.

> **Not:** Diğer öğreticiler için bkz.  *Tutorials* .

---

## ✅ Prerequisites

* Bir miktar programlama deneyimi. Buradaki kod oldukça basit, ancak fonksiyonlar, döngüler ve diziler hakkında biraz bilgi sahibi olmak faydalıdır.
* Kodunuzu düzenlemek için bir araç. Herhangi bir metin düzenleyici işinizi görür. Çoğu metin düzenleyici Go için iyi destek sunar. En popüler olanlar: **VSCode** (ücretsiz), **GoLand** (ücretli) ve **Vim** (ücretsiz).
* Bir komut terminali. Go; Linux ve Mac’te herhangi bir terminalde, Windows’ta ise PowerShell veya `cmd` üzerinde iyi çalışır.

---

## 🚀 Start a module that others can use

Bir Go modülü oluşturarak başlayın. Bir modülde, ayrık ve faydalı bir fonksiyon kümesi için bir veya daha fazla ilişkili paketi bir araya getirirsiniz. Örneğin, finansal analiz yapmak için fonksiyonlar içeren paketlere sahip bir modül oluşturabilirsiniz; böylece finansal uygulamalar yazan başkaları sizin çalışmanızı kullanabilir. Modül geliştirme hakkında daha fazlası için bkz.  *Developing and publishing modules* .

Go kodu paketler halinde gruplanır ve paketler modüller halinde gruplanır. Modülünüz, kodunuzu çalıştırmak için gereken bağımlılıkları belirtir; buna Go sürümü ve ihtiyaç duyduğu diğer modüllerin kümesi dahildir.

Modülünüzde işlevsellik ekledikçe veya geliştirdikçe, modülün yeni sürümlerini yayımlarsınız. Modülünüzdeki fonksiyonları çağıran geliştiriciler, modülün güncellenmiş paketlerini içe aktarabilir ve üretimde kullanmadan önce yeni sürümle test edebilir.

Komut istemini açın ve ev dizininize (`home directory`) gidin.

Linux veya Mac’te:

```bash
cd
```

Windows’ta:

```bash
cd %HOMEPATH%
```

Go modülünüzün kaynak kodu için bir `greetings` dizini oluşturun. Örneğin, ev dizininizden aşağıdaki komutları kullanın:

```bash
mkdir greetings
cd greetings
```

### 📦 Modülünüzü `go mod init` ile başlatın

`go mod init` komutunu çalıştırın ve ona modül yolunuzu ( *module path* ) verin — burada `example.com/greetings` kullanın. Bir modül yayımlarsanız, bu yol Go araçlarının modülünüzü indirebileceği bir yol olmalıdır. Bu da kodunuzun deposu ( *repository* ) olacaktır.

Modülünüzü bir *module path* ile adlandırma hakkında daha fazla bilgi için bkz.  *Managing dependencies* .

```bash
$ go mod init example.com/greetings
go: creating new go.mod: module example.com/greetings
```

`go mod init` komutu, kodunuzun bağımlılıklarını izlemek için bir `go.mod` dosyası oluşturur. Şu ana kadar dosya yalnızca modülünüzün adını ve kodunuzun desteklediği Go sürümünü içerir. Ancak bağımlılıklar ekledikçe, `go.mod` dosyası kodunuzun bağlı olduğu sürümleri listeler. Bu, derlemeleri yeniden üretilebilir tutar ve hangi modül sürümlerini kullanacağınıza doğrudan kontrol sağlar.

Metin düzenleyicinizde kodunuzu yazmak için bir dosya oluşturun ve adını `greetings.go` koyun.

Aşağıdaki kodu `greetings.go` dosyanıza yapıştırın ve kaydedin:

```go
package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
```

Bu, modülünüz için ilk koddur. Bir selamlama isteyen herhangi bir çağırana selamlama döndürür. Bir sonraki adımda bu fonksiyonu çağıran kodu yazacaksınız.

Bu kodda şunları yaparsınız:

* İlişkili fonksiyonları bir araya toplamak için bir `greetings` paketi tanımlarsınız.
* Selamlamayı döndürmek için bir `Hello` fonksiyonu uygularsınız.
  Bu fonksiyon `string` türünde bir `name` parametresi alır. Fonksiyon ayrıca bir `string` döndürür. Go’da adı büyük harfle başlayan bir fonksiyon, aynı pakette olmayan bir fonksiyon tarafından çağrılabilir. Bu, Go’da **exported name** olarak bilinir. Exported adlar hakkında daha fazlası için bkz.  *Exported names in the Go tour* .
* Selamlamayı tutmak için bir `message` değişkeni tanımlarsınız.
  Go’da `:=` operatörü, bir değişkeni tek satırda bildirmek ve başlatmak için bir kısayoldur (Go, sağ taraftaki değeri kullanarak değişkenin türünü belirler). Uzun yoldan şöyle yazmış olabilirdiniz:

```go
var message string
message = fmt.Sprintf("Hi, %v. Welcome!", name)
```

* Selamlama mesajını oluşturmak için `fmt` paketinin `Sprintf` fonksiyonunu kullanırsınız. İlk argüman bir biçim dizesidir ( *format string* ) ve `Sprintf`, `name` parametresinin değerini `%v` biçim fiili ( *format verb* ) yerine koyar. `name` parametresinin değerini eklemek, selamlama metnini tamamlar.
* Biçimlendirilmiş selamlama metnini çağırana döndürürsünüz.

Bir sonraki adımda, bu fonksiyonu başka bir modülden çağıracaksınız.
