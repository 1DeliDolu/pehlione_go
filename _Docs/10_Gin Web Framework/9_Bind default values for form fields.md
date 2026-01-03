
## 🧩 Form Alanları İçin Varsayılan Değerleri Bind Etmek

Bazen istemci bir değer göndermediğinde, alanın varsayılan bir değere geri dönmesini istersiniz. Gin’in form binding’i, struct tag içindeki `form` etiketinde `default` seçeneğiyle varsayılanları destekler. Bu, skalerler için ve **Gin v1.11’den itibaren**, açık koleksiyon formatlarıyla koleksiyonlar (*slice/array*) için çalışır.

---

## ✅ Temel Noktalar

* Varsayılanı, `form` anahtarından hemen sonra yazın: `form:"name,default=William"`.
* Koleksiyonlar için değerlerin nasıl bölüneceğini `collection_format:"multi|csv|ssv|tsv|pipes"` ile belirtin.
* `multi` ve `csv` için varsayılanda değerleri ayırmak üzere **noktalı virgül** kullanın (ör. `default=1;2;3`). Gin, tag parser’ı belirsiz kalmasın diye bunları dahili olarak virgüle çevirir.
* `ssv` (boşluk), `tsv` (tab) ve `pipes` (`|`) için varsayılanda doğal ayırıcıyı kullanın.

---

## 🧪 Örnek

```go
package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

type Person struct {
  Name      string    `form:"name,default=William"`
  Age       int       `form:"age,default=10"`
  Friends   []string  `form:"friends,default=Will;Bill"`                     // multi/csv: varsayılanlarda ; kullan
  Addresses [2]string `form:"addresses,default=foo bar" collection_format:"ssv"`
  LapTimes  []int     `form:"lap_times,default=1;2;3" collection_format:"csv"`
}

func main() {
  r := gin.Default()
  r.POST("/person", func(c *gin.Context) {
    var req Person
    if err := c.ShouldBind(&req); err != nil { // Content-Type’a göre binder’ı çıkarır
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    c.JSON(http.StatusOK, req)
  })
  _ = r.Run(":8080")
}
```

---

## 📮 Body Olmadan POST Atarsanız

Body olmadan POST ederseniz Gin varsayılanlarla yanıt verir:

Terminal penceresi

```bash
curl -X POST http://localhost:8080/person
```

Yanıt (örnek):

```json
{
  "Name": "William",
  "Age": 10,
  "Friends": ["Will", "Bill"],
  "Addresses": ["foo", "bar"],
  "LapTimes": [1, 2, 3]
}
```

---

## 📝 Notlar ve Dikkat Edilecekler

* Go struct tag sözdizimi, seçenekleri ayırmak için virgül kullanır; varsayılan değerlerin içinde virgül kullanmaktan kaçının.
* `multi` ve `csv` için noktalı virgüller varsayılan değerleri ayırır; bu formatlarda tekil varsayılanların içine noktalı virgül koymayın.
* Geçersiz `collection_format` değerleri bir binding hatasına yol açar.

---

## 🔄 İlgili Değişiklikler

* Form binding için koleksiyon formatları (`multi`, `csv`, `ssv`, `tsv`, `pipes`) v1.11 civarında geliştirildi.
* Koleksiyonlar için varsayılan değerler v1.11’de eklendi (PR #4048).

