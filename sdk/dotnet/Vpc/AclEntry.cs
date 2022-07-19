// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vpc
{
    /// <summary>
    /// Provides a resource to manage acl entry
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
    ///         var fooAcl = new Volcengine.Vpc.Acl("fooAcl", new Volcengine.Vpc.AclArgs
    ///         {
    ///             AclName = "tf-test-3",
    ///             Description = "tf-test",
    ///         });
    ///         var fooAclEntry = new Volcengine.Vpc.AclEntry("fooAclEntry", new Volcengine.Vpc.AclEntryArgs
    ///         {
    ///             AclId = fooAcl.Id,
    ///             Description = "tf acl entry desc demo",
    ///             Entry = "192.2.2.1/32",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// AclEntry can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:Vpc/aclEntry:AclEntry default ID is a string concatenated with colons(AclId:Entry)
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:Vpc/aclEntry:AclEntry")]
    public partial class AclEntry : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of Acl.
        /// </summary>
        [Output("aclId")]
        public Output<string> AclId { get; private set; } = null!;

        /// <summary>
        /// The description of the AclEntry.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The content of the AclEntry.
        /// </summary>
        [Output("entry")]
        public Output<string> Entry { get; private set; } = null!;


        /// <summary>
        /// Create a AclEntry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AclEntry(string name, AclEntryArgs args, CustomResourceOptions? options = null)
            : base("volcengine:Vpc/aclEntry:AclEntry", name, args ?? new AclEntryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AclEntry(string name, Input<string> id, AclEntryState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:Vpc/aclEntry:AclEntry", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AclEntry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AclEntry Get(string name, Input<string> id, AclEntryState? state = null, CustomResourceOptions? options = null)
        {
            return new AclEntry(name, id, state, options);
        }
    }

    public sealed class AclEntryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of Acl.
        /// </summary>
        [Input("aclId", required: true)]
        public Input<string> AclId { get; set; } = null!;

        /// <summary>
        /// The description of the AclEntry.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The content of the AclEntry.
        /// </summary>
        [Input("entry", required: true)]
        public Input<string> Entry { get; set; } = null!;

        public AclEntryArgs()
        {
        }
    }

    public sealed class AclEntryState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of Acl.
        /// </summary>
        [Input("aclId")]
        public Input<string>? AclId { get; set; }

        /// <summary>
        /// The description of the AclEntry.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The content of the AclEntry.
        /// </summary>
        [Input("entry")]
        public Input<string>? Entry { get; set; }

        public AclEntryState()
        {
        }
    }
}
