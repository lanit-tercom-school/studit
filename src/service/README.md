#Backend Team

---
##Backend Requirements    
Для успешного запуска локального сервера требуется
* [golang](https://golang.org/)
* [postgresql](https://www.postgresql.org/)

---
##Start Service   
Запускаем сервер:
1. Клонируем репозиторий
* Переходим на ветку develop
* Создаём бд с именем `studit`, для которой нужно выполнить содержимое файла `/schema/data_base_init.sql` для создания таблиц (username: `postgres`, password: `postgres`)
* Переходим в папку `Path/To/Repository/src/service/`
* Запускаем консоль (этот и пункт выше можно поменять местами)
* Выполнить
```
InstallPackages.cmd
```
* Выполнить
 ```
 cd schema
 ```
* Выполнить
```
go run db_init_data.go
```
* Выполнить
```
cd ..
```
* Выполнить
```
go run.cmd
```
* ...
* Profit
