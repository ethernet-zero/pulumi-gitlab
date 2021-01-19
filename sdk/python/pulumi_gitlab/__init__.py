# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .branch_protection import *
from .deploy_key import *
from .deploy_key_enable import *
from .deploy_token import *
from .get_group import *
from .get_group_membership import *
from .get_project import *
from .get_projects import *
from .get_user import *
from .get_users import *
from .group import *
from .group_cluster import *
from .group_label import *
from .group_ldap_link import *
from .group_membership import *
from .group_share_group import *
from .group_variable import *
from .instance_cluster import *
from .instance_variable import *
from .label import *
from .pipeline_schedule import *
from .pipeline_schedule_variable import *
from .pipeline_trigger import *
from .project import *
from .project_approval_rule import *
from .project_cluster import *
from .project_hook import *
from .project_level_mr_approvals import *
from .project_membership import *
from .project_mirror import *
from .project_share_group import *
from .project_variable import *
from .provider import *
from .service_github import *
from .service_jira import *
from .service_pipelines_email import *
from .service_slack import *
from .tag_protection import *
from .user import *
from ._inputs import *
from . import outputs

# Make subpackages available:
from . import (
    config,
)

def _register_module():
    import pulumi
    from . import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "gitlab:index/branchProtection:BranchProtection":
                return BranchProtection(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/deployKey:DeployKey":
                return DeployKey(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/deployKeyEnable:DeployKeyEnable":
                return DeployKeyEnable(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/deployToken:DeployToken":
                return DeployToken(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/group:Group":
                return Group(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/groupCluster:GroupCluster":
                return GroupCluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/groupLabel:GroupLabel":
                return GroupLabel(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/groupLdapLink:GroupLdapLink":
                return GroupLdapLink(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/groupMembership:GroupMembership":
                return GroupMembership(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/groupShareGroup:GroupShareGroup":
                return GroupShareGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/groupVariable:GroupVariable":
                return GroupVariable(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/instanceCluster:InstanceCluster":
                return InstanceCluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/instanceVariable:InstanceVariable":
                return InstanceVariable(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/label:Label":
                return Label(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/pipelineSchedule:PipelineSchedule":
                return PipelineSchedule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable":
                return PipelineScheduleVariable(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/pipelineTrigger:PipelineTrigger":
                return PipelineTrigger(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/project:Project":
                return Project(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/projectApprovalRule:ProjectApprovalRule":
                return ProjectApprovalRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/projectCluster:ProjectCluster":
                return ProjectCluster(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/projectHook:ProjectHook":
                return ProjectHook(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals":
                return ProjectLevelMrApprovals(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/projectMembership:ProjectMembership":
                return ProjectMembership(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/projectMirror:ProjectMirror":
                return ProjectMirror(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/projectShareGroup:ProjectShareGroup":
                return ProjectShareGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/projectVariable:ProjectVariable":
                return ProjectVariable(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/serviceGithub:ServiceGithub":
                return ServiceGithub(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/serviceJira:ServiceJira":
                return ServiceJira(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/servicePipelinesEmail:ServicePipelinesEmail":
                return ServicePipelinesEmail(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/serviceSlack:ServiceSlack":
                return ServiceSlack(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/tagProtection:TagProtection":
                return TagProtection(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "gitlab:index/user:User":
                return User(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("gitlab", "index/branchProtection", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/deployKey", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/deployKeyEnable", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/deployToken", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/group", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/groupCluster", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/groupLabel", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/groupLdapLink", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/groupMembership", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/groupShareGroup", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/groupVariable", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/instanceCluster", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/instanceVariable", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/label", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/pipelineSchedule", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/pipelineScheduleVariable", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/pipelineTrigger", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/project", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/projectApprovalRule", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/projectCluster", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/projectHook", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/projectLevelMrApprovals", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/projectMembership", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/projectMirror", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/projectShareGroup", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/projectVariable", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/serviceGithub", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/serviceJira", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/servicePipelinesEmail", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/serviceSlack", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/tagProtection", _module_instance)
    pulumi.runtime.register_resource_module("gitlab", "index/user", _module_instance)


    class Package(pulumi.runtime.ResourcePackage):
        _version = _utilities.get_semver_version()

        def version(self):
            return Package._version

        def construct_provider(self, name: str, typ: str, urn: str) -> pulumi.ProviderResource:
            if typ != "pulumi:providers:gitlab":
                raise Exception(f"unknown provider type {typ}")
            return Provider(name, pulumi.ResourceOptions(urn=urn))


    pulumi.runtime.register_resource_package("gitlab", Package())

_register_module()
