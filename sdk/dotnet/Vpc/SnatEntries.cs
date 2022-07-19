// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vpc
{
    public static class SnatEntries
    {
        /// <summary>
        /// Use this data source to query detailed information of snat entries
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
        ///         var @default = Output.Create(Volcengine.Vpc.SnatEntries.InvokeAsync(new Volcengine.Vpc.SnatEntriesArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "snat-274zl8b1kxzb47fap8u35uune",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<SnatEntriesResult> InvokeAsync(SnatEntriesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<SnatEntriesResult>("volcengine:Vpc/snatEntries:SnatEntries", args ?? new SnatEntriesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of snat entries
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
        ///         var @default = Output.Create(Volcengine.Vpc.SnatEntries.InvokeAsync(new Volcengine.Vpc.SnatEntriesArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "snat-274zl8b1kxzb47fap8u35uune",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<SnatEntriesResult> Invoke(SnatEntriesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<SnatEntriesResult>("volcengine:Vpc/snatEntries:SnatEntries", args ?? new SnatEntriesInvokeArgs(), options.WithDefaults());
    }


    public sealed class SnatEntriesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// An id of the public ip address used by the SNAT entry.
        /// </summary>
        [Input("eipId")]
        public string? EipId { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of SNAT entry ids.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// An id of the nat gateway to which the entry belongs.
        /// </summary>
        [Input("natGatewayId")]
        public string? NatGatewayId { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// A name of SNAT entry.
        /// </summary>
        [Input("snatEntryName")]
        public string? SnatEntryName { get; set; }

        /// <summary>
        /// An id of the subnet that is required to access the Internet.
        /// </summary>
        [Input("subnetId")]
        public string? SubnetId { get; set; }

        public SnatEntriesArgs()
        {
        }
    }

    public sealed class SnatEntriesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// An id of the public ip address used by the SNAT entry.
        /// </summary>
        [Input("eipId")]
        public Input<string>? EipId { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of SNAT entry ids.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// An id of the nat gateway to which the entry belongs.
        /// </summary>
        [Input("natGatewayId")]
        public Input<string>? NatGatewayId { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// A name of SNAT entry.
        /// </summary>
        [Input("snatEntryName")]
        public Input<string>? SnatEntryName { get; set; }

        /// <summary>
        /// An id of the subnet that is required to access the Internet.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public SnatEntriesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class SnatEntriesResult
    {
        /// <summary>
        /// The id of the public ip address used by the SNAT entry.
        /// </summary>
        public readonly string? EipId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// The id of the nat gateway to which the entry belongs.
        /// </summary>
        public readonly string? NatGatewayId;
        public readonly string? OutputFile;
        /// <summary>
        /// The collection of snat entries.
        /// </summary>
        public readonly ImmutableArray<Outputs.SnatEntriesSnatEntryResult> SnatEntries;
        /// <summary>
        /// The name of the SNAT entry.
        /// </summary>
        public readonly string? SnatEntryName;
        /// <summary>
        /// The id of the subnet that is required to access the internet.
        /// </summary>
        public readonly string? SubnetId;
        /// <summary>
        /// The total count of snat entries query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private SnatEntriesResult(
            string? eipId,

            string id,

            ImmutableArray<string> ids,

            string? natGatewayId,

            string? outputFile,

            ImmutableArray<Outputs.SnatEntriesSnatEntryResult> snatEntries,

            string? snatEntryName,

            string? subnetId,

            int totalCount)
        {
            EipId = eipId;
            Id = id;
            Ids = ids;
            NatGatewayId = natGatewayId;
            OutputFile = outputFile;
            SnatEntries = snatEntries;
            SnatEntryName = snatEntryName;
            SubnetId = subnetId;
            TotalCount = totalCount;
        }
    }
}
