# 🌐 Web Uygulamaları Yazma

## 📚 İçindekiler

Giriş
Başlarken
Veri Yapıları
`net/http` paketine giriş (kısa ara)
`net/http` ile wiki sayfaları sunma
Sayfaları Düzenleme
`html/template` paketi
Var olmayan sayfaları ele alma
Sayfaları Kaydetme
Hata yönetimi
Şablon önbellekleme
Doğrulama
Fonksiyon literal’leri ve closure’lara giriş
Deneyin!
Diğer görevler

---

## 🧭 Giriş¶

Bu eğitimde kapsananlar:

* Yükleme ve kaydetme metotlarına sahip bir veri yapısı oluşturma
* `net/http` paketini kullanarak web uygulamaları geliştirme
* HTML şablonlarını işlemek için `html/template` paketini kullanma
* Kullanıcı girdisini doğrulamak için `regexp` paketini kullanma
* Closure’ları kullanma

Varsayılan bilgi düzeyi:

* Programlama deneyimi
* Temel web teknolojilerini anlama (HTTP, HTML)
* Biraz UNIX/DOS komut satırı bilgisi

---

## 🚀 Başlarken¶

Şu anda Go çalıştırmak için FreeBSD, Linux, macOS veya Windows bir makineye sahip olmanız gerekir. Komut istemini temsil etmek için `$` kullanacağız.

Go’yu kurun (Installation Instructions’a bakın).

Bu eğitim için GOPATH’iniz içinde yeni bir dizin oluşturun ve içine girin:

```bash
$ mkdir gowiki
$ cd gowiki
```

`wiki.go` adında bir dosya oluşturun, favori editörünüzle açın ve aşağıdaki satırları ekleyin:

```go
package main

import (
    "fmt"
    "os"
)
```

Go standart kütüphanesinden `fmt` ve `os` paketlerini içe aktarıyoruz. Daha sonra ek işlevleri uyguladıkça, bu import bildirimine daha fazla paket ekleyeceğiz.

---

## 🧱 Veri Yapıları¶

Veri yapılarıyla başlayalım. Bir wiki, birbirine bağlı sayfalardan oluşur; her sayfanın bir başlığı ve bir gövdesi (sayfa içeriği) vardır. Burada, başlık ve gövdeyi temsil eden iki alanla bir `Page` struct’ı tanımlıyoruz.

```go
type Page struct {
    Title string
    Body  []byte
}
```

`[]byte` türü “byte dilimi (slice)” anlamına gelir. (`slice`’lar hakkında daha fazlası için  *Slices: usage and internals* ’a bakın.) `Body` öğesi `string` yerine `[]byte`’tır; çünkü aşağıda göreceğiniz gibi, kullanacağımız `io` kütüphanelerinin beklediği tür budur.

`Page` struct’ı, sayfa verisinin bellekte nasıl tutulacağını açıklar. Peki kalıcı depolama (persistent storage) ne olacak? Bunu, `Page` üzerinde bir `save` metodu oluşturarak çözebiliriz:

```go
func (p *Page) save() error {
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
}
```

Bu metodun imzası şöyle okunur: “Bu, alıcı (receiver) olarak `Page` işaretçisi olan `p`’yi alan `save` adlı bir metottur. Parametre almaz ve `error` türünde bir değer döndürür.”

Bu metod, `Page`’in `Body` alanını bir metin dosyasına kaydeder. Basitlik için, dosya adı olarak `Title` kullanacağız.

`save` metodu bir `error` değeri döndürür; çünkü bu, `WriteFile`’ın (byte dilimini dosyaya yazan standart kütüphane fonksiyonu) dönüş türüdür. `save` metodu bu hata değerini döndürür; böylece yazma sırasında bir şeyler ters giderse uygulama bunu ele alabilir. Her şey yolunda giderse `Page.save()` `nil` döndürür (işaretçiler, arayüzler ve bazı diğer türler için sıfır-değer).

`WriteFile`’a üçüncü parametre olarak verilen sekizlik (octal) tamsayı literal’i `0600`, dosyanın yalnızca mevcut kullanıcı için okuma-yazma izinleriyle oluşturulacağını belirtir. (Detaylar için Unix `open(2)` man sayfasına bakın.)

Sayfaları kaydetmenin yanında, sayfaları yüklemek de isteyeceğiz:

```go
func loadPage(title string) *Page {
    filename := title + ".txt"
    body, _ := os.ReadFile(filename)
    return &Page{Title: title, Body: body}
}
```

`loadPage` fonksiyonu, `title` parametresinden dosya adını oluşturur, dosyanın içeriğini `body` adlı yeni bir değişkene okur ve doğru `title` ve `body` değerleriyle oluşturulmuş bir `Page` literal’ına işaretçi döndürür.

Fonksiyonlar birden fazla değer döndürebilir. Standart kütüphanedeki `os.ReadFile` fonksiyonu `[]byte` ve `error` döndürür. `loadPage` içinde hata henüz ele alınmıyor; alt çizgi (`_`) ile gösterilen “blank identifier”, hata dönüş değerini atmak için kullanılır (özünde, değeri hiçbir şeye atamak).

Ama `ReadFile` hata ile karşılaşırsa ne olur? Örneğin dosya mevcut olmayabilir. Bu hataları görmezden gelmemeliyiz. Fonksiyonu `*Page` ve `error` döndürecek şekilde değiştirelim.

```go
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}
```

Bu fonksiyonun çağıranları artık ikinci parametreyi kontrol edebilir; eğer `nil` ise bir `Page` başarıyla yüklenmiştir. Değilse, çağıran tarafından ele alınabilecek bir hata olacaktır (detaylar için dil spesifikasyonuna bakın).

Bu noktada basit bir veri yapımız ve dosyaya kaydedip dosyadan yükleyebilme yeteneğimiz var. Yazdıklarımızı test etmek için bir `main` fonksiyonu yazalım:

```go
func main() {
    p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    p1.save()
    p2, _ := loadPage("TestPage")
    fmt.Println(string(p2.Body))
}
```

Bu kod derlenip çalıştırıldıktan sonra, `TestPage.txt` adlı bir dosya oluşturulur ve `p1`’in içerikleri bu dosyaya yazılır. Ardından dosya `p2` struct’ına okunur ve `Body` alanı ekrana basılır.

Programı şöyle derleyip çalıştırabilirsiniz:

```bash
$ go build wiki.go
$ ./wiki
```

```text
This is a sample Page.
```

(Windows kullanıyorsanız programı çalıştırmak için `./wiki` yerine `"wiki"` yazmanız gerekir.)

Buraya kadar yazdığımız kodu görmek için buraya tıklayın.

---

## 🛰️ `net/http` paketine giriş (kısa ara)¶

İşte basit bir web sunucusunun çalışan tam bir örneği:

```go
//go:build ignore

package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

`main` fonksiyonu, `http.HandleFunc` çağrısıyla başlar; bu, `http` paketine web köküne (`"/"`) gelen tüm istekleri `handler` ile işlemesini söyler.

Ardından `http.ListenAndServe` çağrılır; bu çağrı, herhangi bir arayüzde 8080 portunu dinlemesini belirtir (`":8080"`). (Şimdilik ikinci parametre olan `nil` konusunda endişelenmeyin.) Bu fonksiyon, program sonlandırılana kadar bloklar.

`ListenAndServe` her zaman bir hata döndürür; çünkü yalnızca beklenmeyen bir hata oluştuğunda döner. Bu hatayı loglamak için fonksiyon çağrısını `log.Fatal` ile sararız.

`handler` fonksiyonu `http.HandlerFunc` türündedir. Argüman olarak `http.ResponseWriter` ve `http.Request` alır.

Bir `http.ResponseWriter` değeri, HTTP sunucusunun yanıtını oluşturur; ona yazarak HTTP istemcisine veri göndeririz.

Bir `http.Request`, istemci HTTP isteğini temsil eden bir veri yapısıdır. `r.URL.Path`, istek URL’sinin path bileşenidir. Sondaki `[1:]`, “Path’in 1. karakterinden sonuna kadar bir alt dilim oluştur” demektir. Bu, path adından baştaki `"/"` karakterini düşürür.

Bu programı çalıştırır ve şu URL’ye erişirseniz:

```text
http://localhost:8080/monkeys
```

Program şu içeriği içeren bir sayfa sunar:

```text
Hi there, I love monkeys!
```

---

## 🧩 `net/http` ile wiki sayfaları sunma¶

`net/http` paketini kullanmak için import edilmesi gerekir:

```go
import (
    "fmt"
    "os"
    "log"
    "net/http"
)
```

Kullanıcıların bir wiki sayfasını görüntülemesine izin verecek bir handler oluşturalım: `viewHandler`. Bu handler, `"/view/"` ile başlayan URL’leri ele alacaktır.

```go
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
```

Yine, `loadPage`’den dönen hata değerini yok saymak için `_` kullanıldığına dikkat edin. Bu, burada basitlik için yapılmıştır ve genel olarak kötü bir uygulama kabul edilir. Buna daha sonra değineceğiz.

Önce, bu fonksiyon sayfa başlığını `r.URL.Path`’ten (istek URL’sinin path bileşeni) çıkarır. Path, `[len("/view/"):]` ile tekrar dilimlenir ve path’in başındaki `"/view/"` bileşeni atılır. Bunun nedeni, path’in her zaman `"/view/"` ile başlayacak olmasıdır; bu kısım sayfanın başlığının parçası değildir.

Fonksiyon daha sonra sayfa verisini yükler, sayfayı basit bir HTML string’i ile formatlar ve `http.ResponseWriter` olan `w`’ye yazar.

Bu handler’ı kullanmak için, `main` fonksiyonumuzu `http`’yi başlatacak ve `/view/` altındaki tüm isteklerde `viewHandler`’ı kullanacak şekilde yeniden yazarız.

```go
func main() {
    http.HandleFunc("/view/", viewHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Buraya kadar yazdığımız kodu görmek için buraya tıklayın.

Bir test sayfa verisi oluşturalım (`test.txt` olarak), kodumuzu derleyelim ve bir wiki sayfası sunmayı deneyelim.

Editörünüzde `test.txt` dosyasını açın ve içine `"Hello world"` (tırnaklar olmadan) yazıp kaydedin.

```bash
$ go build wiki.go
$ ./wiki
```

(Windows kullanıyorsanız programı çalıştırmak için `./wiki` yerine `"wiki"` yazmanız gerekir.)

Bu web sunucusu çalışırken, `http://localhost:8080/view/test` adresine yapılan bir ziyaret, başlığı `"test"` olan ve içinde `"Hello world"` yazan bir sayfa göstermelidir.

---

## ✏️ Sayfaları Düzenleme¶

Bir wiki, sayfaları düzenleme yeteneği olmadan wiki değildir. İki yeni handler oluşturalım: biri düzenleme formunu göstermek için `editHandler`, diğeri formdan girilen veriyi kaydetmek için `saveHandler`.

Önce, bunları `main()` içine ekleriz:

```go
func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    http.HandleFunc("/save/", saveHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

`editHandler` fonksiyonu sayfayı yükler (ya da eğer sayfa yoksa boş bir `Page` struct’ı oluşturur) ve bir HTML formu gösterir.

```go
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    fmt.Fprintf(w, "<h1>Editing %s</h1>"+
        "<form action=\"/save/%s\" method=\"POST\">"+
        "<textarea name=\"body\">%s</textarea><br>"+
        "<input type=\"submit\" value=\"Save\">"+
        "</form>",
        p.Title, p.Title, p.Body)
}
```

Bu fonksiyon çalışır; ancak tüm o hard-coded HTML çirkindir. Elbette daha iyi bir yol var.

---

## 🧾 `html/template` paketi¶

`html/template` paketi Go standart kütüphanesinin bir parçasıdır. HTML’yi ayrı bir dosyada tutmak için `html/template` kullanabiliriz; böylece alttaki Go kodunu değiştirmeden edit sayfamızın düzenini değiştirebiliriz.

Önce, `html/template`’i import listesine eklemeliyiz. Ayrıca artık `fmt` kullanmayacağız; dolayısıyla onu kaldırmamız gerekir.

```go
import (
    "html/template"
    "os"
    "net/http"
)
```

HTML formunu içeren bir şablon dosyası oluşturalım. `edit.html` adlı yeni bir dosya açın ve aşağıdaki satırları ekleyin:

```html
<h1>Editing {{.Title}}</h1>

<form action="/save/{{.Title}}" method="POST">
<div><textarea name="body" rows="20" cols="80">{{printf "%s" .Body}}</textarea></div>
<div><input type="submit" value="Save"></div>
</form>
```

`editHandler`’ı, hard-coded HTML yerine şablonu kullanacak şekilde değiştirin:

```go
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    t, _ := template.ParseFiles("edit.html")
    t.Execute(w, p)
}
```

`template.ParseFiles` fonksiyonu `edit.html` içeriğini okuyacak ve bir `*template.Template` döndürecektir.

`t.Execute` metodu şablonu çalıştırır ve üretilen HTML’yi `http.ResponseWriter`’a yazar. `.Title` ve `.Body` noktalı tanımlayıcıları, `p.Title` ve `p.Body`’ye referans verir.

Şablon yönergeleri (directive) çift süslü parantezler içinde yazılır. `printf "%s" .Body` talimatı, `.Body`’yi byte akışı yerine string olarak çıktılar; bu, `fmt.Printf` çağrısının aynısıdır. `html/template` paketi, şablon aksiyonlarının yalnızca güvenli ve doğru görünümlü HTML üretmesini garanti etmeye yardımcı olur. Örneğin, otomatik olarak `>` işaretini escape eder ve `&gt;` ile değiştirir; böylece kullanıcı verisi form HTML’sini bozmaz.

Şablonlarla çalıştığımıza göre, `viewHandler` için `view.html` adlı bir şablon oluşturalım:

```html
<h1>{{.Title}}</h1>

<p>[<a href="/edit/{{.Title}}">edit</a>]</p>

<div>{{printf "%s" .Body}}</div>
```

`viewHandler`’ı buna göre değiştirin:

```go
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    t, _ := template.ParseFiles("view.html")
    t.Execute(w, p)
}
```

Her iki handler’da da neredeyse aynı şablon kodunu kullandığımıza dikkat edin. Bu tekrarları, şablonlama kodunu ayrı bir fonksiyona taşıyarak kaldıralım:

```go
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles(tmpl + ".html")
    t.Execute(w, p)
}
```

Handler’ları bu fonksiyonu kullanacak şekilde değiştirin:

```go
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    renderTemplate(w, "view", p)
}
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}
```

Eğer `main` içindeki uygulanmamış `saveHandler` kaydını yorum satırı yaparsak, programımızı yeniden derleyip test edebiliriz. Buraya kadar yazdığımız kodu görmek için buraya tıklayın.

---

## 🚫 Var olmayan sayfaları ele alma¶

Peki `/view/APageThatDoesntExist` adresini ziyaret ederseniz ne olur? HTML içeren bir sayfa görürsünüz. Bunun nedeni, `loadPage`’in hata dönüş değerini yok sayması ve veri olmadan şablonu doldurmaya devam etmesidir. Bunun yerine, istenen `Page` mevcut değilse, içerik oluşturulabilmesi için istemciyi edit sayfasına yönlendirmelidir:

```go
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}
```

`http.Redirect` fonksiyonu, HTTP yanıtına `http.StatusFound` (302) durum kodunu ve bir `Location` header’ını ekler.

---

## 💾 Sayfaları Kaydetme¶

`saveHandler` fonksiyonu, edit sayfalarında bulunan formların gönderimini ele alır. `main` içindeki ilgili satırın yorumunu kaldırdıktan sonra handler’ı uygulayalım:

```go
func saveHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/save/"):]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    p.save()
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
```

Sayfa başlığı (URL’de sağlanan) ve formun tek alanı olan `Body`, yeni bir `Page` içinde tutulur. Ardından `save()` metodu çağrılarak veri dosyaya yazılır ve istemci `/view/` sayfasına yönlendirilir.

`FormValue`’nun döndürdüğü değer `string` türündedir. Bu değerin `Page` struct’ına sığabilmesi için `[]byte`’a çevrilmesi gerekir. Dönüşüm için `[]byte(body)` kullanırız.

---

## 🧯 Hata yönetimi¶

Programımızda hataların yok sayıldığı birkaç yer var. Bu kötü bir pratiktir; en azından, bir hata oluştuğunda program beklenmeyen davranış sergiler. Daha iyi bir çözüm, hataları ele almak ve kullanıcıya bir hata mesajı döndürmektir. Böylece bir şeyler ters gittiğinde sunucu tam olarak istediğimiz gibi çalışır ve kullanıcı bilgilendirilebilir.

Önce, `renderTemplate` içindeki hataları ele alalım:

```go
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, err := template.ParseFiles(tmpl + ".html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
```

`http.Error` fonksiyonu, belirtilen bir HTTP yanıt kodu (bu durumda “Internal Server Error”) ve hata mesajını gönderir. Bunu ayrı bir fonksiyona koyma kararı şimdiden karşılığını veriyor.

Şimdi `saveHandler`’ı düzeltelim:

```go
func saveHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/save/"):]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
```

`p.save()` sırasında oluşan hatalar kullanıcıya raporlanacaktır.

---

## 🗂️ Şablon önbellekleme¶

Bu kodda bir verimsizlik var: `renderTemplate`, bir sayfa render edildiğinde her seferinde `ParseFiles` çağırır. Daha iyi bir yaklaşım, program başlatılırken `ParseFiles`’ı bir kez çağırıp tüm şablonları tek bir `*Template` içine parse etmektir. Sonra belirli bir şablonu render etmek için `ExecuteTemplate` metodunu kullanabiliriz.

Önce `templates` adlı global bir değişken oluşturur ve `ParseFiles` ile başlatırız.

```go
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
```

`template.Must`, kendisine `nil` olmayan bir hata değeri verildiğinde panic atan; aksi halde `*Template`’ı değiştirmeden döndüren bir kolaylık sarmalayıcısıdır. Burada panic uygundur; şablonlar yüklenemiyorsa yapılabilecek tek mantıklı şey programdan çıkmaktır.

`ParseFiles`, şablon dosyalarını tanımlayan bir veya daha fazla string argüman alır ve bu dosyaları temel dosya adına göre isimlendirilmiş şablonlar halinde parse eder. Programımıza daha fazla şablon ekleseydik, adlarını `ParseFiles` çağrısının argümanlarına eklerdik.

Ardından `renderTemplate` fonksiyonunu, uygun şablon adını kullanarak `templates.ExecuteTemplate` metodunu çağıracak şekilde değiştiririz:

```go
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
```

Şablon adı dosya adı olduğu için, `tmpl` argümanına `".html"` eklememiz gerekir.

---

## ✅ Doğrulama¶

Fark etmiş olabileceğiniz gibi, bu programda ciddi bir güvenlik açığı vardır: bir kullanıcı, sunucuda okunacak/yazılacak keyfi (arbitrary) bir path sağlayabilir. Bunu azaltmak için, başlığı bir regular expression ile doğrulayan bir fonksiyon yazabiliriz.

Önce import listesine `"regexp"` ekleyin. Sonra doğrulama ifademizi saklamak için bir global değişken oluşturabiliriz:

```go
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
```

`regexp.MustCompile` fonksiyonu regular expression’ı parse edip derler ve bir `regexp.Regexp` döndürür. `MustCompile`, derleme başarısız olursa panic atmasıyla `Compile`’dan ayrılır; `Compile` ise ikinci değer olarak hata döndürür.

Şimdi, `validPath` ifadesini kullanarak path’i doğrulayan ve sayfa başlığını çıkaran bir fonksiyon yazalım:

```go
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        http.NotFound(w, r)
        return "", errors.New("invalid Page Title")
    }
    return m[2], nil // The title is the second subexpression.
}
```

Başlık geçerliyse, `nil` hata değeriyle birlikte döndürülür. Başlık geçersizse, fonksiyon HTTP bağlantısına `http.NotFound` ile bir “404 Not Found” hatası yazar ve handler’a bir hata döndürür. Yeni bir hata oluşturmak için `errors` paketini import etmemiz gerekir.

Handler’ların her birine `getTitle` çağrısı ekleyelim:

```go
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title, err := getTitle(w, r)
    if err != nil {
        return
    }
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}
func editHandler(w http.ResponseWriter, r *http.Request) {
    title, err := getTitle(w, r)
    if err != nil {
        return
    }
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}
func saveHandler(w http.ResponseWriter, r *http.Request) {
    title, err := getTitle(w, r)
    if err != nil {
        return
    }
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err = p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
```

---

## 🧠 Fonksiyon literal’leri ve closure’lara giriş¶

Her handler’da hata koşulunu yakalamak çok fazla tekrar eden kod getirir. Peki her bir handler’ı, bu doğrulamayı ve hata kontrolünü yapan bir fonksiyonla sarmalayabilsek ne olurdu? Go’nun fonksiyon literal’leri, burada bize yardımcı olabilecek güçlü bir soyutlama aracıdır.

Önce handler’ların her birinin fonksiyon tanımını, bir `title string` kabul edecek şekilde yeniden yazarız:

```go
func viewHandler(w http.ResponseWriter, r *http.Request, title string)
func editHandler(w http.ResponseWriter, r *http.Request, title string)
func saveHandler(w http.ResponseWriter, r *http.Request, title string)
```

Şimdi, yukarıdaki türde bir fonksiyon alan ve `http.HandleFunc`’a verilmeye uygun `http.HandlerFunc` türünde bir fonksiyon döndüren bir wrapper fonksiyon tanımlayalım:

```go
func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Here we will extract the page title from the Request,
        // and call the provided handler 'fn'
    }
}
```

Dönen fonksiyon, dışarıda tanımlanan değerleri “içine aldığı” için closure olarak adlandırılır. Bu durumda `fn` değişkeni (makeHandler’a verilen tek argüman) closure tarafından içine alınır. `fn`, bizim `save`, `edit` veya `view` handler’larımızdan biri olacaktır.

Şimdi `getTitle` içindeki kodu alıp burada kullanabiliriz (bazı küçük değişikliklerle):

```go
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }
        fn(w, r, m[2])
    }
}
```

`makeHandler`’ın döndürdüğü closure, `http.ResponseWriter` ve `http.Request` alan bir fonksiyondur (yani bir `http.HandlerFunc`). Closure, başlığı istek path’inden çıkarır ve `validPath` regexp’i ile doğrular. Başlık geçersizse, `http.NotFound` fonksiyonu ile `ResponseWriter`’a bir hata yazılır. Başlık geçerliyse, içine alınan handler fonksiyonu `fn`, `ResponseWriter`, `Request` ve `title` argümanlarıyla çağrılır.

Artık handler fonksiyonlarını `main` içinde, `http` paketine kaydetmeden önce `makeHandler` ile sarabiliriz:

```go
func main() {
    http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))

    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Son olarak, handler fonksiyonlarından `getTitle` çağrılarını kaldırırız; bu da onları çok daha basit hale getirir:

```go
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
```

---

## 🧪 Deneyin!¶

Kodun son halini görmek için buraya tıklayın.

Kodu yeniden derleyin ve uygulamayı çalıştırın:

```bash
$ go build wiki.go
$ ./wiki
```

`http://localhost:8080/view/ANewPage` adresini ziyaret etmek, size sayfa düzenleme formunu göstermelidir. Sonrasında bir metin girip ‘Save’ düğmesine tıklayabilir ve yeni oluşturulan sayfaya yönlendirilmelisiniz.

---

## 🛠️ Diğer görevler¶

Kendi başınıza ele almak isteyebileceğiniz bazı basit görevler:

* Şablonları `tmpl/` içinde, sayfa verisini `data/` içinde saklayın.
* Web kökünün `/view/FrontPage`’e yönlendirilmesini sağlayan bir handler ekleyin.
* Sayfa şablonlarını geçerli HTML yapıp bazı CSS kuralları ekleyerek güzelleştirin.
* Sayfalar arası linklemeyi, `[PageName]` örneklerini
  `<a href="/view/PageName">PageName</a>`
  biçimine çevirerek uygulayın. (İpucu: bunun için `regexp.ReplaceAllFunc` kullanabilirsiniz.)
