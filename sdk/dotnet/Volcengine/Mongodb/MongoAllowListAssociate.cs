// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Mongodb
{
    /// <summary>
    /// Provides a resource to manage mongodb allow list associate
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Volcengine = Volcengine.Pulumi.Volcengine;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var foo = new Volcengine.Mongodb.MongoAllowListAssociate("foo", new()
    ///     {
    ///         AllowListId = "acl-9e307ce4efe843fb9ffd8cb6a6cb225f",
    ///         InstanceId = "mongo-replica-f16e9298b121",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// mongodb allow list associate can be imported using the instanceId:allowListId, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:mongodb/mongoAllowListAssociate:MongoAllowListAssociate default mongo-replica-e405f8e2****:acl-d1fd76693bd54e658912e7337d5b****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:mongodb/mongoAllowListAssociate:MongoAllowListAssociate")]
    public partial class MongoAllowListAssociate : global::Pulumi.CustomResource
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
        /// Create a MongoAllowListAssociate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MongoAllowListAssociate(string name, MongoAllowListAssociateArgs args, CustomResourceOptions? options = null)
            : base("volcengine:mongodb/mongoAllowListAssociate:MongoAllowListAssociate", name, args ?? new MongoAllowListAssociateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MongoAllowListAssociate(string name, Input<string> id, MongoAllowListAssociateState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:mongodb/mongoAllowListAssociate:MongoAllowListAssociate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MongoAllowListAssociate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MongoAllowListAssociate Get(string name, Input<string> id, MongoAllowListAssociateState? state = null, CustomResourceOptions? options = null)
        {
            return new MongoAllowListAssociate(name, id, state, options);
        }
    }

    public sealed class MongoAllowListAssociateArgs : global::Pulumi.ResourceArgs
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

        public MongoAllowListAssociateArgs()
        {
        }
        public static new MongoAllowListAssociateArgs Empty => new MongoAllowListAssociateArgs();
    }

    public sealed class MongoAllowListAssociateState : global::Pulumi.ResourceArgs
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

        public MongoAllowListAssociateState()
        {
        }
        public static new MongoAllowListAssociateState Empty => new MongoAllowListAssociateState();
    }
}
