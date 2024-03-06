package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// IUserUsecaseはユーザーに関するユースケースのインターフェース
type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

// NewUserUsecaseはIUserUsecaseを実装した構造体を返す
func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

/*
SignUp関数の概要
1. パスワードをハッシュ化
2. ユーザーを作成
3. レスポンスを返す 
*/
func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	/*
	 * bycryptとは、パスワードのハッシュ化を行うためのライブラリ
	 * GenerateFromPasswordの引数は、ハッシュ化したいパスワードとコストパラメータ
	*/
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{Email: user.Email, Password: string(hash)}
	// CreateUserメソッドを呼び出し、新しいユーザーを作成
	// CreateUserメソッドは、リポジトリ層のメソッドを呼び出している
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID: newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

/*
Login関数の概要
1. ユーザーを取得
2. パスワードを検証
3. JWTトークンを作成
*/
func (uu *userUsecase) Login(user model.User) (string, error) {
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	// expはトークンの有効期限を設定するためのもの
	// time.Now().Add(time.Hour * 12)は、現在時刻から12時間後を表す
	// トークンの役割は、認証されたユーザーを識別すること
	jst, err := time.LoadLocation("Asia/Tokyo")
    if err != nil {
        panic(err)
    }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp": time.Now().In(jst).Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}