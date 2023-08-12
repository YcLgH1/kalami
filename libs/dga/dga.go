package dga

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"strconv"
	"strings"
	"time"
)

// Check Error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Add 0 prefix if otp less than length
func AddPrefix0(otp string, length int) string {
	if len(otp) == length {
		return otp
	}
	for i := (length - len(otp)); i > 0; i-- {
		otp = "0" + otp
	}
	return otp
}

// Parse number string to Char string
// src: num-string, 6-times-length
func Num2Char(src string, length int) string {
	if len(src) != length*6 {
		return ""
	}

	tar := []byte{}
	for i := 0; i < 6; i++ {
		key := src[i*length : i*length+length]
		num, err := strconv.Atoi(key)
		check(err)

		num = num % 62

		if num < 10 {
			num += 48
		} else if num < 36 {
			// mum += 65 - 10
			num += 55
		} else {
			// num += 97 - 36
			num += 61
		}

		tar = append(tar, byte(num))
	}

	return string(tar)
}

// Get H-OTP Token
// secret string: 40 chars hex string [0-9, a-e]
// time period: valid time period, eg: 3600 stands for 3600s
func GetHOTPToken(secret string, timePoint int64) string {
	key, err := base64.StdEncoding.DecodeString(strings.ToUpper(secret))
	check(err)

	// int64 is 8*8bit => 8byte
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(timePoint))

	// use sha256(32byte) instead of sha1(20byte)
	hash := hmac.New(sha256.New, key)
	hash.Write(bs)
	h := hash.Sum(nil)

	// sha256 is 32 byte, and maximum decimal 15, we need 16 byte
	o := (h[31] & 15)

	var header uint64
	r := bytes.NewReader(h[o : o+16])
	err = binary.Read(r, binary.BigEndian, &header)
	check(err)

	h12 := uint(header) % 1000000000000000000
	otp := strconv.Itoa(int(h12))

	return AddPrefix0(otp, 18)
}

// Get T-OTP Token <1 hour> refresh
// secret string: 40 chars hex string [0-9, a-e]
func GetTOTPToken(secret string, offset int64) string {
	interval := 3600
	timePoint := time.Now().Unix()/int64(interval) + offset

	return GetHOTPToken(secret, timePoint)
}

// Get A 6-char-length Domain from secret, valid for 1 hour
// secret string: 40 chars hex string [0-9, a-e]
func GetDomain(secret string, offset int64) string {
	return Num2Char(GetTOTPToken(secret, offset), 3)
}
