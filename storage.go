package main

import (
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/labstack/echo/v4"
)

var storage_base_folder = "data_storage"

func StorageHandlerSet(ctx echo.Context) error {
	NameFile := ctx.Param("name") + ".json"
	var dir_file = path.Join(storage_base_folder, NameFile)
	os.Mkdir(storage_base_folder, 0777)
	os.Remove(dir_file)

	var data, err = ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}
	ioutil.WriteFile(dir_file, data, 0777)
	return nil
}

func StorageHandlerGet(ctx echo.Context) error {
	NameFile := ctx.Param("name") + ".json"
	var dir_file = path.Join(storage_base_folder, NameFile)

	var file, err = os.Open(dir_file)
	if err != nil {
		ctx.JSON(400, nil)
		return err
	}
	defer file.Close()
	io.Copy(ctx.Response().Writer, file)
	return nil
}
