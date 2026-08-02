## Запуск проекта в Docker

### Что нужно перед запуском

- Установленный Docker
- Открытый Docker daemon: Docker Desktop, OrbStack или другой совместимый runtime
- Запуск команд из корня проекта, где лежит `docker-compose.yml`

### Сначала проверьте Docker context

Перед запуском контейнеров нужно понять, какой Docker context сейчас активен. Контейнеры будут созданы и отображены именно в нем.

Проверить текущий context:

```bash
docker context ls
```

Если нужен запуск через Docker Desktop:

```bash
docker context use desktop-linux
```

Если нужен запуск через OrbStack:

```bash
docker context use orbstack
```

### Быстрый запуск

```bash
docker compose up -d --build
```

### Проверка, что контейнеры поднялись

```bash
docker compose ps
```

Ожидаемые сервисы:

- `monitoring_api`
- `monitoring_db`
- `monitoring_redis`
- `monitoring_scheduler`
- `monitoring_checker`

### Проверка API

```bash
curl http://localhost:8000/
curl http://localhost:8000/health
```

Ожидаемые ответы:

```json
{"status":"ok","service":"api","message":"container is running"}
```

```json
{"status":"healthy"}
```

### Где смотреть контейнеры

Контейнеры отображаются в том приложении, чей Docker context был активен в момент запуска.

### Остановка проекта

```bash
docker compose down
```

### Полезно знать

- API доступен на `localhost:8000`
- PostgreSQL доступен на `localhost:5433`
- Redis доступен на `localhost:6380`
- `scheduler` и `checker` сейчас запускаются как контейнеры, но их бизнес-логика пока не реализована полностью
