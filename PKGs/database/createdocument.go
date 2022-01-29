package database

import (
	"fmt"
)


type Database struct {
	Name							string								`json:"name"`
	NumberOfCollections				int									`json:"number_of_collections"`
	Size							string								`json:"size"`
	Collections						[]Collection						`json:"collections"`
}

type Collection struct {
	Name							string								`json:"name"`
	NumberOfDocuments				int									`json:"number_of_documents"`
	Size							string								`json:"size"`
	Documents						[]Document 							`json:"documents"`
}

type Document struct {
	ID								string								`json:"_id"`
	Size							string								`json:"size"`
	Data							interface{}							`json:"data"`
}


func CreateOneDocument(dbName, colName, data string) {
	
}




func GenerateSize(b int64) string {
	const unit = 1000
    if b < unit {
		return fmt.Sprintf("%d B", b)
    }
    div, exp := int64(unit), 0
    for n := b / unit; n >= unit; n /= unit {
		div *= unit
        exp++
    }
    return fmt.Sprintf("%.1f %cB",
	float64(b)/float64(div), "kMGTPE"[exp])
}




func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}