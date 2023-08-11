// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage privatelink vpc endpoint
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const endpoint = new volcengine.privatelink.VpcEndpoint("endpoint", {
 *     securityGroupIds: ["sg-2d5z8cr53k45c58ozfdum****"],
 *     serviceId: "epsvc-2byz5nzgiansw2dx0eehh****",
 *     endpointName: "tf-test-ep",
 *     description: "tf-test",
 * });
 * const zone = new volcengine.privatelink.VpcEndpointZone("zone", {
 *     endpointId: endpoint.id,
 *     subnetId: "subnet-2bz47q19zhx4w2dx0eevn****",
 *     privateIpAddress: "172.16.0.252",
 * });
 * ```
 *
 * ## Import
 *
 * VpcEndpoint can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:privatelink/vpcEndpoint:VpcEndpoint default ep-3rel74u229dz45zsk2i6l****
 * ```
 */
export class VpcEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing VpcEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcEndpointState, opts?: pulumi.CustomResourceOptions): VpcEndpoint {
        return new VpcEndpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:privatelink/vpcEndpoint:VpcEndpoint';

    /**
     * Returns true if the given object is an instance of VpcEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcEndpoint.__pulumiType;
    }

    /**
     * Whether the vpc endpoint is locked.
     */
    public /*out*/ readonly businessStatus!: pulumi.Output<string>;
    /**
     * The connection  status of vpc endpoint.
     */
    public /*out*/ readonly connectionStatus!: pulumi.Output<string>;
    /**
     * The create time of vpc endpoint.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The delete time of vpc endpoint.
     */
    public /*out*/ readonly deletedTime!: pulumi.Output<string>;
    /**
     * The description of vpc endpoint.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The domain of vpc endpoint.
     */
    public /*out*/ readonly endpointDomain!: pulumi.Output<string>;
    /**
     * The name of vpc endpoint.
     */
    public readonly endpointName!: pulumi.Output<string>;
    /**
     * The type of vpc endpoint.
     */
    public /*out*/ readonly endpointType!: pulumi.Output<string>;
    /**
     * the security group ids of vpc endpoint.
     */
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * The id of vpc endpoint service.
     */
    public readonly serviceId!: pulumi.Output<string>;
    /**
     * The name of vpc endpoint service.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * The status of vpc endpoint.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The update time of vpc endpoint.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * The vpc id of vpc endpoint.
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a VpcEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcEndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcEndpointArgs | VpcEndpointState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcEndpointState | undefined;
            resourceInputs["businessStatus"] = state ? state.businessStatus : undefined;
            resourceInputs["connectionStatus"] = state ? state.connectionStatus : undefined;
            resourceInputs["creationTime"] = state ? state.creationTime : undefined;
            resourceInputs["deletedTime"] = state ? state.deletedTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["endpointDomain"] = state ? state.endpointDomain : undefined;
            resourceInputs["endpointName"] = state ? state.endpointName : undefined;
            resourceInputs["endpointType"] = state ? state.endpointType : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["serviceId"] = state ? state.serviceId : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as VpcEndpointArgs | undefined;
            if ((!args || args.securityGroupIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupIds'");
            }
            if ((!args || args.serviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["endpointName"] = args ? args.endpointName : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["businessStatus"] = undefined /*out*/;
            resourceInputs["connectionStatus"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["deletedTime"] = undefined /*out*/;
            resourceInputs["endpointDomain"] = undefined /*out*/;
            resourceInputs["endpointType"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcEndpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcEndpoint resources.
 */
export interface VpcEndpointState {
    /**
     * Whether the vpc endpoint is locked.
     */
    businessStatus?: pulumi.Input<string>;
    /**
     * The connection  status of vpc endpoint.
     */
    connectionStatus?: pulumi.Input<string>;
    /**
     * The create time of vpc endpoint.
     */
    creationTime?: pulumi.Input<string>;
    /**
     * The delete time of vpc endpoint.
     */
    deletedTime?: pulumi.Input<string>;
    /**
     * The description of vpc endpoint.
     */
    description?: pulumi.Input<string>;
    /**
     * The domain of vpc endpoint.
     */
    endpointDomain?: pulumi.Input<string>;
    /**
     * The name of vpc endpoint.
     */
    endpointName?: pulumi.Input<string>;
    /**
     * The type of vpc endpoint.
     */
    endpointType?: pulumi.Input<string>;
    /**
     * the security group ids of vpc endpoint.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The id of vpc endpoint service.
     */
    serviceId?: pulumi.Input<string>;
    /**
     * The name of vpc endpoint service.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * The status of vpc endpoint.
     */
    status?: pulumi.Input<string>;
    /**
     * The update time of vpc endpoint.
     */
    updateTime?: pulumi.Input<string>;
    /**
     * The vpc id of vpc endpoint.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcEndpoint resource.
 */
export interface VpcEndpointArgs {
    /**
     * The description of vpc endpoint.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of vpc endpoint.
     */
    endpointName?: pulumi.Input<string>;
    /**
     * the security group ids of vpc endpoint.
     */
    securityGroupIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The id of vpc endpoint service.
     */
    serviceId: pulumi.Input<string>;
    /**
     * The name of vpc endpoint service.
     */
    serviceName?: pulumi.Input<string>;
}
