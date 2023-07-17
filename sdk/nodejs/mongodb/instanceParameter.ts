// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage mongodb instance parameter
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * //    该资源无法创建，需先import资源
 * //    $ terraform import volcengine_mongodb_instance_parameter.default param:mongo-replica-f16e9298b121:connPoolMaxConnsPerHost
 * //    请注意instance_id和parameter_name需与上述import的ID对应
 * const defaultInstanceParameter = new volcengine.mongodb.InstanceParameter("default", {
 *     instanceId: "mongo-replica-f16e9298b121", // 必填 import之后不允许修改
 *     parameterName: "connPoolMaxConnsPerHost", // 必填 import之后不允许修改
 *     parameterRole: "Node", // 必填
 *     parameterValue: "600", // 必填
 * });
 * ```
 *
 * ## Import
 *
 * mongodb parameter can be imported using the param:instanceId:parameterName, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:mongodb/instanceParameter:InstanceParameter default param:mongo-replica-e405f8e2****:connPoolMaxConnsPerHost
 * ```
 *
 *  NoteThis resource must be imported before it can be used. Please note that instance_id and parameter_name must correspond to the ID of the above import.
 */
export class InstanceParameter extends pulumi.CustomResource {
    /**
     * Get an existing InstanceParameter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceParameterState, opts?: pulumi.CustomResourceOptions): InstanceParameter {
        return new InstanceParameter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:mongodb/instanceParameter:InstanceParameter';

    /**
     * Returns true if the given object is an instance of InstanceParameter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceParameter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceParameter.__pulumiType;
    }

    /**
     * The instance ID. This field cannot be modified after the resource is imported.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The name of parameter. This field cannot be modified after the resource is imported.
     */
    public readonly parameterName!: pulumi.Output<string>;
    /**
     * The node type to which the parameter belongs. The value range is as follows: Node, Shard, ConfigServer, Mongos.
     */
    public readonly parameterRole!: pulumi.Output<string>;
    /**
     * The value of parameter.
     */
    public readonly parameterValue!: pulumi.Output<string>;

    /**
     * Create a InstanceParameter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceParameterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceParameterArgs | InstanceParameterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceParameterState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["parameterName"] = state ? state.parameterName : undefined;
            resourceInputs["parameterRole"] = state ? state.parameterRole : undefined;
            resourceInputs["parameterValue"] = state ? state.parameterValue : undefined;
        } else {
            const args = argsOrState as InstanceParameterArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.parameterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parameterName'");
            }
            if ((!args || args.parameterRole === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parameterRole'");
            }
            if ((!args || args.parameterValue === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parameterValue'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["parameterName"] = args ? args.parameterName : undefined;
            resourceInputs["parameterRole"] = args ? args.parameterRole : undefined;
            resourceInputs["parameterValue"] = args ? args.parameterValue : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstanceParameter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceParameter resources.
 */
export interface InstanceParameterState {
    /**
     * The instance ID. This field cannot be modified after the resource is imported.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The name of parameter. This field cannot be modified after the resource is imported.
     */
    parameterName?: pulumi.Input<string>;
    /**
     * The node type to which the parameter belongs. The value range is as follows: Node, Shard, ConfigServer, Mongos.
     */
    parameterRole?: pulumi.Input<string>;
    /**
     * The value of parameter.
     */
    parameterValue?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceParameter resource.
 */
export interface InstanceParameterArgs {
    /**
     * The instance ID. This field cannot be modified after the resource is imported.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The name of parameter. This field cannot be modified after the resource is imported.
     */
    parameterName: pulumi.Input<string>;
    /**
     * The node type to which the parameter belongs. The value range is as follows: Node, Shard, ConfigServer, Mongos.
     */
    parameterRole: pulumi.Input<string>;
    /**
     * The value of parameter.
     */
    parameterValue: pulumi.Input<string>;
}
