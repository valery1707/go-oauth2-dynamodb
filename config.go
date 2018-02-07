package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Config dynamodb configuration parameters
type Config struct {
	SESSION  *session.Session
	TABLE    *TableConfig
	ENDPOINT string
}

type TableConfig struct {
	BasicCname   string
	AccessCName  string
	RefreshCName string
}

// NewConfig create dynamodb configuration
func NewConfig(region string, endpoint string, access_key string, secret string) (config *Config, err error) {
	awsConfig := aws.NewConfig()
	if len(region) > 0 {
		awsConfig.Region = aws.String(region)
	}
	if len(access_key) > 0 && len(secret) > 0 {
		awsConfig.Credentials = credentials.NewStaticCredentials(access_key, secret, "")
	}
	if len(endpoint) > 0 {
		awsConfig.Endpoint = aws.String(endpoint)
	}
	newSession, err := session.NewSession(awsConfig)
	if err != nil {
		return
	}
	config = &Config{
		SESSION: newSession,
		TABLE: &TableConfig{
			BasicCname:   "oauth2_basic",
			AccessCName:  "oauth2_access",
			RefreshCName: "oauth2_refresh",
		},
		ENDPOINT: endpoint,
	}
	return
}
