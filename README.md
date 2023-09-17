# go-gorm
This is my second GOLANG project, i made this project by using :
- GORM
- GIN
- POSTGRES

### How to run :
- create a .env file : 

```
PORT=3000
DB_URL="host=localhost user= password= dbname=go-gorm port=5432"
JWT_SECRET="skmeasionrs1391-3ns"
```

- migrate :
```
go run migrate/migrate.go
```

- run compile daemon :
```
CompileDaemon -command="./go-gorm"
```
you're all set ;D