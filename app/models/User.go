package models

import (
	"github.com/revel/revel"
	"regexp"
	"time"
)

type User struct {
	UID            int64       `gorm:"column:user_id;primary_key;AUTO_INCREMENT"`
	Nama           string      `gorm:"column:nama;type:varchar(45)"`
	Email          string      `gorm:"column:email;type:varchar(40);unique"`
	Username       string      `gorm:"column:username;type:varchar(40);unique"`
	HashedPassword []byte      `gorm:"column:password"`
	FbId           string      `gorm:"column:fbid;"`
	GplusId        string      `gorm:"column:gplusid;"`
	TwitId         string      `gorm:"column:twitid"`
	Verify         bool        `gorm:"column:verify;type:bool"`
	Agent          AgentTravel `gorm:"ForeignKey:UserId"`
	TokenUser      UserToken   `gorm:"ForeignKey:UserId"`
	Foto           UserFoto    `gorm:"ForeignKey:UserId"`
}

type UserToken struct {
	IdToken     int64  `gorm:"primary_key;AUTO_INCREMENT"`
	AccessToken string `gorm:"type:varchar(255)"`
	Expiry      time.Time
	Used        bool  `gorm:"type:bool"`
	UserId      int64 `gorm:"column:user_id;index"`
}

type UserFoto struct {
	IdFoto int64 `gorm:"primary_key;AUTO_INCREMENT"`
	Foto   string
	Width  int
	Height int
	Size   int
	Format string
	UserId int64  `gorm:"column:user_id;index"`
	Dir    string `json:"dir, omitempty"`
}

/*
type UserSosmed struct {
	IDUsersos int    `db:"idUsersosmed, primarykey, autoincrement"`
	FbId      string `db:"fbId"`
	GplusId   string `db:"gplusId"`
	TwitId    string `db:"twitId"`
	UserId    int    `db:"idUser"`
}
*/

var emailregex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

//var userregex = regexp.MustCompile(`^\\w*$`)
func (user *User) Validation(r *revel.Validation) {
	r.Check(user.Nama,
		revel.Required{},
		revel.MinSize{4},
		revel.MaxSize{20},
	)
	r.Check(user.Username,
		revel.Required{},
		revel.MinSize{4},
		revel.MaxSize{20},

		//revel.Match{userregex},
	)
	r.Check(user.Email,
		revel.Required{},
		revel.MinSize{4},
		revel.Match{emailregex},
	)

}
