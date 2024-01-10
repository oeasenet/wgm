package wgm

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DefaultModel represents a default model for MongoDB documents.
type DefaultModel struct {
	Id             primitive.ObjectID `bson:"_id" json:"id"`
	CreateTime     int64              `bson:"create_time" json:"create_time"`
	LastModifyTime int64              `bson:"last_modify_time" json:"last_modify_time"`
}

// IDefaultModel represents an interface for a default model in MongoDB documents.
type IDefaultModel interface {
	ColName() string
	GetId() string
	GetObjectID() primitive.ObjectID
	PutId(id string)
	setDefaultCreateTime()
	setDefaultLastModifyTime()
	setDefaultId()
}

// ColName returns the name of the collection to which the DefaultModel belongs.
func (m *DefaultModel) ColName() string {
	return ""
}

// GetId returns the hexadecimal string representation of the Id field of the DefaultModel.
// It uses the Hex() method from the Id field to convert the Id to its string representation.
func (m *DefaultModel) GetId() string {
	return m.Id.Hex()
}

// PutId updates the "Id" field of the DefaultModel with the specified value.
// The input "id" is expected to be a hexadecimal string representing an ObjectId.
// If the conversion from the hexadecimal string to an ObjectId fails, the "Id" field will not be modified.
// Example:
//
//	doc := &Doc{}
//	hexId := "63632c7dfc826378c8abd802"
//	doc.PutId(hexId)
//	hex, _ := primitive.ObjectIDFromHex(hexId)
//	require.Equal(t, hex, doc.Id)
//
// Doc declaration:
//
//	type Doc struct {
//	    wgm.DefaultModel `bson:",inline"`
//	    Name             string `bson:"name"`
//	    Age              int    `bson:"age"`
//	}
//	func (d *Doc) ColName() string {
//	    return "Docs"
//	}
func (m *DefaultModel) PutId(id string) {
	hex, _ := primitive.ObjectIDFromHex(id)
	m.Id = hex
}

func (m *DefaultModel) setDefaultCreateTime() {
	m.CreateTime = time.Now().UnixMilli()
}

func (m *DefaultModel) setDefaultLastModifyTime() {
	m.LastModifyTime = time.Now().UnixMilli()
}

func (m *DefaultModel) setDefaultId() {
	if m.Id.IsZero() {
		m.Id = primitive.NewObjectID()
	}
}

// BeforeInsert sets the default values for the Id, CreateTime, and LastModifyTime fields of the DefaultModel
// before inserting it into the database. It is called before inserting a new document.
// This method should be called within the context of a transaction.
// It does not return an error.
func (m *DefaultModel) BeforeInsert(ctx context.Context) error {
	m.setDefaultId()
	m.setDefaultCreateTime()
	m.setDefaultLastModifyTime()
	return nil
}

// BeforeUpdate updates the last modify time of the DefaultModel instance to the current time.
func (m *DefaultModel) BeforeUpdate(ctx context.Context) error {
	m.setDefaultLastModifyTime()
	return nil
}

// BeforeUpsert sets defaults for Id, CreateTime, and LastModifyTime before upserting the DefaultModel.
func (m *DefaultModel) BeforeUpsert(ctx context.Context) error {
	m.setDefaultId()
	m.setDefaultCreateTime()
	m.setDefaultLastModifyTime()
	return nil
}

// GetObjectID returns the ObjectID of the DefaultModel.
// The ObjectID is used to uniquely identify the DefaultModel in the database.
// It is retrieved from the `Id` field of the DefaultModel struct.
func (m *DefaultModel) GetObjectID() primitive.ObjectID {
	return m.Id
}
