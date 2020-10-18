# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import outputs

__all__ = [
    'GetGroupMembershipResult',
    'AwaitableGetGroupMembershipResult',
    'get_group_membership',
]

@pulumi.output_type
class GetGroupMembershipResult:
    """
    A collection of values returned by getGroupMembership.
    """
    def __init__(__self__, access_level=None, full_path=None, group_id=None, id=None, members=None):
        if access_level and not isinstance(access_level, str):
            raise TypeError("Expected argument 'access_level' to be a str")
        pulumi.set(__self__, "access_level", access_level)
        if full_path and not isinstance(full_path, str):
            raise TypeError("Expected argument 'full_path' to be a str")
        pulumi.set(__self__, "full_path", full_path)
        if group_id and not isinstance(group_id, int):
            raise TypeError("Expected argument 'group_id' to be a int")
        pulumi.set(__self__, "group_id", group_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if members and not isinstance(members, list):
            raise TypeError("Expected argument 'members' to be a list")
        pulumi.set(__self__, "members", members)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> str:
        """
        One of five levels of access to the group.
        """
        return pulumi.get(self, "access_level")

    @property
    @pulumi.getter(name="fullPath")
    def full_path(self) -> str:
        return pulumi.get(self, "full_path")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> int:
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def members(self) -> Sequence['outputs.GetGroupMembershipMemberResult']:
        """
        The list of group members.
        """
        return pulumi.get(self, "members")


class AwaitableGetGroupMembershipResult(GetGroupMembershipResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupMembershipResult(
            access_level=self.access_level,
            full_path=self.full_path,
            group_id=self.group_id,
            id=self.id,
            members=self.members)


def get_group_membership(access_level: Optional[str] = None,
                         full_path: Optional[str] = None,
                         group_id: Optional[int] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupMembershipResult:
    """
    ## # gitlab\_group\_membership

    Provides details about a list of group members in the gitlab provider. The results include id, username, name and more about the requested members.

    ## Example Usage

    **By group's ID**

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    example = gitlab.get_group_membership(group_id=123)
    ```

    **By group's full path**

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    example = gitlab.get_group_membership(full_path="foo/bar")
    ```


    :param str access_level: Only return members with the desidered access level. Acceptable values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
    :param str full_path: The full path of the group.
    :param int group_id: The ID of the group.
    """
    __args__ = dict()
    __args__['accessLevel'] = access_level
    __args__['fullPath'] = full_path
    __args__['groupId'] = group_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gitlab:index/getGroupMembership:getGroupMembership', __args__, opts=opts, typ=GetGroupMembershipResult).value

    return AwaitableGetGroupMembershipResult(
        access_level=__ret__.access_level,
        full_path=__ret__.full_path,
        group_id=__ret__.group_id,
        id=__ret__.id,
        members=__ret__.members)
