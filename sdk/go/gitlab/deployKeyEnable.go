// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # gitlab\_deploy\_key\_enable
//
// This resource allows you to enable pre-existing deploy keys for your GitLab projects.
//
// **the GITLAB KEY_ID for the deploy key must be known**
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gitlab/sdk/v3/go/gitlab"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		parentProject, err := gitlab.NewProject(ctx, "parentProject", nil)
// 		if err != nil {
// 			return err
// 		}
// 		fooProject, err := gitlab.NewProject(ctx, "fooProject", nil)
// 		if err != nil {
// 			return err
// 		}
// 		parentDeployKey, err := gitlab.NewDeployKey(ctx, "parentDeployKey", &gitlab.DeployKeyArgs{
// 			Key:     pulumi.String("ssh-rsa AAAA..."),
// 			Project: parentProject.ID(),
// 			Title:   pulumi.String("Example deploy key"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = gitlab.NewDeployKeyEnable(ctx, "fooDeployKeyEnable", &gitlab.DeployKeyEnableArgs{
// 			KeyId:   parentDeployKey.ID(),
// 			Project: fooProject.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DeployKeyEnable struct {
	pulumi.CustomResourceState

	CanPush pulumi.BoolOutput   `pulumi:"canPush"`
	Key     pulumi.StringOutput `pulumi:"key"`
	// The Gitlab key id for the pre-existing deploy key
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// The name or id of the project to add the deploy key to.
	Project pulumi.StringOutput `pulumi:"project"`
	Title   pulumi.StringOutput `pulumi:"title"`
}

// NewDeployKeyEnable registers a new resource with the given unique name, arguments, and options.
func NewDeployKeyEnable(ctx *pulumi.Context,
	name string, args *DeployKeyEnableArgs, opts ...pulumi.ResourceOption) (*DeployKeyEnable, error) {
	if args == nil || args.KeyId == nil {
		return nil, errors.New("missing required argument 'KeyId'")
	}
	if args == nil || args.Project == nil {
		return nil, errors.New("missing required argument 'Project'")
	}
	if args == nil {
		args = &DeployKeyEnableArgs{}
	}
	var resource DeployKeyEnable
	err := ctx.RegisterResource("gitlab:index/deployKeyEnable:DeployKeyEnable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeployKeyEnable gets an existing DeployKeyEnable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeployKeyEnable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeployKeyEnableState, opts ...pulumi.ResourceOption) (*DeployKeyEnable, error) {
	var resource DeployKeyEnable
	err := ctx.ReadResource("gitlab:index/deployKeyEnable:DeployKeyEnable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeployKeyEnable resources.
type deployKeyEnableState struct {
	CanPush *bool   `pulumi:"canPush"`
	Key     *string `pulumi:"key"`
	// The Gitlab key id for the pre-existing deploy key
	KeyId *string `pulumi:"keyId"`
	// The name or id of the project to add the deploy key to.
	Project *string `pulumi:"project"`
	Title   *string `pulumi:"title"`
}

type DeployKeyEnableState struct {
	CanPush pulumi.BoolPtrInput
	Key     pulumi.StringPtrInput
	// The Gitlab key id for the pre-existing deploy key
	KeyId pulumi.StringPtrInput
	// The name or id of the project to add the deploy key to.
	Project pulumi.StringPtrInput
	Title   pulumi.StringPtrInput
}

func (DeployKeyEnableState) ElementType() reflect.Type {
	return reflect.TypeOf((*deployKeyEnableState)(nil)).Elem()
}

type deployKeyEnableArgs struct {
	CanPush *bool   `pulumi:"canPush"`
	Key     *string `pulumi:"key"`
	// The Gitlab key id for the pre-existing deploy key
	KeyId string `pulumi:"keyId"`
	// The name or id of the project to add the deploy key to.
	Project string  `pulumi:"project"`
	Title   *string `pulumi:"title"`
}

// The set of arguments for constructing a DeployKeyEnable resource.
type DeployKeyEnableArgs struct {
	CanPush pulumi.BoolPtrInput
	Key     pulumi.StringPtrInput
	// The Gitlab key id for the pre-existing deploy key
	KeyId pulumi.StringInput
	// The name or id of the project to add the deploy key to.
	Project pulumi.StringInput
	Title   pulumi.StringPtrInput
}

func (DeployKeyEnableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deployKeyEnableArgs)(nil)).Elem()
}
