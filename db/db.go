package db

import (
	"bufio"
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"mime/multipart"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Client the Mongo client
var Client = new(mongo.Client)

var c = new(mongo.ChangeStream)

//ConnectMongo ...
func ConnectMongo(address string) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(address))
	if err != nil {
		panic(err.Error())

	}
	log.Println("Successfully connected to the address provided")
	Client = client

}

type model interface {
	DBMate() map[string]string
}

//C the mongo Collection
func C(m model) *mongo.Collection {
	return Client.Database(m.DBMate()["dbName"]).Collection(m.DBMate()["cName"])
}

//InsertOne ...
func InsertOne(m model) (*mongo.InsertOneResult, error) {

	data, err := bson.Marshal(m)

	if err != nil {
		return nil, err
	}

	bm := bson.M{}
	bson.Unmarshal(data, &bm)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return C(m).InsertOne(ctx, bm)
}

//UpdateOne ...
func UpdateOne(m model, filter interface{}) (*mongo.UpdateResult, error) {

	data, err := bson.Marshal(m)

	if err != nil {
		return nil, err
	}

	bm := bson.M{}
	bson.Unmarshal(data, &bm)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return C(m).UpdateOne(ctx, filter, bson.M{"$set": bm})
}

//SessionInsertOne ...
func SessionInsertOne(ctx mongo.SessionContext, m model) (*mongo.InsertOneResult, error) {

	data, err := bson.Marshal(m)

	if err != nil {
		return nil, err
	}

	bm := bson.M{}
	bson.Unmarshal(data, &bm)

	return C(m).InsertOne(ctx, bm)
}

//RunSession ...
func RunSession(fn func(mongo.SessionContext) error) error {
	var err error
	session, err := Client.StartSession()
	if err != nil {
		return err
	}
	if err = session.StartTransaction(); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err = fn(sc)
		if err != nil {
			return err
		}
		if err = session.AbortTransaction(sc); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	session.EndSession(ctx)
	return nil

}

//SaveFsFile ...
func SaveFsFile(dbName string, file *multipart.FileHeader) (primitive.ObjectID, error) {
	var err error
	var bucket *gridfs.Bucket
	var ustream *gridfs.UploadStream
	if bucket, err = gridfs.NewBucket(Client.Database(dbName), options.GridFSBucket().SetName("fs_files")); err != nil {
		return primitive.NilObjectID, err
	}

	opts := options.GridFSUpload()
	opts.SetMetadata(map[string]string{"content-type": file.Header.Get("content-type")})

	t := strconv.FormatInt(time.Now().Unix(), 10)
	if ustream, err = bucket.OpenUploadStream(t+file.Filename, opts); err != nil {
		return primitive.NilObjectID, err
	}

	defer ustream.Close()

	fileContent, _ := file.Open()
	byteContainer, _ := ioutil.ReadAll(fileContent)

	if _, err = ustream.Write(byteContainer); err != nil {

		return primitive.NilObjectID, err
	}
	fileID := ustream.FileID.(primitive.ObjectID)
	return fileID, nil

}

//SaveFsFormBuffer ...
func SaveFsFormBuffer(dbName, contentType string, buffer *bytes.Buffer) (primitive.ObjectID, error) {
	var err error
	var bucket *gridfs.Bucket
	var ustream *gridfs.UploadStream

	if bucket, err = gridfs.NewBucket(Client.Database(dbName), options.GridFSBucket().SetName("fs_files")); err != nil {
		return primitive.NilObjectID, err
	}

	opts := options.GridFSUpload()
	opts.SetMetadata(map[string]string{"content-type": contentType})

	name := primitive.NewObjectID().Hex()
	if ustream, err = bucket.OpenUploadStream(name, opts); err != nil {
		return primitive.NilObjectID, err
	}

	defer ustream.Close()

	if _, err = ustream.Write(buffer.Bytes()); err != nil {

		return primitive.NilObjectID, err
	}
	fileID := ustream.FileID.(primitive.ObjectID)
	return fileID, nil

}

// FsFile ...
type FsFile struct {
	ID         interface{} `bson:"_id"`
	ChunkSize  int         `bson:"chunkSize"`
	UploadDate time.Time   `bson:"uploadDate"`
	Length     int64       ",minsize"
	MD5        string
	Filename   string            ",omitempty"
	Metadata   map[string]string ",omitempty"
	Buffer     bytes.Buffer      ",omitempty"
}

//Bytes  ...
func (fs *FsFile) Bytes() []byte {
	return fs.Buffer.Bytes()
}

//DelFsFile ...
func DelFsFile(dbName string, fid string) error {

	var err error
	var bucket *gridfs.Bucket
	if bucket, err = gridfs.NewBucket(Client.Database(dbName), options.GridFSBucket().SetName("fs_files")); err != nil {
		return err
	}

	fileID, _ := primitive.ObjectIDFromHex(fid)

	return bucket.Delete(fileID)

}

//GetFsFile ...
func GetFsFile(dbName string, fid string) (*FsFile, error) {

	var err error
	var bucket *gridfs.Bucket
	var doc FsFile

	doc.Buffer = bytes.Buffer{}

	if bucket, err = gridfs.NewBucket(Client.Database(dbName), options.GridFSBucket().SetName("fs_files")); err != nil {
		return &doc, err
	}

	fileID, _ := primitive.ObjectIDFromHex(fid)

	limit := int32(1)
	opts := options.GridFSFind()
	opts.Limit = &limit
	cursor, err := bucket.Find(bson.M{"_id": fileID}, opts)

	if err != nil {
		return &doc, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for cursor.Next(ctx) {
		cursor.Decode(&doc)
	}

	w := bufio.NewWriter(&doc.Buffer)

	if _, err = bucket.DownloadToStream(fileID, w); err != nil {
		return &doc, err
	}

	return &doc, nil

}
