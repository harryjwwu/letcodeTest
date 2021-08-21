package bytedance
// 基础 一星
/*
     * 1. 从牌顶拿出一张牌，放到桌子上
     * 2. 从牌顶拿出一张牌，放在牌的底部
     * 3. 重复第一步，第二步操作，知道所有的牌都放到了桌子上
     *
     * 问：已知桌子上牌的顺序是1,2,3,4,5,6,7,8,9,10,11,12,13
     * 牌原来的顺序是什么
 */
func CardMove(location []int32) []int32 {
	if len(location) == 0 {
		return nil
	}
	// 倒过来
	// 从牌底拿一张牌，放到牌顶
	// 从桌子最后的牌，拿一张牌，放到牌顶
	// 直到桌子上没有牌为止
	var result []int32
	for i := len(location)-1; i >= 0; i -- {
		if len(result) > 0 {
			result = append([]int32{result[len(result)-1]}, result...) //将尾部元素添加到顶部
			result = result[:len(result)-1] //删除最后一个元素
		}
		result = append([]int32{location[i]}, result...) //将桌子上的牌放到顶部
	}
	return result
}

