// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// A Repository File can be imported using an id made up of `<project-id>:<branch-name>:<file-path>`, e.g.
//
// ```sh
//
//	$ pulumi import gitlab:index/repositoryFile:RepositoryFile this 1:main:foo/bar.txt
//
// ```
type RepositoryFile struct {
	pulumi.CustomResourceState

	// Email of the commit author.
	AuthorEmail pulumi.StringPtrOutput `pulumi:"authorEmail"`
	// Name of the commit author.
	AuthorName pulumi.StringPtrOutput `pulumi:"authorName"`
	// The blob id.
	BlobId pulumi.StringOutput `pulumi:"blobId"`
	// Name of the branch to which to commit to.
	Branch pulumi.StringOutput `pulumi:"branch"`
	// The commit id.
	CommitId pulumi.StringOutput `pulumi:"commitId"`
	// Commit message.
	CommitMessage pulumi.StringOutput `pulumi:"commitMessage"`
	// File content. If the content is not yet base64 encoded, it will be encoded automatically. No other encoding is currently supported, because of a [GitLab API bug](https://gitlab.com/gitlab-org/gitlab/-/issues/342430).
	Content pulumi.StringOutput `pulumi:"content"`
	// File content sha256 digest.
	ContentSha256 pulumi.StringOutput `pulumi:"contentSha256"`
	// The file content encoding.
	Encoding pulumi.StringOutput `pulumi:"encoding"`
	// Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
	ExecuteFilemode pulumi.BoolPtrOutput `pulumi:"executeFilemode"`
	// The filename.
	FileName pulumi.StringOutput `pulumi:"fileName"`
	// The full path of the file. It must be relative to the root of the project without a leading slash `/`.
	FilePath pulumi.StringOutput `pulumi:"filePath"`
	// The last known commit id.
	LastCommitId pulumi.StringOutput `pulumi:"lastCommitId"`
	// The name or ID of the project.
	Project pulumi.StringOutput `pulumi:"project"`
	// The name of branch, tag or commit.
	Ref pulumi.StringOutput `pulumi:"ref"`
	// The file size.
	Size pulumi.IntOutput `pulumi:"size"`
	// Name of the branch to start the new commit from.
	StartBranch pulumi.StringPtrOutput `pulumi:"startBranch"`
}

// NewRepositoryFile registers a new resource with the given unique name, arguments, and options.
func NewRepositoryFile(ctx *pulumi.Context,
	name string, args *RepositoryFileArgs, opts ...pulumi.ResourceOption) (*RepositoryFile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Branch == nil {
		return nil, errors.New("invalid value for required argument 'Branch'")
	}
	if args.CommitMessage == nil {
		return nil, errors.New("invalid value for required argument 'CommitMessage'")
	}
	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.FilePath == nil {
		return nil, errors.New("invalid value for required argument 'FilePath'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource RepositoryFile
	err := ctx.RegisterResource("gitlab:index/repositoryFile:RepositoryFile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepositoryFile gets an existing RepositoryFile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepositoryFile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryFileState, opts ...pulumi.ResourceOption) (*RepositoryFile, error) {
	var resource RepositoryFile
	err := ctx.ReadResource("gitlab:index/repositoryFile:RepositoryFile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RepositoryFile resources.
type repositoryFileState struct {
	// Email of the commit author.
	AuthorEmail *string `pulumi:"authorEmail"`
	// Name of the commit author.
	AuthorName *string `pulumi:"authorName"`
	// The blob id.
	BlobId *string `pulumi:"blobId"`
	// Name of the branch to which to commit to.
	Branch *string `pulumi:"branch"`
	// The commit id.
	CommitId *string `pulumi:"commitId"`
	// Commit message.
	CommitMessage *string `pulumi:"commitMessage"`
	// File content. If the content is not yet base64 encoded, it will be encoded automatically. No other encoding is currently supported, because of a [GitLab API bug](https://gitlab.com/gitlab-org/gitlab/-/issues/342430).
	Content *string `pulumi:"content"`
	// File content sha256 digest.
	ContentSha256 *string `pulumi:"contentSha256"`
	// The file content encoding.
	Encoding *string `pulumi:"encoding"`
	// Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
	ExecuteFilemode *bool `pulumi:"executeFilemode"`
	// The filename.
	FileName *string `pulumi:"fileName"`
	// The full path of the file. It must be relative to the root of the project without a leading slash `/`.
	FilePath *string `pulumi:"filePath"`
	// The last known commit id.
	LastCommitId *string `pulumi:"lastCommitId"`
	// The name or ID of the project.
	Project *string `pulumi:"project"`
	// The name of branch, tag or commit.
	Ref *string `pulumi:"ref"`
	// The file size.
	Size *int `pulumi:"size"`
	// Name of the branch to start the new commit from.
	StartBranch *string `pulumi:"startBranch"`
}

type RepositoryFileState struct {
	// Email of the commit author.
	AuthorEmail pulumi.StringPtrInput
	// Name of the commit author.
	AuthorName pulumi.StringPtrInput
	// The blob id.
	BlobId pulumi.StringPtrInput
	// Name of the branch to which to commit to.
	Branch pulumi.StringPtrInput
	// The commit id.
	CommitId pulumi.StringPtrInput
	// Commit message.
	CommitMessage pulumi.StringPtrInput
	// File content. If the content is not yet base64 encoded, it will be encoded automatically. No other encoding is currently supported, because of a [GitLab API bug](https://gitlab.com/gitlab-org/gitlab/-/issues/342430).
	Content pulumi.StringPtrInput
	// File content sha256 digest.
	ContentSha256 pulumi.StringPtrInput
	// The file content encoding.
	Encoding pulumi.StringPtrInput
	// Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
	ExecuteFilemode pulumi.BoolPtrInput
	// The filename.
	FileName pulumi.StringPtrInput
	// The full path of the file. It must be relative to the root of the project without a leading slash `/`.
	FilePath pulumi.StringPtrInput
	// The last known commit id.
	LastCommitId pulumi.StringPtrInput
	// The name or ID of the project.
	Project pulumi.StringPtrInput
	// The name of branch, tag or commit.
	Ref pulumi.StringPtrInput
	// The file size.
	Size pulumi.IntPtrInput
	// Name of the branch to start the new commit from.
	StartBranch pulumi.StringPtrInput
}

func (RepositoryFileState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryFileState)(nil)).Elem()
}

type repositoryFileArgs struct {
	// Email of the commit author.
	AuthorEmail *string `pulumi:"authorEmail"`
	// Name of the commit author.
	AuthorName *string `pulumi:"authorName"`
	// Name of the branch to which to commit to.
	Branch string `pulumi:"branch"`
	// Commit message.
	CommitMessage string `pulumi:"commitMessage"`
	// File content. If the content is not yet base64 encoded, it will be encoded automatically. No other encoding is currently supported, because of a [GitLab API bug](https://gitlab.com/gitlab-org/gitlab/-/issues/342430).
	Content string `pulumi:"content"`
	// Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
	ExecuteFilemode *bool `pulumi:"executeFilemode"`
	// The full path of the file. It must be relative to the root of the project without a leading slash `/`.
	FilePath string `pulumi:"filePath"`
	// The name or ID of the project.
	Project string `pulumi:"project"`
	// Name of the branch to start the new commit from.
	StartBranch *string `pulumi:"startBranch"`
}

// The set of arguments for constructing a RepositoryFile resource.
type RepositoryFileArgs struct {
	// Email of the commit author.
	AuthorEmail pulumi.StringPtrInput
	// Name of the commit author.
	AuthorName pulumi.StringPtrInput
	// Name of the branch to which to commit to.
	Branch pulumi.StringInput
	// Commit message.
	CommitMessage pulumi.StringInput
	// File content. If the content is not yet base64 encoded, it will be encoded automatically. No other encoding is currently supported, because of a [GitLab API bug](https://gitlab.com/gitlab-org/gitlab/-/issues/342430).
	Content pulumi.StringInput
	// Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
	ExecuteFilemode pulumi.BoolPtrInput
	// The full path of the file. It must be relative to the root of the project without a leading slash `/`.
	FilePath pulumi.StringInput
	// The name or ID of the project.
	Project pulumi.StringInput
	// Name of the branch to start the new commit from.
	StartBranch pulumi.StringPtrInput
}

func (RepositoryFileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryFileArgs)(nil)).Elem()
}

type RepositoryFileInput interface {
	pulumi.Input

	ToRepositoryFileOutput() RepositoryFileOutput
	ToRepositoryFileOutputWithContext(ctx context.Context) RepositoryFileOutput
}

func (*RepositoryFile) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryFile)(nil)).Elem()
}

func (i *RepositoryFile) ToRepositoryFileOutput() RepositoryFileOutput {
	return i.ToRepositoryFileOutputWithContext(context.Background())
}

func (i *RepositoryFile) ToRepositoryFileOutputWithContext(ctx context.Context) RepositoryFileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryFileOutput)
}

// RepositoryFileArrayInput is an input type that accepts RepositoryFileArray and RepositoryFileArrayOutput values.
// You can construct a concrete instance of `RepositoryFileArrayInput` via:
//
//	RepositoryFileArray{ RepositoryFileArgs{...} }
type RepositoryFileArrayInput interface {
	pulumi.Input

	ToRepositoryFileArrayOutput() RepositoryFileArrayOutput
	ToRepositoryFileArrayOutputWithContext(context.Context) RepositoryFileArrayOutput
}

type RepositoryFileArray []RepositoryFileInput

func (RepositoryFileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryFile)(nil)).Elem()
}

func (i RepositoryFileArray) ToRepositoryFileArrayOutput() RepositoryFileArrayOutput {
	return i.ToRepositoryFileArrayOutputWithContext(context.Background())
}

func (i RepositoryFileArray) ToRepositoryFileArrayOutputWithContext(ctx context.Context) RepositoryFileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryFileArrayOutput)
}

// RepositoryFileMapInput is an input type that accepts RepositoryFileMap and RepositoryFileMapOutput values.
// You can construct a concrete instance of `RepositoryFileMapInput` via:
//
//	RepositoryFileMap{ "key": RepositoryFileArgs{...} }
type RepositoryFileMapInput interface {
	pulumi.Input

	ToRepositoryFileMapOutput() RepositoryFileMapOutput
	ToRepositoryFileMapOutputWithContext(context.Context) RepositoryFileMapOutput
}

type RepositoryFileMap map[string]RepositoryFileInput

func (RepositoryFileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryFile)(nil)).Elem()
}

func (i RepositoryFileMap) ToRepositoryFileMapOutput() RepositoryFileMapOutput {
	return i.ToRepositoryFileMapOutputWithContext(context.Background())
}

func (i RepositoryFileMap) ToRepositoryFileMapOutputWithContext(ctx context.Context) RepositoryFileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryFileMapOutput)
}

type RepositoryFileOutput struct{ *pulumi.OutputState }

func (RepositoryFileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryFile)(nil)).Elem()
}

func (o RepositoryFileOutput) ToRepositoryFileOutput() RepositoryFileOutput {
	return o
}

func (o RepositoryFileOutput) ToRepositoryFileOutputWithContext(ctx context.Context) RepositoryFileOutput {
	return o
}

// Email of the commit author.
func (o RepositoryFileOutput) AuthorEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringPtrOutput { return v.AuthorEmail }).(pulumi.StringPtrOutput)
}

// Name of the commit author.
func (o RepositoryFileOutput) AuthorName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringPtrOutput { return v.AuthorName }).(pulumi.StringPtrOutput)
}

// The blob id.
func (o RepositoryFileOutput) BlobId() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringOutput { return v.BlobId }).(pulumi.StringOutput)
}

// Name of the branch to which to commit to.
func (o RepositoryFileOutput) Branch() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringOutput { return v.Branch }).(pulumi.StringOutput)
}

// The commit id.
func (o RepositoryFileOutput) CommitId() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringOutput { return v.CommitId }).(pulumi.StringOutput)
}

// Commit message.
func (o RepositoryFileOutput) CommitMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringOutput { return v.CommitMessage }).(pulumi.StringOutput)
}

// File content. If the content is not yet base64 encoded, it will be encoded automatically. No other encoding is currently supported, because of a [GitLab API bug](https://gitlab.com/gitlab-org/gitlab/-/issues/342430).
func (o RepositoryFileOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

// File content sha256 digest.
func (o RepositoryFileOutput) ContentSha256() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringOutput { return v.ContentSha256 }).(pulumi.StringOutput)
}

// The file content encoding.
func (o RepositoryFileOutput) Encoding() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringOutput { return v.Encoding }).(pulumi.StringOutput)
}

// Enables or disables the execute flag on the file. **Note**: requires GitLab 14.10 or newer.
func (o RepositoryFileOutput) ExecuteFilemode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.BoolPtrOutput { return v.ExecuteFilemode }).(pulumi.BoolPtrOutput)
}

// The filename.
func (o RepositoryFileOutput) FileName() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringOutput { return v.FileName }).(pulumi.StringOutput)
}

// The full path of the file. It must be relative to the root of the project without a leading slash `/`.
func (o RepositoryFileOutput) FilePath() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringOutput { return v.FilePath }).(pulumi.StringOutput)
}

// The last known commit id.
func (o RepositoryFileOutput) LastCommitId() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringOutput { return v.LastCommitId }).(pulumi.StringOutput)
}

// The name or ID of the project.
func (o RepositoryFileOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The name of branch, tag or commit.
func (o RepositoryFileOutput) Ref() pulumi.StringOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringOutput { return v.Ref }).(pulumi.StringOutput)
}

// The file size.
func (o RepositoryFileOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// Name of the branch to start the new commit from.
func (o RepositoryFileOutput) StartBranch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryFile) pulumi.StringPtrOutput { return v.StartBranch }).(pulumi.StringPtrOutput)
}

type RepositoryFileArrayOutput struct{ *pulumi.OutputState }

func (RepositoryFileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RepositoryFile)(nil)).Elem()
}

func (o RepositoryFileArrayOutput) ToRepositoryFileArrayOutput() RepositoryFileArrayOutput {
	return o
}

func (o RepositoryFileArrayOutput) ToRepositoryFileArrayOutputWithContext(ctx context.Context) RepositoryFileArrayOutput {
	return o
}

func (o RepositoryFileArrayOutput) Index(i pulumi.IntInput) RepositoryFileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RepositoryFile {
		return vs[0].([]*RepositoryFile)[vs[1].(int)]
	}).(RepositoryFileOutput)
}

type RepositoryFileMapOutput struct{ *pulumi.OutputState }

func (RepositoryFileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RepositoryFile)(nil)).Elem()
}

func (o RepositoryFileMapOutput) ToRepositoryFileMapOutput() RepositoryFileMapOutput {
	return o
}

func (o RepositoryFileMapOutput) ToRepositoryFileMapOutputWithContext(ctx context.Context) RepositoryFileMapOutput {
	return o
}

func (o RepositoryFileMapOutput) MapIndex(k pulumi.StringInput) RepositoryFileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RepositoryFile {
		return vs[0].(map[string]*RepositoryFile)[vs[1].(string)]
	}).(RepositoryFileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryFileInput)(nil)).Elem(), &RepositoryFile{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryFileArrayInput)(nil)).Elem(), RepositoryFileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryFileMapInput)(nil)).Elem(), RepositoryFileMap{})
	pulumi.RegisterOutputType(RepositoryFileOutput{})
	pulumi.RegisterOutputType(RepositoryFileArrayOutput{})
	pulumi.RegisterOutputType(RepositoryFileMapOutput{})
}
