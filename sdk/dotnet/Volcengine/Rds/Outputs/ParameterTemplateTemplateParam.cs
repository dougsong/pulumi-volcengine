// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Rds.Outputs
{

    [OutputType]
    public sealed class ParameterTemplateTemplateParam
    {
        /// <summary>
        /// Parameter name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Parameter running value.
        /// </summary>
        public readonly string? RunningValue;

        [OutputConstructor]
        private ParameterTemplateTemplateParam(
            string? name,

            string? runningValue)
        {
            Name = name;
            RunningValue = runningValue;
        }
    }
}
