# Simple Login with Golang

یک سیستم ساده لاگین و ثبت‌نام با **Go** بدون استفاده از فریم‌ورک‌های جانبی — تنها با کتابخانه‌های استاندارد و کمی کدنویسی امن.

---

## ✨ ویژگی‌ها
- ثبت‌نام (Signup) کاربران
- ورود (Login) با اعتبارسنجی رمز عبور
- ذخیره‌ی امن رمز با **bcrypt**
- استفاده از SQLite (یا هر دیتابیس دیگر)
- ساختار ماژولار:
  - `configs/`
  - `utils/`
  - `apps/user/`
  - `middlewares/`
- مدیریت JWT برای احراز هویت (اختیاری)

---

## 📂 ساختار پروژه
```
simple-login-with-golang/
├── configs/          # تنظیمات دیتابیس و کانکشن
├── utils/            # توابع عمومی (هش کردن رمز و ...)
├── apps/user/        # روترها و هندلرهای کاربر (Login, Signup)
├── middlewares/      # میدل‌ورها (مثل JSON parser)
├── server/           # راه‌اندازی سرور HTTP
├── go.mod
├── go.sum
└── README.md
```

---

## 🚀 شروع سریع

### 1. کلون کردن پروژه
```bash
git clone https://github.com/MahdiMalvandi/simple-login-with-golang.git
cd simple-login-with-golang
```

### 2. نصب وابستگی‌ها
```bash
go mod tidy
```

### 3. پیکربندی دیتابیس  
(مثلاً در فایل `configs/database.go` یا فایل `.env`)

### 4. اجرای برنامه
```bash
go run server/main.go
```

---

## 🔑 استفاده از API

### ثبت‌نام (Signup)  
`POST /signup`

**Request**
```json
{
  "first_name": "Ali",
  "last_name": "Reza",
  "username": "ali123",
  "password": "mypassword"
}
```

**Response**
```json
{
  "message": "User created successfully"
}
```

---

### ورود (Login)  
`POST /login`

**Request**
```json
{
  "username": "ali123",
  "password": "mypassword"
}
```

**Response**
```json
{
  "token": "<JWT_TOKEN_HERE>"
}
```

---

## 🛠️ مسیرها (Routers)
```
POST /signup       ->  SignUpHandler
POST /login        ->  LoginHandler
POST /check-token  ->  CheckJwtTokenHandler (اختیاری)
GET  /users        ->  ListUsersHandler (فقط برای admin)
```

---

## 🔒 نکات امنیتی
- رمزهای عبور با **bcrypt** هش می‌شوند.
- فیلد `Password` با تگ `json:"-"` در مدل کاربر مخفی است.
- در صورت استفاده از JWT، تنها **توکن** به کاربر برگردانده می‌شود.
