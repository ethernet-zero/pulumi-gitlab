// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `Project` data source allows details of a project to be retrieved by either its ID or its path with namespace.
//
// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#get-single-project)
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
//			_, err = gitlab.LookupProject(ctx, &GetProjectArgs{
//				Id: pulumi.StringRef("foo/bar/baz"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	var rv LookupProjectResult
	err := ctx.Invoke("gitlab:index/getProject:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProject.
type LookupProjectArgs struct {
	// Default number of revisions for shallow cloning.
	CiDefaultGitDepth *int `pulumi:"ciDefaultGitDepth"`
	// The integer or path with namespace that uniquely identifies the project within the gitlab install.
	Id *string `pulumi:"id"`
	// The path of the repository with namespace.
	PathWithNamespace *string `pulumi:"pathWithNamespace"`
	// If true, jobs can be viewed by non-project members.
	PublicBuilds *bool `pulumi:"publicBuilds"`
}

// A collection of values returned by getProject.
type LookupProjectResult struct {
	// Set the analytics access level. Valid values are `disabled`, `private`, `enabled`.
	AnalyticsAccessLevel string `pulumi:"analyticsAccessLevel"`
	// Whether the project is in read-only mode (archived).
	Archived bool `pulumi:"archived"`
	// Auto-cancel pending pipelines. This isn’t a boolean, but enabled/disabled.
	AutoCancelPendingPipelines string `pulumi:"autoCancelPendingPipelines"`
	// Auto Deploy strategy. Valid values are `continuous`, `manual`, `timedIncremental`.
	AutoDevopsDeployStrategy string `pulumi:"autoDevopsDeployStrategy"`
	// Enable Auto DevOps for this project.
	AutoDevopsEnabled bool `pulumi:"autoDevopsEnabled"`
	// Set whether auto-closing referenced issues on default branch.
	AutocloseReferencedIssues bool `pulumi:"autocloseReferencedIssues"`
	// The Git strategy. Defaults to fetch.
	BuildGitStrategy string `pulumi:"buildGitStrategy"`
	// The maximum amount of time, in seconds, that a job can run.
	BuildTimeout int `pulumi:"buildTimeout"`
	// Set the builds access level. Valid values are `disabled`, `private`, `enabled`.
	BuildsAccessLevel string `pulumi:"buildsAccessLevel"`
	// CI config file path for the project.
	CiConfigPath string `pulumi:"ciConfigPath"`
	// Default number of revisions for shallow cloning.
	CiDefaultGitDepth int `pulumi:"ciDefaultGitDepth"`
	// Set the image cleanup policy for this project. **Note**: this field is sometimes named `containerExpirationPolicyAttributes` in the GitLab Upstream API.
	ContainerExpirationPolicies []GetProjectContainerExpirationPolicy `pulumi:"containerExpirationPolicies"`
	// Set visibility of container registry, for this project. Valid values are `disabled`, `private`, `enabled`.
	ContainerRegistryAccessLevel string `pulumi:"containerRegistryAccessLevel"`
	// The default branch for the project.
	DefaultBranch string `pulumi:"defaultBranch"`
	// A description of the project.
	Description string `pulumi:"description"`
	// Disable email notifications.
	EmailsDisabled bool `pulumi:"emailsDisabled"`
	// The classification label for the project.
	ExternalAuthorizationClassificationLabel string `pulumi:"externalAuthorizationClassificationLabel"`
	// Set the forking access level. Valid values are `disabled`, `private`, `enabled`.
	ForkingAccessLevel string `pulumi:"forkingAccessLevel"`
	// URL that can be provided to `git clone` to clone the
	HttpUrlToRepo string `pulumi:"httpUrlToRepo"`
	// The integer or path with namespace that uniquely identifies the project within the gitlab install.
	Id string `pulumi:"id"`
	// Set the issues access level. Valid values are `disabled`, `private`, `enabled`.
	IssuesAccessLevel string `pulumi:"issuesAccessLevel"`
	// Enable issue tracking for the project.
	IssuesEnabled bool `pulumi:"issuesEnabled"`
	// Enable LFS for the project.
	LfsEnabled bool `pulumi:"lfsEnabled"`
	// Template used to create merge commit message in merge requests. (Introduced in GitLab 14.5.)
	MergeCommitTemplate string `pulumi:"mergeCommitTemplate"`
	// Enable or disable merge pipelines.
	MergePipelinesEnabled bool `pulumi:"mergePipelinesEnabled"`
	// Set the merge requests access level. Valid values are `disabled`, `private`, `enabled`.
	MergeRequestsAccessLevel string `pulumi:"mergeRequestsAccessLevel"`
	// Enable merge requests for the project.
	MergeRequestsEnabled bool `pulumi:"mergeRequestsEnabled"`
	// Enable or disable merge trains.
	MergeTrainsEnabled bool `pulumi:"mergeTrainsEnabled"`
	// The name of the project.
	Name string `pulumi:"name"`
	// The namespace (group or user) of the project. Defaults to your user.
	NamespaceId int `pulumi:"namespaceId"`
	// Set the operations access level. Valid values are `disabled`, `private`, `enabled`.
	OperationsAccessLevel string `pulumi:"operationsAccessLevel"`
	// The path of the repository.
	Path string `pulumi:"path"`
	// The path of the repository with namespace.
	PathWithNamespace string `pulumi:"pathWithNamespace"`
	// Enable pipelines for the project.
	PipelinesEnabled bool `pulumi:"pipelinesEnabled"`
	// Show link to create/view merge request when pushing from the command line
	PrintingMergeRequestLinkEnabled bool `pulumi:"printingMergeRequestLinkEnabled"`
	// If true, jobs can be viewed by non-project members.
	PublicBuilds *bool `pulumi:"publicBuilds"`
	// Push rules for the project.
	PushRules GetProjectPushRules `pulumi:"pushRules"`
	// Enable `Delete source branch` option by default for all new merge requests
	RemoveSourceBranchAfterMerge bool `pulumi:"removeSourceBranchAfterMerge"`
	// Set the repository access level. Valid values are `disabled`, `private`, `enabled`.
	RepositoryAccessLevel string `pulumi:"repositoryAccessLevel"`
	// Which storage shard the repository is on. (administrator only)
	RepositoryStorage string `pulumi:"repositoryStorage"`
	// Allow users to request member access.
	RequestAccessEnabled bool `pulumi:"requestAccessEnabled"`
	// Set the requirements access level. Valid values are `disabled`, `private`, `enabled`.
	RequirementsAccessLevel string `pulumi:"requirementsAccessLevel"`
	// Automatically resolve merge request diffs discussions on lines changed with a push.
	ResolveOutdatedDiffDiscussions bool `pulumi:"resolveOutdatedDiffDiscussions"`
	// Registration token to use during runner setup.
	RunnersToken string `pulumi:"runnersToken"`
	// Set the security and compliance access level. Valid values are `disabled`, `private`, `enabled`.
	SecurityAndComplianceAccessLevel string `pulumi:"securityAndComplianceAccessLevel"`
	// Set the snippets access level. Valid values are `disabled`, `private`, `enabled`.
	SnippetsAccessLevel string `pulumi:"snippetsAccessLevel"`
	// Enable snippets for the project.
	SnippetsEnabled bool `pulumi:"snippetsEnabled"`
	// Template used to create squash commit message in merge requests. (Introduced in GitLab 14.6.)
	SquashCommitTemplate string `pulumi:"squashCommitTemplate"`
	// URL that can be provided to `git clone` to clone the
	SshUrlToRepo string `pulumi:"sshUrlToRepo"`
	// The list of topics for the project.
	Topics []string `pulumi:"topics"`
	// Repositories are created as private by default.
	VisibilityLevel string `pulumi:"visibilityLevel"`
	// URL that can be used to find the project in a browser.
	WebUrl string `pulumi:"webUrl"`
	// Set the wiki access level. Valid values are `disabled`, `private`, `enabled`.
	WikiAccessLevel string `pulumi:"wikiAccessLevel"`
	// Enable wiki for the project.
	WikiEnabled bool `pulumi:"wikiEnabled"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectResult, error) {
			args := v.(LookupProjectArgs)
			r, err := LookupProject(ctx, &args, opts...)
			var s LookupProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectResultOutput)
}

// A collection of arguments for invoking getProject.
type LookupProjectOutputArgs struct {
	// Default number of revisions for shallow cloning.
	CiDefaultGitDepth pulumi.IntPtrInput `pulumi:"ciDefaultGitDepth"`
	// The integer or path with namespace that uniquely identifies the project within the gitlab install.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The path of the repository with namespace.
	PathWithNamespace pulumi.StringPtrInput `pulumi:"pathWithNamespace"`
	// If true, jobs can be viewed by non-project members.
	PublicBuilds pulumi.BoolPtrInput `pulumi:"publicBuilds"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

// A collection of values returned by getProject.
type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

// Set the analytics access level. Valid values are `disabled`, `private`, `enabled`.
func (o LookupProjectResultOutput) AnalyticsAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.AnalyticsAccessLevel }).(pulumi.StringOutput)
}

// Whether the project is in read-only mode (archived).
func (o LookupProjectResultOutput) Archived() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.Archived }).(pulumi.BoolOutput)
}

// Auto-cancel pending pipelines. This isn’t a boolean, but enabled/disabled.
func (o LookupProjectResultOutput) AutoCancelPendingPipelines() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.AutoCancelPendingPipelines }).(pulumi.StringOutput)
}

// Auto Deploy strategy. Valid values are `continuous`, `manual`, `timedIncremental`.
func (o LookupProjectResultOutput) AutoDevopsDeployStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.AutoDevopsDeployStrategy }).(pulumi.StringOutput)
}

// Enable Auto DevOps for this project.
func (o LookupProjectResultOutput) AutoDevopsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.AutoDevopsEnabled }).(pulumi.BoolOutput)
}

// Set whether auto-closing referenced issues on default branch.
func (o LookupProjectResultOutput) AutocloseReferencedIssues() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.AutocloseReferencedIssues }).(pulumi.BoolOutput)
}

// The Git strategy. Defaults to fetch.
func (o LookupProjectResultOutput) BuildGitStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.BuildGitStrategy }).(pulumi.StringOutput)
}

// The maximum amount of time, in seconds, that a job can run.
func (o LookupProjectResultOutput) BuildTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectResult) int { return v.BuildTimeout }).(pulumi.IntOutput)
}

// Set the builds access level. Valid values are `disabled`, `private`, `enabled`.
func (o LookupProjectResultOutput) BuildsAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.BuildsAccessLevel }).(pulumi.StringOutput)
}

// CI config file path for the project.
func (o LookupProjectResultOutput) CiConfigPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.CiConfigPath }).(pulumi.StringOutput)
}

// Default number of revisions for shallow cloning.
func (o LookupProjectResultOutput) CiDefaultGitDepth() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectResult) int { return v.CiDefaultGitDepth }).(pulumi.IntOutput)
}

// Set the image cleanup policy for this project. **Note**: this field is sometimes named `containerExpirationPolicyAttributes` in the GitLab Upstream API.
func (o LookupProjectResultOutput) ContainerExpirationPolicies() GetProjectContainerExpirationPolicyArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []GetProjectContainerExpirationPolicy {
		return v.ContainerExpirationPolicies
	}).(GetProjectContainerExpirationPolicyArrayOutput)
}

// Set visibility of container registry, for this project. Valid values are `disabled`, `private`, `enabled`.
func (o LookupProjectResultOutput) ContainerRegistryAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.ContainerRegistryAccessLevel }).(pulumi.StringOutput)
}

// The default branch for the project.
func (o LookupProjectResultOutput) DefaultBranch() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.DefaultBranch }).(pulumi.StringOutput)
}

// A description of the project.
func (o LookupProjectResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Description }).(pulumi.StringOutput)
}

// Disable email notifications.
func (o LookupProjectResultOutput) EmailsDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.EmailsDisabled }).(pulumi.BoolOutput)
}

// The classification label for the project.
func (o LookupProjectResultOutput) ExternalAuthorizationClassificationLabel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.ExternalAuthorizationClassificationLabel }).(pulumi.StringOutput)
}

// Set the forking access level. Valid values are `disabled`, `private`, `enabled`.
func (o LookupProjectResultOutput) ForkingAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.ForkingAccessLevel }).(pulumi.StringOutput)
}

// URL that can be provided to `git clone` to clone the
func (o LookupProjectResultOutput) HttpUrlToRepo() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.HttpUrlToRepo }).(pulumi.StringOutput)
}

// The integer or path with namespace that uniquely identifies the project within the gitlab install.
func (o LookupProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

// Set the issues access level. Valid values are `disabled`, `private`, `enabled`.
func (o LookupProjectResultOutput) IssuesAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.IssuesAccessLevel }).(pulumi.StringOutput)
}

// Enable issue tracking for the project.
func (o LookupProjectResultOutput) IssuesEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.IssuesEnabled }).(pulumi.BoolOutput)
}

// Enable LFS for the project.
func (o LookupProjectResultOutput) LfsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.LfsEnabled }).(pulumi.BoolOutput)
}

// Template used to create merge commit message in merge requests. (Introduced in GitLab 14.5.)
func (o LookupProjectResultOutput) MergeCommitTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.MergeCommitTemplate }).(pulumi.StringOutput)
}

// Enable or disable merge pipelines.
func (o LookupProjectResultOutput) MergePipelinesEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.MergePipelinesEnabled }).(pulumi.BoolOutput)
}

// Set the merge requests access level. Valid values are `disabled`, `private`, `enabled`.
func (o LookupProjectResultOutput) MergeRequestsAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.MergeRequestsAccessLevel }).(pulumi.StringOutput)
}

// Enable merge requests for the project.
func (o LookupProjectResultOutput) MergeRequestsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.MergeRequestsEnabled }).(pulumi.BoolOutput)
}

// Enable or disable merge trains.
func (o LookupProjectResultOutput) MergeTrainsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.MergeTrainsEnabled }).(pulumi.BoolOutput)
}

// The name of the project.
func (o LookupProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

// The namespace (group or user) of the project. Defaults to your user.
func (o LookupProjectResultOutput) NamespaceId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectResult) int { return v.NamespaceId }).(pulumi.IntOutput)
}

// Set the operations access level. Valid values are `disabled`, `private`, `enabled`.
func (o LookupProjectResultOutput) OperationsAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.OperationsAccessLevel }).(pulumi.StringOutput)
}

// The path of the repository.
func (o LookupProjectResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Path }).(pulumi.StringOutput)
}

// The path of the repository with namespace.
func (o LookupProjectResultOutput) PathWithNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.PathWithNamespace }).(pulumi.StringOutput)
}

// Enable pipelines for the project.
func (o LookupProjectResultOutput) PipelinesEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.PipelinesEnabled }).(pulumi.BoolOutput)
}

// Show link to create/view merge request when pushing from the command line
func (o LookupProjectResultOutput) PrintingMergeRequestLinkEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.PrintingMergeRequestLinkEnabled }).(pulumi.BoolOutput)
}

// If true, jobs can be viewed by non-project members.
func (o LookupProjectResultOutput) PublicBuilds() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *bool { return v.PublicBuilds }).(pulumi.BoolPtrOutput)
}

// Push rules for the project.
func (o LookupProjectResultOutput) PushRules() GetProjectPushRulesOutput {
	return o.ApplyT(func(v LookupProjectResult) GetProjectPushRules { return v.PushRules }).(GetProjectPushRulesOutput)
}

// Enable `Delete source branch` option by default for all new merge requests
func (o LookupProjectResultOutput) RemoveSourceBranchAfterMerge() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.RemoveSourceBranchAfterMerge }).(pulumi.BoolOutput)
}

// Set the repository access level. Valid values are `disabled`, `private`, `enabled`.
func (o LookupProjectResultOutput) RepositoryAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.RepositoryAccessLevel }).(pulumi.StringOutput)
}

// Which storage shard the repository is on. (administrator only)
func (o LookupProjectResultOutput) RepositoryStorage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.RepositoryStorage }).(pulumi.StringOutput)
}

// Allow users to request member access.
func (o LookupProjectResultOutput) RequestAccessEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.RequestAccessEnabled }).(pulumi.BoolOutput)
}

// Set the requirements access level. Valid values are `disabled`, `private`, `enabled`.
func (o LookupProjectResultOutput) RequirementsAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.RequirementsAccessLevel }).(pulumi.StringOutput)
}

// Automatically resolve merge request diffs discussions on lines changed with a push.
func (o LookupProjectResultOutput) ResolveOutdatedDiffDiscussions() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.ResolveOutdatedDiffDiscussions }).(pulumi.BoolOutput)
}

// Registration token to use during runner setup.
func (o LookupProjectResultOutput) RunnersToken() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.RunnersToken }).(pulumi.StringOutput)
}

// Set the security and compliance access level. Valid values are `disabled`, `private`, `enabled`.
func (o LookupProjectResultOutput) SecurityAndComplianceAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.SecurityAndComplianceAccessLevel }).(pulumi.StringOutput)
}

// Set the snippets access level. Valid values are `disabled`, `private`, `enabled`.
func (o LookupProjectResultOutput) SnippetsAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.SnippetsAccessLevel }).(pulumi.StringOutput)
}

// Enable snippets for the project.
func (o LookupProjectResultOutput) SnippetsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.SnippetsEnabled }).(pulumi.BoolOutput)
}

// Template used to create squash commit message in merge requests. (Introduced in GitLab 14.6.)
func (o LookupProjectResultOutput) SquashCommitTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.SquashCommitTemplate }).(pulumi.StringOutput)
}

// URL that can be provided to `git clone` to clone the
func (o LookupProjectResultOutput) SshUrlToRepo() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.SshUrlToRepo }).(pulumi.StringOutput)
}

// The list of topics for the project.
func (o LookupProjectResultOutput) Topics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []string { return v.Topics }).(pulumi.StringArrayOutput)
}

// Repositories are created as private by default.
func (o LookupProjectResultOutput) VisibilityLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.VisibilityLevel }).(pulumi.StringOutput)
}

// URL that can be used to find the project in a browser.
func (o LookupProjectResultOutput) WebUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.WebUrl }).(pulumi.StringOutput)
}

// Set the wiki access level. Valid values are `disabled`, `private`, `enabled`.
func (o LookupProjectResultOutput) WikiAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.WikiAccessLevel }).(pulumi.StringOutput)
}

// Enable wiki for the project.
func (o LookupProjectResultOutput) WikiEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.WikiEnabled }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
