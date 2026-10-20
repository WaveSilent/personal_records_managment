# Personal Records Management API

Сервис для управления записями сотрудников (Personal Records), написанный на Go. Проект полностью контейнеризирован и использует PostgreSQL для хранения данных.

## 🛠️ Стек технологий
*   **Язык:** Go (Golang)
*   **База данных:** PostgreSQL 18
*   **Контейнеризация:** Docker, Docker Compose
*   **Миграции:** golang-migrate
*   **Архитектура:** Слоистая (Models -> Database -> Business Logic -> HTTP)

## 🏗️ Структура проекта
*   `models/` — доменные модели данных (структура `Worker`).
*   `database/simple_sql/` — слой работы с базой данных (SQL-запросы).
*   `personal_records/` — бизнес-логика (управление записями).
*   `http/` — HTTP-обработчики и роутинг.
*   `migrations/` — SQL-файлы для миграций базы данных.
