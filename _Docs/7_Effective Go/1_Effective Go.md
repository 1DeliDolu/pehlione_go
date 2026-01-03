
## 🧩 Kontrol Yapıları 

Go’nun kontrol yapıları C’dekilerle ilişkilidir, ancak önemli açılardan farklılık gösterir. `do` veya `while` döngüsü yoktur; yalnızca biraz genelleştirilmiş bir `for` vardır. `switch` daha esnektir; `if` ve `switch`, `for`’daki gibi isteğe bağlı bir başlatma deyimi (*initialization statement*) kabul eder; `break` ve `continue` deyimleri, hangi yapıyı kıracağını veya devam ettireceğini belirlemek için isteğe bağlı bir etiket (*label*) alır; ayrıca *type switch* ve çoklu iletişim çoklayıcısı (*select*) gibi yeni kontrol yapıları da vardır. Sözdizimi de biraz farklıdır: parantez yoktur ve gövdeler her zaman süslü parantezlerle sınırlandırılmalıdır.

---

## 🔀 If

Go’da basit bir `if` şöyle görünür:

```go
if x > 0 {
    return y
}
```

Zorunlu süslü parantezler, basit `if` deyimlerini birden fazla satırda yazmayı teşvik eder. Özellikle gövde `return` veya `break` gibi bir kontrol deyimi içerdiğinde, bunu yapmak zaten iyi bir stildir.

`if` ve `switch` bir başlatma deyimi kabul ettiğinden, yerel bir değişkeni hazırlamak için bunun kullanıldığı sıkça görülür.

```go
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}
```

Go kütüphanelerinde, bir `if` deyimi sonraki deyime “akmıyorsa”—yani gövde `break`, `continue`, `goto` veya `return` ile bitiyorsa—gereksiz `else`’nin atlandığını görürsünüz.

```go
f, err := os.Open(name)
if err != nil {
    return err
}
codeUsing(f)
```

Bu, kodun bir dizi hata koşuluna karşı korunması gereken yaygın bir duruma örnektir. Denetimin başarılı akışı sayfanın aşağısına doğru ilerleyip hata durumlarını ortaya çıktıkça elediğinde kod iyi okunur. Hata durumları genellikle `return` deyimleriyle bittiğinden, ortaya çıkan kod `else` deyimlerine ihtiyaç duymaz.

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

## 🔁 Yeniden Bildirim ve Yeniden Atama

Bir parantez: Önceki bölümdeki son örnek, `:=` kısa bildirim biçiminin nasıl çalıştığına dair bir ayrıntıyı gösterir. `os.Open` çağrısını yapan bildirim şöyle okunur:

```go
f, err := os.Open(name)
```

Bu deyim iki değişken bildirir: `f` ve `err`. Birkaç satır sonra `f.Stat` çağrısı şöyle okunur:

```go
d, err := f.Stat()
```

Bu, `d` ve `err`’yi bildiriyormuş gibi görünür. Ancak `err`’nin her iki deyimde de göründüğüne dikkat edin. Bu tekrar yasaldır: `err` ilk deyim tarafından bildirilir, ancak ikinci deyimde yalnızca yeniden atanır. Bu, `f.Stat` çağrısının yukarıda bildirilen mevcut `err` değişkenini kullandığı ve ona yalnızca yeni bir değer verdiği anlamına gelir.

Bir `:=` bildiriminde bir `v` değişkeni, daha önce bildirilmiş olsa bile şu koşullarla görünebilir:

* Bu bildirim, `v`’nin mevcut bildirimiyle aynı kapsamda (*scope*) olmalıdır (eğer `v` dış bir kapsamda zaten bildirilmişse, bu bildirim yeni bir değişken oluşturur §),
* Başlatmadaki karşılık gelen değer `v`’ye atanabilir olmalıdır ve
* Bildirim tarafından en az bir başka değişken oluşturulmalıdır.

Bu alışılmadık özellik, tamamen pragmatizmdir; örneğin uzun bir `if-else` zincirinde tek bir `err` değerini kullanmayı kolaylaştırır. Bunu sıkça kullanıldığını göreceksiniz.

§ Burada, Go’da fonksiyon parametreleri ile dönüş değerlerinin kapsamının, sözcüksel olarak gövdeyi saran süslü parantezlerin dışında görünseler bile, fonksiyon gövdesiyle aynı olduğunu belirtmekte fayda var.

---

## 🔁 For

Go’daki `for` döngüsü C’ninkine benzer—ama aynı değildir. `for` ve `while`’ı birleştirir ve `do-while` yoktur. Üç biçimi vardır; bunlardan yalnızca biri noktalı virgül içerir.

```go
// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }
```

Kısa bildirimler, indeks değişkenini doğrudan döngü içinde bildirmeyi kolaylaştırır.

```go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```

Bir dizi (*array*), dilim (*slice*), string veya harita (*map*) üzerinde dönüyorsanız ya da bir kanaldan (*channel*) okuyorsanız, bir `range` cümleciği döngüyü yönetebilir.

```go
for key, value := range oldMap {
    newMap[key] = value
}
```

Aralıktaki yalnızca ilk öğeye (anahtar veya indeks) ihtiyacınız varsa, ikincisini bırakın:

```go
for key := range m {
    if key.expired() {
        delete(m, key)
    }
}
```

Aralıktaki yalnızca ikinci öğeye (değer) ihtiyacınız varsa, ilkini atmak için boş tanımlayıcıyı, yani alt çizgiyi kullanın:

```go
sum := 0
for _, value := range array {
    sum += value
}
```

Boş tanımlayıcının daha sonra açıklanan birçok kullanımı vardır.

String’lerde `range`, UTF-8’i ayrıştırarak tek tek Unicode kod noktalarını çıkarmak gibi daha fazla iş yapar. Hatalı kodlamalar bir bayt tüketir ve yerine koyma rünü (*replacement rune*) olan `U+FFFD` üretir. (`rune` adı (ilişkili yerleşik türle birlikte) Go terminolojisinde tek bir Unicode kod noktası anlamına gelir. Ayrıntılar için dil belirtimine bakın.) Şu döngü:

```go
for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
    fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}
```

şunu yazdırır:

```text
character U+65E5 '日' starts at byte position 0
character U+672C '本' starts at byte position 3
character U+FFFD '�' starts at byte position 6
character U+8A9E '語' starts at byte position 7
```

Son olarak, Go’da virgül operatörü yoktur ve `++` ile `--` ifadeler (*expressions*) değil, deyimlerdir (*statements*). Dolayısıyla bir `for` içinde birden fazla değişkeni ilerletmek istiyorsanız paralel atama kullanmalısınız (ancak bu `++` ve `--`’yi devre dışı bırakır).

```go
// Reverse a
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}
```

---

## 🎛️ Switch

Go’nun `switch`’i C’dekinden daha geneldir. İfadelerin sabit (*constant*) ya da hatta tamsayı olması gerekmez; `case`’ler yukarıdan aşağıya değerlendirilir ve bir eşleşme bulunana kadar ilerlenir; ayrıca `switch`’in ifadesi yoksa `true` üzerinden anahtarlanır. Bu nedenle bir `if-else-if-else` zincirini `switch` olarak yazmak mümkündür—ve *idiomatik*tir.

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

Otomatik *fall through* yoktur; ancak `case`’ler virgülle ayrılmış listeler halinde sunulabilir.

```go
func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}
```

Go’da bazı diğer C-benzeri dillere kıyasla neredeyse bu kadar yaygın olmasalar da, `break` deyimleri bir `switch`’i erken sonlandırmak için kullanılabilir. Ancak bazen `switch`’ten değil, onu saran bir döngüden çıkmak gerekir ve Go’da bu, döngüye bir etiket koyup o etikete `break` yaparak gerçekleştirilebilir. Bu örnek her iki kullanımı da gösterir.

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

Elbette `continue` deyimi de isteğe bağlı bir etiket kabul eder, ancak yalnızca döngülere uygulanır.

Bu bölümü kapatmak için, bayt dilimlerini (*byte slices*) karşılaştıran bir karşılaştırma rutini burada; iki `switch` deyimi kullanır:

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

## 🧬 Type Switch

Bir `switch`, bir arayüz (*interface*) değişkeninin dinamik türünü keşfetmek için de kullanılabilir. Böyle bir *type switch*, parantez içinde `type` anahtar sözcüğü bulunan bir tür doğrulamasının (*type assertion*) sözdizimini kullanır. `switch` ifadesinde bir değişken bildirirse, değişken her bir `case` cümlesinde karşılık gelen türe sahip olur. Bu tür durumlarda ismi yeniden kullanmak da *idiomatik*tir; böylece her bir `case` içinde aynı isimle ama farklı türle yeni bir değişken bildirilmiş olur.

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


## 🧠 Fonksiyonlar 

### 🔁 Çoklu dönüş değerleri

Go’nun alışılmadık özelliklerinden biri, fonksiyon ve metotların birden fazla değer döndürebilmesidir. Bu biçim, C programlarındaki birkaç hantal deyimi iyileştirmek için kullanılabilir: örneğin, EOF için `-1` gibi bant-içi (*in-band*) hata dönüşleri ve adres üzerinden geçirilen bir argümanı değiştirerek sonuç döndürmek.

C’de bir yazma hatası, negatif bir sayaçla bildirilir; hata kodu ise uçucu (*volatile*) bir konumda gizlenir. Go’da `Write` bir sayaç ve bir hata döndürebilir: “Evet, bazı baytlar yazdın ama hepsini değil; çünkü aygıtı doldurdun.” `os` paketindeki dosyaların `Write` metodunun imzası şöyledir:

```go
func (file *File) Write(b []byte) (n int, err error)
```

Dokümantasyonun söylediği gibi, `n != len(b)` olduğunda yazılan bayt sayısını ve `nil` olmayan bir hata döndürür. Bu yaygın bir stildir; daha fazla örnek için hata işleme bölümüne bakın.

Benzer bir yaklaşım, bir başvuru (*reference*) parametresi simüle etmek için bir dönüş değerine işaretçi geçme ihtiyacını da ortadan kaldırır. Aşağıda, bir bayt dilimindeki (*byte slice*) bir konumdan bir sayı alan, sayıyı ve bir sonraki konumu döndüren basitçe yazılmış bir fonksiyon var:

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

Bunu, bir girdi dilimi `b` içindeki sayıları şöyle taramak için kullanabilirsiniz:

```go
for i := 0; i < len(b); {
    x, i = nextInt(b, i)
    fmt.Println(x)
}
```

---

### 🏷️ İsimlendirilmiş sonuç parametreleri

Bir Go fonksiyonunun dönüş veya sonuç “parametreleri” isimlendirilebilir ve tıpkı gelen parametreler gibi normal değişkenler olarak kullanılabilir. İsimlendirildiklerinde, fonksiyon başladığında türlerinin sıfır değerlerine (*zero value*) atanırlar; fonksiyon argümansız bir `return` deyimi çalıştırırsa, sonuç parametrelerinin o anki değerleri dönüş değerleri olarak kullanılır.

İsimler zorunlu değildir; ancak kodu daha kısa ve daha net hâle getirebilirler: bir tür dokümantasyondurlar. `nextInt` fonksiyonunun sonuçlarını isimlendirirsek, dönen iki `int`’in hangisinin hangisi olduğu açıkça belli olur.

```go
func nextInt(b []byte, pos int) (value, nextPos int) {
```

İsimlendirilmiş sonuçlar, başlatıldıkları ve “süslenmemiş” (*unadorned*) `return` ile bağlı oldukları için hem sadeleştirebilir hem de açıklığa kavuşturabilir. Aşağıda, bunu iyi kullanan `io.ReadFull`’un bir sürümü var:

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

## ⏳ Defer 

Go’nun `defer` deyimi, bir fonksiyon çağrısını (ertelenmiş fonksiyon) `defer`’i çalıştıran fonksiyon dönmeden hemen önce çalıştırılmak üzere planlar. Bu, bir fonksiyon hangi yoldan dönerse dönsün serbest bırakılması gereken kaynaklar gibi durumlarla baş etmek için alışılmadık ama etkili bir yoldur. Kanonik örnekler bir mutex kilidini açmak veya bir dosyayı kapatmaktır.

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

`Close` gibi bir fonksiyon çağrısını ertelemenin iki avantajı vardır. Birincisi, dosyayı kapatmayı asla unutmayacağınızı garanti eder; bu, fonksiyonu daha sonra yeni bir dönüş yolu eklemek üzere düzenlerseniz kolayca yapılabilecek bir hatadır. İkincisi, kapatma işleminin açma işleminin yanında yer almasıdır; bu da kapanışı fonksiyonun sonuna koymaktan çok daha anlaşılırdır.

Ertelenen fonksiyonun argümanları (fonksiyon bir metotsa alıcı (*receiver*) dahil), çağrı çalıştığında değil, `defer` çalıştığında değerlendirilir. Bu, fonksiyon çalışırken değişken değerlerinin değişmesi konusunda endişeleri azaltmasının yanında, tek bir ertelenmiş çağrı noktasının birden fazla fonksiyon yürütmesini erteleyebilmesi anlamına gelir. Aşağıda saçma bir örnek var:

```go
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
}
```

Ertelenmiş fonksiyonlar LIFO sırasıyla çalıştırılır; bu yüzden bu kod, fonksiyon döndüğünde `4 3 2 1 0` yazdırır. Daha makul bir örnek, program boyunca fonksiyon yürütümünü izlemek için basit bir yoldur. Şöyle iki basit izleme rutini yazabiliriz:

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

`defer`’e verilen argümanların `defer` çalıştığında değerlendirildiği gerçeğinden yararlanarak daha iyisini yapabiliriz. İzleme rutini, izlemeyi kaldırma (*untracing*) rutinine gidecek argümanı hazırlayabilir. Bu örnek:

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

şunu yazdırır:

```text
entering: b
in b
entering: a
in a
leaving: a
leaving: b
```

Diğer dillerdeki blok düzeyi kaynak yönetimine alışkın programcılar için `defer` tuhaf görünebilir; ancak onun en ilginç ve güçlü uygulamaları, tam da blok tabanlı değil fonksiyon tabanlı olmasından gelir. *panic* ve *recover* bölümünde, olanaklarına dair başka bir örnek daha göreceğiz.

---

## 🧱 Veri 

### 🧮 `new` ile ayırma (*allocation*)

Go’nun iki ayırma ilkel (*primitive*) fonksiyonu vardır: yerleşik `new` ve `make`. Farklı işler yaparlar ve farklı türlere uygulanırlar; bu kafa karıştırıcı olabilir, ancak kurallar basittir. Önce `new`’dan bahsedelim. `new`, bellek ayırır; fakat bazı diğer dillerdeki benzerlerinin aksine belleği başlatmaz, yalnızca sıfırlar. Yani `new(T)`, `T` türünden yeni bir öğe için sıfırlanmış depolama alanı ayırır ve onun adresini, `*T` türünden bir değer olarak döndürür. Go terminolojisinde, `T` türünün yeni ayrılmış sıfır değerine işaretçi döndürür.

`new`’un döndürdüğü bellek sıfırlandığı için, veri yapılarınızı tasarlarken her türün sıfır değerinin ek başlatma olmadan kullanılabilir olmasını sağlamak faydalıdır. Bu, veri yapısını kullanan kişinin `new` ile bir tane oluşturup hemen işe başlaması anlamına gelir. Örneğin, `bytes.Buffer` dokümantasyonu “`Buffer` için sıfır değer, kullanıma hazır boş bir arabellek” der. Benzer şekilde `sync.Mutex`, açık bir kurucu (*constructor*) ya da `Init` metodu içermez. Bunun yerine, `sync.Mutex` için sıfır değer, kilidi açılmış bir mutex olarak tanımlanır.

“Sıfır değer faydalıdır” özelliği geçişkendir (*transitively*). Şu tür bildirimini düşünün:

```go
type SyncedBuffer struct {
    lock   sync.Mutex
    buffer bytes.Buffer
}
```

`SyncedBuffer` türündeki değerler de, ayırma ya da yalnızca bildirim sonrasında, ek bir düzenleme olmadan hemen kullanılmaya hazırdır. Aşağıdaki kesitte hem `p` hem de `v`, daha fazla ayarlama olmadan doğru şekilde çalışacaktır:

```go
p := new(SyncedBuffer)  // type *SyncedBuffer
var v SyncedBuffer      // type  SyncedBuffer
```

---

### 🏗️ Kurucular ve bileşik sabitler (*composite literals*)

Bazen sıfır değer yeterince iyi değildir ve başlatıcı bir kurucu gerekir; aşağıdaki örnek `os` paketinden türetilmiştir:

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}
```

Burada çok fazla kalıp kod (*boilerplate*) vardır. Bunu, değerlendirildiği her seferinde yeni bir örnek oluşturan bir ifade olan bileşik sabit (*composite literal*) kullanarak sadeleştirebiliriz:

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &f
}
```

C’den farklı olarak, yerel bir değişkenin adresini döndürmek tamamen sorun değildir; değişkenle ilişkili depolama alanı fonksiyon döndükten sonra da yaşamaya devam eder. Aslında, bir bileşik sabitin adresini almak, her değerlendirilişinde yeni bir örnek ayırır; bu yüzden son iki satırı birleştirebiliriz:

```go
return &File{fd, name, nil, 0}
```

Bir bileşik sabitin alanları sırayla yerleştirilir ve hepsi mevcut olmalıdır. Ancak elemanları açıkça `alan:değer` çiftleri olarak etiketleyerek, başlatıcılar herhangi bir sırada görünebilir; eksik olanlar kendi sıfır değerlerinde bırakılır. Dolayısıyla şöyle diyebiliriz:

```go
return &File{fd: fd, name: name}
```

Sınır durum olarak, bir bileşik sabit hiç alan içermezse, tür için bir sıfır değer oluşturur. `new(File)` ve `&File{}` ifadeleri eşdeğerdir.

Bileşik sabitler diziler (*arrays*), dilimler (*slices*) ve haritalar (*maps*) için de oluşturulabilir; burada alan etiketleri uygun şekilde indeksler veya harita anahtarlarıdır. Bu örneklerde, `Enone`, `Eio` ve `Einval` değerleri farklı olduğu sürece (değerleri ne olursa olsun) başlatmalar çalışır:

```go
a := [...]string   {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
s := []string      {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
```

---

### 🛠️ `make` ile ayırma (*allocation*)

Ayırmaya geri dönelim. Yerleşik `make(T, args)` fonksiyonu `new(T)`’den farklı bir amaca hizmet eder. Yalnızca dilimleri (*slices*), haritaları (*maps*) ve kanalları (*channels*) oluşturur ve `T` türünden başlatılmış (sıfırlanmış değil) bir değer döndürür (`*T` değil). Ayrımın nedeni şudur: bu üç tür, perde arkasında kullanımdan önce başlatılması gereken veri yapılarına referansları temsil eder. Örneğin bir dilim, (bir dizinin içindeki) veriye işaretçi, uzunluk ve kapasite içeren üç öğeli bir tanımlayıcıdır ve bu öğeler başlatılana kadar dilim `nil`’dir. Dilimler, haritalar ve kanallar için `make`, iç veri yapısını başlatır ve değeri kullanıma hazırlar. Örneğin:

```go
make([]int, 10, 100)
```

100 `int`’ten oluşan bir dizi ayırır ve ardından dizinin ilk 10 elemanına işaret eden, uzunluğu 10 ve kapasitesi 100 olan bir dilim yapısı oluşturur. (Bir dilim oluştururken kapasite atlanabilir; daha fazla bilgi için dilimler bölümüne bakın.) Buna karşılık, `new([]int)` yeni ayrılmış, sıfırlanmış bir dilim yapısına işaretçi döndürür; yani `nil` bir dilim değerine işaretçi.

Bu örnekler `new` ve `make` arasındaki farkı gösterir:

```go
var p *[]int = new([]int)       // dilim yapısını ayırır; *p == nil; nadiren kullanışlı
var v []int = make([]int, 100)  // dilim v artık 100 elemanlı yeni bir diziye referans verir
```

```go
// Gereksiz derecede karmaşık:
var p *[]int = new([]int)
*p = make([]int, 100, 100)
```

```go
// İdiomatik:
v := make([]int, 100)
```

`make`’in yalnızca haritalara, dilimlere ve kanallara uygulandığını ve işaretçi döndürmediğini unutmayın. Açık bir işaretçi elde etmek için `new` ile ayırın veya bir değişkenin adresini açıkça alın.

---

## 🧱 Diziler (*arrays*) 

Diziler, belleğin ayrıntılı yerleşimini planlarken yararlıdır ve bazen ayırmayı (*allocation*) önlemeye yardımcı olabilir; ancak esas olarak bir sonraki bölümün konusu olan dilimler için bir yapı taşıdır. Bu konuya temel oluşturmak için, diziler hakkında birkaç kelime.

Go’da dizilerin çalışma biçimi ile C’deki arasında büyük farklar vardır. Go’da:

* Diziler değerlerdir. Bir diziyi diğerine atamak tüm elemanları kopyalar.
* Özellikle, bir diziyi bir fonksiyona geçirirseniz, fonksiyon dizinin bir kopyasını alır; ona işaretçi almaz.
* Bir dizinin boyutu türünün parçasıdır. `[10]int` ve `[20]int` türleri farklıdır.
* Değer olma özelliği yararlı olabilir ama pahalı da olabilir; C-benzeri davranış ve verimlilik istiyorsanız, diziye işaretçi geçebilirsiniz.

```go
func Sum(a *[3]float64) (sum float64) {
    for _, v := range *a {
        sum += v
    }
    return
}

array := [...]float64{7.0, 8.5, 9.1}
x := Sum(&array)  // Açık adres alma operatörüne dikkat edin
```

Ama bu stil bile *idiomatik* Go değildir. Bunun yerine dilimleri kullanın.

---

## 🧩 Dilimler (*slices*) 

Dilimler, dizileri sararak veri dizilerine daha genel, güçlü ve kullanışlı bir arayüz sağlar. Dönüşüm matrisleri gibi açık boyutlu öğeler dışında, Go’daki çoğu dizi programlama, basit diziler yerine dilimlerle yapılır.

Dilimler, alttaki bir diziye referans tutar ve bir dilimi diğerine atarsanız ikisi de aynı diziye referans verir. Bir fonksiyon bir dilim argümanı alırsa, dilimin elemanlarında yaptığı değişiklikler çağırana görünür; bu, alttaki diziye işaretçi geçirmekle benzerlik gösterir. Dolayısıyla bir `Read` fonksiyonu, bir işaretçi ve bir sayaç yerine bir dilim argümanı kabul edebilir; dilimin uzunluğu, ne kadar veri okunacağına dair üst sınırı belirler. `os` paketindeki `File` türünün `Read` metodunun imzası şöyledir:

```go
func (f *File) Read(buf []byte) (n int, err error)
```

Metot, okunan bayt sayısını ve varsa bir hata değerini döndürür. Daha büyük bir arabellek `buf` içindeki ilk 32 bayta okumak için arabelleği dilimleyin (burada dilimleme fiil olarak kullanılıyor):

```go
n, err := f.Read(buf[0:32])
```

Böyle dilimleme yaygındır ve verimlidir. Hatta verimliliği bir kenara bırakırsak, aşağıdaki kesit de arabelleğin ilk 32 baytını okur:

```go
var n int
var err error
for i := 0; i < 32; i++ {
    nbytes, e := f.Read(buf[i:i+1])  // Bir bayt oku.
    n += nbytes
    if nbytes == 0 || e != nil {
        err = e
        break
    }
}
```

Bir dilimin uzunluğu, alttaki dizinin sınırları içinde kaldığı sürece değiştirilebilir; yalnızca onu kendisinin bir dilimine atayın. Bir dilimin kapasitesi, yerleşik `cap` fonksiyonuyla erişilir ve dilimin alabileceği azami uzunluğu bildirir. Aşağıda, bir dilime veri eklemek için bir fonksiyon var. Veri kapasiteyi aşarsa, dilim yeniden ayırılır. Ortaya çıkan dilim döndürülür. Fonksiyon, `len` ve `cap`’in `nil` dilime uygulandığında yasal olduğu ve `0` döndürdüğü gerçeğini kullanır.

```go
func Append(slice, data []byte) []byte {
    l := len(slice)
    if l+len(data) > cap(slice) {  // yeniden ayır
        // Gelecekteki büyüme için gerekenin iki katını ayır.
        newSlice := make([]byte, (l+len(data))*2)
        // copy fonksiyonu önceden bildirilmiştir ve her dilim türü için çalışır.
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0 : l+len(data)]
    copy(slice[l:], data)
    return slice
}
```

Sonrasında dilimi döndürmemiz gerekir; çünkü `Append`, `slice`’ın elemanlarını değiştirebilse de, `slice`’ın kendisi (işaretçi, uzunluk ve kapasiteyi tutan çalışma zamanı veri yapısı) değer olarak geçirilir.

Bir dilime ekleme fikri o kadar kullanışlıdır ki, `append` yerleşik fonksiyonuyla yakalanmıştır. Ancak bu fonksiyonun tasarımını anlamak için biraz daha bilgiye ihtiyacımız var; bu yüzden daha sonra tekrar döneceğiz.

---

## 🧮 İki boyutlu dilimler 

Go’nun dizileri ve dilimleri tek boyutludur. 2B dizi veya dilim eşdeğerini oluşturmak için, aşağıdaki gibi bir “dizilerin dizisi” veya “dilimlerin dilimi” tanımlamak gerekir:

```go
type Transform [3][3]float64  // 3x3 bir dizi; aslında dizilerin dizisi.
type LinesOfText [][]byte     // Bayt dilimlerinden oluşan bir dilim.
```

Dilimler değişken uzunluklu olduğu için, her iç dilimin farklı bir uzunluğa sahip olması mümkündür. Bu, `LinesOfText` örneğimizde olduğu gibi yaygın bir durum olabilir: her satırın bağımsız bir uzunluğu vardır.

```go
text := LinesOfText{
    []byte("Now is the time"),
    []byte("for all good gophers"),
    []byte("to bring some fun to the party."),
}
```

Bazen 2B bir dilim ayırmak gerekir; örneğin piksel tarama satırlarını işlerken bu durum ortaya çıkabilir. Bunu başarmanın iki yolu vardır. Biri her dilimi bağımsız ayırmaktır; diğeri tek bir dizi ayırıp tek tek dilimleri onun içine işaret ettirmektir. Hangisinin kullanılacağı uygulamanıza bağlıdır. Dilimler büyüyebilir ya da küçülebilirse, bir sonraki satırın üzerine yazmayı önlemek için bağımsız ayırmak gerekir; değilse, tek bir ayırmayla nesneyi kurmak daha verimli olabilir. Referans olması için, iki yöntemin de taslakları burada. Önce satır satır:

```go
// Üst seviye dilimi ayır.
picture := make([][]uint8, YSize) // Y’nin her bir birimi için bir satır.
// Satırlar üzerinde döngü kur; her satır için dilimi ayır.
for i := range picture {
    picture[i] = make([]uint8, XSize)
}
```

Şimdi tek bir ayırma ile, satırlara dilimlenmiş biçimde:

```go
// Üst seviye dilimi ayır, öncekiyle aynı.
picture := make([][]uint8, YSize) // Y’nin her bir birimi için bir satır.
// Tüm pikselleri tutacak tek büyük bir dilim ayır.
pixels := make([]uint8, XSize*YSize) // picture [][]uint8 olsa da türü []uint8’tir.
// Satırlar üzerinde döngü kur; her satırı kalan pixels diliminin önünden dilimle.
for i := range picture {
    picture[i], pixels = pixels[:XSize], pixels[XSize:]
}
```


## 🗺️ Haritalar 

Haritalar (*map*), bir türden değerleri (anahtar, *key*) başka bir türden değerlerle (öğe ya da değer, *element/value*) ilişkilendiren, kullanışlı ve güçlü yerleşik bir veri yapısıdır. Anahtar, eşitlik operatörünün tanımlı olduğu herhangi bir tür olabilir; örneğin tamsayılar, kayan noktalı ve karmaşık sayılar, string’ler, işaretçiler (*pointers*), arayüzler (*interfaces*; dinamik tür eşitliği desteklediği sürece), yapılar (*structs*) ve diziler (*arrays*). Dilimler (*slices*) harita anahtarı olarak kullanılamaz; çünkü onlar için eşitlik tanımlı değildir. Dilimler gibi, haritalar da altta yatan bir veri yapısına referans tutar. Bir haritayı içeriğini değiştiren bir fonksiyona geçirirseniz, değişiklikler çağırana görünür.

Haritalar, iki nokta üst üste ile ayrılmış anahtar-değer çiftleriyle olağan bileşik sabit (*composite literal*) sözdizimi kullanılarak oluşturulabilir; bu yüzden başlatma sırasında kurmaları kolaydır.

```go
var timeZone = map[string]int{
    "UTC":  0*60*60,
    "EST": -5*60*60,
    "CST": -6*60*60,
    "MST": -7*60*60,
    "PST": -8*60*60,
}
```

Harita değerlerine atama ve haritadan değer çekme, diziler ve dilimler için yapılanla sözdizimsel olarak tamamen aynıdır; tek fark, indeksin tamsayı olmasının gerekmemesidir.

```go
offset := timeZone["EST"]
```

Haritada bulunmayan bir anahtarla değer çekmeye çalışmak, girişlerin türü için sıfır değeri (*zero value*) döndürür. Örneğin harita tamsayılar içeriyorsa, var olmayan bir anahtarı aramak `0` döndürür. Bir küme (*set*), değer türü `bool` olan bir harita olarak uygulanabilir. Değeri kümeye koymak için harita girişini `true` yapın; sonra basit indeksleme ile test edin.

```go
attended := map[string]bool{
    "Ann": true,
    "Joe": true,
    ...
}

if attended[person] { // person haritada yoksa false olur
    fmt.Println(person, "was at the meeting")
}
```

Bazen eksik bir giriş ile sıfır değerini ayırt etmek gerekir. `"UTC"` için bir giriş mi var, yoksa bu `0` değeri, haritada hiç olmadığı için mi? Bunu çoklu atamanın bir biçimiyle ayırt edebilirsiniz.

```go
var seconds int
var ok bool
seconds, ok = timeZone[tz]
```

Açık nedenlerle buna “comma ok” deyimi (*idiom*) denir. Bu örnekte, `tz` mevcutsa `seconds` uygun şekilde ayarlanır ve `ok` `true` olur; değilse `seconds` sıfır olur ve `ok` `false` olur. İşte bunu hoş bir hata raporuyla birleştiren bir fonksiyon:

```go
func offset(tz string) int {
    if seconds, ok := timeZone[tz]; ok {
        return seconds
    }
    log.Println("unknown time zone:", tz)
    return 0
}
```

Gerçek değeri önemsemeden yalnızca haritada varlığını test etmek için, değer için normal değişken yerine boş tanımlayıcıyı (`_`) kullanabilirsiniz.

```go
_, present := timeZone[tz]
```

Bir harita girişini silmek için, argümanları harita ve silinecek anahtar olan yerleşik `delete` fonksiyonunu kullanın. Anahtar zaten yoksa bile bunu yapmak güvenlidir.

```go
delete(timeZone, "PDT")  // Now on Standard Time
```

---

## 🖨️ Yazdırma 

Go’da biçimli yazdırma (*formatted printing*), C’nin `printf` ailesine benzer bir stil kullanır; ancak daha zengin ve daha geneldir. Bu fonksiyonlar `fmt` paketindedir ve adları büyük harflidir: `fmt.Printf`, `fmt.Fprintf`, `fmt.Sprintf` vb. String fonksiyonları (`Sprintf` vb.) sağlanan bir arabelleği doldurmak yerine bir string döndürür.

Bir biçim string’i sağlamak zorunda değilsiniz. `Printf`, `Fprintf` ve `Sprintf` için, örneğin `Print` ve `Println` gibi başka birer fonksiyon çifti daha vardır. Bu fonksiyonlar biçim string’i almaz; bunun yerine her argüman için varsayılan bir biçim üretir. `Println` sürümleri ayrıca argümanlar arasına bir boşluk koyar ve sona yeni satır ekler; `Print` sürümleri ise yalnızca, iki yandaki işlenenlerden (*operand*) hiçbiri string değilse boşluk ekler. Bu örnekte her satır aynı çıktıyı üretir.

```go
fmt.Printf("Hello %d\n", 23)
fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
fmt.Println("Hello", 23)
fmt.Println(fmt.Sprint("Hello ", 23))
```

Biçimli yazdırma fonksiyonları `fmt.Fprint` ve benzerleri ilk argüman olarak `io.Writer` arayüzünü uygulayan herhangi bir nesneyi alır; `os.Stdout` ve `os.Stderr` değişkenleri bunun tanıdık örnekleridir.

Burada işler C’den ayrışmaya başlar. İlk olarak, `%d` gibi sayısal biçimler işaret (*signedness*) veya boyut için bayrak (*flag*) almaz; bunun yerine yazdırma rutinleri bu özellikleri belirlemek için argümanın türünü kullanır.

```go
var x uint64 = 1<<64 - 1
fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))
```

şunu yazdırır:

```text
18446744073709551615 ffffffffffffffff; -1 -1
```

Sadece varsayılan dönüşümü istiyorsanız (ör. tamsayılar için ondalık), “value” anlamına gelen kapsayıcı (*catchall*) biçim `%v`’yi kullanabilirsiniz; sonuç, `Print` ve `Println`’in üreteceği çıktının aynısıdır. Üstelik bu biçim, diziler, dilimler, yapılar ve haritalar dahil herhangi bir değeri yazdırabilir. İşte önceki bölümde tanımlanan saat dilimi haritası için bir yazdırma deyimi:

```go
fmt.Printf("%v\n", timeZone)  // veya sadece fmt.Println(timeZone)
```

bu çıktıyı verir:

```text
map[CST:-21600 EST:-18000 MST:-25200 PST:-28800 UTC:0]
```

Haritalar için, `Printf` ve benzerleri çıktıyı anahtara göre sözlük sırasıyla (*lexicographically*) sıralar.

Bir yapı (*struct*) yazdırırken, değiştirilmiş `%+v` biçimi, alanları adlarıyla etiketler; herhangi bir değer için alternatif `%#v` biçimi ise değeri tam Go sözdizimiyle yazdırır.

```go
type T struct {
    a int
    b float64
    c string
}
t := &T{ 7, -2.35, "abc\tdef" }
fmt.Printf("%v\n", t)
fmt.Printf("%+v\n", t)
fmt.Printf("%#v\n", t)
fmt.Printf("%#v\n", timeZone)
```

şunu yazdırır:

```text
&{7 -2.35 abc   def}
&{a:7 b:-2.35 c:abc     def}
&main.T{a:7, b:-2.35, c:"abc\tdef"}
map[string]int{"CST":-21600, "EST":-18000, "MST":-25200, "PST":-28800, "UTC":0}
```

(Ampersand’lara dikkat edin.) Tırnaklı string biçimi, `string` veya `[]byte` türündeki bir değere uygulandığında `%q` ile de kullanılabilir. Alternatif `%#q` biçimi, mümkünse bunun yerine backquote kullanır. (`%q` biçimi ayrıca tamsayılara ve `rune`’lara da uygulanır; tek tırnaklı bir `rune` sabiti üretir.) `%x` de, tamsayıların yanı sıra string’lerde, bayt dizilerinde ve bayt dilimlerinde çalışır; uzun bir onaltılık string üretir ve biçimde boşluk varsa (`% x`) baytların arasına boşluk koyar.

Bir diğer kullanışlı biçim `%T`’dir; bir değerin türünü yazdırır.

```go
fmt.Printf("%T\n", timeZone)
```

şunu yazdırır:

```text
map[string]int
```

Özel bir tür için varsayılan biçimi kontrol etmek istiyorsanız, tek gereksinim tür üzerinde `String() string` imzasına sahip bir metot tanımlamaktır. Basit `T` türümüz için bu şöyle görünebilir:

```go
func (t *T) String() string {
    return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}
fmt.Printf("%v\n", t)
```

şu biçimde yazdırmak için:

```text
7/-2.35/"abc\tdef"
```

(`T` türündeki değerleri ve `T` işaretçilerini yazdırmanız gerekiyorsa, `String` için alıcı (*receiver*) değer türünde olmalıdır; bu örnek işaretçi kullandı çünkü bu, yapı türleri için daha verimli ve *idiomatik*tir. Daha fazla bilgi için aşağıdaki “pointers vs. value receivers” bölümüne bakın.)

`String` metodumuz `Sprintf` çağırabilir; çünkü yazdırma rutinleri tamamen yeniden girişli (*reentrant*)dir ve bu şekilde sarılabilir. Ancak bu yaklaşımın önemli bir ayrıntısı vardır: alıcıyı (*receiver*) doğrudan string olarak yazdırmaya çalışacak ve böylece `String` metodunu sonsuza dek tekrar çağıracak biçimde `Sprintf` çağırarak bir `String` metodu kurmayın. Bu, aşağıdaki örneğin gösterdiği gibi yaygın ve kolay yapılan bir hatadır.

```go
type MyString string

func (m MyString) String() string {
    return fmt.Sprintf("MyString=%s", m) // Error: will recur forever.
}
```

Bunu düzeltmek de kolaydır: argümanı, metodu olmayan temel `string` türüne dönüştürün.

```go
type MyString string
func (m MyString) String() string {
    return fmt.Sprintf("MyString=%s", string(m)) // OK: note conversion.
}
```

Başlatma (*initialization*) bölümünde bu özyinelemeyi (*recursion*) önleyen başka bir teknik daha göreceğiz.

Başka bir yazdırma tekniği, bir yazdırma rutininin argümanlarını doğrudan başka bir rutine aktarmaktır. `Printf` imzası, son argümanı olarak `...interface{}` türünü kullanır; bu, biçim string’inden sonra keyfi sayıda (ve keyfi türde) parametre gelebileceğini belirtir.

```go
func Printf(format string, v ...interface{}) (n int, err error) {
```

`Printf` içinde, `v` aslında `[]interface{}` türünden bir değişken gibi davranır; ancak başka bir değişken-argümanlı (*variadic*) fonksiyona geçirilirse, normal bir argüman listesi gibi davranır. Daha önce kullandığımız `log.Println` fonksiyonunun uygulaması burada. Gerçek biçimlendirme için argümanlarını doğrudan `fmt.Sprintln`’e geçirir.

```go
// Println prints to the standard logger in the manner of fmt.Println.
func Println(v ...interface{}) {
    std.Output(2, fmt.Sprintln(v...))  // Output takes parameters (int, string)
}
```

İç içe çağrıda `v`’nin ardından `...` yazarak, derleyiciye `v`’yi bir argüman listesi olarak ele almasını söyleriz; aksi halde `v`’yi tek bir dilim argümanı olarak geçirirdi.

Yazdırma hakkında burada kapsadığımızdan daha fazlası vardır. Ayrıntılar için `fmt` paketi `godoc` dokümantasyonuna bakın.

Bu arada, `...` parametresi belirli bir türde olabilir; örneğin bir tamsayı listesinden en küçüğünü seçen `Min` fonksiyonu için `...int`:

```go
func Min(a ...int) int {
    min := int(^uint(0) >> 1)  // largest int
    for _, i := range a {
        if i < min {
            min = i
        }
    }
    return min
}
```

---

## ➕ Append 

Artık yerleşik `append` fonksiyonunun tasarımını açıklamak için gereken eksik parçaya sahibiz. `append` imzası, yukarıdaki kendi yazdığımız `Append` fonksiyonundan farklıdır. Şematik olarak şöyledir:

```go
func append(slice []T, elements ...T) []T
```

Burada `T`, herhangi bir verilen tür için bir yer tutucudur. Çağıranın belirlediği `T` türüne göre çalışan bir Go fonksiyonunu gerçekten yazamazsınız. `append`’in yerleşik olmasının nedeni budur: derleyiciden destek gerekir.

`append`’in yaptığı şey, öğeleri dilimin sonuna ekleyip sonucu döndürmektir. Sonuç döndürülmelidir; çünkü bizim elle yazdığımız `Append`’de olduğu gibi alttaki dizi değişebilir. Şu basit örnek:

```go
x := []int{1,2,3}
x = append(x, 4, 5, 6)
fmt.Println(x)
```

`[1 2 3 4 5 6]` yazdırır. Dolayısıyla `append`, biraz `Printf` gibi çalışır; keyfi sayıda argüman toplar.

Peki bir dilime bir dilim eklemek isteseydik—yani bizim `Append`’imizin yaptığı gibi? Kolay: `Output` çağrısında yaptığımız gibi, çağrı noktasında `...` kullanın. Bu parça, yukarıdakiyle aynı çıktıyı üretir:

```go
x := []int{1,2,3}
y := []int{4,5,6}
x = append(x, y...)
fmt.Println(x)
```

Bu `...` olmadan derlenmezdi; çünkü türler yanlış olurdu: `y`, `int` türünde değildir.


## 🧱 Başlatma 

Go’da başlatma, yüzeysel olarak C veya C++’taki başlatmadan çok farklı görünmese de, daha güçlüdür. Karmaşık yapılar başlatma sırasında inşa edilebilir ve başlatılmış nesneler arasındaki sıralama sorunları—farklı paketler arasında bile—doğru şekilde ele alınır.

---

## 🧮 Sabitler 

Go’daki sabitler (*constants*) tam anlamıyla sabittir. Fonksiyonların içinde yerel (*local*) olarak tanımlansalar bile derleme zamanında oluşturulurlar ve yalnızca sayılar, karakterler (*rune*’lar), string’ler veya boolean’lar olabilirler. Derleme zamanı kısıtlaması nedeniyle, onları tanımlayan ifadelerin sabit ifadeler olması gerekir; yani derleyici tarafından değerlendirilebilir olmalıdır. Örneğin, `1<<3` bir sabit ifadedir; ancak `math.Sin(math.Pi/4)` değildir, çünkü `math.Sin` fonksiyon çağrısı çalışma zamanında gerçekleşmelidir.

Go’da numaralandırılmış sabitler (*enumerated constants*), `iota` numaralandırıcısı (*enumerator*) kullanılarak oluşturulur. `iota` bir ifadenin parçası olabildiği ve ifadeler örtük olarak tekrar edilebildiği için, karmaşık değer kümeleri oluşturmak kolaydır.

```go
type ByteSize float64

const (
    _           = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)
```

Herhangi bir kullanıcı tanımlı türe `String` gibi bir metot bağlayabilmek, keyfi değerlerin yazdırma için kendilerini otomatik olarak biçimlendirmesini mümkün kılar. Bunu en sık `struct`’lara uygulanmış hâliyle görseniz de, bu teknik `ByteSize` gibi kayan noktalı türler gibi skaler türler için de faydalıdır.

```go
func (b ByteSize) String() string {
    switch {
    case b >= YB:
        return fmt.Sprintf("%.2fYB", b/YB)
    case b >= ZB:
        return fmt.Sprintf("%.2fZB", b/ZB)
    case b >= EB:
        return fmt.Sprintf("%.2fEB", b/EB)
    case b >= PB:
        return fmt.Sprintf("%.2fPB", b/PB)
    case b >= TB:
        return fmt.Sprintf("%.2fTB", b/TB)
    case b >= GB:
        return fmt.Sprintf("%.2fGB", b/GB)
    case b >= MB:
        return fmt.Sprintf("%.2fMB", b/MB)
    case b >= KB:
        return fmt.Sprintf("%.2fKB", b/KB)
    }
    return fmt.Sprintf("%.2fB", b)
}
```

`YB` ifadesi `1.00YB` olarak yazdırılırken, `ByteSize(1e13)` `9.09TB` olarak yazdırılır.

Burada `Sprintf`’in `ByteSize`’ın `String` metodunu uygulamak için güvenli kullanımı (sonsuz tekrar etmeyi önlemesi), bir dönüştürmeden dolayı değil, `%f` ile `Sprintf` çağırmasından kaynaklanır; çünkü bu bir string biçimi değildir: `Sprintf`, `String` metodunu yalnızca bir string istediğinde çağırır ve `%f` kayan noktalı bir değer ister.

---

## 🧷 Değişkenler 

Değişkenler, sabitler gibi başlatılabilir; ancak başlatıcı, çalışma zamanında hesaplanan genel bir ifade olabilir.

```go
var (
    home   = os.Getenv("HOME")
    user   = os.Getenv("USER")
    gopath = os.Getenv("GOPATH")
)
```

---

## 🧨 `init` fonksiyonu 

Son olarak, her kaynak dosya, gereken herhangi bir durumu (*state*) kurmak için kendi argümansız (*niladic*) `init` fonksiyonunu tanımlayabilir. (Aslında her dosya birden fazla `init` fonksiyonu tanımlayabilir.) Ve “son olarak” gerçekten “son olarak” demektir: `init`, paketteki tüm değişken bildirimleri başlatıcılarını değerlendirdikten sonra çağrılır ve bu başlatıcılar da yalnızca içe aktarılan tüm paketler başlatıldıktan sonra değerlendirilir.

Bildirimlerle ifade edilemeyen başlatmaların yanında, `init` fonksiyonlarının yaygın bir kullanımı, gerçek yürütme başlamadan önce program durumunun doğruluğunu doğrulamak veya onarmaktır.

```go
func init() {
    if user == "" {
        log.Fatal("$USER not set")
    }
    if home == "" {
        home = "/home/" + user
    }
    if gopath == "" {
        gopath = home + "/go"
    }
    // gopath may be overridden by --gopath flag on command line.
    flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}
```

---

## 🧩 Metotlar 

### 🧷 İşaretçiler (*pointers*) ve değerler (*values*)

`ByteSize` ile gördüğümüz gibi, metotlar (bir işaretçi veya bir arayüz (*interface*) hariç) adlandırılmış herhangi bir tür için tanımlanabilir; alıcı (*receiver*) bir `struct` olmak zorunda değildir.

Yukarıdaki dilimler (*slices*) tartışmasında bir `Append` fonksiyonu yazmıştık. Bunu, dilimler üzerinde bir metot olarak da tanımlayabiliriz. Bunu yapmak için önce metodu bağlayabileceğimiz adlandırılmış bir tür bildirir ve sonra alıcıyı (*receiver*) o türden bir değer yaparız.

```go
type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte {
    // Body exactly the same as the Append function defined above.
}
```

Bu hâlâ metodun güncellenmiş dilimi döndürmesini gerektirir. Metodu, alıcı olarak `ByteSlice` işaretçisi alacak şekilde yeniden tanımlayarak bu hantallığı ortadan kaldırabiliriz; böylece metot çağıranın dilimini üzerine yazabilir.

```go
func (p *ByteSlice) Append(data []byte) {
    slice := *p
    // Body as above, without the return.
    *p = slice
}
```

Aslında daha da iyisini yapabiliriz. Eğer fonksiyonumuzu standart bir `Write` metodu gibi görünmesi için şöyle değiştirirsek,

```go
func (p *ByteSlice) Write(data []byte) (n int, err error) {
    slice := *p
    // Again as above.
    *p = slice
    return len(data), nil
}
```

o zaman `*ByteSlice` türü standart arayüz `io.Writer`’ı sağlar (*satisfies*), ki bu kullanışlıdır. Örneğin içine yazdırabiliriz.

```go
var b ByteSlice
fmt.Fprintf(&b, "This hour has %d days\n", 7)
```

`ByteSlice`’ın adresini geçiririz; çünkü `io.Writer`’ı yalnızca `*ByteSlice` sağlar. Alıcılar için işaretçi mi değer mi kuralı şudur: değer metotları, işaretçiler ve değerler üzerinde çağrılabilir; ancak işaretçi metotları yalnızca işaretçiler üzerinde çağrılabilir.

Bu kural, işaretçi metotların alıcıyı değiştirebilmesinden doğar; onları bir değer üzerinde çağırmak, metoda değerin bir kopyasını verir ve yapılan değişiklikler atılır. Bu yüzden dil, bu hatayı engeller. Yine de kullanışlı bir istisna vardır. Değer adreslenebilir (*addressable*) olduğunda, dil, bir değer üzerinde bir işaretçi metodu çağırma gibi yaygın durumu, adres operatörünü otomatik ekleyerek halleder. Örneğimizde `b` değişkeni adreslenebilir olduğu için, `Write` metodunu yalnızca `b.Write` ile çağırabiliriz. Derleyici bunu bizim için `(&b).Write` olarak yeniden yazar.

Bu arada, bir bayt dilimi üzerinde `Write` kullanma fikri, `bytes.Buffer`’ın uygulanmasının merkezindedir.

---

## 🧬 Arayüzler ve diğer türler 

### 🧾 Arayüzler (*interfaces*)

Go’daki arayüzler, bir nesnenin davranışını belirtmenin bir yolunu sağlar: bir şey bunu yapabiliyorsa, burada kullanılabilir. Zaten birkaç basit örnek gördük; özel yazıcılar (*custom printers*) bir `String` metodu ile uygulanabilirken, `Fprintf` bir `Write` metodu olan her şeye çıktı üretebilir. Bir ya da iki metotlu arayüzler Go kodunda yaygındır ve genellikle metottan türetilen bir ad alırlar; `Write` uygulayan bir şey için `io.Writer` gibi.

Bir tür birden fazla arayüzü uygulayabilir. Örneğin, bir koleksiyon `sort` paketindeki yordamlarla, `Len()`, `Less(i, j int) bool` ve `Swap(i, j int)` içeren `sort.Interface`’ı uyguluyorsa sıralanabilir ve ayrıca özel bir biçimleyiciye (*formatter*) sahip olabilir. Bu uydurma örnekte `Sequence` her ikisini de sağlar.

```go
type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
    return len(s)
}
func (s Sequence) Less(i, j int) bool {
    return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

// Copy returns a copy of the Sequence.
func (s Sequence) Copy() Sequence {
    copy := make(Sequence, 0, len(s))
    return append(copy, s...)
}

// Method for printing - sorts the elements before printing.
func (s Sequence) String() string {
    s = s.Copy() // Make a copy; don't overwrite argument.
    sort.Sort(s)
    str := "["
    for i, elem := range s { // Loop is O(N²); will fix that in next example.
        if i > 0 {
            str += " "
        }
        str += fmt.Sprint(elem)
    }
    return str + "]"
}
```

---

## 🔄 Dönüştürmeler 

`Sequence`’ın `String` metodu, `Sprint`’in dilimler için zaten yaptığı işi yeniden yapıyor. (Ayrıca karmaşıklığı O(N²), ki kötüdür.) `Sequence`’ı `Sprint` çağırmadan önce sıradan bir `[]int`’e dönüştürürsek, çabayı paylaşabiliriz (ve ayrıca hızlandırabiliriz).

```go
func (s Sequence) String() string {
    s = s.Copy()
    sort.Sort(s)
    return fmt.Sprint([]int(s))
}
```

Bu metot, bir `String` metodundan `Sprintf`’i güvenli çağırmak için kullanılan dönüştürme tekniğine bir başka örnektir. İki tür (Sequence ve []int), tür adını yok sayarsak aynı olduğundan, aralarında dönüştürme yapmak yasaldır. Dönüştürme yeni bir değer oluşturmaz; yalnızca mevcut değer geçici olarak yeni bir türe sahipmiş gibi davranır. (Tamsayıdan kayan noktalıya gibi, yeni bir değer oluşturan başka yasal dönüştürmeler de vardır.)

Go programlarında, farklı bir metot kümesine erişmek için bir ifadenin türünü dönüştürmek bir deyimdir (*idiom*). Örnek olarak, tüm örneği şuna indirgemek için mevcut `sort.IntSlice` türünü kullanabiliriz:

```go
type Sequence []int

// Method for printing - sorts the elements before printing
func (s Sequence) String() string {
    s = s.Copy()
    sort.IntSlice(s).Sort()
    return fmt.Sprint([]int(s))
}
```

Artık `Sequence`’ın birden fazla arayüzü (sıralama ve yazdırma) uygulaması yerine, bir veri öğesinin birden fazla türe (Sequence, sort.IntSlice ve []int) dönüştürülebilmesi yeteneğini kullanıyoruz; bunların her biri işin bir parçasını yapıyor. Bu pratikte daha az yaygındır ama etkili olabilir.

---

## 🧷 Arayüz dönüştürmeleri ve tür doğrulamaları (*type assertions*) 

*Type switch*’ler, bir dönüşüm biçimidir: bir arayüz alırlar ve `switch` içindeki her `case` için onu, bir anlamda, o `case`’in türüne dönüştürürler. Aşağıda, `fmt.Printf` altındaki kodun bir değeri tür anahtarı (*type switch*) kullanarak string’e nasıl çevirdiğine dair basitleştirilmiş bir sürüm var. Zaten bir string ise, arayüzün tuttuğu gerçek string değerini isteriz; `String` metodu varsa, metodun çağrılmasının sonucunu isteriz.

```go
type Stringer interface {
    String() string
}

var value interface{} // Value provided by caller.
switch str := value.(type) {
case string:
    return str
case Stringer:
    return str.String()
}
```

İlk `case` fark edilir (*concrete*) bir değeri bulur; ikincisi arayüzü başka bir arayüze dönüştürür. Türleri bu şekilde karıştırmak tamamen uygundur.

Peki yalnızca önemsediğimiz tek bir tür varsa? Değerin bir string tuttuğunu biliyorsak ve yalnızca onu çıkarmak istiyorsak? Tek `case`’li bir *type switch* olurdu ama bir tür doğrulaması (*type assertion*) da olur. Bir tür doğrulaması, bir arayüz değerini alır ve ondan belirtilen açık türdeki (*explicit*) değeri çıkarır. Sözdizimi, bir *type switch*’in açılış cümlesinden ödünç alınmıştır; ancak `type` anahtar sözcüğü yerine açık bir tür vardır:

```go
value.(typeName)
```

ve sonuç, statik türü `typeName` olan yeni bir değerdir. Bu tür ya arayüzün tuttuğu somut tür olmalı ya da değerin dönüştürülebileceği ikinci bir arayüz türü olmalıdır. Değerin içinde olduğunu bildiğimiz string’i çıkarmak için şunu yazabiliriz:

```go
str := value.(string)
```

Ama değer bir string içermiyorsa, program çalışma zamanı hatasıyla çöker. Buna karşı korunmak için, değerin bir string olup olmadığını güvenle sınamak üzere “comma, ok” deyimini (*idiom*) kullanın:

```go
str, ok := value.(string)
if ok {
    fmt.Printf("string value is: %q\n", str)
} else {
    fmt.Printf("value is not a string\n")
}
```

Tür doğrulaması başarısız olursa, `str` yine de var olur ve türü `string` olur; fakat sıfır değere, yani boş string’e sahip olur.

Yeteneği göstermek için, işte bu bölümün başını açan *type switch*’e eşdeğer bir `if-else` deyimi:

```go
if str, ok := value.(string); ok {
    return str
} else if str, ok := value.(Stringer); ok {
    return str.String()
}
```

---

## 🧠 Genellik 

Bir tür yalnızca bir arayüzü uygulamak için varsa ve o arayüzün ötesinde dışa aktarılan (*exported*) metotları asla olmayacaksa, türün kendisini dışa aktarmaya gerek yoktur. Yalnızca arayüzü dışa aktarmak, değerin burada tanımlanan davranışın ötesinde ilginç bir davranışının olmadığını açık eder. Ayrıca, yaygın bir metodun her örneğinde dokümantasyonu tekrar etme ihtiyacını ortadan kaldırır.

Bu tür durumlarda, kurucunun (*constructor*) uygulayan tür yerine bir arayüz değeri döndürmesi gerekir. Örneğin, hash kütüphanelerinde hem `crc32.NewIEEE` hem de `adler32.New`, arayüz türü `hash.Hash32` döndürür. Bir Go programında Adler-32 yerine CRC-32 algoritmasını kullanmak, yalnızca kurucu çağrısını değiştirmeyi gerektirir; kodun geri kalanı, algoritma değişikliğinden etkilenmez.

Benzer bir yaklaşım, çeşitli `crypto` paketlerindeki akış şifreleme (*streaming cipher*) algoritmalarının, zincirledikleri blok şifrelerden ayrılmasını sağlar. `crypto/cipher` paketindeki `Block` arayüzü, tek bir veri bloğunu şifreleyen bir blok şifrenin davranışını belirtir. Sonra, `bufio` paketine benzetmeyle, bu arayüzü uygulayan `cipher` paketleri, blok şifrelemenin ayrıntılarını bilmeden, `Stream` arayüzüyle temsil edilen akış şifrelerini kurmak için kullanılabilir.

`crypto/cipher` arayüzleri şöyle görünür:

```go
type Block interface {
    BlockSize() int
    Encrypt(dst, src []byte)
    Decrypt(dst, src []byte)
}

type Stream interface {
    XORKeyStream(dst, src []byte)
}
```

İşte bir blok şifreyi akış şifresine dönüştüren sayaç modu (CTR) akışının tanımı; blok şifrenin ayrıntılarının nasıl soyutlandığına dikkat edin:

```go
// NewCTR returns a Stream that encrypts/decrypts using the given Block in
// counter mode. The length of iv must be the same as the Block's block size.
func NewCTR(block Block, iv []byte) Stream
```

`NewCTR`, yalnızca belirli bir şifreleme algoritmasına ve veri kaynağına değil; `Block` arayüzünün herhangi bir uygulamasına ve herhangi bir `Stream`’e uygulanır. Arayüz değerleri döndürdükleri için, CTR şifrelemeyi diğer şifreleme modlarıyla değiştirmek yerel bir değişikliktir. Kurucu çağrıları düzenlenmelidir; ancak çevredeki kod sonuçla yalnızca bir `Stream` olarak çalışmak zorunda olduğundan, farkı görmez.

---

## 🌐 Arayüzler ve metotlar 

Neredeyse her şeye metot bağlanabildiği için, neredeyse her şey bir arayüzü sağlayabilir. Aydınlatıcı bir örnek `http` paketindedir; `Handler` arayüzünü tanımlar. `Handler` uygulayan her nesne HTTP isteklerini sunabilir.

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

`ResponseWriter`, istemciye yanıt döndürmek için gereken metotlara erişim sağlayan bir arayüzdür. Bu metotlar standart `Write` metodunu içerir; bu yüzden bir `http.ResponseWriter`, `io.Writer` beklenen her yerde kullanılabilir. `Request`, istemciden gelen isteğin ayrıştırılmış gösterimini içeren bir `struct`’tır.

Kısalık için, POST’ları yok sayalım ve HTTP isteklerinin her zaman GET olduğunu varsayalım; bu basitleştirme, handler’ların nasıl kurulduğunu etkilemez. İşte sayfanın kaç kez ziyaret edildiğini sayan basit bir handler uygulaması.

```go
// Simple counter server.
type Counter struct {
    n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    ctr.n++
    fmt.Fprintf(w, "counter = %d\n", ctr.n)
}
```

(Temamıza uygun olarak, `Fprintf`’in bir `http.ResponseWriter`’a yazdırabildiğine dikkat edin.) Gerçek bir sunucuda `ctr.n`’e erişim, eşzamanlı erişime karşı korunmalıdır. Öneriler için `sync` ve `atomic` paketlerine bakın.

Referans için, böyle bir sunucunun URL ağacındaki bir düğüme nasıl bağlanacağını burada görebilirsiniz.

```go
import "net/http"
...
ctr := new(Counter)
http.Handle("/counter", ctr)
```

Peki neden `Counter` bir `struct` olsun? Gereken tek şey bir tamsayıdır. (Artırımın çağırana görünür olması için alıcının bir işaretçi olması gerekir.)

```go
// Simpler counter server.
type Counter int

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    *ctr++
    fmt.Fprintf(w, "counter = %d\n", *ctr)
}
```

Programınızda bir sayfanın ziyaret edildiğinin bildirilmesi gereken bir iç durum varsa ne olur? Web sayfasına bir kanal (*channel*) bağlayın.

```go
// A channel that sends a notification on each visit.
// (Probably want the channel to be buffered.)
type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    ch <- req
    fmt.Fprint(w, "notification sent")
}
```

Son olarak, diyelim ki `/args` üzerinde sunucu ikilisini (*server binary*) çalıştırırken kullanılan argümanları göstermek istedik. Argümanları yazdıran bir fonksiyon yazmak kolaydır.

```go
func ArgServer() {
    fmt.Println(os.Args)
}
```

Bunu nasıl HTTP sunucusuna dönüştürürüz? `ArgServer`’ı, değerini yok saydığımız bir türün metodu yapabilirdik; ancak daha temiz bir yol var. İşaretçiler ve arayüzler dışında her tür için metot tanımlayabildiğimizden, bir fonksiyon için metot yazabiliriz. `http` paketi şu kodu içerir:

```go
// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers.  If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler object that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, req).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, req *Request) {
    f(w, req)
}
```

`HandlerFunc`, `ServeHTTP` adında bir metodu olan bir türdür; dolayısıyla bu türün değerleri HTTP isteklerini sunabilir. Metodun uygulanışına bakın: alıcı, bir fonksiyon olan `f`’dir ve metot `f`’yi çağırır. Bu garip görünebilir ama alıcının bir kanal olup metotta kanala gönderme yapılmasından çok da farklı değildir.

`ArgServer`’ı bir HTTP sunucusuna dönüştürmek için, önce onu doğru imzaya (*signature*) sahip olacak şekilde değiştiririz.

```go
// Argument server.
func ArgServer(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, os.Args)
}
```

`ArgServer` artık `HandlerFunc` ile aynı imzaya sahiptir; dolayısıyla `IntSlice.Sort`’a erişmek için `Sequence`’ı `IntSlice`’a dönüştürdüğümüz gibi, metotlarına erişmek için `HandlerFunc` türüne dönüştürülebilir. Kurulum kodu özlüdür:

```go
http.Handle("/args", http.HandlerFunc(ArgServer))
```

Birisi `/args` sayfasını ziyaret ettiğinde, o sayfaya kurulan handler’ın değeri `ArgServer`, türü `HandlerFunc` olur. HTTP sunucusu bu türün `ServeHTTP` metodunu, alıcı olarak `ArgServer` ile çağırır; bu da `HandlerFunc.ServeHTTP` içindeki `f(w, req)` çağrısı üzerinden `ArgServer`’ı çağırır. Böylece argümanlar görüntülenir.

Bu bölümde bir `struct`, bir tamsayı, bir kanal ve bir fonksiyondan bir HTTP sunucusu yaptık; bunun tamamı, arayüzlerin yalnızca metot kümeleri olması ve (neredeyse) her tür için metot tanımlanabilmesi sayesindedir.

---

## 🕳️ Boş tanımlayıcı (*blank identifier*) 

Boş tanımlayıcıdan, `for range` döngüleri ve haritalar bağlamında birkaç kez bahsettik. Boş tanımlayıcı, herhangi bir türden herhangi bir değeri alarak atama yapılabilir ya da bildirilebilir ve değer zararsız biçimde yok sayılır. Unix’teki `/dev/null` dosyasına yazmak gibidir: bir değişkenin gerektiği ama gerçek değerin ilgisiz olduğu yerlerde yer tutucu (*place-holder*) olarak kullanılacak, yalnızca yazılabilir (*write-only*) bir değeri temsil eder. Gördüklerimizin ötesinde de kullanımları vardır.

### 🔁 Çoklu atamada boş tanımlayıcı

`for range` döngüsünde boş tanımlayıcı kullanımı, genel bir durumun özel bir hâlidir: çoklu atama (*multiple assignment*).

Bir atama sol tarafta birden fazla değer gerektiriyor, ancak bu değerlerden biri program tarafından kullanılmayacaksa, sol tarafta boş tanımlayıcı kullanmak sahte bir değişken oluşturma ihtiyacını ortadan kaldırır ve değerin atılacağını açıkça belirtir. Örneğin, bir değer ve bir hata döndüren bir fonksiyon çağrılırken, yalnızca hata önemliyse, ilgisiz değeri atmak için boş tanımlayıcıyı kullanın.

```go
if _, err := os.Stat(path); os.IsNotExist(err) {
    fmt.Printf("%s does not exist\n", path)
}
```

Bazen hatayı yok saymak için hata değerini atan kod görürsünüz; bu korkunç bir pratiktir. Hata dönüşlerini her zaman kontrol edin; bir sebeple sağlanırlar.

```go
// Bad! This code will crash if path does not exist.
fi, _ := os.Stat(path)
if fi.IsDir() {
    fmt.Printf("%s is a directory\n", path)
}
```

---

## 🚫 Kullanılmayan import’lar ve değişkenler 

Bir paketi içe aktarıp kullanmamak veya bir değişkeni bildirip kullanmamak hatadır. Kullanılmayan import’lar programı şişirir ve derlemeyi yavaşlatır; başlatılan ama kullanılmayan bir değişken ise en azından boşa bir hesaplamadır ve belki daha büyük bir hatanın göstergesidir. Ancak bir program aktif geliştirme altındayken, kullanılmayan import’lar ve değişkenler sıkça ortaya çıkar ve sadece derlemenin devam etmesi için onları silmek can sıkıcı olabilir; sonra tekrar gerekeceklerdir. Boş tanımlayıcı bir çözüm sağlar.

Bu yarım yazılmış programda iki kullanılmayan import (`fmt` ve `io`) ve kullanılmayan bir değişken (`fd`) vardır; bu yüzden derlenmez, ama o ana kadarki kodun doğru olup olmadığını görmek güzel olurdu.

```go
package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

func main() {
    fd, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err)
    }
    // TODO: use fd.
}
```

Kullanılmayan import şikâyetlerini susturmak için, içe aktarılan paketten bir sembole referans vermek üzere boş tanımlayıcı kullanın. Benzer şekilde, kullanılmayan `fd` değişkenini boş tanımlayıcıya atamak, kullanılmayan değişken hatasını susturur. Programın bu sürümü derlenir.

```go
package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

var _ = fmt.Printf // For debugging; delete when done.
var _ io.Reader    // For debugging; delete when done.

func main() {
    fd, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err)
    }
    // TODO: use fd.
    _ = fd
}
```

Gelenek olarak, import hatalarını susturmak için konan global bildirimler, hem bulunmalarını kolaylaştırmak hem de sonra temizlenmeleri gerektiğini hatırlatmak için import’lardan hemen sonra gelmeli ve yorumlanmalıdır.

---

## 🧪 Yan etki (*side effect*) için import 

Önceki örnekteki `fmt` veya `io` gibi kullanılmayan bir import eninde sonunda kullanılmalı veya kaldırılmalıdır: boş atamalar, kodun çalışma aşamasında olduğunu işaret eder. Ancak bazen, hiçbir şeyi açıkça kullanmadan yalnızca yan etkileri için bir paketi içe aktarmak faydalıdır. Örneğin, `net/http/pprof` paketi `init` fonksiyonu sırasında hata ayıklama (*debugging*) bilgisi sağlayan HTTP handler’larını kaydeder. Dışa aktarılan bir API’si vardır, ancak çoğu istemci yalnızca handler kaydına ihtiyaç duyar ve veriye bir web sayfası üzerinden erişir. Paketi yalnızca yan etkileri için içe aktarmak üzere, paket adını boş tanımlayıcıya yeniden adlandırın:

```go
import _ "net/http/pprof"
```

Bu import biçimi, paketin yalnızca yan etkileri için içe aktarıldığını açık eder; çünkü paketin başka bir olası kullanımı yoktur: bu dosyada bir adı yoktur. (Bir adı olsaydı ve biz o adı kullanmasaydık, derleyici programı reddederdi.)

---

## ✅ Arayüz denetimleri (*interface checks*) 

Yukarıdaki arayüz tartışmasında gördüğümüz gibi, bir türün bir arayüzü uyguladığını açıkça bildirmesi gerekmez. Bunun yerine, tür yalnızca arayüzün metotlarını uygulayarak arayüzü sağlar. Pratikte, çoğu arayüz dönüştürmesi statiktir ve bu yüzden derleme zamanında denetlenir. Örneğin, `*os.File`’ı `io.Reader` bekleyen bir fonksiyona geçirmek, `*os.File` `io.Reader` arayüzünü uygulamıyorsa derlenmez.

Ancak bazı arayüz denetimleri çalışma zamanında olur. Bunun bir örneği `encoding/json` paketindedir; `Marshaler` arayüzünü tanımlar. JSON kodlayıcı, bu arayüzü uygulayan bir değer aldığında, standart dönüşümü yapmak yerine değerin *marshaling* metodunu çağırarak onu JSON’a dönüştürür. Kodlayıcı bu özelliği çalışma zamanında şöyle bir tür doğrulamasıyla denetler:

```go
m, ok := val.(json.Marshaler)
```

Eğer yalnızca bir türün bir arayüzü sağlayıp sağlamadığını sormak gerekiyorsa—arayüzü gerçekten kullanmadan—örneğin bir hata denetiminin parçası olarak, tür doğrulamasıyla elde edilen değeri yok saymak için boş tanımlayıcıyı kullanın:

```go
if _, ok := val.(json.Marshaler); ok {
    fmt.Printf("value %v of type %T implements json.Marshaler\n", val, val)
}
```

Bu durumun ortaya çıktığı yerlerden biri, türü uygulayan paket içinde, türün gerçekten arayüzü sağladığını garanti etmek gerektiğindedir. Eğer bir tür—örneğin `json.RawMessage`—özel bir JSON temsiline ihtiyaç duyuyorsa, `json.Marshaler`’ı uygulamalıdır; ancak derleyicinin bunu otomatik olarak doğrulayacağı statik dönüştürmeler mevcut olmayabilir. Tür yanlışlıkla arayüzü sağlamaz hâle gelirse, JSON kodlayıcı yine çalışır, ama özel uygulamayı kullanmaz. Uygulamanın doğru olduğunu garanti etmek için, pakette boş tanımlayıcı kullanarak bir global bildirim yapılabilir:

```go
var _ json.Marshaler = (*RawMessage)(nil)
```

Bu bildirimde, `*RawMessage`’ın `Marshaler`’a dönüştürülmesiyle yapılan atama, `*RawMessage`’ın `Marshaler`’ı sağladığını gerektirir ve bu özellik derleme zamanında denetlenir. `json.Marshaler` arayüzü değişirse, bu paket artık derlenmez ve güncellenmesi gerektiği anlaşılır.

Bu yapıda boş tanımlayıcının görünmesi, bildirimin bir değişken oluşturmak için değil, yalnızca tür denetimi için var olduğunu gösterir. Yine de, bir arayüzü sağlayan her tür için bunu yapmayın. Gelenek olarak, bu tür bildirimler yalnızca kodda zaten mevcut statik dönüştürmeler olmadığında kullanılır; bu nadir bir durumdur.


## 🧩 Gömme 

Go, tipik “tür güdümlü” alt sınıflama (*subclassing*) anlayışını sağlamaz; ancak bir `struct` veya arayüz (*interface*) içine türleri gömerek bir uygulamanın parçalarını “ödünç alma” yeteneğine sahiptir.

### 🧷 Arayüz gömme

Arayüz gömme çok basittir. Daha önce `io.Reader` ve `io.Writer` arayüzlerinden bahsetmiştik; tanımları şöyledir:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

`io` paketi ayrıca bu tür metotların birkaçını birden uygulayabilen nesneleri belirten başka arayüzler de dışa aktarır. Örneğin, hem `Read` hem `Write` içeren `io.ReadWriter` vardır. `io.ReadWriter`’ı iki metodu açıkça yazarak da belirtebiliriz; ama iki arayüzü gömerek yeni arayüzü oluşturmak daha kolay ve daha çağrışım yapıcıdır:

```go
// ReadWriter is the interface that combines the Reader and Writer interfaces.
type ReadWriter interface {
    Reader
    Writer
}
```

Bu, göründüğü gibi bir şey söyler: bir `ReadWriter`, bir `Reader`’ın yaptığını yapabilir ve bir `Writer`’ın yaptığını yapabilir; gömülü arayüzlerin bir birleşimidir. Yalnızca arayüzler, arayüzlerin içine gömülebilir.

### 🧱 Struct gömme

Aynı temel fikir `struct`’lar için de geçerlidir; ancak çok daha geniş kapsamlı sonuçlara sahiptir. `bufio` paketinde iki `struct` türü vardır: `bufio.Reader` ve `bufio.Writer`; elbette her biri `io` paketindeki karşılık gelen arayüzleri uygular. `bufio` ayrıca gömme kullanarak bir okuyucu ve bir yazıcıyı tek bir `struct` içinde birleştiren arabellekli bir okuyucu/yazıcı da uygular: `struct` içinde türleri listeler ama alan adı vermez.

```go
// ReadWriter stores pointers to a Reader and a Writer.
// It implements io.ReadWriter.
type ReadWriter struct {
    *Reader  // *bufio.Reader
    *Writer  // *bufio.Writer
}
```

Gömülü öğeler `struct` işaretçileridir ve elbette kullanılmadan önce geçerli `struct`’lara işaret edecek şekilde başlatılmaları gerekir. `ReadWriter` `struct`’ı şu şekilde de yazılabilirdi:

```go
type ReadWriter struct {
    reader *Reader
    writer *Writer
}
```

ancak o zaman alanların metotlarını yukarı taşımak (*promote*) ve `io` arayüzlerini sağlamak için ayrıca yönlendirme metotları yazmamız gerekirdi; örneğin:

```go
func (rw *ReadWriter) Read(p []byte) (n int, err error) {
    return rw.reader.Read(p)
}
```

`struct`’ları doğrudan gömerek bu angaryadan kaçınırız. Gömülü türlerin metotları “bedavaya” gelir; bu da `bufio.ReadWriter`’ın yalnızca `bufio.Reader` ve `bufio.Writer` metotlarına sahip olmadığı, aynı zamanda üç arayüzü de sağladığı anlamına gelir: `io.Reader`, `io.Writer` ve `io.ReadWriter`.

Gömmenin, alt sınıflamadan farklılaştığı önemli bir nokta vardır. Bir türü gömdüğümüzde, o türün metotları dış türün metotları hâline gelir; ancak çağrıldıklarında metot alıcısı (*receiver*) dış tür değil, iç türdür. Örneğimizde, `bufio.ReadWriter`’ın `Read` metodu çağrıldığında, yukarıda elle yazılan yönlendirme metoduyla aynı etkiyi üretir; alıcı `ReadWriter`’ın kendisi değil, onun `reader` alanıdır.

Gömme ayrıca basit bir kolaylık da olabilir. Bu örnek, gömülü bir alanı normal, adlandırılmış bir alanla birlikte gösterir.

```go
type Job struct {
    Command string
    *log.Logger
}
```

`Job` türü artık `*log.Logger`’ın `Print`, `Printf`, `Println` ve diğer metotlarına sahiptir. Elbette `Logger`’a bir alan adı da verebilirdik; ama bunu yapmak şart değildir. Ve artık başlatıldıktan sonra, `Job` üzerinden log yazabiliriz:

```go
job.Println("starting now...")
```

`Logger`, `Job` `struct`’ının normal bir alanıdır; bu yüzden `Job` için kurucu içinde alışıldık şekilde başlatabiliriz:

```go
func NewJob(command string, logger *log.Logger) *Job {
    return &Job{command, logger}
}
```

veya bir bileşik sabitle:

```go
job := &Job{command, log.New(os.Stderr, "Job: ", log.Ldate)}
```

Gömülü bir alana doğrudan referans vermemiz gerekirse, paket niteleyicisini yok sayarak alanın tür adı, alan adı işlevi görür; `ReadWriter` `struct`’ındaki `Read` metodunda olduğu gibi. Burada, bir `Job` değişkeni `job` için `*log.Logger` alanına erişmek istersek `job.Logger` yazarız; bu, `Logger` metotlarını rafine etmek istersek faydalı olur.

```go
func (job *Job) Printf(format string, args ...interface{}) {
    job.Logger.Printf("%q: %s", job.Command, fmt.Sprintf(format, args...))
}
```

Tür gömme, ad çakışması problemine yol açar; ancak çözüm kuralları basittir. Birincisi, alan veya metot `X`, türün daha derin iç içe kısımlarındaki diğer `X` öğelerini gizler. Eğer `log.Logger` içinde `Command` adlı bir alan veya metot olsaydı, `Job`’un `Command` alanı onu bastırırdı.

İkincisi, aynı ad aynı gömme seviyesinde görünürse genellikle hatadır; `Job` `struct`’ı başka bir `Logger` alanı veya metodu içeriyorsa `log.Logger`’ı gömmek hatalı olurdu. Ancak, kopya ad programda tür tanımı dışında hiç anılmazsa sorun yoktur. Bu nitelik, dışarıdan gömülen türlerde yapılan değişikliklere karşı bir miktar koruma sağlar; hiçbir alanın kullanılmadığı bir programda, alt türlerden birine eklenen bir alanın başka bir alt türdeki alanla çakışması problem değildir.

---

## ⚙️ Eşzamanlılık 

### 🔁 İletişerek paylaş

Eşzamanlı programlama büyük bir konudur; burada yalnızca Go’ya özgü bazı önemli noktalar için yer var.

Birçok ortamda eşzamanlı programlamayı zorlaştıran şey, paylaşılan değişkenlere doğru erişimi uygulamak için gereken inceliklerdir. Go, paylaşılan değerlerin kanallar (*channels*) üzerinden dolaştırıldığı ve aslında ayrı yürütme iş parçacıkları tarafından aktif biçimde paylaşılmadığı farklı bir yaklaşımı teşvik eder. Herhangi bir anda değere yalnızca bir goroutine erişir. Veri yarışları (*data races*), tasarım gereği oluşamaz. Bu düşünme biçimini teşvik etmek için bunu bir slogana indirdik:

Do not communicate by sharing memory; instead, share memory by communicating.

Bu yaklaşım aşırıya kaçabilir. Örneğin, referans sayımı bir `int` değişkenin etrafına bir mutex koyarak en iyi yapılabilir. Ancak üst düzey bir yaklaşım olarak, erişimi kontrol etmek için kanalları kullanmak, açık ve doğru programlar yazmayı kolaylaştırır.

Bu modeli düşünmenin bir yolu, tek bir CPU üzerinde çalışan tipik bir tek iş parçacıklı programu ele almaktır. Senkronizasyon ilkellerine ihtiyaç duymaz. Şimdi bunun bir başka örneğini çalıştırın; onun da senkronizasyona ihtiyacı yoktur. Şimdi bu ikisinin iletişim kurmasına izin verin; iletişimin kendisi senkronize ediciyse hâlâ başka bir senkronizasyona ihtiyaç yoktur. Unix boru hatları (*pipelines*) bu modele mükemmel uyar. Go’nun eşzamanlılık yaklaşımı Hoare’un *Communicating Sequential Processes* (CSP) fikrinden doğsa da, aynı zamanda Unix borularının tür-güvenli (*type-safe*) bir genellemesi olarak da görülebilir.

### 🧵 Goroutine’ler

Bunlara *goroutine* denir; çünkü mevcut terimler—*threads*, *coroutines*, *processes* vb.—yanıltıcı çağrışımlar taşır. Bir goroutine’in basit bir modeli vardır: aynı adres alanında, diğer goroutine’lerle eşzamanlı yürütülen bir fonksiyondur. Hafiftir; maliyeti yığın alanı ayırmaktan biraz fazladır. Yığınlar küçük başlar; bu yüzden ucuzdurlar ve gerektiğinde heap depolaması ayırarak (ve serbest bırakarak) büyürler.

Goroutine’ler birden çok işletim sistemi iş parçacığı üzerine çoklanır (*multiplexed*); bu yüzden biri I/O beklerken bloklansa bile diğerleri çalışmaya devam eder. Tasarımları, iş parçacığı oluşturma ve yönetiminin karmaşıklıklarının çoğunu gizler.

Bir fonksiyon veya metot çağrısının başına `go` anahtar sözcüğünü ekleyerek, çağrıyı yeni bir goroutine içinde çalıştırırsınız. Çağrı bittiğinde goroutine sessizce sonlanır. (Etkisi, Unix kabuğundaki bir komutu arka planda çalıştırmak için kullanılan `&` notasyonuna benzer.)

```go
go list.Sort()  // run list.Sort concurrently; don't wait for it.
```

Bir goroutine çağrısında fonksiyon sabiti (*function literal*) kullanmak işe yarayabilir.

```go
func Announce(message string, delay time.Duration) {
    go func() {
        time.Sleep(delay)
        fmt.Println(message)
    }()  // Note the parentheses - must call the function.
}
```

Go’da fonksiyon sabitleri *closure*’dır: uygulama, fonksiyonun referans verdiği değişkenlerin aktif oldukları sürece yaşamalarını sağlar.

Bu örnekler pek pratik değildir; çünkü fonksiyonların tamamlandığını bildirecek bir yolu yoktur. Bunun için kanallara ihtiyacımız var.

### 📡 Kanallar

Haritalar gibi, kanallar da `make` ile ayrılır ve ortaya çıkan değer alttaki bir veri yapısına referans gibi davranır. İsteğe bağlı bir tamsayı parametresi verilirse, kanalın arabellek boyutunu ayarlar. Varsayılan değer sıfırdır; yani arabelleklenmemiş (*unbuffered*) ya da eşzamanlı (*synchronous*) kanal.

```go
ci := make(chan int)            // unbuffered channel of integers
cj := make(chan int, 0)         // unbuffered channel of integers
cs := make(chan *os.File, 100)  // buffered channel of pointers to Files
```

Arabelleklenmemiş kanallar, iletişimi—bir değerin değişimini—senkronizasyonla birleştirir; yani iki hesaplamanın (goroutine’in) bilinen bir durumda olmasını garanti eder.

Kanallarla kullanılan birçok güzel deyim (*idiom*) vardır. Başlamak için bir tane: önceki bölümde bir sıralamayı arka planda başlatmıştık. Bir kanal, başlatan goroutine’in sıralamanın bitmesini beklemesini sağlayabilir.

```go
c := make(chan int)  // Allocate a channel.
// Start the sort in a goroutine; when it completes, signal on the channel.
go func() {
    list.Sort()
    c <- 1  // Send a signal; value does not matter.
}()
doSomethingForAWhile()
<-c   // Wait for sort to finish; discard sent value.
```

Alıcılar her zaman veri alana kadar bloklanır. Kanal arabelleklenmemişse, gönderici alıcının değeri almasına kadar bloklanır. Kanalın arabelleği varsa, gönderici yalnızca değerin arabelleğe kopyalanmasına kadar bloklanır; arabellek doluysa, bu da bir alıcının arabellekten bir değer almasını beklemek anlamına gelir.

Arabellekli bir kanal örneğin bir semafor gibi kullanılabilir; örneğin iş hacmini sınırlamak için. Bu örnekte, gelen istekler `handle`’a aktarılır; `handle` kanala bir değer gönderir, isteği işler ve sonra “semaforu” bir sonraki tüketici için hazır hâle getirmek üzere kanaldan bir değer alır. Kanal arabelleğinin kapasitesi, eşzamanlı `process` çağrılarının sayısını sınırlar.

```go
var sem = make(chan int, MaxOutstanding)

func handle(r *Request) {
    sem <- 1    // Wait for active queue to drain.
    process(r)  // May take a long time.
    <-sem       // Done; enable next request to run.
}

func Serve(queue chan *Request) {
    for {
        req := <-queue
        go handle(req)  // Don't wait for handle to finish.
    }
}
```

`MaxOutstanding` adet `handle` `process` çalıştırırken, daha fazlası dolu arabelleğe gönderme yapmaya çalışırken bloklanır; mevcut handler’lardan biri bitip arabellekten alıncaya kadar.

Ne var ki bu tasarımın bir problemi vardır: `Serve`, yalnızca `MaxOutstanding` tanesi aynı anda çalışabilse bile gelen her istek için yeni bir goroutine oluşturur. Sonuç olarak, istekler çok hızlı gelirse program sınırsız kaynak tüketebilir. Bu eksikliği, goroutine oluşturmayı kapılayacak (*gate*) şekilde `Serve`’ü değiştirerek giderebiliriz:

```go
func Serve(queue chan *Request) {
    for req := range queue {
        sem <- 1
        go func() {
            process(req)
            <-sem
        }()
    }
}
```

(Not: Go 1.22’den önceki sürümlerde bu kodda bir hata vardır: döngü değişkeni tüm goroutine’ler arasında paylaşılır. Ayrıntılar için Go wiki’ye bakın.)

Kaynakları iyi yöneten bir başka yaklaşım da, isteklere bakan sabit sayıda `handle` goroutine’i başlatıp hepsinin istek kanalını okumasını sağlamaktır. Goroutine sayısı, `process`’in eşzamanlı çağrı sayısını sınırlar. Bu `Serve` fonksiyonu ayrıca çıkması gerektiğinde kendisine haber verilecek bir kanal da kabul eder; goroutine’leri başlattıktan sonra o kanaldan alım yaparak bloklanır.

```go
func handle(queue chan *Request) {
    for r := range queue {
        process(r)
    }
}

func Serve(clientRequests chan *Request, quit chan bool) {
    // Start handlers
    for i := 0; i < MaxOutstanding; i++ {
        go handle(clientRequests)
    }
    <-quit  // Wait to be told to exit.
}
```

### 🔁 Kanalların kanalları

Go’nun en önemli özelliklerinden biri, kanalın birinci sınıf (*first-class*) bir değer olmasıdır; tıpkı diğer her şey gibi ayrılabilir ve dolaştırılabilir. Bu özelliğin yaygın bir kullanımı, güvenli ve paralel çoklayıcıdan çıkarma (*demultiplexing*) uygulamaktır.

Önceki bölümdeki örnekte `handle`, bir isteğin idealize edilmiş handler’ıydı ama ele aldığı türü tanımlamadık. Eğer bu tür, yanıt vermek için bir kanal içerirse, her istemci yanıt için kendi yolunu sağlayabilir. İşte `Request` türünün şematik tanımı.

```go
type Request struct {
    args        []int
    f           func([]int) int
    resultChan  chan int
}
```

İstemci, fonksiyonunu ve argümanlarını, ayrıca yanıtı almak üzere `request` nesnesinin içindeki bir kanalı sağlar.

```go
func sum(a []int) (s int) {
    for _, v := range a {
        s += v
    }
    return
}

request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
// Send request
clientRequests <- request
// Wait for response.
fmt.Printf("answer: %d\n", <-request.resultChan)
```

Sunucu tarafında, değişen tek şey handler fonksiyonudur.

```go
func handle(queue chan *Request) {
    for req := range queue {
        req.resultChan <- req.f(req.args)
    }
}
```

Bunu gerçekçi yapmak için açıkça daha yapılacak çok şey var; ama bu kod, hız sınırlamalı (*rate-limited*), paralel, bloklamayan (*non-blocking*) bir RPC sistemi için bir çerçevedir ve ortada bir tane bile mutex yoktur.

### 🧮 Paralelleştirme

Bu fikirlerin bir diğer uygulaması, pahalı bir hesabı birden çok CPU çekirdeğine paralelleştirmektir. Hesap bağımsız çalışabilecek parçalara bölünebiliyorsa, paralelleştirilebilir; her parça bittiğinde sinyallemek için bir kanal kullanılır.

Diyelim ki bir öğe vektörü üzerinde pahalı bir işlem yapmak istiyoruz ve her öğe üzerindeki işlemin sonucu birbirinden bağımsız; şu idealize edilmiş örnekte olduğu gibi.

```go
type Vector []float64

// Apply the operation to v[i], v[i+1] ... up to v[n-1].
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
    for ; i < n; i++ {
        v[i] += u.Op(v[i])
    }
    c <- 1    // signal that this piece is done
}
```

Parçaları bir döngüde, CPU başına bir tane olacak şekilde bağımsız başlatırız. Herhangi bir sırayla tamamlanabilirler ama önemli değildir; tüm goroutine’leri başlattıktan sonra kanalı boşaltarak (*drain*) tamamlanma sinyallerini sayarız.

```go
const numCPU = 4 // number of CPU cores

func (v Vector) DoAll(u Vector) {
    c := make(chan int, numCPU)  // Buffering optional but sensible.
    for i := 0; i < numCPU; i++ {
        go v.DoSome(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
    }
    // Drain the channel.
    for i := 0; i < numCPU; i++ {
        <-c    // wait for one task to complete
    }
    // All done.
}
```

`numCPU` için sabit bir değer tanımlamak yerine, çalışma zamanına uygun değeri sorabiliriz. `runtime.NumCPU` fonksiyonu makinedeki donanımsal CPU çekirdek sayısını döndürür; dolayısıyla şöyle yazabiliriz:

```go
var numCPU = runtime.NumCPU()
```

Ayrıca `runtime.GOMAXPROCS` adlı bir fonksiyon vardır; bir Go programının aynı anda çalıştırabileceği (ya da ayarlayabileceği) çekirdek sayısını bildirir. Varsayılanı `runtime.NumCPU` değeridir; ancak benzer adlı bir kabuk ortam değişkeni ayarlanarak veya fonksiyon pozitif bir sayı ile çağrılarak değiştirilebilir. Sıfırla çağırmak yalnızca değeri sorgular. Dolayısıyla kullanıcının kaynak isteğini onurlandırmak istiyorsak şöyle yazmalıyız:

```go
var numCPU = runtime.GOMAXPROCS(0)
```

Eşzamanlılık (*concurrency*) ile paralellik (*parallelism*) fikirlerini karıştırmamaya dikkat edin: eşzamanlılık, bir programı bağımsız yürütülen bileşenler olarak yapılandırmaktır; paralellik ise verimlilik için hesapları birden çok CPU üzerinde paralel yürütmektir. Go’nun eşzamanlılık özellikleri bazı problemleri paralel hesaplar olarak yapılandırmayı kolaylaştırsa da, Go paralel bir dil değil, eşzamanlı bir dildir; tüm paralelleştirme problemleri Go’nun modeline uymaz. Ayrımın tartışması için, bu blog gönderisinde alıntılanan konuşmaya bakın.

### 🪣 Sızdıran bir tampon

Eşzamanlı programlamanın araçları, eşzamanlı olmayan fikirleri bile ifade etmeyi kolaylaştırabilir. İşte bir RPC paketinden soyutlanmış bir örnek. İstemci goroutine’i, belki bir ağ üzerinden, bir kaynaktan veri alarak döner. Arabellek ayırıp serbest bırakmaktan kaçınmak için bir serbest liste (*free list*) tutar ve bunu temsil etmek için arabellekli bir kanal kullanır. Kanal boşsa, yeni bir arabellek ayrılır. Mesaj arabelleği hazır olunca sunucuya `serverChan` üzerinden gönderilir.

```go
var freeList = make(chan *Buffer, 100)
var serverChan = make(chan *Buffer)

func client() {
    for {
        var b *Buffer
        // Grab a buffer if available; allocate if not.
        select {
        case b = <-freeList:
            // Got one; nothing more to do.
        default:
            // None free, so allocate a new one.
            b = new(Buffer)
        }
        load(b)              // Read next message from the net.
        serverChan <- b      // Send to server.
    }
}
```

Sunucu döngüsü, istemciden gelen her mesajı alır, işler ve arabelleği serbest listeye geri koyar.

```go
func server() {
    for {
        b := <-serverChan    // Wait for work.
        process(b)
        // Reuse buffer if there's room.
        select {
        case freeList <- b:
            // Buffer on free list; nothing more to do.
        default:
            // Free list full, just carry on.
        }
    }
}
```

İstemci, `freeList`’ten bir arabellek almaya çalışır; yoksa taze bir tane ayırır. Sunucunun `freeList`’e gönderimi, liste dolu değilse `b`’yi serbest listeye geri koyar; liste doluysa arabellek yere düşürülür ve çöp toplayıcı (*garbage collector*) tarafından geri kazanılır. (`select` deyimlerindeki `default` cümleleri, başka hiçbir `case` hazır olmadığında çalışır; yani `select`’ler asla bloklanmaz.) Bu uygulama, birkaç satırda sızdıran kova (*leaky bucket*) serbest listesini kurar; arabellekli kanala ve çöp toplayıcıya muhasebe için yaslanır.

---

## ❗ Hatalar 

Kütüphane yordamları sıklıkla çağırana bir tür hata göstergesi döndürmek zorundadır. Daha önce belirtildiği gibi, Go’nun çoklu dönüş değeri, normal dönüş değerinin yanında ayrıntılı bir hata açıklaması döndürmeyi kolaylaştırır. Bu özelliği ayrıntılı hata bilgisi sağlamak için kullanmak iyi bir stildir. Örneğin göreceğimiz gibi, `os.Open` başarısız olduğunda yalnızca `nil` bir işaretçi döndürmez; aynı zamanda neyin yanlış gittiğini açıklayan bir `error` değeri de döndürür.

Gelenek gereği, hatalar `error` türündedir; bu basit bir yerleşik arayüzdür.

```go
type error interface {
    Error() string
}
```

Bir kütüphane yazarı, bu arayüzü perde arkasında daha zengin bir modelle uygulamakta serbesttir; bu yalnızca hatayı görmekle kalmayıp bir bağlam (*context*) da sağlamayı mümkün kılar. Söylendiği gibi, `os.Open` olağan `*os.File` dönüş değerinin yanında bir `error` değeri de döndürür. Dosya başarıyla açılırsa hata `nil` olur; ama bir problem olduğunda, `os.PathError` tutar:

```go
// PathError records an error and the operation and
// file path that caused it.
type PathError struct {
    Op string    // "open", "unlink", etc.
    Path string  // The associated file.
    Err error    // Returned by the system call.
}

func (e *PathError) Error() string {
    return e.Op + " " + e.Path + ": " + e.Err.Error()
}
```

`PathError`’ın `Error` metodu şöyle bir string üretir:

```text
open /etc/passwx: no such file or directory
```

Böyle bir hata—problemli dosya adını, işlemi ve tetiklenen işletim sistemi hatasını içeren—hata üreten çağrıdan çok uzakta bile yazdırılsa faydalıdır; düz “no such file or directory” ifadesinden çok daha bilgilendiricidir.

Mümkün olduğunda, hata string’leri kökenlerini belirtmelidir; örneğin hatayı üreten işlemi veya paketi adlandıran bir önekle. Örneğin `image` paketinde, bilinmeyen bir biçim yüzünden oluşan çözme (*decoding*) hatasının string gösterimi “image: unknown format”tır.

Hata ayrıntılarıyla hassas biçimde ilgilenen çağıranlar, belirli hataları aramak ve ayrıntıları çıkarmak için bir *type switch* veya tür doğrulaması (*type assertion*) kullanabilir. `PathError` için bu, toparlanabilir hataları görmek üzere içteki `Err` alanını incelemeyi içerebilir.

```go
for try := 0; try < 2; try++ {
    file, err = os.Create(filename)
    if err == nil {
        return
    }
    if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
        deleteTempFiles()  // Recover some space.
        continue
    }
    return
}
```

Buradaki ikinci `if` deyimi başka bir tür doğrulamasıdır. Başarısız olursa `ok` `false` olur ve `e` `nil` olur. Başarılı olursa `ok` `true` olur; bu, hatanın `*os.PathError` türünde olduğu anlamına gelir ve o zaman `e` de öyledir; böylece hatayla ilgili daha fazla bilgi için onu inceleyebiliriz.

---

## 💥 Panic 

Çağırana hata raporlamanın olağan yolu, ekstra bir dönüş değeri olarak `error` döndürmektir. Kanonik `Read` metodu bunun iyi bilinen bir örneğidir; bir bayt sayacı ve bir hata döndürür. Peki hata toparlanamazsa ne olur? Bazen program basitçe devam edemez.

Bu amaçla, yerleşik `panic` fonksiyonu vardır; etkisi, programı durduracak bir çalışma zamanı hatası oluşturmaktır (ama bir sonraki bölüme bakın). Fonksiyon, keyfi türde tek bir argüman alır—çoğunlukla bir string—ve program ölürken yazdırılır. Ayrıca, sonsuz bir döngüden çıkmak gibi imkânsız bir şeyin olduğunu belirtmenin de bir yoludur.

```go
// A toy implementation of cube root using Newton's method.
func CubeRoot(x float64) float64 {
    z := x/3   // Arbitrary initial value
    for i := 0; i < 1e6; i++ {
        prevz := z
        z -= (z*z*z-x) / (3*z*z)
        if veryClose(z, prevz) {
            return z
        }
    }
    // A million iterations has not converged; something is wrong.
    panic(fmt.Sprintf("CubeRoot(%g) did not converge", x))
}
```

Bu yalnızca bir örnektir; gerçek kütüphane fonksiyonları `panic`’ten kaçınmalıdır. Problem maskelenebiliyor veya etrafından dolaşılabiliyorsa, tüm programı aşağı çekmek yerine çalışmaya devam etmek her zaman daha iyidir. Olası bir karşı örnek, başlatma (*initialization*) sırasıdır: kütüphane gerçekten kendini kuramıyorsa, tabiri caizse `panic` makul olabilir.

```go
var user = os.Getenv("USER")

func init() {
    if user == "" {
        panic("no value for $USER")
    }
}
```

---

## 🛟 Recover 

`panic` çağrıldığında—dilim sınırı dışında indeksleme veya başarısız bir tür doğrulaması gibi çalışma zamanı hataları da buna dahildir—mevcut fonksiyonun yürütmesini hemen durdurur ve goroutine’in yığınını (*stack*) açmaya (*unwinding*) başlar; bu sırada ertelenmiş (*deferred*) fonksiyonları çalıştırır. Bu açma işlemi goroutine’in yığınının tepesine ulaşırsa program ölür. Ancak yerleşik `recover` fonksiyonunu kullanarak goroutine’in kontrolünü yeniden ele almak ve normal yürütmeye dönmek mümkündür.

`recover` çağrısı, yığın açmayı durdurur ve `panic`’e verilen argümanı döndürür. Açma sırasında çalışan tek kod ertelenmiş fonksiyonların içindeki kod olduğu için, `recover` yalnızca ertelenmiş fonksiyonların içinde faydalıdır.

`recover`’ın bir uygulaması, bir sunucuda diğer goroutine’leri öldürmeden, arızalanan bir goroutine’i kapatmaktır.

```go
func server(workChan <-chan *Work) {
    for work := range workChan {
        go safelyDo(work)
    }
}

func safelyDo(work *Work) {
    defer func() {
        if err := recover(); err != nil {
            log.Println("work failed:", err)
        }
    }()
    do(work)
}
```

Bu örnekte, `do(work)` `panic` yaparsa sonuç loglanır ve goroutine, diğerlerini rahatsız etmeden temizce çıkar. Ertelenmiş *closure* içinde başka bir şey yapmaya gerek yoktur; `recover`’ı çağırmak durumu tamamen ele alır.

`recover`, yalnızca ertelenmiş bir fonksiyonun içinden doğrudan çağrılmadıkça her zaman `nil` döndürdüğünden, ertelenmiş kod kendi içinde `panic` ve `recover` kullanan kütüphane yordamlarını güvenle çağırabilir. Örneğin `safelyDo` içindeki ertelenmiş fonksiyon, `recover` çağırmadan önce bir log fonksiyonunu çağırabilir; bu log kodu panik durumundan etkilenmeden çalışır.

Bu kurtarma kalıbı (*recovery pattern*) yerindeyken, `do` fonksiyonu (ve çağırdığı her şey) `panic` çağırarak herhangi bir kötü durumdan temizce çıkabilir. Bu fikri, karmaşık yazılımlarda hata işlemeyi sadeleştirmek için kullanabiliriz. Bir regexp paketinin idealize edilmiş bir sürümüne bakalım; ayrıştırma hatalarını yerel bir hata türüyle `panic` yaparak raporlar. `Error` tanımı, bir hata metodu ve `Compile` fonksiyonu aşağıdadır.

```go
// Error is the type of a parse error; it satisfies the error interface.
type Error string
func (e Error) Error() string {
    return string(e)
}

// error is a method of *Regexp that reports parsing errors by
// panicking with an Error.
func (regexp *Regexp) error(err string) {
    panic(Error(err))
}

// Compile returns a parsed representation of the regular expression.
func Compile(str string) (regexp *Regexp, err error) {
    regexp = new(Regexp)
    // doParse will panic if there is a parse error.
    defer func() {
        if e := recover(); e != nil {
            regexp = nil    // Clear return value.
            err = e.(Error) // Will re-panic if not a parse error.
        }
    }()
    return regexp.doParse(str), nil
}
```

`doParse` `panic` yaparsa, kurtarma bloğu dönüş değerini `nil` yapar—ertelenmiş fonksiyonlar isimlendirilmiş dönüş değerlerini değiştirebilir. Sonra, `err`’ye atamada, problemin yerel `Error` türünde olduğunu, bir tür doğrulamasıyla denetler. Eğer değilse, tür doğrulaması başarısız olur; bu da yığın açmanın, sanki hiç kesintiye uğramamış gibi, çalışma zamanı hatasıyla devam etmesine neden olur. Bu denetim, örneğin sınır dışında indeksleme gibi beklenmeyen bir şey olursa, `panic` ve `recover` ayrıştırma hatalarını ele almak için kullanılsa bile kodun yine de başarısız olmasını sağlar.

Hata işleme yerindeyken, `error` metodu (bir türe bağlanmış bir metot olduğu için, yerleşik `error` türüyle aynı ada sahip olması sorun değildir, hatta doğaldır) ayrıştırma hatalarını, ayrıştırma yığınını (*parse stack*) elle açma konusunda endişelenmeden raporlamayı kolaylaştırır:

```go
if pos == 0 {
    re.error("'*' illegal at start of expression")
}
```

Bu kalıp ne kadar faydalı olsa da, yalnızca bir paket içinde kullanılmalıdır. `Parse`, iç `panic` çağrılarını `error` değerlerine çevirir; panikleri istemcisine açığa çıkarmaz. Bu izlenmesi gereken iyi bir kuraldır.

Bu arada, bu *re-panic* deyimi, gerçek bir hata oluşursa `panic` değerini değiştirir. Ancak hem özgün hem de yeni başarısızlık çöküş raporunda sunulur; dolayısıyla problemin kök nedeni hâlâ görünür olur. Bu basit *re-panic* yaklaşımı genellikle yeterlidir—sonuçta bu bir çöküştür—ama yalnızca özgün değeri göstermek isterseniz, beklenmeyen problemleri filtreleyip özgün hatayla yeniden `panic` yapmak için biraz daha kod yazabilirsiniz. Bunu okura bir alıştırma olarak bırakıyoruz.

---

## 🌐 Bir web sunucusu 

Tam bir Go programıyla bitirelim: bir web sunucusu. Bu aslında bir tür web yeniden-sunucusudur (*re-server*). Google, `chart.apis.google.com` adresinde veriyi grafik ve diyagramlara otomatik biçimlendiren bir servis sağlar. Ancak etkileşimli kullanmak zordur; çünkü veriyi URL içinde bir sorgu olarak koymanız gerekir. Buradaki program, bir veri biçimi için daha hoş bir arayüz sağlar: kısa bir metin alır, chart sunucusunu çağırarak bir QR kodu üretir; bu, metni kodlayan bir kutucuk matrisi şeklindeki bir görüntüdür. Bu görüntü, cep telefonunuzun kamerasıyla alınabilir ve örneğin bir URL olarak yorumlanabilir; telefonun küçük klavyesine URL yazma zahmetini azaltır.

İşte tam program. Ardından açıklaması gelir.

```go
package main

import (
    "flag"
    "html/template"
    "log"
    "net/http"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {
    flag.Parse()
    http.Handle("/", http.HandlerFunc(QR))
    err := http.ListenAndServe(*addr, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

func QR(w http.ResponseWriter, req *http.Request) {
    templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET">
    <input maxLength=1024 size=70 name=s value="" title="Text to QR Encode">
    <input type=submit value="Show QR" name=qr>
</form>
</body>
</html>
`
```

`main`’e kadar olan parçaları takip etmek kolay olmalı. Tek bayrak, sunucumuz için varsayılan bir HTTP portu ayarlar. `templ` şablon değişkeni, eğlencenin gerçekleştiği yerdir. Sunucu tarafından çalıştırılıp sayfayı gösterecek bir HTML şablonu inşa eder; birazdan daha fazlası.

`main` fonksiyonu bayrakları ayrıştırır ve yukarıda konuştuğumuz mekanizmayı kullanarak `QR` fonksiyonunu sunucunun kök yoluna bağlar. Ardından sunucuyu başlatmak için `http.ListenAndServe` çağrılır; sunucu çalıştığı sürece bloklanır.

`QR`, form verisini içeren isteği alır ve şablonu, `s` adlı form değeri içindeki veri üzerinde çalıştırır.

`html/template` paketi güçlüdür; bu program yalnızca yeteneklerine değinir. Özünde, bir HTML metnini, `templ.Execute`’a geçirilen veri öğelerinden türetilen öğelerle değiştirerek çalışma zamanında yeniden yazar; burada veri, form değeridir. Şablon metni içinde (`templateStr`) çift süslü parantezle çevrili parçalar şablon eylemlerini (*template actions*) gösterir. `{{if .}}` ile `{{end}}` arasındaki parça, yalnızca geçerli veri öğesinin değeri—`.` (dot) adı verilir—boş değilse çalışır. Yani string boşken şablonun bu parçası bastırılır.

İki adet `{{.}}` parçası, şablona sunulan verinin—sorgu string’inin—web sayfasında gösterilmesini söyler. HTML şablon paketi, metnin güvenle görüntülenebilmesi için uygun kaçışlamayı (*escaping*) otomatik sağlar.

Şablon string’inin geri kalanı, sayfa yüklendiğinde gösterilecek HTML’dir. Bu açıklama çok hızlı geldiyse, daha kapsamlı bir tartışma için `template` paketinin dokümantasyonuna bakın.

Ve işte bu kadar: birkaç satır kod ve biraz veri güdümlü HTML metniyle faydalı bir web sunucusu. Go, az satırla çok şey olmasını sağlayacak kadar güçlüdür.

