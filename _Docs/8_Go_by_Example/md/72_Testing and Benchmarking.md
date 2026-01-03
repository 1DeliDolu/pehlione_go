
## 🧪 Go by Example: Test ve Benchmarking

Unit test, prensipli Go programları yazmanın önemli bir parçasıdır. `testing` paketi unit test yazmak için ihtiyaç duyduğumuz araçları sağlar ve `go test` komutu testleri çalıştırır.

Gösterim amacıyla bu kod `package main` içindedir, ancak herhangi bir paket de olabilir. Test kodu tipik olarak test ettiği kodla aynı pakette bulunur.

### ▶️ Çalıştırma

```go
package main

import (
    "fmt"
    "testing"
)

// Bu basit tamsayı minimum (min) implementasyonunu test edeceğiz.
// Tipik olarak test ettiğimiz kod intutils.go gibi bir kaynak dosyada olurdu
// ve buna ait test dosyası da intutils_test.go olarak adlandırılırdı.
func IntMin(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// Bir test, adı Test ile başlayan bir fonksiyon yazarak oluşturulur.
func TestIntMinBasic(t *testing.T) {
    ans := IntMin(2, -2)
    if ans != -2 {
        // t.Error* test başarısızlıklarını raporlar ama testi çalıştırmaya devam eder.
        // t.Fatal* test başarısızlıklarını raporlar ve testi hemen durdurur.
        t.Errorf("IntMin(2, -2) = %d; want -2", ans)
    }
}

// Test yazmak tekrar edici olabilir, bu yüzden “table-driven” (tablo güdümlü) stil kullanmak idiomatiktir.
// Test girdileri ve beklenen çıktılar bir tabloda listelenir ve tek bir döngü bu tabloyu gezerek test mantığını uygular.
func TestIntMinTableDriven(t *testing.T) {
    var tests = []struct {
        a, b int
        want int
    }{
        {0, 1, 0},
        {1, 0, 0},
        {2, -2, -2},
        {0, -1, -1},
        {-1, 0, -1},
    }

    // t.Run, her tablo girdisi için birer tane olacak şekilde “subtest” çalıştırmayı sağlar.
    // Bunlar go test -v çalıştırıldığında ayrı ayrı gösterilir.
    for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
            ans := IntMin(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}

// Benchmark testleri tipik olarak _test.go dosyalarında bulunur ve Benchmark ile başlayan adlara sahiptir.
// Benchmark’ın çalışması için gereken ama ölçülmemesi gereken kod, bu döngüden önce yer alır.
func BenchmarkIntMin(b *testing.B) {
    for b.Loop() {
        // Benchmark çalıştırıcısı, tek bir iterasyonun çalışma süresine makul bir tahmin yapmak için
        // bu döngü gövdesini otomatik olarak çok kez çalıştırır.
        IntMin(1, 2)
    }
}
```

### ✅ Mevcut Projedeki Tüm Testleri Verbose Modda Çalıştırma

```bash
$ go test -v
== RUN   TestIntMinBasic
--- PASS: TestIntMinBasic (0.00s)
=== RUN   TestIntMinTableDriven
=== RUN   TestIntMinTableDriven/0,1
=== RUN   TestIntMinTableDriven/1,0
=== RUN   TestIntMinTableDriven/2,-2
=== RUN   TestIntMinTableDriven/0,-1
=== RUN   TestIntMinTableDriven/-1,0
--- PASS: TestIntMinTableDriven (0.00s)
    --- PASS: TestIntMinTableDriven/0,1 (0.00s)
    --- PASS: TestIntMinTableDriven/1,0 (0.00s)
    --- PASS: TestIntMinTableDriven/2,-2 (0.00s)
    --- PASS: TestIntMinTableDriven/0,-1 (0.00s)
    --- PASS: TestIntMinTableDriven/-1,0 (0.00s)
PASS
ok      examples/testing-and-benchmarking    0.023s
```

### 📏 Mevcut Projedeki Tüm Benchmark’ları Çalıştırma

Tüm testler, benchmark’lardan önce çalıştırılır. `-bench` bayrağı, benchmark fonksiyon adlarını bir regexp ile filtreler.

```bash
$ go test -bench=.
goos: darwin
goarch: arm64
pkg: examples/testing
BenchmarkIntMin-8 1000000000 0.3136 ns/op
PASS
ok      examples/testing-and-benchmarking    0.351s
```

## ⏭️ Sonraki Örnek: Komut Satırı Argümanları

