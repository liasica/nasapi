// Copyright (C) liasica. 2024-present.
//
// Created at 2024-02-08
// Based on qnap-apis by liasica, magicrolan@qq.com.

package qnap

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// doc https://petstore.swagger.io/?url=https://download.qnap.com/apidoc/qvrpro/qvr_pro_api_1.0.0.yaml

func (q *Qnap) QvrProCameraList() {
	var res QvrCameraList
	_, err := q.client.R().
		SetQueryParams(map[string]string{
			"sid": q.sid,
			"ver": "1.0.0",
		}).
		SetResult(&res).
		Get("/qvrpro/camera/list")

	if err != nil {
		log.Errorf("")
	}
	fmt.Printf("%#v", res)
}
