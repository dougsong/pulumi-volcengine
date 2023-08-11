// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage vke addon
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.vke.Addon("foo", {
 *     clusterId: "cccctv1vqtofp49d96ujg",
 *     config: "{\"xxx\":\"true\"}",
 *     deployMode: "Unmanaged",
 *     deployNodeType: "Node",
 *     version: "v0.1.3",
 * });
 * ```
 *
 * ## Import
 *
 * VkeAddon can be imported using the clusterId:Name, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:vke/addon:Addon default cc9l74mvqtofjnoj5****:nginx-ingress
 * ```
 *
 *  Notice Some kind of VKEAddon can not be removed from volcengine, and it will make a forbidden error when try to destroy. If you want to remove it from terraform state, please use command $ terraform state rm volcengine_vke_addon.${name}
 */
export class Addon extends pulumi.CustomResource {
    /**
     * Get an existing Addon resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AddonState, opts?: pulumi.CustomResourceOptions): Addon {
        return new Addon(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:vke/addon:Addon';

    /**
     * Returns true if the given object is an instance of Addon.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Addon {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Addon.__pulumiType;
    }

    /**
     * The cluster id of the addon.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * The config info of addon. Please notice that `ingress-nginx` component prohibits updating config, can only works on the web console.
     */
    public readonly config!: pulumi.Output<string>;
    /**
     * The deploy mode.
     */
    public readonly deployMode!: pulumi.Output<string>;
    /**
     * The deploy node type.
     */
    public readonly deployNodeType!: pulumi.Output<string>;
    /**
     * The name of the addon.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The version info of the cluster.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a Addon resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AddonArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AddonArgs | AddonState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AddonState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["deployMode"] = state ? state.deployMode : undefined;
            resourceInputs["deployNodeType"] = state ? state.deployNodeType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as AddonArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["deployMode"] = args ? args.deployMode : undefined;
            resourceInputs["deployNodeType"] = args ? args.deployNodeType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Addon.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Addon resources.
 */
export interface AddonState {
    /**
     * The cluster id of the addon.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The config info of addon. Please notice that `ingress-nginx` component prohibits updating config, can only works on the web console.
     */
    config?: pulumi.Input<string>;
    /**
     * The deploy mode.
     */
    deployMode?: pulumi.Input<string>;
    /**
     * The deploy node type.
     */
    deployNodeType?: pulumi.Input<string>;
    /**
     * The name of the addon.
     */
    name?: pulumi.Input<string>;
    /**
     * The version info of the cluster.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Addon resource.
 */
export interface AddonArgs {
    /**
     * The cluster id of the addon.
     */
    clusterId: pulumi.Input<string>;
    /**
     * The config info of addon. Please notice that `ingress-nginx` component prohibits updating config, can only works on the web console.
     */
    config?: pulumi.Input<string>;
    /**
     * The deploy mode.
     */
    deployMode?: pulumi.Input<string>;
    /**
     * The deploy node type.
     */
    deployNodeType?: pulumi.Input<string>;
    /**
     * The name of the addon.
     */
    name?: pulumi.Input<string>;
    /**
     * The version info of the cluster.
     */
    version?: pulumi.Input<string>;
}
