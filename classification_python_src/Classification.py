import pandas as pd
import matplotlib.pyplot as plt
from sklearn.cluster import KMeans
from sklearn.preprocessing import StandardScaler
import os

class Classification:
    def read_csv_data(self,file_path):
        """CSVデータを読み込みDataFrameに変換する"""
        return pd.read_csv(file_path)

    def preprocess_data(self,data):
        """心拍数データを抽出し、標準化する"""
        heart_rate = data['bpm'].values.reshape(-1, 1)
        scaler = StandardScaler()
        return scaler.fit_transform(heart_rate)

    def perform_clustering(self,data, num_clusters=2):

        heart_rate_scaled = self.preprocess_data(data)

        """KMeansを用いてクラスタリングを行う"""
        kmeans = KMeans(n_clusters=num_clusters, n_init=10 ,random_state=42)
        kmeans.fit(heart_rate_scaled)
        data['cluster'] = kmeans.labels_
        return data

    def identify_extreme_clusters(self,data):
        """最高値と最低値のクラスタを特定する"""
        high_bpm_cluster = data[data['bpm'] == data['bpm'].max()]['cluster'].values[0]
        low_bpm_cluster = data[data['bpm'] == data['bpm'].min()]['cluster'].values[0]
        return high_bpm_cluster, low_bpm_cluster

    def filter_states(self,data, high_cluster, low_cluster):
        """驚愕と安堵のデータを抽出する"""
        df_surprise = data[(data['diff'] > 1) & (data['cluster'] == high_cluster)]
        df_relief = data[(data['diff'] < -1) & (data['cluster'] == low_cluster)]
        return df_surprise, df_relief

    def update_states(self,data, surprise_data, relief_data, high_cluster, low_cluster):
        """条件に合致したデータの状態を変更する"""
        # clasterがhigh_bpm_clusterの時はnervous、low_bpm_clusterの時はnormalにする
        data.loc[data['cluster'] == high_cluster, 'state'] = 'nervous'
        data.loc[data['cluster'] == low_cluster, 'state'] = 'normal'

        for index in surprise_data.index:
            data.loc[index:index + 5, 'state'] = "surprise"
            data['state'] = data['state'].astype(str)

        for index in relief_data.index:
            data.loc[index:index + 5, 'state'] = "relief"
            data['state'] = data['state'].astype(str)  # 'state'列を文字列型に変換

        return data

    def plot_clusters(self,data):
        """クラスタリング結果をグラフで表示する"""
        plt.figure(figsize=(10, 6))
        for cluster in data['cluster'].unique():
            plt.scatter(data[data['cluster'] == cluster]['time'], data[data['cluster'] == cluster]['bpm'],
                        label=f'Cluster {cluster}')

        plt.scatter(data[data['state'] == 'surprise']['time'], data[data['state'] == 'surprise']['bpm'],
                    label='Surprise', color='red')
        plt.scatter(data[data['state'] == 'relief']['time'], data[data['state'] == 'relief']['bpm'], label='Relief',
                    color='blue')

        plt.xlabel('Time')
        plt.ylabel('Heart Rate (bpm)')
        plt.title('Heart Rate Clustering with Modified Clusters')
        plt.legend()
        plt.show()

    def main(self, df: pd.DataFrame) -> str:

        df['state'] = ''  # 'state'列を空で初期化

        # KMeansでクラスタリングを行う
        clustered_data = self.perform_clustering(df)
        # print(clustered_data)

        clustered_data['diff'] = clustered_data['bpm'].diff()

        # クラスタの特定と状態の更新
        high_cluster, low_cluster = self.identify_extreme_clusters(clustered_data)

        surprise_states, relief_states = self.filter_states(clustered_data, high_cluster, low_cluster)

        updated_data = self.update_states(clustered_data, surprise_states, relief_states, high_cluster, low_cluster)
        # print(updated_data)

        # グラフの描画
        # self.plot_clusters(updated_data)

        # dataの最後の値を取得
        return df.tail(1)["state"].values[0]
