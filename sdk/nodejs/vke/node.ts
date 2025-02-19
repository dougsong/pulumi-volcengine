// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage vke node
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooVpc = new volcengine.vpc.Vpc("fooVpc", {
 *     vpcName: "acc-test-project1",
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const fooSubnet = new volcengine.vpc.Subnet("fooSubnet", {
 *     subnetName: "acc-subnet-test-2",
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: "cn-beijing-a",
 *     vpcId: fooVpc.id,
 * });
 * const fooSecurityGroup = new volcengine.vpc.SecurityGroup("fooSecurityGroup", {
 *     vpcId: fooVpc.id,
 *     securityGroupName: "acc-test-security-group2",
 * });
 * const fooInstance = new volcengine.ecs.Instance("fooInstance", {
 *     imageId: "image-ybqi99s7yq8rx7mnk44b",
 *     instanceType: "ecs.g1ie.large",
 *     instanceName: "acc-test-ecs-name2",
 *     password: "93f0cb0614Aab12",
 *     instanceChargeType: "PostPaid",
 *     systemVolumeType: "ESSD_PL0",
 *     systemVolumeSize: 40,
 *     subnetId: fooSubnet.id,
 *     securityGroupIds: [fooSecurityGroup.id],
 *     projectName: "default",
 * });
 * const fooCluster = new volcengine.vke.Cluster("fooCluster", {
 *     description: "created by terraform",
 *     deleteProtectionEnabled: false,
 *     clusterConfig: {
 *         subnetIds: [fooSubnet.id],
 *         apiServerPublicAccessEnabled: true,
 *         apiServerPublicAccessConfig: {
 *             publicAccessNetworkConfig: {
 *                 billingType: "PostPaidByBandwidth",
 *                 bandwidth: 1,
 *             },
 *         },
 *         resourcePublicAccessDefaultEnabled: true,
 *     },
 *     podsConfig: {
 *         podNetworkMode: "VpcCniShared",
 *         vpcCniConfig: {
 *             subnetIds: [fooSubnet.id],
 *         },
 *     },
 *     servicesConfig: {
 *         serviceCidrsv4s: ["172.30.0.0/18"],
 *     },
 *     tags: [{
 *         key: "tf-k1",
 *         value: "tf-v1",
 *     }],
 * });
 * const fooNodePool = new volcengine.vke.NodePool("fooNodePool", {
 *     clusterId: fooCluster.id,
 *     nodeConfig: {
 *         instanceTypeIds: ["ecs.g1ie.large"],
 *         subnetIds: [fooSubnet.id],
 *         security: {
 *             login: {
 *                 password: "UHdkMTIzNDU2",
 *             },
 *             securityGroupIds: [fooSecurityGroup.id],
 *         },
 *         instanceChargeType: "PostPaid",
 *         period: 1,
 *     },
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
 *         cordon: false,
 *     },
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * const fooNode = new volcengine.vke.Node("fooNode", {
 *     clusterId: fooCluster.id,
 *     instanceId: fooInstance.id,
 *     keepInstanceName: true,
 *     additionalContainerStorageEnabled: false,
 *     containerStoragePath: "",
 *     nodePoolId: fooNodePool.id,
 *     kubernetesConfig: {
 *         labels: [
 *             {
 *                 key: "tf-key1",
 *                 value: "tf-value1",
 *             },
 *             {
 *                 key: "tf-key2",
 *                 value: "tf-value2",
 *             },
 *         ],
 *         taints: [
 *             {
 *                 key: "tf-key3",
 *                 value: "tf-value3",
 *                 effect: "NoSchedule",
 *             },
 *             {
 *                 key: "tf-key4",
 *                 value: "tf-value4",
 *                 effect: "NoSchedule",
 *             },
 *         ],
 *         cordon: true,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * VKE node can be imported using the node id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:vke/node:Node default nc5t5epmrsf****
 * ```
 */
export class Node extends pulumi.CustomResource {
    /**
     * Get an existing Node resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NodeState, opts?: pulumi.CustomResourceOptions): Node {
        return new Node(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:vke/node:Node';

    /**
     * Returns true if the given object is an instance of Node.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Node {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Node.__pulumiType;
    }

    /**
     * The flag of additional container storage enable, the value is `true` or `false`.
     */
    public readonly additionalContainerStorageEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The client token.
     */
    public readonly clientToken!: pulumi.Output<string>;
    /**
     * The cluster id.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * The container storage path.
     */
    public readonly containerStoragePath!: pulumi.Output<string>;
    /**
     * The ImageId of NodeConfig.
     */
    public readonly imageId!: pulumi.Output<string>;
    /**
     * The initializeScript of Node.
     */
    public readonly initializeScript!: pulumi.Output<string>;
    /**
     * The instance id.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The flag of keep instance name, the value is `true` or `false`.
     */
    public readonly keepInstanceName!: pulumi.Output<boolean | undefined>;
    /**
     * The KubernetesConfig of Node.
     */
    public readonly kubernetesConfig!: pulumi.Output<outputs.vke.NodeKubernetesConfig>;
    /**
     * The node pool id.
     */
    public readonly nodePoolId!: pulumi.Output<string>;

    /**
     * Create a Node resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NodeArgs | NodeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NodeState | undefined;
            resourceInputs["additionalContainerStorageEnabled"] = state ? state.additionalContainerStorageEnabled : undefined;
            resourceInputs["clientToken"] = state ? state.clientToken : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["containerStoragePath"] = state ? state.containerStoragePath : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
            resourceInputs["initializeScript"] = state ? state.initializeScript : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["keepInstanceName"] = state ? state.keepInstanceName : undefined;
            resourceInputs["kubernetesConfig"] = state ? state.kubernetesConfig : undefined;
            resourceInputs["nodePoolId"] = state ? state.nodePoolId : undefined;
        } else {
            const args = argsOrState as NodeArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["additionalContainerStorageEnabled"] = args ? args.additionalContainerStorageEnabled : undefined;
            resourceInputs["clientToken"] = args ? args.clientToken : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["containerStoragePath"] = args ? args.containerStoragePath : undefined;
            resourceInputs["imageId"] = args ? args.imageId : undefined;
            resourceInputs["initializeScript"] = args ? args.initializeScript : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["keepInstanceName"] = args ? args.keepInstanceName : undefined;
            resourceInputs["kubernetesConfig"] = args ? args.kubernetesConfig : undefined;
            resourceInputs["nodePoolId"] = args ? args.nodePoolId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Node.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Node resources.
 */
export interface NodeState {
    /**
     * The flag of additional container storage enable, the value is `true` or `false`.
     */
    additionalContainerStorageEnabled?: pulumi.Input<boolean>;
    /**
     * The client token.
     */
    clientToken?: pulumi.Input<string>;
    /**
     * The cluster id.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The container storage path.
     */
    containerStoragePath?: pulumi.Input<string>;
    /**
     * The ImageId of NodeConfig.
     */
    imageId?: pulumi.Input<string>;
    /**
     * The initializeScript of Node.
     */
    initializeScript?: pulumi.Input<string>;
    /**
     * The instance id.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The flag of keep instance name, the value is `true` or `false`.
     */
    keepInstanceName?: pulumi.Input<boolean>;
    /**
     * The KubernetesConfig of Node.
     */
    kubernetesConfig?: pulumi.Input<inputs.vke.NodeKubernetesConfig>;
    /**
     * The node pool id.
     */
    nodePoolId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Node resource.
 */
export interface NodeArgs {
    /**
     * The flag of additional container storage enable, the value is `true` or `false`.
     */
    additionalContainerStorageEnabled?: pulumi.Input<boolean>;
    /**
     * The client token.
     */
    clientToken?: pulumi.Input<string>;
    /**
     * The cluster id.
     */
    clusterId: pulumi.Input<string>;
    /**
     * The container storage path.
     */
    containerStoragePath?: pulumi.Input<string>;
    /**
     * The ImageId of NodeConfig.
     */
    imageId?: pulumi.Input<string>;
    /**
     * The initializeScript of Node.
     */
    initializeScript?: pulumi.Input<string>;
    /**
     * The instance id.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The flag of keep instance name, the value is `true` or `false`.
     */
    keepInstanceName?: pulumi.Input<boolean>;
    /**
     * The KubernetesConfig of Node.
     */
    kubernetesConfig?: pulumi.Input<inputs.vke.NodeKubernetesConfig>;
    /**
     * The node pool id.
     */
    nodePoolId?: pulumi.Input<string>;
}
