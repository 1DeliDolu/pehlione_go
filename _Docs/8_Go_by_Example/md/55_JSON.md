
## 🧩 Go by Example: JSON

Go, yerleşik ve özel veri tiplerine **JSON encode/decode** (kodlama/çözme) desteği sunar.

---

## ▶️ Çalıştırma

```go
package main
import (
    "encoding/json"
    "fmt"
    "os"
    "strings"
)
```

Aşağıdaki iki struct’ı, özel tiplerin encode ve decode edilmesini göstermek için kullanacağız.

```go
type response1 struct {
    Page   int
    Fruits []string
}
```

JSON’da yalnızca *exported* alanlar encode/decode edilir. Alanların exported olması için büyük harfle başlaması gerekir.

```go
type response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}
func main() {
```

---

## 🧾 Temel Tipleri JSON’a Encode Etme

Önce temel veri tiplerini JSON string’lerine encode etmeye bakalım. İşte atomik değerler için bazı örnekler.

```go
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))
    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))
    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))
    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))
```

Slice ve map örnekleri de aşağıdaki gibi; beklendiği üzere JSON array ve object olarak encode edilirler.

```go
    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))
    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))
```

JSON paketi, özel veri tiplerinizi otomatik olarak encode edebilir. Encode edilen çıktıya yalnızca exported alanları dahil eder ve varsayılan olarak bu alan adlarını JSON key’leri olarak kullanır.

```go
    res1D := &response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))
```

Struct alan deklarasyonlarında tag kullanarak encode edilen JSON key isimlerini özelleştirebilirsiniz. Örnek için yukarıdaki `response2` tanımına bakın.

```go
    res2D := &response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))
```

---

## 🧩 JSON’ı Go Değerlerine Decode Etme

Şimdi JSON verisini Go değerlerine decode etmeye bakalım. İşte generic bir veri yapısı için bir örnek.

```go
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
```

JSON paketinin decode edilmiş veriyi koyabileceği bir değişken sağlamamız gerekir. Bu `map[string]interface{}` string’lerden rastgele veri tiplerine bir map tutacaktır.

```go
    var dat map[string]interface{}
```

Asıl decode işlemi ve ilgili hata kontrolü:

```go
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)
```

Decode edilen map içindeki değerleri kullanmak için, uygun tiplerine dönüştürmemiz gerekir. Örneğin burada `num` içindeki değeri beklenen `float64` tipine dönüştürüyoruz.

```go
    num := dat["num"].(float64)
    fmt.Println(num)
```

İç içe veriye erişmek, bir dizi dönüşüm gerektirir.

```go
    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1)
```

JSON’ı özel veri tiplerine de decode edebiliriz. Bunun avantajı, programlarımıza ek *type-safety* katması ve decode edilen verilere erişirken type assertion ihtiyacını ortadan kaldırmasıdır.

```go
    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])
```

Yukarıdaki örneklerde, veriler ile standard output’taki JSON temsili arasında her zaman bytes ve string’leri ara katman olarak kullandık. JSON encode’larını doğrudan `os.Stdout` gibi `os.Writer`’lara veya hatta HTTP response body’lerine stream edebiliriz.

```go
    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
```

`os.Stdin` veya HTTP request body’leri gibi `os.Reader`’lardan stream okuma, `json.Decoder` ile yapılır.

```go
    dec := json.NewDecoder(strings.NewReader(str))
    res1 := response2{}
    dec.Decode(&res1)
    fmt.Println(res1)
}
```

---

## 💻 CLI Çıktısı

```bash
$ go run json.go
true
1
2.34
"gopher"
["apple","peach","pear"]
{"apple":5,"lettuce":7}
{"Page":1,"Fruits":["apple","peach","pear"]}
{"page":1,"fruits":["apple","peach","pear"]}
map[num:6.13 strs:[a b]]
6.13
a
{1 [apple peach]}
apple
{"apple":5,"lettuce":7}
{1 [apple peach]}
```

---

## 📚 Ek Not

Go’da JSON’un temellerini burada ele aldık; ancak daha fazlası için JSON and Go blog yazısına ve JSON paket dokümanlarına göz atın.

---

## ⏭️ Sonraki Örnek

Sonraki örnek: XML.

