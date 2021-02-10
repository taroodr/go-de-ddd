package repositories

import (
	"github.com/shimar/go-de-ddd/internal/app/domain/entities"
	"github.com/shimar/go-de-ddd/internal/app/infra/notifications"
	"gopkg.in/gorp.v1"
)

type userRepository struct {
	db gorp.SqlExecutor
}

func (u *userRepository) Save(user entities.User)  {
	// 通知オブジェクトを生成
	userDataModelBuilder := &notifications.UserDataModelBuilder{}
	// 通知オブジェクトを引き渡し内部データを取得
	user.Notify(userDataModelBuilder)

	// 通知された内部データからデータモデルを生成
	userDataModel := userDataModelBuilder.Build()

	u.db.Insert(userDataModel)
}