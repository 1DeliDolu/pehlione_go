
## 🔌 TCP Sunucusu

`net` paketi, TCP soket sunucularını kolayca oluşturmak için ihtiyaç duyduğumuz araçları sağlar.

---

## ▶️ Çalıştırma

```go
package main
```

```go
import (
    "bufio"
    "fmt"
    "log"
    "net"
    "strings"
)
```

```go
func main() {
```

---

## 🟦 Sunucuyu Başlatma

`net.Listen`, verilen ağ (*network* — TCP) ve adres (*address* — tüm arayüzlerde 8090 portu) üzerinde sunucuyu başlatır.

```go
    listener, err := net.Listen("tcp", ":8090")
    if err != nil {
        log.Fatal("Error listening:", err)
    }
```

Uygulama çıktığında portu serbest bırakmak için listener’ı kapatın.

```go
    defer listener.Close()
```

---

## 🔁 Bağlantıları Kabul Etme Döngüsü

Yeni istemci bağlantılarını kabul etmek için sonsuz döngüye girin.

```go
    for {
```

Bir bağlantı bekleyin.

```go
        conn, err := listener.Accept()
        if err != nil {
            log.Println("Error accepting conn:", err)
            continue
        }
```

Burada bağlantıyı ele almak için bir goroutine kullanıyoruz; böylece ana döngü daha fazla bağlantı kabul etmeye devam edebilir.

```go
        go handleConnection(conn)
    }
}
```

---

## 🧵 Tek Bir Bağlantıyı Ele Alma

`handleConnection`, tek bir istemci bağlantısını ele alır; istemciden bir satır metin okur ve bir yanıt döndürür.

```go
func handleConnection(conn net.Conn) {
```

İstemciyle etkileşimimiz bittiğinde kaynakları serbest bırakmak için bağlantıyı kapatırız.

```go
    defer conn.Close()
```

İstemciden (newline ile sonlanan) bir satır veri okumak için `bufio.NewReader` kullanın.

```go
    reader := bufio.NewReader(conn)
    message, err := reader.ReadString('\n')
    if err != nil {
        log.Printf("Read error: %v", err)
        return
    }
```

İki yönlü iletişimi göstermek için istemciye bir yanıt oluşturup geri gönderin.

```go
    ackMsg := strings.ToUpper(strings.TrimSpace(message))
    response := fmt.Sprintf("ACK: %s\n", ackMsg)
    _, err = conn.Write([]byte(response))
    if err != nil {
        log.Printf("Server write error: %v", err)
    }
}
```

---

## 🧪 Sunucuyu Çalıştırma

TCP sunucusunu arka planda çalıştırın.

```bash
$ go run tcp-server.go &
```

`netcat` kullanarak veri gönderin ve yanıtı yakalayın.

```bash
$ echo "Hello from netcat" | nc localhost 8090
ACK: HELLO FROM NETCAT
```

---

## ➡️ Sonraki Örnek

Sonraki örnek: **Context**.

