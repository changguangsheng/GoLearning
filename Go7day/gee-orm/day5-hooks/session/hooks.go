package session

import (
	"geeorm/log"
	"reflect"
)

const (
	BeforeQuery  = "BeforeQuery"
	AfterQuery   = "AfterQuery"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate  = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete  = "AfterDelete"
	BeforeInsert = "BeforeInsert"
	AfterInsert  = "AfterInsert"
)

func (s *Session) CallMethod(method string, value interface{}) {
	fm := reflect.ValueOf(s.RefTable().Model).MethodByName(method)
	if value != nil {
		fm = reflect.ValueOf(value).MethodByName(method)
	}
	//s.RefTable().Model 或 value 即当前会话正在操作的对象，使用 MethodByName 方法反射得到该对象的方法。
	param := []reflect.Value{reflect.ValueOf(s)}
	if fm.IsValid() {
		if v := fm.Call(param); len(v) > 0 {
			if err, ok := v[0].Interface().(error); ok {
				log.Error(err)
			}
		}
	}
	return
}

type Account struct {
	ID       int `geeorm:"PRIMARY KEY"`
	Password string
}

func (account *Account) BeforeInsert(s *Session) error {
	log.Info("before insert")
	account.Password = "******"
	return nil
}
//func TestSession_CallMethod(t *testing.T) {
//	s := NewSession().Model(&Account{})
//	_ = s.DropTable()
//	_ = s.CreateTable()
//	_, _ = s.Insert(&Account{1, "123456", &Account{2, "qwerty"}})
//
//	u := &Account{}
//
//	err := s.First(u)
//	if err != nil || u.ID != 1001 || u.Password != "******" {
//		t.Fatal("Failed to call hooks after query, got", u)
//	}
//}
