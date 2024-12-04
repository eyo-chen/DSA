package main

type RecentCounter struct {
	Request []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		Request: []int{},
	}
}

// Brute Force
func (rc *RecentCounter) Ping(t int) int {
	// Add the request to the queue
	rc.Request = append(rc.Request, t)

	// Calculate the range of requests within the last 3000ms
	// The range is [t-3000, t]
	min := t - 3000
	max := t
	ans := 0

	// Iterate through the requests and count how many fall within the range
	for _, r := range rc.Request {
		// If the request falls within the range, increment the answer
		if r >= min && r <= max {
			ans++
		}
	}

	return ans
}

// Using Queue Behavior
func (rc *RecentCounter) Ping2(t int) int {
	// Add the request to the queue
	rc.Request = append(rc.Request, t)

	// Calculate the range of requests within the last 3000ms
	min := t - 3000

	// Remove requests that are outside the range
	// Keep removing the first element if there's still requests in the queue
	// AND the first element is outside the range
	for len(rc.Request) > 0 && rc.Request[0] < min {
		rc.Request = rc.Request[1:]
	}

	// Return the number of requests within the range
	return len(rc.Request)
}
