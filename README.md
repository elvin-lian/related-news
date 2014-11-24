# 相关资讯
------

## 算法

1) 将最近2天的资讯的关键字保存到 bigMap:
```
{
  k1: [newsId1, newsId2],
  k2: [newsId2, newsId3, newsId4]
}

```
此bigMap每天凌晨更新一次

2) 相关文章

如果某篇资讯的关键字与其它资讯的有n (n=7) 个相同的，表示是相关的文章

## 依赖

```
go get gopkg.in/mgo.v2
go get gopkg.in/mgo.v2/bson
```

## 接口

1. /v1/news/:newsId  返回资讯的相关文章IDs

2. /v1/news/:newsId/append 将新的资讯添加到bigMap中

3. /v1/news/analyze 生成bigMap

4. /v1/news/init_news 生成最近2天的资讯的相关文章ID

5. /v1/news/len  返回bigMap的长度

6. /v1/news/add?pk=primaryKey  将资讯添加到bigMap并返回资讯的相关文章IDs

7. /v1/news/analyze_dedup  SimHash过虑