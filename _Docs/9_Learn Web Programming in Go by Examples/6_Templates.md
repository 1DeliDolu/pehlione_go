
## 🧩 Templates

### 📌 Introduction

Go’nun `html/template` paketi, HTML şablonları için zengin bir şablonlama dili sunar. Genellikle web uygulamalarında, verileri istemcinin tarayıcısında yapılandırılmış bir şekilde göstermek için kullanılır. Go’nun şablon dilinin en büyük avantajlarından biri verilerin otomatik olarak *escape* edilmesidir. Go HTML şablonunu ayrıştırır ve tarayıcıda görüntülemeden önce tüm girdileri *escape* ettiği için XSS saldırıları konusunda endişelenmeye gerek yoktur.

---

### 🧪 First Template

Go’da bir şablon yazmak oldukça basittir. Bu örnek, HTML’de sırasız liste (`ul`) olarak yazılmış bir TODO listesini gösterir.

Şablonlar render edilirken içeri verilen veri, Go’nun herhangi bir veri yapısı olabilir. Basit bir string veya sayı olabileceği gibi, aşağıdaki örnekte olduğu gibi iç içe geçmiş veri yapıları da olabilir.

Bir şablonda veriye erişmek için en üst değişkene `{{.}}` ile erişilir. Süslü parantez içindeki nokta (*dot*), *pipeline* olarak adlandırılır ve verinin kök (*root*) elemanını temsil eder.

```go
data := TodoPageData{
    PageTitle: "My TODO list",
    Todos: []Todo{
        {Title: "Task 1", Done: false},
        {Title: "Task 2", Done: true},
        {Title: "Task 3", Done: true},
    },
}
```

```html
<h1>{{.PageTitle}}</h1>
<ul>
    {{range .Todos}}
        {{if .Done}}
            <li class="done">{{.Title}}</li>
        {{else}}
            <li>{{.Title}}</li>
        {{end}}
    {{end}}
</ul>
```

---

### 🧠 Control Structures

Şablon dili, HTML render etmek için zengin bir kontrol yapıları seti içerir. Burada en sık kullanılanların bir özetini göreceksiniz. Tüm olası yapıların detaylı listesi için ziyaret edin: `text/template`

| Control Structure                | Definition                                                             |
| -------------------------------- | ---------------------------------------------------------------------- |
| `{{/* a comment */}}`            | Bir yorum tanımlar                                                     |
| `{{.}}`                          | Kök elemanı render eder                                                |
| `{{.Title}}`                     | İç içe bir elemandaki “Title” alanını render eder                      |
| `{{if .Done}} {{else}} {{end}}`  | Bir if ifadesi tanımlar                                                |
| `{{range .Todos}} {{.}} {{end}}` | Tüm “Todos” üzerinde döngü yapar ve her birini `{{.}}` ile render eder |
| `{{block "content" .}} {{end}}`  | “content” adlı bir blok tanımlar                                       |

---

### 📂 Parsing Templates from Files

Şablonlar bir string’den veya disk üzerindeki bir dosyadan parse edilebilir. Genellikle şablonlar diskten parse edildiği için bu örnek bunu nasıl yapacağınızı gösterir.

Bu örnekte, Go programıyla aynı dizinde `layout.html` adlı bir şablon dosyası vardır.

```go
tmpl, err := template.ParseFiles("layout.html")
// or
tmpl := template.Must(template.ParseFiles("layout.html"))
```

---

### 🌐 Execute a Template in a Request Handler

Şablon diskten parse edildikten sonra request handler içinde kullanılmaya hazırdır.

`Execute` fonksiyonu, şablonu yazdırmak için bir `io.Writer` ve şablona veri geçirmek için bir `interface{}` kabul eder. Fonksiyon bir `http.ResponseWriter` üzerinde çağrıldığında, HTTP response içinde `Content-Type` header’ı otomatik olarak `Content-Type: text/html; charset=utf-8` olarak ayarlanır.

```go
func(w http.ResponseWriter, r *http.Request) {
    tmpl.Execute(w, "data goes here")
}
```

---

### 🧾 The Code (for copy/paste)

Bu, bu örnekte öğrendiklerinizi denemek için kullanabileceğiniz eksiksiz koddur.

```go
package main


import (
    "html/template"
    "net/http"
)



type Todo struct {
    Title string
    Done  bool
}



type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}



func main() {
    tmpl := template.Must(template.ParseFiles("layout.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := TodoPageData{
            PageTitle: "My TODO list",
            Todos: []Todo{
                {Title: "Task 1", Done: false},
                {Title: "Task 2", Done: true},
                {Title: "Task 3", Done: true},
            },
        }
        tmpl.Execute(w, data)
    })
    http.ListenAndServe(":80", nil)
}
```

```html
<h1>{{.PageTitle}}</h1>
<ul>
    {{range .Todos}}
        {{if .Done}}
            <li class="done">{{.Title}}</li>
        {{else}}
            <li>{{.Title}}</li>
        {{end}}
    {{end}}
</ul>
```

