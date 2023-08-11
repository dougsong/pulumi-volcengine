// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooVolume = new volcengine.ebs.Volume("fooVolume", {
 *     volumeName: "terraform-test",
 *     zoneId: "cn-xx-a",
 *     volumeType: "ESSD_PL0",
 *     kind: "data",
 *     size: 40,
 *     volumeChargeType: "PostPaid",
 *     projectName: "default",
 * });
 * const fooVolumeAttach = new volcengine.ebs.VolumeAttach("fooVolumeAttach", {
 *     volumeId: fooVolume.id,
 *     instanceId: "i-yc8pfhbafwijutv6s1fv",
 * });
 * const foo2 = new volcengine.ebs.Volume("foo2", {
 *     volumeName: "terraform-test3",
 *     zoneId: "cn-beijing-b",
 *     volumeType: "ESSD_PL0",
 *     kind: "data",
 *     size: 40,
 *     volumeChargeType: "PrePaid",
 *     instanceId: "i-yc8pfhbafwijutv6s1fv",
 * });
 * ```
 *
 * ## Import
 *
 * Volume can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:ebs/volume:Volume default vol-mizl7m1kqccg5smt1bdpijuj
 * ```
 */
export class Volume extends pulumi.CustomResource {
    /**
     * Get an existing Volume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeState, opts?: pulumi.CustomResourceOptions): Volume {
        return new Volume(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:ebs/volume:Volume';

    /**
     * Returns true if the given object is an instance of Volume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Volume {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Volume.__pulumiType;
    }

    /**
     * Creation time of Volume.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Delete Volume with Attached Instance.
     */
    public readonly deleteWithInstance!: pulumi.Output<boolean>;
    /**
     * The description of the Volume.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the instance to which the created volume is automatically attached. Please note this field needs to ask the system administrator to apply for a whitelist.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The kind of Volume, the value is `data`.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * The ProjectName of the Volume.
     */
    public readonly projectName!: pulumi.Output<string | undefined>;
    /**
     * The size of Volume.
     */
    public readonly size!: pulumi.Output<number>;
    /**
     * Status of Volume.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Status of Trade.
     */
    public /*out*/ readonly tradeStatus!: pulumi.Output<number>;
    /**
     * The charge type of the Volume, the value is `PostPaid` or `PrePaid`. The `PrePaid` volume cannot be detached. Cannot convert `PrePaid` volume to `PostPaid`.Please note that `PrePaid` type needs to ask the system administrator to apply for a whitelist.
     */
    public readonly volumeChargeType!: pulumi.Output<string | undefined>;
    /**
     * The name of Volume.
     */
    public readonly volumeName!: pulumi.Output<string>;
    /**
     * The type of Volume, the value is `PTSSD` or `ESSD_PL0` or `ESSD_PL1` or `ESSD_PL2` or `ESSD_FlexPL`.
     */
    public readonly volumeType!: pulumi.Output<string>;
    /**
     * The id of the Zone.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a Volume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VolumeArgs | VolumeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VolumeState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["deleteWithInstance"] = state ? state.deleteWithInstance : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["kind"] = state ? state.kind : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tradeStatus"] = state ? state.tradeStatus : undefined;
            resourceInputs["volumeChargeType"] = state ? state.volumeChargeType : undefined;
            resourceInputs["volumeName"] = state ? state.volumeName : undefined;
            resourceInputs["volumeType"] = state ? state.volumeType : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as VolumeArgs | undefined;
            if ((!args || args.kind === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kind'");
            }
            if ((!args || args.size === undefined) && !opts.urn) {
                throw new Error("Missing required property 'size'");
            }
            if ((!args || args.volumeName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'volumeName'");
            }
            if ((!args || args.volumeType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'volumeType'");
            }
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            resourceInputs["deleteWithInstance"] = args ? args.deleteWithInstance : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["volumeChargeType"] = args ? args.volumeChargeType : undefined;
            resourceInputs["volumeName"] = args ? args.volumeName : undefined;
            resourceInputs["volumeType"] = args ? args.volumeType : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tradeStatus"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Volume.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Volume resources.
 */
export interface VolumeState {
    /**
     * Creation time of Volume.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Delete Volume with Attached Instance.
     */
    deleteWithInstance?: pulumi.Input<boolean>;
    /**
     * The description of the Volume.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the instance to which the created volume is automatically attached. Please note this field needs to ask the system administrator to apply for a whitelist.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The kind of Volume, the value is `data`.
     */
    kind?: pulumi.Input<string>;
    /**
     * The ProjectName of the Volume.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The size of Volume.
     */
    size?: pulumi.Input<number>;
    /**
     * Status of Volume.
     */
    status?: pulumi.Input<string>;
    /**
     * Status of Trade.
     */
    tradeStatus?: pulumi.Input<number>;
    /**
     * The charge type of the Volume, the value is `PostPaid` or `PrePaid`. The `PrePaid` volume cannot be detached. Cannot convert `PrePaid` volume to `PostPaid`.Please note that `PrePaid` type needs to ask the system administrator to apply for a whitelist.
     */
    volumeChargeType?: pulumi.Input<string>;
    /**
     * The name of Volume.
     */
    volumeName?: pulumi.Input<string>;
    /**
     * The type of Volume, the value is `PTSSD` or `ESSD_PL0` or `ESSD_PL1` or `ESSD_PL2` or `ESSD_FlexPL`.
     */
    volumeType?: pulumi.Input<string>;
    /**
     * The id of the Zone.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Volume resource.
 */
export interface VolumeArgs {
    /**
     * Delete Volume with Attached Instance.
     */
    deleteWithInstance?: pulumi.Input<boolean>;
    /**
     * The description of the Volume.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the instance to which the created volume is automatically attached. Please note this field needs to ask the system administrator to apply for a whitelist.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The kind of Volume, the value is `data`.
     */
    kind: pulumi.Input<string>;
    /**
     * The ProjectName of the Volume.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The size of Volume.
     */
    size: pulumi.Input<number>;
    /**
     * The charge type of the Volume, the value is `PostPaid` or `PrePaid`. The `PrePaid` volume cannot be detached. Cannot convert `PrePaid` volume to `PostPaid`.Please note that `PrePaid` type needs to ask the system administrator to apply for a whitelist.
     */
    volumeChargeType?: pulumi.Input<string>;
    /**
     * The name of Volume.
     */
    volumeName: pulumi.Input<string>;
    /**
     * The type of Volume, the value is `PTSSD` or `ESSD_PL0` or `ESSD_PL1` or `ESSD_PL2` or `ESSD_FlexPL`.
     */
    volumeType: pulumi.Input<string>;
    /**
     * The id of the Zone.
     */
    zoneId: pulumi.Input<string>;
}
