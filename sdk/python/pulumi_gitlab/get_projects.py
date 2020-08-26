# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables
from . import outputs

__all__ = [
    'GetProjectsResult',
    'AwaitableGetProjectsResult',
    'get_projects',
]

@pulumi.output_type
class GetProjectsResult:
    """
    A collection of values returned by getProjects.
    """
    def __init__(__self__, archived=None, group_id=None, id=None, include_subgroups=None, max_queryable_pages=None, membership=None, min_access_level=None, order_by=None, owned=None, page=None, per_page=None, projects=None, search=None, simple=None, sort=None, starred=None, statistics=None, visibility=None, with_custom_attributes=None, with_issues_enabled=None, with_merge_requests_enabled=None, with_programming_language=None, with_shared=None):
        if archived and not isinstance(archived, bool):
            raise TypeError("Expected argument 'archived' to be a bool")
        pulumi.set(__self__, "archived", archived)
        if group_id and not isinstance(group_id, float):
            raise TypeError("Expected argument 'group_id' to be a float")
        pulumi.set(__self__, "group_id", group_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if include_subgroups and not isinstance(include_subgroups, bool):
            raise TypeError("Expected argument 'include_subgroups' to be a bool")
        pulumi.set(__self__, "include_subgroups", include_subgroups)
        if max_queryable_pages and not isinstance(max_queryable_pages, float):
            raise TypeError("Expected argument 'max_queryable_pages' to be a float")
        pulumi.set(__self__, "max_queryable_pages", max_queryable_pages)
        if membership and not isinstance(membership, bool):
            raise TypeError("Expected argument 'membership' to be a bool")
        pulumi.set(__self__, "membership", membership)
        if min_access_level and not isinstance(min_access_level, float):
            raise TypeError("Expected argument 'min_access_level' to be a float")
        pulumi.set(__self__, "min_access_level", min_access_level)
        if order_by and not isinstance(order_by, str):
            raise TypeError("Expected argument 'order_by' to be a str")
        pulumi.set(__self__, "order_by", order_by)
        if owned and not isinstance(owned, bool):
            raise TypeError("Expected argument 'owned' to be a bool")
        pulumi.set(__self__, "owned", owned)
        if page and not isinstance(page, float):
            raise TypeError("Expected argument 'page' to be a float")
        pulumi.set(__self__, "page", page)
        if per_page and not isinstance(per_page, float):
            raise TypeError("Expected argument 'per_page' to be a float")
        pulumi.set(__self__, "per_page", per_page)
        if projects and not isinstance(projects, list):
            raise TypeError("Expected argument 'projects' to be a list")
        pulumi.set(__self__, "projects", projects)
        if search and not isinstance(search, str):
            raise TypeError("Expected argument 'search' to be a str")
        pulumi.set(__self__, "search", search)
        if simple and not isinstance(simple, bool):
            raise TypeError("Expected argument 'simple' to be a bool")
        pulumi.set(__self__, "simple", simple)
        if sort and not isinstance(sort, str):
            raise TypeError("Expected argument 'sort' to be a str")
        pulumi.set(__self__, "sort", sort)
        if starred and not isinstance(starred, bool):
            raise TypeError("Expected argument 'starred' to be a bool")
        pulumi.set(__self__, "starred", starred)
        if statistics and not isinstance(statistics, bool):
            raise TypeError("Expected argument 'statistics' to be a bool")
        pulumi.set(__self__, "statistics", statistics)
        if visibility and not isinstance(visibility, str):
            raise TypeError("Expected argument 'visibility' to be a str")
        pulumi.set(__self__, "visibility", visibility)
        if with_custom_attributes and not isinstance(with_custom_attributes, bool):
            raise TypeError("Expected argument 'with_custom_attributes' to be a bool")
        pulumi.set(__self__, "with_custom_attributes", with_custom_attributes)
        if with_issues_enabled and not isinstance(with_issues_enabled, bool):
            raise TypeError("Expected argument 'with_issues_enabled' to be a bool")
        pulumi.set(__self__, "with_issues_enabled", with_issues_enabled)
        if with_merge_requests_enabled and not isinstance(with_merge_requests_enabled, bool):
            raise TypeError("Expected argument 'with_merge_requests_enabled' to be a bool")
        pulumi.set(__self__, "with_merge_requests_enabled", with_merge_requests_enabled)
        if with_programming_language and not isinstance(with_programming_language, str):
            raise TypeError("Expected argument 'with_programming_language' to be a str")
        pulumi.set(__self__, "with_programming_language", with_programming_language)
        if with_shared and not isinstance(with_shared, bool):
            raise TypeError("Expected argument 'with_shared' to be a bool")
        pulumi.set(__self__, "with_shared", with_shared)

    @property
    @pulumi.getter
    def archived(self) -> Optional[bool]:
        return pulumi.get(self, "archived")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[float]:
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="includeSubgroups")
    def include_subgroups(self) -> Optional[bool]:
        return pulumi.get(self, "include_subgroups")

    @property
    @pulumi.getter(name="maxQueryablePages")
    def max_queryable_pages(self) -> Optional[float]:
        return pulumi.get(self, "max_queryable_pages")

    @property
    @pulumi.getter
    def membership(self) -> Optional[bool]:
        return pulumi.get(self, "membership")

    @property
    @pulumi.getter(name="minAccessLevel")
    def min_access_level(self) -> Optional[float]:
        return pulumi.get(self, "min_access_level")

    @property
    @pulumi.getter(name="orderBy")
    def order_by(self) -> Optional[str]:
        return pulumi.get(self, "order_by")

    @property
    @pulumi.getter
    def owned(self) -> Optional[bool]:
        return pulumi.get(self, "owned")

    @property
    @pulumi.getter
    def page(self) -> Optional[float]:
        return pulumi.get(self, "page")

    @property
    @pulumi.getter(name="perPage")
    def per_page(self) -> Optional[float]:
        return pulumi.get(self, "per_page")

    @property
    @pulumi.getter
    def projects(self) -> List['outputs.GetProjectsProjectResult']:
        """
        A list containing the projects matching the supplied arguments
        """
        return pulumi.get(self, "projects")

    @property
    @pulumi.getter
    def search(self) -> Optional[str]:
        return pulumi.get(self, "search")

    @property
    @pulumi.getter
    def simple(self) -> Optional[bool]:
        return pulumi.get(self, "simple")

    @property
    @pulumi.getter
    def sort(self) -> Optional[str]:
        return pulumi.get(self, "sort")

    @property
    @pulumi.getter
    def starred(self) -> Optional[bool]:
        return pulumi.get(self, "starred")

    @property
    @pulumi.getter
    def statistics(self) -> Optional[bool]:
        return pulumi.get(self, "statistics")

    @property
    @pulumi.getter
    def visibility(self) -> Optional[str]:
        """
        The visibility of the project.
        """
        return pulumi.get(self, "visibility")

    @property
    @pulumi.getter(name="withCustomAttributes")
    def with_custom_attributes(self) -> Optional[bool]:
        return pulumi.get(self, "with_custom_attributes")

    @property
    @pulumi.getter(name="withIssuesEnabled")
    def with_issues_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "with_issues_enabled")

    @property
    @pulumi.getter(name="withMergeRequestsEnabled")
    def with_merge_requests_enabled(self) -> Optional[bool]:
        return pulumi.get(self, "with_merge_requests_enabled")

    @property
    @pulumi.getter(name="withProgrammingLanguage")
    def with_programming_language(self) -> Optional[str]:
        return pulumi.get(self, "with_programming_language")

    @property
    @pulumi.getter(name="withShared")
    def with_shared(self) -> Optional[bool]:
        return pulumi.get(self, "with_shared")


class AwaitableGetProjectsResult(GetProjectsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectsResult(
            archived=self.archived,
            group_id=self.group_id,
            id=self.id,
            include_subgroups=self.include_subgroups,
            max_queryable_pages=self.max_queryable_pages,
            membership=self.membership,
            min_access_level=self.min_access_level,
            order_by=self.order_by,
            owned=self.owned,
            page=self.page,
            per_page=self.per_page,
            projects=self.projects,
            search=self.search,
            simple=self.simple,
            sort=self.sort,
            starred=self.starred,
            statistics=self.statistics,
            visibility=self.visibility,
            with_custom_attributes=self.with_custom_attributes,
            with_issues_enabled=self.with_issues_enabled,
            with_merge_requests_enabled=self.with_merge_requests_enabled,
            with_programming_language=self.with_programming_language,
            with_shared=self.with_shared)


def get_projects(archived: Optional[bool] = None,
                 group_id: Optional[float] = None,
                 include_subgroups: Optional[bool] = None,
                 max_queryable_pages: Optional[float] = None,
                 membership: Optional[bool] = None,
                 min_access_level: Optional[float] = None,
                 order_by: Optional[str] = None,
                 owned: Optional[bool] = None,
                 page: Optional[float] = None,
                 per_page: Optional[float] = None,
                 search: Optional[str] = None,
                 simple: Optional[bool] = None,
                 sort: Optional[str] = None,
                 starred: Optional[bool] = None,
                 statistics: Optional[bool] = None,
                 visibility: Optional[str] = None,
                 with_custom_attributes: Optional[bool] = None,
                 with_issues_enabled: Optional[bool] = None,
                 with_merge_requests_enabled: Optional[bool] = None,
                 with_programming_language: Optional[str] = None,
                 with_shared: Optional[bool] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectsResult:
    """
    Provides details about a list of projects in the Gitlab provider. Listing all projects and group projects with [project filtering](https://docs.gitlab.com/ee/api/projects.html#list-user-projects) or [group project filtering](https://docs.gitlab.com/ee/api/groups.html#list-a-groups-projects) is supported.

    > NOTE: This data source supports all available filters exposed by the `xanzy/go-gitlab` package, which might not expose all available filters exposed by the Gitlab APIs.

    ## Example Usage
    ### List projects within a group tree

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    mygroup = gitlab.get_group(full_path="mygroup")
    group_projects = gitlab.get_projects(group_id=mygroup.id,
        order_by="name",
        include_subgroups=True,
        with_shared=False)
    ```
    ### List projects using the search syntax

    ```python
    import pulumi
    import pulumi_gitlab as gitlab

    projects = gitlab.get_projects(search="postgresql",
        visibility="private")
    ```


    :param bool archived: Limit by archived status.
    :param float group_id: The ID of the group owned by the authenticated user to look projects for within. Cannot be used with `min_access_level`, `with_programming_language` or `statistics`.
    :param bool include_subgroups: Include projects in subgroups of this group. Default is `false`. Needs `group_id`.
    :param float max_queryable_pages: Prevents overloading your Gitlab instance in case of a misconfiguration. Default is `10`.
    :param bool membership: Limit by projects that the current user is a member of.
    :param float min_access_level: Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/ee/api/members.html) for values. Cannot be used with `group_id`.
    :param str order_by: Return projects ordered by `id`, `name`, `path`, `created_at`, `updated_at`, or `last_activity_at` fields. Default is `created_at`.
    :param bool owned: Limit by projects owned by the current user.
    :param str search: Return list of authorized projects matching the search criteria.
    :param bool simple: Return only the ID, URL, name, and path of each project.
    :param str sort: Return projects sorted in `asc` or `desc` order. Default is `desc`.
    :param bool starred: Limit by projects starred by the current user.
    :param bool statistics: Include project statistics. Cannot be used with `group_id`.
    :param str visibility: Limit by visibility `public`, `internal`, or `private`.
    :param bool with_custom_attributes: Include custom attributes in response _(admins only)_.
    :param bool with_issues_enabled: Limit by projects with issues feature enabled. Default is `false`.
    :param bool with_merge_requests_enabled: Limit by projects with merge requests feature enabled. Default is `false`.
    :param str with_programming_language: Limit by projects which use the given programming language. Cannot be used with `group_id`.
    :param bool with_shared: Include projects shared to this group. Default is `true`. Needs `group_id`.
    """
    __args__ = dict()
    __args__['archived'] = archived
    __args__['groupId'] = group_id
    __args__['includeSubgroups'] = include_subgroups
    __args__['maxQueryablePages'] = max_queryable_pages
    __args__['membership'] = membership
    __args__['minAccessLevel'] = min_access_level
    __args__['orderBy'] = order_by
    __args__['owned'] = owned
    __args__['page'] = page
    __args__['perPage'] = per_page
    __args__['search'] = search
    __args__['simple'] = simple
    __args__['sort'] = sort
    __args__['starred'] = starred
    __args__['statistics'] = statistics
    __args__['visibility'] = visibility
    __args__['withCustomAttributes'] = with_custom_attributes
    __args__['withIssuesEnabled'] = with_issues_enabled
    __args__['withMergeRequestsEnabled'] = with_merge_requests_enabled
    __args__['withProgrammingLanguage'] = with_programming_language
    __args__['withShared'] = with_shared
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gitlab:index/getProjects:getProjects', __args__, opts=opts, typ=GetProjectsResult).value

    return AwaitableGetProjectsResult(
        archived=__ret__.archived,
        group_id=__ret__.group_id,
        id=__ret__.id,
        include_subgroups=__ret__.include_subgroups,
        max_queryable_pages=__ret__.max_queryable_pages,
        membership=__ret__.membership,
        min_access_level=__ret__.min_access_level,
        order_by=__ret__.order_by,
        owned=__ret__.owned,
        page=__ret__.page,
        per_page=__ret__.per_page,
        projects=__ret__.projects,
        search=__ret__.search,
        simple=__ret__.simple,
        sort=__ret__.sort,
        starred=__ret__.starred,
        statistics=__ret__.statistics,
        visibility=__ret__.visibility,
        with_custom_attributes=__ret__.with_custom_attributes,
        with_issues_enabled=__ret__.with_issues_enabled,
        with_merge_requests_enabled=__ret__.with_merge_requests_enabled,
        with_programming_language=__ret__.with_programming_language,
        with_shared=__ret__.with_shared)
