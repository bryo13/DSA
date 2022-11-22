// bubble sort compares two neighbouring elements and swaps them if they are
// larger of smaller depending on the mode of sorting, i.e acending or descending
// has a O(n*n) - uses two  for loops

// isSwapped is a variable that stops the comparing if its sorted
package bubble

func SortArray(a [5]int16) [5]int16 {
	isSwapped := true
	len := len(a)

	for isSwapped {
		isSwapped = false

		for i := 1; i < len; i++ {
			temp := a[i]

			if a[i] < a[i-1] {
				a[i] = a[i-1]
				a[i-1] = temp
				isSwapped = true
			}
		}
	}
	return a
}
