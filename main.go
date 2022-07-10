package main

import (
	"fmt"

	"cdk.tf/go/stack/generated/hashicorp/aws"
	"cdk.tf/go/stack/generated/hashicorp/aws/datasources"
	"cdk.tf/go/stack/generated/hashicorp/aws/kms"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

const (
	env     = "Devops"
	company = "yourCompany"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)
	aws.NewAwsProvider(stack, jsii.String("AWS"), &aws.AwsProviderConfig{
		Region: jsii.String("us-west-1"),
	})

	awsAccount := datasources.NewDataAwsCallerIdentity(stack, jsii.String("aws-id"), nil)

	awsKmsKey := kms.NewKmsKey(stack, jsii.String("aws-kms"), &kms.KmsKeyConfig{
		Description:       jsii.String("Create KMS encryption for creating encryption on volume"),
		EnableKeyRotation: true,
		Policy:            jsii.String(makePolicy(*awsAccount.Id())),
		Tags: &map[string]*string{
			"Name":    jsii.String("CDKtf-Golang-Demo-KMS-key"),
			"Team":    jsii.String(env),
			"Company": jsii.String(company),
		},
	})

	kms.NewKmsAlias(stack, jsii.String("aws-kms-alias"), &kms.KmsAliasConfig{
		Name:        jsii.String(fmt.Sprintf("alias/%s", env)),
		TargetKeyId: awsKmsKey.Id(),
	})

	cdktf.NewTerraformOutput(stack, jsii.String("aws-kms-id"), &cdktf.TerraformOutputConfig{
		Value: awsKmsKey.Id(),
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	stack := NewMyStack(app, "cdktf-go-aws-kms")
	cdktf.NewRemoteBackend(stack, &cdktf.RemoteBackendProps{
		Hostname:     jsii.String("app.terraform.io"),
		Organization: jsii.String("jigsaw373"),
		Workspaces:   cdktf.NewNamedRemoteWorkspace(jsii.String("cdktf-go-aws-kms")),
	})

	app.Synth()
}
