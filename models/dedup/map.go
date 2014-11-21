package dedup

// 去重
var ContMap  map[uint16][]int64
var TitleMap map[uint16][]int64
var NewsMap map[int64][]uint16

func init() {
	ContMap = make(map[uint16][]int64)
	TitleMap = make(map[uint16][]int64)
	NewsMap = make(map[int64][]uint16)
}

func AppendToNewsMap(news * News) {
	if news.Sh1 > 0 {
		ContMap[news.Sh1] = append(ContMap[news.Sh1], news.Id)
		ContMap[news.Sh2] = append(ContMap[news.Sh2], news.Id)
		ContMap[news.Sh3] = append(ContMap[news.Sh3], news.Id)
		ContMap[news.Sh4] = append(ContMap[news.Sh4], news.Id)
	}

	if news.ShT1 > 0 {
		TitleMap[news.ShT1] = append(TitleMap[news.ShT1], news.Id)
		TitleMap[news.ShT2] = append(TitleMap[news.ShT2], news.Id)
		TitleMap[news.ShT3] = append(TitleMap[news.ShT3], news.Id)
		TitleMap[news.ShT4] = append(TitleMap[news.ShT4], news.Id)
	}

	NewsMap[news.Id] = []uint16{news.Sh1, news.Sh2, news.Sh3, news.Sh4, news.ShT1, news.ShT2, news.ShT3, news.ShT4}
}

func CleanMap() {
	ContMap = map[uint16][]int64{}
	TitleMap = map[uint16][]int64{}
	NewsMap = map[int64][]uint16{}
}

func ContMapLen() int {
	return len(ContMap)
}

func TitleMapLen() int {
	return len(TitleMap)
}

func NewsMapLen() int {
	return len(NewsMap)
}
