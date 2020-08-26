# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Project']


class Project(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 approvals_before_merge: Optional[pulumi.Input[float]] = None,
                 archived: Optional[pulumi.Input[bool]] = None,
                 container_registry_enabled: Optional[pulumi.Input[bool]] = None,
                 default_branch: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 initialize_with_readme: Optional[pulumi.Input[bool]] = None,
                 issues_enabled: Optional[pulumi.Input[bool]] = None,
                 lfs_enabled: Optional[pulumi.Input[bool]] = None,
                 merge_method: Optional[pulumi.Input[str]] = None,
                 merge_requests_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[float]] = None,
                 only_allow_merge_if_all_discussions_are_resolved: Optional[pulumi.Input[bool]] = None,
                 only_allow_merge_if_pipeline_succeeds: Optional[pulumi.Input[bool]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 pipelines_enabled: Optional[pulumi.Input[bool]] = None,
                 remove_source_branch_after_merge: Optional[pulumi.Input[bool]] = None,
                 request_access_enabled: Optional[pulumi.Input[bool]] = None,
                 shared_runners_enabled: Optional[pulumi.Input[bool]] = None,
                 shared_with_groups: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ProjectSharedWithGroupArgs']]]]] = None,
                 snippets_enabled: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 visibility_level: Optional[pulumi.Input[str]] = None,
                 wiki_enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Project resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] approvals_before_merge: Number of merge request approvals required for merging. Default is 0.
        :param pulumi.Input[bool] archived: Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
        :param pulumi.Input[bool] container_registry_enabled: Enable container registry for the project.
        :param pulumi.Input[str] default_branch: The default branch for the project.
        :param pulumi.Input[str] description: A description of the project.
        :param pulumi.Input[bool] initialize_with_readme: Create master branch with first commit containing a README.md file.
        :param pulumi.Input[bool] issues_enabled: Enable issue tracking for the project.
        :param pulumi.Input[bool] lfs_enabled: Enable LFS for the project.
        :param pulumi.Input[str] merge_method: Set to `ff` to create fast-forward merges
               Valid values are `merge`, `rebase_merge`, `ff`
               Repositories are created with `merge` by default
        :param pulumi.Input[bool] merge_requests_enabled: Enable merge requests for the project.
        :param pulumi.Input[str] name: The name of the project.
        :param pulumi.Input[float] namespace_id: The namespace (group or user) of the project. Defaults to your user.
               See `Group` for an example.
        :param pulumi.Input[bool] only_allow_merge_if_all_discussions_are_resolved: Set to true if you want allow merges only if all discussions are resolved.
        :param pulumi.Input[bool] only_allow_merge_if_pipeline_succeeds: Set to true if you want allow merges only if a pipeline succeeds.
        :param pulumi.Input[str] path: The path of the repository.
        :param pulumi.Input[bool] pipelines_enabled: Enable pipelines for the project.
        :param pulumi.Input[bool] remove_source_branch_after_merge: Enable `Delete source branch` option by default for all new merge requests.
        :param pulumi.Input[bool] request_access_enabled: Allow users to request member access.
        :param pulumi.Input[bool] shared_runners_enabled: Enable shared runners for this project.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ProjectSharedWithGroupArgs']]]] shared_with_groups: Enable sharing the project with a list of groups (maps).
        :param pulumi.Input[bool] snippets_enabled: Enable snippets for the project.
        :param pulumi.Input[List[pulumi.Input[str]]] tags: Tags (topics) of the project.
        :param pulumi.Input[str] visibility_level: Set to `public` to create a public project.
               Valid values are `private`, `internal`, `public`.
               Repositories are created as private by default.
        :param pulumi.Input[bool] wiki_enabled: Enable wiki for the project.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['approvals_before_merge'] = approvals_before_merge
            __props__['archived'] = archived
            __props__['container_registry_enabled'] = container_registry_enabled
            __props__['default_branch'] = default_branch
            __props__['description'] = description
            __props__['initialize_with_readme'] = initialize_with_readme
            __props__['issues_enabled'] = issues_enabled
            __props__['lfs_enabled'] = lfs_enabled
            __props__['merge_method'] = merge_method
            __props__['merge_requests_enabled'] = merge_requests_enabled
            __props__['name'] = name
            __props__['namespace_id'] = namespace_id
            __props__['only_allow_merge_if_all_discussions_are_resolved'] = only_allow_merge_if_all_discussions_are_resolved
            __props__['only_allow_merge_if_pipeline_succeeds'] = only_allow_merge_if_pipeline_succeeds
            __props__['path'] = path
            __props__['pipelines_enabled'] = pipelines_enabled
            __props__['remove_source_branch_after_merge'] = remove_source_branch_after_merge
            __props__['request_access_enabled'] = request_access_enabled
            __props__['shared_runners_enabled'] = shared_runners_enabled
            __props__['shared_with_groups'] = shared_with_groups
            __props__['snippets_enabled'] = snippets_enabled
            __props__['tags'] = tags
            __props__['visibility_level'] = visibility_level
            __props__['wiki_enabled'] = wiki_enabled
            __props__['http_url_to_repo'] = None
            __props__['runners_token'] = None
            __props__['ssh_url_to_repo'] = None
            __props__['web_url'] = None
        super(Project, __self__).__init__(
            'gitlab:index/project:Project',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            approvals_before_merge: Optional[pulumi.Input[float]] = None,
            archived: Optional[pulumi.Input[bool]] = None,
            container_registry_enabled: Optional[pulumi.Input[bool]] = None,
            default_branch: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            http_url_to_repo: Optional[pulumi.Input[str]] = None,
            initialize_with_readme: Optional[pulumi.Input[bool]] = None,
            issues_enabled: Optional[pulumi.Input[bool]] = None,
            lfs_enabled: Optional[pulumi.Input[bool]] = None,
            merge_method: Optional[pulumi.Input[str]] = None,
            merge_requests_enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace_id: Optional[pulumi.Input[float]] = None,
            only_allow_merge_if_all_discussions_are_resolved: Optional[pulumi.Input[bool]] = None,
            only_allow_merge_if_pipeline_succeeds: Optional[pulumi.Input[bool]] = None,
            path: Optional[pulumi.Input[str]] = None,
            pipelines_enabled: Optional[pulumi.Input[bool]] = None,
            remove_source_branch_after_merge: Optional[pulumi.Input[bool]] = None,
            request_access_enabled: Optional[pulumi.Input[bool]] = None,
            runners_token: Optional[pulumi.Input[str]] = None,
            shared_runners_enabled: Optional[pulumi.Input[bool]] = None,
            shared_with_groups: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ProjectSharedWithGroupArgs']]]]] = None,
            snippets_enabled: Optional[pulumi.Input[bool]] = None,
            ssh_url_to_repo: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            visibility_level: Optional[pulumi.Input[str]] = None,
            web_url: Optional[pulumi.Input[str]] = None,
            wiki_enabled: Optional[pulumi.Input[bool]] = None) -> 'Project':
        """
        Get an existing Project resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] approvals_before_merge: Number of merge request approvals required for merging. Default is 0.
        :param pulumi.Input[bool] archived: Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
        :param pulumi.Input[bool] container_registry_enabled: Enable container registry for the project.
        :param pulumi.Input[str] default_branch: The default branch for the project.
        :param pulumi.Input[str] description: A description of the project.
        :param pulumi.Input[str] http_url_to_repo: URL that can be provided to `git clone` to clone the
               repository via HTTP.
        :param pulumi.Input[bool] initialize_with_readme: Create master branch with first commit containing a README.md file.
        :param pulumi.Input[bool] issues_enabled: Enable issue tracking for the project.
        :param pulumi.Input[bool] lfs_enabled: Enable LFS for the project.
        :param pulumi.Input[str] merge_method: Set to `ff` to create fast-forward merges
               Valid values are `merge`, `rebase_merge`, `ff`
               Repositories are created with `merge` by default
        :param pulumi.Input[bool] merge_requests_enabled: Enable merge requests for the project.
        :param pulumi.Input[str] name: The name of the project.
        :param pulumi.Input[float] namespace_id: The namespace (group or user) of the project. Defaults to your user.
               See `Group` for an example.
        :param pulumi.Input[bool] only_allow_merge_if_all_discussions_are_resolved: Set to true if you want allow merges only if all discussions are resolved.
        :param pulumi.Input[bool] only_allow_merge_if_pipeline_succeeds: Set to true if you want allow merges only if a pipeline succeeds.
        :param pulumi.Input[str] path: The path of the repository.
        :param pulumi.Input[bool] pipelines_enabled: Enable pipelines for the project.
        :param pulumi.Input[bool] remove_source_branch_after_merge: Enable `Delete source branch` option by default for all new merge requests.
        :param pulumi.Input[bool] request_access_enabled: Allow users to request member access.
        :param pulumi.Input[str] runners_token: Registration token to use during runner setup.
        :param pulumi.Input[bool] shared_runners_enabled: Enable shared runners for this project.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ProjectSharedWithGroupArgs']]]] shared_with_groups: Enable sharing the project with a list of groups (maps).
        :param pulumi.Input[bool] snippets_enabled: Enable snippets for the project.
        :param pulumi.Input[str] ssh_url_to_repo: URL that can be provided to `git clone` to clone the
               repository via SSH.
        :param pulumi.Input[List[pulumi.Input[str]]] tags: Tags (topics) of the project.
        :param pulumi.Input[str] visibility_level: Set to `public` to create a public project.
               Valid values are `private`, `internal`, `public`.
               Repositories are created as private by default.
        :param pulumi.Input[str] web_url: URL that can be used to find the project in a browser.
        :param pulumi.Input[bool] wiki_enabled: Enable wiki for the project.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["approvals_before_merge"] = approvals_before_merge
        __props__["archived"] = archived
        __props__["container_registry_enabled"] = container_registry_enabled
        __props__["default_branch"] = default_branch
        __props__["description"] = description
        __props__["http_url_to_repo"] = http_url_to_repo
        __props__["initialize_with_readme"] = initialize_with_readme
        __props__["issues_enabled"] = issues_enabled
        __props__["lfs_enabled"] = lfs_enabled
        __props__["merge_method"] = merge_method
        __props__["merge_requests_enabled"] = merge_requests_enabled
        __props__["name"] = name
        __props__["namespace_id"] = namespace_id
        __props__["only_allow_merge_if_all_discussions_are_resolved"] = only_allow_merge_if_all_discussions_are_resolved
        __props__["only_allow_merge_if_pipeline_succeeds"] = only_allow_merge_if_pipeline_succeeds
        __props__["path"] = path
        __props__["pipelines_enabled"] = pipelines_enabled
        __props__["remove_source_branch_after_merge"] = remove_source_branch_after_merge
        __props__["request_access_enabled"] = request_access_enabled
        __props__["runners_token"] = runners_token
        __props__["shared_runners_enabled"] = shared_runners_enabled
        __props__["shared_with_groups"] = shared_with_groups
        __props__["snippets_enabled"] = snippets_enabled
        __props__["ssh_url_to_repo"] = ssh_url_to_repo
        __props__["tags"] = tags
        __props__["visibility_level"] = visibility_level
        __props__["web_url"] = web_url
        __props__["wiki_enabled"] = wiki_enabled
        return Project(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="approvalsBeforeMerge")
    def approvals_before_merge(self) -> Optional[float]:
        """
        Number of merge request approvals required for merging. Default is 0.
        """
        return pulumi.get(self, "approvals_before_merge")

    @property
    @pulumi.getter
    def archived(self) -> Optional[bool]:
        """
        Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
        """
        return pulumi.get(self, "archived")

    @property
    @pulumi.getter(name="containerRegistryEnabled")
    def container_registry_enabled(self) -> Optional[bool]:
        """
        Enable container registry for the project.
        """
        return pulumi.get(self, "container_registry_enabled")

    @property
    @pulumi.getter(name="defaultBranch")
    def default_branch(self) -> Optional[str]:
        """
        The default branch for the project.
        """
        return pulumi.get(self, "default_branch")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A description of the project.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="httpUrlToRepo")
    def http_url_to_repo(self) -> str:
        """
        URL that can be provided to `git clone` to clone the
        repository via HTTP.
        """
        return pulumi.get(self, "http_url_to_repo")

    @property
    @pulumi.getter(name="initializeWithReadme")
    def initialize_with_readme(self) -> Optional[bool]:
        """
        Create master branch with first commit containing a README.md file.
        """
        return pulumi.get(self, "initialize_with_readme")

    @property
    @pulumi.getter(name="issuesEnabled")
    def issues_enabled(self) -> Optional[bool]:
        """
        Enable issue tracking for the project.
        """
        return pulumi.get(self, "issues_enabled")

    @property
    @pulumi.getter(name="lfsEnabled")
    def lfs_enabled(self) -> Optional[bool]:
        """
        Enable LFS for the project.
        """
        return pulumi.get(self, "lfs_enabled")

    @property
    @pulumi.getter(name="mergeMethod")
    def merge_method(self) -> Optional[str]:
        """
        Set to `ff` to create fast-forward merges
        Valid values are `merge`, `rebase_merge`, `ff`
        Repositories are created with `merge` by default
        """
        return pulumi.get(self, "merge_method")

    @property
    @pulumi.getter(name="mergeRequestsEnabled")
    def merge_requests_enabled(self) -> Optional[bool]:
        """
        Enable merge requests for the project.
        """
        return pulumi.get(self, "merge_requests_enabled")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the project.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> float:
        """
        The namespace (group or user) of the project. Defaults to your user.
        See `Group` for an example.
        """
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter(name="onlyAllowMergeIfAllDiscussionsAreResolved")
    def only_allow_merge_if_all_discussions_are_resolved(self) -> Optional[bool]:
        """
        Set to true if you want allow merges only if all discussions are resolved.
        """
        return pulumi.get(self, "only_allow_merge_if_all_discussions_are_resolved")

    @property
    @pulumi.getter(name="onlyAllowMergeIfPipelineSucceeds")
    def only_allow_merge_if_pipeline_succeeds(self) -> Optional[bool]:
        """
        Set to true if you want allow merges only if a pipeline succeeds.
        """
        return pulumi.get(self, "only_allow_merge_if_pipeline_succeeds")

    @property
    @pulumi.getter
    def path(self) -> Optional[str]:
        """
        The path of the repository.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="pipelinesEnabled")
    def pipelines_enabled(self) -> Optional[bool]:
        """
        Enable pipelines for the project.
        """
        return pulumi.get(self, "pipelines_enabled")

    @property
    @pulumi.getter(name="removeSourceBranchAfterMerge")
    def remove_source_branch_after_merge(self) -> Optional[bool]:
        """
        Enable `Delete source branch` option by default for all new merge requests.
        """
        return pulumi.get(self, "remove_source_branch_after_merge")

    @property
    @pulumi.getter(name="requestAccessEnabled")
    def request_access_enabled(self) -> Optional[bool]:
        """
        Allow users to request member access.
        """
        return pulumi.get(self, "request_access_enabled")

    @property
    @pulumi.getter(name="runnersToken")
    def runners_token(self) -> str:
        """
        Registration token to use during runner setup.
        """
        return pulumi.get(self, "runners_token")

    @property
    @pulumi.getter(name="sharedRunnersEnabled")
    def shared_runners_enabled(self) -> bool:
        """
        Enable shared runners for this project.
        """
        return pulumi.get(self, "shared_runners_enabled")

    @property
    @pulumi.getter(name="sharedWithGroups")
    def shared_with_groups(self) -> Optional[List['outputs.ProjectSharedWithGroup']]:
        """
        Enable sharing the project with a list of groups (maps).
        """
        return pulumi.get(self, "shared_with_groups")

    @property
    @pulumi.getter(name="snippetsEnabled")
    def snippets_enabled(self) -> Optional[bool]:
        """
        Enable snippets for the project.
        """
        return pulumi.get(self, "snippets_enabled")

    @property
    @pulumi.getter(name="sshUrlToRepo")
    def ssh_url_to_repo(self) -> str:
        """
        URL that can be provided to `git clone` to clone the
        repository via SSH.
        """
        return pulumi.get(self, "ssh_url_to_repo")

    @property
    @pulumi.getter
    def tags(self) -> Optional[List[str]]:
        """
        Tags (topics) of the project.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="visibilityLevel")
    def visibility_level(self) -> Optional[str]:
        """
        Set to `public` to create a public project.
        Valid values are `private`, `internal`, `public`.
        Repositories are created as private by default.
        """
        return pulumi.get(self, "visibility_level")

    @property
    @pulumi.getter(name="webUrl")
    def web_url(self) -> str:
        """
        URL that can be used to find the project in a browser.
        """
        return pulumi.get(self, "web_url")

    @property
    @pulumi.getter(name="wikiEnabled")
    def wiki_enabled(self) -> Optional[bool]:
        """
        Enable wiki for the project.
        """
        return pulumi.get(self, "wiki_enabled")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

