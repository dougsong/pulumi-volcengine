// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage privatelink vpc endpoint service
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.privatelink.VpcEndpointService("foo", {
 *     autoAcceptEnabled: true,
 *     description: "tftest",
 *     resources: [{
 *         resourceId: "clb-2bzxccdjo9uyo2dx0eg0orzla",
 *         resourceType: "CLB",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * VpcEndpointService can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:privatelink/vpcEndpointService:VpcEndpointService default epsvc-2fe630gurkl37k5gfuy33****
 * ```
 */
export class VpcEndpointService extends pulumi.CustomResource {
    /**
     * Get an existing VpcEndpointService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcEndpointServiceState, opts?: pulumi.CustomResourceOptions): VpcEndpointService {
        return new VpcEndpointService(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:privatelink/vpcEndpointService:VpcEndpointService';

    /**
     * Returns true if the given object is an instance of VpcEndpointService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcEndpointService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcEndpointService.__pulumiType;
    }

    /**
     * Whether auto accept node connect.
     */
    public readonly autoAcceptEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The create time of service.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The description of service.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The resources info. When create vpc endpoint service, the resource must exist.
     */
    public readonly resources!: pulumi.Output<outputs.privatelink.VpcEndpointServiceResource[]>;
    /**
     * The domain of service.
     */
    public /*out*/ readonly serviceDomain!: pulumi.Output<string>;
    /**
     * The Id of service.
     */
    public /*out*/ readonly serviceId!: pulumi.Output<string>;
    /**
     * The name of service.
     */
    public /*out*/ readonly serviceName!: pulumi.Output<string>;
    /**
     * The resource type of service.
     */
    public /*out*/ readonly serviceResourceType!: pulumi.Output<string>;
    /**
     * The type of service.
     */
    public /*out*/ readonly serviceType!: pulumi.Output<string>;
    /**
     * The status of service.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The update time of service.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * The IDs of zones.
     */
    public /*out*/ readonly zoneIds!: pulumi.Output<string[]>;

    /**
     * Create a VpcEndpointService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcEndpointServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcEndpointServiceArgs | VpcEndpointServiceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcEndpointServiceState | undefined;
            resourceInputs["autoAcceptEnabled"] = state ? state.autoAcceptEnabled : undefined;
            resourceInputs["creationTime"] = state ? state.creationTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["resources"] = state ? state.resources : undefined;
            resourceInputs["serviceDomain"] = state ? state.serviceDomain : undefined;
            resourceInputs["serviceId"] = state ? state.serviceId : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["serviceResourceType"] = state ? state.serviceResourceType : undefined;
            resourceInputs["serviceType"] = state ? state.serviceType : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
            resourceInputs["zoneIds"] = state ? state.zoneIds : undefined;
        } else {
            const args = argsOrState as VpcEndpointServiceArgs | undefined;
            if ((!args || args.resources === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resources'");
            }
            resourceInputs["autoAcceptEnabled"] = args ? args.autoAcceptEnabled : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["resources"] = args ? args.resources : undefined;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["serviceDomain"] = undefined /*out*/;
            resourceInputs["serviceId"] = undefined /*out*/;
            resourceInputs["serviceName"] = undefined /*out*/;
            resourceInputs["serviceResourceType"] = undefined /*out*/;
            resourceInputs["serviceType"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["zoneIds"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcEndpointService.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcEndpointService resources.
 */
export interface VpcEndpointServiceState {
    /**
     * Whether auto accept node connect.
     */
    autoAcceptEnabled?: pulumi.Input<boolean>;
    /**
     * The create time of service.
     */
    creationTime?: pulumi.Input<string>;
    /**
     * The description of service.
     */
    description?: pulumi.Input<string>;
    /**
     * The resources info. When create vpc endpoint service, the resource must exist.
     */
    resources?: pulumi.Input<pulumi.Input<inputs.privatelink.VpcEndpointServiceResource>[]>;
    /**
     * The domain of service.
     */
    serviceDomain?: pulumi.Input<string>;
    /**
     * The Id of service.
     */
    serviceId?: pulumi.Input<string>;
    /**
     * The name of service.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * The resource type of service.
     */
    serviceResourceType?: pulumi.Input<string>;
    /**
     * The type of service.
     */
    serviceType?: pulumi.Input<string>;
    /**
     * The status of service.
     */
    status?: pulumi.Input<string>;
    /**
     * The update time of service.
     */
    updateTime?: pulumi.Input<string>;
    /**
     * The IDs of zones.
     */
    zoneIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a VpcEndpointService resource.
 */
export interface VpcEndpointServiceArgs {
    /**
     * Whether auto accept node connect.
     */
    autoAcceptEnabled?: pulumi.Input<boolean>;
    /**
     * The description of service.
     */
    description?: pulumi.Input<string>;
    /**
     * The resources info. When create vpc endpoint service, the resource must exist.
     */
    resources: pulumi.Input<pulumi.Input<inputs.privatelink.VpcEndpointServiceResource>[]>;
}
