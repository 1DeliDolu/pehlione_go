
## 🧬 Süreç Başlatma

Bazen Go programlarımızın başka süreçleri (*processes*) başlatması gerekir.

---

## ▶️ Çalıştırma

```go
package main
```

```go
import (
    "errors"
    "fmt"
    "io"
    "os/exec"
)
```

```go
func main() {
```

---

## 📅 Basit Bir Komut Çalıştırma

Argüman veya girdi almayan ve sadece stdout’a bir şey yazdıran basit bir komutla başlayacağız. `exec.Command` yardımcı fonksiyonu, bu harici süreci temsil eden bir nesne oluşturur.

```go
    dateCmd := exec.Command("date")
```

`Output` metodu komutu çalıştırır, bitmesini bekler ve standart çıktısını toplar. Hata yoksa `dateOut`, tarih bilgisini içeren byte’ları tutar.

```go
    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> date")
    fmt.Println(string(dateOut))
```

---

## ⚠️ Hata Türlerini Ayırt Etme

`Command`’ın `Output` ve diğer metotları; komutu çalıştırırken bir sorun olursa (ör. yanlış path) `*exec.Error`, komut çalışıp sıfır olmayan bir dönüş koduyla (*return code*) çıkarsa `*exec.ExitError` döndürür.

```go
    _, err = exec.Command("date", "-x").Output()
    if err != nil {
        var execErr *exec.Error
        var exitErr *exec.ExitError
        switch {
        case errors.As(err, &execErr):
            fmt.Println("failed executing:", err)
        case errors.As(err, &exitErr):
            exitCode := exitErr.ExitCode()
            fmt.Println("command exit rc =", exitCode)
        default:
            panic(err)
        }
    }
```

---

## 🔁 stdin’e Veri Pipe Etme, stdout’tan Sonucu Alma

Şimdi, harici sürecin stdin’ine veri pipe ettiğimiz ve sonuçları stdout’tan topladığımız biraz daha kapsamlı bir örneğe bakacağız.

```go
    grepCmd := exec.Command("grep", "hello")
```

Burada giriş/çıkış pipe’larını açıkça alırız, süreci başlatırız, ona biraz girdi yazarız, çıkan çıktıyı okuruz ve son olarak sürecin çıkmasını bekleriz.

```go
    grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()
    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, _ := io.ReadAll(grepOut)
    grepCmd.Wait()
```

Yukarıdaki örnekte hata kontrollerini atladık; ancak hepsi için alışıldık `if err != nil` desenini kullanabilirsiniz. Ayrıca yalnızca `StdoutPipe` sonuçlarını topluyoruz, fakat `StderrPipe`’ı da aynı şekilde toplayabilirsiniz.

```go
    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))
```

---

## 🧾 Tek Bir Komut Satırı String’i Yerine Argüman Dizisi

Komut başlatırken, tek bir komut satırı string’i geçmek yerine, açıkça ayrılmış bir komut ve argüman dizisi sağlamamız gerektiğine dikkat edin. Eğer bir string ile tam bir komut çalıştırmak isterseniz, `bash`’in `-c` seçeneğini kullanabilirsiniz:

```go
    lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))
}
```

---

## 🧪 Örnek Çıktı

Başlatılan programlar, onları doğrudan komut satırından çalıştırmışız gibi aynı çıktıyı döndürür.

```bash
$ go run spawning-processes.go 
> date
Thu 05 May 2022 10:10:12 PM PDT
```

`date` komutunda `-x` bayrağı yoktur; bu yüzden bir hata mesajı ve sıfır olmayan bir dönüş koduyla çıkar.

```text
command exited with rc = 1
```

```bash
> grep hello
hello grep
> ls -a -l -h
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 spawning-processes.go
```

---

## ➡️ Sonraki Örnek

Sonraki örnek: **Exec'ing Processes**.

