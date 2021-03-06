package eip

import (
	qcservice "github.com/yunify/qingcloud-sdk-go/service"
)

type EIP struct {
	Name        string
	ID          string
	Address     string
	Status      string
	Bandwidth   int
	BillingMode string
}

type EIPHelper interface {
	GetEIPByID(id string) (*EIP, error)
	GetEIPByAddr(addr string) (*EIP, error)
	ReleaseEIP(id string) error
	GetAvaliableOrAllocateEIP() (*EIP, error)
	AllocateEIP() (*EIP, error)
	GetAvaliableEIPs() ([]*EIP, error)
}

func (e *EIP) ToQingCloudEIP() *qcservice.EIP {
	return &qcservice.EIP{
		EIPID:   &e.ID,
		EIPAddr: &e.Address,
		EIPName: &e.Name,
		Status:  &e.Status,
	}
}
