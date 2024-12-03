package main

func AsteroidCollision(asteroids []int) []int {
	stack := []int{}

	for _, asteroid := range asteroids {
		// Keep track if current asteroid survives
		// This is used to check if we need to add the current asteroid to the stack in the end
		survive := true

		// While we have a negative asteroid and stack has positive asteroids
		for len(stack) > 0 && asteroid < 0 && stack[len(stack)-1] > 0 {
			// If equal size, both are destroyed
			if stack[len(stack)-1] == -asteroid {
				stack = stack[:len(stack)-1]
				survive = false
				break
			}

			// If top asteroid is bigger, current asteroid is destroyed
			if stack[len(stack)-1] > -asteroid {
				survive = false
				break
			}

			// Top asteroid is smaller, remove it and continue checking
			stack = stack[:len(stack)-1]
		}

		if survive {
			stack = append(stack, asteroid)
		}
	}

	return stack
}
