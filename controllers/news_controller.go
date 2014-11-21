package controllers

import (
	"github.com/astaxie/beego"
	"related-news/models/bigmap"
	"related-news/models/dedup"
	"fmt"
)

type NewsController struct{
	beego.Controller
}

func (this *NewsController) Get() {
	data := make(map[string]interface{})
	newsId, err := this.GetInt(":id")
	if err == nil && newsId > 0 {
		news, err := bigmap.GetNews(newsId)
		if err != nil {
			data["code"] = 0
			data["message"] = err
		} else {
			relatedIds := news.RelatedIds
			//if len(relatedIds) == 0 {
			relatedIds = bigmap.GetSimilarNewsIds(news.Id, news.Tags)
			err := bigmap.UpdateNewsRelatedIds(news.Id, relatedIds)
			data["message"] = err
			//}
			data["code"] = 1
			data["id"] = newsId
			data["ids"] = relatedIds
		}
	}
	this.RenderJson(1, data, "")
}

func (this *NewsController) Append() {
	newsId, err := this.GetInt(":id")
	if err == nil && newsId > 0 {
		news, err := bigmap.GetNews(newsId)
		if err == nil {
			for _, k := range news.Tags {
				bigmap.AppendToBigMap(k, news.Id)
			}
			this.RenderJson(1, nil, fmt.Sprintf("big map len: %d", bigmap.BigMapLen()))

		}else {
			this.RenderJson(0, nil, err.Error())
		}

	}else {
		this.RenderJson(0, nil, err.Error())
	}
}

func (this *NewsController) Add() {
	pk := this.GetString("pk")
	news, err := bigmap.GetNewsByPk(pk)
	if err == nil {
		for _, k := range news.Tags {
			bigmap.AppendToBigMap(k, news.Id)
		}
		relatedIds := bigmap.GetSimilarNewsIds(news.Id, news.Tags)
		err := bigmap.UpdateNewsRelatedIds(news.Id, relatedIds)

		mess := ""
		if err != nil {
			mess = err.Error()
		}
		data := make(map[string]interface{})
		data["code"] = 1
		data["id"] = news.Id
		data["ids"] = relatedIds

		beego.Debug(bigmap.BigMapLen())
		this.RenderJson(1, data, mess)
	}else {
		this.RenderJson(0, nil, err.Error())
	}
}

func (this *NewsController) Analyze() {
	bigmap.AnalyzeNews()
	this.RenderJson(1, nil, fmt.Sprintf("big map len: %d", bigmap.BigMapLen()))
}

func (this *NewsController) InitNews() {
	bigmap.InitNewsRelated(0)
	this.RenderJson(1, nil, "")
}

func (this *NewsController) Len() {
	data := make(map[string]interface{})
	data["bigMapLen"] = bigmap.BigMapLen()
	this.RenderJson(1, data, "")
}

func (this *NewsController) DedupAnalyze() {
	dedup.AnalyzeNews()
	this.RenderJson(1, nil,
		fmt.Sprintf("title map len: %d \n content map len: %d \n news map len: %d",
			dedup.TitleMapLen(),
			dedup.ContMapLen(),
			dedup.NewsMapLen()))
}



func (this *NewsController) RenderJson(code int, data map[string]interface{}, mess string) {
	json := make(map[string]interface{})

	json["code"] = code
	if "" != mess {
		json["message"] = mess
	}

	if nil != data {
		json["data"] = data
	}

	this.Data["json"] = json
	this.ServeJson()
	this.StopRun()
}
