from loading.parking import getParkingDataFrame
from loading.traffic import getTrafficDataFrame
from processing.random_forest_parking import random_forest_regression
from processing.parkingtraffic import correlate_tp
from processing.plot_parking import plot_parking
from processing.plot_traffic import plot_traffic

parking = getParkingDataFrame('../../data')
traffic = getTrafficDataFrame('../../data')

correlate_tp(parking, traffic)

random_forest_regression(parking, traffic, 2)

plot_parking(parking)
plot_traffic(traffic)
