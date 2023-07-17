// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Privatelink.Outputs
{

    [OutputType]
    public sealed class VpcEndpointServicePermissionsPermissionResult
    {
        /// <summary>
        /// The Id of permit account.
        /// </summary>
        public readonly string PermitAccountId;

        [OutputConstructor]
        private VpcEndpointServicePermissionsPermissionResult(string permitAccountId)
        {
            PermitAccountId = permitAccountId;
        }
    }
}
