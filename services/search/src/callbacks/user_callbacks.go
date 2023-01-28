package callbacks

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/A-Siam/bracker/search/src/database"
	"github.com/A-Siam/bracker/search/src/dto"
	es "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func OnUserCreated(userDto dto.UserDto, ok chan bool) {
	db := database.GetDB()
	err := indexUser(db, userDto)
	if err != nil {
		ok <- false
		return
	}
	ok <- true
}

func indexUser(db *es.Client, userDto dto.UserDto) error {
	data, err := json.Marshal(userDto)
	if err != nil {
		return err
	}
	req := esapi.IndexRequest{
		Index:      "users",
		DocumentID: userDto.ID,
		Body:       bytes.NewReader(data),
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), db)
	if err != nil {
		return err
	}
	res.Body.Close()
	return nil
}
