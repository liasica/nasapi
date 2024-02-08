// Copyright (C) liasica. 2024-present.
//
// Created at 2024-02-08
// Based on qnap-apis by liasica, magicrolan@qq.com.

package qnap

import (
	"encoding/base64"
	"encoding/xml"
	"sync"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// docs: https://qnap-dev.github.io/container-station-api/system.html

type Qnap struct {
	username       string
	password       string
	base64Password string
	baseUrl        string
	client         *resty.Client
	sid            string
}

var (
	instance *Qnap
	once     sync.Once
)

func GetInstance() *Qnap {
	once.Do(func() {
		pass := viper.GetString("qnap.password")
		baseUrl := viper.GetString("qnap.baseUrl")
		instance = &Qnap{
			username:       viper.GetString("qnap.username"),
			password:       pass,
			baseUrl:        baseUrl,
			base64Password: base64.StdEncoding.EncodeToString([]byte(pass)),
			client: resty.New().
				SetBaseURL(baseUrl).
				SetXMLMarshaler(xml.Marshal).SetXMLUnmarshaler(xml.Unmarshal).
				SetJSONMarshaler(jsoniter.Marshal).SetJSONUnmarshaler(jsoniter.Unmarshal),
		}
		// 首次登录
		instance.cgiAuthLogin()
	})
	return instance
}

func (q *Qnap) cgiAuthLogin() {
	var res AuthLoginResponse
	_, err := q.client.R().
		SetQueryParams(map[string]string{
			"user":       q.username,
			"pwd":        q.base64Password,
			"serviceKey": "1",
		}).
		SetResult(&res).
		Get("/cgi-bin/authLogin.cgi")
	if err != nil {
		log.Errorf("[QNAP] 认证请求失败: %s", err.Error())
	}
	if res.AuthSid == "" {
		log.Error("[QNAP] 认证失败")
	}

	q.sid = res.AuthSid
}

func (q *Qnap) getUrl(uri string) string {
	return q.baseUrl + uri
}
