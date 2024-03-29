package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a public VPC subnet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicSubnet := awscdk.Aws_ec2.NewPublicSubnet(this, jsii.String("MyPublicSubnet"), &PublicSubnetProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	CidrBlock: jsii.String("cidrBlock"),
//   	VpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	AssignIpv6AddressOnCreation: jsii.Boolean(false),
//   	Ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   	MapPublicIpOnLaunch: jsii.Boolean(false),
//   })
//
type PublicSubnet interface {
	Subnet
	IPublicSubnet
	// The Availability Zone the subnet is located in.
	AvailabilityZone() *string
	// Parts of this VPC subnet.
	DependencyElements() *[]constructs.IDependable
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// Dependable that can be depended upon to force internet connectivity established on the VPC.
	InternetConnectivityEstablished() constructs.IDependable
	// The IPv4 CIDR block for this subnet.
	Ipv4CidrBlock() *string
	// Network ACL associated with this Subnet.
	//
	// Upon creation, this is the default ACL which allows all traffic, except
	// explicit DENY entries that you add.
	//
	// You can replace it with a custom ACL which denies all traffic except
	// the explicit ALLOW entries that you add by creating a `NetworkAcl`
	// object and calling `associateNetworkAcl()`.
	NetworkAcl() INetworkAcl
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
	// The routeTableId attached to this subnet.
	RouteTable() IRouteTable
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	SubnetAvailabilityZone() *string
	// The subnetId for this particular subnet.
	SubnetId() *string
	SubnetIpv6CidrBlocks() *[]*string
	SubnetNetworkAclAssociationId() *string
	// The Amazon Resource Name (ARN) of the Outpost for this subnet (if one exists).
	SubnetOutpostArn() *string
	SubnetVpcId() *string
	// Create a default route that points to a passed IGW, with a dependency on the IGW's attachment to the VPC.
	AddDefaultInternetRoute(gatewayId *string, gatewayAttachment constructs.IDependable)
	// Adds an entry to this subnets route table that points to the passed NATGatewayId.
	AddDefaultNatRoute(natGatewayId *string)
	// Create a default IPv6 route that points to a passed EIGW.
	AddIpv6DefaultEgressOnlyInternetRoute(gatewayId *string)
	// Create a default IPv6 route that points to a passed IGW.
	AddIpv6DefaultInternetRoute(gatewayId *string)
	// Adds an entry to this subnets route table that points to the passed NATGatewayId.
	//
	// Uses the known 64:ff9b::/96 prefix.
	AddIpv6Nat64Route(natGatewayId *string)
	// Creates a new managed NAT gateway attached to this public subnet.
	//
	// Also adds the EIP for the managed NAT.
	//
	// Returns: A ref to the the NAT Gateway ID.
	AddNatGateway(eipAllocationId *string) CfnNatGateway
	// Adds an entry to this subnets route table.
	AddRoute(id *string, options *AddRouteOptions)
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
	// Associate a Network ACL with this subnet.
	AssociateNetworkAcl(id *string, networkAcl INetworkAcl)
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
}

// The jsii proxy struct for PublicSubnet
type jsiiProxy_PublicSubnet struct {
	jsiiProxy_Subnet
	jsiiProxy_IPublicSubnet
}

func (j *jsiiProxy_PublicSubnet) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicSubnet) DependencyElements() *[]constructs.IDependable {
	var returns *[]constructs.IDependable
	_jsii_.Get(
		j,
		"dependencyElements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicSubnet) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicSubnet) InternetConnectivityEstablished() constructs.IDependable {
	var returns constructs.IDependable
	_jsii_.Get(
		j,
		"internetConnectivityEstablished",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicSubnet) Ipv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicSubnet) NetworkAcl() INetworkAcl {
	var returns INetworkAcl
	_jsii_.Get(
		j,
		"networkAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicSubnet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicSubnet) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicSubnet) RouteTable() IRouteTable {
	var returns IRouteTable
	_jsii_.Get(
		j,
		"routeTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicSubnet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicSubnet) SubnetAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicSubnet) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicSubnet) SubnetIpv6CidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIpv6CidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicSubnet) SubnetNetworkAclAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetNetworkAclAssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicSubnet) SubnetOutpostArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetOutpostArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicSubnet) SubnetVpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetVpcId",
		&returns,
	)
	return returns
}


func NewPublicSubnet(scope constructs.Construct, id *string, props *PublicSubnetProps) PublicSubnet {
	_init_.Initialize()

	if err := validateNewPublicSubnetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PublicSubnet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.PublicSubnet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPublicSubnet_Override(p PublicSubnet, scope constructs.Construct, id *string, props *PublicSubnetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.PublicSubnet",
		[]interface{}{scope, id, props},
		p,
	)
}

func PublicSubnet_FromPublicSubnetAttributes(scope constructs.Construct, id *string, attrs *PublicSubnetAttributes) IPublicSubnet {
	_init_.Initialize()

	if err := validatePublicSubnet_FromPublicSubnetAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IPublicSubnet

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.PublicSubnet",
		"fromPublicSubnetAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

func PublicSubnet_FromSubnetAttributes(scope constructs.Construct, id *string, attrs *SubnetAttributes) ISubnet {
	_init_.Initialize()

	if err := validatePublicSubnet_FromSubnetAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns ISubnet

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.PublicSubnet",
		"fromSubnetAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import existing subnet from id.
func PublicSubnet_FromSubnetId(scope constructs.Construct, id *string, subnetId *string) ISubnet {
	_init_.Initialize()

	if err := validatePublicSubnet_FromSubnetIdParameters(scope, id, subnetId); err != nil {
		panic(err)
	}
	var returns ISubnet

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.PublicSubnet",
		"fromSubnetId",
		[]interface{}{scope, id, subnetId},
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
func PublicSubnet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePublicSubnet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.PublicSubnet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func PublicSubnet_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validatePublicSubnet_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.PublicSubnet",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func PublicSubnet_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validatePublicSubnet_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.PublicSubnet",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func PublicSubnet_IsVpcSubnet(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePublicSubnet_IsVpcSubnetParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.PublicSubnet",
		"isVpcSubnet",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicSubnet) AddDefaultInternetRoute(gatewayId *string, gatewayAttachment constructs.IDependable) {
	if err := p.validateAddDefaultInternetRouteParameters(gatewayId, gatewayAttachment); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addDefaultInternetRoute",
		[]interface{}{gatewayId, gatewayAttachment},
	)
}

func (p *jsiiProxy_PublicSubnet) AddDefaultNatRoute(natGatewayId *string) {
	if err := p.validateAddDefaultNatRouteParameters(natGatewayId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addDefaultNatRoute",
		[]interface{}{natGatewayId},
	)
}

func (p *jsiiProxy_PublicSubnet) AddIpv6DefaultEgressOnlyInternetRoute(gatewayId *string) {
	if err := p.validateAddIpv6DefaultEgressOnlyInternetRouteParameters(gatewayId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addIpv6DefaultEgressOnlyInternetRoute",
		[]interface{}{gatewayId},
	)
}

func (p *jsiiProxy_PublicSubnet) AddIpv6DefaultInternetRoute(gatewayId *string) {
	if err := p.validateAddIpv6DefaultInternetRouteParameters(gatewayId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addIpv6DefaultInternetRoute",
		[]interface{}{gatewayId},
	)
}

func (p *jsiiProxy_PublicSubnet) AddIpv6Nat64Route(natGatewayId *string) {
	if err := p.validateAddIpv6Nat64RouteParameters(natGatewayId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addIpv6Nat64Route",
		[]interface{}{natGatewayId},
	)
}

func (p *jsiiProxy_PublicSubnet) AddNatGateway(eipAllocationId *string) CfnNatGateway {
	var returns CfnNatGateway

	_jsii_.Invoke(
		p,
		"addNatGateway",
		[]interface{}{eipAllocationId},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicSubnet) AddRoute(id *string, options *AddRouteOptions) {
	if err := p.validateAddRouteParameters(id, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addRoute",
		[]interface{}{id, options},
	)
}

func (p *jsiiProxy_PublicSubnet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := p.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (p *jsiiProxy_PublicSubnet) AssociateNetworkAcl(id *string, networkAcl INetworkAcl) {
	if err := p.validateAssociateNetworkAclParameters(id, networkAcl); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"associateNetworkAcl",
		[]interface{}{id, networkAcl},
	)
}

func (p *jsiiProxy_PublicSubnet) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicSubnet) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := p.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicSubnet) GetResourceNameAttribute(nameAttr *string) *string {
	if err := p.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicSubnet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

