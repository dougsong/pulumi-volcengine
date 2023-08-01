// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Privatelink
{
    /// <summary>
    /// Provides a resource to manage privatelink vpc endpoint service resource
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
    ///         var foo = new Volcengine.Privatelink.VpcEndpointServiceResource("foo", new Volcengine.Privatelink.VpcEndpointServiceResourceArgs
    ///         {
    ///             ResourceId = "clb-3reii8qfbp7gg5zsk2hsrbe3c",
    ///             ServiceId = "epsvc-3rel73uf2ewao5zsk2j2l58ro",
    ///         });
    ///         var foo1 = new Volcengine.Privatelink.VpcEndpointServiceResource("foo1", new Volcengine.Privatelink.VpcEndpointServiceResourceArgs
    ///         {
    ///             ResourceId = "clb-2d6sfye98rzls58ozfducee1o",
    ///             ServiceId = "epsvc-3rel73uf2ewao5zsk2j2l58ro",
    ///         });
    ///         var foo2 = new Volcengine.Privatelink.VpcEndpointServiceResource("foo2", new Volcengine.Privatelink.VpcEndpointServiceResourceArgs
    ///         {
    ///             ResourceId = "clb-3refkvae02gow5zsk2ilaev5y",
    ///             ServiceId = "epsvc-3rel73uf2ewao5zsk2j2l58ro",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// VpcEndpointServiceResource can be imported using the serviceId:resourceId, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource default epsvc-2fe630gurkl37k5gfuy33****:clb-bp1o94dp5i6ea****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource")]
    public partial class VpcEndpointServiceResource : Pulumi.CustomResource
    {
        /// <summary>
        /// The id of resource.
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        /// <summary>
        /// The id of service.
        /// </summary>
        [Output("serviceId")]
        public Output<string> ServiceId { get; private set; } = null!;


        /// <summary>
        /// Create a VpcEndpointServiceResource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcEndpointServiceResource(string name, VpcEndpointServiceResourceArgs args, CustomResourceOptions? options = null)
            : base("volcengine:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource", name, args ?? new VpcEndpointServiceResourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcEndpointServiceResource(string name, Input<string> id, VpcEndpointServiceResourceState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcEndpointServiceResource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcEndpointServiceResource Get(string name, Input<string> id, VpcEndpointServiceResourceState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcEndpointServiceResource(name, id, state, options);
        }
    }

    public sealed class VpcEndpointServiceResourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of resource.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// The id of service.
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public VpcEndpointServiceResourceArgs()
        {
        }
    }

    public sealed class VpcEndpointServiceResourceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of resource.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// The id of service.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        public VpcEndpointServiceResourceState()
        {
        }
    }
}
