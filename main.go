package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func main() {
	ctx := context.Background()

	rootCfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     "access-key",
				SecretAccessKey: "secret-key",
				Source:          "example hard coded credentials",
			},
		}))
	if err != nil {
		panic(err)
	}
	rootCfg.Region = "us-east-1"

	stsSvc := sts.NewFromConfig(rootCfg)
	genCreds := stscreds.NewAssumeRoleProvider(stsSvc, fmt.Sprintf("arn:aws:iam::%d:role/%s", 111111111111, "role-name"))

	cfg := rootCfg.Copy()
	cfg.Credentials = aws.NewCredentialsCache(genCreds)

	creds, err := cfg.Credentials.Retrieve(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(creds)
}
