package main

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	base1 := []byte("1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	rand.Seed(1)
	for i := 0; i < 10; i++ {
		x := rand.Intn(len(base1))
		y := rand.Intn(len(base1))
		base1[x], base1[y] = base1[y], base1[x]
	}
	var username string
	fmt.Println("请输入用户名：")
	_, _ = fmt.Scanln(&username)
	var password []byte
	fmt.Println("请输入注册码：")
	_, _ = fmt.Scanln(&password)
	var result uint64
	for key, value := range []byte(username) {
		result += uint64(math.Pow(float64(bytes.IndexByte(base1, value)), float64(key+1)))
	}
	if len(password) != len(strconv.FormatUint(result, 10)) {
		fmt.Println("注册码长度错误！")
		os.Exit(1)
	}
	key := 0
	for result > 0 {
		value := int(math.Mod(float64(result), 10))
		result = result / 10
		if base1[value] != password[key] {
			fmt.Println("认证失败！")
			os.Exit(1)
		}
		key += 1
	}
	fmt.Println("成功注册！你太牛逼了！")
	if username=="wuaipojie" {
		fmt.Println("在 2020 年 9 月 16 日 00:00 前，私发支付宝号 + 本次注册码获取红包")
	}
}
