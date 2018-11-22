package articles

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Storage is a data storage to keep articles.
type Storage interface {
	// Save saves an article to storage.
	Save(*Article) error

	// Get gets an article by id.
	Get(id string) (*Article, error)

	// List returns all articles from storage.
	List() ([]*Article, error)

	// Close closes underlying database connection.
	Close() error
}

// MongoStorage implements Storage.
type MongoStorage struct {
	// articleC is articles collection.
	articleC string

	//dbName is database name.
	dbName string

	ss *mgo.Session
}

// NewMongoStorage returns a new MongoDB storage with given MongoDB address and database.
func NewMongoStorage(addr, dbName string) (Storage, error) {
	s := &MongoStorage{
		articleC: "article",
		dbName:   dbName,
	}

	// dial to mongo.
	var err error
	s.ss, err = mgo.Dial(addr)
	if err != nil {
		return nil, err
	}

	// set consistency mode.
	s.ss.SetMode(mgo.Strong, true)

	return s, nil
}

// Save saves an article to storage.
func (s *MongoStorage) Save(a *Article) error {
	ss := s.ss.Clone()
	defer ss.Close()
	return ss.DB(s.dbName).C(s.articleC).Insert(a)
}

// Get gets an article by id.
func (s *MongoStorage) Get(id string) (*Article, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("not found")
	}
	ss := s.ss.Clone()
	defer ss.Close()
	var a *Article
	return a, ss.DB(s.dbName).C(s.articleC).Find(bson.M{"_id": bson.ObjectIdHex("id")}).One(&a)
}

// List returns all articles from storage.
func (s *MongoStorage) List() ([]*Article, error) {
	ss := s.ss.Clone()
	defer ss.Close()
	var articles []*Article
	return articles, ss.DB(s.dbName).C(s.articleC).Find(bson.M{}).All(&articles)
}

// Close closes underlying database connection.
func (s *MongoStorage) Close() error {
	s.ss.Close()
	return nil
}
