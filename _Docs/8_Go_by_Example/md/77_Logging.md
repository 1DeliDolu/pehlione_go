
## 🧾 Loglama

Go standart kütüphanesi, Go programlarından log çıktısı almak için doğrudan araçlar sağlar: serbest biçimli çıktı için `log` paketi ve yapılandırılmış çıktı için `log/slog` paketi.

---

## ▶️ Çalıştırma

```go
package main
```

```go
import (
    "bytes"
    "fmt"
    "log"
    "os"
    "log/slog"
)
```

```go
func main() {
```

---

## 🪵 Standart Logger ile Log Basma

`log` paketinden `Println` gibi fonksiyonları doğrudan çağırmak, zaten `os.Stderr`’a makul bir log çıktısı verecek şekilde önceden yapılandırılmış olan standart logger’ı kullanır. `Fatal*` veya `Panic*` gibi ek metotlar ise logladıktan sonra programdan çıkacaktır.

```go
    log.Println("standard logger")
```

---

## ⚙️ Çıktı Formatını Bayraklarla Ayarlama

Logger’lar, çıktı formatlarını belirlemek için bayraklarla (*flags*) yapılandırılabilir. Varsayılan olarak standart logger’da `log.Ldate` ve `log.Ltime` bayrakları ayarlıdır ve bunlar `log.LstdFlags` içinde toplanır. Örneğin, zamanı mikrosaniye doğruluğuyla üretmesi için bayrakları değiştirebiliriz.

```go
    log.SetFlags(log.LstdFlags | log.Lmicroseconds)
    log.Println("with micro")
```

Ayrıca log fonksiyonunun çağrıldığı dosya adını ve satır numarasını üretmeyi de destekler.

```go
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    log.Println("with file/line")
```

---

## 🧰 Özel Logger Oluşturma

Özel bir logger oluşturup etrafta dolaştırmak faydalı olabilir. Yeni bir logger oluştururken, çıktısını diğer logger’lardan ayırmak için bir prefix ayarlayabiliriz.

```go
    mylog := log.New(os.Stdout, "my:", log.LstdFlags)
    mylog.Println("from mylog")
```

Mevcut logger’larda (standart logger dahil) `SetPrefix` metodu ile prefix ayarlanabilir.

```go
    mylog.SetPrefix("ohmy:")
    mylog.Println("from mylog")
```

---

## 🎯 Özel Çıktı Hedefi Kullanma

Logger’lar özel çıktı hedeflerine sahip olabilir; herhangi bir `io.Writer` çalışır.

```go
    var buf bytes.Buffer
    buflog := log.New(&buf, "buf:", log.LstdFlags)
```

Bu çağrı, log çıktısını `buf` içine yazar.

```go
    buflog.Println("hello")
```

Bu da, onu standart çıktıda gerçekten gösterecektir.

```go
    fmt.Print("from buflog:", buf.String())
```

---

## 🧱 Yapılandırılmış Log: slog

`slog` paketi, yapılandırılmış log çıktısı sağlar. Örneğin JSON formatında loglamak oldukça doğrudandır.

```go
    jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
    myslog := slog.New(jsonHandler)
    myslog.Info("hi there")
```

Mesaja ek olarak, `slog` çıktısı key=value çiftlerinden oluşan keyfi sayıda alan içerebilir.

```go
    myslog.Info("hello again", "key", "val", "age", 25)
}
```

---

## 🧪 Örnek Çıktı

Örnek çıktı; basılan tarih ve saat, örneğin ne zaman çalıştırıldığına bağlı olacaktır.

```bash
$ go run logging.go
2023/08/22 10:45:16 standard logger
2023/08/22 10:45:16.904141 with micro
2023/08/22 10:45:16 logging.go:40: with file/line
my:2023/08/22 10:45:16 from mylog
ohmy:2023/08/22 10:45:16 from mylog
from buflog:buf:2023/08/22 10:45:16 hello
```

Bunlar, web sitesinde sunumun daha anlaşılır olması için satır kırılarak gösterilmiştir; gerçekte tek satır halinde basılırlar.

```json
{"time":"2023-08-22T10:45:16.904166391-07:00",
 "level":"INFO","msg":"hi there"}
{"time":"2023-08-22T10:45:16.904178985-07:00",
    "level":"INFO","msg":"hello again",
    "key":"val","age":25}
```

---

## ➡️ Sonraki Örnek

Sonraki örnek: **HTTP Client**.

