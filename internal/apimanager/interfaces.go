package apimanager

import (
	"github.com/s-yakubovskiy/whereami/pkg/ipapi"
	"github.com/s-yakubovskiy/whereami/pkg/ipqualityscore"
)

type IpApiInterface interface {
	Lookup(ip string) (*ipapi.IpApiLocation, error)
}

type IpQualityInterface interface {
	Lookup(ip string) (*ipqualityscore.IpQualityScoreLocation, error)
}

// type Ip
