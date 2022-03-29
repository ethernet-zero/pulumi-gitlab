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
    /// The `gitlab.GroupLdapLink` resource allows to manage the lifecycle of an LDAP integration with a group.
    /// 
    /// **Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#ldap-group-links)
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
    ///         var test = new GitLab.GroupLdapLink("test", new GitLab.GroupLdapLinkArgs
    ///         {
    ///             Cn = "testuser",
    ///             GroupAccess = "developer",
    ///             GroupId = "12345",
    ///             LdapProvider = "ldapmain",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// # GitLab group ldap links can be imported using an id made up of `group_id:ldap_provider:cn`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gitlab:index/groupLdapLink:GroupLdapLink test "12345:ldapmain:testuser"
    /// ```
    /// </summary>
    [GitLabResourceType("gitlab:index/groupLdapLink:GroupLdapLink")]
    public partial class GroupLdapLink : Pulumi.CustomResource
    {
        /// <summary>
        /// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        /// </summary>
        [Output("accessLevel")]
        public Output<string?> AccessLevel { get; private set; } = null!;

        /// <summary>
        /// The CN of the LDAP group to link with.
        /// </summary>
        [Output("cn")]
        public Output<string> Cn { get; private set; } = null!;

        /// <summary>
        /// If true, then delete and replace an existing LDAP link if one exists.
        /// </summary>
        [Output("force")]
        public Output<bool?> Force { get; private set; } = null!;

        /// <summary>
        /// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        /// </summary>
        [Output("groupAccess")]
        public Output<string?> GroupAccess { get; private set; } = null!;

        /// <summary>
        /// The id of the GitLab group.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// The name of the LDAP provider as stored in the GitLab database.
        /// </summary>
        [Output("ldapProvider")]
        public Output<string> LdapProvider { get; private set; } = null!;


        /// <summary>
        /// Create a GroupLdapLink resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupLdapLink(string name, GroupLdapLinkArgs args, CustomResourceOptions? options = null)
            : base("gitlab:index/groupLdapLink:GroupLdapLink", name, args ?? new GroupLdapLinkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupLdapLink(string name, Input<string> id, GroupLdapLinkState? state = null, CustomResourceOptions? options = null)
            : base("gitlab:index/groupLdapLink:GroupLdapLink", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupLdapLink resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupLdapLink Get(string name, Input<string> id, GroupLdapLinkState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupLdapLink(name, id, state, options);
        }
    }

    public sealed class GroupLdapLinkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        /// </summary>
        [Input("accessLevel")]
        public Input<string>? AccessLevel { get; set; }

        /// <summary>
        /// The CN of the LDAP group to link with.
        /// </summary>
        [Input("cn", required: true)]
        public Input<string> Cn { get; set; } = null!;

        /// <summary>
        /// If true, then delete and replace an existing LDAP link if one exists.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        /// </summary>
        [Input("groupAccess")]
        public Input<string>? GroupAccess { get; set; }

        /// <summary>
        /// The id of the GitLab group.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// The name of the LDAP provider as stored in the GitLab database.
        /// </summary>
        [Input("ldapProvider", required: true)]
        public Input<string> LdapProvider { get; set; } = null!;

        public GroupLdapLinkArgs()
        {
        }
    }

    public sealed class GroupLdapLinkState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        /// </summary>
        [Input("accessLevel")]
        public Input<string>? AccessLevel { get; set; }

        /// <summary>
        /// The CN of the LDAP group to link with.
        /// </summary>
        [Input("cn")]
        public Input<string>? Cn { get; set; }

        /// <summary>
        /// If true, then delete and replace an existing LDAP link if one exists.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// Minimum access level for members of the LDAP group. Valid values are: `no one`, `minimal`, `guest`, `reporter`, `developer`, `maintainer`, `owner`, `master`
        /// </summary>
        [Input("groupAccess")]
        public Input<string>? GroupAccess { get; set; }

        /// <summary>
        /// The id of the GitLab group.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The name of the LDAP provider as stored in the GitLab database.
        /// </summary>
        [Input("ldapProvider")]
        public Input<string>? LdapProvider { get; set; }

        public GroupLdapLinkState()
        {
        }
    }
}
