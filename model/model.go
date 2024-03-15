package model

import (
	"fmt"

	"github.com/google/uuid"
)

type Dog struct {
    Id string
    Name string
    Breed string
}

type DogMap map[string]Dog

func AddDog(dogMap DogMap, name string, breed string) Dog {
	fmt.Println("AddDog: name =", name)
	fmt.Println("AddDog: breed =", breed)
	dog := Dog{
		Id: uuid.NewString(),
		Name: name,
		Breed: breed,
	}
	dogMap[dog.Id] = dog
	return dog;
}
