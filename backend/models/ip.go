package models

import "net"

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
