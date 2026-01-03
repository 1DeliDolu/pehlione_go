
## 🧩 Request Body’yi Farklı Struct’lara Bind Etmeyi Denemek

Request body’yi bind etmek için kullanılan normal yöntemler `c.Request.Body`’yi tüketir ve bu yüzden birden fazla kez çağrılamazlar.

```go
type formA struct {
  Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
  Bar string `json:"bar" xml:"bar" binding:"required"`
}

func SomeHandler(c *gin.Context) {
  objA := formA{}
  objB := formB{}
  // Bu c.ShouldBind, c.Request.Body’yi tüketir ve yeniden kullanılamaz.
  if errA := c.ShouldBind(&objA); errA == nil {
    c.String(http.StatusOK, `the body should be formA`)
  // c.Request.Body artık EOF olduğu için burada her zaman hata oluşur.
  } else if errB := c.ShouldBind(&objB); errB == nil {
    c.String(http.StatusOK, `the body should be formB`)
  } else {
    ...
  }
}
```

Bunun için `c.ShouldBindBodyWith` kullanabilirsiniz.

```go
func SomeHandler(c *gin.Context) {
  objA := formA{}
  objB := formB{}
  // Bu, c.Request.Body’yi okur ve sonucu context içine kaydeder.
  if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
    c.String(http.StatusOK, `the body should be formA`)
  // Bu noktada, context içine kaydedilmiş body yeniden kullanılır.
  } else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
    c.String(http.StatusOK, `the body should be formB JSON`)
  // Ayrıca diğer formatları da kabul edebilir
  } else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
    c.String(http.StatusOK, `the body should be formB XML`)
  } else {
    ...
  }
}
```

`c.ShouldBindBodyWith`, bind işleminden önce body’yi context içine kaydeder. Bunun performansa küçük bir etkisi vardır; bu nedenle tek seferde binding çağırmanız yeterliyse bu yöntemi kullanmamalısınız.

Bu özellik yalnızca bazı formatlar için gereklidir — JSON, XML, MsgPack, ProtoBuf. Diğer formatlar için (*Query*, *Form*, *FormPost*, *FormMultipart*), `c.ShouldBind()` performansa zarar vermeden birden fazla kez çağrılabilir (Bkz. #1341).

