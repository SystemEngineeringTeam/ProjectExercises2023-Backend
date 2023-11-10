import os

import pandas as pd
import numpy as np
from matplotlib import pyplot as plt
from enum import Enum
from typing import List

class ReadCSVController:

    def __init__(self):
        self.path = "../output_csv/"

    def getBpmDf(self, board_id: int) -> List[pd.DataFrame]:
        folder_name = self.path + "/" + str(board_id)

        df_north = pd.read_csv(os.path.join(folder_name, "north_bpm.csv"))
        df_west = pd.read_csv(os.path.join(folder_name, "west_bpm.csv"))
        df_south = pd.read_csv(os.path.join(folder_name, "south_bpm.csv"))
        df_east = pd.read_csv(os.path.join(folder_name, "east_bpm.csv"))

        df_list = [df_north, df_west, df_south, df_east]
        return df_list

    def getEmotionDf(self, board_id: int) -> List[pd.DataFrame]:
        folder_name = self.path + "/" + str(board_id)

        df_emotion_north = pd.read_csv(os.path.join(folder_name, "north_emotion.csv"))
        df_emotion_west = pd.read_csv(os.path.join(folder_name, "west_emotion.csv"))
        df_emotion_south = pd.read_csv(os.path.join(folder_name, "south_emotion.csv"))
        df_emotion_east = pd.read_csv(os.path.join(folder_name, "east_emotion.csv"))

        df_emotion_list = [df_emotion_north, df_emotion_west, df_emotion_south, df_emotion_east]
        return df_emotion_list
