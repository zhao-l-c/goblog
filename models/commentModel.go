package models
import (
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
    ReplyTime string
    // 评论内容
    Content string `orm:"size(1000)"`
}

/*
tid: 文章id
content：评论内容
*/
func AddComment(tid, content string) error {
    titleId, err := strconv.ParseInt(tid, 10, 64)
    if err != nil {
        return err
    }
    orm := orm.NewOrm()
    comment := &Comment{
        Tid: titleId,
        Content: content,
        ReplyTime: CurrentTime(),
    }
    _, err = orm.Insert(comment)
    if err != nil {
        return err
    }
    // 文章评论数增一
    topic := &Topic{
        Id: titleId,
    }
    if orm.Read(topic) == nil {
        topic.ReplyCount++
        topic.LastReplyTime = CurrentTime()
        // TODO set current user id
        // topic.ReplyLastUserId = current_user_id
        orm.Update(topic)
    }
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

func DeleteComment(cid, tid string) error {
    commentId, err := strconv.ParseInt(cid, 10, 64)
    if err != nil {
        return err
    }
    orm := orm.NewOrm()
    comment := &Comment{Id: commentId}
    _, err = orm.Delete(comment)
    if err != nil {
        return err
    }
    topicId, err := strconv.ParseInt(tid, 10, 64)
    if err != nil {
        return err
    }
    comments := make([]*Comment, 0)
    _, err = orm.QueryTable("comment").Filter("tid", topicId).OrderBy("-replyTime").All(&comments)
    if err != nil {
        return err
    }
    topic := &Topic{Id: topicId}
    if orm.Read(topic) == nil && len(comments) > 0 {
        topic.ReplyCount--
        topic.LastReplyTime = comments[0].ReplyTime
        // TODO set last user id
        // topic.ReplyLastUserId = comments[0].UserId
        orm.Update(topic)
    }
    return err
}