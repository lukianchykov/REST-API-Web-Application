# REST API Web Application для менеджемента списка дел (TODO list) на Go

## Technologies:
- <b> REST API Architecture</b>
- <b> Gin FrameWork</b>
- <b> Postgresql</b> 
- <b> Migrating Database Technology</b>
- <b> Docker</b>
- <b> Viper configuration</b>
- <b> Sqlx for db</b>
- <b> JWT. Middleware</b>

### Для запуска приложения:

- Необходимо запустить <b>Docker container</b> с <b>database</b> https://hub.docker.com/repository/docker/mandarin4ek/webapp
- Версия: <b>postgresv1.0</b>
- Запустить сам проект

### В случаи пустой базы данных, необходимо сделать миграцию:


```
migrate -path ./schema -database 'mandarin4ek/webapp:postgresv1.0://postgres:rootdocker@localhost:5436/postgres?sslmode=disable' up
migrate -path ./schema -database 'mandarin4ek/webapp:postgresv1.0://postgres:rootdocker@localhost:5436/postgres?sslmode=disable' down
```

