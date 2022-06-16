package otp_generator

import "github.com/xlzd/gotp"

func OtpGenerator() string {
	secretLength := 4
	data := gotp.RandomSecret(secretLength)
	return data
}
