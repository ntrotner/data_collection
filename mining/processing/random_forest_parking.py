import pandas as pd
from sklearn.ensemble import RandomForestRegressor
from sklearn.model_selection import train_test_split
from sklearn import metrics
from sklearn import tree

import plotly.express as px
from plotly.subplots import make_subplots
import plotly.graph_objects as go

def random_forest_regression(parking, traffic, offset):
    combined = parking.join(traffic, how='outer').fillna(0)
    
    x = []
    for time, a, b, c, d, f in zip(combined['1'].index.to_list(), combined['1'].to_list(), combined['2'].to_list(), combined['3'].to_list(), combined['4'].to_list(), combined['5'].to_list()):
        x.append([time.minute, time.hour, time.day, time.month, time.year, time.dayofweek, time.day_of_week, int(time.timestamp()), int(a), int(b), int(c), int(d), int(f), time])
    for _ in range(offset):
        x.pop()
    results = {}
    
    for i in parking.columns:
        y = combined['Kunsthalle, Tiefgarage'].to_list()
        y = [int(yi) for yi in y][offset:]
        X_train, X_test, y_train, y_test = train_test_split(x, y, test_size=0.1, shuffle=False)
        
        clf = RandomForestRegressor(n_estimators=750)
        clf = clf.fit([z[:-1] for z in X_train], y_train)
        
        #Predict the response for test dataset
        y_pred = clf.predict([z[:-1] for z in X_test])
        r2 = metrics.r2_score(y_test, y_pred)
        results[i] = r2
        print(f'{i}: {r2}')
        
        
        df = pd.DataFrame(dict(
            x = [z[-1] + pd.DateOffset(hours=(15*offset)//60, minutes=(15*offset)%60) for z in X_test],
            y_p = y_pred.tolist(),
            y_t = y_test
        )).sort_values(by="x")

        fig = make_subplots(specs=[[{"secondary_y": True}]])

        fig.add_trace(
            go.Scatter(x=df["x"], y=df["y_p"], name=f'Predicted Data'),
            secondary_y=False
        )

        fig.add_trace(
            go.Scatter(x=df["x"], y=df["y_t"], name=f'Real Data'),
            secondary_y=False
        )

        fig.update_layout(title_text=f'{i} with R²: {r2}')
        fig.write_html(f'./resources/{i}.html')
        fig.update_layout(height=1080, width=1920)
        fig.write_image(f'./resources/{i}.png')