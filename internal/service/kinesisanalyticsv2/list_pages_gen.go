// Code generated by "internal/generate/listpages/main.go -ListOps=ListApplications -ContextOnly"; DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesisanalyticsv2"
	"github.com/aws/aws-sdk-go/service/kinesisanalyticsv2/kinesisanalyticsv2iface"
)

func listApplicationsPages(ctx context.Context, conn kinesisanalyticsv2iface.KinesisAnalyticsV2API, input *kinesisanalyticsv2.ListApplicationsInput, fn func(*kinesisanalyticsv2.ListApplicationsOutput, bool) bool) error {
	for {
		output, err := conn.ListApplicationsWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
