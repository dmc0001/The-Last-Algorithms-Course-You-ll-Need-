package heap

type item struct {
	v        any
	priority int
}

type Heap struct {
	data []item
	idx  int
}

func (h *Heap) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *Heap) Peek() (any, bool) {
	if h.IsEmpty() {
		return nil, false
	}
	return h.data[0].v, true
}

func (h *Heap) Delete() (any, bool) {
	if h.IsEmpty() {
		return nil, false
	}
	removed := h.data[0].v
	last := len(h.data) - 1
	h.data[0] = h.data[last]
	h.data = h.data[:last]
	h.idx = 0
	h.heapfyDown()
	return removed, true
}

func (h *Heap) Insert(v any, priority int) {
	h.data = append(h.data, item{v: v, priority: priority})
	h.idx = len(h.data) - 1
	h.heapfyUp()
}

func (h *Heap) heapfyDown() {
	r := h.righChild()
	l := h.leftChild()
	n := len(h.data)
	smallest := h.idx
	if l < n && h.data[smallest].priority > h.data[l].priority {
		smallest = l
	}
	if r < n && h.data[smallest].priority > h.data[r].priority {
		smallest = r
	}
	//base case
	if smallest == h.idx {
		return
	}
	h.data[h.idx], h.data[smallest] = h.data[smallest], h.data[h.idx]
	h.idx = smallest
	h.heapfyDown()

}

func (h *Heap) heapfyUp() {
	// get parent of the current idx
	p := h.parent()
	// get indx
	curr := h.idx
	// compare them
	if h.data[p].priority <= h.data[curr].priority {
		return
	}
	// do the swap
	h.data[curr], h.data[p] = h.data[p], h.data[curr]
	h.idx = p
	h.heapfyUp()
}

func (h *Heap) parent() int {
	return int((h.idx - 1) / 2)
}

func (h *Heap) righChild() int {
	return h.idx*2 + 2
}

func (h *Heap) leftChild() int {
	return h.idx*2 + 1
}
