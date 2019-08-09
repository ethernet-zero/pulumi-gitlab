// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/project.html.markdown.
 */
export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * Number of merge request approvals required for merging. Default is 0.
     */
    public readonly approvalsBeforeMerge!: pulumi.Output<number | undefined>;
    /**
     * Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
     */
    public readonly archived!: pulumi.Output<boolean | undefined>;
    /**
     * Enable container registry for the project.
     */
    public readonly containerRegistryEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The default branch for the project.
     */
    public readonly defaultBranch!: pulumi.Output<string | undefined>;
    /**
     * A description of the project.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * URL that can be provided to `git clone` to clone the
     * repository via HTTP.
     */
    public /*out*/ readonly httpUrlToRepo!: pulumi.Output<string>;
    /**
     * Enable issue tracking for the project.
     */
    public readonly issuesEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Set to `ff` to create fast-forward merges
     * Valid values are `merge`, `rebaseMerge`, `ff`
     * Repositories are created with `merge` by default
     */
    public readonly mergeMethod!: pulumi.Output<string | undefined>;
    /**
     * Enable merge requests for the project.
     */
    public readonly mergeRequestsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the project.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The namespace (group or user) of the project. Defaults to your user.
     * See `gitlab..Group` for an example.
     */
    public readonly namespaceId!: pulumi.Output<number>;
    /**
     * Set to true if you want allow merges only if all discussions are resolved.
     */
    public readonly onlyAllowMergeIfAllDiscussionsAreResolved!: pulumi.Output<boolean | undefined>;
    /**
     * Set to true if you want allow merges only if a pipeline succeeds.
     */
    public readonly onlyAllowMergeIfPipelineSucceeds!: pulumi.Output<boolean | undefined>;
    /**
     * The path of the repository.
     */
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * Registration token to use during runner setup.
     */
    public /*out*/ readonly runnersToken!: pulumi.Output<string>;
    /**
     * Enable shared runners for this project.
     */
    public readonly sharedRunnersEnabled!: pulumi.Output<boolean>;
    /**
     * Enable sharing the project with a list of groups (maps).
     */
    public readonly sharedWithGroups!: pulumi.Output<{ groupAccessLevel: string, groupId: number, groupName: string }[] | undefined>;
    /**
     * Enable snippets for the project.
     */
    public readonly snippetsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * URL that can be provided to `git clone` to clone the
     * repository via SSH.
     */
    public /*out*/ readonly sshUrlToRepo!: pulumi.Output<string>;
    /**
     * Tags (topics) of the project.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * Set to `public` to create a public project.
     * Valid values are `private`, `internal`, `public`.
     * Repositories are created as private by default.
     */
    public readonly visibilityLevel!: pulumi.Output<string | undefined>;
    /**
     * URL that can be used to find the project in a browser.
     */
    public /*out*/ readonly webUrl!: pulumi.Output<string>;
    /**
     * Enable wiki for the project.
     */
    public readonly wikiEnabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProjectState | undefined;
            inputs["approvalsBeforeMerge"] = state ? state.approvalsBeforeMerge : undefined;
            inputs["archived"] = state ? state.archived : undefined;
            inputs["containerRegistryEnabled"] = state ? state.containerRegistryEnabled : undefined;
            inputs["defaultBranch"] = state ? state.defaultBranch : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["httpUrlToRepo"] = state ? state.httpUrlToRepo : undefined;
            inputs["issuesEnabled"] = state ? state.issuesEnabled : undefined;
            inputs["mergeMethod"] = state ? state.mergeMethod : undefined;
            inputs["mergeRequestsEnabled"] = state ? state.mergeRequestsEnabled : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namespaceId"] = state ? state.namespaceId : undefined;
            inputs["onlyAllowMergeIfAllDiscussionsAreResolved"] = state ? state.onlyAllowMergeIfAllDiscussionsAreResolved : undefined;
            inputs["onlyAllowMergeIfPipelineSucceeds"] = state ? state.onlyAllowMergeIfPipelineSucceeds : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["runnersToken"] = state ? state.runnersToken : undefined;
            inputs["sharedRunnersEnabled"] = state ? state.sharedRunnersEnabled : undefined;
            inputs["sharedWithGroups"] = state ? state.sharedWithGroups : undefined;
            inputs["snippetsEnabled"] = state ? state.snippetsEnabled : undefined;
            inputs["sshUrlToRepo"] = state ? state.sshUrlToRepo : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["visibilityLevel"] = state ? state.visibilityLevel : undefined;
            inputs["webUrl"] = state ? state.webUrl : undefined;
            inputs["wikiEnabled"] = state ? state.wikiEnabled : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            inputs["approvalsBeforeMerge"] = args ? args.approvalsBeforeMerge : undefined;
            inputs["archived"] = args ? args.archived : undefined;
            inputs["containerRegistryEnabled"] = args ? args.containerRegistryEnabled : undefined;
            inputs["defaultBranch"] = args ? args.defaultBranch : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["issuesEnabled"] = args ? args.issuesEnabled : undefined;
            inputs["mergeMethod"] = args ? args.mergeMethod : undefined;
            inputs["mergeRequestsEnabled"] = args ? args.mergeRequestsEnabled : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namespaceId"] = args ? args.namespaceId : undefined;
            inputs["onlyAllowMergeIfAllDiscussionsAreResolved"] = args ? args.onlyAllowMergeIfAllDiscussionsAreResolved : undefined;
            inputs["onlyAllowMergeIfPipelineSucceeds"] = args ? args.onlyAllowMergeIfPipelineSucceeds : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["sharedRunnersEnabled"] = args ? args.sharedRunnersEnabled : undefined;
            inputs["sharedWithGroups"] = args ? args.sharedWithGroups : undefined;
            inputs["snippetsEnabled"] = args ? args.snippetsEnabled : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["visibilityLevel"] = args ? args.visibilityLevel : undefined;
            inputs["wikiEnabled"] = args ? args.wikiEnabled : undefined;
            inputs["httpUrlToRepo"] = undefined /*out*/;
            inputs["runnersToken"] = undefined /*out*/;
            inputs["sshUrlToRepo"] = undefined /*out*/;
            inputs["webUrl"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Project.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * Number of merge request approvals required for merging. Default is 0.
     */
    readonly approvalsBeforeMerge?: pulumi.Input<number>;
    /**
     * Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
     */
    readonly archived?: pulumi.Input<boolean>;
    /**
     * Enable container registry for the project.
     */
    readonly containerRegistryEnabled?: pulumi.Input<boolean>;
    /**
     * The default branch for the project.
     */
    readonly defaultBranch?: pulumi.Input<string>;
    /**
     * A description of the project.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * URL that can be provided to `git clone` to clone the
     * repository via HTTP.
     */
    readonly httpUrlToRepo?: pulumi.Input<string>;
    /**
     * Enable issue tracking for the project.
     */
    readonly issuesEnabled?: pulumi.Input<boolean>;
    /**
     * Set to `ff` to create fast-forward merges
     * Valid values are `merge`, `rebaseMerge`, `ff`
     * Repositories are created with `merge` by default
     */
    readonly mergeMethod?: pulumi.Input<string>;
    /**
     * Enable merge requests for the project.
     */
    readonly mergeRequestsEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the project.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The namespace (group or user) of the project. Defaults to your user.
     * See `gitlab..Group` for an example.
     */
    readonly namespaceId?: pulumi.Input<number>;
    /**
     * Set to true if you want allow merges only if all discussions are resolved.
     */
    readonly onlyAllowMergeIfAllDiscussionsAreResolved?: pulumi.Input<boolean>;
    /**
     * Set to true if you want allow merges only if a pipeline succeeds.
     */
    readonly onlyAllowMergeIfPipelineSucceeds?: pulumi.Input<boolean>;
    /**
     * The path of the repository.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * Registration token to use during runner setup.
     */
    readonly runnersToken?: pulumi.Input<string>;
    /**
     * Enable shared runners for this project.
     */
    readonly sharedRunnersEnabled?: pulumi.Input<boolean>;
    /**
     * Enable sharing the project with a list of groups (maps).
     */
    readonly sharedWithGroups?: pulumi.Input<pulumi.Input<{ groupAccessLevel: pulumi.Input<string>, groupId: pulumi.Input<number>, groupName?: pulumi.Input<string> }>[]>;
    /**
     * Enable snippets for the project.
     */
    readonly snippetsEnabled?: pulumi.Input<boolean>;
    /**
     * URL that can be provided to `git clone` to clone the
     * repository via SSH.
     */
    readonly sshUrlToRepo?: pulumi.Input<string>;
    /**
     * Tags (topics) of the project.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set to `public` to create a public project.
     * Valid values are `private`, `internal`, `public`.
     * Repositories are created as private by default.
     */
    readonly visibilityLevel?: pulumi.Input<string>;
    /**
     * URL that can be used to find the project in a browser.
     */
    readonly webUrl?: pulumi.Input<string>;
    /**
     * Enable wiki for the project.
     */
    readonly wikiEnabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * Number of merge request approvals required for merging. Default is 0.
     */
    readonly approvalsBeforeMerge?: pulumi.Input<number>;
    /**
     * Whether the project is in read-only mode (archived). Repositories can be archived/unarchived by toggling this parameter.
     */
    readonly archived?: pulumi.Input<boolean>;
    /**
     * Enable container registry for the project.
     */
    readonly containerRegistryEnabled?: pulumi.Input<boolean>;
    /**
     * The default branch for the project.
     */
    readonly defaultBranch?: pulumi.Input<string>;
    /**
     * A description of the project.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Enable issue tracking for the project.
     */
    readonly issuesEnabled?: pulumi.Input<boolean>;
    /**
     * Set to `ff` to create fast-forward merges
     * Valid values are `merge`, `rebaseMerge`, `ff`
     * Repositories are created with `merge` by default
     */
    readonly mergeMethod?: pulumi.Input<string>;
    /**
     * Enable merge requests for the project.
     */
    readonly mergeRequestsEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the project.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The namespace (group or user) of the project. Defaults to your user.
     * See `gitlab..Group` for an example.
     */
    readonly namespaceId?: pulumi.Input<number>;
    /**
     * Set to true if you want allow merges only if all discussions are resolved.
     */
    readonly onlyAllowMergeIfAllDiscussionsAreResolved?: pulumi.Input<boolean>;
    /**
     * Set to true if you want allow merges only if a pipeline succeeds.
     */
    readonly onlyAllowMergeIfPipelineSucceeds?: pulumi.Input<boolean>;
    /**
     * The path of the repository.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * Enable shared runners for this project.
     */
    readonly sharedRunnersEnabled?: pulumi.Input<boolean>;
    /**
     * Enable sharing the project with a list of groups (maps).
     */
    readonly sharedWithGroups?: pulumi.Input<pulumi.Input<{ groupAccessLevel: pulumi.Input<string>, groupId: pulumi.Input<number>, groupName?: pulumi.Input<string> }>[]>;
    /**
     * Enable snippets for the project.
     */
    readonly snippetsEnabled?: pulumi.Input<boolean>;
    /**
     * Tags (topics) of the project.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set to `public` to create a public project.
     * Valid values are `private`, `internal`, `public`.
     * Repositories are created as private by default.
     */
    readonly visibilityLevel?: pulumi.Input<string>;
    /**
     * Enable wiki for the project.
     */
    readonly wikiEnabled?: pulumi.Input<boolean>;
}
