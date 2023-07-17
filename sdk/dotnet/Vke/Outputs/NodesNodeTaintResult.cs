// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vke.Outputs
{

    [OutputType]
    public sealed class NodesNodeTaintResult
    {
        /// <summary>
        /// The Effect of Taint.
        /// </summary>
        public readonly string Effect;
        /// <summary>
        /// The Key of Taint.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The Value of Taint.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private NodesNodeTaintResult(
            string effect,

            string key,

            string value)
        {
            Effect = effect;
            Key = key;
            Value = value;
        }
    }
}
