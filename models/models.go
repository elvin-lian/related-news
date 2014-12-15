package models

import (
	"time"

	"related-news/utils/mongo"

	"gopkg.in/mgo.v2/bson"
)

type News struct {
	Id         int64
	CateId     int `bson:"cate_id"`
	Content    string
	Tags       []string
	Sh1        uint16 `bson:"sh1"`
	Sh2        uint16 `bson:"sh2"`
	Sh3        uint16 `bson:"sh3"`
	Sh4        uint16 `bson:"sh4"`
	ShT1       uint16 `bson:"sh_t1"`
	ShT2       uint16 `bson:"sh_t2"`
	ShT3       uint16 `bson:"sh_t3"`
	ShT4       uint16 `bson:"sh_t4"`
	CreatedAt  time.Time `bson:"created_at"`
	RelatedIds []int64   `bson:"related_ids"`
}

func GetNewsByPk(pk string) (news News, err error) {
	mongodb := mongo.Collection("news")
	news = News{}
	err = mongodb.Find(bson.M{"_id": bson.ObjectIdHex(pk)}).
	Select(bson.M{"id":1, "content":1, "title":1, "tags":1, "related_ids": 1, "sh1":1, "sh2":1, "sh3":1, "sh4": 1, "sh_t1":1, "sh_t2":1, "sh_t3":1, "sh_t4": 1}).
	One(&news)
	return
}
