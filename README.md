# Uptime Monitor (Go)

Простой сервис мониторинга доступности сайтов, написанный на Go. Он регулярно проверяет указанные URL-адреса, отслеживает изменения состояния (UP ↔ DOWN), отправляет уведомления в Telegram и предоставляет метрики для Prometheus.

 Учебный pet-проект для изучения Go, DevOps-интеграции (Prometheus, Docker) и Telegram API.

## ⚙️ Возможности

- Периодическая проверка HTTP/HTTPS-сайтов
- Хранение текущих статусов в памяти
- Уведомления в Telegram при изменении состояния
- Метрики `/metrics` для Prometheus
- Конфигурация через YAML
- Поддержка Docker и docker-compose


## 📝 Конфигурация

Файл: `configs/config.yaml`

```yaml
interval_seconds: 60

telegram_token: "your-telegram-bot-token"
telegram_chat_id: 123456789

targets:
  - name: "Google"
    url: "https://www.google.com"
  - name: "GitHub"
    url: "https://github.com"

---