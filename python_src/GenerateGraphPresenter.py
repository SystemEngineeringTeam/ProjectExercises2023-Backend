from typing import List

import pandas as pd

import GraphController as GC
import ReadCSVController as RCC
import SensingController as SC

class GenerateGraphPresenter:

    def __init__(self, boardId: int):
        self.boardId = boardId
        self.readCSVController = RCC.ReadCSVController()
        self.graphController = GC.GraphController()
        self.sensingController = SC.SensingController()

    def diff(self, df_list: List[pd.DataFrame]) -> List[pd.DataFrame]:
        df_diff_list = []

        for df in df_list:
            heart_rate_df_diff = df.copy()
            heart_rate_df_diff["bpm"] = heart_rate_df_diff["bpm"].diff()
            # ソートする
            heart_rate_df_diff = heart_rate_df_diff.sort_values("time")
            heart_rate_df_diff = self.sensingController.low_filter(heart_rate_df_diff, 10)
            df_diff_list.append(heart_rate_df_diff)

        return df_diff_list

    def conditionClassification(self, df_emotion_list: List[pd.DataFrame]) -> List[pd.DataFrame]:
        color_list = []

        for df in df_emotion_list:
            color_intervals = self.sensingController.classification(df)
            print(color_intervals)

            color_list.append(pd.DataFrame(
                data=color_intervals,
                columns=["start_time", "end_time", "color"]
            ))

        return color_list

    def generateGraph(self) -> None:

        # CSVファイルを読み込む
        df_list = self.readCSVController.getBpmDf(self.boardId)
        df_emotion_list = self.readCSVController.getEmotionDf(self.boardId)

        # 微分を行う
        df_diff_list = self.diff(df_list)

        # 状態を分類する
        color_list = self.conditionClassification(df_emotion_list)

        self.graphController.plot(
            plt_lists=df_list,
            boardId=self.boardId,
            color_list=color_list
        )