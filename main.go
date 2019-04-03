package main
import(
	"algorithm/search"
	"fmt"
)

func main(){
	arr := []int{0,1,2,3,4,5,6,7,8,9}
	index := search.BinarySearchEx(arr,5)

	fmt.Printf("Index:%d\n",index)

	index = search.BinarySearchEx(arr,10)

	fmt.Printf("Index:%d\n",index)

	search.DemoRevertm()
}

