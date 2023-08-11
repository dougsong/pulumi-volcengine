// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage network acl associate
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.vpc.NetworkAcl("foo", {
 *     vpcId: "vpc-ru0wv9alfoxsu3nuld85rpp",
 *     networkAclName: "tf-test-acl",
 * });
 * const foo1 = new volcengine.vpc.NetworkAclAssociate("foo1", {
 *     networkAclId: foo.id,
 *     resourceId: "subnet-637jxq81u5mon3gd6ivc7rj",
 * });
 * ```
 *
 * ## Import
 *
 * NetworkAcl associate can be imported using the network_acl_id:resource_id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:vpc/networkAclAssociate:NetworkAclAssociate default nacl-172leak37mi9s4d1w33pswqkh:subnet-637jxq81u5mon3gd6ivc7rj
 * ```
 */
export class NetworkAclAssociate extends pulumi.CustomResource {
    /**
     * Get an existing NetworkAclAssociate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkAclAssociateState, opts?: pulumi.CustomResourceOptions): NetworkAclAssociate {
        return new NetworkAclAssociate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:vpc/networkAclAssociate:NetworkAclAssociate';

    /**
     * Returns true if the given object is an instance of NetworkAclAssociate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkAclAssociate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkAclAssociate.__pulumiType;
    }

    /**
     * The id of Network Acl.
     */
    public readonly networkAclId!: pulumi.Output<string>;
    /**
     * The resource id of Network Acl.
     */
    public readonly resourceId!: pulumi.Output<string>;

    /**
     * Create a NetworkAclAssociate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkAclAssociateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkAclAssociateArgs | NetworkAclAssociateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkAclAssociateState | undefined;
            resourceInputs["networkAclId"] = state ? state.networkAclId : undefined;
            resourceInputs["resourceId"] = state ? state.resourceId : undefined;
        } else {
            const args = argsOrState as NetworkAclAssociateArgs | undefined;
            if ((!args || args.networkAclId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkAclId'");
            }
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            resourceInputs["networkAclId"] = args ? args.networkAclId : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetworkAclAssociate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkAclAssociate resources.
 */
export interface NetworkAclAssociateState {
    /**
     * The id of Network Acl.
     */
    networkAclId?: pulumi.Input<string>;
    /**
     * The resource id of Network Acl.
     */
    resourceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkAclAssociate resource.
 */
export interface NetworkAclAssociateArgs {
    /**
     * The id of Network Acl.
     */
    networkAclId: pulumi.Input<string>;
    /**
     * The resource id of Network Acl.
     */
    resourceId: pulumi.Input<string>;
}
