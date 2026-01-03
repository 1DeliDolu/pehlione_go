
## 🧩 Encoding

### 🏗️ JSON Değişimi ile Build

Gin, varsayılan JSON paketi olarak `encoding/json` kullanır; ancak başka tag’lerle build ederek bunu değiştirebilirsiniz.

---

## ⚙️ go-json

Terminal penceresi

```bash
go build -tags=go_json .
```

---

## ⚙️ jsoniter

Terminal penceresi

```bash
go build -tags=jsoniter .
```

---

## ⚙️ sonic *(CPU’nuzun avx komutunu desteklediğinden emin olmanız gerekir.)*

Terminal penceresi

```bash
$ go build -tags="sonic avx" .
```

