// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Vke.Inputs
{

    public sealed class ClusterServicesConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("serviceCidrsv4s", required: true)]
        private InputList<string>? _serviceCidrsv4s;

        /// <summary>
        /// The IPv4 private network address exposed by the service.
        /// </summary>
        public InputList<string> ServiceCidrsv4s
        {
            get => _serviceCidrsv4s ?? (_serviceCidrsv4s = new InputList<string>());
            set => _serviceCidrsv4s = value;
        }

        public ClusterServicesConfigGetArgs()
        {
        }
        public static new ClusterServicesConfigGetArgs Empty => new ClusterServicesConfigGetArgs();
    }
}
