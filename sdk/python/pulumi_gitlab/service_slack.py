# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class ServiceSlack(pulumi.CustomResource):
    confidential_issue_channel: pulumi.Output[str]
    """
    The name of the channel to receive confidential issue events notifications.
    """
    confidential_issues_events: pulumi.Output[bool]
    """
    Enable notifications for confidential issues events.
    """
    confidential_note_events: pulumi.Output[bool]
    """
    Enable notifications for confidential note events.
    """
    issue_channel: pulumi.Output[str]
    """
    The name of the channel to receive issue events notifications.
    """
    issues_events: pulumi.Output[bool]
    """
    Enable notifications for issues events.
    """
    job_events: pulumi.Output[bool]
    merge_request_channel: pulumi.Output[str]
    """
    The name of the channel to receive merge request events notifications.
    """
    merge_requests_events: pulumi.Output[bool]
    """
    Enable notifications for merge requests events.
    """
    note_channel: pulumi.Output[str]
    """
    The name of the channel to receive note events notifications.
    """
    note_events: pulumi.Output[bool]
    """
    Enable notifications for note events.
    """
    notify_only_broken_pipelines: pulumi.Output[bool]
    """
    Send notifications for broken pipelines.
    """
    notify_only_default_branch: pulumi.Output[bool]
    """
    Send notifications only for the default branch.
    """
    pipeline_channel: pulumi.Output[str]
    """
    The name of the channel to receive pipeline events notifications.
    """
    pipeline_events: pulumi.Output[bool]
    """
    Enable notifications for pipeline events.
    """
    project: pulumi.Output[str]
    """
    ID of the project you want to activate integration on.
    """
    push_channel: pulumi.Output[str]
    """
    The name of the channel to receive push events notifications.
    """
    push_events: pulumi.Output[bool]
    """
    Enable notifications for push events.
    """
    tag_push_channel: pulumi.Output[str]
    """
    The name of the channel to receive tag push events notifications.
    """
    tag_push_events: pulumi.Output[bool]
    """
    Enable notifications for tag push events.
    """
    username: pulumi.Output[str]
    """
    Username to use.
    """
    webhook: pulumi.Output[str]
    """
    Webhook URL (ex.: https://hooks.slack.com/services/...)
    """
    wiki_page_channel: pulumi.Output[str]
    """
    The name of the channel to receive wiki page events notifications.
    """
    wiki_page_events: pulumi.Output[bool]
    """
    Enable notifications for wiki page events.
    """
    def __init__(__self__, resource_name, opts=None, confidential_issue_channel=None, confidential_issues_events=None, confidential_note_events=None, issue_channel=None, issues_events=None, merge_request_channel=None, merge_requests_events=None, note_channel=None, note_events=None, notify_only_broken_pipelines=None, notify_only_default_branch=None, pipeline_channel=None, pipeline_events=None, project=None, push_channel=None, push_events=None, tag_push_channel=None, tag_push_events=None, username=None, webhook=None, wiki_page_channel=None, wiki_page_events=None, __props__=None, __name__=None, __opts__=None):
        """
        This resource allows you to manage Slack notifications integration.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] confidential_issue_channel: The name of the channel to receive confidential issue events notifications.
        :param pulumi.Input[bool] confidential_issues_events: Enable notifications for confidential issues events.
        :param pulumi.Input[bool] confidential_note_events: Enable notifications for confidential note events.
        :param pulumi.Input[str] issue_channel: The name of the channel to receive issue events notifications.
        :param pulumi.Input[bool] issues_events: Enable notifications for issues events.
        :param pulumi.Input[str] merge_request_channel: The name of the channel to receive merge request events notifications.
        :param pulumi.Input[bool] merge_requests_events: Enable notifications for merge requests events.
        :param pulumi.Input[str] note_channel: The name of the channel to receive note events notifications.
        :param pulumi.Input[bool] note_events: Enable notifications for note events.
        :param pulumi.Input[bool] notify_only_broken_pipelines: Send notifications for broken pipelines.
        :param pulumi.Input[bool] notify_only_default_branch: Send notifications only for the default branch.
        :param pulumi.Input[str] pipeline_channel: The name of the channel to receive pipeline events notifications.
        :param pulumi.Input[bool] pipeline_events: Enable notifications for pipeline events.
        :param pulumi.Input[str] project: ID of the project you want to activate integration on.
        :param pulumi.Input[str] push_channel: The name of the channel to receive push events notifications.
        :param pulumi.Input[bool] push_events: Enable notifications for push events.
        :param pulumi.Input[str] tag_push_channel: The name of the channel to receive tag push events notifications.
        :param pulumi.Input[bool] tag_push_events: Enable notifications for tag push events.
        :param pulumi.Input[str] username: Username to use.
        :param pulumi.Input[str] webhook: Webhook URL (ex.: https://hooks.slack.com/services/...)
        :param pulumi.Input[str] wiki_page_channel: The name of the channel to receive wiki page events notifications.
        :param pulumi.Input[bool] wiki_page_events: Enable notifications for wiki page events.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/service_slack.html.markdown.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['confidential_issue_channel'] = confidential_issue_channel
            __props__['confidential_issues_events'] = confidential_issues_events
            __props__['confidential_note_events'] = confidential_note_events
            __props__['issue_channel'] = issue_channel
            __props__['issues_events'] = issues_events
            __props__['merge_request_channel'] = merge_request_channel
            __props__['merge_requests_events'] = merge_requests_events
            __props__['note_channel'] = note_channel
            __props__['note_events'] = note_events
            __props__['notify_only_broken_pipelines'] = notify_only_broken_pipelines
            __props__['notify_only_default_branch'] = notify_only_default_branch
            __props__['pipeline_channel'] = pipeline_channel
            __props__['pipeline_events'] = pipeline_events
            if project is None:
                raise TypeError("Missing required property 'project'")
            __props__['project'] = project
            __props__['push_channel'] = push_channel
            __props__['push_events'] = push_events
            __props__['tag_push_channel'] = tag_push_channel
            __props__['tag_push_events'] = tag_push_events
            __props__['username'] = username
            if webhook is None:
                raise TypeError("Missing required property 'webhook'")
            __props__['webhook'] = webhook
            __props__['wiki_page_channel'] = wiki_page_channel
            __props__['wiki_page_events'] = wiki_page_events
            __props__['job_events'] = None
        super(ServiceSlack, __self__).__init__(
            'gitlab:index/serviceSlack:ServiceSlack',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, confidential_issue_channel=None, confidential_issues_events=None, confidential_note_events=None, issue_channel=None, issues_events=None, job_events=None, merge_request_channel=None, merge_requests_events=None, note_channel=None, note_events=None, notify_only_broken_pipelines=None, notify_only_default_branch=None, pipeline_channel=None, pipeline_events=None, project=None, push_channel=None, push_events=None, tag_push_channel=None, tag_push_events=None, username=None, webhook=None, wiki_page_channel=None, wiki_page_events=None):
        """
        Get an existing ServiceSlack resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] confidential_issue_channel: The name of the channel to receive confidential issue events notifications.
        :param pulumi.Input[bool] confidential_issues_events: Enable notifications for confidential issues events.
        :param pulumi.Input[bool] confidential_note_events: Enable notifications for confidential note events.
        :param pulumi.Input[str] issue_channel: The name of the channel to receive issue events notifications.
        :param pulumi.Input[bool] issues_events: Enable notifications for issues events.
        :param pulumi.Input[str] merge_request_channel: The name of the channel to receive merge request events notifications.
        :param pulumi.Input[bool] merge_requests_events: Enable notifications for merge requests events.
        :param pulumi.Input[str] note_channel: The name of the channel to receive note events notifications.
        :param pulumi.Input[bool] note_events: Enable notifications for note events.
        :param pulumi.Input[bool] notify_only_broken_pipelines: Send notifications for broken pipelines.
        :param pulumi.Input[bool] notify_only_default_branch: Send notifications only for the default branch.
        :param pulumi.Input[str] pipeline_channel: The name of the channel to receive pipeline events notifications.
        :param pulumi.Input[bool] pipeline_events: Enable notifications for pipeline events.
        :param pulumi.Input[str] project: ID of the project you want to activate integration on.
        :param pulumi.Input[str] push_channel: The name of the channel to receive push events notifications.
        :param pulumi.Input[bool] push_events: Enable notifications for push events.
        :param pulumi.Input[str] tag_push_channel: The name of the channel to receive tag push events notifications.
        :param pulumi.Input[bool] tag_push_events: Enable notifications for tag push events.
        :param pulumi.Input[str] username: Username to use.
        :param pulumi.Input[str] webhook: Webhook URL (ex.: https://hooks.slack.com/services/...)
        :param pulumi.Input[str] wiki_page_channel: The name of the channel to receive wiki page events notifications.
        :param pulumi.Input[bool] wiki_page_events: Enable notifications for wiki page events.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/service_slack.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["confidential_issue_channel"] = confidential_issue_channel
        __props__["confidential_issues_events"] = confidential_issues_events
        __props__["confidential_note_events"] = confidential_note_events
        __props__["issue_channel"] = issue_channel
        __props__["issues_events"] = issues_events
        __props__["job_events"] = job_events
        __props__["merge_request_channel"] = merge_request_channel
        __props__["merge_requests_events"] = merge_requests_events
        __props__["note_channel"] = note_channel
        __props__["note_events"] = note_events
        __props__["notify_only_broken_pipelines"] = notify_only_broken_pipelines
        __props__["notify_only_default_branch"] = notify_only_default_branch
        __props__["pipeline_channel"] = pipeline_channel
        __props__["pipeline_events"] = pipeline_events
        __props__["project"] = project
        __props__["push_channel"] = push_channel
        __props__["push_events"] = push_events
        __props__["tag_push_channel"] = tag_push_channel
        __props__["tag_push_events"] = tag_push_events
        __props__["username"] = username
        __props__["webhook"] = webhook
        __props__["wiki_page_channel"] = wiki_page_channel
        __props__["wiki_page_events"] = wiki_page_events
        return ServiceSlack(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

