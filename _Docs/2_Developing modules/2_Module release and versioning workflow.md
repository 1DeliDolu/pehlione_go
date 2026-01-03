## 🚢 Modül Yayınlama ve Sürümleme İş Akışı

Diğer geliştiricilerin kullanması için modüller geliştirdiğinizde, modülü kullanan geliştiriciler için güvenilir ve tutarlı bir deneyim sağlamaya yardımcı olan bir iş akışını takip edebilirsiniz. Bu konu, bu iş akışındaki üst düzey adımları açıklar.

Modül geliştirmeye genel bir bakış için ***Modülleri geliştirme ve yayınlama*** konusuna bakın.

---

## 🔎 Ayrıca bakınız

Eğer yalnızca kodunuzda harici paketleri kullanmak istiyorsanız, mutlaka ***Bağımlılıkları yönetme*** konusuna bakın.

Her yeni sürümle birlikte, modülünüzdeki değişiklikleri sürüm numarasıyla işaret edersiniz. Daha fazlası için ***Modül sürüm numaralandırma*** konusuna bakın.

---

## 🧭 Yaygın İş Akışı Adımları

Aşağıdaki sıra, örnek bir yeni modül için yayınlama ve sürümleme iş akışı adımlarını gösterir. Her adım hakkında daha fazla bilgi için bu konudaki bölümlere bakın.

1. Bir modül başlatın ve kaynaklarını, geliştiricilerin kullanmasını kolaylaştıracak ve sizin bakımınızı kolaylaştıracak şekilde organize edin.
   * Modül geliştirmede tamamen yeniyseniz, ***Eğitim: Bir Go modülü oluştur*** konusuna göz atın.
   * Go’nun merkeziyetsiz modül yayınlama sisteminde kodunuzu nasıl organize ettiğiniz önemlidir. Daha fazlası için ***Modül kaynağını yönetme*** konusuna bakın.
2. Yayınlanmamış modüldeki fonksiyonları çağıran yerel istemci kodu yazmak için ortamı hazırlayın.
   * Bir modül yayınlamadan önce, **`go get`** gibi komutları kullanan tipik bağımlılık yönetimi iş akışı için modül kullanılamaz. Bu aşamada modül kodunuzu test etmenin iyi bir yolu, onu çağıran kodunuza yerel bir dizinde bulunurken denemektir.
   * Yerel geliştirme hakkında daha fazla bilgi için ***Yayınlanmamış bir modüle karşı kod yazma*** bölümüne bakın.
3. Modülün kodu diğer geliştiricilerin denemesi için hazır olduğunda, *alpha* ve *beta* gibi **`v0`** ön sürümlerini yayınlamaya başlayın. Daha fazlası için ***Ön sürüm sürümleri yayınlama*** bölümüne bakın.
4. Kararlı olması garanti edilmeyen, ancak kullanıcıların deneyebileceği bir **`v0`** sürümü yayınlayın. Daha fazlası için ***İlk (kararsız) sürümü yayınlama*** bölümüne bakın.
5. **`v0`** sürümünüz yayınlandıktan sonra, onun yeni sürümlerini yayınlamaya devam edebilirsiniz (ve etmelisiniz!).
   * Bu yeni sürümler; hata düzeltmeleri (patch sürümleri), modülün  *public API* ’sine eklemeler (minor sürümler) ve hatta kırıcı değişiklikler içerebilir. **`v0`** sürümü kararlılık ya da geriye dönük uyumluluk garantisi vermediği için, sürümlerinde kırıcı değişiklikler yapabilirsiniz.
   * Daha fazlası için ***Hata düzeltmelerini yayınlama*** ve ***Geriye dönük uyumlu API değişikliklerini yayınlama*** bölümlerine bakın.
6. Kararlı bir sürümü yayınlamaya hazırlanırken, *alpha* ve *beta* olarak ön sürümler yayınlarsınız. Daha fazlası için ***Ön sürüm sürümleri yayınlama*** bölümüne bakın.
7. İlk kararlı sürüm olarak **`v1`** yayınlayın.
   * Bu, modülün kararlılığıyla ilgili taahhütler veren ilk sürümdür. Daha fazlası için ***İlk kararlı sürümü yayınlama*** bölümüne bakın.
8. **`v1`** sürümünde, hataları düzeltmeye ve gerektiğinde modülün  *public API* ’sine eklemeler yapmaya devam edin.
   * Daha fazlası için ***Hata düzeltmelerini yayınlama*** ve ***Geriye dönük uyumlu API değişikliklerini yayınlama*** bölümlerine bakın.
9. Kaçınılamadığında, kırıcı değişiklikleri yeni bir **major** sürümde yayınlayın.
   * **`v1.x.x`** ’ten  **`v2.x.x`** ’e gibi bir major sürüm güncellemesi, modülünüzün kullanıcıları için çok yıkıcı bir yükseltme olabilir. Bu son çare olmalıdır. Daha fazlası için ***Geriye dönük uyumsuz API değişikliklerini yayınlama*** bölümüne bakın.

---

## 🧪 Yayınlanmamış Bir Modüle Karşı Kod Yazma

Bir modül veya bir modülün yeni bir sürümünü geliştirmeye başladığınızda, henüz onu yayınlamamış olursunuz. Bir modül yayınlamadan önce, Go komutlarını kullanarak modülü bir bağımlılık olarak ekleyemezsiniz. Bunun yerine, ilk etapta, yayınlanmamış modüldeki fonksiyonları çağıran istemci kodunu farklı bir modül içinde yazarken, modülün yerel dosya sistemindeki bir kopyasına referans vermeniz gerekir.

İstemci modülün **`go.mod`** dosyasında **`replace`** yönergesini kullanarak modülü yerel olarak referans gösterebilirsiniz. Daha fazla bilgi için ***Yerel bir dizindeki modül kodunu gerektirme*** bölümüne bakın.

---

## 🏷️ Ön Sürüm (Pre-release) Sürümleri Yayınlama

Bir modülü başkalarının denemesi ve size geri bildirim vermesi için kullanılabilir hâle getirmek üzere ön sürümler yayınlayabilirsiniz. Bir ön sürüm, kararlılık garantisi içermez.

Ön sürüm numaraları, bir ön sürüm tanımlayıcısıyla eklenir. Sürüm numaraları hakkında daha fazlası için ***Modül sürüm numaralandırma*** konusuna bakın.

İki örnek:

* `v0.2.1-beta.1`
* `v1.2.3-alpha`

Bir ön sürümü kullanılabilir hâle getirirken, bu ön sürümü kullanan geliştiricilerin **`go get`** komutuyla sürümü açıkça belirtmeleri gerekeceğini unutmayın. Bunun nedeni, varsayılan olarak **`go`** komutunun modülü ararken ön sürümler yerine yayın (release) sürümlerini tercih etmesidir. Bu yüzden geliştiriciler ön sürümü şu örnekte olduğu gibi açıkça belirterek almalıdır:

```bash
go get example.com/theirmodule@v1.2.3-alpha
```

Bir ön sürümü, deponuzda modül kodunu etiketleyerek (tag) ve etikette ön sürüm tanımlayıcısını belirterek yayınlarsınız. Daha fazlası için ***Bir modül yayınlama*** konusuna bakın.

---

## 🧱 İlk (Kararsız) Sürümü Yayınlama

Ön sürüm yayınladığınızda olduğu gibi, kararlılık veya geriye dönük uyumluluk garantisi vermeyen; ancak kullanıcılarınıza modülü deneme ve geri bildirim verme fırsatı sağlayan yayın sürümleri de yayınlayabilirsiniz.

Kararsız yayınlar, sürüm numaraları **`v0.x.x`** aralığında olan sürümlerdir. **`v0`** sürümü, kararlılık ya da geriye dönük uyumluluk garantisi vermez. Ancak **`v1`** ve sonrasında kararlılık taahhütleri vermeden önce geri bildirim almak ve API’nizi rafine etmek için bir yol sunar. Daha fazlası için ***Modül sürüm numaralandırma*** konusuna bakın.

Yayınlanmış diğer sürümlerde olduğu gibi, kararlı bir **`v1`** sürümünü yayınlamaya doğru ilerlerken değişiklik yaptıkça **`v0`** sürüm numarasının minor ve patch kısımlarını artırabilirsiniz. Örneğin **`v0.0.0.0`** yayınladıktan sonra, ilk hata düzeltmeleri setiyle **`v0.0.1`** yayınlayabilirsiniz.

Bir örnek sürüm numarası:

* `v0.1.3`

Kararsız bir yayını, deponuzda modül kodunu **`v0`** sürüm numarasıyla etiketleyerek (tag) yayınlarsınız. Daha fazlası için ***Bir modül yayınlama*** konusuna bakın.

---

## ✅ İlk Kararlı Sürümü Yayınlama

İlk kararlı sürümünüz **`v1.x.x`** sürüm numarasına sahip olacaktır. İlk kararlı sürüm, geri bildirim aldığınız, hataları düzelttiğiniz ve modülü kullanıcılar için stabilize ettiğiniz ön sürümleri ve **`v0`** sürümlerini takip eder.

**`v1`** sürümüyle birlikte, modülünüzü kullanan geliştiricilere şu taahhütlerde bulunursunuz:

* Major sürümün sonraki minor ve patch sürümlerine kendi kodlarını bozmadan yükseltebilirler.
* Modülün  *public API* ’sinde — fonksiyon ve metot imzaları dahil — geriye dönük uyumluluğu bozan başka değişiklikler yapmayacaksınız.
* Dışa aktarılan (exported) türleri kaldırmayacaksınız; bu, geriye dönük uyumluluğu bozar.
* API’nizdeki gelecekteki değişiklikler (örneğin bir  *struct* ’a yeni bir alan eklemek gibi) geriye dönük uyumlu olacaktır ve yeni bir minor sürümde yer alacaktır.
* Hata düzeltmeleri (örneğin bir güvenlik düzeltmesi) bir patch sürümünde veya bir minor sürümün parçası olarak yer alacaktır.

Not: İlk major sürümünüz bir **`v0`** sürümü olabilir; ancak **`v0`** sürümü kararlılık veya geriye dönük uyumluluk garantisi vermez. Sonuç olarak  **`v0`** ’dan  **`v1`** ’e artırırken, **`v0`** sürümü kararlı kabul edilmediği için geriye dönük uyumluluğu bozma konusunda dikkatli olmanız gerekmez.

Sürüm numaraları hakkında daha fazlası için ***Modül sürüm numaralandırma*** konusuna bakın.

Kararlı bir sürüm numarası örneği:

* `v1.0.0`

İlk kararlı sürümü, deponuzda modül kodunu **`v1`** sürüm numarasıyla etiketleyerek (tag) yayınlarsınız. Daha fazlası için ***Bir modül yayınlama*** konusuna bakın.

---

## 🐞 Hata Düzeltmelerini Yayınlama

Değişikliklerin yalnızca hata düzeltmeleriyle sınırlı olduğu bir yayın yayınlayabilirsiniz. Bu, **patch sürümü** olarak bilinir.

Bir patch sürümü yalnızca küçük değişiklikler içerir. Özellikle, modülün  *public API* ’sinde hiçbir değişiklik içermez. Tüketici (consuming) kodu geliştiren geliştiriciler bu sürüme güvenle ve kodlarını değiştirmeye gerek kalmadan yükseltebilir.

Not: Patch sürümünüz, o modülün kendi geçişli bağımlılıklarını ( *transitive dependencies* ) patch sürümünden daha fazla yükseltmemeye çalışmalıdır. Aksi hâlde birisi modülünüzün patch sürümüne yükseltirken, kullandığı bir geçişli bağımlılıkta yanlışlıkla daha kapsamlı bir değişikliği içeri çekebilir.

Patch sürümü, modül sürüm numarasının patch kısmını artırır. Daha fazlası için ***Modül sürüm numaralandırma*** konusuna bakın.

Aşağıdaki örnekte **`v1.0.1`** bir patch sürümüdür.

Eski sürüm: `v1.0.0`
Yeni sürüm: `v1.0.1`

Bir patch sürümünü, deponuzda modül kodunu etiketleyerek (tag) ve etikette patch sürüm numarasını artırarak yayınlarsınız. Daha fazlası için ***Bir modül yayınlama*** konusuna bakın.

---

## 🧩 Geriye Dönük Uyumlu API Değişikliklerini Yayınlama

Modülünüzün  *public API* ’sinde kırıcı olmayan değişiklikler yapabilir ve bu değişiklikleri bir minor sürüm yayınıyla yayınlayabilirsiniz.

Bu sürüm API’yi değiştirir, ancak çağıran kodu bozacak şekilde değil. Bu; modülün kendi bağımlılıklarındaki değişiklikleri veya yeni fonksiyonlar, metotlar, *struct* alanları ya da türler eklenmesini içerebilir. İçerdiği değişikliklere rağmen, bu tür bir yayın modülün fonksiyonlarını çağıran mevcut kod için geriye dönük uyumluluk ve kararlılık garantisi verir.

Minor sürüm, modül sürüm numarasının minor kısmını artırır. Daha fazlası için ***Modül sürüm numaralandırma*** konusuna bakın.

Aşağıdaki örnekte **`v1.1.0`** bir minor sürümüdür.

Eski sürüm: `v1.0.1`
Yeni sürüm: `v1.1.0`

Bir minor sürümünü, deponuzda modül kodunu etiketleyerek (tag) ve etikette minor sürüm numarasını artırarak yayınlarsınız. Daha fazlası için ***Bir modül yayınlama*** konusuna bakın.

---

## 💥 Geriye Dönük Uyumsuz API Değişikliklerini Yayınlama

Geriye dönük uyumluluğu bozan bir sürümü, bir major sürüm yayını yayınlayarak yayınlayabilirsiniz.

Bir major sürüm yayını geriye dönük uyumluluk garantisi vermez; tipik olarak bunun nedeni, modülün  *public API* ’sinde modülün önceki sürümlerini kullanan kodu bozacak değişiklikler içermesidir.

Major sürüm yükseltmesinin, modüle dayanan kod üzerindeki yıkıcı etkisi nedeniyle, mümkünse major sürüm güncellemesinden kaçınmalısınız. Major sürüm güncellemeleri hakkında daha fazlası için ***Büyük bir sürüm güncellemesi geliştirme*** konusuna bakın. Kırıcı değişiklik yapmaktan kaçınma stratejileri için, Go blogundaki ***Modüllerinizi uyumlu tutma*** yazısına bakın.

Diğer sürüm türlerini yayınlamak temelde modül kodunu sürüm numarasıyla etiketlemeyi gerektirirken, bir major sürüm güncellemesi yayınlamak daha fazla adım gerektirir.

Yeni major sürümün geliştirilmesine başlamadan önce, deponuzda yeni sürümün kaynağı için bir yer oluşturun.

Bunu yapmanın bir yolu, deponuzda özellikle yeni major sürüm ve onun sonraki minor ve patch sürümleri için yeni bir dal (branch) oluşturmaktır. Daha fazlası için ***Modül kaynağını yönetme*** konusuna bakın.

Modülün **`go.mod`** dosyasında, modül yolunu yeni major sürüm numarasını ekleyerek güncelleyin; aşağıdaki örnekte olduğu gibi:

```text
example.com/mymodule/v2
```

Modül yolu, modülün tanımlayıcısı olduğundan, bu değişiklik fiilen yeni bir modül oluşturur. Ayrıca paket yolunu değiştirerek geliştiricilerin kodlarını bozan bir sürümü istemeden import etmeyeceğini garanti eder. Bunun yerine, yükseltmek isteyenler eski yoldaki kullanımları açıkça yeni yol ile değiştirecektir.

Kodunuzda, güncellediğiniz modüldeki paketleri import ettiğiniz tüm paket yollarını — güncellediğiniz modül içindeki paketler dahil — değiştirin. Bunu yapmanız gerekir; çünkü modül yolunu değiştirdiniz.

Her yeni yayında olduğu gibi, resmi bir yayın yayınlamadan önce geri bildirim ve hata raporları almak için ön sürümler yayınlamalısınız.

Yeni major sürümü, deponuzda modül kodunu etiketleyerek (tag) ve etikette major sürüm numarasını artırarak yayınlayın — örneğin  **`v1.5.2`** ’den  **`v2.0.0`** ’a.

Daha fazlası için ***Bir modül yayınlama*** konusuna bakın.
