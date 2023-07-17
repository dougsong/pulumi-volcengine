// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tls
{
    /// <summary>
    /// Provides a resource to manage tls rule applier
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
    ///         var foo = new Volcengine.Tls.RuleApplier("foo", new Volcengine.Tls.RuleApplierArgs
    ///         {
    ///             HostGroupId = "a2a9e8c5-9835-434f-b866-2c1cfa82887d",
    ///             RuleId = "25104b0f-28b7-4a5c-8339-7f9c431d77c8",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// tls rule applier can be imported using the rule id and host group id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:tls/ruleApplier:RuleApplier default fa************:bcb*******
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:tls/ruleApplier:RuleApplier")]
    public partial class RuleApplier : Pulumi.CustomResource
    {
        /// <summary>
        /// The id of the host group.
        /// </summary>
        [Output("hostGroupId")]
        public Output<string> HostGroupId { get; private set; } = null!;

        /// <summary>
        /// The id of the rule.
        /// </summary>
        [Output("ruleId")]
        public Output<string> RuleId { get; private set; } = null!;


        /// <summary>
        /// Create a RuleApplier resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RuleApplier(string name, RuleApplierArgs args, CustomResourceOptions? options = null)
            : base("volcengine:tls/ruleApplier:RuleApplier", name, args ?? new RuleApplierArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RuleApplier(string name, Input<string> id, RuleApplierState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:tls/ruleApplier:RuleApplier", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RuleApplier resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RuleApplier Get(string name, Input<string> id, RuleApplierState? state = null, CustomResourceOptions? options = null)
        {
            return new RuleApplier(name, id, state, options);
        }
    }

    public sealed class RuleApplierArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the host group.
        /// </summary>
        [Input("hostGroupId", required: true)]
        public Input<string> HostGroupId { get; set; } = null!;

        /// <summary>
        /// The id of the rule.
        /// </summary>
        [Input("ruleId", required: true)]
        public Input<string> RuleId { get; set; } = null!;

        public RuleApplierArgs()
        {
        }
    }

    public sealed class RuleApplierState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the host group.
        /// </summary>
        [Input("hostGroupId")]
        public Input<string>? HostGroupId { get; set; }

        /// <summary>
        /// The id of the rule.
        /// </summary>
        [Input("ruleId")]
        public Input<string>? RuleId { get; set; }

        public RuleApplierState()
        {
        }
    }
}
