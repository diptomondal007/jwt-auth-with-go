package psql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
	"jwtauthwithgo/auth"
	"jwtauthwithgo/auth/models"
	"log"
)

type psqlRepository struct {
	db *gorm.DB
}

func (r psqlRepository) Login(username string) (*models.User, error) {
	var user models.User
	err := r.db.Find(&user,"username=?",username).Error
	if err != nil{
		return nil, err
	}
	return &user, nil
}

func (r psqlRepository) Register(user *models.User) error {
	pass := user.Password
	passToByte := []byte(pass)
	hashedPassword, _ := bcrypt.GenerateFromPassword(passToByte, bcrypt.MinCost)
	r.db.Create(&models.User{
		Username:  user.Username,
		Password:  string(hashedPassword),
		CreatedAt: user.CreatedAt,
	})
	return nil
}

func NewPsqlRepository(connection string) (auth.Repository, error) {
	repo := &psqlRepository{}
	db, err := gorm.Open("postgres", connection)
	if err != nil{
		return nil, err
	}else {
		log.Println("Connected To Database!")
	}
	db.AutoMigrate(&models.User{})
	repo.db = db
	return repo, nil
}
