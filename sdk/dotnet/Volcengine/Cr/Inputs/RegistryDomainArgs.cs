// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Cr.Inputs
{

    public sealed class RegistryDomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain of registry.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The type of registry.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public RegistryDomainArgs()
        {
        }
        public static new RegistryDomainArgs Empty => new RegistryDomainArgs();
    }
}
