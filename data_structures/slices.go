package data_structures

type mapOperation func(int32) int32
type filterOperation func(int32) bool

func mapInts(op mapOperation, vals []int32) []int32 {
	result := make([]int32, 0, 0)
	for _, item := range vals {
		result = append(result, op(item))
	}
	return result
}

func filterInts(op filterOperation, vals []int32) []int32 {
	result := make([]int32, 0, 0)
	for _, item := range vals {
		if op(item) {
			result = append(result, item)
		}
	}
	return result
}

func concatenate(dest []string, newValues ...string) []string {
	return nil
}

func equals(list1 []string, list2 []string) bool {
	return false
}

func partialReverse(src []int, from, to int) []int {
	return nil
}
