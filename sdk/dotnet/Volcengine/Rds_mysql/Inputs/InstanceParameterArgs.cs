// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Rds_mysql.Inputs
{

    public sealed class InstanceParameterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Parameter name.
        /// </summary>
        [Input("parameterName", required: true)]
        public Input<string> ParameterName { get; set; } = null!;

        /// <summary>
        /// Parameter value.
        /// </summary>
        [Input("parameterValue", required: true)]
        public Input<string> ParameterValue { get; set; } = null!;

        public InstanceParameterArgs()
        {
        }
        public static new InstanceParameterArgs Empty => new InstanceParameterArgs();
    }
}
