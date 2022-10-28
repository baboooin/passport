package main

import (
	"time"
)

type UserRPC int

type Timestamp time.Time

type ErrorReply struct {
	Code    int
	Message string
}

type SuccesReply struct {
	Code    int
	Message string
	Token   string
}

type Reply struct {
	Code     int       `json:"code"`     //
	Message  string    `json:"message"`  //
	Password string    `json:"password"` //
	Expire   time.Time `json:"expire"`   //
	Token    string    `json:"token"`
}

type Role int64

const (
	Au Role = iota
	Wu
	Du
	Su
)

type TokenRecord struct {
	UId string `json:"id"`
}

type UserCompanyLink struct {
	CId  string `json:"cid"`
	UId  string `json:"uid"`
	Role Role   `json:"role"`
}

type UserWarehouseLink struct {
	CId  string `json:"cid"`
	WId  string `json:"wid"`
	Role Role   `json:"role"`
}

type CompanyRecord struct {
	CId    string `json:"cid"`
	INN    string `json:"inn"`
	KPP    string `json:"kpp"`
	Name   string `json:"name"`
	Parent string `json:"parent"`
}

type GeoPoint struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type WarehouseRecord struct {
	WId     string   `json:"wid"`
	CId     string   `json:"cid"`
	Name    string   `json:"name"`
	Point   GeoPoint `json:"geoPoint"`
	Country string   `json:"country"`
	County  string   `json:"county"`
	City    string   `json:"city"`
	House   string   `json:"house"`
}

type UserRecord struct {
	UId       string `json:"uid"`
	FistsName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	Hash      string `json:"hash"`
}

type LoginRec struct {
	Login    string
	Password string
}

type MobileLoginRec struct {
	Mobile string
	Code   string
}

type RegisterRec struct {
	FistsName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	TaxNum    string `json:"taxnum"`
}

type Truck struct {
	TId string `json:"tid"`
}
