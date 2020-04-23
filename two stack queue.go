package main

type MyQueue struct {
	StackOne []int
	StackTwo []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{StackOne: []int{}, StackTwo: []int{}}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.StackOne = append(this.StackOne, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {

	stackOneLen := len(this.StackOne)
	stackTwoLen := len(this.StackTwo)

	if stackTwoLen == 0 {
		for i := 0; i < stackOneLen; i++ {
			ele := this.StackOne[stackOneLen-1-i]
			this.StackTwo = append(this.StackTwo, ele)
			this.StackOne = this.StackOne[:stackOneLen-1-i]
		}
	}

	ret := this.StackTwo[len(this.StackTwo)-1]
	this.StackTwo = this.StackTwo[:len(this.StackTwo)-1]
	return ret
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	stackOneLen := len(this.StackOne)
	stackTwoLen := len(this.StackTwo)

	if stackTwoLen == 0 {
		for i := 0; i < stackOneLen; i++ {
			ele := this.StackOne[stackOneLen-1-i]
			this.StackTwo = append(this.StackTwo, ele)
			this.StackOne = this.StackOne[:stackOneLen-1-i]
		}
	}

	ret := this.StackTwo[len(this.StackTwo)-1]
	return ret
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	stackOneLen := len(this.StackOne)
	stackTwoLen := len(this.StackTwo)

	if stackOneLen == 0 && stackTwoLen == 0 {
		return true
	}

	return false
}
