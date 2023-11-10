import os

import pandas as pd
import numpy as np
from matplotlib import pyplot as plt
from enum import Enum
from typing import List

class SensingController:

    """
    状態を計算する
    @param list データフレーム
    """
    def classification(self, heart_rate_df_diff: pd.DataFrame) -> List:
        current_interval = None
        color_intervals = []  # ここで color_intervals を初期化

        for index, row in heart_rate_df_diff.iterrows():
            if current_interval is None:
                current_interval = [row['time'], None, row['status']]
            elif current_interval[2] != row['status']:
                current_interval[1] = row['time']
                color_intervals.append(current_interval)
                current_interval = [row['time'], None, row['status']]

        if current_interval is not None:
            current_interval[1] = heart_rate_df_diff['time'].iloc[-1]
            color_intervals.append(current_interval)

        return color_intervals

    """
    ローパスフィルタ
    @param list データフレーム
    @param window_size ウィンドウサイズ
    @return フィルタ後のデータフレーム
    """
    def low_filter(self, low_filter_list: pd.DataFrame, window_size: int) -> pd.DataFrame:

        for column in low_filter_list:
            if column == 'time':
                continue
            low_filter_list[column] = low_filter_list[column].rolling(window_size).mean()

        return low_filter_list