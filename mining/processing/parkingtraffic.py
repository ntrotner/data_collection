import numpy
import plotly.express as px
from plotly.subplots import make_subplots
import plotly.graph_objects as go

def correlate_tp(parking, traffic):
    combined = parking.join(traffic, how='outer').fillna(0)

    corr = combined.replace([numpy.inf, -numpy.inf], numpy.nan).fillna(
        0).corr()[parking.columns].loc[traffic.columns.tolist(), :].fillna(0)


    corr_pct = combined.pct_change().replace([numpy.inf, -numpy.inf], numpy.nan).fillna(
        0).corr()[parking.columns].loc[traffic.columns.tolist(), :].fillna(0)

    fig = make_subplots(rows=2, cols=1, subplot_titles=(
        "Absolute Values", "Changes in Percent"), shared_xaxes=True, vertical_spacing=0.15)

    fig.add_trace(go.Heatmap(x=parking.columns, y=traffic.columns,
                  z=corr, texttemplate="%{text}", colorbar=dict(y=.75,len=.5)), row=1, col=1)
    fig.add_trace(go.Heatmap(x=parking.columns, y=traffic.columns,
                  z=corr_pct, texttemplate="%{text}", colorbar=dict(y=.25,len=.5)), row=2, col=1)

    fig.write_html("./resources/correlation_tp.html")
    fig.write_image("./resources/correlation_tp.png")
