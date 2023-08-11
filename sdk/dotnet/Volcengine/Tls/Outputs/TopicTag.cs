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
    public sealed class TopicTag
    {
        /// <summary>
        /// The Key of Tags.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The Value of Tags.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private TopicTag(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
