package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"go-zero-cloud-disk/core/models"

	_ "github.com/go-sql-driver/mysql"

	"xorm.io/xorm"
)

func TestXorm(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(localhost:3306)/cloud_disk?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		t.Fatal(err)
	}

	data := make([]*models.UserBasic, 0)
	err = engine.Find(&data)
	if err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dst.String())
}
