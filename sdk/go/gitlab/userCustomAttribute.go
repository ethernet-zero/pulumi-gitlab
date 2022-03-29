// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `UserCustomAttribute` resource allows to manage custom attributes for a user.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/custom_attributes.html)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gitlab/sdk/v4/go/gitlab"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := gitlab.NewUserCustomAttribute(ctx, "attr", &gitlab.UserCustomAttributeArgs{
// 			Key:   pulumi.String("location"),
// 			User:  pulumi.Int(42),
// 			Value: pulumi.String("Greenland"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// # You can import a user custom attribute using an id made up of `{user-id}:{key}`, e.g.
//
// ```sh
//  $ pulumi import gitlab:index/userCustomAttribute:UserCustomAttribute attr 42:location
// ```
type UserCustomAttribute struct {
	pulumi.CustomResourceState

	// Key for the Custom Attribute.
	Key pulumi.StringOutput `pulumi:"key"`
	// The id of the user.
	User pulumi.IntOutput `pulumi:"user"`
	// Value for the Custom Attribute.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewUserCustomAttribute registers a new resource with the given unique name, arguments, and options.
func NewUserCustomAttribute(ctx *pulumi.Context,
	name string, args *UserCustomAttributeArgs, opts ...pulumi.ResourceOption) (*UserCustomAttribute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.User == nil {
		return nil, errors.New("invalid value for required argument 'User'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	var resource UserCustomAttribute
	err := ctx.RegisterResource("gitlab:index/userCustomAttribute:UserCustomAttribute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserCustomAttribute gets an existing UserCustomAttribute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserCustomAttribute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserCustomAttributeState, opts ...pulumi.ResourceOption) (*UserCustomAttribute, error) {
	var resource UserCustomAttribute
	err := ctx.ReadResource("gitlab:index/userCustomAttribute:UserCustomAttribute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserCustomAttribute resources.
type userCustomAttributeState struct {
	// Key for the Custom Attribute.
	Key *string `pulumi:"key"`
	// The id of the user.
	User *int `pulumi:"user"`
	// Value for the Custom Attribute.
	Value *string `pulumi:"value"`
}

type UserCustomAttributeState struct {
	// Key for the Custom Attribute.
	Key pulumi.StringPtrInput
	// The id of the user.
	User pulumi.IntPtrInput
	// Value for the Custom Attribute.
	Value pulumi.StringPtrInput
}

func (UserCustomAttributeState) ElementType() reflect.Type {
	return reflect.TypeOf((*userCustomAttributeState)(nil)).Elem()
}

type userCustomAttributeArgs struct {
	// Key for the Custom Attribute.
	Key string `pulumi:"key"`
	// The id of the user.
	User int `pulumi:"user"`
	// Value for the Custom Attribute.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a UserCustomAttribute resource.
type UserCustomAttributeArgs struct {
	// Key for the Custom Attribute.
	Key pulumi.StringInput
	// The id of the user.
	User pulumi.IntInput
	// Value for the Custom Attribute.
	Value pulumi.StringInput
}

func (UserCustomAttributeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userCustomAttributeArgs)(nil)).Elem()
}

type UserCustomAttributeInput interface {
	pulumi.Input

	ToUserCustomAttributeOutput() UserCustomAttributeOutput
	ToUserCustomAttributeOutputWithContext(ctx context.Context) UserCustomAttributeOutput
}

func (*UserCustomAttribute) ElementType() reflect.Type {
	return reflect.TypeOf((**UserCustomAttribute)(nil)).Elem()
}

func (i *UserCustomAttribute) ToUserCustomAttributeOutput() UserCustomAttributeOutput {
	return i.ToUserCustomAttributeOutputWithContext(context.Background())
}

func (i *UserCustomAttribute) ToUserCustomAttributeOutputWithContext(ctx context.Context) UserCustomAttributeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserCustomAttributeOutput)
}

// UserCustomAttributeArrayInput is an input type that accepts UserCustomAttributeArray and UserCustomAttributeArrayOutput values.
// You can construct a concrete instance of `UserCustomAttributeArrayInput` via:
//
//          UserCustomAttributeArray{ UserCustomAttributeArgs{...} }
type UserCustomAttributeArrayInput interface {
	pulumi.Input

	ToUserCustomAttributeArrayOutput() UserCustomAttributeArrayOutput
	ToUserCustomAttributeArrayOutputWithContext(context.Context) UserCustomAttributeArrayOutput
}

type UserCustomAttributeArray []UserCustomAttributeInput

func (UserCustomAttributeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserCustomAttribute)(nil)).Elem()
}

func (i UserCustomAttributeArray) ToUserCustomAttributeArrayOutput() UserCustomAttributeArrayOutput {
	return i.ToUserCustomAttributeArrayOutputWithContext(context.Background())
}

func (i UserCustomAttributeArray) ToUserCustomAttributeArrayOutputWithContext(ctx context.Context) UserCustomAttributeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserCustomAttributeArrayOutput)
}

// UserCustomAttributeMapInput is an input type that accepts UserCustomAttributeMap and UserCustomAttributeMapOutput values.
// You can construct a concrete instance of `UserCustomAttributeMapInput` via:
//
//          UserCustomAttributeMap{ "key": UserCustomAttributeArgs{...} }
type UserCustomAttributeMapInput interface {
	pulumi.Input

	ToUserCustomAttributeMapOutput() UserCustomAttributeMapOutput
	ToUserCustomAttributeMapOutputWithContext(context.Context) UserCustomAttributeMapOutput
}

type UserCustomAttributeMap map[string]UserCustomAttributeInput

func (UserCustomAttributeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserCustomAttribute)(nil)).Elem()
}

func (i UserCustomAttributeMap) ToUserCustomAttributeMapOutput() UserCustomAttributeMapOutput {
	return i.ToUserCustomAttributeMapOutputWithContext(context.Background())
}

func (i UserCustomAttributeMap) ToUserCustomAttributeMapOutputWithContext(ctx context.Context) UserCustomAttributeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserCustomAttributeMapOutput)
}

type UserCustomAttributeOutput struct{ *pulumi.OutputState }

func (UserCustomAttributeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserCustomAttribute)(nil)).Elem()
}

func (o UserCustomAttributeOutput) ToUserCustomAttributeOutput() UserCustomAttributeOutput {
	return o
}

func (o UserCustomAttributeOutput) ToUserCustomAttributeOutputWithContext(ctx context.Context) UserCustomAttributeOutput {
	return o
}

type UserCustomAttributeArrayOutput struct{ *pulumi.OutputState }

func (UserCustomAttributeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserCustomAttribute)(nil)).Elem()
}

func (o UserCustomAttributeArrayOutput) ToUserCustomAttributeArrayOutput() UserCustomAttributeArrayOutput {
	return o
}

func (o UserCustomAttributeArrayOutput) ToUserCustomAttributeArrayOutputWithContext(ctx context.Context) UserCustomAttributeArrayOutput {
	return o
}

func (o UserCustomAttributeArrayOutput) Index(i pulumi.IntInput) UserCustomAttributeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserCustomAttribute {
		return vs[0].([]*UserCustomAttribute)[vs[1].(int)]
	}).(UserCustomAttributeOutput)
}

type UserCustomAttributeMapOutput struct{ *pulumi.OutputState }

func (UserCustomAttributeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserCustomAttribute)(nil)).Elem()
}

func (o UserCustomAttributeMapOutput) ToUserCustomAttributeMapOutput() UserCustomAttributeMapOutput {
	return o
}

func (o UserCustomAttributeMapOutput) ToUserCustomAttributeMapOutputWithContext(ctx context.Context) UserCustomAttributeMapOutput {
	return o
}

func (o UserCustomAttributeMapOutput) MapIndex(k pulumi.StringInput) UserCustomAttributeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserCustomAttribute {
		return vs[0].(map[string]*UserCustomAttribute)[vs[1].(string)]
	}).(UserCustomAttributeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserCustomAttributeInput)(nil)).Elem(), &UserCustomAttribute{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserCustomAttributeArrayInput)(nil)).Elem(), UserCustomAttributeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserCustomAttributeMapInput)(nil)).Elem(), UserCustomAttributeMap{})
	pulumi.RegisterOutputType(UserCustomAttributeOutput{})
	pulumi.RegisterOutputType(UserCustomAttributeArrayOutput{})
	pulumi.RegisterOutputType(UserCustomAttributeMapOutput{})
}
