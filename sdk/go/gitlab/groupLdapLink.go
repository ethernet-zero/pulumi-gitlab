// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `GroupLdapLink` resource allows to manage the lifecycle of an LDAP integration with a group.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#ldap-group-links)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gitlab/sdk/v4/go/gitlab"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gitlab.NewGroupLdapLink(ctx, "test", &gitlab.GroupLdapLinkArgs{
//				Cn:           pulumi.String("testuser"),
//				GroupAccess:  pulumi.String("developer"),
//				GroupId:      pulumi.String("12345"),
//				LdapProvider: pulumi.String("ldapmain"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// GitLab group ldap links can be imported using an id made up of `group_id:ldap_provider:cn`, e.g.
//
// ```sh
//
//	$ pulumi import gitlab:index/groupLdapLink:GroupLdapLink test "12345:ldapmain:testuser"
//
// ```
type GroupLdapLink struct {
	pulumi.CustomResourceState

	// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	//
	// Deprecated: Use `group_access` instead of the `access_level` attribute.
	AccessLevel pulumi.StringPtrOutput `pulumi:"accessLevel"`
	// The CN of the LDAP group to link with.
	Cn pulumi.StringOutput `pulumi:"cn"`
	// If true, then delete and replace an existing LDAP link if one exists.
	Force pulumi.BoolPtrOutput `pulumi:"force"`
	// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	GroupAccess pulumi.StringPtrOutput `pulumi:"groupAccess"`
	// The id of the GitLab group.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
	LdapProvider pulumi.StringOutput `pulumi:"ldapProvider"`
}

// NewGroupLdapLink registers a new resource with the given unique name, arguments, and options.
func NewGroupLdapLink(ctx *pulumi.Context,
	name string, args *GroupLdapLinkArgs, opts ...pulumi.ResourceOption) (*GroupLdapLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cn == nil {
		return nil, errors.New("invalid value for required argument 'Cn'")
	}
	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.LdapProvider == nil {
		return nil, errors.New("invalid value for required argument 'LdapProvider'")
	}
	var resource GroupLdapLink
	err := ctx.RegisterResource("gitlab:index/groupLdapLink:GroupLdapLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupLdapLink gets an existing GroupLdapLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupLdapLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupLdapLinkState, opts ...pulumi.ResourceOption) (*GroupLdapLink, error) {
	var resource GroupLdapLink
	err := ctx.ReadResource("gitlab:index/groupLdapLink:GroupLdapLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupLdapLink resources.
type groupLdapLinkState struct {
	// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	//
	// Deprecated: Use `group_access` instead of the `access_level` attribute.
	AccessLevel *string `pulumi:"accessLevel"`
	// The CN of the LDAP group to link with.
	Cn *string `pulumi:"cn"`
	// If true, then delete and replace an existing LDAP link if one exists.
	Force *bool `pulumi:"force"`
	// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	GroupAccess *string `pulumi:"groupAccess"`
	// The id of the GitLab group.
	GroupId *string `pulumi:"groupId"`
	// The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
	LdapProvider *string `pulumi:"ldapProvider"`
}

type GroupLdapLinkState struct {
	// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	//
	// Deprecated: Use `group_access` instead of the `access_level` attribute.
	AccessLevel pulumi.StringPtrInput
	// The CN of the LDAP group to link with.
	Cn pulumi.StringPtrInput
	// If true, then delete and replace an existing LDAP link if one exists.
	Force pulumi.BoolPtrInput
	// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	GroupAccess pulumi.StringPtrInput
	// The id of the GitLab group.
	GroupId pulumi.StringPtrInput
	// The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
	LdapProvider pulumi.StringPtrInput
}

func (GroupLdapLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupLdapLinkState)(nil)).Elem()
}

type groupLdapLinkArgs struct {
	// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	//
	// Deprecated: Use `group_access` instead of the `access_level` attribute.
	AccessLevel *string `pulumi:"accessLevel"`
	// The CN of the LDAP group to link with.
	Cn string `pulumi:"cn"`
	// If true, then delete and replace an existing LDAP link if one exists.
	Force *bool `pulumi:"force"`
	// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	GroupAccess *string `pulumi:"groupAccess"`
	// The id of the GitLab group.
	GroupId string `pulumi:"groupId"`
	// The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
	LdapProvider string `pulumi:"ldapProvider"`
}

// The set of arguments for constructing a GroupLdapLink resource.
type GroupLdapLinkArgs struct {
	// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	//
	// Deprecated: Use `group_access` instead of the `access_level` attribute.
	AccessLevel pulumi.StringPtrInput
	// The CN of the LDAP group to link with.
	Cn pulumi.StringInput
	// If true, then delete and replace an existing LDAP link if one exists.
	Force pulumi.BoolPtrInput
	// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
	GroupAccess pulumi.StringPtrInput
	// The id of the GitLab group.
	GroupId pulumi.StringInput
	// The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
	LdapProvider pulumi.StringInput
}

func (GroupLdapLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupLdapLinkArgs)(nil)).Elem()
}

type GroupLdapLinkInput interface {
	pulumi.Input

	ToGroupLdapLinkOutput() GroupLdapLinkOutput
	ToGroupLdapLinkOutputWithContext(ctx context.Context) GroupLdapLinkOutput
}

func (*GroupLdapLink) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupLdapLink)(nil)).Elem()
}

func (i *GroupLdapLink) ToGroupLdapLinkOutput() GroupLdapLinkOutput {
	return i.ToGroupLdapLinkOutputWithContext(context.Background())
}

func (i *GroupLdapLink) ToGroupLdapLinkOutputWithContext(ctx context.Context) GroupLdapLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupLdapLinkOutput)
}

// GroupLdapLinkArrayInput is an input type that accepts GroupLdapLinkArray and GroupLdapLinkArrayOutput values.
// You can construct a concrete instance of `GroupLdapLinkArrayInput` via:
//
//	GroupLdapLinkArray{ GroupLdapLinkArgs{...} }
type GroupLdapLinkArrayInput interface {
	pulumi.Input

	ToGroupLdapLinkArrayOutput() GroupLdapLinkArrayOutput
	ToGroupLdapLinkArrayOutputWithContext(context.Context) GroupLdapLinkArrayOutput
}

type GroupLdapLinkArray []GroupLdapLinkInput

func (GroupLdapLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupLdapLink)(nil)).Elem()
}

func (i GroupLdapLinkArray) ToGroupLdapLinkArrayOutput() GroupLdapLinkArrayOutput {
	return i.ToGroupLdapLinkArrayOutputWithContext(context.Background())
}

func (i GroupLdapLinkArray) ToGroupLdapLinkArrayOutputWithContext(ctx context.Context) GroupLdapLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupLdapLinkArrayOutput)
}

// GroupLdapLinkMapInput is an input type that accepts GroupLdapLinkMap and GroupLdapLinkMapOutput values.
// You can construct a concrete instance of `GroupLdapLinkMapInput` via:
//
//	GroupLdapLinkMap{ "key": GroupLdapLinkArgs{...} }
type GroupLdapLinkMapInput interface {
	pulumi.Input

	ToGroupLdapLinkMapOutput() GroupLdapLinkMapOutput
	ToGroupLdapLinkMapOutputWithContext(context.Context) GroupLdapLinkMapOutput
}

type GroupLdapLinkMap map[string]GroupLdapLinkInput

func (GroupLdapLinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupLdapLink)(nil)).Elem()
}

func (i GroupLdapLinkMap) ToGroupLdapLinkMapOutput() GroupLdapLinkMapOutput {
	return i.ToGroupLdapLinkMapOutputWithContext(context.Background())
}

func (i GroupLdapLinkMap) ToGroupLdapLinkMapOutputWithContext(ctx context.Context) GroupLdapLinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupLdapLinkMapOutput)
}

type GroupLdapLinkOutput struct{ *pulumi.OutputState }

func (GroupLdapLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupLdapLink)(nil)).Elem()
}

func (o GroupLdapLinkOutput) ToGroupLdapLinkOutput() GroupLdapLinkOutput {
	return o
}

func (o GroupLdapLinkOutput) ToGroupLdapLinkOutputWithContext(ctx context.Context) GroupLdapLinkOutput {
	return o
}

// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
//
// Deprecated: Use `group_access` instead of the `access_level` attribute.
func (o GroupLdapLinkOutput) AccessLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupLdapLink) pulumi.StringPtrOutput { return v.AccessLevel }).(pulumi.StringPtrOutput)
}

// The CN of the LDAP group to link with.
func (o GroupLdapLinkOutput) Cn() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupLdapLink) pulumi.StringOutput { return v.Cn }).(pulumi.StringOutput)
}

// If true, then delete and replace an existing LDAP link if one exists.
func (o GroupLdapLinkOutput) Force() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupLdapLink) pulumi.BoolPtrOutput { return v.Force }).(pulumi.BoolPtrOutput)
}

// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
func (o GroupLdapLinkOutput) GroupAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupLdapLink) pulumi.StringPtrOutput { return v.GroupAccess }).(pulumi.StringPtrOutput)
}

// The id of the GitLab group.
func (o GroupLdapLinkOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupLdapLink) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// The name of the LDAP provider as stored in the GitLab database. Note that this is NOT the value of the `label` attribute as shown in the web UI. In most cases this will be `ldapmain` but you may use the [LDAP check rake task](https://docs.gitlab.com/ee/administration/raketasks/ldap.html#check) for receiving the LDAP server name: `LDAP: ... Server: ldapmain`
func (o GroupLdapLinkOutput) LdapProvider() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupLdapLink) pulumi.StringOutput { return v.LdapProvider }).(pulumi.StringOutput)
}

type GroupLdapLinkArrayOutput struct{ *pulumi.OutputState }

func (GroupLdapLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GroupLdapLink)(nil)).Elem()
}

func (o GroupLdapLinkArrayOutput) ToGroupLdapLinkArrayOutput() GroupLdapLinkArrayOutput {
	return o
}

func (o GroupLdapLinkArrayOutput) ToGroupLdapLinkArrayOutputWithContext(ctx context.Context) GroupLdapLinkArrayOutput {
	return o
}

func (o GroupLdapLinkArrayOutput) Index(i pulumi.IntInput) GroupLdapLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GroupLdapLink {
		return vs[0].([]*GroupLdapLink)[vs[1].(int)]
	}).(GroupLdapLinkOutput)
}

type GroupLdapLinkMapOutput struct{ *pulumi.OutputState }

func (GroupLdapLinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GroupLdapLink)(nil)).Elem()
}

func (o GroupLdapLinkMapOutput) ToGroupLdapLinkMapOutput() GroupLdapLinkMapOutput {
	return o
}

func (o GroupLdapLinkMapOutput) ToGroupLdapLinkMapOutputWithContext(ctx context.Context) GroupLdapLinkMapOutput {
	return o
}

func (o GroupLdapLinkMapOutput) MapIndex(k pulumi.StringInput) GroupLdapLinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GroupLdapLink {
		return vs[0].(map[string]*GroupLdapLink)[vs[1].(string)]
	}).(GroupLdapLinkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupLdapLinkInput)(nil)).Elem(), &GroupLdapLink{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupLdapLinkArrayInput)(nil)).Elem(), GroupLdapLinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupLdapLinkMapInput)(nil)).Elem(), GroupLdapLinkMap{})
	pulumi.RegisterOutputType(GroupLdapLinkOutput{})
	pulumi.RegisterOutputType(GroupLdapLinkArrayOutput{})
	pulumi.RegisterOutputType(GroupLdapLinkMapOutput{})
}
