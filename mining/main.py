from loading.parking import getParkingDataFrame
from loading.traffic import getTrafficDataFrame
from processing.parkingtraffic import correlate_tp
from processing.plot_parking import plot_parking
from processing.plot_traffic import plot_traffic

parking = getParkingDataFrame()
traffic = getTrafficDataFrame()

correlate_tp(parking, traffic)

plot_parking(parking)
plot_traffic(traffic)
