package entity

import "time"

type Base struct {
	ID        int64     `bson:"_id" json:"id,string"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	DeletedAt int64     `bson:"deleted_at" json:"deleted_at"`
}
