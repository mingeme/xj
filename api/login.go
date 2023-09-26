package api

import (
	"fmt"
	"github.com/tradlwa/xj/api/urlcodec"
	"strings"
)

var ErrBadCredentials = fmt.Errorf("invalid username or password")

type LoginResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

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

	values := urlcodec.StructToValues(loginRequest)
	var response LoginResponse
	err := client.Post("login", strings.NewReader(values.Encode()), &response)
	if err != nil {
		return err
	}
	if response.Code == 200 {
		return nil
	}
	return ErrBadCredentials
}
