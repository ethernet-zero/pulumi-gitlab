# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities, _tables

__all__ = ['ProjectLevelMrApprovalsArgs', 'ProjectLevelMrApprovals']

@pulumi.input_type
class ProjectLevelMrApprovalsArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[int],
                 disable_overriding_approvers_per_merge_request: Optional[pulumi.Input[bool]] = None,
                 merge_requests_author_approval: Optional[pulumi.Input[bool]] = None,
                 merge_requests_disable_committers_approval: Optional[pulumi.Input[bool]] = None,
                 reset_approvals_on_push: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a ProjectLevelMrApprovals resource.
        :param pulumi.Input[int] project_id: The ID of the project to change MR approval configuration.
        :param pulumi.Input[bool] disable_overriding_approvers_per_merge_request: By default, users are able to edit the approval rules in merge requests. If set to true,
               the approval rules for all new merge requests will be determined by the default approval rules. Default is `false`.
        :param pulumi.Input[bool] merge_requests_author_approval: Set to `true` if you want to allow merge request authors to self-approve merge requests. Authors
               also need to be included in the approvers list in order to be able to approve their merge request. Default is `false`.
        :param pulumi.Input[bool] merge_requests_disable_committers_approval: Set to `true` if you want to prevent approval of merge requests by merge request committers. Default is `false`.
        :param pulumi.Input[bool] reset_approvals_on_push: Set to `true` if you want to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
        """
        pulumi.set(__self__, "project_id", project_id)
        if disable_overriding_approvers_per_merge_request is not None:
            pulumi.set(__self__, "disable_overriding_approvers_per_merge_request", disable_overriding_approvers_per_merge_request)
        if merge_requests_author_approval is not None:
            pulumi.set(__self__, "merge_requests_author_approval", merge_requests_author_approval)
        if merge_requests_disable_committers_approval is not None:
            pulumi.set(__self__, "merge_requests_disable_committers_approval", merge_requests_disable_committers_approval)
        if reset_approvals_on_push is not None:
            pulumi.set(__self__, "reset_approvals_on_push", reset_approvals_on_push)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[int]:
        """
        The ID of the project to change MR approval configuration.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="disableOverridingApproversPerMergeRequest")
    def disable_overriding_approvers_per_merge_request(self) -> Optional[pulumi.Input[bool]]:
        """
        By default, users are able to edit the approval rules in merge requests. If set to true,
        the approval rules for all new merge requests will be determined by the default approval rules. Default is `false`.
        """
        return pulumi.get(self, "disable_overriding_approvers_per_merge_request")

    @disable_overriding_approvers_per_merge_request.setter
    def disable_overriding_approvers_per_merge_request(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_overriding_approvers_per_merge_request", value)

    @property
    @pulumi.getter(name="mergeRequestsAuthorApproval")
    def merge_requests_author_approval(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` if you want to allow merge request authors to self-approve merge requests. Authors
        also need to be included in the approvers list in order to be able to approve their merge request. Default is `false`.
        """
        return pulumi.get(self, "merge_requests_author_approval")

    @merge_requests_author_approval.setter
    def merge_requests_author_approval(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "merge_requests_author_approval", value)

    @property
    @pulumi.getter(name="mergeRequestsDisableCommittersApproval")
    def merge_requests_disable_committers_approval(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` if you want to prevent approval of merge requests by merge request committers. Default is `false`.
        """
        return pulumi.get(self, "merge_requests_disable_committers_approval")

    @merge_requests_disable_committers_approval.setter
    def merge_requests_disable_committers_approval(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "merge_requests_disable_committers_approval", value)

    @property
    @pulumi.getter(name="resetApprovalsOnPush")
    def reset_approvals_on_push(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` if you want to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
        """
        return pulumi.get(self, "reset_approvals_on_push")

    @reset_approvals_on_push.setter
    def reset_approvals_on_push(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "reset_approvals_on_push", value)


class ProjectLevelMrApprovals(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_overriding_approvers_per_merge_request: Optional[pulumi.Input[bool]] = None,
                 merge_requests_author_approval: Optional[pulumi.Input[bool]] = None,
                 merge_requests_disable_committers_approval: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 reset_approvals_on_push: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a ProjectLevelMrApprovals resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disable_overriding_approvers_per_merge_request: By default, users are able to edit the approval rules in merge requests. If set to true,
               the approval rules for all new merge requests will be determined by the default approval rules. Default is `false`.
        :param pulumi.Input[bool] merge_requests_author_approval: Set to `true` if you want to allow merge request authors to self-approve merge requests. Authors
               also need to be included in the approvers list in order to be able to approve their merge request. Default is `false`.
        :param pulumi.Input[bool] merge_requests_disable_committers_approval: Set to `true` if you want to prevent approval of merge requests by merge request committers. Default is `false`.
        :param pulumi.Input[int] project_id: The ID of the project to change MR approval configuration.
        :param pulumi.Input[bool] reset_approvals_on_push: Set to `true` if you want to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectLevelMrApprovalsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ProjectLevelMrApprovals resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ProjectLevelMrApprovalsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectLevelMrApprovalsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_overriding_approvers_per_merge_request: Optional[pulumi.Input[bool]] = None,
                 merge_requests_author_approval: Optional[pulumi.Input[bool]] = None,
                 merge_requests_disable_committers_approval: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[int]] = None,
                 reset_approvals_on_push: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            __props__['disable_overriding_approvers_per_merge_request'] = disable_overriding_approvers_per_merge_request
            __props__['merge_requests_author_approval'] = merge_requests_author_approval
            __props__['merge_requests_disable_committers_approval'] = merge_requests_disable_committers_approval
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__['project_id'] = project_id
            __props__['reset_approvals_on_push'] = reset_approvals_on_push
        super(ProjectLevelMrApprovals, __self__).__init__(
            'gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            disable_overriding_approvers_per_merge_request: Optional[pulumi.Input[bool]] = None,
            merge_requests_author_approval: Optional[pulumi.Input[bool]] = None,
            merge_requests_disable_committers_approval: Optional[pulumi.Input[bool]] = None,
            project_id: Optional[pulumi.Input[int]] = None,
            reset_approvals_on_push: Optional[pulumi.Input[bool]] = None) -> 'ProjectLevelMrApprovals':
        """
        Get an existing ProjectLevelMrApprovals resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disable_overriding_approvers_per_merge_request: By default, users are able to edit the approval rules in merge requests. If set to true,
               the approval rules for all new merge requests will be determined by the default approval rules. Default is `false`.
        :param pulumi.Input[bool] merge_requests_author_approval: Set to `true` if you want to allow merge request authors to self-approve merge requests. Authors
               also need to be included in the approvers list in order to be able to approve their merge request. Default is `false`.
        :param pulumi.Input[bool] merge_requests_disable_committers_approval: Set to `true` if you want to prevent approval of merge requests by merge request committers. Default is `false`.
        :param pulumi.Input[int] project_id: The ID of the project to change MR approval configuration.
        :param pulumi.Input[bool] reset_approvals_on_push: Set to `true` if you want to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["disable_overriding_approvers_per_merge_request"] = disable_overriding_approvers_per_merge_request
        __props__["merge_requests_author_approval"] = merge_requests_author_approval
        __props__["merge_requests_disable_committers_approval"] = merge_requests_disable_committers_approval
        __props__["project_id"] = project_id
        __props__["reset_approvals_on_push"] = reset_approvals_on_push
        return ProjectLevelMrApprovals(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="disableOverridingApproversPerMergeRequest")
    def disable_overriding_approvers_per_merge_request(self) -> pulumi.Output[Optional[bool]]:
        """
        By default, users are able to edit the approval rules in merge requests. If set to true,
        the approval rules for all new merge requests will be determined by the default approval rules. Default is `false`.
        """
        return pulumi.get(self, "disable_overriding_approvers_per_merge_request")

    @property
    @pulumi.getter(name="mergeRequestsAuthorApproval")
    def merge_requests_author_approval(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to `true` if you want to allow merge request authors to self-approve merge requests. Authors
        also need to be included in the approvers list in order to be able to approve their merge request. Default is `false`.
        """
        return pulumi.get(self, "merge_requests_author_approval")

    @property
    @pulumi.getter(name="mergeRequestsDisableCommittersApproval")
    def merge_requests_disable_committers_approval(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to `true` if you want to prevent approval of merge requests by merge request committers. Default is `false`.
        """
        return pulumi.get(self, "merge_requests_disable_committers_approval")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[int]:
        """
        The ID of the project to change MR approval configuration.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="resetApprovalsOnPush")
    def reset_approvals_on_push(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to `true` if you want to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
        """
        return pulumi.get(self, "reset_approvals_on_push")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

