// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage Group Access Token for your GitLab Groups. (Introduced in GitLab 14.7)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const exampleGroupAccessToken = new gitlab.GroupAccessToken("exampleGroupAccessToken", {
 *     group: "25",
 *     expiresAt: "2020-03-14",
 *     accessLevel: "developer",
 *     scopes: ["api"],
 * });
 * const exampleGroupVariable = new gitlab.GroupVariable("exampleGroupVariable", {
 *     group: "25",
 *     key: "gat",
 *     value: exampleGroupAccessToken.token,
 * });
 * ```
 *
 * ## Import
 *
 * # A GitLab Group Access Token can be imported using a key composed of `<group-id>:<token-id>`, e.g.
 *
 * ```sh
 *  $ pulumi import gitlab:index/groupAccessToken:GroupAccessToken example "12345:1"
 * ```
 *
 * # ATTENTIONthe `token` resource attribute is not available for imported resources as this information cannot be read from the GitLab API.
 */
export class GroupAccessToken extends pulumi.CustomResource {
    /**
     * Get an existing GroupAccessToken resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupAccessTokenState, opts?: pulumi.CustomResourceOptions): GroupAccessToken {
        return new GroupAccessToken(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/groupAccessToken:GroupAccessToken';

    /**
     * Returns true if the given object is an instance of GroupAccessToken.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupAccessToken {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupAccessToken.__pulumiType;
    }

    /**
     * The access level for the group access token. Valid values are: `guest`, `reporter`, `developer`, `maintainer`.
     */
    public readonly accessLevel!: pulumi.Output<string | undefined>;
    /**
     * True if the token is active.
     */
    public /*out*/ readonly active!: pulumi.Output<boolean>;
    /**
     * Time the token has been created, RFC3339 format.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD. Default is never.
     */
    public readonly expiresAt!: pulumi.Output<string | undefined>;
    /**
     * The ID or path of the group to add the group access token to.
     */
    public readonly group!: pulumi.Output<string>;
    /**
     * The name of the group access token.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * True if the token is revoked.
     */
    public /*out*/ readonly revoked!: pulumi.Output<boolean>;
    /**
     * The scope for the group access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`.
     */
    public readonly scopes!: pulumi.Output<string[]>;
    /**
     * The group access token. This is only populated when creating a new group access token. This attribute is not available for imported resources.
     */
    public /*out*/ readonly token!: pulumi.Output<string>;
    /**
     * The user id associated to the token.
     */
    public /*out*/ readonly userId!: pulumi.Output<number>;

    /**
     * Create a GroupAccessToken resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupAccessTokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupAccessTokenArgs | GroupAccessTokenState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupAccessTokenState | undefined;
            resourceInputs["accessLevel"] = state ? state.accessLevel : undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["group"] = state ? state.group : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["revoked"] = state ? state.revoked : undefined;
            resourceInputs["scopes"] = state ? state.scopes : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as GroupAccessTokenArgs | undefined;
            if ((!args || args.group === undefined) && !opts.urn) {
                throw new Error("Missing required property 'group'");
            }
            if ((!args || args.scopes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scopes'");
            }
            resourceInputs["accessLevel"] = args ? args.accessLevel : undefined;
            resourceInputs["expiresAt"] = args ? args.expiresAt : undefined;
            resourceInputs["group"] = args ? args.group : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["scopes"] = args ? args.scopes : undefined;
            resourceInputs["active"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["revoked"] = undefined /*out*/;
            resourceInputs["token"] = undefined /*out*/;
            resourceInputs["userId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GroupAccessToken.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupAccessToken resources.
 */
export interface GroupAccessTokenState {
    /**
     * The access level for the group access token. Valid values are: `guest`, `reporter`, `developer`, `maintainer`.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * True if the token is active.
     */
    active?: pulumi.Input<boolean>;
    /**
     * Time the token has been created, RFC3339 format.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD. Default is never.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The ID or path of the group to add the group access token to.
     */
    group?: pulumi.Input<string>;
    /**
     * The name of the group access token.
     */
    name?: pulumi.Input<string>;
    /**
     * True if the token is revoked.
     */
    revoked?: pulumi.Input<boolean>;
    /**
     * The scope for the group access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`.
     */
    scopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The group access token. This is only populated when creating a new group access token. This attribute is not available for imported resources.
     */
    token?: pulumi.Input<string>;
    /**
     * The user id associated to the token.
     */
    userId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a GroupAccessToken resource.
 */
export interface GroupAccessTokenArgs {
    /**
     * The access level for the group access token. Valid values are: `guest`, `reporter`, `developer`, `maintainer`.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * The token expires at midnight UTC on that date. The date must be in the format YYYY-MM-DD. Default is never.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The ID or path of the group to add the group access token to.
     */
    group: pulumi.Input<string>;
    /**
     * The name of the group access token.
     */
    name?: pulumi.Input<string>;
    /**
     * The scope for the group access token. It determines the actions which can be performed when authenticating with this token. Valid values are: `api`, `readApi`, `readRegistry`, `writeRegistry`, `readRepository`, `writeRepository`.
     */
    scopes: pulumi.Input<pulumi.Input<string>[]>;
}
