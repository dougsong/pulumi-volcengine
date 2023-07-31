// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = new volcengine.vpn.Gateway("foo", {
 *     bandwidth: 20,
 *     description: "tf-test",
 *     period: 2,
 *     projectName: "default",
 *     subnetId: "subnet-12bh8g2d7fshs17q7y2nx82uk",
 *     vpcId: "vpc-12b31m7z2kc8w17q7y2fih9ts",
 *     vpnGatewayName: "tf-test",
 * });
 * ```
 *
 * ## Import
 *
 * VpnGateway can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:vpn/gateway:Gateway default vgw-273zkshb2qayo7fap8t2****
 * ```
 */
export class Gateway extends pulumi.CustomResource {
    /**
     * Get an existing Gateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GatewayState, opts?: pulumi.CustomResourceOptions): Gateway {
        return new Gateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:vpn/gateway:Gateway';

    /**
     * Returns true if the given object is an instance of Gateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Gateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Gateway.__pulumiType;
    }

    /**
     * The account ID of the VPN gateway.
     */
    public /*out*/ readonly accountId!: pulumi.Output<string>;
    /**
     * The bandwidth of the VPN gateway. Unit: Mbps. Values: 5, 10, 20, 50, 100, 200, 500.
     */
    public readonly bandwidth!: pulumi.Output<number>;
    /**
     * The BillingType of the VPN gateway. Only support `PrePaid`. Terraform will only remove the PrePaid VPN gateway from the
     * state file, not actually remove.
     */
    public readonly billingType!: pulumi.Output<string | undefined>;
    /**
     * The business status of the VPN gateway.
     */
    public /*out*/ readonly businessStatus!: pulumi.Output<string>;
    /**
     * The connection count of the VPN gateway.
     */
    public /*out*/ readonly connectionCount!: pulumi.Output<number>;
    /**
     * The create time of VPN gateway.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The deleted time of the VPN gateway.
     */
    public /*out*/ readonly deletedTime!: pulumi.Output<string>;
    /**
     * The description of the VPN gateway.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The expired time of the VPN gateway.
     */
    public /*out*/ readonly expiredTime!: pulumi.Output<string>;
    /**
     * The IP address of the VPN gateway.
     */
    public /*out*/ readonly ipAddress!: pulumi.Output<string>;
    /**
     * The lock reason of the VPN gateway.
     */
    public /*out*/ readonly lockReason!: pulumi.Output<string>;
    /**
     * The Period of the VPN gateway. Default value is 12. This parameter is only useful when creating vpn gateway. Default period unit is Month.
     * Value range: 1~9, 12, 24, 36. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The project name of the VPN gateway.
     */
    public readonly projectName!: pulumi.Output<string | undefined>;
    /**
     * The renew type of the VPN gateway.
     */
    public /*out*/ readonly renewType!: pulumi.Output<string>;
    /**
     * The route count of the VPN gateway.
     */
    public /*out*/ readonly routeCount!: pulumi.Output<number>;
    /**
     * The status of the VPN gateway.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The ID of the subnet where you want to create the VPN gateway.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * Tags.
     */
    public readonly tags!: pulumi.Output<outputs.vpn.GatewayTag[] | undefined>;
    /**
     * The update time of VPN gateway.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * The ID of the VPC where you want to create the VPN gateway.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * The ID of the VPN gateway.
     */
    public /*out*/ readonly vpnGatewayId!: pulumi.Output<string>;
    /**
     * The name of the VPN gateway.
     */
    public readonly vpnGatewayName!: pulumi.Output<string>;

    /**
     * Create a Gateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GatewayArgs | GatewayState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GatewayState | undefined;
            resourceInputs["accountId"] = state ? state.accountId : undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["billingType"] = state ? state.billingType : undefined;
            resourceInputs["businessStatus"] = state ? state.businessStatus : undefined;
            resourceInputs["connectionCount"] = state ? state.connectionCount : undefined;
            resourceInputs["creationTime"] = state ? state.creationTime : undefined;
            resourceInputs["deletedTime"] = state ? state.deletedTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["expiredTime"] = state ? state.expiredTime : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["lockReason"] = state ? state.lockReason : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["renewType"] = state ? state.renewType : undefined;
            resourceInputs["routeCount"] = state ? state.routeCount : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vpnGatewayId"] = state ? state.vpnGatewayId : undefined;
            resourceInputs["vpnGatewayName"] = state ? state.vpnGatewayName : undefined;
        } else {
            const args = argsOrState as GatewayArgs | undefined;
            if ((!args || args.bandwidth === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bandwidth'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["billingType"] = args ? args.billingType : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["vpnGatewayName"] = args ? args.vpnGatewayName : undefined;
            resourceInputs["accountId"] = undefined /*out*/;
            resourceInputs["businessStatus"] = undefined /*out*/;
            resourceInputs["connectionCount"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["deletedTime"] = undefined /*out*/;
            resourceInputs["expiredTime"] = undefined /*out*/;
            resourceInputs["ipAddress"] = undefined /*out*/;
            resourceInputs["lockReason"] = undefined /*out*/;
            resourceInputs["renewType"] = undefined /*out*/;
            resourceInputs["routeCount"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["vpnGatewayId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Gateway.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Gateway resources.
 */
export interface GatewayState {
    /**
     * The account ID of the VPN gateway.
     */
    accountId?: pulumi.Input<string>;
    /**
     * The bandwidth of the VPN gateway. Unit: Mbps. Values: 5, 10, 20, 50, 100, 200, 500.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * The BillingType of the VPN gateway. Only support `PrePaid`. Terraform will only remove the PrePaid VPN gateway from the
     * state file, not actually remove.
     */
    billingType?: pulumi.Input<string>;
    /**
     * The business status of the VPN gateway.
     */
    businessStatus?: pulumi.Input<string>;
    /**
     * The connection count of the VPN gateway.
     */
    connectionCount?: pulumi.Input<number>;
    /**
     * The create time of VPN gateway.
     */
    creationTime?: pulumi.Input<string>;
    /**
     * The deleted time of the VPN gateway.
     */
    deletedTime?: pulumi.Input<string>;
    /**
     * The description of the VPN gateway.
     */
    description?: pulumi.Input<string>;
    /**
     * The expired time of the VPN gateway.
     */
    expiredTime?: pulumi.Input<string>;
    /**
     * The IP address of the VPN gateway.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * The lock reason of the VPN gateway.
     */
    lockReason?: pulumi.Input<string>;
    /**
     * The Period of the VPN gateway. Default value is 12. This parameter is only useful when creating vpn gateway. Default period unit is Month.
     * Value range: 1~9, 12, 24, 36. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    period?: pulumi.Input<number>;
    /**
     * The project name of the VPN gateway.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The renew type of the VPN gateway.
     */
    renewType?: pulumi.Input<string>;
    /**
     * The route count of the VPN gateway.
     */
    routeCount?: pulumi.Input<number>;
    /**
     * The status of the VPN gateway.
     */
    status?: pulumi.Input<string>;
    /**
     * The ID of the subnet where you want to create the VPN gateway.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.vpn.GatewayTag>[]>;
    /**
     * The update time of VPN gateway.
     */
    updateTime?: pulumi.Input<string>;
    /**
     * The ID of the VPC where you want to create the VPN gateway.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The ID of the VPN gateway.
     */
    vpnGatewayId?: pulumi.Input<string>;
    /**
     * The name of the VPN gateway.
     */
    vpnGatewayName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Gateway resource.
 */
export interface GatewayArgs {
    /**
     * The bandwidth of the VPN gateway. Unit: Mbps. Values: 5, 10, 20, 50, 100, 200, 500.
     */
    bandwidth: pulumi.Input<number>;
    /**
     * The BillingType of the VPN gateway. Only support `PrePaid`. Terraform will only remove the PrePaid VPN gateway from the
     * state file, not actually remove.
     */
    billingType?: pulumi.Input<string>;
    /**
     * The description of the VPN gateway.
     */
    description?: pulumi.Input<string>;
    /**
     * The Period of the VPN gateway. Default value is 12. This parameter is only useful when creating vpn gateway. Default period unit is Month.
     * Value range: 1~9, 12, 24, 36. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    period?: pulumi.Input<number>;
    /**
     * The project name of the VPN gateway.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The ID of the subnet where you want to create the VPN gateway.
     */
    subnetId: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.vpn.GatewayTag>[]>;
    /**
     * The ID of the VPC where you want to create the VPN gateway.
     */
    vpcId: pulumi.Input<string>;
    /**
     * The name of the VPN gateway.
     */
    vpnGatewayName?: pulumi.Input<string>;
}
