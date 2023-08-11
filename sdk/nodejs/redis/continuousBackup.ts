// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage redis continuous backup
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.redis.ContinuousBackup("foo", {instanceId: "redis-cnlficlt4974s****"});
 * ```
 *
 * ## Import
 *
 * Redis Continuous Backup can be imported using the continuous:instanceId, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:redis/continuousBackup:ContinuousBackup default continuous:redis-asdljioeixxxx
 * ```
 */
export class ContinuousBackup extends pulumi.CustomResource {
    /**
     * Get an existing ContinuousBackup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContinuousBackupState, opts?: pulumi.CustomResourceOptions): ContinuousBackup {
        return new ContinuousBackup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:redis/continuousBackup:ContinuousBackup';

    /**
     * Returns true if the given object is an instance of ContinuousBackup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContinuousBackup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContinuousBackup.__pulumiType;
    }

    /**
     * The Id of redis instance.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a ContinuousBackup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContinuousBackupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContinuousBackupArgs | ContinuousBackupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContinuousBackupState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as ContinuousBackupArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ContinuousBackup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContinuousBackup resources.
 */
export interface ContinuousBackupState {
    /**
     * The Id of redis instance.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ContinuousBackup resource.
 */
export interface ContinuousBackupArgs {
    /**
     * The Id of redis instance.
     */
    instanceId: pulumi.Input<string>;
}
