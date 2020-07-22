package listops

//IntList is a custom type used in the methods
type IntList []int

type unaryFunc func(x int) int
type predFunc func(n int) bool
type binFunc func(x, y int) int

//Length returns the length of a int list
func (il IntList) Length() int {
	return len(il)
}

//Reverse returns the reverse list from a previews int list
func (il IntList) Reverse() IntList {
	newList := IntList{}
	lenght := len(il)

	for i := lenght; i > 0; i-- {
		newList = append(newList, il[i-1])
	}

	return newList
}

//Append items from a list in the constructor list
func (il IntList) Append(originalList IntList) IntList {
	for _, item := range originalList {
		il = append(il, item)
	}
	return il
}

//Concat items from a list of lists in the constructor list
func (il IntList) Concat(list []IntList) IntList {
	for _, items := range list {
		for _, item := range items {
			il = append(il, item)
		}
	}
	return il
}

//Map returns a new list with an function applied in items from the constructor list
func (il IntList) Map(fn unaryFunc) IntList {
	newList := IntList{}
	for _, item := range il {
		newList = append(newList, fn(item))
	}
	return newList
}

//Filter returns a new list with an function applied in items from the constructor lists where the function applied is true
func (il IntList) Filter(fn predFunc) IntList {
	newList := make(IntList, 0, len(il))
	for _, item := range il {
		if fn(item) {
			newList = append(newList, item)
		}
	}
	return newList
}

//Foldr reduces each item from the right side of a list using an acumulator
func (il IntList) Foldr(fn binFunc, initial int) int {
	for i := len(il) - 1; i >= 0; i-- {
		initial = fn(il[i], initial)
	}
	return initial
}

//Foldl reduces each item from the left side of a list using an acumulator
func (il IntList) Foldl(fn binFunc, initial int) int {
	for _, item := range il {
		initial = fn(initial, item)
	}
	return initial
}
