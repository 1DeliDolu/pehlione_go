
## 🧩 Go by Example: XML

Go, `encoding/xml` paketiyle XML ve XML benzeri formatlar için yerleşik destek sunar.

---

## ▶️ Çalıştırma

```go
package main
import (
    "encoding/xml"
    "fmt"
)
```

---

## 🌿 Struct’ı XML’e Eşleme

`Plant` XML’e map edilecektir. JSON örneklerine benzer şekilde, alan tag’leri encoder ve decoder için yönergeler içerir. Burada XML paketinin bazı özel özelliklerini kullanıyoruz: `XMLName` alan adı, bu struct’ı temsil eden XML elemanının adını belirler; `id,attr` ise `Id` alanının iç içe bir eleman yerine bir XML attribute’u olduğunu ifade eder.

```go
type Plant struct {
    XMLName xml.Name `xml:"plant"`
    Id      int      `xml:"id,attr"`
    Name    string   `xml:"name"`
    Origin  []string `xml:"origin"`
}
func (p Plant) String() string {
    return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
        p.Id, p.Name, p.Origin)
}
func main() {
    coffee := &Plant{Id: 27, Name: "Coffee"}
    coffee.Origin = []string{"Ethiopia", "Brazil"}
```

---

## 🧾 XML Üretme ve Pretty Output

Bitkimizi temsil eden XML’i üretelim; daha okunabilir bir çıktı üretmek için `MarshalIndent` kullanıyoruz.

```go
    out, _ := xml.MarshalIndent(coffee, " ", "  ")
    fmt.Println(string(out))
```

Çıktıya genel bir XML header eklemek için, bunu açıkça sonuna ekleyin.

```go
    fmt.Println(xml.Header + string(out))
```

---

## 🧩 XML’i Unmarshal ile Parse Etme

`Unmarshal` ile XML içeren bir byte akışını bir veri yapısına parse edebiliriz. XML hatalı biçimlendirilmişse veya `Plant` üzerine map edilemiyorsa, açıklayıcı bir hata döndürülür.

```go
    var p Plant
    if err := xml.Unmarshal(out, &p); err != nil {
        panic(err)
    }
    fmt.Println(p)
    tomato := &Plant{Id: 81, Name: "Tomato"}
    tomato.Origin = []string{"Mexico", "California"}
```

---

## 🧬 İç İçe XML Yapısı

`parent>child>plant` alan tag’i, encoder’a tüm bitkileri `<parent><child>...` altında iç içe yerleştirmesini söyler.

```go
    type Nesting struct {
        XMLName xml.Name `xml:"nesting"`
        Plants  []*Plant `xml:"parent>child>plant"`
    }
    nesting := &Nesting{}
    nesting.Plants = []*Plant{coffee, tomato}
    out, _ = xml.MarshalIndent(nesting, " ", "  ")
    fmt.Println(string(out))
}
```

---

## 💻 CLI Çıktısı

```bash
$ go run xml.go
 <plant id="27">
   <name>Coffee</name>
   <origin>Ethiopia</origin>
   <origin>Brazil</origin>
 </plant>
<?xml version="1.0" encoding="UTF-8"?>
 <plant id="27">
   <name>Coffee</name>
   <origin>Ethiopia</origin>
   <origin>Brazil</origin>
 </plant>
Plant id=27, name=Coffee, origin=[Ethiopia Brazil]
 <nesting>
   <parent>
     <child>
       <plant id="27">
         <name>Coffee</name>
         <origin>Ethiopia</origin>
         <origin>Brazil</origin>
       </plant>
       <plant id="81">
         <name>Tomato</name>
         <origin>Mexico</origin>
         <origin>California</origin>
       </plant>
     </child>
   </parent>
 </nesting>
```

---

## ⏭️ Sonraki Örnek

Sonraki örnek: Time.

