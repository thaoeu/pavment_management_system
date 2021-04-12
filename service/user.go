/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-01-14
 * Time: 22:25
 */
package service

import (
	"github.com/izghua/zgh"
	"github.com/thaoeu/pavment_management_system/conf"
	"github.com/thaoeu/pavment_management_system/entity"
)

func GetUserById(userId int) (*entity.ZUsers, error) {
	user := new(entity.ZUsers)
	_, err := conf.SqlServer.Id(userId).Cols("name", "email").Get(user)
	if err != nil {
		zgh.ZLog().Error("message", "service.GetUserById", "error", err.Error())
		return user, err
	}
	return user, nil
}

func UserCnt() (cnt int64, err error) {
	user := new(entity.ZUsers)
	cnt, err = conf.SqlServer.Count(user)
	return
}
