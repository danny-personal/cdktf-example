package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-google-go/google/v7/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// The code that defines your stack goes here
	provider.NewGoogleProvider(stack, jsii.String("google"), &provider.GoogleProviderConfig{
		Region:  jsii.String("asia-northeast1"),
		Project: jsii.String(""),
		Zone:    jsii.String(""),
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "cdktf-example")

	app.Synth()
}
