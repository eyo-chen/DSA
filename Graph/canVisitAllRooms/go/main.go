package main

// using DFS(recursion)
func CanVisitAllRooms(rooms [][]int) bool {
	isRoomVisited := map[int]bool{
		0: true,
	}
	helper(0, rooms, isRoomVisited)

	return len(isRoomVisited) == len(rooms)
}

func helper(key int, rooms [][]int, isRoomVisited map[int]bool) {
	for _, curRoom := range rooms[key] {
		// check if the room is already visited
		if _, ok := isRoomVisited[curRoom]; ok {
			continue
		}

		isRoomVisited[curRoom] = true
		helper(curRoom, rooms, isRoomVisited)
	}
}

// using BFS
func CanVisitAllRooms1(rooms [][]int) bool {
	queue := []int{0}
	hashTable := map[int]bool{}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if hashTable[n] {
			continue
		}

		hashTable[n] = true

		// Directly put all the keys in the current room into the queue
		// Or
		// We can manually check each key in the current room
		// and put the key into the queue if the room is not visited
		queue = append(queue, rooms[n]...)
	}

	return len(rooms) == len(hashTable)
}
