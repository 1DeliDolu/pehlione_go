# 🧬 Eğitim: Generics ile Başlangıç

İçindekiler
Ön Koşullar
Kodunuz için bir klasör oluşturun
Generic olmayan fonksiyonlar ekleyin
Birden çok türü işlemek için generic bir fonksiyon ekleyin
Generic fonksiyonu çağırırken tür argümanlarını kaldırın
Bir tür kısıtı (type constraint) tanımlayın
Sonuç
Tamamlanmış kod

Bu eğitim, Go’da  *generics* ’in temellerini tanıtır. *Generics* ile, çağıran kod tarafından sağlanan bir tür kümesindeki herhangi bir türle çalışacak şekilde yazılmış fonksiyonlar veya türler tanımlayıp kullanabilirsiniz.

Bu eğitimde, iki basit generic olmayan fonksiyon tanımlayacak, ardından aynı mantığı tek bir generic fonksiyonda toplayacaksınız.

Aşağıdaki bölümlerden ilerleyeceksiniz:

* Kodunuz için bir klasör oluşturun.
* Generic olmayan fonksiyonlar ekleyin.
* Birden çok türü işlemek için generic bir fonksiyon ekleyin.
* Generic fonksiyonu çağırırken tür argümanlarını kaldırın.
* Bir tür kısıtı tanımlayın.

> Not: Diğer eğitimler için Eğitimler’e bakın.
> Not: İsterseniz, programınızı düzenlemek ve çalıştırmak için “Go dev branch” modunda Go playground’u kullanabilirsiniz.

---

## ✅ Ön Koşullar¶

* Go 1.18 veya üzeri bir kurulum. Kurulum talimatları için Installing Go’ya bakın.
* Kodunuzu düzenlemek için bir araç. Sahip olduğunuz herhangi bir metin düzenleyici gayet uygundur.
* Bir komut terminali. Go, Linux ve Mac’te herhangi bir terminalle; Windows’ta ise PowerShell veya cmd ile iyi çalışır.

---

## 📁 Kodunuz için bir klasör oluşturun¶

Başlamak için, yazacağınız kod için bir klasör oluşturun.

Komut istemini açın ve home dizininize geçin.

**Linux veya Mac’te:**

```bash
$ cd
```

**Windows’ta:**

```text
C:\> cd %HOMEPATH%
```

Bu eğitimin geri kalanında, komut satırı göstergesi olarak `$` kullanılacaktır. Kullandığınız komutlar Windows’ta da çalışacaktır.

Komut isteminden, `generics` adlı bir dizin oluşturun.

```bash
$ mkdir generics
$ cd generics
```

Kodunuzu tutacak bir modül oluşturun.

Yeni kodunuzun modül yolunu vererek `go mod init` komutunu çalıştırın.

```bash
$ go mod init example/generics
go: creating new go.mod: module example/generics
```

> Not: Üretim kodunda, kendi ihtiyaçlarınıza daha uygun, daha spesifik bir modül yolu belirlersiniz. Daha fazlası için Managing dependencies’e bakın.

Sıradaki adımda, map’lerle çalışmak için basit bazı kodlar ekleyeceksiniz.

---

## 🧱 Generic olmayan fonksiyonlar ekleyin¶

Bu adımda, bir map’in değerlerini toplayıp toplamı döndüren iki fonksiyon ekleyeceksiniz.

İki fonksiyon tanımlıyorsunuz (tek bir fonksiyon yerine), çünkü iki farklı map türüyle çalışıyorsunuz: biri `int64` değerleri saklıyor, diğeri `float64` değerleri saklıyor.

### ✍️ Kodu yazın¶

Metin düzenleyicinizi kullanarak, `generics` dizininde `main.go` adlı bir dosya oluşturun. Go kodunuzu bu dosyaya yazacaksınız.

`main.go` dosyasında, dosyanın en üstüne aşağıdaki package bildirimini yapıştırın.

```go
package main
```

Bağımsız bir program (bir kütüphanenin aksine) her zaman `package main` içindedir.

Package bildiriminden sonra, aşağıdaki iki fonksiyon bildirimini yapıştırın.

```go
// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}
```

Bu kodda şunları yaparsınız:

* Bir map’in değerlerini toplayıp toplamı döndüren iki fonksiyon tanımlarsınız.
* `SumFloats`, `string` → `float64` değerlerinden oluşan bir map alır.
* `SumInts`, `string` → `int64` değerlerinden oluşan bir map alır.

`main.go` dosyasının en üstünde, package bildiriminden sonra, aşağıdaki `main` fonksiyonunu yapıştırarak iki map’i başlatın ve onları bir önceki adımda tanımladığınız fonksiyonları çağırırken argüman olarak kullanın.

```go
func main() {
    // Initialize a map for the integer values
    ints := map[string]int64{
        "first":  34,
        "second": 12,
    }

    // Initialize a map for the float values
    floats := map[string]float64{
        "first":  35.98,
        "second": 26.99,
    }

    fmt.Printf("Non-Generic Sums: %v and %v\n",
        SumInts(ints),
        SumFloats(floats))
}
```

Bu kodda şunları yaparsınız:

* Her biri iki girişe sahip olacak şekilde `float64` değerlerinden bir map ve `int64` değerlerinden bir map başlatırsınız.
* Her map’in değerleri toplamını bulmak için daha önce tanımladığınız iki fonksiyonu çağırırsınız.
* Sonucu yazdırırsınız.

`main.go` dosyasının en üstüne, package bildiriminden hemen sonra, az önce yazdığınız kodu desteklemek için ihtiyaç duyacağınız paketi import edin.

Kodun ilk satırları şöyle görünmelidir:

```go
package main

import "fmt"
```

`main.go` dosyasını kaydedin.

### ▶️ Kodu çalıştırın¶

`main.go` dosyasının bulunduğu dizinde komut satırından kodu çalıştırın.

```bash
$ go run .
Non-Generic Sums: 46 and 62.97
```

*Generics* ile burada iki fonksiyon yerine tek bir fonksiyon yazabilirsiniz. Sırada, `int` veya `float` değerleri içeren map’ler için tek bir generic fonksiyon ekleyeceksiniz.

---

## 🧩 Birden çok türü işlemek için generic bir fonksiyon ekleyin¶

Bu bölümde, `int` veya `float` değerleri içeren bir map alabilen tek bir generic fonksiyon ekleyeceksiniz; bu da az önce yazdığınız iki fonksiyonu etkin biçimde tek bir fonksiyonla değiştirecek.

Her iki türdeki değerleri desteklemek için, bu tek fonksiyonun hangi türleri desteklediğini bildirecek bir yol gerekir. Öte yandan çağıran kodun, bir tamsayı map’i mi yoksa bir float map’i mi ile çağırdığını belirtmesi için de bir yol gerekir.

Bunu desteklemek için, normal fonksiyon parametrelerine ek olarak  *type parameter* ’lar bildiren bir fonksiyon yazacaksınız. Bu  *type parameter* ’lar fonksiyonu generic hale getirir ve farklı türlerde argümanlarla çalışmasını sağlar. Fonksiyonu  *type argument* ’lar ve normal fonksiyon argümanlarıyla çağıracaksınız.

Her  *type parameter* ’ın bir  *type constraint* ’i vardır; bu, *type parameter* için bir tür *meta-type* gibi davranır. Her  *type constraint* , çağıran kodun ilgili *type parameter* için kullanabileceği izinli  *type argument* ’ları belirtir.

Bir  *type parameter* ’ın kısıtı tipik olarak bir tür kümesini temsil etse de, derleme zamanında *type parameter* tek bir türün yerini tutar — çağıran kodun *type argument* olarak sağladığı türün. Eğer  *type argument* ’ın türü ilgili *type parameter* kısıtı tarafından izinli değilse, kod derlenmez.

Unutmayın: Bir  *type parameter* , generic kodun üzerinde yaptığı tüm işlemleri desteklemelidir. Örneğin fonksiyonunuz, kısıtı sayısal türleri içeren bir *type parameter* üzerinde string işlemleri (ör. indeksleme) yapmaya çalışsaydı, kod derlenmezdi.

Birazdan yazacağınız kodda, tamsayı veya float türlerine izin veren bir kısıt kullanacaksınız.

### ✍️ Kodu yazın¶

Daha önce eklediğiniz iki fonksiyonun altına aşağıdaki generic fonksiyonu yapıştırın.

```go
// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
```

Bu kodda şunları yaparsınız:

* İki  *type parameter* ’a (köşeli parantez içindekiler) sahip bir `SumIntsOrFloats` fonksiyonu bildirirsiniz: `K` ve `V`; ayrıca  *type parameter* ’ları kullanan tek bir argüman (`m`), `map[K]V` türündedir. Fonksiyon `V` türünde bir değer döndürür.
* `K`  *type parameter* ’ı için `comparable`  *type constraint* ’ini belirtirsiniz. Bu, özellikle bu tür durumlar için amaçlanmış, Go’da önceden bildirilmiş bir kısıttır. Değerleri `==` ve `!=` karşılaştırma operatörlerinin operandı olabilen her türlü türe izin verir. Go, map anahtarlarının karşılaştırılabilir olmasını ister. Bu nedenle `K`’yi `comparable` olarak bildirmek, `K`’yi map değişkeninde anahtar olarak kullanabilmeniz için gereklidir. Ayrıca çağıran kodun map anahtarları için izinli bir tür kullanmasını garanti eder.
* `V`  *type parameter* ’ı için iki türün birleşimi olan bir kısıt belirtirsiniz: `int64` ve `float64`. `|` kullanımı iki türün birleşimini belirtir; yani bu kısıt her iki türe de izin verir. Çağıran kodda her iki tür de derleyici tarafından argüman olarak kabul edilir.
* `m` argümanının, daha önce *type parameter* olarak belirtilen `K` ve `V` türlerini kullanan `map[K]V` türünde olduğunu belirtirsiniz. `K` karşılaştırılabilir bir tür olduğu için `map[K]V`’nin geçerli bir map türü olduğunu biliyoruz. Eğer `K`’yi `comparable` olarak bildirmeseydik, derleyici `map[K]V` referansını reddederdi.

`main.go` dosyasında, mevcut kodunuzun altına aşağıdaki kodu yapıştırın.

```go
fmt.Printf("Generic Sums: %v and %v\n",
    SumIntsOrFloats[string, int64](ints),
    SumIntsOrFloats[string, float64](floats))
```

Bu kodda şunları yaparsınız:

* Az önce bildirdiğiniz generic fonksiyonu çağırır ve oluşturduğunuz map’lerin her birini geçirirsiniz.
* Fonksiyondaki  *type parameter* ’ların yerine hangi türlerin geçmesi gerektiğini açıkça belirtmek için  *type argument* ’lar (köşeli parantez içindeki tür adları) verirsiniz.
* Bir sonraki bölümde göreceğiniz gibi, çoğu zaman bu  *type argument* ’ları yazmayabilirsiniz. Go, çoğu durumda bunları kodunuzdan çıkarım yoluyla belirleyebilir.
* Fonksiyonun döndürdüğü toplamları yazdırırsınız.

### ▶️ Kodu çalıştırın¶

`main.go` dosyasının bulunduğu dizinde komut satırından kodu çalıştırın.

```bash
$ go run .
Non-Generic Sums: 46 and 62.97
Generic Sums: 46 and 62.97
```

Kodu çalıştırmak için, her çağrıda derleyici  *type parameter* ’ları, o çağrıda belirtilen somut türlerle değiştirdi.

Yazdığınız generic fonksiyonu çağırırken, derleyiciye fonksiyonun  *type parameter* ’larının yerine hangi türleri koyacağını söyleyen  *type argument* ’ları belirttiniz. Bir sonraki bölümde göreceğiniz gibi, çoğu durumda bu  *type argument* ’ları yazmayabilirsiniz; çünkü derleyici bunları çıkarım yoluyla belirleyebilir.

---

## 🧠 Generic fonksiyonu çağırırken tür argümanlarını kaldırın¶

Bu bölümde, generic fonksiyon çağrısının değiştirilmiş bir sürümünü ekleyecek ve çağıran kodu basitleştirmek için küçük bir değişiklik yapacaksınız. Bu durumda gerekli olmadıkları için  *type argument* ’ları kaldıracaksınız.

Go derleyicisinin, kullanmak istediğiniz türleri çıkarabildiği durumlarda çağıran kodda  *type argument* ’ları atlayabilirsiniz. Derleyici,  *type argument* ’ları fonksiyon argümanlarının türlerinden çıkarır.

Bunun her zaman mümkün olmadığını unutmayın. Örneğin, argümanı olmayan bir generic fonksiyonu çağırmanız gerekseydi, fonksiyon çağrısında  *type argument* ’ları yazmanız gerekirdi.

### ✍️ Kodu yazın¶

`main.go` dosyasında, mevcut kodunuzun altına aşağıdaki kodu yapıştırın.

```go
fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
    SumIntsOrFloats(ints),
    SumIntsOrFloats(floats))
```

Bu kodda şunları yaparsınız:

* *Type argument* ’ları yazmadan generic fonksiyonu çağırırsınız.

### ▶️ Kodu çalıştırın¶

`main.go` dosyasının bulunduğu dizinde komut satırından kodu çalıştırın.

```bash
$ go run .
Non-Generic Sums: 46 and 62.97
Generic Sums: 46 and 62.97
Generic Sums, type parameters inferred: 46 and 62.97
```

Sırada, tamsayılar ve float’ların birleşimini yeniden kullanılabilir bir *type constraint* içine alarak fonksiyonu daha da basitleştireceksiniz; böylece örneğin diğer kodlardan da yeniden kullanabilirsiniz.

---

## 🧾 Bir tür kısıtı tanımlayın¶

Bu son bölümde, daha önce tanımladığınız kısıtı kendi arayüzü (interface) içine taşıyacak ve böylece birden çok yerde yeniden kullanabileceksiniz. Kısıtları bu şekilde bildirmek, özellikle kısıt daha karmaşık olduğunda, kodu daha derli toplu hale getirmeye yardımcı olur.

Bir  *type constraint* ’i bir interface olarak bildirirsiniz. Kısıt, interface’i uygulayan herhangi bir türe izin verir. Örneğin, üç metoda sahip bir interface *type constraint* bildirir ve bunu bir generic fonksiyonda *type parameter* ile kullanırsanız, fonksiyonu çağırmak için kullanılan  *type argument* ’ların bu metodların hepsine sahip olması gerekir.

 *Kısıt interface* ’leri belirli türlere de referans verebilir; bu bölümde göreceksiniz.

### ✍️ Kodu yazın¶

`main` fonksiyonunun hemen üstünde, import ifadelerinin hemen ardından, bir *type constraint* bildirmek için aşağıdaki kodu yapıştırın.

```go
type Number interface {
    int64 | float64
}
```

Bu kodda şunları yaparsınız:

* *Type constraint* olarak kullanılacak `Number` interface türünü bildirirsiniz.
* Interface içinde `int64` ve `float64` birleşimini bildirirsiniz.
* Esasen birleşimi fonksiyon bildiriminden alıp yeni bir *type constraint* içine taşırsınız. Böylece bir  *type parameter* ’ı `int64` veya `float64` ile kısıtlamak istediğinizde, `int64 | float64` yazmak yerine bu `Number`  *type constraint* ’ini kullanabilirsiniz.

Daha önce eklediğiniz fonksiyonların altına aşağıdaki generic `SumNumbers` fonksiyonunu yapıştırın.

```go
// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
```

Bu kodda şunları yaparsınız:

* Daha önce bildirdiğiniz generic fonksiyonla aynı mantığa sahip bir generic fonksiyon tanımlarsınız; ancak *type constraint* olarak birleşim yerine yeni interface türünü kullanırsınız. Daha önce olduğu gibi, argüman ve dönüş türleri için  *type parameter* ’ları kullanırsınız.

`main.go` dosyasında, mevcut kodunuzun altına aşağıdaki kodu yapıştırın.

```go
fmt.Printf("Generic Sums with Constraint: %v and %v\n",
    SumNumbers(ints),
    SumNumbers(floats))
```

Bu kodda şunları yaparsınız:

* Her map ile `SumNumbers` fonksiyonunu çağırır ve değerlerin toplamını yazdırırsınız.
* Bir önceki bölümde olduğu gibi, generic fonksiyon çağrılarında  *type argument* ’ları (köşeli parantez içindeki tür adları) yazmazsınız. Go derleyicisi,  *type argument* ’ı diğer argümanlardan çıkarabilir.

### ▶️ Kodu çalıştırın¶

`main.go` dosyasının bulunduğu dizinde komut satırından kodu çalıştırın.

```bash
$ go run .
Non-Generic Sums: 46 and 62.97
Generic Sums: 46 and 62.97
Generic Sums, type parameters inferred: 46 and 62.97
Generic Sums with Constraint: 46 and 62.97
```

---

## 🎉 Sonuç¶

Harika! Go’da *generics* ile yeni tanıştınız.

Önerilen sonraki konular:

* Go Tour, Go temellerine adım adım harika bir giriş sunar.
* Effective Go ve How to write Go code içinde faydalı Go en iyi uygulamalarını bulabilirsiniz.

---

## ✅ Tamamlanmış kod¶

Bu programı Go playground’da çalıştırabilirsiniz. Playground’da sadece Run düğmesine tıklayın.

```go
package main

import "fmt"

type Number interface {
    int64 | float64
}

func main() {
    // Initialize a map for the integer values
    ints := map[string]int64{
        "first": 34,
        "second": 12,
    }

    // Initialize a map for the float values
    floats := map[string]float64{
        "first": 35.98,
        "second": 26.99,
    }

    fmt.Printf("Non-Generic Sums: %v and %v\n",
        SumInts(ints),
        SumFloats(floats))

    fmt.Printf("Generic Sums: %v and %v\n",
        SumIntsOrFloats[string, int64](ints),
        SumIntsOrFloats[string, float64](floats))

    fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
        SumIntsOrFloats(ints),
        SumIntsOrFloats(floats))

    fmt.Printf("Generic Sums with Constraint: %v and %v\n",
        SumNumbers(ints),
        SumNumbers(floats))
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}

// SumIntsOrFloats sums the values of map m. It supports both floats and integers
// as map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

// SumNumbers sums the values of map m. Its supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
```
