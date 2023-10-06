package api

import (
	"fmt"

	"github.com/tradlwa/xj/api/urlcodec"
)

var ErrBadCredentials = fmt.Errorf("invalid username or password")

func Login(client *Client, username string, password string) error {
	loginRequest := struct {
		Username   string `url:"userName"`
		Password   string `url:"password"`
		IfRemember string `url:"ifRemember"`
	}{
		Username:   username,
		Password:   password,
		IfRemember: "on",
	}

	var response BaseResponse
	err := client.Post("login", urlcodec.StructToStringReader(loginRequest), &response)
	if err != nil {
		return err
	}
	if response.Code == 200 {
		return nil
	}
	return ErrBadCredentials
}
