import mysql.connector
from sqlalchemy.pool import QueuePool


class SqlHandler:
    def __init__(self):
        try:
            self.conn = mysql.connector.connect(
                user='root',  # ユーザー名
                password='admin',  # パスワード
                host='127.0.0.1',  # ホスト名(IPアドレス）
                db='mahjong_sensing',
                port='3309'
            )

            # self.cur = self.conn.cursor()

            self.cnxpool = QueuePool(lambda: self.conn, pool_size=10)
            # 接続成功ログ出力
            print('接続できました。')

        except:
            # print("DB接続失敗")
            exit(1)

    def show_tables(self):
        #  テーブルを全部出力
        self.cur.execute("show tables")
        for row in self.cur:
            print(row)
