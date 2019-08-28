package domain

import "encoding/base64"

type OEM string
type OEMPageURL string
type OEMPageResultUrl string

func GetOEMURL() string {
	return DecodeBase64("aHR0cHM6Ly93d3cuYXV0b3RyYWRlci5jby56YS9zZWFyY2gvbWFrZW1vZGVsc2F1dG9jb21wbGV0ZQ==")
}

func DecodeBase64(in string) string {
	sDec, _ := base64.StdEncoding.DecodeString(in)
	return string(sDec)
}

func GetOEMPagesURL() string {
	return DecodeBase64("aHR0cHM6Ly93d3cuYXV0b3RyYWRlci5jby56YS9jYXJzLWZvci1zYWxlLw==")
}

func GetBaseURL() string {
	return DecodeBase64("aHR0cHM6Ly93d3cuYXV0b3RyYWRlci5jby56YQ==")
}
