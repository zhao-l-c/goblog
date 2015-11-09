package models

import (
    "github.com/astaxie/beego/orm"
    "strconv"
    "strings"
    "GoWeb/beego/commons"
    "path"
    "os"
)

type Topic struct {
    Id              int64
    Uid             int64
    Title           string
    Content         string `orm:"size(5000)"`
    Tags            string `orm:"size(1000)"`
    Attachment      string
    CreateTime      string `orm:"index"`
    UpdateTime      string `orm:"index"`
    // 文章分类
    Category        string    `orm:"index"`
    // 浏览数
    Views           int64     `orm:"index"`
    // 作者
    Author          string
    // 最后回复时间
    LastReplyTime   string `orm:"index"`
    // 回复数
    ReplyCount      int64
    // 最后回复用户id
    LastReplyUserId int64
}

func AddTopic(title, content, tags, cid, attachName string) error {
    orm := orm.NewOrm()
    categoryId, err := strconv.ParseInt(cid, 10, 64)
    if err != nil {
        return err
    }
    category := &Category{Id: categoryId}
    err = orm.QueryTable("category").Filter("id", categoryId).One(category)
    if err != nil {
        return err
    }
    // java scala linux => $java#$scala#$linux#
    tagArray := strings.Split(strings.Trim(tags, " "), " ")
    tags = "$" + strings.Join(tagArray, "#$")  + "#"

    topic := &Topic{
        Title: title,
        Content: content,
        Tags: tags,
        Attachment: attachName,
        CreateTime: CurrentTime(),
        UpdateTime: CurrentTime(),
        LastReplyTime: CurrentTime(),
        Category: category.Title,
    }
    // add one topic
    _, err = orm.Insert(topic)
    if err != nil {
        return err
    }
    // update category's count
    category.TopicCount++
    category.TopicTime = CurrentTime()
    _, err = orm.Update(category)
    if err != nil {
        return err
    }
    // insert or update tag
    InsertOrUpdateTags(tagArray)
    return err
}

func AllTopics(orderBy bool, category, tag string) ([]*Topic, error) {
    orm := orm.NewOrm()
    topics := make([]*Topic, 0)
    var err error
    query := orm.QueryTable("topic")
    if len(category) > 0 {
        query = query.Filter("category", category)
    }
    if len(tag) > 0 {
        query = query.Filter("tags__contains", "$" + tag + "#")
    }
    if orderBy {
        query = query.OrderBy("-createTime")
    }
    _, err = query.All(&topics)
    return topics, err
}

func GetTopic(id string) (*Topic, error) {
    tid, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        return nil, err
    }
    orm := orm.NewOrm()

    topic := &Topic{
        Id: tid,
    }
    err = orm.QueryTable("topic").Filter("id", tid).One(topic)
    if err != nil {
        return nil, err
    }
    // update views
    topic.Views++
    orm.Update(topic)
    topic.Tags = strings.Replace(strings.Replace(topic.Tags, "#", " ", -1), "$", "", -1)
    return topic, nil
}

func UpdateTopic(id, title, content, tags, cid, attachName string, attached bool) error {
    tid, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        return err
    }
    categoryId, err := strconv.ParseInt(cid, 10, 64)
    if err != nil {
        return err
    }

    orm := orm.NewOrm()
    category := &Category{Id: categoryId}
    err = orm.QueryTable("category").Filter("id", categoryId).One(category)
    if err != nil {
        return err
    }
    topic := &Topic{Id: tid}
    tagArray := strings.Split(strings.Trim(tags, " "), " ")
    tags = "$" + strings.Join(tagArray, "#$") + "#"
    var oldTags string
    var oldTitle string
    var oldAttachName string
    // valida exist and modify
    err = orm.Read(topic)
    if err != nil {
        return err
    } else {
        oldTitle = topic.Category
        oldTags = topic.Tags
        oldAttachName = topic.Attachment

        topic.Title = title
        topic.Content = content
        topic.Tags = tags
        if attached {
            topic.Attachment = attachName
        }
        topic.UpdateTime = CurrentTime()
        topic.Category = category.Title
        _, err = orm.Update(topic)
        if err != nil {
            return err
        }
    }

    // delete old file
    if attached {
        oldAttachName = path.Join("attachment", oldAttachName)
        if _, err := os.Stat(oldAttachName); err == nil {
            os.Remove(oldAttachName)
        }
    }
    // update category's topicCount
    if oldTitle != category.Title {
        // new category plus one
        category.TopicCount++
        _, err = orm.Update(category)
        if err != nil {
            return err
        }
        // TODO 如果旧的分类被删除，则报错，错误信息是：<QuerySeter> no row found
        count, err := orm.QueryTable("category").Filter("title", oldTitle).Count()
        if err != nil {
            return err
        }
        // 分类存在，则分类数目减一
        if count == 1 {
            oldCategory := &Category{Title: oldTitle}
            err = orm.QueryTable("category").Filter("title", oldTitle).One(oldCategory)
            if err != nil {
                return err
            }
            oldCategory.TopicCount--
            if oldCategory.TopicCount >= 0 {
                _, err = orm.Update(oldCategory)
                if err != nil {
                    return err
                }
            }
        }
    }
    oldTags = strings.Replace(strings.Replace(oldTags, "#", " ", -1), "$", "", -1)
    oldTagArray := strings.Split(strings.Trim(oldTags, " "), " ")
    return ModifyTags(commons.Diff(tagArray, oldTagArray))
}

func DeleteTopic(id string) error {
    tid, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        return err
    }
    orm := orm.NewOrm()
    topic := &Topic{
        Id: tid,
    }
    _, err = orm.Delete(topic)
    return err
}