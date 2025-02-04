// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getGroupHooks` data source allows to retrieve details about hooks in a group.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#list-group-hooks)
func GetGroupHooks(ctx *pulumi.Context, args *GetGroupHooksArgs, opts ...pulumi.InvokeOption) (*GetGroupHooksResult, error) {
	var rv GetGroupHooksResult
	err := ctx.Invoke("gitlab:index/getGroupHooks:getGroupHooks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroupHooks.
type GetGroupHooksArgs struct {
	// The ID or full path of the group.
	Group string `pulumi:"group"`
}

// A collection of values returned by getGroupHooks.
type GetGroupHooksResult struct {
	// The ID or full path of the group.
	Group string `pulumi:"group"`
	// The list of hooks.
	Hooks []GetGroupHooksHook `pulumi:"hooks"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}

func GetGroupHooksOutput(ctx *pulumi.Context, args GetGroupHooksOutputArgs, opts ...pulumi.InvokeOption) GetGroupHooksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGroupHooksResult, error) {
			args := v.(GetGroupHooksArgs)
			r, err := GetGroupHooks(ctx, &args, opts...)
			var s GetGroupHooksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGroupHooksResultOutput)
}

// A collection of arguments for invoking getGroupHooks.
type GetGroupHooksOutputArgs struct {
	// The ID or full path of the group.
	Group pulumi.StringInput `pulumi:"group"`
}

func (GetGroupHooksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupHooksArgs)(nil)).Elem()
}

// A collection of values returned by getGroupHooks.
type GetGroupHooksResultOutput struct{ *pulumi.OutputState }

func (GetGroupHooksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGroupHooksResult)(nil)).Elem()
}

func (o GetGroupHooksResultOutput) ToGetGroupHooksResultOutput() GetGroupHooksResultOutput {
	return o
}

func (o GetGroupHooksResultOutput) ToGetGroupHooksResultOutputWithContext(ctx context.Context) GetGroupHooksResultOutput {
	return o
}

// The ID or full path of the group.
func (o GetGroupHooksResultOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupHooksResult) string { return v.Group }).(pulumi.StringOutput)
}

// The list of hooks.
func (o GetGroupHooksResultOutput) Hooks() GetGroupHooksHookArrayOutput {
	return o.ApplyT(func(v GetGroupHooksResult) []GetGroupHooksHook { return v.Hooks }).(GetGroupHooksHookArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGroupHooksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGroupHooksResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGroupHooksResultOutput{})
}
