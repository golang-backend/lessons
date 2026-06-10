# 1-dars Vazifalari: O'zgaruvchilar va Ma'lumot Turlari

Salom! Bugungi darsda Go tilida **o'zgaruvchilar** (`var`, `:=`) va **ma'lumot turlari**
(`int`, `float`, `string`, `bool`) bilan tanishdik. Quyidagi masalalar shularni
mustahkamlash uchun. Har bir masalani alohida faylda yeching (masalan, `masala1.go`).

> **Eslatma:** Har bir faylda `package main` va `func main()` bo'lishi kerak.
> Dasturni `go run masala1.go` buyrug'i bilan ishga tushirasiz.

---

## 1-masala: O'zingiz haqingizda
Quyidagi o'zgaruvchilarni e'lon qiling va ekranga chiqaring:
- `ism` (string) — ismingiz
- `yosh` (int) — yoshingiz
- `boy` (float64) — bo'yingiz metrda (masalan, 1.75)
- `talabami` (bool) — talabamisiz yoki yo'q

Hammasini `fmt.Println` bilan chiqaring.

---

## 2-masala: To'g'ri tur tanlash
Quyidagi qiymatlar uchun **eng mos** (eng kam joy egallaydigan) butun son turini tanlang
va o'zgaruvchi e'lon qiling. Nima uchun shu turni tanlaganingizni izoh (`//`) bilan yozing:
- Odamning yoshi: 25
- Bir yildagi kunlar soni: 365
- O'zbekiston aholisi (taxminan): 36000000
- Haroratning manfiy qiymati: -40

---

## 3-masala: `var` va `:=` farqi
Bir xil ikkita o'zgaruvchi e'lon qiling:
- Birinchisini `var` yordamida (turini ko'rsatib).
- Ikkinchisini `:=` (qisqa e'lon) yordamida.

Ikkalasini ham ekranga chiqaring va `:=` ni faqat `func main()` ichida ishlatish
mumkinligini izohda yozing.

---

## 4-masala: Ikki sonni qo'shish
Ikkita `int` o'zgaruvchi e'lon qiling (`a = 17`, `b = 8`). Ularning:
- yig'indisini (`a + b`)
- ayirmasini (`a - b`)
- ko'paytmasini (`a * b`)
- bo'linmasini (`a / b`)

hisoblab, har birini ekranga chiqaring. Bo'linma natijasi nega butun son
chiqayotganini o'ylab ko'ring.

---

## 5-masala: Float bilan ishlash
4-masaladagi `a` va `b` ni endi `float64` qilib e'lon qiling va `a / b` ni
hisoblang. Natija qanday farq qildi? Izohda yozing.

---

## 6-masala: Doira yuzasi
Radius `r = 5` (float64) bo'lsa, doira yuzasini hisoblang.
Formula: `yuza = 3.14159 * r * r`. Natijani ekranga chiqaring.

---

## 7-masala: Maksimal qiymatlar
Quyidagi turlarning **eng katta** qiymatini o'zgaruvchiga berib, ekranga chiqaring:
- `int8`
- `uint8`
- `int16`
- `uint16`

Agar bir birlik oshirsangiz nima bo'lishini taxmin qiling (hozircha kod yozmasangiz
ham bo'ladi, faqat o'ylab ko'ring).

---

## 8-masala: Mantiqiy qiymatlar
`bool` turidagi uchta o'zgaruvchi e'lon qiling:
- `yomgirYogyaptimi` = true
- `quyoshChiqdimi` = false
- `darsTugadimi` = true

Hammasini ekranga chiqaring.

---

## 9-masala: Matnlarni birlashtirish
Ikkita `string` o'zgaruvchi e'lon qiling:
- `ism = "Umidjon"`
- `familiya = "Karimov"`

Ularni `+` yordamida birlashtirib, to'liq ismni bitta o'zgaruvchiga saqlang
(orasida bo'sh joy bo'lsin) va ekranga chiqaring.
Natija: `Umidjon Karimov`

---

## 10-masala: Talaba ma'lumotnomasi
Bitta o'quvchi haqida quyidagi barcha turlardan foydalanib ma'lumot to'plang
va chiroyli ko'rinishda chiqaring:
- `ism` (string)
- `yosh` (int)
- `bahosi` (float64) — masalan, 4.8
- `imtihondanOtdimi` (bool)

`fmt.Println` o'rniga `fmt.Printf` dan foydalanib ko'ring. Masalan:
```go
fmt.Printf("Ism: %s, Yosh: %d, Baho: %.1f, O'tdi: %t\n", ism, yosh, bahosi, imtihondanOtdimi)
```

---

### Qo'shimcha (ixtiyoriy) — qiynalganlar uchun emas, qiziqqanlar uchun:
- 7-masaladagi `int8` ning eng katta qiymatiga `+1` qo'shib, `go run` qilib ko'ring.
  Natija nima bo'ldi? Bu hodisa **overflow** (to'lib ketish) deyiladi.

Omad! 🚀
