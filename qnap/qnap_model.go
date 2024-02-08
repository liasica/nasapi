// Copyright (C) liasica. 2024-present.
//
// Created at 2024-02-08
// Based on nasapi by liasica, magicrolan@qq.com.

package qnap

type AuthLoginResponse struct {
	AuthSid    string `xml:"authSid"`
	ErrorValue string `xml:"errorValue"`
}
