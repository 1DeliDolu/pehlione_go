# 📝 Go Kodu Nasıl Yazılır

## 📚 İçindekiler

Giriş
Kod organizasyonu
İlk programınız
Modülünüzden paketleri import etme
Uzak modüllerden paketleri import etme
Test
Sırada ne var
Yardım alma

---

## 🧭 Giriş¶

Bu doküman, bir modül içinde basit bir Go paketinin geliştirilmesini gösterir ve Go modüllerini, paketlerini ve komutlarını getirmek (fetch), derlemek (build) ve kurmak (install) için standart yol olan `go` aracını tanıtır.

---

## 🗂️ Kod organizasyonu¶

Go programları paketler (packages) halinde organize edilir. Bir paket, aynı dizindeki ve birlikte derlenen kaynak dosyalarının bir koleksiyonudur. Bir kaynak dosyada tanımlanan fonksiyonlar, türler, değişkenler ve sabitler, aynı paket içindeki diğer tüm kaynak dosyalara görünürdür.

Bir depo (repository) bir veya daha fazla modül içerir. Bir modül, birlikte yayınlanan ilgili Go paketlerinin bir koleksiyonudur. Bir Go deposu tipik olarak yalnızca bir modül içerir ve bu modül deponun kök dizininde yer alır. Oradaki `go.mod` adlı bir dosya, modül yolunu bildirir: modül içindeki tüm paketler için import yolunun (import path) önekidir. Modül; `go.mod` dosyasını içeren dizindeki paketleri ve o dizinin alt dizinlerini (eğer varsa başka bir `go.mod` içeren bir sonraki alt dizine kadar) içerir.

Kodunuzu derleyebilmek için onu uzak bir depoya yayınlamanız gerekmediğini unutmayın. Bir modül, bir depoya ait olmadan yerel olarak da tanımlanabilir. Bununla birlikte, bir gün yayınlayacakmışsınız gibi kodunuzu organize etmek iyi bir alışkanlıktır.

Her modülün yolu, yalnızca paketleri için bir import yolu öneki olarak hizmet etmez; aynı zamanda `go` komutunun onu indirmek için nerede araması gerektiğini de belirtir. Örneğin `golang.org/x/tools` modülünü indirmek için, `go` komutu `https://golang.org/x/tools` tarafından işaret edilen depoya danışır (daha fazlası burada açıklanmıştır).

Bir import yolu (import path), bir paketi import etmek için kullanılan bir string’dir. Bir paketin import yolu, modül yolunun modül içindeki alt dizinle birleştirilmiş halidir. Örneğin `github.com/google/go-cmp` modülü, `cmp/` dizininde bir paket içerir. Bu paketin import yolu `github.com/google/go-cmp/cmp` olur. Standart kütüphanedeki paketlerin modül yolu öneki yoktur.

---

## 👋 İlk programınız¶

Basit bir programı derlemek ve çalıştırmak için önce bir modül yolu seçin (biz `example/user/hello` kullanacağız) ve bunu bildiren bir `go.mod` dosyası oluşturun:

```bash
$ mkdir hello # Alternatively, clone it if it already exists in version control.
$ cd hello
$ go mod init example/user/hello
go: creating new go.mod: module example/user/hello
$ cat go.mod
module example/user/hello

go 1.16
$
```

Bir Go kaynak dosyasındaki ilk ifade (statement) paket adı (package name) olmalıdır. Çalıştırılabilir komutlar (executable commands) her zaman `package main` kullanmalıdır.

Ardından, o dizinde `hello.go` adlı bir dosya oluşturun ve aşağıdaki Go kodunu içine koyun:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world.")
}
```

Artık bu programı `go` aracıyla derleyip kurabilirsiniz:

```bash
$ go install example/user/hello
$
```

Bu komut `hello` komutunu derler ve çalıştırılabilir bir ikili (binary) üretir. Sonra bu ikiliyi `$HOME/go/bin/hello` (veya Windows’ta `%USERPROFILE%\go\bin\hello.exe`) olarak kurar.

Kurulum dizini, `GOPATH` ve `GOBIN` ortam değişkenleri tarafından kontrol edilir. `GOBIN` ayarlanmışsa, ikililer bu dizine kurulur. `GOPATH` ayarlanmışsa, ikililer `GOPATH` listesindeki ilk dizinin `bin` alt dizinine kurulur. Aksi halde, ikililer varsayılan `GOPATH`’in (`$HOME/go` veya `%USERPROFILE%\go`) `bin` alt dizinine kurulur.

Gelecekteki `go` komutları için bir ortam değişkeninin varsayılan değerini taşınabilir (portable) biçimde ayarlamak üzere `go env` komutunu kullanabilirsiniz:

```bash
$ go env -w GOBIN=/somewhere/else/bin
$
```

`go env -w` ile daha önce ayarladığınız bir değişkeni kaldırmak için `go env -u` kullanın:

```bash
$ go env -u GOBIN
$
```

`go install` gibi komutlar, mevcut çalışma dizinini içeren modül bağlamında uygulanır. Çalışma dizini `example/user/hello` modülünün içinde değilse, `go install` başarısız olabilir.

Kolaylık için, `go` komutları çalışma dizinine göre göreli yolları kabul eder ve başka bir yol verilmezse mevcut çalışma dizinindeki paketi varsayar. Bu nedenle çalışma dizinimizde aşağıdaki komutların hepsi eşdeğerdir:

```bash
$ go install example/user/hello
$ go install .
$ go install
```

Şimdi programın çalıştığından emin olmak için çalıştıralım. Ek kolaylık olarak, ikilileri kolay çalıştırmak için kurulum dizinini `PATH`’inize ekleyelim:

```bash
# Windows users should consult /wiki/SettingGOPATH
# for setting %PATH%.
$ export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
$ hello
Hello, world.
$
```

Bir kaynak kontrol sistemi kullanıyorsanız, şimdi bir depo başlatmak, dosyaları eklemek ve ilk değişikliğinizi commit etmek için iyi bir zaman olabilir. Bu adım yine opsiyoneldir: Go kodu yazmak için kaynak kontrol kullanmanız gerekmez.

```bash
$ git init
Initialized empty Git repository in /home/user/hello/.git/
$ git add go.mod hello.go
$ git commit -m "initial commit"
[master (root-commit) 0b4507d] initial commit
 1 file changed, 7 insertion(+)
 create mode 100644 go.mod hello.go
$
```

`go` komutu, verilen bir modül yolunu içeren depoyu; karşılık gelen bir HTTPS URL’si isteyerek ve HTML yanıtına gömülü metadatayı okuyarak bulur (`go help importpath`’e bakın). Go kodu içeren depolar için birçok hosting servisi bu metadatayı zaten sağlar; bu nedenle modülünüzü başkalarının kullanabilmesi için kullanılabilir kılmanın en kolay yolu, modül yolunuzun depo URL’siyle eşleşmesidir.

---

## 🧩 Modülünüzden paketleri import etme¶

Bir `morestrings` paketi yazalım ve onu `hello` programından kullanalım. Önce, paket için `$HOME/hello/morestrings` adlı bir dizin oluşturun ve sonra bu dizinde `reverse.go` adlı bir dosya oluşturup aşağıdaki içeriği koyun:

```go
// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package morestrings

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
```

`ReverseRunes` fonksiyonumuz büyük harfle başladığı için dışa aktarılmıştır (exported) ve `morestrings` paketini import eden diğer paketlerde kullanılabilir.

Paketin `go build` ile derlendiğini test edelim:

```bash
$ cd $HOME/hello/morestrings
$ go build
$
```

Bu, bir çıktı dosyası üretmez. Bunun yerine, derlenmiş paketi yerel derleme önbelleğine (local build cache) kaydeder.

`morestrings` paketinin derlendiğini doğruladıktan sonra, onu `hello` programından kullanalım. Bunun için, orijinal `$HOME/hello/hello.go` dosyanızı `morestrings` paketini kullanacak şekilde değiştirin:

```go
package main

import (
    "fmt"

    "example/user/hello/morestrings"
)

func main() {
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
}
```

`hello` programını kurun:

```bash
$ go install example/user/hello
```

Programın yeni sürümünü çalıştırdığınızda, yeni ve ters çevrilmiş bir mesaj görmelisiniz:

```bash
$ hello
Hello, Go!
```

---

## 🌍 Uzak modüllerden paketleri import etme¶

Bir import yolu, bir sürüm kontrol sistemi (Git veya Mercurial gibi) kullanarak paket kaynak kodunun nasıl elde edileceğini tanımlayabilir. `go` aracı bu özelliği, paketleri uzak depolardan otomatik olarak çekmek için kullanır. Örneğin programınızda `github.com/google/go-cmp/cmp` kullanmak için:

```go
package main

import (
    "fmt"

    "example/user/hello/morestrings"
    "github.com/google/go-cmp/cmp"
)

func main() {
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
    fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
```

Artık harici bir modüle bağımlılığınız olduğuna göre, bu modülü indirmeniz ve sürümünü `go.mod` dosyanıza kaydetmeniz gerekir. `go mod tidy` komutu; import edilen paketler için eksik modül gereksinimlerini ekler ve artık kullanılmayan modüller üzerindeki gereksinimleri kaldırır.

```bash
$ go mod tidy
go: finding module for package github.com/google/go-cmp/cmp
go: found github.com/google/go-cmp/cmp in github.com/google/go-cmp v0.5.4
$ go install example/user/hello
$ hello
Hello, Go!
  string(
-     "Hello World",
+     "Hello Go",
  )
$ cat go.mod
module example/user/hello

go 1.16

require github.com/google/go-cmp v0.5.4
$
```

Modül bağımlılıkları, `GOPATH` ortam değişkeninin işaret ettiği dizinin `pkg/mod` alt dizinine otomatik olarak indirilir. Belirli bir modül sürümü için indirilen içerikler, o sürümü gerektiren diğer tüm modüller arasında paylaşılır; bu nedenle `go` komutu bu dosya ve dizinleri salt-okunur (read-only) olarak işaretler. İndirilen tüm modülleri kaldırmak için `go clean` komutuna `-modcache` bayrağını verebilirsiniz:

```bash
$ go clean -modcache
$
```

---

## ✅ Test¶

Go; `go test` komutu ve `testing` paketinden oluşan hafif bir test çerçevesine sahiptir.

Bir test yazmak için, adı `_test.go` ile biten ve imzası `func (t *testing.T)` olan `TestXXX` adlı fonksiyonlar içeren bir dosya oluşturursunuz. Test çerçevesi her bir bu fonksiyonu çalıştırır; eğer fonksiyon `t.Error` veya `t.Fail` gibi bir başarısızlık fonksiyonunu çağırırsa test başarısız kabul edilir.

Aşağıdaki Go kodunu içeren `$HOME/hello/morestrings/reverse_test.go` dosyasını oluşturarak `morestrings` paketine bir test ekleyin.

```go
package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
    cases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {"Hello, 世界", "界世 ,olleH"},
        {"", ""},
    }
    for _, c := range cases {
        got := ReverseRunes(c.in)
        if got != c.want {
            t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}
```

Sonra testi `go test` ile çalıştırın:

```bash
$ cd $HOME/hello/morestrings
$ go test
PASS
ok  	example/user/hello/morestrings 0.165s
$
```

Daha fazla ayrıntı için `go help test` çalıştırın ve `testing` paketi dokümantasyonuna bakın.

---

## ⏭️ Sırada ne var¶

Yeni bir stabil Go sürümü yayınlandığında haberdar edilmek için `golang-announce` posta listesine abone olun.

Açık, idiomatik Go kodu yazmak için ipuçları için  *Effective Go* ’ya bakın.

Dil hakkında öğrenmek için  *A Tour of Go* ’yu takip edin.

Go diliyle, kütüphaneleriyle ve araçlarıyla ilgili derinlemesine makaleler için dokümantasyon sayfasını ziyaret edin.

---

## 🆘 Yardım alma¶

Gerçek zamanlı yardım için, topluluk tarafından işletilen gophers Slack sunucusundaki yardımsever gopher’lara danışın (davet bağlantısını buradan alın).

Go diliyle ilgili tartışmalar için resmi posta listesi  *Go Nuts* ’tır.

Hataları Go issue tracker kullanarak bildirin.
