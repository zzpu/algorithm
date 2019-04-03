package search

import "fmt"

type ListItem struct{
	Value int
	NextItem *ListItem
}

func Revertm(items *ListItem,m int) *ListItem{

	var dst,head *ListItem

	head = items
	dst  = items
	for i := 1;head.NextItem != nil;i++{
		//间隔m才赋值
		if i >= m{
			dst = dst.NextItem
		}
		//头部往前移动
		head = head.NextItem
	}

	return dst
}

func DemoRevertm(){

	var items,p *ListItem

	items = new(ListItem)

	p = items
	for i:= 1;i<=100;i++{

		p.NextItem = new(ListItem)

		p = p.NextItem
		p.Value = i

	}

	item := Revertm(items,10)

	if item != nil{
		fmt.Printf("item:%d\n",item.Value)
	}else{
		fmt.Printf("not fount")
	}


}
