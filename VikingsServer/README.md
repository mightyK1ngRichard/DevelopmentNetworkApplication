# Лабораторная работа №3
 <div>
 <img src="https://img.shields.io/badge/language-GoLang-blue.svg" alt="Go Language">
 </div>

- **Цель работы**: создание веб-сервиса в бэкенде для использования его в `SPA`
- **Порядок показа**: выполнить GET списка, сделать POST новой записи, показать новые данные через select. Объяснить модели, сериализаторы, контроллеры, роутеры для методов веб-сервиса
- **Контрольные вопросы**: веб-сервис, REST, RPC, HTTP, OSI ISO
- **Диаграмма классов** с детализацией бэкенда (домены методов по `url` с интерфейсами, модели, таблицы БД) + insomnia/postman
- **Задание**: Создание веб-сервиса со всей итоговой бизнес логикой, но без авторизации, подключение его к БД и тестирование в `insomnia`/`swagger`/`postman`

Создание **веб-сервиса** для получения/редактирования данных из вашей БД. Для изображений `услуг` использовать `Minio` или хранение файлов картинок в бинарном виде в БД.

Требуется разработать все методы для реализации итоговой бизнес логики вашего приложения. Методы и `url` в `API` должны соответствовать `REST`. Для списка `услуг` и `заявок` нужно предусмотреть фильтрацию на бэкенде. Для логических действий в приложении (оплата, подтверждение, завершение) предусмотреть отдельные методы для обновления конкретных полей (статусы нельзя менять с любого на любой).

## Выбранная тема:
Походы викингов. Услуги - города, заявки - походы викингов

## Эндпоинты:

```http
GET     http://localhost:7070/api/v3/cities
GET     http://localhost:7070/api/v3/cities?city=1
POST    http://localhost:7070/api/v3/cities
PUT     http://localhost:7070/api/v3/cities
DELETE  http://localhost:7070/api/v3/cities
```

```http
GET     http://localhost:7070/api/v3/hikes
GET     http://localhost:7070/api/v3/hikes?hike=2
POST    http://localhost:7070/api/v3/hikes
PUT     http://localhost:7070/api/v3/hikes
DELETE  http://localhost:7070/api/v3/hikes
```

```http
GET     http://localhost:7070/api/v3/vikings
GET     http://localhost:7070/api/v3/vikings?viking=2
POST    http://localhost:7070/api/v3/vikings
PUT     http://localhost:7070/api/v3/vikings
DELETE  http://localhost:7070/api/v3/vikings
```
