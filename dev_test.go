package HealHero

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/argon2"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

var db = module.MongoConnect("MONGOSTRING", "keekons")

func TestGetUserFromEmail(t *testing.T) {
	username := "ardvprw"
	hasil, err := module.GetUserFromEmail(username, db)
	if err != nil {
		t.Errorf("Error TestGetUserFromEmail: %v", err)
	} else {
		fmt.Println(hasil)
	}
}

// Insert-Tiket
func TestInsertOneReservasi(t *testing.T) {
	var doc model.Reservasi
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

type Userr struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string             `bson:"username,omitempty" json:"username,omitempty"`
	Role  string             `bson:"role,omitempty" json:"role,omitempty"`
}

func TestGetAllDoc(t *testing.T) {
	hasil := module.GetAllDocs(db, "user", []Userr{})
	fmt.Println(hasil)
}

func TestInsertUser(t *testing.T) {
	var doc model.User
	doc.Username = "Gabriella"
	password := "bela123"
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		t.Errorf("kesalahan server : salt")
	} else {
		hashedPassword := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
		user := bson.M{
			"username":    doc.Username,
			"password": hex.EncodeToString(hashedPassword),
		}
		_, err = module.InsertOneDoc(db, "user", user)
		if err != nil {
			t.Errorf("gagal insert")
		} else {
			fmt.Println("berhasil insert")
		}
	}
}

func TestSignUpRegistrasi(t *testing.T) {
	var doc model.Registrasi
	doc.NamaLengkap = "Gabriella"
	doc.NomorHP = "6287825683284"
	doc.TanggalLahir = "25 Mei 2003"
	doc.Alamat = "Wastukencana Blok 7 No 11"
	doc.NIM = "1214027"
	err := module.SignUpRegistrasi(db, doc)
	if err != nil {
		t.Errorf("Error inserting document: %v", err)
	} else {
		fmt.Println("Data berhasil disimpan dengan nama :", doc.NamaLengkap)
	}
}

func TestLogIn(t *testing.T) {
	var doc model.User
	doc.Username = "Gabriela"
	doc.Password = "gabril12"
	user, err := module.LogIn(db, doc)
	if err != nil {
		t.Errorf("Error getting document: %v", err)
	} else {
		fmt.Println("Selamat datang Driver:", user)
	}
}

func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := module.GenerateKey()
	fmt.Println("ini private key :", privateKey)
	fmt.Println("ini public key :", publicKey)
	id := "6569a026a943657839880665"
	objectId, err := primitive.ObjectIDFromHex(id)
	role := "registrasi"
	if err != nil {
		t.Fatalf("error converting id to objectID: %v", err)
	}
	hasil, err := module.Encode(objectId, role, privateKey)
	fmt.Println("ini hasil :", hasil, err)
}


func TestWatoken(t *testing.T) {
	body, err := module.Decode("fca3dbba6c382d6e937d33837f7428c1211e01a9928cbbbc0b86bb8351c02407", " v4.public.eyJleHAiOiIyMDIzLTEyLTAxVDE4OjU4OjE1KzA4OjAwIiwiaWF0IjoiMjAyMy0xMi0wMVQxNjo1ODoxNSswODowMCIsImlkIjoiNjU1YzNiOWExZDY1MjRmMmYxMjAwZmM2IiwibmJmIjoiMjAyMy0xMi0wMVQxNjo1ODoxNSswODowMCIsInJvbGUiOiJwZW5nZ3VuYSJ9GIKgKcp8gj4lzPH_NFvpx3GR2kBZ2qsDquYMKQdQ1PFpvHKlDy-FeO1umIGCaMuYyACP5jd-Y0at1WCOrsNRCA")
	fmt.Println("isi : ", body, err)
}

// test Tiket
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
