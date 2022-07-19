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
    /// Provides a resource to manage network interface
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
    ///         var foo = new Volcengine.Vpc.NetworkInterface("foo", new Volcengine.Vpc.NetworkInterfaceArgs
    ///         {
    ///             Description = "tf-test-up",
    ///             NetworkInterfaceName = "tf-test-up",
    ///             PortSecurityEnabled = false,
    ///             PrimaryIpAddress = "192.168.0.253",
    ///             SecurityGroupIds = 
    ///             {
    ///                 "sg-2744hspo7jbpc7fap8t7lef1p",
    ///             },
    ///             SubnetId = "subnet-2744ht7fhjthc7fap8tm10eqg",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Network interface can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:Vpc/networkInterface:NetworkInterface default eni-bp1fgnh68xyz9****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:Vpc/networkInterface:NetworkInterface")]
    public partial class NetworkInterface : Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the ENI.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the ENI.
        /// </summary>
        [Output("networkInterfaceName")]
        public Output<string> NetworkInterfaceName { get; private set; } = null!;

        /// <summary>
        /// Set port security enable or disable.
        /// </summary>
        [Output("portSecurityEnabled")]
        public Output<bool> PortSecurityEnabled { get; private set; } = null!;

        /// <summary>
        /// The primary IP address of the ENI.
        /// </summary>
        [Output("primaryIpAddress")]
        public Output<string> PrimaryIpAddress { get; private set; } = null!;

        /// <summary>
        /// The list of the security group id to which the secondary ENI belongs.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// The status of the ENI.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The id of the subnet to which the ENI is connected.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkInterface resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkInterface(string name, NetworkInterfaceArgs args, CustomResourceOptions? options = null)
            : base("volcengine:Vpc/networkInterface:NetworkInterface", name, args ?? new NetworkInterfaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkInterface(string name, Input<string> id, NetworkInterfaceState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:Vpc/networkInterface:NetworkInterface", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkInterface resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkInterface Get(string name, Input<string> id, NetworkInterfaceState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkInterface(name, id, state, options);
        }
    }

    public sealed class NetworkInterfaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the ENI.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the ENI.
        /// </summary>
        [Input("networkInterfaceName")]
        public Input<string>? NetworkInterfaceName { get; set; }

        /// <summary>
        /// Set port security enable or disable.
        /// </summary>
        [Input("portSecurityEnabled")]
        public Input<bool>? PortSecurityEnabled { get; set; }

        /// <summary>
        /// The primary IP address of the ENI.
        /// </summary>
        [Input("primaryIpAddress")]
        public Input<string>? PrimaryIpAddress { get; set; }

        [Input("securityGroupIds", required: true)]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The list of the security group id to which the secondary ENI belongs.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// The id of the subnet to which the ENI is connected.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public NetworkInterfaceArgs()
        {
        }
    }

    public sealed class NetworkInterfaceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the ENI.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the ENI.
        /// </summary>
        [Input("networkInterfaceName")]
        public Input<string>? NetworkInterfaceName { get; set; }

        /// <summary>
        /// Set port security enable or disable.
        /// </summary>
        [Input("portSecurityEnabled")]
        public Input<bool>? PortSecurityEnabled { get; set; }

        /// <summary>
        /// The primary IP address of the ENI.
        /// </summary>
        [Input("primaryIpAddress")]
        public Input<string>? PrimaryIpAddress { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The list of the security group id to which the secondary ENI belongs.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// The status of the ENI.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The id of the subnet to which the ENI is connected.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public NetworkInterfaceState()
        {
        }
    }
}
