package services

import (
	"github.com/jinzhu/copier"

	"digin/adapters/ipinfo"
	"digin/dtos"
)

type PublicIPService interface {
	GetPublicIP() *dtos.GetPublicIPResponse
}

type publicIPService struct {
	ipInfoAdapter ipinfo.Adapter
}

func NewPublicIPService(ipInfoAdapter ipinfo.Adapter) PublicIPService {
	return &publicIPService{ipInfoAdapter: ipInfoAdapter}
}

func (_this *publicIPService) GetPublicIP() *dtos.GetPublicIPResponse {
	var (
		ipInfo = _this.ipInfoAdapter.GetIPInfo()
		resp   dtos.GetPublicIPResponse
		_      = copier.Copy(&resp, ipInfo)
	)
	return &resp
}
