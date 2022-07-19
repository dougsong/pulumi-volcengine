// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage vke node pool
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const vkeTest = new volcengine.Vke.NodePool("vke_test", {
 *     clusterId: "ccah01nnqtofnluts98j0",
 *     kubernetesConfig: {
 *         labels: [
 *             {
 *                 key: "aa",
 *                 value: "bb",
 *             },
 *             {
 *                 key: "cccc",
 *                 value: "dddd",
 *             },
 *         ],
 *     },
 *     nodeConfig: {
 *         dataVolumes: [{
 *             size: 60,
 *             type: "ESSD_PL0",
 *         }],
 *         instanceTypeIds: ["ecs.r1.large"],
 *         security: {
 *             login: {
 *                 //      ssh_key_pair_name = "ssh-6fbl66fxqm"
 *                 password: "UHdkMTIzNDU2",
 *             },
 *         },
 *         subnetIds: ["subnet-3recgzi7hfim85zsk2i8l9ve7"],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * NodePool can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:Vke/nodePool:NodePool default pcabe57vqtofgrbln3dp0
 * ```
 */
export class NodePool extends pulumi.CustomResource {
    /**
     * Get an existing NodePool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NodePoolState, opts?: pulumi.CustomResourceOptions): NodePool {
        return new NodePool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:Vke/nodePool:NodePool';

    /**
     * Returns true if the given object is an instance of NodePool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NodePool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NodePool.__pulumiType;
    }

    /**
     * The node pool elastic scaling configuration information.
     */
    public readonly autoScaling!: pulumi.Output<outputs.Vke.NodePoolAutoScaling | undefined>;
    /**
     * Is enabled of AutoScaling.
     */
    public readonly autoScalingEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The ClientToken of NodePool.
     */
    public readonly clientToken!: pulumi.Output<string | undefined>;
    /**
     * The ClusterId of NodePool.
     */
    public readonly clusterId!: pulumi.Output<string | undefined>;
    /**
     * The ClusterIds of NodePool.
     */
    public readonly clusterIds!: pulumi.Output<string[] | undefined>;
    /**
     * The CreateClientToken of NodePool.
     */
    public /*out*/ readonly createClientToken!: pulumi.Output<string>;
    /**
     * The IDs of NodePool.
     */
    public readonly ids!: pulumi.Output<string[] | undefined>;
    /**
     * The KubernetesConfig of NodeConfig.
     */
    public readonly kubernetesConfig!: pulumi.Output<outputs.Vke.NodePoolKubernetesConfig | undefined>;
    /**
     * The Name of NodePool.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Config of NodePool.
     */
    public readonly nodeConfig!: pulumi.Output<outputs.Vke.NodePoolNodeConfig | undefined>;
    /**
     * The Status of NodePool.
     */
    public readonly statuses!: pulumi.Output<outputs.Vke.NodePoolStatus[] | undefined>;
    /**
     * The UpdateClientToken of NodePool.
     */
    public /*out*/ readonly updateClientToken!: pulumi.Output<string>;

    /**
     * Create a NodePool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NodePoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NodePoolArgs | NodePoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NodePoolState | undefined;
            resourceInputs["autoScaling"] = state ? state.autoScaling : undefined;
            resourceInputs["autoScalingEnabled"] = state ? state.autoScalingEnabled : undefined;
            resourceInputs["clientToken"] = state ? state.clientToken : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["clusterIds"] = state ? state.clusterIds : undefined;
            resourceInputs["createClientToken"] = state ? state.createClientToken : undefined;
            resourceInputs["ids"] = state ? state.ids : undefined;
            resourceInputs["kubernetesConfig"] = state ? state.kubernetesConfig : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodeConfig"] = state ? state.nodeConfig : undefined;
            resourceInputs["statuses"] = state ? state.statuses : undefined;
            resourceInputs["updateClientToken"] = state ? state.updateClientToken : undefined;
        } else {
            const args = argsOrState as NodePoolArgs | undefined;
            resourceInputs["autoScaling"] = args ? args.autoScaling : undefined;
            resourceInputs["autoScalingEnabled"] = args ? args.autoScalingEnabled : undefined;
            resourceInputs["clientToken"] = args ? args.clientToken : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["clusterIds"] = args ? args.clusterIds : undefined;
            resourceInputs["ids"] = args ? args.ids : undefined;
            resourceInputs["kubernetesConfig"] = args ? args.kubernetesConfig : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodeConfig"] = args ? args.nodeConfig : undefined;
            resourceInputs["statuses"] = args ? args.statuses : undefined;
            resourceInputs["createClientToken"] = undefined /*out*/;
            resourceInputs["updateClientToken"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NodePool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NodePool resources.
 */
export interface NodePoolState {
    /**
     * The node pool elastic scaling configuration information.
     */
    autoScaling?: pulumi.Input<inputs.Vke.NodePoolAutoScaling>;
    /**
     * Is enabled of AutoScaling.
     */
    autoScalingEnabled?: pulumi.Input<boolean>;
    /**
     * The ClientToken of NodePool.
     */
    clientToken?: pulumi.Input<string>;
    /**
     * The ClusterId of NodePool.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The ClusterIds of NodePool.
     */
    clusterIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The CreateClientToken of NodePool.
     */
    createClientToken?: pulumi.Input<string>;
    /**
     * The IDs of NodePool.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The KubernetesConfig of NodeConfig.
     */
    kubernetesConfig?: pulumi.Input<inputs.Vke.NodePoolKubernetesConfig>;
    /**
     * The Name of NodePool.
     */
    name?: pulumi.Input<string>;
    /**
     * The Config of NodePool.
     */
    nodeConfig?: pulumi.Input<inputs.Vke.NodePoolNodeConfig>;
    /**
     * The Status of NodePool.
     */
    statuses?: pulumi.Input<pulumi.Input<inputs.Vke.NodePoolStatus>[]>;
    /**
     * The UpdateClientToken of NodePool.
     */
    updateClientToken?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NodePool resource.
 */
export interface NodePoolArgs {
    /**
     * The node pool elastic scaling configuration information.
     */
    autoScaling?: pulumi.Input<inputs.Vke.NodePoolAutoScaling>;
    /**
     * Is enabled of AutoScaling.
     */
    autoScalingEnabled?: pulumi.Input<boolean>;
    /**
     * The ClientToken of NodePool.
     */
    clientToken?: pulumi.Input<string>;
    /**
     * The ClusterId of NodePool.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The ClusterIds of NodePool.
     */
    clusterIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The IDs of NodePool.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The KubernetesConfig of NodeConfig.
     */
    kubernetesConfig?: pulumi.Input<inputs.Vke.NodePoolKubernetesConfig>;
    /**
     * The Name of NodePool.
     */
    name?: pulumi.Input<string>;
    /**
     * The Config of NodePool.
     */
    nodeConfig?: pulumi.Input<inputs.Vke.NodePoolNodeConfig>;
    /**
     * The Status of NodePool.
     */
    statuses?: pulumi.Input<pulumi.Input<inputs.Vke.NodePoolStatus>[]>;
}
