# widgetsensor
widget sensor log processor

Accepts from `stdin` the following sensor log data example:

```text
reference 70.0 45.0
thermometer temp-1
2007-04-05T22:00 temp-1 72.4
2007-04-05T22:01 temp-1 76.0
thermometer temp-2
2007-04-05T22:01 temp-2 69.5
2007-04-05T22:02 temp-2 70.1
humidity hum-1
2007-04-05T22:04 hum-1 45.2
2007-04-05T22:05 hum-1 45.3
humidity hum-2
2007-04-05T22:04 hum-2 44.4
2007-04-05T22:05 hum-2 43.9
```

Will produce an output for each sensor example:
```text
temp-1: precise
temp-2: ultra precise
hum-1: OK
hum-2: discard
```

For a thermometer, it is branded “ultra precise” if the mean of the readings is within 0.5 degrees of the known temperature, and the standard deviation is less than 3.
It is branded “very precise” if the mean is within 0.5 degrees of the room, and the standard deviation is under 5.
Otherwise, it’s sold as “precise”.

For a humidity sensor it is discarded unless it is within 1% of the reference value for all readings.

# testing

`make test`

# running

`cat sensor.dat | go run cmd/main/main.go`
