// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage bioos cluster
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.bioos.Cluster("foo", {
 *     description: "test-description",
 *     sharedConfigs: [{
 *         enable: true,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Cluster can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:bioos/cluster:Cluster default *****
 * ```
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:bioos/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * The id of the vke cluster.
     */
    public /*out*/ readonly clusterId!: pulumi.Output<string>;
    /**
     * The description of the cluster.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the cluster.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The configuration of the shared cluster.
     */
    public readonly sharedConfigs!: pulumi.Output<outputs.bioos.ClusterSharedConfig[] | undefined>;
    /**
     * The configuration of the vke cluster.
     */
    public readonly vkeConfigs!: pulumi.Output<outputs.bioos.ClusterVkeConfig[] | undefined>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["sharedConfigs"] = state ? state.sharedConfigs : undefined;
            resourceInputs["vkeConfigs"] = state ? state.vkeConfigs : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["sharedConfigs"] = args ? args.sharedConfigs : undefined;
            resourceInputs["vkeConfigs"] = args ? args.vkeConfigs : undefined;
            resourceInputs["clusterId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * The id of the vke cluster.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The description of the cluster.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * The configuration of the shared cluster.
     */
    sharedConfigs?: pulumi.Input<pulumi.Input<inputs.bioos.ClusterSharedConfig>[]>;
    /**
     * The configuration of the vke cluster.
     */
    vkeConfigs?: pulumi.Input<pulumi.Input<inputs.bioos.ClusterVkeConfig>[]>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * The description of the cluster.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * The configuration of the shared cluster.
     */
    sharedConfigs?: pulumi.Input<pulumi.Input<inputs.bioos.ClusterSharedConfig>[]>;
    /**
     * The configuration of the vke cluster.
     */
    vkeConfigs?: pulumi.Input<pulumi.Input<inputs.bioos.ClusterVkeConfig>[]>;
}
