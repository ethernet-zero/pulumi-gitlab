// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Group struct {
	pulumi.CustomResourceState

	// Boolean, defaults to false.  Default to Auto
	// DevOps pipeline for all projects within this group.
	AutoDevopsEnabled pulumi.BoolPtrOutput `pulumi:"autoDevopsEnabled"`
	// The description of the group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Boolean, defaults to false.  Disable email notifications
	EmailsDisabled pulumi.BoolPtrOutput `pulumi:"emailsDisabled"`
	// The full name of the group.
	FullName pulumi.StringOutput `pulumi:"fullName"`
	// The full path of the group.
	FullPath pulumi.StringOutput `pulumi:"fullPath"`
	// Boolean, defaults to true.  Whether to enable LFS
	// support for projects in this group.
	LfsEnabled pulumi.BoolPtrOutput `pulumi:"lfsEnabled"`
	// Boolean, defaults to false.  Disable the capability
	// of a group from getting mentioned
	MentionsDisabled pulumi.BoolPtrOutput `pulumi:"mentionsDisabled"`
	// The name of this group.
	Name pulumi.StringOutput `pulumi:"name"`
	// Integer, id of the parent group (creates a nested group).
	ParentId pulumi.IntPtrOutput `pulumi:"parentId"`
	// The path of the group.
	Path pulumi.StringOutput `pulumi:"path"`
	// , defaults to Maintainer.
	// Determine if developers can create projects
	// in the group. Can be noone (No one), maintainer (Maintainers),
	// or developer (Developers + Maintainers).
	ProjectCreationLevel pulumi.StringPtrOutput `pulumi:"projectCreationLevel"`
	// Boolean, defaults to false.  Whether to
	// enable users to request access to the group.
	RequestAccessEnabled pulumi.BoolPtrOutput `pulumi:"requestAccessEnabled"`
	// Boolean, defaults to false.
	// equire all users in this group to setup Two-factor authentication.
	RequireTwoFactorAuthentication pulumi.BoolPtrOutput `pulumi:"requireTwoFactorAuthentication"`
	// The group level registration token to use during runner setup.
	RunnersToken pulumi.StringOutput `pulumi:"runnersToken"`
	// Boolean, defaults to false.  Prevent sharing
	// a project with another group within this group.
	ShareWithGroupLock pulumi.BoolPtrOutput `pulumi:"shareWithGroupLock"`
	// , defaults to Owner.
	// Allowed to create subgroups.
	// Can be owner (Owners), or maintainer (Maintainers).
	SubgroupCreationLevel pulumi.StringPtrOutput `pulumi:"subgroupCreationLevel"`
	// Int, defaults to 48.
	// Time before Two-factor authentication is enforced (in hours).
	TwoFactorGracePeriod pulumi.IntPtrOutput `pulumi:"twoFactorGracePeriod"`
	// The group's visibility. Can be `private`, `internal`, or `public`.
	VisibilityLevel pulumi.StringOutput `pulumi:"visibilityLevel"`
	// Web URL of the group.
	WebUrl pulumi.StringOutput `pulumi:"webUrl"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	var resource Group
	err := ctx.RegisterResource("gitlab:index/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("gitlab:index/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// Boolean, defaults to false.  Default to Auto
	// DevOps pipeline for all projects within this group.
	AutoDevopsEnabled *bool `pulumi:"autoDevopsEnabled"`
	// The description of the group.
	Description *string `pulumi:"description"`
	// Boolean, defaults to false.  Disable email notifications
	EmailsDisabled *bool `pulumi:"emailsDisabled"`
	// The full name of the group.
	FullName *string `pulumi:"fullName"`
	// The full path of the group.
	FullPath *string `pulumi:"fullPath"`
	// Boolean, defaults to true.  Whether to enable LFS
	// support for projects in this group.
	LfsEnabled *bool `pulumi:"lfsEnabled"`
	// Boolean, defaults to false.  Disable the capability
	// of a group from getting mentioned
	MentionsDisabled *bool `pulumi:"mentionsDisabled"`
	// The name of this group.
	Name *string `pulumi:"name"`
	// Integer, id of the parent group (creates a nested group).
	ParentId *int `pulumi:"parentId"`
	// The path of the group.
	Path *string `pulumi:"path"`
	// , defaults to Maintainer.
	// Determine if developers can create projects
	// in the group. Can be noone (No one), maintainer (Maintainers),
	// or developer (Developers + Maintainers).
	ProjectCreationLevel *string `pulumi:"projectCreationLevel"`
	// Boolean, defaults to false.  Whether to
	// enable users to request access to the group.
	RequestAccessEnabled *bool `pulumi:"requestAccessEnabled"`
	// Boolean, defaults to false.
	// equire all users in this group to setup Two-factor authentication.
	RequireTwoFactorAuthentication *bool `pulumi:"requireTwoFactorAuthentication"`
	// The group level registration token to use during runner setup.
	RunnersToken *string `pulumi:"runnersToken"`
	// Boolean, defaults to false.  Prevent sharing
	// a project with another group within this group.
	ShareWithGroupLock *bool `pulumi:"shareWithGroupLock"`
	// , defaults to Owner.
	// Allowed to create subgroups.
	// Can be owner (Owners), or maintainer (Maintainers).
	SubgroupCreationLevel *string `pulumi:"subgroupCreationLevel"`
	// Int, defaults to 48.
	// Time before Two-factor authentication is enforced (in hours).
	TwoFactorGracePeriod *int `pulumi:"twoFactorGracePeriod"`
	// The group's visibility. Can be `private`, `internal`, or `public`.
	VisibilityLevel *string `pulumi:"visibilityLevel"`
	// Web URL of the group.
	WebUrl *string `pulumi:"webUrl"`
}

type GroupState struct {
	// Boolean, defaults to false.  Default to Auto
	// DevOps pipeline for all projects within this group.
	AutoDevopsEnabled pulumi.BoolPtrInput
	// The description of the group.
	Description pulumi.StringPtrInput
	// Boolean, defaults to false.  Disable email notifications
	EmailsDisabled pulumi.BoolPtrInput
	// The full name of the group.
	FullName pulumi.StringPtrInput
	// The full path of the group.
	FullPath pulumi.StringPtrInput
	// Boolean, defaults to true.  Whether to enable LFS
	// support for projects in this group.
	LfsEnabled pulumi.BoolPtrInput
	// Boolean, defaults to false.  Disable the capability
	// of a group from getting mentioned
	MentionsDisabled pulumi.BoolPtrInput
	// The name of this group.
	Name pulumi.StringPtrInput
	// Integer, id of the parent group (creates a nested group).
	ParentId pulumi.IntPtrInput
	// The path of the group.
	Path pulumi.StringPtrInput
	// , defaults to Maintainer.
	// Determine if developers can create projects
	// in the group. Can be noone (No one), maintainer (Maintainers),
	// or developer (Developers + Maintainers).
	ProjectCreationLevel pulumi.StringPtrInput
	// Boolean, defaults to false.  Whether to
	// enable users to request access to the group.
	RequestAccessEnabled pulumi.BoolPtrInput
	// Boolean, defaults to false.
	// equire all users in this group to setup Two-factor authentication.
	RequireTwoFactorAuthentication pulumi.BoolPtrInput
	// The group level registration token to use during runner setup.
	RunnersToken pulumi.StringPtrInput
	// Boolean, defaults to false.  Prevent sharing
	// a project with another group within this group.
	ShareWithGroupLock pulumi.BoolPtrInput
	// , defaults to Owner.
	// Allowed to create subgroups.
	// Can be owner (Owners), or maintainer (Maintainers).
	SubgroupCreationLevel pulumi.StringPtrInput
	// Int, defaults to 48.
	// Time before Two-factor authentication is enforced (in hours).
	TwoFactorGracePeriod pulumi.IntPtrInput
	// The group's visibility. Can be `private`, `internal`, or `public`.
	VisibilityLevel pulumi.StringPtrInput
	// Web URL of the group.
	WebUrl pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// Boolean, defaults to false.  Default to Auto
	// DevOps pipeline for all projects within this group.
	AutoDevopsEnabled *bool `pulumi:"autoDevopsEnabled"`
	// The description of the group.
	Description *string `pulumi:"description"`
	// Boolean, defaults to false.  Disable email notifications
	EmailsDisabled *bool `pulumi:"emailsDisabled"`
	// Boolean, defaults to true.  Whether to enable LFS
	// support for projects in this group.
	LfsEnabled *bool `pulumi:"lfsEnabled"`
	// Boolean, defaults to false.  Disable the capability
	// of a group from getting mentioned
	MentionsDisabled *bool `pulumi:"mentionsDisabled"`
	// The name of this group.
	Name *string `pulumi:"name"`
	// Integer, id of the parent group (creates a nested group).
	ParentId *int `pulumi:"parentId"`
	// The path of the group.
	Path string `pulumi:"path"`
	// , defaults to Maintainer.
	// Determine if developers can create projects
	// in the group. Can be noone (No one), maintainer (Maintainers),
	// or developer (Developers + Maintainers).
	ProjectCreationLevel *string `pulumi:"projectCreationLevel"`
	// Boolean, defaults to false.  Whether to
	// enable users to request access to the group.
	RequestAccessEnabled *bool `pulumi:"requestAccessEnabled"`
	// Boolean, defaults to false.
	// equire all users in this group to setup Two-factor authentication.
	RequireTwoFactorAuthentication *bool `pulumi:"requireTwoFactorAuthentication"`
	// Boolean, defaults to false.  Prevent sharing
	// a project with another group within this group.
	ShareWithGroupLock *bool `pulumi:"shareWithGroupLock"`
	// , defaults to Owner.
	// Allowed to create subgroups.
	// Can be owner (Owners), or maintainer (Maintainers).
	SubgroupCreationLevel *string `pulumi:"subgroupCreationLevel"`
	// Int, defaults to 48.
	// Time before Two-factor authentication is enforced (in hours).
	TwoFactorGracePeriod *int `pulumi:"twoFactorGracePeriod"`
	// The group's visibility. Can be `private`, `internal`, or `public`.
	VisibilityLevel *string `pulumi:"visibilityLevel"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// Boolean, defaults to false.  Default to Auto
	// DevOps pipeline for all projects within this group.
	AutoDevopsEnabled pulumi.BoolPtrInput
	// The description of the group.
	Description pulumi.StringPtrInput
	// Boolean, defaults to false.  Disable email notifications
	EmailsDisabled pulumi.BoolPtrInput
	// Boolean, defaults to true.  Whether to enable LFS
	// support for projects in this group.
	LfsEnabled pulumi.BoolPtrInput
	// Boolean, defaults to false.  Disable the capability
	// of a group from getting mentioned
	MentionsDisabled pulumi.BoolPtrInput
	// The name of this group.
	Name pulumi.StringPtrInput
	// Integer, id of the parent group (creates a nested group).
	ParentId pulumi.IntPtrInput
	// The path of the group.
	Path pulumi.StringInput
	// , defaults to Maintainer.
	// Determine if developers can create projects
	// in the group. Can be noone (No one), maintainer (Maintainers),
	// or developer (Developers + Maintainers).
	ProjectCreationLevel pulumi.StringPtrInput
	// Boolean, defaults to false.  Whether to
	// enable users to request access to the group.
	RequestAccessEnabled pulumi.BoolPtrInput
	// Boolean, defaults to false.
	// equire all users in this group to setup Two-factor authentication.
	RequireTwoFactorAuthentication pulumi.BoolPtrInput
	// Boolean, defaults to false.  Prevent sharing
	// a project with another group within this group.
	ShareWithGroupLock pulumi.BoolPtrInput
	// , defaults to Owner.
	// Allowed to create subgroups.
	// Can be owner (Owners), or maintainer (Maintainers).
	SubgroupCreationLevel pulumi.StringPtrInput
	// Int, defaults to 48.
	// Time before Two-factor authentication is enforced (in hours).
	TwoFactorGracePeriod pulumi.IntPtrInput
	// The group's visibility. Can be `private`, `internal`, or `public`.
	VisibilityLevel pulumi.StringPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((*Group)(nil))
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

type GroupOutput struct {
	*pulumi.OutputState
}

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Group)(nil))
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GroupOutput{})
}
