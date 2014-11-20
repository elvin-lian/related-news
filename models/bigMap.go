package models

var BigMap map[string][]int64

func init() {
	BigMap = make(map[string][]int64)
}

func AppendToBigMap(key string, id int64) {
	BigMap[key] = append(BigMap[key], id)
}

func CountBigMap(keywords []string) (idsMap map[int64]int) {
	idsMap = make(map[int64]int)
	for _, k := range keywords {
		if tmpIds, ok := BigMap[k]; ok {
			for _, id := range tmpIds {
				if _, ok := idsMap[id]; !ok {
					idsMap[id] = 0;
				}
				idsMap[id] ++
			}
		}
	}
	return
}

func CleanBigMap(){
	BigMap = map[string][]int64{}
}

func BigMapLen() int{
	return len(BigMap)
}
