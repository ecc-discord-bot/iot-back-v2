package models

type User struct {
	UserID     string `gorm:"primary_key"`
	DiscordId  string
	StudentsId string
	Name       string
	Class      string
	Signature  string
	NowTime    int64
	IsPaid     bool
	IsAgreed   bool
}

// ユーザーを取得する
func GetUser(userid string) (User, error) {
	var user User
	result := dbconn.Where(&User{UserID: userid}).First(&user)

	// エラー処理
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func CreateUser(user User) error {
	return dbconn.Create(&user).Error
}

func UpdateUser(user User) error {
	return dbconn.Save(&user).Error
}
