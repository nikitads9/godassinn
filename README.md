# godassinn
Букинг сервис.
## Инструкция по запуску
Конфигурационный файл **.env.example** переименовываем в **.env**
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
