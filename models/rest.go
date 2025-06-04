package models

type Response struct {
	Code         int         `json:"code"`
	ResponseCode string      `json:"responseCode"`
	ResponseDesc string      `json:"responseDesc"`
	ResponseData interface{} `json:"responseData,omitempty"`
}
