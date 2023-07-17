// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage ecs deployment set
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultDeploymentSet = new volcengine.ecs.DeploymentSet("default", {
 *     deploymentSetName: "tf-test",
 *     description: "test1",
 *     granularity: "host",
 * });
 * ```
 *
 * ## Import
 *
 * ECS deployment set can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:ecs/deploymentSet:DeploymentSet default i-mizl7m1kqccg5smt1bdpijuj
 * ```
 */
export class DeploymentSet extends pulumi.CustomResource {
    /**
     * Get an existing DeploymentSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeploymentSetState, opts?: pulumi.CustomResourceOptions): DeploymentSet {
        return new DeploymentSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:ecs/deploymentSet:DeploymentSet';

    /**
     * Returns true if the given object is an instance of DeploymentSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeploymentSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeploymentSet.__pulumiType;
    }

    /**
     * The ID of ECS DeploymentSet.
     */
    public /*out*/ readonly deploymentSetId!: pulumi.Output<string>;
    /**
     * The name of ECS DeploymentSet.
     */
    public readonly deploymentSetName!: pulumi.Output<string>;
    /**
     * The description of ECS DeploymentSet.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The granularity of ECS DeploymentSet.Valid values: switch, host, rack,Default is host.
     */
    public readonly granularity!: pulumi.Output<string | undefined>;
    /**
     * The strategy of ECS DeploymentSet.Valid values: Availability.Default is Availability.
     */
    public readonly strategy!: pulumi.Output<string | undefined>;

    /**
     * Create a DeploymentSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeploymentSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeploymentSetArgs | DeploymentSetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeploymentSetState | undefined;
            resourceInputs["deploymentSetId"] = state ? state.deploymentSetId : undefined;
            resourceInputs["deploymentSetName"] = state ? state.deploymentSetName : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["granularity"] = state ? state.granularity : undefined;
            resourceInputs["strategy"] = state ? state.strategy : undefined;
        } else {
            const args = argsOrState as DeploymentSetArgs | undefined;
            if ((!args || args.deploymentSetName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deploymentSetName'");
            }
            resourceInputs["deploymentSetName"] = args ? args.deploymentSetName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["granularity"] = args ? args.granularity : undefined;
            resourceInputs["strategy"] = args ? args.strategy : undefined;
            resourceInputs["deploymentSetId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DeploymentSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DeploymentSet resources.
 */
export interface DeploymentSetState {
    /**
     * The ID of ECS DeploymentSet.
     */
    deploymentSetId?: pulumi.Input<string>;
    /**
     * The name of ECS DeploymentSet.
     */
    deploymentSetName?: pulumi.Input<string>;
    /**
     * The description of ECS DeploymentSet.
     */
    description?: pulumi.Input<string>;
    /**
     * The granularity of ECS DeploymentSet.Valid values: switch, host, rack,Default is host.
     */
    granularity?: pulumi.Input<string>;
    /**
     * The strategy of ECS DeploymentSet.Valid values: Availability.Default is Availability.
     */
    strategy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DeploymentSet resource.
 */
export interface DeploymentSetArgs {
    /**
     * The name of ECS DeploymentSet.
     */
    deploymentSetName: pulumi.Input<string>;
    /**
     * The description of ECS DeploymentSet.
     */
    description?: pulumi.Input<string>;
    /**
     * The granularity of ECS DeploymentSet.Valid values: switch, host, rack,Default is host.
     */
    granularity?: pulumi.Input<string>;
    /**
     * The strategy of ECS DeploymentSet.Valid values: Availability.Default is Availability.
     */
    strategy?: pulumi.Input<string>;
}
