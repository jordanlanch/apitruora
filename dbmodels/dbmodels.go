package dbmodels

import (
	"time"
)


type Model struct {
	ID        uint       `gorm:"primary_key" json:"-"`
	CreatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type Response struct {
	Model
	Servers       []Servers `json:"servers" gorm:"PRELOAD:true;foreignkey:ID"`
	ServersChanged   bool   `json:"servers_changed"`
	SslGrade         string `json:"ssl_grade"`
	PreviousSslGrade string `json:"previous_ssl_grade"`
	Logo             string `json:" logo"`
	Title            string `json:" title"`
	IsDown           bool   `json:"is_down"`
}

type Servers struct {
	Model
	Address  string `json:"address"`
	SslGrade string `json:"ssl_grade"`
	Country  string `json:"country"`
	Owner    string `json:"owner"`
} 

type Items struct {
	Model
	Domain string `json:"domain"`
	Response []Response `json:"response" gorm:"PRELOAD:true;foreignkey:ID"`
}