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
    public sealed class RuleUserDefineRuleParsePathRule
    {
        /// <summary>
        /// A list of field names. Log Service will parse the path sample (PathSample) into multiple fields according to the regular expression (Regex), and Keys is used to specify the field name of each field.
        /// </summary>
        public readonly ImmutableArray<string> Keys;
        /// <summary>
        /// Sample capture path for a real scene.
        /// </summary>
        public readonly string? PathSample;
        /// <summary>
        /// Regular expression for extracting path fields. It must match the collection path sample, otherwise it cannot be extracted successfully.
        /// </summary>
        public readonly string? Regex;

        [OutputConstructor]
        private RuleUserDefineRuleParsePathRule(
            ImmutableArray<string> keys,

            string? pathSample,

            string? regex)
        {
            Keys = keys;
            PathSample = pathSample;
            Regex = regex;
        }
    }
}
