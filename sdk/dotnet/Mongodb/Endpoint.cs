// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Mongodb
{
    /// <summary>
    /// Provides a resource to manage mongodb endpoint
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Volcengine = Pulumi.Volcengine;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Volcengine.Mongodb.Endpoint("foo", new Volcengine.Mongodb.EndpointArgs
    ///         {
    ///             EipIds = 
    ///             {
    ///                 "eip-3rfe12dvmz8qo5zsk2h91q05p",
    ///             },
    ///             InstanceId = "mongo-replica-38cf5badeb9e",
    ///             NetworkType = "Public",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// mongodb endpoint can be imported using the instanceId:endpointId `instanceId`represents the instance that endpoint related to. `endpointId`the id of endpoint. e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:mongodb/endpoint:Endpoint default mongo-replica-e405f8e2****:BRhFA0pDAk0XXkxCZQ
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:mongodb/endpoint:Endpoint")]
    public partial class Endpoint : Pulumi.CustomResource
    {
        /// <summary>
        /// A list of EIP IDs that need to be bound when applying for endpoint.
        /// </summary>
        [Output("eipIds")]
        public Output<ImmutableArray<string>> EipIds { get; private set; } = null!;

        /// <summary>
        /// The id of endpoint.
        /// </summary>
        [Output("endpointId")]
        public Output<string> EndpointId { get; private set; } = null!;

        /// <summary>
        /// The instance where the endpoint resides.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// A list of the Mongos node that needs to apply for the endpoint.
        /// </summary>
        [Output("mongosNodeIds")]
        public Output<ImmutableArray<string>> MongosNodeIds { get; private set; } = null!;

        /// <summary>
        /// The network type of endpoint.
        /// </summary>
        [Output("networkType")]
        public Output<string?> NetworkType { get; private set; } = null!;

        /// <summary>
        /// The object ID corresponding to the endpoint.
        /// </summary>
        [Output("objectId")]
        public Output<string> ObjectId { get; private set; } = null!;


        /// <summary>
        /// Create a Endpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Endpoint(string name, EndpointArgs args, CustomResourceOptions? options = null)
            : base("volcengine:mongodb/endpoint:Endpoint", name, args ?? new EndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Endpoint(string name, Input<string> id, EndpointState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:mongodb/endpoint:Endpoint", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Endpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Endpoint Get(string name, Input<string> id, EndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new Endpoint(name, id, state, options);
        }
    }

    public sealed class EndpointArgs : Pulumi.ResourceArgs
    {
        [Input("eipIds")]
        private InputList<string>? _eipIds;

        /// <summary>
        /// A list of EIP IDs that need to be bound when applying for endpoint.
        /// </summary>
        public InputList<string> EipIds
        {
            get => _eipIds ?? (_eipIds = new InputList<string>());
            set => _eipIds = value;
        }

        /// <summary>
        /// The instance where the endpoint resides.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("mongosNodeIds")]
        private InputList<string>? _mongosNodeIds;

        /// <summary>
        /// A list of the Mongos node that needs to apply for the endpoint.
        /// </summary>
        public InputList<string> MongosNodeIds
        {
            get => _mongosNodeIds ?? (_mongosNodeIds = new InputList<string>());
            set => _mongosNodeIds = value;
        }

        /// <summary>
        /// The network type of endpoint.
        /// </summary>
        [Input("networkType")]
        public Input<string>? NetworkType { get; set; }

        /// <summary>
        /// The object ID corresponding to the endpoint.
        /// </summary>
        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        public EndpointArgs()
        {
        }
    }

    public sealed class EndpointState : Pulumi.ResourceArgs
    {
        [Input("eipIds")]
        private InputList<string>? _eipIds;

        /// <summary>
        /// A list of EIP IDs that need to be bound when applying for endpoint.
        /// </summary>
        public InputList<string> EipIds
        {
            get => _eipIds ?? (_eipIds = new InputList<string>());
            set => _eipIds = value;
        }

        /// <summary>
        /// The id of endpoint.
        /// </summary>
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        /// <summary>
        /// The instance where the endpoint resides.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("mongosNodeIds")]
        private InputList<string>? _mongosNodeIds;

        /// <summary>
        /// A list of the Mongos node that needs to apply for the endpoint.
        /// </summary>
        public InputList<string> MongosNodeIds
        {
            get => _mongosNodeIds ?? (_mongosNodeIds = new InputList<string>());
            set => _mongosNodeIds = value;
        }

        /// <summary>
        /// The network type of endpoint.
        /// </summary>
        [Input("networkType")]
        public Input<string>? NetworkType { get; set; }

        /// <summary>
        /// The object ID corresponding to the endpoint.
        /// </summary>
        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        public EndpointState()
        {
        }
    }
}
