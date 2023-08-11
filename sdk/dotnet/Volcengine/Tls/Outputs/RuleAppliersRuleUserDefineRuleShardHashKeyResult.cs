// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Tls.Outputs
{

    [OutputType]
    public sealed class RuleAppliersRuleUserDefineRuleShardHashKeyResult
    {
        /// <summary>
        /// The HashKey of the log group is used to specify the partition (shard) to be written to by the current log group.
        /// </summary>
        public readonly string HashKey;

        [OutputConstructor]
        private RuleAppliersRuleUserDefineRuleShardHashKeyResult(string hashKey)
        {
            HashKey = hashKey;
        }
    }
}
