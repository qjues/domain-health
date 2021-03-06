package model

import (
	"encoding/json"
	"time"
)

type CertInfo struct {
	ExpireTime int64
	CommonName string
}

type Domain struct {
	Address       string
	LastCheckTime int64
	CheckError    error `json:"-"`
	CheckErrorStr string
	OriginError   error `json:"-"`
	CreatedTime   int64
	From          string
	CertInfo      CertInfo
}

func NewDomain() (d *Domain) {
	return &Domain{
		CreatedTime: time.Now().Unix(),
	}
}

func NewDomainFromBytes(b []byte) (d *Domain) {
	d = &Domain{}
	err := d.UnmarshalFrom(b)
	if err != nil {
		return nil
	}

	return
}

func (d *Domain) UnmarshalFrom(b []byte) error {
	return json.Unmarshal(b, d)
}

func (d *Domain) Marshal() []byte {
	if d.CheckErrorStr == "" && d.CheckError != nil {
		d.CheckErrorStr = d.CheckError.Error()
	}
	bytes, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return bytes
}
