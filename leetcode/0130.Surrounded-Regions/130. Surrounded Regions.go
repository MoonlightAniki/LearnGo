package solution0130

/*
Given an m x n matrix board containing 'X' and 'O', capture all regions that are 4-directionally surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.



Example 1:
Input: board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
Output: [["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
Explanation: Surrounded regions should not be on the border, which means that any 'O' on the border of the board are not flipped to 'X'. Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'. Two cells are connected if they are adjacent cells connected horizontally or vertically.


Example 2:
Input: board = [["X"]]
Output: [["X"]]


Constraints:
m == board.length
n == board[i].length
1 <= m, n <= 200
board[i][j] is 'X' or 'O'.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/surrounded-regions
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
var m = 0
var n = 0

var directions = [][2]int{
	{0, -1}, //Up
	{1, 0},  //Right
	{0, 1},  //Down
	{-1, 0}, //Left
}

func inArea(x, y int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}

func bfs(board [][]byte, visited [][]bool, record *[][2]int, x int, y int) bool {
	queue := [][2]int{{x, y}}
	visited[x][y] = true
	ret := true
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		*record = append(*record, front)
		for _, d := range directions {
			newX := front[0] + d[0]
			newY := front[1] + d[1]
			if !inArea(newX, newY) {
				ret = false
			} else if !visited[newX][newY] && board[newX][newY] == 'O' {
				queue = append(queue, [2]int{newX, newY})
				visited[newX][newY] = true
			}
		}
	}
	return ret
}

func solve(board [][]byte) {
	m = len(board)
	if m == 0 {
		return
	}
	n = len(board[0])
	if n == 0 {
		return
	}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !visited[i][j] {
				record := make([][2]int, 0)
				if bfs(board, visited, &record, i, j) {
					for _, r := range record {
						board[r[0]][r[1]] = 'X'
					}
				}
			}
		}
	}
}

// 执行用时：20 ms , 在所有 Go 提交中击败了 62.33% 的用户
