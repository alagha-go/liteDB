package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tidwall/pretty"
	"github.com/tidwall/sjson"
)

type Database struct{
	Name 						string
	Collections					[]Collection
}

type DB struct {
	Database					string
	Collection					string
	Path						string
}

type Collection struct {
	Name 						string
	Documents					[]string
}

type Cursor struct {
	Data 						[]interface{}
}

const (
	DBPath = "./DBs/"
)


func CreateDatabase(name string) DB {
	var Database DB
	database := fmt.Sprintf("%s%s", DBPath, name)
	exists := DBexists(database)
	if !exists {
		err := os.Mkdir(database, 0755)
		if err != nil {
			fmt.Println(err)
		}else{
			fmt.Printf("database %s successfully created\n", name)
		}
	}else{
		fmt.Printf("database %s already exists\n", name)
	}
	Database.Database = name[:]
	return Database
}


func (database DB)CreateCollection(name string) DB {
	collection := fmt.Sprintf("%s%s/%s", DBPath, database.Database, name)
	database.Path = collection
	exists := DBexists(collection)
	if !exists {
		err := os.Mkdir(collection, 0755)
		if err != nil {
			fmt.Println(err)
		}else{
			fmt.Printf("colleeection %s successfully created\n", name)
		}
	}else{
		fmt.Printf("collection %s already exists\n", name)
	}
	database.Collection = name
	fmt.Printf("Database: %s\n",database.Database)
	fmt.Printf("Collection: %s\n", database.Collection)

	return	database
}


func (db DB)CreateDocument(document interface{}) Cursor {
	var createdDocument interface{}
	var cursor Cursor

	ID := GenerateId()

	documentName := fmt.Sprintf("%s.json", ID)

	FileName := fmt.Sprintf("%s%s/%s/%s", DBPath, db.Database, db.Collection, documentName)

	exist := DBexists(FileName)
	if exist {
		return db.CreateDocument(document)
	}
	byteDocument, err := json.Marshal(document)
	HandleError(err)

	jsonDocument := string(byteDocument)

	jsonDocument, err = sjson.Set(jsonDocument, "_id", ID)
	HandleError(err)
	

	byteDocument = pretty.Pretty([]byte(jsonDocument))

	// byteDocument = []byte(jsonDocument)

	err = ioutil.WriteFile(FileName, byteDocument,0644)
	HandleError(err)

	err = json.Unmarshal(byteDocument, &createdDocument)
	HandleError(err)

	cursor.Data = append(cursor.Data, createdDocument)

	return cursor
}





func (db DB)GetAllDocuments() Cursor {
	var cursor Cursor
	var alldocuments []interface{}
	files, err := ioutil.ReadDir(db.Path)
    HandleError(err)

	for _, file := range files{
		var document interface{}
		filePath := fmt.Sprintf("%s/%s", db.Path, file.Name())
		content, err := LoadFile(filePath)
		HandleError(err)
		err = json.Unmarshal(content, &document)
		HandleError(err)
		alldocuments = append(alldocuments, document)
	}
	cursor.Data = alldocuments
	return cursor
}


func LoadFile(file string) ([]byte, error) {
	fileContent, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return fileContent, nil
}


func (db DB) GetOne()


func (c Cursor) Decode(Type interface{}) {
	byteDocument, err := json.Marshal(c.Data[0])
	HandleError(err)
	err = json.Unmarshal(byteDocument, &Type)
	HandleError(err)
}

func (c Cursor) DecodeMany(Type interface{}){
	byteDocuments, err := json.Marshal(c.Data)
	HandleError(err)
	err = json.Unmarshal(byteDocuments, &Type)
	HandleError(err)
}


func DBexists(DBFile string) bool {
	if _, err := os.Stat(DBFile); os.IsNotExist(err){
		return false
	}
	return true
}