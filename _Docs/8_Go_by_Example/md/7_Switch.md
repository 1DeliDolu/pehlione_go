
## 🔀 Go by Example: Switch

`switch` deyimleri, çok sayıda dal (*branch*) üzerinden koşullu mantığı ifade eder.

---

## 🧾 Kaynak Kod

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // İşte temel bir switch.
    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    // Aynı case deyiminde birden fazla ifadeyi virgülle ayırabilirsiniz.
    // Bu örnekte isteğe bağlı default case’i de kullanıyoruz.
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

    // İfade (*expression*) olmadan switch, if/else mantığını ifade etmenin alternatif bir yoludur.
    // Burada ayrıca case ifadelerinin sabit (*constant*) olmak zorunda olmadığını da gösteriyoruz.
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    // Type switch, değerler yerine türleri karşılaştırır.
    // Bunu bir interface değerinin türünü keşfetmek için kullanabilirsiniz.
    // Bu örnekte, t değişkeni bulunduğu case cümlesine karşılık gelen türe sahip olacaktır.
    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}
```

---

## ▶️ Çalıştırma

```bash
$ go run switch.go 
Write 2 as two
It's a weekday
It's after noon
I'm a bool
I'm an int
Don't know type string
```

---

## 📌 Sonraki Örnek

**Next example:** *Arrays.*

