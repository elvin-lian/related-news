package controllers

import (
	"github.com/astaxie/beego"
	"related-news/models"
	"fmt"
)

type NewsController struct{
	beego.Controller
}

func (this *NewsController) Get() {
	data := make(map[string]interface{})
	newsId, err := this.GetInt(":id")
	if err == nil && newsId > 0 {
		news, err := models.GetNews(newsId)
		if err != nil {
			data["code"] = 0
			data["message"] = err
		} else {
			relatedIds := news.RelatedIds
			//if len(relatedIds) == 0 {
			relatedIds = models.GetSimilarNewsIds(news.Id, news.Tags)
			err := models.UpdateNewsRelatedIds(news.Id, relatedIds)
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
		news, err := models.GetNews(newsId)
		if err == nil {
			for _, k := range news.Tags {
				models.AppendToBigMap(k, news.Id)
			}
			this.RenderJson(1, nil, fmt.Sprintf("big map len: %d", models.BigMapLen()))

		}else {
			this.RenderJson(0, nil, err.Error())
		}

	}else {
		this.RenderJson(0, nil, err.Error())
	}
}

func (this *NewsController) Add() {
	pk := this.GetString("pk")
	news, err := models.GetNewsByPk(pk)
	if err == nil {
		for _, k := range news.Tags {
			models.AppendToBigMap(k, news.Id)
		}
		relatedIds := models.GetSimilarNewsIds(news.Id, news.Tags)
		err := models.UpdateNewsRelatedIds(news.Id, relatedIds)

		mess := ""
		if err != nil {
			mess = err.Error()
		}
		data := make(map[string]interface{})
		data["code"] = 1
		data["id"] = news.Id
		data["ids"] = relatedIds

		beego.Debug(models.BigMapLen())
		this.RenderJson(1, data, mess)
	}else {
		this.RenderJson(0, nil, err.Error())
	}
}

func (this *NewsController) Analyze() {
	models.AnalyzeNews(0)
	this.RenderJson(1, nil, fmt.Sprintf("big map len: %d", models.BigMapLen()))
}

func (this *NewsController) InitNews() {
	models.InitNewsRelated(0)
	this.RenderJson(1, nil, "")
}

func (this *NewsController) Len() {
	data := make(map[string]interface{})
	data["bigMapLen"] = models.BigMapLen()
	this.RenderJson(1, data, "")
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
