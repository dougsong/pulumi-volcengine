// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Mongodb.Inputs
{

    public sealed class InstancesTagArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Key of Tags.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        /// <summary>
        /// The Value of Tags.
        /// </summary>
        [Input("value", required: true)]
        public string Value { get; set; } = null!;

        public InstancesTagArgs()
        {
        }
    }
}
