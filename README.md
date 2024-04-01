# godassinn

Букинг сервис.

## Инструкция по запуску

Конфигурационный файл **.env.example** переименовываем в **.env** <br />
В классическом случае (без ELK-стека) для запуска применяем:

```
set -o allexport && source ./.env && set +o allexport
docker-compose up -d
```
<details>
<summary> 
Запуск с сбором логов в ELK
</summary>
В случае, если хотим запустить версию с ELK,то необходимо раскомментировать следующие строчки в файле **docker-compose.yml** в конфигурации Jaeger:
    
- `- SPAN_STORAGE_TYPE=elasticsearch`
- `- ES_TAGS_AS_FIELDS_ALL=true`
- `- ES_SERVER_URLS=http://elasticsearch:9200`
- `- ES_USERNAME=elastic`
- `- ES_PASSWORD=${ELASTIC_PASSWORD}`
    
```
set -o allexport && source ./.env && set +o allexport
docker-compose -f docker-compose-elk.yml  up setup -d
docker-compose -f docker-compose-elk.yml  up -d
```
Команду `docker-compose-elk up setup -d` нужно применять только при первоначальной настройке.


Пароль _"changeme"_ , установленный по умолчанию в **.env** файле **небезопасен**. Для того, чтобы сгенерировать случайные пароли, нужно выполнить следующие шаги:

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

2. Заменяем пароли в конфигурационном файле **.env** на сгенерированные ранее.
</details>

## Инструкция по изменению структуры базы данных

Структура базы данных описана в двух местах: 
- ` deploy/database/init.sql ` 
- ` deploy/migrations/*.sql ` 

Первый файл не используется при создании базы данных в среде разработки и тестирования, он нужен только для наглядности и общего понимания текущей структуры базы данных. Если вы изменили стрктуру базы данных, то, пожалуйста, измените ее и в этом файле. <br />
Если вы хотите изменить структуру БД: добавить новые таблицы или поменять текущие, то необходимо создать новый файл формата **ГГГГММДДЧЧММСС_описание_коммита.sql** в директории ` deploy/migrations `. <br />
Перед перечислением изменений, которые вы хотите внести, необходимо прописать ` -- +goose Up `. <br />
Также необходимо прописать и откат этих изменений (обратные операции), предварив их ` -- +goose Down `. <br />

Чтобы применить изменения, просто пропишите, хотите ли вы накатить миграцию (up) или откатить миграцию (down) в скрипте ` deploy/migrations/migration.sh ` и запустите ранее созданный контейнер **migrator**. В логах контейнера вы увидите статус исполнения миграции.