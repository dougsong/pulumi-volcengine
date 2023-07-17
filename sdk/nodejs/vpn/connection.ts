// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage vpn connection
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = new volcengine.vpn.Connection("foo", {
 *     customerGatewayId: "cgw-12ayj1s157gn417q7y29bixqy",
 *     description: "tf-test",
 *     dpdAction: "none",
 *     ikeConfigAuthAlg: "md5",
 *     ikeConfigDhGroup: "group2",
 *     ikeConfigEncAlg: "aes",
 *     ikeConfigLifetime: 100,
 *     ikeConfigLocalId: "tf_test",
 *     ikeConfigMode: "main",
 *     ikeConfigPsk: "tftest@!3",
 *     ikeConfigRemoteId: "tf_test",
 *     ikeConfigVersion: "ikev1",
 *     ipsecConfigAuthAlg: "sha256",
 *     ipsecConfigDhGroup: "group2",
 *     ipsecConfigEncAlg: "aes",
 *     ipsecConfigLifetime: 100,
 *     localSubnets: ["192.168.0.0/22"],
 *     natTraversal: true,
 *     projectName: "default",
 *     remoteSubnets: ["192.161.0.0/20"],
 *     vpnConnectionName: "tf-test",
 *     vpnGatewayId: "vgw-2feq19gnyc9hc59gp68914u6o",
 * });
 * ```
 *
 * ## Import
 *
 * VpnConnection can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:vpn/connection:Connection default vgc-3tex2x1cwd4c6c0v****
 * ```
 */
export class Connection extends pulumi.CustomResource {
    /**
     * Get an existing Connection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConnectionState, opts?: pulumi.CustomResourceOptions): Connection {
        return new Connection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:vpn/connection:Connection';

    /**
     * Returns true if the given object is an instance of Connection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Connection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Connection.__pulumiType;
    }

    /**
     * The account ID of the VPN connection.
     */
    public /*out*/ readonly accountId!: pulumi.Output<string>;
    /**
     * The IPsec attach status.
     */
    public /*out*/ readonly attachStatus!: pulumi.Output<string>;
    /**
     * The attach type of the VPN connection, the value can be VpnGateway or TransitRouter.
     */
    public readonly attachType!: pulumi.Output<string | undefined>;
    /**
     * The business status of IPsec connection, valid when the attach type is 'TransitRouter'.
     */
    public /*out*/ readonly businessStatus!: pulumi.Output<string>;
    /**
     * The connect status of the VPN connection.
     */
    public /*out*/ readonly connectStatus!: pulumi.Output<string>;
    /**
     * The create time of VPN connection.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The ID of the customer gateway.
     */
    public readonly customerGatewayId!: pulumi.Output<string>;
    /**
     * The delete time of resource, valid when the attach type is 'TransitRouter'.
     */
    public /*out*/ readonly deletedTime!: pulumi.Output<string>;
    /**
     * The description of the VPN connection.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The dpd action of the VPN connection.
     */
    public readonly dpdAction!: pulumi.Output<string | undefined>;
    /**
     * The auth alg of the ike config of the VPN connection.
     */
    public readonly ikeConfigAuthAlg!: pulumi.Output<string | undefined>;
    /**
     * The dk group of the ike config of the VPN connection.
     */
    public readonly ikeConfigDhGroup!: pulumi.Output<string | undefined>;
    /**
     * The enc alg of the ike config of the VPN connection.
     */
    public readonly ikeConfigEncAlg!: pulumi.Output<string | undefined>;
    /**
     * The lifetime of the ike config of the VPN connection.
     */
    public readonly ikeConfigLifetime!: pulumi.Output<number | undefined>;
    /**
     * The localId of the ike config of the VPN connection.
     */
    public readonly ikeConfigLocalId!: pulumi.Output<string>;
    /**
     * The mode of the ike config of the VPN connection.
     */
    public readonly ikeConfigMode!: pulumi.Output<string | undefined>;
    /**
     * The psk of the ike config of the VPN connection.
     */
    public readonly ikeConfigPsk!: pulumi.Output<string>;
    /**
     * The remote id of the ike config of the VPN connection.
     */
    public readonly ikeConfigRemoteId!: pulumi.Output<string>;
    /**
     * The version of the ike config of the VPN connection.
     */
    public readonly ikeConfigVersion!: pulumi.Output<string | undefined>;
    /**
     * The ip address of transit router, valid when the attach type is 'TransitRouter'.
     */
    public /*out*/ readonly ipAddress!: pulumi.Output<string>;
    /**
     * The auth alg of the ipsec config of the VPN connection.
     */
    public readonly ipsecConfigAuthAlg!: pulumi.Output<string | undefined>;
    /**
     * The dh group of the ipsec config of the VPN connection.
     */
    public readonly ipsecConfigDhGroup!: pulumi.Output<string | undefined>;
    /**
     * The enc alg of the ipsec config of the VPN connection.
     */
    public readonly ipsecConfigEncAlg!: pulumi.Output<string | undefined>;
    /**
     * The ipsec config of the ike config of the VPN connection.
     */
    public readonly ipsecConfigLifetime!: pulumi.Output<number | undefined>;
    /**
     * The local subnet of the VPN connection.
     */
    public readonly localSubnets!: pulumi.Output<string[]>;
    /**
     * Whether to enable the connection log.
     */
    public /*out*/ readonly logEnabled!: pulumi.Output<boolean>;
    /**
     * The nat traversal of the VPN connection.
     */
    public readonly natTraversal!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to initiate negotiation mode immediately.
     */
    public readonly negotiateInstantly!: pulumi.Output<boolean | undefined>;
    /**
     * The overdue time of resource, valid when the attach type is 'TransitRouter'.
     */
    public /*out*/ readonly overdueTime!: pulumi.Output<string>;
    /**
     * The project name of the VPN connection.
     */
    public readonly projectName!: pulumi.Output<string | undefined>;
    /**
     * The remote subnet of the VPN connection.
     */
    public readonly remoteSubnets!: pulumi.Output<string[]>;
    /**
     * The status of the VPN connection.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The id of transit router, valid when the attach type is 'TransitRouter'.
     */
    public /*out*/ readonly transitRouterId!: pulumi.Output<string>;
    /**
     * The update time of VPN connection.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * The ID of the VPN connection.
     */
    public /*out*/ readonly vpnConnectionId!: pulumi.Output<string>;
    /**
     * The name of the VPN connection.
     */
    public readonly vpnConnectionName!: pulumi.Output<string>;
    /**
     * The ID of the vpn gateway.
     */
    public readonly vpnGatewayId!: pulumi.Output<string | undefined>;
    /**
     * The zone id of transit router, valid when the attach type is 'TransitRouter'.
     */
    public /*out*/ readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a Connection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConnectionArgs | ConnectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConnectionState | undefined;
            resourceInputs["accountId"] = state ? state.accountId : undefined;
            resourceInputs["attachStatus"] = state ? state.attachStatus : undefined;
            resourceInputs["attachType"] = state ? state.attachType : undefined;
            resourceInputs["businessStatus"] = state ? state.businessStatus : undefined;
            resourceInputs["connectStatus"] = state ? state.connectStatus : undefined;
            resourceInputs["creationTime"] = state ? state.creationTime : undefined;
            resourceInputs["customerGatewayId"] = state ? state.customerGatewayId : undefined;
            resourceInputs["deletedTime"] = state ? state.deletedTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dpdAction"] = state ? state.dpdAction : undefined;
            resourceInputs["ikeConfigAuthAlg"] = state ? state.ikeConfigAuthAlg : undefined;
            resourceInputs["ikeConfigDhGroup"] = state ? state.ikeConfigDhGroup : undefined;
            resourceInputs["ikeConfigEncAlg"] = state ? state.ikeConfigEncAlg : undefined;
            resourceInputs["ikeConfigLifetime"] = state ? state.ikeConfigLifetime : undefined;
            resourceInputs["ikeConfigLocalId"] = state ? state.ikeConfigLocalId : undefined;
            resourceInputs["ikeConfigMode"] = state ? state.ikeConfigMode : undefined;
            resourceInputs["ikeConfigPsk"] = state ? state.ikeConfigPsk : undefined;
            resourceInputs["ikeConfigRemoteId"] = state ? state.ikeConfigRemoteId : undefined;
            resourceInputs["ikeConfigVersion"] = state ? state.ikeConfigVersion : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["ipsecConfigAuthAlg"] = state ? state.ipsecConfigAuthAlg : undefined;
            resourceInputs["ipsecConfigDhGroup"] = state ? state.ipsecConfigDhGroup : undefined;
            resourceInputs["ipsecConfigEncAlg"] = state ? state.ipsecConfigEncAlg : undefined;
            resourceInputs["ipsecConfigLifetime"] = state ? state.ipsecConfigLifetime : undefined;
            resourceInputs["localSubnets"] = state ? state.localSubnets : undefined;
            resourceInputs["logEnabled"] = state ? state.logEnabled : undefined;
            resourceInputs["natTraversal"] = state ? state.natTraversal : undefined;
            resourceInputs["negotiateInstantly"] = state ? state.negotiateInstantly : undefined;
            resourceInputs["overdueTime"] = state ? state.overdueTime : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["remoteSubnets"] = state ? state.remoteSubnets : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["transitRouterId"] = state ? state.transitRouterId : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
            resourceInputs["vpnConnectionId"] = state ? state.vpnConnectionId : undefined;
            resourceInputs["vpnConnectionName"] = state ? state.vpnConnectionName : undefined;
            resourceInputs["vpnGatewayId"] = state ? state.vpnGatewayId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as ConnectionArgs | undefined;
            if ((!args || args.customerGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customerGatewayId'");
            }
            if ((!args || args.ikeConfigPsk === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ikeConfigPsk'");
            }
            if ((!args || args.localSubnets === undefined) && !opts.urn) {
                throw new Error("Missing required property 'localSubnets'");
            }
            if ((!args || args.remoteSubnets === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remoteSubnets'");
            }
            resourceInputs["attachType"] = args ? args.attachType : undefined;
            resourceInputs["customerGatewayId"] = args ? args.customerGatewayId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dpdAction"] = args ? args.dpdAction : undefined;
            resourceInputs["ikeConfigAuthAlg"] = args ? args.ikeConfigAuthAlg : undefined;
            resourceInputs["ikeConfigDhGroup"] = args ? args.ikeConfigDhGroup : undefined;
            resourceInputs["ikeConfigEncAlg"] = args ? args.ikeConfigEncAlg : undefined;
            resourceInputs["ikeConfigLifetime"] = args ? args.ikeConfigLifetime : undefined;
            resourceInputs["ikeConfigLocalId"] = args ? args.ikeConfigLocalId : undefined;
            resourceInputs["ikeConfigMode"] = args ? args.ikeConfigMode : undefined;
            resourceInputs["ikeConfigPsk"] = args ? args.ikeConfigPsk : undefined;
            resourceInputs["ikeConfigRemoteId"] = args ? args.ikeConfigRemoteId : undefined;
            resourceInputs["ikeConfigVersion"] = args ? args.ikeConfigVersion : undefined;
            resourceInputs["ipsecConfigAuthAlg"] = args ? args.ipsecConfigAuthAlg : undefined;
            resourceInputs["ipsecConfigDhGroup"] = args ? args.ipsecConfigDhGroup : undefined;
            resourceInputs["ipsecConfigEncAlg"] = args ? args.ipsecConfigEncAlg : undefined;
            resourceInputs["ipsecConfigLifetime"] = args ? args.ipsecConfigLifetime : undefined;
            resourceInputs["localSubnets"] = args ? args.localSubnets : undefined;
            resourceInputs["natTraversal"] = args ? args.natTraversal : undefined;
            resourceInputs["negotiateInstantly"] = args ? args.negotiateInstantly : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["remoteSubnets"] = args ? args.remoteSubnets : undefined;
            resourceInputs["vpnConnectionName"] = args ? args.vpnConnectionName : undefined;
            resourceInputs["vpnGatewayId"] = args ? args.vpnGatewayId : undefined;
            resourceInputs["accountId"] = undefined /*out*/;
            resourceInputs["attachStatus"] = undefined /*out*/;
            resourceInputs["businessStatus"] = undefined /*out*/;
            resourceInputs["connectStatus"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["deletedTime"] = undefined /*out*/;
            resourceInputs["ipAddress"] = undefined /*out*/;
            resourceInputs["logEnabled"] = undefined /*out*/;
            resourceInputs["overdueTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["transitRouterId"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["vpnConnectionId"] = undefined /*out*/;
            resourceInputs["zoneId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Connection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Connection resources.
 */
export interface ConnectionState {
    /**
     * The account ID of the VPN connection.
     */
    accountId?: pulumi.Input<string>;
    /**
     * The IPsec attach status.
     */
    attachStatus?: pulumi.Input<string>;
    /**
     * The attach type of the VPN connection, the value can be VpnGateway or TransitRouter.
     */
    attachType?: pulumi.Input<string>;
    /**
     * The business status of IPsec connection, valid when the attach type is 'TransitRouter'.
     */
    businessStatus?: pulumi.Input<string>;
    /**
     * The connect status of the VPN connection.
     */
    connectStatus?: pulumi.Input<string>;
    /**
     * The create time of VPN connection.
     */
    creationTime?: pulumi.Input<string>;
    /**
     * The ID of the customer gateway.
     */
    customerGatewayId?: pulumi.Input<string>;
    /**
     * The delete time of resource, valid when the attach type is 'TransitRouter'.
     */
    deletedTime?: pulumi.Input<string>;
    /**
     * The description of the VPN connection.
     */
    description?: pulumi.Input<string>;
    /**
     * The dpd action of the VPN connection.
     */
    dpdAction?: pulumi.Input<string>;
    /**
     * The auth alg of the ike config of the VPN connection.
     */
    ikeConfigAuthAlg?: pulumi.Input<string>;
    /**
     * The dk group of the ike config of the VPN connection.
     */
    ikeConfigDhGroup?: pulumi.Input<string>;
    /**
     * The enc alg of the ike config of the VPN connection.
     */
    ikeConfigEncAlg?: pulumi.Input<string>;
    /**
     * The lifetime of the ike config of the VPN connection.
     */
    ikeConfigLifetime?: pulumi.Input<number>;
    /**
     * The localId of the ike config of the VPN connection.
     */
    ikeConfigLocalId?: pulumi.Input<string>;
    /**
     * The mode of the ike config of the VPN connection.
     */
    ikeConfigMode?: pulumi.Input<string>;
    /**
     * The psk of the ike config of the VPN connection.
     */
    ikeConfigPsk?: pulumi.Input<string>;
    /**
     * The remote id of the ike config of the VPN connection.
     */
    ikeConfigRemoteId?: pulumi.Input<string>;
    /**
     * The version of the ike config of the VPN connection.
     */
    ikeConfigVersion?: pulumi.Input<string>;
    /**
     * The ip address of transit router, valid when the attach type is 'TransitRouter'.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * The auth alg of the ipsec config of the VPN connection.
     */
    ipsecConfigAuthAlg?: pulumi.Input<string>;
    /**
     * The dh group of the ipsec config of the VPN connection.
     */
    ipsecConfigDhGroup?: pulumi.Input<string>;
    /**
     * The enc alg of the ipsec config of the VPN connection.
     */
    ipsecConfigEncAlg?: pulumi.Input<string>;
    /**
     * The ipsec config of the ike config of the VPN connection.
     */
    ipsecConfigLifetime?: pulumi.Input<number>;
    /**
     * The local subnet of the VPN connection.
     */
    localSubnets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to enable the connection log.
     */
    logEnabled?: pulumi.Input<boolean>;
    /**
     * The nat traversal of the VPN connection.
     */
    natTraversal?: pulumi.Input<boolean>;
    /**
     * Whether to initiate negotiation mode immediately.
     */
    negotiateInstantly?: pulumi.Input<boolean>;
    /**
     * The overdue time of resource, valid when the attach type is 'TransitRouter'.
     */
    overdueTime?: pulumi.Input<string>;
    /**
     * The project name of the VPN connection.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The remote subnet of the VPN connection.
     */
    remoteSubnets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The status of the VPN connection.
     */
    status?: pulumi.Input<string>;
    /**
     * The id of transit router, valid when the attach type is 'TransitRouter'.
     */
    transitRouterId?: pulumi.Input<string>;
    /**
     * The update time of VPN connection.
     */
    updateTime?: pulumi.Input<string>;
    /**
     * The ID of the VPN connection.
     */
    vpnConnectionId?: pulumi.Input<string>;
    /**
     * The name of the VPN connection.
     */
    vpnConnectionName?: pulumi.Input<string>;
    /**
     * The ID of the vpn gateway.
     */
    vpnGatewayId?: pulumi.Input<string>;
    /**
     * The zone id of transit router, valid when the attach type is 'TransitRouter'.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Connection resource.
 */
export interface ConnectionArgs {
    /**
     * The attach type of the VPN connection, the value can be VpnGateway or TransitRouter.
     */
    attachType?: pulumi.Input<string>;
    /**
     * The ID of the customer gateway.
     */
    customerGatewayId: pulumi.Input<string>;
    /**
     * The description of the VPN connection.
     */
    description?: pulumi.Input<string>;
    /**
     * The dpd action of the VPN connection.
     */
    dpdAction?: pulumi.Input<string>;
    /**
     * The auth alg of the ike config of the VPN connection.
     */
    ikeConfigAuthAlg?: pulumi.Input<string>;
    /**
     * The dk group of the ike config of the VPN connection.
     */
    ikeConfigDhGroup?: pulumi.Input<string>;
    /**
     * The enc alg of the ike config of the VPN connection.
     */
    ikeConfigEncAlg?: pulumi.Input<string>;
    /**
     * The lifetime of the ike config of the VPN connection.
     */
    ikeConfigLifetime?: pulumi.Input<number>;
    /**
     * The localId of the ike config of the VPN connection.
     */
    ikeConfigLocalId?: pulumi.Input<string>;
    /**
     * The mode of the ike config of the VPN connection.
     */
    ikeConfigMode?: pulumi.Input<string>;
    /**
     * The psk of the ike config of the VPN connection.
     */
    ikeConfigPsk: pulumi.Input<string>;
    /**
     * The remote id of the ike config of the VPN connection.
     */
    ikeConfigRemoteId?: pulumi.Input<string>;
    /**
     * The version of the ike config of the VPN connection.
     */
    ikeConfigVersion?: pulumi.Input<string>;
    /**
     * The auth alg of the ipsec config of the VPN connection.
     */
    ipsecConfigAuthAlg?: pulumi.Input<string>;
    /**
     * The dh group of the ipsec config of the VPN connection.
     */
    ipsecConfigDhGroup?: pulumi.Input<string>;
    /**
     * The enc alg of the ipsec config of the VPN connection.
     */
    ipsecConfigEncAlg?: pulumi.Input<string>;
    /**
     * The ipsec config of the ike config of the VPN connection.
     */
    ipsecConfigLifetime?: pulumi.Input<number>;
    /**
     * The local subnet of the VPN connection.
     */
    localSubnets: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The nat traversal of the VPN connection.
     */
    natTraversal?: pulumi.Input<boolean>;
    /**
     * Whether to initiate negotiation mode immediately.
     */
    negotiateInstantly?: pulumi.Input<boolean>;
    /**
     * The project name of the VPN connection.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The remote subnet of the VPN connection.
     */
    remoteSubnets: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the VPN connection.
     */
    vpnConnectionName?: pulumi.Input<string>;
    /**
     * The ID of the vpn gateway.
     */
    vpnGatewayId?: pulumi.Input<string>;
}
