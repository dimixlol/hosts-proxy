package models

import (
	"fmt"
	"github.com/gofrs/uuid"
	"golang.org/x/net/idna"
	"net"
	"strings"
	"time"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Host struct {
	BaseModel
	Host  string `gorm:"not null" json:"host"`
	Times int    `gorm:"default:0"`
}

func (h *Host) IncrementTimes() {
	h.Times++
}

func NewHost(host string) (*Host, error) {
	h, err := idna.ToASCII(host)
	if err != nil {
		return nil, err
	}
	if strings.HasSuffix(h, ".") {
		h = h[:len(h)-1]
	}

	if parts := strings.Split(h, "."); len(parts) < 2 {
		return nil, fmt.Errorf("host must have at least 2 parts")
	}
	return &Host{Host: h}, nil
}

type IP struct {
	BaseModel
	IP    string `gorm:"not null" json:"ip"`
	Times int    `gorm:"default:0"`
}

func (ip *IP) IncrementTimes() {
	ip.Times++
}

func NewIP(ip string) *IP {
	if parsedIP := net.ParseIP(ip); parsedIP != nil {
		return &IP{IP: ip}
	}
	return nil
}

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