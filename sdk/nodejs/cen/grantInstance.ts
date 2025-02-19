// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage cen grant instance
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.cen.GrantInstance("foo", {
 *     cenId: "cen-2d6zdn0c1z5s058ozfcyf4lee",
 *     cenOwnerId: "210000****",
 *     instanceId: "vpc-2bysvq1xx543k2dx0eeulpeiv",
 *     instanceRegionId: "cn-guilin-boe",
 *     instanceType: "VPC",
 * });
 * ```
 *
 * ## Import
 *
 * Cen grant instance can be imported using the CenId:CenOwnerId:InstanceId:InstanceType:RegionId, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:cen/grantInstance:GrantInstance default cen-7qthudw0ll6jmc***:210000****:vpc-2fexiqjlgjif45oxruvso****:VPC:cn-beijing
 * ```
 */
export class GrantInstance extends pulumi.CustomResource {
    /**
     * Get an existing GrantInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GrantInstanceState, opts?: pulumi.CustomResourceOptions): GrantInstance {
        return new GrantInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:cen/grantInstance:GrantInstance';

    /**
     * Returns true if the given object is an instance of GrantInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GrantInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GrantInstance.__pulumiType;
    }

    /**
     * The ID of the cen.
     */
    public readonly cenId!: pulumi.Output<string>;
    /**
     * The owner ID of the cen.
     */
    public readonly cenOwnerId!: pulumi.Output<string>;
    /**
     * The ID of the instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The region ID of the instance.
     */
    public readonly instanceRegionId!: pulumi.Output<string>;
    /**
     * The type of the instance.
     */
    public readonly instanceType!: pulumi.Output<string>;

    /**
     * Create a GrantInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GrantInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GrantInstanceArgs | GrantInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GrantInstanceState | undefined;
            resourceInputs["cenId"] = state ? state.cenId : undefined;
            resourceInputs["cenOwnerId"] = state ? state.cenOwnerId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["instanceRegionId"] = state ? state.instanceRegionId : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
        } else {
            const args = argsOrState as GrantInstanceArgs | undefined;
            if ((!args || args.cenId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cenId'");
            }
            if ((!args || args.cenOwnerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cenOwnerId'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.instanceRegionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceRegionId'");
            }
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            resourceInputs["cenId"] = args ? args.cenId : undefined;
            resourceInputs["cenOwnerId"] = args ? args.cenOwnerId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["instanceRegionId"] = args ? args.instanceRegionId : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GrantInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GrantInstance resources.
 */
export interface GrantInstanceState {
    /**
     * The ID of the cen.
     */
    cenId?: pulumi.Input<string>;
    /**
     * The owner ID of the cen.
     */
    cenOwnerId?: pulumi.Input<string>;
    /**
     * The ID of the instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The region ID of the instance.
     */
    instanceRegionId?: pulumi.Input<string>;
    /**
     * The type of the instance.
     */
    instanceType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GrantInstance resource.
 */
export interface GrantInstanceArgs {
    /**
     * The ID of the cen.
     */
    cenId: pulumi.Input<string>;
    /**
     * The owner ID of the cen.
     */
    cenOwnerId: pulumi.Input<string>;
    /**
     * The ID of the instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The region ID of the instance.
     */
    instanceRegionId: pulumi.Input<string>;
    /**
     * The type of the instance.
     */
    instanceType: pulumi.Input<string>;
}
