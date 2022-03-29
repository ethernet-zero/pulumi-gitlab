// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `DeployToken` resource allows to manage the lifecycle of group and project deploy tokens.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/deploy_tokens.html)
type DeployToken struct {
	pulumi.CustomResourceState

	// Time the token will expire it, RFC3339 format. Will not expire per default.
	ExpiresAt pulumi.StringPtrOutput `pulumi:"expiresAt"`
	// The name or id of the group to add the deploy token to.
	Group pulumi.StringPtrOutput `pulumi:"group"`
	// A name to describe the deploy token with.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name or id of the project to add the deploy token to.
	Project pulumi.StringPtrOutput `pulumi:"project"`
	// Valid values: `readRepository`, `readRegistry`, `readPackageRegistry`, `writeRegistry`, `writePackageRegistry`.
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// The secret token. This is only populated when creating a new deploy token.
	Token pulumi.StringOutput `pulumi:"token"`
	// A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewDeployToken registers a new resource with the given unique name, arguments, and options.
func NewDeployToken(ctx *pulumi.Context,
	name string, args *DeployTokenArgs, opts ...pulumi.ResourceOption) (*DeployToken, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scopes == nil {
		return nil, errors.New("invalid value for required argument 'Scopes'")
	}
	var resource DeployToken
	err := ctx.RegisterResource("gitlab:index/deployToken:DeployToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeployToken gets an existing DeployToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeployToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeployTokenState, opts ...pulumi.ResourceOption) (*DeployToken, error) {
	var resource DeployToken
	err := ctx.ReadResource("gitlab:index/deployToken:DeployToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeployToken resources.
type deployTokenState struct {
	// Time the token will expire it, RFC3339 format. Will not expire per default.
	ExpiresAt *string `pulumi:"expiresAt"`
	// The name or id of the group to add the deploy token to.
	Group *string `pulumi:"group"`
	// A name to describe the deploy token with.
	Name *string `pulumi:"name"`
	// The name or id of the project to add the deploy token to.
	Project *string `pulumi:"project"`
	// Valid values: `readRepository`, `readRegistry`, `readPackageRegistry`, `writeRegistry`, `writePackageRegistry`.
	Scopes []string `pulumi:"scopes"`
	// The secret token. This is only populated when creating a new deploy token.
	Token *string `pulumi:"token"`
	// A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
	Username *string `pulumi:"username"`
}

type DeployTokenState struct {
	// Time the token will expire it, RFC3339 format. Will not expire per default.
	ExpiresAt pulumi.StringPtrInput
	// The name or id of the group to add the deploy token to.
	Group pulumi.StringPtrInput
	// A name to describe the deploy token with.
	Name pulumi.StringPtrInput
	// The name or id of the project to add the deploy token to.
	Project pulumi.StringPtrInput
	// Valid values: `readRepository`, `readRegistry`, `readPackageRegistry`, `writeRegistry`, `writePackageRegistry`.
	Scopes pulumi.StringArrayInput
	// The secret token. This is only populated when creating a new deploy token.
	Token pulumi.StringPtrInput
	// A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
	Username pulumi.StringPtrInput
}

func (DeployTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*deployTokenState)(nil)).Elem()
}

type deployTokenArgs struct {
	// Time the token will expire it, RFC3339 format. Will not expire per default.
	ExpiresAt *string `pulumi:"expiresAt"`
	// The name or id of the group to add the deploy token to.
	Group *string `pulumi:"group"`
	// A name to describe the deploy token with.
	Name *string `pulumi:"name"`
	// The name or id of the project to add the deploy token to.
	Project *string `pulumi:"project"`
	// Valid values: `readRepository`, `readRegistry`, `readPackageRegistry`, `writeRegistry`, `writePackageRegistry`.
	Scopes []string `pulumi:"scopes"`
	// A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a DeployToken resource.
type DeployTokenArgs struct {
	// Time the token will expire it, RFC3339 format. Will not expire per default.
	ExpiresAt pulumi.StringPtrInput
	// The name or id of the group to add the deploy token to.
	Group pulumi.StringPtrInput
	// A name to describe the deploy token with.
	Name pulumi.StringPtrInput
	// The name or id of the project to add the deploy token to.
	Project pulumi.StringPtrInput
	// Valid values: `readRepository`, `readRegistry`, `readPackageRegistry`, `writeRegistry`, `writePackageRegistry`.
	Scopes pulumi.StringArrayInput
	// A username for the deploy token. Default is `gitlab+deploy-token-{n}`.
	Username pulumi.StringPtrInput
}

func (DeployTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deployTokenArgs)(nil)).Elem()
}

type DeployTokenInput interface {
	pulumi.Input

	ToDeployTokenOutput() DeployTokenOutput
	ToDeployTokenOutputWithContext(ctx context.Context) DeployTokenOutput
}

func (*DeployToken) ElementType() reflect.Type {
	return reflect.TypeOf((**DeployToken)(nil)).Elem()
}

func (i *DeployToken) ToDeployTokenOutput() DeployTokenOutput {
	return i.ToDeployTokenOutputWithContext(context.Background())
}

func (i *DeployToken) ToDeployTokenOutputWithContext(ctx context.Context) DeployTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployTokenOutput)
}

// DeployTokenArrayInput is an input type that accepts DeployTokenArray and DeployTokenArrayOutput values.
// You can construct a concrete instance of `DeployTokenArrayInput` via:
//
//          DeployTokenArray{ DeployTokenArgs{...} }
type DeployTokenArrayInput interface {
	pulumi.Input

	ToDeployTokenArrayOutput() DeployTokenArrayOutput
	ToDeployTokenArrayOutputWithContext(context.Context) DeployTokenArrayOutput
}

type DeployTokenArray []DeployTokenInput

func (DeployTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeployToken)(nil)).Elem()
}

func (i DeployTokenArray) ToDeployTokenArrayOutput() DeployTokenArrayOutput {
	return i.ToDeployTokenArrayOutputWithContext(context.Background())
}

func (i DeployTokenArray) ToDeployTokenArrayOutputWithContext(ctx context.Context) DeployTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployTokenArrayOutput)
}

// DeployTokenMapInput is an input type that accepts DeployTokenMap and DeployTokenMapOutput values.
// You can construct a concrete instance of `DeployTokenMapInput` via:
//
//          DeployTokenMap{ "key": DeployTokenArgs{...} }
type DeployTokenMapInput interface {
	pulumi.Input

	ToDeployTokenMapOutput() DeployTokenMapOutput
	ToDeployTokenMapOutputWithContext(context.Context) DeployTokenMapOutput
}

type DeployTokenMap map[string]DeployTokenInput

func (DeployTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeployToken)(nil)).Elem()
}

func (i DeployTokenMap) ToDeployTokenMapOutput() DeployTokenMapOutput {
	return i.ToDeployTokenMapOutputWithContext(context.Background())
}

func (i DeployTokenMap) ToDeployTokenMapOutputWithContext(ctx context.Context) DeployTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeployTokenMapOutput)
}

type DeployTokenOutput struct{ *pulumi.OutputState }

func (DeployTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeployToken)(nil)).Elem()
}

func (o DeployTokenOutput) ToDeployTokenOutput() DeployTokenOutput {
	return o
}

func (o DeployTokenOutput) ToDeployTokenOutputWithContext(ctx context.Context) DeployTokenOutput {
	return o
}

type DeployTokenArrayOutput struct{ *pulumi.OutputState }

func (DeployTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeployToken)(nil)).Elem()
}

func (o DeployTokenArrayOutput) ToDeployTokenArrayOutput() DeployTokenArrayOutput {
	return o
}

func (o DeployTokenArrayOutput) ToDeployTokenArrayOutputWithContext(ctx context.Context) DeployTokenArrayOutput {
	return o
}

func (o DeployTokenArrayOutput) Index(i pulumi.IntInput) DeployTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DeployToken {
		return vs[0].([]*DeployToken)[vs[1].(int)]
	}).(DeployTokenOutput)
}

type DeployTokenMapOutput struct{ *pulumi.OutputState }

func (DeployTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeployToken)(nil)).Elem()
}

func (o DeployTokenMapOutput) ToDeployTokenMapOutput() DeployTokenMapOutput {
	return o
}

func (o DeployTokenMapOutput) ToDeployTokenMapOutputWithContext(ctx context.Context) DeployTokenMapOutput {
	return o
}

func (o DeployTokenMapOutput) MapIndex(k pulumi.StringInput) DeployTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DeployToken {
		return vs[0].(map[string]*DeployToken)[vs[1].(string)]
	}).(DeployTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeployTokenInput)(nil)).Elem(), &DeployToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeployTokenArrayInput)(nil)).Elem(), DeployTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeployTokenMapInput)(nil)).Elem(), DeployTokenMap{})
	pulumi.RegisterOutputType(DeployTokenOutput{})
	pulumi.RegisterOutputType(DeployTokenArrayOutput{})
	pulumi.RegisterOutputType(DeployTokenMapOutput{})
}
