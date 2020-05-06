package heap

type MaxHeap struct {
    Element []int
    count   int
}

func (M *MaxHeap) Push(data int) {
    M.Element = append(M.Element, data)
    M.count += 1
    if M.count > 1 {
        M.shifup(M.count)
    }
}

func (M *MaxHeap) Extract() int {
    if M.count > 0 {
        data := M.Element[0]
        M.Element[0] = M.Element[M.count-1]
        M.Element = M.Element[:M.count-1]
        M.count -= 1
        if M.count > 1 {
            M.shifdown(1)
        }
        return data
    }
}

func (M *MaxHeap) shifup(index int) {
    parent_index := (index >> 1) - 1
    current_index := index - 1
    if parent_index >= 0 && M.Element[parent_index] < M.Element[current_index] {
        M.Element[parent_index], M.Element[current_index] = M.Element[current_index], M.Element[parent_index]
        M.shifup(parent_index + 1)
    }

}

func (M *MaxHeap) shifdown(index int) {
    var swich_index int
    current_index := index - 1
    left_child_index := (index << 1) - 1
    right_child_index := index << 1
    max_index := M.count - 1
    if right_child_index <= max_index {
        if M.Element[left_child_index] > M.Element[right_child_index] {
            swich_index = left_child_index
        } else {
            swich_index = right_child_index
        }
    } else {
        swich_index = left_child_index
    }
    if swich_index <= max_index && M.Element[current_index] < M.Element[swich_index] {
        M.Element[current_index], M.Element[swich_index] = M.Element[swich_index], M.Element[current_index]
        M.shifdown(swich_index)
    }
}
