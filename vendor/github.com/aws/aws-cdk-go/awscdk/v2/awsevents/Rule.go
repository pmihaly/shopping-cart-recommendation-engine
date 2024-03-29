package awsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Defines an EventBridge Rule in this stack.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunc"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromInline(jsii.String("exports.handler = handler.toString()")),
//   })
//
//   rule := events.NewRule(this, jsii.String("rule"), &RuleProps{
//   	EventPattern: &EventPattern{
//   		Source: []*string{
//   			jsii.String("aws.ec2"),
//   		},
//   	},
//   })
//
//   queue := sqs.NewQueue(this, jsii.String("Queue"))
//
//   rule.AddTarget(targets.NewLambdaFunction(fn, &LambdaFunctionProps{
//   	DeadLetterQueue: queue,
//   	 // Optional: add a dead letter queue
//   	MaxEventAge: awscdk.Duration_Hours(jsii.Number(2)),
//   	 // Optional: set the maxEventAge retry policy
//   	RetryAttempts: jsii.Number(2),
//   }))
//
type Rule interface {
	awscdk.Resource
	IRule
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
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
	// The value of the event rule Amazon Resource Name (ARN), such as arn:aws:events:us-east-2:123456789012:rule/example.
	RuleArn() *string
	// The name event rule.
	RuleName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Adds an event pattern filter to this rule.
	//
	// If a pattern was already specified,
	// these values are merged into the existing pattern.
	//
	// For example, if the rule already contains the pattern:
	//
	//    {
	//      "resources": [ "r1" ],
	//      "detail": {
	//        "hello": [ 1 ]
	//      }
	//    }
	//
	// And `addEventPattern` is called with the pattern:
	//
	//    {
	//      "resources": [ "r2" ],
	//      "detail": {
	//        "foo": [ "bar" ]
	//      }
	//    }
	//
	// The resulting event pattern will be:
	//
	//    {
	//      "resources": [ "r1", "r2" ],
	//      "detail": {
	//        "hello": [ 1 ],
	//        "foo": [ "bar" ]
	//      }
	// }.
	AddEventPattern(eventPattern *EventPattern)
	// Adds a target to the rule. The abstract class RuleTarget can be extended to define new targets.
	//
	// No-op if target is undefined.
	AddTarget(target IRuleTarget)
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
	// Returns a string representation of this construct.
	ToString() *string
	ValidateRule() *[]*string
}

// The jsii proxy struct for Rule
type jsiiProxy_Rule struct {
	internal.Type__awscdkResource
	jsiiProxy_IRule
}

func (j *jsiiProxy_Rule) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rule) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rule) RuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rule) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewRule(scope constructs.Construct, id *string, props *RuleProps) Rule {
	_init_.Initialize()

	if err := validateNewRuleParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Rule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events.Rule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewRule_Override(r Rule, scope constructs.Construct, id *string, props *RuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events.Rule",
		[]interface{}{scope, id, props},
		r,
	)
}

// Import an existing EventBridge Rule provided an ARN.
func Rule_FromEventRuleArn(scope constructs.Construct, id *string, eventRuleArn *string) IRule {
	_init_.Initialize()

	if err := validateRule_FromEventRuleArnParameters(scope, id, eventRuleArn); err != nil {
		panic(err)
	}
	var returns IRule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Rule",
		"fromEventRuleArn",
		[]interface{}{scope, id, eventRuleArn},
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
func Rule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Rule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Rule_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateRule_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Rule",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Rule_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateRule_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Rule",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rule) AddEventPattern(eventPattern *EventPattern) {
	if err := r.validateAddEventPatternParameters(eventPattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addEventPattern",
		[]interface{}{eventPattern},
	)
}

func (r *jsiiProxy_Rule) AddTarget(target IRuleTarget) {
	_jsii_.InvokeVoid(
		r,
		"addTarget",
		[]interface{}{target},
	)
}

func (r *jsiiProxy_Rule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := r.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_Rule) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rule) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := r.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rule) GetResourceNameAttribute(nameAttr *string) *string {
	if err := r.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rule) ValidateRule() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"validateRule",
		nil, // no parameters
		&returns,
	)

	return returns
}

