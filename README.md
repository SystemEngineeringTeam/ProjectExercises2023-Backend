麻雀センシングのバックエンド
===

# 環境構築
## mysqlを実行させる場合
```bash
$ docker-compose up -d
```

## GoのModuleをダウンロードする
```bash
$ cd ./src
$ go mod tidy
```

# Goの実行
```bash
$ cd ./src
$ go run main.go
```

# swaggerの確認
下記URLを確認する
URL: http://localhost:8080/api/v1/swagger/index.html

## 登録の仕方
コメントを記載して、下記コマンドを実行する
```bash
$ swag init
```
# 動作確認
## ゲーム開始
http://localhost:8080/api/v1/start/

## ゲーム終了
http://localhost:8080/api/v1/finish/

# MySQLモニタの立ち上げ
```bash
$ cd ./src
$ docker exec -it mahjong_sensing_DB bash
$ mysql -u root -p
```