import("slices")
type MinStack struct {
	min int
	stack []int
	length int
}

func Constructor() MinStack {
	var minstack MinStack
	return minstack
}

func (this *MinStack) Push(val int) {
	if this.length == 0 || val < this.min {
		this.min = val
	}
	this.stack = append(this.stack, val)
	this.length++
}

func (this *MinStack) Pop() {
	if this.length > 0 {
		var tmp int
		if this.length == 1 {
			this.stack = []int{}
		} else {
			tmp = this.stack[this.length-1]
			this.stack = this.stack[:this.length-1]
		}
		this.length--
		if tmp == this.min && this.length > 0 {
			this.min = slices.Min(this.stack)
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[this.length-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
