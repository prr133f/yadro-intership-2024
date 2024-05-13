# Тестовое задание Yadro 2024
## Инструкция по запуску в Docker
Для начала нужно скопировать репозиторий:
```sh
git clone git@github.com:prr133f/yadro-intership-2024.git
```
Расположите в папке `/static/` файл с входными данными в формате `.txt`

Затем соберите образ приложения:
```sh
make build-image
```
И запустите его 
```sh
make run file=$(FILE_FROM_STATIC_DIR)
# Пример запуска
# make run file=data.txt
```