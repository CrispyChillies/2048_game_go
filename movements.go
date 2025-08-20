package main

// := declare a new variable, = used to assign a value to a declared variable

func reverseRow(row []int) []int {
	newRow := make([]int, Size)
	for i := range Size {
		newRow[i] = row[Size-1-i]
	}
	return newRow
}

func mergeRowLeft(row []int) []int {
	newRow := make([]int, 0, Size)

	// Loop by using index and value in the row list
	for _, v := range row {
		if v != 0 {
			newRow = append(newRow, v)
		}
	}

	for len(newRow) < Size {
		newRow = append(newRow, 0)
	}

	for i := 0; i < Size-1; i++ {
		if newRow[i] != 0 && newRow[i] == newRow[i+1] {
			newRow[i] *= 2
			newRow[i+1] = 0
		}
	}

	// 3. slide again after merging
	finalRow := make([]int, 0, Size)
	for _, v := range newRow {
		if v != 0 {
			finalRow = append(finalRow, v)
		}
	}
	for len(finalRow) < Size {
		finalRow = append(finalRow, 0)
	}
	return finalRow
}
