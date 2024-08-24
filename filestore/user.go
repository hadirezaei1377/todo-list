package filestore

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo-list/constant"
	"todo-list/entity"
)

type FileStore struct {
	filePath          string
	serializationMode string
}

// constructor
func New(path, serializationMode string) FileStore {
	return FileStore{filePath: path, serializationMode: serializationMode}
}

func (f FileStore) Save(u entity.User) {
	f.writeUserToFile(u)
}

func (f FileStore) Load() []entity.User {
	var uStore []entity.User

	file, err := os.Open(f.filePath)
	if err != nil {
		fmt.Println("can't open the file", err)
	}

	var data = make([]byte, 1024)
	_, oErr := file.Read(data)
	if oErr != nil {
		fmt.Println("can't read from the file", oErr)

		return nil
	}

	var dataStr = string(data)

	userSlice := strings.Split(dataStr, "\n")
	fmt.Println("len userSlice", len(userSlice), f.serializationMode)
	for _, u := range userSlice {
		var userStruct = entity.User{}

		switch f.serializationMode {
		case constant.ManDarAvardiSerializationMode:
			var dErr error
			userStruct, dErr = deserilizeFromManDaravardi(u)
			if dErr != nil {
				fmt.Println("can't deserialize user record to user struct", dErr)

				return nil
			}
		case constant.JsonSerializationMode:
			if u[0] != '{' && u[len(u)-1] != '}' {
				continue
			}

			uErr := json.Unmarshal([]byte(u), &userStruct)
			if uErr != nil {
				fmt.Println("can't deserialize user record to user struct with json mode", uErr)

				return nil
			}
		default:
			fmt.Println("invalid serialization mode")

			return nil
		}

		uStore = append(uStore, userStruct)
	}

	return uStore
}

func (f FileStore) writeUserToFile(user entity.User) {
	var file *os.File

	file, err := os.OpenFile(f.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("can't create or open file", err)

		return
	}
	defer file.Close()

	var data []byte
	// serialize the user struct/object
	if f.serializationMode == constant.ManDarAvardiSerializationMode {
		data = []byte(fmt.Sprintf("id: %d, name: %s, email: %s, password: %s\n", user.ID, user.Name,
			user.Email, user.Password))
	} else if f.serializationMode == constant.JsonSerializationMode {
		//json

		var jErr error
		data, jErr = json.Marshal(user)
		if jErr != nil {
			fmt.Println("can't marshal user struct to json", jErr)

			return
		}

		data = append(data, []byte("\n")...)

	} else {
		fmt.Println("invalid serialization mode")

		return
	}

	numberOfWrittenBytes, wErr := file.Write(data)
	if wErr != nil {
		fmt.Printf("can't write to the file %v\n", wErr)

		return
	}

	fmt.Println("numberOfWrittenBytes", numberOfWrittenBytes)
}

func deserilizeFromManDaravardi(userStr string) (entity.User, error) {

	if userStr == "" {
		return entity.User{}, errors.New("user string is empty")
	}

	var user = entity.User{}

	userFields := strings.Split(userStr, ",")
	for _, field := range userFields {
		values := strings.Split(field, ": ")
		if len(values) != 2 {
			fmt.Println("field is not valid, skipping...", len(values))

			continue
		}
		fieldName := strings.ReplaceAll(values[0], " ", "")
		fieldValue := values[1]

		switch fieldName {
		case "id":
			id, err := strconv.Atoi(fieldValue)
			if err != nil {
				return entity.User{}, errors.New("strconv error")
			}
			user.ID = id
		case "name":
			user.Name = fieldValue
		case "email":
			user.Email = fieldValue
		case "password":
			user.Password = fieldValue
		}

	}

	return user, nil
}
