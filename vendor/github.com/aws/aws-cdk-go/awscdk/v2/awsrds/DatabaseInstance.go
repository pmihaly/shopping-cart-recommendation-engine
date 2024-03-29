package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
)

// A database instance.
//
// Example:
//   var vpc iVpc
//
//
//   instance1 := rds.NewDatabaseInstance(this, jsii.String("PostgresInstance1"), &DatabaseInstanceProps{
//   	Engine: rds.DatabaseInstanceEngine_POSTGRES(),
//   	// Generate the secret with admin username `postgres` and random password
//   	Credentials: rds.Credentials_FromGeneratedSecret(jsii.String("postgres")),
//   	Vpc: Vpc,
//   })
//   // Templated secret with username and password fields
//   templatedSecret := secretsmanager.NewSecret(this, jsii.String("TemplatedSecret"), &SecretProps{
//   	GenerateSecretString: &SecretStringGenerator{
//   		SecretStringTemplate: jSON.stringify(map[string]*string{
//   			"username": jsii.String("postgres"),
//   		}),
//   		GenerateStringKey: jsii.String("password"),
//   		ExcludeCharacters: jsii.String("/@\""),
//   	},
//   })
//   // Using the templated secret as credentials
//   instance2 := rds.NewDatabaseInstance(this, jsii.String("PostgresInstance2"), &DatabaseInstanceProps{
//   	Engine: rds.DatabaseInstanceEngine_POSTGRES(),
//   	Credentials: map[string]interface{}{
//   		"username": templatedSecret.secretValueFromJson(jsii.String("username")).toString(),
//   		"password": templatedSecret.secretValueFromJson(jsii.String("password")),
//   	},
//   	Vpc: Vpc,
//   })
//
type DatabaseInstance interface {
	DatabaseInstanceBase
	IDatabaseInstance
	// The log group is created when `cloudwatchLogsExports` is set.
	//
	// Each export value will create a separate log group.
	CloudwatchLogGroups() *map[string]awslogs.ILogGroup
	// Access to network connections.
	Connections() awsec2.Connections
	// The instance endpoint address.
	DbInstanceEndpointAddress() *string
	// The instance endpoint port.
	DbInstanceEndpointPort() *string
	EnableIamAuthentication() *bool
	SetEnableIamAuthentication(val *bool)
	// The engine of this database Instance.
	//
	// May be not known for imported Instances if it wasn't provided explicitly,
	// or for read replicas.
	Engine() IInstanceEngine
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The instance arn.
	InstanceArn() *string
	// The instance endpoint.
	InstanceEndpoint() Endpoint
	// The instance identifier.
	InstanceIdentifier() *string
	// The AWS Region-unique, immutable identifier for the DB instance.
	//
	// This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB instance is accessed.
	InstanceResourceId() *string
	InstanceType() awsec2.InstanceType
	NewCfnProps() *CfnDBInstanceProps
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The AWS Secrets Manager secret attached to the instance.
	Secret() awssecretsmanager.ISecret
	SourceCfnProps() *CfnDBInstanceProps
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The VPC where this database instance is deployed.
	Vpc() awsec2.IVpc
	VpcPlacement() *awsec2.SubnetSelection
	// Add a new db proxy to this instance.
	AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy
	// Adds the multi user rotation to this instance.
	AddRotationMultiUser(id *string, options *RotationMultiUserOptions) awssecretsmanager.SecretRotation
	// Adds the single user rotation of the master password to this instance.
	AddRotationSingleUser(options *RotationSingleUserOptions) awssecretsmanager.SecretRotation
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Renders the secret attachment target specifications.
	AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grant the given identity connection access to the database.
	GrantConnect(grantee awsiam.IGrantable, dbUser *string) awsiam.Grant
	// Return the given named metric for this DBInstance.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The percentage of CPU utilization.
	//
	// Average over 5 minutes.
	MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of database connections in use.
	//
	// Average over 5 minutes.
	MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of available random access memory.
	//
	// Average over 5 minutes.
	MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of available storage space.
	//
	// Average over 5 minutes.
	MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The average number of disk write I/O operations per second.
	//
	// Average over 5 minutes.
	MetricReadIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The average number of disk read I/O operations per second.
	//
	// Average over 5 minutes.
	MetricWriteIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Defines a CloudWatch event rule which triggers for instance events.
	//
	// Use
	// `rule.addEventPattern(pattern)` to specify a filter.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	SetLogRetention()
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for DatabaseInstance
type jsiiProxy_DatabaseInstance struct {
	jsiiProxy_DatabaseInstanceBase
	jsiiProxy_IDatabaseInstance
}

func (j *jsiiProxy_DatabaseInstance) CloudwatchLogGroups() *map[string]awslogs.ILogGroup {
	var returns *map[string]awslogs.ILogGroup
	_jsii_.Get(
		j,
		"cloudwatchLogGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) DbInstanceEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) DbInstanceEndpointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) EnableIamAuthentication() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableIamAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Engine() IInstanceEngine {
	var returns IInstanceEngine
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) InstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) InstanceEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"instanceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) InstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) InstanceResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) InstanceType() awsec2.InstanceType {
	var returns awsec2.InstanceType
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) NewCfnProps() *CfnDBInstanceProps {
	var returns *CfnDBInstanceProps
	_jsii_.Get(
		j,
		"newCfnProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Secret() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) SourceCfnProps() *CfnDBInstanceProps {
	var returns *CfnDBInstanceProps
	_jsii_.Get(
		j,
		"sourceCfnProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) VpcPlacement() *awsec2.SubnetSelection {
	var returns *awsec2.SubnetSelection
	_jsii_.Get(
		j,
		"vpcPlacement",
		&returns,
	)
	return returns
}


func NewDatabaseInstance(scope constructs.Construct, id *string, props *DatabaseInstanceProps) DatabaseInstance {
	_init_.Initialize()

	if err := validateNewDatabaseInstanceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseInstance{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDatabaseInstance_Override(d DatabaseInstance, scope constructs.Construct, id *string, props *DatabaseInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseInstance",
		[]interface{}{scope, id, props},
		d,
	)
}

func (j *jsiiProxy_DatabaseInstance)SetEnableIamAuthentication(val *bool) {
	_jsii_.Set(
		j,
		"enableIamAuthentication",
		val,
	)
}

// Import an existing database instance.
func DatabaseInstance_FromDatabaseInstanceAttributes(scope constructs.Construct, id *string, attrs *DatabaseInstanceAttributes) IDatabaseInstance {
	_init_.Initialize()

	if err := validateDatabaseInstance_FromDatabaseInstanceAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IDatabaseInstance

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstance",
		"fromDatabaseInstanceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DatabaseInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func DatabaseInstance_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateDatabaseInstance_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstance",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func DatabaseInstance_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateDatabaseInstance_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstance",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy {
	if err := d.validateAddProxyParameters(id, options); err != nil {
		panic(err)
	}
	var returns DatabaseProxy

	_jsii_.Invoke(
		d,
		"addProxy",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) AddRotationMultiUser(id *string, options *RotationMultiUserOptions) awssecretsmanager.SecretRotation {
	if err := d.validateAddRotationMultiUserParameters(id, options); err != nil {
		panic(err)
	}
	var returns awssecretsmanager.SecretRotation

	_jsii_.Invoke(
		d,
		"addRotationMultiUser",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) AddRotationSingleUser(options *RotationSingleUserOptions) awssecretsmanager.SecretRotation {
	if err := d.validateAddRotationSingleUserParameters(options); err != nil {
		panic(err)
	}
	var returns awssecretsmanager.SecretRotation

	_jsii_.Invoke(
		d,
		"addRotationSingleUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := d.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (d *jsiiProxy_DatabaseInstance) AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps {
	var returns *awssecretsmanager.SecretAttachmentTargetProps

	_jsii_.Invoke(
		d,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := d.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) GetResourceNameAttribute(nameAttr *string) *string {
	if err := d.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) GrantConnect(grantee awsiam.IGrantable, dbUser *string) awsiam.Grant {
	if err := d.validateGrantConnectParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantConnect",
		[]interface{}{grantee, dbUser},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricCPUUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricDatabaseConnectionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricDatabaseConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricFreeableMemoryParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFreeableMemory",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricFreeStorageSpaceParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFreeStorageSpace",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) MetricReadIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricReadIOPSParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricReadIOPS",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) MetricWriteIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricWriteIOPSParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricWriteIOPS",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := d.validateOnEventParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		d,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) SetLogRetention() {
	_jsii_.InvokeVoid(
		d,
		"setLogRetention",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
