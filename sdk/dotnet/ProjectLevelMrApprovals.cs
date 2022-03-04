// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    /// <summary>
    /// This resource allows you to configure project-level MR approvals. for your GitLab projects.
    /// For further information on merge request approvals, consult the [GitLab API documentation](https://docs.gitlab.com/ee/api/merge_request_approvals.html#project-level-mr-approvals).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using GitLab = Pulumi.GitLab;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var fooProject = new GitLab.Project("fooProject", new GitLab.ProjectArgs
    ///         {
    ///             Description = "My example project",
    ///         });
    ///         var fooProjectLevelMrApprovals = new GitLab.ProjectLevelMrApprovals("fooProjectLevelMrApprovals", new GitLab.ProjectLevelMrApprovalsArgs
    ///         {
    ///             ProjectId = fooProject.Id,
    ///             ResetApprovalsOnPush = true,
    ///             DisableOverridingApproversPerMergeRequest = false,
    ///             MergeRequestsAuthorApproval = false,
    ///             MergeRequestsDisableCommittersApproval = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals # You can import an approval configuration state using `&lt;resource&gt; &lt;project_id&gt;`.
    /// ```
    /// 
    /// # # For example
    /// 
    /// ```sh
    ///  $ pulumi import gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals foo 1234
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals")]
    public partial class ProjectLevelMrApprovals : Pulumi.CustomResource
    {
        /// <summary>
        /// By default, users are able to edit the approval rules in merge requests. If set to true,
        /// </summary>
        [Output("disableOverridingApproversPerMergeRequest")]
        public Output<bool?> DisableOverridingApproversPerMergeRequest { get; private set; } = null!;

        /// <summary>
        /// Set to `true` if you want to allow merge request authors to self-approve merge requests. Authors
        /// </summary>
        [Output("mergeRequestsAuthorApproval")]
        public Output<bool?> MergeRequestsAuthorApproval { get; private set; } = null!;

        /// <summary>
        /// Set to `true` if you want to prevent approval of merge requests by merge request committers.
        /// </summary>
        [Output("mergeRequestsDisableCommittersApproval")]
        public Output<bool?> MergeRequestsDisableCommittersApproval { get; private set; } = null!;

        /// <summary>
        /// The ID of the project to change MR approval configuration.
        /// </summary>
        [Output("projectId")]
        public Output<int> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Set to `true` if you want to require authentication when approving a merge request.
        /// </summary>
        [Output("requirePasswordToApprove")]
        public Output<bool?> RequirePasswordToApprove { get; private set; } = null!;

        /// <summary>
        /// Set to `true` if you want to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
        /// </summary>
        [Output("resetApprovalsOnPush")]
        public Output<bool?> ResetApprovalsOnPush { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectLevelMrApprovals resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectLevelMrApprovals(string name, ProjectLevelMrApprovalsArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals", name, args ?? new ProjectLevelMrApprovalsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectLevelMrApprovals(string name, Input<string> id, ProjectLevelMrApprovalsState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProjectLevelMrApprovals resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectLevelMrApprovals Get(string name, Input<string> id, ProjectLevelMrApprovalsState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectLevelMrApprovals(name, id, state, options);
        }
    }

    public sealed class ProjectLevelMrApprovalsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// By default, users are able to edit the approval rules in merge requests. If set to true,
        /// </summary>
        [Input("disableOverridingApproversPerMergeRequest")]
        public Input<bool>? DisableOverridingApproversPerMergeRequest { get; set; }

        /// <summary>
        /// Set to `true` if you want to allow merge request authors to self-approve merge requests. Authors
        /// </summary>
        [Input("mergeRequestsAuthorApproval")]
        public Input<bool>? MergeRequestsAuthorApproval { get; set; }

        /// <summary>
        /// Set to `true` if you want to prevent approval of merge requests by merge request committers.
        /// </summary>
        [Input("mergeRequestsDisableCommittersApproval")]
        public Input<bool>? MergeRequestsDisableCommittersApproval { get; set; }

        /// <summary>
        /// The ID of the project to change MR approval configuration.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        /// <summary>
        /// Set to `true` if you want to require authentication when approving a merge request.
        /// </summary>
        [Input("requirePasswordToApprove")]
        public Input<bool>? RequirePasswordToApprove { get; set; }

        /// <summary>
        /// Set to `true` if you want to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
        /// </summary>
        [Input("resetApprovalsOnPush")]
        public Input<bool>? ResetApprovalsOnPush { get; set; }

        public ProjectLevelMrApprovalsArgs()
        {
        }
    }

    public sealed class ProjectLevelMrApprovalsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// By default, users are able to edit the approval rules in merge requests. If set to true,
        /// </summary>
        [Input("disableOverridingApproversPerMergeRequest")]
        public Input<bool>? DisableOverridingApproversPerMergeRequest { get; set; }

        /// <summary>
        /// Set to `true` if you want to allow merge request authors to self-approve merge requests. Authors
        /// </summary>
        [Input("mergeRequestsAuthorApproval")]
        public Input<bool>? MergeRequestsAuthorApproval { get; set; }

        /// <summary>
        /// Set to `true` if you want to prevent approval of merge requests by merge request committers.
        /// </summary>
        [Input("mergeRequestsDisableCommittersApproval")]
        public Input<bool>? MergeRequestsDisableCommittersApproval { get; set; }

        /// <summary>
        /// The ID of the project to change MR approval configuration.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Set to `true` if you want to require authentication when approving a merge request.
        /// </summary>
        [Input("requirePasswordToApprove")]
        public Input<bool>? RequirePasswordToApprove { get; set; }

        /// <summary>
        /// Set to `true` if you want to remove all approvals in a merge request when new commits are pushed to its source branch. Default is `true`.
        /// </summary>
        [Input("resetApprovalsOnPush")]
        public Input<bool>? ResetApprovalsOnPush { get; set; }

        public ProjectLevelMrApprovalsState()
        {
        }
    }
}
