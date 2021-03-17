package main

import "fmt"

type node struct {
	data interface{}
	next *node
}

func (n *node) String() string {
	s := ""
	for n.next != nil {
		s += fmt.Sprintf("%v -> ", n.data)
		n = n.next
	}
	return s
}

var list1 = &node{1, &node{3, &node{5, &node{7, nil}}}}
var list2 = &node{2, &node{4, &node{6, &node{8, nil}}}}
var list3 = &node{0, &node{9, &node{10, &node{11, nil}}}}

// var list4 = &node{2, &node{3, &node{6, &node{7, nil}}}}

func merge(listOfList ...*node) *node {

	output := listOfList[0]
	for i, list := range listOfList[1:] {
		i++
		if list.next == nil {
			break
		}
		if output.data.(int) >= listOfList[i].data.(int) {
			listOfList[i] = list.next
			list.next = output
			output = list
		} else {
			for output.next != nil {

				if output.next.data.(int) >= list.data.(int) {
					listOfList[i] = list.next
					list.next = output.next
					output.next = list
					break
				}

				output = output.next

				if output.next == nil {
					listOfList[i] = list.next
					list.next = nil
					output.next = list
					output.next.next = nil
					break
				}
			}
		}
	}
	return output
}

func main() {
	merged := merge(list1, list2, list3)
	fmt.Printf("%s\n", merged)
}
