// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage volume attach
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = new volcengine.Ebs.VolumeAttach("foo", {
 *     instanceId: "i-4ay59ww7dq8dt9c29hd4",
 *     volumeId: "vol-3tzl52wubz3b9fciw7ev",
 * });
 * ```
 *
 * ## Import
 *
 * VolumeAttach can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:Ebs/volumeAttach:VolumeAttach default vol-abc12345:i-abc12345
 * ```
 */
export class VolumeAttach extends pulumi.CustomResource {
    /**
     * Get an existing VolumeAttach resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeAttachState, opts?: pulumi.CustomResourceOptions): VolumeAttach {
        return new VolumeAttach(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:Ebs/volumeAttach:VolumeAttach';

    /**
     * Returns true if the given object is an instance of VolumeAttach.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VolumeAttach {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VolumeAttach.__pulumiType;
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
     * The Id of Instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Status of Volume.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Update time of Volume.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * The Id of Volume.
     */
    public readonly volumeId!: pulumi.Output<string>;

    /**
     * Create a VolumeAttach resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeAttachArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VolumeAttachArgs | VolumeAttachState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VolumeAttachState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["deleteWithInstance"] = state ? state.deleteWithInstance : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["volumeId"] = state ? state.volumeId : undefined;
        } else {
            const args = argsOrState as VolumeAttachArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.volumeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'volumeId'");
            }
            resourceInputs["deleteWithInstance"] = args ? args.deleteWithInstance : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["volumeId"] = args ? args.volumeId : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VolumeAttach.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VolumeAttach resources.
 */
export interface VolumeAttachState {
    /**
     * Creation time of Volume.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Delete Volume with Attached Instance.
     */
    deleteWithInstance?: pulumi.Input<boolean>;
    /**
     * The Id of Instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Status of Volume.
     */
    status?: pulumi.Input<string>;
    /**
     * Update time of Volume.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * The Id of Volume.
     */
    volumeId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VolumeAttach resource.
 */
export interface VolumeAttachArgs {
    /**
     * Delete Volume with Attached Instance.
     */
    deleteWithInstance?: pulumi.Input<boolean>;
    /**
     * The Id of Instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The Id of Volume.
     */
    volumeId: pulumi.Input<string>;
}
