from typing import NamedTuple
import datetime
from enum import Enum


class Azimuth(Enum):
    north = "north"
    west = "west"
    south = "south"
    east = "east"


class HeartRateDataModel(NamedTuple):
    id: int
    board_surface_id: int
    azimuth: Azimuth
    bpm: int
    time: int


class HeartRateData():

    def __init__(self, cur):
        self.cur = cur

    def get_last_board_id(self) -> int:
        self.cur.execute("SELECT MAX(board_surface_id) FROM heart_rate_data")
        result = self.cur.fetchall()  # 結果を取得
        return result[0][0] if result else None

    def get_all_heart_rate_data(self, azimuth: Azimuth) -> list[HeartRateDataModel]:
        board_id = self.get_last_board_id()
        self.cur.execute(
            "SELECT * FROM heart_rate_data WHERE board_surface_id = %s AND azimuth = %s",
            (board_id, azimuth.value)
        )

        heart_rate_data_list = []
        for row in self.cur.fetchall():
            heart_rate_data_list.append(HeartRateDataModel(
                id=row[0],
                board_surface_id=row[1],
                time=row[2].timestamp(),
                azimuth=Azimuth(row[3]),
                bpm=row[4],

            ))
        return heart_rate_data_list