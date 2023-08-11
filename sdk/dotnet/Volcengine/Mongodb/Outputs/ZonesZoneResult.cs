// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Mongodb.Outputs
{

    [OutputType]
    public sealed class ZonesZoneResult
    {
        /// <summary>
        /// The id of the zone.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The id of the zone.
        /// </summary>
        public readonly string ZoneId;
        /// <summary>
        /// The name of the zone.
        /// </summary>
        public readonly string ZoneName;

        [OutputConstructor]
        private ZonesZoneResult(
            string id,

            string zoneId,

            string zoneName)
        {
            Id = id;
            ZoneId = zoneId;
            ZoneName = zoneName;
        }
    }
}
