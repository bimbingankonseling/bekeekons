package module

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/badoux/checkmail"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/argon2"
)

// var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(MongoString, dbname string) *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv(MongoString)))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

// crud
func GetAllDocs(db *mongo.Database, col string, docs interface{}) interface{} {
	collection := db.Collection(col)
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error GetAllDocs %s: %s", col, err)
	}
	err = cursor.All(context.TODO(), &docs)
	if err != nil {
		return err
	}
	return docs
}

func InsertOneDoc(db *mongo.Database, col string, doc interface{}) (insertedID primitive.ObjectID, err error) {
	result, err := db.Collection(col).InsertOne(context.Background(), doc)
	if err != nil {
		return insertedID, fmt.Errorf("kesalahan server : insert")
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func UpdateOneDoc(id primitive.ObjectID, db *mongo.Database, col string, doc interface{}) (err error) {
	filter := bson.M{"_id": id}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, bson.M{"$set": doc})
	if err != nil {
		return fmt.Errorf("error update: %v", err)
	}
	if result.ModifiedCount == 0 {
		err = fmt.Errorf("tidak ada data yang diubah")
		return
	}
	return nil
}

func DeleteOneDoc(_id primitive.ObjectID, db *mongo.Database, col string) error {
	collection := db.Collection(col)
	filter := bson.M{"_id": _id}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

// signup
func SignUpRegistrasi(db *mongo.Database, insertedDoc model.Registrasi) error {
	objectId := primitive.NewObjectID() 
	if insertedDoc.NamaLengkap == "" || insertedDoc.NomorHP == "" || insertedDoc.TanggalLahir == "" || insertedDoc.Alamat == "" || insertedDoc.NIM == "" {
		return fmt.Errorf("Dimohon untuk melengkapi data")
	} 
	hashedPassword := argon2.IDKey([]byte(insertedDoc.Akun.Password))
	user := bson.M{
		"_id": objectId,
		"username": insertedDoc.Akun.Email,
		"password": hex.EncodeToString(hashedPassword),
	}
	registrasi := bson.M{
		"namalengkap": insertedDoc.NamaLengkap,
		"nomorhp": insertedDoc.NomorHP,
		"tanggallahir": insertedDoc.TanggalLahir,
		"alamat": insertedDoc.Alamat,
		"nim": insertedDoc.NIM,
	}
	_, err = InsertOneDoc(db, "user", user)
	if err != nil {
		return fmt.Errorf("kesalahan server")
	}
	_, err = InsertOneDoc(db, "registrasi", registrasi)
	if err != nil {
		return fmt.Errorf("kesalahan server")
	}
	return nil
}

// login
func LogIn(db *mongo.Database, insertedDoc model.User) (user model.User, err error) {
	if insertedDoc.Username == "" || insertedDoc.Password == "" {
		return user, fmt.Errorf("Dimohon untuk melengkapi data")
	} 
	if err = checkmail.ValidateFormat(insertedDoc.Username); err != nil {
		return user, fmt.Errorf("Username tidak valid")
	} 
	existsDoc, err := GetUserFromEmail(insertedDoc.Username, db)
	if err != nil {
		return 
	}
	salt, err := hex.DecodeString(existsDoc.Salt)
	if err != nil {
		return user, fmt.Errorf("kesalahan server : salt")
	}
	hash := argon2.IDKey([]byte(insertedDoc.Password))
	if hex.EncodeToString(hash) != existsDoc.Password {
		return user, fmt.Errorf("password salah")
	}
	return existsDoc, nil
}

//Reservasi
func InsertReservasi(iduser primitive.ObjectID, db *mongo.Database, insertedDoc model.Reservasi) error {
	if insertedDoc.Nama == "" || insertedDoc.Notelp == "" || insertedDoc.TTL == "" || insertedDoc.Status == "" || insertedDoc.Keluhan == "" {
		return fmt.Errorf("mohon untuk melengkapi data")
	}

	reser := bson.M{
		"nama":    insertedDoc.Nama,
		"notelp":   insertedDoc.Notelp,
		"ttl":   insertedDoc.TTL,
		"status":        insertedDoc.Status,
		"keluhan":        insertedDoc.Keluhan,
	}

	_, err := InsertOneDoc(db, "reservasi", reser)
	if err != nil {
		return fmt.Errorf("error saat menyimpan data reservasi: %s", err)
	}
	return nil
}

func UpdateReservasi(idparam, iduser primitive.ObjectID, db *mongo.Database, insertedDoc model.Reservasi) error {
	_, err := GetReservasiFromID(idparam, db)
	if err != nil {
		return err
	}
	if insertedDoc.Nama == "" || insertedDoc.Notelp == "" || insertedDoc.TTL == "" || insertedDoc.Status == "" || insertedDoc.Keluhan == "" {
		return fmt.Errorf("mohon untuk melengkapi data")
	}
	reser := bson.M{
		"nama":    insertedDoc.Nama,
		"notelp":   insertedDoc.Notelp,
		"ttl":   insertedDoc.TTL,
		"status":        insertedDoc.Status,
		"keluhan":        insertedDoc.Keluhan,
	}

	err = UpdateOneDoc(idparam, db, "reservasi", reser)
	if err != nil {
		return err
	}
	return nil
}


func DeleteReservasi(idparam, iduser primitive.ObjectID, db *mongo.Database) error {
	_, err := GetReservasiFromID(idparam, db)
	if err != nil {
		return err
	}
	err = DeleteOneDoc(idparam, db, "reservasi")
	if err != nil {
		return err
	}
	return nil
}

func GetAllReservasi(db *mongo.Database) (tiket []model.Reservasi, err error) {
	collection := db.Collection("reservasi")
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return reservasi, fmt.Errorf("error GetAllReservasi mongo: %s", err)
	}
	err = cursor.All(context.TODO(), &tiket)
	if err != nil {
		return reservasi, fmt.Errorf("error GetAllReservasi context: %s", err)
	}
	return reservasi, nil
}


func GetReservasiFromID(_id primitive.ObjectID, db *mongo.Database) (doc model.Reservasi, err error) {
	collection := db.Collection("reservasi")
	filter := bson.M{"_id": _id}
	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return doc, fmt.Errorf("_id tidak ditemukan")
		}
		return doc, fmt.Errorf("kesalahan server")
	}
	return doc, nil
}
