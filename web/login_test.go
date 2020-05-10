
package web

import(
	"testing"
	"errors"
	"github.com/wusidn/qiaqia/dao"
	"github.com/golang/mock/gomock"
	"github.com/wusidn/qiaqia/mockdao"
)

func TestLogin(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDao := mockdao.NewMockDao(ctrl)
	mockTableUsers := mockdao.NewMockTableUsers(ctrl)

	mockDao.EXPECT().GetTableUsers().Return(mockTableUsers).AnyTimes()

	email := "wangqiaqia@163.com"
	encodePassWD := "123123"

	mockTableUsers.EXPECT().QueryUserInfoByEmail(email).Return(dao.UserInfo{
			Id: 0, 
			Locked: 0,
			Email: email,
			PhoneNumber: "13211111111",
			EncodePassword: encodePassWD,
			IsSystem: 1,
		}, nil)

	userInfoNotExistErr := errors.New("db not find UserInfo by email[wangbaichi@163.com]")
	mockTableUsers.EXPECT().QueryUserInfoByEmail("wangbaichi@163.com").Return(dao.UserInfo{}, userInfoNotExistErr)
	
	service := New(mockDao)

	//test login success
	_, err := service.login(email, "123123")
	if err != nil{
		t.Logf("test user login fail [%v]", err.Error())
		t.Fail()
	}

	//test login fail email not exist
	_, err = service.login("wangbaichi@163.com", "123123")
	if err == nil{
		t.FailNow()
	}

}