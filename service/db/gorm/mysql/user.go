package db

import (
	"github.com/qinhan-shu/gp-server/model/gorm"
)

// GetUserByID : get uer model by user id
func (m *MysqlDriver) GetUserByID(id int64) (*model.User, error) {
	var user model.User
	if err := m.conn.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// CheckPlayer : get uer model by username and password
func (m *MysqlDriver) CheckPlayer(username, password string) (*model.User, error) {
	var user model.User
	if err := m.conn.Where("username = ? and password = ?", username, password).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUsersByRole : get users by required role
func (m *MysqlDriver) GetUsersByRole(role int64) ([]*model.User, error) {
	var users []*model.User
	var err error
	if role < 0 {
		// role < 0 : get all users
		err = m.conn.Find(&users).Error
	} else {
		err = m.conn.Where("role = ? ", role).Find(&users).Error
	}

	if err != nil {
		return nil, err
	}
	return users, nil
}

// AddUser : add new record
func (m *MysqlDriver) AddUser(user *model.User) error {
	return m.conn.Create(user).Error
}

// UpdateUser : update user
func (m *MysqlDriver) UpdateUser(user *model.User) error {
	originUser, err := m.GetUserByID(user.ID)
	if err != nil {
		return err
	}
	return m.conn.Model(originUser).Updates(*user).Error
}
