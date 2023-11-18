import HeartRateDate as HRD
import SqlHandler as SH
import pandas as pd
import Classification as CL
import sys
from flask import Flask, g, request, jsonify

app = Flask(__name__)
sqlHandler = SH.SqlHandler()

@app.route('/sensing', methods=['GET'])
def get_sensing():

    print("get_sensing")
    req = request.args
    azimuth = req.get("azimuth")

    heartRateData = HRD.HeartRateData(sqlHandler.cur)
    classification = CL.Classification()

    getAllHeatRateData = heartRateData.get_all_heart_rate_data(HRD.Azimuth(azimuth))
    # print(getAllHeatRateData)

    #     dfに変換
    df = pd.DataFrame(getAllHeatRateData)
    state = classification.main(df)

    # 状態を出力
    print(state)

    data = {
        "state": state
    }

    return jsonify(data)

if __name__ == "__main__":
    app.run(port=10080, debug=True)
