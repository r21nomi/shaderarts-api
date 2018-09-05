# shaderarts-api
API for ShaderArts

## Get started
### Setup for Firebase
Place `serviceAccountKey.json` file under root dir. 

### Setup for MySQL
```
$ export RDS_USER=xxxx
$ export RDS_PASS=xxxx
$ export RDS_PROTOCOL=xxxx
$ export RDS_DBNAME=xxxx
```
```
$ mysql.server start
```

### Start
```
$ go run main.go
```