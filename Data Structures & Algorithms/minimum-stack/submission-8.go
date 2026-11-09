import("slices")
type MinStack struct {
	min int
	stack []int
}

func Constructor() MinStack {
	var minstack MinStack
	return minstack
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 || val < this.min {
		this.min = val
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	s_len := len(this.stack)
	if s_len > 0 {
		var tmp int
		if s_len == 1 {
			this.stack = []int{}
		} else {
			tmp = this.stack[s_len-1]
			this.stack = this.stack[:s_len-1]
		}
		if tmp == this.min && len(this.stack) > 0 {
			this.min = slices.Min(this.stack)
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
