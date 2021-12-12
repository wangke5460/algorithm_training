package week04

import "container/heap"

var timeStep int

//消息存储
type ListNode struct {
	timeID int
	id     int
	next   *ListNode
}

//关注人
type set map[int]bool

//维护一个大顶堆 合并k个链表
type maxHeap []*ListNode

type Twitter struct {
	//要储存某个人关注的列表
	//要有一个保存某个人信息的 存储方式 链表和数值都可以的
	//取信息时候要将本人及其关注的信息，合并k个链表或者数组
	followUser map[int]set
	storeMess  map[int]*ListNode
}

func Constructor() Twitter {
	return Twitter{
		followUser: map[int]set{},
		storeMess:  map[int]*ListNode{},
	}
}

//储存用户信息
func (this *Twitter) PostTweet(userId int, tweetId int) {
	timeStep++
	newNode := &ListNode{timeID: timeStep, id: tweetId}
	newNode.next = this.storeMess[userId]
	this.storeMess[userId] = newNode
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	h := new(maxHeap)
	if this.storeMess[userId] != nil {
		heap.Push(h, this.storeMess[userId])
	}
	for user, _ := range this.followUser[userId] {
		if this.storeMess[user] != nil {
			heap.Push(h, this.storeMess[user])
		}
	}

	ans := []int{}
	for h.Len() > 0 && len(ans) < 10 {
		node := heap.Pop(h).(*ListNode)
		ans = append(ans, node.id)
		if node.next == nil {
			continue
		}
		heap.Push(h, node.next)
	}
	return ans
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.followUser[followerId] == nil {
		this.followUser[followerId] = set{}
	}
	this.followUser[followerId][followeeId] = true
	return
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.followUser[followerId], followeeId)
}

func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Less(i, j int) bool {
	return m[i].timeID > m[j].timeID
}

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(*ListNode))
}

func (m *maxHeap) Pop() (v interface{}) {
	v, *m = (*m)[len(*m)-1], (*m)[:len(*m)-1]
	return v
}
