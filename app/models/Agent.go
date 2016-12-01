package models

import (
	"github.com/revel/revel"
	"regexp"
)

type AgentTravel struct {
	IdAgent   int64          `gorm:"primary_key;AUTO_INCREMENT"`
	NamaAgent string         `gorm:"column:nama_agent;type:varchar(45)"`
	Alamat    string         `gorm:"column:alamat_agent;type:varchar(255)"`
	Notelp    string         `gorm:"column:notelp_agent;type:varchar(12)"`
	Email     string         `gorm:"column:email_agent;type:varchar(45);unique"`
	Website   string         `gorm:"column:website_agent;type:varchar(45);unique"`
	UserId    int64          `gorm:"column:user_id;index"`
	Service   []AgentService `gorm:"ForeignKey:IdAgent"`
}

type AgentService struct {
	IdService int64          `gorm:"primary_key;AUTO_INCREMENT"; json:"id, omitempty"`
	IdAgent   int64          `gorm:"column:agent_id;index"; json:"idAgent, omitempty"`
	Service   string         `gorm:"column:service;type:varchar(45)"; json:"service, omitempty"`
	Kategori  string         `gorm:"column:kategori;type:varchar(45)"; json:"kategori, omitempty"`
	Price     float32        `gorm:"column:price"; json:"price, omitempty"`
	Desc      string         `gorm:"column:desc;type:text"; json:"desc, omitempty"`
	Foto      FotoService    `gorm:"ForeignKey:IdService"`
	AddOn     []AddOnService `gorm:"ForeignKey:IdService"`
	Status    bool           `gorm:"column:status;type:bool"`
}

type AddOnService struct {
	IdAddOnService int64   `gorm:"primary_key;AUTO_INCREMENT"`
	IdService      int64   `gorm:"column:service_id;index"`
	Service        string  `gorm:"column:service;type:varchar(45)"`
	Price          float32 `gorm:"column:price"`
}

type FotoService struct {
	IdFoto    int64  `gorm:"primary_key;AUTO_INCREMENT"; json:"id, omitempty"`
	Foto      string `json:"foto, omitempty"`
	Width     int    `json:"width, omitempty"`
	Height    int    `json:"height, omitempty"`
	Size      int    `json:"size, omitempty"`
	Format    string `json:"format, omitempty"`
	IdService int64  `gorm:"column:service_id;index"`
	Dir       string `json:"dir, omitempty"`
}

var emailreg = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

//var urlreg = regexp.MustCompile(`/^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$/ `)

func (at AgentTravel) Validation(r *revel.Validation) {
	r.Check(at.NamaAgent,
		revel.Required{},
		revel.MinSize{4},
		revel.MaxSize{25},
	)

	r.Check(at.Alamat, revel.Required{})
	r.Check(at.Notelp,
		revel.Required{},
		revel.MinSize{8},
		revel.MaxSize{12},
	)

	r.Check(at.Email,
		revel.Required{},
		revel.MinSize{4},
		revel.Match{emailreg},
	)
	r.Check(at.Website,
		revel.MinSize{6},
		//		revel.Match{urlreg},
	)

}
