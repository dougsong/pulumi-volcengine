// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage listener
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/clb"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooZones, err := ecs.Zones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			fooVpc, err := vpc.NewVpc(ctx, "fooVpc", &vpc.VpcArgs{
//				VpcName:   pulumi.String("acc-test-vpc"),
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			fooSubnet, err := vpc.NewSubnet(ctx, "fooSubnet", &vpc.SubnetArgs{
//				SubnetName: pulumi.String("acc-test-subnet"),
//				CidrBlock:  pulumi.String("172.16.0.0/24"),
//				ZoneId:     *pulumi.String(fooZones.Zones[0].Id),
//				VpcId:      fooVpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			fooClb, err := clb.NewClb(ctx, "fooClb", &clb.ClbArgs{
//				Type:             pulumi.String("public"),
//				SubnetId:         fooSubnet.ID(),
//				LoadBalancerSpec: pulumi.String("small_1"),
//				Description:      pulumi.String("acc0Demo"),
//				LoadBalancerName: pulumi.String("acc-test-create"),
//				EipBillingConfig: &clb.ClbEipBillingConfigArgs{
//					Isp:            pulumi.String("BGP"),
//					EipBillingType: pulumi.String("PostPaidByBandwidth"),
//					Bandwidth:      pulumi.Int(1),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			fooServerGroup, err := clb.NewServerGroup(ctx, "fooServerGroup", &clb.ServerGroupArgs{
//				LoadBalancerId:  fooClb.ID(),
//				ServerGroupName: pulumi.String("acc-test-create"),
//				Description:     pulumi.String("hello demo11"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = clb.NewListener(ctx, "fooListener", &clb.ListenerArgs{
//				LoadBalancerId: fooClb.ID(),
//				ListenerName:   pulumi.String("acc-test-listener"),
//				Protocol:       pulumi.String("HTTP"),
//				Port:           pulumi.Int(90),
//				ServerGroupId:  fooServerGroup.ID(),
//				HealthCheck: &clb.ListenerHealthCheckArgs{
//					Enabled:            pulumi.String("on"),
//					Interval:           pulumi.Int(10),
//					Timeout:            pulumi.Int(3),
//					HealthyThreshold:   pulumi.Int(5),
//					UnHealthyThreshold: pulumi.Int(2),
//					Domain:             pulumi.String("volcengine.com"),
//					HttpCode:           pulumi.String("http_2xx"),
//					Method:             pulumi.String("GET"),
//					Uri:                pulumi.String("/"),
//				},
//				Enabled: pulumi.String("on"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Listener can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:clb/listener:Listener default lsn-273yv0mhs5xj47fap8sehiiso
//
// ```
type Listener struct {
	pulumi.CustomResourceState

	// The id list of the Acl.
	AclIds pulumi.StringArrayOutput `pulumi:"aclIds"`
	// The enable status of Acl. Optional choice contains `on`, `off`.
	AclStatus pulumi.StringOutput `pulumi:"aclStatus"`
	// The type of the Acl. Optional choice contains `white`, `black`.
	AclType pulumi.StringOutput `pulumi:"aclType"`
	// The certificate id associated with the listener.
	CertificateId pulumi.StringPtrOutput `pulumi:"certificateId"`
	// The description of the Listener.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The enable status of the Listener. Optional choice contains `on`, `off`.
	Enabled pulumi.StringOutput `pulumi:"enabled"`
	// The connection timeout of the Listener.
	EstablishedTimeout pulumi.IntOutput `pulumi:"establishedTimeout"`
	// The config of health check.
	HealthCheck ListenerHealthCheckOutput `pulumi:"healthCheck"`
	// The ID of the Listener.
	ListenerId pulumi.StringOutput `pulumi:"listenerId"`
	// The name of the Listener.
	ListenerName pulumi.StringOutput `pulumi:"listenerName"`
	// The region of the request.
	LoadBalancerId pulumi.StringOutput `pulumi:"loadBalancerId"`
	// The port receiving request of the Listener, the value range in 1~65535.
	Port pulumi.IntOutput `pulumi:"port"`
	// The protocol of the Listener. Optional choice contains `TCP`, `UDP`, `HTTP`, `HTTPS`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The scheduling algorithm of the Listener. Optional choice contains `wrr`, `wlc`, `sh`.
	Scheduler pulumi.StringOutput `pulumi:"scheduler"`
	// The server group id associated with the listener.
	ServerGroupId pulumi.StringOutput `pulumi:"serverGroupId"`
}

// NewListener registers a new resource with the given unique name, arguments, and options.
func NewListener(ctx *pulumi.Context,
	name string, args *ListenerArgs, opts ...pulumi.ResourceOption) (*Listener, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoadBalancerId == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerId'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ServerGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ServerGroupId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Listener
	err := ctx.RegisterResource("volcengine:clb/listener:Listener", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListener gets an existing Listener resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListener(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListenerState, opts ...pulumi.ResourceOption) (*Listener, error) {
	var resource Listener
	err := ctx.ReadResource("volcengine:clb/listener:Listener", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Listener resources.
type listenerState struct {
	// The id list of the Acl.
	AclIds []string `pulumi:"aclIds"`
	// The enable status of Acl. Optional choice contains `on`, `off`.
	AclStatus *string `pulumi:"aclStatus"`
	// The type of the Acl. Optional choice contains `white`, `black`.
	AclType *string `pulumi:"aclType"`
	// The certificate id associated with the listener.
	CertificateId *string `pulumi:"certificateId"`
	// The description of the Listener.
	Description *string `pulumi:"description"`
	// The enable status of the Listener. Optional choice contains `on`, `off`.
	Enabled *string `pulumi:"enabled"`
	// The connection timeout of the Listener.
	EstablishedTimeout *int `pulumi:"establishedTimeout"`
	// The config of health check.
	HealthCheck *ListenerHealthCheck `pulumi:"healthCheck"`
	// The ID of the Listener.
	ListenerId *string `pulumi:"listenerId"`
	// The name of the Listener.
	ListenerName *string `pulumi:"listenerName"`
	// The region of the request.
	LoadBalancerId *string `pulumi:"loadBalancerId"`
	// The port receiving request of the Listener, the value range in 1~65535.
	Port *int `pulumi:"port"`
	// The protocol of the Listener. Optional choice contains `TCP`, `UDP`, `HTTP`, `HTTPS`.
	Protocol *string `pulumi:"protocol"`
	// The scheduling algorithm of the Listener. Optional choice contains `wrr`, `wlc`, `sh`.
	Scheduler *string `pulumi:"scheduler"`
	// The server group id associated with the listener.
	ServerGroupId *string `pulumi:"serverGroupId"`
}

type ListenerState struct {
	// The id list of the Acl.
	AclIds pulumi.StringArrayInput
	// The enable status of Acl. Optional choice contains `on`, `off`.
	AclStatus pulumi.StringPtrInput
	// The type of the Acl. Optional choice contains `white`, `black`.
	AclType pulumi.StringPtrInput
	// The certificate id associated with the listener.
	CertificateId pulumi.StringPtrInput
	// The description of the Listener.
	Description pulumi.StringPtrInput
	// The enable status of the Listener. Optional choice contains `on`, `off`.
	Enabled pulumi.StringPtrInput
	// The connection timeout of the Listener.
	EstablishedTimeout pulumi.IntPtrInput
	// The config of health check.
	HealthCheck ListenerHealthCheckPtrInput
	// The ID of the Listener.
	ListenerId pulumi.StringPtrInput
	// The name of the Listener.
	ListenerName pulumi.StringPtrInput
	// The region of the request.
	LoadBalancerId pulumi.StringPtrInput
	// The port receiving request of the Listener, the value range in 1~65535.
	Port pulumi.IntPtrInput
	// The protocol of the Listener. Optional choice contains `TCP`, `UDP`, `HTTP`, `HTTPS`.
	Protocol pulumi.StringPtrInput
	// The scheduling algorithm of the Listener. Optional choice contains `wrr`, `wlc`, `sh`.
	Scheduler pulumi.StringPtrInput
	// The server group id associated with the listener.
	ServerGroupId pulumi.StringPtrInput
}

func (ListenerState) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerState)(nil)).Elem()
}

type listenerArgs struct {
	// The id list of the Acl.
	AclIds []string `pulumi:"aclIds"`
	// The enable status of Acl. Optional choice contains `on`, `off`.
	AclStatus *string `pulumi:"aclStatus"`
	// The type of the Acl. Optional choice contains `white`, `black`.
	AclType *string `pulumi:"aclType"`
	// The certificate id associated with the listener.
	CertificateId *string `pulumi:"certificateId"`
	// The description of the Listener.
	Description *string `pulumi:"description"`
	// The enable status of the Listener. Optional choice contains `on`, `off`.
	Enabled *string `pulumi:"enabled"`
	// The connection timeout of the Listener.
	EstablishedTimeout *int `pulumi:"establishedTimeout"`
	// The config of health check.
	HealthCheck *ListenerHealthCheck `pulumi:"healthCheck"`
	// The name of the Listener.
	ListenerName *string `pulumi:"listenerName"`
	// The region of the request.
	LoadBalancerId string `pulumi:"loadBalancerId"`
	// The port receiving request of the Listener, the value range in 1~65535.
	Port int `pulumi:"port"`
	// The protocol of the Listener. Optional choice contains `TCP`, `UDP`, `HTTP`, `HTTPS`.
	Protocol string `pulumi:"protocol"`
	// The scheduling algorithm of the Listener. Optional choice contains `wrr`, `wlc`, `sh`.
	Scheduler *string `pulumi:"scheduler"`
	// The server group id associated with the listener.
	ServerGroupId string `pulumi:"serverGroupId"`
}

// The set of arguments for constructing a Listener resource.
type ListenerArgs struct {
	// The id list of the Acl.
	AclIds pulumi.StringArrayInput
	// The enable status of Acl. Optional choice contains `on`, `off`.
	AclStatus pulumi.StringPtrInput
	// The type of the Acl. Optional choice contains `white`, `black`.
	AclType pulumi.StringPtrInput
	// The certificate id associated with the listener.
	CertificateId pulumi.StringPtrInput
	// The description of the Listener.
	Description pulumi.StringPtrInput
	// The enable status of the Listener. Optional choice contains `on`, `off`.
	Enabled pulumi.StringPtrInput
	// The connection timeout of the Listener.
	EstablishedTimeout pulumi.IntPtrInput
	// The config of health check.
	HealthCheck ListenerHealthCheckPtrInput
	// The name of the Listener.
	ListenerName pulumi.StringPtrInput
	// The region of the request.
	LoadBalancerId pulumi.StringInput
	// The port receiving request of the Listener, the value range in 1~65535.
	Port pulumi.IntInput
	// The protocol of the Listener. Optional choice contains `TCP`, `UDP`, `HTTP`, `HTTPS`.
	Protocol pulumi.StringInput
	// The scheduling algorithm of the Listener. Optional choice contains `wrr`, `wlc`, `sh`.
	Scheduler pulumi.StringPtrInput
	// The server group id associated with the listener.
	ServerGroupId pulumi.StringInput
}

func (ListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerArgs)(nil)).Elem()
}

type ListenerInput interface {
	pulumi.Input

	ToListenerOutput() ListenerOutput
	ToListenerOutputWithContext(ctx context.Context) ListenerOutput
}

func (*Listener) ElementType() reflect.Type {
	return reflect.TypeOf((**Listener)(nil)).Elem()
}

func (i *Listener) ToListenerOutput() ListenerOutput {
	return i.ToListenerOutputWithContext(context.Background())
}

func (i *Listener) ToListenerOutputWithContext(ctx context.Context) ListenerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerOutput)
}

// ListenerArrayInput is an input type that accepts ListenerArray and ListenerArrayOutput values.
// You can construct a concrete instance of `ListenerArrayInput` via:
//
//	ListenerArray{ ListenerArgs{...} }
type ListenerArrayInput interface {
	pulumi.Input

	ToListenerArrayOutput() ListenerArrayOutput
	ToListenerArrayOutputWithContext(context.Context) ListenerArrayOutput
}

type ListenerArray []ListenerInput

func (ListenerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Listener)(nil)).Elem()
}

func (i ListenerArray) ToListenerArrayOutput() ListenerArrayOutput {
	return i.ToListenerArrayOutputWithContext(context.Background())
}

func (i ListenerArray) ToListenerArrayOutputWithContext(ctx context.Context) ListenerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerArrayOutput)
}

// ListenerMapInput is an input type that accepts ListenerMap and ListenerMapOutput values.
// You can construct a concrete instance of `ListenerMapInput` via:
//
//	ListenerMap{ "key": ListenerArgs{...} }
type ListenerMapInput interface {
	pulumi.Input

	ToListenerMapOutput() ListenerMapOutput
	ToListenerMapOutputWithContext(context.Context) ListenerMapOutput
}

type ListenerMap map[string]ListenerInput

func (ListenerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Listener)(nil)).Elem()
}

func (i ListenerMap) ToListenerMapOutput() ListenerMapOutput {
	return i.ToListenerMapOutputWithContext(context.Background())
}

func (i ListenerMap) ToListenerMapOutputWithContext(ctx context.Context) ListenerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerMapOutput)
}

type ListenerOutput struct{ *pulumi.OutputState }

func (ListenerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Listener)(nil)).Elem()
}

func (o ListenerOutput) ToListenerOutput() ListenerOutput {
	return o
}

func (o ListenerOutput) ToListenerOutputWithContext(ctx context.Context) ListenerOutput {
	return o
}

// The id list of the Acl.
func (o ListenerOutput) AclIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringArrayOutput { return v.AclIds }).(pulumi.StringArrayOutput)
}

// The enable status of Acl. Optional choice contains `on`, `off`.
func (o ListenerOutput) AclStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.AclStatus }).(pulumi.StringOutput)
}

// The type of the Acl. Optional choice contains `white`, `black`.
func (o ListenerOutput) AclType() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.AclType }).(pulumi.StringOutput)
}

// The certificate id associated with the listener.
func (o ListenerOutput) CertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringPtrOutput { return v.CertificateId }).(pulumi.StringPtrOutput)
}

// The description of the Listener.
func (o ListenerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The enable status of the Listener. Optional choice contains `on`, `off`.
func (o ListenerOutput) Enabled() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.Enabled }).(pulumi.StringOutput)
}

// The connection timeout of the Listener.
func (o ListenerOutput) EstablishedTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Listener) pulumi.IntOutput { return v.EstablishedTimeout }).(pulumi.IntOutput)
}

// The config of health check.
func (o ListenerOutput) HealthCheck() ListenerHealthCheckOutput {
	return o.ApplyT(func(v *Listener) ListenerHealthCheckOutput { return v.HealthCheck }).(ListenerHealthCheckOutput)
}

// The ID of the Listener.
func (o ListenerOutput) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.ListenerId }).(pulumi.StringOutput)
}

// The name of the Listener.
func (o ListenerOutput) ListenerName() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.ListenerName }).(pulumi.StringOutput)
}

// The region of the request.
func (o ListenerOutput) LoadBalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.LoadBalancerId }).(pulumi.StringOutput)
}

// The port receiving request of the Listener, the value range in 1~65535.
func (o ListenerOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Listener) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// The protocol of the Listener. Optional choice contains `TCP`, `UDP`, `HTTP`, `HTTPS`.
func (o ListenerOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// The scheduling algorithm of the Listener. Optional choice contains `wrr`, `wlc`, `sh`.
func (o ListenerOutput) Scheduler() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.Scheduler }).(pulumi.StringOutput)
}

// The server group id associated with the listener.
func (o ListenerOutput) ServerGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Listener) pulumi.StringOutput { return v.ServerGroupId }).(pulumi.StringOutput)
}

type ListenerArrayOutput struct{ *pulumi.OutputState }

func (ListenerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Listener)(nil)).Elem()
}

func (o ListenerArrayOutput) ToListenerArrayOutput() ListenerArrayOutput {
	return o
}

func (o ListenerArrayOutput) ToListenerArrayOutputWithContext(ctx context.Context) ListenerArrayOutput {
	return o
}

func (o ListenerArrayOutput) Index(i pulumi.IntInput) ListenerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Listener {
		return vs[0].([]*Listener)[vs[1].(int)]
	}).(ListenerOutput)
}

type ListenerMapOutput struct{ *pulumi.OutputState }

func (ListenerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Listener)(nil)).Elem()
}

func (o ListenerMapOutput) ToListenerMapOutput() ListenerMapOutput {
	return o
}

func (o ListenerMapOutput) ToListenerMapOutputWithContext(ctx context.Context) ListenerMapOutput {
	return o
}

func (o ListenerMapOutput) MapIndex(k pulumi.StringInput) ListenerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Listener {
		return vs[0].(map[string]*Listener)[vs[1].(string)]
	}).(ListenerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerInput)(nil)).Elem(), &Listener{})
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerArrayInput)(nil)).Elem(), ListenerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ListenerMapInput)(nil)).Elem(), ListenerMap{})
	pulumi.RegisterOutputType(ListenerOutput{})
	pulumi.RegisterOutputType(ListenerArrayOutput{})
	pulumi.RegisterOutputType(ListenerMapOutput{})
}
