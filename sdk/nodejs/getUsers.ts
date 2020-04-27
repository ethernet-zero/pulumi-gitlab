// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides details about a list of users in the gitlab provider. The results include id, username, email, name and more about the requested users. Users can also be sorted and filtered using several options.
 * 
 * **NOTE**: Some of the available options require administrator privileges. Please visit [Gitlab API documentation][usersForAdmins] for more information.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 * 
 * const example = pulumi.output(gitlab.getUsers({
 *     createdBefore: "2019-01-01",
 *     orderBy: "name",
 *     sort: "desc",
 * }, { async: true }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/d/users.html.markdown.
 */
export function getUsers(args?: GetUsersArgs, opts?: pulumi.InvokeOptions): Promise<GetUsersResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gitlab:index/getUsers:getUsers", {
        "active": args.active,
        "blocked": args.blocked,
        "createdAfter": args.createdAfter,
        "createdBefore": args.createdBefore,
        "externProvider": args.externProvider,
        "externUid": args.externUid,
        "orderBy": args.orderBy,
        "search": args.search,
        "sort": args.sort,
    }, opts);
}

/**
 * A collection of arguments for invoking getUsers.
 */
export interface GetUsersArgs {
    /**
     * Filter users that are active.
     */
    readonly active?: boolean;
    /**
     * Filter users that are blocked.
     */
    readonly blocked?: boolean;
    /**
     * Search for users created after a specific date. (Requires administrator privileges)
     */
    readonly createdAfter?: string;
    /**
     * Search for users created before a specific date. (Requires administrator privileges)
     */
    readonly createdBefore?: string;
    /**
     * Lookup users by external provider. (Requires administrator privileges)
     */
    readonly externProvider?: string;
    /**
     * Lookup users by external UID. (Requires administrator privileges)
     */
    readonly externUid?: string;
    /**
     * Order the users' list by `id`, `name`, `username`, `createdAt` or `updatedAt`. (Requires administrator privileges)
     */
    readonly orderBy?: string;
    /**
     * Search users by username, name or email.
     */
    readonly search?: string;
    /**
     * Sort users' list in asc or desc order. (Requires administrator privileges)
     */
    readonly sort?: string;
}

/**
 * A collection of values returned by getUsers.
 */
export interface GetUsersResult {
    readonly active?: boolean;
    readonly blocked?: boolean;
    readonly createdAfter?: string;
    readonly createdBefore?: string;
    readonly externProvider?: string;
    /**
     * The external UID of the user.
     */
    readonly externUid?: string;
    readonly orderBy?: string;
    readonly search?: string;
    readonly sort?: string;
    /**
     * The list of users.
     */
    readonly users: outputs.GetUsersUser[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
