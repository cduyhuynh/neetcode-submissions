func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]int, 9)
	cols := make([]map[byte]int, 9)
	squares := make([]map[byte]int, 9)
	for i:=0; i<len(board); i++{
		for j:=0; j<len(board[0]); j++{
			cell := board[i][j]
			if cell == '.' {
				continue
			}
			fmt.Println(i,j, cell)

			if len(rows[i]) == 0 {
				rows[i] = map[byte]int{cell: 1}
			} else {
				rows[i][cell]++
				if rows[i][cell] > 1 {
					return false
				}
			}
			
			if len(cols[j]) == 0 {
				cols[j] = map[byte]int{cell: 1}
			} else {
				cols[j][cell]++
				if cols[j][cell] > 1 {
					return false
				}
			}
			square_idx := (i/3)*3+(j/3)
			if len(squares[square_idx]) == 0 {
				squares[square_idx] = map[byte]int{cell: 1}
			} else {
				squares[square_idx][cell]++
				if squares[square_idx][cell] > 1 {
					return false
				}
			}
		}
	}

	return true
}
