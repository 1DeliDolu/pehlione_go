
## 🔁 Süreçleri Exec Etme

Önceki örnekte harici süreçleri (*external processes*) başlatmaya baktık. Bunu, çalışan bir Go sürecinin erişebileceği harici bir sürece ihtiyaç duyduğumuzda yaparız. Bazen ise mevcut Go sürecini tamamen başka bir (belki Go olmayan) süreçle değiştirmek isteriz. Bunu yapmak için, Go’nun klasik `exec` fonksiyonunun implementasyonunu kullanacağız.

---

## ▶️ Çalıştırma

```go
package main
```

```go
import (
    "os"
    "os/exec"
    "syscall"
)
```

```go
func main() {
```

---

## 🔎 Çalıştırılacak Binary’nin Yolunu Bulma

Örneğimizde `ls` komutunu `exec` edeceğiz. Go, çalıştırmak istediğimiz binary için mutlak bir yol (*absolute path*) ister; bu yüzden onu bulmak için `exec.LookPath` kullanacağız (muhtemelen `/bin/ls`).

```go
    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }
```

---

## 🧾 Argümanları Slice Olarak Verme

`exec`, argümanları slice biçiminde ister (tek bir büyük string yerine). `ls`’e birkaç yaygın argüman vereceğiz. İlk argümanın program adı olması gerektiğine dikkat edin.

```go
    args := []string{"ls", "-a", "-l", "-h"}
```

---

## 🌿 Ortam Değişkenlerini Sağlama

`exec`, kullanılacak bir ortam değişkenleri (*environment variables*) kümesine de ihtiyaç duyar. Burada sadece mevcut ortamımızı sağlıyoruz.

```go
    env := os.Environ()
```

---

## ⚙️ syscall.Exec Çağrısı

İşte asıl `syscall.Exec` çağrısı. Bu çağrı başarılı olursa, sürecimizin yürütülmesi burada sona erer ve `/bin/ls -a -l -h` süreciyle değiştirilir. Bir hata varsa bir dönüş değeri alırız.

```go
    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}
```

---

## 🧪 Örnek Çıktı

Programı çalıştırdığımızda, süreç `ls` ile değiştirilir.

```bash
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go
```

---

## 📝 Not

Go, klasik Unix `fork` fonksiyonunu sunmaz. Ancak genellikle bu bir sorun değildir; çünkü goroutine başlatmak, süreç başlatmak (*spawning processes*) ve süreçleri `exec` etmek, `fork` için çoğu kullanım senaryosunu kapsar.

---

## ➡️ Sonraki Örnek

Sonraki örnek: **Signals**.

