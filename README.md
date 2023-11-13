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
$ go run ./main.go
```

# pythonでグラフを作成する方法

基本的には、Go側から自動的に呼び出されるが、手動で実験するための方法である。
CSVのデータを`output_csv`に保存する。
main.pyの後についている番号はボードIDである

```bash
$ cd ./python_src
$ python ./main.py 3
```

# swaggerの確認
下記URLを確認する
URL: http://localhost:8080/api/v1/swagger/index.html

# PostmanのURL
https://project-exer-2023-sysken.postman.co/workspace/My-Workspace~700b9801-fec4-4f25-8b56-157aacad80cb/request/30829807-797c3f20-51f9-4d4e-a1e2-2b05991ce420

## 登録の仕方
コメントを記載して、下記コマンドを実行する
```bash
$ cd src
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