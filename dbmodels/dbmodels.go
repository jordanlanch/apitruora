package dbmodels

import (
	"time"
)


type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"-"`
}

type Response struct {
	Model
	Servers       []Servers `json:"servers" gorm:"many2many:order_servers"`
	ServersChanged   bool   `json:"servers_changed"`
	SslGrade         string `json:"ssl_grade"`
	PreviousSslGrade string `json:"previous_ssl_grade"`
	Logo             string `json:" logo"`
	IsDown           bool   `json:"is_down"`
}

type Servers struct {
	Model
	Address  string `json:"address"`
	SslGrade string `json:"ssl_grade"`
	Country  string `json:"country"`
	Owner    string `json:"owner"`
} 