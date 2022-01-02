package lesson1

//node: 链表节点
/*field
value: 节点值
prev: 前驱节点
next: 后驱节点
*/
type node struct {
	value int
	prev  *node
	next  *node
}

//MyCircularDeque: 循环队列
/*field
cap: 队列容量
count: 队列大小
head: 队列头部节点
tail: 队列尾端节点
*/
type MyCircularDeque struct {
	cap   int
	count int
	head  *node
	tail  *node
}

//Constructor: 构造循环队列
/*parameter
k: 队列容量
*/
func Constructor(k int) MyCircularDeque {

	head := &node{
		value: -100,
	}

	tail := &node{
		value: -100,
	}

	head.next = tail
	tail.prev = head

	d := MyCircularDeque{
		cap:   k,
		count: 0,
		head:  head,
		tail:  tail,
	}

	return d
}

//InsertFront: 插入新的头部节点
/*parameter
value: 插入值
return: 插入成功为true, 否则false
*/
func (this *MyCircularDeque) InsertFront(value int) bool {

	//已满不能插入
	if this.IsFull() {

		return false
	}

	newNode := &node{
		value: value,
		prev:  this.head,
		next:  this.head.next,
	}

	this.head.next.prev = newNode
	this.head.next = newNode
	this.count++

	return true
}

//InsertLast: 插入新的尾端节点
/*parameter
value: 插入值
return: 插入成功为true, 否则false
*/
func (this *MyCircularDeque) InsertLast(value int) bool {

	if this.IsFull() {

		return false
	}

	newNode := &node{
		value: value,
		prev:  this.tail.prev,
		next:  this.tail,
	}

	this.tail.prev.next = newNode
	this.tail.prev = newNode
	this.count++

	return true
}

//DeleteFront：删除头部节点
/*parameter
return:  删除成功为true, 否则false
*/
func (this *MyCircularDeque) DeleteFront() bool {

	//空列表不能删除
	if this.IsEmpty() {

		return false
	}

	this.head.next.next.prev = this.head
	this.head.next = this.head.next.next
	this.count--

	return true
}

//DeleteLast: 删除尾端节点
/*parameter
return:  删除成功为true, 否则false
*/
func (this *MyCircularDeque) DeleteLast() bool {

	if this.IsEmpty() {

		return false
	}

	this.tail.prev.prev.next = this.tail
	this.tail.prev = this.tail.prev.prev
	this.count--

	return true
}

//GetFront: 获取头部节点值
/*parameter
return: 头部节点的值
*/
func (this *MyCircularDeque) GetFront() int {

	if this.IsEmpty() {

		return -1
	}

	return this.head.next.value
}

//GetRear: 获取尾端节点值
/*parameter
return: 尾端节点的值
*/
func (this *MyCircularDeque) GetRear() int {

	if this.IsEmpty() {

		return -1
	}

	return this.tail.prev.value
}

//IsEmpty: 队列是否为空？
/*parameter
return: 空队列为true，否则false
*/
func (this *MyCircularDeque) IsEmpty() bool {

	//队列节点数为0则判断为空队列
	return this.count <= 0
}

//IsFull: 队列是否已满？
/*parameter
return: 满队列为true，否则false
*/
func (this *MyCircularDeque) IsFull() bool {

	//队列节点数大于等于队列容量为已满
	return this.count >= this.cap
}
