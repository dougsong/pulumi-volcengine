// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Tls
{
    /// <summary>
    /// Provides a resource to manage tls kafka consumer
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Volcengine = Volcengine.PulumiPackage.Volcengine;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Volcengine.Tls.KafkaConsumer("foo", new Volcengine.Tls.KafkaConsumerArgs
    ///         {
    ///             TopicId = "cfb5c08b-0c7a-44fa-8971-8afc12f1b123",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Tls Kafka Consumer can be imported using the kafka:topic_id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:tls/kafkaConsumer:KafkaConsumer default kafka:edf051ed-3c46-49ba-9339-bea628fedc15
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:tls/kafkaConsumer:KafkaConsumer")]
    public partial class KafkaConsumer : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether allow consume.
        /// </summary>
        [Output("allowConsume")]
        public Output<bool> AllowConsume { get; private set; } = null!;

        /// <summary>
        /// The topic of consume.
        /// </summary>
        [Output("consumeTopic")]
        public Output<string> ConsumeTopic { get; private set; } = null!;

        /// <summary>
        /// The id of topic.
        /// </summary>
        [Output("topicId")]
        public Output<string> TopicId { get; private set; } = null!;


        /// <summary>
        /// Create a KafkaConsumer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KafkaConsumer(string name, KafkaConsumerArgs args, CustomResourceOptions? options = null)
            : base("volcengine:tls/kafkaConsumer:KafkaConsumer", name, args ?? new KafkaConsumerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KafkaConsumer(string name, Input<string> id, KafkaConsumerState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:tls/kafkaConsumer:KafkaConsumer", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/volcengine",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing KafkaConsumer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KafkaConsumer Get(string name, Input<string> id, KafkaConsumerState? state = null, CustomResourceOptions? options = null)
        {
            return new KafkaConsumer(name, id, state, options);
        }
    }

    public sealed class KafkaConsumerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of topic.
        /// </summary>
        [Input("topicId", required: true)]
        public Input<string> TopicId { get; set; } = null!;

        public KafkaConsumerArgs()
        {
        }
    }

    public sealed class KafkaConsumerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether allow consume.
        /// </summary>
        [Input("allowConsume")]
        public Input<bool>? AllowConsume { get; set; }

        /// <summary>
        /// The topic of consume.
        /// </summary>
        [Input("consumeTopic")]
        public Input<string>? ConsumeTopic { get; set; }

        /// <summary>
        /// The id of topic.
        /// </summary>
        [Input("topicId")]
        public Input<string>? TopicId { get; set; }

        public KafkaConsumerState()
        {
        }
    }
}
