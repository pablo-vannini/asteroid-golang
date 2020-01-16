package main

import "fmt"

func main() {
	count := 10
	for a := 0; a < count; a++ {
		for e := 0; e < count; e++ {
			for h := 0; h < count; h++ {
				for n := 0; n < count; n++ {
					for r := 0; r < count; r++ {
						for s := 0; s < count; s++ {
							for t := 0; t < count; t++ {
								for u := 0; u < count; u++ {
									for v := 0; v < count; v++ {
										if result := test(a, e, h, n, r, s, t, u, v); result {
											fmt.Printf("Eureka: %d %d %d %d %d %d %d %d %d\n", a, e, h, n, r, s, t, u, v)
											fmt.Printf("EARTH: %d %d %d %d %d\n\n", e, a, r, t, h)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func test(a int, e int, h int, n int, r int, s int, t int, u int, v int) bool {
	// Earth value
	earth := e*10000 + a*1000 + r*100 + t*10 + h
	venus := v*10000 + e*1000 + n*100 + u*10 + s
	uranus := u*100000 + r*10000 + a*1000 + n*100 + u*10 + s
	saturn := s*100000 + a*10000 + t*1000 + u*100 + r*10 + n
	return (earth+venus+uranus == saturn) && areDiferent(a, e, h, n, r, s, t, u, v)
}

func areDiferent(a int, e int, h int, n int, r int, s int, t int, u int, v int) bool {
	set := make(map[int]int)
	set[a] = a
	set[e] = e
	set[h] = h
	set[n] = n
	set[r] = r
	set[s] = s
	set[t] = t
	set[u] = u
	set[v] = v
	return len(set) == 9
}
