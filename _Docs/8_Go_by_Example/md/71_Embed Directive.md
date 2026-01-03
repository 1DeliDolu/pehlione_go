
## 🧷 Go by Example: Embed Direktifi

`//go:embed`, derleme zamanında programların Go binary’sine rastgele dosya ve klasörleri dahil etmesine izin veren bir derleyici direktifidir. Embed direktifi hakkında daha fazlasını buradan okuyun.

### ▶️ Çalıştırma

```go
package main

// embed paketini içe aktarın; bu paketten dışa aktarılan (exported) herhangi bir tanımlayıcı kullanmıyorsanız,
// _ "embed" ile blank import da yapabilirsiniz.
import (
    "embed"
)

// embed direktifleri, Go kaynak dosyasını içeren dizine göre göreli (relative) yollar kabul eder.
// Bu direktif, dosyanın içeriğini hemen ardından gelen string değişkenine gömer.

//go:embed folder/single_file.txt
var fileString string

// Ya da dosyanın içeriğini bir []byte içine gömün.

//go:embed folder/single_file.txt
var fileByte []byte

// Wildcard kullanarak birden fazla dosyayı hatta klasörleri gömebiliriz.
// Bu, basit bir sanal dosya sistemini (virtual file system) uygulayan embed.FS türünde bir değişken kullanır.

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {
    // single_file.txt içeriğini yazdır.
    print(fileString)
    print(string(fileByte))

    // Gömülü klasörden bazı dosyaları al.
    content1, _ := folder.ReadFile("folder/file1.hash")
    print(string(content1))
    content2, _ := folder.ReadFile("folder/file2.hash")
    print(string(content2))
}
```

### 💻 CLI

Bu komutları örneği çalıştırmak için kullanın. (Not: Go playground sınırlaması nedeniyle bu örnek yalnızca yerel makinenizde çalıştırılabilir.)

```bash
$ mkdir -p folder
$ echo "hello go" > folder/single_file.txt
$ echo "123" > folder/file1.hash
$ echo "456" > folder/file2.hash
$ go run embed-directive.go
hello go
hello go
123
456
```

## ⏭️ Sonraki Örnek: Test ve Benchmarking

