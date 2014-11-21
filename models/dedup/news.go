package dedup

import (
	"related-news/utils/mongo"
	"time"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	mgo "gopkg.in/mgo.v2"
)

type News struct {
	Id         int64
	CateId     int `bson:"cate_id"`
	Sh1        uint16 `bson:"sh1"`
	Sh2        uint16 `bson:"sh2"`
	Sh3        uint16 `bson:"sh3"`
	Sh4        uint16 `bson:"sh4"`
	ShT1       uint16 `bson:"sh_t1"`
	ShT2       uint16 `bson:"sh_t2"`
	ShT3       uint16 `bson:"sh_t3"`
	ShT4       uint16 `bson:"sh_t4"`
	CreatedAt  time.Time `bson:"created_at"`
}

/**
 * 初始化最近days天的数据到BigMap里
 */
func AnalyzeNews() {
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

	beego.Debug("beginAt: ", beginAt)

	var lastId int64 = 0
	var allNews []News
	var query *mgo.Query

	for i := 0; i < maxLoop; i++ {
		beego.Debug(i, " times")

		allNews = []News{}

		if lastId == 0 {
			query = mongodb.Find(bson.M{"created_at": bson.M{"$gte": beginAt}, "status":1, "cate_id": bson.M{"$ne": 9}})
		}else {
			query = mongodb.Find(bson.M{"created_at": bson.M{"$gte": beginAt}, "status":1, "id": bson.M{"$lt": lastId}, "cate_id": bson.M{"$ne": 9}})
		}

		err := query.Select(bson.M{"id":1, "sh1":1, "sh2":1, "sh3":1, "sh4": 1, "sh_t1":1, "sh_t2":1, "sh_t3":1, "sh_t4": 1}).Limit(limit).Sort("-id").All(&allNews)
		if err != nil {
			beego.Error("query news error: ", err.Error())
		}else {
			if len(allNews) == 0 {
				break;
			}

			for _, news := range allNews {
				AppendToNewsMap(&news)
				lastId = news.Id
			}
		}
	}

	beego.Debug("Content Map len: ", ContMapLen())
	beego.Debug("Title Map len: ", TitleMapLen())
	beego.Debug("News Map len: ", NewsMapLen())
}

//
//func GetSimilarNewsIds(id int64, keywords []string) (ids []int64) {
//	ids = []int64{}
//
//	if len(keywords) > 0 {
//		weight, err := beego.AppConfig.Int("similarWeight")
//		if err != nil {
//			weight = 5
//		}
//
//		idsMap := CountBigMap(keywords)
//		idsMapUseful := make(map[int64]int)
//		for k, v := range idsMap {
//			if v > weight && k != id {
//				idsMapUseful[k] = v
//			}
//		}
//		if len(idsMapUseful) > 0 {
//			idsTmp := sortMap(idsMapUseful)
//			if (len(idsTmp) > 6) {
//				ids = idsTmp[:6]
//			}else {
//				ids = idsTmp
//			}
//		}
//	}
//	return;
//}

