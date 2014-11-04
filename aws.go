package goenv

import (
	"github.com/crowdmob/goamz/aws"
	"log"
)

func (goenv *Goenv) GetAwsAuth() (auth aws.Auth) {
	auth.AccessKey = goenv.RequireFailToEnv("AWS_ACCESS_KEY_ID")
	auth.SecretKey = goenv.RequireFailToEnv("AWS_SECRET_ACCESS_KEY")
	return
}

func (goenv *Goenv) GetAwsRegion() (region aws.Region) {
	region_name := goenv.GetFailToEnv("AWS_REGION", "us-east-1")
	region, ok := aws.Regions[region_name]
	if !ok {
		log.Panicf("goenv Require couldn't find default aws region", goenv.environment, "")
	}
	return region
}

func (goenv *Goenv) GetAwsBucketName() (bucketName string) {
	bucketName = goenv.RequireFailToEnv("AWS_BUCKET_NAME")
	return bucketName
}
