package model

import (

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string             `bson:"Username,omitempty" json:"Username,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
}

type Registrasi struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NamaLengkap  string             `bson:"namalengkap,omitempty" json:"namalengkap,omitempty"`
	NomorHP      string             `bson:"nomorhp,omitempty" json:"nomorhp,omitempty"`
	TanggalLahir string             `bson:"tanggallahir,omitempty" json:"tanggallahir,omitempty"`
	Alamat       string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	NIM          string             `bson:"nim,omitempty" json:"nim,omitempty"`
}

type Reservasi struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama 		  string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Notelp		  string             `bson:"notelp,omitempty" json:"notelp,omitempty"`
	TTL     	  string             `bson:"ttl,omitempty" json:"ttl,omitempty"`
	Status	      string             `bson:"status,omitempty" json:"status,omitempty"`
	Keluhan	      string             `bson:"keluhan,omitempty" json:"keluhan,omitempty"`
}

type Response struct {
	Status  bool   `json:"status" bson:"status"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type ReservasiResponse struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message,omitempty" bson:"message,omitempty"`
	Data    Reservasi   `json:"data" bson:"data"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	Data    User   `bson:"data,omitempty" json:"data,omitempty"`
}


