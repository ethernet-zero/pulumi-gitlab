// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## # gitlab\_project
 *
 * Provide details about a specific project in the gitlab provider. The results include the name of the project, path, description, default branch, etc.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = pulumi.output(gitlab.getProject({
 *     id: "30",
 * }));
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = pulumi.output(gitlab.getProject({
 *     id: "foo/bar/baz",
 * }));
 * ```
 */
export function getProject(args: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gitlab:index/getProject:getProject", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectArgs {
    /**
     * The integer or path with namespace that uniquely identifies the project within the gitlab install.
     */
    id: string;
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
    readonly id: string;
    /**
     * Enable issue tracking for the project.
     */
    readonly issuesEnabled: boolean;
    /**
     * Enable LFS for the project.
     */
    readonly lfsEnabled: boolean;
    /**
     * Enable merge requests for the project.
     */
    readonly mergeRequestsEnabled: boolean;
    readonly name: string;
    /**
     * The namespace (group or user) of the project. Defaults to your user.
     * See `gitlab.Group` for an example.
     */
    readonly namespaceId: number;
    /**
     * The path of the repository.
     */
    readonly path: string;
    /**
     * The path of the repository with namespace.
     */
    readonly pathWithNamespace: string;
    /**
     * Enable pipelines for the project.
     */
    readonly pipelinesEnabled: boolean;
    readonly pushRules: outputs.GetProjectPushRules;
    /**
     * Enable `Delete source branch` option by default for all new merge requests
     */
    readonly removeSourceBranchAfterMerge: boolean;
    /**
     * Allow users to request member access.
     */
    readonly requestAccessEnabled: boolean;
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

export function getProjectOutput(args: GetProjectOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectResult> {
    return pulumi.output(args).apply(a => getProject(a, opts))
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectOutputArgs {
    /**
     * The integer or path with namespace that uniquely identifies the project within the gitlab install.
     */
    id: pulumi.Input<string>;
}
