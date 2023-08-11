// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Cen
{
    public static class BandwidthPackages
    {
        /// <summary>
        /// Use this data source to query detailed information of cen bandwidth packages
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
        ///     var foo = Volcengine.Cen.BandwidthPackages.Invoke(new()
        ///     {
        ///         CenId = "cen-2bzrl3srxsv0g2dx0efyoojn3",
        ///         Ids = new[]
        ///         {
        ///             "cbp-2bzeew3s8p79c2dx0eeohej4x",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<BandwidthPackagesResult> InvokeAsync(BandwidthPackagesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<BandwidthPackagesResult>("volcengine:cen/bandwidthPackages:BandwidthPackages", args ?? new BandwidthPackagesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cen bandwidth packages
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
        ///     var foo = Volcengine.Cen.BandwidthPackages.Invoke(new()
        ///     {
        ///         CenId = "cen-2bzrl3srxsv0g2dx0efyoojn3",
        ///         Ids = new[]
        ///         {
        ///             "cbp-2bzeew3s8p79c2dx0eeohej4x",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<BandwidthPackagesResult> Invoke(BandwidthPackagesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<BandwidthPackagesResult>("volcengine:cen/bandwidthPackages:BandwidthPackages", args ?? new BandwidthPackagesInvokeArgs(), options.WithDefaults());
    }


    public sealed class BandwidthPackagesArgs : global::Pulumi.InvokeArgs
    {
        [Input("cenBandwidthPackageNames")]
        private List<string>? _cenBandwidthPackageNames;

        /// <summary>
        /// A list of cen bandwidth package names.
        /// </summary>
        public List<string> CenBandwidthPackageNames
        {
            get => _cenBandwidthPackageNames ?? (_cenBandwidthPackageNames = new List<string>());
            set => _cenBandwidthPackageNames = value;
        }

        /// <summary>
        /// A cen id.
        /// </summary>
        [Input("cenId")]
        public string? CenId { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of cen bandwidth package IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A local geographic region set id.
        /// </summary>
        [Input("localGeographicRegionSetId")]
        public string? LocalGeographicRegionSetId { get; set; }

        /// <summary>
        /// A Name Regex of cen bandwidth package.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// A peer geographic region set id.
        /// </summary>
        [Input("peerGeographicRegionSetId")]
        public string? PeerGeographicRegionSetId { get; set; }

        [Input("tags")]
        private List<Inputs.BandwidthPackagesTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public List<Inputs.BandwidthPackagesTagArgs> Tags
        {
            get => _tags ?? (_tags = new List<Inputs.BandwidthPackagesTagArgs>());
            set => _tags = value;
        }

        public BandwidthPackagesArgs()
        {
        }
        public static new BandwidthPackagesArgs Empty => new BandwidthPackagesArgs();
    }

    public sealed class BandwidthPackagesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("cenBandwidthPackageNames")]
        private InputList<string>? _cenBandwidthPackageNames;

        /// <summary>
        /// A list of cen bandwidth package names.
        /// </summary>
        public InputList<string> CenBandwidthPackageNames
        {
            get => _cenBandwidthPackageNames ?? (_cenBandwidthPackageNames = new InputList<string>());
            set => _cenBandwidthPackageNames = value;
        }

        /// <summary>
        /// A cen id.
        /// </summary>
        [Input("cenId")]
        public Input<string>? CenId { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of cen bandwidth package IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A local geographic region set id.
        /// </summary>
        [Input("localGeographicRegionSetId")]
        public Input<string>? LocalGeographicRegionSetId { get; set; }

        /// <summary>
        /// A Name Regex of cen bandwidth package.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// A peer geographic region set id.
        /// </summary>
        [Input("peerGeographicRegionSetId")]
        public Input<string>? PeerGeographicRegionSetId { get; set; }

        [Input("tags")]
        private InputList<Inputs.BandwidthPackagesTagInputArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.BandwidthPackagesTagInputArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.BandwidthPackagesTagInputArgs>());
            set => _tags = value;
        }

        public BandwidthPackagesInvokeArgs()
        {
        }
        public static new BandwidthPackagesInvokeArgs Empty => new BandwidthPackagesInvokeArgs();
    }


    [OutputType]
    public sealed class BandwidthPackagesResult
    {
        /// <summary>
        /// The collection of cen bandwidth package query.
        /// </summary>
        public readonly ImmutableArray<Outputs.BandwidthPackagesBandwidthPackageResult> BandwidthPackages;
        public readonly ImmutableArray<string> CenBandwidthPackageNames;
        public readonly string? CenId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// The local geographic region set id of the cen bandwidth package.
        /// </summary>
        public readonly string? LocalGeographicRegionSetId;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The peer geographic region set id of the cen bandwidth package.
        /// </summary>
        public readonly string? PeerGeographicRegionSetId;
        /// <summary>
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.BandwidthPackagesTagResult> Tags;
        /// <summary>
        /// The total count of cen bandwidth package query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private BandwidthPackagesResult(
            ImmutableArray<Outputs.BandwidthPackagesBandwidthPackageResult> bandwidthPackages,

            ImmutableArray<string> cenBandwidthPackageNames,

            string? cenId,

            string id,

            ImmutableArray<string> ids,

            string? localGeographicRegionSetId,

            string? nameRegex,

            string? outputFile,

            string? peerGeographicRegionSetId,

            ImmutableArray<Outputs.BandwidthPackagesTagResult> tags,

            int totalCount)
        {
            BandwidthPackages = bandwidthPackages;
            CenBandwidthPackageNames = cenBandwidthPackageNames;
            CenId = cenId;
            Id = id;
            Ids = ids;
            LocalGeographicRegionSetId = localGeographicRegionSetId;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            PeerGeographicRegionSetId = peerGeographicRegionSetId;
            Tags = tags;
            TotalCount = totalCount;
        }
    }
}
