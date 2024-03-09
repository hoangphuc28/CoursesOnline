package mysql

import (
	"fmt"
	"github.com/hoangphuc28/CoursesOnline/Cart-Service/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(c *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Mysql.User,
		c.Mysql.Password,
		c.Mysql.Host,
		c.Mysql.Port,
		c.Mysql.DBName,
	)
	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return gormDb, nil
}
