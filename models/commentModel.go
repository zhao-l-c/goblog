package models
import (
    "time"
    "github.com/astaxie/beego/orm"
    "strconv"
)

type Comment struct {
    Id int64
    // 文章id
    Tid int64
    // 评论人id
    UserId string
    // 评论人姓名
    UserName string
    // 评论时间
    ReplyTime time.Time
    // 评论内容
    Content string `orm:"size(1000)"`
}

func AddComment(tid, content string) error {
    titleId, err := strconv.ParseInt(tid, 10, 64)
    if err != nil {
        return err
    }
    orm := orm.NewOrm()
    comment := &Comment{
        Tid: titleId,
        Content: content,
        ReplyTime: time.Now(),
    }
    _, err = orm.Insert(comment)
    return err
}

func QueryCommentsByTid(tid string) ([]*Comment, error) {
    titleId, err := strconv.ParseInt(tid, 10, 64)
    if err != nil {
        return nil, err
    }
    orm := orm.NewOrm()
    comments := make([]*Comment, 0)
    _, err = orm.QueryTable("comment").Filter("tid", titleId).All(&comments)
    return comments, err
}

func DeleteComment(id string) error {
    cid, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        return err
    }
    orm := orm.NewOrm()
    comment := &Comment{Id: cid}
    _, err = orm.Delete(comment)
    return err
}