// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage mongodb allow list associate
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = new volcengine.mongodb.MongoAllowListAssociate("foo", {
 *     allowListId: "acl-9e307ce4efe843fb9ffd8cb6a6cb225f",
 *     instanceId: "mongo-replica-f16e9298b121",
 * });
 * ```
 *
 * ## Import
 *
 * mongodb allow list associate can be imported using the instanceId:allowListId, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:mongodb/mongoAllowListAssociate:MongoAllowListAssociate default mongo-replica-e405f8e2****:acl-d1fd76693bd54e658912e7337d5b****
 * ```
 */
export class MongoAllowListAssociate extends pulumi.CustomResource {
    /**
     * Get an existing MongoAllowListAssociate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MongoAllowListAssociateState, opts?: pulumi.CustomResourceOptions): MongoAllowListAssociate {
        return new MongoAllowListAssociate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:mongodb/mongoAllowListAssociate:MongoAllowListAssociate';

    /**
     * Returns true if the given object is an instance of MongoAllowListAssociate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MongoAllowListAssociate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MongoAllowListAssociate.__pulumiType;
    }

    /**
     * Id of allow list to associate.
     */
    public readonly allowListId!: pulumi.Output<string>;
    /**
     * Id of instance to associate.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a MongoAllowListAssociate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MongoAllowListAssociateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MongoAllowListAssociateArgs | MongoAllowListAssociateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MongoAllowListAssociateState | undefined;
            resourceInputs["allowListId"] = state ? state.allowListId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as MongoAllowListAssociateArgs | undefined;
            if ((!args || args.allowListId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allowListId'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["allowListId"] = args ? args.allowListId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MongoAllowListAssociate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MongoAllowListAssociate resources.
 */
export interface MongoAllowListAssociateState {
    /**
     * Id of allow list to associate.
     */
    allowListId?: pulumi.Input<string>;
    /**
     * Id of instance to associate.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MongoAllowListAssociate resource.
 */
export interface MongoAllowListAssociateArgs {
    /**
     * Id of allow list to associate.
     */
    allowListId: pulumi.Input<string>;
    /**
     * Id of instance to associate.
     */
    instanceId: pulumi.Input<string>;
}
