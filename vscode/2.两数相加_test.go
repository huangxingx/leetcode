package vscode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	t1 := []int{2, 4, 3}
	t2 := []int{5, 6, 4}
	except := []int{7, 0, 8}
	retListNode := addTwoNumbers(makeListNode(t1), makeListNode(t2))
	retList := changeListNode2List(retListNode)
	if !reflect.DeepEqual(retList, except) {
		t.Error("错误")
	}
	fmt.Println(retList)
}
