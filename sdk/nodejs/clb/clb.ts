// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage clb
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = new volcengine.Clb.Clb("foo", {
 *     description: "Demo",
 *     loadBalancerSpec: "small_1",
 *     regionId: "cn-north-3",
 *     subnetId: "subnet-2744i7u9alnnk7fap8tkq8aft",
 *     type: "public",
 * });
 * ```
 *
 * ## Import
 *
 * CLB can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:Clb/clb:Clb default clb-273y2ok6ets007fap8txvf6us
 * ```
 */
export class Clb extends pulumi.CustomResource {
    /**
     * Get an existing Clb resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClbState, opts?: pulumi.CustomResourceOptions): Clb {
        return new Clb(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:Clb/clb:Clb';

    /**
     * Returns true if the given object is an instance of Clb.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Clb {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Clb.__pulumiType;
    }

    /**
     * The description of the CLB.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The eni address of the CLB.
     */
    public readonly eniAddress!: pulumi.Output<string>;
    /**
     * The billing type of the CLB.
     */
    public readonly loadBalancerBillingType!: pulumi.Output<string>;
    /**
     * The name of the CLB.
     */
    public readonly loadBalancerName!: pulumi.Output<string>;
    /**
     * The specification of the CLB.
     */
    public readonly loadBalancerSpec!: pulumi.Output<string>;
    /**
     * The reason of the console modification protection.
     */
    public readonly modificationProtectionReason!: pulumi.Output<string | undefined>;
    /**
     * The status of the console modification protection.
     */
    public readonly modificationProtectionStatus!: pulumi.Output<string | undefined>;
    /**
     * The region of the request.
     */
    public readonly regionId!: pulumi.Output<string>;
    /**
     * The id of the Subnet.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * The type of the CLB. And optional choice contains `public` or `private`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The id of the VPC.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a Clb resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClbArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClbArgs | ClbState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClbState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["eniAddress"] = state ? state.eniAddress : undefined;
            resourceInputs["loadBalancerBillingType"] = state ? state.loadBalancerBillingType : undefined;
            resourceInputs["loadBalancerName"] = state ? state.loadBalancerName : undefined;
            resourceInputs["loadBalancerSpec"] = state ? state.loadBalancerSpec : undefined;
            resourceInputs["modificationProtectionReason"] = state ? state.modificationProtectionReason : undefined;
            resourceInputs["modificationProtectionStatus"] = state ? state.modificationProtectionStatus : undefined;
            resourceInputs["regionId"] = state ? state.regionId : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as ClbArgs | undefined;
            if ((!args || args.loadBalancerSpec === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loadBalancerSpec'");
            }
            if ((!args || args.regionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'regionId'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["eniAddress"] = args ? args.eniAddress : undefined;
            resourceInputs["loadBalancerBillingType"] = args ? args.loadBalancerBillingType : undefined;
            resourceInputs["loadBalancerName"] = args ? args.loadBalancerName : undefined;
            resourceInputs["loadBalancerSpec"] = args ? args.loadBalancerSpec : undefined;
            resourceInputs["modificationProtectionReason"] = args ? args.modificationProtectionReason : undefined;
            resourceInputs["modificationProtectionStatus"] = args ? args.modificationProtectionStatus : undefined;
            resourceInputs["regionId"] = args ? args.regionId : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Clb.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Clb resources.
 */
export interface ClbState {
    /**
     * The description of the CLB.
     */
    description?: pulumi.Input<string>;
    /**
     * The eni address of the CLB.
     */
    eniAddress?: pulumi.Input<string>;
    /**
     * The billing type of the CLB.
     */
    loadBalancerBillingType?: pulumi.Input<string>;
    /**
     * The name of the CLB.
     */
    loadBalancerName?: pulumi.Input<string>;
    /**
     * The specification of the CLB.
     */
    loadBalancerSpec?: pulumi.Input<string>;
    /**
     * The reason of the console modification protection.
     */
    modificationProtectionReason?: pulumi.Input<string>;
    /**
     * The status of the console modification protection.
     */
    modificationProtectionStatus?: pulumi.Input<string>;
    /**
     * The region of the request.
     */
    regionId?: pulumi.Input<string>;
    /**
     * The id of the Subnet.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The type of the CLB. And optional choice contains `public` or `private`.
     */
    type?: pulumi.Input<string>;
    /**
     * The id of the VPC.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Clb resource.
 */
export interface ClbArgs {
    /**
     * The description of the CLB.
     */
    description?: pulumi.Input<string>;
    /**
     * The eni address of the CLB.
     */
    eniAddress?: pulumi.Input<string>;
    /**
     * The billing type of the CLB.
     */
    loadBalancerBillingType?: pulumi.Input<string>;
    /**
     * The name of the CLB.
     */
    loadBalancerName?: pulumi.Input<string>;
    /**
     * The specification of the CLB.
     */
    loadBalancerSpec: pulumi.Input<string>;
    /**
     * The reason of the console modification protection.
     */
    modificationProtectionReason?: pulumi.Input<string>;
    /**
     * The status of the console modification protection.
     */
    modificationProtectionStatus?: pulumi.Input<string>;
    /**
     * The region of the request.
     */
    regionId: pulumi.Input<string>;
    /**
     * The id of the Subnet.
     */
    subnetId: pulumi.Input<string>;
    /**
     * The type of the CLB. And optional choice contains `public` or `private`.
     */
    type: pulumi.Input<string>;
    /**
     * The id of the VPC.
     */
    vpcId?: pulumi.Input<string>;
}
