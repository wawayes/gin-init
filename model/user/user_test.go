package user

import (
	"gin-init/common/database"
	"gin-init/test"
	"gin-init/util"
	"gorm.io/gorm"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cu = User{}

func TestModelUser(t *testing.T) {
	test.InitSetting()
	database.GetInstanceConnection().Init()
	db := database.GetInstanceConnection().GetDB()
	userId := util.NewShortIDString("user")
	user := &User{
		UserID: userId,
	}
	err := cu.CreateUser(db, user)
	assert.Nil(t, err)

	err = cu.DeleteUserByUserId(db, userId)
	assert.Nil(t, err)

	user.UserAccount = "test"
	user.UserPassword = "testtest"
	err = cu.UpdateUser(db, user)
	assert.Nil(t, err)

	queryUser, err := cu.QueryUserInfoByUserId(db, userId)
	assert.Nil(t, err)
	assert.NotNil(t, queryUser)

	notExistUserId := util.NewShortIDString("user")
	queryUser, err = cu.QueryUserInfoByUserId(db, notExistUserId)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	assert.NotNil(t, queryUser)

	userList, err := cu.QueryUserList(db, 1, 20)
	assert.Nil(t, err)
	assert.NotNil(t, userList)

}
