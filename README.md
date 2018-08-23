# 環境
- docker-compose 1.6.2

# 手順
- 起動
```
$ docker-compose up
```
- 起動確認
```
$ curl localhost:5000/api/v1/heartBeat
```

#
- DBサーバーに接続
```
$ docker-compose exec go-api-db bash
# psql
```
- APIサーバーに接続
```
$ docker-compose exec go-api bash
```

- APIを叩く
```
$ curl -X GET    http://localhost:5000/api/v1/accounts
$ curl -X POST   http://localhost:5000/api/v1/accounts?name=hoge
$ curl -X PUT    http://localhost:5000/api/v1/accounts?name=hoge\&id=1
$ curl -X DELETE http://localhost:5000/api/v1/accounts?id=1
```