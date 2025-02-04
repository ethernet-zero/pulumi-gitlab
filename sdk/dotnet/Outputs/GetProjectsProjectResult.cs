// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Outputs
{

    [OutputType]
    public sealed class GetProjectsProjectResult
    {
        public readonly ImmutableDictionary<string, string> _links;
        public readonly bool AllowMergeOnSkippedPipeline;
        public readonly string AnalyticsAccessLevel;
        public readonly int ApprovalsBeforeMerge;
        /// <summary>
        /// Limit by archived status.
        /// </summary>
        public readonly bool Archived;
        public readonly string AutoCancelPendingPipelines;
        public readonly string AutoDevopsDeployStrategy;
        public readonly bool AutoDevopsEnabled;
        public readonly bool AutocloseReferencedIssues;
        public readonly string AvatarUrl;
        public readonly string BuildCoverageRegex;
        public readonly string BuildGitStrategy;
        public readonly int BuildTimeout;
        public readonly string BuildsAccessLevel;
        public readonly string CiConfigPath;
        public readonly int CiDefaultGitDepth;
        public readonly bool CiForwardDeploymentEnabled;
        public readonly ImmutableArray<Outputs.GetProjectsProjectContainerExpirationPolicyResult> ContainerExpirationPolicies;
        public readonly string ContainerRegistryAccessLevel;
        public readonly bool ContainerRegistryEnabled;
        public readonly string CreatedAt;
        public readonly int CreatorId;
        public readonly ImmutableArray<ImmutableDictionary<string, string>> CustomAttributes;
        public readonly string DefaultBranch;
        public readonly string Description;
        public readonly bool EmailsDisabled;
        public readonly string ExternalAuthorizationClassificationLabel;
        public readonly Outputs.GetProjectsProjectForkedFromProjectResult ForkedFromProject;
        public readonly string ForkingAccessLevel;
        public readonly int ForksCount;
        public readonly string HttpUrlToRepo;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly int Id;
        public readonly string ImportError;
        public readonly string ImportStatus;
        public readonly string IssuesAccessLevel;
        public readonly bool IssuesEnabled;
        public readonly bool JobsEnabled;
        public readonly string LastActivityAt;
        public readonly bool LfsEnabled;
        public readonly string MergeCommitTemplate;
        public readonly string MergeMethod;
        public readonly bool MergePipelinesEnabled;
        public readonly string MergeRequestsAccessLevel;
        public readonly bool MergeRequestsEnabled;
        public readonly bool MergeTrainsEnabled;
        public readonly bool Mirror;
        public readonly bool MirrorOverwritesDivergedBranches;
        public readonly bool MirrorTriggerBuilds;
        public readonly int MirrorUserId;
        public readonly string Name;
        public readonly string NameWithNamespace;
        public readonly Outputs.GetProjectsProjectNamespaceResult Namespace;
        public readonly bool OnlyAllowMergeIfAllDiscussionsAreResolved;
        public readonly bool OnlyAllowMergeIfPipelineSucceeds;
        public readonly bool OnlyMirrorProtectedBranches;
        public readonly int OpenIssuesCount;
        public readonly string OperationsAccessLevel;
        public readonly Outputs.GetProjectsProjectOwnerResult Owner;
        public readonly bool PackagesEnabled;
        public readonly string Path;
        public readonly string PathWithNamespace;
        public readonly Outputs.GetProjectsProjectPermissionsResult Permissions;
        public readonly bool Public;
        public readonly bool PublicBuilds;
        public readonly string ReadmeUrl;
        public readonly string RepositoryAccessLevel;
        public readonly string RepositoryStorage;
        public readonly bool RequestAccessEnabled;
        public readonly string RequirementsAccessLevel;
        public readonly bool ResolveOutdatedDiffDiscussions;
        public readonly string RunnersToken;
        public readonly string SecurityAndComplianceAccessLevel;
        public readonly bool SharedRunnersEnabled;
        public readonly ImmutableArray<Outputs.GetProjectsProjectSharedWithGroupResult> SharedWithGroups;
        public readonly string SnippetsAccessLevel;
        public readonly bool SnippetsEnabled;
        public readonly string SquashCommitTemplate;
        public readonly string SshUrlToRepo;
        public readonly int StarCount;
        /// <summary>
        /// Include project statistics. Cannot be used with `group_id`.
        /// </summary>
        public readonly ImmutableDictionary<string, int> Statistics;
        public readonly ImmutableArray<string> TagLists;
        public readonly ImmutableArray<string> Topics;
        /// <summary>
        /// Limit by visibility `public`, `internal`, or `private`.
        /// </summary>
        public readonly string Visibility;
        public readonly string WebUrl;
        public readonly string WikiAccessLevel;
        public readonly bool WikiEnabled;

        [OutputConstructor]
        private GetProjectsProjectResult(
            ImmutableDictionary<string, string> _links,

            bool allowMergeOnSkippedPipeline,

            string analyticsAccessLevel,

            int approvalsBeforeMerge,

            bool archived,

            string autoCancelPendingPipelines,

            string autoDevopsDeployStrategy,

            bool autoDevopsEnabled,

            bool autocloseReferencedIssues,

            string avatarUrl,

            string buildCoverageRegex,

            string buildGitStrategy,

            int buildTimeout,

            string buildsAccessLevel,

            string ciConfigPath,

            int ciDefaultGitDepth,

            bool ciForwardDeploymentEnabled,

            ImmutableArray<Outputs.GetProjectsProjectContainerExpirationPolicyResult> containerExpirationPolicies,

            string containerRegistryAccessLevel,

            bool containerRegistryEnabled,

            string createdAt,

            int creatorId,

            ImmutableArray<ImmutableDictionary<string, string>> customAttributes,

            string defaultBranch,

            string description,

            bool emailsDisabled,

            string externalAuthorizationClassificationLabel,

            Outputs.GetProjectsProjectForkedFromProjectResult forkedFromProject,

            string forkingAccessLevel,

            int forksCount,

            string httpUrlToRepo,

            int id,

            string importError,

            string importStatus,

            string issuesAccessLevel,

            bool issuesEnabled,

            bool jobsEnabled,

            string lastActivityAt,

            bool lfsEnabled,

            string mergeCommitTemplate,

            string mergeMethod,

            bool mergePipelinesEnabled,

            string mergeRequestsAccessLevel,

            bool mergeRequestsEnabled,

            bool mergeTrainsEnabled,

            bool mirror,

            bool mirrorOverwritesDivergedBranches,

            bool mirrorTriggerBuilds,

            int mirrorUserId,

            string name,

            string nameWithNamespace,

            Outputs.GetProjectsProjectNamespaceResult @namespace,

            bool onlyAllowMergeIfAllDiscussionsAreResolved,

            bool onlyAllowMergeIfPipelineSucceeds,

            bool onlyMirrorProtectedBranches,

            int openIssuesCount,

            string operationsAccessLevel,

            Outputs.GetProjectsProjectOwnerResult owner,

            bool packagesEnabled,

            string path,

            string pathWithNamespace,

            Outputs.GetProjectsProjectPermissionsResult permissions,

            bool @public,

            bool publicBuilds,

            string readmeUrl,

            string repositoryAccessLevel,

            string repositoryStorage,

            bool requestAccessEnabled,

            string requirementsAccessLevel,

            bool resolveOutdatedDiffDiscussions,

            string runnersToken,

            string securityAndComplianceAccessLevel,

            bool sharedRunnersEnabled,

            ImmutableArray<Outputs.GetProjectsProjectSharedWithGroupResult> sharedWithGroups,

            string snippetsAccessLevel,

            bool snippetsEnabled,

            string squashCommitTemplate,

            string sshUrlToRepo,

            int starCount,

            ImmutableDictionary<string, int> statistics,

            ImmutableArray<string> tagLists,

            ImmutableArray<string> topics,

            string visibility,

            string webUrl,

            string wikiAccessLevel,

            bool wikiEnabled)
        {
            this._links = _links;
            AllowMergeOnSkippedPipeline = allowMergeOnSkippedPipeline;
            AnalyticsAccessLevel = analyticsAccessLevel;
            ApprovalsBeforeMerge = approvalsBeforeMerge;
            Archived = archived;
            AutoCancelPendingPipelines = autoCancelPendingPipelines;
            AutoDevopsDeployStrategy = autoDevopsDeployStrategy;
            AutoDevopsEnabled = autoDevopsEnabled;
            AutocloseReferencedIssues = autocloseReferencedIssues;
            AvatarUrl = avatarUrl;
            BuildCoverageRegex = buildCoverageRegex;
            BuildGitStrategy = buildGitStrategy;
            BuildTimeout = buildTimeout;
            BuildsAccessLevel = buildsAccessLevel;
            CiConfigPath = ciConfigPath;
            CiDefaultGitDepth = ciDefaultGitDepth;
            CiForwardDeploymentEnabled = ciForwardDeploymentEnabled;
            ContainerExpirationPolicies = containerExpirationPolicies;
            ContainerRegistryAccessLevel = containerRegistryAccessLevel;
            ContainerRegistryEnabled = containerRegistryEnabled;
            CreatedAt = createdAt;
            CreatorId = creatorId;
            CustomAttributes = customAttributes;
            DefaultBranch = defaultBranch;
            Description = description;
            EmailsDisabled = emailsDisabled;
            ExternalAuthorizationClassificationLabel = externalAuthorizationClassificationLabel;
            ForkedFromProject = forkedFromProject;
            ForkingAccessLevel = forkingAccessLevel;
            ForksCount = forksCount;
            HttpUrlToRepo = httpUrlToRepo;
            Id = id;
            ImportError = importError;
            ImportStatus = importStatus;
            IssuesAccessLevel = issuesAccessLevel;
            IssuesEnabled = issuesEnabled;
            JobsEnabled = jobsEnabled;
            LastActivityAt = lastActivityAt;
            LfsEnabled = lfsEnabled;
            MergeCommitTemplate = mergeCommitTemplate;
            MergeMethod = mergeMethod;
            MergePipelinesEnabled = mergePipelinesEnabled;
            MergeRequestsAccessLevel = mergeRequestsAccessLevel;
            MergeRequestsEnabled = mergeRequestsEnabled;
            MergeTrainsEnabled = mergeTrainsEnabled;
            Mirror = mirror;
            MirrorOverwritesDivergedBranches = mirrorOverwritesDivergedBranches;
            MirrorTriggerBuilds = mirrorTriggerBuilds;
            MirrorUserId = mirrorUserId;
            Name = name;
            NameWithNamespace = nameWithNamespace;
            Namespace = @namespace;
            OnlyAllowMergeIfAllDiscussionsAreResolved = onlyAllowMergeIfAllDiscussionsAreResolved;
            OnlyAllowMergeIfPipelineSucceeds = onlyAllowMergeIfPipelineSucceeds;
            OnlyMirrorProtectedBranches = onlyMirrorProtectedBranches;
            OpenIssuesCount = openIssuesCount;
            OperationsAccessLevel = operationsAccessLevel;
            Owner = owner;
            PackagesEnabled = packagesEnabled;
            Path = path;
            PathWithNamespace = pathWithNamespace;
            Permissions = permissions;
            Public = @public;
            PublicBuilds = publicBuilds;
            ReadmeUrl = readmeUrl;
            RepositoryAccessLevel = repositoryAccessLevel;
            RepositoryStorage = repositoryStorage;
            RequestAccessEnabled = requestAccessEnabled;
            RequirementsAccessLevel = requirementsAccessLevel;
            ResolveOutdatedDiffDiscussions = resolveOutdatedDiffDiscussions;
            RunnersToken = runnersToken;
            SecurityAndComplianceAccessLevel = securityAndComplianceAccessLevel;
            SharedRunnersEnabled = sharedRunnersEnabled;
            SharedWithGroups = sharedWithGroups;
            SnippetsAccessLevel = snippetsAccessLevel;
            SnippetsEnabled = snippetsEnabled;
            SquashCommitTemplate = squashCommitTemplate;
            SshUrlToRepo = sshUrlToRepo;
            StarCount = starCount;
            Statistics = statistics;
            TagLists = tagLists;
            Topics = topics;
            Visibility = visibility;
            WebUrl = webUrl;
            WikiAccessLevel = wikiAccessLevel;
            WikiEnabled = wikiEnabled;
        }
    }
}
