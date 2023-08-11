// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage dnat entry
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.nat.DnatEntry("foo", {
 *     dnatEntryName: "terraform-test2",
 *     externalIp: "10.249.186.68",
 *     externalPort: "23",
 *     internalIp: "193.168.1.1",
 *     internalPort: "24",
 *     natGatewayId: "ngw-imw3aej7e96o8gbssxkfbybv",
 *     protocol: "tcp",
 * });
 * ```
 *
 * ## Import
 *
 * Dnat entry can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:nat/dnatEntry:DnatEntry default dnat-3fvhk47kf56****
 * ```
 */
export class DnatEntry extends pulumi.CustomResource {
    /**
     * Get an existing DnatEntry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DnatEntryState, opts?: pulumi.CustomResourceOptions): DnatEntry {
        return new DnatEntry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:nat/dnatEntry:DnatEntry';

    /**
     * Returns true if the given object is an instance of DnatEntry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DnatEntry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DnatEntry.__pulumiType;
    }

    /**
     * The id of the DNAT rule.
     */
    public /*out*/ readonly dnatEntryId!: pulumi.Output<string>;
    /**
     * The name of the DNAT rule.
     */
    public readonly dnatEntryName!: pulumi.Output<string | undefined>;
    /**
     * Provides the public IP address for public network access.
     */
    public readonly externalIp!: pulumi.Output<string>;
    /**
     * The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
     */
    public readonly externalPort!: pulumi.Output<string>;
    /**
     * Provides the internal IP address.
     */
    public readonly internalIp!: pulumi.Output<string>;
    /**
     * The port or port segment on which the cloud server instance provides services to the public network.
     */
    public readonly internalPort!: pulumi.Output<string>;
    /**
     * The id of the nat gateway to which the entry belongs.
     */
    public readonly natGatewayId!: pulumi.Output<string>;
    /**
     * The network protocol.
     */
    public readonly protocol!: pulumi.Output<string>;

    /**
     * Create a DnatEntry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DnatEntryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DnatEntryArgs | DnatEntryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DnatEntryState | undefined;
            resourceInputs["dnatEntryId"] = state ? state.dnatEntryId : undefined;
            resourceInputs["dnatEntryName"] = state ? state.dnatEntryName : undefined;
            resourceInputs["externalIp"] = state ? state.externalIp : undefined;
            resourceInputs["externalPort"] = state ? state.externalPort : undefined;
            resourceInputs["internalIp"] = state ? state.internalIp : undefined;
            resourceInputs["internalPort"] = state ? state.internalPort : undefined;
            resourceInputs["natGatewayId"] = state ? state.natGatewayId : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
        } else {
            const args = argsOrState as DnatEntryArgs | undefined;
            if ((!args || args.externalIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'externalIp'");
            }
            if ((!args || args.externalPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'externalPort'");
            }
            if ((!args || args.internalIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'internalIp'");
            }
            if ((!args || args.internalPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'internalPort'");
            }
            if ((!args || args.natGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'natGatewayId'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            resourceInputs["dnatEntryName"] = args ? args.dnatEntryName : undefined;
            resourceInputs["externalIp"] = args ? args.externalIp : undefined;
            resourceInputs["externalPort"] = args ? args.externalPort : undefined;
            resourceInputs["internalIp"] = args ? args.internalIp : undefined;
            resourceInputs["internalPort"] = args ? args.internalPort : undefined;
            resourceInputs["natGatewayId"] = args ? args.natGatewayId : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["dnatEntryId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DnatEntry.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DnatEntry resources.
 */
export interface DnatEntryState {
    /**
     * The id of the DNAT rule.
     */
    dnatEntryId?: pulumi.Input<string>;
    /**
     * The name of the DNAT rule.
     */
    dnatEntryName?: pulumi.Input<string>;
    /**
     * Provides the public IP address for public network access.
     */
    externalIp?: pulumi.Input<string>;
    /**
     * The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
     */
    externalPort?: pulumi.Input<string>;
    /**
     * Provides the internal IP address.
     */
    internalIp?: pulumi.Input<string>;
    /**
     * The port or port segment on which the cloud server instance provides services to the public network.
     */
    internalPort?: pulumi.Input<string>;
    /**
     * The id of the nat gateway to which the entry belongs.
     */
    natGatewayId?: pulumi.Input<string>;
    /**
     * The network protocol.
     */
    protocol?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DnatEntry resource.
 */
export interface DnatEntryArgs {
    /**
     * The name of the DNAT rule.
     */
    dnatEntryName?: pulumi.Input<string>;
    /**
     * Provides the public IP address for public network access.
     */
    externalIp: pulumi.Input<string>;
    /**
     * The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
     */
    externalPort: pulumi.Input<string>;
    /**
     * Provides the internal IP address.
     */
    internalIp: pulumi.Input<string>;
    /**
     * The port or port segment on which the cloud server instance provides services to the public network.
     */
    internalPort: pulumi.Input<string>;
    /**
     * The id of the nat gateway to which the entry belongs.
     */
    natGatewayId: pulumi.Input<string>;
    /**
     * The network protocol.
     */
    protocol: pulumi.Input<string>;
}
