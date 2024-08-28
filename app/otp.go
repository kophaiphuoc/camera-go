package app

import (
	"fmt"
	"github.com/pquerna/otp/totp"
	"time"
)

func GenerateOTP(secret string) (string, int64, error) {

	// gen OTP for now time
	otp, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return "", 0, err
	}

	// default period of time
	timeStep := 30

	// get now time
	currentTime := time.Now().Unix()

	// calculator time compared to present
	elapsed := currentTime % int64(timeStep)

	// Calculate remaining time in current cycle
	timeLeft := int64(timeStep) - elapsed

	return otp, timeLeft, nil
}

func startOTPRoutine(secret string) {
	// create a ticker to call the function back every 30 seconds
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		otp, timeLeft, err := GenerateOTP(secret)
		if err != nil {
			fmt.Printf("Lỗi khi tạo mã OTP: %v\n", err)
			return
		}

		fmt.Printf("Mã OTP: %s\n", otp)
		fmt.Printf("Thời gian còn lại trong chu kỳ 30 giây: %d giây\n", timeLeft)

		<-ticker.C
	}
}