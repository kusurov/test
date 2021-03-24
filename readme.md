## Описание
Выполненное тестовое задание.

## Использованные либы
- [ozzo-validation v3.6.0](https://github.com/go-ozzo/ozzo-validation)
- [gorilla/mux v1.8.0](https://github.com/gorilla/mux)
- [gorilla/sessions v1.2.1](https://github.com/gorilla/sessions)
- [logrus v1.8.1](https://github.com/sirupsen/logrus)
- [crypto](https://pkg.go.dev/golang.org/x/crypto)
- [yaml v2.4.0](gopkg.in/yaml.v2)

## Миграции
Для применения миграций использую [golang-migrate](https://github.com/golang-migrate/migrate)
```sh
migrate -path migrations -database "mysql://uname:upass@tcp(host:port)/dbname" up
```
## Запуск проекта
```sh
go run cmd/main.go
```

## Затраченное рвемя
В общем на изучение материалов и разработку ушло примерно 24 часа
