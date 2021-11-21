package mongodb

import (
	"fmt"
	"testing"
)

func TestInitMongo(t *testing.T) {
	_, err := InitMongo()
	fmt.Println(err)
}
