// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tls.Outputs
{

    [OutputType]
    public sealed class RulesRuleExcludePathResult
    {
        /// <summary>
        /// The type of the log template.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Collection path.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private RulesRuleExcludePathResult(
            string type,

            string value)
        {
            Type = type;
            Value = value;
        }
    }
}
