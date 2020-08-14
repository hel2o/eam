package model

import "time"

type EAM struct {
	Id          int       `json:"id" xorm:"notnull pk autoincr INT(11)"`
	Code        string    `json:"code" xorm:"notnull default '' VARCHAR(11)"`
	Level       string    `json:"level" xorm:"notnull default '' VARCHAR(11)"`
	Ip          string    `json:"ip" xorm:"notnull default '' VARCHAR(64)"`
	Name        string    `json:"name" xorm:"notnull default '' VARCHAR(64)"`
	Place       string    `json:"place" xorm:"notnull default '' VARCHAR(64)"`
	Company     string    `json:"company" xorm:"notnull default '' VARCHAR(64)"`
	Department  string    `json:"department" xorm:"notnull default '' VARCHAR(64)"`
	Os          string    `json:"os" xorm:"notnull default '' VARCHAR(64)"`
	AppService  string    `json:"appService" xorm:"notnull default '' VARCHAR(64)"`
	SqlType     string    `json:"sqlType" xorm:"notnull default '' VARCHAR(64)"`
	ServerModel string    `json:"serverModel" xorm:"notnull default '' VARCHAR(64)"`
	ServerPara  string    `json:"serverPara" xorm:"notnull default '' VARCHAR(64)"`
	Developer   string    `json:"developer" xorm:"notnull default '' VARCHAR(64)"`
	ServerCode  string    `json:"serverCode" xorm:"notnull default '' VARCHAR(64)"`
	PD          string    `json:"PD" xorm:"notnull default '' VARCHAR(64)"`
	Person      string    `json:"person" xorm:"notnull default '' VARCHAR(64)"`
	IsVirtual   string    `json:"isVirtual" xorm:"notnull default '' VARCHAR(12)"`
	RealHost    string    `json:"realHost" xorm:"notnull default '' VARCHAR(64)"`
	Remarks     string    `json:"remarks" xorm:"notnull default '' VARCHAR(128)"`
	UpdatedAt   time.Time `json:"updateAt" xorm:"updated"`
	CreatedAt   time.Time `json:"createAt" xorm:"created"`
}
