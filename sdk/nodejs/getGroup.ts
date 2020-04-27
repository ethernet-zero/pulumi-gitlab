// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides details about a specific group in the gitlab provider.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 * 
 * const foo = pulumi.output(gitlab.getGroup({
 *     groupId: 123,
 * }, { async: true }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/d/group.html.markdown.
 */
export function getGroup(args?: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gitlab:index/getGroup:getGroup", {
        "fullPath": args.fullPath,
        "groupId": args.groupId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupArgs {
    /**
     * The full path of the group.
     */
    readonly fullPath?: string;
    /**
     * The ID of the group.
     */
    readonly groupId?: number;
}

/**
 * A collection of values returned by getGroup.
 */
export interface GetGroupResult {
    /**
     * The description of the group.
     */
    readonly description: string;
    /**
     * The full name of the group.
     */
    readonly fullName: string;
    /**
     * The full path of the group.
     */
    readonly fullPath: string;
    readonly groupId: number;
    /**
     * Boolean, is LFS enabled for projects in this group.
     */
    readonly lfsEnabled: boolean;
    /**
     * The name of this group.
     */
    readonly name: string;
    /**
     * Integer, ID of the parent group.
     */
    readonly parentId: number;
    /**
     * The path of the group.
     */
    readonly path: string;
    /**
     * Boolean, is request for access enabled to the group.
     */
    readonly requestAccessEnabled: boolean;
    /**
     * The group level registration token to use during runner setup.
     */
    readonly runnersToken: string;
    /**
     * Visibility level of the group. Possible values are `private`, `internal`, `public`.
     */
    readonly visibilityLevel: string;
    /**
     * Web URL of the group.
     */
    readonly webUrl: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
