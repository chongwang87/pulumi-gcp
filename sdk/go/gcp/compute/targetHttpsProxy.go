// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Represents a TargetHttpsProxy resource, which is used by one or more
// global forwarding rule to route incoming HTTPS requests to a URL map.
// 
// 
// To get more information about TargetHttpsProxy, see:
// 
// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetHttpsProxies)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/target-proxies)
type TargetHttpsProxy struct {
	s *pulumi.ResourceState
}

// NewTargetHttpsProxy registers a new resource with the given unique name, arguments, and options.
func NewTargetHttpsProxy(ctx *pulumi.Context,
	name string, args *TargetHttpsProxyArgs, opts ...pulumi.ResourceOpt) (*TargetHttpsProxy, error) {
	if args == nil || args.SslCertificates == nil {
		return nil, errors.New("missing required argument 'SslCertificates'")
	}
	if args == nil || args.UrlMap == nil {
		return nil, errors.New("missing required argument 'UrlMap'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["quicOverride"] = nil
		inputs["sslCertificates"] = nil
		inputs["sslPolicy"] = nil
		inputs["urlMap"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["quicOverride"] = args.QuicOverride
		inputs["sslCertificates"] = args.SslCertificates
		inputs["sslPolicy"] = args.SslPolicy
		inputs["urlMap"] = args.UrlMap
	}
	inputs["creationTimestamp"] = nil
	inputs["proxyId"] = nil
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:compute/targetHttpsProxy:TargetHttpsProxy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TargetHttpsProxy{s: s}, nil
}

// GetTargetHttpsProxy gets an existing TargetHttpsProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetHttpsProxy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *TargetHttpsProxyState, opts ...pulumi.ResourceOpt) (*TargetHttpsProxy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["description"] = state.Description
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["proxyId"] = state.ProxyId
		inputs["quicOverride"] = state.QuicOverride
		inputs["selfLink"] = state.SelfLink
		inputs["sslCertificates"] = state.SslCertificates
		inputs["sslPolicy"] = state.SslPolicy
		inputs["urlMap"] = state.UrlMap
	}
	s, err := ctx.ReadResource("gcp:compute/targetHttpsProxy:TargetHttpsProxy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TargetHttpsProxy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *TargetHttpsProxy) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *TargetHttpsProxy) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *TargetHttpsProxy) CreationTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

func (r *TargetHttpsProxy) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *TargetHttpsProxy) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *TargetHttpsProxy) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

func (r *TargetHttpsProxy) ProxyId() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["proxyId"])
}

func (r *TargetHttpsProxy) QuicOverride() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["quicOverride"])
}

// The URI of the created resource.
func (r *TargetHttpsProxy) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

func (r *TargetHttpsProxy) SslCertificates() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["sslCertificates"])
}

func (r *TargetHttpsProxy) SslPolicy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sslPolicy"])
}

func (r *TargetHttpsProxy) UrlMap() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["urlMap"])
}

// Input properties used for looking up and filtering TargetHttpsProxy resources.
type TargetHttpsProxyState struct {
	CreationTimestamp interface{}
	Description interface{}
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	ProxyId interface{}
	QuicOverride interface{}
	// The URI of the created resource.
	SelfLink interface{}
	SslCertificates interface{}
	SslPolicy interface{}
	UrlMap interface{}
}

// The set of arguments for constructing a TargetHttpsProxy resource.
type TargetHttpsProxyArgs struct {
	Description interface{}
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	QuicOverride interface{}
	SslCertificates interface{}
	SslPolicy interface{}
	UrlMap interface{}
}
