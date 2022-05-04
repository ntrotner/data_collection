import plotly.express as px
from plotly.subplots import make_subplots
import plotly.graph_objects as go

def plot_traffic(traffic):
    fig = go.Figure()

    for i in traffic.columns:
        fig.add_trace(go.Scatter(x=traffic.index,
                      y=traffic[i], mode="lines", name=i))
    fig.update_layout(margin=dict(l=5, r=5, t=35, b=5))
    fig.write_html("./resources/traffic.html")
    fig.update_layout(height=1080, width=1920)
    fig.write_image("./resources/traffic.png")
