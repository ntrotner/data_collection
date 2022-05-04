import pandas as pd
import glob
import os


def getParkingDataFrame(pathToData) -> pd.DataFrame:
    all_files = glob.glob(f'{pathToData}/parking/*.csv')
    csvs = []

    for filename in all_files:
        df = pd.read_csv(filename, index_col=None, header=0)
        csvs.append(df)

    frame = pd.concat(csvs, axis=0, ignore_index=True)
    
    names = {}
    for i in list(set(frame["identifier"].iloc)):
        df = pd.read_csv(
            f'{pathToData}/parking/meta/{i}', index_col=None, header=0)
        names[i] = df["title"].iloc[0]

    frame["identifier"] = frame["identifier"].map(names)
    frame['date'] = pd.to_datetime(frame['date'])
    frame = frame.set_index('date')
    frame.index = pd.to_datetime(frame.index, utc=True)
    return frame.groupby('identifier').resample("15T").mean().reset_index().pivot(index="date", columns=["identifier"], values="free_slots")
