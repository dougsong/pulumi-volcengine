// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Redis.Outputs
{

    [OutputType]
    public sealed class InstancesInstanceParamOptionResult
    {
        /// <summary>
        /// The description of this option item.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The Value of Tags.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private InstancesInstanceParamOptionResult(
            string description,

            string value)
        {
            Description = description;
            Value = value;
        }
    }
}
