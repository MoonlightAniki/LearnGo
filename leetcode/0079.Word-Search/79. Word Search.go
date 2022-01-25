package solution0079

import "fmt"

/*
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.



Example 1:
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true

Example 2:
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
Output: true

Example 3:
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
Output: false


Constraints:
m == board.length
n = board[i].length
1 <= m, n <= 6
1 <= word.length <= 15
board and word consists of only lowercase and uppercase English letters.


Follow up: Could you use search pruning to make your solution faster with a larger board?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

var m, n int
var direction = [][]int{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}
var visited [][]bool

func inArea(x, y int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}

// 从 board[startX][startY] 开始查找字符串 word[index...len(word))
func searchWord(board [][]byte, startX int, startY int, word []byte, index int) bool {
	if index == len(word)-1 {
		return board[startX][startY] == word[index]
	}
	if board[startX][startY] == word[index] {
		visited[startX][startY] = true
		for _, d := range direction {
			newX, newY := startX+d[0], startY+d[1]
			if inArea(newX, newY) && !visited[newX][newY] && searchWord(board, newX, newY, word, index+1) {
				return true
			}
		}
		visited[startX][startY] = false
	}
	return false
}

func exist(board [][]byte, word string) bool {
	m = len(board)
	n = len(board[0])
	visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if searchWord(board, i, j, []byte(word), 0) {
				return true
			}
		}
	}
	return false
}

func Test() {
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED"))
}
