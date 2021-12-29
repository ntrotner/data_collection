import pandas as pd
import glob
import os

abbreviations = {
    "49.474627,8.480784-49.481042,8.471755": "1",
    "49.487982,8.476265-49.493906,8.462725": "2",
    "49.503059,8.497637-49.491722,8.481986": "3",
    "49.474701,8.500074-49.483018,8.478691": "4",
    "49.502159,8.459277-49.494586,8.460765": "5"
}


def getTrafficDataFrame() -> pd.DataFrame:
    all_files = glob.glob(os.path.join(os.path.dirname(__file__),
                                       f'{pathToData}/traffic/*.csv'))
    csvs = []

    for filename in all_files:
        df = pd.read_csv(filename, index_col=None, header=0)
        csvs.append(df)

    frame = pd.concat(csvs, axis=0, ignore_index=True)
    frame['date'] = pd.to_datetime(frame['date'])
    frame["id"] = frame["id"].map(abbreviations)

    return frame[["date", "id", "durationInTraffic"]].set_index('date').groupby('id').resample("15T").mean().reset_index().pivot(index="date", columns=["id"], values="durationInTraffic")
