package easy

import (
	. "dsa/structs"
	"fmt"
)

func LevelOrder(root *TreeNode) [][]int {
	var depth int
	levels := make([][]int, 0)

	var order func(root *TreeNode)
	order = func(root *TreeNode) {
		if root == nil {
			return
		}
		levels = doesSliceLevelExist(levels, depth)
		levels[depth] = append(levels[depth], root.Val)

		depth++
		order(root.Left)
		order(root.Right)
		depth--
	}

	order(root)
	fmt.Println(levels)

	return levels
}

func doesSliceLevelExist(levelsSlice [][]int, depth int) [][]int {

	if len(levelsSlice) <= depth {
		newlevel := make([]int, 0)
		levelsSlice = append(levelsSlice, newlevel)

	}
	return levelsSlice

}

func doesMapLevelExist(levelsMap map[int][]int, level int) {
	if _, ok := levelsMap[level]; !ok {
		//levelsMap[level] = make([]int, 0, int(math.Pow(2.0, float64(level))))
		levelsMap[level] = make([]int, 0)
	}
}
