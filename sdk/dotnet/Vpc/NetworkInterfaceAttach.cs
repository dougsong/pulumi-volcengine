// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vpc
{
    /// <summary>
    /// Provides a resource to manage network interface attach
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
    ///         var foo = new Volcengine.Vpc.NetworkInterfaceAttach("foo", new Volcengine.Vpc.NetworkInterfaceAttachArgs
    ///         {
    ///             InstanceId = "i-72q20hi6s082wcafdem8",
    ///             NetworkInterfaceId = "eni-274ecj646ylts7fap8t6xbba1",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Network interface attach can be imported using the network_interface_id:instance_id.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:vpc/networkInterfaceAttach:NetworkInterfaceAttach default eni-bp1fg655nh68xyz9***:i-wijfn35c****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:vpc/networkInterfaceAttach:NetworkInterfaceAttach")]
    public partial class NetworkInterfaceAttach : Pulumi.CustomResource
    {
        /// <summary>
        /// The id of the instance to which the ENI is bound.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The id of the ENI.
        /// </summary>
        [Output("networkInterfaceId")]
        public Output<string> NetworkInterfaceId { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkInterfaceAttach resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkInterfaceAttach(string name, NetworkInterfaceAttachArgs args, CustomResourceOptions? options = null)
            : base("volcengine:vpc/networkInterfaceAttach:NetworkInterfaceAttach", name, args ?? new NetworkInterfaceAttachArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkInterfaceAttach(string name, Input<string> id, NetworkInterfaceAttachState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:vpc/networkInterfaceAttach:NetworkInterfaceAttach", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkInterfaceAttach resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkInterfaceAttach Get(string name, Input<string> id, NetworkInterfaceAttachState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkInterfaceAttach(name, id, state, options);
        }
    }

    public sealed class NetworkInterfaceAttachArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the instance to which the ENI is bound.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The id of the ENI.
        /// </summary>
        [Input("networkInterfaceId", required: true)]
        public Input<string> NetworkInterfaceId { get; set; } = null!;

        public NetworkInterfaceAttachArgs()
        {
        }
    }

    public sealed class NetworkInterfaceAttachState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the instance to which the ENI is bound.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The id of the ENI.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        public NetworkInterfaceAttachState()
        {
        }
    }
}
