// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage hooks for your GitLab projects.
 * For further information on hooks, consult the [gitlab
 * documentation](https://docs.gitlab.com/ce/user/project/integrations/webhooks.html).
 * 
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 * 
 * const example = new gitlab.ProjectHook("example", {
 *     mergeRequestsEvents: true,
 *     project: "example/hooked",
 *     url: "https://example.com/hook/example",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/project_hook.html.markdown.
 */
export class ProjectHook extends pulumi.CustomResource {
    /**
     * Get an existing ProjectHook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectHookState, opts?: pulumi.CustomResourceOptions): ProjectHook {
        return new ProjectHook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/projectHook:ProjectHook';

    /**
     * Returns true if the given object is an instance of ProjectHook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectHook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectHook.__pulumiType;
    }

    /**
     * Enable ssl verification when invoking
     * the hook.
     */
    public readonly enableSslVerification!: pulumi.Output<boolean | undefined>;
    /**
     * Invoke the hook for issues events.
     */
    public readonly issuesEvents!: pulumi.Output<boolean | undefined>;
    /**
     * Invoke the hook for job events.
     */
    public readonly jobEvents!: pulumi.Output<boolean | undefined>;
    /**
     * Invoke the hook for merge requests.
     */
    public readonly mergeRequestsEvents!: pulumi.Output<boolean | undefined>;
    /**
     * Invoke the hook for notes events.
     */
    public readonly noteEvents!: pulumi.Output<boolean | undefined>;
    /**
     * Invoke the hook for pipeline events.
     */
    public readonly pipelineEvents!: pulumi.Output<boolean | undefined>;
    /**
     * The name or id of the project to add the hook to.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Invoke the hook for push events.
     */
    public readonly pushEvents!: pulumi.Output<boolean | undefined>;
    /**
     * Invoke the hook for tag push events.
     */
    public readonly tagPushEvents!: pulumi.Output<boolean | undefined>;
    /**
     * A token to present when invoking the hook.
     */
    public readonly token!: pulumi.Output<string | undefined>;
    /**
     * The url of the hook to invoke.
     */
    public readonly url!: pulumi.Output<string>;
    /**
     * Invoke the hook for wiki page events.
     */
    public readonly wikiPageEvents!: pulumi.Output<boolean | undefined>;

    /**
     * Create a ProjectHook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectHookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectHookArgs | ProjectHookState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProjectHookState | undefined;
            inputs["enableSslVerification"] = state ? state.enableSslVerification : undefined;
            inputs["issuesEvents"] = state ? state.issuesEvents : undefined;
            inputs["jobEvents"] = state ? state.jobEvents : undefined;
            inputs["mergeRequestsEvents"] = state ? state.mergeRequestsEvents : undefined;
            inputs["noteEvents"] = state ? state.noteEvents : undefined;
            inputs["pipelineEvents"] = state ? state.pipelineEvents : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["pushEvents"] = state ? state.pushEvents : undefined;
            inputs["tagPushEvents"] = state ? state.tagPushEvents : undefined;
            inputs["token"] = state ? state.token : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["wikiPageEvents"] = state ? state.wikiPageEvents : undefined;
        } else {
            const args = argsOrState as ProjectHookArgs | undefined;
            if (!args || args.project === undefined) {
                throw new Error("Missing required property 'project'");
            }
            if (!args || args.url === undefined) {
                throw new Error("Missing required property 'url'");
            }
            inputs["enableSslVerification"] = args ? args.enableSslVerification : undefined;
            inputs["issuesEvents"] = args ? args.issuesEvents : undefined;
            inputs["jobEvents"] = args ? args.jobEvents : undefined;
            inputs["mergeRequestsEvents"] = args ? args.mergeRequestsEvents : undefined;
            inputs["noteEvents"] = args ? args.noteEvents : undefined;
            inputs["pipelineEvents"] = args ? args.pipelineEvents : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["pushEvents"] = args ? args.pushEvents : undefined;
            inputs["tagPushEvents"] = args ? args.tagPushEvents : undefined;
            inputs["token"] = args ? args.token : undefined;
            inputs["url"] = args ? args.url : undefined;
            inputs["wikiPageEvents"] = args ? args.wikiPageEvents : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ProjectHook.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectHook resources.
 */
export interface ProjectHookState {
    /**
     * Enable ssl verification when invoking
     * the hook.
     */
    readonly enableSslVerification?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for issues events.
     */
    readonly issuesEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for job events.
     */
    readonly jobEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for merge requests.
     */
    readonly mergeRequestsEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for notes events.
     */
    readonly noteEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for pipeline events.
     */
    readonly pipelineEvents?: pulumi.Input<boolean>;
    /**
     * The name or id of the project to add the hook to.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Invoke the hook for push events.
     */
    readonly pushEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for tag push events.
     */
    readonly tagPushEvents?: pulumi.Input<boolean>;
    /**
     * A token to present when invoking the hook.
     */
    readonly token?: pulumi.Input<string>;
    /**
     * The url of the hook to invoke.
     */
    readonly url?: pulumi.Input<string>;
    /**
     * Invoke the hook for wiki page events.
     */
    readonly wikiPageEvents?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ProjectHook resource.
 */
export interface ProjectHookArgs {
    /**
     * Enable ssl verification when invoking
     * the hook.
     */
    readonly enableSslVerification?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for issues events.
     */
    readonly issuesEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for job events.
     */
    readonly jobEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for merge requests.
     */
    readonly mergeRequestsEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for notes events.
     */
    readonly noteEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for pipeline events.
     */
    readonly pipelineEvents?: pulumi.Input<boolean>;
    /**
     * The name or id of the project to add the hook to.
     */
    readonly project: pulumi.Input<string>;
    /**
     * Invoke the hook for push events.
     */
    readonly pushEvents?: pulumi.Input<boolean>;
    /**
     * Invoke the hook for tag push events.
     */
    readonly tagPushEvents?: pulumi.Input<boolean>;
    /**
     * A token to present when invoking the hook.
     */
    readonly token?: pulumi.Input<string>;
    /**
     * The url of the hook to invoke.
     */
    readonly url: pulumi.Input<string>;
    /**
     * Invoke the hook for wiki page events.
     */
    readonly wikiPageEvents?: pulumi.Input<boolean>;
}
