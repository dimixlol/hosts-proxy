package models

import (
	"github.com/gofrs/uuid"
)

type URL struct {
	BaseModel
	Slug   string `gorm:"type:varchar(20);unique;not null" json:"slug"`
	HostID uuid.UUID
	Host   *Host `json:"host"`
	IPID   uuid.UUID
	IP     *IP `json:"ip"`
}

func (u *URL) GetIP() string {
	return u.IP.IP
}

func (u *URL) GetHost() string {
	return u.Host.Host
}

func (u *URL) GetSlug() string {
	return u.Slug
}
