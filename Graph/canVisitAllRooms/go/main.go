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
