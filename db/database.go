package db

import (
	"errors"
	"time"

	"come.archroid.pixelgolang/entities"
	"come.archroid.pixelgolang/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Init() {

	var err error
	Database, err = gorm.Open(sqlite.Open("./userdata/db.sqlite"), &gorm.Config{})

	if err != nil {
		panic("failed to create Database")
	}

	Database.AutoMigrate(&entities.User{}, &entities.Post{})
}
func CreateUser(username string, password string, email string) (entities.User, error) {

	//check for same email
	user := entities.User{
		Email: email,
	}
	if Database.Where(&user).Find(&user).RowsAffected != 0 {
		return user, errors.New("email already taken")
	}

	user = entities.User{
		Username:  username,
		Email:     email,
		Password:  password,
		AuthToken: utils.GenerateSecureToken(20),
	}

	Database.Create(&user)
	return user, nil
}

func LoginUser(email string, password string) (entities.User, error) {
	user := entities.User{
		Email: email,
	}

	if Database.Where(&user).Find(&user).RowsAffected == 0 {
		return user, errors.New("Email not found!")
	}

	if user.Password != password {
		return user, errors.New("Invalid password!")
	}

	return user, nil

}

func CreatePost(token string, imageurl string, caption string, id string) (entities.Post, error) {

	post := entities.Post{}

	// find user by token
	user := entities.User{
		AuthToken: token,
	}
	if Database.Where(&user).Find(&user).RowsAffected == 0 {
		return post, errors.New("Email not found!")
	}

	post = entities.Post{

		Email:       user.Email,
		Caption:     caption,
		Image_Url:   imageurl,
		Date:        time.Now().Unix(),
		Likes_Count: 0,
		ID:          id,
	}

	Database.Create(&post)

	return post, nil
}

func GetPostById(id string) (entities.Post, error) {
	post := entities.Post{
		ID: id,
	}

	if Database.Where(&post).Find(&post).RowsAffected == 0 {
		return post, errors.New("Post not found!")
	}

	return post, nil
}
