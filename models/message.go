package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Message struct {
	Id      int
	Content string
	AddTime int64
}

type MessageUser struct {
	Id        int
	UserId    int
	MessageId int64
	AddTime   int64
	Status    int
}

func init() {
	orm.RegisterModel(new(Message), new(MessageUser))
}

// SendMsg 保存消息
func SendMsg(content string) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(&Message{
		Content: content,
		AddTime: time.Now().Unix(),
	})
	return
}

// SendUser 保存消息接收人
func Sender(userID int, messageID int64) error {
	o := orm.NewOrm()
	_, err := o.Insert(&MessageUser{
		UserId:    userID,
		MessageId: messageID,
		AddTime:   time.Now().Unix(),
		Status:    1,
	})
	return err
}
