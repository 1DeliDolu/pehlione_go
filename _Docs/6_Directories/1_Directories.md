
# 📚 Go Standart Kütüphane Paketleri 

## 🗂️ `archive`

## 📦 `tar`

`tar` paketi, *tar* arşivlerine erişimi uygular.

## 📦 `zip`

`zip` paketi, ZIP arşivlerini okuma ve yazma için destek sağlar.

## 📦 `bufio`

`bufio` paketi, arabellekli G/Ç uygular. Bir `io.Reader` veya `io.Writer` nesnesini sararak, arayüzü yine uygulayan fakat arabellekleme ve metinsel G/Ç için bazı yardımcılar sağlayan başka bir nesne (`Reader` veya `Writer`) oluşturur.

## 📦 `builtin`

`builtin` paketi, Go’nun önceden bildirilmiş tanımlayıcıları için dokümantasyon sağlar.

## 📦 `bytes`

`bytes` paketi, byte slice’larının işlenmesi için fonksiyonlar uygular.

## 📦 `cmp`

`cmp` paketi, sıralı değerleri karşılaştırmayla ilgili türler ve fonksiyonlar sağlar.

## 🗂️ `compress`

## 📦 `bzip2`

`bzip2` paketi, *bzip2* açmayı (*decompression*) uygular.

## 📦 `flate`

`flate` paketi, RFC 1951’de tanımlanan DEFLATE sıkıştırılmış veri biçimini uygular.

## 📦 `gzip`

`gzip` paketi, RFC 1952’de belirtildiği gibi *gzip* biçimli sıkıştırılmış dosyaları okuma ve yazmayı uygular.

## 📦 `lzw`

`lzw` paketi, T. A. Welch’in “A Technique for High-Performance Data Compression”, *Computer*, 17(6) (June 1984), ss. 8–19 çalışmasında tanımlanan Lempel-Ziv-Welch sıkıştırılmış veri biçimini uygular.

## 📦 `zlib`

`zlib` paketi, RFC 1950’de belirtildiği gibi *zlib* biçimli sıkıştırılmış verileri okuma ve yazmayı uygular.

## 🗂️ `container`

## 📦 `heap`

`heap` paketi, `heap.Interface` uygulayan herhangi bir tür için heap işlemleri sağlar.

## 📦 `list`

`list` paketi, çift bağlı bir liste uygular.

## 📦 `ring`

`ring` paketi, dairesel listeler üzerinde işlemler uygular.

## 📦 `context`

`context` paketi, API sınırları boyunca ve süreçler arasında son tarihleri (*deadline*), iptal sinyallerini ve diğer istek-kapsamlı değerleri taşıyan `Context` türünü tanımlar.

## 📦 `crypto`

`crypto` paketi, yaygın kriptografik sabitleri toplar.

## 📦 `aes`

`aes` paketi, ABD Federal Information Processing Standards Publication 197’de tanımlandığı gibi AES şifrelemesini (eski adıyla Rijndael) uygular.

## 📦 `cipher`

`cipher` paketi, düşük seviyeli blok şifre uygulamalarının etrafına sarılabilen standart blok şifre kiplerini (*modes*) uygular.

## 📦 `des`

`des` paketi, ABD Federal Information Processing Standards Publication 46-3’te tanımlandığı gibi Data Encryption Standard (DES) ve Triple Data Encryption Algorithm (TDEA) algoritmalarını uygular.

## 📦 `dsa`

`dsa` paketi, FIPS 186-3’te tanımlandığı gibi Digital Signature Algorithm (DSA) algoritmasını uygular.

## 📦 `ecdh`

`ecdh` paketi, NIST eğrileri ve Curve25519 üzerinde Elliptic Curve Diffie-Hellman uygular.

## 📦 `ecdsa`

`ecdsa` paketi, [FIPS 186-5]’te tanımlandığı gibi Elliptic Curve Digital Signature Algorithm (ECDSA) algoritmasını uygular.

## 📦 `ed25519`

`ed25519` paketi, Ed25519 imza algoritmasını uygular.

## 📦 `elliptic`

`elliptic` paketi, asal alanlar üzerinde standart NIST P-224, P-256, P-384 ve P-521 eliptik eğrilerini uygular.

## 🗂️ `fips140`

## 📦 `hkdf`

`hkdf` paketi, RFC 5869’da tanımlanan HMAC tabanlı Extract-and-Expand Key Derivation Function (HKDF) fonksiyonunu uygular.

## 📦 `hmac`

`hmac` paketi, ABD Federal Information Processing Standards Publication 198’de tanımlandığı gibi Keyed-Hash Message Authentication Code (HMAC) algoritmasını uygular.

## 📦 `md5`

`md5` paketi, RFC 1321’de tanımlandığı gibi MD5 hash algoritmasını uygular.

## 📦 `mlkem`

`mlkem` paketi, [NIST FIPS 203]’te belirtildiği gibi kuantuma dayanıklı anahtar kapsülleme yöntemi ML-KEM’i (eski adıyla Kyber) uygular.

## 📦 `pbkdf2`

`pbkdf2` paketi, RFC 8018’de (PKCS #5 v2.1) tanımlanan PBKDF2 anahtar türetme fonksiyonunu uygular.

## 📦 `rand`

`rand` paketi, kriptografik olarak güvenli bir rastgele sayı üreteci uygular.

## 📦 `rc4`

`rc4` paketi, Bruce Schneier’in *Applied Cryptography* eserinde tanımlandığı gibi RC4 şifrelemesini uygular.

## 📦 `rsa`

`rsa` paketi, PKCS #1 ve RFC 8017’de belirtildiği gibi RSA şifrelemesini uygular.

## 📦 `sha1`

`sha1` paketi, RFC 3174’te tanımlandığı gibi SHA-1 hash algoritmasını uygular.

## 📦 `sha256`

`sha256` paketi, FIPS 180-4’te tanımlandığı gibi SHA224 ve SHA256 hash algoritmalarını uygular.

## 📦 `sha3`

`sha3` paketi, FIPS 202’de tanımlanan SHA-3 hash algoritmalarını ve SHAKE genişletilebilir çıktı (*extendable output*) fonksiyonlarını uygular.

## 📦 `sha512`

`sha512` paketi, FIPS 180-4’te tanımlandığı gibi SHA-384, SHA-512, SHA-512/224 ve SHA-512/256 hash algoritmalarını uygular.

## 📦 `subtle`

`subtle` paketi, kriptografik kodda sıklıkla yararlı olan, ancak doğru kullanımı dikkatli düşünmeyi gerektiren fonksiyonlar uygular.

## 📦 `tls`

`tls` paketi, RFC 5246’da belirtildiği gibi TLS 1.2’yi ve RFC 8446’da belirtildiği gibi TLS 1.3’ü kısmen uygular.

## 📦 `x509`

`x509` paketi, X.509 standardının bir alt kümesini uygular.

## 📦 `x509/pkix`

`pkix` paketi, X.509 sertifikalarının ASN.1 ayrıştırma ve serileştirmesinde kullanılan düşük seviyeli paylaşılan yapıları (CRL ve OCSP dâhil) içerir.

## 🗂️ `database`

## 📦 `sql`

`sql` paketi, SQL (veya SQL-benzeri) veritabanları etrafında genel bir arayüz sağlar.

## 📦 `sql/driver`

`driver` paketi, `sql` paketi tarafından kullanılan veritabanı sürücülerinin uygulaması gereken arayüzleri tanımlar.

## 🗂️ `debug`

## 📦 `buildinfo`

`buildinfo` paketi, bir Go ikilisinin nasıl derlendiğine ilişkin Go binary’sine gömülü bilgilere erişim sağlar.

## 📦 `dwarf`

`dwarf` paketi, [http://dwarfstd.org/doc/dwarf-2.0.0.pdf](http://dwarfstd.org/doc/dwarf-2.0.0.pdf) adresindeki DWARF 2.0 Standardında tanımlandığı gibi, yürütülebilir dosyalardan yüklenen DWARF hata ayıklama bilgisine erişim sağlar.

## 📦 `elf`

`elf` paketi, ELF nesne dosyalarına erişimi uygular.

## 📦 `gosym`

`gosym` paketi, `gc` derleyicileri tarafından üretilen Go ikililerinin içine gömülen Go sembol ve satır numarası tablolarına erişimi uygular.

## 📦 `macho`

`macho` paketi, Mach-O nesne dosyalarına erişimi uygular.

## 📦 `pe`

`pe` paketi, PE (Microsoft Windows Portable Executable) dosyalarına erişimi uygular.

## 📦 `plan9obj`

`plan9obj` paketi, Plan 9 `a.out` nesne dosyalarına erişimi uygular.

## 📦 `embed`

`embed` paketi, çalışan Go programına gömülü dosyalara erişim sağlar.

## 📦 `encoding`

`encoding` paketi, byte seviyesinde ve metinsel temsiller arasında veri dönüştüren diğer paketlerin paylaştığı arayüzleri tanımlar.

## 📦 `ascii85`

`ascii85` paketi, `btoa` aracında ve Adobe’nin PostScript ve PDF doküman biçimlerinde kullanılan *ascii85* kodlamasını uygular.

## 📦 `asn1`

`asn1` paketi, ITU-T Rec X.690’da tanımlandığı gibi DER-kodlu ASN.1 veri yapılarının ayrıştırılmasını uygular.

## 📦 `base32`

`base32` paketi, RFC 4648’de belirtildiği gibi *base32* kodlamasını uygular.

## 📦 `base64`

`base64` paketi, RFC 4648’de belirtildiği gibi *base64* kodlamasını uygular.

## 📦 `binary`

`binary` paketi, sayılar ile byte dizileri arasında basit çeviri ile `varint` kodlama ve çözmeyi uygular.

## 📦 `csv`

`csv` paketi, virgülle ayrılmış değerler (CSV) dosyalarını okur ve yazar.

## 📦 `gob`

`gob` paketi, `gob` akışlarını yönetir — `Encoder` (gönderici) ve `Decoder` (alıcı) arasında değiş tokuş edilen ikili değerler.

## 📦 `hex`

`hex` paketi, onaltılık (*hexadecimal*) kodlama ve çözmeyi uygular.

## 📦 `json`

`json` paketi, RFC 7159’da tanımlandığı gibi JSON kodlama ve çözmeyi uygular.

## 📦 `json/jsontext`

`jsontext` paketi, RFC 4627, RFC 7159, RFC 7493, RFC 8259 ve RFC 8785’te belirtildiği gibi JSON’un sözdizimsel (*syntactic*) işlenmesini uygular.

## 📦 `json/v2`

`json` paketi, RFC 8259’da belirtildiği gibi JSON’un anlamsal (*semantic*) işlenmesini uygular.

## 📦 `pem`

`pem` paketi, Privacy Enhanced Mail’de ortaya çıkan PEM veri kodlamasını uygular.

## 📦 `xml`

`xml` paketi, XML ad alanlarını (*name spaces*) anlayan basit bir XML 1.0 ayrıştırıcısı uygular.

## 📦 `errors`

`errors` paketi, hataları (*errors*) işlemek için fonksiyonlar sağlar.

## 📦 `expvar`

`expvar` paketi, sunuculardaki işlem sayaçları gibi genel değişkenler için standartlaştırılmış bir arayüz sağlar.

## 📦 `flag`

`flag` paketi, komut satırı bayraklarını (*flags*) ayrıştırmayı uygular.

## 📦 `fmt`

`fmt` paketi, C’nin `printf` ve `scanf` fonksiyonlarına benzer işlevlere sahip fonksiyonlarla biçimlendirilmiş G/Ç uygular.

## 🗂️ `go`

## 📦 `ast`

`ast` paketi, Go paketleri için sözdizim ağaçlarını (*syntax trees*) temsil etmekte kullanılan türleri bildirir.

## 📦 `build`

`build` paketi, Go paketleri hakkında bilgi toplar.

## 📦 `build/constraint`

`constraint` paketi, build constraint satırlarının ayrıştırılmasını ve değerlendirilmesini uygular.

## 📦 `constant`

`constant` paketi, tipsiz Go sabitlerini temsil eden `Value` değerlerini ve bunlara karşılık gelen işlemleri uygular.

## 📦 `doc`

`doc` paketi, Go AST’den kaynak kod dokümantasyonunu çıkarır.

## 📦 `doc/comment`

`comment` paketi, doc comment’lerin (dokümantasyon yorumları) ayrıştırılmasını ve yeniden biçimlendirilmesini uygular.

## 📦 `format`

`format` paketi, Go kaynak kodu için standart biçimlendirmeyi uygular.

## 📦 `importer`

`importer` paketi, *export data importer*’lara erişim sağlar.

## 📦 `parser`

`parser` paketi, Go kaynak dosyaları için bir ayrıştırıcı uygular.

## 📦 `printer`

`printer` paketi, AST düğümlerini yazdırmayı (*printing*) uygular.

## 📦 `scanner`

`scanner` paketi, Go kaynak metni için bir tarayıcı (*scanner*) uygular.

## 📦 `token`

`token` paketi, sözlüksel token’ları temsil eden sabitleri ve temel işlemleri (yazdırma, kestirimler) tanımlar.

## 📦 `types`

`types` paketi, veri türlerini bildirir ve Go paketlerinin tür denetimi (*type-checking*) için algoritmaları uygular.

## 📦 `version`

`version` paketi, [Go sürümleri] ve [Go toolchain adlandırma sözdizimi] ile ilgili işlemler sağlar: `"go1.20"`, `"go1.21.0"`, `"go1.22rc2"`, `"go1.23.4-bigcorp"` gibi string’ler.

## 📦 `hash`

`hash` paketi, hash fonksiyonları için arayüzler sağlar.

## 📦 `adler32`

`adler32` paketi, Adler-32 sağlama toplamını (*checksum*) uygular.

## 📦 `crc32`

`crc32` paketi, 32-bit döngüsel artıklık denetimini (*CRC-32 checksum*) uygular.

## 📦 `crc64`

`crc64` paketi, 64-bit döngüsel artıklık denetimini (*CRC-64 checksum*) uygular.

## 📦 `fnv`

`fnv` paketi, Glenn Fowler, Landon Curt Noll ve Phong Vo tarafından oluşturulan kriptografik olmayan FNV-1 ve FNV-1a hash fonksiyonlarını uygular.

## 📦 `maphash`

`maphash` paketi, byte dizileri ve karşılaştırılabilir değerler üzerinde hash fonksiyonları sağlar.

## 📦 `html`

`html` paketi, HTML metnini escape ve unescape etmek için fonksiyonlar sağlar.

## 📦 `template`

`template` paketi (`html/template`), kod enjeksiyonuna karşı güvenli HTML çıktısı üretmek için veri güdümlü (*data-driven*) şablonlar uygular.

## 📦 `image`

`image` paketi, temel bir 2B (*2-D*) görüntü kütüphanesi uygular.

## 📦 `color`

`color` paketi, temel bir renk kütüphanesi uygular.

## 📦 `color/palette`

`palette` paketi, standart renk paletleri sağlar.

## 📦 `draw`

`draw` paketi, görüntü birleştirmeyi (*image composition*) uygular.

## 📦 `gif`

`gif` paketi, GIF çözücü (*decoder*) ve kodlayıcı (*encoder*) uygular.

## 📦 `jpeg`

`jpeg` paketi, JPEG çözücü ve kodlayıcı uygular.

## 📦 `png`

`png` paketi, PNG çözücü ve kodlayıcı uygular.

## 🗂️ `index`

## 📦 `suffixarray`

`suffixarray` paketi, bellek içi (*in-memory*) suffix array kullanarak logaritmik zamanda alt string araması uygular.

## 📦 `io`

`io` paketi, G/Ç ilkel(ler)i (*I/O primitives*) için temel arayüzler sağlar.

## 📦 `fs`

`fs` paketi, bir dosya sistemi için temel arayüzleri tanımlar.

## 📦 `ioutil`

`ioutil` paketi, bazı G/Ç yardımcı fonksiyonlarını uygular.

## 📦 `iter`

`iter` paketi, diziler (*sequences*) üzerinde iterator’larla ilgili temel tanımlar ve işlemler sağlar.

## 📦 `log`

`log` paketi, basit bir loglama paketi uygular.

## 📦 `slog`

`slog` paketi, yapılandırılmış loglama (*structured logging*) uygular; log kayıtları, anahtar-değer çiftleri olarak ifade edilen bir mesaj, önem derecesi (*severity*) ve diğer nitelikleri içerir.

## 📦 `syslog`

`syslog` paketi, sistem log servisine basit bir arayüz sağlar.

## 📦 `maps`

`maps` paketi, herhangi bir türdeki map’lerle yararlı çeşitli fonksiyonlar tanımlar.

## 📦 `math`

`math` paketi, temel sabitleri ve matematiksel fonksiyonları sağlar.

## 📦 `big`

`big` paketi, keyfî hassasiyetli aritmetiği (*arbitrary-precision arithmetic*, büyük sayılar) uygular.

## 📦 `bits`

`bits` paketi, önceden bildirilmiş unsigned tamsayı türleri için bit sayma ve bit manipülasyonu uygular.

## 📦 `cmplx`

`cmplx` paketi, kompleks sayılar için temel sabitleri ve matematiksel fonksiyonları sağlar.

## 📦 `rand`

`rand` paketi, benzetim (*simulation*) gibi görevler için uygun sözde-rastgele sayı üreteçleri uygular; ancak güvenlik-hassas işlerde kullanılmamalıdır.

## 📦 `rand/v2`

`rand` paketi, benzetim (*simulation*) gibi görevler için uygun sözde-rastgele sayı üreteçleri uygular; ancak güvenlik-hassas işlerde kullanılmamalıdır.

## 📦 `mime`

`mime` paketi, MIME belirtiminin (*spec*) bazı kısımlarını uygular.

## 📦 `multipart`

`multipart` paketi, RFC 2046’da tanımlandığı gibi MIME multipart ayrıştırmayı uygular.

## 📦 `quotedprintable`

`quotedprintable` paketi, RFC 2045’te belirtildiği gibi quoted-printable kodlamasını uygular.

## 📦 `net`

`net` paketi, TCP/IP, UDP, alan adı çözümleme ve Unix domain socket’ler dâhil ağ G/Ç için taşınabilir bir arayüz sağlar.

## 📦 `http`

`http` paketi, HTTP istemci ve sunucu uygulamaları sağlar.

## 📦 `http/cgi`

`cgi` paketi, RFC 3875’te belirtildiği gibi CGI (Common Gateway Interface) uygular.

## 📦 `http/cookiejar`

`cookiejar` paketi, bellek içi RFC 6265-uyumlu bir `http.CookieJar` uygular.

## 📦 `http/fcgi`

`fcgi` paketi, FastCGI protokolünü uygular.

## 📦 `http/httptest`

`httptest` paketi, HTTP testleri için yardımcılar sağlar.

## 📦 `http/httptrace`

`httptrace` paketi, HTTP istemci istekleri içindeki olayları izlemek (*trace*) için mekanizmalar sağlar.

## 📦 `http/httputil`

`httputil` paketi, `net/http` içindekileri tamamlayan HTTP yardımcı fonksiyonları sağlar.

## 📦 `http/pprof`

`pprof` paketi, `pprof` görselleştirme aracının beklediği biçimde çalışma zamanı profil verisini kendi HTTP sunucusu üzerinden sunar.

## 📦 `mail`

`mail` paketi, e-posta iletilerini ayrıştırmayı uygular.

## 📦 `netip`

`netip` paketi, küçük bir değer türü (*small value type*) olan bir IP adresi türü tanımlar.

## 📦 `rpc`

`rpc` paketi, bir nesnenin dışa aktarılan (*exported*) metotlarına, ağ veya diğer G/Ç bağlantıları üzerinden erişim sağlar.

## 📦 `rpc/jsonrpc`

`jsonrpc` paketi, `rpc` paketi için JSON-RPC 1.0 `ClientCodec` ve `ServerCodec` uygular.

## 📦 `smtp`

`smtp` paketi, RFC 5321’de tanımlandığı gibi Simple Mail Transfer Protocol (SMTP) uygular.

## 📦 `textproto`

`textproto` paketi, HTTP, NNTP ve SMTP tarzı metin tabanlı istek/yanıt protokolleri için genel destek sağlar.

## 📦 `url`

`url` paketi, URL’leri ayrıştırır ve query escape işlemlerini uygular.

## 📦 `os`

`os` paketi, işletim sistemi işlevselliğine platform-bağımsız bir arayüz sağlar.

## 📦 `exec`

`exec` paketi, harici komutları çalıştırır.

## 📦 `signal`

`signal` paketi, gelen sinyallere erişim sağlar.

## 📦 `user`

`user` paketi, kullanıcı hesaplarını ada veya kimliğe göre sorgulamaya izin verir.

## 📦 `path`

`path` paketi, `/` ile ayrılmış yol (*slash-separated paths*) manipülasyonu için yardımcı rutinler uygular.

## 📦 `filepath`

`filepath` paketi, dosya adlarını hedef işletim sistemi tarafından tanımlanan yollarla uyumlu şekilde manipüle eden yardımcı rutinler uygular.

## 📦 `plugin`

`plugin` paketi, Go eklentilerini (*plugins*) yüklemeyi ve sembol çözümlemeyi (*symbol resolution*) uygular.

## 📦 `reflect`

`reflect` paketi, çalışma zamanı yansımasını (*run-time reflection*) uygular; bir programın keyfî türlere sahip nesneleri manipüle etmesine izin verir.

## 📦 `regexp`

`regexp` paketi, düzenli ifade (*regular expression*) aramasını uygular.

## 📦 `syntax`

`syntax` paketi, düzenli ifadeleri ayrıştırma ağaçlarına (*parse trees*) ayrıştırır ve ayrıştırma ağaçlarını programlara derler.

## 📦 `runtime`

`runtime` paketi, goroutine’leri kontrol etmek gibi Go’nun çalışma zamanı sistemiyle etkileşen işlemleri içerir.

## 📦 `cgo`

`cgo` paketi, `cgo` aracı tarafından üretilen kod için çalışma zamanı desteği içerir.

## 📦 `coverage`

`coverage` paketi, `os.Exit` ile sonlanmayan uzun süre çalışan veya sunucu programlarında çalışma zamanında coverage profili yazmak için API’ler içerir.

## 📦 `debug`

`debug` paketi, programların çalışırken kendilerini hata ayıklamasına yardımcı olan imkânlar içerir.

## 📦 `metrics`

`metrics` paketi, Go çalışma zamanı tarafından dışa aktarılan uygulamaya-özgü metriklere erişmek için kararlı bir arayüz sağlar.

## 📦 `pprof`

`pprof` paketi, çalışma zamanı profil verisini yazar.

## 📦 `race`

`race` paketi, veri yarışı (*data race*) tespit mantığını içerir.

## 📦 `trace`

`trace` paketi, Go yürütüm izleyicisi (*execution tracer*) için iz (*traces*) üretmeye yönelik imkânlar içerir.

## 📦 `slices`

`slices` paketi, herhangi bir türdeki slice’larla yararlı çeşitli fonksiyonlar tanımlar.

## 📦 `sort`

`sort` paketi, slice’ları ve kullanıcı tanımlı koleksiyonları sıralamak için temel yapı taşları sağlar.

## 📦 `strconv`

`strconv` paketi, temel veri türlerinin string temsillerine ve bu temsillerden dönüşümleri uygular.

## 📦 `strings`

`strings` paketi, UTF-8 kodlu string’leri manipüle etmek için basit fonksiyonlar uygular.

## 📦 `structs`

`structs` paketi, bir struct’ın özelliklerini değiştirmek için struct alanı olarak kullanılabilecek işaretleyici türler (*marker types*) tanımlar.

## 📦 `sync`

`sync` paketi, karşılıklı dışlama kilitleri (*mutual exclusion locks*) gibi temel senkronizasyon ilkel(ler)ini sağlar.

## 📦 `atomic`

`atomic` paketi, senkronizasyon algoritmalarında yararlı düşük seviyeli atomik bellek ilkel(ler)ini sağlar.

## 📦 `syscall`

`syscall` paketi, düşük seviyeli işletim sistemi ilkel(ler)i için bir arayüz içerir.

## 📦 `js`

`js` paketi, `js/wasm` mimarisi kullanılırken WebAssembly ana ortamına (*host environment*) erişim sağlar.

## 📦 `testing`

`testing` paketi, Go paketlerinin otomatik testini destekler.

## 📦 `fstest`

`fstest` paketi, dosya sistemi uygulamalarını ve dosya sistemi kullanıcılarını test etmek için destek sağlar.

## 📦 `iotest`

`iotest` paketi, çoğunlukla test amaçlı yararlı `Reader` ve `Writer` uygulamaları sağlar.

## 📦 `quick`

`quick` paketi, black box testleri için yardımcı fonksiyonlar uygular.

## 📦 `slogtest`

`slogtest` paketi, `log/slog.Handler` uygulamalarını test etmek için destek sağlar.

## 📦 `synctest`

`synctest` paketi, eşzamanlı (*concurrent*) kodu test etmek için destek sağlar.

## 🗂️ `text`

## 📦 `scanner`

`scanner` paketi, UTF-8 kodlu metin için bir tarayıcı (*scanner*) ve tokenleştirici sağlar.

## 📦 `tabwriter`

`tabwriter` paketi, girdideki tab ile ayrılmış sütunları düzgün hizalanmış metne çeviren bir yazma filtresi (`tabwriter.Writer`) uygular.

## 📦 `template`

`template` paketi, metinsel çıktı üretmek için veri güdümlü (*data-driven*) şablonlar uygular.

## 📦 `template/parse`

`parse` paketi, `text/template` ve `html/template` tarafından tanımlandığı gibi şablonlar için ayrıştırma ağaçlarını (*parse trees*) inşa eder.

## 📦 `time`

`time` paketi, zamanı ölçmek ve görüntülemek için işlevsellik sağlar.

## 📦 `tzdata`

`tzdata` paketi, timezone veritabanının gömülü bir kopyasını sağlar.

## 📦 `unicode`

`unicode` paketi, Unicode kod noktalarının bazı özelliklerini sınamak için veri ve fonksiyonlar sağlar.

## 📦 `utf16`

`utf16` paketi, UTF-16 dizilerinin kodlanmasını ve çözülmesini uygular.

## 📦 `utf8`

`utf8` paketi, UTF-8 kodlu metin için fonksiyonlar ve sabitler sağlar.

## 📦 `unique`

`unique` paketi, karşılaştırılabilir değerleri kanonikleştirmek (“interning”) için imkânlar sağlar.

## 📦 `unsafe`

`unsafe` paketi, Go programlarının tür güvenliğini (*type safety*) devre dışı bırakacak işlemler içerir.

## 📦 `weak`

`weak` paketi, belleği zayıf biçimde (yani geri kazanımını engellemeden) güvenli şekilde referanslamak için yöntemler sağlar.

