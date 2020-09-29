# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables

__all__ = [
    'ProjectPushRulesArgs',
]

@pulumi.input_type
class ProjectPushRulesArgs:
    def __init__(__self__, *,
                 author_email_regex: Optional[pulumi.Input[str]] = None,
                 branch_name_regex: Optional[pulumi.Input[str]] = None,
                 commit_committer_check: Optional[pulumi.Input[bool]] = None,
                 commit_message_negative_regex: Optional[pulumi.Input[str]] = None,
                 commit_message_regex: Optional[pulumi.Input[str]] = None,
                 deny_delete_tag: Optional[pulumi.Input[bool]] = None,
                 file_name_regex: Optional[pulumi.Input[str]] = None,
                 max_file_size: Optional[pulumi.Input[float]] = None,
                 member_check: Optional[pulumi.Input[bool]] = None,
                 prevent_secrets: Optional[pulumi.Input[bool]] = None,
                 reject_unsigned_commits: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[str] author_email_regex: All commit author emails must match this regex, e.g. `@my-company.com$`.
        :param pulumi.Input[str] branch_name_regex: All branch names must match this regex, e.g. `(feature|hotfix)\/*`.
        :param pulumi.Input[bool] commit_committer_check: Users can only push commits to this repository that were committed with one of their own verified emails.
        :param pulumi.Input[str] commit_message_negative_regex: No commit message is allowed to match this regex, for example `ssh\:\/\/`.
        :param pulumi.Input[str] commit_message_regex: All commit messages must match this regex, e.g. `Fixed \d+\..*`.
        :param pulumi.Input[bool] deny_delete_tag: Deny deleting a tag.
        :param pulumi.Input[str] file_name_regex: All commited filenames must not match this regex, e.g. `(jar|exe)$`.
        :param pulumi.Input[float] max_file_size: Maximum file size (MB).
        :param pulumi.Input[bool] member_check: Restrict commits by author (email) to existing GitLab users.
        :param pulumi.Input[bool] prevent_secrets: GitLab will reject any files that are likely to contain secrets.
        :param pulumi.Input[bool] reject_unsigned_commits: Reject commit when it’s not signed through GPG.
        """
        if author_email_regex is not None:
            pulumi.set(__self__, "author_email_regex", author_email_regex)
        if branch_name_regex is not None:
            pulumi.set(__self__, "branch_name_regex", branch_name_regex)
        if commit_committer_check is not None:
            pulumi.set(__self__, "commit_committer_check", commit_committer_check)
        if commit_message_negative_regex is not None:
            pulumi.set(__self__, "commit_message_negative_regex", commit_message_negative_regex)
        if commit_message_regex is not None:
            pulumi.set(__self__, "commit_message_regex", commit_message_regex)
        if deny_delete_tag is not None:
            pulumi.set(__self__, "deny_delete_tag", deny_delete_tag)
        if file_name_regex is not None:
            pulumi.set(__self__, "file_name_regex", file_name_regex)
        if max_file_size is not None:
            pulumi.set(__self__, "max_file_size", max_file_size)
        if member_check is not None:
            pulumi.set(__self__, "member_check", member_check)
        if prevent_secrets is not None:
            pulumi.set(__self__, "prevent_secrets", prevent_secrets)
        if reject_unsigned_commits is not None:
            pulumi.set(__self__, "reject_unsigned_commits", reject_unsigned_commits)

    @property
    @pulumi.getter(name="authorEmailRegex")
    def author_email_regex(self) -> Optional[pulumi.Input[str]]:
        """
        All commit author emails must match this regex, e.g. `@my-company.com$`.
        """
        return pulumi.get(self, "author_email_regex")

    @author_email_regex.setter
    def author_email_regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "author_email_regex", value)

    @property
    @pulumi.getter(name="branchNameRegex")
    def branch_name_regex(self) -> Optional[pulumi.Input[str]]:
        """
        All branch names must match this regex, e.g. `(feature|hotfix)\/*`.
        """
        return pulumi.get(self, "branch_name_regex")

    @branch_name_regex.setter
    def branch_name_regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "branch_name_regex", value)

    @property
    @pulumi.getter(name="commitCommitterCheck")
    def commit_committer_check(self) -> Optional[pulumi.Input[bool]]:
        """
        Users can only push commits to this repository that were committed with one of their own verified emails.
        """
        return pulumi.get(self, "commit_committer_check")

    @commit_committer_check.setter
    def commit_committer_check(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "commit_committer_check", value)

    @property
    @pulumi.getter(name="commitMessageNegativeRegex")
    def commit_message_negative_regex(self) -> Optional[pulumi.Input[str]]:
        """
        No commit message is allowed to match this regex, for example `ssh\:\/\/`.
        """
        return pulumi.get(self, "commit_message_negative_regex")

    @commit_message_negative_regex.setter
    def commit_message_negative_regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "commit_message_negative_regex", value)

    @property
    @pulumi.getter(name="commitMessageRegex")
    def commit_message_regex(self) -> Optional[pulumi.Input[str]]:
        """
        All commit messages must match this regex, e.g. `Fixed \d+\..*`.
        """
        return pulumi.get(self, "commit_message_regex")

    @commit_message_regex.setter
    def commit_message_regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "commit_message_regex", value)

    @property
    @pulumi.getter(name="denyDeleteTag")
    def deny_delete_tag(self) -> Optional[pulumi.Input[bool]]:
        """
        Deny deleting a tag.
        """
        return pulumi.get(self, "deny_delete_tag")

    @deny_delete_tag.setter
    def deny_delete_tag(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "deny_delete_tag", value)

    @property
    @pulumi.getter(name="fileNameRegex")
    def file_name_regex(self) -> Optional[pulumi.Input[str]]:
        """
        All commited filenames must not match this regex, e.g. `(jar|exe)$`.
        """
        return pulumi.get(self, "file_name_regex")

    @file_name_regex.setter
    def file_name_regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_name_regex", value)

    @property
    @pulumi.getter(name="maxFileSize")
    def max_file_size(self) -> Optional[pulumi.Input[float]]:
        """
        Maximum file size (MB).
        """
        return pulumi.get(self, "max_file_size")

    @max_file_size.setter
    def max_file_size(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "max_file_size", value)

    @property
    @pulumi.getter(name="memberCheck")
    def member_check(self) -> Optional[pulumi.Input[bool]]:
        """
        Restrict commits by author (email) to existing GitLab users.
        """
        return pulumi.get(self, "member_check")

    @member_check.setter
    def member_check(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "member_check", value)

    @property
    @pulumi.getter(name="preventSecrets")
    def prevent_secrets(self) -> Optional[pulumi.Input[bool]]:
        """
        GitLab will reject any files that are likely to contain secrets.
        """
        return pulumi.get(self, "prevent_secrets")

    @prevent_secrets.setter
    def prevent_secrets(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "prevent_secrets", value)

    @property
    @pulumi.getter(name="rejectUnsignedCommits")
    def reject_unsigned_commits(self) -> Optional[pulumi.Input[bool]]:
        """
        Reject commit when it’s not signed through GPG.
        """
        return pulumi.get(self, "reject_unsigned_commits")

    @reject_unsigned_commits.setter
    def reject_unsigned_commits(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "reject_unsigned_commits", value)


