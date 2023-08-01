// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Vpn
{
    /// <summary>
    /// Provides a resource to manage customer gateway
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
    ///         var foo = new Volcengine.Vpn.CustomerGateway("foo", new Volcengine.Vpn.CustomerGatewayArgs
    ///         {
    ///             CustomerGatewayName = "tf-test",
    ///             Description = "tf-test",
    ///             IpAddress = "192.0.1.3",
    ///             ProjectName = "default",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CustomerGateway can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:vpn/customerGateway:CustomerGateway default cgw-2byswc356dybk2dx0eed2****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:vpn/customerGateway:CustomerGateway")]
    public partial class CustomerGateway : Pulumi.CustomResource
    {
        /// <summary>
        /// The account ID of the customer gateway.
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// The connection count of the customer gateway.
        /// </summary>
        [Output("connectionCount")]
        public Output<int> ConnectionCount { get; private set; } = null!;

        /// <summary>
        /// The create time of customer gateway.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// The ID of the customer gateway.
        /// </summary>
        [Output("customerGatewayId")]
        public Output<string> CustomerGatewayId { get; private set; } = null!;

        /// <summary>
        /// The name of the customer gateway.
        /// </summary>
        [Output("customerGatewayName")]
        public Output<string> CustomerGatewayName { get; private set; } = null!;

        /// <summary>
        /// The description of the customer gateway.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The IP address of the customer gateway.
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// The project name of the VPN customer gateway.
        /// </summary>
        [Output("projectName")]
        public Output<string?> ProjectName { get; private set; } = null!;

        /// <summary>
        /// The status of the customer gateway.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The update time of customer gateway.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a CustomerGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomerGateway(string name, CustomerGatewayArgs args, CustomResourceOptions? options = null)
            : base("volcengine:vpn/customerGateway:CustomerGateway", name, args ?? new CustomerGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomerGateway(string name, Input<string> id, CustomerGatewayState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:vpn/customerGateway:CustomerGateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CustomerGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomerGateway Get(string name, Input<string> id, CustomerGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new CustomerGateway(name, id, state, options);
        }
    }

    public sealed class CustomerGatewayArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the customer gateway.
        /// </summary>
        [Input("customerGatewayName")]
        public Input<string>? CustomerGatewayName { get; set; }

        /// <summary>
        /// The description of the customer gateway.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The IP address of the customer gateway.
        /// </summary>
        [Input("ipAddress", required: true)]
        public Input<string> IpAddress { get; set; } = null!;

        /// <summary>
        /// The project name of the VPN customer gateway.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        public CustomerGatewayArgs()
        {
        }
    }

    public sealed class CustomerGatewayState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account ID of the customer gateway.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// The connection count of the customer gateway.
        /// </summary>
        [Input("connectionCount")]
        public Input<int>? ConnectionCount { get; set; }

        /// <summary>
        /// The create time of customer gateway.
        /// </summary>
        [Input("creationTime")]
        public Input<string>? CreationTime { get; set; }

        /// <summary>
        /// The ID of the customer gateway.
        /// </summary>
        [Input("customerGatewayId")]
        public Input<string>? CustomerGatewayId { get; set; }

        /// <summary>
        /// The name of the customer gateway.
        /// </summary>
        [Input("customerGatewayName")]
        public Input<string>? CustomerGatewayName { get; set; }

        /// <summary>
        /// The description of the customer gateway.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The IP address of the customer gateway.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The project name of the VPN customer gateway.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The status of the customer gateway.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The update time of customer gateway.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public CustomerGatewayState()
        {
        }
    }
}
