package keekonseling

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/bimbingankonseling/bekeekons/model"
	"github.com/bimbingankonseling/bekeekons/module"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/argon2"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

var db = module.MongoConnect("MONGOSTRING", "keekons")

func TestGeneratePasswordHash(t *testing.T) {
	password := "bellaa"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity
	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}
func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("gabril", privateKey)
	fmt.Println(hasil, err)
}

func TestHashFunction(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "keekons")
	var userdata User
	userdata.Username = "gabril"
	userdata.Password = "bellaa"

	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mconn, "user", filter)
	fmt.Println("Mongo User Result: ", res)
	hash, _ := HashPassword(userdata.Password)
	fmt.Println("Hash Password : ", hash)
	match := CheckPasswordHash(userdata.Password, res.Password)
	fmt.Println("Match:   ", match)

}

func TestIsPasswordValid(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "keekons")
	var userdata User
	userdata.Username = "gabril"
	userdata.Password = "bellaa"

	anu := IsPasswordValid(mconn, "user", userdata)
	fmt.Println(anu)
}

func TestInsertUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "keekons")
	var userdata User
	userdata.Username = "gabril"
	userdata.Password = "bellaa"

	nama := InsertUser(mconn, "user", userdata)
	fmt.Println(nama)
}

//Insert-Tiket

func TestInsertOneReservasi(t *testing.T) {
	var doc Reservasi
	doc.Nama = "Gabriella"
	doc.Notelp = "6287825683284"
	doc.TTL = "25 Mei 2023"
	doc.Status = "Mahasiswa"
	doc.Keluhan = "Cape"
	if doc.Nama == "" || doc.Notelp == "" || doc.TTL == "" || doc.Status == "" || doc.Keluhan == "" {
		t.Errorf("mohon untuk melengkapi data")
	} else {
		insertedID, err := module.InsertOneDoc(db, "reservasi", doc)
		if err != nil {
			t.Errorf("Error inserting document: %v", err)
			fmt.Println("Data tidak berhasil disimpan")
		} else {
			fmt.Println("Data berhasil disimpan dengan id :", insertedID.Hex())
		}
	}
}

func TestGetAllDoc(t *testing.T) {
	hasil := module.GetAllDocs(db, "user", []Userr{})
	fmt.Println(hasil)
}

// test reservasi
func TestInsertReservasi(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "keekons")
	payload, err := module.Decode("fca3dbba6c382d6e937d33837f7428c1211e01a9928cbbbc0b86bb8351c02407", "v4.public.eyJleHAiOiIyMDIzLTEyLTAxVDE4OjU4OjE1KzA4OjAwIiwiaWF0IjoiMjAyMy0xMi0wMVQxNjo1ODoxNSswODowMCIsImlkIjoiNjU1YzNiOWExZDY1MjRmMmYxMjAwZmM2IiwibmJmIjoiMjAyMy0xMi0wMVQxNjo1ODoxNSswODowMCIsInJvbGUiOiJwZW5nZ3VuYSJ9GIKgKcp8gj4lzPH_NFvpx3GR2kBZ2qsDquYMKQdQ1PFpvHKlDy-FeO1umIGCaMuYyACP5jd-Y0at1WCOrsNRCA")
	if err != nil {
		t.Errorf("Error decode token: %v", err)
	}
	// if payload.Role != "mitra" {
	// 	t.Errorf("Error role: %v", err)
	// }
	var datatiket model.Reservasi
	datatiket.Nama = "Bella"
	datatiket.Notelp = "6287815683284"
	datatiket.TTL = "25 Mei 2003"
	datatiket.Status = "Mahasiswa"
	datatiket.Keluhan = "Cape"
	err = module.InsertReservasi(payload.Id, conn, datatiket)
	if err != nil {
		t.Errorf("Error insert : %v", err)
	} else {
		fmt.Println("Success!!!")
	}
}

func TestUpdateReservasi(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "keekons")
	payload, err := module.Decode("fca3dbba6c382d6e937d33837f7428c1211e01a9928cbbbc0b86bb8351c02407", "v4.public.eyJleHAiOiIyMDIzLTEyLTAxVDE4OjU4OjE1KzA4OjAwIiwiaWF0IjoiMjAyMy0xMi0wMVQxNjo1ODoxNSswODowMCIsImlkIjoiNjU1YzNiOWExZDY1MjRmMmYxMjAwZmM2IiwibmJmIjoiMjAyMy0xMi0wMVQxNjo1ODoxNSswODowMCIsInJvbGUiOiJwZW5nZ3VuYSJ9GIKgKcp8gj4lzPH_NFvpx3GR2kBZ2qsDquYMKQdQ1PFpvHKlDy-FeO1umIGCaMuYyACP5jd-Y0at1WCOrsNRCA")
	if err != nil {
		t.Errorf("Error decode token: %v", err)
	}
	
	var datatiket model.Reservasi
	datatiket.Nama = "Bela"
	datatiket.Notelp = "6287815683284"
	datatiket.TTL = "25 Mei 2003"
	datatiket.Status = "Mahasiswa"
	datatiket.Keluhan = "Cape Banget"
	id := "6569a53d783c6970079a560b"
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to objectID: %v", err)
	}
	err = module.UpdateReservasi(objectId, payload.Id, conn, datatiket)
	if err != nil {
		t.Errorf("Error update : %v", err)
	} else {
		fmt.Println("Success!!!")
	}
}

func TestDeleteReservasi(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "keekons")
	payload, err := module.Decode("fca3dbba6c382d6e937d33837f7428c1211e01a9928cbbbc0b86bb8351c02407", "v4.public.eyJleHAiOiIyMDIzLTEyLTAxVDE4OjU4OjE1KzA4OjAwIiwiaWF0IjoiMjAyMy0xMi0wMVQxNjo1ODoxNSswODowMCIsImlkIjoiNjU1YzNiOWExZDY1MjRmMmYxMjAwZmM2IiwibmJmIjoiMjAyMy0xMi0wMVQxNjo1ODoxNSswODowMCIsInJvbGUiOiJwZW5nZ3VuYSJ9GIKgKcp8gj4lzPH_NFvpx3GR2kBZ2qsDquYMKQdQ1PFpvHKlDy-FeO1umIGCaMuYyACP5jd-Y0at1WCOrsNRCA")
	if err != nil {
		t.Errorf("Error decode token: %v", err)
	}
	id := "6569a53d783c6970079a560b"
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to objectID: %v", err)
	}
	err = module.DeleteReservasi(objectId, payload.Id, conn)
	if err != nil {
		t.Errorf("Error delete : %v", err)
	} else {
		fmt.Println("Success!!!")
	}
}

func TestGetAllReservasi(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "keekons")
	data, err := module.GetAllReservasi(conn)
	if err != nil {
		t.Errorf("Error get all : %v", err)
	} else {
		fmt.Println(data)
	}
}

func TestGetReservasiFromID(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "keekons")
	id := "6569a025a943657839880661"
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to objectID: %v", err)
	}
	reservasi, err := module.GetReservasiFromID(objectId, conn)
	if err != nil {
		t.Errorf("Error get Tiket : %v", err)
	} else {
		fmt.Println(reservasi)
	}
}
