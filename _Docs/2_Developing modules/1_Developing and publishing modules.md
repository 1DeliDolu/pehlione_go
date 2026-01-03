## 🚀 Modülleri Geliştirme ve Yayınlama

İlgili paketleri modüller içinde toplayabilir, ardından modülleri diğer geliştiricilerin kullanması için yayınlayabilirsiniz. Bu konu, modülleri geliştirme ve yayınlamaya dair genel bir bakış sunar.

Modülleri geliştirmeyi, yayınlamayı ve kullanmayı desteklemek için şunları kullanırsınız:

* Zaman içinde yeni sürümlerle revize ederek modüller geliştirdiğiniz ve yayınladığınız bir iş akışı. **Bkz.  *Modülleri geliştirme ve yayınlama iş akışı* .**
* Bir modülün kullanıcılarının modülü anlamasına ve yeni sürümlere kararlı bir şekilde yükseltmesine yardımcı olan tasarım pratikleri. **Bkz.  *Tasarım ve geliştirme* .**
* Modülleri yayınlamak ve kodlarını almak için merkeziyetsiz bir sistem. Modülünüzü kendi deponuzdan diğer geliştiricilerin kullanımına sunar ve bir sürüm numarasıyla yayınlarsınız. **Bkz.  *Merkeziyetsiz yayınlama* .**
* Geliştiricilerin modülünüzü bulabileceği bir paket arama motoru ve dokümantasyon tarayıcısı ( ***pkg.go.dev*** ). **Bkz.  *Paket keşfi* .**
* Modülünüzü kullanan geliştiricilere kararlılık ve geriye dönük uyumluluk beklentilerini iletmek için bir modül sürüm numaralandırma kuralı. **Bkz.  *Sürümleme* .**
* Diğer geliştiricilerin bağımlılıkları yönetmesini kolaylaştıran Go araçları; modülünüzün kaynağını alma, yükseltme vb. dahil. **Bkz.  *Bağımlılıkları yönetme* .**

### 🔎 Ayrıca bakınız

Eğer sadece başkalarının geliştirdiği paketleri kullanmakla ilgileniyorsanız, bu konu size göre değildir. Bunun yerine ***Bağımlılıkları yönetme*** konusuna bakın.

Modül geliştirmeye dair birkaç temel içeren bir eğitim için ***Eğitim: Bir Go modülü oluştur*** konusuna bakın.

---

## 🧭 Modülleri Geliştirme ve Yayınlama İş Akışı

Başkaları için modüllerinizi yayınlamak istediğinizde, bu modüllerin kullanılmasını kolaylaştırmak için birkaç konvansiyonu benimsersiniz.

Aşağıdaki üst düzey adımlar, ***Modül sürümleme ve yayınlama iş akışı*** içinde daha ayrıntılı olarak açıklanır:

1. Modülün içereceği paketleri tasarlayın ve kodlayın.
2. Go araçları aracılığıyla başkaları için erişilebilir olmasını sağlayan konvansiyonlarla kodu deponuza commit edin.
3. Modülü geliştiriciler tarafından keşfedilebilir hâle getirmek için yayınlayın.
4. Zaman içinde, her sürümün kararlılığını ve geriye dönük uyumluluğunu işaret eden bir sürüm numaralandırma kuralını kullanan sürümlerle modülü revize edin.

---

## 🧩 Tasarım ve Geliştirme

Modülünüz, içindeki fonksiyonlar ve paketler tutarlı bir bütün oluşturuyorsa, geliştiricilerin onu bulması ve kullanması daha kolay olur. Bir modülün  **public API** ’sini tasarlarken, işlevselliğini odaklı ve ayrık tutmaya çalışın.

Ayrıca modülünüzü geriye dönük uyumluluğu dikkate alarak tasarlamak ve geliştirmek, kullanıcılarının yükseltme yaparken kendi kodlarında oluşacak değişim ihtiyacını (churn) en aza indirmesine yardımcı olur. Geriye dönük uyumluluğu bozan bir sürüm yayınlamaktan kaçınmak için kod içinde belirli teknikler kullanabilirsiniz. Bu teknikler hakkında daha fazla bilgi için Go blogunda ***Modüllerinizi uyumlu tutma*** konusuna bakın.

Bir modülü yayınlamadan önce, **`replace`** yönergesini kullanarak onu yerel dosya sistemi üzerinde referans alabilirsiniz. Bu, modül hâlâ geliştirilirken modüldeki fonksiyonları çağıran istemci kodu yazmayı kolaylaştırır. Daha fazla bilgi için ***Modül sürümleme ve yayınlama iş akışı*** içindeki **“Yayınlanmamış bir modüle karşı kod yazma”** bölümüne bakın.

---

## 🌐 Merkeziyetsiz Yayınlama

Go’da modülünüzü yayınlamak için, deponuzdaki kodu etiketleyerek (tag) diğer geliştiricilerin kullanmasına uygun hâle getirirsiniz. Modülünüzü merkezi bir servise göndermeniz gerekmez; çünkü Go araçları modülünüzü doğrudan deponuzdan (modülün yolu kullanılarak bulunur; bu yol şeması çıkarılmış bir URL’dir) veya bir proxy sunucusundan indirebilir.

Geliştiriciler kodlarında paketinizi import ettikten sonra, derleme sırasında kullanılacak modül kodunu indirmek için Go araçlarını (bunlara **`go get`** komutu da dahildir) kullanır. Bu modeli desteklemek için, Go araçlarının (başka bir geliştirici adına) deponuzdan modül kaynak kodunu alabilmesini mümkün kılan konvansiyonları ve en iyi uygulamaları izlersiniz. Örneğin Go araçları, modülü yayınlamak için kullandığınız sürüm numarasıyla birlikte belirttiğiniz modül yolunu (module path) kullanarak modülünüzün konumunu belirler ve kullanıcıları için indirir.

Kaynak ve yayınlama konvansiyonları ile en iyi uygulamalar hakkında daha fazla bilgi için ***Modül kaynağını yönetme*** konusuna bakın.

Bir modül yayınlamaya yönelik adım adım talimatlar için ***Bir modül yayınlama*** konusuna bakın.

---

## 🔍 Paket Keşfi

Modülünüzü yayınladıktan ve biri onu Go araçlarıyla çektikten sonra, ***pkg.go.dev*** üzerindeki Go paket keşif sitesinde görünür hâle gelir. Orada geliştiriciler site içinde arama yaparak modülünüzü bulabilir ve dokümantasyonunu okuyabilir.

Modülü kullanmaya başlamak için, bir geliştirici modülden paketleri import eder, ardından kaynak kodunu derleme sırasında kullanmak üzere indirmek için **`go get`** komutunu çalıştırır.

Geliştiricilerin modülleri nasıl bulduğu ve kullandığı hakkında daha fazla bilgi için ***Bağımlılıkları yönetme*** konusuna bakın.

---

## 🏷️ Sürümleme

Modülünüzü zaman içinde revize edip iyileştirirken, her sürümün kararlılığını ve geriye dönük uyumluluğunu işaret edecek şekilde tasarlanmış (semantik sürümleme modeline dayalı) sürüm numaraları atarsınız. Bu, modülünüzü kullanan geliştiricilerin modülün ne zaman kararlı olduğunu ve bir yükseltmenin davranışta önemli değişiklikler içerip içermeyeceğini belirlemesine yardımcı olur. Modülün sürüm numarasını, deponuzdaki modül kaynak kodunu bu numarayla etiketleyerek (tag) belirtirsiniz.

Büyük sürüm güncellemeleri geliştirme hakkında daha fazlası için ***Büyük bir sürüm güncellemesi geliştirme*** konusuna bakın.

Go modüllerinde semantik sürümleme modelini nasıl kullandığınız hakkında daha fazla bilgi için ***Modül sürüm numaralandırma*** konusuna bakın.
