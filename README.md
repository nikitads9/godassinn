# godassinn
Букинг сервис.
## Инструкция по запуску
Конфигурационный файл .env.example переименовываем в .env
В классическом случае (без ELK-стека) для запуска применяем
```
set -o allexport && source ./.env && set +o allexport
docker-compose up -d
```

В случае, если хотим запустить версию с ELK,
```
set -o allexport && source ./.env && set +o allexport
docker-compose-elk up setup -d
docker-compose-elk up -d
```
Команду `docker-compose-elk up setup -d` нужно применять только при первоначальной настройке.


Пароль _"changeme"_ , установленный по умолчанию в env файле **небезопасен**. Для того, чтобы сгенерировать случайные пароли, нужно выполнить следующие шаги:

1. Сбросить пароли для встроенных пользователей

    Команды ниже сбрасывают пароли встроенных пользователей `elastic`, `logstash_internal` и `kibana_system` и возвращают строку с новыми паролями.

    ```sh
    docker-compose exec elasticsearch bin/elasticsearch-reset-password --batch --user elastic
    ```

    ```sh
    docker-compose exec elasticsearch bin/elasticsearch-reset-password --batch --user logstash_internal
    ```

    ```sh
    docker-compose exec elasticsearch bin/elasticsearch-reset-password --batch --user kibana_system
    ```

2. Заменяем пароли в конфигурационном файле .env на сгенерированные ранее.
