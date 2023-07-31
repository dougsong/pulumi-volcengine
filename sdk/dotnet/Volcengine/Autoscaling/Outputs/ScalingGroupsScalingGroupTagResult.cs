// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Autoscaling.Outputs
{

    [OutputType]
    public sealed class ScalingGroupsScalingGroupTagResult
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
        private ScalingGroupsScalingGroupTagResult(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
