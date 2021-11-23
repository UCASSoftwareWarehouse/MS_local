package mysql

import (
	"fmt"
	"testing"
)

func TestInitMysql(t *testing.T) {
	_, err := InitMysql()
	fmt.Println(err)
}
