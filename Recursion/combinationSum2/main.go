package main

import "sort"

// Using Hast Table and unique candidates list
func CombinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	frequencyMap := make(map[int]int)
	uniqueCandidates := []int{}

	// Count frequencies and build unique candidates list
	for _, c := range candidates {
		if frequencyMap[c] == 0 {
			uniqueCandidates = append(uniqueCandidates, c)
		}
		frequencyMap[c]++
	}

	backtrack(uniqueCandidates, target, 0, []int{}, &result, frequencyMap)
	return result
}

func backtrack(candidates []int, target, index int, current []int, result *[][]int, frequencyMap map[int]int) {
	if target == 0 {
		*result = append(*result, append([]int(nil), current...))
		return
	}

	if target < 0 {
		return
	}

	for i := index; i < len(candidates); i++ {
		candidate := candidates[i]

		if frequencyMap[candidate] == 0 {
			continue
		}

		frequencyMap[candidate]--
		current = append(current, candidate)
		backtrack(candidates, target-candidate, i, current, result, frequencyMap)
		current = current[:len(current)-1]
		frequencyMap[candidate]++
	}
}

// Using Sorting and skipping duplicates
func CombinationSum22(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	ans := [][]int{}
	helper2(candidates, target, 0, []int{}, &ans)
	return ans
}

func helper2(candidates []int, target int, index int, cur []int, ans *[][]int) {
	if target < 0 {
		return
	}

	if target == 0 {
		t := make([]int, len(cur))
		copy(t, cur)
		*ans = append(*ans, t)
		return
	}

	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}

		cur = append(cur, candidates[i])
		helper2(candidates, target-candidates[i], i+1, cur, ans)
		cur = cur[:len(cur)-1]
	}
}
