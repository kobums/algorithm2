// IntHeap is a max-heap of integers.
type IntHeap []int

func (h IntHeap) Len() int            {
    return len(h)
}

func (h IntHeap) Less(i, j int) bool  {
    return h[i] > h[j]
} // Reverse order to make it a max-heap

func (h IntHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	h := &IntHeap{}
	heap.Init(h)
	
	// Add all stones to the heap
	for _, stone := range stones {
		heap.Push(h, stone)
	}

	// Smash stones until only one or zero stones are left
	for h.Len() > 1 {
		// Remove the two largest stones
		stone1 := heap.Pop(h).(int)
		stone2 := heap.Pop(h).(int)

		if stone1 != stone2 {
			// If they are not equal, push the difference back into the heap
			heap.Push(h, stone1-stone2)
		}
	}

	// If there's no stone left, return 0. Otherwise, return the last stone.
	if h.Len() == 0 {
		return 0
	}
	return heap.Pop(h).(int)
}