// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServicePipelinesEmail struct {
	pulumi.CustomResourceState

	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `defaultAndProtected`. Default is `default`
	BranchesToBeNotified pulumi.StringPtrOutput `pulumi:"branchesToBeNotified"`
	// Notify only broken pipelines. Default is true.
	NotifyOnlyBrokenPipelines pulumi.BoolPtrOutput `pulumi:"notifyOnlyBrokenPipelines"`
	// ID of the project you want to activate integration on.
	Project pulumi.StringOutput `pulumi:"project"`
	// ) email addresses where notifications are sent.
	Recipients pulumi.StringArrayOutput `pulumi:"recipients"`
}

// NewServicePipelinesEmail registers a new resource with the given unique name, arguments, and options.
func NewServicePipelinesEmail(ctx *pulumi.Context,
	name string, args *ServicePipelinesEmailArgs, opts ...pulumi.ResourceOption) (*ServicePipelinesEmail, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.Recipients == nil {
		return nil, errors.New("invalid value for required argument 'Recipients'")
	}
	var resource ServicePipelinesEmail
	err := ctx.RegisterResource("gitlab:index/servicePipelinesEmail:ServicePipelinesEmail", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServicePipelinesEmail gets an existing ServicePipelinesEmail resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicePipelinesEmail(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicePipelinesEmailState, opts ...pulumi.ResourceOption) (*ServicePipelinesEmail, error) {
	var resource ServicePipelinesEmail
	err := ctx.ReadResource("gitlab:index/servicePipelinesEmail:ServicePipelinesEmail", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServicePipelinesEmail resources.
type servicePipelinesEmailState struct {
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `defaultAndProtected`. Default is `default`
	BranchesToBeNotified *string `pulumi:"branchesToBeNotified"`
	// Notify only broken pipelines. Default is true.
	NotifyOnlyBrokenPipelines *bool `pulumi:"notifyOnlyBrokenPipelines"`
	// ID of the project you want to activate integration on.
	Project *string `pulumi:"project"`
	// ) email addresses where notifications are sent.
	Recipients []string `pulumi:"recipients"`
}

type ServicePipelinesEmailState struct {
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `defaultAndProtected`. Default is `default`
	BranchesToBeNotified pulumi.StringPtrInput
	// Notify only broken pipelines. Default is true.
	NotifyOnlyBrokenPipelines pulumi.BoolPtrInput
	// ID of the project you want to activate integration on.
	Project pulumi.StringPtrInput
	// ) email addresses where notifications are sent.
	Recipients pulumi.StringArrayInput
}

func (ServicePipelinesEmailState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePipelinesEmailState)(nil)).Elem()
}

type servicePipelinesEmailArgs struct {
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `defaultAndProtected`. Default is `default`
	BranchesToBeNotified *string `pulumi:"branchesToBeNotified"`
	// Notify only broken pipelines. Default is true.
	NotifyOnlyBrokenPipelines *bool `pulumi:"notifyOnlyBrokenPipelines"`
	// ID of the project you want to activate integration on.
	Project string `pulumi:"project"`
	// ) email addresses where notifications are sent.
	Recipients []string `pulumi:"recipients"`
}

// The set of arguments for constructing a ServicePipelinesEmail resource.
type ServicePipelinesEmailArgs struct {
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, and `defaultAndProtected`. Default is `default`
	BranchesToBeNotified pulumi.StringPtrInput
	// Notify only broken pipelines. Default is true.
	NotifyOnlyBrokenPipelines pulumi.BoolPtrInput
	// ID of the project you want to activate integration on.
	Project pulumi.StringInput
	// ) email addresses where notifications are sent.
	Recipients pulumi.StringArrayInput
}

func (ServicePipelinesEmailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePipelinesEmailArgs)(nil)).Elem()
}

type ServicePipelinesEmailInput interface {
	pulumi.Input

	ToServicePipelinesEmailOutput() ServicePipelinesEmailOutput
	ToServicePipelinesEmailOutputWithContext(ctx context.Context) ServicePipelinesEmailOutput
}

func (*ServicePipelinesEmail) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePipelinesEmail)(nil))
}

func (i *ServicePipelinesEmail) ToServicePipelinesEmailOutput() ServicePipelinesEmailOutput {
	return i.ToServicePipelinesEmailOutputWithContext(context.Background())
}

func (i *ServicePipelinesEmail) ToServicePipelinesEmailOutputWithContext(ctx context.Context) ServicePipelinesEmailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePipelinesEmailOutput)
}

func (i *ServicePipelinesEmail) ToServicePipelinesEmailPtrOutput() ServicePipelinesEmailPtrOutput {
	return i.ToServicePipelinesEmailPtrOutputWithContext(context.Background())
}

func (i *ServicePipelinesEmail) ToServicePipelinesEmailPtrOutputWithContext(ctx context.Context) ServicePipelinesEmailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePipelinesEmailPtrOutput)
}

type ServicePipelinesEmailPtrInput interface {
	pulumi.Input

	ToServicePipelinesEmailPtrOutput() ServicePipelinesEmailPtrOutput
	ToServicePipelinesEmailPtrOutputWithContext(ctx context.Context) ServicePipelinesEmailPtrOutput
}

type servicePipelinesEmailPtrType ServicePipelinesEmailArgs

func (*servicePipelinesEmailPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePipelinesEmail)(nil))
}

func (i *servicePipelinesEmailPtrType) ToServicePipelinesEmailPtrOutput() ServicePipelinesEmailPtrOutput {
	return i.ToServicePipelinesEmailPtrOutputWithContext(context.Background())
}

func (i *servicePipelinesEmailPtrType) ToServicePipelinesEmailPtrOutputWithContext(ctx context.Context) ServicePipelinesEmailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePipelinesEmailPtrOutput)
}

// ServicePipelinesEmailArrayInput is an input type that accepts ServicePipelinesEmailArray and ServicePipelinesEmailArrayOutput values.
// You can construct a concrete instance of `ServicePipelinesEmailArrayInput` via:
//
//          ServicePipelinesEmailArray{ ServicePipelinesEmailArgs{...} }
type ServicePipelinesEmailArrayInput interface {
	pulumi.Input

	ToServicePipelinesEmailArrayOutput() ServicePipelinesEmailArrayOutput
	ToServicePipelinesEmailArrayOutputWithContext(context.Context) ServicePipelinesEmailArrayOutput
}

type ServicePipelinesEmailArray []ServicePipelinesEmailInput

func (ServicePipelinesEmailArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicePipelinesEmail)(nil)).Elem()
}

func (i ServicePipelinesEmailArray) ToServicePipelinesEmailArrayOutput() ServicePipelinesEmailArrayOutput {
	return i.ToServicePipelinesEmailArrayOutputWithContext(context.Background())
}

func (i ServicePipelinesEmailArray) ToServicePipelinesEmailArrayOutputWithContext(ctx context.Context) ServicePipelinesEmailArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePipelinesEmailArrayOutput)
}

// ServicePipelinesEmailMapInput is an input type that accepts ServicePipelinesEmailMap and ServicePipelinesEmailMapOutput values.
// You can construct a concrete instance of `ServicePipelinesEmailMapInput` via:
//
//          ServicePipelinesEmailMap{ "key": ServicePipelinesEmailArgs{...} }
type ServicePipelinesEmailMapInput interface {
	pulumi.Input

	ToServicePipelinesEmailMapOutput() ServicePipelinesEmailMapOutput
	ToServicePipelinesEmailMapOutputWithContext(context.Context) ServicePipelinesEmailMapOutput
}

type ServicePipelinesEmailMap map[string]ServicePipelinesEmailInput

func (ServicePipelinesEmailMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicePipelinesEmail)(nil)).Elem()
}

func (i ServicePipelinesEmailMap) ToServicePipelinesEmailMapOutput() ServicePipelinesEmailMapOutput {
	return i.ToServicePipelinesEmailMapOutputWithContext(context.Background())
}

func (i ServicePipelinesEmailMap) ToServicePipelinesEmailMapOutputWithContext(ctx context.Context) ServicePipelinesEmailMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePipelinesEmailMapOutput)
}

type ServicePipelinesEmailOutput struct{ *pulumi.OutputState }

func (ServicePipelinesEmailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePipelinesEmail)(nil))
}

func (o ServicePipelinesEmailOutput) ToServicePipelinesEmailOutput() ServicePipelinesEmailOutput {
	return o
}

func (o ServicePipelinesEmailOutput) ToServicePipelinesEmailOutputWithContext(ctx context.Context) ServicePipelinesEmailOutput {
	return o
}

func (o ServicePipelinesEmailOutput) ToServicePipelinesEmailPtrOutput() ServicePipelinesEmailPtrOutput {
	return o.ToServicePipelinesEmailPtrOutputWithContext(context.Background())
}

func (o ServicePipelinesEmailOutput) ToServicePipelinesEmailPtrOutputWithContext(ctx context.Context) ServicePipelinesEmailPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServicePipelinesEmail) *ServicePipelinesEmail {
		return &v
	}).(ServicePipelinesEmailPtrOutput)
}

type ServicePipelinesEmailPtrOutput struct{ *pulumi.OutputState }

func (ServicePipelinesEmailPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePipelinesEmail)(nil))
}

func (o ServicePipelinesEmailPtrOutput) ToServicePipelinesEmailPtrOutput() ServicePipelinesEmailPtrOutput {
	return o
}

func (o ServicePipelinesEmailPtrOutput) ToServicePipelinesEmailPtrOutputWithContext(ctx context.Context) ServicePipelinesEmailPtrOutput {
	return o
}

func (o ServicePipelinesEmailPtrOutput) Elem() ServicePipelinesEmailOutput {
	return o.ApplyT(func(v *ServicePipelinesEmail) ServicePipelinesEmail {
		if v != nil {
			return *v
		}
		var ret ServicePipelinesEmail
		return ret
	}).(ServicePipelinesEmailOutput)
}

type ServicePipelinesEmailArrayOutput struct{ *pulumi.OutputState }

func (ServicePipelinesEmailArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServicePipelinesEmail)(nil))
}

func (o ServicePipelinesEmailArrayOutput) ToServicePipelinesEmailArrayOutput() ServicePipelinesEmailArrayOutput {
	return o
}

func (o ServicePipelinesEmailArrayOutput) ToServicePipelinesEmailArrayOutputWithContext(ctx context.Context) ServicePipelinesEmailArrayOutput {
	return o
}

func (o ServicePipelinesEmailArrayOutput) Index(i pulumi.IntInput) ServicePipelinesEmailOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServicePipelinesEmail {
		return vs[0].([]ServicePipelinesEmail)[vs[1].(int)]
	}).(ServicePipelinesEmailOutput)
}

type ServicePipelinesEmailMapOutput struct{ *pulumi.OutputState }

func (ServicePipelinesEmailMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ServicePipelinesEmail)(nil))
}

func (o ServicePipelinesEmailMapOutput) ToServicePipelinesEmailMapOutput() ServicePipelinesEmailMapOutput {
	return o
}

func (o ServicePipelinesEmailMapOutput) ToServicePipelinesEmailMapOutputWithContext(ctx context.Context) ServicePipelinesEmailMapOutput {
	return o
}

func (o ServicePipelinesEmailMapOutput) MapIndex(k pulumi.StringInput) ServicePipelinesEmailOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ServicePipelinesEmail {
		return vs[0].(map[string]ServicePipelinesEmail)[vs[1].(string)]
	}).(ServicePipelinesEmailOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePipelinesEmailInput)(nil)).Elem(), &ServicePipelinesEmail{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePipelinesEmailPtrInput)(nil)).Elem(), &ServicePipelinesEmail{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePipelinesEmailArrayInput)(nil)).Elem(), ServicePipelinesEmailArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePipelinesEmailMapInput)(nil)).Elem(), ServicePipelinesEmailMap{})
	pulumi.RegisterOutputType(ServicePipelinesEmailOutput{})
	pulumi.RegisterOutputType(ServicePipelinesEmailPtrOutput{})
	pulumi.RegisterOutputType(ServicePipelinesEmailArrayOutput{})
	pulumi.RegisterOutputType(ServicePipelinesEmailMapOutput{})
}
