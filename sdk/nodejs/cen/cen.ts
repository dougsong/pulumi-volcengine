// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage cen
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = new volcengine.cen.Cen("foo", {
 *     cenName: "tf-test-3",
 *     description: "tf-test",
 *     projectName: "default",
 * });
 * ```
 *
 * ## Import
 *
 * Cen can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:cen/cen:Cen default cen-7qthudw0ll6jmc****
 * ```
 */
export class Cen extends pulumi.CustomResource {
    /**
     * Get an existing Cen resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CenState, opts?: pulumi.CustomResourceOptions): Cen {
        return new Cen(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:cen/cen:Cen';

    /**
     * Returns true if the given object is an instance of Cen.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cen {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cen.__pulumiType;
    }

    /**
     * The account ID of the cen.
     */
    public /*out*/ readonly accountId!: pulumi.Output<string>;
    /**
     * A list of bandwidth package IDs of the cen.
     */
    public /*out*/ readonly cenBandwidthPackageIds!: pulumi.Output<string[]>;
    /**
     * The ID of the cen.
     */
    public /*out*/ readonly cenId!: pulumi.Output<string>;
    /**
     * The name of the cen.
     */
    public readonly cenName!: pulumi.Output<string>;
    /**
     * The create time of the cen.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The description of the cen.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The ProjectName of the cen instance.
     */
    public readonly projectName!: pulumi.Output<string | undefined>;
    /**
     * The status of the cen.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Tags.
     */
    public readonly tags!: pulumi.Output<outputs.cen.CenTag[] | undefined>;
    /**
     * The update time of the cen.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Cen resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CenArgs | CenState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CenState | undefined;
            resourceInputs["accountId"] = state ? state.accountId : undefined;
            resourceInputs["cenBandwidthPackageIds"] = state ? state.cenBandwidthPackageIds : undefined;
            resourceInputs["cenId"] = state ? state.cenId : undefined;
            resourceInputs["cenName"] = state ? state.cenName : undefined;
            resourceInputs["creationTime"] = state ? state.creationTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as CenArgs | undefined;
            resourceInputs["cenName"] = args ? args.cenName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["accountId"] = undefined /*out*/;
            resourceInputs["cenBandwidthPackageIds"] = undefined /*out*/;
            resourceInputs["cenId"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cen.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cen resources.
 */
export interface CenState {
    /**
     * The account ID of the cen.
     */
    accountId?: pulumi.Input<string>;
    /**
     * A list of bandwidth package IDs of the cen.
     */
    cenBandwidthPackageIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the cen.
     */
    cenId?: pulumi.Input<string>;
    /**
     * The name of the cen.
     */
    cenName?: pulumi.Input<string>;
    /**
     * The create time of the cen.
     */
    creationTime?: pulumi.Input<string>;
    /**
     * The description of the cen.
     */
    description?: pulumi.Input<string>;
    /**
     * The ProjectName of the cen instance.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The status of the cen.
     */
    status?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.cen.CenTag>[]>;
    /**
     * The update time of the cen.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cen resource.
 */
export interface CenArgs {
    /**
     * The name of the cen.
     */
    cenName?: pulumi.Input<string>;
    /**
     * The description of the cen.
     */
    description?: pulumi.Input<string>;
    /**
     * The ProjectName of the cen instance.
     */
    projectName?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.cen.CenTag>[]>;
}
