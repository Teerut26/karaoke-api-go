package services

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func CleanHLS() {
	fmt.Println(time.Now().Format("[ 2006-01-02 15:04:05 ]"), "Checking hls folder")
	if _, err := os.Stat("hls"); os.IsNotExist(err) {
		return
	}

	files, err := os.ReadDir("hls")
	if err != nil {
		return
	}

	for _, file := range files {
		r, _ := regexp.Compile(`(?m)_(\d+)`)

		if !r.MatchString(file.Name()) {
			continue
		}

		fileExpireTime, _ := strconv.ParseInt(r.FindStringSubmatch(file.Name())[1], 10, 64)

		if time.Now().Unix() > fileExpireTime {
			os.RemoveAll("hls/" + file.Name())
			fmt.Println("Deleted", file.Name())
		}
	}
}
