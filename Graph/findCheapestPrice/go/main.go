package main

import (
	"math"
)

type flightInfo struct {
	stop   int
	price  int
	flight int
}

// BFS without optimization
func FindCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// create adjacency list
	adj := map[int][][]int{}
	for _, f := range flights { // f[0] = from, f[1] = to, f[2] = price
		adj[f[0]] = append(adj[f[0]], []int{f[1], f[2]}) // adj[from] = [to, price]
	}

	q := []flightInfo{{0, 0, src}}
	minPrice := math.MaxInt

	for len(q) > 0 {
		f := q[0]
		q = q[1:]

		// if stop is greater than k, then skip
		if f.stop-1 > k {
			continue
		}

		// if price is greater than minPrice, then skip
		if f.price > minPrice {
			continue
		}

		// if destination is reached, then update minPrice
		if f.flight == dst {
			minPrice = f.price
			continue
		}

		// loop through all the flights(adjacent flights) from the current flight
		for _, a := range adj[f.flight] { // a[0] = to, a[1] = price
			newStops := f.stop + 1
			newPrice := f.price + a[1]
			to := a[0]

			q = append(q, flightInfo{newStops, newPrice, to})
		}
	}

	if minPrice == math.MaxInt {
		return -1
	}

	return minPrice
}

// BFS with optimization(using prices array)
func FindCheapestPrice1(n int, flights [][]int, src int, dst int, k int) int {
	// create adjacency list
	adj := map[int][][]int{}
	for _, f := range flights { // f[0] = from, f[1] = to, f[2] = price
		adj[f[0]] = append(adj[f[0]], []int{f[1], f[2]}) // adj[from] = [to, price]
	}

	// create prices array
	// prices[i] represents the minimum price to reach location i
	prices := make([]int, n)
	for i := range prices {
		prices[i] = math.MaxInt // set all prices to max
	}

	q := []flightInfo{{0, 0, src}}
	minPrice := math.MaxInt

	for len(q) > 0 {
		f := q[0]
		q = q[1:]

		// if stop is greater than k, then skip
		if f.stop-1 > k {
			continue
		}

		// if destination is reached, then update minPrice
		if f.flight == dst {
			minPrice = f.price
			continue
		}

		// loop through all the flights(adjacent flights) from the current flight
		for _, adjFlight := range adj[f.flight] { // a[0] = to, a[1] = price
			newPrice := f.price + adjFlight[1]
			newStops := f.stop + 1
			to := adjFlight[0]

			// note that prices[i] is the minimum price to reach location i
			// if the new price is greater than the existing price
			// then there's no need to explore this path
			if newPrice >= prices[to] {
				continue
			}

			q = append(q, flightInfo{newStops, newPrice, to})

			// update the price to reach location
			prices[to] = newPrice

		}
	}

	if minPrice == math.MaxInt {
		return -1
	}

	return minPrice
}

type flightInfo1 struct {
	price  int
	flight int
}

// BFS with optimization(using prices array)
// BFS with layer by layer traversal
func FindCheapestPrice2(n int, flights [][]int, src int, dst int, k int) int {
	// create adjacency list
	adj := map[int][][]int{}
	for _, f := range flights { // f[0] = from, f[1] = to, f[2] = price
		adj[f[0]] = append(adj[f[0]], []int{f[1], f[2]}) // adj[from] = [to, price]
	}

	// create prices array
	// prices[i] represents the minimum price to reach location i
	prices := make([]int, n)
	for i := range prices {
		prices[i] = math.MaxInt // set all prices to max
	}

	q := []flightInfo1{{0, src}}
	minPrice := math.MaxInt
	stops := -1

	// note that stops <= k is used to ensure that we don't exceed the number of stops
	for len(q) > 0 && stops <= k {
		for _, f := range q {
			q = q[1:]

			// if destination is reached, then update minPrice
			if f.flight == dst {
				minPrice = f.price
				continue
			}

			// loop through all the flights(adjacent flights) from the current flight
			for _, a := range adj[f.flight] { // a[0] = to, a[1] = price
				newPrice := f.price + a[1]
				to := a[0]

				// note that prices[i] is the minimum price to reach location i
				// if the new price is greater than the existing price
				// then there's no need to explore this path
				if newPrice > prices[a[0]] {
					continue
				}

				q = append(q, flightInfo1{newPrice, to})

				// update the price to reach location a[0]
				prices[a[0]] = newPrice
			}
		}

		stops++
	}

	if minPrice == math.MaxInt {
		return -1
	}

	return minPrice
}

// Using Bellman Ford Algorithm
func FindCheapestPrice3(n int, flights [][]int, src int, dst int, k int) int {
	// Initialize the prices array
	prices := make([]int, n)
	for i := 0; i < n; i++ {
		prices[i] = math.MaxInt
	}

	// Mark the source price as 0
	prices[src] = 0

	// Loop through the number of stops
	for i := 0; i <= k; i++ {
		tempPrices := make([]int, n)
		copy(tempPrices, prices)

		// Loop through the number of stops
		// Note that we loop through k + 1 times
		for _, f := range flights {
			// Create a temporary prices array(copy of the original prices array)
			from, to, price := f[0], f[1], f[2]

			// Loop through all the adjacent flights
			if prices[from] == math.MaxInt {
				continue
			}

			// If the price of the source is INT_MAX, then we skip this flight
			// It means the `from` node is not reachable yet
			newPrice := prices[from] + price
			if newPrice > tempPrices[to] {
				continue
			}

			// Update the price of the destination node
			tempPrices[to] = newPrice
		}

		// Update the prices array
		copy(prices, tempPrices)
	}

	if prices[dst] == math.MaxInt {
		return -1
	}

	return prices[dst]
}
