# Asteroid problem solve

This repo is for solve the a asteroid problem that you can find [here](https://www.linkedin.com/posts/globant_an-asteroid-is-heading-towards-the-earth-activity-6623303226718507008-x0vs/)

This simple program loop around all posible values of the digits and check when they solve the problem, the output is:
```
Eureka: 2 7 4 0 5 8 1 6 9
EARTH: 7 2 5 1 4

Eureka: 6 9 4 8 5 2 3 1 0
EARTH: 9 6 5 3 4
```

This takes almost 6 seconds to check all scenarios
```bash
$ time ./earth
Eureka: 2 7 4 0 5 8 1 6 9
EARTH: 7 2 5 1 4

Eureka: 6 9 4 8 5 2 3 1 0
EARTH: 9 6 5 3 4

./earth  5.56s user 0.05s system 99% cpu 5.609 total
```
