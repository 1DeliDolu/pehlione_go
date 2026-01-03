# 🧪 Eğitim: Fuzzing ile Başlangıç

Bu eğitim, Go’da  *fuzzing* ’in temellerini tanıtır. *Fuzzing* ile, zafiyetleri veya çökme oluşturan girdileri bulmaya çalışmak için testinize rastgele veriler çalıştırılır. *Fuzzing* ile bulunabilecek bazı zafiyet örnekleri; SQL injection, buffer overflow, denial of service ve cross-site scripting saldırılarıdır.

Bu eğitimde, basit bir fonksiyon için bir fuzz testi yazacak, `go` komutunu çalıştıracak ve koddaki sorunları ayıklayıp düzelteceksiniz.

Bu eğitim boyunca kullanılan terminoloji için Go Fuzzing sözlüğüne bakın.

Aşağıdaki bölümlerden ilerleyeceksiniz:

* Kodunuz için bir klasör oluşturun.
* Test edilecek kodu ekleyin.
* Birim testi (unit test) ekleyin.
* Fuzz testi ekleyin.
* İki hatayı düzeltin.
* Ek kaynakları keşfedin.

> Not: Diğer eğitimler için Eğitimler’e bakın.
> Not: Go fuzzing şu anda, Go Fuzzing dokümanlarında listelenen yerleşik (built-in) türlerin bir alt kümesini destekler; gelecekte daha fazla yerleşik tür için destek eklenecektir.

---

## ✅ Ön Koşullar¶

* Go 1.18 veya üzeri bir kurulum. Kurulum talimatları için Installing Go’ya bakın.
* Kodunuzu düzenlemek için bir araç. Sahip olduğunuz herhangi bir metin düzenleyici gayet uygundur.
* Bir komut terminali. Go, Linux ve Mac’te herhangi bir terminalle; Windows’ta ise PowerShell veya cmd ile iyi çalışır.
* *Fuzzing* ’i destekleyen bir ortam. Kapsama (coverage) enstrümantasyonu ile Go fuzzing şu anda yalnızca AMD64 ve ARM64 mimarilerinde kullanılabilir.

---

## 📁 Kodunuz için bir klasör oluşturun¶

Başlamak için, yazacağınız kod için bir klasör oluşturun.

Komut istemini açın ve home dizininize geçin.

**Linux veya Mac’te:**

```bash
$ cd
```

**Windows’ta:**

```text
C:\> cd %HOMEPATH%
```

Bu eğitimin geri kalanında, komut satırı göstergesi olarak `$` kullanılacaktır. Kullandığınız komutlar Windows’ta da çalışacaktır.

Komut isteminden, `fuzz` adlı bir dizin oluşturun.

```bash
$ mkdir fuzz
$ cd fuzz
```

Kodunuzu tutacak bir modül oluşturun.

Yeni kodunuzun modül yolunu vererek `go mod init` komutunu çalıştırın.

```bash
$ go mod init example/fuzz
go: creating new go.mod: module example/fuzz
```

> Not: Üretim kodunda, kendi ihtiyaçlarınıza daha uygun, daha spesifik bir modül yolu belirlersiniz. Daha fazlası için Managing dependencies’e bakın.

Sıradaki adımda, daha sonra fuzz edeceğimiz bir string’i tersine çevirmek için basit bir kod ekleyeceksiniz.

---

## 🧩 Test edilecek kodu ekleyin¶

Bu adımda, bir string’i tersine çevirmek için bir fonksiyon ekleyeceksiniz.

### ✍️ Kodu yazın¶

Metin düzenleyicinizi kullanarak, `fuzz` dizininde `main.go` adlı bir dosya oluşturun.

`main.go` dosyasında, dosyanın en üstüne aşağıdaki package bildirimini yapıştırın.

```go
package main
```

Bağımsız bir program (bir kütüphanenin aksine) her zaman `package main` içindedir.

Package bildiriminden sonra, aşağıdaki fonksiyon bildirimini yapıştırın.

```go
func Reverse(s string) string {
    b := []byte(s)
    for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
        b[i], b[j] = b[j], b[i]
    }
    return string(b)
}
```

Bu fonksiyon bir string kabul eder, onu byte byte dolaşır ve sonunda tersine çevrilmiş string’i döndürür.

> Not: Bu kod, `golang.org/x/example` içindeki `stringutil.Reverse` fonksiyonunu temel alır.

`main.go` dosyasının en üstünde, package bildiriminden sonra, bir string’i başlatmak, tersine çevirmek, çıktıyı yazdırmak ve tekrarlamak için aşağıdaki `main` fonksiyonunu yapıştırın.

```go
func main() {
    input := "The quick brown fox jumped over the lazy dog"
    rev := Reverse(input)
    doubleRev := Reverse(rev)
    fmt.Printf("original: %q\n", input)
    fmt.Printf("reversed: %q\n", rev)
    fmt.Printf("reversed again: %q\n", doubleRev)
}
```

Bu fonksiyon birkaç `Reverse` işlemi çalıştırır, ardından çıktıyı komut satırına yazdırır. Bu, kodu çalışırken görmek ve potansiyel olarak hata ayıklamak için faydalı olabilir.

`main` fonksiyonu `fmt` paketini kullandığı için onu import etmeniz gerekir.

`main.go` dosyasının ilk satırları şöyle görünmelidir:

```go
package main

import "fmt"
```

### ▶️ Kodu çalıştırın¶

Komut satırından, `main.go` dosyasını içeren dizinde kodu çalıştırın.

```bash
$ go run .
original: "The quick brown fox jumped over the lazy dog"
reversed: "god yzal eht revo depmuj xof nworb kciuq ehT"
reversed again: "The quick brown fox jumped over the lazy dog"
```

Orijinal string’i, onu tersine çevirmenin sonucunu ve ardından tekrar tersine çevirmenin sonucunu görürsünüz; bu sonuncusu orijinal ile eşdeğerdir.

Artık kod çalıştığına göre, onu test etme zamanı.

---

## 🧪 Birim testi ekleyin¶

Bu adımda, `Reverse` fonksiyonu için temel bir birim testi yazacaksınız.

### ✍️ Kodu yazın¶

Metin düzenleyicinizi kullanarak, `fuzz` dizininde `reverse_test.go` adlı bir dosya oluşturun.

Aşağıdaki kodu `reverse_test.go` dosyasına yapıştırın.

```go
package main

import (
    "testing"
)

func TestReverse(t *testing.T) {
    testcases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {" ", " "},
        {"!12345", "54321!"},
    }
    for _, tc := range testcases {
        rev := Reverse(tc.in)
        if rev != tc.want {
                t.Errorf("Reverse: %q, want %q", rev, tc.want)
        }
    }
}
```

Bu basit test, listelenen girdi string’lerinin doğru şekilde tersine çevrileceğini doğrular.

### ▶️ Kodu çalıştırın¶

Birim testini `go test` ile çalıştırın.

```bash
$ go test
PASS
ok      example/fuzz  0.013s
```

Sırada, birim testini bir fuzz testine dönüştüreceksiniz.

---

## 🎲 Fuzz testi ekleyin¶

Birim testinin kısıtları vardır; özellikle her girdinin geliştirici tarafından teste eklenmesi gerekir.  *Fuzzing* ’in bir faydası, kodunuz için girdileri kendisinin üretmesi ve sizin düşündüğünüz test vakalarının ulaşmadığı uç durumları tespit edebilmesidir.

Bu bölümde, daha az iş ile daha fazla girdi üretebilmek için birim testini bir fuzz testine dönüştüreceksiniz.

Unit test’leri, benchmark’ları ve fuzz test’lerini aynı `*_test.go` dosyasında tutabileceğinizi unutmayın; ancak bu örnekte birim testini fuzz testine dönüştüreceksiniz.

### ✍️ Kodu yazın¶

Metin düzenleyicinizde, `reverse_test.go` içindeki birim testini aşağıdaki fuzz testi ile değiştirin.

```go
func FuzzReverse(f *testing.F) {
    testcases := []string{"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, orig string) {
        rev := Reverse(orig)
        doubleRev := Reverse(rev)
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}
```

 *Fuzzing* ’in de bazı kısıtları vardır. Birim testinizde, `Reverse` fonksiyonunun beklenen çıktısını öngörebilir ve gerçek çıktının bu beklentileri karşıladığını doğrulayabilirdiniz.

Örneğin, `Reverse("Hello, world")` test vakasında birim testi dönüşü `"dlrow ,olleH"` olarak belirtir.

 *Fuzzing* ’de ise girdiler üzerinde kontrolünüz olmadığı için beklenen çıktıyı öngöremezsiniz.

Bununla birlikte, `Reverse` fonksiyonunun fuzz testinde doğrulayabileceğiniz bazı özellikleri vardır. Bu fuzz testinde kontrol edilen iki özellik şunlardır:

* Bir string’i iki kez tersine çevirmek orijinal değeri korur.
* Tersine çevrilmiş string, geçerli UTF-8 olma durumunu korur.

Birim testi ile fuzz testi arasındaki söz dizimi farklarına dikkat edin:

* Fonksiyon `TestXxx` yerine `FuzzXxx` ile başlar ve `*testing.T` yerine `*testing.F` alır.
* `t.Run` çalıştırması görmeyi bekleyeceğiniz yerde, bunun yerine bir fuzz hedef fonksiyonu alan `f.Fuzz` vardır; bu fonksiyonun parametreleri `*testing.T` ve fuzz edilecek türlerdir. Birim testinizden gelen girdiler, `f.Add` kullanılarak *seed corpus* girdileri olarak sağlanır.

Yeni paket olan `unicode/utf8`’in import edildiğinden emin olun.

```go
package main

import (
    "testing"
    "unicode/utf8"
)
```

Birim testi fuzz testine dönüştürüldüğüne göre, testi yeniden çalıştırma zamanı.

### ▶️ Kodu çalıştırın¶

Seed girdilerin geçtiğinden emin olmak için fuzzing yapmadan fuzz testini çalıştırın.

```bash
$ go test
PASS
ok      example/fuzz  0.013s
```

Bu dosyada başka testleriniz varsa ve yalnızca fuzz testini çalıştırmak istiyorsanız `go test -run=FuzzReverse` komutunu da kullanabilirsiniz.

Rastgele üretilen string girdilerinin bir hataya sebep olup olmadığını görmek için `FuzzReverse`’ü fuzzing ile çalıştırın. Bu, `go test` komutunda yeni bir bayrak olan `-fuzz` bayrağını `Fuzz` parametresi ile ayarlayarak yürütülür. Aşağıdaki komutu kopyalayın.

```bash
$ go test -fuzz=Fuzz
```

Bir başka yararlı bayrak `-fuzztime`’dır; fuzzing’in süresini sınırlar. Örneğin aşağıdaki testte `-fuzztime 10s` belirtmek, daha erken bir hata oluşmadığı sürece testin varsayılan olarak 10 saniye geçtikten sonra çıkacağı anlamına gelir. Diğer test bayrakları için `cmd/go` dokümantasyonunun bu bölümüne bakın.

Şimdi az önce kopyaladığınız komutu çalıştırın.

```text
$ go test -fuzz=Fuzz
fuzz: elapsed: 0s, gathering baseline coverage: 0/3 completed
fuzz: elapsed: 0s, gathering baseline coverage: 3/3 completed, now fuzzing with 8 workers
fuzz: minimizing 38-byte failing input file...
--- FAIL: FuzzReverse (0.01s)
    --- FAIL: FuzzReverse (0.00s)
        reverse_test.go:20: Reverse produced invalid UTF-8 string "\x9c\xdd"

    Failing input written to testdata/fuzz/FuzzReverse/af69258a12129d6cbba438df5d5f25ba0ec050461c116f777e77ea7c9a0d217a
    To re-run:
    go test -run=FuzzReverse/af69258a12129d6cbba438df5d5f25ba0ec050461c116f777e77ea7c9a0d217a
FAIL
exit status 1
FAIL    example/fuzz  0.030s
```

Fuzzing sırasında bir hata oluştu ve soruna neden olan girdi, bir sonraki `go test` çağrısında `-fuzz` bayrağı olmasa bile çalıştırılacak olan bir *seed corpus* dosyasına yazıldı. Soruna neden olan girdiyi görmek için, `testdata/fuzz/FuzzReverse` dizinine yazılan corpus dosyasını bir metin düzenleyicide açın. Sizin *seed corpus* dosyanız farklı bir string içerebilir; ancak format aynı olacaktır.

```text
go test fuzz v1
string("泃")
```

Corpus dosyasının ilk satırı kodlama sürümünü belirtir. Sonraki her satır, corpus girdisini oluşturan her türün değerini temsil eder. Fuzz hedefi yalnızca 1 girdi aldığı için, sürümden sonra yalnızca 1 değer vardır.

`-fuzz` bayrağı olmadan `go test` komutunu tekrar çalıştırın; yeni başarısız *seed corpus* girdisi kullanılacaktır:

```text
$ go test
--- FAIL: FuzzReverse (0.00s)
    --- FAIL: FuzzReverse/af69258a12129d6cbba438df5d5f25ba0ec050461c116f777e77ea7c9a0d217a (0.00s)
        reverse_test.go:20: Reverse produced invalid string
FAIL
exit status 1
FAIL    example/fuzz  0.016s
```

Testimiz başarısız olduğuna göre, hata ayıklama zamanı.

---

## 🧯 Geçersiz string hatasını düzeltin¶

Bu bölümde, hatayı ayıklayacak ve bug’ı düzelteceksiniz.

İsterseniz devam etmeden önce biraz düşünerek sorunu kendiniz çözmeyi deneyebilirsiniz.

### 🔎 Hatayı teşhis edin¶

Bu hatayı ayıklamanın birkaç farklı yolu vardır. Metin düzenleyiciniz olarak VS Code kullanıyorsanız, araştırmak için debugger’ınızı ayarlayabilirsiniz.

Bu eğitimde, terminalinize faydalı hata ayıklama bilgileri loglayacağız.

Önce, `utf8.ValidString` için dokümanları göz önünde bulundurun.

> `ValidString`, `s`’nin tamamen geçerli UTF-8 olarak kodlanmış  *rune* ’lardan oluşup oluşmadığını bildirir.

Mevcut `Reverse` fonksiyonu string’i byte byte tersine çevirir ve problemimiz de burada yatar. Orijinal string’in UTF-8 olarak kodlanmış  *rune* ’larını korumak için, string’i byte byte değil rune rune tersine çevirmeliyiz.

Girdinin (bu durumda Çince karakter `泃`) tersine çevrildiğinde `Reverse`’ün neden geçersiz bir string ürettiğini incelemek için, ters çevrilmiş string’deki *rune* sayısını inceleyebilirsiniz.

### ✍️ Kodu yazın¶

Metin düzenleyicinizde, `FuzzReverse` içindeki fuzz hedefini aşağıdaki ile değiştirin.

```go
f.Fuzz(func(t *testing.T, orig string) {
    rev := Reverse(orig)
    doubleRev := Reverse(rev)
    t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
    if orig != doubleRev {
        t.Errorf("Before: %q, after: %q", orig, doubleRev)
    }
    if utf8.ValidString(orig) && !utf8.ValidString(rev) {
        t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
    }
})
```

Bu `t.Logf` satırı, bir hata oluşursa veya testi `-v` ile çalıştırırsanız komut satırına yazdırır; bu da bu özel problemi ayıklamanıza yardımcı olabilir.

### ▶️ Kodu çalıştırın¶

Testi `go test` ile çalıştırın.

```text
$ go test
--- FAIL: FuzzReverse (0.00s)
    --- FAIL: FuzzReverse/28f36ef487f23e6c7a81ebdaa9feffe2f2b02b4cddaa6252e87f69863046a5e0 (0.00s)
        reverse_test.go:16: Number of runes: orig=1, rev=3, doubleRev=1
        reverse_test.go:21: Reverse produced invalid UTF-8 string "\x83\xb3\xe6"
FAIL
exit status 1
FAIL    example/fuzz    0.598s
```

Tüm *seed corpus* tek byte’lık karakterlerden oluşan string’ler kullanıyordu. Ancak `泃` gibi karakterler birkaç byte gerektirebilir. Dolayısıyla string’i byte byte tersine çevirmek çok byte’lı karakterleri geçersiz hale getirir.

> Not: Go’nun string’leri nasıl ele aldığıyla ilgileniyorsanız, daha derin bir anlayış için “Strings, bytes, runes and characters in Go” blog yazısını okuyun.

Bug’ı daha iyi anladıktan sonra, `Reverse` fonksiyonundaki hatayı düzeltin.

### 🛠️ Hatayı düzeltin¶

`Reverse` fonksiyonunu düzeltmek için, string üzerinde byte’lar yerine  *rune* ’lar üzerinden dolaşalım.

#### ✍️ Kodu yazın¶

Metin düzenleyicinizde, mevcut `Reverse()` fonksiyonunu aşağıdaki ile değiştirin.

```go
func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
```

Temel fark, `Reverse`’ün artık string’deki her byte yerine her *rune* üzerinde dönmesidir. Bunun yalnızca bir örnek olduğunu ve birleştiren (combining) karakterleri doğru ele almadığını unutmayın.

#### ▶️ Kodu çalıştırın¶

Testi `go test` ile çalıştırın.

```bash
$ go test
PASS
ok      example/fuzz  0.016s
```

Test artık geçiyor.

`go test -fuzz` ile tekrar fuzz edin ve yeni bug’lar olup olmadığına bakın.

```text
$ go test -fuzz=Fuzz
fuzz: elapsed: 0s, gathering baseline coverage: 0/37 completed
fuzz: minimizing 506-byte failing input file...
fuzz: elapsed: 0s, gathering baseline coverage: 5/37 completed
--- FAIL: FuzzReverse (0.02s)
    --- FAIL: FuzzReverse (0.00s)
        reverse_test.go:33: Before: "\x91", after: "�"

    Failing input written to testdata/fuzz/FuzzReverse/1ffc28f7538e29d79fce69fef20ce5ea72648529a9ca10bea392bcff28cd015c
    To re-run:
    go test -run=FuzzReverse/1ffc28f7538e29d79fce69fef20ce5ea72648529a9ca10bea392bcff28cd015c
FAIL
exit status 1
FAIL    example/fuzz  0.032s
```

String’in iki kez tersine çevrildikten sonra orijinalden farklı olduğunu görüyoruz. Bu kez girdinin kendisi geçersiz unicode. Peki string’lerle fuzz ediyorsak bu nasıl mümkün?

Tekrar hata ayıklayalım.

---

## 🔁 Çift tersine çevirme hatasını düzeltin¶

Bu bölümde, çift tersine çevirme hatasını ayıklayacak ve bug’ı düzelteceksiniz.

İsterseniz devam etmeden önce biraz düşünerek sorunu kendiniz çözmeyi deneyebilirsiniz.

### 🔎 Hatayı teşhis edin¶

Daha önce olduğu gibi, bu başarısızlığı ayıklamak için birkaç farklı yol vardır. Bu durumda, bir debugger kullanmak harika bir yaklaşım olurdu.

Bu eğitimde, `Reverse` fonksiyonunda faydalı hata ayıklama bilgileri loglayacağız.

Hatayı fark etmek için tersine çevrilmiş string’e yakından bakın. Go’da bir string, yalnızca okunur (read only) bir byte slice’ıdır ve geçerli UTF-8 olmayan byte’lar içerebilir. Orijinal string, tek bir byte olan `'\x91'` içeren bir byte slice’ıdır. Girdi string’i `[]rune`’a dönüştürüldüğünde Go, byte slice’ı UTF-8’e kodlar ve byte’ı UTF-8 karakteri `�` ile değiştirir. Yerine konan UTF-8 karakterini girdi byte slice’ı ile karşılaştırdığımızda, açıkça eşit olmadıkları görülür.

### ✍️ Kodu yazın¶

Metin düzenleyicinizde, `Reverse` fonksiyonunu aşağıdaki ile değiştirin.

```go
func Reverse(s string) string {
    fmt.Printf("input: %q\n", s)
    r := []rune(s)
    fmt.Printf("runes: %q\n", r)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
```

Bu, string’i bir *rune* dilimine dönüştürürken nelerin yanlış gittiğini anlamamıza yardımcı olur.

### ▶️ Kodu çalıştırın¶

Bu sefer, logları incelemek için yalnızca başarısız olan testi çalıştırmak istiyoruz. Bunu yapmak için `go test -run` kullanacağız.

`FuzzXxx/testdata` içindeki belirli bir corpus girdisini çalıştırmak için, `-run` bayrağına `{FuzzTestName}/{filename}` sağlayabilirsiniz. Bu, hata ayıklarken faydalı olabilir. Bu durumda, `-run` bayrağını başarısız testin tam hash’ine eşitleyin. Terminalinizden benzersiz hash’i kopyalayıp yapıştırın; aşağıdakinden farklı olacaktır.

```text
$ go test -run=FuzzReverse/28f36ef487f23e6c7a81ebdaa9feffe2f2b02b4cddaa6252e87f69863046a5e0
input: "\x91"
runes: ['�']
input: "�"
runes: ['�']
--- FAIL: FuzzReverse (0.00s)
    --- FAIL: FuzzReverse/28f36ef487f23e6c7a81ebdaa9feffe2f2b02b4cddaa6252e87f69863046a5e0 (0.00s)
        reverse_test.go:16: Number of runes: orig=1, rev=1, doubleRev=1
        reverse_test.go:18: Before: "\x91", after: "�"
FAIL
exit status 1
FAIL    example/fuzz    0.145s
```

Girdinin geçersiz unicode olduğunu bildiğimize göre, `Reverse` fonksiyonumuzdaki hatayı düzeltelim.

### 🛠️ Hatayı düzeltin¶

Bu sorunu çözmek için, `Reverse`’e giren string geçerli UTF-8 değilse bir hata döndürelim.

#### ✍️ Kodu yazın¶

Metin düzenleyicinizde, mevcut `Reverse` fonksiyonunu aşağıdaki ile değiştirin.

```go
func Reverse(s string) (string, error) {
    if !utf8.ValidString(s) {
        return s, errors.New("input is not valid UTF-8")
    }
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r), nil
}
```

Bu değişiklik, girdi string’i geçerli UTF-8 olmayan karakterler içeriyorsa bir hata döndürür.

`Reverse` fonksiyonu artık bir hata döndürdüğü için, `main` fonksiyonunu ek hata değerini göz ardı edecek şekilde değiştirin. Mevcut `main` fonksiyonunu aşağıdaki ile değiştirin.

```go
func main() {
    input := "The quick brown fox jumped over the lazy dog"
    rev, revErr := Reverse(input)
    doubleRev, doubleRevErr := Reverse(rev)
    fmt.Printf("original: %q\n", input)
    fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
    fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}
```

Bu `Reverse` çağrıları, girdi string’i geçerli UTF-8 olduğu için `nil` hata döndürmelidir.

`errors` ve `unicode/utf8` paketlerini import etmeniz gerekecek. `main.go` içindeki import ifadesi aşağıdaki gibi görünmelidir.

```go
import (
    "errors"
    "fmt"
    "unicode/utf8"
)
```

`reverse_test.go` dosyasını, hataları kontrol edecek ve dönüşle hata üretildiyse testi atlayacak şekilde değiştirin.

```go
func FuzzReverse(f *testing.F) {
    testcases := []string {"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, orig string) {
        rev, err1 := Reverse(orig)
        if err1 != nil {
            return
        }
        doubleRev, err2 := Reverse(rev)
        if err2 != nil {
             return
        }
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}
```

`return` etmek yerine, bir fuzz girdisinin yürütmesini durdurmak için `t.Skip()` de çağırabilirsiniz.

#### ▶️ Kodu çalıştırın¶

Testi `go test` ile çalıştırın.

```bash
$ go test
PASS
ok      example/fuzz  0.019s
```

`go test -fuzz=Fuzz` ile fuzz edin; ardından birkaç saniye geçtikten sonra `ctrl-C` ile fuzzing’i durdurun. Fuzz testi, `-fuzztime` bayrağını geçmediğiniz sürece, başarısız bir girdi ile karşılaşana kadar çalışır. Varsayılan davranış, hiçbir hata oluşmazsa sonsuza kadar çalışmasıdır; süreç `ctrl-C` ile kesilebilir.

```text
$ go test -fuzz=Fuzz
fuzz: elapsed: 0s, gathering baseline coverage: 0/38 completed
fuzz: elapsed: 0s, gathering baseline coverage: 38/38 completed, now fuzzing with 4 workers
fuzz: elapsed: 3s, execs: 86342 (28778/sec), new interesting: 2 (total: 35)
fuzz: elapsed: 6s, execs: 193490 (35714/sec), new interesting: 4 (total: 37)
fuzz: elapsed: 9s, execs: 304390 (36961/sec), new interesting: 4 (total: 37)
...
fuzz: elapsed: 3m45s, execs: 7246222 (32357/sec), new interesting: 8 (total: 41)
^Cfuzz: elapsed: 3m48s, execs: 7335316 (31648/sec), new interesting: 8 (total: 41)
PASS
ok      example/fuzz  228.000s
```

Hata bulunmadıysa 30 saniye sonra çıkacak şekilde `go test -fuzz=Fuzz -fuzztime 30s` ile fuzz edin.

```text
$ go test -fuzz=Fuzz -fuzztime 30s
fuzz: elapsed: 0s, gathering baseline coverage: 0/5 completed
fuzz: elapsed: 0s, gathering baseline coverage: 5/5 completed, now fuzzing with 4 workers
fuzz: elapsed: 3s, execs: 80290 (26763/sec), new interesting: 12 (total: 12)
fuzz: elapsed: 6s, execs: 210803 (43501/sec), new interesting: 14 (total: 14)
fuzz: elapsed: 9s, execs: 292882 (27360/sec), new interesting: 14 (total: 14)
fuzz: elapsed: 12s, execs: 371872 (26329/sec), new interesting: 14 (total: 14)
fuzz: elapsed: 15s, execs: 517169 (48433/sec), new interesting: 15 (total: 15)
fuzz: elapsed: 18s, execs: 663276 (48699/sec), new interesting: 15 (total: 15)
fuzz: elapsed: 21s, execs: 771698 (36143/sec), new interesting: 15 (total: 15)
fuzz: elapsed: 24s, execs: 924768 (50990/sec), new interesting: 16 (total: 16)
fuzz: elapsed: 27s, execs: 1082025 (52427/sec), new interesting: 17 (total: 17)
fuzz: elapsed: 30s, execs: 1172817 (30281/sec), new interesting: 17 (total: 17)
fuzz: elapsed: 31s, execs: 1172817 (0/sec), new interesting: 17 (total: 17)
PASS
ok      example/fuzz  31.025s
```

Fuzzing geçti!

`-fuzz` bayrağına ek olarak, `go test`’e birkaç yeni bayrak daha eklenmiştir ve bunlar dokümantasyonda görülebilir.

Fuzzing çıktısında kullanılan terimler hakkında daha fazla bilgi için Go Fuzzing’e bakın. Örneğin, “new interesting”, mevcut fuzz test corpus’unun kod kapsamasını (code coverage) genişleten girdileri ifade eder. “new interesting” girdilerinin sayısının fuzzing başlarken keskin biçimde artması, yeni kod yolları keşfedildikçe birkaç kez sıçraması, ardından zamanla azalması beklenir.

---

## ✅ Sonuç¶

Güzel iş! Go’da *fuzzing* ile yeni tanıştınız.

Bir sonraki adım, kodunuzda fuzz etmek istediğiniz bir fonksiyon seçmek ve denemek. Eğer fuzzing kodunuzda bir bug bulursa, onu trophy case’e eklemeyi düşünün.

Herhangi bir problem yaşarsanız veya bir özellik fikriniz varsa, bir issue açın.

Özellik hakkında tartışma ve genel geri bildirim için, Gophers Slack’teki `#fuzzing` kanalına da katılabilirsiniz.

Daha fazla okuma için `go.dev/security/fuzz` adresindeki dokümantasyona göz atın.

---

## 🧾 Tamamlanmış kod¶

— `main.go` —

```go
package main

import (
    "errors"
    "fmt"
    "unicode/utf8"
)

func main() {
    input := "The quick brown fox jumped over the lazy dog"
    rev, revErr := Reverse(input)
    doubleRev, doubleRevErr := Reverse(rev)
    fmt.Printf("original: %q\n", input)
    fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
    fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}

func Reverse(s string) (string, error) {
    if !utf8.ValidString(s) {
        return s, errors.New("input is not valid UTF-8")
    }
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r), nil
}
```

— `reverse_test.go` —

```go
package main

import (
    "testing"
    "unicode/utf8"
)

func FuzzReverse(f *testing.F) {
    testcases := []string{"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc) // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, orig string) {
        rev, err1 := Reverse(orig)
        if err1 != nil {
            return
        }
        doubleRev, err2 := Reverse(rev)
        if err2 != nil {
            return
        }
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}
```

Yukarı dön
