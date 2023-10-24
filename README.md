麻雀センシングのバックエンド
===

# 環境構築
## mysqlを実行させる場合
```bash
$ docker-compose up -d
```

## GoのModuleをダウンロードする
```bash
$ go mod tidy
```

# Goの実行
```bash
$ cd backend/src
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
