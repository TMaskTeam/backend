# T-Mask

## Описание продукта

T-Mask — это платформа лояльности, интегрированная в экосистему Т-Банка, которая позволяет малому бизнесу создавать и управлять программами лояльности, а клиентам — участвовать в них без необходимости устанавливать отдельные приложения или запоминать коды.

---

### Стек технологий

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

### Архитектурные принципы

1. **Clean Architecture** - Handler - Service - Repository
2. **Dependency Injection**
3. **JWT Authentication** - jwt храниться в HttpOnly cookies (XSS protection)
4. **Role-Based Access** - `business_owner`, `client`, `admin` роли

---

## Старт веб-приложения

### Запуск через Docker

```bash
git clone https://github.com/TMaskTeam/backend.git
cd backend

docker-compose up -d --build

sh src/scripts/migrations.sh --up

```

### Локальный запуск

```bash
go mod download

sh src/scripts/migrations.sh --up

go run src/cmd/main.go
```

### Переменные окружения (.env)

```env
SERVER_PORT=8080

DATABASE_USER=test_user
DATABASE_PASSWORD=secret
DATABASE_DBNAME=test_db
DATABASE_HOST=localhost
DATABASE_PORT=5432

JWT_SECRET=secret
```

---

## Эндпоинты API

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

Вот **все недостающие эндпоинты** в том же формате. Добавь их в конец секции API.

---


#### GET /businesses/{id} — Получить бизнес по ID

**Request:** (требует авторизации — владелец бизнеса)
```
GET /api/v1/businesses/1
Cookie: token=<JWT_TOKEN>
```

**Response (200 OK):**
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

#### PUT /businesses/{id} — Обновить бизнес

**Request:** (требует авторизации — владелец бизнеса)
```json
{
    "name": "Кофе Хаус Updated",
    "address": "ул. Тверская, 20"
}
```
*Все поля опциональны*

**Response (200 OK):**
```json
{
    "business_id": 1,
    "owner_id": 1,
    "name": "Кофе Хаус Updated",
    "address": "ул. Тверская, 20",
    "created_at": "2025-05-17T10:30:00Z",
    "updated_at": "2025-05-17T12:00:00Z"
}
```

---


#### GET /programs/{program_id} — Получить программу по ID

**Request:**
```
GET /api/v1/programs/1
```

**Response (200 OK):**
```json
{
    "program_id": 1,
    "business_id": 1,
    "business_name": "Кофе Хаус",
    "program_name": "Кофейная карта",
    "token_name": "кофеины",
    "created_at": "2025-05-17T10:30:00Z",
    "updated_at": "2025-05-17T10:30:00Z"
}
```

---

#### PUT /programs/{program_id} — Обновить программу

**Request:** (требует авторизации — владелец бизнеса)
```json
{
    "program_name": "Кофейная карта Platinum",
    "token_name": "платина"
}
```
*Все поля опциональны*

**Response (200 OK):**
```json
{
    "program_id": 1,
    "business_id": 1,
    "business_name": "Кофе Хаус",
    "program_name": "Кофейная карта Platinum",
    "token_name": "платина",
    "created_at": "2025-05-17T10:30:00Z",
    "updated_at": "2025-05-17T12:00:00Z"
}
```

---

#### DELETE /programs/{program_id} — Удалить программу

**Request:** (требует авторизации — владелец бизнеса)
```
DELETE /api/v1/programs/1
```

**Response:** `204 No Content`

---


#### POST /programs/{program_id}/settings — Добавить настройки программы

**Request:** (требует авторизации — владелец бизнеса)
```json
{
    "visit_tokens": 1,
    "percentage_purchase_tokens": 5,
    "register_tokens": 10,
    "birthday_tokens": 50,
    "friend_invite_tokens": 20,
    "minimum_receipt_limit": 500
}
```

**Response (201 Created):**
```json
{
    "program_info_id": 1,
    "program_id": 1,
    "visit_tokens": 1,
    "percentage_purchase_tokens": 5,
    "register_tokens": 10,
    "birthday_tokens": 50,
    "friend_invite_tokens": 20,
    "minimum_receipt_limit": 500,
    "created_at": "2025-05-17T10:30:00Z",
    "updated_at": "2025-05-17T10:30:00Z"
}
```

---

#### GET /programs/{program_id}/settings — Получить настройки программы

**Request:**

**Response (200 OK):**
```json
{
    "program_info_id": 1,
    "program_id": 1,
    "visit_tokens": 1,
    "percentage_purchase_tokens": 5,
    "register_tokens": 10,
    "birthday_tokens": 50,
    "friend_invite_tokens": 20,
    "minimum_receipt_limit": 500,
    "created_at": "2025-05-17T10:30:00Z",
    "updated_at": "2025-05-17T10:30:00Z"
}
```

---

#### PUT /programs/{program_id}/settings — Обновить настройки программы

**Request:** (требует авторизации — владелец бизнеса)
```json
{
    "visit_tokens": 2,
    "percentage_purchase_tokens": 7,
    "minimum_receipt_limit": 300
}
```
*Все поля опциональны*

**Response (200 OK):** Тот же формат, что и GET, с обновлёнными данными

---


#### GET /client/programs — Мои программы (для клиента)

**Request:** (требует авторизации — роль `client`)
```
Cookie: token=<JWT_TOKEN>
```

**Response (200 OK):**
```json
{
    "programs": [
        {
            "program_id": 1,
            "business_name": "Кофе Хаус",
            "program_name": "Кофейная карта",
            "token_name": "кофеины",
            "balance": 125,
            "joined_at": "2025-05-17T10:30:00Z"
        }
    ]
}
```

---

#### GET /client/programs/{program_id} — Моя карта программы (баланс, скидка)

**Request:** (требует авторизации — роль `client`)
```
GET /api/v1/client/programs/1
Cookie: token=<JWT_TOKEN>
```

**Response (200 OK):**
```json
{
    "program_id": 1,
    "business_name": "Кофе Хаус",
    "program_name": "Кофейная карта",
    "token_name": "кофеины",
    "balance": 125,
    "total_earned": 450,
    "total_spent": 325,
    "total_visits": 23,
    "current_discount": 1.2,
    "joined_at": "2025-05-17T10:30:00Z",
    "last_visit_at": "2025-05-17T15:30:00Z"
}
```

---


#### GET /programs/{program_id}/participants — Список участников программы

**Request:** (требует авторизации — владелец бизнеса)
```
GET /api/v1/programs/1/participants
Cookie: token=<JWT_TOKEN>
```

**Response (200 OK):**
```json
{
    "participants": [
        {
            "client_id": 1,
            "first_name": "Анна",
            "last_name": "Сидорова",
            "phone_number": "+79009999999",
            "email": "anna@example.com",
            "balance": 125,
            "total_visits": 23,
            "joined_at": "2025-05-17T10:30:00Z",
            "last_visit_at": "2025-05-17T15:30:00Z"
        }
    ],
    "total": 1
}
```

---

#### GET /programs/{program_id}/participants/{client_id} — Данные участника в программе

**Request:** (требует авторизации — владелец бизнеса)
```
GET /api/v1/programs/1/participants/1
Cookie: token=<JWT_TOKEN>
```

**Response (200 OK):**
```json
{
    "client_id": 1,
    "first_name": "Анна",
    "last_name": "Сидорова",
    "phone_number": "+79009999999",
    "email": "anna@example.com",
    "balance": 125,
    "total_earned": 450,
    "total_spent": 325,
    "total_visits": 23,
    "current_discount": 1.2,
    "joined_at": "2025-05-17T10:30:00Z",
    "last_visit_at": "2025-05-17T15:30:00Z"
}
```

---


#### POST /loyalty/earn — Начислить токены

**Request:** (требует авторизации — webhook или админ)
```json
{
    "client_id": 1,
    "program_id": 1,
    "ruble_amount": 500,
    "description": "Оплата в кафе"
}
```
*Либо `ruble_amount`, либо `tokens`*

**Response (200 OK):**
```json
{
    "transaction_id": 1,
    "client_id": 1,
    "program_id": 1,
    "amount": 25,
    "new_balance": 150,
    "description": "Оплата 500₽, начислено 25 кофеинов"
}
```

---

#### POST /loyalty/spend — Списать токены

**Request:** (требует авторизации — клиент или кассир)
```json
{
    "client_id": 1,
    "program_id": 1,
    "tokens": 50,
    "description": "Скидка 50₽ на чек"
}
```

**Response (200 OK):**
```json
{
    "transaction_id": 2,
    "client_id": 1,
    "program_id": 1,
    "amount": -50,
    "new_balance": 100,
    "description": "Списано 50 кофеинов"
}
```

---

#### GET /client/programs/{program_id}/balance — Баланс токенов

**Request:** (требует авторизации — клиент)
```
GET /api/v1/client/programs/1/balance
Cookie: token=<JWT_TOKEN>
```

**Response (200 OK):**
```json
{
    "program_id": 1,
    "program_name": "Кофейная карта",
    "token_name": "кофеины",
    "balance": 125,
    "total_earned": 450,
    "total_spent": 325
}
```

---

### Формат ошибки

```json
{
    "error": "inn is already used"
}
```

---

## Структура базы данных

### Основные таблицы

| Таблица | Миграция |
|---------|----------|
| `business_owner` | [`20260515071000_create_business_owner_table.sql`](../migration/20260515071000_create_business_owner_table.sql) |
| `client` | [`20260516145306_create_client_table.sql`](../migration/20260516145306_create_client_table.sql) |
| `business` | [`20260515132902_business_table.sql`](../migration/20260515132902_business_table.sql) |
| `bonus_program` | [`20260515132947_bonus_program_table.sql`](../migration/20260515132947_bonus_program_table.sql) |
| `bonus_program_info` | [`20260515133020_bonus_program_info_table.sql`](../migration/20260515133020_bonus_program_info_table.sql) |
| `client_bonus_program` | [`20260516145347_create_client_bonus_table.sql`](../migration/20260516145347_create_client_bonus_table.sql) |
---

## Безопасность

| Аспект | Реализация |
|--------|------------|
| Пароли | bcrypt хэширование |
| Токены | JWT в HttpOnly cookies (XSS защита) |
| Срок жизни токена | 24 часа |
| Доступ по ролям | Проверка `role` в JWT |
| CORS | Настраиваемые allowed origins |

---



## Лицензия

Лицензия: [MIT](LICENSE)

2025 T-Mask Team.
```