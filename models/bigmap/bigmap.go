package bigmap

import "sync"

var BigMap map[string][]int64

var lock *sync.RWMutex

func init() {
	lock = &sync.RWMutex{}
	BigMap = make(map[string][]int64)
}

func AppendToBigMap(key string, id int64) {
	lock.Lock()
	BigMap[key] = append(BigMap[key], id)
	lock.Unlock()
}

func CountBigMap(keywords []string) (idsMap map[int64]int) {
	idsMap = make(map[int64]int)
	for _, k := range keywords {
		if tmpIds, ok := BigMap[k]; ok {
			for _, id := range tmpIds {
				if _, ok := idsMap[id]; !ok {
					idsMap[id] = 0
				}
				idsMap[id]++
			}
		}
	}
	return
}

func CleanBigMap() {
	lock.Lock()
	BigMap = map[string][]int64{}
	lock.Unlock()
}

func BigMapLen() int {
	return len(BigMap)
}
