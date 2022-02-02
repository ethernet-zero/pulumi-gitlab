// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## # gitlab\_pipeline\_schedule
 *
 * This resource allows you to create and manage pipeline schedules.
 * For further information on clusters, consult the [gitlab
 * documentation](https://docs.gitlab.com/ce/user/project/pipelines/schedules.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = new gitlab.PipelineSchedule("example", {
 *     cron: "0 1 * * *",
 *     description: "Used to schedule builds",
 *     project: "12345",
 *     ref: "master",
 * });
 * ```
 *
 * ## Import
 *
 * GitLab pipeline schedules can be imported using an id made up of `{project_id}:{pipeline_schedule_id}`, e.g.
 *
 * ```sh
 *  $ pulumi import gitlab:index/pipelineSchedule:PipelineSchedule test 1:3
 * ```
 */
export class PipelineSchedule extends pulumi.CustomResource {
    /**
     * Get an existing PipelineSchedule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PipelineScheduleState, opts?: pulumi.CustomResourceOptions): PipelineSchedule {
        return new PipelineSchedule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/pipelineSchedule:PipelineSchedule';

    /**
     * Returns true if the given object is an instance of PipelineSchedule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PipelineSchedule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PipelineSchedule.__pulumiType;
    }

    /**
     * The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
     */
    public readonly active!: pulumi.Output<boolean | undefined>;
    /**
     * The cron (e.g. `0 1 * * *`).
     */
    public readonly cron!: pulumi.Output<string>;
    /**
     * The timezone.
     */
    public readonly cronTimezone!: pulumi.Output<string | undefined>;
    /**
     * The description of the pipeline schedule.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The name or id of the project to add the schedule to.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The branch/tag name to be triggered.
     */
    public readonly ref!: pulumi.Output<string>;

    /**
     * Create a PipelineSchedule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PipelineScheduleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PipelineScheduleArgs | PipelineScheduleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PipelineScheduleState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["cron"] = state ? state.cron : undefined;
            resourceInputs["cronTimezone"] = state ? state.cronTimezone : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["ref"] = state ? state.ref : undefined;
        } else {
            const args = argsOrState as PipelineScheduleArgs | undefined;
            if ((!args || args.cron === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cron'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.ref === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ref'");
            }
            resourceInputs["active"] = args ? args.active : undefined;
            resourceInputs["cron"] = args ? args.cron : undefined;
            resourceInputs["cronTimezone"] = args ? args.cronTimezone : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["ref"] = args ? args.ref : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PipelineSchedule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PipelineSchedule resources.
 */
export interface PipelineScheduleState {
    /**
     * The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
     */
    active?: pulumi.Input<boolean>;
    /**
     * The cron (e.g. `0 1 * * *`).
     */
    cron?: pulumi.Input<string>;
    /**
     * The timezone.
     */
    cronTimezone?: pulumi.Input<string>;
    /**
     * The description of the pipeline schedule.
     */
    description?: pulumi.Input<string>;
    /**
     * The name or id of the project to add the schedule to.
     */
    project?: pulumi.Input<string>;
    /**
     * The branch/tag name to be triggered.
     */
    ref?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PipelineSchedule resource.
 */
export interface PipelineScheduleArgs {
    /**
     * The activation of pipeline schedule. If false is set, the pipeline schedule will deactivated initially.
     */
    active?: pulumi.Input<boolean>;
    /**
     * The cron (e.g. `0 1 * * *`).
     */
    cron: pulumi.Input<string>;
    /**
     * The timezone.
     */
    cronTimezone?: pulumi.Input<string>;
    /**
     * The description of the pipeline schedule.
     */
    description: pulumi.Input<string>;
    /**
     * The name or id of the project to add the schedule to.
     */
    project: pulumi.Input<string>;
    /**
     * The branch/tag name to be triggered.
     */
    ref: pulumi.Input<string>;
}
