# Сервис управления сегментами пользователей

## Описание

Сервис для управления сегментами пользователей

* **Framework**: Fiber
* **ORM**: Gorm
* **Logs**: Zero-log
* **SQL dialect**: pgSql

## Сервис работает со следующими сущностями:

* Пользователь
* Сегмент

## Запуск

### Config

Конфиг задается переменными окружения

| Name                 | Description                       | Default value            | Expected value     | Requiered |
|:---------------------|:----------------------------------|:-------------------------|:-------------------|:---------:|
| APP_PORT             | Port which servers`s listening on | 8080                     | app port           |    ✔️     |
| REPORT_DIRECTORY     | Directory to save reports         | ./reports                | directory path     |    ✔️     |
| DB_HOST              | Database sever host               | localhost                | any string         |    ✔️     |
| DB_PORT              | Database sever port               | 5432                     | db port            |    ✔️     |
| DB_NAME              | Database name                     | -                        | db name            |    ✔️     |
| DB_USER              | Username for database             | -                        | db admin username  |    ✔️     |
| DB_PWD               | Password for database             | -                        | db admin password  |    ✔️     |
| DB_SSL_MODE          | Disable or enable SSL mode        | disable                  | `enable`/`disable` |    ✔️     |
| DB_TIMEZONE          | Database timezone                 | Europe/Moscow            | timezone           |    ✔️     |
| DB_DNS               | Database DNS                      | -                        | any string         |     ❌     |

Для запуска в docker необходимо указать переменные в файле .env

Изначально в нем заданы все переменные, необходимые для запуска, но при желании их можно изменить

### Запуск сервиса

1. Клонируем проект
```Shell
git clone https://github.com/vindosVP/users-segments-service.git
```
2. Запускаем

```Shell
cd users-segments-service
```

```Shell
docker-compose up
```
### API

Документация к апи: https://documenter.getpostman.com/view/20758355/2s9Y5ZwhNY

Сваггер можно найти в директории `src/docs`
