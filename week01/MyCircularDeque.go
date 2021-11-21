package week01

type MyCircularDeque struct {
	queue     []int
	headIndex int
	size      int
	capacity  int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		queue:     make([]int, k),
		headIndex: 0,
		size:      0,
		capacity:  k,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.headIndex = (this.headIndex - 1 + this.capacity) % this.capacity
	this.queue[this.headIndex] = value
	this.size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.queue[(this.headIndex+this.size)%this.capacity] = value
	this.size++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.headIndex = (this.headIndex + 1) % this.capacity
	this.size--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.headIndex]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[(this.headIndex+this.size-1)%this.capacity]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.size == this.capacity
}
