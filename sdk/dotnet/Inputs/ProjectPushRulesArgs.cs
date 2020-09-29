// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab.Inputs
{

    public sealed class ProjectPushRulesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// All commit author emails must match this regex, e.g. `@my-company.com$`.
        /// </summary>
        [Input("authorEmailRegex")]
        public Input<string>? AuthorEmailRegex { get; set; }

        /// <summary>
        /// All branch names must match this regex, e.g. `(feature|hotfix)\/*`.
        /// </summary>
        [Input("branchNameRegex")]
        public Input<string>? BranchNameRegex { get; set; }

        /// <summary>
        /// Users can only push commits to this repository that were committed with one of their own verified emails.
        /// </summary>
        [Input("commitCommitterCheck")]
        public Input<bool>? CommitCommitterCheck { get; set; }

        /// <summary>
        /// No commit message is allowed to match this regex, for example `ssh\:\/\/`.
        /// </summary>
        [Input("commitMessageNegativeRegex")]
        public Input<string>? CommitMessageNegativeRegex { get; set; }

        /// <summary>
        /// All commit messages must match this regex, e.g. `Fixed \d+\..*`.
        /// </summary>
        [Input("commitMessageRegex")]
        public Input<string>? CommitMessageRegex { get; set; }

        /// <summary>
        /// Deny deleting a tag.
        /// </summary>
        [Input("denyDeleteTag")]
        public Input<bool>? DenyDeleteTag { get; set; }

        /// <summary>
        /// All commited filenames must not match this regex, e.g. `(jar|exe)$`.
        /// </summary>
        [Input("fileNameRegex")]
        public Input<string>? FileNameRegex { get; set; }

        /// <summary>
        /// Maximum file size (MB).
        /// </summary>
        [Input("maxFileSize")]
        public Input<int>? MaxFileSize { get; set; }

        /// <summary>
        /// Restrict commits by author (email) to existing GitLab users.
        /// </summary>
        [Input("memberCheck")]
        public Input<bool>? MemberCheck { get; set; }

        /// <summary>
        /// GitLab will reject any files that are likely to contain secrets.
        /// </summary>
        [Input("preventSecrets")]
        public Input<bool>? PreventSecrets { get; set; }

        /// <summary>
        /// Reject commit when it’s not signed through GPG.
        /// </summary>
        [Input("rejectUnsignedCommits")]
        public Input<bool>? RejectUnsignedCommits { get; set; }

        public ProjectPushRulesArgs()
        {
        }
    }
}
