package models

import (
	"fmt"
	"golang.org/x/net/idna"
	"strings"
)

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
