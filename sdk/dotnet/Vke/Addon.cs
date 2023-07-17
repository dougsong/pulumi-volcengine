// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vke
{
    /// <summary>
    /// Provides a resource to manage vke addon
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
    ///         var foo = new Volcengine.Vke.Addon("foo", new Volcengine.Vke.AddonArgs
    ///         {
    ///             ClusterId = "cccctv1vqtofp49d96ujg",
    ///             Config = "{\"xxx\":\"true\"}",
    ///             DeployMode = "Unmanaged",
    ///             DeployNodeType = "Node",
    ///             Version = "v0.1.3",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// VkeAddon can be imported using the clusterId:Name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:vke/addon:Addon default cc9l74mvqtofjnoj5****:nginx-ingress
    /// ```
    /// 
    ///  Notice Some kind of VKEAddon can not be removed from volcengine, and it will make a forbidden error when try to destroy. If you want to remove it from terraform state, please use command $ terraform state rm volcengine_vke_addon.${name}
    /// </summary>
    [VolcengineResourceType("volcengine:vke/addon:Addon")]
    public partial class Addon : Pulumi.CustomResource
    {
        /// <summary>
        /// The cluster id of the addon.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// The config info of addon. Please notice that `ingress-nginx` component prohibits updating config, can only works on the web console.
        /// </summary>
        [Output("config")]
        public Output<string> Config { get; private set; } = null!;

        /// <summary>
        /// The deploy mode.
        /// </summary>
        [Output("deployMode")]
        public Output<string> DeployMode { get; private set; } = null!;

        /// <summary>
        /// The deploy node type.
        /// </summary>
        [Output("deployNodeType")]
        public Output<string> DeployNodeType { get; private set; } = null!;

        /// <summary>
        /// The name of the addon.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The version info of the cluster.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Addon resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Addon(string name, AddonArgs args, CustomResourceOptions? options = null)
            : base("volcengine:vke/addon:Addon", name, args ?? new AddonArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Addon(string name, Input<string> id, AddonState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:vke/addon:Addon", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Addon resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Addon Get(string name, Input<string> id, AddonState? state = null, CustomResourceOptions? options = null)
        {
            return new Addon(name, id, state, options);
        }
    }

    public sealed class AddonArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cluster id of the addon.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The config info of addon. Please notice that `ingress-nginx` component prohibits updating config, can only works on the web console.
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        /// <summary>
        /// The deploy mode.
        /// </summary>
        [Input("deployMode")]
        public Input<string>? DeployMode { get; set; }

        /// <summary>
        /// The deploy node type.
        /// </summary>
        [Input("deployNodeType")]
        public Input<string>? DeployNodeType { get; set; }

        /// <summary>
        /// The name of the addon.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The version info of the cluster.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public AddonArgs()
        {
        }
    }

    public sealed class AddonState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cluster id of the addon.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The config info of addon. Please notice that `ingress-nginx` component prohibits updating config, can only works on the web console.
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        /// <summary>
        /// The deploy mode.
        /// </summary>
        [Input("deployMode")]
        public Input<string>? DeployMode { get; set; }

        /// <summary>
        /// The deploy node type.
        /// </summary>
        [Input("deployNodeType")]
        public Input<string>? DeployNodeType { get; set; }

        /// <summary>
        /// The name of the addon.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The version info of the cluster.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public AddonState()
        {
        }
    }
}
