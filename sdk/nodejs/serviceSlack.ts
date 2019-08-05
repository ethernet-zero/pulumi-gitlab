// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to manage Slack notifications integration.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 * 
 * const awesomeProject = new gitlab.Project("awesome_project", {
 *     description: "My awesome project.",
 *     visibilityLevel: "public",
 * });
 * const slack = new gitlab.ServiceSlack("slack", {
 *     project: awesomeProject.id,
 *     pushChannel: "push_chan",
 *     pushEvents: true,
 *     username: "myuser",
 *     webhook: "https://webhook.com",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/r/service_slack.html.markdown.
 */
export class ServiceSlack extends pulumi.CustomResource {
    /**
     * Get an existing ServiceSlack resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceSlackState, opts?: pulumi.CustomResourceOptions): ServiceSlack {
        return new ServiceSlack(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/serviceSlack:ServiceSlack';

    /**
     * Returns true if the given object is an instance of ServiceSlack.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceSlack {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceSlack.__pulumiType;
    }

    /**
     * The name of the channel to receive confidential issue events notifications.
     */
    public readonly confidentialIssueChannel!: pulumi.Output<string | undefined>;
    /**
     * Enable notifications for confidential issues events.
     */
    public readonly confidentialIssuesEvents!: pulumi.Output<boolean>;
    /**
     * Enable notifications for confidential note events.
     */
    public readonly confidentialNoteEvents!: pulumi.Output<boolean>;
    /**
     * The name of the channel to receive issue events notifications.
     */
    public readonly issueChannel!: pulumi.Output<string | undefined>;
    /**
     * Enable notifications for issues events.
     */
    public readonly issuesEvents!: pulumi.Output<boolean>;
    public /*out*/ readonly jobEvents!: pulumi.Output<boolean>;
    /**
     * The name of the channel to receive merge request events notifications.
     */
    public readonly mergeRequestChannel!: pulumi.Output<string | undefined>;
    /**
     * Enable notifications for merge requests events.
     */
    public readonly mergeRequestsEvents!: pulumi.Output<boolean>;
    /**
     * The name of the channel to receive note events notifications.
     */
    public readonly noteChannel!: pulumi.Output<string | undefined>;
    /**
     * Enable notifications for note events.
     */
    public readonly noteEvents!: pulumi.Output<boolean>;
    /**
     * Send notifications for broken pipelines.
     */
    public readonly notifyOnlyBrokenPipelines!: pulumi.Output<boolean>;
    /**
     * Send notifications only for the default branch.
     */
    public readonly notifyOnlyDefaultBranch!: pulumi.Output<boolean>;
    /**
     * The name of the channel to receive pipeline events notifications.
     */
    public readonly pipelineChannel!: pulumi.Output<string | undefined>;
    /**
     * Enable notifications for pipeline events.
     */
    public readonly pipelineEvents!: pulumi.Output<boolean>;
    /**
     * ID of the project you want to activate integration on.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The name of the channel to receive push events notifications.
     */
    public readonly pushChannel!: pulumi.Output<string | undefined>;
    /**
     * Enable notifications for push events.
     */
    public readonly pushEvents!: pulumi.Output<boolean>;
    /**
     * The name of the channel to receive tag push events notifications.
     */
    public readonly tagPushChannel!: pulumi.Output<string | undefined>;
    /**
     * Enable notifications for tag push events.
     */
    public readonly tagPushEvents!: pulumi.Output<boolean>;
    /**
     * Username to use.
     */
    public readonly username!: pulumi.Output<string | undefined>;
    /**
     * Webhook URL (ex.: https://hooks.slack.com/services/...)
     */
    public readonly webhook!: pulumi.Output<string>;
    /**
     * The name of the channel to receive wiki page events notifications.
     */
    public readonly wikiPageChannel!: pulumi.Output<string | undefined>;
    /**
     * Enable notifications for wiki page events.
     */
    public readonly wikiPageEvents!: pulumi.Output<boolean>;

    /**
     * Create a ServiceSlack resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceSlackArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceSlackArgs | ServiceSlackState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServiceSlackState | undefined;
            inputs["confidentialIssueChannel"] = state ? state.confidentialIssueChannel : undefined;
            inputs["confidentialIssuesEvents"] = state ? state.confidentialIssuesEvents : undefined;
            inputs["confidentialNoteEvents"] = state ? state.confidentialNoteEvents : undefined;
            inputs["issueChannel"] = state ? state.issueChannel : undefined;
            inputs["issuesEvents"] = state ? state.issuesEvents : undefined;
            inputs["jobEvents"] = state ? state.jobEvents : undefined;
            inputs["mergeRequestChannel"] = state ? state.mergeRequestChannel : undefined;
            inputs["mergeRequestsEvents"] = state ? state.mergeRequestsEvents : undefined;
            inputs["noteChannel"] = state ? state.noteChannel : undefined;
            inputs["noteEvents"] = state ? state.noteEvents : undefined;
            inputs["notifyOnlyBrokenPipelines"] = state ? state.notifyOnlyBrokenPipelines : undefined;
            inputs["notifyOnlyDefaultBranch"] = state ? state.notifyOnlyDefaultBranch : undefined;
            inputs["pipelineChannel"] = state ? state.pipelineChannel : undefined;
            inputs["pipelineEvents"] = state ? state.pipelineEvents : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["pushChannel"] = state ? state.pushChannel : undefined;
            inputs["pushEvents"] = state ? state.pushEvents : undefined;
            inputs["tagPushChannel"] = state ? state.tagPushChannel : undefined;
            inputs["tagPushEvents"] = state ? state.tagPushEvents : undefined;
            inputs["username"] = state ? state.username : undefined;
            inputs["webhook"] = state ? state.webhook : undefined;
            inputs["wikiPageChannel"] = state ? state.wikiPageChannel : undefined;
            inputs["wikiPageEvents"] = state ? state.wikiPageEvents : undefined;
        } else {
            const args = argsOrState as ServiceSlackArgs | undefined;
            if (!args || args.project === undefined) {
                throw new Error("Missing required property 'project'");
            }
            if (!args || args.webhook === undefined) {
                throw new Error("Missing required property 'webhook'");
            }
            inputs["confidentialIssueChannel"] = args ? args.confidentialIssueChannel : undefined;
            inputs["confidentialIssuesEvents"] = args ? args.confidentialIssuesEvents : undefined;
            inputs["confidentialNoteEvents"] = args ? args.confidentialNoteEvents : undefined;
            inputs["issueChannel"] = args ? args.issueChannel : undefined;
            inputs["issuesEvents"] = args ? args.issuesEvents : undefined;
            inputs["mergeRequestChannel"] = args ? args.mergeRequestChannel : undefined;
            inputs["mergeRequestsEvents"] = args ? args.mergeRequestsEvents : undefined;
            inputs["noteChannel"] = args ? args.noteChannel : undefined;
            inputs["noteEvents"] = args ? args.noteEvents : undefined;
            inputs["notifyOnlyBrokenPipelines"] = args ? args.notifyOnlyBrokenPipelines : undefined;
            inputs["notifyOnlyDefaultBranch"] = args ? args.notifyOnlyDefaultBranch : undefined;
            inputs["pipelineChannel"] = args ? args.pipelineChannel : undefined;
            inputs["pipelineEvents"] = args ? args.pipelineEvents : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["pushChannel"] = args ? args.pushChannel : undefined;
            inputs["pushEvents"] = args ? args.pushEvents : undefined;
            inputs["tagPushChannel"] = args ? args.tagPushChannel : undefined;
            inputs["tagPushEvents"] = args ? args.tagPushEvents : undefined;
            inputs["username"] = args ? args.username : undefined;
            inputs["webhook"] = args ? args.webhook : undefined;
            inputs["wikiPageChannel"] = args ? args.wikiPageChannel : undefined;
            inputs["wikiPageEvents"] = args ? args.wikiPageEvents : undefined;
            inputs["jobEvents"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ServiceSlack.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceSlack resources.
 */
export interface ServiceSlackState {
    /**
     * The name of the channel to receive confidential issue events notifications.
     */
    readonly confidentialIssueChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for confidential issues events.
     */
    readonly confidentialIssuesEvents?: pulumi.Input<boolean>;
    /**
     * Enable notifications for confidential note events.
     */
    readonly confidentialNoteEvents?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive issue events notifications.
     */
    readonly issueChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for issues events.
     */
    readonly issuesEvents?: pulumi.Input<boolean>;
    readonly jobEvents?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive merge request events notifications.
     */
    readonly mergeRequestChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for merge requests events.
     */
    readonly mergeRequestsEvents?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive note events notifications.
     */
    readonly noteChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for note events.
     */
    readonly noteEvents?: pulumi.Input<boolean>;
    /**
     * Send notifications for broken pipelines.
     */
    readonly notifyOnlyBrokenPipelines?: pulumi.Input<boolean>;
    /**
     * Send notifications only for the default branch.
     */
    readonly notifyOnlyDefaultBranch?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive pipeline events notifications.
     */
    readonly pipelineChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for pipeline events.
     */
    readonly pipelineEvents?: pulumi.Input<boolean>;
    /**
     * ID of the project you want to activate integration on.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The name of the channel to receive push events notifications.
     */
    readonly pushChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for push events.
     */
    readonly pushEvents?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive tag push events notifications.
     */
    readonly tagPushChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for tag push events.
     */
    readonly tagPushEvents?: pulumi.Input<boolean>;
    /**
     * Username to use.
     */
    readonly username?: pulumi.Input<string>;
    /**
     * Webhook URL (ex.: https://hooks.slack.com/services/...)
     */
    readonly webhook?: pulumi.Input<string>;
    /**
     * The name of the channel to receive wiki page events notifications.
     */
    readonly wikiPageChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for wiki page events.
     */
    readonly wikiPageEvents?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ServiceSlack resource.
 */
export interface ServiceSlackArgs {
    /**
     * The name of the channel to receive confidential issue events notifications.
     */
    readonly confidentialIssueChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for confidential issues events.
     */
    readonly confidentialIssuesEvents?: pulumi.Input<boolean>;
    /**
     * Enable notifications for confidential note events.
     */
    readonly confidentialNoteEvents?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive issue events notifications.
     */
    readonly issueChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for issues events.
     */
    readonly issuesEvents?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive merge request events notifications.
     */
    readonly mergeRequestChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for merge requests events.
     */
    readonly mergeRequestsEvents?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive note events notifications.
     */
    readonly noteChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for note events.
     */
    readonly noteEvents?: pulumi.Input<boolean>;
    /**
     * Send notifications for broken pipelines.
     */
    readonly notifyOnlyBrokenPipelines?: pulumi.Input<boolean>;
    /**
     * Send notifications only for the default branch.
     */
    readonly notifyOnlyDefaultBranch?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive pipeline events notifications.
     */
    readonly pipelineChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for pipeline events.
     */
    readonly pipelineEvents?: pulumi.Input<boolean>;
    /**
     * ID of the project you want to activate integration on.
     */
    readonly project: pulumi.Input<string>;
    /**
     * The name of the channel to receive push events notifications.
     */
    readonly pushChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for push events.
     */
    readonly pushEvents?: pulumi.Input<boolean>;
    /**
     * The name of the channel to receive tag push events notifications.
     */
    readonly tagPushChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for tag push events.
     */
    readonly tagPushEvents?: pulumi.Input<boolean>;
    /**
     * Username to use.
     */
    readonly username?: pulumi.Input<string>;
    /**
     * Webhook URL (ex.: https://hooks.slack.com/services/...)
     */
    readonly webhook: pulumi.Input<string>;
    /**
     * The name of the channel to receive wiki page events notifications.
     */
    readonly wikiPageChannel?: pulumi.Input<string>;
    /**
     * Enable notifications for wiki page events.
     */
    readonly wikiPageEvents?: pulumi.Input<boolean>;
}
