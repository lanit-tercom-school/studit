# StudIT WIKI

## Backend

### Setup

1. Установить [PostgreSQL 9.6 (Database)](https://www.postgresql.org/download/), (!!! обязательно проверьте, что первые 2 цифры версии, которую вы устанавливаете : 9.6.)  во время установки введите password: *postgres* 
2. Установить [Go](https://golang.org/dl/)
3. Изменить переменную PATH (изменить → создать) `%GOPATH%\bin\ ` [ссылка на инструкцию](https://www.java.com/ru/download/help/path.xml) 
4. Создать системную переменную 
[GOPATH](https://github.com/golang/go/wiki/GOPATH), написав в графе значение путь до *папки, куда скачан проект*\backend {например: C:\Repos\studit\backend} (!!! Обязательно после 3-го и 4-го пунктов перезагрузить консоль, иначе эти изменения не подхватятся)
5. Установить [Bee (Beego client)](https://github.com/beego/bee), введя в консоль команду `go get -u github.com/beego/bee`
6. Установить [Dep (Go package manager)](https://github.com/golang/dep), введя в консоль команду `go get -u github.com/golang/dep/cmd/dep`
7. Установить пакеты, запустив файл: installpkg.cmd (studit\backend\src)
8. Инициализировать базы данных, запустив файл: initdbs.cmd (studit\backend\src)

### Develop launch
1. Запустить backend, открыв файл: run-backend.cmd (studit\backend\src)

## Frontend

### Setup

1. Установить [NodeJS](https://nodejs.org/en/) (с npm)
1. Установить все frontend packages: из папки *repository_path*\frontend выполнить команду `npm install`

### Develop launch

1. Запустить frontend: из папки *repository_path*\frontend командой `npm start`
1. [Открыть frontend в браузере](http://localhost:4200) 

asd
asda
sdas
da
sda
s

