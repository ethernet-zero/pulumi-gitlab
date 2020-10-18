# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['BranchProtection']


class BranchProtection(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 code_owner_approval_required: Optional[pulumi.Input[bool]] = None,
                 merge_access_level: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 push_access_level: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## # gitlab\_branch_protection

        This resource allows you to protect a specific branch by an access level so that the user with less access level cannot Merge/Push to the branch. GitLab EE features to protect by group or user are not supported.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        branch_protect = gitlab.BranchProtection("branchProtect",
            branch="BranchProtected",
            merge_access_level="developer",
            project="12345",
            push_access_level="developer")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] branch: Name of the branch.
        :param pulumi.Input[bool] code_owner_approval_required: Bool, defaults to false. Can be set to true to require code owner approval before merging.
        :param pulumi.Input[str] merge_access_level: One of five levels of access to the project.
        :param pulumi.Input[str] project: The id of the project.
        :param pulumi.Input[str] push_access_level: One of five levels of access to the project.
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

            if branch is None:
                raise TypeError("Missing required property 'branch'")
            __props__['branch'] = branch
            __props__['code_owner_approval_required'] = code_owner_approval_required
            if merge_access_level is None:
                raise TypeError("Missing required property 'merge_access_level'")
            __props__['merge_access_level'] = merge_access_level
            if project is None:
                raise TypeError("Missing required property 'project'")
            __props__['project'] = project
            if push_access_level is None:
                raise TypeError("Missing required property 'push_access_level'")
            __props__['push_access_level'] = push_access_level
        super(BranchProtection, __self__).__init__(
            'gitlab:index/branchProtection:BranchProtection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            branch: Optional[pulumi.Input[str]] = None,
            code_owner_approval_required: Optional[pulumi.Input[bool]] = None,
            merge_access_level: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            push_access_level: Optional[pulumi.Input[str]] = None) -> 'BranchProtection':
        """
        Get an existing BranchProtection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] branch: Name of the branch.
        :param pulumi.Input[bool] code_owner_approval_required: Bool, defaults to false. Can be set to true to require code owner approval before merging.
        :param pulumi.Input[str] merge_access_level: One of five levels of access to the project.
        :param pulumi.Input[str] project: The id of the project.
        :param pulumi.Input[str] push_access_level: One of five levels of access to the project.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["branch"] = branch
        __props__["code_owner_approval_required"] = code_owner_approval_required
        __props__["merge_access_level"] = merge_access_level
        __props__["project"] = project
        __props__["push_access_level"] = push_access_level
        return BranchProtection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def branch(self) -> pulumi.Output[str]:
        """
        Name of the branch.
        """
        return pulumi.get(self, "branch")

    @property
    @pulumi.getter(name="codeOwnerApprovalRequired")
    def code_owner_approval_required(self) -> pulumi.Output[Optional[bool]]:
        """
        Bool, defaults to false. Can be set to true to require code owner approval before merging.
        """
        return pulumi.get(self, "code_owner_approval_required")

    @property
    @pulumi.getter(name="mergeAccessLevel")
    def merge_access_level(self) -> pulumi.Output[str]:
        """
        One of five levels of access to the project.
        """
        return pulumi.get(self, "merge_access_level")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The id of the project.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="pushAccessLevel")
    def push_access_level(self) -> pulumi.Output[str]:
        """
        One of five levels of access to the project.
        """
        return pulumi.get(self, "push_access_level")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

