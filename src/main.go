package main

import (
	"app/src/config"
	"app/src/logger"
	"app/src/redis"
	"app/src/router"

	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	region   = "ap-northeast-1"
	confPath = flag.String("c", "", "config file path (required).")
	conf     *config.Config
)

func main() {
	flag.Parse()
	if *confPath == "" {
		flag.Usage()
		return
	}

	var err error
	conf, err = config.New(*confPath)
	if err != nil {
		fmt.Printf("[ERROR] Can not read configuration file: %v", err)
		return
	}

	if logErr := logger.Init(conf); logErr != nil {
		fmt.Println("[ERROR] Failed to initialize logger instance", err)
		return
	}

	cred := credentials.NewStaticCredentials(
		os.Getenv("AWS_ACCESS_KEY"),
		os.Getenv("AWS_SECRET_ACCESS_KEY"),
		"",
	)
	conf := &aws.Config{
		Credentials: cred,
		Region:      &region,
	}

	sess, awsErr := session.NewSession(conf)
	if awsErr != nil {
		logger.Error("Can not create new aws sessoin. error: %v", awsErr)
		return
	}

	rd, err := redis.New()
	if err != nil {
		logger.Error("Failed to create redis instance. error: %v", err)
		return
	}

	r := router.New(sess, rd)
	r.Run(":8080")
}
