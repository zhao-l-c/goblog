package models

import (
    "time"
    "github.com/astaxie/beego/orm"
    "strconv"
)

type Topic struct {
    Id              int64
    Uid             int64
    Title           string
    Content         string `orm:"size(5000)"`
    Attachment      string
    CreateTime      time.Time `orm:"index"`
    UpdateTime      time.Time `orm:"index"`
    Category        string    `orm:"index"`
    Views           int64     `orm:"index"`
    Author          string
    ReplyTime       time.Time `orm:"index"`
    ReplyCount      int64
    ReplyLastUserId int64
}

func AddTopic(title, content , cid string) error {
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

    topic := &Topic{
        Title: title,
        Content: content,
        CreateTime: time.Now(),
        UpdateTime: time.Now(),
        ReplyTime: time.Now(),
        Category: category.Title,
    }
    // add one topic
    _, err = orm.Insert(topic)
    // update category's count
    category.TopicCount++
    category.TopicTime = time.Now()
    orm.Update(category)
    return err
}

func AllTopics(orderBy bool) ([]*Topic, error) {
    orm := orm.NewOrm()
    topics := make([]*Topic, 0)
    var err error
    if orderBy {
        _, err = orm.QueryTable("topic").OrderBy("-createTime").All(&topics)
    } else {
        _, err = orm.QueryTable("topic").All(&topics)
    }
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
    return topic, nil
}

func UpdateTopic(id, title, content, cid string) error {
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
    // valida exist and modify
    var oldCategoryTitle string
    if orm.Read(topic) == nil {
        oldCategoryTitle = topic.Category
        topic.Title = title
        topic.Content = content
        topic.UpdateTime = time.Now()
        topic.Category = category.Title
        orm.Update(topic)
    }


    // update category's topicCount
    if oldCategoryTitle != category.Title {
        // new category plus one
        category.TopicCount++
        orm.Update(category)
        // old category minus one
        oldCategory := &Category{Title: oldCategoryTitle}
        err = orm.QueryTable("category").Filter("title", oldCategoryTitle).One(oldCategory)
        if err != nil {
            return err
        }
        oldCategory.TopicCount--
        if oldCategory.TopicCount >= 0 {
            orm.Update(oldCategory)
        }
    }

    return nil
}

func DeleteTopic(id string) error {
    tid, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        return err
    }
    orm := orm.NewOrm()
    topic := &Topic {
        Id: tid,
    }
    _, err = orm.Delete(topic)
    return err
}