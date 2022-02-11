package repositories

import (
	"restapi/src/dtos"
	"restapi/src/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRepository is struct
type UserRepository struct {
	DBConnection *gorm.DB
}

func NewUserRepository(dbConnection *gorm.DB) *UserRepository {
	return &UserRepository{
		DBConnection: dbConnection,
	}
}

// VerifyCredential is function
func (r *UserRepository) VerifyCredential(email string, password string) interface{} {
	// Get the first record ordered by primary key
	// db.First(&user)
	// SELECT * FROM users ORDER BY id LIMIT 1;

	// Get one record, no specified order
	// db.Take(&user)
	// SELECT * FROM users LIMIT 1;

	// Get last record, ordered by primary key desc
	// db.Last(&user)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	var user models.User
	res := r.DBConnection.Model(&user).
		Where("email = ?", email).
		First(&user)

	if res.Error == nil && checkPassword(user.Password, password) {
		return user
	}

	return nil
}

// CreateUser is function
func (r *UserRepository) CreateUser(input dtos.UserCreateInput) (*models.User, error) {

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashAndSalt([]byte(input.Password)),
	}

	result := r.DBConnection.Model(&models.User{}).Create(&user)

	if result.RowsAffected == 1 && result.Error == nil {
		return &user, nil
	}

	return nil, result.Error
}

// hashAndSalt is function hash password
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic("Failed to hash a password")
	}
	return string(hash)
}

// checkPassword is function
func checkPassword(hashPwd string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))
	return err == nil
}
