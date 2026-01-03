
## 📘 Giriş¶

Bu, Go programlama dili için başvuru kılavuzudur. Daha fazla bilgi ve diğer belgeler için go.dev’e bakın. 

Go, *sistem programlama* göz önünde bulundurularak tasarlanmış genel amaçlı bir dildir. *Güçlü tipli*dir ve *çöp toplayıcılı*dır; ayrıca eşzamanlı programlama için açık desteğe sahiptir. Programlar, özellikleri bağımlılıkların verimli biçimde yönetilmesine olanak tanıyan paketlerden (*package*) oluşturulur.

Sözdizimi kompakt ve ayrıştırılması kolaydır; bu da tümleşik geliştirme ortamları gibi otomatik araçlarla kolay analiz yapılmasını sağlar.

---

## 🧩 Gösterim¶

Sözdizimi, *Extended Backus-Naur Form (EBNF)*’nin bir varyantı kullanılarak belirtilir:

```ebnf
Syntax      = { Production } .
Production  = production_name "=" [ Expression ] "." .
Expression  = Term { "|" Term } .
Term        = Factor { Factor } .
Factor      = production_name | token [ "…" token ] | Group | Option | Repetition .
Group       = "(" Expression ")" .
Option      = "[" Expression "]" .
Repetition  = "{" Expression "}" .
```

Üretimler, terimlerden ve aşağıdaki işleçlerden (artan öncelik sırasıyla) oluşturulan ifadelerdir:

```text
|   alternation
()  grouping
[]  option (0 or 1 times)
{}  repetition (0 to n times)
```

Küçük harfli üretim adları, leksik (terminal) belirteçleri tanımlamak için kullanılır. Terminal olmayanlar *CamelCase* biçimindedir. Leksik belirteçler çift tırnak `""` veya ters tırnak ` `` ` içine alınır.

`a … b` biçimi, a’dan b’ye kadar olan karakter kümesini alternatifler olarak temsil eder. Yatay üç nokta `…` başka yerlerde de daha fazla belirtilmeyen çeşitli enumerasyonları veya kod parçacıklarını gayriresmî olarak göstermek için kullanılır. `…` karakteri (üç karakter `...`’ın aksine) Go dilinin bir belirteci değildir.

`[Go 1.xx]` biçimindeki bir bağlantı, açıklanan bir dil özelliğinin (veya onun bir yönünün) dil sürümü 1.xx ile değiştirildiğini ya da eklendiğini ve dolayısıyla derlemek için en az o sürümün gerektiğini belirtir. Ayrıntılar için ekte bağlantısı verilen bölüme bakın.

---

## 🗂️ Kaynak kodu gösterimi¶

Kaynak kodu, *UTF-8* ile kodlanmış *Unicode* metnidir. Metin kanonikleştirilmez; bu nedenle tek bir aksanlı kod noktası, aksan ile bir harfin birleşiminden oluşturulmuş aynı karakterden farklıdır; bunlar iki ayrı kod noktası olarak ele alınır. Basitlik için, bu belge kaynak metindeki bir *Unicode kod noktasına* atıfla niteliksiz “karakter” terimini kullanır.

Her kod noktası ayrıdır; örneğin büyük ve küçük harfler farklı karakterlerdir.

Uygulama kısıtlaması: Diğer araçlarla uyumluluk için, bir derleyici kaynak metinde *NUL* karakterini (U+0000) yasaklayabilir.

Uygulama kısıtlaması: Diğer araçlarla uyumluluk için, bir derleyici *UTF-8* ile kodlanmış *byte order mark*’ı (U+FEFF) ilk Unicode kod noktası ise yok sayabilir. Bir *byte order mark* başka herhangi bir yerde yasaklanabilir.

---

## 🔤 Karakterler¶

Aşağıdaki terimler belirli Unicode karakter kategorilerini belirtmek için kullanılır:

```ebnf
newline        = /* the Unicode code point U+000A */ .
unicode_char   = /* an arbitrary Unicode code point except newline */ .
unicode_letter = /* a Unicode code point categorized as "Letter" */ .
unicode_digit  = /* a Unicode code point categorized as "Number, decimal digit" */ .
```

*The Unicode Standard 8.0*’da, Bölüm 4.5 “General Category” bir karakter kategorileri kümesi tanımlar. Go, Letter kategorilerinden herhangi birinde (Lu, Ll, Lt, Lm veya Lo) olan tüm karakterleri *Unicode harfleri* olarak; Number kategorisi Nd’de olanları ise *Unicode rakamları* olarak kabul eder.

---

## 🔡 Harfler ve rakamlar¶

Alt çizgi karakteri `_` (U+005F) küçük harf olarak kabul edilir.

```ebnf
letter        = unicode_letter | "_" .
decimal_digit = "0" … "9" .
binary_digit  = "0" | "1" .
octal_digit   = "0" … "7" .
hex_digit     = "0" … "9" | "A" … "F" | "a" … "f" .
```

---

## 🧱 Leksik öğeler¶

### 💬 Yorumlar¶

Yorumlar, program belgeleri olarak hizmet eder. İki biçimi vardır:

* Satır yorumları `//` karakter dizisiyle başlar ve satır sonunda biter.
* Genel yorumlar `/*` karakter dizisiyle başlar ve onu izleyen ilk `*/` karakter dizisiyle biter.

Bir yorum, bir *rune* ya da *string literal* içinde veya başka bir yorumun içinde başlayamaz. Yeni satır içermeyen bir genel yorum boşluk gibi davranır. Diğer tüm yorumlar yeni satır gibi davranır.

### 🧩 Belirteçler (Tokens)¶

Belirteçler, Go dilinin söz varlığını oluşturur. Dört sınıf vardır: *identifier*’lar, *keyword*’ler, *operator*’ler ve *punctuation*’lar ile *literal*’lar. Boşluk (U+0020), yatay sekme (U+0009), satır başı (U+000D) ve yeni satırdan (U+000A) oluşan beyaz alan, aksi halde tek bir belirteçte birleşecek belirteçleri ayırması dışında yok sayılır. Ayrıca, bir yeni satır veya dosya sonu noktalı virgül eklenmesini tetikleyebilir. Girdi belirteçlere bölünürken, sıradaki belirteç geçerli bir belirteç oluşturan en uzun karakter dizisidir.

### 🪡 Noktalı virgüller¶

Biçimsel sözdizimi, birtakım üretimlerde sonlandırıcı olarak noktalı virgül `";"` kullanır. Go programları, aşağıdaki iki kural ile bu noktalı virgüllerin çoğunu atlayabilir:

1. Girdi belirteçlere bölünürken, bir satırın son belirteci aşağıdakilerden biriyse, belirteç akışına o belirtecin hemen ardından otomatik olarak bir noktalı virgül eklenir:

* bir *identifier*
* bir tamsayı, kayan nokta, sanal (imaginary), *rune* veya *string literal*
* şu anahtar sözcüklerden biri: `break`, `continue`, `fallthrough` veya `return`
* şu işleç ve noktalama işaretlerinden biri: `++`, `--`, `)`, `]` veya `}`

2. Karmaşık ifadelerin tek satırı kaplayabilmesi için, kapanış `")"` veya `"}"` öncesinde noktalı virgül atlanabilir.

İdiomatik kullanımı yansıtmak için, bu belgedeki kod örnekleri bu kuralları kullanarak noktalı virgülleri eler.

### 🏷️ Tanımlayıcılar (Identifiers)¶

Tanımlayıcılar, değişkenler ve türler gibi program varlıklarını adlandırır. Bir tanımlayıcı, bir veya daha fazla harf ve rakam dizisidir. Bir tanımlayıcının ilk karakteri harf olmalıdır.

```ebnf
identifier = letter { letter | unicode_digit } .
```

Örnekler:

```text
a
_x9
ThisVariableIsExported
αβ
```

Bazı tanımlayıcılar önceden tanımlanmıştır.

### 🔑 Anahtar sözcükler (Keywords)¶

Aşağıdaki anahtar sözcükler ayrılmıştır ve tanımlayıcı olarak kullanılamaz.

```text
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

### ➕ İşleçler ve noktalama¶

Aşağıdaki karakter dizileri işleçleri (atama işleçleri dahil) ve noktalama işaretlerini temsil eder [Go 1.18]:

```text
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=          ~
```

---

## 🔢 Tamsayı sabitleri¶

Bir tamsayı sabiti, bir tamsayı sabitini temsil eden rakamlar dizisidir. İsteğe bağlı bir önek, onluk olmayan tabanı belirler: ikilik için `0b` veya `0B`, sekizlik için `0`, `0o` veya `0O`, onaltılık için `0x` veya `0X` [Go 1.13]. Tek başına `0` onluk sıfır kabul edilir. Onaltılık sabitlerde a–f ve A–F harfleri 10–15 değerlerini temsil eder.

Okunabilirlik için, alt çizgi karakteri `_` bir taban önekinden sonra veya ardışık rakamların arasında görünebilir; bu alt çizgiler sabitin değerini değiştirmez.

```ebnf
int_lit        = decimal_lit | binary_lit | octal_lit | hex_lit .
decimal_lit    = "0" | ( "1" … "9" ) [ [ "_" ] decimal_digits ] .
binary_lit     = "0" ( "b" | "B" ) [ "_" ] binary_digits .
octal_lit      = "0" [ "o" | "O" ] [ "_" ] octal_digits .
hex_lit        = "0" ( "x" | "X" ) [ "_" ] hex_digits .

decimal_digits = decimal_digit { [ "_" ] decimal_digit } .
binary_digits  = binary_digit { [ "_" ] binary_digit } .
octal_digits   = octal_digit { [ "_" ] octal_digit } .
hex_digits     = hex_digit { [ "_" ] hex_digit } .
```

Örnekler:

```text
42
4_2
0600
0_600
0o600
0O600       // ikinci karakter büyük harf 'O'dur
0xBadFace
0xBad_Face
0x_67_7a_2f_cc_40_c6
170141183460469231731687303715884105727
170_141183_460469_231731_687303_715884_105727

_42         // bir tanımlayıcıdır, tamsayı sabiti değildir
42_         // geçersiz: _ ardışık rakamları ayırmalıdır
4__2        // geçersiz: aynı anda yalnızca bir _
0_xBadFace  // geçersiz: _ ardışık rakamları ayırmalıdır
```

---

## 🔣 Kayan nokta sabitleri¶

Bir kayan nokta sabiti, bir kayan nokta sabitinin onluk veya onaltılık gösterimidir.

Onluk kayan nokta sabiti; bir tamsayı kısmı (onluk rakamlar), bir ondalık nokta, bir kesir kısmı (onluk rakamlar) ve bir üs kısmından (isteğe bağlı işaret ve onluk rakamları izleyen `e` veya `E`) oluşur. Tamsayı kısmı veya kesir kısmından biri atlanabilir; ondalık nokta veya üs kısmından biri de atlanabilir. Üs değeri `exp`, mantisayı (tamsayı ve kesir kısmı) `10exp` ile ölçekler.

Onaltılık kayan nokta sabiti; `0x` veya `0X` öneği, bir tamsayı kısmı (onaltılık rakamlar), bir radix noktası, bir kesir kısmı (onaltılık rakamlar) ve bir üs kısmından (isteğe bağlı işaret ve onluk rakamları izleyen `p` veya `P`) oluşur. Tamsayı kısmı veya kesir kısmından biri atlanabilir; radix noktası da atlanabilir, ancak üs kısmı zorunludur. (Bu sözdizimi IEEE 754-2008 §5.12.3’te verilenle eşleşir.) Üs değeri `exp`, mantisayı (tamsayı ve kesir kısmı) `2exp` ile ölçekler [Go 1.13].

Okunabilirlik için, alt çizgi karakteri `_` bir taban önekinden sonra veya ardışık rakamların arasında görünebilir; bu alt çizgiler sabitin değerini değiştirmez.

```ebnf
float_lit         = decimal_float_lit | hex_float_lit .

decimal_float_lit = decimal_digits "." [ decimal_digits ] [ decimal_exponent ] |
                    decimal_digits decimal_exponent |
                    "." decimal_digits [ decimal_exponent ] .
decimal_exponent  = ( "e" | "E" ) [ "+" | "-" ] decimal_digits .

hex_float_lit     = "0" ( "x" | "X" ) hex_mantissa hex_exponent .
hex_mantissa      = [ "_" ] hex_digits "." [ hex_digits ] |
                    [ "_" ] hex_digits |
                    "." hex_digits .
hex_exponent      = ( "p" | "P" ) [ "+" | "-" ] decimal_digits .
```

Örnekler:

```text
0.
72.40
072.40       // == 72.40
2.71828
1.e+0
6.67428e-11
1E6
.25
.12345E+5
1_5.         // == 15.0
0.15e+0_2    // == 15.0

0x1p-2       // == 0.25
0x2.p10      // == 2048.0
0x1.Fp+0     // == 1.9375
0X.8p-0      // == 0.5
0X_1FFFP-16  // == 0.1249847412109375
0x15e-2      // == 0x15e - 2 (tamsayı çıkarma)

0x.p1        // geçersiz: mantisada rakam yok
1p-2         // geçersiz: p üssü onaltılık mantisa gerektirir
0x1.5e-2     // geçersiz: onaltılık mantisa p üssü gerektirir
1_.5         // geçersiz: _ ardışık rakamları ayırmalıdır
1._5         // geçersiz: _ ardışık rakamları ayırmalıdır
1.5_e1       // geçersiz: _ ardışık rakamları ayırmalıdır
1.5e_1       // geçersiz: _ ardışık rakamları ayırmalıdır
1.5e1_       // geçersiz: _ ardışık rakamları ayırmalıdır
```

---

## 🧠 Sanal (Imaginary) sabitler¶

Bir sanal sabit, bir kompleks sabitin sanal kısmını temsil eder. Bir tamsayı veya kayan nokta sabitini izleyen küçük harf `i`’den oluşur. Bir sanal sabitin değeri, ilgili tamsayı veya kayan nokta sabitinin değeri ile sanal birim `i`’nin çarpımıdır [Go 1.13]

```ebnf
imaginary_lit = (decimal_digits | int_lit | float_lit) "i" .
```

Geriye dönük uyumluluk için, tamamen onluk rakamlardan (ve muhtemelen alt çizgilerden) oluşan sanal sabitin tamsayı kısmı, başında sıfır olsa bile onluk tamsayı kabul edilir.

```text
0i
0123i         // geriye dönük uyumluluk için == 123i
0o123i        // == 0o123 * 1i == 83i
0xabci        // == 0xabc * 1i == 2748i
0.i
2.71828i
1.e+0i
6.67428e-11i
1E6i
.25i
.12345E+5i
0x1p-2i       // == 0x1p-2 * 1i == 0.25i
```

---

## 🪪 Rune sabitleri¶

Bir *rune* sabiti, bir *Unicode kod noktasını* tanımlayan tamsayı değerli bir rune sabitini temsil eder. Bir rune sabiti, `'x'` veya `'\n'` gibi tek tırnaklar içine alınmış bir veya daha fazla karakter olarak ifade edilir. Tırnaklar içinde, yeni satır ve kaçırılmamış tek tırnak dışında herhangi bir karakter görünebilir. Tek tırnaklı tek karakter, karakterin kendisinin Unicode değerini temsil ederken; ters eğik çizgiyle başlayan çok karakterli diziler değerleri çeşitli biçimlerde kodlar.

En basit biçim, tırnaklar içindeki tek karakteri temsil eder; Go kaynak metni UTF-8 ile kodlanmış Unicode karakterleri olduğundan, bir tamsayı değeri temsil etmek için birden çok UTF-8 baytı gerekebilir. Örneğin `'a'` sabiti, Unicode U+0061 olan ve değeri 0x61 olan `a` harfini temsil eden tek bir bayt tutarken; `'ä'` sabiti, U+00E4 olan ve değeri 0xe4 olan `a-umlaut` karakterini temsil eden iki bayt (0xc3 0xa4) tutar.

Birkaç ters eğik çizgi kaçışı, keyfi değerlerin ASCII metin olarak kodlanmasına izin verir. Tamsayı değerini sayısal sabit olarak temsil etmenin dört yolu vardır: Tam olarak iki onaltılık rakamı izleyen `\x`; tam olarak dört onaltılık rakamı izleyen `\u`; tam olarak sekiz onaltılık rakamı izleyen `\U`; ve tam olarak üç sekizlik rakamı izleyen yalın ters eğik çizgi `\`. Her durumda sabitin değeri, ilgili tabandaki rakamların temsil ettiği değerdir.

Bu gösterimlerin tümü tamsayı ürettiği hâlde, geçerli aralıkları farklıdır. Sekizlik kaçışlar 0 ile 255 (dahil) arasında bir değeri temsil etmelidir. Onaltılık kaçışlar yapıları gereği bu koşulu sağlar. `\u` ve `\U` kaçışları Unicode kod noktalarını temsil eder; dolayısıyla bunların içinde bazı değerler yasadışıdır; özellikle 0x10FFFF üzerindekiler ve surrogate yarımlar.

Bir ters eğik çizgiden sonra, belirli tek karakterli kaçışlar özel değerleri temsil eder:

```text
\a   U+0007 alert veya zil
\b   U+0008 backspace
\f   U+000C form feed
\n   U+000A line feed veya yeni satır
\r   U+000D carriage return
\t   U+0009 yatay sekme
\v   U+000B dikey sekme
\\   U+005C ters eğik çizgi
\'   U+0027 tek tırnak  (yalnızca rune sabitleri içinde geçerli kaçış)
\"   U+0022 çift tırnak (yalnızca string sabitleri içinde geçerli kaçış)
```

Bir rune sabitinde ters eğik çizgiyi izleyen tanınmayan bir karakter yasadışıdır.

```ebnf
rune_lit         = "'" ( unicode_value | byte_value ) "'" .
unicode_value    = unicode_char | little_u_value | big_u_value | escaped_char .
byte_value       = octal_byte_value | hex_byte_value .
octal_byte_value = `\` octal_digit octal_digit octal_digit .
hex_byte_value   = `\` "x" hex_digit hex_digit .
little_u_value   = `\` "u" hex_digit hex_digit hex_digit hex_digit .
big_u_value      = `\` "U" hex_digit hex_digit hex_digit hex_digit
                           hex_digit hex_digit hex_digit hex_digit .
escaped_char     = `\` ( "a" | "b" | "f" | "n" | "r" | "t" | "v" | `\` | "'" | `"` ) .
```

Örnekler:

```text
'a'
'ä'
'本'
'\t'
'\000'
'\007'
'\377'
'\x07'
'\xff'
'\u12e4'
'\U00101234'
'\''         // tek tırnak karakteri içeren rune sabiti
'aa'         // yasadışı: çok fazla karakter
'\k'         // yasadışı: ters eğik çizgiden sonra k tanınmaz
'\xa'        // yasadışı: çok az onaltılık rakam
'\0'         // yasadışı: çok az sekizlik rakam
'\400'       // yasadışı: sekizlik değer 255’i aşıyor
'\uDFFF'     // yasadışı: surrogate yarım
'\U00110000' // yasadışı: geçersiz Unicode kod noktası
```

---

## 🧵 String sabitleri¶

Bir string sabiti, bir karakter dizisinin birleştirilmesinden elde edilen bir string sabitini temsil eder. İki biçimi vardır: ham (*raw*) string sabitleri ve yorumlanan (*interpreted*) string sabitleri.

Ham string sabitleri, `` `foo` `` gibi ters tırnaklar arasındaki karakter dizileridir. Tırnaklar içinde, ters tırnak dışında herhangi bir karakter görünebilir. Ham string sabitinin değeri, tırnaklar arasındaki yorumlanmamış (örtük olarak UTF-8 kodlu) karakterlerden oluşan stringdir; özellikle ters eğik çizgilerin özel anlamı yoktur ve string yeni satırlar içerebilir. Ham string sabitleri içindeki carriage return karakterleri (`'\r'`) ham string değerinden çıkarılır.

Yorumlanan string sabitleri, `"bar"` gibi çift tırnaklar arasındaki karakter dizileridir. Tırnaklar içinde, yeni satır ve kaçırılmamış çift tırnak dışında herhangi bir karakter görünebilir. Tırnaklar arasındaki metin sabitin değerini oluşturur; ters eğik çizgi kaçışları rune sabitlerindeki gibi yorumlanır (ancak `\'` yasadışıdır ve `\"` yasaldır) ve aynı kısıtlamalara tabidir. Üç basamaklı sekizlik (`\nnn`) ve iki basamaklı onaltılık (`\xnn`) kaçışlar sonuç stringinde tek tek baytları temsil eder; diğer tüm kaçışlar tek tek karakterlerin (muhtemelen çok baytlı) UTF-8 kodlamasını temsil eder. Bu nedenle bir string sabiti içinde `\377` ve `\xFF` değeri 0xFF=255 olan tek bir baytı temsil ederken; `ÿ`, `\u00FF`, `\U000000FF` ve `\xc3\xbf` U+00FF karakterinin UTF-8 kodlaması olan iki baytı (0xc3 0xbf) temsil eder.

```ebnf
string_lit             = raw_string_lit | interpreted_string_lit .
raw_string_lit         = "`" { unicode_char | newline } "`" .
interpreted_string_lit = `"` { unicode_value | byte_value } `"` .
```

Örnekler:

```text
`abc`                // "abc" ile aynıdır
`\n
\n`                  // "\\n\n\\n" ile aynıdır
"\n"
"\""                 // `"` ile aynıdır
"Hello, world!\n"
"日本語"
"\u65e5本\U00008a9e"
"\xff\u00FF"
"\uD800"             // yasadışı: surrogate yarım
"\U00110000"         // yasadışı: geçersiz Unicode kod noktası
```

Bu örneklerin tümü aynı stringi temsil eder:

```text
"日本語"                                 // UTF-8 girdi metni
`日本語`                                 // ham sabit olarak UTF-8 girdi metni
"\u65e5\u672c\u8a9e"                    // açık Unicode kod noktaları
"\U000065e5\U0000672c\U00008a9e"        // açık Unicode kod noktaları
"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"  // açık UTF-8 baytları
```

Kaynak kodu bir karakteri, örneğin bir aksan ve bir harf içeren birleşik bir form gibi, iki kod noktası olarak temsil ediyorsa; rune sabitine yerleştirildiğinde sonuç hata olur (tek bir kod noktası değildir) ve string sabitine yerleştirildiğinde iki kod noktası olarak görünür.

---

## 🧾 Sabitler¶

Boolean sabitler, rune sabitleri, tamsayı sabitleri, kayan nokta sabitleri, kompleks sabitler ve string sabitleri vardır. Rune, tamsayı, kayan nokta ve kompleks sabitlere topluca sayısal sabitler denir.

Bir sabit değer; bir rune, tamsayı, kayan nokta, sanal veya string sabitiyle; bir sabiti gösteren bir tanımlayıcıyla; bir sabit ifadeyle; sonucu sabit olan bir dönüşümle; ya da `min` veya `max` gibi yerleşik fonksiyonların sabit argümanlara uygulanmasının sonucu ile, belirli değerlere `unsafe.Sizeof` uygulanmasıyla, bazı ifadelere `cap` veya `len` uygulanmasıyla, bir kompleks sabite `real` ve `imag` uygulanmasıyla ve sayısal sabitlere `complex` uygulanmasıyla temsil edilir. Boolean doğruluk değerleri, önceden tanımlı `true` ve `false` sabitleriyle temsil edilir. Önceden tanımlı `iota` tanımlayıcısı bir tamsayı sabitini belirtir.

Genel olarak, kompleks sabitler bir sabit ifade biçimidir ve o bölümde ele alınır.

Sayısal sabitler, keyfi hassasiyette kesin değerleri temsil eder ve taşma yapmaz. Sonuç olarak, IEEE 754 negatif sıfır, sonsuzluk ve sayı-değil (NaN) değerlerini gösteren sabitler yoktur.

Sabitler tipli veya tipsiz olabilir. Literal sabitler, `true`, `false`, `iota` ve yalnızca tipsiz sabit operandları içeren belirli sabit ifadeler tipsizdir.

Bir sabite tür, bir sabit bildirimi veya dönüşüm ile açıkça; ya da bir değişken bildiriminde, bir atama ifadesinde veya bir ifadenin operandı olarak kullanıldığında örtük biçimde verilebilir. Sabit değer ilgili türün değeri olarak temsil edilemiyorsa hata olur. Tür bir tür parametresi ise, sabit tür parametresinin sabit olmayan bir değerine dönüştürülür.

Tipsiz bir sabitin varsayılan türü, tipli bir değerin gerektiği bağlamlarda sabitin örtük olarak dönüştürüldüğü türdür; örneğin açık tür olmayan `i := 0` gibi kısa değişken bildiriminde. Tipsiz bir sabitin varsayılan türü, sırasıyla boolean, rune, tamsayı, kayan nokta, kompleks veya string sabit olmasına bağlı olarak `bool`, `rune`, `int`, `float64`, `complex128` veya `string`’dir.

Uygulama kısıtlaması: Sayısal sabitler dilde keyfi hassasiyete sahip olsa da, bir derleyici onları sınırlı hassasiyette bir iç gösterimle uygulayabilir. Bununla birlikte, her uygulama şunları yapmak zorundadır:

* Tamsayı sabitlerini en az 256 bit ile temsil etmek.
* Kayan nokta sabitlerini (kompleks sabitin kısımları dahil) en az 256 bit mantisa ve en az 16 bit işaretli ikilik üs ile temsil etmek.
* Bir tamsayı sabitini kesin olarak temsil edemiyorsa hata vermek.
* Taşma nedeniyle bir kayan nokta veya kompleks sabiti temsil edemiyorsa hata vermek.
* Hassasiyet sınırları nedeniyle bir kayan nokta veya kompleks sabiti temsil edemiyorsa, en yakın temsil edilebilir sabite yuvarlamak.

Bu gereksinimler hem literal sabitlere hem de sabit ifadelerin değerlendirilmesinin sonucuna uygulanır.

---

## 📦 Değişkenler¶

Bir değişken, bir değeri tutmak için kullanılan bir saklama konumudur. İzin verilen değerler kümesi, değişkenin türü tarafından belirlenir.

Bir değişken bildirimi veya (fonksiyon parametreleri ve sonuçları için) bir fonksiyon bildirimi ya da fonksiyon literal’inin imzası, adlandırılmış bir değişken için saklama alanı ayırır. Yerleşik `new` fonksiyonunun çağrılması veya bir bileşik literal’in adresinin alınması, çalışma zamanında bir değişken için saklama alanı ayırır. Böyle bir anonim değişkene (muhtemelen örtük) bir pointer dolaylı erişimi üzerinden başvurulur.

Dizi (*array*), dilim (*slice*) ve yapı (*struct*) türlerinden oluşan yapılandırılmış değişkenlerin, tek tek adreslenebilen elemanları ve alanları vardır. Her böyle eleman bir değişken gibi davranır.

Bir değişkenin statik türü (veya kısaca türü), bildiriminde verilen tür; `new` çağrısında veya bileşik literal’de sağlanan tür; ya da yapılandırılmış bir değişkenin bir elemanının türüdür. *Interface* türündeki değişkenlerin ayrıca farklı bir dinamik türü vardır; bu, çalışma zamanında değişkene atanan değerin (interface olmayan) türüdür (değer önceden tanımlı `nil` tanımlayıcısı olmadıkça; `nil`’in türü yoktur). Dinamik tür yürütme sırasında değişebilir; ancak interface değişkenlerinde saklanan değerler her zaman değişkenin statik türüne atanabilirdir.

```go
var x interface{}  // x nil'dir ve statik türü interface{}'dir
var v *T           // v'nin değeri nil'dir, statik türü *T'dir
x = 42             // x'in değeri 42'dir ve dinamik türü int'tir
x = v              // x'in değeri (*T)(nil)'dir ve dinamik türü *T'dir
```

Bir değişkenin değeri, bir ifadede değişkene başvurularak alınır; bu, değişkene en son atanan değerdir. Bir değişkene henüz bir değer atanmamışsa, değeri türü için sıfır değerdir.


## 🧬 Türler¶

Bir *tür (type)*, belirli değerlere özgü işlemler ve yöntemlerle birlikte bir değerler kümesini belirler. Bir tür, eğer bir *tür adı (type name)* varsa onunla gösterilebilir; tür jenerikse, tür adını *tür argümanları (type arguments)* izlemelidir. Bir tür ayrıca mevcut türlerden bir tür oluşturan bir *tür literal’i (type literal)* kullanılarak da belirtilebilir. 

```ebnf
Type     = TypeName [ TypeArgs ] | TypeLit | "(" Type ")" .
TypeName = identifier | QualifiedIdent .
TypeArgs = "[" TypeList [ "," ] "]" .
TypeList = Type { "," Type } .
TypeLit  = ArrayType | StructType | PointerType | FunctionType | InterfaceType |
           SliceType | MapType | ChannelType .
```

Dil, bazı tür adlarını önceden tanımlar. Diğerleri, tür bildirimleri veya tür parametresi listeleri ile tanıtılır. Bileşik türler—*array*, *struct*, *pointer*, *function*, *interface*, *slice*, *map* ve *channel* türleri—tür literal’leri kullanılarak oluşturulabilir.

Önceden tanımlı türler, tanımlı türler ve tür parametrelerine *adlandırılmış türler (named types)* denir. Bir takma ad (alias), eğer takma ad bildiriminde verilen tür bir adlandırılmış türse, bir adlandırılmış türü belirtir.

---

## ✅ Boolean türleri¶

Bir *boolean* türü, önceden tanımlı `true` ve `false` sabitleriyle gösterilen Boolean doğruluk değerleri kümesini temsil eder. Önceden tanımlı boolean türü `bool`’dur; bu, tanımlı bir türdür.

---

## 🔢 Sayısal türler¶

Bir *tamsayı (integer)*, *kayan nokta (floating-point)* veya *kompleks (complex)* tür, sırasıyla tamsayı, kayan nokta veya kompleks değerler kümesini temsil eder. Topluca *sayısal türler (numeric types)* olarak adlandırılırlar. Önceden tanımlı, mimariden bağımsız sayısal türler şunlardır:

```text
uint8       tüm işaretsiz  8-bit tamsayılar kümesi (0 ile 255)
uint16      tüm işaretsiz 16-bit tamsayılar kümesi (0 ile 65535)
uint32      tüm işaretsiz 32-bit tamsayılar kümesi (0 ile 4294967295)
uint64      tüm işaretsiz 64-bit tamsayılar kümesi (0 ile 18446744073709551615)

int8        tüm işaretli  8-bit tamsayılar kümesi (-128 ile 127)
int16       tüm işaretli 16-bit tamsayılar kümesi (-32768 ile 32767)
int32       tüm işaretli 32-bit tamsayılar kümesi (-2147483648 ile 2147483647)
int64       tüm işaretli 64-bit tamsayılar kümesi (-9223372036854775808 ile 9223372036854775807)

float32     tüm IEEE 754 32-bit kayan nokta sayıları kümesi
float64     tüm IEEE 754 64-bit kayan nokta sayıları kümesi

complex64   gerçek ve sanal kısımları float32 olan tüm kompleks sayılar kümesi
complex128  gerçek ve sanal kısımları float64 olan tüm kompleks sayılar kümesi

byte        uint8 için takma ad
rune        int32 için takma ad
```

*n*-bit bir tamsayının değeri *n* bit genişliğindedir ve *two's complement* aritmetiği kullanılarak temsil edilir.

Ayrıca, uygulamaya özgü boyutlara sahip önceden tanımlı bir tamsayı türleri kümesi de vardır:

```text
uint     32 ya da 64 bit
int      uint ile aynı boyut
uintptr  bir pointer değerinin yorumlanmamış bitlerini saklayacak kadar büyük bir işaretsiz tamsayı
```

Taşınabilirlik sorunlarından kaçınmak için, tüm sayısal türler tanımlı türlerdir ve bu nedenle `byte`’ın `uint8` için bir takma ad olması ve `rune`’un `int32` için bir takma ad olması dışında birbirinden ayrıdır. Farklı sayısal türler bir ifadede veya atamada karıştırıldığında açık dönüşümler gerekir. Örneğin `int32` ve `int`, belirli bir mimaride aynı boyutta olsalar bile aynı tür değildir.

---

## 🧵 String türleri¶

Bir *string* türü, string değerleri kümesini temsil eder. Bir string değeri, (muhtemelen boş) bir bayt dizisidir. Bayt sayısına string’in uzunluğu (*length*) denir ve asla negatif değildir. String’ler değiştirilemez (*immutable*): bir kez oluşturulduktan sonra, içeriğini değiştirmek imkânsızdır. Önceden tanımlı string türü `string`’dir; bu, tanımlı bir türdür.

Bir string `s`’nin uzunluğu yerleşik `len` fonksiyonu kullanılarak bulunabilir. String bir sabitse, uzunluğu derleme zamanı sabitidir. Bir string’in baytlarına, `0` ile `len(s)-1` arasındaki tamsayı indeksleriyle erişilebilir. Böyle bir elemanın adresini almak yasadışıdır; `s[i]` string’in *i*’inci baytı ise, `&s[i]` geçersizdir.

---

## 🧱 Dizi (Array) türleri¶

Bir *array*, eleman türü denilen tek bir türden elemanların numaralandırılmış bir dizisidir. Eleman sayısına array’in uzunluğu (*length*) denir ve asla negatif değildir.

```ebnf
ArrayType   = "[" ArrayLength "]" ElementType .
ArrayLength = Expression .
ElementType = Type .
```

Uzunluk, array’in türünün bir parçasıdır; türü `int` olan bir değerle temsil edilebilen, negatif olmayan bir sabite değerlendirilmelidir. Array `a`’nın uzunluğu yerleşik `len` fonksiyonuyla bulunabilir. Elemanlara `0` ile `len(a)-1` arasındaki tamsayı indeksleriyle adreslenebilir. Array türleri her zaman tek boyutludur ancak çok boyutlu türler oluşturmak için birleştirilebilir.

```text
[32]byte
[2*N] struct { x, y int32 }
[1000]*float64
[3][5]int
[2][2][2]float64  // [2]([2]([2]float64)) ile aynı
```

Bir array türü `T`, yalnızca array veya struct türleri olan kapsayan türler aracılığıyla doğrudan veya dolaylı biçimde, eleman türü olarak `T` türüne veya bileşen olarak `T` içeren bir türe sahip olamaz.

```go
// geçersiz array türleri
type (
	T1 [10]T1                 // T1'in eleman türü T1
	T2 [10]struct{ f T2 }     // T2, bir struct bileşeni olarak T2 içerir
	T3 [10]T4                 // T3, T4 içindeki bir struct bileşeni olarak T3 içerir
	T4 struct{ f T3 }         // T4, bir struct içinde array T3'ün bileşeni olarak T4 içerir
)

// geçerli array türleri
type (
	T5 [10]*T5                // T5, bir pointer bileşeni olarak T5 içerir
	T6 [10]func() T6          // T6, bir function türü bileşeni olarak T6 içerir
	T7 [10]struct{ f []T7 }   // T7, bir struct içinde slice bileşeni olarak T7 içerir
)
```

---

## 🧩 Dilim (Slice) türleri¶

Bir *slice*, alttaki (*underlying*) bir array’in bitişik bir parçası için bir tanımlayıcıdır ve o array’den numaralandırılmış bir eleman dizisine erişim sağlar. Bir slice türü, eleman türünün array’lerinden tüm slice’ların kümesini belirtir. Eleman sayısına slice’ın uzunluğu (*length*) denir ve asla negatif değildir. Başlatılmamış bir slice’ın değeri `nil`’dir.

```ebnf
SliceType = "[" "]" ElementType .
```

Slice `s`’nin uzunluğu yerleşik `len` fonksiyonu ile bulunabilir; array’lerden farklı olarak yürütme sırasında değişebilir. Elemanlara `0` ile `len(s)-1` arasındaki tamsayı indeksleriyle adreslenebilir. Belirli bir elemanın slice indeksi, alttaki array’deki aynı elemanın indeksinden daha küçük olabilir.

Bir slice, bir kez başlatıldıktan sonra, elemanlarını tutan bir alttaki array ile her zaman ilişkilidir. Dolayısıyla bir slice, saklamayı array’i ve aynı array’in diğer slice’ları ile paylaşır; buna karşılık ayrı array’ler her zaman ayrı saklama alanını temsil eder.

Bir slice’ın altındaki array, slice’ın sonunu aşacak şekilde uzayabilir. *Capacity*, bu uzantının bir ölçüsüdür: slice uzunluğu ile slice’ın ötesindeki array uzunluğunun toplamıdır; bu kapasiteye kadar olan uzunlukta bir slice, orijinal slice’dan yeni bir slice kesilerek oluşturulabilir. Slice `a`’nın kapasitesi yerleşik `cap(a)` fonksiyonu ile bulunabilir.

Belirli bir eleman türü `T` için yeni, başlatılmış bir slice değeri yerleşik `make` fonksiyonuyla yapılabilir; `make`, bir slice türü ve uzunluğu ve isteğe bağlı olarak kapasiteyi belirten parametreleri alır. `make` ile oluşturulan bir slice her zaman yeni, gizli bir array tahsis eder ve dönen slice değeri o array’e referans verir. Yani, şu ifade:

```text
make([]T, length, capacity)
```

bir array tahsis edip onu dilimlemekle aynı slice’ı üretir; dolayısıyla şu iki ifade eşdeğerdir:

```text
make([]int, 50, 100)
new([100]int)[0:50]
```

Array’ler gibi, slice’lar da her zaman tek boyutludur ancak daha yüksek boyutlu nesneler oluşturmak için birleştirilebilir. Array’lerin array’leri ile, iç array’ler yapısal olarak her zaman aynı uzunluktadır; ancak slice’ların slice’ları (veya slice’ların array’leri) ile iç uzunluklar dinamik olarak değişebilir. Ayrıca iç slice’lar ayrı ayrı başlatılmalıdır.

---

## 🏗️ Yapı (Struct) türleri¶

Bir *struct*, her birinin bir adı ve türü olan, alan (*field*) denilen adlandırılmış elemanların bir dizisidir. Alan adları açıkça (*IdentifierList*) veya örtük olarak (*EmbeddedField*) belirtilebilir. Bir struct içinde, boş olmayan alan adları benzersiz olmalıdır.

```ebnf
StructType    = "struct" "{" { FieldDecl ";" } "}" .
FieldDecl     = (IdentifierList Type | EmbeddedField) [ Tag ] .
EmbeddedField = [ "*" ] TypeName [ TypeArgs ] .
Tag           = string_lit .
```

```go
// Boş bir struct.
struct {}

// 6 alanlı bir struct.
struct {
	x, y int
	u float32
	_ float32  // padding
	A *[]int
	F func()
}
```

Bir türle bildirilip açık bir alan adı olmayan alana *embedded field* denir. Bir embedded field, bir tür adı `T` veya interface olmayan bir tür adı `*T`’ye pointer olarak belirtilmelidir ve `T`’nin kendisi bir pointer türü veya tür parametresi olamaz. Niteliksiz tür adı, alan adı olarak davranır.

```go
// Türleri T1, *T2, P.T3 ve *P.T4 olan dört embedded field içeren bir struct
struct {
	T1        // alan adı T1'dir
	*T2       // alan adı T2'dir
	P.T3      // alan adı T3'tür
	*P.T4     // alan adı T4'tür
	x, y int  // alan adları x ve y'dir
}
```

Aşağıdaki bildirim yasadışıdır çünkü alan adları bir struct türünde benzersiz olmalıdır:

```go
struct {
	T     // embedded field *T ve *P.T ile çakışır
	*T    // embedded field T ve *P.T ile çakışır
	*P.T  // embedded field T ve *T ile çakışır
}
```

Bir struct `x` içinde bir embedded field’in `f` alanı veya yöntemi, `x.f` yasal bir seçici (*selector*) olup o alanı veya yöntemi `f` gösteriyorsa, *promoted* olarak adlandırılır.

Promoted alanlar, bir struct’ın sıradan alanları gibi davranır; ancak struct’ın bileşik literal’lerinde alan adı olarak kullanılamazlar.

Bir struct türü `S` ve bir tür adı `T` verildiğinde, promoted yöntemler struct’ın *method set*’ine aşağıdaki gibi dahil edilir:

* `S`, `T` türünde bir embedded field içeriyorsa, `S` ve `*S`’nin method set’leri alıcı (*receiver*) türü `T` olan promoted yöntemleri içerir. `*S`’nin method set’i ayrıca alıcı türü `*T` olan promoted yöntemleri de içerir.
* `S`, `*T` türünde bir embedded field içeriyorsa, `S` ve `*S`’nin method set’leri alıcı türü `T` veya `*T` olan promoted yöntemleri içerir.

Bir alan bildiriminin ardından isteğe bağlı bir string literal etiket (*tag*) gelebilir; bu, ilgili alan bildirimindeki tüm alanlar için bir öznitelik olur. Boş bir etiket string’i, etiketin olmamasıyla eşdeğerdir. Etiketler bir yansıma (*reflection*) arayüzü üzerinden görünürdür ve struct’lar için tür kimliğine (*type identity*) dahil olur, ancak bunun dışında yok sayılır.

```go
struct {
	x, y float64 ""  // boş etiket string'i, etiketin olmaması gibidir
	name string  "any string is permitted as a tag"
	_    [4]byte "ceci n'est pas un champ de structure"
}

// Bir TimeStamp protocol buffer'ına karşılık gelen bir struct.
// Etiket string'leri protocol buffer alan numaralarını tanımlar;
// reflect paketinde özetlenen geleneği izlerler.
struct {
	microsec  uint64 `protobuf:"1"`
	serverIP6 uint64 `protobuf:"2"`
}
```

Bir struct türü `T`, yalnızca array veya struct türleri olan kapsayan türler aracılığıyla doğrudan veya dolaylı biçimde, türü `T` olan bir alan veya bileşen olarak `T` içeren bir tür barındıramaz.

```go
// geçersiz struct türleri
type (
	T1 struct{ T1 }            // T1, T1 türünde bir alan içerir
	T2 struct{ f [10]T2 }      // T2, bir array bileşeni olarak T2 içerir
	T3 struct{ T4 }            // T3, T4 içindeki bir array bileşeni olarak T3 içerir
	T4 struct{ f [10]T3 }      // T4, bir array içinde struct T3'ün bileşeni olarak T4 içerir
)

// geçerli struct türleri
type (
	T5 struct{ f *T5 }         // T5, bir pointer bileşeni olarak T5 içerir
	T6 struct{ f func() T6 }   // T6, bir function türü bileşeni olarak T6 içerir
	T7 struct{ f [10][]T7 }    // T7, bir array içinde slice bileşeni olarak T7 içerir
)
```

---

## 📌 Pointer türleri¶

Bir pointer türü, temel tür (*base type*) denilen belirli bir türdeki değişkenlere yönelik tüm pointer’ların kümesini belirtir. Başlatılmamış bir pointer’ın değeri `nil`’dir.

```ebnf
PointerType = "*" BaseType .
BaseType    = Type .
```

```text
*Point
*[4]int
```

---

## 🧰 Function türleri¶

Bir function türü, aynı parametre ve sonuç türlerine sahip tüm function’ların kümesini belirtir. Function türünde başlatılmamış bir değişkenin değeri `nil`’dir.

```ebnf
FunctionType  = "func" Signature .
Signature     = Parameters [ Result ] .
Result        = Parameters | Type .
Parameters    = "(" [ ParameterList [ "," ] ] ")" .
ParameterList = ParameterDecl { "," ParameterDecl } .
ParameterDecl = [ IdentifierList ] [ "..." ] Type .
```

Parametreler veya sonuçlar listesinde, adlar (*IdentifierList*) ya tamamen bulunmalı ya da tamamen bulunmamalıdır. Bulunuyorsa, her ad belirtilen türden bir öğeyi (parametre veya sonuç) temsil eder ve imzadaki tüm boş olmayan adlar benzersiz olmalıdır. Bulunmuyorsa, her tür o türden bir öğeyi temsil eder. Parametre ve sonuç listeleri her zaman parantez içindedir; ancak tam olarak bir tane adsız sonuç varsa, parantezsiz bir tür olarak yazılabilir.

Bir function imzasındaki son gelen parametre, `...` ile öneklenen bir türe sahip olabilir. Böyle bir parametreye sahip bir function’a *variadic* denir ve bu parametre için sıfır veya daha fazla argümanla çağrılabilir.

```go
func()
func(x int) int
func(a, _ int, z float32) bool
func(a, b int, z float32) (bool)
func(prefix string, values ...int)
func(a, b int, z float64, opt ...interface{}) (success bool)
func(int, int, float64) (float64, *[]int)
func(n int) func(p *T)
```

---

## 🧩 Interface türleri¶

Bir interface türü bir *type set* tanımlar. Interface türündeki bir değişken, interface’in type set’inde bulunan herhangi bir türden bir değer saklayabilir. Böyle bir türün interface’i *implement* ettiği söylenir. Başlatılmamış bir interface değişkeninin değeri `nil`’dir.

```ebnf
InterfaceType  = "interface" "{" { InterfaceElem ";" } "}" .
InterfaceElem  = MethodElem | TypeElem .
MethodElem     = MethodName Signature .
MethodName     = identifier .
TypeElem       = TypeTerm { "|" TypeTerm } .
TypeTerm       = Type | UnderlyingType .
UnderlyingType = "~" Type .
```

Bir interface türü, bir interface elemanları listesiyle belirtilir. Bir interface elemanı ya bir yöntemdir ya da bir tür elemanıdır; tür elemanı, bir veya daha fazla tür teriminin birleşimidir. Bir tür terimi, ya tek bir türdür ya da tek bir alttaki (*underlying*) türdür.

---

## 🧩 Temel interface’ler¶

En temel biçimde bir interface, (muhtemelen boş) bir yöntemler listesi belirtir. Böyle bir interface’in tanımladığı type set, bu yöntemlerin tümünü *implement* eden türler kümesidir ve karşılık gelen method set, tam olarak interface’de belirtilen yöntemlerden oluşur. Type set’leri tamamen bir yöntemler listesiyle tanımlanabilen interface’lere *basic interface* denir.

```go
// Basit bir File interface'i.
interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	Close() error
}
```

Açıkça belirtilen her yöntemin adı benzersiz olmalı ve boş olmamalıdır.

```go
interface {
	String() string
	String() string  // yasadışı: String benzersiz değil
	_(x int)         // yasadışı: yöntem boş olmayan bir ada sahip olmalı
}
```

Bir interface’i birden fazla tür implement edebilir. Örneğin, iki tür `S1` ve `S2` aşağıdaki method set’e sahipse

```go
func (p T) Read(p []byte) (n int, err error)
func (p T) Write(p []byte) (n int, err error)
func (p T) Close() error
```

(`T`, `S1` veya `S2`’den birini temsil ederken) `S1` ve `S2`, başka hangi yöntemlere sahip olurlarsa olsunlar ya da hangi yöntemleri paylaşsalar da `File` interface’ini implement eder.

Bir interface’in type set’inin üyesi olan her tür, o interface’i implement eder. Herhangi bir tür, birkaç farklı interface’i implement edebilir. Örneğin, tüm türler, tüm (interface olmayan) türler kümesini temsil eden boş interface’i implement eder:

```go
interface{}
```

Kolaylık için, önceden tanımlı `any` türü, boş interface için bir takma addır. [Go 1.18]

Benzer şekilde, bir tür bildiriminin içinde `Locker` adlı bir interface tanımlamak için görünen şu interface belirtimini düşünün:

```go
type Locker interface {
	Lock()
	Unlock()
}
```

Eğer `S1` ve `S2` ayrıca şunları da implement ediyorsa

```go
func (p T) Lock() { … }
func (p T) Unlock() { … }
```

`Locker` interface’ini de (ve `File` interface’ini de) implement ederler.

---

## 🧩 Embedded interface’ler¶

Biraz daha genel bir biçimde, bir interface `T`, bir interface elemanı olarak (muhtemelen nitelikli) bir interface tür adı `E` kullanabilir. Buna, `E` interface’inin `T` içine gömülmesi (*embedding*) denir. [Go 1.14] `T`’nin type set’i, `T`’nin açıkça bildirilen yöntemleriyle tanımlanan type set’ler ile `T`’nin embedded interface’lerinin type set’lerinin kesişimidir. Başka bir deyişle, `T`’nin type set’i, `T`’nin açıkça bildirilen tüm yöntemlerini ve ayrıca `E`’nin tüm yöntemlerini implement eden tüm türlerin kümesidir. [Go 1.18]

```go
type Reader interface {
	Read(p []byte) (n int, err error)
	Close() error
}

type Writer interface {
	Write(p []byte) (n int, err error)
	Close() error
}

// ReadWriter'ın yöntemleri Read, Write ve Close'tur.
type ReadWriter interface {
	Reader  // Reader'ın yöntemlerini ReadWriter'ın method set'ine dahil eder
	Writer  // Writer'ın yöntemlerini ReadWriter'ın method set'ine dahil eder
}
```

Interface’ler gömülürken, aynı ada sahip yöntemler özdeş imzalara sahip olmalıdır.

```go
type ReadCloser interface {
	Reader   // Reader'ın yöntemlerini ReadCloser'ın method set'ine dahil eder
	Close()  // yasadışı: Reader.Close ile Close imzaları farklı
}
```

---

## 🧩 Genel interface’ler¶

En genel biçimlerinde, bir interface elemanı aynı zamanda keyfi bir tür terimi `T`, alttaki türü `T` belirten `~T` biçiminde bir terim veya `t1|t2|…|tn` terimlerinden oluşan bir birleşim (*union*) olabilir. [Go 1.18] Yöntem belirtimleriyle birlikte, bu elemanlar bir interface’in type set’inin aşağıdaki gibi kesin tanımlanmasını sağlar:

* Boş interface’in type set’i, tüm interface olmayan türler kümesidir.
* Boş olmayan bir interface’in type set’i, interface elemanlarının type set’lerinin kesişimidir.
* Bir yöntem belirtiminin type set’i, method set’leri o yöntemi içeren tüm interface olmayan türler kümesidir.
* Interface olmayan bir tür teriminin type set’i, yalnızca o türden oluşan kümedir.
* `~T` biçimindeki bir terimin type set’i, alttaki türü `T` olan tüm türler kümesidir.
* `t1|t2|…|tn` birleşiminin type set’i, terimlerin type set’lerinin birleşimidir.

“tüm interface olmayan türler kümesi” nicemlemesi, yalnızca eldeki programda bildirilen tüm (interface olmayan) türleri değil, tüm olası programlarda tüm olası türleri ifade eder ve dolayısıyla sonsuzdur. Benzer şekilde, belirli bir yöntemi implement eden tüm interface olmayan türlerin kümesi verildiğinde, bu türlerin method set’lerinin kesişimi, eldeki programdaki tüm türler o yöntemi her zaman başka bir yöntemle eşlese bile, tam olarak o yöntemi içerecektir.

Yapı gereği, bir interface’in type set’i hiçbir zaman bir interface türü içermez.

```go
// Yalnızca int türünü temsil eden bir interface.
interface {
	int
}

// Alttaki türü int olan tüm türleri temsil eden bir interface.
interface {
	~int
}

// Alttaki türü int olan ve String yöntemini implement eden tüm türleri temsil eden bir interface.
interface {
	~int
	String() string
}

// Boş bir type set'i temsil eden bir interface: hem int hem string olan hiçbir tür yoktur.
interface {
	int
	string
}
```

`~T` biçimindeki bir terimde, `T`’nin alttaki türü kendisi olmalıdır ve `T` bir interface olamaz.

```go
type MyInt int

interface {
	~[]byte  // []byte'ın alttaki türü kendisidir
	~MyInt   // yasadışı: MyInt'in alttaki türü MyInt değildir
	~error   // yasadışı: error bir interface'tir
}
```

Union elemanlar type set’lerin birleşimlerini belirtir:

```go
// Float interface'i tüm kayan nokta türlerini temsil eder
// (alttaki türleri float32 veya float64 olan tüm adlandırılmış türler dahil).
type Float interface {
	~float32 | ~float64
}
```

`T` veya `~T` biçimindeki bir terimdeki `T` bir tür parametresi olamaz ve interface olmayan tüm terimlerin type set’leri ikişer ikişer ayrık olmalıdır (type set’lerin ikili kesişimi boş olmalıdır). Bir tür parametresi `P` verildiğinde:

```go
interface {
	P                // yasadışı: P bir tür parametresi
	int | ~P         // yasadışı: P bir tür parametresi
	~int | MyInt     // yasadışı: ~int ve MyInt için type set'ler ayrık değil (~int, MyInt'i içerir)
	float32 | Float  // çakışan type set'ler ama Float bir interface
}
```

Uygulama kısıtlaması: (Birden fazla terimli) bir union, önceden tanımlı `comparable` tanımlayıcısını veya yöntem belirten interface’leri ya da `comparable`’ı veya yöntem belirten interface’leri gömemez.

*Basic* olmayan interface’ler yalnızca tür kısıtı (*type constraint*) olarak veya kısıt olarak kullanılan diğer interface’lerin elemanı olarak kullanılabilir. Değerlerin veya değişkenlerin türü olamazlar ya da başka, interface olmayan türlerin bileşeni olamazlar.

```go
var x Float                     // yasadışı: Float bir basic interface değil

var x interface{} = Float(nil)  // yasadışı

type Floatish struct {
	f Float                 // yasadışı
}
```

Bir interface türü `T`, doğrudan veya dolaylı biçimde, `T` olan, `T` içeren veya `T`’yi gömen bir tür elemanını gömemez.

```go
// yasadışı: Bad kendisini gömemez
type Bad interface {
	Bad
}

// yasadışı: Bad1, Bad2 kullanarak kendisini gömemez
type Bad1 interface {
	Bad2
}
type Bad2 interface {
	Bad1
}

// yasadışı: Bad3, Bad3 içeren bir union'ı gömemez
type Bad3 interface {
	~int | ~string | Bad3
}

// yasadışı: Bad4, eleman türü olarak Bad4 içeren bir array'i gömemez
type Bad4 interface {
	[10]Bad4
}
```

### 🧩 Bir interface’i implement etmek¶

Bir tür `T`, bir interface `I`’yi şu durumlarda implement eder:

* `T` bir interface değildir ve `I`’nin type set’inin bir elemanıdır; veya
* `T` bir interface’tir ve `T`’nin type set’i `I`’nin type set’inin bir alt kümesidir.

Türü `T` olan bir değer, eğer `T` o interface’i implement ediyorsa, interface’i implement eder.

---

## 🗺️ Map türleri¶

Bir *map*, başka bir türün benzersiz anahtarları kümesiyle indekslenen, tek bir türden elemanların sırasız bir grubudur. Anahtarların türüne *key type*, elemanların türüne *element type* denir. Başlatılmamış bir map’in değeri `nil`’dir.

```ebnf
MapType = "map" "[" KeyType "]" ElementType .
KeyType = Type .
```

Anahtar türündeki operandlar için `==` ve `!=` karşılaştırma işleçleri tamamen tanımlı olmalıdır; bu nedenle anahtar türü bir function, map veya slice olamaz. Anahtar türü bir interface türüyse, bu karşılaştırma işleçleri dinamik anahtar değerleri için tanımlı olmalıdır; aksi hâlde çalışma zamanı *panic* oluşur.

```go
map[string]int
map[*T]struct{ x, y float64 }
map[string]interface{}
```

Map elemanlarının sayısına uzunluk (*length*) denir. Bir map `m` için, bu `len` yerleşik fonksiyonu ile bulunabilir ve yürütme sırasında değişebilir. Elemanlar atamalarla yürütme sırasında eklenebilir ve indeks ifadeleriyle alınabilir; `delete` ve `clear` yerleşik fonksiyonlarıyla kaldırılabilir.

Yeni, boş bir map değeri `make` yerleşik fonksiyonuyla yapılır; `make`, map türünü ve isteğe bağlı bir kapasite ipucunu argüman olarak alır:

```go
make(map[string]int)
make(map[string]int, 100)
```

Başlangıç kapasitesi boyutunu sınırlamaz: map’ler, depolanan öğe sayısını karşılayacak şekilde büyür; `nil` map’ler istisnadır. `nil` map, boş bir map’e eşdeğerdir; ancak ona eleman eklenemez.

---

## 📡 Channel türleri¶

Bir *channel*, eşzamanlı çalışan function’ların, belirtilen eleman türündeki değerleri gönderip alarak iletişim kurması için bir mekanizma sağlar. Başlatılmamış bir channel’ın değeri `nil`’dir.

```ebnf
ChannelType = ( "chan" | "chan" "<-" | "<-" "chan" ) ElementType .
```

İsteğe bağlı `<-` işleci channel yönünü belirtir: gönderme (*send*) veya alma (*receive*). Yön verilirse channel yönlüdür; aksi hâlde çift yönlüdür. Bir channel, atama veya açık dönüşüm ile yalnızca gönderme ya da yalnızca alma ile sınırlandırılabilir.

```go
chan T          // T türündeki değerleri göndermek ve almak için kullanılabilir
chan<- float64  // yalnızca float64 gönderilebilir
<-chan int      // yalnızca int alınabilir
```

`<-` işleci, mümkün olan en soldaki `chan` ile bağlanır:

```go
chan<- chan int    // chan<- (chan int) ile aynı
chan<- <-chan int  // chan<- (<-chan int) ile aynı
<-chan <-chan int  // <-chan (<-chan int) ile aynı
chan (<-chan int)
```

Yeni, başlatılmış bir channel değeri `make` yerleşik fonksiyonuyla yapılabilir; `make`, channel türünü ve isteğe bağlı kapasiteyi argüman olarak alır:

```go
make(chan int, 100)
```

Kapasite, eleman sayısı olarak channel’daki tamponun boyutunu belirler. Kapasite sıfırsa veya yoksa, channel *unbuffered*’dır ve iletişim yalnızca hem gönderen hem alan hazır olduğunda başarılı olur. Aksi hâlde channel *buffered*’dır ve tampon dolu değilse (gönderimler) veya boş değilse (alımlar) bloklamadan başarılı olur. `nil` bir channel iletişim için asla hazır değildir.

Bir channel, `close` yerleşik fonksiyonuyla kapatılabilir. Alım işleci için çok değerli atama biçimi, alınan bir değerin channel kapanmadan önce gönderilip gönderilmediğini raporlar.

Tek bir channel, daha fazla eşzamanlama olmaksızın, herhangi sayıda goroutine tarafından gönderme ifadelerinde, alım işlemlerinde ve `cap` ile `len` yerleşik fonksiyonlarına çağrılarda kullanılabilir. Channel’lar *first-in-first-out* kuyruklar gibi davranır. Örneğin, bir goroutine bir channel’a değerler gönderir ve ikinci bir goroutine bunları alırsa, değerler gönderildikleri sırayla alınır.

---

## 🧩 Türlerin ve değerlerin özellikleri¶

### 🧱 Değerlerin temsili¶

Önceden tanımlı türlerin (aşağıda `any` ve `error` interface’leri için bakın), array’lerin ve struct’ların değerleri kendi kendine yeterlidir: Her böyle değer, tüm verisinin eksiksiz bir kopyasını içerir ve bu türlerdeki değişkenler tüm değeri saklar. Örneğin, bir array değişkeni, array’in tüm elemanları için saklamayı (değişkenleri) sağlar. İlgili sıfır değerler, değerin türlerine özgüdür; asla `nil` değildir.

`nil` olmayan pointer, function, slice, map ve channel değerleri, birden çok değer tarafından paylaşılabilecek alttaki verilere referanslar içerir:

* Bir pointer değeri, pointer temel türü değerini tutan değişkene bir referanstır.
* Bir function değeri, (muhtemelen anonim) function’a ve kapalı değişkenlere referanslar içerir.
* Bir slice değeri, slice uzunluğunu, kapasitesini ve alttaki array’ine bir referansı içerir.
* Bir map veya channel değeri, map veya channel’ın uygulamaya özgü veri yapısına bir referanstır.
* Bir interface değeri, interface’in dinamik türüne bağlı olarak kendi kendine yeterli olabilir veya alttaki verilere referanslar içerebilir. Önceden tanımlı `nil` tanımlayıcısı, değerleri referans içerebilen türler için sıfır değerdir.

Birden çok değer alttaki veriyi paylaştığında, bir değerin değiştirilmesi başka bir değeri değiştirebilir. Örneğin, bir slice’ın bir elemanını değiştirmek, aynı array’i paylaşan tüm slice’lar için alttaki array’deki o elemanı değiştirir.

---

## 🧬 Alttaki türler¶

Her tür `T`’nin bir alttaki türü (*underlying type*) vardır: Eğer `T`, önceden tanımlı boolean, sayısal veya string türlerinden biriyse ya da bir tür literal’iyse, karşılık gelen alttaki tür `T`’nin kendisidir. Aksi hâlde, `T`’nin alttaki türü, `T`’nin bildiriminde başvurduğu türün alttaki türüdür. Bir tür parametresi için, alttaki tür, tür kısıtının alttaki türüdür; bu her zaman bir interface’tir.

```go
type (
	A1 = string
	A2 = A1
)

type (
	B1 string
	B2 B1
	B3 []B1
	B4 B3
)

func f[P any](x P) { … }
```

`string`, `A1`, `A2`, `B1` ve `B2`’nin alttaki türü `string`’dir. `[]B1`, `B3` ve `B4`’ün alttaki türü `[]B1`’dir. `P`’nin alttaki türü `interface{}`’dir.

---

## 🆔 Tür kimliği¶

İki tür ya özdeştir (“aynıdır”) ya da farklıdır.

Adlandırılmış bir tür her zaman herhangi bir başka türden farklıdır. Aksi hâlde, iki tür, alttaki tür literal’leri yapısal olarak eşdeğer ise özdeştir; yani aynı literal yapıya sahiptirler ve karşılık gelen bileşenler özdeş türlere sahiptir. Ayrıntıda:

* İki array türü, eleman türleri özdeşse ve array uzunlukları aynıysa özdeştir.
* İki slice türü, eleman türleri özdeşse özdeştir.
* İki struct türü, alan dizileri aynıysa ve karşılık gelen alan çiftlerinin adları, türleri ve etiketleri özdeşse ve her ikisi de ya embedded ya da embedded değilse özdeştir. Farklı paketlerden dışa aktarılmamış (*non-exported*) alan adları her zaman farklıdır.
* İki pointer türü, temel türleri özdeşse özdeştir.
* İki function türü, parametre ve sonuç sayıları aynıysa, karşılık gelen parametre ve sonuç türleri özdeşse ve her iki function ya variadic ya da değilse özdeştir. Parametre ve sonuç adlarının eşleşmesi gerekmez.
* İki interface türü, aynı type set’i tanımlıyorsa özdeştir.
* İki map türü, key ve element türleri özdeşse özdeştir.
* İki channel türü, eleman türleri özdeşse ve yönleri aynıysa özdeştir.
* İki instantiate edilmiş tür, tanımlı türleri ve tüm tür argümanları özdeşse özdeştir.

Aşağıdaki bildirimler verildiğinde:

```go
type (
	A0 = []string
	A1 = A0
	A2 = struct{ a, b int }
	A3 = int
	A4 = func(A3, float64) *A0
	A5 = func(x int, _ float64) *[]string

	B0 A0
	B1 []string
	B2 struct{ a, b int }
	B3 struct{ a, c int }
	B4 func(int, float64) *B0
	B5 func(x int, y float64) *A1

	C0 = B0
	D0[P1, P2 any] struct{ x P1; y P2 }
	E0 = D0[int, string]
)
```

şu türler özdeştir:

```text
A0, A1 ve []string
A2 ve struct{ a, b int }
A3 ve int
A4, func(int, float64) *[]string ve A5

B0 ve C0
D0[int, string] ve E0
[]int ve []int
struct{ a, b *B5 } ve struct{ a, b *B5 }
func(x int, y float64) *[]string, func(int, float64) (result *[]string) ve A5
```

`B0` ve `B1` farklıdır çünkü ayrı tür tanımlarıyla oluşturulan yeni türlerdir; `func(int, float64) *B0` ve `func(x int, y float64) *[]string` farklıdır çünkü `B0`, `[]string`’den farklıdır; ve `P1` ile `P2` farklıdır çünkü farklı tür parametreleridir. `D0[int, string]` ile `struct{ x int; y string }` farklıdır çünkü ilki instantiate edilmiş bir tanımlı türken ikincisi bir tür literal’idir (ancak yine de atanabilirdir).

---

## 🧷 Atanabilirlik¶

Türü `V` olan bir değer `x`, türü `T` olan bir değişkene atanabilirdir (“x, T’ye atanabilir”) eğer aşağıdaki koşullardan biri geçerliyse:

* `V` ve `T` özdeştir.
* `V` ve `T` özdeş alttaki türlere sahiptir ancak tür parametresi değildir ve `V` veya `T`’den en az biri adlandırılmış tür değildir.
* `V` ve `T` özdeş eleman türlerine sahip channel türleridir, `V` çift yönlü bir channel’dır ve `V` veya `T`’den en az biri adlandırılmış tür değildir.
* `T` bir interface türüdür, ancak bir tür parametresi değildir ve `x`, `T`’yi implement eder.
* `x`, önceden tanımlı `nil` tanımlayıcısıdır ve `T` bir pointer, function, slice, map, channel veya interface türüdür, ancak bir tür parametresi değildir.
* `x`, `T` türündeki bir değerle temsil edilebilen tipsiz bir sabittir.

Ek olarak, eğer `x`’in türü `V` veya `T` tür parametreleriyse, `x` aşağıdaki koşullardan biri geçerliyse `T` türündeki bir değişkene atanabilirdir:

* `x`, önceden tanımlı `nil` tanımlayıcısıdır, `T` bir tür parametresidir ve `x`, `T`’nin type set’indeki her türe atanabilirdir.
* `V` adlandırılmış bir tür değildir, `T` bir tür parametresidir ve `x`, `T`’nin type set’indeki her türe atanabilirdir.
* `V` bir tür parametresidir ve `T` adlandırılmış bir tür değildir ve `V`’nin type set’indeki her türden değerler `T`’ye atanabilirdir.

---

## 🧾 Temsil edilebilirlik¶

Bir sabit `x`, tür parametresi olmayan `T` türünde bir değerle temsil edilebilir (*representable*) eğer aşağıdaki koşullardan biri geçerliyse:

* `x`, `T` tarafından belirlenen değerler kümesindedir.
* `T` bir kayan nokta türüdür ve `x`, taşma olmaksızın `T`’nin hassasiyetine yuvarlanabilir. Yuvarlama, IEEE 754 *round-to-even* kurallarını kullanır, ancak IEEE negatif sıfır daha da basitleştirilerek işaretsiz sıfıra dönüştürülür. Not: sabit değerler hiçbir zaman IEEE negatif sıfır, NaN veya sonsuzluk üretmez.
* `T` bir kompleks türdür ve `x`’in bileşenleri `real(x)` ve `imag(x)`, `T`’nin bileşen türünün (`float32` veya `float64`) değerleriyle temsil edilebilirdir.

Eğer `T` bir tür parametresiyse, `x`, `T` türündeki bir değerle temsil edilebilir eğer `x`, `T`’nin type set’indeki her türün değeriyle temsil edilebiliyorsa.

```text
x                   T           x, T türündeki bir değerle temsil edilebilir çünkü

'a'                 byte        97 byte değerleri kümesindedir
97                  rune        rune, int32 için bir takma addır ve 97 32-bit tamsayılar kümesindedir
"foo"               string      "foo" string değerleri kümesindedir
1024                int16       1024 16-bit tamsayılar kümesindedir
42.0                byte        42 işaretsiz 8-bit tamsayılar kümesindedir
1e10                uint64      10000000000 işaretsiz 64-bit tamsayılar kümesindedir
2.718281828459045   float32     2.718281828459045 2.7182817'ye yuvarlanır ve float32 değerleri kümesindedir
-1e-1000            float64     -1e-1000 IEEE -0.0'a yuvarlanır ve daha da 0.0'a basitleştirilir
0i                  int         0 bir tamsayı değeridir
(42 + 0i)           float32     42.0 (sıfır sanal kısımla) float32 değerleri kümesindedir

x                   T           x, T türündeki bir değerle temsil edilemez çünkü

0                   bool        0 boolean değerleri kümesinde değildir
'a'                 string      'a' bir rune'dur, string değerleri kümesinde değildir
1024                byte        1024 işaretsiz 8-bit tamsayılar kümesinde değildir
-1                  uint16      -1 işaretsiz 16-bit tamsayılar kümesinde değildir
1.1                 int         1.1 bir tamsayı değeri değildir
42i                 float32     (0 + 42i) float32 değerleri kümesinde değildir
1e1000              float64     1e1000 yuvarlama sonrası IEEE +Inf'e taşar
```

---

## 🧩 Method set’ler¶

Bir türün *method set’i*, o türden bir operand üzerinde hangi yöntemlerin çağrılabileceğini belirler. Her türün, onunla ilişkili (muhtemelen boş) bir method set’i vardır:

* Tanımlı bir tür `T`’nin method set’i, alıcı türü `T` olan tüm yöntemlerden oluşur.
* Tanımlı bir tür `T`’ye pointer’ın method set’i (`T` ne pointer ne de interface iken), alıcı türü `*T` veya `T` olan tüm yöntemler kümesidir.
* Bir interface türünün method set’i, interface’in type set’indeki her türün method set’lerinin kesişimidir (çıkan method set genellikle interface’de bildirilen yöntemler kümesidir).

Struct’lar (ve struct’lara pointer’lar) için embedded field içeren durumlarda, struct türleri bölümünde açıklandığı üzere ek kurallar uygulanır. Diğer herhangi bir türün method set’i boştur.

Bir method set içinde, her yöntem benzersiz ve boş olmayan bir yöntem adına sahip olmalıdır.

---

## 🧱 Bloklar¶

Bir blok, eşleşen süslü parantezler içinde bildiriler ve ifadelerden oluşan (muhtemelen boş) bir dizidir.

```ebnf
Block         = "{" StatementList "}" .
StatementList = { Statement ";" } .
```

Kaynak kodundaki açık bloklara ek olarak, örtük bloklar da vardır:

* *Universe block*, tüm Go kaynak metnini kapsar.
* Her paket, o paket için tüm Go kaynak metnini içeren bir *package block*’a sahiptir.
* Her dosya, o dosyadaki tüm Go kaynak metnini içeren bir *file block*’a sahiptir.
* Her `"if"`, `"for"` ve `"switch"` ifadesi kendi örtük bloğu içinde kabul edilir.
* Bir `"switch"` veya `"select"` ifadesindeki her bir clause, örtük bir blok gibi davranır.

Bloklar iç içe geçer ve kapsamı (*scoping*) etkiler.


## 📜 Bildirimler ve kapsam¶

Bir bildirim, boş olmayan bir tanımlayıcıyı bir sabite, türe, tür parametresine, değişkene, fonksiyona, etikete veya pakete bağlar. Bir programdaki her tanımlayıcı bildirilmiş olmalıdır. Aynı blokta hiçbir tanımlayıcı iki kez bildirilemez ve hiçbir tanımlayıcı hem dosya hem de paket bloğunda bildirilemez. 

Boş tanımlayıcı, bir bildirimde herhangi bir tanımlayıcı gibi kullanılabilir; ancak bir bağlama (binding) getirmez ve dolayısıyla bildirilmiş sayılmaz. Paket bloğunda, `init` tanımlayıcısı yalnızca `init` fonksiyon bildirimleri için kullanılabilir ve boş tanımlayıcı gibi yeni bir bağlama getirmez. 

```ebnf
Declaration  = ConstDecl | TypeDecl | VarDecl .
TopLevelDecl = Declaration | FunctionDecl | MethodDecl .
```

Bir bildirilen tanımlayıcının kapsamı, tanımlayıcının belirtilen sabiti, türü, değişkeni, fonksiyonu, etiketi veya paketi gösterdiği kaynak metnin kapsamıdır. 

Go, bloklar kullanarak leksik kapsam (*lexically scoped*) uygular: 

* Önceden bildirilmiş (*predeclared*) bir tanımlayıcının kapsamı *universe block*’tur.
* En üst düzeyde (herhangi bir fonksiyonun dışında) bildirilen bir sabit, tür, değişken veya fonksiyonu (ancak yöntemi değil) gösteren bir tanımlayıcının kapsamı *package block*’tur.
* İçe aktarılan bir paketin paket adının kapsamı, import bildirimini içeren dosyanın *file block*’udur.
* Bir yöntem alıcısını (*method receiver*), fonksiyon parametresini veya sonuç değişkenini gösteren bir tanımlayıcının kapsamı fonksiyon gövdesidir.
* Bir fonksiyonun tür parametresini veya bir yöntem alıcısı tarafından bildirilen bir tür parametresini gösteren bir tanımlayıcının kapsamı, fonksiyon adından sonra başlar ve fonksiyon gövdesinin sonunda biter.
* Bir türün tür parametresini gösteren bir tanımlayıcının kapsamı, tür adından sonra başlar ve `TypeSpec`’in sonunda biter.
* Bir fonksiyon içinde bildirilen bir sabit veya değişken tanımlayıcısının kapsamı, `ConstSpec` veya `VarSpec`’in sonunda (kısa değişken bildirimleri için `ShortVarDecl`) başlar ve en içteki kapsayan bloğun sonunda biter.
* Bir fonksiyon içinde bildirilen bir tür tanımlayıcısının kapsamı, `TypeSpec` içindeki tanımlayıcıda başlar ve en içteki kapsayan bloğun sonunda biter.
* Bir blokta bildirilen bir tanımlayıcı, bir iç blokta yeniden bildirilebilir. İç bildirimin tanımlayıcısı kapsamdayken, iç bildirimin bildirdiği varlığı ifade eder.

`package` maddesi bir bildirim değildir; paket adı hiçbir kapsamda görünmez. Amacı, aynı pakete ait dosyaları tanımlamak ve import bildirimleri için varsayılan paket adını belirtmektir. 

---

## 🏷️ Etiket kapsamları¶

Etiketler (*labels*), etiketli ifadeler tarafından bildirilir ve `"break"`, `"continue"` ve `"goto"` ifadelerinde kullanılır. Hiç kullanılmayan bir etiketi tanımlamak yasadışıdır. Diğer tanımlayıcıların aksine, etiketler blok kapsamlı değildir ve etiket olmayan tanımlayıcılarla çakışmaz. Bir etiketin kapsamı, bildirildiği fonksiyonun gövdesidir ve iç içe geçmiş herhangi bir fonksiyonun gövdesini dışlar. 

---

## 🕳️ Boş tanımlayıcı¶

Boş tanımlayıcı, alt çizgi karakteri `_` ile temsil edilir. Normal (boş olmayan) bir tanımlayıcı yerine anonim bir yer tutucu olarak hizmet eder ve bildirimlerde, operand olarak ve atama ifadelerinde özel anlama sahiptir. 

---

## 🧩 Önceden bildirilmiş tanımlayıcılar¶

Aşağıdaki tanımlayıcılar *universe block* içinde örtük olarak bildirilmiştir [Go 1.18] [Go 1.21]: 

**Türler:**

```text
any bool byte comparable
complex64 complex128 error float32 float64
int int8 int16 int32 int64 rune string
uint uint8 uint16 uint32 uint64 uintptr
```

**Sabitler:**

```text
true false iota
```

**Sıfır değer:**

```text
nil
```

**Fonksiyonlar:**

```text
append cap clear close complex copy delete imag len
make max min new panic print println real recover
```

---

## 📤 Dışa aktarılan tanımlayıcılar¶

Bir tanımlayıcı, başka bir paketten erişime izin vermek için dışa aktarılabilir (*exported*) olabilir. Bir tanımlayıcı, şu iki koşul da sağlanıyorsa dışa aktarılır: 

* Tanımlayıcı adının ilk karakteri bir Unicode büyük harfidir (Unicode karakter kategorisi Lu); ve
* Tanımlayıcı paket bloğunda bildirilmiştir ya da bir alan adı (*field name*) veya yöntem adıdır (*method name*).

Diğer tüm tanımlayıcılar dışa aktarılmaz. 

---

## 🧷 Tanımlayıcıların benzersizliği¶

Bir tanımlayıcı kümesi verildiğinde, bir tanımlayıcı, kümedeki her diğer tanımlayıcıdan farklıysa benzersiz (*unique*) olarak adlandırılır. İki tanımlayıcı şu durumlarda farklıdır: farklı yazılırlarsa veya farklı paketlerde görünüp dışa aktarılmamışlarsa. Aksi hâlde aynıdırlar. 

---

## 🧾 Sabit bildirimleri¶

Bir sabit bildirimi, bir tanımlayıcı listesini (sabit adlarını) bir sabit ifade listesinin değerlerine bağlar. Tanımlayıcı sayısı ifade sayısına eşit olmalıdır ve soldaki n’inci tanımlayıcı sağdaki n’inci ifadenin değerine bağlanır. 

```ebnf
ConstDecl      = "const" ( ConstSpec | "(" { ConstSpec ";" } ")" ) .
ConstSpec      = IdentifierList [ [ Type ] "=" ExpressionList ] .

IdentifierList = identifier { "," identifier } .
ExpressionList = Expression { "," Expression } .
```

Tür mevcutsa, tüm sabitler belirtilen türü alır ve ifadeler o türe atanabilir olmalıdır; bu tür bir tür parametresi olmamalıdır. Tür atlanırsa, sabitler karşılık gelen ifadelerin tek tek türlerini alır. İfade değerleri tipsiz sabitse, bildirilen sabitler tipsiz kalır ve sabit tanımlayıcıları sabit değerleri ifade eder. Örneğin, ifade bir kayan nokta literal’iyse, sabit tanımlayıcı, literal’in kesir kısmı sıfır olsa bile bir kayan nokta sabitini ifade eder. 

```go
const Pi float64 = 3.14159265358979323846
const zero = 0.0         // tipsiz kayan nokta sabiti
const (
	size int64 = 1024
	eof        = -1  // tipsiz tamsayı sabiti
)
const a, b, c = 3, 4, "foo"  // a = 3, b = 4, c = "foo", tipsiz tamsayı ve string sabitler
const u, v float32 = 0, 3    // u = 0.0, v = 3.0
```

Parantezli bir `const` bildirim listesinde, ifade listesi ilk `ConstSpec` dışındakilerde atlanabilir. Böyle boş bir liste, kendisinden önce gelen ilk boş olmayan ifade listesinin ve varsa türünün metinsel ikamesine eşdeğerdir. Dolayısıyla ifade listesini atlamak, bir önceki listeyi tekrarlamaya eşdeğerdir. Tanımlayıcı sayısı, bir önceki listedeki ifade sayısına eşit olmalıdır. `iota` sabit üreteci ile birlikte bu mekanizma ardışık değerlerin hafif bir şekilde bildirilmesine izin verir: 

```go
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Partyday
	numberOfDays  // bu sabit dışa aktarılmaz
)
```

---

## 🧮 Iota¶

Bir sabit bildirimi içinde, önceden bildirilmiş `iota` tanımlayıcısı ardışık tipsiz tamsayı sabitlerini temsil eder. Değeri, ilgili sabit bildirimindeki ilgili `ConstSpec`’in indeksidir ve sıfırdan başlar. İlişkili bir sabitler kümesi oluşturmak için kullanılabilir: 

```go
const (
	c0 = iota  // c0 == 0
	c1 = iota  // c1 == 1
	c2 = iota  // c2 == 2
)

const (
	a = 1 << iota  // a == 1  (iota == 0)
	b = 1 << iota  // b == 2  (iota == 1)
	c = 3          // c == 3  (iota == 2, kullanılmadı)
	d = 1 << iota  // d == 8  (iota == 3)
)

const (
	u         = iota * 42  // u == 0     (tipsiz tamsayı sabiti)
	v float64 = iota * 42  // v == 42.0  (float64 sabiti)
	w         = iota * 42  // w == 84    (tipsiz tamsayı sabiti)
)

const x = iota  // x == 0
const y = iota  // y == 0
```

Tanım gereği, aynı `ConstSpec` içindeki `iota`’nın birden fazla kullanımı aynı değere sahiptir: 

```go
const (
	bit0, mask0 = 1 << iota, 1<<iota - 1  // bit0 == 1, mask0 == 0  (iota == 0)
	bit1, mask1                           // bit1 == 2, mask1 == 1  (iota == 1)
	_, _                                  //                        (iota == 2, kullanılmadı)
	bit3, mask3                           // bit3 == 8, mask3 == 7  (iota == 3)
)
```

Bu son örnek, son boş olmayan ifade listesinin örtük tekrarını kullanır.

---

## 🧾 Tür bildirimleri¶

Bir tür bildirimi, bir tanımlayıcıyı (tür adı) bir türe bağlar. Tür bildirimleri iki biçimde gelir: takma ad (*alias*) bildirimleri ve tür tanımları (*type definitions*). 

```ebnf
TypeDecl = "type" ( TypeSpec | "(" { TypeSpec ";" } ")" ) .
TypeSpec = AliasDecl | TypeDef .
```

---

## 🏷️ Takma ad (Alias) bildirimleri¶

Bir takma ad bildirimi, bir tanımlayıcıyı verilen türe bağlar [Go 1.9]. 

```ebnf
AliasDecl = identifier [ TypeParameters ] "=" Type .
```

Tanımlayıcının kapsamı içinde, verilen tür için bir takma ad olarak hizmet eder. 

```go
type (
	nodeList = []*Node  // nodeList ve []*Node özdeş türlerdir
	Polar    = polar    // Polar ve polar özdeş türleri gösterir
)
```

Takma ad bildirimi tür parametreleri belirtiyorsa [Go 1.24], tür adı jenerik bir takma adı ifade eder. Jenerik takma adlar kullanıldıklarında instantiate edilmelidir. 

```go
type set[P comparable] = map[P]bool
```

Bir takma ad bildiriminde verilen tür bir tür parametresi olamaz. 

```go
type A[P any] = P    // yasadışı: P bir tür parametresidir
```

---

## 🧬 Tür tanımları¶

Bir tür tanımı, verilen türle aynı alttaki (*underlying*) türe ve işlemlere sahip yeni, farklı bir tür oluşturur ve bir tanımlayıcıyı (tür adı) buna bağlar. 

```ebnf
TypeDef = identifier [ TypeParameters ] Type .
```

Yeni türe *tanımlı tür (defined type)* denir. Oluşturulduğu tür dahil, başka herhangi bir türden farklıdır. 

```go
type (
	Point struct{ x, y float64 }  // Point ve struct{ x, y float64 } farklı türlerdir
	polar Point                   // polar ve Point farklı türleri gösterir
)

type TreeNode struct {
	left, right *TreeNode
	value any
}

type Block interface {
	BlockSize() int
	Encrypt(src, dst []byte)
	Decrypt(src, dst []byte)
}
```

Tanımlı bir türe yöntemler (*methods*) iliştirilebilir. Verilen türe bağlı yöntemleri miras almaz; ancak bir interface türünün ya da bileşik bir türün elemanlarının method set’i değişmeden kalır: 

```go
// Mutex, Lock ve Unlock olmak üzere iki yöntemi olan bir veri türüdür.
type Mutex struct         { /* Mutex alanları */ }
func (m *Mutex) Lock()    { /* Lock uygulaması */ }
func (m *Mutex) Unlock()  { /* Unlock uygulaması */ }

// NewMutex, Mutex ile aynı bileşime sahiptir fakat method set'i boştur.
type NewMutex Mutex

// PtrMutex'in alttaki türü *Mutex'in method set'i değişmeden kalır,
// fakat PtrMutex'in method set'i boştur.
type PtrMutex *Mutex

// *PrintableMutex'in method set'i, embedded field Mutex'e bağlı
// Lock ve Unlock yöntemlerini içerir.
type PrintableMutex struct {
	Mutex
}

// MyBlock, Block ile aynı method set'ine sahip bir interface türüdür.
type MyBlock Block
```

Tür tanımları, farklı boolean, sayısal veya string türleri tanımlamak ve bunlarla yöntemler iliştirmek için kullanılabilir: 

```go
type TimeZone int

const (
	EST TimeZone = -(5 + iota)
	CST
	MST
	PST
)

func (tz TimeZone) String() string {
	return fmt.Sprintf("GMT%+dh", tz)
}
```

Tür tanımı tür parametreleri belirtiyorsa, tür adı jenerik bir türü ifade eder. Jenerik türler kullanıldıklarında instantiate edilmelidir. 

```go
type List[T any] struct {
	next  *List[T]
	value T
}
```

Bir tür tanımında verilen tür bir tür parametresi olamaz. 

```go
type T[P any] P    // yasadışı: P bir tür parametresidir

func f[T any]() {
	type L T   // yasadışı: T, kapsayan fonksiyon tarafından bildirilen bir tür parametresidir
}
```

Bir jenerik türün yöntemleri de olabilir. Bu durumda, yöntem alıcıları jenerik tür tanımında bulunan tür parametreleriyle aynı sayıda tür parametresi bildirmelidir. 

```go
// Len yöntemi, bağlı liste l'deki eleman sayısını döndürür.
func (l *List[T]) Len() int  { … }
```

---

## 🧷 Tür parametresi bildirimleri¶

Bir tür parametresi listesi, jenerik bir fonksiyonun veya tür bildiriminin tür parametrelerini bildirir. Tür parametresi listesi, sıradan bir fonksiyon parametre listesi gibi görünür; farkları, tür parametresi adlarının hepsinin mevcut olması ve listenin parantez yerine köşeli parantez içine alınmasıdır [Go 1.18]. 

```ebnf
TypeParameters = "[" TypeParamList [ "," ] "]" .
TypeParamList  = TypeParamDecl { "," TypeParamDecl } .
TypeParamDecl  = IdentifierList TypeConstraint .
```

Listedeki tüm boş olmayan adlar benzersiz olmalıdır. Her ad, bildirimin içinde henüz bilinmeyen bir tür için yer tutucu olarak davranan yeni ve farklı bir adlandırılmış tür olan bir tür parametresi bildirir. Tür parametresi, jenerik fonksiyon veya tür instantiate edildiğinde bir tür argümanı ile değiştirilir. 

```go
[P any]
[S interface{ ~[]byte|string }]
[S ~[]E, E any]
[P Constraint[int]]
[_ any]
```

Her sıradan fonksiyon parametresinin bir parametre türü olduğu gibi, her tür parametresinin de *tür kısıtı (type constraint)* denilen karşılık gelen (meta-)türü vardır. 

Tek bir tür parametresi `P`’yi, `P C` metninin geçerli bir ifade oluşturacağı şekilde bir kısıt `C` ile bildiren jenerik bir tür için tür parametresi listesinde bir ayrıştırma belirsizliği ortaya çıkar: 

```go
type T[P *C] …
type T[P (C)] …
type T[P *C|Q] …
…
```

Bu nadir durumlarda, tür parametresi listesi bir ifadeden ayırt edilemez ve tür bildirimi bir array tür bildirimi olarak ayrıştırılır. Belirsizliği çözmek için, kısıtı bir interface içine gömün veya sonda virgül kullanın: 

```go
type T[P interface{*C}] …
type T[P *C,] …
```

Tür parametreleri, bir jenerik türe bağlı bir yöntem bildiriminin alıcı (*receiver*) belirtimiyle de bildirilebilir. 

Bir jenerik tür `T`’nin tür parametresi listesi içinde, bir tür kısıtı `T`’ye (doğrudan ya da dolaylı olarak, başka bir jenerik türün tür parametresi listesi üzerinden) başvuramaz. 

```go
type T1[P T1[P]] …                    // yasadışı: T1 kendisine başvuruyor
type T2[P interface{ T2[int] }] …     // yasadışı: T2 kendisine başvuruyor
type T3[P interface{ m(T3[int])}] …   // yasadışı: T3 kendisine başvuruyor
type T4[P T5[P]] …                    // yasadışı: T4, T5'e başvuruyor ve
type T5[P T4[P]] …                    //          T5, T4'e başvuruyor

type T6[P int] struct{ f *T6[P] }     // ok: T6'ya başvuru tür parametresi listesinde değil
```

---

## 🧩 Tür kısıtları¶

Bir tür kısıtı, ilgili tür parametresi için izin verilen tür argümanları kümesini tanımlayan ve o tür parametresinin değerleri tarafından desteklenen işlemleri kontrol eden bir interface’tir [Go 1.18]. 

```ebnf
TypeConstraint = TypeElem .
```

Kısıt, `interface{E}` biçiminde bir interface literal’iyse ve `E` gömülü bir tür elemanıysa (yöntem değilse), tür parametresi listesinde kolaylık için dıştaki `interface{ … }` atlanabilir: 

```go
[T []P]                      // = [T interface{[]P}]
[T ~int]                     // = [T interface{~int}]
[T int|string]               // = [T interface{int|string}]
type Constraint ~int         // yasadışı: ~int bir tür parametresi listesinde değil
```

Önceden tanımlı `comparable` interface türü, katı biçimde karşılaştırılabilir (*strictly comparable*) tüm interface olmayan türlerin kümesini ifade eder [Go 1.18]. 

Interface olmayan tür parametreleri olmayan interface’ler karşılaştırılabilir olsalar da, katı biçimde karşılaştırılabilir değildir ve bu nedenle `comparable`’ı implement etmezler. Ancak `comparable`’ı sağlarlar. 

```text
int                          // comparable'ı implement eder (int katı biçimde karşılaştırılabilir)
[]byte                       // comparable'ı implement etmez (slice'lar karşılaştırılamaz)
interface{}                  // comparable'ı implement etmez (yukarıya bakın)
interface{ ~int | ~string }  // yalnızca tür parametresi: comparable'ı implement eder (int, string türleri katı biçimde karşılaştırılabilir)
interface{ comparable }      // yalnızca tür parametresi: comparable'ı implement eder (comparable kendisini implement eder)
interface{ ~int | ~[]byte }  // yalnızca tür parametresi: comparable'ı implement etmez (slice'lar karşılaştırılabilir değildir)
interface{ ~struct{ any } }  // yalnızca tür parametresi: comparable'ı implement etmez (alan any katı biçimde karşılaştırılabilir değildir)
```

`comparable` interface’i ve (doğrudan veya dolaylı olarak) `comparable`’ı gömen interface’ler yalnızca tür kısıtı olarak kullanılabilir. Değerlerin veya değişkenlerin türleri olamazlar ya da başka, interface olmayan türlerin bileşenleri olamazlar. 

---

## ✅ Bir tür kısıtını sağlamak¶

Bir tür argümanı `T`, bir tür kısıtı `C`’yi, `C` tarafından tanımlanan type set’in bir elemanıysa (başka bir deyişle, `C`’yi implement ediyorsa) sağlar. Bir istisna olarak, katı biçimde karşılaştırılabilir bir tür kısıtı, karşılaştırılabilir (katı biçimde karşılaştırılabilir olması gerekmeyen) bir tür argümanı tarafından da sağlanabilir [Go 1.20]. Daha kesin olarak: 

Bir tür `T`, bir kısıt `C`’yi şu durumlarda sağlar:

* `T`, `C`’yi implement eder; veya
* `C`, `interface{ comparable; E }` biçiminde yazılabiliyorsa; burada `E` basic bir interface’tir ve `T` karşılaştırılabilir ve `E`’yi implement eder.

```text
tür argümanı      tür kısıtı                     // kısıt sağlama

int               interface{ ~int }              // sağlanır: int, interface{ ~int }'yi implement eder
string            comparable                     // sağlanır: string comparable'ı implement eder (string katı biçimde karşılaştırılabilir)
[]byte            comparable                     // sağlanmaz: slice'lar karşılaştırılabilir değildir
any               interface{ comparable; int }   // sağlanmaz: any, interface{ int }'yi implement etmez
any               comparable                     // sağlanır: any karşılaştırılabilirdir ve basic interface any'yi implement eder
struct{f any}     comparable                     // sağlanır: struct{f any} karşılaştırılabilirdir ve basic interface any'yi implement eder
any               interface{ comparable; m() }   // sağlanmaz: any, basic interface interface{ m() }'yi implement etmez
interface{ m() }  interface{ comparable; m() }   // sağlanır: interface{ m() } karşılaştırılabilirdir ve basic interface interface{ m() }'yi implement eder
```

Kısıt sağlama kuralındaki istisna nedeniyle, tür parametresi türündeki operandların karşılaştırılması çalışma zamanında *panic* edebilir (comparable tür parametreleri her zaman katı biçimde karşılaştırılabilir olsa bile). 

---

## 📦 Değişken bildirimleri¶

Bir değişken bildirimi bir veya daha fazla değişken oluşturur, karşılık gelen tanımlayıcıları onlara bağlar ve her birine bir tür ve bir başlangıç değeri verir. 

```ebnf
VarDecl = "var" ( VarSpec | "(" { VarSpec ";" } ")" ) .
VarSpec = IdentifierList ( Type [ "=" ExpressionList ] | "=" ExpressionList ) .
```

```go
var i int
var U, V, W float64
var k = 0
var x, y float32 = -1, -2
var (
	i       int
	u, v, s = 2.0, 3.0, "bar"
)
var re, im = complexSqrt(-1)
var _, found = entries[name]  // map araması; yalnızca "found" ile ilgileniliyor
```

Bir ifadeler listesi verilirse, değişkenler atama ifadeleri kurallarına göre bu ifadelerle başlatılır. Aksi hâlde, her değişken sıfır değeriyle başlatılır. 

Tür mevcutsa, her değişkene o tür verilir. Aksi hâlde, her değişkene karşılık gelen başlangıç değerinin türü verilir. Bu değer tipsiz bir sabitse, önce varsayılan türüne örtük olarak dönüştürülür; tipsiz bir boolean değerse, önce `bool` türüne örtük olarak dönüştürülür. Önceden bildirilmiş `nil` tanımlayıcısı, açık türü olmayan bir değişkeni başlatmak için kullanılamaz. 

```go
var d = math.Sin(0.5)  // d float64'tür
var i = 42             // i int'tir
var t, ok = x.(T)      // t T'dir, ok bool'dur
var n = nil            // yasadışı
```

Uygulama kısıtlaması: Bir derleyici, bir fonksiyon gövdesi içinde bildirilen ve hiç kullanılmayan bir değişkeni bildirmenin yasadışı olmasını sağlayabilir. 

---

## ⚡ Kısa değişken bildirimleri¶

Kısa değişken bildirimi şu sözdizimini kullanır: 

```ebnf
ShortVarDecl = IdentifierList ":=" ExpressionList .
```

Bu, başlatıcı ifadelerle fakat türler olmadan yapılan normal bir değişken bildiriminin kısaltmasıdır:

```text
"var" IdentifierList "=" ExpressionList .
```

```go
i, j := 0, 10
f := func() int { return 7 }
ch := make(chan int)
r, w, _ := os.Pipe()  // os.Pipe() bağlı bir File çifti ve bir hata döndürür (varsa)
_, y, _ := coord(p)   // coord() üç değer döndürür; yalnızca y koordinatıyla ilgileniliyor
```

Normal değişken bildirimlerinden farklı olarak, kısa değişken bildirimi; değişkenler başlangıçta aynı blokta (veya blok fonksiyon gövdesiyse parametre listelerinde) daha önce bildirilmişse, aynı türle yeniden bildirim (*redeclare*) yapabilir ve boş olmayan değişkenlerden en az biri yeni olmalıdır. Sonuç olarak, yeniden bildirim yalnızca çok değişkenli kısa bildirimde görünebilir. Yeniden bildirim yeni bir değişken tanıtmaz; yalnızca özgün değişkene yeni bir değer atar. `:=`’nin sol tarafındaki boş olmayan değişken adları benzersiz olmalıdır. 

```go
field1, offset := nextField(str, 0)
field2, offset := nextField(str, offset)  // offset'i yeniden bildirir
x, y, x := 1, 2, 3                        // yasadışı: x, := sol tarafında tekrar ediyor
```

Kısa değişken bildirimleri yalnızca fonksiyonların içinde görünebilir. `"if"`, `"for"` veya `"switch"` ifadelerinin başlatıcıları gibi bazı bağlamlarda, yerel geçici değişkenleri bildirmek için kullanılabilirler. 

---

## 🧰 Fonksiyon bildirimleri¶

Bir fonksiyon bildirimi, bir tanımlayıcıyı (fonksiyon adı) bir fonksiyona bağlar. 

```ebnf
FunctionDecl = "func" FunctionName [ TypeParameters ] Signature [ FunctionBody ] .
FunctionName = identifier .
FunctionBody = Block .
```

Fonksiyon imzası sonuç parametreleri bildiriyorsa, fonksiyon gövdesinin ifade listesi sonlandırıcı bir ifadeyle bitmelidir. 

```go
func IndexRune(s string, r rune) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}
	// geçersiz: return ifadesi eksik
}
```

Fonksiyon bildirimi tür parametreleri belirtiyorsa, fonksiyon adı jenerik bir fonksiyonu ifade eder. Jenerik bir fonksiyon çağrılmadan veya bir değer olarak kullanılmadan önce instantiate edilmelidir. 

```go
func min[T ~int|~float64](x, y T) T {
	if x < y {
		return x
	}
	return y
}
```

Tür parametreleri olmayan bir fonksiyon bildirimi, gövdeyi atlayabilir. Böyle bir bildirim, bir assembly rutini gibi Go dışındaki bir yerde uygulanmış bir fonksiyon için imza sağlar. 

```go
func flushICache(begin, end uintptr)  // haricen uygulanır
```

---

## 🧷 Yöntem (Method) bildirimleri¶

Bir yöntem, bir alıcıya (*receiver*) sahip bir fonksiyondur. Bir yöntem bildirimi, bir tanımlayıcıyı (yöntem adı) bir yönteme bağlar ve yöntemi alıcının temel türüyle ilişkilendirir. 

```ebnf
MethodDecl = "func" Receiver MethodName Signature [ FunctionBody ] .
Receiver   = Parameters .
```

Alıcı, yöntem adından önce gelen ekstra bir parametre bölümü ile belirtilir. Bu parametre bölümü, tek bir variadic olmayan parametre olan alıcıyı bildirmelidir. Türü; tanımlı bir tür `T` veya tanımlı bir tür `T`’ye bir pointer olmalıdır; ayrıca köşeli parantez içinde, tür parametresi adlarının bir listesi `[P1, P2, …]` bunu izleyebilir. `T`’ye alıcı temel türü (*receiver base type*) denir. Bir alıcı temel türü bir pointer veya interface türü olamaz ve yöntemle aynı pakette tanımlanmış olmalıdır. Yöntemin alıcı temel türüne bağlandığı söylenir ve yöntem adı yalnızca `T` veya `*T` türü için seçiciler (*selectors*) içinde görünür. 

Boş olmayan bir alıcı tanımlayıcısı yöntem imzasında benzersiz olmalıdır. Alıcının değeri yöntem gövdesi içinde başvurulmuyorsa, tanımlayıcısı bildirimde atlanabilir. Aynı durum, fonksiyonların ve yöntemlerin parametreleri için de genel olarak geçerlidir. 

Bir temel tür için, ona bağlı yöntemlerin boş olmayan adları benzersiz olmalıdır. Temel tür bir struct türüyse, boş olmayan yöntem ve alan adları ayrı olmalıdır. 

Tanımlı `Point` türü verildiğinde aşağıdaki bildirimler: 

```go
func (p *Point) Length() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

func (p *Point) Scale(factor float64) {
	p.x *= factor
	p.y *= factor
}
```

alıcı türü `*Point` olan `Length` ve `Scale` yöntemlerini `Point` temel türüne bağlar.

Alıcı temel türü jenerik bir türse, alıcı belirtimi yöntemin kullanacağı karşılık gelen tür parametrelerini bildirmelidir. Bu, alıcı tür parametrelerini yöntem için kullanılabilir hâle getirir. Sözdizimsel olarak, bu tür parametresi bildirimi, alıcı temel türün instantiate edilmesine benzer: tür argümanları, alıcı temel tür tanımındaki her tür parametresi için bir tane olacak şekilde, bildirilen tür parametrelerini gösteren tanımlayıcılar olmalıdır. Tür parametresi adlarının, alıcı temel tür tanımındaki karşılık gelen parametre adlarıyla eşleşmesi gerekmez ve tüm boş olmayan parametre adları alıcı parametre bölümünde ve yöntem imzasında benzersiz olmalıdır. Alıcı tür parametresi kısıtları, alıcı temel tür tanımı tarafından ima edilir: karşılık gelen tür parametreleri karşılık gelen kısıtlara sahiptir. 

```go
type Pair[A, B any] struct {
	a A
	b B
}

func (p Pair[A, B]) Swap() Pair[B, A]  { … }  // alıcı A, B'yi bildirir
func (p Pair[First, _]) First() First  { … }  // alıcı First'ü bildirir, Pair içindeki A'ya karşılık gelir
```

Alıcı türü (bir pointer’a) bir takma ad (*alias*) ile gösteriliyorsa, takma ad jenerik olmamalı ve instantiate edilmiş bir jenerik türü—doğrudan veya dolaylı olarak (başka bir takma ad üzerinden) ve pointer dolaylılıklarından bağımsız olarak—göstermemelidir. 

```go
type GPoint[P any] = Point
type HPoint        = *GPoint[int]
type IPair         = Pair[int, int]

func (*GPoint[P]) Draw(P)   { … }  // yasadışı: takma ad jenerik olmamalı
func (HPoint) Draw(P)       { … }  // yasadışı: takma ad instantiate edilmiş tür GPoint[int]'i göstermemeli
func (*IPair) Second() int  { … }  // yasadışı: takma ad instantiate edilmiş tür Pair[int, int]'i göstermemeli
```


## ⚖️ Karşılaştırma işleçleri¶

Karşılaştırma işleçleri iki operandı karşılaştırır ve *tipsiz (untyped)* bir boolean değer üretir. 

```text
==    eşit
!=    eşit değil
<     küçük
<=    küçük veya eşit
>     büyük
>=    büyük veya eşit
```

Herhangi bir karşılaştırmada, birinci operand ikinci operandın türüne *atanabilir (assignable)* olmalıdır ya da tersi geçerli olmalıdır. 

Eşitlik işleçleri `==` ve `!=`, *karşılaştırılabilir (comparable)* türlerdeki operandlara uygulanır. Sıralama işleçleri `<`, `<=`, `>`, `>=`, *sıralı (ordered)* türlerdeki operandlara uygulanır. Bu terimler ve karşılaştırma sonuçları aşağıdaki gibi tanımlanır: 

* Boolean türleri karşılaştırılabilirdir. İki boolean değeri, ya ikisi de `true` ya da ikisi de `false` ise eşittir.
* Tamsayı türleri karşılaştırılabilir ve sıralıdır. İki tamsayı değeri alışıldık biçimde karşılaştırılır.
* Kayan nokta türleri karşılaştırılabilir ve sıralıdır. İki kayan nokta değeri IEEE 754 standardınca tanımlandığı gibi karşılaştırılır.
* Kompleks türler karşılaştırılabilirdir. İki kompleks sayı `u` ve `v`, hem `real(u) == real(v)` hem de `imag(u) == imag(v)` ise eşittir.
* String türleri karşılaştırılabilir ve sıralıdır. İki string değeri, bayt düzeyinde sözlüksel (lexical) olarak karşılaştırılır.
* Pointer türleri karşılaştırılabilirdir. İki pointer, aynı değişkeni gösteriyorlarsa veya ikisinin değeri de `nil` ise eşittir. Ayrı sıfır-boyutlu değişkenlere işaret eden pointer’lar eşit olabilir de olmayabilir de.
* Channel türleri karşılaştırılabilirdir. İki channel değeri, `make` çağrısının aynı çağrısıyla oluşturuldularsa veya ikisi de `nil` ise eşittir.
* Tür parametresi olmayan interface türleri karşılaştırılabilirdir. İki interface değeri, özdeş dinamik türlere ve eşit dinamik değerlere sahipse veya ikisi de `nil` ise eşittir.
* Interface olmayan `X` türünde bir değer `x` ile interface `T` türünde bir değer `t`, `X` karşılaştırılabilir ve `X`, `T`’yi implement ediyorsa karşılaştırılabilir. `t`’nin dinamik türü `X` ile özdeşse ve `t`’nin dinamik değeri `x` ile eşitse eşittirler.
* Struct türleri, tüm alan türleri karşılaştırılabilirse karşılaştırılabilirdir. İki struct değeri, karşılık gelen boş olmayan alan değerleri eşitse eşittir. Alanlar kaynak sırasına göre karşılaştırılır; ilk farklılıkta durulur (ya da tüm alanlar karşılaştırıldıysa biter).
* Array türleri, eleman türleri karşılaştırılabilirse karşılaştırılabilirdir. İki array değeri, karşılık gelen eleman değerleri eşitse eşittir. Elemanlar artan indeks sırasıyla karşılaştırılır; ilk farklılıkta durulur (ya da tüm elemanlar karşılaştırıldıysa biter).
* Tür parametreleri, *katı biçimde karşılaştırılabilir (strictly comparable)* iseler karşılaştırılabilirdir (aşağıya bakın).
* Özdeş dinamik türlere sahip iki interface değerinin karşılaştırılması, o dinamik tür karşılaştırılabilir değilse çalışma zamanında *panic* oluşturur. Bu davranış yalnızca doğrudan interface değer karşılaştırmalarında değil, interface değerlerinden oluşan array’leri veya interface alanlı struct’ları karşılaştırırken de geçerlidir. 

Slice, map ve function türleri karşılaştırılabilir değildir. Ancak özel bir durum olarak, bir slice, map veya function değeri önceden bildirilmiş `nil` tanımlayıcısıyla karşılaştırılabilir. Pointer, channel ve interface değerlerinin `nil` ile karşılaştırılması da izinlidir ve yukarıdaki genel kurallardan çıkar. 

```go
const c = 3 < 4            // c, tipsiz boolean sabiti true'dur

type MyBool bool
var x, y int
var (
	// Bir karşılaştırmanın sonucu tipsiz bir boolean'dır.
	// Olağan atama kuralları geçerlidir.
	b3        = x == y // b3'ün türü bool'dur
	b4 bool   = x == y // b4'ün türü bool'dur
	b5 MyBool = x == y // b5'in türü MyBool'dur
)
```

Bir tür, karşılaştırılabilir olup interface türü değilse ve interface türlerinden oluşmuyorsa *katı biçimde karşılaştırılabilir*dir. Özellikle: 

* Boolean, sayısal, string, pointer ve channel türleri katı biçimde karşılaştırılabilirdir.
* Struct türleri, tüm alan türleri katı biçimde karşılaştırılabilirse katı biçimde karşılaştırılabilirdir.
* Array türleri, eleman türleri katı biçimde karşılaştırılabilirse katı biçimde karşılaştırılabilirdir.
* Tür parametreleri, type set’lerindeki tüm türler katı biçimde karşılaştırılabilirse katı biçimde karşılaştırılabilirdir.

---

## 🧠 Mantıksal işleçler¶

Mantıksal işleçler boolean değerlere uygulanır ve operandlarla aynı türde bir sonuç üretir. Sol operand değerlendirilir; sonra koşul gerektiriyorsa sağ operand değerlendirilir. 

```text
&&    koşullu AND    p && q  =  "eğer p ise q, değilse false"
||    koşullu OR     p || q  =  "eğer p ise true, değilse q"
!     NOT            !p      =  "p'nin değili"
```

---

## 🧭 Adres işleçleri¶

`T` türünde bir operand `x` için, adres alma işlemi `&x`, `x`’e işaret eden `*T` türünde bir pointer üretir. Operand adreslenebilir olmalıdır; yani bir değişken, pointer dereference işlemi, slice indeksleme işlemi, adreslenebilir bir struct operandının alan seçicisi veya adreslenebilir bir array’in indeksleme işlemi olmalıdır. Adreslenebilirlik kuralına istisna olarak, `x` bir (muhtemelen parantezlenmiş) bileşik literal de olabilir. `x`’in değerlendirilmesi çalışma zamanında *panic* üretecekse, `&x`’in değerlendirilmesi de *panic* üretir. 

`*T` pointer türünde bir operand `x` için, pointer dereference işlemi `*x`, `x`’in işaret ettiği `T` türündeki değişkeni belirtir. `x` `nil` ise `*x`’i değerlendirme girişimi çalışma zamanında *panic* oluşturur. 

```go
&x
&a[f(2)]
&Point{2, 3}
*p
*pf(x)

var x *int = nil
*x   // çalışma zamanında panic oluşturur
&*x  // çalışma zamanında panic oluşturur
```

---

## 📩 Alma (receive) işleci¶

Channel türünde bir operand `ch` için, alma işlemi `<-ch`’nin değeri, `ch` channel’ından alınan değerdir. Channel yönü alma işlemlerine izin vermelidir ve alma işleminin türü, channel’ın eleman türüdür. İfade, bir değer kullanılabilir olana kadar bloklanır. `nil` bir channel’dan almak sonsuza dek bloklanır. Kapalı bir channel’da alma işlemi her zaman anında ilerleyebilir; daha önce gönderilmiş değerler alındıktan sonra, eleman türünün sıfır değerini üretir. 

```go
v1 := <-ch
v2 = <-ch
f(<-ch)
<-strobe  // saat darbesini bekle ve alınan değeri at
```

Operand türü bir tür parametresiyse, type set’teki tüm türler alma işlemlerine izin veren channel türleri olmalı ve hepsinin eleman türü aynı olmalıdır; bu eleman türü, alma işleminin türüdür. 

Özel biçimdeki bir atama ifadesinde veya başlatmada kullanılan bir alma ifadesi:

```go
x, ok = <-ch
x, ok := <-ch
var x, ok = <-ch
var x, ok T = <-ch
```

iletişimin başarılı olup olmadığını raporlayan ek bir *tipsiz* boolean sonuç üretir. `ok` değeri, alınan değer başarılı bir gönderim işlemiyle iletildiyse `true`, channel kapalı ve boş olduğu için üretilen bir sıfır değerse `false` olur. 

---

## 🔄 Dönüşümler (conversions)¶

Bir dönüşüm, bir ifadenin türünü dönüşümde belirtilen türe değiştirir. Dönüşüm kaynak kodda açıkça görünebilir veya bir ifadenin göründüğü bağlam tarafından ima edilebilir. 

Açık bir dönüşüm, `T(x)` biçiminde bir ifadedir; burada `T` bir tür, `x` ise `T` türüne dönüştürülebilen bir ifadedir. 

```ebnf
Conversion = Type "(" Expression [ "," ] ")" .
```

Tür `*` veya `<-` işleciyle başlıyorsa ya da tür `func` anahtar sözcüğüyle başlıyor ve sonuç listesi yoksa, belirsizliği önlemek için gerektiğinde parantezlenmelidir: 

```go
*Point(p)        // *(Point(p)) ile aynı
(*Point)(p)      // p, *Point türüne dönüştürülür
<-chan int(c)    // <-(chan int(c)) ile aynı
(<-chan int)(c)  // c, <-chan int türüne dönüştürülür
func()(x)        // func() x imzalı fonksiyon
(func())(x)      // x, func() türüne dönüştürülür
(func() int)(x)  // x, func() int türüne dönüştürülür
func() int(x)    // x, func() int türüne dönüştürülür (belirsiz değil)
```

Bir sabit değer `x`, `x`, `T` türünde bir değerle *temsil edilebiliyorsa (representable)* `T` türüne dönüştürülebilir. Özel bir durum olarak, bir tamsayı sabiti `x`, sabit olmayan `x` için kullanılan kuralla aynı kuralı kullanarak açıkça bir string türüne dönüştürülebilir. 

Tür parametresi olmayan bir türe sabit dönüşümü, *türlü (typed)* bir sabit üretir. 

```go
uint(iota)               // türü uint olan iota değeri
float32(2.718281828)     // türü float32 olan 2.718281828
complex128(1)            // türü complex128 olan 1.0 + 0.0i
float32(0.49999999)      // türü float32 olan 0.5
float64(-1e-1000)        // türü float64 olan 0.0
string('x')              // türü string olan "x"
string(0x266c)           // türü string olan "♬"
myString("foo" + "bar")  // türü myString olan "foobar"
string([]byte{'a'})      // sabit değil: []byte{'a'} sabit değildir
(*int)(nil)              // sabit değil: nil sabit değildir, *int boolean/sayısal/string türü değildir
int(1.2)                 // yasadışı: 1.2 bir int ile temsil edilemez
string(65.0)             // yasadışı: 65.0 bir tamsayı sabiti değildir
```

Bir sabiti bir tür parametresine dönüştürmek, o türden sabit olmayan bir değer üretir; değer, tür parametresinin instantiate edildiği tür argümanı değer temsiliyle temsil edilir. Örneğin, şu fonksiyonda: 

```go
func f[P ~float32|~float64]() {
	… P(1.1) …
}
```

`P(1.1)` dönüşümü `P` türünde sabit olmayan bir değer üretir ve `1.1` değeri, `f` için tür argümanına bağlı olarak `float32` veya `float64` olarak temsil edilir. Buna göre, `f` bir `float32` türüyle instantiate edilirse, `P(1.1) + 1.2` ifadesinin sayısal hesabı, karşılık gelen sabit olmayan `float32` toplamanın hassasiyetiyle yapılır. 

Sabit olmayan bir değer `x`, aşağıdaki durumlardan herhangi birinde `T` türüne dönüştürülebilir: 

* `x`, `T`’ye atanabilirdir.
* Struct etiketleri yok sayılarak (aşağıya bakın), `x`’in türü ve `T` tür parametresi değildir ve özdeş alttaki türlere sahiptir.
* Struct etiketleri yok sayılarak, `x`’in türü ve `T`, adlandırılmış tür olmayan pointer türleridir; pointer temel türleri tür parametresi değildir ve özdeş alttaki türlere sahiptir.
* `x`’in türü ve `T` hem tamsayı hem de kayan nokta türleridir.
* `x`’in türü ve `T` her ikisi de kompleks türlerdir.
* `x` bir tamsayı veya bayt/rune slice’ıdır ve `T` bir string türüdür.
* `x` bir string’dir ve `T` bayt veya rune slice’ıdır.
* `x` bir slice’tır, `T` bir array `[Go 1.20]` veya array pointer’ıdır `[Go 1.17]` ve slice ile array türlerinin eleman türleri özdeştir.

Ek olarak, `T` veya `x`’in türü `V` tür parametresiyse, aşağıdakilerden biri geçerliyse `x`, `T` türüne dönüştürülebilir: 

* Hem `V` hem `T` tür parametresidir ve `V`’nin type set’indeki her türden bir değer, `T`’nin type set’indeki her türe dönüştürülebilir.
* Yalnızca `V` tür parametresidir ve `V`’nin type set’indeki her türden bir değer, `T`’ye dönüştürülebilir.
* Yalnızca `T` tür parametresidir ve `x`, `T`’nin type set’indeki her türe dönüştürülebilir.

Dönüşüm amacıyla struct tür kimliği karşılaştırmalarında struct etiketleri yok sayılır: 

```go
type Person struct {
	Name    string
	Address *struct {
		Street string
		City   string
	}
}

var data *struct {
	Name    string `json:"name"`
	Address *struct {
		Street string `json:"street"`
		City   string `json:"city"`
	} `json:"address"`
}

var person = (*Person)(data)  // etiketleri yok sayınca alttaki türler özdeştir
```

Sayısal türler arasında veya bir string türüne (ya da string türünden) yapılan (sabit olmayan) dönüşümler için özel kurallar geçerlidir. Bu dönüşümler `x`’in temsilini değiştirebilir ve çalışma zamanı maliyeti doğurabilir. Diğer tüm dönüşümler yalnızca türü değiştirir, `x`’in temsilini değiştirmez. 

Pointer ve tamsayılar arasında dönüşüm için dilde bir mekanizma yoktur. `unsafe` paketi bu işlevi, kısıtlı koşullar altında sağlar. 

### 🧮 Sayısal türler arasında dönüşümler

Sabit olmayan sayısal değerlerin dönüşümünde şu kurallar geçerlidir: 

* Tamsayı türleri arasında dönüşümde, değer işaretliyse örtük sonsuz hassasiyete işaret genişletmesi (*sign extension*), işaretsizse sıfır genişletmesi (*zero extension*) yapılır. Ardından sonuç türünün boyutuna sığacak şekilde kırpılır (*truncate*). Örneğin `v := uint16(0x10F0)` ise `uint32(int8(v)) == 0xFFFFFFF0`. Dönüşüm her zaman geçerli bir değer üretir; taşma göstergesi yoktur.
* Kayan noktadan tamsayıya dönüşümde kesir kısmı atılır (sıfıra doğru kırpma).
* Tamsayı veya kayan nokta sayıdan kayan nokta türüne ya da bir kompleks sayıdan başka bir kompleks türe dönüşümde, sonuç değeri hedef türün hassasiyetine yuvarlanır. Örneğin `float32` türünde bir değişken `x`, IEEE 754 32-bit sayıdan daha yüksek hassasiyetle saklanabilir; ancak `float32(x)`, `x`’in değerini 32-bit hassasiyete yuvarlama sonucunu temsil eder. Benzer şekilde, `x + 0.1` 32 bitten fazla hassasiyet kullanabilir; ancak `float32(x + 0.1)` kullanmaz.
* Kayan nokta veya kompleks değer içeren tüm sabit olmayan dönüşümlerde, sonuç türü değeri temsil edemiyorsa dönüşüm başarılı olur, fakat sonuç değeri *uygulamaya bağlıdır (implementation-dependent)*.

### 🧵 String türüne ve string türünden dönüşümler¶

Bayt slice’ını string türüne dönüştürmek, ardışık baytları slice elemanları olan bir string üretir: 

```go
string([]byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'})   // "hellø"
string([]byte{})                                     // ""
string([]byte(nil))                                  // ""

type bytes []byte
string(bytes{'h', 'e', 'l', 'l', '\xc3', '\xb8'})    // "hellø"

type myByte byte
string([]myByte{'w', 'o', 'r', 'l', 'd', '!'})       // "world!"
myString([]myByte{'\xf0', '\x9f', '\x8c', '\x8d'})   // "🌍"
```

Rune slice’ını string türüne dönüştürmek, tek tek rune değerlerinin string’e dönüştürülüp birleştirilmesiyle oluşan bir string üretir: 

```go
string([]rune{0x767d, 0x9d6c, 0x7fd4})   // "\u767d\u9d6c\u7fd4" == "白鵬翔"
string([]rune{})                         // ""
string([]rune(nil))                      // ""

type runes []rune
string(runes{0x767d, 0x9d6c, 0x7fd4})    // "\u767d\u9d6c\u7fd4" == "白鵬翔"

type myRune rune
string([]myRune{0x266b, 0x266c})         // "\u266b\u266c" == "♫♬"
myString([]myRune{0x1f30e})              // "\U0001f30e" == "🌎"
```

String türündeki bir değeri bayt slice türüne dönüştürmek, ardışık elemanları string’in baytları olan `nil` olmayan bir slice üretir. Oluşan slice’ın kapasitesi uygulamaya özeldir ve slice uzunluğundan büyük olabilir: 

```go
[]byte("hellø")             // []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}
[]byte("")                  // []byte{}

bytes("hellø")              // []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}

[]myByte("world!")          // []myByte{'w', 'o', 'r', 'l', 'd', '!'}
[]myByte(myString("🌏"))    // []myByte{'\xf0', '\x9f', '\x8c', '\x8f'}
```

String türündeki bir değeri rune slice türüne dönüştürmek, string’in Unicode kod noktalarından oluşan bir slice üretir. Oluşan slice’ın kapasitesi uygulamaya özeldir ve slice uzunluğundan büyük olabilir: 

```go
[]rune(myString("白鵬翔"))   // []rune{0x767d, 0x9d6c, 0x7fd4}
[]rune("")                  // []rune{}

runes("白鵬翔")              // []rune{0x767d, 0x9d6c, 0x7fd4}

[]myRune("♫♬")              // []myRune{0x266b, 0x266c}
[]myRune(myString("🌐"))    // []myRune{0x1f310}
```

Son olarak, tarihsel nedenlerle, bir tamsayı değer string türüne dönüştürülebilir. Bu dönüşüm biçimi, verilen tamsayı değerinin temsil ettiği Unicode kod noktasının (muhtemelen çok baytlı) UTF-8 gösterimini içeren bir string üretir. Geçerli Unicode kod noktaları aralığının dışındaki değerler `"\uFFFD"`’ye dönüştürülür. 

```go
string('a')          // "a"
string(65)           // "A"
string('\xf8')       // "\u00f8" == "ø" == "\xc3\xb8"
string(-1)           // "\ufffd" == "\xef\xbf\xbd"

type myString string
myString('\u65e5')   // "\u65e5" == "日" == "\xe6\x97\xa5"
```

Not: Bu dönüşüm biçimi ileride dilden kaldırılabilir. `go vet` aracı, bazı tamsayıdan string’e dönüşümleri olası hata olarak işaretler. Bunun yerine `utf8.AppendRune` veya `utf8.EncodeRune` gibi kütüphane fonksiyonları kullanılmalıdır. 

### 🧩 Slice’tan array’e veya array pointer’ına dönüşümler¶

Bir slice’ı array’e dönüştürmek, slice’ın alttaki array’inin elemanlarını içeren bir array üretir. Benzer şekilde, bir slice’ı array pointer’ına dönüştürmek, slice’ın alttaki array’ine işaret eden bir pointer üretir. Her iki durumda da, slice uzunluğu array uzunluğundan küçükse çalışma zamanında *panic* oluşur. 

```go
s := make([]byte, 2, 4)

a0 := [0]byte(s)
a1 := [1]byte(s[1:])     // a1[0] == s[1]
a2 := [2]byte(s)         // a2[0] == s[0]
a4 := [4]byte(s)         // panic: len([4]byte) > len(s)

s0 := (*[0]byte)(s)      // s0 != nil
s1 := (*[1]byte)(s[1:])  // &s1[0] == &s[1]
s2 := (*[2]byte)(s)      // &s2[0] == &s[0]
s4 := (*[4]byte)(s)      // panic: len([4]byte) > len(s)

var t []string
t0 := [0]string(t)       // nil slice t için ok
t1 := (*[0]string)(t)    // t1 == nil
t2 := (*[1]string)(t)    // panic: len([1]string) > len(t)

u := make([]byte, 0)
u0 := (*[0]byte)(u)      // u0 != nil
```

---

## 🧾 Sabit ifadeler¶

Sabit ifadeler yalnızca sabit operandlar içerebilir ve derleme zamanında değerlendirilir. 

Tipsiz boolean, sayısal ve string sabitleri; sırasıyla boolean, sayısal veya string türünde operand kullanmanın yasal olduğu her yerde operand olarak kullanılabilir.

Bir sabit karşılaştırma her zaman tipsiz bir boolean sabiti üretir. Bir sabit kaydırma (*shift*) ifadesinin sol operandı tipsiz bir sabitse, sonuç bir tamsayı sabitidir; aksi hâlde, sol operandla aynı türde bir sabittir ve sol operand tamsayı türünde olmalıdır. 

Tipsiz sabitlerle yapılan diğer tüm işlemler, aynı “kind”de tipsiz bir sabit üretir; yani boolean, tamsayı, kayan nokta, kompleks veya string sabiti. (Kaydırma dışındaki) ikili bir işlemde tipsiz operandların “kind”i farklıysa, sonuç bu listedeki daha sonra gelen operandın “kind”ine sahip olur: `integer`, `rune`, `floating-point`, `complex`. Örneğin, tipsiz bir tamsayı sabitinin tipsiz bir kompleks sabite bölünmesi tipsiz bir kompleks sabit üretir. 

```go
const a = 2 + 3.0          // a == 5.0   (tipsiz kayan nokta sabiti)
const b = 15 / 4           // b == 3     (tipsiz tamsayı sabiti)
const c = 15 / 4.0         // c == 3.75  (tipsiz kayan nokta sabiti)
const Θ float64 = 3/2      // Θ == 1.0   (tür float64, 3/2 tamsayı bölmesidir)
const Π float64 = 3/2.     // Π == 1.5   (tür float64, 3/2. float bölmesidir)
const d = 1 << 3.0         // d == 8     (tipsiz tamsayı sabiti)
const e = 1.0 << 3         // e == 8     (tipsiz tamsayı sabiti)
const f = int32(1) << 33   // yasadışı  (sabit 8589934592 int32 taşması yapar)
const g = float64(2) >> 1  // yasadışı  (float64(2) türlü kayan nokta sabitidir)
const h = "foo" > "bar"    // h == true (tipsiz boolean sabiti)
const j = true             // j == true (tipsiz boolean sabiti)
const k = 'w' + 1          // k == 'x'  (tipsiz rune sabiti)
const l = "hi"             // l == "hi" (tipsiz string sabiti)
const m = string(k)        // m == "x"  (tür string)
const Σ = 1 - 0.707i       //
const Δ = Σ + 2.0e-4       //
const Φ = iota*1i - 1/1i   //
```

Yerleşik `complex` fonksiyonunu tipsiz tamsayı, rune veya kayan nokta sabitlerine uygulamak tipsiz bir kompleks sabit üretir. 

```go
const ic = complex(0, c)   // ic == 3.75i (tipsiz kompleks sabit)
const iΘ = complex(0, Θ)   // iΘ == 1i    (tür complex128)
```

Sabit ifadeler her zaman tam olarak değerlendirilir; ara değerler ve sabitlerin kendileri, dilde önceden tanımlı herhangi bir türün desteklediğinden çok daha yüksek hassasiyet gerektirebilir. Aşağıdakiler yasal bildirimlerdir: 

```go
const Huge = 1 << 100         // Huge == 1267650600228229401496703205376 (tipsiz tamsayı sabiti)
const Four int8 = Huge >> 98  // Four == 4                               (tür int8)
```

Sabit bölme veya kalan işleminin böleni sıfır olmamalıdır: 

```go
3.14 / 0.0   // yasadışı: sıfıra bölme
```

Türlü sabitlerin değerleri her zaman sabit türünün değerleriyle doğru biçimde temsil edilebilir olmalıdır. Aşağıdaki sabit ifadeler yasadışıdır: 

```go
uint(-1)     // -1 bir uint ile temsil edilemez
int(3.14)    // 3.14 bir int ile temsil edilemez
int64(Huge)  // 1267650600228229401496703205376 bir int64 ile temsil edilemez
Four * 300   // operand 300 int8 ile temsil edilemez (Four'un türü)
Four * 100   // çarpım 400 int8 ile temsil edilemez (Four'un türü)
```

Unary bit düzeyi tümleme işleci `^` için kullanılan maske, sabit olmayanlar için olan kuralla eşleşir: işaretsiz sabitlerde maske tüm bitleri `1`, işaretli ve tipsiz sabitlerde maske `-1`’dir. 

```go
^1         // tipsiz tamsayı sabiti, -2'ye eşit
uint8(^1)  // yasadışı: uint8(-2) ile aynı, -2 uint8 ile temsil edilemez
^uint8(1)  // türlü uint8 sabiti, 0xFF ^ uint8(1) = uint8(0xFE) ile aynı
int8(^1)   // int8(-2) ile aynı
^int8(1)   // -1 ^ int8(1) = -2 ile aynı
```

Uygulama kısıtlaması: Bir derleyici, tipsiz kayan nokta veya kompleks sabit ifadeleri hesaplanırken yuvarlama kullanabilir; sabitler bölümündeki uygulama kısıtlamasına bakın. Bu yuvarlama, sonsuz hassasiyetle hesaplandığında tamsayı olabilecek bir kayan nokta sabit ifadesinin tamsayı bağlamında geçersiz olmasına (veya tersi) neden olabilir. 

---

## ⏱️ Değerlendirme sırası¶

Paket düzeyinde, başlatma bağımlılıkları değişken bildirimlerindeki tek tek başlatma ifadelerinin değerlendirme sırasını belirler. Bunun dışında; bir ifadenin, atamanın veya `return` ifadesinin operandları değerlendirilirken, tüm fonksiyon çağrıları, yöntem çağrıları, alma işlemleri ve ikili mantıksal işlemler leksik olarak soldan sağa değerlendirilir. 

Örneğin, (fonksiyon-içi) şu atamada:

```go
y[f()], ok = g(z || h(), i()+x[j()], <-c), k()
```

fonksiyon çağrıları ve iletişim şu sırayla olur: `f()`, `h()` (eğer `z` `false` değerlendirilirse), `i()`, `j()`, `<-c`, `g()`, `k()`. Ancak bu olayların; `x`’in değerlendirilmesi ve indekslenmesine ve `y` ile `z`’nin değerlendirilmesine göre göreli sırası, leksik gereklilikler dışında belirtilmemiştir. Örneğin `g`, argümanları değerlendirilmeden çağrılamaz. 

```go
a := 1
f := func() int { a++; return a }
x := []int{a, f()}            // x [1, 2] veya [2, 2] olabilir: a ile f() arasında değerlendirme sırası belirtilmemiştir
m := map[int]int{a: 1, a: 2}  // m {2: 1} veya {2: 2} olabilir: iki map ataması arasında değerlendirme sırası belirtilmemiştir
n := map[int]int{a: f()}      // n {2: 3} veya {3: 3} olabilir: anahtar ile değer arasında değerlendirme sırası belirtilmemiştir
```

Paket düzeyinde, başlatma bağımlılıkları tek tek başlatma ifadeleri için soldan sağa kuralını geçersiz kılar; ancak her ifadenin içindeki operandlar için kuralı geçersiz kılmaz: 

```go
var a, b, c = f() + v(), g(), sqr(u()) + v()

func f() int        { return c }
func g() int        { return a }
func sqr(x int) int { return x*x }

// u ve v fonksiyonları diğer tüm değişken ve fonksiyonlardan bağımsızdır
```

Fonksiyon çağrıları şu sırayla gerçekleşir: `u()`, `sqr()`, `v()`, `f()`, `v()`, `g()`.

Tek bir ifade içindeki kayan nokta işlemleri, işleçlerin bağlaşıklığına (*associativity*) göre değerlendirilir. Açık parantezler, varsayılan bağlaşıklığı geçersiz kılar. `x + (y + z)` ifadesinde, `y + z` toplaması `x` ile toplamadan önce yapılır. 

---

## 🧰 İşleçler (operators)¶

İşleçler operandları ifadeler hâlinde birleştirir. 

```ebnf
Expression = UnaryExpr | Expression binary_op Expression .
UnaryExpr  = PrimaryExpr | unary_op UnaryExpr .

binary_op  = "||" | "&&" | rel_op | add_op | mul_op .
rel_op     = "==" | "!=" | "<" | "<=" | ">" | ">=" .
add_op     = "+" | "-" | "|" | "^" .
mul_op     = "*" | "/" | "%" | "<<" | ">>" | "&" | "&^" .

unary_op   = "+" | "-" | "!" | "^" | "*" | "&" | "<-" .
```

Karşılaştırmalar başka bir yerde ele alınır. Diğer ikili işleçler için, işlem kaydırmaları veya tipsiz sabitleri içermedikçe operand türleri özdeş olmalıdır. Yalnızca sabitleri içeren işlemler için sabit ifadeler bölümüne bakın. 

Kaydırma işlemleri hariç, operandlardan biri tipsiz bir sabit ve diğeri tipsiz değilse, sabit diğer operandın türüne örtük olarak dönüştürülür. 

Bir kaydırma ifadesinin sağ operandı `int` türünde olmalıdır [Go 1.13] veya `uint` türünde bir değerle temsil edilebilen tipsiz bir sabit olmalıdır. Sabit olmayan bir kaydırma ifadesinin sol operandı tipsiz bir sabitse, önce kaydırma ifadesi yerine yalnızca sol operand varmış gibi alacağı türe örtük olarak dönüştürülür. 

```go
var a [1024]byte
var s uint = 33

// Aşağıdaki örneklerin sonuçları 64-bit int'ler içindir.
var i = 1<<s                   // 1'in türü int'tir
var j int32 = 1<<s             // 1'in türü int32; j == 0
var k = uint64(1<<s)           // 1'in türü uint64; k == 1<<33
var m int = 1.0<<s             // 1.0'ın türü int; m == 1<<33
var n = 1.0<<s == j            // 1.0'ın türü int32; n == true
var o = 1<<s == 2<<s           // 1 ve 2'nin türü int; o == false
var p = 1<<s == 1<<33          // 1'in türü int; p == true
var u = 1.0<<s                 // yasadışı: 1.0'ın türü float64, kaydırılamaz
var u1 = 1.0<<s != 0           // yasadışı: 1.0'ın türü float64, kaydırılamaz
var u2 = 1<<s != 1.0           // yasadışı: 1'in türü float64, kaydırılamaz
var v1 float32 = 1<<s          // yasadışı: 1'in türü float32, kaydırılamaz
var v2 = string(1<<s)          // yasadışı: 1 string'e dönüştürülmüş, kaydırılamaz
var w int64 = 1.0<<33          // 1.0<<33 bir sabit kaydırma ifadesidir; w == 1<<33
var x = a[1.0<<s]              // panic: 1.0'ın türü int, fakat 1<<33 array sınırlarını aşar
var b = make([]byte, 1.0<<s)   // 1.0'ın türü int; len(b) == 1<<33

// Aşağıdaki örneklerin sonuçları 32-bit int'ler içindir,
// bu da kaydırmaların taşmasına yol açar.
var mm int = 1.0<<s            // 1.0'ın türü int; mm == 0
var oo = 1<<s == 2<<s          // 1 ve 2'nin türü int; oo == true
var pp = 1<<s == 1<<33         // yasadışı: 1'in türü int, fakat 1<<33 int taşması yapar
var xx = a[1.0<<s]             // 1.0'ın türü int; xx == a[0]
var bb = make([]byte, 1.0<<s)  // 1.0'ın türü int; len(bb) == 0
```

---

## 🎚️ İşleç önceliği¶

Unary işleçler en yüksek önceliğe sahiptir. `++` ve `--` işleçleri ifade değil, birer bildirim oluşturduğundan işleç hiyerarşisinin dışındadır. Sonuç olarak `*p++` ifadesi `(*p)++` ile aynıdır. 

İkili işleçler için beş öncelik düzeyi vardır. Çarpma işleçleri en sıkı bağlanır; ardından toplama işleçleri, sonra karşılaştırma işleçleri, sonra `&&` (mantıksal AND) ve son olarak `||` (mantıksal OR) gelir: 

| Öncelik | İşleçler                       |       |   |
| ------: | ------------------------------ | ----- | - |
|       5 | `*` `/` `%` `<<` `>>` `&` `&^` |       |   |
|       4 | `+` `-` `                      | ` `^` |   |
|       3 | `==` `!=` `<` `<=` `>` `>=`    |       |   |
|       2 | `&&`                           |       |   |
|       1 | `                              |       | ` |

Aynı öncelikteki ikili işleçler soldan sağa bağlanır. Örneğin `x / y * z`, `(x / y) * z` ile aynıdır. 

```go
+x                         // x
42 + a - b                 // (42 + a) - b
23 + 3*x[i]                // 23 + (3 * x[i])
x <= f()                   // x <= f()
^a >> b                    // (^a) >> b
f() || g()                 // f() || g()
x == y+1 && <-chanInt > 0  // (x == (y+1)) && ((<-chanInt) > 0)
```

---

## ➗ Aritmetik işleçler¶

Aritmetik işleçler sayısal değerlere uygulanır ve ilk operandla aynı türde bir sonuç üretir. Dört standart aritmetik işleç (`+`, `-`, `*`, `/`) tamsayı, kayan nokta ve kompleks türlere uygulanır; `+` ayrıca string’lere de uygulanır. Bit düzeyi mantıksal ve kaydırma işleçleri yalnızca tamsayılara uygulanır. 

```text
+    toplam                 tamsayılar, float'lar, kompleks değerler, string'ler
-    fark                   tamsayılar, float'lar, kompleks değerler
*    çarpım                 tamsayılar, float'lar, kompleks değerler
/    bölüm                  tamsayılar, float'lar, kompleks değerler
%    kalan                  tamsayılar

&    bit düzeyi AND         tamsayılar
|    bit düzeyi OR          tamsayılar
^    bit düzeyi XOR         tamsayılar
&^   bit temizleme (AND NOT) tamsayılar

<<   sola kaydırma          tamsayı << tamsayı >= 0
>>   sağa kaydırma          tamsayı >> tamsayı >= 0
```

Operand türü bir tür parametresiyse, işleç o type set’teki her türe uygulanabilir olmalıdır. Operandlar, tür parametresinin instantiate edildiği tür argümanının değerleri olarak temsil edilir ve işlem, o tür argümanının hassasiyetiyle hesaplanır. Örneğin: 

```go
func dotProduct[F ~float32|~float64](v1, v2 []F) F {
	var s F
	for i, x := range v1 {
		y := v2[i]
		s += x * y
	}
	return s
}
```

`x * y` çarpımı ve `s += x * y` toplaması, `F` için tür argümanına bağlı olarak sırasıyla `float32` veya `float64` hassasiyetinde hesaplanır.

### 🧮 Tamsayı işleçleri¶

İki tamsayı değeri `x` ve `y` için, tamsayı bölümü `q = x / y` ve kalan `r = x % y` şu ilişkileri sağlar: 

```text
x = q*y + r  ve  |r| < |y|
```

Burada `x / y`, sıfıra doğru kırpılır (“truncated division”). 

```text
 x     y     x / y     x % y
 5     3       1         2
-5     3      -1        -2
 5    -3      -1         2
-5    -3       1        -2
```

Bu kuralın tek istisnası şudur: Bölünen `x`, `x`’in `int` türü için en küçük (en negatif) değer ise, iki’nin tümleyeni (*two’s complement*) tamsayı taşması nedeniyle `q = x / -1` bölümü `x`’e eşittir (ve `r = 0`). 

```text
                         x, q
int8                     -128
int16                  -32768
int32             -2147483648
int64    -9223372036854775808
```

Bölen bir sabitse sıfır olmamalıdır. Bölen çalışma zamanında sıfırsa *panic* oluşur. Bölünen negatif değilse ve bölen 2’nin sabit bir kuvvetiyse, bölme sağa kaydırma ile, kalan hesaplama ise bit düzeyi AND ile değiştirilebilir: 

```text
 x     x / 4     x % 4     x >> 2     x & 3
 11      2         3         2          3
-11     -2        -3        -3          1
```

Kaydırma işleçleri, sol operandı sağ operandın belirttiği kaydırma sayısı kadar kaydırır; sağ operand negatif olmamalıdır. Kaydırma sayısı çalışma zamanında negatifse *panic* oluşur. Kaydırma işleçleri, sol operand işaretliyse aritmetik kaydırma, işaretsizse mantıksal kaydırma uygular. Kaydırma sayısı için üst sınır yoktur. Kaydırmalar, `n` kaydırma sayısı için operandın 1 ile `n` kez kaydırılması gibi davranır. Sonuç olarak `x << 1`, `x*2` ile aynı; `x >> 1` ise negatif sonsuza doğru kırpılmış `x/2` ile aynıdır. 

Tamsayı operandlar için unary işleçler `+`, `-` ve `^` şu şekilde tanımlıdır: 

```text
+x                          = 0 + x
-x    (negasyon)            = 0 - x
^x    (bit düzeyi tümleme)  = m ^ x  (işaretsiz x için m = "tüm bitler 1")
                              ve işaretli x için m = -1
```

### 🧯 Tamsayı taşması¶

İşaretsiz tamsayılarda `+`, `-`, `*` ve `<<` işlemleri `2^n` modunda hesaplanır; burada `n`, işaretsiz tamsayı türünün bit genişliğidir. Kabaca, taşmada yüksek bitler atılır ve programlar “wrap around” davranışına güvenebilir. 

İşaretli tamsayılarda `+`, `-`, `*`, `/` ve `<<` işlemleri yasal olarak taşabilir ve sonuç değeri; işaretli tamsayı temsili, işlem ve operandlar tarafından deterministik olarak tanımlanır. Taşma *panic* oluşturmaz. Bir derleyici, taşma olmayacağı varsayımıyla kodu optimize edemez; örneğin `x < x + 1` ifadesinin her zaman doğru olduğunu varsayamaz. 

### 🌊 Kayan nokta işleçleri¶

Kayan nokta ve kompleks sayılar için `+x`, `x` ile aynı; `-x`, `x`’in negasyonudur. Kayan nokta veya kompleks sayının sıfıra bölünmesinin sonucu, IEEE 754 standardının ötesinde belirtilmemiştir; çalışma zamanında *panic* olup olmayacağı uygulamaya bağlıdır. 

Bir uygulama, birden çok kayan nokta işlemini tek bir birleştirilmiş işleme (*fused operation*) dönüştürebilir, hatta bildirimler arasında da, ve tek tek talimatları yürütüp yuvarlayarak elde edilen değerden farklı bir sonuç üretebilir. Açık bir kayan nokta tür dönüşümü, hedef türün hassasiyetine yuvarlar; böylece bu yuvarlamayı atlayacak bir birleştirmeyi engeller. 

Örneğin, bazı mimariler `x*y + z`’yi `x*y` ara sonucunu yuvarlamadan hesaplayan bir *fused multiply and add* (FMA) talimatı sağlar. Bu örnekler, Go uygulamasının bu talimatı ne zaman kullanabileceğini gösterir: 

```go
// r'yi hesaplarken FMA'ya izin var; çünkü x*y açıkça yuvarlanmıyor:
r  = x*y + z
r  = z;   r += x*y
t  = x*y; r = t + z
*p = x*y; r = *p + z
r  = x*y + float64(z)

// r'yi hesaplarken FMA yasak; çünkü x*y'nin yuvarlamasını atlar:
r  = float64(x*y) + z
r  = z; r += float64(x*y)
t  = float64(x*y); r = t + z
```

---

## 🔗 String birleştirme¶

String’ler `+` işleci veya `+=` atama işleci kullanılarak birleştirilebilir: 

```go
s := "hi" + string(c)
s += " and good bye"
```

String toplaması, operandları birleştirerek yeni bir string oluşturur. 


## 🧾 Deyimler¶

Deyimler yürütmeyi kontrol eder. 

```ebnf
Statement  = Declaration | LabeledStmt | SimpleStmt |
             GoStmt | ReturnStmt | BreakStmt | ContinueStmt | GotoStmt |
             FallthroughStmt | Block | IfStmt | SwitchStmt | SelectStmt | ForStmt |
             DeferStmt .

SimpleStmt = EmptyStmt | ExpressionStmt | SendStmt | IncDecStmt | Assignment | ShortVarDecl .
```

---

## 🛑 Sonlandırıcı deyimler¶

Bir *sonlandırıcı deyim (terminating statement)*, bir blok içindeki olağan kontrol akışını kesintiye uğratır. Aşağıdaki deyimler sonlandırıcıdır:

* Bir `"return"` veya `"goto"` deyimi.
* Yerleşik `panic` fonksiyonuna bir çağrı.
* Deyim listesi bir sonlandırıcı deyim ile biten bir blok.
* Şu koşullarda bir `"if"` deyimi:

  * `"else"` dalı vardır ve
  * her iki dal da sonlandırıcı deyimlerdir.
* Şu koşullarda bir `"for"` deyimi:

  * `"for"` deyimine başvuran hiçbir `"break"` deyimi yoktur ve
  * döngü koşulu yoktur ve
  * `"for"` deyimi bir *range clause* kullanmaz.
* Şu koşullarda bir `"switch"` deyimi:

  * `"switch"` deyimine başvuran hiçbir `"break"` deyimi yoktur,
  * bir *default* case vardır ve
  * *default* dâhil her case’teki deyim listeleri bir sonlandırıcı deyimle veya (muhtemelen etiketli) bir `"fallthrough"` deyimiyle biter.
* Şu koşullarda bir `"select"` deyimi:

  * `"select"` deyimine başvuran hiçbir `"break"` deyimi yoktur ve
  * her case’teki deyim listeleri (varsa default dâhil) bir sonlandırıcı deyimle biter.
* Sonlandırıcı bir deyimi etiketleyen bir etiketli deyim.

Diğer tüm deyimler sonlandırıcı değildir.

Bir deyim listesi, boş değilse ve içindeki son boş olmayan deyim sonlandırıcıysa, sonlandırıcı bir deyimle biter.

---

## ⛔ Boş deyimler¶

Boş deyim hiçbir şey yapmaz.

```ebnf
EmptyStmt = .
```

---

## 🏷️ Etiketli deyimler¶

Etiketli bir deyim, bir *goto*, *break* veya *continue* deyiminin hedefi olabilir.

```ebnf
LabeledStmt = Label ":" Statement .
Label       = identifier .
```

```go
Error: log.Panic("error encountered")
```

---

## 🧩 İfade deyimleri¶

Belirli yerleşik fonksiyonların istisnası dışında, fonksiyon ve yöntem çağrıları ile alma (*receive*) işlemleri deyim bağlamında görünebilir. Bu tür deyimler parantezlenebilir.

```ebnf
ExpressionStmt = Expression .
```

Aşağıdaki yerleşik fonksiyonlara deyim bağlamında izin verilmez:

```text
append cap complex imag len make new real
unsafe.Add unsafe.Alignof unsafe.Offsetof unsafe.Sizeof unsafe.Slice unsafe.SliceData unsafe.String unsafe.StringData
```

```go
h(x+y)
f.Close()
<-ch
(<-ch)
len("foo")  // illegal if len is the built-in function
```

---

## 📤 Gönderme deyimleri¶

Bir gönderme deyimi, bir channel üzerinde bir değer gönderir. Channel ifadesi channel türünde olmalı, channel yönü gönderme işlemlerine izin vermeli ve gönderilecek değerin türü channel’ın eleman türüne atanabilir olmalıdır.

```ebnf
SendStmt = Channel "<-" Expression .
Channel  = Expression .
```

İletişim başlamadan önce hem channel hem de değer ifadesi değerlendirilir. İletişim, gönderme işlemi ilerleyene kadar bloklanır. Tamponsuz (*unbuffered*) bir channel’a gönderme, bir alıcı hazırsa ilerleyebilir. Tamponlu (*buffered*) bir channel’a gönderme, tampondaki yer varsa ilerleyebilir. Kapalı bir channel’a gönderme, çalışma zamanında *panic* oluşturarak ilerler. `nil` bir channel’a gönderme sonsuza dek bloklanır.

```go
ch <- 3  // send value 3 to channel ch
```

Eğer channel ifadesinin türü bir tür parametresiyse, type set’teki tüm türler gönderme işlemlerine izin veren channel türleri olmalı, hepsinin eleman türü aynı olmalı ve gönderilecek değerin türü de o eleman türüne atanabilir olmalıdır.

---

## ➕➖ Artırma/Azaltma deyimleri¶

`"++"` ve `"--"` deyimleri, operandlarını tipsiz sabit `1` ile artırır veya azaltır. Bir atamada olduğu gibi operand adreslenebilir olmalı veya bir map indeks ifadesi olmalıdır.

```ebnf
IncDecStmt = Expression ( "++" | "--" ) .
```

Aşağıdaki atama deyimleri anlamsal olarak eşdeğerdir:

```text
IncDec statement    Assignment
x++                 x += 1
x--                 x -= 1
```

---

## 📝 Atama deyimleri¶

Bir atama, bir değişkende saklanan mevcut değeri bir ifadenin belirttiği yeni bir değerle değiştirir. Bir atama deyimi, tek bir değeri tek bir değişkene veya birden fazla değeri eşleşen sayıda değişkene atayabilir.

```ebnf
Assignment = ExpressionList assign_op ExpressionList .

assign_op  = [ add_op | mul_op ] "=" .
```

Her sol taraf operandı adreslenebilir, bir map indeks ifadesi veya (yalnızca `=` atamaları için) boş tanımlayıcı olmalıdır. Operandlar parantezlenebilir.

```go
x = 1
*p = f()
a[i] = 23
(k) = <-ch  // same as: k = <-ch
```

`op` bir ikili aritmetik işleçken `x op= y` biçimindeki bir atama işlemi, `x = x op (y)` ile eşdeğerdir; ancak `x` yalnızca bir kez değerlendirilir. `op=` yapısı tek bir token’dır. Atama işlemlerinde, hem sol hem de sağ ifade listeleri tam olarak bir tek-değerli ifade içermelidir ve sol ifade boş tanımlayıcı olmamalıdır.

```go
a[i] <<= 2
i &^= 1<<n
```

Bir demet (*tuple*) ataması, çok değerli bir işlemin tek tek elemanlarını bir değişken listesine atar. İki biçimi vardır. İlkinde, sağ taraf operandı; bir fonksiyon çağrısı, bir channel veya map işlemi ya da bir tür doğrulaması (*type assertion*) gibi tek bir çok değerli ifadedir. Sol taraftaki operand sayısı değer sayısıyla eşleşmelidir. Örneğin, `f` iki değer döndüren bir fonksiyonsa,

```go
x, y = f()
```

ilk değeri `x`’e, ikinci değeri `y`’ye atar. İkinci biçimde, soldaki operand sayısı sağdaki ifade sayısına eşit olmalıdır; sağdaki her ifade tek-değerlidir ve sağdaki n’inci ifade soldaki n’inci operanda atanır:

```go
one, two, three = '一', '二', '三'
```

Boş tanımlayıcı, bir atamada sağ taraftaki değerleri yok saymanın bir yolunu sağlar:

```go
_ = x       // evaluate x but ignore it
x, _ = f()  // evaluate f() but ignore second result value
```

Atama iki aşamada ilerler. Önce, sol taraftaki indeks ifadelerinin operandları ve pointer dolaylılıkları (seçicilerdeki örtük pointer dolaylılıkları dâhil) ile sağ taraftaki ifadelerin tümü olağan sırayla değerlendirilir. İkinci olarak, atamalar soldan sağa sırayla gerçekleştirilir.

```go
a, b = b, a  // exchange a and b

x := []int{1, 2, 3}
i := 0
i, x[i] = 1, 2  // set i = 1, x[0] = 2

i = 0
x[i], i = 2, 1  // set x[0] = 2, i = 1

x[0], x[0] = 1, 2  // set x[0] = 1, then x[0] = 2 (so x[0] == 2 at end)

x[1], x[3] = 4, 5  // set x[1] = 4, then panic setting x[3] = 5.

type Point struct { x, y int }
var p *Point
x[2], p.x = 6, 7  // set x[2] = 6, then panic setting p.x = 7

i = 2
x = []int{3, 5, 7}
for i, x[i] = range x {  // set i, x[2] = 0, x[0]
	break
}
// after this loop, i == 0 and x is []int{3, 5, 3}
```

Atamalarda, her değer, atandığı operandın türüne atanabilir olmalıdır; şu özel durumlar dışında:

* Her türlü (typed) değer boş tanımlayıcıya atanabilir.
* Tipsiz bir sabit, interface türünde bir değişkene veya boş tanımlayıcıya atanıyorsa, sabit önce varsayılan türüne örtük olarak dönüştürülür.
* Tipsiz bir boolean değer, interface türünde bir değişkene veya boş tanımlayıcıya atanıyorsa, önce `bool` türüne örtük olarak dönüştürülür.
* Bir değer bir değişkene atanırken, yalnızca değişkende saklanan veri değiştirilir. Eğer değer bir referans içeriyorsa, atama referansı kopyalar ancak referans verilen verinin bir kopyasını yapmaz (örneğin bir slice’ın alttaki array’i).

```go
var s1 = []int{1, 2, 3}
var s2 = s1                    // s2 stores the slice descriptor of s1
s1 = s1[:1]                    // s1's length is 1 but it still shares its underlying array with s2
s2[0] = 42                     // setting s2[0] changes s1[0] as well
fmt.Println(s1, s2)            // prints [42] [42 2 3]

var m1 = make(map[string]int)
var m2 = m1                    // m2 stores the map descriptor of m1
m1["foo"] = 42                 // setting m1["foo"] changes m2["foo"] as well
fmt.Println(m2["foo"])         // prints 42
```

---

## 🔀 If deyimleri¶

`"if"` deyimleri, bir boolean ifadenin değerine göre iki daldan birinin koşullu yürütümünü belirtir. İfade `true` değerlendirilirse `"if"` dalı yürütülür; aksi hâlde, varsa `"else"` dalı yürütülür.

```ebnf
IfStmt = "if" [ SimpleStmt ";" ] Expression Block [ "else" ( IfStmt | Block ) ] .
```

```go
if x > max {
	x = max
}
```

İfadenin önünde, ifade değerlendirilmeden önce yürütülen bir basit deyim bulunabilir.

```go
if x := f(); x < y {
	return x
} else if x > z {
	return z
} else {
	return y
}
```

---

## 🧭 Switch deyimleri¶

`"switch"` deyimleri çok yollu yürütüm sağlar. Bir ifade veya tür, hangi dalın yürütüleceğini belirlemek için `"switch"` içindeki `"case"`’lerle karşılaştırılır.

```ebnf
SwitchStmt = ExprSwitchStmt | TypeSwitchStmt .
```

İki biçimi vardır: *expression switch* ve *type switch*. Bir expression switch’te, case’ler switch ifadesinin değerine karşı karşılaştırılan ifadeler içerir. Bir type switch’te, case’ler özel olarak işaretlenmiş bir switch ifadesinin türüyle karşılaştırılan türleri içerir. Switch ifadesi bir switch deyiminde tam olarak bir kez değerlendirilir.

---

## 🧩 Expression switch’ler¶

Bir expression switch’te, switch ifadesi değerlendirilir ve case ifadeleri (sabit olmak zorunda değildir) soldan sağa ve yukarıdan aşağıya değerlendirilir; switch ifadesine eşit olan ilk case, ilgili case’in deyimlerinin yürütümünü tetikler; diğer case’ler atlanır. Hiçbir case eşleşmezse ve bir `"default"` case varsa, onun deyimleri yürütülür. En fazla bir default case olabilir ve `"switch"` deyiminin içinde herhangi bir yerde bulunabilir. Eksik bir switch ifadesi, boolean değeri `true` ile eşdeğerdir.

```ebnf
ExprSwitchStmt = "switch" [ SimpleStmt ";" ] [ Expression ] "{" { ExprCaseClause } "}" .
ExprCaseClause = ExprSwitchCase ":" StatementList .
ExprSwitchCase = "case" ExpressionList | "default" .
```

Switch ifadesi tipsiz bir sabite değerlendirilirse, önce varsayılan türüne örtük olarak dönüştürülür. Önceden bildirilmiş tipsiz değer `nil`, switch ifadesi olarak kullanılamaz. Switch ifadesinin türü karşılaştırılabilir olmalıdır.

Bir case ifadesi tipsizse, önce switch ifadesinin türüne örtük olarak dönüştürülür. Her (muhtemelen dönüştürülmüş) case ifadesi `x` ve switch ifadesinin değeri `t` için, `x == t` geçerli bir karşılaştırma olmalıdır.

Başka bir deyişle, switch ifadesi; açık tür olmadan geçici bir `t` değişkeni bildirip başlatmak için kullanılıyormuş gibi ele alınır; her case ifadesi `x`, eşitlik testi için bu `t` değeriyle sınanır.

Bir case veya default clause içinde, son boş olmayan deyim; kontrolün bu clause’un sonundan bir sonraki clause’un ilk deyimine akmasını belirtmek için (muhtemelen etiketli) bir `"fallthrough"` deyimi olabilir. Aksi hâlde kontrol `"switch"` deyiminin sonuna akar. `"fallthrough"` deyimi, expression switch’te son clause hariç tüm clause’ların son deyimi olarak görünebilir.

Switch ifadesinin önünde, ifade değerlendirilmeden önce yürütülen bir basit deyim bulunabilir.

```go
switch tag {
default: s3()
case 0, 1, 2, 3: s1()
case 4, 5, 6, 7: s2()
}

switch x := f(); {  // missing switch expression means "true"
case x < 0: return -x
default: return x
}

switch {
case x < y: f1()
case x < z: f2()
case x == 4: f3()
}
```

Uygulama kısıtlaması: Bir derleyici, aynı sabite değerlendirilmiş birden çok case ifadesini yasaklayabilir. Örneğin, mevcut derleyiciler case ifadelerinde yinelenen tamsayı, kayan nokta veya string sabitlerine izin vermez.

---

## 🧷 Type switch’ler¶

Bir type switch, değerler yerine türleri karşılaştırır. Diğer yönleriyle bir expression switch’e benzer. `type` anahtar sözcüğünü kullanan özel bir tür doğrulaması (*type assertion*) biçimindeki bir switch ifadesiyle işaretlenir; gerçek bir tür yerine `type` kullanılır:

```go
switch x.(type) {
// cases
}
```

Case’ler, ifadenin `x` dinamik türüne karşı gerçek türleri `T` eşleştirir. Tür doğrulamalarında olduğu gibi, `x` interface türünde olmalı, fakat bir tür parametresi olmamalıdır ve bir case’te listelenen her interface olmayan tür `T`, `x`’in türünü implement etmelidir. Type switch case’lerinde listelenen türlerin tümü birbirinden farklı olmalıdır.

```ebnf
TypeSwitchStmt  = "switch" [ SimpleStmt ";" ] TypeSwitchGuard "{" { TypeCaseClause } "}" .
TypeSwitchGuard = [ identifier ":=" ] PrimaryExpr "." "(" "type" ")" .
TypeCaseClause  = TypeSwitchCase ":" StatementList .
TypeSwitchCase  = "case" TypeList | "default" .
```

TypeSwitchGuard, kısa değişken bildirimi içerebilir. Bu biçim kullanıldığında, değişken her clause’un örtük bloğunda, TypeSwitchCase’in sonunda bildirilir. Tam olarak bir tür içeren case’lerde değişken o türe sahiptir; aksi hâlde değişken, TypeSwitchGuard’daki ifadenin türüne sahiptir.

Bir tür yerine, case önceden bildirilmiş tanımlayıcı `nil`’i kullanabilir; bu case, TypeSwitchGuard’daki ifade `nil` bir interface değeri olduğunda seçilir. En fazla bir `nil` case olabilir.

Türü `interface{}` olan bir `x` ifadesi verildiğinde, aşağıdaki type switch:

```go
switch i := x.(type) {
case nil:
	printString("x is nil")                // type of i is type of x (interface{})
case int:
	printInt(i)                            // type of i is int
case float64:
	printFloat64(i)                        // type of i is float64
case func(int) float64:
	printFunction(i)                       // type of i is func(int) float64
case bool, string:
	printString("type is bool or string")  // type of i is type of x (interface{})
default:
	printString("don't know the type")     // type of i is type of x (interface{})
}
```

şu şekilde yeniden yazılabilir:

```go
v := x  // x is evaluated exactly once
if v == nil {
	i := v                                 // type of i is type of x (interface{})
	printString("x is nil")
} else if i, isInt := v.(int); isInt {
	printInt(i)                            // type of i is int
} else if i, isFloat64 := v.(float64); isFloat64 {
	printFloat64(i)                        // type of i is float64
} else if i, isFunc := v.(func(int) float64); isFunc {
	printFunction(i)                       // type of i is func(int) float64
} else {
	_, isBool := v.(bool)
	_, isString := v.(string)
	if isBool || isString {
		i := v                         // type of i is type of x (interface{})
		printString("type is bool or string")
	} else {
		i := v                         // type of i is type of x (interface{})
		printString("don't know the type")
	}
}
```

Bir tür parametresi veya jenerik bir tür, bir case içinde tür olarak kullanılabilir. Instantiate edildiğinde bu tür, switch’teki başka bir girdiyi yinelerse, ilk eşleşen case seçilir.

```go
func f[P any](x any) int {
	switch x.(type) {
	case P:
		return 0
	case string:
		return 1
	case []P:
		return 2
	case []byte:
		return 3
	default:
		return 4
	}
}

var v1 = f[string]("foo")   // v1 == 0
var v2 = f[byte]([]byte{})  // v2 == 2
```

Type switch guard’ın önünde, guard değerlendirilmeden önce yürütülen bir basit deyim bulunabilir.

`"fallthrough"` deyimine bir type switch’te izin verilmez.

---

## 🔁 For deyimleri¶

Bir `"for"` deyimi, bir bloğun yinelenen yürütümünü belirtir. Üç biçim vardır: yineleme tek bir koşulla, bir `"for"` clause ile veya bir `"range"` clause ile kontrol edilebilir.

```ebnf
ForStmt   = "for" [ Condition | ForClause | RangeClause ] Block .
Condition = Expression .
```

---

## ✅ Tek koşullu for deyimleri¶

En basit biçiminde, bir `"for"` deyimi, bir boolean koşul `true` değerlendirdikçe bir bloğun yinelenen yürütümünü belirtir. Koşul her yinelemeden önce değerlendirilir. Koşul yoksa, boolean değeri `true` ile eşdeğerdir.

```go
for a < b {
	a *= 2
}
```

---

## 🧷 For clause’lu for deyimleri¶

Bir `"for"` deyimi bir ForClause ile kontrol ediliyorsa, yine koşul tarafından yönetilir; ancak ek olarak atama, artırma veya azaltma deyimi gibi bir *init* ve bir *post* deyimi belirtebilir. Init deyimi kısa değişken bildirimi olabilir, fakat post deyimi olmamalıdır.

```ebnf
ForClause = [ InitStmt ] ";" [ Condition ] ";" [ PostStmt ] .
InitStmt  = SimpleStmt .
PostStmt  = SimpleStmt .
```

```go
for i := 0; i < 10; i++ {
	f(i)
}
```

Boş değilse, init deyimi ilk yinelemenin koşulu değerlendirilmeden önce bir kez yürütülür; post deyimi ise bloğun her yürütümünden sonra (ve yalnızca blok yürütüldüyse) yürütülür. ForClause’un herhangi bir elemanı boş olabilir; ancak yalnızca bir koşul varsa semikolonlar gerekmez. Koşul yoksa, boolean değeri `true` ile eşdeğerdir.

```text
for cond { S() }    is the same as    for ; cond ; { S() }
for      { S() }    is the same as    for true     { S() }
```

Her yinelemenin kendi ayrı bildirilen değişkeni (veya değişkenleri) vardır [Go 1.22]. İlk yinelemenin kullandığı değişken init deyimi tarafından bildirilir. Sonraki her yinelemenin kullandığı değişken, post deyimini yürütmeden önce örtük olarak bildirilir ve o anda bir önceki yinelemenin değişken değerine başlatılır.

```go
var prints []func()
for i := 0; i < 5; i++ {
	prints = append(prints, func() { println(i) })
	i++
}
for _, p := range prints {
	p()
}
prints

1
3
5
```

[Go 1.22] öncesinde, yinelemeler ayrı değişkenler yerine tek bir değişken kümesini paylaşır. Bu durumda, yukarıdaki örnek şunu yazdırır:

```text
6
6
6
```

---

## 🧭 Range clause’lu for deyimleri¶

Bir `"range"` clause’lu `"for"` deyimi; bir array, slice, string veya map içindeki tüm girdiler; bir channel’dan alınan değerler; sıfırdan bir üst sınıra kadar tamsayı değerler [Go 1.22]; veya bir iterator fonksiyonunun yield fonksiyonuna geçirilen değerler [Go 1.23] üzerinde yineleme yapar. Her giriş için, varsa yineleme değerlerini ilgili yineleme değişkenlerine atar ve ardından bloğu yürütür.

```ebnf
RangeClause = [ ExpressionList "=" | IdentifierList ":=" ] "range" Expression .
```

`"range"` clause’un sağındaki ifadeye *range expression* denir; bu, bir array, array’e pointer, slice, string, map, alma işlemlerine izin veren bir channel, bir tamsayı veya belirli bir imzaya sahip bir fonksiyon olabilir (aşağıya bakın). Bir atamada olduğu gibi, varsa soldaki operandlar adreslenebilir veya map indeks ifadeleri olmalıdır; bunlar yineleme değişkenlerini belirtir. Range expression bir fonksiyon ise, yineleme değişkenlerinin azami sayısı fonksiyon imzasına bağlıdır. Range expression bir channel veya tamsayıysa, en fazla bir yineleme değişkenine izin verilir; aksi hâlde en fazla iki olabilir. Son yineleme değişkeni boş tanımlayıcıysa, range clause bu tanımlayıcı olmadan yazılmış aynı clause ile eşdeğerdir.

Range expression `x`, döngü başlamadan önce değerlendirilir; bir istisna dışında: en fazla bir yineleme değişkeni varsa ve `x` veya `len(x)` sabitse, range expression değerlendirilmez.

Soldaki fonksiyon çağrıları, yineleme başına bir kez değerlendirilir. Her yineleme için, ilgili yineleme değişkenleri mevcutsa yineleme değerleri aşağıdaki gibi üretilir:

```text
Range expression                                       1st value                2nd value

array or slice      a  [n]E, *[n]E, or []E             index    i  int          a[i]       E
string              s  string type                     index    i  int          see below  rune
map                 m  map[K]V                         key      k  K            m[k]       V
channel             c  chan E, <-chan E                element  e  E
integer value       n  integer type, or untyped int    value    i  see below
function, 0 values  f  func(func() bool)
function, 1 value   f  func(func(V) bool)              value    v  V
function, 2 values  f  func(func(K, V) bool)           key      k  K            v          V
```

Array, array’e pointer veya slice değeri `a` için indeks yineleme değerleri, eleman indeks 0’dan başlayarak artan sırada üretilir. En fazla bir yineleme değişkeni varsa, range döngüsü 0’dan `len(a)-1`’e kadar yineleme değerleri üretir ve array veya slice’ın kendisini indekslemez. `nil` bir slice için yineleme sayısı 0’dır.

String değeri için `"range"` clause, string içindeki Unicode kod noktaları üzerinde bayt indeks 0’dan başlayarak yineleme yapar. Ardışık yinelemelerde indeks değeri, string içindeki ardışık UTF-8 kodlanmış kod noktalarının ilk baytının indeksi olur ve türü `rune` olan ikinci değer, karşılık gelen kod noktasının değeridir. Yineleme geçersiz bir UTF-8 dizisiyle karşılaşırsa, ikinci değer Unicode yerine koyma karakteri olan `0xFFFD` olur ve bir sonraki yineleme string içinde tek bir bayt ilerler.

Map’ler üzerinde yineleme sırası belirtilmemiştir ve bir yinelemeden diğerine aynı olacağı garanti edilmez. Yineleme sırasında henüz ulaşılmamış bir map girdisi silinirse, karşılık gelen yineleme değeri üretilmez. Yineleme sırasında bir map girdisi oluşturulursa, o girdi yineleme sırasında üretilebilir veya atlanabilir. Seçim, oluşturulan her girdi için ve yinelemeler arasında değişebilir. Map `nil` ise, yineleme sayısı 0’dır.

Channel’lar için yineleme değerleri, channel kapatılana kadar channel’a gönderilen ardışık değerlerdir. Channel `nil` ise, range expression sonsuza dek bloklanır.

Tamsayı değeri `n` için (burada `n` tamsayı türünde veya tipsiz bir tamsayı sabitidir), 0’dan `n-1`’e kadar yineleme değerleri artan sırada üretilir. `n` tamsayı türündeyse, yineleme değerleri aynı türe sahiptir. Aksi hâlde, `n`’in türü, yineleme değişkenine atanıyormuş gibi belirlenir. Özellikle: yineleme değişkeni önceden var olan bir değişkense, yineleme değerlerinin türü o yineleme değişkeninin türüdür ve bu tür tamsayı türünde olmalıdır. Aksi hâlde, yineleme değişkeni `"range"` clause tarafından bildiriliyorsa veya yoksa, yineleme değerlerinin türü `n` için varsayılan türdür. `n <= 0` ise, döngü hiç yineleme yapmaz.

Fonksiyon `f` için, yineleme; `f`’yi, argüman olarak yeni, sentezlenmiş bir yield fonksiyonu ile çağırarak ilerler. `f` dönmeden önce yield çağrılırsa, yield’e verilen argümanlar döngü gövdesini bir kez yürütmek için yineleme değerleri olur. Her ardışık döngü yinelemesinden sonra yield `true` döndürür ve döngüyü sürdürmek için yeniden çağrılabilir. Döngü gövdesi sonlandırmadıkça, `"range"` clause bu şekilde her yield çağrısı için `f` dönene kadar yineleme değerleri üretmeye devam eder. Döngü gövdesi sonlanırsa (örneğin bir `break` deyimiyle), yield `false` döndürür ve tekrar çağrılmamalıdır.

Yineleme değişkenleri, kısa değişken bildirimi (`:=`) biçimi kullanılarak `"range"` clause tarafından bildirilebilir. Bu durumda kapsamları `"for"` deyiminin bloğudur ve her yinelemenin kendi yeni değişkenleri vardır [Go 1.22] (bkz. ForClause’lu `"for"` deyimleri). Değişkenler, kendi yineleme değerlerinin türlerine sahiptir.

Yineleme değişkenleri `"range"` clause tarafından açıkça bildirilmemişse, önceden var olmalıdır. Bu durumda yineleme değerleri, bir atama deyiminde olduğu gibi ilgili değişkenlere atanır.

```go
var testdata *struct {
	a *[7]int
}
for i, _ := range testdata.a {
	// testdata.a is never evaluated; len(testdata.a) is constant
	// i ranges from 0 to 6
	f(i)
}

var a [10]string
for i, s := range a {
	// type of i is int
	// type of s is string
	// s == a[i]
	g(i, s)
}

var key string
var val interface{}  // element type of m is assignable to val
m := map[string]int{"mon":0, "tue":1, "wed":2, "thu":3, "fri":4, "sat":5, "sun":6}
for key, val = range m {
	h(key, val)
}
// key == last map key encountered in iteration
// val == map[key]

var ch chan Work = producer()
for w := range ch {
	doWork(w)
}

// empty a channel
for range ch {}

// call f(0), f(1), ... f(9)
for i := range 10 {
	// type of i is int (default type for untyped constant 10)
	f(i)
}

// invalid: 256 cannot be assigned to uint8
var u uint8
for u = range 256 {
}

// invalid: 1e3 is a floating-point constant
for range 1e3 {
}

// fibo generates the Fibonacci sequence
fibo := func(yield func(x int) bool) {
	f0, f1 := 0, 1
	for yield(f0) {
		f0, f1 = f1, f0+f1
	}
}

// print the Fibonacci numbers below 1000:
for x := range fibo {
	if x >= 1000 {
		break
	}
	fmt.Printf("%d ", x)
}
// output: 0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987

// iteration support for a recursive tree data structure
type Tree[K cmp.Ordered, V any] struct {
	left, right *Tree[K, V]
	key         K
	value       V
}

func (t *Tree[K, V]) walk(yield func(key K, val V) bool) bool {
	return t == nil || t.left.walk(yield) && yield(t.key, t.value) && t.right.walk(yield)
}

func (t *Tree[K, V]) Walk(yield func(key K, val V) bool) {
	t.walk(yield)
}

// walk tree t in-order
var t Tree[string, int]
for k, v := range t.Walk {
	// process k, v
}
```

Range expression’ın türü bir tür parametresiyse, type set’teki tüm türler aynı alttaki türe sahip olmalı ve range expression o tür için geçerli olmalıdır; veya type set channel türleri içeriyorsa, yalnızca özdeş eleman türlerine sahip channel türlerini içermeli ve tüm channel türleri alma işlemlerine izin vermelidir.

---

## 🚀 Go deyimleri¶

Bir `"go"` deyimi, bir fonksiyon çağrısının yürütümünü aynı adres uzayı içinde bağımsız bir eşzamanlı kontrol iş parçacığı (veya *goroutine*) olarak başlatır.

```ebnf
GoStmt = "go" Expression .
```

İfade bir fonksiyon veya yöntem çağrısı olmalıdır; parantezlenemez. Yerleşik fonksiyon çağrıları, ifade deyimleri için geçerli olduğu şekilde kısıtlanmıştır.

Fonksiyon değeri ve parametreler, çağıran goroutine’de olağan şekilde değerlendirilir; ancak normal çağrıdan farklı olarak, program yürütümü çağrılan fonksiyonun tamamlanmasını beklemez. Bunun yerine fonksiyon, yeni bir goroutine içinde bağımsız olarak yürütülmeye başlar. Fonksiyon sonlandığında, goroutine’i de sonlanır. Fonksiyonun herhangi bir dönüş değeri varsa, fonksiyon tamamlandığında bunlar yok sayılır.

```go
go Server()
go func(ch chan<- bool) { for { sleep(10); ch <- true }} (c)
```

---

## 🎛️ Select deyimleri¶

Bir `"select"` deyimi, bir dizi olası gönderme veya alma işlemi arasından hangisinin ilerleyeceğini seçer. `"switch"` deyimine benzer görünür; ancak case’lerin tümü iletişim işlemlerine atıf yapar.

```ebnf
SelectStmt = "select" "{" { CommClause } "}" .
CommClause = CommCase ":" StatementList .
CommCase   = "case" ( SendStmt | RecvStmt ) | "default" .
RecvStmt   = [ ExpressionList "=" | IdentifierList ":=" ] RecvExpr .
RecvExpr   = Expression .
```

Bir RecvStmt içeren case, bir RecvExpr sonucunu bir veya iki değişkene atayabilir; bu değişkenler kısa değişken bildirimiyle bildirilebilir. RecvExpr, (muhtemelen parantezlenmiş) bir alma işlemi olmalıdır. En fazla bir default case olabilir ve case listesi içinde herhangi bir yerde bulunabilir.

Bir `"select"` deyiminin yürütümü birkaç adımda ilerler:

1. Deyimdeki tüm case’ler için, alma işlemlerinin channel operandları ile gönderme deyimlerinin channel ve sağ-taraf ifadeleri, `"select"` deyimine girildiğinde kaynak sırasına göre tam olarak bir kez değerlendirilir. Sonuç; alınacak veya gönderilecek bir channel kümesi ve gönderilecek karşılık gelen değerler kümesidir. Bu değerlendirmedeki yan etkiler, hangi (varsa) iletişim işlemi seçilirse seçilsin gerçekleşir. Kısa değişken bildirimi veya atama içeren bir RecvStmt’in sol taraf ifadeleri henüz değerlendirilmez.
2. Bir veya daha fazla iletişim işlemi ilerleyebiliyorsa, ilerleyebilenlerden biri, uniform pseudo-random seçimle seçilir. Aksi hâlde, bir default case varsa o seçilir. Default case yoksa, `"select"` deyimi en az bir iletişim işlemi ilerleyebilene kadar bloklanır.
3. Seçilen case default değilse, ilgili iletişim işlemi yürütülür.
4. Seçilen case bir kısa değişken bildirimi veya bir atama içeren RecvStmt ise, sol taraf ifadeleri değerlendirilir ve alınan değer (veya değerler) atanır.
5. Seçilen case’in deyim listesi yürütülür.

`nil` channel’larda iletişim asla ilerleyemeyeceğinden, yalnızca `nil` channel’lardan oluşan ve default case içermeyen bir select sonsuza dek bloklanır.

```go
var a []int
var c, c1, c2, c3, c4 chan int
var i1, i2 int
select {
case i1 = <-c1:
	print("received ", i1, " from c1\n")
case c2 <- i2:
	print("sent ", i2, " to c2\n")
case i3, ok := (<-c3):  // same as: i3, ok := <-c3
	if ok {
		print("received ", i3, " from c3\n")
	} else {
		print("c3 is closed\n")
	}
case a[f()] = <-c4:
	// same as:
	// case t := <-c4
	//	a[f()] = t
default:
	print("no communication\n")
}

for {  // send random sequence of bits to c
	select {
	case c <- 0:  // note: no statement, no fallthrough, no folding of cases
	case c <- 1:
	}
}

select {}  // block forever
```

---

## ↩️ Return deyimleri¶

Bir `F` fonksiyonu içinde bir `"return"` deyimi, `F`’nin yürütümünü sonlandırır ve isteğe bağlı olarak bir veya daha fazla sonuç değeri sağlar. `F` tarafından ertelenmiş (*deferred*) tüm fonksiyonlar, `F` çağırana dönmeden önce yürütülür.

```ebnf
ReturnStmt = "return" [ ExpressionList ] .
```

Sonuç türü olmayan bir fonksiyonda, bir `"return"` deyimi herhangi bir sonuç değeri belirtmemelidir.

```go
func noResult() {
	return
}
```

Sonuç türü olan bir fonksiyondan değer döndürmenin üç yolu vardır:

1. Dönüş değeri veya değerleri `"return"` deyiminde açıkça listelenebilir. Her ifade tek-değerli olmalı ve fonksiyonun sonuç türünün karşılık gelen elemanına atanabilir olmalıdır.

```go
func simpleF() int {
	return 2
}

func complexF1() (re float64, im float64) {
	return -7.0, -4.0
}
```

2. `"return"` deyimindeki ifade listesi, çok değerli bir fonksiyona tek bir çağrı olabilir. Etkisi, o fonksiyonun döndürdüğü her değerin ilgili türde geçici bir değişkene atanması, ardından bu değişkenleri listeleyen bir `"return"` deyiminin yürütülmesi gibidir; bu noktada birinci durumun kuralları geçerlidir.

```go
func complexF2() (re float64, im float64) {
	return complexF1()
}
```

3. Fonksiyonun sonuç türü sonuç parametrelerine adlar veriyorsa, ifade listesi boş olabilir. Sonuç parametreleri sıradan yerel değişkenler gibi davranır ve fonksiyon gerektiğinde bunlara değer atayabilir. `"return"` deyimi bu değişkenlerin değerlerini döndürür.

```go
func complexF3() (re float64, im float64) {
	re = 7.0
	im = 4.0
	return
}

func (devnull) Write(p []byte) (n int, _ error) {
	n = len(p)
	return
}
```

Nasıl bildirilmiş olursa olsun, tüm sonuç değerleri fonksiyona girişte kendi türlerinin sıfır değerleriyle başlatılır. Sonuç belirten bir `"return"` deyimi, ertelenmiş fonksiyonlar yürütülmeden önce sonuç parametrelerini ayarlar.

Uygulama kısıtlaması: Bir derleyici, dönüş noktasında sonuç parametresiyle aynı adda farklı bir varlık (sabit, tür veya değişken) kapsamdaysa, `"return"` deyiminde boş ifade listesini yasaklayabilir.

```go
func f(n int) (res int, err error) {
	if _, err := f(n-1); err != nil {
		return  // invalid return statement: err is shadowed
	}
	return
}
```

---

## 🧨 Break deyimleri¶

Bir `"break"` deyimi, aynı fonksiyon içinde en içteki `"for"`, `"switch"` veya `"select"` deyiminin yürütümünü sonlandırır.

```ebnf
BreakStmt = "break" [ Label ] .
```

Bir etiket varsa, bu, kapsayan bir `"for"`, `"switch"` veya `"select"` deyiminin etiketi olmalıdır ve sonlandırılan yürütüm o deyiminki olur.

```go
OuterLoop:
	for i = 0; i < n; i++ {
		for j = 0; j < m; j++ {
			switch a[i][j] {
			case nil:
				state = Error
				break OuterLoop
			case item:
				state = Found
				break OuterLoop
			}
		}
	}
```

---

## ⏭️ Continue deyimleri¶

Bir `"continue"` deyimi, aynı fonksiyon içinde en içteki kapsayan `"for"` döngüsünün bir sonraki yinelemesine, kontrolü döngü bloğunun sonuna ilerleterek başlar. `"for"` döngüsü aynı fonksiyon içinde olmalıdır.

```ebnf
ContinueStmt = "continue" [ Label ] .
```

Bir etiket varsa, bu, kapsayan bir `"for"` deyiminin etiketi olmalıdır ve yürütümü ilerletilen döngü o olur.

```go
RowLoop:
	for y, row := range rows {
		for x, data := range row {
			if data == endOfRow {
				continue RowLoop
			}
			row[x] = data + bias(x, y)
		}
	}
```

---

## 🧭 Goto deyimleri¶

Bir `"goto"` deyimi, kontrolü aynı fonksiyon içinde karşılık gelen etiketin bulunduğu deyime aktarır.

```ebnf
GotoStmt = "goto" Label .
```

```go
goto Error
```

`"goto"` deyimini yürütmek, goto noktasında kapsamda olmayan değişkenlerin sonradan kapsama girmesine neden olmamalıdır. Örneğin, şu örnek:

```go
	goto L  // BAD
	v := 3
L:
```

hatalıdır; çünkü `L` etiketine atlama, `v`’nin oluşturulmasını atlar.

Bir blok dışındaki `"goto"` deyimi, o bloğun içindeki bir etikete atlayamaz. Örneğin, şu örnek:

```go
if n%2 == 1 {
	goto L1
}
for n > 0 {
	f()
	n--
L1:
	f()
	n--
}
```

hatalıdır; çünkü `L1` etiketi `"for"` deyiminin bloğu içindedir, fakat `"goto"` dışarıdadır.

---

## 🪂 Fallthrough deyimleri¶

Bir `"fallthrough"` deyimi, bir expression `"switch"` deyiminde kontrolü bir sonraki case clause’un ilk deyimine aktarır. Yalnızca bu tür bir clause içinde, son boş olmayan deyim olarak kullanılabilir.

```ebnf
FallthroughStmt = "fallthrough" .
```

---

## 🧯 Defer deyimleri¶

Bir `"defer"` deyimi, yürütümü çevreleyen fonksiyon döndüğü anda ertelenen bir fonksiyonu çağırır; bu, çevreleyen fonksiyonun bir return deyimi yürütmesiyle, fonksiyon gövdesinin sonuna ulaşmasıyla veya ilgili goroutine’in panic’e girmesiyle olabilir.

```ebnf
DeferStmt = "defer" Expression .
```

İfade bir fonksiyon veya yöntem çağrısı olmalıdır; parantezlenemez. Yerleşik fonksiyon çağrıları, ifade deyimleri için geçerli olduğu şekilde kısıtlanmıştır.

Her `"defer"` deyimi yürütüldüğünde, fonksiyon değeri ve çağrı parametreleri olağan şekilde değerlendirilir ve her seferinde yeniden saklanır; ancak gerçek fonksiyon çağrılmaz. Bunun yerine, ertelenmiş fonksiyonlar, çevreleyen fonksiyon dönmeden hemen önce, ertelendikleri sıranın tersiyle çağrılır. Yani, çevreleyen fonksiyon açık bir return deyimiyle dönerse, ertelenmiş fonksiyonlar; return deyiminin sonuç parametrelerini ayarlamasından sonra, fakat çağırana dönmeden önce yürütülür. Ertelenen fonksiyon değeri `nil`’e değerlendirilirse, yürütüm panic’e, fonksiyon çağrıldığında girer; `"defer"` deyimi yürütülürken değil.

Örneğin, ertelenen fonksiyon bir fonksiyon literal’i ise ve çevreleyen fonksiyonun, literal içinde kapsamda olan adlandırılmış sonuç parametreleri varsa, ertelenmiş fonksiyon bu sonuç parametrelerine erişebilir ve döndürülmeden önce onları değiştirebilir. Ertelenmiş fonksiyonun dönüş değerleri varsa, fonksiyon tamamlandığında yok sayılır. (Ayrıca panic’leri ele alma bölümüne bakın.)

```go
lock(l)
defer unlock(l)  // unlocking happens before surrounding function returns

// prints 3 2 1 0 before surrounding function returns
for i := 0; i <= 3; i++ {
	defer fmt.Print(i)
}

// f returns 42
func f() (result int) {
	defer func() {
		// result is accessed after it was set to 6 by the return statement
		result *= 7
	}()
	return 6
}
```


## 🧰 Yerleşik fonksiyonlar¶

Yerleşik fonksiyonlar önceden bildirilmiştir. Diğer tüm fonksiyonlar gibi çağrılırlar; ancak bazılarında ilk argüman olarak bir ifade yerine bir tür kabul edilir.

Yerleşik fonksiyonların standart Go türleri yoktur; bu yüzden yalnızca çağrı ifadelerinde görünebilirler; fonksiyon değeri olarak kullanılamazlar. 

---

## 🧩 Slice’lara ekleme ve kopyalama¶

Yerleşik `append` ve `copy` fonksiyonları, yaygın slice işlemlerine yardımcı olur. Her iki fonksiyonda da sonuç, argümanların işaret ettiği bellek alanlarının çakışıp çakışmamasından bağımsızdır. 

### ➕ `append`

Variadic olan `append`, `S` türünde bir slice `s`’e sıfır veya daha fazla `x` değeri ekler ve elde edilen slice’ı yine `S` türünde döndürür. `x` değerleri, `S`’nin eleman türü `E` olmak üzere `...E` türünde bir parametreye geçirilir ve ilgili parametre geçirme kuralları uygulanır. Özel bir durum olarak, `append`; ilk argümanı `[]byte` türüne atanabilir ve ikinci argümanı `string` türünde olmak üzere, ardından `...` ile de kabul eder. Bu biçim string’in baytlarını ekler. 

```go
append(s S, x ...E) S  // E, S'nin eleman türüdür
```

`S` bir tür parametresiyse, type set’indeki tüm türler aynı alttaki slice türü `[]E`’ye sahip olmalıdır. 

`s`’in kapasitesi ek değerleri sığdırmaya yetmiyorsa, `append` hem mevcut slice elemanlarını hem de ek değerleri alacak kadar büyük yeni bir alttaki array tahsis eder. Aksi hâlde, `append` mevcut alttaki array’i yeniden kullanır. 

```go
s0 := []int{0, 0}
s1 := append(s0, 2)                // tek eleman ekle         s1 == []int{0, 0, 2}
s2 := append(s1, 3, 5, 7)          // çok eleman ekle         s2 == []int{0, 0, 2, 3, 5, 7}
s3 := append(s2, s0...)            // bir slice ekle          s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}
s4 := append(s3[3:6], s3[2:]...)   // çakışan slice ekle      s4 == []int{3, 5, 7, 2, 3, 5, 7, 0, 0}

var t []interface{}
t = append(t, 42, 3.1415, "foo")   //                          t == []interface{}{42, 3.1415, "foo"}

var b []byte
b = append(b, "bar"...)            // string içeriğini ekle     b == []byte{'b', 'a', 'r'}
```

### 📋 `copy`

`copy`, slice elemanlarını bir kaynak `src`’den bir hedef `dst`’ye kopyalar ve kopyalanan eleman sayısını döndürür. Her iki argümanın eleman türü `E` özdeş olmalı ve `[]E` türünde bir slice’a atanabilir olmalıdır. Kopyalanan eleman sayısı `len(src)` ve `len(dst)` değerlerinin minimumudur. Özel bir durum olarak, `copy`; hedef argümanı `[]byte` türüne atanabilir ve kaynak argümanı `string` türünde olacak şekilde de kabul eder. Bu biçim string’in baytlarını byte slice’a kopyalar. 

```go
copy(dst, src []T) int
copy(dst []byte, src string) int
```

Argümanlardan birinin veya ikisinin türü bir tür parametresiyse, ilgili type set’lerindeki tüm türler aynı alttaki slice türü `[]E`’ye sahip olmalıdır. 

```go
var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
var s = make([]int, 6)
var b = make([]byte, 5)
n1 := copy(s, a[0:])            // n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
n2 := copy(s, s[2:])            // n2 == 4, s == []int{2, 3, 4, 5, 4, 5}
n3 := copy(b, "Hello, World!")  // n3 == 5, b == []byte("Hello")
```

---

## 🧹 Clear¶

Yerleşik `clear` fonksiyonu, map, slice veya tür parametresi türünde bir argüman alır ve tüm elemanları siler veya sıfırlar [Go 1.21]. 

| Çağrı      | Argüman türü    | Sonuç                                                                 |
| ---------- | --------------- | --------------------------------------------------------------------- |
| `clear(m)` | `map[K]T`       | tüm girdileri siler, boş bir map elde edilir (`len(m) == 0`)          |
| `clear(s)` | `[]T`           | `s`’in uzunluğuna kadar tüm elemanları `T`’nin sıfır değerine ayarlar |
| `clear(t)` | tür parametresi | aşağıya bakın                                                         |

Argümanın türü bir tür parametresiyse, type set’teki tüm türler map veya slice olmalı ve `clear`, gerçek tür argümanına karşılık gelen işlemi yapar. Map veya slice `nil` ise `clear` no-op’tur. 

---

## 🔒 Close¶

Bir channel `ch` için yerleşik `close(ch)` fonksiyonu, o channel’a artık değer gönderilmeyeceğini kaydeder. `ch`’nin yalnızca-alma (*receive-only*) channel olması hatadır. Kapalı bir channel’a göndermek veya kapalı bir channel’ı kapatmak çalışma zamanında *panic* oluşturur. `nil` channel’ı kapatmak da çalışma zamanında *panic* oluşturur. `close` çağrısından sonra ve daha önce gönderilmiş tüm değerler alındıktan sonra, alma işlemleri bloklanmadan channel türünün sıfır değerini döndürür. Çok-değerli alma işlemi, alınan değerle birlikte channel’ın kapalı olup olmadığına dair bir gösterge döndürür. 

Argümanın türü bir tür parametresiyse, type set’teki tüm türler aynı eleman türüne sahip channel’lar olmalıdır. Bu channel’lardan herhangi birinin yalnızca-alma channel olması hatadır. 

---

## 🧪 Kompleks sayılarla çalışma¶

Üç fonksiyon kompleks sayıları kurar ve parçalarına ayırır. `complex`, kayan nokta gerçek (*real*) ve sanal (*imaginary*) kısımdan kompleks değer oluşturur; `real` ve `imag`, kompleks değerin gerçek ve sanal kısımlarını çıkarır. 

```go
complex(realPart, imaginaryPart floatT) complexT
real(complexT) floatT
imag(complexT) floatT
```

Argüman ve dönüş türleri birbirine karşılık gelir. `complex` için iki argüman aynı kayan nokta türünde olmalıdır ve dönüş türü, karşılık gelen kayan nokta bileşenlerine sahip kompleks türdür: `float32` argümanları için `complex64`, `float64` argümanları için `complex128`. Argümanlardan biri tipsiz sabitse, önce diğerinin türüne örtük olarak dönüştürülür. Her iki argüman da tipsiz sabitse, kompleks olmayan sayılar olmalı ya da sanal kısımları sıfır olmalıdır; bu durumda dönüş değeri tipsiz kompleks sabittir. 

`real` ve `imag` için argüman kompleks türde olmalı ve dönüş türü karşılık gelen kayan nokta türüdür: `complex64` argümanı için `float32`, `complex128` argümanı için `float64`. Argüman tipsiz sabitse, bir sayı olmalı ve dönüş değeri tipsiz kayan nokta sabitidir. 

`real` ve `imag`, `complex`’in tersini oluşturur; dolayısıyla kompleks tür `Z`’de bir `z` değeri için:

`z == Z(complex(real(z), imag(z)))`. 

Operandların tümü sabitse, dönüş değeri sabittir. 

```go
var a = complex(2, -2)             // complex128
const b = complex(1.0, -1.4)       // tipsiz kompleks sabit 1 - 1.4i
x := float32(math.Cos(math.Pi/2))  // float32
var c64 = complex(5, -x)           // complex64
var s int = complex(1, 0)          // tipsiz kompleks sabit 1 + 0i, int'e dönüştürülebilir
_ = complex(1, 2<<s)               // yasadışı: 2 kayan nokta türünü varsayar, kaydırılamaz
var rl = real(c64)                 // float32
var im = imag(a)                   // float64
const c = imag(b)                  // tipsiz sabit -1.4
_ = imag(3 << s)                   // yasadışı: 3 kompleks türü varsayar, kaydırılamaz
```

Tür parametresi türündeki argümanlara izin verilmez. 

---

## 🗑️ Map elemanlarını silme¶

Yerleşik `delete`, `m` map’inden anahtarı `k` olan elemanı kaldırır. `k` değeri, `m`’in anahtar türüne atanabilir olmalıdır. 

```go
delete(m, k)  // m map'inden m[k] elemanını kaldır
```

`m`’in türü bir tür parametresiyse, type set’teki tüm türler map olmalı ve hepsinin anahtar türleri özdeş olmalıdır. 

Map `m` `nil` ise veya `m[k]` elemanı yoksa, `delete` no-op’tur. 

---

## 📏 Uzunluk ve kapasite¶

Yerleşik `len` ve `cap` fonksiyonları çeşitli türlerde argüman alır ve `int` türünde sonuç döndürür. Uygulama, sonucun her zaman bir `int` içine sığacağını garanti eder. 

### 📌 `len`

| Çağrı    | Argüman türü    | Sonuç                                           |
| -------- | --------------- | ----------------------------------------------- |
| `len(s)` | `string`        | string uzunluğu (bayt cinsinden)                |
|          | `[n]T`, `*[n]T` | array uzunluğu (`== n`)                         |
|          | `[]T`           | slice uzunluğu                                  |
|          | `map[K]T`       | map uzunluğu (tanımlı anahtar sayısı)           |
|          | `chan T`        | channel tamponundaki kuyruklanmış eleman sayısı |
|          | tür parametresi | aşağıya bakın                                   |

### 📌 `cap`

| Çağrı    | Argüman türü    | Sonuç                     |
| -------- | --------------- | ------------------------- |
| `cap(s)` | `[n]T`, `*[n]T` | array uzunluğu (`== n`)   |
|          | `[]T`           | slice kapasitesi          |
|          | `chan T`        | channel tampon kapasitesi |
|          | tür parametresi | aşağıya bakın             |

Argüman türü bir tür parametresi `P` ise, `len(e)` (veya `cap(e)`) çağrısı `P`’nin type set’indeki her tür için geçerli olmalıdır. Sonuç, `P`’nin instantiate edildiği tür argümanına karşılık gelen argümanın uzunluğu (veya kapasitesi) olur. 

Bir slice’ın kapasitesi, alttaki array’de yer ayrılmış eleman sayısıdır. Her zaman şu ilişki geçerlidir: 

```text
0 <= len(s) <= cap(s)
```

`nil` slice, map veya channel’ın uzunluğu 0’dır. `nil` slice veya channel’ın kapasitesi 0’dır. 

`len(s)` ifadesi, `s` bir string sabitiyse sabittir. `len(s)` ve `cap(s)`, `s` bir array (veya array pointer’ı) türündeyse ve `s` ifadesi channel alma işlemleri veya (sabit olmayan) fonksiyon çağrıları içermiyorsa sabittir; bu durumda `s` değerlendirilmez. Aksi hâlde `len` ve `cap` çağrıları sabit değildir ve `s` değerlendirilir. 

```go
const (
	c1 = imag(2i)                    // imag(2i) = 2.0 bir sabittir
	c2 = len([10]float64{2})         // [10]float64{2} fonksiyon çağrısı içermez
	c3 = len([10]float64{c1})        // [10]float64{c1} fonksiyon çağrısı içermez
	c4 = len([10]float64{imag(2i)})  // imag(2i) sabittir ve fonksiyon çağrısı yapılmaz
	c5 = len([10]float64{imag(z)})   // geçersiz: imag(z) (sabit olmayan) bir fonksiyon çağrısıdır
)
var z complex128
```

---

## 🧱 Slice, map ve channel oluşturma¶

Yerleşik `make`, bir tür `T` alır; bu tür slice, map veya channel türü ya da bir tür parametresi olmalıdır; isteğe bağlı olarak tür-özel bir ifade listesi izleyebilir. `T` türünde bir değer döndürür ( `*T` değil ). Bellek, başlangıç değerleri bölümünde anlatıldığı gibi başlatılır. 

| Çağrı           | Tür `T`         | Sonuç                                             |
| --------------- | --------------- | ------------------------------------------------- |
| `make(T, n)`    | slice           | `len == n`, `cap == n` olan `T` türünde slice     |
| `make(T, n, m)` | slice           | `len == n`, `cap == m` olan `T` türünde slice     |
| `make(T)`       | map             | `T` türünde map                                   |
| `make(T, n)`    | map             | yaklaşık `n` eleman için başlangıç alanı olan map |
| `make(T)`       | channel         | `T` türünde tamponsuz channel                     |
| `make(T, n)`    | channel         | tampon boyutu `n` olan kanal                      |
| `make(T, n)`    | tür parametresi | aşağıya bakın                                     |
| `make(T, n, m)` | tür parametresi | aşağıya bakın                                     |

İlk argüman bir tür parametresiyse, type set’teki tüm türler aynı alttaki türe sahip olmalıdır; bu tür bir slice veya map olmalı; ya da channel türleri varsa, yalnızca channel türleri olmalı, hepsi aynı eleman türüne sahip olmalı ve channel yönleri çakışmamalıdır. 

Boyut argümanlarının (`n` ve `m`) her biri tamsayı türünde olmalı, yalnızca tamsayı türleri içeren bir type set’e sahip olmalı veya tipsiz bir sabit olmalıdır. Sabit bir boyut argümanı negatif olmamalı ve `int` türünde bir değerle temsil edilebilir olmalıdır; tipsiz sabitse `int` türü verilir. Hem `n` hem `m` veriliyorsa ve sabitse, `n`, `m`’den büyük olmamalıdır. Slice ve channel’lar için, çalışma zamanında `n` negatifse veya `m`’den büyükse *panic* oluşur. 

```go
s := make([]int, 10, 100)       // len(s) == 10, cap(s) == 100
s := make([]int, 1e3)           // len(s) == cap(s) == 1000
s := make([]int, 1<<63)         // yasadışı: len(s) bir int ile temsil edilemez
s := make([]int, 10, 0)         // yasadışı: len(s) > cap(s)
c := make(chan int, 10)         // tampon boyutu 10 olan channel
m := make(map[string]int, 100)  // yaklaşık 100 eleman için başlangıç alanı olan map
```

Bir map türüyle `make` çağrısı ve boyut ipucu `n`, `n` map elemanını tutacak başlangıç alanıyla bir map oluşturur. Tam davranış uygulamaya bağlıdır. 

---

## 🔽 Min ve max¶

Yerleşik `min` ve `max`, sıralı (*ordered*) türlerde sabit sayıda argümanın en küçük veya en büyük değerini hesaplar. En az bir argüman olmalıdır [Go 1.21]. 

İşleçler için olan aynı tür kuralları uygulanır: sıralı argümanlar `x` ve `y` için `min(x, y)`, `x + y` geçerliyse geçerlidir ve `min(x, y)`’nin türü, `x + y`’nin türüdür ( `max` için de benzer). Tüm argümanlar sabitse sonuç sabittir. 

```go
var x, y int
m := min(x)                 // m == x
m := min(x, y)              // m x ve y'nin küçüğüdür
m := max(x, y, 10)          // m x ve y'nin büyüğüdür ama en az 10'dur
c := max(1, 2.0, 10)        // c == 10.0 (kayan nokta kind)
f := max(0, float32(x))     // f'nin türü float32'dir
var s []string
_ = min(s...)               // geçersiz: slice argümanlarına izin verilmez
t := max("", "foo", "bar")  // t == "foo" (string kind)
```

Sayısal argümanlar için, tüm NaN’lerin eşit olduğu varsayımıyla `min` ve `max` değişme ve birleşme özelliklerine sahiptir: 

```text
min(x, y)    == min(y, x)
min(x, y, z) == min(min(x, y), z) == min(x, min(y, z))
```

Kayan nokta argümanlarda negatif sıfır, NaN ve sonsuzluk için şu kurallar geçerlidir: 

```text
   x        y    min(x, y)    max(x, y)

  -0.0    0.0         -0.0          0.0    // negatif sıfır, (negatif olmayan) sıfırdan küçüktür
  -Inf      y         -Inf            y    // negatif sonsuzluk, diğer tüm sayılardan küçüktür
  +Inf      y            y         +Inf    // pozitif sonsuzluk, diğer tüm sayılardan büyüktür
   NaN      y          NaN          NaN    // argümanlardan biri NaN ise, sonuç NaN olur
```

String argümanlar için `min` sonucu; bayt düzeyinde sözlüksel karşılaştırmaya göre en küçük değere sahip ilk argümandır ( `max` için en büyük): 

```text
min(x, y)    == eğer x <= y ise x, değilse y
min(x, y, z) == min(min(x, y), z)
```

---

## 🧿 Tahsis (allocation)¶

Yerleşik `new`, bir tür `T` alır, çalışma zamanında o türde bir değişken için depolama alanı tahsis eder ve onu işaret eden `*T` türünde bir değer döndürür. Değişken, başlangıç değerleri bölümünde anlatıldığı gibi başlatılır. 

```go
new(T)
```

Örneğin:

```go
type S struct { a int; b float64 }
new(S)
```

`S` türünde bir değişken için depolama alanı tahsis eder, başlatır (`a=0`, `b=0.0`) ve o konumun adresini içeren `*S` türünde bir değer döndürür. 

---

## 🧨 Panic’leri ele alma¶

İki yerleşik fonksiyon, `panic` ve `recover`, çalışma zamanı panic’lerini ve programın tanımladığı hata koşullarını raporlamaya ve ele almaya yardımcı olur. 

```go
func panic(interface{})
func recover() interface{}
```

Bir fonksiyon `F` yürütülürken, `panic`’e açık bir çağrı veya çalışma zamanı panic’i `F`’nin yürütümünü sonlandırır. Ardından `F` tarafından ertelenen (*deferred*) tüm fonksiyonlar olağan şekilde yürütülür. Sonra `F`’yi çağıranın ertelenmiş fonksiyonları yürütülür ve böylece yürütülen goroutine’deki en üst düzey fonksiyonun ertelenmiş fonksiyonlarına kadar sürer. Bu noktada program sonlandırılır ve hata koşulu, `panic` argümanının değeri dâhil raporlanır. Bu sonlandırma dizisine *panicking* denir. 

```go
panic(42)
panic("unreachable")
panic(Error("cannot parse"))
```

`recover`, panic’e giren bir goroutine’in davranışını yönetmesine izin verir. `G` fonksiyonunun, `recover` çağıran bir `D` fonksiyonunu `defer` ettiğini ve `G`’nin yürütüldüğü aynı goroutine’de bir panic oluştuğunu varsayın. Ertelenmiş fonksiyonların yürütümü `D`’ye ulaştığında, `recover` çağrısının dönüş değeri, `panic` çağrısına geçirilen değer olur. `D` normal dönerse ve yeni bir panic başlatmazsa, panicking dizisi durur. Bu durumda, `G` ile `panic` çağrısı arasında çağrılan fonksiyonların durumu atılır ve normal yürütüm devam eder. `G`’nin `D`’den önce ertelediği diğer fonksiyonlar daha sonra yürütülür ve `G` çağırana dönerek sonlanır. 

Goroutine panic’te değilse veya `recover` doğrudan ertelenmiş bir fonksiyon tarafından çağrılmadıysa, `recover`’ın dönüş değeri `nil` olur. Tersine, goroutine panic’teyse ve `recover` doğrudan ertelenmiş bir fonksiyon tarafından çağrıldıysa, `recover`’ın dönüş değerinin `nil` olmayacağı garanti edilir. Bunu sağlamak için, `nil` bir interface değeriyle (veya tipsiz `nil` ile) `panic` çağırmak çalışma zamanında panic oluşturur. 

Aşağıdaki örnekteki `protect` fonksiyonu, fonksiyon argümanı `g`’yi çağırır ve çağıranları `g`’nin neden olduğu çalışma zamanı panic’lerinden korur: 

```go
func protect(g func()) {
	defer func() {
		log.Println("done")  // Println, panic olsa bile normal çalışır
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()
	log.Println("start")
	g()
}
```

---

## 🪜 Bootstrapping¶

Mevcut uygulamalar, bootstrapping sırasında yararlı olan bazı yerleşik fonksiyonlar sağlar. Bu fonksiyonlar bütünlük için belgelenmiştir; ancak dilde kalacağı garanti edilmez. Sonuç döndürmezler. 

| Fonksiyon | Davranış                                                                         |
| --------- | -------------------------------------------------------------------------------- |
| `print`   | tüm argümanları yazar; biçimlendirme uygulamaya özeldir                          |
| `println` | `print` gibidir; ancak argümanlar arasına boşluk koyar ve sonda yeni satır basar |

Uygulama kısıtlaması: `print` ve `println`, keyfî argüman türlerini kabul etmek zorunda değildir; ancak boolean, sayısal ve string türlerini yazdırma desteklenmelidir. 



