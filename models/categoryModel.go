package models

import (
    "time"
    "github.com/astaxie/beego/orm"
    "strconv"
)


type Category struct {
    Id              int64
    Title           string
    Views           int64     `orm:"index"`
    CreateTime      time.Time `orm:"index"`
    TopicTime       time.Time `orm:"index"`
    TopicCount      int64
    TopicLastUserId int64
}

func AddCategory(title string) error {
    orm := orm.NewOrm()
    category := &Category{
        Title: title,
        CreateTime: time.Now(),
        TopicTime: time.Now(),
    }
    // query if duplicated
    err := orm.QueryTable("category").Filter("title", title).One(category)
    if err == nil {
        return err
    }

    // insert new one
    _, err = orm.Insert(category)
    if err != nil {
        return err
    }
    return nil
}

// query all
func AllCategories(orderBy bool) ([]*Category, error) {
    orm := orm.NewOrm()
    categories := make([]*Category, 0)
    var err error
    if orderBy {
        _, err = orm.QueryTable("category").OrderBy("-createTime").All(&categories)
    } else {
        _, err = orm.QueryTable("category").All(&categories)
    }
    return categories, err
}

// delete
func DeleteCategory(id string) error {
    cid, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        return err
    }
    orm := orm.NewOrm()
    category := &Category{Id: cid}
    _, err = orm.Delete(category)
    return err
}
