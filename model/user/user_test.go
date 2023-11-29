package user

import (
	"gin-init/common/database"
	"gin-init/test"
	"gin-init/util"
	"gorm.io/gorm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModelUser(t *testing.T) {
	test.InitSetting()
	database.GetInstanceConnection().Init()
	db := database.GetInstanceConnection().GetDB()
	userId := util.NewShortIDString("user")
	user := &User{
		UserID: userId,
	}
	err := CreateUser(db, user)
	assert.Nil(t, err)

	err = DeleteUserByUserId(db, userId)
	assert.Nil(t, err)

	user.UserAccount = "test"
	user.UserPassword = "testtest"
	err = UpdateUser(db, user)
	assert.Nil(t, err)

	queryUser, err := QueryUserInfoByUserId(db, userId)
	assert.Nil(t, err)
	assert.NotNil(t, queryUser)

	notExistUserId := util.NewShortIDString("user")
	queryUser, err = QueryUserInfoByUserId(db, notExistUserId)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	assert.NotNil(t, queryUser)

	userList, err := QueryUserList(db, 1, 20)
	assert.Nil(t, err)
	assert.NotNil(t, userList)

}
