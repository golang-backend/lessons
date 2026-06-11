# Daraja Tekshiruv Masalalari: struct, array, slice, map, interface

Bu masalalar Go asoslarini (data types, generics gacha) o'rgangan o'quvchining
**amaliy darajasini aniqlash** uchun. Har bir masala bitta asosiy tushunchani
sinaydi. Kodni `go run` bilan ishlata olishingiz va natijani tushuntira olishingiz
muhim.

> **Maslahat:** Faqat ishlaydigan kod yetarli emas — har bir masalada "nega
> aynan shu turni (slice / map / interface) tanladingiz?" degan savolga javob
> bera olishingiz kerak.

---

## A daraja — Array va Slice

### 1-masala: Slice bilan o'rtacha qiymat
Butun sonlardan iborat slice bering: `[]int{12, 7, 25, 3, 18}`.
- Slice ichidagi barcha sonlar yig'indisini hisoblang.
- O'rtacha (`average`) qiymatni toping.
- `len()` dan foydalaning, sonlar sonini qattiq (hardcoded) yozmang.

### 2-masala: Slice'ni o'stirish
Bo'sh slice (`var s []int`) yarating va `append` yordamida 1 dan 10 gacha
sonlarni qo'shing. So'ng faqat **juft** sonlarni yangi slice'ga ajratib oling
va chiqaring.

> Tekshiriladi: `append`, `for` sikli, `%` operatori, slice'ni filtrlash.

---

## B daraja — Map

### 3-masala: So'zlarni sanash
Quyidagi slice berilgan:
```go
mevalar := []string{"olma", "anor", "olma", "uzum", "anor", "olma"}
```
`map[string]int` yarating va har bir meva nechta marta uchraganini sanang.
Natija: `olma: 3, anor: 2, uzum: 1`.

### 4-masala: Telefon kitobchasi
`map[string]string` ko'rinishida telefon kitobchasi yarating (ism → raqam).
- Kamida 3 ta kontakt qo'shing.
- Berilgan ism bo'yicha raqamni qidiring va **topilmasa** "Topilmadi" deb chiqaring
  (`value, ok := m[key]` shaklini ishlating).
- Bitta kontaktni `delete` bilan o'chiring.

> Tekshiriladi: map yaratish, `comma-ok` idiomasi, `delete`.

---

## C daraja — Struct

### 5-masala: Talaba structi
`Student` nomli struct yarating: `Ism string`, `Yosh int`, `Baho float64`.
- 3 ta talaba yarating va ularni `[]Student` slice'ga joylang.
- Slice bo'ylab aylanib (`for range`), bahosi 4.0 dan yuqori talabalarni chiqaring.

### 6-masala: Method bilan struct
5-masaladagi `Student` ga method qo'shing:
```go
func (s Student) Otdimi() bool { ... }
```
Agar baho 3.0 dan baland bo'lsa `true` qaytarsin. Har bir talaba uchun
`talaba.Otdimi()` ni chaqirib natijani chiqaring.

> Tekshiriladi: method receiver, value vs pointer tushunchasi.

### 7-masala: Ichma-ich (nested) struct
`Manzil` structi yarating (`Shahar`, `Kocha`). Keyin `Student` ichiga
`Manzil` ni maydon sifatida joylang. Bitta talaba yaratib, uning shahrini
`talaba.Manzil.Shahar` orqali chiqaring.

---

## D daraja — Interface (eng muhim qism)

### 8-masala: Shakl yuzasi (klassik)
`Shape` (Shakl) interfeysini yarating:
```go
type Shape interface {
    Yuza() float64
}
```
- `Doira` (radius bilan) va `Tortburchak` (eni va bo'yi bilan) structlarini yarating.
- Ikkalasi ham `Yuza()` methodini amalga oshirsin.
- `[]Shape` slice yaratib, har bir shaklning yuzasini bitta `for` sikli ichida chiqaring.

> Bu masala interfeysni **chinakam tushunganini** ko'rsatadi — har xil structlarni
> bitta interface orqali bir xil ishlata olish.

### 9-masala: `interface{}` va tur tekshiruvi
`[]interface{}{42, "salom", 3.14, true}` slice'ini bering.
Har bir element bo'ylab aylaning va `switch v := x.(type)` (type switch)
yordamida har bir elementning **turini aniqlab**, mos xabar chiqaring.
Masalan: `42 — bu butun son (int)`.

> Tekshiriladi: empty interface, type assertion / type switch.

---

## E daraja — Hammasini birlashtirish (yakuniy)

### 10-masala: Mini do'kon tizimi
Quyidagilarni birlashtiring:
- `Mahsulot` structi: `Nomi string`, `Narxi float64`, `Soni int`.
- `[]Mahsulot` slice — kamida 4 ta mahsulot.
- `map[string]float64` — mahsulot nomi → umumiy qiymati (`Narxi * Soni`).
- Eng qimmat mahsulotni topadigan funksiya yozing.
- `Qimmatbaho() bool` methodi: narxi 100000 dan oshsa `true`.

Hammasini `main` ichida ishlatib, do'kon hisobotini chiroyli chiqaring.

> Bu masalani toza yecha olsa — o'quvchi struct, slice, map va method'larni
> **amalda erkin qo'llay oladi** degani.

---

## Daraja baholash mezoni (o'qituvchi uchun)

| Yechilgan masalalar | Daraja |
|---------------------|--------|
| 1–4 | Asoslarni biladi, slice/map bilan ishlay oladi |
| 5–7 | Struct va method'larni tushunadi |
| 8–9 | Interface'ni real tushunadi (kuchli o'quvchi) |
| 10  | Mustaqil loyiha yoza oladi, keyingi mavzularga tayyor |

> **Eslatma:** O'quvchi 8 va 9-masalalarni mustaqil yecha olsa, "data types'ni
> generics gacha o'rgandim" degan da'vosi haqiqatga to'g'ri keladi. Aks holda
> interface mavzusini qaytarish kerak.
