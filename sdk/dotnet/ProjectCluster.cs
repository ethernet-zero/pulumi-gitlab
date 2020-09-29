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
    /// ## # gitlab\_project\_cluster
    /// 
    /// This resource allows you to create and manage project clusters for your GitLab projects.
    /// For further information on clusters, consult the [gitlab
    /// documentation](https://docs.gitlab.com/ce/user/project/clusters/index.html).
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
    ///         var foo = new GitLab.Project("foo", new GitLab.ProjectArgs
    ///         {
    ///         });
    ///         var bar = new GitLab.ProjectCluster("bar", new GitLab.ProjectClusterArgs
    ///         {
    ///             Domain = "example.com",
    ///             Enabled = true,
    ///             EnvironmentScope = "*",
    ///             KubernetesApiUrl = "https://124.124.124",
    ///             KubernetesAuthorizationType = "rbac",
    ///             KubernetesCaCert = "some-cert",
    ///             KubernetesNamespace = "namespace",
    ///             KubernetesToken = "some-token",
    ///             ManagementProjectId = "123456",
    ///             Project = foo.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class ProjectCluster : Pulumi.CustomResource
    {
        [Output("clusterType")]
        public Output<string> ClusterType { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The base domain of the cluster.
        /// </summary>
        [Output("domain")]
        public Output<string?> Domain { get; private set; } = null!;

        /// <summary>
        /// Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The associated environment to the cluster. Defaults to `*`.
        /// </summary>
        [Output("environmentScope")]
        public Output<string?> EnvironmentScope { get; private set; } = null!;

        /// <summary>
        /// The URL to access the Kubernetes API.
        /// </summary>
        [Output("kubernetesApiUrl")]
        public Output<string> KubernetesApiUrl { get; private set; } = null!;

        /// <summary>
        /// The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
        /// </summary>
        [Output("kubernetesAuthorizationType")]
        public Output<string?> KubernetesAuthorizationType { get; private set; } = null!;

        /// <summary>
        /// TLS certificate (needed if API is using a self-signed TLS certificate).
        /// </summary>
        [Output("kubernetesCaCert")]
        public Output<string?> KubernetesCaCert { get; private set; } = null!;

        /// <summary>
        /// The unique namespace related to the project.
        /// </summary>
        [Output("kubernetesNamespace")]
        public Output<string?> KubernetesNamespace { get; private set; } = null!;

        /// <summary>
        /// The token to authenticate against Kubernetes.
        /// </summary>
        [Output("kubernetesToken")]
        public Output<string> KubernetesToken { get; private set; } = null!;

        /// <summary>
        /// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
        /// </summary>
        [Output("managed")]
        public Output<bool?> Managed { get; private set; } = null!;

        /// <summary>
        /// The ID of the management project for the cluster.
        /// </summary>
        [Output("managementProjectId")]
        public Output<string?> ManagementProjectId { get; private set; } = null!;

        /// <summary>
        /// The name of cluster.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("platformType")]
        public Output<string> PlatformType { get; private set; } = null!;

        /// <summary>
        /// The id of the project to add the cluster to.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("providerType")]
        public Output<string> ProviderType { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectCluster(string name, ProjectClusterArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/projectCluster:ProjectCluster", name, args ?? new ProjectClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectCluster(string name, Input<string> id, ProjectClusterState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/projectCluster:ProjectCluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectCluster Get(string name, Input<string> id, ProjectClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectCluster(name, id, state, options);
        }
    }

    public sealed class ProjectClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The base domain of the cluster.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The associated environment to the cluster. Defaults to `*`.
        /// </summary>
        [Input("environmentScope")]
        public Input<string>? EnvironmentScope { get; set; }

        /// <summary>
        /// The URL to access the Kubernetes API.
        /// </summary>
        [Input("kubernetesApiUrl", required: true)]
        public Input<string> KubernetesApiUrl { get; set; } = null!;

        /// <summary>
        /// The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
        /// </summary>
        [Input("kubernetesAuthorizationType")]
        public Input<string>? KubernetesAuthorizationType { get; set; }

        /// <summary>
        /// TLS certificate (needed if API is using a self-signed TLS certificate).
        /// </summary>
        [Input("kubernetesCaCert")]
        public Input<string>? KubernetesCaCert { get; set; }

        /// <summary>
        /// The unique namespace related to the project.
        /// </summary>
        [Input("kubernetesNamespace")]
        public Input<string>? KubernetesNamespace { get; set; }

        /// <summary>
        /// The token to authenticate against Kubernetes.
        /// </summary>
        [Input("kubernetesToken", required: true)]
        public Input<string> KubernetesToken { get; set; } = null!;

        /// <summary>
        /// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
        /// </summary>
        [Input("managed")]
        public Input<bool>? Managed { get; set; }

        /// <summary>
        /// The ID of the management project for the cluster.
        /// </summary>
        [Input("managementProjectId")]
        public Input<string>? ManagementProjectId { get; set; }

        /// <summary>
        /// The name of cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The id of the project to add the cluster to.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        public ProjectClusterArgs()
        {
        }
    }

    public sealed class ProjectClusterState : Pulumi.ResourceArgs
    {
        [Input("clusterType")]
        public Input<string>? ClusterType { get; set; }

        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The base domain of the cluster.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Determines if cluster is active or not. Defaults to `true`. This attribute cannot be read.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The associated environment to the cluster. Defaults to `*`.
        /// </summary>
        [Input("environmentScope")]
        public Input<string>? EnvironmentScope { get; set; }

        /// <summary>
        /// The URL to access the Kubernetes API.
        /// </summary>
        [Input("kubernetesApiUrl")]
        public Input<string>? KubernetesApiUrl { get; set; }

        /// <summary>
        /// The cluster authorization type. Valid values are `rbac`, `abac`, `unknown_authorization`. Defaults to `rbac`.
        /// </summary>
        [Input("kubernetesAuthorizationType")]
        public Input<string>? KubernetesAuthorizationType { get; set; }

        /// <summary>
        /// TLS certificate (needed if API is using a self-signed TLS certificate).
        /// </summary>
        [Input("kubernetesCaCert")]
        public Input<string>? KubernetesCaCert { get; set; }

        /// <summary>
        /// The unique namespace related to the project.
        /// </summary>
        [Input("kubernetesNamespace")]
        public Input<string>? KubernetesNamespace { get; set; }

        /// <summary>
        /// The token to authenticate against Kubernetes.
        /// </summary>
        [Input("kubernetesToken")]
        public Input<string>? KubernetesToken { get; set; }

        /// <summary>
        /// Determines if cluster is managed by gitlab or not. Defaults to `true`. This attribute cannot be read.
        /// </summary>
        [Input("managed")]
        public Input<bool>? Managed { get; set; }

        /// <summary>
        /// The ID of the management project for the cluster.
        /// </summary>
        [Input("managementProjectId")]
        public Input<string>? ManagementProjectId { get; set; }

        /// <summary>
        /// The name of cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("platformType")]
        public Input<string>? PlatformType { get; set; }

        /// <summary>
        /// The id of the project to add the cluster to.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("providerType")]
        public Input<string>? ProviderType { get; set; }

        public ProjectClusterState()
        {
        }
    }
}
