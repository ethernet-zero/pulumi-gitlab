// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Project struct {
	pulumi.CustomResourceState

	// Number of merge request approvals required for merging. Default is 0.
	ApprovalsBeforeMerge pulumi.IntPtrOutput `pulumi:"approvalsBeforeMerge"`
	// Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
	Archived pulumi.BoolPtrOutput `pulumi:"archived"`
	// Enable container registry for the project.
	ContainerRegistryEnabled pulumi.BoolPtrOutput `pulumi:"containerRegistryEnabled"`
	// The default branch for the project.
	DefaultBranch pulumi.StringPtrOutput `pulumi:"defaultBranch"`
	// A description of the project.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// For group-level custom templates, specifies ID of group from which all the custom project templates are sourced. Leave empty for instance-level templates. Requires useCustomTemplate to be true (enterprise edition).
	GroupWithProjectTemplatesId pulumi.IntPtrOutput `pulumi:"groupWithProjectTemplatesId"`
	// URL that can be provided to `git clone` to clone the
	// repository via HTTP.
	HttpUrlToRepo pulumi.StringOutput `pulumi:"httpUrlToRepo"`
	// Git URL to a repository to be imported.
	ImportUrl pulumi.StringPtrOutput `pulumi:"importUrl"`
	// Create master branch with first commit containing a README.md file.
	InitializeWithReadme pulumi.BoolPtrOutput `pulumi:"initializeWithReadme"`
	// Enable issue tracking for the project.
	IssuesEnabled pulumi.BoolPtrOutput `pulumi:"issuesEnabled"`
	// Enable LFS for the project.
	LfsEnabled pulumi.BoolPtrOutput `pulumi:"lfsEnabled"`
	// Set to `ff` to create fast-forward merges
	// Valid values are `merge`, `rebaseMerge`, `ff`
	// Repositories are created with `merge` by default
	MergeMethod pulumi.StringPtrOutput `pulumi:"mergeMethod"`
	// Enable merge requests for the project.
	MergeRequestsEnabled pulumi.BoolPtrOutput `pulumi:"mergeRequestsEnabled"`
	// Enables pull mirroring in a project. Default is `false`. For further information on mirroring,
	// consult the [gitlab documentation](https://docs.gitlab.com/ee/user/project/repository/repository_mirroring.html#repository-mirroring).
	Mirror pulumi.BoolPtrOutput `pulumi:"mirror"`
	// Pull mirror overwrites diverged branches.
	MirrorOverwritesDivergedBranches pulumi.BoolPtrOutput `pulumi:"mirrorOverwritesDivergedBranches"`
	// Pull mirroring triggers builds. Default is `false`.
	MirrorTriggerBuilds pulumi.BoolPtrOutput `pulumi:"mirrorTriggerBuilds"`
	// The name of the project.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace (group or user) of the project. Defaults to your user.
	// See `Group` for an example.
	NamespaceId pulumi.IntOutput `pulumi:"namespaceId"`
	// Set to true if you want allow merges only if all discussions are resolved.
	OnlyAllowMergeIfAllDiscussionsAreResolved pulumi.BoolPtrOutput `pulumi:"onlyAllowMergeIfAllDiscussionsAreResolved"`
	// Set to true if you want allow merges only if a pipeline succeeds.
	OnlyAllowMergeIfPipelineSucceeds pulumi.BoolPtrOutput `pulumi:"onlyAllowMergeIfPipelineSucceeds"`
	// Only mirror protected branches.
	OnlyMirrorProtectedBranches pulumi.BoolPtrOutput `pulumi:"onlyMirrorProtectedBranches"`
	// Enable packages repository for the project.
	PackagesEnabled pulumi.BoolPtrOutput `pulumi:"packagesEnabled"`
	// Enable pages access control
	// Valid values are `disabled`, `private`, `enabled`, `public`.
	// `private` is the default.
	PagesAccessLevel pulumi.StringPtrOutput `pulumi:"pagesAccessLevel"`
	// The path of the repository.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// The path of the repository with namespace.
	PathWithNamespace pulumi.StringOutput `pulumi:"pathWithNamespace"`
	// Enable pipelines for the project.
	PipelinesEnabled pulumi.BoolPtrOutput `pulumi:"pipelinesEnabled"`
	// Push rules for the project (documented below).
	PushRules ProjectPushRulesOutput `pulumi:"pushRules"`
	// Enable `Delete source branch` option by default for all new merge requests.
	RemoveSourceBranchAfterMerge pulumi.BoolPtrOutput `pulumi:"removeSourceBranchAfterMerge"`
	// Allow users to request member access.
	RequestAccessEnabled pulumi.BoolPtrOutput `pulumi:"requestAccessEnabled"`
	// Registration token to use during runner setup.
	RunnersToken pulumi.StringOutput `pulumi:"runnersToken"`
	// Enable shared runners for this project.
	SharedRunnersEnabled pulumi.BoolOutput `pulumi:"sharedRunnersEnabled"`
	// Enable snippets for the project.
	SnippetsEnabled pulumi.BoolPtrOutput `pulumi:"snippetsEnabled"`
	// URL that can be provided to `git clone` to clone the
	// repository via SSH.
	SshUrlToRepo pulumi.StringOutput `pulumi:"sshUrlToRepo"`
	// Tags (topics) of the project.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// When used without use_custom_template, name of a built-in project template. When used with use_custom_template, name of a custom project template. This option is mutually exclusive with `templateProjectId`.
	TemplateName pulumi.StringPtrOutput `pulumi:"templateName"`
	// When used with use_custom_template, project ID of a custom project template. This is preferable to using templateName since templateName may be ambiguous (enterprise edition). This option is mutually exclusive with `templateName`.
	TemplateProjectId pulumi.IntPtrOutput `pulumi:"templateProjectId"`
	// Use either custom instance or group (with group_with_project_templates_id) project template (enterprise edition).
	UseCustomTemplate pulumi.BoolPtrOutput `pulumi:"useCustomTemplate"`
	// Set to `public` to create a public project.
	// Valid values are `private`, `internal`, `public`.
	// Repositories are created as private by default.
	VisibilityLevel pulumi.StringPtrOutput `pulumi:"visibilityLevel"`
	// URL that can be used to find the project in a browser.
	WebUrl pulumi.StringOutput `pulumi:"webUrl"`
	// Enable wiki for the project.
	WikiEnabled pulumi.BoolPtrOutput `pulumi:"wikiEnabled"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		args = &ProjectArgs{}
	}

	var resource Project
	err := ctx.RegisterResource("gitlab:index/project:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("gitlab:index/project:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
	// Number of merge request approvals required for merging. Default is 0.
	ApprovalsBeforeMerge *int `pulumi:"approvalsBeforeMerge"`
	// Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
	Archived *bool `pulumi:"archived"`
	// Enable container registry for the project.
	ContainerRegistryEnabled *bool `pulumi:"containerRegistryEnabled"`
	// The default branch for the project.
	DefaultBranch *string `pulumi:"defaultBranch"`
	// A description of the project.
	Description *string `pulumi:"description"`
	// For group-level custom templates, specifies ID of group from which all the custom project templates are sourced. Leave empty for instance-level templates. Requires useCustomTemplate to be true (enterprise edition).
	GroupWithProjectTemplatesId *int `pulumi:"groupWithProjectTemplatesId"`
	// URL that can be provided to `git clone` to clone the
	// repository via HTTP.
	HttpUrlToRepo *string `pulumi:"httpUrlToRepo"`
	// Git URL to a repository to be imported.
	ImportUrl *string `pulumi:"importUrl"`
	// Create master branch with first commit containing a README.md file.
	InitializeWithReadme *bool `pulumi:"initializeWithReadme"`
	// Enable issue tracking for the project.
	IssuesEnabled *bool `pulumi:"issuesEnabled"`
	// Enable LFS for the project.
	LfsEnabled *bool `pulumi:"lfsEnabled"`
	// Set to `ff` to create fast-forward merges
	// Valid values are `merge`, `rebaseMerge`, `ff`
	// Repositories are created with `merge` by default
	MergeMethod *string `pulumi:"mergeMethod"`
	// Enable merge requests for the project.
	MergeRequestsEnabled *bool `pulumi:"mergeRequestsEnabled"`
	// Enables pull mirroring in a project. Default is `false`. For further information on mirroring,
	// consult the [gitlab documentation](https://docs.gitlab.com/ee/user/project/repository/repository_mirroring.html#repository-mirroring).
	Mirror *bool `pulumi:"mirror"`
	// Pull mirror overwrites diverged branches.
	MirrorOverwritesDivergedBranches *bool `pulumi:"mirrorOverwritesDivergedBranches"`
	// Pull mirroring triggers builds. Default is `false`.
	MirrorTriggerBuilds *bool `pulumi:"mirrorTriggerBuilds"`
	// The name of the project.
	Name *string `pulumi:"name"`
	// The namespace (group or user) of the project. Defaults to your user.
	// See `Group` for an example.
	NamespaceId *int `pulumi:"namespaceId"`
	// Set to true if you want allow merges only if all discussions are resolved.
	OnlyAllowMergeIfAllDiscussionsAreResolved *bool `pulumi:"onlyAllowMergeIfAllDiscussionsAreResolved"`
	// Set to true if you want allow merges only if a pipeline succeeds.
	OnlyAllowMergeIfPipelineSucceeds *bool `pulumi:"onlyAllowMergeIfPipelineSucceeds"`
	// Only mirror protected branches.
	OnlyMirrorProtectedBranches *bool `pulumi:"onlyMirrorProtectedBranches"`
	// Enable packages repository for the project.
	PackagesEnabled *bool `pulumi:"packagesEnabled"`
	// Enable pages access control
	// Valid values are `disabled`, `private`, `enabled`, `public`.
	// `private` is the default.
	PagesAccessLevel *string `pulumi:"pagesAccessLevel"`
	// The path of the repository.
	Path *string `pulumi:"path"`
	// The path of the repository with namespace.
	PathWithNamespace *string `pulumi:"pathWithNamespace"`
	// Enable pipelines for the project.
	PipelinesEnabled *bool `pulumi:"pipelinesEnabled"`
	// Push rules for the project (documented below).
	PushRules *ProjectPushRules `pulumi:"pushRules"`
	// Enable `Delete source branch` option by default for all new merge requests.
	RemoveSourceBranchAfterMerge *bool `pulumi:"removeSourceBranchAfterMerge"`
	// Allow users to request member access.
	RequestAccessEnabled *bool `pulumi:"requestAccessEnabled"`
	// Registration token to use during runner setup.
	RunnersToken *string `pulumi:"runnersToken"`
	// Enable shared runners for this project.
	SharedRunnersEnabled *bool `pulumi:"sharedRunnersEnabled"`
	// Enable snippets for the project.
	SnippetsEnabled *bool `pulumi:"snippetsEnabled"`
	// URL that can be provided to `git clone` to clone the
	// repository via SSH.
	SshUrlToRepo *string `pulumi:"sshUrlToRepo"`
	// Tags (topics) of the project.
	Tags []string `pulumi:"tags"`
	// When used without use_custom_template, name of a built-in project template. When used with use_custom_template, name of a custom project template. This option is mutually exclusive with `templateProjectId`.
	TemplateName *string `pulumi:"templateName"`
	// When used with use_custom_template, project ID of a custom project template. This is preferable to using templateName since templateName may be ambiguous (enterprise edition). This option is mutually exclusive with `templateName`.
	TemplateProjectId *int `pulumi:"templateProjectId"`
	// Use either custom instance or group (with group_with_project_templates_id) project template (enterprise edition).
	UseCustomTemplate *bool `pulumi:"useCustomTemplate"`
	// Set to `public` to create a public project.
	// Valid values are `private`, `internal`, `public`.
	// Repositories are created as private by default.
	VisibilityLevel *string `pulumi:"visibilityLevel"`
	// URL that can be used to find the project in a browser.
	WebUrl *string `pulumi:"webUrl"`
	// Enable wiki for the project.
	WikiEnabled *bool `pulumi:"wikiEnabled"`
}

type ProjectState struct {
	// Number of merge request approvals required for merging. Default is 0.
	ApprovalsBeforeMerge pulumi.IntPtrInput
	// Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
	Archived pulumi.BoolPtrInput
	// Enable container registry for the project.
	ContainerRegistryEnabled pulumi.BoolPtrInput
	// The default branch for the project.
	DefaultBranch pulumi.StringPtrInput
	// A description of the project.
	Description pulumi.StringPtrInput
	// For group-level custom templates, specifies ID of group from which all the custom project templates are sourced. Leave empty for instance-level templates. Requires useCustomTemplate to be true (enterprise edition).
	GroupWithProjectTemplatesId pulumi.IntPtrInput
	// URL that can be provided to `git clone` to clone the
	// repository via HTTP.
	HttpUrlToRepo pulumi.StringPtrInput
	// Git URL to a repository to be imported.
	ImportUrl pulumi.StringPtrInput
	// Create master branch with first commit containing a README.md file.
	InitializeWithReadme pulumi.BoolPtrInput
	// Enable issue tracking for the project.
	IssuesEnabled pulumi.BoolPtrInput
	// Enable LFS for the project.
	LfsEnabled pulumi.BoolPtrInput
	// Set to `ff` to create fast-forward merges
	// Valid values are `merge`, `rebaseMerge`, `ff`
	// Repositories are created with `merge` by default
	MergeMethod pulumi.StringPtrInput
	// Enable merge requests for the project.
	MergeRequestsEnabled pulumi.BoolPtrInput
	// Enables pull mirroring in a project. Default is `false`. For further information on mirroring,
	// consult the [gitlab documentation](https://docs.gitlab.com/ee/user/project/repository/repository_mirroring.html#repository-mirroring).
	Mirror pulumi.BoolPtrInput
	// Pull mirror overwrites diverged branches.
	MirrorOverwritesDivergedBranches pulumi.BoolPtrInput
	// Pull mirroring triggers builds. Default is `false`.
	MirrorTriggerBuilds pulumi.BoolPtrInput
	// The name of the project.
	Name pulumi.StringPtrInput
	// The namespace (group or user) of the project. Defaults to your user.
	// See `Group` for an example.
	NamespaceId pulumi.IntPtrInput
	// Set to true if you want allow merges only if all discussions are resolved.
	OnlyAllowMergeIfAllDiscussionsAreResolved pulumi.BoolPtrInput
	// Set to true if you want allow merges only if a pipeline succeeds.
	OnlyAllowMergeIfPipelineSucceeds pulumi.BoolPtrInput
	// Only mirror protected branches.
	OnlyMirrorProtectedBranches pulumi.BoolPtrInput
	// Enable packages repository for the project.
	PackagesEnabled pulumi.BoolPtrInput
	// Enable pages access control
	// Valid values are `disabled`, `private`, `enabled`, `public`.
	// `private` is the default.
	PagesAccessLevel pulumi.StringPtrInput
	// The path of the repository.
	Path pulumi.StringPtrInput
	// The path of the repository with namespace.
	PathWithNamespace pulumi.StringPtrInput
	// Enable pipelines for the project.
	PipelinesEnabled pulumi.BoolPtrInput
	// Push rules for the project (documented below).
	PushRules ProjectPushRulesPtrInput
	// Enable `Delete source branch` option by default for all new merge requests.
	RemoveSourceBranchAfterMerge pulumi.BoolPtrInput
	// Allow users to request member access.
	RequestAccessEnabled pulumi.BoolPtrInput
	// Registration token to use during runner setup.
	RunnersToken pulumi.StringPtrInput
	// Enable shared runners for this project.
	SharedRunnersEnabled pulumi.BoolPtrInput
	// Enable snippets for the project.
	SnippetsEnabled pulumi.BoolPtrInput
	// URL that can be provided to `git clone` to clone the
	// repository via SSH.
	SshUrlToRepo pulumi.StringPtrInput
	// Tags (topics) of the project.
	Tags pulumi.StringArrayInput
	// When used without use_custom_template, name of a built-in project template. When used with use_custom_template, name of a custom project template. This option is mutually exclusive with `templateProjectId`.
	TemplateName pulumi.StringPtrInput
	// When used with use_custom_template, project ID of a custom project template. This is preferable to using templateName since templateName may be ambiguous (enterprise edition). This option is mutually exclusive with `templateName`.
	TemplateProjectId pulumi.IntPtrInput
	// Use either custom instance or group (with group_with_project_templates_id) project template (enterprise edition).
	UseCustomTemplate pulumi.BoolPtrInput
	// Set to `public` to create a public project.
	// Valid values are `private`, `internal`, `public`.
	// Repositories are created as private by default.
	VisibilityLevel pulumi.StringPtrInput
	// URL that can be used to find the project in a browser.
	WebUrl pulumi.StringPtrInput
	// Enable wiki for the project.
	WikiEnabled pulumi.BoolPtrInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// Number of merge request approvals required for merging. Default is 0.
	ApprovalsBeforeMerge *int `pulumi:"approvalsBeforeMerge"`
	// Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
	Archived *bool `pulumi:"archived"`
	// Enable container registry for the project.
	ContainerRegistryEnabled *bool `pulumi:"containerRegistryEnabled"`
	// The default branch for the project.
	DefaultBranch *string `pulumi:"defaultBranch"`
	// A description of the project.
	Description *string `pulumi:"description"`
	// For group-level custom templates, specifies ID of group from which all the custom project templates are sourced. Leave empty for instance-level templates. Requires useCustomTemplate to be true (enterprise edition).
	GroupWithProjectTemplatesId *int `pulumi:"groupWithProjectTemplatesId"`
	// Git URL to a repository to be imported.
	ImportUrl *string `pulumi:"importUrl"`
	// Create master branch with first commit containing a README.md file.
	InitializeWithReadme *bool `pulumi:"initializeWithReadme"`
	// Enable issue tracking for the project.
	IssuesEnabled *bool `pulumi:"issuesEnabled"`
	// Enable LFS for the project.
	LfsEnabled *bool `pulumi:"lfsEnabled"`
	// Set to `ff` to create fast-forward merges
	// Valid values are `merge`, `rebaseMerge`, `ff`
	// Repositories are created with `merge` by default
	MergeMethod *string `pulumi:"mergeMethod"`
	// Enable merge requests for the project.
	MergeRequestsEnabled *bool `pulumi:"mergeRequestsEnabled"`
	// Enables pull mirroring in a project. Default is `false`. For further information on mirroring,
	// consult the [gitlab documentation](https://docs.gitlab.com/ee/user/project/repository/repository_mirroring.html#repository-mirroring).
	Mirror *bool `pulumi:"mirror"`
	// Pull mirror overwrites diverged branches.
	MirrorOverwritesDivergedBranches *bool `pulumi:"mirrorOverwritesDivergedBranches"`
	// Pull mirroring triggers builds. Default is `false`.
	MirrorTriggerBuilds *bool `pulumi:"mirrorTriggerBuilds"`
	// The name of the project.
	Name *string `pulumi:"name"`
	// The namespace (group or user) of the project. Defaults to your user.
	// See `Group` for an example.
	NamespaceId *int `pulumi:"namespaceId"`
	// Set to true if you want allow merges only if all discussions are resolved.
	OnlyAllowMergeIfAllDiscussionsAreResolved *bool `pulumi:"onlyAllowMergeIfAllDiscussionsAreResolved"`
	// Set to true if you want allow merges only if a pipeline succeeds.
	OnlyAllowMergeIfPipelineSucceeds *bool `pulumi:"onlyAllowMergeIfPipelineSucceeds"`
	// Only mirror protected branches.
	OnlyMirrorProtectedBranches *bool `pulumi:"onlyMirrorProtectedBranches"`
	// Enable packages repository for the project.
	PackagesEnabled *bool `pulumi:"packagesEnabled"`
	// Enable pages access control
	// Valid values are `disabled`, `private`, `enabled`, `public`.
	// `private` is the default.
	PagesAccessLevel *string `pulumi:"pagesAccessLevel"`
	// The path of the repository.
	Path *string `pulumi:"path"`
	// Enable pipelines for the project.
	PipelinesEnabled *bool `pulumi:"pipelinesEnabled"`
	// Push rules for the project (documented below).
	PushRules *ProjectPushRules `pulumi:"pushRules"`
	// Enable `Delete source branch` option by default for all new merge requests.
	RemoveSourceBranchAfterMerge *bool `pulumi:"removeSourceBranchAfterMerge"`
	// Allow users to request member access.
	RequestAccessEnabled *bool `pulumi:"requestAccessEnabled"`
	// Enable shared runners for this project.
	SharedRunnersEnabled *bool `pulumi:"sharedRunnersEnabled"`
	// Enable snippets for the project.
	SnippetsEnabled *bool `pulumi:"snippetsEnabled"`
	// Tags (topics) of the project.
	Tags []string `pulumi:"tags"`
	// When used without use_custom_template, name of a built-in project template. When used with use_custom_template, name of a custom project template. This option is mutually exclusive with `templateProjectId`.
	TemplateName *string `pulumi:"templateName"`
	// When used with use_custom_template, project ID of a custom project template. This is preferable to using templateName since templateName may be ambiguous (enterprise edition). This option is mutually exclusive with `templateName`.
	TemplateProjectId *int `pulumi:"templateProjectId"`
	// Use either custom instance or group (with group_with_project_templates_id) project template (enterprise edition).
	UseCustomTemplate *bool `pulumi:"useCustomTemplate"`
	// Set to `public` to create a public project.
	// Valid values are `private`, `internal`, `public`.
	// Repositories are created as private by default.
	VisibilityLevel *string `pulumi:"visibilityLevel"`
	// Enable wiki for the project.
	WikiEnabled *bool `pulumi:"wikiEnabled"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// Number of merge request approvals required for merging. Default is 0.
	ApprovalsBeforeMerge pulumi.IntPtrInput
	// Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
	Archived pulumi.BoolPtrInput
	// Enable container registry for the project.
	ContainerRegistryEnabled pulumi.BoolPtrInput
	// The default branch for the project.
	DefaultBranch pulumi.StringPtrInput
	// A description of the project.
	Description pulumi.StringPtrInput
	// For group-level custom templates, specifies ID of group from which all the custom project templates are sourced. Leave empty for instance-level templates. Requires useCustomTemplate to be true (enterprise edition).
	GroupWithProjectTemplatesId pulumi.IntPtrInput
	// Git URL to a repository to be imported.
	ImportUrl pulumi.StringPtrInput
	// Create master branch with first commit containing a README.md file.
	InitializeWithReadme pulumi.BoolPtrInput
	// Enable issue tracking for the project.
	IssuesEnabled pulumi.BoolPtrInput
	// Enable LFS for the project.
	LfsEnabled pulumi.BoolPtrInput
	// Set to `ff` to create fast-forward merges
	// Valid values are `merge`, `rebaseMerge`, `ff`
	// Repositories are created with `merge` by default
	MergeMethod pulumi.StringPtrInput
	// Enable merge requests for the project.
	MergeRequestsEnabled pulumi.BoolPtrInput
	// Enables pull mirroring in a project. Default is `false`. For further information on mirroring,
	// consult the [gitlab documentation](https://docs.gitlab.com/ee/user/project/repository/repository_mirroring.html#repository-mirroring).
	Mirror pulumi.BoolPtrInput
	// Pull mirror overwrites diverged branches.
	MirrorOverwritesDivergedBranches pulumi.BoolPtrInput
	// Pull mirroring triggers builds. Default is `false`.
	MirrorTriggerBuilds pulumi.BoolPtrInput
	// The name of the project.
	Name pulumi.StringPtrInput
	// The namespace (group or user) of the project. Defaults to your user.
	// See `Group` for an example.
	NamespaceId pulumi.IntPtrInput
	// Set to true if you want allow merges only if all discussions are resolved.
	OnlyAllowMergeIfAllDiscussionsAreResolved pulumi.BoolPtrInput
	// Set to true if you want allow merges only if a pipeline succeeds.
	OnlyAllowMergeIfPipelineSucceeds pulumi.BoolPtrInput
	// Only mirror protected branches.
	OnlyMirrorProtectedBranches pulumi.BoolPtrInput
	// Enable packages repository for the project.
	PackagesEnabled pulumi.BoolPtrInput
	// Enable pages access control
	// Valid values are `disabled`, `private`, `enabled`, `public`.
	// `private` is the default.
	PagesAccessLevel pulumi.StringPtrInput
	// The path of the repository.
	Path pulumi.StringPtrInput
	// Enable pipelines for the project.
	PipelinesEnabled pulumi.BoolPtrInput
	// Push rules for the project (documented below).
	PushRules ProjectPushRulesPtrInput
	// Enable `Delete source branch` option by default for all new merge requests.
	RemoveSourceBranchAfterMerge pulumi.BoolPtrInput
	// Allow users to request member access.
	RequestAccessEnabled pulumi.BoolPtrInput
	// Enable shared runners for this project.
	SharedRunnersEnabled pulumi.BoolPtrInput
	// Enable snippets for the project.
	SnippetsEnabled pulumi.BoolPtrInput
	// Tags (topics) of the project.
	Tags pulumi.StringArrayInput
	// When used without use_custom_template, name of a built-in project template. When used with use_custom_template, name of a custom project template. This option is mutually exclusive with `templateProjectId`.
	TemplateName pulumi.StringPtrInput
	// When used with use_custom_template, project ID of a custom project template. This is preferable to using templateName since templateName may be ambiguous (enterprise edition). This option is mutually exclusive with `templateName`.
	TemplateProjectId pulumi.IntPtrInput
	// Use either custom instance or group (with group_with_project_templates_id) project template (enterprise edition).
	UseCustomTemplate pulumi.BoolPtrInput
	// Set to `public` to create a public project.
	// Valid values are `private`, `internal`, `public`.
	// Repositories are created as private by default.
	VisibilityLevel pulumi.StringPtrInput
	// Enable wiki for the project.
	WikiEnabled pulumi.BoolPtrInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (*Project) ElementType() reflect.Type {
	return reflect.TypeOf((*Project)(nil))
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

func (i *Project) ToProjectPtrOutput() ProjectPtrOutput {
	return i.ToProjectPtrOutputWithContext(context.Background())
}

func (i *Project) ToProjectPtrOutputWithContext(ctx context.Context) ProjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectPtrOutput)
}

type ProjectPtrInput interface {
	pulumi.Input

	ToProjectPtrOutput() ProjectPtrOutput
	ToProjectPtrOutputWithContext(ctx context.Context) ProjectPtrOutput
}

type projectPtrType ProjectArgs

func (*projectPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil))
}

func (i *projectPtrType) ToProjectPtrOutput() ProjectPtrOutput {
	return i.ToProjectPtrOutputWithContext(context.Background())
}

func (i *projectPtrType) ToProjectPtrOutputWithContext(ctx context.Context) ProjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectPtrOutput)
}

// ProjectArrayInput is an input type that accepts ProjectArray and ProjectArrayOutput values.
// You can construct a concrete instance of `ProjectArrayInput` via:
//
//          ProjectArray{ ProjectArgs{...} }
type ProjectArrayInput interface {
	pulumi.Input

	ToProjectArrayOutput() ProjectArrayOutput
	ToProjectArrayOutputWithContext(context.Context) ProjectArrayOutput
}

type ProjectArray []ProjectInput

func (ProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Project)(nil))
}

func (i ProjectArray) ToProjectArrayOutput() ProjectArrayOutput {
	return i.ToProjectArrayOutputWithContext(context.Background())
}

func (i ProjectArray) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectArrayOutput)
}

// ProjectMapInput is an input type that accepts ProjectMap and ProjectMapOutput values.
// You can construct a concrete instance of `ProjectMapInput` via:
//
//          ProjectMap{ "key": ProjectArgs{...} }
type ProjectMapInput interface {
	pulumi.Input

	ToProjectMapOutput() ProjectMapOutput
	ToProjectMapOutputWithContext(context.Context) ProjectMapOutput
}

type ProjectMap map[string]ProjectInput

func (ProjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Project)(nil))
}

func (i ProjectMap) ToProjectMapOutput() ProjectMapOutput {
	return i.ToProjectMapOutputWithContext(context.Background())
}

func (i ProjectMap) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMapOutput)
}

type ProjectOutput struct {
	*pulumi.OutputState
}

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Project)(nil))
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectPtrOutput() ProjectPtrOutput {
	return o.ToProjectPtrOutputWithContext(context.Background())
}

func (o ProjectOutput) ToProjectPtrOutputWithContext(ctx context.Context) ProjectPtrOutput {
	return o.ApplyT(func(v Project) *Project {
		return &v
	}).(ProjectPtrOutput)
}

type ProjectPtrOutput struct {
	*pulumi.OutputState
}

func (ProjectPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil))
}

func (o ProjectPtrOutput) ToProjectPtrOutput() ProjectPtrOutput {
	return o
}

func (o ProjectPtrOutput) ToProjectPtrOutputWithContext(ctx context.Context) ProjectPtrOutput {
	return o
}

type ProjectArrayOutput struct{ *pulumi.OutputState }

func (ProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Project)(nil))
}

func (o ProjectArrayOutput) ToProjectArrayOutput() ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) Index(i pulumi.IntInput) ProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Project {
		return vs[0].([]Project)[vs[1].(int)]
	}).(ProjectOutput)
}

type ProjectMapOutput struct{ *pulumi.OutputState }

func (ProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Project)(nil))
}

func (o ProjectMapOutput) ToProjectMapOutput() ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) MapIndex(k pulumi.StringInput) ProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Project {
		return vs[0].(map[string]Project)[vs[1].(string)]
	}).(ProjectOutput)
}

func init() {
	pulumi.RegisterOutputType(ProjectOutput{})
	pulumi.RegisterOutputType(ProjectPtrOutput{})
	pulumi.RegisterOutputType(ProjectArrayOutput{})
	pulumi.RegisterOutputType(ProjectMapOutput{})
}
