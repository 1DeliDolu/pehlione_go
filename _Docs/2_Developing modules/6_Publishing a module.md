## 📦 Bir Modül Yayınlama

Diğer geliştiricilerin kullanması için bir modülü kullanılabilir hâle getirmek istediğinizde, onu Go araçları tarafından görülebilir olacak şekilde yayınlarsınız. Modülü yayınladıktan sonra, paketlerini import eden geliştiriciler **`go get`** gibi komutları çalıştırarak modüle bağımlılığı çözümleyebilir.

Not: Bir modülün etiketlenmiş (tag’lenmiş) bir sürümünü yayınladıktan sonra değiştirmeyin. Modülü kullanan geliştiriciler için Go araçları, indirilen bir modülü ilk indirilen kopyaya karşı doğrular. İkisi farklıysa, Go araçları bir güvenlik hatası döndürür. Daha önce yayınlanmış bir sürümün kodunu değiştirmek yerine, yeni bir sürüm yayınlayın.

---

## 🔎 Ayrıca bakınız

* Modül geliştirmeye genel bir bakış için ***Modülleri geliştirme ve yayınlama*** konusuna bakın
* Yayınlamayı da içeren üst düzey modül geliştirme iş akışı için ***Modül yayınlama ve sürümleme iş akışı*** konusuna bakın.

---

## 🪜 Yayınlama Adımları

Bir modülü yayınlamak için aşağıdaki adımları izleyin.

### 🧭 1) Modül kök dizinine geçin

Bir komut istemi açın ve yerel depodaki modülünüzün kök dizinine geçin.

### 🧹 2) `go mod tidy` çalıştırın

Modülün artık gerekli olmayan ve zamanla birikmiş olabilecek bağımlılıklarını kaldırır.

```bash
$ go mod tidy
```

### 🧪 3) Son kez testleri çalıştırın: `go test ./...`

Her şeyin çalıştığından emin olmak için tekrar test edin.

Bu, Go test çerçevesini kullanarak yazdığınız birim testlerini çalıştırır.

```bash
$ go test ./...
ok      example.com/mymodule       0.015s
```

### 🏷️ 4) Yeni bir sürüm numarasıyla etiketi (tag) oluşturun

**`git tag`** komutunu kullanarak projeyi yeni bir sürüm numarasıyla etiketleyin.

Sürüm numarası olarak, bu yayındaki değişikliklerin niteliğini kullanıcılara işaret eden bir numara kullanın. Daha fazlası için ***Modül sürüm numaralandırma*** konusuna bakın.

```bash
$ git commit -m "mymodule: changes for v0.1.0"
$ git tag v0.1.0
```

### 📤 5) Yeni etiketi origin depoya itin

```bash
$ git push origin v0.1.0
```

### 🌐 6) `go list` ile modülü görünür hâle getirin

Yayınladığınız modül hakkında Go’nun modül indeksini güncellemesini tetiklemek için **`go list`** komutunu çalıştırarak modülü kullanılabilir hâle getirin.

Komutun önüne, **`GOPROXY`** ortam değişkenini bir Go proxy’sine ayarlayan bir ifade ekleyin. Bu, isteğinizin proxy’ye ulaşmasını sağlar.

```bash
$ GOPROXY=proxy.golang.org go list -m example.com/mymodule@v0.1.0
```

Modülünüzle ilgilenen geliştiriciler, ondan bir paket import eder ve diğer modüllerde olduğu gibi **`go get`** komutunu çalıştırır. En son sürümler için **`go get`** çalıştırabilirler veya aşağıdaki örnekte olduğu gibi belirli bir sürümü belirtebilirler:

```bash
$ go get example.com/mymodule@v0.1.0
```
