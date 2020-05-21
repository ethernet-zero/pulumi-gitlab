// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage CI/CD variables for your GitLab groups.
 * For further information on variables, consult the [gitlab
 * documentation](https://docs.gitlab.com/ce/ci/variables/README.html#variables).
 *
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gitlab from "@pulumi/gitlab";
 *
 * const example = new gitlab.GroupVariable("example", {
 *     group: "12345",
 *     key: "groupVariableKey",
 *     protected: false,
 *     value: "groupVariableValue",
 * });
 * ```
 */
export class GroupVariable extends pulumi.CustomResource {
    /**
     * Get an existing GroupVariable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupVariableState, opts?: pulumi.CustomResourceOptions): GroupVariable {
        return new GroupVariable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gitlab:index/groupVariable:GroupVariable';

    /**
     * Returns true if the given object is an instance of GroupVariable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupVariable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupVariable.__pulumiType;
    }

    /**
     * The name or id of the group to add the hook to.
     */
    public readonly group!: pulumi.Output<string>;
    /**
     * The name of the variable.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
     */
    public readonly protected!: pulumi.Output<boolean | undefined>;
    /**
     * The value of the variable.
     */
    public readonly value!: pulumi.Output<string>;
    /**
     * The type of a variable. Available types are: envVar (default) and file.
     */
    public readonly variableType!: pulumi.Output<string | undefined>;

    /**
     * Create a GroupVariable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupVariableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupVariableArgs | GroupVariableState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GroupVariableState | undefined;
            inputs["group"] = state ? state.group : undefined;
            inputs["key"] = state ? state.key : undefined;
            inputs["protected"] = state ? state.protected : undefined;
            inputs["value"] = state ? state.value : undefined;
            inputs["variableType"] = state ? state.variableType : undefined;
        } else {
            const args = argsOrState as GroupVariableArgs | undefined;
            if (!args || args.group === undefined) {
                throw new Error("Missing required property 'group'");
            }
            if (!args || args.key === undefined) {
                throw new Error("Missing required property 'key'");
            }
            if (!args || args.value === undefined) {
                throw new Error("Missing required property 'value'");
            }
            inputs["group"] = args ? args.group : undefined;
            inputs["key"] = args ? args.key : undefined;
            inputs["protected"] = args ? args.protected : undefined;
            inputs["value"] = args ? args.value : undefined;
            inputs["variableType"] = args ? args.variableType : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GroupVariable.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupVariable resources.
 */
export interface GroupVariableState {
    /**
     * The name or id of the group to add the hook to.
     */
    readonly group?: pulumi.Input<string>;
    /**
     * The name of the variable.
     */
    readonly key?: pulumi.Input<string>;
    /**
     * If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
     */
    readonly protected?: pulumi.Input<boolean>;
    /**
     * The value of the variable.
     */
    readonly value?: pulumi.Input<string>;
    /**
     * The type of a variable. Available types are: envVar (default) and file.
     */
    readonly variableType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupVariable resource.
 */
export interface GroupVariableArgs {
    /**
     * The name or id of the group to add the hook to.
     */
    readonly group: pulumi.Input<string>;
    /**
     * The name of the variable.
     */
    readonly key: pulumi.Input<string>;
    /**
     * If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.
     */
    readonly protected?: pulumi.Input<boolean>;
    /**
     * The value of the variable.
     */
    readonly value: pulumi.Input<string>;
    /**
     * The type of a variable. Available types are: envVar (default) and file.
     */
    readonly variableType?: pulumi.Input<string>;
}
