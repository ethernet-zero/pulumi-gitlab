// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `GroupBadge` resource allows to mange the lifecycle of group badges.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/user/project/badges.html#group-badges)
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
// 		foo, err := gitlab.NewGroup(ctx, "foo", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = gitlab.NewGroupBadge(ctx, "example", &gitlab.GroupBadgeArgs{
// 			Group:    foo.ID(),
// 			LinkUrl:  pulumi.String("https://example.com/badge-123"),
// 			ImageUrl: pulumi.String("https://example.com/badge-123.svg"),
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
// # GitLab group badges can be imported using an id made up of `{group_id}:{badge_id}`, e.g.
//
// ```sh
//  $ pulumi import gitlab:index/groupBadge:GroupBadge foo 1:3
// ```
type GroupBadge struct {
	pulumi.CustomResourceState

	// The id of the group to add the badge to.
	Group pulumi.StringOutput `pulumi:"group"`
	// The image url which will be presented on group overview.
	ImageUrl pulumi.StringOutput `pulumi:"imageUrl"`
	// The url linked with the badge.
	LinkUrl pulumi.StringOutput `pulumi:"linkUrl"`
	// The image_url argument rendered (in case of use of placeholders).
	RenderedImageUrl pulumi.StringOutput `pulumi:"renderedImageUrl"`
	// The link_url argument rendered (in case of use of placeholders).
	RenderedLinkUrl pulumi.StringOutput `pulumi:"renderedLinkUrl"`
}

// NewGroupBadge registers a new resource with the given unique name, arguments, and options.
func NewGroupBadge(ctx *pulumi.Context,
	name string, args *GroupBadgeArgs, opts ...pulumi.ResourceOption) (*GroupBadge, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	if args.ImageUrl == nil {
		return nil, errors.New("invalid value for required argument 'ImageUrl'")
	}
	if args.LinkUrl == nil {
		return nil, errors.New("invalid value for required argument 'LinkUrl'")
	}
	var resource GroupBadge
	err := ctx.RegisterResource("gitlab:index/groupBadge:GroupBadge", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupBadge gets an existing GroupBadge resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupBadge(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupBadgeState, opts ...pulumi.ResourceOption) (*GroupBadge, error) {
	var resource GroupBadge
	err := ctx.ReadResource("gitlab:index/groupBadge:GroupBadge", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupBadge resources.
type groupBadgeState struct {
	// The id of the group to add the badge to.
	Group *string `pulumi:"group"`
	// The image url which will be presented on group overview.
	ImageUrl *string `pulumi:"imageUrl"`
	// The url linked with the badge.
	LinkUrl *string `pulumi:"linkUrl"`
	// The image_url argument rendered (in case of use of placeholders).
	RenderedImageUrl *string `pulumi:"renderedImageUrl"`
	// The link_url argument rendered (in case of use of placeholders).
	RenderedLinkUrl *string `pulumi:"renderedLinkUrl"`
}

type GroupBadgeState struct {
	// The id of the group to add the badge to.
	Group pulumi.StringPtrInput
	// The image url which will be presented on group overview.
	ImageUrl pulumi.StringPtrInput
	// The url linked with the badge.
	LinkUrl pulumi.StringPtrInput
	// The image_url argument rendered (in case of use of placeholders).
	RenderedImageUrl pulumi.StringPtrInput
	// The link_url argument rendered (in case of use of placeholders).
	RenderedLinkUrl pulumi.StringPtrInput
}

func (GroupBadgeState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupBadgeState)(nil)).Elem()
}

type groupBadgeArgs struct {
	// The id of the group to add the badge to.
	Group string `pulumi:"group"`
	// The image url which will be presented on group overview.
	ImageUrl string `pulumi:"imageUrl"`
	// The url linked with the badge.
	LinkUrl string `pulumi:"linkUrl"`
}

// The set of arguments for constructing a GroupBadge resource.
type GroupBadgeArgs struct {
	// The id of the group to add the badge to.
	Group pulumi.StringInput
	// The image url which will be presented on group overview.
	ImageUrl pulumi.StringInput
	// The url linked with the badge.
	LinkUrl pulumi.StringInput
}

func (GroupBadgeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupBadgeArgs)(nil)).Elem()
}

type GroupBadgeInput interface {
	pulumi.Input

	ToGroupBadgeOutput() GroupBadgeOutput
	ToGroupBadgeOutputWithContext(ctx context.Context) GroupBadgeOutput
}

func (*GroupBadge) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupBadge)(nil)).Elem()
}

func (i *GroupBadge) ToGroupBadgeOutput() GroupBadgeOutput {
	return i.ToGroupBadgeOutputWithContext(context.Background())
}

func (i *GroupBadge) ToGroupBadgeOutputWithContext(ctx context.Context) GroupBadgeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupBadgeOutput)
}

// GroupBadgeArrayInput is an input type that accepts GroupBadgeArray and GroupBadgeArrayOutput values.
// You can construct a concrete instance of `GroupBadgeArrayInput` via:
//
//          GroupBadgeArray{ GroupBadgeArgs{...} }
type GroupBadgeArrayInput interface {
	pulumi.Input

	ToGroupBadgeArrayOutput() GroupBadgeArrayOutput
	ToGroupBadgeArrayOutputWithContext(context.Context) GroupBadgeArrayOutput
}

type GroupBadgeArray []GroupBadgeInput

func (GroupBadgeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupBadge)(nil)).Elem()
}

func (i GroupBadgeArray) ToGroupBadgeArrayOutput() GroupBadgeArrayOutput {
	return i.ToGroupBadgeArrayOutputWithContext(context.Background())
}

func (i GroupBadgeArray) ToGroupBadgeArrayOutputWithContext(ctx context.Context) GroupBadgeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupBadgeArrayOutput)
}

// GroupBadgeMapInput is an input type that accepts GroupBadgeMap and GroupBadgeMapOutput values.
// You can construct a concrete instance of `GroupBadgeMapInput` via:
//
//          GroupBadgeMap{ "key": GroupBadgeArgs{...} }
type GroupBadgeMapInput interface {
	pulumi.Input

	ToGroupBadgeMapOutput() GroupBadgeMapOutput
	ToGroupBadgeMapOutputWithContext(context.Context) GroupBadgeMapOutput
}

type GroupBadgeMap map[string]GroupBadgeInput

func (GroupBadgeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupBadge)(nil)).Elem()
}

func (i GroupBadgeMap) ToGroupBadgeMapOutput() GroupBadgeMapOutput {
	return i.ToGroupBadgeMapOutputWithContext(context.Background())
}

func (i GroupBadgeMap) ToGroupBadgeMapOutputWithContext(ctx context.Context) GroupBadgeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupBadgeMapOutput)
}

type GroupBadgeOutput struct{ *pulumi.OutputState }

func (GroupBadgeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupBadge)(nil)).Elem()
}

func (o GroupBadgeOutput) ToGroupBadgeOutput() GroupBadgeOutput {
	return o
}

func (o GroupBadgeOutput) ToGroupBadgeOutputWithContext(ctx context.Context) GroupBadgeOutput {
	return o
}

// The id of the group to add the badge to.
func (o GroupBadgeOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupBadge) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

// The image url which will be presented on group overview.
func (o GroupBadgeOutput) ImageUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupBadge) pulumi.StringOutput { return v.ImageUrl }).(pulumi.StringOutput)
}

// The url linked with the badge.
func (o GroupBadgeOutput) LinkUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupBadge) pulumi.StringOutput { return v.LinkUrl }).(pulumi.StringOutput)
}

// The image_url argument rendered (in case of use of placeholders).
func (o GroupBadgeOutput) RenderedImageUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupBadge) pulumi.StringOutput { return v.RenderedImageUrl }).(pulumi.StringOutput)
}

// The link_url argument rendered (in case of use of placeholders).
func (o GroupBadgeOutput) RenderedLinkUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupBadge) pulumi.StringOutput { return v.RenderedLinkUrl }).(pulumi.StringOutput)
}

type GroupBadgeArrayOutput struct{ *pulumi.OutputState }

func (GroupBadgeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupBadge)(nil)).Elem()
}

func (o GroupBadgeArrayOutput) ToGroupBadgeArrayOutput() GroupBadgeArrayOutput {
	return o
}

func (o GroupBadgeArrayOutput) ToGroupBadgeArrayOutputWithContext(ctx context.Context) GroupBadgeArrayOutput {
	return o
}

func (o GroupBadgeArrayOutput) Index(i pulumi.IntInput) GroupBadgeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupBadge {
		return vs[0].([]*GroupBadge)[vs[1].(int)]
	}).(GroupBadgeOutput)
}

type GroupBadgeMapOutput struct{ *pulumi.OutputState }

func (GroupBadgeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupBadge)(nil)).Elem()
}

func (o GroupBadgeMapOutput) ToGroupBadgeMapOutput() GroupBadgeMapOutput {
	return o
}

func (o GroupBadgeMapOutput) ToGroupBadgeMapOutputWithContext(ctx context.Context) GroupBadgeMapOutput {
	return o
}

func (o GroupBadgeMapOutput) MapIndex(k pulumi.StringInput) GroupBadgeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupBadge {
		return vs[0].(map[string]*GroupBadge)[vs[1].(string)]
	}).(GroupBadgeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupBadgeInput)(nil)).Elem(), &GroupBadge{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupBadgeArrayInput)(nil)).Elem(), GroupBadgeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupBadgeMapInput)(nil)).Elem(), GroupBadgeMap{})
	pulumi.RegisterOutputType(GroupBadgeOutput{})
	pulumi.RegisterOutputType(GroupBadgeArrayOutput{})
	pulumi.RegisterOutputType(GroupBadgeMapOutput{})
}
