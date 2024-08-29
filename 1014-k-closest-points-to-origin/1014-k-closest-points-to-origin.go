type PointHeap [][]int

func (h PointHeap) Len() int {
    return len(h)
}

func (h PointHeap) Less(i, j int) bool {
    return distance(h[i]) > distance(h[j])
}

func (h PointHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *PointHeap) Push(x interface{}) {
    *h = append(*h, x.([]int))
}

func (h *PointHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func distance(point []int) int {
    return point[0]*point[0] + point[1]*point[1]
}

func kClosest(points [][]int, k int) [][]int {
    h := &PointHeap{}
    heap.Init(h)

    for _, point := range points {
        heap.Push(h, point)
        if h.Len() > k {
            heap.Pop(h)
        }
    }

    result := make([][]int, k)
    for i := 0; i < k; i++ {
        result[i] = heap.Pop(h).([]int)
    }
    return result
}