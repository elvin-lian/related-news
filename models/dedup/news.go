package dedup

import (
	"errors"
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/astaxie/beego"

	. "github.com/elvin-lian/related-news/models"
	"github.com/elvin-lian/related-news/utils/mongo"
)

func init() {
	ok, err := beego.AppConfig.Bool("autoInitNews")
	if err == nil && ok {
		AnalyzeNews()
	}
}

/**
 * 初始化最近days天的数据到BigMap里
 */
func AnalyzeNews() {
	beego.Debug("准备去重相关数据")

	CleanMap()

	days := 7
	daysConf, err := beego.AppConfig.Int("dedupMaxDays")
	if err == nil {
		days = daysConf
	}

	mongodb := mongo.Collection("news")

	limit := 5000
	maxLoop, err := beego.AppConfig.Int("dedupMaxLoop")
	if err != nil {
		maxLoop = 100
	}

	now := time.Now()
	beginOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	beginAt := beginOfToday.AddDate(0, 0, -1*days) // days ago

	beego.Informational("beginAt: ", beginAt)

	var lastId int64 = 0
	var allNews []News
	var query *mgo.Query

	for i := 0; i < maxLoop; i++ {
		allNews = []News{}

		if lastId == 0 {
			query = mongodb.Find(bson.M{"created_at": bson.M{"$gte": beginAt}, "status": 1, "cate_id": bson.M{"$ne": 9}})
		} else {
			query = mongodb.Find(bson.M{"created_at": bson.M{"$gte": beginAt}, "status": 1, "id": bson.M{"$lt": lastId}, "cate_id": bson.M{"$ne": 9}})
		}

		err := query.Select(bson.M{"id": 1, "sh1": 1, "sh2": 1, "sh3": 1, "sh4": 1, "sh_t1": 1, "sh_t2": 1, "sh_t3": 1, "sh_t4": 1}).Limit(limit).Sort("-id").All(&allNews)
		if err != nil {
			beego.Error("query news error: ", err.Error())
		} else {
			if len(allNews) == 0 {
				break
			}

			for _, news := range allNews {
				AppendToNewsMap(&news)
				lastId = news.Id
			}
		}
	}
}

// 当有重复时，返回true
func Check(sh *[8]uint16) int8 {
	if sh[0] == 0 {
		return -1
	}

	ok := checkContent(sh)
	if !ok && sh[4] != 0 {
		ok = checkTitle(sh)
	}

	if ok {
		return 1
	} else {
		return 0
	}
}

// 如果是在指定的距离内，返回true (表示有重复)
func checkContent(sh *[8]uint16) (resp bool) {
	maxDist, err := beego.AppConfig.Int("haimingDistanceWeight")
	if err != nil {
		maxDist = 10
	}

	newsIds := []int64{}
	s1 := ""
	for i := 0; i < 4; i++ {
		s1 += dexbin(sh[i])

		if ids, ok := ContMap[sh[i]]; ok {
			for j := 0; j < len(ids); j++ {
				newsIds = append(newsIds, ids[j])
			}
		}
	}

	newsIds = uniqSliceInt64(newsIds)

	s2 := ""
	for i := 0; i < len(newsIds); i++ {
		if news, ok := NewsMap[newsIds[i]]; ok {
			s2 = dexbin(news[0]) + dexbin(news[1]) + dexbin(news[2]) + dexbin(news[3])
			dict, err := hamming(s1, s2)

			if err == nil && dict < maxDist {
				resp = true
				break
			}
		}
	}
	return
}

// 如果是在指定的距离内，返回true (表示有重复)
func checkTitle(sh *[8]uint16) (hasDeup bool) {
	maxDist, err := beego.AppConfig.Int("haimingDistance")
	if err != nil {
		maxDist = 12
	}

	newsIds := []int64{}
	s1 := ""
	for i := 4; i < 8; i++ {
		s1 += dexbin(sh[i])
		if ids, ok := TitleMap[sh[i]]; ok {
			for j := 0; j < len(ids); j++ {
				newsIds = append(newsIds, ids[j])
			}
		}
	}

	newsIds = uniqSliceInt64(newsIds)

	s2 := ""
	for i := 0; i < len(newsIds); i++ {
		if news, ok := NewsMap[newsIds[i]]; ok {
			s2 = dexbin(news[4]) + dexbin(news[5]) + dexbin(news[6]) + dexbin(news[7])
			dict, err := hamming(s1, s2)
			if err == nil && dict < maxDist {
				hasDeup = true
				break
			}
		}
	}
	return
}

func hamming(s1 string, s2 string) (distance int, err error) {
	// index by code point, not byte
	r1 := []rune(s1)
	r2 := []rune(s2)

	if len(r1) != len(r2) {
		err = errors.New("Hamming distance of different sized strings.")
		return
	}

	for i, v := range r1 {
		if r2[i] != v {
			distance += 1
		}
	}
	return
}

func uniqSliceInt64(ids []int64) (res []int64) {
	if len(ids) < 2 {
		res = ids
	} else {
		res = []int64{}
		tmp := make(map[int64]bool)
		for _, id := range ids {
			if _, ok := tmp[id]; !ok {
				tmp[id] = true
				res = append(res, id)
			}
		}
	}
	return
}

func dexbin(i uint16) string {
	return fmt.Sprintf("%016b", i)
}
