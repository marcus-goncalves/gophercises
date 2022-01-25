package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	UserNotFound  = errors.New("models: Resources not found")
	InvalidId     = errors.New("models: ID provided is invalid")
	InvalidPwd    = errors.New("models: incorrect user or pwd provided")
	userPwdPepper = "mqtoc"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService() (*UserService, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "photo_app:photo_app_pass@tcp(127.0.0.1:3306)/photo_gallerie?charset=utf8mb4&parseTime=True&loc=Local",
	}))
	if err != nil {
		return nil, err
	}

	return &UserService{
		db: db,
	}, nil

}

func (us *UserService) AutoMigrate() error {
	err := us.db.Migrator().AutoMigrate(&User{})
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) DestructiveReset() error {
	err := us.db.Migrator().DropTable(&User{})
	if err != nil {
		return err
	}

	return us.AutoMigrate()
}

func (us *UserService) Close() {
	us.Close()
}

type User struct {
	gorm.Model
	Name      string
	Email     string `gorm:"not null; unique_index"`
	Password  string `gorm:"-"`
	HashedPwd string `gorm:"not null"`
}

func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return UserNotFound
	}

	return err
}

func (us *UserService) ByID(id uint) (*User, error) {
	var user User

	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *UserService) ByEmail(email string) (*User, error) {
	var user User

	db := us.db.Where("email = ?", email)
	err := first(db, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *UserService) Create(user *User) error {
	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password+userPwdPepper), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.HashedPwd = string(hashedBytes)
	user.Password = ""
	return us.db.Create(&user).Error
}

func (us *UserService) Update(user *User) error {
	return us.db.Save(&user).Error
}

func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return InvalidId
	}

	user := User{Model: gorm.Model{ID: id}}

	return us.db.Delete(&user).Error
}

func (us *UserService) Authenticate(email, pwd string) (*User, error) {
	foundUser, err := us.ByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.HashedPwd), []byte(pwd+userPwdPepper))
	switch err {
	case nil:
		return foundUser, nil
	case bcrypt.ErrMismatchedHashAndPassword:
		return nil, InvalidPwd
	default:
		return nil, err
	}
}
