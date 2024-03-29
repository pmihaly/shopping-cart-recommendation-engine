//go:build ignore_iac

package main

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type IacStackProps struct {
	awscdk.StackProps
}

func NewIacStack(scope constructs.Construct, id string, props *IacStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	vpc := awsec2.NewVpc(stack, jsii.String("Vpc"), &awsec2.VpcProps{
		IpAddresses: awsec2.IpAddresses_Cidr(jsii.String("10.0.0.0/16")),
	})

	dbSecret := awssecretsmanager.NewSecret(stack, jsii.String("TemplatedSecret"), &awssecretsmanager.SecretProps{
		GenerateSecretString: &awssecretsmanager.SecretStringGenerator{
			SecretStringTemplate: jsii.String(fmt.Sprintf("{\"username\":\"%s\"}", "postgres")),
			GenerateStringKey:    jsii.String("password"),
			ExcludeCharacters:    jsii.String("/@\""),
		},
	})

	dbInstance := awsrds.NewDatabaseInstance(stack, jsii.String("PostgresInstance1"), &awsrds.DatabaseInstanceProps{
		Engine:           awsrds.DatabaseInstanceEngine_POSTGRES(),
		Credentials:      awsrds.Credentials_FromSecret(dbSecret, jsii.String("postgres")),
		Vpc:              vpc,
		AllocatedStorage: jsii.Number(20),
		DatabaseName:     jsii.String("CartRecomm"),
		InstanceType:     awsec2.InstanceType_Of(awsec2.InstanceClass_T4G, awsec2.InstanceSize_MICRO),
	})

	lambdaSecurityGroup := awsec2.NewSecurityGroup(stack, jsii.String("LambdaSecurityGroup"), &awsec2.SecurityGroupProps{
		Vpc: vpc,
	})

	dbInstance.Connections().AllowFrom(lambdaSecurityGroup, awsec2.Port_AllTraffic(), jsii.String("Allow inbound traffic from Lambda"))

	lambda := awslambda.NewFunction(stack, jsii.String("MyLambdaFunction"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Code:    awslambda.Code_FromAsset(jsii.String("result/lambda.zip"), &awss3assets.AssetOptions{}),
		Handler: jsii.String("bootstrap.main"),
		Vpc:     vpc,
		Environment: &map[string]*string{
			*jsii.String("DATABASE_URL"): jsii.String(fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
				*dbSecret.SecretValueFromJson(jsii.String("username")).UnsafeUnwrap(),
				*dbSecret.SecretValueFromJson(jsii.String("password")).UnsafeUnwrap(),
				*jsii.String("postgres"),
				*dbInstance.DbInstanceEndpointAddress(),
				dbInstance.DbInstanceEndpointPort(),
				*jsii.String("cart-recommendation"),
			)),
		},
		SecurityGroups: &[]awsec2.ISecurityGroup{lambdaSecurityGroup},
	})

	api := awsapigateway.NewLambdaRestApi(stack, jsii.String("CartRecommApi"), &awsapigateway.LambdaRestApiProps{
		Handler: lambda,
		Proxy:   jsii.Bool(true),
	})

	api.Root().AddMethod(jsii.String("GET"), awsapigateway.NewLambdaIntegration(lambda, &awsapigateway.LambdaIntegrationOptions{
		PassthroughBehavior: awsapigateway.PassthroughBehavior_WHEN_NO_MATCH,
		RequestTemplates: &map[string]*string{
			*jsii.String("application/json"): jsii.String("{\"statusCode\":200}"),
		},
	}), &awsapigateway.MethodOptions{
		AuthorizationType: awsapigateway.AuthorizationType_NONE,
		ApiKeyRequired:    jsii.Bool(false),
	})

	awscdk.NewCfnOutput(stack, jsii.String("ApiUrl"), &awscdk.CfnOutputProps{
		Value: api.Url(),
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewIacStack(app, "CartRecommStack", &IacStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String("353435981812"),
		Region:  jsii.String("eu-central-1"),
	}
}
