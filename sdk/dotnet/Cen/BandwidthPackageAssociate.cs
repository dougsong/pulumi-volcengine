// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cen
{
    /// <summary>
    /// Provides a resource to manage cen bandwidth package associate
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
    ///         var foo = new Volcengine.Cen.BandwidthPackageAssociate("foo", new Volcengine.Cen.BandwidthPackageAssociateArgs
    ///         {
    ///             CenBandwidthPackageId = "cbp-2bzeew3s8p79c2dx0eeohej4x",
    ///             CenId = "cen-2bzrl3srxsv0g2dx0efyoojn3",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cen bandwidth package associate can be imported using the CenBandwidthPackageId:CenId, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:cen/bandwidthPackageAssociate:BandwidthPackageAssociate default cbp-4c2zaavbvh5fx****:cen-7qthudw0ll6jmc****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:cen/bandwidthPackageAssociate:BandwidthPackageAssociate")]
    public partial class BandwidthPackageAssociate : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the cen bandwidth package.
        /// </summary>
        [Output("cenBandwidthPackageId")]
        public Output<string> CenBandwidthPackageId { get; private set; } = null!;

        /// <summary>
        /// The ID of the cen.
        /// </summary>
        [Output("cenId")]
        public Output<string> CenId { get; private set; } = null!;


        /// <summary>
        /// Create a BandwidthPackageAssociate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BandwidthPackageAssociate(string name, BandwidthPackageAssociateArgs args, CustomResourceOptions? options = null)
            : base("volcengine:cen/bandwidthPackageAssociate:BandwidthPackageAssociate", name, args ?? new BandwidthPackageAssociateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BandwidthPackageAssociate(string name, Input<string> id, BandwidthPackageAssociateState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:cen/bandwidthPackageAssociate:BandwidthPackageAssociate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BandwidthPackageAssociate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BandwidthPackageAssociate Get(string name, Input<string> id, BandwidthPackageAssociateState? state = null, CustomResourceOptions? options = null)
        {
            return new BandwidthPackageAssociate(name, id, state, options);
        }
    }

    public sealed class BandwidthPackageAssociateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the cen bandwidth package.
        /// </summary>
        [Input("cenBandwidthPackageId", required: true)]
        public Input<string> CenBandwidthPackageId { get; set; } = null!;

        /// <summary>
        /// The ID of the cen.
        /// </summary>
        [Input("cenId", required: true)]
        public Input<string> CenId { get; set; } = null!;

        public BandwidthPackageAssociateArgs()
        {
        }
    }

    public sealed class BandwidthPackageAssociateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the cen bandwidth package.
        /// </summary>
        [Input("cenBandwidthPackageId")]
        public Input<string>? CenBandwidthPackageId { get; set; }

        /// <summary>
        /// The ID of the cen.
        /// </summary>
        [Input("cenId")]
        public Input<string>? CenId { get; set; }

        public BandwidthPackageAssociateState()
        {
        }
    }
}
