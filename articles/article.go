package articles

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Article represents an article.
type Article struct {
	// ID is unique id of article.
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`

	// Human readable & SEO friendly URL path of article.
	Path string `json:"path" bson:"path"`

	// Title of the article.
	Title string `json:"title"  bson:"title" validate:"required"`

	// Content of the article.
	Content string `json:"content" bson:"content" validate:"required"`

	// Location info of where article is posted from.
	Location Location `json:"location" bson:"location,omitempty"`

	// CreatedAt is the creation time of article.
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

// Location info of where article is posted from.
type Location struct {
	City    string `json:"city" bson:"city"`
	Country string `json:"country" bson:"country"`
}
