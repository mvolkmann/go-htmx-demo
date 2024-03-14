package dogs

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

func Add(dogMap DogMap, name string, breed string) Dog {
	fmt.Println("name =", name)
	fmt.Println("breed =", breed)
	dog := Dog{
		Id: uuid.NewString(),
		Name: name,
		Breed: breed,
	}
	dogMap[dog.Id] = dog
	return dog;
}
