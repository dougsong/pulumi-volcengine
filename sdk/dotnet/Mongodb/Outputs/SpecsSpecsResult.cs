// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Mongodb.Outputs
{

    [OutputType]
    public sealed class SpecsSpecsResult
    {
        /// <summary>
        /// The collection of mongos node specs.
        /// </summary>
        public readonly ImmutableArray<Outputs.SpecsSpecsMongosNodeSpecResult> MongosNodeSpecs;
        /// <summary>
        /// The collection of node specs.
        /// </summary>
        public readonly ImmutableArray<Outputs.SpecsSpecsNodeSpecResult> NodeSpecs;
        /// <summary>
        /// The collection of shard node specs.
        /// </summary>
        public readonly ImmutableArray<Outputs.SpecsSpecsShardNodeSpecResult> ShardNodeSpecs;

        [OutputConstructor]
        private SpecsSpecsResult(
            ImmutableArray<Outputs.SpecsSpecsMongosNodeSpecResult> mongosNodeSpecs,

            ImmutableArray<Outputs.SpecsSpecsNodeSpecResult> nodeSpecs,

            ImmutableArray<Outputs.SpecsSpecsShardNodeSpecResult> shardNodeSpecs)
        {
            MongosNodeSpecs = mongosNodeSpecs;
            NodeSpecs = nodeSpecs;
            ShardNodeSpecs = shardNodeSpecs;
        }
    }
}
