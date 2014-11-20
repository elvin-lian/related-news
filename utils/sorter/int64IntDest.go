package sorter

type item struct {
	Key int64
	Val int
}

type Int64IntDescItem []item

func Int64IntDesc(m map[int64]int) Int64IntDescItem{
	ms := make(Int64IntDescItem, 0, len(m))

	for k, v := range m {
		ms = append(ms, item{k, v})
	}

	return ms
}

func (ms Int64IntDescItem) Len() int {
	return len(ms)
}

func (ms Int64IntDescItem) Less(i, j int) bool {
	return ms[i].Val > ms[j].Val // 按值大到小排序
}

func (ms Int64IntDescItem) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}
