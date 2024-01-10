package wgm

import (
	"context"
	"github.com/oeasenet/jog"
	"github.com/qiniu/qmgo"
	"time"
)

// instance is a global variable of type *wgm.
// It is used to store an instance of the wgm struct, which is responsible for managing a connection to a database and providing methods for interacting with the database.
// Example usage:
// To initialize the instance, you can call the InitWgm function, passing in a connection URI and a database name.
//
//	err := InitWgm(connectionUri, databaseName)
//
// To close the connection to the database, you can call the CloseAll function.
//
//	CloseAll()
//
// To perform a database ping, you can call the Ping function.
//
//	err := Ping()
//
// To retrieve a qmgo.Collection object for a specific collection name, you can call the Col function.
//
//	collection := Col(collectionName)
//
// To retrieve the context associated with the instance, you can call the Ctx function.
//
//	ctx := Ctx()
//
// To paginate through a collection of documents, you can call the FindPage function.
//
//	totalDoc, totalPage := FindPage(model, filter, result, pageSize, currentPage)
//
// To paginate through a collection of documents with additional options, you can call the FindPageWithOption function.
//
//	totalDoc, totalPage := FindPageWithOption(model, filter, result, pageSize, currentPage, option)
//
// To check if a document exists in the collection, you can call the FindOne function.
//
//	hasResult := FindOne(model, filter)
//
// To find a document by its ID, you can call the FindById function.
//
//	found, err := FindById(collectionName, id, result)
//
// To insert a document into the collection, you can call the Insert function.
//
//	insertResult, err := Insert(model)
//
// The `wgm` type is a struct that has the following methods:
// - GetModelCollection: returns a qmgo.Collection object for a specific model.
// - GetCollection: returns a qmgo.Collection object for a specific collection name.
// - Ctx: returns the context associated with the wgm instance.
// - newCtxWithTimeout: creates a new context with a timeout duration.
// - newCtx: creates a new context with a default timeout duration.
// The `IDefaultModel` type is an interface that defines several methods related to working with database models.
var instance *wgm

type wgm struct {
	client *qmgo.Client
	dbName string
}

// InitWgm initializes the connection to the specified database using the provided connection URI and database name.
// Parameters:
// - connectionUri: The URI of the MongoDB connection.
// - databaseName: The name of the database to connect to.
// Returns:
// - error: An error if the connection could not be established.
// Usage example:
//
//	err := InitWgm("mongodb://localhost:27017/", "wshop_test")
//	if err != nil {
//		// Handle error
//	}
//
// Usage example:
//
//	err := InitWgm("wrong://localhost:27017/", "wshop_test")
//	if err != nil {
//		// Handle error
//	}
//
// Usage example:
//
//	SetupDefaultConnection()
func InitWgm(connectionUri string, databaseName string) error {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{
		Uri:      connectionUri,
		Database: databaseName,
	})
	if err != nil {
		return err
	}
	instance = &wgm{
		client: client,
		dbName: databaseName,
	}
	return nil
}

// newCtxWithTimeout creates a new context with the specified timeout.
// Parameters:
// - timeout: The duration of the timeout.
// Returns:
// - context.Context: The newly created context.
func (w *wgm) newCtxWithTimeout(timeout time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	return ctx
}

// newCtx creates a new context with a default timeout of 10 seconds.
// This method is used internally by the Ctx() method.
//
// Returns:
// - context.Context: The newly created context.
func (w *wgm) newCtx() context.Context {
	return w.newCtxWithTimeout(10 * time.Second)
}

func CloseAll() {
	if instance == nil {
		jog.Fatal("must initialize WGM first, by calling InitWgm() method")
	}

	if err := instance.client.Close(instance.Ctx()); err != nil {
		jog.Error(err)
	}
}

// Ping checks if the WGM instance is initialized and then performs a ping operation on the MongoDB connection.
// If the WGM instance is not initialized, it logs an error message and returns an error.
func Ping() error {
	if instance == nil {
		jog.Fatal("must initialize WGM first, by calling InitWgm() method")
	}
	return instance.client.Ping(10)
}
