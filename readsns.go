package main

import (

	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {

	// Connect to DynamoDB
	config := aws.NewConfig()
	config.Region = aws.String("eu-west-1")
	sess := session.Must(session.NewSession(config))
	dyn := dynamodb.New(sess, config)
	params := &dynamodb.DescribeTableInput{
    			TableName: aws.String("TableName"), // Required
		}
	resp, err := dyn.DescribeTable(params)

	if err != nil {
	    // Print the error, cast err to awserr.Error to get the Code and
	    // Message from an error.
	    fmt.Println(err.Error())
	    return
	}

	// Pretty-print the response data.
	fmt.Println(resp)

}
