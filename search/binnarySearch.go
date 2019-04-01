package search

//递归实现
func BinarySearch(arr []int,dst int) int{
	mid := len(arr)/2

	//递归最重要的是确定终止条件
	if len(arr) == 0{
		return -1
	}
	//查找成功
	if arr[mid]== dst{
		return mid
	}else if dst < arr[mid]{
		arr = arr[:mid - 1]
	}else {
		arr = arr[mid+1:]
	}
	//继续递归
	return BinarySearch(arr,dst)
}
//非递归实现
func BinarySearchEx(arr []int,dst int) int{

	//递归最重要的是确定终止条件
	if len(arr) == 0{
		return -1
	}

	for len(arr)!=0{
		mid := len(arr)/2
		if dst == arr[mid]{
			return mid
		}else if dst < arr[mid]{
			arr = arr[:mid-1]
		}else{
			arr = arr[mid+1:]
		}

	}
	return -1
}