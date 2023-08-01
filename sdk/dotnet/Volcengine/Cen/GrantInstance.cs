// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Cen
{
    /// <summary>
    /// Provides a resource to manage cen grant instance
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
    ///         var foo = new Volcengine.Cen.GrantInstance("foo", new Volcengine.Cen.GrantInstanceArgs
    ///         {
    ///             CenId = "cen-2d6zdn0c1z5s058ozfcyf4lee",
    ///             CenOwnerId = "210000****",
    ///             InstanceId = "vpc-2bysvq1xx543k2dx0eeulpeiv",
    ///             InstanceRegionId = "cn-guilin-boe",
    ///             InstanceType = "VPC",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cen grant instance can be imported using the CenId:CenOwnerId:InstanceId:InstanceType:RegionId, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:cen/grantInstance:GrantInstance default cen-7qthudw0ll6jmc***:210000****:vpc-2fexiqjlgjif45oxruvso****:VPC:cn-beijing
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:cen/grantInstance:GrantInstance")]
    public partial class GrantInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the cen.
        /// </summary>
        [Output("cenId")]
        public Output<string> CenId { get; private set; } = null!;

        /// <summary>
        /// The owner ID of the cen.
        /// </summary>
        [Output("cenOwnerId")]
        public Output<string> CenOwnerId { get; private set; } = null!;

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The region ID of the instance.
        /// </summary>
        [Output("instanceRegionId")]
        public Output<string> InstanceRegionId { get; private set; } = null!;

        /// <summary>
        /// The type of the instance.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;


        /// <summary>
        /// Create a GrantInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GrantInstance(string name, GrantInstanceArgs args, CustomResourceOptions? options = null)
            : base("volcengine:cen/grantInstance:GrantInstance", name, args ?? new GrantInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GrantInstance(string name, Input<string> id, GrantInstanceState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:cen/grantInstance:GrantInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GrantInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GrantInstance Get(string name, Input<string> id, GrantInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new GrantInstance(name, id, state, options);
        }
    }

    public sealed class GrantInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the cen.
        /// </summary>
        [Input("cenId", required: true)]
        public Input<string> CenId { get; set; } = null!;

        /// <summary>
        /// The owner ID of the cen.
        /// </summary>
        [Input("cenOwnerId", required: true)]
        public Input<string> CenOwnerId { get; set; } = null!;

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The region ID of the instance.
        /// </summary>
        [Input("instanceRegionId", required: true)]
        public Input<string> InstanceRegionId { get; set; } = null!;

        /// <summary>
        /// The type of the instance.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        public GrantInstanceArgs()
        {
        }
    }

    public sealed class GrantInstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the cen.
        /// </summary>
        [Input("cenId")]
        public Input<string>? CenId { get; set; }

        /// <summary>
        /// The owner ID of the cen.
        /// </summary>
        [Input("cenOwnerId")]
        public Input<string>? CenOwnerId { get; set; }

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The region ID of the instance.
        /// </summary>
        [Input("instanceRegionId")]
        public Input<string>? InstanceRegionId { get; set; }

        /// <summary>
        /// The type of the instance.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        public GrantInstanceState()
        {
        }
    }
}
