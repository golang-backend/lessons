# Shahbozbek — Tekshiruv natijasi (feedback)

`checking_level_task.md` dagi 10 masala bo'yicha tekshirildi. Umumiy holat: **yaxshi**,
asosiy tushunchalar (slice, map, struct, method, interface) o'zlashtirilgan. Quyida
har bir fayl bo'yicha kamchiliklar va yaxshilash takliflari.

---

## Umumiy kuzatuvlar

1. **Fayl nomlari mazmuniga mos emas.** Masalan, `task3.go` ichida 5/6-masala,
   `task4.go` ichida 7/8-masala, `task5.go` ichida 10-masala bor. `go run task3.go`
   desangiz aslida 5-masala ishlaydi — chalkash. Fayllarni mazmuniga qarab nomlang
   (`task5_6.go`, `task7_8.go`, `task10.go`) yoki har masalani alohida papkaga oling.

2. **9-masala umuman yo'q.** `interface{}` + `switch v := x.(type)` (type switch)
   masalasi yozilmagan. Bu D darajadagi muhim masala — uni qo'shing.

3. Har fayl `package main` va o'z `main()` iga ega. Bu to'g'ri, lekin bitta papkada
   bir nechta `main()` bo'lgani uchun `go run .` ishlamaydi — faqat `go run <fayl>.go`
   ishlaydi. Buni bilib qo'ying.

---

## task1.go — 1 & 2-masala

**1-masala (average): muhim xato.**
```go
func average(slice []int) int {   // ← int qaytaryapti
	...
	return sum / len(slice)        // ← butun sonli bo'lish, kasr yo'qoladi
}
```
`12+7+25+3+18 = 65`, `65/5 = 13` chiqdi — bu yerda tasodifan butun. Lekin
o'rtacha odatda kasrli bo'ladi (mas. `[1,2] → 1.5`). Bu kod `1` qaytaradi.

**Tuzatish:**
```go
func average(slice []int) float64 {
	sum := 0
	for _, num := range slice {
		sum += num
	}
	return float64(sum) / float64(len(slice))
}
```

**2-masala:** to'g'ri ishlaydi 👍. Kichik eslatma: funksiya nomi `A` —
mazmunli nom bering (`juftlarniAjrat` yoki `filterEven`).

---

## task2.go — 3 & 4-masala

**3-masala (sanash): to'g'ri** 👍. Faqat funksiya nomida xato — `caunts` →
`counts` (imlo).

**4-masala (telefon kitobchasi): ikkita muammo bor.**

Xato natija: `topildi  true` chiqyapti, raqam ko'rinmadi. Sababi:
```go
value, ok := contacts[ism]
if ok {
	fmt.Println("topildi", contacts[value], ok)  // ← contacts[value] xato
}
```
`value` allaqachon raqam (`+998991112233`). Siz `contacts[value]` deb yana shu
raqamni *kalit* sifatida qidiryapsiz — natija bo'sh string. To'g'risi shunchaki
`value` ni chop etish:
```go
if ok {
	fmt.Println("topildi:", value)
} else {
	fmt.Println("Topilmadi")
}
```

**`delete` yo'q.** Masala sharti aniq aytadi: "Bitta kontaktni `delete` bilan
o'chiring." Bu qism umuman yozilmagan. Qo'shing:
```go
delete(contacts, "Vali")
```
va o'chirilgandan keyin yana qidirib, "Topilmadi" chiqishini ko'rsating.

---

## task3.go — 5 & 6-masala (fayl nomi chalkash)

**Spec bilan tur mos emas.** Masala `Baho float64` deydi, sizda `Baho int`:
```go
type Student struct {
	Ism  string
	Yosh int
	Baho int   // ← float64 bo'lishi kerak
}
```
Hozir `Baho: 5.0` yozsangiz ham `int` ga `5` bo'lib tushyapti. `4.5` kabi kasr
baho qo'ya olmaysiz. `float64` qiling — shunda `Otdimi()` dagi `s.Baho > 3` ham
`3.0` bilan to'g'ri solishtiriladi.

Qolgani to'g'ri ishlaydi 👍 (5-masalada `>= 4.0` filtr, 6-masalada method).

> Eslatma (value vs pointer): `Otdimi()` faqat o'qiyapti, qiymatni
> o'zgartirmaydi — shu uchun value receiver `(s Student)` to'g'ri tanlov. Buni
> tushuntira olsangiz, masala maqsadiga yetdingiz.

---

## task4.go — 7 & 8-masala (fayl nomi chalkash)

**7-masala (nested struct): to'g'ri** 👍.

**8-masala (Shape interface): to'g'ri ishlaydi** 👍 — interfeysni tushunganingiz
ko'rinib turibdi (har xil struct bitta `[]Shape` da).

Ikki kichik tavsiya:
- `eni`, `boyi` ni `int` qilib, keyin `float64(t.eni * t.boyi)` ga aylantiryapsiz.
  Yuza odatda kasrli — to'g'ridan-to'g'ri `float64` maydon qiling.
- `3.14` o'rniga `math.Pi` ishlatsangiz aniqroq bo'ladi.

---

## task5.go — 10-masala (fayl nomi chalkash)

Ko'p qismi to'g'ri, lekin **3 ta masala sharti bajarilmagan**:

1. **`map[string]float64` yo'q.** Spec: "mahsulot nomi → umumiy qiymati
   (`Narxi * Soni`)" mapда saqlanishi kerak. Siz mapni umuman yaratmay, `for`
   ichida bevosita chop etyapsiz. Map yarating:
   ```go
   qiymatlar := make(map[string]float64)
   for _, m := range mahsulotlar {
   	qiymatlar[m.Nomi] = m.Narxi * float64(m.Soni)
   }
   ```

2. **`mostExpensive` natijasi ishlatilmagan.** `mostExpensive(mahsulotlar)`
   chaqirilyapti, lekin qaytgan qiymat hech qayerda chop etilmaydi. Natijani
   oling va ko'rsating:
   ```go
   eng := mostExpensive(mahsulotlar)
   fmt.Println("Eng qimmat:", eng.Nomi)
   ```
   Yana: funksiya *ichida* `fmt.Println` qilish noto'g'ri uslub — funksiya faqat
   topib qaytarsin, chop etishni `main` qilsin.

3. **`Narxi * Soni` da tur yo'qotilyapti.** `int(mahsulot.Narxi)` — narx
   `float64` bo'lsa kasr qismi tashlanadi. To'g'risi:
   `mahsulot.Narxi * float64(mahsulot.Soni)`.

`Qimmatbaho()` methodi to'g'ri 👍.

---

## Xulosa (daraja)

| Masala | Holat |
|--------|-------|
| 1 | ⚠️ ishlaydi, lekin `int` o'rniga `float64` bo'lishi kerak |
| 2 | ✅ |
| 3 | ✅ |
| 4 | ❌ qidiruv natijasi xato + `delete` yo'q |
| 5 | ⚠️ `Baho` turi `float64` bo'lsin |
| 6 | ✅ |
| 7 | ✅ |
| 8 | ✅ |
| 9 | ❌ umuman yozilmagan |
| 10 | ⚠️ map yo'q, mostExpensive natijasi ishlatilmagan, tur konvertatsiya xatosi |

**Baho:** Interface va structlarni tushungan — bu **kuchli o'quvchi** belgisi.
Lekin spec talablariga to'liq amal qilish (map, delete, float64, type switch)
yetishmadi. Avval **4, 9, 10-masalalarni** to'g'irlang — shundan keyin "data
types'ni generics gacha o'rgandim" degan da'vo to'liq tasdiqlanadi.
