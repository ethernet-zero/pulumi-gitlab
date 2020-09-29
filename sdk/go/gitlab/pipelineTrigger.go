// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # gitlab\_pipeline\_trigger
//
// This resource allows you to create and manage pipeline triggers
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gitlab/sdk/v2/go/gitlab"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := gitlab.NewPipelineTrigger(ctx, "example", &gitlab.PipelineTriggerArgs{
// 			Description: pulumi.String("Used to trigger builds"),
// 			Project:     pulumi.String("12345"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type PipelineTrigger struct {
	pulumi.CustomResourceState

	// The description of the pipeline trigger.
	Description pulumi.StringOutput `pulumi:"description"`
	// The name or id of the project to add the trigger to.
	Project pulumi.StringOutput `pulumi:"project"`
	Token   pulumi.StringOutput `pulumi:"token"`
}

// NewPipelineTrigger registers a new resource with the given unique name, arguments, and options.
func NewPipelineTrigger(ctx *pulumi.Context,
	name string, args *PipelineTriggerArgs, opts ...pulumi.ResourceOption) (*PipelineTrigger, error) {
	if args == nil || args.Description == nil {
		return nil, errors.New("missing required argument 'Description'")
	}
	if args == nil || args.Project == nil {
		return nil, errors.New("missing required argument 'Project'")
	}
	if args == nil {
		args = &PipelineTriggerArgs{}
	}
	var resource PipelineTrigger
	err := ctx.RegisterResource("gitlab:index/pipelineTrigger:PipelineTrigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipelineTrigger gets an existing PipelineTrigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipelineTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineTriggerState, opts ...pulumi.ResourceOption) (*PipelineTrigger, error) {
	var resource PipelineTrigger
	err := ctx.ReadResource("gitlab:index/pipelineTrigger:PipelineTrigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PipelineTrigger resources.
type pipelineTriggerState struct {
	// The description of the pipeline trigger.
	Description *string `pulumi:"description"`
	// The name or id of the project to add the trigger to.
	Project *string `pulumi:"project"`
	Token   *string `pulumi:"token"`
}

type PipelineTriggerState struct {
	// The description of the pipeline trigger.
	Description pulumi.StringPtrInput
	// The name or id of the project to add the trigger to.
	Project pulumi.StringPtrInput
	Token   pulumi.StringPtrInput
}

func (PipelineTriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineTriggerState)(nil)).Elem()
}

type pipelineTriggerArgs struct {
	// The description of the pipeline trigger.
	Description string `pulumi:"description"`
	// The name or id of the project to add the trigger to.
	Project string `pulumi:"project"`
}

// The set of arguments for constructing a PipelineTrigger resource.
type PipelineTriggerArgs struct {
	// The description of the pipeline trigger.
	Description pulumi.StringInput
	// The name or id of the project to add the trigger to.
	Project pulumi.StringInput
}

func (PipelineTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineTriggerArgs)(nil)).Elem()
}
