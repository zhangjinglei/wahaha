package metadata

// metadata common key
const (

	// Network
	RemoteIP   = "remote_ip"
	RemotePort = "remote_port"
	ServerAddr = "server_addr"
	ClientAddr = "client_addr"

	// Router
	Cluster = "cluster"
	Color   = "color"

	// Trace
	Trace  = "trace"
	Caller = "caller"

	// Timeout
	Timeout = "timeout"

	// Dispatch
	CPUUsage = "cpu_usage"
	Errors   = "errors"
	Requests = "requests"

	// Mirror
	Mirror = "mirror"

	// Mid 外网账户用户id
	Mid = "mid" // NOTE: ！！！业务可重新修改key名！！！

	// Device 客户端信息
	Device = "device"

	// Criticality 重要性
	Criticality = "criticality"
)

var outgoingKey = map[string]struct{}{
	Color:       struct{}{},
	RemoteIP:    struct{}{},
	RemotePort:  struct{}{},
	Mirror:      struct{}{},
	Criticality: struct{}{},
}

var incomingKey = map[string]struct{}{
	Caller: struct{}{},
}

// IsOutgoingKey represent this key should propagate by rpc.
func IsOutgoingKey(key string) bool {
	_, ok := outgoingKey[key]
	return ok
}

// IsIncomingKey represent this key should extract from rpc metadata.
func IsIncomingKey(key string) (ok bool) {
	_, ok = outgoingKey[key]
	if ok {
		return
	}
	_, ok = incomingKey[key]
	return
}

var Incomingheaders = []string{
	// All applications should propagate x-request-id. This header is
	// included in access log statements and is used for consistent trace
	// sampling and log sampling decisions in Istio.
	"x-request-id",

	// Lightstep tracing header. Propagate this if you use lightstep tracing
	// in Istio (see
	// https://istio.io/latest/docs/tasks/observability/distributed-tracing/lightstep/)
	// Note: this should probably be changed to use B3 or W3C TRACE_CONTEXT.
	// Lightstep recommends using B3 or TRACE_CONTEXT and most application
	// libraries from lightstep do not support x-ot-span-context.
	"x-ot-span-context",

	// Datadog tracing header. Propagate these headers if you use Datadog
	// tracing.
	"x-datadog-trace-id",
	"x-datadog-parent-id",
	"x-datadog-sampling-priority",

	// W3C Trace Context. Compatible with OpenCensusAgent and Stackdriver Istio
	// configurations.
	"traceparent",
	"tracestate",

	// Cloud trace context. Compatible with OpenCensusAgent and Stackdriver Istio
	// configurations.
	"x-cloud-trace-context",

	// Grpc binary trace context. Compatible with OpenCensusAgent nad
	// Stackdriver Istio configurations.
	"grpc-trace-bin",

	// b3 trace headers. Compatible with Zipkin, OpenCensusAgent, and
	// Stackdriver Istio configurations. Commented out since they are
	// propagated by the OpenTracing tracer above.
	"x-b3-traceid",
	"x-b3-spanid",
	"x-b3-parentspanid",
	"x-b3-sampled",
	"x-b3-flags",

	// Application-specific headers to forward.
	"end-user",
	"user-agent",
}
