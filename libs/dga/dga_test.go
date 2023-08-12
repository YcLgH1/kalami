package dga

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestOTP(t *testing.T) {
	secret := "dummysecretdummydummysecretdummy"
	otp := GetTOTPToken(secret, 0)
	domain := GetDomain(secret, 0)

	t.Logf("secret %s otp %s domain %s", secret, otp, domain)

	//Copies the otp generated to your clipboard
	err := exec.Command("bash", "-c", fmt.Sprintf("echo %s | tr -d \"\n, \" | pbcopy", otp)).Run()
	check(err)
}

func TestNum2Str(t *testing.T) {
	for i := 0; i < 1000; i++ {
		// num := i % 36
		// if num < 26 {
		// 	num += 97
		// } else {
		// 	// num += 48 - 26
		// 	num += 22
		// }

		num := i % 62
		if num < 10 {
			num += 48
		} else if num < 36 {
			// mum += 65 - 10
			num += 55
		} else {
			// num += 97 - 36
			num += 61
		}

		fmt.Printf("%s", string(byte(num)))
	}
}
