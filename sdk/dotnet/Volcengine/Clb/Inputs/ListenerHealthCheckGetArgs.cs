// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Clb.Inputs
{

    public sealed class ListenerHealthCheckGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain of health check.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The enable status of health check function. Optional choice contains `on`, `off`.
        /// </summary>
        [Input("enabled")]
        public Input<string>? Enabled { get; set; }

        /// <summary>
        /// The healthy threshold of health check, default 3, range in 2~10.
        /// </summary>
        [Input("healthyThreshold")]
        public Input<int>? HealthyThreshold { get; set; }

        /// <summary>
        /// The normal http status code of health check, the value can be `http_2xx` or `http_3xx` or `http_4xx` or `http_5xx`.
        /// </summary>
        [Input("httpCode")]
        public Input<string>? HttpCode { get; set; }

        /// <summary>
        /// The interval executing health check, default 2, range in 1~300.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// The method of health check, the value can be `GET` or `HEAD`.
        /// </summary>
        [Input("method")]
        public Input<string>? Method { get; set; }

        /// <summary>
        /// The response timeout of health check, default 2, range in 1~60..
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The UDP expect of health check. This field must be specified simultaneously with field `udp_request`.
        /// </summary>
        [Input("udpExpect")]
        public Input<string>? UdpExpect { get; set; }

        /// <summary>
        /// The UDP request of health check. This field must be specified simultaneously with field `udp_expect`.
        /// </summary>
        [Input("udpRequest")]
        public Input<string>? UdpRequest { get; set; }

        /// <summary>
        /// The unhealthy threshold of health check, default 3, range in 2~10.
        /// </summary>
        [Input("unHealthyThreshold")]
        public Input<int>? UnHealthyThreshold { get; set; }

        /// <summary>
        /// The uri of health check.
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public ListenerHealthCheckGetArgs()
        {
        }
        public static new ListenerHealthCheckGetArgs Empty => new ListenerHealthCheckGetArgs();
    }
}
