// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # gitlab\_project_mirror
 *
 * This resource allows you to add a mirror target for the repository, all changes will be synced to the remote target.
 *
 * > This is for *pushing* changes to a remote repository. *Pull Mirroring* can be configured using a combination of the
 * `importUrl`, `mirror`, and `mirrorTriggerBuilds` properties on the `gitlab.Project` resource.
 *
 * For further information on mirroring, consult the
 * [gitlab documentation](https://docs.gitlab.com/ee/user/project/repository/repository_mirroring.html#repository-mirroring).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const foo = new gitlab.ProjectMirror("foo", {
 *     project: "1",
 *     url: "https://username:password@github.com/org/repository.git",
 * });
 * ```
 *
 * ## Import
 *
 * GitLab project mirror can be imported using an id made up of `project_id:mirror_id`, e.g.
 *
 * ```sh
 *  $ pulumi import gitlab:index/projectMirror:ProjectMirror foo "12345:1337"
 * ```
 */
export class ProjectMirror extends pulumi.CustomResource {
    /**
     * Get an existing ProjectMirror resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectMirrorState, opts?: pulumi.CustomResourceOptions): ProjectMirror {
        return new ProjectMirror(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectMirror:ProjectMirror';

    /**
     * Returns true if the given object is an instance of ProjectMirror.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectMirror {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectMirror.__pulumiType;
    }

    /**
     * Determines if the mirror is enabled.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Determines if divergent refs are skipped.
     */
    public readonly keepDivergentRefs!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly mirrorId!: pulumi.Output<number>;
    /**
     * Determines if only protected branches are mirrored.
     */
    public readonly onlyProtectedBranches!: pulumi.Output<boolean | undefined>;
    /**
     * The id of the project.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URL of the remote repository to be mirrored.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a ProjectMirror resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectMirrorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectMirrorArgs | ProjectMirrorState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProjectMirrorState | undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["keepDivergentRefs"] = state ? state.keepDivergentRefs : undefined;
            inputs["mirrorId"] = state ? state.mirrorId : undefined;
            inputs["onlyProtectedBranches"] = state ? state.onlyProtectedBranches : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as ProjectMirrorArgs | undefined;
            if ((!args || args.project === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.url === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'url'");
            }
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["keepDivergentRefs"] = args ? args.keepDivergentRefs : undefined;
            inputs["onlyProtectedBranches"] = args ? args.onlyProtectedBranches : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["url"] = args ? args.url : undefined;
            inputs["mirrorId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ProjectMirror.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectMirror resources.
 */
export interface ProjectMirrorState {
    /**
     * Determines if the mirror is enabled.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Determines if divergent refs are skipped.
     */
    readonly keepDivergentRefs?: pulumi.Input<boolean>;
    readonly mirrorId?: pulumi.Input<number>;
    /**
     * Determines if only protected branches are mirrored.
     */
    readonly onlyProtectedBranches?: pulumi.Input<boolean>;
    /**
     * The id of the project.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URL of the remote repository to be mirrored.
     */
    readonly url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectMirror resource.
 */
export interface ProjectMirrorArgs {
    /**
     * Determines if the mirror is enabled.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Determines if divergent refs are skipped.
     */
    readonly keepDivergentRefs?: pulumi.Input<boolean>;
    /**
     * Determines if only protected branches are mirrored.
     */
    readonly onlyProtectedBranches?: pulumi.Input<boolean>;
    /**
     * The id of the project.
     */
    readonly project: pulumi.Input<string>;
    /**
     * The URL of the remote repository to be mirrored.
     */
    readonly url: pulumi.Input<string>;
}
