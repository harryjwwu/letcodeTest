package main
import (
"fmt"
	bd "project/algorithm/bytedance"

tc "project/algorithm/tencent"
)

func main() {
	fmt.Printf("result:%v", bd.CardMove([]int32{1,2,3,4,5,6,7,8,9,10,11,12,13}))
	fmt.Printf("index:%v", tc.TotalSum([]int32{3, 1, 2 ,6}, 9))
}
