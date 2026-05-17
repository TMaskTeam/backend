```markdown
# T-Mask

## Описание продукта

T-Mask — это платформа лояльности, интегрированная в экосистему Т-Банка, которая позволяет малому бизнесу создавать и управлять программами лояльности, а клиентам — участвовать в них без необходимости устанавливать отдельные приложения или запоминать коды.

---

### Technology Stack

| Layer | Technology |
|-------|------------|
| Language | Go 1.21+ |
| Web Framework | Fiber v3 |
| ORM | GORM |
| Database | PostgreSQL 14 |
| Auth | JWT (HttpOnly cookies) |
| Password Hashing | bcrypt |
| Migrations | Goose |
| Container | Docker + Docker Compose |

### Architectural Principles

1. **Clean Architecture** - Handler - Service - Repository
2. **Dependency Injection**
3. **JWT Authentication** - jwt храниться в HttpOnly cookies (XSS protection)
4. **Role-Based Access** - `business_owner`, `client`, `admin` роли

---

## 🚀 Быстрый старт

### Запуск через Docker

```bash
git clone https://github.com/TMaskTeam/backend.git
cd backend

docker-compose up -d --build

sh src/scripts/migrations.sh --up

```

### Локальный запуск

```bash
# Установить зависимости
go mod download

# Накатить миграции
sh src/scripts/migrations.sh --up

# Запустить приложение
go run src/cmd/main.go
```

### Переменные окружения (.env)

```env
# Server
SERVER_PORT=8080

# Database
DATABASE_USER=test_user
DATABASE_PASSWORD=secret
DATABASE_DBNAME=test_db
DATABASE_HOST=localhost
DATABASE_PORT=5432

# JWT
JWT_SECRET=your-super-secret-key-change-in-production
```

---

## 📡 API Endpoints

### Базовый URL

```
http://localhost:8080/api/v1
```

### Аутентификация

Для защищённых эндпоинтов используется **HttpOnly cookie**, которая устанавливается автоматически после логина. Браузер сам отправляет её при каждом запросе.

```
Cookie: token=<JWT_TOKEN>
```

---

### 1. Аутентификация

#### POST /auth/owner/register — Регистрация владельца бизнеса

**Request:**
```json
{
    "first_name": "Иван",
    "last_name": "Петров",
    "middle_name": "Иванович",
    "inn": "123456789012",
    "phone_number": "+79001234567",
    "email": "ivan@example.com",
    "birthday": "1990-01-15",
    "password": "MyP@ssw0rd123"
}
```

**Response (201 Created):**
```json
{
    "owner_id": 1,
    "first_name": "Иван",
    "last_name": "Петров",
    "inn": "123456789012",
    "phone_number": "+79001234567",
    "email": "ivan@example.com",
    "birthday": "1990-01-15T00:00:00Z"
}
```

**Validation Rules:**

| Field | Rules |
|-------|-------|
| `first_name` | Required, min 2 chars |
| `last_name` | Required, min 2 chars |
| `inn` | Required, 10 or 12 digits |
| `phone_number` | Required, valid format |
| `email` | Required, valid email |
| `password` | Required, min 8 chars |

---

#### POST /auth/owner/login — Вход владельца бизнеса

**Request:**
```json
{
    "login": "ivan@example.com",
    "password": "MyP@ssw0rd123"
}
```

**Response (200 OK):**
```json
{
    "owner_id": 1,
    "first_name": "Иван",
    "last_name": "Петров",
    "inn": "123456789012",
    "phone_number": "+79001234567",
    "email": "ivan@example.com",
    "birthday": "1990-01-15T00:00:00Z"
}
```
*JWT token устанавливается как HttpOnly cookie автоматически.*

---

#### POST /auth/client/register — Регистрация клиента

**Request:**
```json
{
    "first_name": "Анна",
    "last_name": "Сидорова",
    "phone_number": "+79009999999",
    "email": "anna@example.com",
    "birthday": "1995-05-20",
    "password": "ClientPass123"
}
```

**Response (201 Created):**
```json
{
    "client_id": 1,
    "first_name": "Анна",
    "last_name": "Сидорова",
    "phone_number": "+79009999999",
    "email": "anna@example.com",
    "birthday": "1995-05-20T00:00:00Z"
}
```

---

#### POST /auth/client/login — Вход клиента

**Request:**
```json
{
    "login": "anna@example.com",
    "password": "ClientPass123"
}
```

**Response (200 OK):**
```json
{
    "client_id": 1,
    "first_name": "Анна",
    "last_name": "Сидорова",
    "phone_number": "+79009999999",
    "email": "anna@example.com",
    "birthday": "1995-05-20T00:00:00Z"
}
```

---

#### POST /auth/logout — Выход

**Request:** (требует авторизации)
```
Cookie: token=<JWT_TOKEN>
```

**Response:** `204 No Content`

---

### 2. Профиль

#### GET /me — Получить текущий профиль

**Request:** (требует авторизации)
```
Cookie: token=<JWT_TOKEN>
```

**Response для владельца бизнеса (200 OK):**
```json
{
    "role": "business_owner",
    "owner_id": 1,
    "first_name": "Иван",
    "last_name": "Петров",
    "inn": "123456789012",
    "phone_number": "+79001234567",
    "email": "ivan@example.com",
    "birthday": "1990-01-15T00:00:00Z"
}
```

**Response для клиента (200 OK):**
```json
{
    "role": "client",
    "client_id": 1,
    "first_name": "Анна",
    "last_name": "Сидорова",
    "phone_number": "+79009999999",
    "email": "anna@example.com",
    "birthday": "1995-05-20T00:00:00Z"
}
```

---

#### PUT /me — Обновить профиль

**Request:** (требует авторизации)
```json
{
    "first_name": "Петр",
    "last_name": "Сидоров",
    "phone_number": "+79001112233",
    "email": "petr@example.com",
    "password": "NewPassword123"
}
```
*Все поля опциональны*

**Response (200 OK):** Тот же формат, что и GET `/me` с обновлёнными данными

---

### 3. Управление бизнесом (только для владельца бизнеса)

#### POST /businesses — Создать бизнес

**Request:** (требует авторизации)
```json
{
    "name": "Кофе Хаус",
    "address": "ул. Тверская, 15"
}
```

**Response (201 Created):**
```json
{
    "business_id": 1,
    "owner_id": 1,
    "name": "Кофе Хаус",
    "address": "ул. Тверская, 15",
    "created_at": "2025-05-17T10:30:00Z",
    "updated_at": "2025-05-17T10:30:00Z"
}
```

---

#### GET /businesses — Получить все мои бизнесы

**Request:** (требует авторизации)

**Response (200 OK):**
```json
{
    "businesses": [
        {
            "business_id": 1,
            "name": "Кофе Хаус",
            "address": "ул. Тверская, 15"
        },
        {
            "business_id": 2,
            "name": "Пиццерия",
            "address": "ул. Арбат, 10"
        }
    ]
}
```

---

#### DELETE /businesses/{id} — Удалить бизнес

**Request:** (требует авторизации)
```
DELETE /api/v1/businesses/1
```

**Response:** `204 No Content`

---

### 4. Бонусные программы

#### POST /businesses/{business_id}/programs — Создать программу лояльности

**Request:** (требует авторизации — владелец бизнеса)
```json
{
    "program_name": "Кофейная карта",
    "token_name": "кофеины"
}
```

**Response (201 Created):**
```json
{
    "program_id": 1,
    "business_id": 1,
    "program_name": "Кофейная карта",
    "token_name": "кофеины"
}
```

---

#### GET /businesses/{business_id}/programs — Получить программы бизнеса

**Request:**

**Response (200 OK):**
```json
{
    "programs": [
        {
            "program_id": 1,
            "business_id": 1,
            "program_name": "Кофейная карта",
            "token_name": "кофеины"
        }
    ]
}
```

---

#### GET /businesses/programs — Получить все программы

**Request:**

**Response (200 OK):**
```json
{
    "programs": [
        {
            "program_id": 1,
            "business_id": 1,
            "business_name": "Кофе Хаус",
            "program_name": "Кофейная карта",
            "token_name": "кофеины"
        }
    ]
}
```

---

#### POST /programs/{program_id}/join — Вступить в программу

**Request:** (требует авторизации — роль `client`)

**Response (200 OK):**
```json
{
    "message": "successfully joined the program"
}
```

---

## 📊 Коды ответов

| HTTP Status | Описание |
|-------------|----------|
| 200 | OK — успешный запрос |
| 201 | Created — ресурс создан |
| 204 | No Content — успешный запрос без тела ответа |
| 400 | Bad Request — неверный запрос |
| 401 | Unauthorized — отсутствует или неверный токен |
| 403 | Forbidden — недостаточно прав |
| 404 | Not Found — ресурс не найден |
| 409 | Conflict — конфликт (дубликат) |
| 500 | Internal Server Error — ошибка сервера |

### Формат ошибки

```json
{
    "error": "inn is already used"
}
```

---

## 🗄️ Структура базы данных

### Основные таблицы

| Таблица | Описание |
|---------|----------|
| `business_owner` | Владельцы бизнеса (ИП) |
| `client` | Клиенты Т-Банка |
| `business` | Точки/заведения бизнеса |
| `bonus_program` | Программы лояльности |
| `bonus_program_info` | Настройки программ |
| `client_bonus_program` | Участие клиентов в программах |

### Связи

```
business_owner (1) ────── (N) business
business (1) ────── (N) bonus_program
client (1) ────── (N) client_bonus_program (N) ────── (1) bonus_program
```

---

## 🔐 Безопасность

| Аспект | Реализация |
|--------|------------|
| Пароли | bcrypt хэширование |
| Токены | JWT в HttpOnly cookies (XSS защита) |
| Срок жизни токена | 24 часа |
| Доступ по ролям | Проверка `role` в JWT |
| CORS | Настраиваемые allowed origins |

---

## 📁 Структура проекта

```
backend/
├── cmd/server/main.go          # Точка входа
├── internal/
│   ├── app/app.go              # Сборка приложения, роуты
│   ├── config/                 # Конфигурация
│   ├── context/                # Адаптер для HandlerContext
│   ├── db/                     # Подключение к БД
│   ├── domain/                 # Бизнес-сущности
│   ├── dto/                    # DTO для API
│   ├── model/                  # GORM модели
│   ├── repository/             # Репозитории (DAO)
│   ├── service/                # Бизнес-логика
│   ├── handler/                # HTTP обработчики
│   ├── middleware/             # Middleware (Auth, Adapt)
│   ├── provider/               # DI контейнер
│   └── validator/              # Валидация
├── pkg/
│   ├── jwt/                    # JWT генерация/валидация
│   └── password/               # bcrypt хэширование
├── migration/                  # SQL миграции
└── scripts/                    # Скрипты (миграции)
```

---

## 👥 Команда

- **Tech Lead** — [Имя]
- **Backend Developer** — [Имя]
- **Frontend Developer** — [Имя]

---

## 📄 Лицензия

© 2025 T-Mask Team. Все права защищены.
```