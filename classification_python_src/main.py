import HeartRateDate as HRD
import SqlHandler as SH
import pandas as pd
import Classification as CL
import sys


def main():
    args = sys.argv

    try:
        azimuth = HRD.Azimuth(str(args[1]))
    except:
        print('azimuthを指定してください')
        return

    sqlHandler = SH.SqlHandler()
    heartRateData = HRD.HeartRateData(sqlHandler.cur)
    classification = CL.Classification()

    getAllHeatRateData = heartRateData.get_all_heart_rate_data(azimuth)
    # print(getAllHeatRateData)

#     dfに変換
    df = pd.DataFrame(getAllHeatRateData)
    state = classification.main(df)

    # 状態を出力
    print(state)



if __name__ == "__main__":
    main()