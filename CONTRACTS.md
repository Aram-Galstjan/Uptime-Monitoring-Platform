# Контракты взаимодействия (Спецификация)

## 1. База данных (Таблицы)

Все сервисы (Python API, Go Scheduler, Go Checker) используют одну базу PostgreSQL.

### Таблица `users`
* `id` (integer, primary key)
* `email` (string, unique)

### Таблица `monitors`
* `id` (integer, primary key) — ID отслеживаемого сайта
* `url` (string) — Ссылка (например, https://example.com)
* `interval` (integer) — Периодичность проверки в секундах (по умолчанию 60)
* `is_active` (boolean) — Включен ли мониторинг (true/false)

### Таблица `checks`
* `id` (integer, primary key)
* `monitor_id` (integer, foreign key) — Ссылка на ID из таблицы monitors
* `status_code` (integer) — HTTP-код ответа (200, 404, 500 и т.д.)
* `response_time_ms` (integer) — Время отклика сайта в миллисекундах
* `created_at` (timestamp) — Дата и время проверки