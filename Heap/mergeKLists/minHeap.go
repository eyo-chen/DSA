package main

type MinHeap struct {
	vals []int
}

func NewMinHeap(size int) *MinHeap {
	return &MinHeap{
		vals: make([]int, 0, size),
	}
}

func (m *MinHeap) Len() int {
	return len(m.vals)
}

// The steps to insert a new value into the min heap are as follows:
// 1. Add the value to the end of the heap
// 2. Bubble up the value until the inserted value is less than its parent
func (m *MinHeap) Insert(val int) {
	m.vals = append(m.vals, val)

	m.siftUp(len(m.vals) - 1)
}

// The steps to pull the root of the min heap are as follows:
// 1. Cache the root value for return
// 2. Move the last element to the root
// 3. Remove the last element
// 4. Sift down the new root until the new root value is less than both children
func (m *MinHeap) Pull() int {
	// cache the result for return
	res := m.vals[0]

	last := len(m.vals) - 1

	// move the last element to the root
	m.vals[0] = m.vals[last]

	// remove the last element
	// note that [low:high] is a slice of the original array
	// low is inclusive, high is exclusive
	// we can think of [:n] returns the new slice with only the first n elements
	m.vals = m.vals[:last]

	// sift down the new root to maintain the heap property
	if len(m.vals) > 0 {
		m.siftDown(0)
	}

	return res
}

// The steps to sift up the new value are as follows:
// 1. Compare the current value with its parent
// 2. If the current value is NOT less than its parent, it means parent is less than child, so we can stop sifting up
// 3. Otherwise, swap the current value with the parent value, and update the current index to the parent index, and repeat the process
func (m *MinHeap) siftUp(i int) {
	for i > 0 {
		parent := m.getParentIndex(i)

		// stop bubble up when parent is less than child
		if m.vals[parent] <= m.vals[i] {
			break
		}

		// swap parent and current idx, also update idx
		m.swap(parent, i)
		i = parent
	}
}

// The steps to sift down the new root are as follows:
// 1. Find the index of the smallest child
// 2. Compare the current index with the smallest child index
// 3. If the current index is EQUAL to the smallest child index, it means the current value is less than both children, so we can stop sifting down
// 4. Otherwise, swap the current index with the smallest child index, and update the current index to the smallest child index, and repeat the process
func (m *MinHeap) siftDown(i int) {
	for {
		minIndex := i
		left, right := 2*i+1, 2*i+2

		// find the index of the smallest child
		// before comparing, we need to make sure the left and right indices are within the bounds of the array
		// in other words, we have to make sure the left and right children exist
		if left < len(m.vals) && m.vals[left] < m.vals[minIndex] {
			minIndex = left
		}
		if right < len(m.vals) && m.vals[right] < m.vals[minIndex] {
			minIndex = right
		}

		// stop sifting down when current value is less than both children
		if minIndex == i {
			break
		}

		m.swap(i, minIndex)
		i = minIndex
	}
}

func (m *MinHeap) swap(i, j int) {
	m.vals[i], m.vals[j] = m.vals[j], m.vals[i]
}

func (m *MinHeap) getParentIndex(i int) int {
	return (i - 1) / 2
}
