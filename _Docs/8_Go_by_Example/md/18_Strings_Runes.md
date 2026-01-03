
## 🔤 Go by Example: String’ler ve Rune’lar

Bir Go *string*’i, salt-okunur (*read-only*) bir *byte slice*’ıdır. Dil ve standart kütüphane, string’leri özel biçimde ele alır: UTF-8 ile kodlanmış metin konteynerleri olarak. Diğer dillerde string’ler “karakterler”den oluşur. Go’da ise karakter kavramına *rune* denir — bu, bir Unicode *code point*’i temsil eden bir tamsayıdır. Bu Go blog yazısı konuya iyi bir giriş niteliğindedir.

### ▶️ Çalıştır

```go
package main
import (
    "fmt"
    "unicode/utf8"
)
func main() {
```

## 📝 UTF-8 String Literal

`s`, Tay dilinde “hello” kelimesini temsil eden bir *literal* değere atanmış bir string’dir. Go string literal’leri UTF-8 kodlu metindir.

```go
    const s = "สวัสดี"
```

## 📏 Byte Uzunluğu

String’ler `[]byte` ile eşdeğer olduğundan, bu ifade string içinde saklanan ham byte’ların uzunluğunu üretir.

```go
    fmt.Println("Len:", len(s))
```

## 🧱 String Indexleme ve Ham Byte’lar

Bir string’i indexlemek, her index’teki ham byte değerlerini üretir. Bu döngü, `s` içindeki code point’leri oluşturan tüm byte’ların hex değerlerini üretir.

```go
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
    fmt.Println()
```

## 🔢 Rune Sayısı

Bir string’de kaç tane *rune* olduğunu saymak için `utf8` paketini kullanabiliriz. `RuneCountInString` çalışma zamanı, string’in boyutuna bağlıdır; çünkü her UTF-8 rune’u sırayla çözümlemek zorundadır. Bazı Tay karakterleri, birden fazla byte’a yayılabilen UTF-8 code point’leri ile temsil edilir; bu yüzden bu sayımın sonucu şaşırtıcı olabilir.

```go
    fmt.Println("Rune count:", utf8.RuneCountInString(s))
```

## 🔁 Range ile Rune Çözümleme

Bir `range` döngüsü string’leri özel olarak ele alır ve her rune’u, string içindeki ofseti ile birlikte çözer.

```go
    for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }
```

## 🧩 DecodeRuneInString ile Aynı İterasyon

`utf8.DecodeRuneInString` fonksiyonunu açıkça kullanarak da aynı iterasyonu yapabiliriz.

```go
    fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width
```

Bu, bir rune değerinin bir fonksiyona aktarılmasını gösterir.

```go
        examineRune(runeValue)
    }
}
func examineRune(r rune) {
```

## 🔎 Rune Literal ve Karşılaştırma

Tek tırnak içine alınmış değerler *rune literal*’leridir. Bir rune değerini bir rune literal ile doğrudan karşılaştırabiliriz.

```go
    if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}
```

## 💻 CLI Çıktısı

```bash
$ go run strings-and-runes.go
Len: 18
e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5 
Rune count: 6
U+0E2A 'ส' starts at 0
U+0E27 'ว' starts at 3
U+0E31 'ั' starts at 6
U+0E2A 'ส' starts at 9
U+0E14 'ด' starts at 12
U+0E35 'ี' starts at 15
Using DecodeRuneInString
U+0E2A 'ส' starts at 0
found so sua
U+0E27 'ว' starts at 3
U+0E31 'ั' starts at 6
U+0E2A 'ส' starts at 9
found so sua
U+0E14 'ด' starts at 12
U+0E35 'ี' starts at 15
```

## ➡️ Sonraki Örnek

Next example: Structs.

