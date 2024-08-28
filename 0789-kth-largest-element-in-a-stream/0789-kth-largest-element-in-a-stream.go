// KthLargest struct to hold the min-heap and the value of k.
type KthLargest struct {
	k    int
	heap IntHeap
}

// Constructor initializes the KthLargest object with the integer k and the stream of test scores nums.
func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{
		k:    k,
		heap: IntHeap{},
	}
	heap.Init(&kl.heap)

	for _, num := range nums {
		kl.Add(num)
	}

	return kl
}

// Add adds a new test score val to the stream and returns the element representing the kth largest element in the pool of test scores so far.
func (this *KthLargest) Add(val int) int {
	if this.heap.Len() < this.k {
		heap.Push(&this.heap, val)
	} else if val > this.heap[0] {
		heap.Pop(&this.heap)
		heap.Push(&this.heap, val)
	}

	return this.heap[0]
}

// IntHeap is a min-heap of integers.
type IntHeap []int

func (h IntHeap) Len() int           { 
    return len(h)
}

func (h IntHeap) Less(i, j int) bool {
    return h[i] < h[j]
}

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