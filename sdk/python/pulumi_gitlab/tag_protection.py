# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['TagProtectionArgs', 'TagProtection']

@pulumi.input_type
class TagProtectionArgs:
    def __init__(__self__, *,
                 create_access_level: pulumi.Input[str],
                 project: pulumi.Input[str],
                 tag: pulumi.Input[str]):
        """
        The set of arguments for constructing a TagProtection resource.
        :param pulumi.Input[str] create_access_level: One of five levels of access to the project.
        :param pulumi.Input[str] project: The id of the project.
        :param pulumi.Input[str] tag: Name of the tag or wildcard.
        """
        pulumi.set(__self__, "create_access_level", create_access_level)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "tag", tag)

    @property
    @pulumi.getter(name="createAccessLevel")
    def create_access_level(self) -> pulumi.Input[str]:
        """
        One of five levels of access to the project.
        """
        return pulumi.get(self, "create_access_level")

    @create_access_level.setter
    def create_access_level(self, value: pulumi.Input[str]):
        pulumi.set(self, "create_access_level", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The id of the project.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def tag(self) -> pulumi.Input[str]:
        """
        Name of the tag or wildcard.
        """
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: pulumi.Input[str]):
        pulumi.set(self, "tag", value)


@pulumi.input_type
class _TagProtectionState:
    def __init__(__self__, *,
                 create_access_level: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TagProtection resources.
        :param pulumi.Input[str] create_access_level: One of five levels of access to the project.
        :param pulumi.Input[str] project: The id of the project.
        :param pulumi.Input[str] tag: Name of the tag or wildcard.
        """
        if create_access_level is not None:
            pulumi.set(__self__, "create_access_level", create_access_level)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if tag is not None:
            pulumi.set(__self__, "tag", tag)

    @property
    @pulumi.getter(name="createAccessLevel")
    def create_access_level(self) -> Optional[pulumi.Input[str]]:
        """
        One of five levels of access to the project.
        """
        return pulumi.get(self, "create_access_level")

    @create_access_level.setter
    def create_access_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_access_level", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the project.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def tag(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the tag or wildcard.
        """
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag", value)


class TagProtection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_access_level: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # gitlab\_tag\_protection

        This resource allows you to protect a specific tag or wildcard by an access level so that the user with less access level cannot Create the tags.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        tag_protect = gitlab.TagProtection("tagProtect",
            create_access_level="developer",
            project="12345",
            tag="TagProtected")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_access_level: One of five levels of access to the project.
        :param pulumi.Input[str] project: The id of the project.
        :param pulumi.Input[str] tag: Name of the tag or wildcard.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TagProtectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # gitlab\_tag\_protection

        This resource allows you to protect a specific tag or wildcard by an access level so that the user with less access level cannot Create the tags.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gitlab as gitlab

        tag_protect = gitlab.TagProtection("tagProtect",
            create_access_level="developer",
            project="12345",
            tag="TagProtected")
        ```

        :param str resource_name: The name of the resource.
        :param TagProtectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TagProtectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_access_level: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TagProtectionArgs.__new__(TagProtectionArgs)

            if create_access_level is None and not opts.urn:
                raise TypeError("Missing required property 'create_access_level'")
            __props__.__dict__["create_access_level"] = create_access_level
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if tag is None and not opts.urn:
                raise TypeError("Missing required property 'tag'")
            __props__.__dict__["tag"] = tag
        super(TagProtection, __self__).__init__(
            'gitlab:index/tagProtection:TagProtection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_access_level: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            tag: Optional[pulumi.Input[str]] = None) -> 'TagProtection':
        """
        Get an existing TagProtection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_access_level: One of five levels of access to the project.
        :param pulumi.Input[str] project: The id of the project.
        :param pulumi.Input[str] tag: Name of the tag or wildcard.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TagProtectionState.__new__(_TagProtectionState)

        __props__.__dict__["create_access_level"] = create_access_level
        __props__.__dict__["project"] = project
        __props__.__dict__["tag"] = tag
        return TagProtection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createAccessLevel")
    def create_access_level(self) -> pulumi.Output[str]:
        """
        One of five levels of access to the project.
        """
        return pulumi.get(self, "create_access_level")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The id of the project.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def tag(self) -> pulumi.Output[str]:
        """
        Name of the tag or wildcard.
        """
        return pulumi.get(self, "tag")

