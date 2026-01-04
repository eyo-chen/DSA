package main

import "sort"

// CarFleet calculates the number of car fleets that will arrive at the target destination.
//
// Approach:
// 1. Calculate the time each car needs to reach the target
// 2. Sort cars by their starting position in descending order (closest to target first)
// 3. Iterate through sorted cars from front to back:
//   - If a car takes longer to reach target than the car ahead, it forms a new fleet
//   - Otherwise, it catches up and joins the fleet ahead
//
// Time Complexity: O(n log n) due to sorting
// Space Complexity: O(n) for storing car data
func CarFleet(target int, position []int, speed []int) int {
	n := len(position)

	// Store each car's position and time to reach target
	type carInfo struct {
		startPosition int
		arrivalTime   float64
	}

	cars := make([]carInfo, n)
	for i := range n {
		// Calculate time to reach target: distance / speed
		distanceToTarget := target - position[i]
		timeToTarget := float64(distanceToTarget) / float64(speed[i])

		cars[i] = carInfo{
			startPosition: position[i],
			arrivalTime:   timeToTarget,
		}
	}

	// Sort cars by starting position in descending order
	// This allows us to process cars from closest to target to farthest
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].startPosition > cars[j].startPosition
	})

	fleetCount := 0
	slowestArrivalTime := -1.0

	// Process cars from front (closest to target) to back
	for _, currentCar := range cars {
		// If current car takes longer than the fleet ahead, it can't catch up
		// so it forms a new fleet
		if currentCar.arrivalTime > slowestArrivalTime {
			fleetCount++
			// This car becomes the slowest (limiting factor) for cars behind it
			slowestArrivalTime = currentCar.arrivalTime
		}
		// If current car arrives earlier or at same time, it catches up
		// and joins the fleet ahead (no increment to fleetCount)
	}

	return fleetCount
}
