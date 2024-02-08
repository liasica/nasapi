// Copyright (C) liasica. 2024-present.
//
// Created at 2024-02-08
// Based on nasapi by liasica, magicrolan@qq.com.

package qnap

type QvrCameraList struct {
	Datas []struct {
		ChannelIndex           int    `json:"channel_index"`
		Name                   string `json:"name"`
		Umsid                  string `json:"umsid"`
		Guid                   string `json:"guid"`
		Brand                  string `json:"brand"`
		Model                  string `json:"model"`
		Mac                    string `json:"mac"`
		Ver                    string `json:"ver"`
		Ip                     string `json:"ip"`
		Port                   string `json:"port"`
		VideoCodecSetting      string `json:"video_codec_setting"`
		VideoResolutionSetting string `json:"video_resolution_setting"`
		FrameRateSetting       string `json:"frame_rate_setting"`
		VideoQualitySetting    string `json:"video_quality_setting"`
		StreamState            []struct {
			Stream                 int    `json:"stream"`
			EnableNormalRecording  int    `json:"enable_normal_recording"`
			EnableAlarmRecording   int    `json:"enable_alarm_recording"`
			VideoCodecSetting      string `json:"video_codec_setting"`
			VideoResolutionSetting string `json:"video_resolution_setting"`
			FrameRateSetting       string `json:"frame_rate_setting"`
			VideoQualitySetting    string `json:"video_quality_setting"`
			Status                 string `json:"status"`
			RecState               string `json:"rec_state"`
			RecStateErrCode        int    `json:"rec_state_err_code"`
			FrameRate              string `json:"frame_rate"`
			BitRate                int    `json:"bit_rate"`
		} `json:"stream_state"`
		Status          string `json:"status"`
		RecState        string `json:"rec_state"`
		RecStateErrCode int    `json:"rec_state_err_code"`
		FrameRate       string `json:"frame_rate"`
		BitRate         int    `json:"bit_rate"`
	} `json:"datas"`
	Success         string `json:"success"`
	TotalChannelNum int    `json:"total_channel_num"`
}
