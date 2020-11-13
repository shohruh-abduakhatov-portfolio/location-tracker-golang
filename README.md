# Location Tracker Go

Location Tracker repo for storing current locations of drivers.



# mongo storage
1. better arr ``` [lon,lat] ``` or in GeoJSON format e.g ``` { "type": "Point", "coordinates": [lon,lat] } ```

2. the cardinality is not too big (remember MongoDB document max size is 16MB). If you have 100K points per user split the collections

3. [with location history](https://www.confluent.io/blog/fleet-management-gps-tracking-with-confluent-cloud-mongodb/)

4. save result as ``` coordinates: [<longitude>, <latitude> ] ```

5. 12



### Links
[interface name](https://play.golang.org/p/Lv6-qqqQsH)
[calling funciton name](https://play.golang.org/)