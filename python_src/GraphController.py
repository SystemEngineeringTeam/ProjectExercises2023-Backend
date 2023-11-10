import os

import pandas as pd
import numpy as np
from matplotlib import pyplot as plt
from enum import Enum
from typing import List

class GraphController:

    def __init__(self):
        self.path = "./output_image/"

    """
    グラフを作成する
    @param list データフレーム
    @param option オプション
    @param axis 軸
    @param filter_num フィルタ数
    """

    def plot(self, plt_lists: List[pd.DataFrame],
             boardId: int, color_list: [pd.DataFrame] = pd.DataFrame(data=[])) -> None:
        fig = plt.figure(figsize=(15, 25))
        fig.subplots_adjust(hspace=0.5)

        for index, list_val in enumerate(plt_lists):
            ax = fig.add_subplot(5, 1, index + 1)

            ax.plot(list_val['time'], list_val["bpm"], label="bpm")

            ax.legend()

            # 色付け
            if len(color_list[index].index) != 0:
                self.__color_paint(color_list[index], ax)

            # タイトル
            if index == 0:
                ax.set_title("bpm" + '_' + "north")
            elif index == 1:
                ax.set_title("bpm" + '_' + "west")
            elif index == 2:
                ax.set_title("bpm" + '_' + "south")
            elif index == 3:
                ax.set_title("bpm" + '_' + "east")

            # 単位の指定
            ax.set_xlabel('time [s]')
            ax.set_ylabel('BPM [bpm]')

            ax.grid()

        plt.savefig(self.path + str(boardId) + '.png')

    """
    色付けを行う
    @param list データフレーム データフレームは、start_time, end_time, colorの3つのカラムを持つ
    @param ax matplotlibのax
    """

    def __color_paint(self, color_list: pd.DataFrame, ax: plt.Axes) -> None:
        for index, color_list in color_list.iterrows():
            # 色付け
            # surprise = 赤
            # relief = 青

            color = None
            if color_list['color'] == 'surprise':
                color = 'red'
            elif color_list['color'] == 'relief':
                color = 'blue'
            else:
                color = 'green'
            ax.axvspan(color_list['start_time'], color_list['end_time'], color=color, alpha=0.5)