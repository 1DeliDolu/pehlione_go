
## 🚩 Komut Satırı Bayrakları

Komut satırı bayrakları (*command-line flags*), komut satırı programları için seçenekleri belirtmenin yaygın bir yoludur. Örneğin `wc -l` komutunda `-l` bir komut satırı bayrağıdır.

---

## ▶️ Çalıştırma

```go
package main
```

Go, temel komut satırı bayraklarını ayrıştırmayı destekleyen bir `flag` paketi sağlar. Bu paketi, örnek komut satırı programımızı uygulamak için kullanacağız.

```go
import (
    "flag"
    "fmt"
)
```

```go
func main() {
```

---

## 🧩 Temel Bayrak Tanımlamaları

Temel bayrak tanımlamaları; *string*, *integer* ve *boolean* seçenekler için mevcuttur. Burada `"foo"` varsayılan değerine ve kısa bir açıklamaya sahip `word` adlı bir *string* bayrağı tanımlıyoruz. Bu `flag.String` fonksiyonu bir *string pointer* döndürür (*string* değer değil); bu pointer’ı nasıl kullanacağımızı aşağıda göreceğiz.

```go
    wordPtr := flag.String("word", "foo", "a string")
```

Bu, `word` bayrağına benzer bir yaklaşımla `numb` ve `fork` bayraklarını tanımlar.

```go
    numbPtr := flag.Int("numb", 42, "an int")
    forkPtr := flag.Bool("fork", false, "a bool")
```

Programın başka bir yerinde tanımlanmış mevcut bir `var` kullanan bir seçenek tanımlamak da mümkündür. Bayrak tanımlama fonksiyonuna bir pointer geçmemiz gerektiğini unutmayın.

```go
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")
```

---

## ✅ Bayrakları Ayrıştırma

Tüm bayraklar tanımlandıktan sonra, komut satırı ayrıştırmasını çalıştırmak için `flag.Parse()` çağrılır.

```go
    flag.Parse()
```

---

## 🧾 Ayrıştırılmış Seçenekleri Yazdırma

Burada sadece ayrıştırılmış seçenekleri ve varsa sondaki konumsal argümanları (*positional arguments*) yazdıracağız. `*wordPtr` gibi pointer’ları dereference etmemiz gerektiğine dikkat edin; böylece gerçek seçenek değerlerini alırız.

```go
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *forkPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}
```

---

## 🛠️ Derleme ve Deneme

Komut satırı bayrakları programını denemek için, önce derlemek ve ardından oluşan binary’yi doğrudan çalıştırmak en iyisidir.

```bash
$ go build command-line-flags.go
```

Önce tüm bayraklara değer vererek derlenen programı deneyin.

```bash
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []
```

---

## 🧷 Varsayılan Değerler

Bayrakları atlarsanız, otomatik olarak varsayılan değerlerini alırlar.

```bash
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []
```

---

## 📎 Sondaki Konumsal Argümanlar

Sondaki konumsal argümanlar, herhangi bir bayraktan sonra verilebilir.

```bash
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]
```

`flag` paketi, tüm bayrakların konumsal argümanlardan önce gelmesini ister (aksi halde bayraklar konumsal argüman olarak yorumlanır).

```bash
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]
```

---

## 🆘 Yardım Metni

Komut satırı programı için otomatik üretilen yardım metnini almak üzere `-h` veya `--help` bayraklarını kullanın.

```bash
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string
```

---

## ❌ Tanımsız Bayrak Hatası

`flag` paketinde belirtilmeyen bir bayrak verirseniz, program bir hata mesajı basar ve yardım metnini yeniden gösterir.

```bash
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...
```

---

## ➡️ Sonraki Örnek

Sonraki örnek: **Command-Line Subcommands**.

