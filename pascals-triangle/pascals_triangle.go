package pascal

func Triangle(n int) [][]int {
	var a = [][]int{{1}}
	if n > 1 {
		for i := 0; i < n-1; i++ {
			a = append(a, CreateNextRow(a[i]))
		}
	}
	return a
}

func CreateNextRow(prev_row []int) []int {
	var next_row []int
	if len(prev_row) == 1 {
		next_row = append(next_row, 1, 1)
	} else {
		next_row = append(next_row, 1)
		for i := 0; i < len(prev_row)-1; i++ {
			next_row = append(next_row, prev_row[i] + prev_row[i+1])
		}
		next_row = append(next_row, 1)
	}
	return next_row
}
