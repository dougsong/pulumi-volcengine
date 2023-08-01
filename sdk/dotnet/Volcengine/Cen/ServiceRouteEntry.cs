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
    /// Provides a resource to manage cen service route entry
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
    ///         var foo = new Volcengine.Cen.ServiceRouteEntry("foo", new Volcengine.Cen.ServiceRouteEntryArgs
    ///         {
    ///             CenId = "cen-12ar8uclj68sg17q7y20v9gil",
    ///             Description = "test-tf",
    ///             DestinationCidrBlock = "100.64.0.0/11",
    ///             PublishMode = "Custom",
    ///             PublishToInstances = 
    ///             {
    ///                 new Volcengine.Cen.Inputs.ServiceRouteEntryPublishToInstanceArgs
    ///                 {
    ///                     InstanceId = "vpc-2fepz36a5ra4g59gp67w197xo",
    ///                     InstanceRegionId = "cn-beijing",
    ///                     InstanceType = "VPC",
    ///                 },
    ///                 new Volcengine.Cen.Inputs.ServiceRouteEntryPublishToInstanceArgs
    ///                 {
    ///                     InstanceId = "vpc-im67wjcikxkw8gbssx8ufpj8",
    ///                     InstanceRegionId = "cn-beijing",
    ///                     InstanceType = "VPC",
    ///                 },
    ///             },
    ///             ServiceRegionId = "cn-beijing",
    ///             ServiceVpcId = "vpc-im67wjcikxkw8gbssx8ufpj8",
    ///         });
    ///         var foo1 = new Volcengine.Cen.ServiceRouteEntry("foo1", new Volcengine.Cen.ServiceRouteEntryArgs
    ///         {
    ///             CenId = "cen-12ar8uclj68sg17q7y20v9gil",
    ///             Description = "test-tf",
    ///             DestinationCidrBlock = "100.64.0.0/10",
    ///             PublishMode = "Custom",
    ///             PublishToInstances = 
    ///             {
    ///                 new Volcengine.Cen.Inputs.ServiceRouteEntryPublishToInstanceArgs
    ///                 {
    ///                     InstanceId = "vpc-2fepz36a5ra4g59gp67w197xo",
    ///                     InstanceRegionId = "cn-beijing",
    ///                     InstanceType = "VPC",
    ///                 },
    ///                 new Volcengine.Cen.Inputs.ServiceRouteEntryPublishToInstanceArgs
    ///                 {
    ///                     InstanceId = "vpc-im67wjcikxkw8gbssx8ufpj8",
    ///                     InstanceRegionId = "cn-beijing",
    ///                     InstanceType = "VPC",
    ///                 },
    ///             },
    ///             ServiceRegionId = "cn-beijing",
    ///             ServiceVpcId = "vpc-im67wjcikxkw8gbssx8ufpj8",
    ///         });
    ///         var foo2 = new Volcengine.Cen.ServiceRouteEntry("foo2", new Volcengine.Cen.ServiceRouteEntryArgs
    ///         {
    ///             CenId = "cen-12ar8uclj68sg17q7y20v9gil",
    ///             Description = "test-tf",
    ///             DestinationCidrBlock = "100.64.0.0/12",
    ///             PublishMode = "Custom",
    ///             PublishToInstances = 
    ///             {
    ///                 new Volcengine.Cen.Inputs.ServiceRouteEntryPublishToInstanceArgs
    ///                 {
    ///                     InstanceId = "vpc-2fepz36a5ra4g59gp67w197xo",
    ///                     InstanceRegionId = "cn-beijing",
    ///                     InstanceType = "VPC",
    ///                 },
    ///                 new Volcengine.Cen.Inputs.ServiceRouteEntryPublishToInstanceArgs
    ///                 {
    ///                     InstanceId = "vpc-im67wjcikxkw8gbssx8ufpj8",
    ///                     InstanceRegionId = "cn-beijing",
    ///                     InstanceType = "VPC",
    ///                 },
    ///             },
    ///             ServiceRegionId = "cn-beijing",
    ///             ServiceVpcId = "vpc-im67wjcikxkw8gbssx8ufpj8",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CenServiceRouteEntry can be imported using the CenId:DestinationCidrBlock:ServiceRegionId:ServiceVpcId, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:cen/serviceRouteEntry:ServiceRouteEntry default cen-2nim00ybaylts7trquyzt****:100.XX.XX.0/24:cn-beijing:vpc-3rlkeggyn6tc010exd32q****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:cen/serviceRouteEntry:ServiceRouteEntry")]
    public partial class ServiceRouteEntry : Pulumi.CustomResource
    {
        /// <summary>
        /// The cen ID of the cen service route entry.
        /// </summary>
        [Output("cenId")]
        public Output<string> CenId { get; private set; } = null!;

        /// <summary>
        /// The create time of the cen service route entry.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// The description of the cen service route entry.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The destination cidr block of the cen service route entry.
        /// </summary>
        [Output("destinationCidrBlock")]
        public Output<string> DestinationCidrBlock { get; private set; } = null!;

        /// <summary>
        /// Publishing scope of cloud service access routes. Valid values are `LocalDCGW`(default), `Custom`.
        /// </summary>
        [Output("publishMode")]
        public Output<string?> PublishMode { get; private set; } = null!;

        /// <summary>
        /// The publish instances. A maximum of 100 can be uploaded in one request. This field needs to be filled in when the `publish_mode` is `Custom`.
        /// </summary>
        [Output("publishToInstances")]
        public Output<ImmutableArray<Outputs.ServiceRouteEntryPublishToInstance>> PublishToInstances { get; private set; } = null!;

        /// <summary>
        /// The service region id of the cen service route entry.
        /// </summary>
        [Output("serviceRegionId")]
        public Output<string> ServiceRegionId { get; private set; } = null!;

        /// <summary>
        /// The service VPC id of the cen service route entry.
        /// </summary>
        [Output("serviceVpcId")]
        public Output<string> ServiceVpcId { get; private set; } = null!;

        /// <summary>
        /// The status of the cen service route entry.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceRouteEntry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceRouteEntry(string name, ServiceRouteEntryArgs args, CustomResourceOptions? options = null)
            : base("volcengine:cen/serviceRouteEntry:ServiceRouteEntry", name, args ?? new ServiceRouteEntryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceRouteEntry(string name, Input<string> id, ServiceRouteEntryState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:cen/serviceRouteEntry:ServiceRouteEntry", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceRouteEntry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceRouteEntry Get(string name, Input<string> id, ServiceRouteEntryState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceRouteEntry(name, id, state, options);
        }
    }

    public sealed class ServiceRouteEntryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cen ID of the cen service route entry.
        /// </summary>
        [Input("cenId", required: true)]
        public Input<string> CenId { get; set; } = null!;

        /// <summary>
        /// The description of the cen service route entry.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The destination cidr block of the cen service route entry.
        /// </summary>
        [Input("destinationCidrBlock", required: true)]
        public Input<string> DestinationCidrBlock { get; set; } = null!;

        /// <summary>
        /// Publishing scope of cloud service access routes. Valid values are `LocalDCGW`(default), `Custom`.
        /// </summary>
        [Input("publishMode")]
        public Input<string>? PublishMode { get; set; }

        [Input("publishToInstances")]
        private InputList<Inputs.ServiceRouteEntryPublishToInstanceArgs>? _publishToInstances;

        /// <summary>
        /// The publish instances. A maximum of 100 can be uploaded in one request. This field needs to be filled in when the `publish_mode` is `Custom`.
        /// </summary>
        public InputList<Inputs.ServiceRouteEntryPublishToInstanceArgs> PublishToInstances
        {
            get => _publishToInstances ?? (_publishToInstances = new InputList<Inputs.ServiceRouteEntryPublishToInstanceArgs>());
            set => _publishToInstances = value;
        }

        /// <summary>
        /// The service region id of the cen service route entry.
        /// </summary>
        [Input("serviceRegionId", required: true)]
        public Input<string> ServiceRegionId { get; set; } = null!;

        /// <summary>
        /// The service VPC id of the cen service route entry.
        /// </summary>
        [Input("serviceVpcId", required: true)]
        public Input<string> ServiceVpcId { get; set; } = null!;

        public ServiceRouteEntryArgs()
        {
        }
    }

    public sealed class ServiceRouteEntryState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cen ID of the cen service route entry.
        /// </summary>
        [Input("cenId")]
        public Input<string>? CenId { get; set; }

        /// <summary>
        /// The create time of the cen service route entry.
        /// </summary>
        [Input("creationTime")]
        public Input<string>? CreationTime { get; set; }

        /// <summary>
        /// The description of the cen service route entry.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The destination cidr block of the cen service route entry.
        /// </summary>
        [Input("destinationCidrBlock")]
        public Input<string>? DestinationCidrBlock { get; set; }

        /// <summary>
        /// Publishing scope of cloud service access routes. Valid values are `LocalDCGW`(default), `Custom`.
        /// </summary>
        [Input("publishMode")]
        public Input<string>? PublishMode { get; set; }

        [Input("publishToInstances")]
        private InputList<Inputs.ServiceRouteEntryPublishToInstanceGetArgs>? _publishToInstances;

        /// <summary>
        /// The publish instances. A maximum of 100 can be uploaded in one request. This field needs to be filled in when the `publish_mode` is `Custom`.
        /// </summary>
        public InputList<Inputs.ServiceRouteEntryPublishToInstanceGetArgs> PublishToInstances
        {
            get => _publishToInstances ?? (_publishToInstances = new InputList<Inputs.ServiceRouteEntryPublishToInstanceGetArgs>());
            set => _publishToInstances = value;
        }

        /// <summary>
        /// The service region id of the cen service route entry.
        /// </summary>
        [Input("serviceRegionId")]
        public Input<string>? ServiceRegionId { get; set; }

        /// <summary>
        /// The service VPC id of the cen service route entry.
        /// </summary>
        [Input("serviceVpcId")]
        public Input<string>? ServiceVpcId { get; set; }

        /// <summary>
        /// The status of the cen service route entry.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ServiceRouteEntryState()
        {
        }
    }
}
