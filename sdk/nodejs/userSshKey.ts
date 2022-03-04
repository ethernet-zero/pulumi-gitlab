// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows to manage GitLab user SSH keys.
 *
 * **Upstream API**: [GitLab API docs](https://docs.gitlab.com/ee/api/users.html#single-ssh-key)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const exampleUser = gitlab.getUser({
 *     username: "example-user",
 * });
 * const exampleUserSshKey = new gitlab.UserSshKey("exampleUserSshKey", {
 *     userId: data.gitlab_user.id,
 *     title: "example-key",
 *     key: "ssh-rsa AAAA...",
 *     expiresAt: "2016-01-21T00:00:00.000Z",
 * });
 * ```
 *
 * ## Import
 *
 * # You can import a user ssh key using an id made up of `{user-id}:{key}`, e.g.
 *
 * ```sh
 *  $ pulumi import gitlab:index/userSshKey:UserSshKey example 42:1
 * ```
 */
export class UserSshKey extends pulumi.CustomResource {
    /**
     * Get an existing UserSshKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserSshKeyState, opts?: pulumi.CustomResourceOptions): UserSshKey {
        return new UserSshKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/userSshKey:UserSshKey';

    /**
     * Returns true if the given object is an instance of UserSshKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserSshKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserSshKey.__pulumiType;
    }

    /**
     * The time when this key was created in GitLab.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The expiration date of the SSH key in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)
     */
    public readonly expiresAt!: pulumi.Output<string | undefined>;
    /**
     * The ssh key. The SSH key `comment` (trailing part) is optional and ignored for diffing, because GitLab overrides it with the username and GitLab hostname.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * The ID of the ssh key.
     */
    public /*out*/ readonly keyId!: pulumi.Output<number>;
    /**
     * The title of the ssh key.
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * The ID of the user to add the ssh key to.
     */
    public readonly userId!: pulumi.Output<number>;

    /**
     * Create a UserSshKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserSshKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserSshKeyArgs | UserSshKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserSshKeyState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["keyId"] = state ? state.keyId : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as UserSshKeyArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.title === undefined) && !opts.urn) {
                throw new Error("Missing required property 'title'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            resourceInputs["expiresAt"] = args ? args.expiresAt : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["keyId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserSshKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserSshKey resources.
 */
export interface UserSshKeyState {
    /**
     * The time when this key was created in GitLab.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The expiration date of the SSH key in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The ssh key. The SSH key `comment` (trailing part) is optional and ignored for diffing, because GitLab overrides it with the username and GitLab hostname.
     */
    key?: pulumi.Input<string>;
    /**
     * The ID of the ssh key.
     */
    keyId?: pulumi.Input<number>;
    /**
     * The title of the ssh key.
     */
    title?: pulumi.Input<string>;
    /**
     * The ID of the user to add the ssh key to.
     */
    userId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a UserSshKey resource.
 */
export interface UserSshKeyArgs {
    /**
     * The expiration date of the SSH key in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The ssh key. The SSH key `comment` (trailing part) is optional and ignored for diffing, because GitLab overrides it with the username and GitLab hostname.
     */
    key: pulumi.Input<string>;
    /**
     * The title of the ssh key.
     */
    title: pulumi.Input<string>;
    /**
     * The ID of the user to add the ssh key to.
     */
    userId: pulumi.Input<number>;
}
