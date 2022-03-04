// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GitLab
{
    public static class GetProject
    {
        /// <summary>
        /// Provide details about a specific project in the gitlab provider. The results include the name of the project, path, description, default branch, etc.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using GitLab = Pulumi.GitLab;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(GitLab.GetProject.InvokeAsync(new GitLab.GetProjectArgs
        ///         {
        ///             Id = "foo/bar/baz",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProjectResult> InvokeAsync(GetProjectArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProjectResult>("gitlab:index/getProject:getProject", args ?? new GetProjectArgs(), options.WithDefaults());

        /// <summary>
        /// Provide details about a specific project in the gitlab provider. The results include the name of the project, path, description, default branch, etc.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using GitLab = Pulumi.GitLab;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(GitLab.GetProject.InvokeAsync(new GitLab.GetProjectArgs
        ///         {
        ///             Id = "foo/bar/baz",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetProjectResult> Invoke(GetProjectInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProjectResult>("gitlab:index/getProject:getProject", args ?? new GetProjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The integer or path with namespace that uniquely identifies the project within the gitlab install.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The path of the repository with namespace.
        /// </summary>
        [Input("pathWithNamespace")]
        public string? PathWithNamespace { get; set; }

        public GetProjectArgs()
        {
        }
    }

    public sealed class GetProjectInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The integer or path with namespace that uniquely identifies the project within the gitlab install.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The path of the repository with namespace.
        /// </summary>
        [Input("pathWithNamespace")]
        public Input<string>? PathWithNamespace { get; set; }

        public GetProjectInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProjectResult
    {
        /// <summary>
        /// Whether the project is in read-only mode (archived).
        /// </summary>
        public readonly bool Archived;
        /// <summary>
        /// The default branch for the project.
        /// </summary>
        public readonly string DefaultBranch;
        /// <summary>
        /// A description of the project.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// URL that can be provided to `git clone` to clone the
        /// </summary>
        public readonly string HttpUrlToRepo;
        /// <summary>
        /// The integer or path with namespace that uniquely identifies the project within the gitlab install.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Enable issue tracking for the project.
        /// </summary>
        public readonly bool IssuesEnabled;
        /// <summary>
        /// Enable LFS for the project.
        /// </summary>
        public readonly bool LfsEnabled;
        /// <summary>
        /// Enable or disable merge pipelines.
        /// </summary>
        public readonly bool MergePipelinesEnabled;
        /// <summary>
        /// Enable merge requests for the project.
        /// </summary>
        public readonly bool MergeRequestsEnabled;
        /// <summary>
        /// Enable or disable merge trains.
        /// </summary>
        public readonly bool MergeTrainsEnabled;
        /// <summary>
        /// The name of the project.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The namespace (group or user) of the project. Defaults to your user.
        /// </summary>
        public readonly int NamespaceId;
        /// <summary>
        /// The path of the repository.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// The path of the repository with namespace.
        /// </summary>
        public readonly string PathWithNamespace;
        /// <summary>
        /// Enable pipelines for the project.
        /// </summary>
        public readonly bool PipelinesEnabled;
        /// <summary>
        /// Show link to create/view merge request when pushing from the command line
        /// </summary>
        public readonly bool PrintingMergeRequestLinkEnabled;
        /// <summary>
        /// Push rules for the project.
        /// </summary>
        public readonly Outputs.GetProjectPushRulesResult PushRules;
        /// <summary>
        /// Enable `Delete source branch` option by default for all new merge requests
        /// </summary>
        public readonly bool RemoveSourceBranchAfterMerge;
        /// <summary>
        /// Allow users to request member access.
        /// </summary>
        public readonly bool RequestAccessEnabled;
        /// <summary>
        /// Registration token to use during runner setup.
        /// </summary>
        public readonly string RunnersToken;
        /// <summary>
        /// Enable snippets for the project.
        /// </summary>
        public readonly bool SnippetsEnabled;
        /// <summary>
        /// URL that can be provided to `git clone` to clone the
        /// </summary>
        public readonly string SshUrlToRepo;
        /// <summary>
        /// Repositories are created as private by default.
        /// </summary>
        public readonly string VisibilityLevel;
        /// <summary>
        /// URL that can be used to find the project in a browser.
        /// </summary>
        public readonly string WebUrl;
        /// <summary>
        /// Enable wiki for the project.
        /// </summary>
        public readonly bool WikiEnabled;

        [OutputConstructor]
        private GetProjectResult(
            bool archived,

            string defaultBranch,

            string description,

            string httpUrlToRepo,

            string id,

            bool issuesEnabled,

            bool lfsEnabled,

            bool mergePipelinesEnabled,

            bool mergeRequestsEnabled,

            bool mergeTrainsEnabled,

            string name,

            int namespaceId,

            string path,

            string pathWithNamespace,

            bool pipelinesEnabled,

            bool printingMergeRequestLinkEnabled,

            Outputs.GetProjectPushRulesResult pushRules,

            bool removeSourceBranchAfterMerge,

            bool requestAccessEnabled,

            string runnersToken,

            bool snippetsEnabled,

            string sshUrlToRepo,

            string visibilityLevel,

            string webUrl,

            bool wikiEnabled)
        {
            Archived = archived;
            DefaultBranch = defaultBranch;
            Description = description;
            HttpUrlToRepo = httpUrlToRepo;
            Id = id;
            IssuesEnabled = issuesEnabled;
            LfsEnabled = lfsEnabled;
            MergePipelinesEnabled = mergePipelinesEnabled;
            MergeRequestsEnabled = mergeRequestsEnabled;
            MergeTrainsEnabled = mergeTrainsEnabled;
            Name = name;
            NamespaceId = namespaceId;
            Path = path;
            PathWithNamespace = pathWithNamespace;
            PipelinesEnabled = pipelinesEnabled;
            PrintingMergeRequestLinkEnabled = printingMergeRequestLinkEnabled;
            PushRules = pushRules;
            RemoveSourceBranchAfterMerge = removeSourceBranchAfterMerge;
            RequestAccessEnabled = requestAccessEnabled;
            RunnersToken = runnersToken;
            SnippetsEnabled = snippetsEnabled;
            SshUrlToRepo = sshUrlToRepo;
            VisibilityLevel = visibilityLevel;
            WebUrl = webUrl;
            WikiEnabled = wikiEnabled;
        }
    }
}
