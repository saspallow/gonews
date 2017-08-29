package model

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"sync"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// News Type
type News struct {
	ID       bson.ObjectId `bson:"_id"`
	Title    string
	Image    string
	Detail   string
	CreateAt time.Time `bson:"createAt"`
	UpdateAt time.Time `bson:"updateAt"`
}

var (
	newsStorage []News
	mutexNews   sync.RWMutex
)

func generateID() string {
	buf := make([]byte, 16)
	rand.Read(buf)
	return base64.StdEncoding.EncodeToString(buf)
}

// CreateNews inserts news into database
func CreateNews(news News) error {
	news.ID = bson.NewObjectId()
	news.CreateAt = time.Now()
	news.UpdateAt = news.CreateAt

	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("news").Insert(&news)
	if err != nil {
		return err
	}
	return nil
}

// ListNews lists all news from database
func ListNews() ([]*News, error) {
	s := mongoSession.Copy()
	defer s.Close()
	var news []*News
	err := s.DB(database).C("news").Find(nil).All(&news)
	if err != nil {
		return nil, err
	}
	return news, nil
}

// GetNews retrives new from database
func GetNews(id string) (*News, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, fmt.Errorf("Invalid ID")
	}
	objectID := bson.ObjectIdHex(id)
	s := mongoSession.Copy()
	defer s.Close()
	var n News
	err := s.DB(database).C("news").FindId(objectID).One(&n)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

// DeleteNews detektes a news from database
func DeleteNews(id string) error {
	objectID := bson.ObjectId(id)
	if !objectID.Valid() {
		return fmt.Errorf("Invalid ID")
	}

	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("news").RemoveId(objectID)
	if err != nil {
		return err
	}
	return nil
}
