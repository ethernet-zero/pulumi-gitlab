// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides details about a specific project in the gitlab provider. The results include the name of the project, path, description, default branch, etc.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 * 
 * const example = pulumi.output(gitlab.getProject({
 *     id: 30,
 * }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/d/project.html.markdown.
 */
export function getProject(args: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> & GetProjectResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetProjectResult> = pulumi.runtime.invoke("gitlab:index/getProject:getProject", {
        "archived": args.archived,
        "defaultBranch": args.defaultBranch,
        "description": args.description,
        "httpUrlToRepo": args.httpUrlToRepo,
        "id": args.id,
        "issuesEnabled": args.issuesEnabled,
        "mergeRequestsEnabled": args.mergeRequestsEnabled,
        "name": args.name,
        "namespaceId": args.namespaceId,
        "path": args.path,
        "runnersToken": args.runnersToken,
        "snippetsEnabled": args.snippetsEnabled,
        "sshUrlToRepo": args.sshUrlToRepo,
        "visibilityLevel": args.visibilityLevel,
        "webUrl": args.webUrl,
        "wikiEnabled": args.wikiEnabled,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectArgs {
    readonly archived?: boolean;
    readonly defaultBranch?: string;
    readonly description?: string;
    readonly httpUrlToRepo?: string;
    /**
     * The integer that uniquely identifies the project within the gitlab install.
     */
    readonly id: number;
    readonly issuesEnabled?: boolean;
    readonly mergeRequestsEnabled?: boolean;
    readonly name?: string;
    readonly namespaceId?: number;
    readonly path?: string;
    readonly runnersToken?: string;
    readonly snippetsEnabled?: boolean;
    readonly sshUrlToRepo?: string;
    readonly visibilityLevel?: string;
    readonly webUrl?: string;
    readonly wikiEnabled?: boolean;
}

/**
 * A collection of values returned by getProject.
 */
export interface GetProjectResult {
    /**
     * Whether the project is in read-only mode (archived).
     */
    readonly archived: boolean;
    /**
     * The default branch for the project.
     */
    readonly defaultBranch: string;
    /**
     * A description of the project.
     */
    readonly description: string;
    /**
     * URL that can be provided to `git clone` to clone the
     * repository via HTTP.
     */
    readonly httpUrlToRepo: string;
    /**
     * Integer that uniquely identifies the project within the gitlab install.
     */
    readonly id: number;
    /**
     * Enable issue tracking for the project.
     */
    readonly issuesEnabled: boolean;
    /**
     * Enable merge requests for the project.
     */
    readonly mergeRequestsEnabled: boolean;
    readonly name: string;
    /**
     * The namespace (group or user) of the project. Defaults to your user.
     * See `gitlab..Group` for an example.
     */
    readonly namespaceId: number;
    /**
     * The path of the repository.
     */
    readonly path: string;
    /**
     * Registration token to use during runner setup.
     */
    readonly runnersToken: string;
    /**
     * Enable snippets for the project.
     */
    readonly snippetsEnabled: boolean;
    /**
     * URL that can be provided to `git clone` to clone the
     * repository via SSH.
     */
    readonly sshUrlToRepo: string;
    /**
     * Repositories are created as private by default.
     */
    readonly visibilityLevel: string;
    /**
     * URL that can be used to find the project in a browser.
     */
    readonly webUrl: string;
    /**
     * Enable wiki for the project.
     */
    readonly wikiEnabled: boolean;
}
