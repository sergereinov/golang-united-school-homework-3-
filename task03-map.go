package homework

import "sort"

type Pair struct {
	int
	string
}

type ByInt []Pair

func (a ByInt) Len() int           { return len(a) }
func (a ByInt) Less(i, j int) bool { return a[i].int < a[j].int }
func (a ByInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func sortMapValues(input map[int]string) (result []string) {

	//get (int,string) pairs
	pairs := make([]Pair, 0, len(input))
	for k, v := range input {
		pairs = append(pairs, Pair{k, v})
	}

	//sort pairs by int
	sort.Sort(ByInt(pairs))

	//copy strings to result
	result = make([]string, 0, len(pairs))
	for _, v := range pairs {
		result = append(result, v.string)
	}
	return
}
