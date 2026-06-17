#!/usr/bin/env python3
# Lesson3 amaliy topshiriq PDF generatori (sklad tizimi).
from fpdf import FPDF

FONT_DIR = "/usr/share/fonts/truetype/dejavu"

class PDF(FPDF):
    def header(self):
        if self.page_no() == 1:
            return
        self.set_font("DejaVu", "", 8)
        self.set_text_color(150, 150, 150)
        self.cell(0, 8, "Lesson 3 — Amaliy topshiriq: Sklad tizimi", align="R")
        self.ln(10)

    def footer(self):
        self.set_y(-15)
        self.set_font("DejaVu", "", 8)
        self.set_text_color(150, 150, 150)
        self.cell(0, 10, f"{self.page_no()}", align="C")


pdf = PDF()
pdf.add_font("DejaVu", "", f"{FONT_DIR}/DejaVuSans.ttf")
pdf.add_font("DejaVu", "B", f"{FONT_DIR}/DejaVuSans-Bold.ttf")
pdf.add_font("Mono", "", f"{FONT_DIR}/DejaVuSansMono.ttf")
pdf.set_auto_page_break(auto=True, margin=18)
pdf.add_page()

W = pdf.epw  # effective page width

def h1(text):
    pdf.set_x(pdf.l_margin)
    pdf.set_font("DejaVu", "B", 18)
    pdf.set_text_color(20, 40, 90)
    pdf.multi_cell(0, 9, text)
    pdf.ln(2)

def h2(text):
    pdf.ln(2)
    pdf.set_x(pdf.l_margin)
    pdf.set_font("DejaVu", "B", 13)
    pdf.set_text_color(20, 40, 90)
    pdf.multi_cell(0, 7, text)
    pdf.ln(1)

def h3(text):
    pdf.set_x(pdf.l_margin)
    pdf.set_font("DejaVu", "B", 11)
    pdf.set_text_color(40, 40, 40)
    pdf.multi_cell(0, 6, text)

def body(text):
    pdf.set_x(pdf.l_margin)
    pdf.set_font("DejaVu", "", 10.5)
    pdf.set_text_color(30, 30, 30)
    pdf.multi_cell(0, 5.6, text)
    pdf.ln(1)

def bullet(text):
    pdf.set_x(pdf.l_margin)
    pdf.set_font("DejaVu", "", 10.5)
    pdf.set_text_color(30, 30, 30)
    pdf.cell(6, 5.6, "•")
    pdf.multi_cell(W - 6, 5.6, text)

def code(text):
    pdf.ln(1)
    pdf.set_x(pdf.l_margin)
    pdf.set_font("Mono", "", 9)
    pdf.set_fill_color(244, 244, 246)
    pdf.set_text_color(20, 20, 20)
    pdf.multi_cell(0, 5, text, fill=True, border=0)
    pdf.ln(1)

def rule():
    pdf.ln(1)
    pdf.set_draw_color(210, 210, 210)
    pdf.line(pdf.l_margin, pdf.get_y(), pdf.l_margin + W, pdf.get_y())
    pdf.ln(3)


# ---------- TITLE ----------
h1("Amaliy topshiriq: Konsol Sklad Tizimi")
pdf.set_font("DejaVu", "", 11)
pdf.set_text_color(90, 90, 90)
pdf.multi_cell(0, 6, "Lesson 3 — struct, interface, map va slice'ni amalda birlashtirish")
pdf.ln(1)
body("Bu topshiriqda siz terminalda ishlaydigan kichik sklad (ombor) boshqaruv "
     "dasturini yozasiz. Dastur ishga tushganda foydalanuvchidan login so'raydi; "
     "login to'g'ri bo'lsagina menyu ochiladi va mahsulotlar bilan ishlash mumkin "
     "bo'ladi. Maqsad — struct va interface'ni real, foydali dasturda qo'llash.")
rule()

# ---------- MAQSAD ----------
h2("Maqsad")
bullet("struct yordamida mahsulot ma'lumotlarini modellashtirish.")
bullet("interface yordamida turli amallarni (qo'shish, ko'rish, o'zgartirish) yagona shaklga keltirish.")
bullet("map va slice'da ma'lumotlarni saqlash va boshqarish.")
bullet("Cheksiz menyu sikli (for) va foydalanuvchi bilan terminal orqali muloqot qilish.")
rule()

# ---------- 1. LOGIN ----------
h2("1-bosqich: Login (kirish)")
body("Dastur ishga tushganda foydalanuvchidan login va parol so'rasin. To'g'ri "
     "qiymatlar kodda oldindan yozilgan bo'lsin (hardcoded):")
code("Login:  admin\nParol:  12345")
body("Talablar:")
bullet("Login yoki parol noto'g'ri bo'lsa, \"Xato! Qaytadan urinib ko'ring\" deb chiqarsin.")
bullet("Foydalanuvchiga 3 marta urinish berilsin. 3 martadan keyin dastur to'xtasin.")
bullet("Login to'g'ri bo'lsa, \"Xush kelibsiz!\" deb yozsin va asosiy menyuni ochsin.")
body("Maslahat: terminaldan o'qish uchun bufio.NewReader(os.Stdin) yoki "
     "fmt.Scanln dan foydalaning.")
rule()

# ---------- 2. STRUCT ----------
h2("2-bosqich: Mahsulot structi")
body("Mahsulot ma'lumotlarini saqlash uchun struct yarating:")
code("type Mahsulot struct {\n"
     "    ID    int\n"
     "    Nomi  string\n"
     "    Narxi float64\n"
     "    Soni  int\n"
     "}")
bullet("Har bir mahsulotning takrorlanmas (unique) ID raqami bo'lsin.")
bullet("Mahsulotlarni []Mahsulot slice'da yoki map[int]Mahsulot da saqlang "
       "(qaysi birini nega tanlaganingizni tushuntira oling).")
rule()

# ---------- 3. INTERFACE ----------
h2("3-bosqich: Interface")
body("Sklad ustidagi amallarni bitta interface ostida birlashtiring. Masalan:")
code("type Sklad interface {\n"
     "    Qoshish(m Mahsulot)\n"
     "    Korish()\n"
     "    NarxOzgartirish(id int, yangiNarx float64)\n"
     "    Ochirish(id int)\n"
     "}")
bullet("Bu interface'ni amalga oshiradigan struct yarating (masalan Ombor).")
bullet("Ombor ichida mahsulotlar ro'yxati saqlansin.")
bullet("Har bir method o'z vazifasini bajarsin (qo'shish, ko'rish va h.k.).")
body("Eslatma: ma'lumotni o'zgartiradigan method'larda pointer receiver "
     "(*Ombor) ishlatish kerakligini o'ylab ko'ring.")
rule()

# ---------- 4. MENU ----------
h2("4-bosqich: Asosiy menyu")
body("Login muvaffaqiyatli bo'lgach, quyidagi menyu cheksiz siklda (for) "
     "ko'rsatilsin. Foydalanuvchi raqam tanlaydi:")
code("===== SKLAD MENYU =====\n"
     "1. Mahsulot qo'shish\n"
     "2. Mahsulotlarni ko'rish\n"
     "3. Narxni o'zgartirish\n"
     "4. Mahsulotni o'chirish\n"
     "5. Eng qimmat mahsulotni topish\n"
     "6. Chiqish\n"
     "Tanlang (1-6): ")
h3("Har bir variant nima qilishi kerak:")
bullet("1 — Foydalanuvchidan nom, narx, son so'rab, yangi mahsulot qo'shsin (ID avtomatik berilsin).")
bullet("2 — Barcha mahsulotlarni jadval ko'rinishida chiroyli chiqarsin (ID, nom, narx, son).")
bullet("3 — ID bo'yicha mahsulotni topib, yangi narx kiritsin. Topilmasa \"Mahsulot topilmadi\" desin.")
bullet("4 — ID bo'yicha mahsulotni o'chirsin.")
bullet("5 — Slice/map bo'ylab aylanib, narxi eng baland mahsulotni topib chiqarsin.")
bullet("6 — \"Dastur tugadi\" deb yozib, sikldan chiqsin (return yoki break).")
body("Noto'g'ri raqam kiritilsa, \"Noto'g'ri tanlov\" deb yozib, menyuni qайta ko'rsatsin.")
rule()

# ---------- 5. QOSHIMCHA ----------
h2("Qo'shimcha (ixtiyoriy — kuchli o'quvchilar uchun)")
bullet("map[string]float64 yordamida har bir mahsulotning umumiy qiymatini (Narxi × Soni) hisoblab chiqaring.")
bullet("\"Umumiy sklad qiymati\" (barcha mahsulotlar jami) ni hisoblaydigan funksiya qo'shing.")
bullet("Qimmatbaho() bool methodi: narxi 100000 dan oshsa true qaytarsin.")
bullet("Mahsulot qo'shganda nom bo'sh yoki narx manfiy bo'lmasligini tekshiring (validatsiya).")
rule()

# ---------- BAHOLASH ----------
h2("Baholash mezoni")
pdf.set_font("DejaVu", "", 10)
pdf.set_text_color(30, 30, 30)
rows = [
    ("Login 3 urinish bilan ishlaydi", "15%"),
    ("Mahsulot structi to'g'ri tuzilgan", "15%"),
    ("Interface yaratilgan va amalga oshirilgan", "25%"),
    ("Menyu va barcha amallar ishlaydi", "30%"),
    ("Kod toza, nomlar mazmunli, xatolik ushlangan", "15%"),
]
pdf.set_fill_color(20, 40, 90)
pdf.set_text_color(255, 255, 255)
pdf.set_font("DejaVu", "B", 10)
pdf.cell(W - 28, 7, " Talab", border=1, fill=True)
pdf.cell(28, 7, "Ball", border=1, fill=True, align="C")
pdf.ln()
pdf.set_text_color(30, 30, 30)
pdf.set_font("DejaVu", "", 10)
fill = False
for name, pct in rows:
    pdf.set_fill_color(238, 241, 248)
    pdf.cell(W - 28, 7, f" {name}", border=1, fill=fill)
    pdf.cell(28, 7, pct, border=1, align="C", fill=fill)
    pdf.ln()
    fill = not fill
pdf.ln(3)

body("Topshirish: faylni sklad.go deb nomlang. Dasturni \"go run sklad.go\" "
     "bilan ishga tushira olishingiz va har bir tanlovni jonli ko'rsata "
     "olishingiz kerak. Nega struct/interface/map ishlatganingizni og'zaki "
     "tushuntira olish — topshiriqning eng muhim qismi.")

out = "/home/murtazo/go/src/github.com/golang-backend/lessons/lesson3/sklad_topshiriq.pdf"
pdf.output(out)
print("Yaratildi:", out)
