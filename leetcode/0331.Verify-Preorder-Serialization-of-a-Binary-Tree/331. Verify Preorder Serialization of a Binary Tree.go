package solution0331

// 执行用时：8 ms, 在所有 Go 提交中击败了 7.35% 的用户
/*func isValidSerialization(preorder string) bool {
	s := preorder
	prevLen := len(s)
	// number,#,# 表示叶子节点，使用正则表达式匹配叶子节点，替换成空节点#, 重复此操作，知道最终只剩一个空节点#
	regex := regexp.MustCompile(`\d+,#,#`)
	for {
		s = regex.ReplaceAllString(s, "#")
		if s == "#" {
			return true
		}
		if len(s) == prevLen {
			return false
		}
		prevLen = len(s)
	}
}*/

// 执行用时： 0 ms, 在所有 Go 提交中击败了 100.00% 的用户
func isValidSerialization(preorder string) bool {
	count := 0 // # 数量
	for i := len(preorder) - 1; i >= 0; i-- {
		if preorder[i] == ',' {
			continue
		} else if preorder[i] == '#' {
			count++
		} else {
			for i >= 0 && preorder[i] != ',' {
				i--
			}
			if count >= 2 {
				// \d+,#,# 叶子替换成成一个空节点#
				count--
			} else {
				return false
			}
		}
	}
	// 最终只剩下一个#则是合法的序列化字符串
	return count == 1
}
