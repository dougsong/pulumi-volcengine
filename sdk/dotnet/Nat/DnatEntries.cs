// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Nat
{
    public static class DnatEntries
    {
        /// <summary>
        /// Use this data source to query detailed information of dnat entries
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var @default = Output.Create(Volcengine.Nat.DnatEntries.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<DnatEntriesResult> InvokeAsync(DnatEntriesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<DnatEntriesResult>("volcengine:nat/dnatEntries:DnatEntries", args ?? new DnatEntriesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dnat entries
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var @default = Output.Create(Volcengine.Nat.DnatEntries.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<DnatEntriesResult> Invoke(DnatEntriesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<DnatEntriesResult>("volcengine:nat/dnatEntries:DnatEntries", args ?? new DnatEntriesInvokeArgs(), options.WithDefaults());
    }


    public sealed class DnatEntriesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the DNAT entry.
        /// </summary>
        [Input("dnatEntryName")]
        public string? DnatEntryName { get; set; }

        /// <summary>
        /// Provides the public IP address for public network access.
        /// </summary>
        [Input("externalIp")]
        public string? ExternalIp { get; set; }

        /// <summary>
        /// The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
        /// </summary>
        [Input("externalPort")]
        public string? ExternalPort { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of DNAT entry ids.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// Provides the internal IP address.
        /// </summary>
        [Input("internalIp")]
        public string? InternalIp { get; set; }

        /// <summary>
        /// The port or port segment on which the cloud server instance provides services to the public network.
        /// </summary>
        [Input("internalPort")]
        public string? InternalPort { get; set; }

        /// <summary>
        /// The id of the NAT gateway.
        /// </summary>
        [Input("natGatewayId")]
        public string? NatGatewayId { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The network protocol.
        /// </summary>
        [Input("protocol")]
        public string? Protocol { get; set; }

        public DnatEntriesArgs()
        {
        }
    }

    public sealed class DnatEntriesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the DNAT entry.
        /// </summary>
        [Input("dnatEntryName")]
        public Input<string>? DnatEntryName { get; set; }

        /// <summary>
        /// Provides the public IP address for public network access.
        /// </summary>
        [Input("externalIp")]
        public Input<string>? ExternalIp { get; set; }

        /// <summary>
        /// The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
        /// </summary>
        [Input("externalPort")]
        public Input<string>? ExternalPort { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of DNAT entry ids.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// Provides the internal IP address.
        /// </summary>
        [Input("internalIp")]
        public Input<string>? InternalIp { get; set; }

        /// <summary>
        /// The port or port segment on which the cloud server instance provides services to the public network.
        /// </summary>
        [Input("internalPort")]
        public Input<string>? InternalPort { get; set; }

        /// <summary>
        /// The id of the NAT gateway.
        /// </summary>
        [Input("natGatewayId")]
        public Input<string>? NatGatewayId { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The network protocol.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public DnatEntriesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class DnatEntriesResult
    {
        /// <summary>
        /// List of DNAT entries.
        /// </summary>
        public readonly ImmutableArray<Outputs.DnatEntriesDnatEntryResult> DnatEntries;
        /// <summary>
        /// The name of the DNAT entry.
        /// </summary>
        public readonly string? DnatEntryName;
        /// <summary>
        /// Provides the public IP address for public network access.
        /// </summary>
        public readonly string? ExternalIp;
        /// <summary>
        /// The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
        /// </summary>
        public readonly string? ExternalPort;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// Provides the internal IP address.
        /// </summary>
        public readonly string? InternalIp;
        /// <summary>
        /// The port or port segment on which the cloud server instance provides services to the public network.
        /// </summary>
        public readonly string? InternalPort;
        /// <summary>
        /// The ID of the NAT gateway.
        /// </summary>
        public readonly string? NatGatewayId;
        public readonly string? OutputFile;
        /// <summary>
        /// The network protocol.
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// The total count of snat entries query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private DnatEntriesResult(
            ImmutableArray<Outputs.DnatEntriesDnatEntryResult> dnatEntries,

            string? dnatEntryName,

            string? externalIp,

            string? externalPort,

            string id,

            ImmutableArray<string> ids,

            string? internalIp,

            string? internalPort,

            string? natGatewayId,

            string? outputFile,

            string? protocol,

            int totalCount)
        {
            DnatEntries = dnatEntries;
            DnatEntryName = dnatEntryName;
            ExternalIp = externalIp;
            ExternalPort = externalPort;
            Id = id;
            Ids = ids;
            InternalIp = internalIp;
            InternalPort = internalPort;
            NatGatewayId = natGatewayId;
            OutputFile = outputFile;
            Protocol = protocol;
            TotalCount = totalCount;
        }
    }
}
