// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Redis
{
    /// <summary>
    /// Provides a resource to manage redis allow list associate
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
    ///         var @default = new Volcengine.Redis.AllowListAssociate("default", new Volcengine.Redis.AllowListAssociateArgs
    ///         {
    ///             AllowListId = "acl-cnlfc5zsxscu1gg2ajh",
    ///             InstanceId = "redis-cnlfbzifs7bpvundz",
    ///         });
    ///         var default1 = new Volcengine.Redis.AllowListAssociate("default1", new Volcengine.Redis.AllowListAssociateArgs
    ///         {
    ///             AllowListId = "acl-cnlff2mb31zo85p5am7",
    ///             InstanceId = "redis-cnlfbzifs7bpvundz",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Redis AllowList Association can be imported using the instanceId:allowListId, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:redis/allowListAssociate:AllowListAssociate default redis-asdljioeixxxx:acl-cn03wk541s55c376xxxx
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:redis/allowListAssociate:AllowListAssociate")]
    public partial class AllowListAssociate : Pulumi.CustomResource
    {
        /// <summary>
        /// Id of allow list to associate.
        /// </summary>
        [Output("allowListId")]
        public Output<string> AllowListId { get; private set; } = null!;

        /// <summary>
        /// Id of instance to associate.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a AllowListAssociate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AllowListAssociate(string name, AllowListAssociateArgs args, CustomResourceOptions? options = null)
            : base("volcengine:redis/allowListAssociate:AllowListAssociate", name, args ?? new AllowListAssociateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AllowListAssociate(string name, Input<string> id, AllowListAssociateState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:redis/allowListAssociate:AllowListAssociate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AllowListAssociate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AllowListAssociate Get(string name, Input<string> id, AllowListAssociateState? state = null, CustomResourceOptions? options = null)
        {
            return new AllowListAssociate(name, id, state, options);
        }
    }

    public sealed class AllowListAssociateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of allow list to associate.
        /// </summary>
        [Input("allowListId", required: true)]
        public Input<string> AllowListId { get; set; } = null!;

        /// <summary>
        /// Id of instance to associate.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public AllowListAssociateArgs()
        {
        }
    }

    public sealed class AllowListAssociateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of allow list to associate.
        /// </summary>
        [Input("allowListId")]
        public Input<string>? AllowListId { get; set; }

        /// <summary>
        /// Id of instance to associate.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public AllowListAssociateState()
        {
        }
    }
}
