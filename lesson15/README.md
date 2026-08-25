# Go'da Interface — to'liq dars qo'llanmasi

---

## Dars rejasi (o'qituvchi uchun)

| # | Bo'lim | Vaqt |
|---|--------|------|
| 1 | Muammo: interfacesiz kod | 10 daq |
| 2 | Interface nima? Ta'rif va sintaksis | 10 daq |
| 3 | Implicit (yashirin) implementatsiya | 10 daq |
| 4 | Interface ichkarida qanday ishlaydi | 15 daq |
| 5 | `nil` tuzoq — eng mashhur xato | 10 daq |
| 6 | Type assertion va type switch | 10 daq |
| 7 | Standart kutubxonadagi misollar | 10 daq |
| 8 | Foydasi va zarari (tahlil) | 10 daq |
| 9 | Best practice va antipattern | 5 daq |
| 10 | Mashqlar | 30 daq |

---

## 1. Muammo: keling, avval interfacesiz yozib ko'ramiz

Vazifa: foydalanuvchiga xabar yuborish. Kanal: email yoki SMS.

```go
package main

import (
    "errors"
    "fmt"
)

type EmailSender struct {
    APIKey string
}

func (e EmailSender) Send(to, msg string) error {
    fmt.Printf("EMAIL -> %s: %s\n", to, msg)
    return nil
}

type SMSSender struct {
    Login string
}

func (s SMSSender) Send(to, msg string) error {
    fmt.Printf("SMS -> %s: %s\n", to, msg)
    return nil
}

// Interfacesiz "dispetcher"
func Notify(kind, to, msg string) error {
    switch kind {
    case "email":
        return EmailSender{APIKey: "..."}.Send(to, msg)
    case "sms":
        return SMSSender{Login: "..."}.Send(to, msg)
    default:
        return errors.New("noma'lum kanal")
    }
}
```

Ko'rinishidan ishlaydi. Endi savol bering:

**Savol 1.** Telegram qo'shsak nima bo'ladi?
→ `Notify` funksiyasini **ochib o'zgartirish** kerak. Har yangi kanal — ishlab turgan kodga tegish. Bu **Open/Closed** prinsipining buzilishi.

**Savol 2.** `Notify` ni test qilib bo'ladimi?
→ Yo'q. Testda ham haqiqiy email ketadi. Konkret tiplar funksiya **ichida** qattiq bog'langan.

**Savol 3.** `Notify` nechta paketni bilishi kerak?
→ Hammasini: email, sms, telegram, push... Yuqori sathdagi kod quyi sathdagi barcha detallarga bog'liq.

Mana shu uchta og'riq — interfacening butun mavjudlik sababi.

---

## 2. Interface nima?

**Ta'rif:** interface — bu *tipni* emas, *xatti-harakatni* tavsiflovchi tip. U faqat metod imzolarini sanaydi, tanasini emas.

```go
type Sender interface {
    Send(to, msg string) error
}
```

Bu shunday deyish: "menga *nima* ekaning qiziq emas, `Send(string, string) error` metodi bo'lsa bas".

Endi `Notify` ni qayta yozamiz:

```go
func Notify(s Sender, to, msg string) error {
    return s.Send(to, msg)
}
```

Butun `switch` yo'qoldi. Telegram qo'shsak:

```go
type TelegramSender struct{ Token string }

func (t TelegramSender) Send(to, msg string) error {
    fmt.Printf("TG -> %s: %s\n", to, msg)
    return nil
}
```

`Notify` ga **bir harf ham tegmadik**. Ishlatish:

```go
func main() {
    Notify(EmailSender{}, "ali@mail.uz", "Salom")
    Notify(SMSSender{}, "+998901234567", "Salom")
    Notify(TelegramSender{}, "@ali", "Salom")
}
```

Test esa endi bir zumda yoziladi:

```go
type MockSender struct {
    Calls []string
    Err   error
}

func (m *MockSender) Send(to, msg string) error {
    m.Calls = append(m.Calls, to+":"+msg)
    return m.Err
}

func TestNotify(t *testing.T) {
    mock := &MockSender{}
    if err := Notify(mock, "ali", "test"); err != nil {
        t.Fatal(err)
    }
    if len(mock.Calls) != 1 {
        t.Fatalf("kutildi 1 ta chaqiruv, bo'ldi %d", len(mock.Calls))
    }
}
```

**Asosiy fikr:** interface — bu *qaramlikni ag'darish* (dependency inversion) vositasi. Yuqori sath quyi sathga emas, ikkalasi ham abstraksiyaga bog'lanadi.

---

## 3. Implicit implementatsiya — Go'ning o'ziga xosligi

Java yoki C#'da shunday yozasiz:

```java
class EmailSender implements Sender { ... }   // OSHKORA
```

Go'da bunday **yo'q**. Tip kerakli metodlarga ega bo'lsa — u interfaceni avtomatik qanoatlantiradi. Buni "structural typing" yoki norasmiy "duck typing" deyishadi.

Bu nimani beradi?

1. **Interfaceni siz e'lon qilasiz, muallif emas.** Boshqa birovning kutubxonasidagi tip uchun o'zingiz interface yozib, uni mock qila olasiz — kutubxonani o'zgartirmasdan.
2. **Paketlar bir-birini import qilmaydi.** `Sender` interfacei `EmailSender` haqida hech narsa bilmaydi.

Kamchiligi ham bor: **kompilyator sizni "men bu interfaceni amalga oshirmoqchi edim" degan niyatingizdan bexabar.** Metod imzosida xato qilsangiz, xato faqat ishlatgan joyingizda chiqadi.

Shuning uchun **compile-time tekshiruv idiomasi**ni o'rgating:

```go
// Bu qator hech narsa qilmaydi, lekin EmailSender Sender'ni
// qanoatlantirmasa, kod KOMPILYATSIYA BO'LMAYDI.
var _ Sender = (*EmailSender)(nil)
var _ Sender = EmailSender{}
```

Bu — professional Go kodida juda ko'p uchraydigan qator.

### Pointer receiver vs value receiver — muhim nuqta

```go
type Counter struct{ n int }

func (c *Counter) Inc() { c.n++ }   // POINTER receiver

type Incrementer interface{ Inc() }

var i Incrementer = &Counter{}  // ✅ ishlaydi
var j Incrementer = Counter{}   // ❌ KOMPILYATSIYA XATOSI
```

**Qoida:** metod pointer receiver bilan yozilgan bo'lsa, interfaceni faqat `*T` qanoatlantiradi, `T` emas. Value receiver bo'lsa — ikkalasi ham.

Sababi: `Counter{}` qiymat sifatida interfacega solinsa, uning nusxasi olinadi va o'zgartirish hech qayerga yetib bormaydi. Go bu xatoni oldindan to'sadi.

---

## 4. Interface ichkarida qanday ishlaydi

Bu qism talabani "sehr" tuyg'usidan xalos qiladi. Runtime'da interface — bu **ikki so'zli struktura**.

Metodli interface (`iface`):

```
┌──────────────┬──────────────┐
│   tab        │    data      │
│  (*itab)     │ (unsafe.Ptr) │
└──────────────┴──────────────┘
      │                │
      │                └──> haqiqiy qiymat (yoki unga pointer)
      └──> itab: { interface tipi, konkret tip, metodlar jadvali }
```

Bo'sh interface (`eface` — ya'ni `any`):

```
┌──────────────┬──────────────┐
│   _type      │    data      │
└──────────────┴──────────────┘
```

Buni amalda ko'rsating:

```go
package main

import (
    "fmt"
    "unsafe"
)

type Sender interface{ Send(to, msg string) error }

func main() {
    var s Sender
    fmt.Println(unsafe.Sizeof(s)) // 16 (64-bit tizimda: 2 * 8 bayt)
}
```

Bundan kelib chiqadigan xulosalar:

- Interface orqali metod chaqirish — bu **itab jadvalidan pointer o'qib, o'sha manzilga o'tish** (dynamic dispatch). Konkret tipdagi to'g'ridan-to'g'ri chaqiruvdan sekinroq.
- Kompilyator interface orqali chaqirilgan metodni odatda **inline qila olmaydi** — chunki qaysi funksiya chaqirilishi kompilyatsiya vaqtida noma'lum.
- Qiymatni interfacega solish ko'pincha uni **heap'ga chiqarib yuboradi** (escape to heap) → GC uchun qo'shimcha yuk.

---

## 5. `nil` tuzoq — Go'dagi eng mashhur xato

Buni albatta o'ting. Deyarli har bir Go dasturchisi bunga bir marta tushgan.

```go
package main

import "fmt"

type MyError struct{ Code int }

func (e *MyError) Error() string { return fmt.Sprintf("xato %d", e.Code) }

func doWork() error {
    var err *MyError = nil   // pointer nil
    return err               // ...lekin interface nil EMAS!
}

func main() {
    if err := doWork(); err != nil {
        fmt.Println("XATO BOR!") // ← MANA SHU CHOP ETILADI
    }
}
```

**Nega?** Interface `nil` bo'lishi uchun **ikkala so'z ham** (`tab` va `data`) nil bo'lishi kerak. Bu yerda:

```
tab  = *MyError tipi haqidagi ma'lumot   ← nil EMAS
data = nil
```

Tip ma'lumoti mavjud → interfacening o'zi nil emas.

**To'g'ri yozilishi:**

```go
func doWork() error {
    var err *MyError = nil
    if err != nil {
        return err
    }
    return nil   // ochiq-oydin nil qaytaramiz
}
```

**Amaliy qoida:** konkret error tipini hech qachon `error` sifatida qaytaradigan o'zgaruvchiga saqlamang. Funksiya imzosi `error` bo'lsa, ichida ham `error` bilan ishlang.

---

## 6. Type assertion va type switch

Interfacedan konkret tipga qaytish kerak bo'lganda:

```go
var s Sender = EmailSender{APIKey: "k"}

// Xavfli forma — mos kelmasa panic
e := s.(EmailSender)

// Xavfsiz forma — HAR DOIM shuni ishlating
if e, ok := s.(EmailSender); ok {
    fmt.Println(e.APIKey)
}
```

Type switch:

```go
func describe(s Sender) string {
    switch v := s.(type) {
    case EmailSender:
        return "email, key=" + v.APIKey
    case SMSSender:
        return "sms, login=" + v.Login
    case nil:
        return "hech narsa"
    default:
        return fmt.Sprintf("noma'lum: %T", v)
    }
}
```

⚠️ **Ogohlantirish talabaga:** agar kodingizda interface ustidan katta `type switch` paydo bo'lsa — bu ko'pincha *interfaceni noto'g'ri loyihalaganingiz* belgisi. Bu holda o'sha xatti-harakatni interfacening o'ziga metod qilib qo'shish kerak.

---

## 7. Standart kutubxonadagi misollar

Interface tushunchasi mavhum tuyulmasligi uchun stdlib'dan ko'rsating.

### `error` — eng ko'p ishlatiladigan interface

```go
type error interface {
    Error() string
}
```

### `fmt.Stringer`

```go
type Stringer interface {
    String() string
}

type Money struct{ Sum int64 }

func (m Money) String() string {
    return fmt.Sprintf("%d.%02d so'm", m.Sum/100, m.Sum%100)
}

fmt.Println(Money{123456})  // 1234.56 so'm
```

### `io.Writer` — kichik interfacening kuchi

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

Bitta metod. Lekin uni quyidagilar qanoatlantiradi: `os.File`, `os.Stdout`, `bytes.Buffer`, `strings.Builder`, `net.Conn`, `http.ResponseWriter`, `gzip.Writer`...

```go
func Report(w io.Writer, data []string) {
    for _, d := range data {
        fmt.Fprintln(w, d)
    }
}

// Prodda:
Report(os.Stdout, data)
Report(httpRespWriter, data)
Report(file, data)

// Testda:
var buf bytes.Buffer
Report(&buf, data)
if buf.String() != "kutilgan" { t.Fail() }
```

Mana shu misol talabaga interfacening qadrini eng yaxshi tushuntiradi.

### `sort.Interface` — uchta metod

```go
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
```

`sort` paketi sizning ma'lumot strukturangiz haqida hech narsa bilmaydi, lekin uni saralay oladi.

---

## 8. Foydasi va zarari — halol tahlil

### ✅ Foydalari

| Foyda | Izoh |
|---|---|
| **Testlanuvchanlik** | Mock/stub yozish oson. Bu — Go'da interfaceni ishlatishning **1-raqamli sababi**. |
| **Qaramlikni ag'darish** | Biznes-mantiq PostgreSQL, Redis, Kafka haqida bilmaydi. Ularni almashtirish arzon. |
| **Kengaytiriluvchanlik** | Yangi implementatsiya = yangi fayl. Eski kodga tegilmaydi. |
| **Polimorfizm** | Bir xil kod turli tiplar bilan ishlaydi, generics'siz ham. |
| **Paketlar ajratilishi** | Import grafigi soddalashadi, sikllar yo'qoladi. |
| **Plugin arxitekturasi** | `http.Handler`, `driver.Driver` kabi kengaytiriladigan tizimlar. |

### ❌ Zararlari va narxi

| Zarar | Izoh |
|---|---|
| **Ishlash tezligi** | Dynamic dispatch + inline qilinmasligi. Issiq sikllarda (millionlab iteratsiya) seziladi. |
| **Heap allokatsiyalar** | Qiymat interfacega solinganda ko'pincha heap'ga qochadi → GC bosimi. |
| **Kodni o'qish qiyinlashadi** | "Bu yerda aslida qaysi implementatsiya ishlayapti?" — IDE'da `Go to definition` bosgan talaba interface ta'rifiga tushadi, kodga emas. |
| **`nil` tuzog'i** | 5-bo'limdagi xato. |
| **Ortiqcha abstraksiya** | Bitta implementatsiyasi bor interface — bu shunchaki keraksiz qatlam. |
| **Semiz interfacelar** | 15 metodli interfaceni mock qilish — azob. |
| **Kompilyator niyatni bilmaydi** | Implicit implementatsiyaning teskari tomoni. |

### Tezlik farqini o'lchash (talabaga topshiriq)

Gapirmang — **o'lchating**. Quyidagi benchmarkni birga ishga tushiring:

```go
package main

import "testing"

type Adder interface{ Add(a, b int) int }

type Impl struct{}

func (Impl) Add(a, b int) int { return a + b }

var sink int

func BenchmarkDirect(b *testing.B) {
    var impl Impl
    for i := 0; i < b.N; i++ {
        sink = impl.Add(i, i)
    }
}

func BenchmarkInterface(b *testing.B) {
    var a Adder = Impl{}
    for i := 0; i < b.N; i++ {
        sink = a.Add(i, i)
    }
}
```

```bash
go test -bench=. -benchmem
```

Talaba o'z ko'zi bilan ko'rsin: to'g'ridan-to'g'ri chaqiruv inline bo'lib deyarli nolga tushadi, interface orqali chaqiruv esa har safar bir necha nanosekund oladi.

Keyin muhim savolni bering: **"Bu farq sizning HTTP handleringizda muhimmi, unda baribir 20 ms baza so'rovi bor?"**

Javob: **yo'q**. Interfacening narxi — bu real, lekin u faqat *issiq siklda* muhim. 99% holatda o'qiluvchanlik va testlanuvchanlik muhimroq.

### Escape'ni ko'rsatish

```bash
go build -gcflags="-m" ./...
```

Chiqishda `escapes to heap` yozuvlarini talabaga ko'rsating — qiymat interfacega solingan joylarda paydo bo'ladi.

---

## 9. Best practice va antipatternlar

### ✅ Qiling

**1. Interfacelar kichik bo'lsin.**
> "The bigger the interface, the weaker the abstraction." — Rob Pike

1–3 metod ideal. `io.Reader` bitta metod bilan butun ekotizimni ushlab turadi.

**2. "Interfaceni qabul qiling, struct qaytaring."**

```go
// ✅ yaxshi
func NewService(db Storage) *Service { ... }

// ❌ yomon
func NewService(db *postgres.DB) Servicer { ... }
```

Sabab: chaqiruvchi qaytgan strukturaning barcha imkoniyatlarini ko'rishi kerak; kirishda esa moslashuvchanlik kerak.

**3. Interfaceni *ishlatuvchi* tomonda e'lon qiling.**

```go
// paket: order  (ISTE'MOLCHI)
package order

type PaymentGateway interface {   // ← interface shu yerda
    Charge(amount int64) error
}

type Service struct{ pg PaymentGateway }
```

```go
// paket: payme  (TA'MINLOVCHI)
package payme

type Client struct{}                            // ← oddiy struct
func (c *Client) Charge(amount int64) error { } // interface haqida bilmaydi
```

Bu Go'da Java'dan asosiy farq. Java'da interface provider paketida turadi, Go'da — consumer paketida.

**4. Compile-time tekshiruv qo'ying:** `var _ Sender = (*EmailSender)(nil)`

**5. Interfaceni ehtiyoj paydo bo'lganda yarating,** oldindan emas. Avval konkret tip bilan yozing; ikkinchi implementatsiya yoki test kerak bo'lganda — interfacega chiqaring.

### ❌ Qilmang

| Antipattern | Nega yomon |
|---|---|
| Har struct uchun interface | Faylni ikki barobar oshiradi, foyda nol |
| `IUserService`, `UserServiceInterface` nomlash | Go idiomasi emas. `Reader`, `Writer`, `Sender` — `-er` qo'shimchasi |
| Bitta implementatsiyali interface | YAGNI. Test uchun kerak bo'lsa — mayli, aks holda o'chiring |
| Hamma joyda `any` (`interface{}`) | Tip xavfsizligini yo'qotasiz. Generics ishlating |
| Interface ustidan ulkan `type switch` | Metod qo'shish kerak edi |
| Semiz interface (10+ metod) | Bo'laklarga bo'ling |

### `any` vs generics

Go 1.18+ dan keyin "har qanday tip" uchun `any` emas, generics ishlating:

```go
// ❌ eski usul — tip xavfsizligi yo'q, allokatsiya bor
func Max(a, b any) any { ... }

// ✅ generics — kompilyatsiya vaqtida tekshiriladi, tezroq
func Max[T int | float64 | string](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

**Farqni tushuntiring:**
- **Interface** = *xatti-harakat* polimorfizmi (turli tiplar, turli mantiq)
- **Generics** = *tip* polimorfizmi (turli tiplar, bir xil mantiq)

---

## 10. Mashqlar

### Mashq 1 (oson)
`Shape` interfacei yozing: `Area() float64` va `Perimeter() float64`. `Circle`, `Rectangle`, `Triangle` uchun amalga oshiring. `TotalArea(shapes []Shape) float64` funksiyasini yozing.

### Mashq 2 (o'rta)
Quyidagi kodni interface yordamida qayta yozing va `switch` ni butunlay yo'qoting:

```go
func Save(format string, data map[string]string) ([]byte, error) {
    switch format {
    case "json":  // ...
    case "csv":   // ...
    case "xml":   // ...
    }
    return nil, errors.New("noma'lum format")
}
```

### Mashq 3 (nil tuzog'i)
Quyidagi kod nega `"xato bor"` chop etadi? Tuzating:

```go
type ValidationError struct{ Field string }

func (v *ValidationError) Error() string { return v.Field + " noto'g'ri" }

func Validate(age int) error {
    var e *ValidationError
    if age < 0 {
        e = &ValidationError{Field: "age"}
    }
    return e
}
```

### Mashq 4 (amaliy — eng muhimi)
`UserService` yozing: u foydalanuvchini bazadan oladi va email yuboradi.
- `UserRepository` va `Mailer` interfacelarini **consumer tomonda** e'lon qiling
- Ikkalasi uchun mock yozing
- `go test` bilan **bazasiz va emailsiz** to'liq test qiling

### Mashq 5 (chuqur)
Struct'ni `io.Writer` qiling — u yozilgan baytlarni sanasin:

```go
type CountingWriter struct {
    w     io.Writer
    Count int64
}
```

Keyin uni `os.Stdout` ustiga o'rang va nechta bayt chiqqanini o'lchang. (Bu — **decorator pattern**, interfacening eng kuchli qo'llanishlaridan biri.)

---

## Dars yakuni: bitta jumla bilan

> **Interface — bu kodingiz o'rtasidagi shartnoma. Uni tezlik uchun emas, o'zgarishga chidamlilik va testlanuvchanlik uchun ishlatasiz. Va faqat haqiqatan kerak bo'lganda.**

### Tekshiruv savollari

1. Interface nol qiymati nima? (`nil`)
2. Nega `nil` pointer qaytarilganda `err != nil` rost bo'ladi?
3. `Counter{}` `Incrementer` ni qanoatlantiradimi, agar `Inc()` pointer receiver bo'lsa? (Yo'q)
4. Interface runtime'da necha bayt egallaydi? (16, 64-bit tizimda)
5. Interfaceni qaysi paketda e'lon qilish kerak — ishlatuvchida yoki amalga oshiruvchida? (Ishlatuvchida)
6. Qachon interface ishlatmaslik kerak?