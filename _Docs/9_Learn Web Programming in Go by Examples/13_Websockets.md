
## 🔌 WebSocket’ler

Bu örnek, Go’da *websocket*’lerle nasıl çalışılacağını gösterir. Gönderdiğimiz her şeyi geri **echo** eden basit bir sunucu oluşturacağız. Bunun için popüler **`gorilla/websocket`** kütüphanesini şu şekilde indirmemiz gerekir:

```bash
$ go get github.com/gorilla/websocket
```

Bundan sonra yazdığımız her uygulama bu kütüphaneyi kullanabilir.

```go
// websockets.go
package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func main() {
    http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
        conn, _ := upgrader.Upgrade(w, r, nil) // basitlik adına hata yok sayıldı

        for {
            // Tarayıcıdan mesaj oku
            msgType, msg, err := conn.ReadMessage()
            if err != nil {
                return
            }

            // Mesajı konsola yazdır
            fmt.Printf("%s gönderdi: %s\n", conn.RemoteAddr(), string(msg))

            // Mesajı tarayıcıya geri yaz
            if err = conn.WriteMessage(msgType, msg); err != nil {
                return
            }
        }
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "websockets.html")
    })

    http.ListenAndServe(":8080", nil)
}
```

```html
<!-- websockets.html -->
<input id="input" type="text" />
<button onclick="send()">Gönder</button>
<pre id="output"></pre>
<script>
    var input = document.getElementById("input");
    var output = document.getElementById("output");
    var socket = new WebSocket("ws://localhost:8080/echo");

    socket.onopen = function () {
        output.innerHTML += "Durum: Bağlandı\n";
    };

    socket.onmessage = function (e) {
        output.innerHTML += "Sunucu: " + e.data + "\n";
    };

    function send() {
        socket.send(input.value);
        input.value = "";
    }
</script>
```

```bash
$ go run websockets.go
[127.0.0.1]:53403 sent: Hello Go Web Examples, you're doing great!
```

```text
Send
Status: Connected
Server: Hello Go Web Examples, you're doing great!
```

