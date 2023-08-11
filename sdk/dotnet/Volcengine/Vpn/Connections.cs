// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Vpn
{
    public static class Connections
    {
        /// <summary>
        /// Use this data source to query detailed information of vpn connections
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = Volcengine.Vpn.Connections.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "vgc-2d5wwids8cdts58ozfe63k2uq",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<ConnectionsResult> InvokeAsync(ConnectionsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ConnectionsResult>("volcengine:vpn/connections:Connections", args ?? new ConnectionsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of vpn connections
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = Volcengine.Vpn.Connections.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "vgc-2d5wwids8cdts58ozfe63k2uq",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<ConnectionsResult> Invoke(ConnectionsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ConnectionsResult>("volcengine:vpn/connections:Connections", args ?? new ConnectionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ConnectionsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// An ID of customer gateway.
        /// </summary>
        [Input("customerGatewayId")]
        public string? CustomerGatewayId { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of VPN connection ids.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of VPN connection.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("vpnConnectionNames")]
        private List<string>? _vpnConnectionNames;

        /// <summary>
        /// A list of VPN connection names.
        /// </summary>
        public List<string> VpnConnectionNames
        {
            get => _vpnConnectionNames ?? (_vpnConnectionNames = new List<string>());
            set => _vpnConnectionNames = value;
        }

        /// <summary>
        /// An ID of VPN gateway.
        /// </summary>
        [Input("vpnGatewayId")]
        public string? VpnGatewayId { get; set; }

        public ConnectionsArgs()
        {
        }
        public static new ConnectionsArgs Empty => new ConnectionsArgs();
    }

    public sealed class ConnectionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// An ID of customer gateway.
        /// </summary>
        [Input("customerGatewayId")]
        public Input<string>? CustomerGatewayId { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of VPN connection ids.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of VPN connection.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("vpnConnectionNames")]
        private InputList<string>? _vpnConnectionNames;

        /// <summary>
        /// A list of VPN connection names.
        /// </summary>
        public InputList<string> VpnConnectionNames
        {
            get => _vpnConnectionNames ?? (_vpnConnectionNames = new InputList<string>());
            set => _vpnConnectionNames = value;
        }

        /// <summary>
        /// An ID of VPN gateway.
        /// </summary>
        [Input("vpnGatewayId")]
        public Input<string>? VpnGatewayId { get; set; }

        public ConnectionsInvokeArgs()
        {
        }
        public static new ConnectionsInvokeArgs Empty => new ConnectionsInvokeArgs();
    }


    [OutputType]
    public sealed class ConnectionsResult
    {
        /// <summary>
        /// The ID of the customer gateway.
        /// </summary>
        public readonly string? CustomerGatewayId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of VPN connection query.
        /// </summary>
        public readonly int TotalCount;
        public readonly ImmutableArray<string> VpnConnectionNames;
        /// <summary>
        /// The collection of VPN connection query.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConnectionsVpnConnectionResult> VpnConnections;
        /// <summary>
        /// The ID of the vpn gateway.
        /// </summary>
        public readonly string? VpnGatewayId;

        [OutputConstructor]
        private ConnectionsResult(
            string? customerGatewayId,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            string? outputFile,

            int totalCount,

            ImmutableArray<string> vpnConnectionNames,

            ImmutableArray<Outputs.ConnectionsVpnConnectionResult> vpnConnections,

            string? vpnGatewayId)
        {
            CustomerGatewayId = customerGatewayId;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            TotalCount = totalCount;
            VpnConnectionNames = vpnConnectionNames;
            VpnConnections = vpnConnections;
            VpnGatewayId = vpnGatewayId;
        }
    }
}
