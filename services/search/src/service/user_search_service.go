package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/A-Siam/bracker/search/src/database"
	"github.com/A-Siam/bracker/search/src/dto"
)

func FindUserByName(name string) ([]dto.UserDto, error) {
	db := database.GetDB()
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"wildcard": map[string]interface{}{
				"name": "*" + name + "*",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return []dto.UserDto{}, nil
	}
	res, err := db.Search(
		db.Search.WithContext(context.Background()),
		db.Search.WithIndex("users"),
		db.Search.WithBody(&buf),
	)
	if err != nil {
		return []dto.UserDto{}, err
	}
	defer res.Body.Close()
	var queryResult map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&queryResult); err != nil {
		return []dto.UserDto{}, err
	}
	if res.IsError() {
		var e map[string]interface{}
		return []dto.UserDto{}, errors.New(fmt.Sprintf("[%s] %s: %s",
			res.Status(),
			e["error"].(map[string]interface{})["type"],
			e["error"].(map[string]interface{})["reason"],
		))
	}
	userDtos := []dto.UserDto{}
	contentJson := queryResult["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hits := range contentJson {
		contentMap := hits.(map[string]interface{})["_source"].(map[string]interface{})
		jsonBody, err := json.Marshal(contentMap)
		if err != nil {
			return []dto.UserDto{}, err
		}
		var userDto dto.UserDto
		if err := json.Unmarshal(jsonBody, &userDto); err != nil {
			return []dto.UserDto{}, err
		}
		userDtos = append(userDtos, userDto)
	}
	return userDtos, nil
}
