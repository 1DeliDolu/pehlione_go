
# 📌 Etkili Go 

## 🗂️ İçindekiler

* Giriş
* Örnekler
* Biçimlendirme
* Yorumlar
* İsimler

  * Paket isimleri
  * Getter’lar
  * Arayüz isimleri
  * MixedCaps
* Noktalı virgüller
* Kontrol yapıları

  * if
  * Yeniden bildirim ve yeniden atama
  * for
  * switch
  * Tür switch’i
* Fonksiyonlar

  * Çoklu dönüş değerleri
  * Adlandırılmış sonuç parametreleri
  * defer
* Veri

  * new ile ayırma
  * Kurucular ve bileşik literaller
  * make ile ayırma
  * Diziler
  * Dilimler
  * İki boyutlu dilimler
  * Map’ler
  * Yazdırma
  * append
* Başlatma

  * Sabitler
  * Değişkenler
  * init fonksiyonu
* Metotlar

  * İşaretçiler vs. değerler
* Arayüzler ve diğer türler

  * Arayüzler
  * Dönüşümler
  * Arayüz dönüşümleri ve tür doğrulamaları
  * Genellik
  * Arayüzler ve metotlar
  * Boş tanımlayıcı
  * Çoklu atamada boş tanımlayıcı
  * Kullanılmayan import’lar ve değişkenler
  * Yan etki için import
  * Arayüz kontrolleri
  * Gömme
* Eşzamanlılık

  * İletişerek paylaş
  * Goroutine’ler
  * Kanallar
  * Kanal kanalları
  * Paralelleştirme
  * Sızdıran bir tampon
* Hatalar

  * panic
  * recover
* Bir web sunucusu

## 🚀 Giriş

Go yeni bir dildir. Var olan dillerden fikirler ödünç alsa da, etkili Go programlarını akrabalarında yazılan programlardan karakter olarak farklı kılan alışılmadık özelliklere sahiptir. Bir C++ veya Java programını doğrudan Go’ya çevirmek muhtemelen tatmin edici bir sonuç üretmez—Java programları Java’da yazılır, Go’da değil. Öte yandan, probleme Go perspektifinden bakmak başarılı ama oldukça farklı bir program ortaya çıkarabilir. Başka bir deyişle, Go’yu iyi yazmak için özelliklerini ve deyimlerini anlamak önemlidir. Ayrıca Go’da isimlendirme, biçimlendirme, program inşası vb. gibi yerleşik programlama geleneklerini bilmek de önemlidir; böylece yazdığınız programlar diğer Go programcıları tarafından kolayca anlaşılır.

Bu belge, açık ve deyimsel Go kodu yazmak için ipuçları verir. Dil belirtimini, Go Turu’nu ve How to Write Go Code belgesini tamamlar; bunları önce okumalısınız.

Ocak 2022’de eklenen not: Bu belge Go’nun 2009’daki sürümü için yazılmıştır ve o zamandan beri önemli ölçüde güncellenmemiştir. Dilin kararlılığı sayesinde dili nasıl kullanacağınızı anlamak için iyi bir rehber olsa da, kütüphaneler hakkında az şey söyler ve Go ekosistemindeki derleme sistemi, test, modüller ve polimorfizm gibi büyük değişiklikler hakkında hiçbir şey söylemez. Güncellemeye dair bir plan yoktur; çünkü çok şey olmuştur ve modern Go kullanımını anlatan genişleyen bir belge, blog ve kitap seti bunu zaten iyi yapmaktadır. *Effective Go* hâlâ yararlıdır, ancak okuyucu bunun eksiksiz bir rehber olmaktan uzak olduğunu bilmelidir. Bağlam için issue 28782’ye bakın.

## 🧪 Örnekler

Go paket kaynakları yalnızca çekirdek kütüphane olarak değil, aynı zamanda dili nasıl kullanacağınızı gösteren örnekler olarak da hizmet etmek üzere tasarlanmıştır. Dahası, pek çok paket, go.dev web sitesinden doğrudan çalıştırabileceğiniz, çalışan, kendi içinde tamamlanmış yürütülebilir örnekler içerir; örneğin şu (gerekirse, açmak için “Example” kelimesine tıklayın). Bir probleme nasıl yaklaşmanız veya bir şeyin nasıl uygulanmış olabileceği konusunda sorunuz varsa, kütüphanedeki dokümantasyon, kod ve örnekler cevaplar, fikirler ve arka plan sağlayabilir.

## 🎨 Biçimlendirme

Biçimlendirme konuları en tartışmalı ama en az sonuç doğurandır. İnsanlar farklı biçimlendirme stillerine uyum sağlayabilir; ancak buna ihtiyaç duymamak daha iyidir ve herkes aynı stile uyduğunda bu konuya daha az zaman harcanır. Sorun, uzun ve buyurgan bir stil rehberi olmadan bu ütopyaya nasıl yaklaşılacağıdır.

Go ile alışılmadık bir yaklaşım benimsiyoruz ve biçimlendirme işinin çoğunu makineye bırakıyoruz. `gofmt` programı (paket düzeyinde çalışan `go fmt` olarak da bulunur; kaynak dosyası düzeyinde değil) bir Go programını okur ve yorumları koruyarak, gerektiğinde yeniden biçimlendirerek standart bir girinti ve dikey hizalama stiliyle kaynağı üretir. Yeni bir yerleşim durumunu nasıl ele alacağınızı bilmek istiyorsanız `gofmt` çalıştırın; cevap doğru görünmüyorsa programınızı (veya dosyanızı) yeniden düzenleyin (ya da `gofmt` hakkında bir hata bildirimi yapın), etrafından dolaşmaya çalışmayın.

Örneğin, bir yapının alanlarındaki yorumları hizalamak için zaman harcamaya gerek yoktur. `gofmt` bunu sizin için yapacaktır. Şu bildirimi ele alalım:

---
```go
type T struct {
    name string // name of the object
    value int // its value
}
```
---

`gofmt` sütunları hizalar:

---
```go
type T struct {
    name    string // name of the object
    value   int    // its value
}
```
---

Standart paketlerdeki tüm Go kodu `gofmt` ile biçimlendirilmiştir.

### 🧩 Biçimlendirmeye dair bazı ayrıntılar

Kısaca:

#### 📐 Girinti

Girinti için sekme kullanırız ve `gofmt` varsayılan olarak bunları üretir. Yalnızca mecbursanız boşluk kullanın.

#### 📏 Satır uzunluğu

Go’nun satır uzunluğu sınırı yoktur. Delikli kart taşmasını dert etmeyin. Bir satır çok uzun geliyorsa, satırı bölün ve fazladan bir sekmeyle girintileyin.

#### 🧷 Parantezler

Go, C ve Java’ya göre daha az paranteze ihtiyaç duyar: kontrol yapılarının (`if`, `for`, `switch`) sözdiziminde parantez yoktur. Ayrıca işlem önceliği hiyerarşisi daha kısa ve daha açıktır; bu yüzden

---
```go
x<<8 + y<<16
```
---

boşluğun ima ettiği şeyi ifade eder; diğer dillerde olduğu gibi farklı bir anlam taşımaz.

## 🗒️ Yorumlar

Go, C tarzı `/* */` blok yorumlar ve C++ tarzı `//` satır yorumları sağlar. Satır yorumları normdur; blok yorumlar çoğunlukla paket yorumları olarak görünür, ancak bir ifade içinde veya büyük kod parçalarını devre dışı bırakmak için yararlıdır.

Üst düzey bildirimlerden önce, arada boş satır olmadan görünen yorumlar, bildirimin kendisini belgeleyen yorumlar olarak kabul edilir. Bu “dokümantasyon yorumları” (*doc comments*), belirli bir Go paketi veya komutu için birincil dokümantasyondur. Dokümantasyon yorumları hakkında daha fazlası için “Go Doc Comments”a bakın.

## 🏷️ İsimler

İsimler Go’da diğer dillerde olduğu kadar önemlidir. Hatta anlamsal etkileri bile vardır: Bir ismin paket dışındaki görünürlüğü, ilk karakterinin büyük harf olup olmamasıyla belirlenir. Bu nedenle Go programlarında isimlendirme gelenekleri hakkında biraz zaman ayırmaya değer.

### 📦 Paket isimleri

Bir paket import edildiğinde, paket adı içeriğe erişmek için bir erişimci olur. Şundan sonra:

---
```go
import "bytes"
```
---

import eden paket `bytes.Buffer` hakkında konuşabilir. Paketi kullanan herkesin içeriğe aynı adla erişebilmesi faydalıdır; bu da paket adının iyi olması gerektiğini ima eder: kısa, öz, çağrışımlı. Gelenek gereği paketlere küçük harfli, tek kelimelik adlar verilir; alt çizgiye veya *mixedCaps*’e gerek olmamalıdır. Kısalığı tercih edin; çünkü paketi kullanan herkes bu adı yazacaktır. Çakışmaları en baştan dert etmeyin. Paket adı yalnızca importlar için varsayılan addır; tüm kaynak kod içinde benzersiz olmak zorunda değildir ve nadir bir çakışma durumunda import eden paket yerelde farklı bir ad seçebilir. Her hâlükârda, karışıklık nadirdir; çünkü import yolundaki dizin adı hangi paketin kullanıldığını belirler.

Bir başka gelenek, paket adının kaynak dizininin temel adı olmasıdır; `src/encoding/base64` içindeki paket `"encoding/base64"` olarak import edilir ama adı `encoding_base64` veya `encodingBase64` değil, `base64`’tür.

Bir paketi import eden, içeriğe referans vermek için adı kullanır; bu yüzden pakette dışa aktarılan adlar bu gerçeği kullanarak tekrarları önleyebilir. (`import .` notasyonunu kullanmayın; bu, paket dışından koşması gereken testleri basitleştirebilir ama aksi hâlde kaçınılmalıdır.) Örneğin `bufio` paketindeki tamponlu okuyucu türünün adı `Reader`’dır, `BufReader` değil; çünkü kullanıcılar onu `bufio.Reader` olarak görür; bu da açık ve öz bir addır. Dahası, import edilen varlıklar her zaman paket adıyla adreslendiğinden, `bufio.Reader` ile `io.Reader` çakışmaz. Benzer şekilde, `ring.Ring` için yeni örnekler üreten fonksiyon—Go’da kurucunun tanımı—normalde `NewRing` diye adlandırılabilirdi; ama `Ring` paketin dışa aktardığı tek tür olduğundan ve paket adı `ring` olduğundan, fonksiyon yalnızca `New` diye adlandırılır; müşteriler bunu `ring.New` olarak görür. İyi isimler seçmek için paket yapısını kullanın.

Bir başka kısa örnek `once.Do`’dur; `once.Do(setup)` iyi okunur ve `once.DoOrWaitUntilDone(setup)` yazmak bunu iyileştirmez. Uzun isimler otomatik olarak daha okunur yapmaz. Faydalı bir dokümantasyon yorumu çoğu zaman ekstra uzun bir isimden daha değerlidir.

### 🔎 Getter’lar

Go, getter ve setter’lar için otomatik destek sağlamaz. Kendiniz getter ve setter yazmanızda bir sakınca yoktur ve çoğu zaman uygundur; ancak getter adında `Get` kullanmak deyimsel değildir ve gerekli de değildir. `owner` adlı bir alanınız varsa (küçük harfli, dışa aktarılmayan), getter metodu `Owner` (büyük harfli, dışa aktarılan) olmalıdır; `GetOwner` değil. Dışa aktarma için büyük harf kullanımı, alanı metottan ayırmak için bir kanca sağlar. Gerekliyse setter fonksiyonu büyük olasılıkla `SetOwner` diye adlandırılır. Her iki ad da pratikte iyi okunur:

---
```go
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
```
---

### 🧩 Arayüz isimleri

Gelenek gereği, tek metotlu arayüzler, metot adı + `-er` son eki veya benzeri bir değişiklikle bir “fail (agent noun)” oluşturacak şekilde adlandırılır: `Reader`, `Writer`, `Formatter`, `CloseNotifier` vb.

Bu tür isimler çoktur ve onları ve yakaladıkları fonksiyon adlarını onurlandırmak verimlidir. `Read`, `Write`, `Close`, `Flush`, `String` vb. kanonik imzalara ve anlamlara sahiptir. Karışıklığı önlemek için, metotunuz aynı imza ve anlamı taşımıyorsa ona bu adlardan birini vermeyin. Tersine, türünüz iyi bilinen bir türdeki metotla aynı anlama sahip bir metot uyguluyorsa, aynı adı ve imzayı verin; string’e dönüştüren metodunuzu `ToString` değil `String` diye adlandırın.

### 🐪 MixedCaps

Son olarak, Go’da gelenek, çok kelimeli adlar yazmak için alt çizgi yerine `MixedCaps` veya `mixedCaps` kullanmaktır.

## ✂️ Noktalı virgüller

C gibi, Go’nun biçimsel dilbilgisi ifadeleri sonlandırmak için noktalı virgül kullanır; ancak C’nin aksine bu noktalı virgüller kaynakta görünmez. Bunun yerine *lexer* tararken basit bir kuralla otomatik olarak noktalı virgül ekler; böylece giriş metni çoğunlukla onlardan arınmış olur.

Kural şudur: Bir satır sonundan önceki son token bir tanımlayıcıysa (içinde `int` ve `float64` gibi kelimeler de vardır), bir temel literal (sayı veya string sabiti gibi) ya da şu token’lardan biriyse

---
```text
break continue fallthrough return ++ -- ) }
```
---

lexer her zaman token’dan sonra bir noktalı virgül ekler. Bu, “satır sonu bir ifadeyi bitirebilecek bir token’dan sonra geliyorsa, noktalı virgül ekle” diye özetlenebilir.

Bir noktalı virgül, kapanış süslü parantezinden hemen önce de atlanabilir; dolayısıyla şu ifade:

---
```go
go func() { for { dst <- <-src } }()
```
---

hiç noktalı virgül gerektirmez. Deyimsel Go programlarında noktalı virgüller yalnızca `for` döngüsü başlıklarında, başlatıcı, koşul ve devam öğelerini ayırmak için bulunur. Ayrıca tek satırda birden fazla ifade yazarsanız, ifadeleri ayırmak için gereklidir.

Noktalı virgül ekleme kurallarının bir sonucu olarak, bir kontrol yapısının (`if`, `for`, `switch` veya `select`) açılış süslü parantezini bir sonraki satıra koyamazsınız. Bunu yaparsanız, süslü parantezden önce noktalı virgül eklenecek ve istenmeyen etkilere yol açabilecektir. Şöyle yazın:

---
```go
if i < f() {
    g()
}
```
---

şöyle değil:

---
```go
if i < f()  // wrong!
{           // wrong!
    g()
}
```
---

## 🧭 Kontrol yapıları

Go’nun kontrol yapıları C’ninkilerle ilişkilidir ama önemli farkları vardır. `do` veya `while` döngüsü yoktur; yalnızca biraz genelleştirilmiş bir `for` vardır; `switch` daha esnektir; `if` ve `switch`, `for`daki gibi isteğe bağlı bir başlatma ifadesi kabul eder; `break` ve `continue` ifadeleri, neyi kırıp neyi sürdüreceğini belirtmek için isteğe bağlı bir etiket alabilir; ayrıca yeni kontrol yapıları vardır: bir tür switch’i ve çok yollu iletişim çoklayıcısı `select`. Sözdizimi de biraz farklıdır: parantez yoktur ve gövdeler her zaman süslü parantezle sınırlandırılmalıdır.

### ✅ if

Go’da basit bir `if` şöyle görünür:

---
```go
if x > 0 {
    return y
}
```
---

Zorunlu süslü parantezler, basit `if` ifadelerini birden çok satıra yazmayı teşvik eder. Özellikle gövde `return` veya `break` gibi bir kontrol ifadesi içeriyorsa, bunu yapmak zaten iyi bir stildir.

`if` ve `switch` bir başlatma ifadesi kabul ettiğinden, yerel bir değişken ayarlamak için bunun kullanıldığını görmek yaygındır:

---
```go
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}
```
---

Go kütüphanelerinde, bir `if` ifadesi bir sonraki ifadeye akmıyorsa—yani gövde `break`, `continue`, `goto` veya `return` ile bitiyorsa—gereksiz `else` atlanır.

---
```go
f, err := os.Open(name)
if err != nil {
    return err
}
codeUsing(f)
```
---

Bu, kodun bir dizi hata durumuna karşı koruma yapmak zorunda olduğu yaygın bir durumdur. Başarılı kontrol akışı sayfada aşağı doğru ilerleyip, hata durumlarını ortaya çıktıkça elediğinde kod iyi okunur. Hata durumları genellikle `return` ile bittiğinden, ortaya çıkan kodda `else` ifadelerine ihtiyaç kalmaz.

---
```go
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat()
if err != nil {
    f.Close()
    return err
}
codeUsing(f, d)
```
---

### 🔁 Yeniden bildirim ve yeniden atama

Bir not: Önceki bölümdeki son örnek, `:=` kısa bildirim biçiminin nasıl çalıştığına dair bir ayrıntı gösterir. `os.Open` çağrısını yapan bildirim şöyleydi:

---
```go
f, err := os.Open(name)
```
---

Bu ifade iki değişken bildirir: `f` ve `err`. Birkaç satır sonra, `f.Stat` çağrısı şöyleydi:

---
```go
d, err := f.Stat()
```
---

Bu da `d` ve `err` bildiriyormuş gibi görünür. Ancak dikkat edin: `err` her iki ifadede de yer alıyor. Bu tekrar yasaldır: `err` ilk ifade tarafından bildirilmiştir, ikincide ise yalnızca yeniden atanmıştır. Bu, `f.Stat` çağrısının yukarıda bildirilen mevcut `err` değişkenini kullandığı ve ona sadece yeni bir değer verdiği anlamına gelir.

Bir `:=` bildiriminde, bir değişken `v` daha önce bildirilmiş olsa bile görünebilir; yeter ki:

* bu bildirim, `v`’nin mevcut bildirimiyle aynı kapsamda olsun (`v` dış bir kapsamda bildirilmişse, bildirim yeni bir değişken oluşturacaktır §),
* başlatmadaki karşılık gelen değer `v`’ye atanabilir olsun ve
* bildirimin en az bir başka değişkeni yeni oluşturuyor olması gereksin.

Bu sıra dışı özellik tamamen pragmatizmdir; örneğin uzun bir `if-else` zincirinde tek bir `err` değerini kullanmayı kolaylaştırır. Bunu sıkça kullanılmış olarak göreceksiniz.

§ Burada şunu not etmeye değer: Go’da fonksiyon parametrelerinin ve dönüş değerlerinin kapsamı, sözcüksel olarak gövdeyi çevreleyen parantezlerin dışında görünseler bile fonksiyon gövdesiyle aynıdır.

### 🔄 for

Go `for` döngüsü C’ninkine benzer—ama aynı değildir. `for` ve `while`’ı birleştirir ve `do-while` yoktur. Üç biçimi vardır; yalnızca birinde noktalı virgül bulunur.

---
```go
// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }
```
---

Kısa bildirimler, indis değişkenini doğrudan döngü içinde bildirmeyi kolaylaştırır.

---
```go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```
---

Bir dizi, dilim, string veya map üzerinde dönüyorsanız ya da bir kanaldan okuyorsanız, `range` ifadesi döngüyü yönetebilir.

---
```go
for key, value := range oldMap {
    newMap[key] = value
}
```
---

`range` içindeki ilk öğeye (anahtar veya indis) ihtiyacınız varsa, ikincisini atın:

---
```go
for key := range m {
    if key.expired() {
        delete(m, key)
    }
}
```
---

Yalnızca ikinci öğeye (değere) ihtiyacınız varsa, birincisini atmak için boş tanımlayıcıyı (`_`) kullanın:

---
```go
sum := 0
for _, value := range array {
    sum += value
}
```
---

Boş tanımlayıcının, daha sonra açıklanan pek çok kullanımı vardır.

String’lerde `range` sizin için daha fazla iş yapar; UTF-8’i ayrıştırarak tek tek Unicode kod noktalarını ayırır. Hatalı kodlamalar bir bayt tüketir ve yerine geçen rune `U+FFFD` üretir. (`rune` adı (ilişkili yerleşik türle birlikte), tek bir Unicode kod noktası için Go terminolojisidir. Ayrıntılar için dil belirtimine bakın.) Şu döngü:

---
```go
for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
    fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}
```
---

şunu yazdırır:

---
```text
character U+65E5 '日' starts at byte position 0
character U+672C '本' starts at byte position 3
character U+FFFD '�' starts at byte position 6
character U+8A9E '語' starts at byte position 7
```
---

Son olarak, Go’da virgül operatörü yoktur ve `++` ile `--` ifade değil, ifadedir (statement). Dolayısıyla `for` içinde birden fazla değişkeni yürütmek istiyorsanız paralel atama kullanmalısınız (ancak bu, `++` ve `--` kullanımını engeller).

---
```go
// Reverse a
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}
```
---

### 🔀 switch

Go’nun `switch`i C’ninkinden daha geneldir. İfadeler sabit veya hatta tamsayı olmak zorunda değildir; `case`ler yukarıdan aşağı değerlendirilir ve ilk eşleşmede durur; `switch` ifadesizse `true` üzerinden switch yapar. Bu nedenle, bir `if-else-if-else` zincirini `switch` olarak yazmak mümkündür—ve deyimseldir.

---
```go
func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }
    return 0
}
```
---

Otomatik *fallthrough* yoktur, ancak `case`ler virgülle ayrılmış listeler hâlinde sunulabilir.

---
```go
func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}
```
---

Go’da bazı diğer C-benzeri diller kadar yaygın olmasalar da, `break` ifadeleri bir `switch`i erken bitirmek için kullanılabilir. Ancak bazen `switch`ten değil, çevreleyen bir döngüden çıkmak gerekir ve Go’da bu, döngüye bir etiket koyup o etikete “break” ederek yapılır. Bu örnek her iki kullanımı da gösterir:

---
```go
Loop:
    for n := 0; n < len(src); n += size {
        switch {
        case src[n] < sizeOne:
            if validateOnly {
                break
            }
            size = 1
            update(src[n])

        case src[n] < sizeTwo:
            if n+1 >= len(src) {
                err = errShortInput
                break Loop
            }
            if validateOnly {
                break
            }
            size = 2
            update(src[n] + src[n+1]<<shift)
        }
    }
```
---

Elbette `continue` ifadesi de isteğe bağlı bir etiket kabul eder, ancak yalnızca döngülere uygulanır.

Bu bölümü kapatmak için, işte iki `switch` ifadesi kullanan bir bayt dilimi karşılaştırma rutini:

---
```go
// Compare returns an integer comparing the two byte slices,
// lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b
func Compare(a, b []byte) int {
    for i := 0; i < len(a) && i < len(b); i++ {
        switch {
        case a[i] > b[i]:
            return 1
        case a[i] < b[i]:
            return -1
        }
    }
    switch {
    case len(a) > len(b):
        return 1
    case len(a) < len(b):
        return -1
    }
    return 0
}
```
---

### 🧬 Tür switch’i

Bir `switch`, bir arayüz değişkeninin dinamik türünü keşfetmek için de kullanılabilir. Böyle bir tür switch’i, parantez içinde `type` anahtar sözcüğü olan bir tür doğrulamasının sözdizimini kullanır. `switch` ifadesi içinde bir değişken bildiriliyorsa, o değişken her `case` içinde karşılık gelen türe sahip olur. Bu tür durumlarda adı yeniden kullanmak da deyimseldir; pratikte her `case`te aynı adla ama farklı türle yeni bir değişken bildirmek gibi olur.

---
```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
case bool:
    fmt.Printf("boolean %t\n", t)             // t has type bool
case int:
    fmt.Printf("integer %d\n", t)             // t has type int
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t has type *int
}
```
---

## 🧰 Fonksiyonlar

### ↩️ Çoklu dönüş değerleri

Go’nun alışılmadık özelliklerinden biri, fonksiyonların ve metotların birden fazla değer döndürebilmesidir. Bu biçim, C programlarındaki iki hantal deyimi iyileştirmek için kullanılabilir: EOF için `-1` gibi bant-içi hata dönüşleri ve adresle geçen bir argümanı değiştirerek dönüş değeri simüle etmek.

C’de bir yazma hatası, negatif bir sayı ile bildirilir ve hata kodu uçucu bir konuma saklanır. Go’da `Write` bir sayı ve bir hata döndürebilir: “Evet, bazı baytlar yazdın ama hepsini değil; çünkü cihazı doldurdun.” `os` paketindeki dosyaların `Write` metodunun imzası şöyledir:

---
```go
func (file *File) Write(b []byte) (n int, err error)
```
---

ve dokümantasyonun söylediği gibi, `n != len(b)` olduğunda yazılan bayt sayısını ve `nil` olmayan bir hata döndürür. Bu yaygın bir stildir; daha fazla örnek için hata işleme bölümüne bakın.

Benzer bir yaklaşım, bir dönüş değerini referans parametresi gibi simüle etmek için işaretçi geçme ihtiyacını ortadan kaldırır. İşte bir bayt diliminde bir konumdan bir sayı çekmek için basit bir fonksiyon; sayıyı ve bir sonraki konumu döndürür:

---
```go
func nextInt(b []byte, i int) (int, int) {
    for ; i < len(b) && !isDigit(b[i]); i++ {
    }
    x := 0
    for ; i < len(b) && isDigit(b[i]); i++ {
        x = x*10 + int(b[i]) - '0'
    }
    return x, i
}
```
---

Bunu, bir giriş dilimi `b` içindeki sayıları şöyle taramak için kullanabilirsiniz:

---
```go
    for i := 0; i < len(b); {
        x, i = nextInt(b, i)
        fmt.Println(x)
    }
```
---

### 🏷️ Adlandırılmış sonuç parametreleri

Bir Go fonksiyonunun dönüş veya sonuç “parametreleri” adlandırılabilir ve gelen parametreler gibi normal değişkenler olarak kullanılabilir. Adlandırıldıklarında, fonksiyon başladığında türlerinin sıfır değerlerine ilklendirilirler; fonksiyon argümansız bir `return` çalıştırırsa, sonuç parametrelerinin o anki değerleri dönüş değerleri olarak kullanılır.

İsimler zorunlu değildir ama kodu daha kısa ve daha açık hâle getirebilir: belgelendirmedir. `nextInt`’in sonuçlarını adlandırırsak, dönen `int`’lerin hangisinin hangisi olduğu açık olur.

---
```go
func nextInt(b []byte, pos int) (value, nextPos int) {
```
---

Adlandırılmış sonuçlar ilklendirildiği ve süssüz bir `return`a bağlı olduğu için, hem basitleştirebilir hem de açıklığa kavuşturabilir. İşte bunları iyi kullanan `io.ReadFull`’un bir sürümü:

---
```go
func ReadFull(r Reader, buf []byte) (n int, err error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    return
}
```
---

### ⏳ defer

Go’nun `defer` ifadesi, bir fonksiyon çağrısının (ertelenmiş fonksiyonun) `defer`i çalıştıran fonksiyon dönmeden hemen önce çalıştırılmasını planlar. Bu, bir fonksiyonun hangi yoldan döndüğüne bakılmaksızın serbest bırakılması gereken kaynaklar gibi durumlarla başa çıkmanın alışılmadık ama etkili bir yoludur. Kanonik örnekler bir mutex’in kilidini açmak veya bir dosyayı kapatmaktır.

---
```go
// Contents returns the file's contents as a string.
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()  // f.Close will run when we're finished.

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...) // append is discussed later.
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", err  // f will be closed if we return here.
        }
    }
    return string(result), nil // f will be closed if we return here.
}
```
---

`Close` gibi bir fonksiyonu ertelemenin iki avantajı vardır. Birincisi, dosyayı kapatmayı asla unutmayacağınızı garanti eder; fonksiyonu daha sonra yeni bir dönüş yolu ekleyecek şekilde düzenlediğinizde bu hata yapmak kolaydır. İkincisi, kapatma işleminin açmanın yakınında durmasıdır; bu, fonksiyonun sonunda olmasından çok daha anlaşılırdır.

Ertelenmiş fonksiyonun argümanları (fonksiyon bir metotsa alıcı dâhil) çağrı çalıştırıldığında değil, `defer` çalıştığı anda değerlendirilir. Bu, fonksiyon çalışırken değişkenlerin değer değiştirmesi konusunda endişeleri azaltmasının yanı sıra, tek bir `defer` çağrı noktasının birden fazla fonksiyon yürütmesini erteleyebilmesi anlamına gelir. İşte saçma bir örnek:

---
```go
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
}
```
---

Ertelenmiş fonksiyonlar LIFO sırasıyla yürütülür; bu yüzden bu kod, fonksiyon döndüğünde `4 3 2 1 0` yazdırır. Daha makul bir örnek, program boyunca fonksiyon yürütümünü izlemek için basit bir yoldur. Şöyle iki basit izleme rutini yazabiliriz:

---
```go
func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

// Use them like this:
func a() {
    trace("a")
    defer untrace("a")
    // do something....
}
```
---

`defer`e verilen argümanların `defer` çalıştığında değerlendirildiği gerçeğinden yararlanarak daha iyisini yapabiliriz. İzleme rutini, izden çıkma rutininin argümanını kurabilir. Şu örnek:

---
```go
func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}

func a() {
    defer un(trace("a"))
    fmt.Println("in a")
}

func b() {
    defer un(trace("b"))
    fmt.Println("in b")
    a()
}

func main() {
    b()
}
```
---

şunu yazdırır:

---
```text
entering: b
in b
entering: a
in a
leaving: a
leaving: b
```
---

Diğer dillerdeki blok düzeyinde kaynak yönetimine alışık programcılar için `defer` tuhaf görünebilir; ancak en ilginç ve güçlü uygulamaları, tam da blok tabanlı değil, fonksiyon tabanlı olmasından gelir. `panic` ve `recover` bölümünde bunun bir başka örneğini göreceğiz.

