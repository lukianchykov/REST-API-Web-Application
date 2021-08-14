# REST API Web Application для менеджемента списка дел (TODO list) на Go

## Technologies:
- <b> REST API Architecture</b>
- <b> Gin FrameWork</b>
- <b> PostgreSQL</b> 
- <b> Migrating Database Technology</b>
- <b> Docker</b>
- <b> Viper Configuration</b>
- <b> Sqlx for db</b>
- <b> JWT. Middleware</b>

## Tasks:
- [x] <b> Sing-up</b>
- [x] <b> Sing-in</b>
- [x] <b> CreateList</b> 
- [x] <b> GetAllLists</b>
- [x] <b> GetListById</b>
- [x] <b> UpdateList</b>
- [x] <b> DeleteList</b>
- [x] <b> CreateItem</b> 
- [x] <b> GetAllItems</b>
- [x] <b> GetItemById</b>
- [x] <b> UpdateItem</b>
- [x] <b> DeleteItem</b>

### First Install Postman and Docker:
- <a href="https://www.postman.com/downloads/">Postman</a>
- <a href="https://www.docker.com/products/docker-desktop">Docker</a>

### Fix `go.sum`:
1) ![](images/gomod.png)
2) ![](images/gomod2.png)
3) ![](images/gomod3.png)

### Fix format file `wait-for-postgres.sh` for unix systems:
1) ![](images/selectfile.png)
2) ![](images/wait-for-postgresUnix.png)

### Configure `Terminal` to `Project Location`:
```
Set-Location -Path E:\GoLangProjects\WebApplication
```
### PowerShell:
![](images/PowerShellConfig.png)
### Goland:
![](images/GolandConfig.png)
## Для запуска приложения:
```
docker-compose up web-app
```
### PowerShell:
![](images/PowerShell.png)
### Goland Terminal:
![](images/Goland.png)
### Docker Desktop:
![](images/Docker1.png)
![](images/Docker2.png)
## В случаи пустой базы данных или первого запуска, необходимо сделать миграцию в другом `terminal`, где не запущено приложение:
```
migrate -path ./schema -database 'mandarin4ek/webapp:postgresv1.0://postgres:rootdocker@localhost:5436/postgres?sslmode=disable' up
migrate -path ./schema -database 'mandarin4ek/webapp:postgresv1.0://postgres:rootdocker@localhost:5436/postgres?sslmode=disable' down
```

### PowerShell:
![](images/PowerShell2.png)
![](images/PowerShell3.png)
### Goland Terminal:
![](images/Goland2.png)
![](images/Goland3.png)
### Работа приложения Postman:
![](images/Postman.png)
![](images/Postman2.png)
![](images/Postman3.png)
![](images/Postman4.png)
![](images/Postman5.png)
![](images/Postman6.png)
![](images/Postman7.png)
![](images/Postman8.png)
![](images/Postman9.png)
![](images/Postman10.png)
![](images/Postman11.png)
