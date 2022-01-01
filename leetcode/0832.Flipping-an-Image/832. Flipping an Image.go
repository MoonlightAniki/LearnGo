// https://leetcode-cn.com/problems/flipping-an-image/

package solution0832

func flipAndInvertImage(image [][]int) [][]int {
	for _, row := range image {
		for m, n := 0, len(row)-1; m < n; {
			row[m], row[n] = row[n], row[m]
			m++
			n--
		}
	}
	for _, row := range image {
		for j, _ := range row {
			row[j] = 1 - row[j]
		}
	}
	return image
}
