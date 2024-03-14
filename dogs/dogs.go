package dogs

import (
	"fmt"

	"github.com/google/uuid"
)

type Dog struct {
    id string
    name string
    breed string
}

type DogMap map[string]Dog

func Add(dogMap DogMap, name string, breed string) Dog {
	fmt.Println("name =", name)
	fmt.Println("breed =", breed)
	dog := Dog{
		id: uuid.NewString(),
		name: name,
		breed: breed,
	}
	dogMap[dog.id] = dog
	return dog;
}
