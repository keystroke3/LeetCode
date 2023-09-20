package main

type Node struct {
	Value int
	Freq  int
	Next  *Node
}

func topKFrequent(nums []int, k int) []int {
	ranks := make(map[int]int, len(nums))
	head := &Node{}
	topK := make([]int, k)
	node := head
	var tail *Node
	for _, n := range nums {
		ranks[n]++
		f := ranks[n]
		if tail != nil && tail.Freq > f {
			continue
		}
		for i := 0; i <= k-1; i++ {
			if f > node.Freq {
				node.Value = n
				node.Freq = f
				topK[i] = n
				break
			}
			if node.Next == nil {
				tail := &Node{Value: n, Freq: f}
				node.Next = tail
				break
			}
			node = node.Next
		}
		node = head
	}
	return topK
}
