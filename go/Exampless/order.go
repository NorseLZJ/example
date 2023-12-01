package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"
)

func generateRandomString(length int) string {
	const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := range result {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		result[i] = chars[randomIndex.Int64()]
	}
	return string(result)
}

func calculateMD5(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

func getOrderNumber(uid int32) string {
	currentTime := time.Now().Format("20060102")
	randomString := generateRandomString(5)
	s := fmt.Sprintf("%s%s%d", currentTime, randomString, uid)
	md5 := calculateMD5(s)
	return fmt.Sprintf("L%s%s%s", md5[:5], currentTime, randomString)
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(getOrderNumber(1))
	}
}
