// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class AutoscalarAutoscalingPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of seconds that the autoscaler should wait before it
        /// starts collecting information from a new instance. This prevents
        /// the autoscaler from collecting information when the instance is
        /// initializing, during which the collected usage would not be
        /// reliable. The default time autoscaler waits is 60 seconds.
        /// Virtual machine initialization times might vary because of
        /// numerous factors. We recommend that you test how long an
        /// instance may take to initialize. To do this, create an instance
        /// and time the startup process.
        /// </summary>
        [Input("cooldownPeriod")]
        public Input<int>? CooldownPeriod { get; set; }

        /// <summary>
        /// Defines the CPU utilization policy that allows the autoscaler to
        /// scale based on the average CPU utilization of a managed instance
        /// group.  Structure is documented below.
        /// </summary>
        [Input("cpuUtilization")]
        public Input<Inputs.AutoscalarAutoscalingPolicyCpuUtilizationArgs>? CpuUtilization { get; set; }

        /// <summary>
        /// Configuration parameters of autoscaling based on a load balancer.  Structure is documented below.
        /// </summary>
        [Input("loadBalancingUtilization")]
        public Input<Inputs.AutoscalarAutoscalingPolicyLoadBalancingUtilizationArgs>? LoadBalancingUtilization { get; set; }

        /// <summary>
        /// The maximum number of instances that the autoscaler can scale up
        /// to. This is required when creating or updating an autoscaler. The
        /// maximum number of replicas should not be lower than minimal number
        /// of replicas.
        /// </summary>
        [Input("maxReplicas", required: true)]
        public Input<int> MaxReplicas { get; set; } = null!;

        [Input("metrics")]
        private InputList<Inputs.AutoscalarAutoscalingPolicyMetricArgs>? _metrics;

        /// <summary>
        /// Configuration parameters of autoscaling based on a custom metric.  Structure is documented below.
        /// </summary>
        public InputList<Inputs.AutoscalarAutoscalingPolicyMetricArgs> Metrics
        {
            get => _metrics ?? (_metrics = new InputList<Inputs.AutoscalarAutoscalingPolicyMetricArgs>());
            set => _metrics = value;
        }

        /// <summary>
        /// The minimum number of replicas that the autoscaler can scale down
        /// to. This cannot be less than 0. If not provided, autoscaler will
        /// choose a default value depending on maximum number of instances
        /// allowed.
        /// </summary>
        [Input("minReplicas", required: true)]
        public Input<int> MinReplicas { get; set; } = null!;

        /// <summary>
        /// Defines operating mode for this policy.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("scaleDownControl")]
        public Input<Inputs.AutoscalarAutoscalingPolicyScaleDownControlArgs>? ScaleDownControl { get; set; }

        public AutoscalarAutoscalingPolicyArgs()
        {
        }
    }
}
