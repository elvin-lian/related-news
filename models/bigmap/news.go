package bigmap

import (
	"related-news/utils/mongo"
	"related-news/utils/sorter"
	"time"
	"sort"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	mgo "gopkg.in/mgo.v2"
)

type News struct {
	Id         int64
	CateId     int `bson:"cate_id"`
	Content    string
	Tags       []string
	CreatedAt  time.Time `bson:"created_at"`
	RelatedIds []int64   `bson:"related_ids"`
}

func GetNews(id int64) (news News, err error) {
	mongodb := mongo.Collection("news")
	news = News{}
	err = mongodb.Find(bson.M{"id": id}).Select(bson.M{"id":1, "content":1, "title":1, "tags":1, "related_ids": 1}).One(&news)
	return
}

func GetNewsByPk(pk string) (news News, err error) {
	mongodb := mongo.Collection("news")
	news = News{}
	err = mongodb.Find(bson.M{"_id": bson.ObjectIdHex(pk)}).Select(bson.M{"id":1, "content":1, "title":1, "tags":1, "related_ids": 1}).One(&news)
	return
}

func UpdateNewsRelatedIds(id int64, relatedIds []int64) (err error) {
	mongodb := mongo.Collection("news")

	err = mongodb.Update(bson.M{"id": id}, bson.M{"$set": bson.M{"related_ids": relatedIds}})
	return
}

/**
 * 初始化最近days天的数据到BigMap里
 */
func AnalyzeNews() {
	CleanBigMap()

	days := 2
	daysConf, err := beego.AppConfig.Int("maxDays")
	if err == nil {
		days = daysConf
	}

	mongodb := mongo.Collection("news")

	limit := 5000
	maxLoop, err := beego.AppConfig.Int("maxLoop")
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

		err := query.Select(bson.M{"id":1, "tags": 1}).Limit(limit).Sort("-id").All(&allNews)
		if err != nil {
			beego.Error("query news error: ", err.Error())
		}else {
			if len(allNews) == 0 {
				break;
			}

			for _, news := range allNews {
				keywords := news.Tags
				for _, k := range keywords {
					AppendToBigMap(k, news.Id)
				}
				lastId = news.Id
			}
		}
	}

	beego.Debug("bigMap len: ", BigMapLen())
}

// 初始化最近days的相关资讯
func InitNewsRelated(days int) {
	if days == 0 {
		daysConf, err := beego.AppConfig.Int("maxDays")
		if err != nil {
			days = 2
		}else {
			days = daysConf
		}
	}

	mongodb := mongo.Collection("news")

	limit := 5000
	maxLoop, err := beego.AppConfig.Int("maxLoop")
	if err != nil {
		maxLoop = 100
	}

	now := time.Now()
	beginOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	beginAt := beginOfToday.AddDate(0, 0, -1*days) // days ago

	var lastId int64 = 0
	var allNews []News
	var query *mgo.Query
	var allNewsCount int = 0
	var hasRelatedNewsCount int = 0

	for i := 0; i < maxLoop; i++ {
		beego.Debug(i, " times")

		allNews = []News{}

		if lastId == 0 {
			query = mongodb.Find(bson.M{"created_at": bson.M{"$gte": beginAt}, "status":1, "cate_id": bson.M{"$ne": 9}})
		}else {
			query = mongodb.Find(bson.M{"created_at": bson.M{"$gte": beginAt}, "status":1, "id": bson.M{"$lt": lastId}, "cate_id": bson.M{"$ne": 9}})
		}

		err := query.Select(bson.M{"id":1, "tags": 1, "related_ids":1}).Limit(limit).Sort("-id").All(&allNews)
		if err != nil {
			beego.Error("query news error: ", err.Error())
		}else {
			if len(allNews) == 0 {
				break;
			}

			for _, news := range allNews {
				allNewsCount ++
				relatedIds := GetSimilarNewsIds(news.Id, news.Tags)
				if len(news.RelatedIds) == 0 && len(relatedIds) == 0 {

				} else {
					err := UpdateNewsRelatedIds(news.Id, relatedIds)
					if err == nil {
						beego.Debug("news.Id: ", news.Id, relatedIds)
						hasRelatedNewsCount ++
					}
				}
				lastId = news.Id
			}
		}
	}

	beego.Debug("allNewsCount: ", allNewsCount)
	beego.Debug("hasRelatedNewsCount: ", hasRelatedNewsCount)
}

func GetSimilarNewsIds(id int64, keywords []string) (ids []int64) {
	ids = []int64{}

	if len(keywords) > 0 {
		weight, err := beego.AppConfig.Int("similarWeight")
		if err != nil {
			weight = 5
		}

		idsMap := CountBigMap(keywords)
		idsMapUseful := make(map[int64]int)
		for k, v := range idsMap {
			if v > weight && k != id {
				idsMapUseful[k] = v
			}
		}
		if len(idsMapUseful) > 0 {
			idsTmp := sortMap(idsMapUseful)
			if (len(idsTmp) > 6) {
				ids = idsTmp[:6]
			}else {
				ids = idsTmp
			}
		}
	}
	return;
}

// 排序
func sortMap(idsMap map[int64]int) (ids []int64) {
	ms := sorter.Int64IntDesc(idsMap)
	sort.Sort(ms)

	for _, item := range ms {
		ids = append(ids, item.Key)
	}
	return
}