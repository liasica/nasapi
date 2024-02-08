// Copyright (C) liasica. 2024-present.
//
// Created at 2024-02-08
// Based on qnap-apis by liasica, magicrolan@qq.com.

package main

import (
	_ "github.com/liasica/nasapi"
	"github.com/liasica/nasapi/qnap"
)

func main() {
	qnap.GetInstance().QvrProCameraList()
}
