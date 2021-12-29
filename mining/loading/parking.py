import pandas as pd
import glob
import os


def getParkingDataFrame() -> pd.DataFrame:
    all_files = glob.glob(os.path.join(os.path.dirname(__file__),
                                       '../../etl/data/parking/*.csv'))
    csvs = []

    for filename in all_files:
        df = pd.read_csv(filename, index_col=None, header=0)
        csvs.append(df)

    frame = pd.concat(csvs, axis=0, ignore_index=True)
    
    names = {}
    for i in list(set(frame["identifier"].iloc)):
        df = pd.read_csv(
            f'../etl/data/parking/meta/{i}', index_col=None, header=0)
        names[i] = df["title"].iloc[0]

    frame["identifier"] = frame["identifier"].map(names)
    frame['date'] = pd.to_datetime(frame['date'])
    return frame.set_index('date').groupby('identifier').resample("15T").mean().reset_index().pivot(index="date", columns=["identifier"], values="free_slots")
