
## 🧩 Go by Example: Metin Şablonları

Go, kullanıcıya dinamik içerik oluşturmak veya özelleştirilmiş çıktı göstermek için `text/template` paketiyle yerleşik destek sunar. `html/template` adında kardeş bir paket aynı API’yi sağlar, ancak ek güvenlik özelliklerine sahiptir ve HTML üretmek için kullanılmalıdır.

---

## ▶️ Çalıştırma

```go
package main
import (
    "os"
    "text/template"
)
func main() {
```

---

## 🧱 Şablon Oluşturma ve Parse Etme

Yeni bir şablon oluşturup gövdesini bir string’den parse edebiliriz. Şablonlar, statik metin ile `{{...}}` içine alınmış ve içeriği dinamik olarak eklemek için kullanılan “action”ların karışımıdır.

```go
    t1 := template.New("t1")
    t1, err := t1.Parse("Value is {{.}}\n")
    if err != nil {
        panic(err)
    }
```

Alternatif olarak, `template.Must` fonksiyonunu kullanarak `Parse` bir hata döndürürse `panic` ettirebiliriz. Bu, özellikle global scope’ta başlatılan şablonlar için faydalıdır.

```go
    t1 = template.Must(t1.Parse("Value: {{.}}\n"))
```

---

## 🧾 Şablonu Execute Etme

Şablonu “execute” ederek, action’ları için belirli değerler ile metnini üretiriz. `{{.}}` action’ı, `Execute` fonksiyonuna parametre olarak geçirilen değer ile değiştirilir.

```go
    t1.Execute(os.Stdout, "some text")
    t1.Execute(os.Stdout, 5)
    t1.Execute(os.Stdout, []string{
        "Go",
        "Rust",
        "C++",
        "C#",
    })
```

---

## 🛠️ Yardımcı Fonksiyon

Aşağıda kullanacağımız yardımcı fonksiyon.

```go
    Create := func(name, t string) *template.Template {
        return template.Must(template.New(name).Parse(t))
    }
```

---

## 🧑‍💼 Struct Alanlarına Erişim

Veri bir struct ise alanlarına erişmek için `{{.FieldName}}` action’ını kullanabiliriz. Şablon execute edilirken erişilebilir olması için alanların *exported* olması gerekir.

```go
    t2 := Create("t2", "Name: {{.Name}}\n")
    t2.Execute(os.Stdout, struct {
        Name string
    }{"Jane Doe"})
```

---

## 🗺️ Map Anahtarlarına Erişim

Aynı şey map’ler için de geçerlidir; map’lerde anahtar isimlerinin büyük/küçük harf zorunluluğu yoktur.

```go
    t2.Execute(os.Stdout, map[string]string{
        "Name": "Mickey Mouse",
    })
```

---

## 🔀 if/else ile Koşullu Çalıştırma

`if/else`, şablonlar için koşullu çalıştırma sağlar. Bir değer, bir tipin varsayılan değeri ise false kabul edilir; örneğin `0`, boş string, `nil` pointer vb. Bu örnek ayrıca şablonların başka bir özelliğini gösterir: action’larda `-` kullanarak boşluk kırpma.

```go
    t3 := Create("t3",
        "{{if . -}} yes {{else -}} no {{end}}\n")
    t3.Execute(os.Stdout, "not empty")
    t3.Execute(os.Stdout, "")
```

---

## 🔁 range ile Döngü

`range` blokları slice, array, map veya channel üzerinde döngü kurmamıza izin verir. `range` bloğunun içinde `{{.}}`, iterasyondaki mevcut öğeye ayarlanır.

```go
    t4 := Create("t4",
        "Range: {{range .}}{{.}} {{end}}\n")
    t4.Execute(os.Stdout,
        []string{
            "Go",
            "Rust",
            "C++",
            "C#",
        })
}
```

---

## 💻 CLI Çıktısı

```bash
$ go run templates.go 
Value: some text
Value: 5
Value: [Go Rust C++ C#]
Name: Jane Doe
Name: Mickey Mouse
yes 
no 
Range: Go Rust C++ C# 
```

---

## ⏭️ Sonraki Örnek

Sonraki örnek: Regular Expressions.

