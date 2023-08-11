// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage tls kafka consumer
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const foo = new volcengine.tls.KafkaConsumer("foo", {topicId: "cfb5c08b-0c7a-44fa-8971-8afc12f1b123"});
 * ```
 *
 * ## Import
 *
 * Tls Kafka Consumer can be imported using the kafka:topic_id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:tls/kafkaConsumer:KafkaConsumer default kafka:edf051ed-3c46-49ba-9339-bea628fedc15
 * ```
 */
export class KafkaConsumer extends pulumi.CustomResource {
    /**
     * Get an existing KafkaConsumer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KafkaConsumerState, opts?: pulumi.CustomResourceOptions): KafkaConsumer {
        return new KafkaConsumer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:tls/kafkaConsumer:KafkaConsumer';

    /**
     * Returns true if the given object is an instance of KafkaConsumer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KafkaConsumer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KafkaConsumer.__pulumiType;
    }

    /**
     * Whether allow consume.
     */
    public /*out*/ readonly allowConsume!: pulumi.Output<boolean>;
    /**
     * The topic of consume.
     */
    public /*out*/ readonly consumeTopic!: pulumi.Output<string>;
    /**
     * The id of topic.
     */
    public readonly topicId!: pulumi.Output<string>;

    /**
     * Create a KafkaConsumer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KafkaConsumerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KafkaConsumerArgs | KafkaConsumerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KafkaConsumerState | undefined;
            resourceInputs["allowConsume"] = state ? state.allowConsume : undefined;
            resourceInputs["consumeTopic"] = state ? state.consumeTopic : undefined;
            resourceInputs["topicId"] = state ? state.topicId : undefined;
        } else {
            const args = argsOrState as KafkaConsumerArgs | undefined;
            if ((!args || args.topicId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topicId'");
            }
            resourceInputs["topicId"] = args ? args.topicId : undefined;
            resourceInputs["allowConsume"] = undefined /*out*/;
            resourceInputs["consumeTopic"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(KafkaConsumer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KafkaConsumer resources.
 */
export interface KafkaConsumerState {
    /**
     * Whether allow consume.
     */
    allowConsume?: pulumi.Input<boolean>;
    /**
     * The topic of consume.
     */
    consumeTopic?: pulumi.Input<string>;
    /**
     * The id of topic.
     */
    topicId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KafkaConsumer resource.
 */
export interface KafkaConsumerArgs {
    /**
     * The id of topic.
     */
    topicId: pulumi.Input<string>;
}
