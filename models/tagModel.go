package models
import (
    "github.com/astaxie/beego/orm"
)

type Tag struct {
    Id int64
    Name string `orm:"index"`
    Count int64
}

func AllTags() ([]*Tag, error) {
    orm := orm.NewOrm()
    tags := make([]*Tag, 0)
    _, err := orm.QueryTable("tag").All(&tags)
    return tags, err
}

func InsertOrUpdateTags(tags []string) error {
    var count int64
    var err error
    for _, name := range tags  {
        count, err = QueryTagByName(name)
        if err == nil {
            if count == 0 {
                err = AddTag(name)
            } else if count == 1 {
                err = IncreaseTagCount(name)
            }
            if err != nil {
                return err
            }
        } else {
            return err
        }
    }
    return nil
}

func DecreaseTagCounts(tags []string) error {
    for _, name := range tags {

    }
}



func AddTag(name string) error {
    orm := orm.NewOrm()
    tag := &Tag{Name: name, Count: 1}
    _, err := orm.Insert(tag)
    return err
}

func IncreaseTagCount(name string) error {
    orm := orm.NewOrm()
    tag := &Tag{Name: name}
    err := orm.QueryTable("tag").Filter("name", name).One(tag)
    if err != nil {
        return err
    }
    tag.Count++
    _, err = orm.Update(tag)
    return err
}

func QueryTagByName(name string) (int64, error) {
    orm := orm.NewOrm()
    count, err := orm.QueryTable("tag").Filter("name", name).Count()
    return count, err
}