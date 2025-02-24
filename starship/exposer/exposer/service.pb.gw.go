// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: exposer/service.proto

/*
Package exposer is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package exposer

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Suppress "imported and not used" errors
var (
	_ codes.Code
	_ io.Reader
	_ status.Status
	_ = errors.New
	_ = runtime.String
	_ = utilities.NewDoubleArray
	_ = metadata.Join
)

func request_Exposer_GetNodeID_0(ctx context.Context, marshaler runtime.Marshaler, client ExposerClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := client.GetNodeID(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_Exposer_GetNodeID_0(ctx context.Context, marshaler runtime.Marshaler, server ExposerServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := server.GetNodeID(ctx, &protoReq)
	return msg, metadata, err
}

func request_Exposer_GetPubKey_0(ctx context.Context, marshaler runtime.Marshaler, client ExposerClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := client.GetPubKey(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_Exposer_GetPubKey_0(ctx context.Context, marshaler runtime.Marshaler, server ExposerServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := server.GetPubKey(ctx, &protoReq)
	return msg, metadata, err
}

func request_Exposer_GetGenesisFile_0(ctx context.Context, marshaler runtime.Marshaler, client ExposerClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := client.GetGenesisFile(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_Exposer_GetGenesisFile_0(ctx context.Context, marshaler runtime.Marshaler, server ExposerServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := server.GetGenesisFile(ctx, &protoReq)
	return msg, metadata, err
}

func request_Exposer_GetGenesisSSZ_0(ctx context.Context, marshaler runtime.Marshaler, client ExposerClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := client.GetGenesisSSZ(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_Exposer_GetGenesisSSZ_0(ctx context.Context, marshaler runtime.Marshaler, server ExposerServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := server.GetGenesisSSZ(ctx, &protoReq)
	return msg, metadata, err
}

func request_Exposer_GetConfigFile_0(ctx context.Context, marshaler runtime.Marshaler, client ExposerClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := client.GetConfigFile(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_Exposer_GetConfigFile_0(ctx context.Context, marshaler runtime.Marshaler, server ExposerServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := server.GetConfigFile(ctx, &protoReq)
	return msg, metadata, err
}

func request_Exposer_GetKeys_0(ctx context.Context, marshaler runtime.Marshaler, client ExposerClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := client.GetKeys(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_Exposer_GetKeys_0(ctx context.Context, marshaler runtime.Marshaler, server ExposerServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := server.GetKeys(ctx, &protoReq)
	return msg, metadata, err
}

func request_Exposer_GetPrivKey_0(ctx context.Context, marshaler runtime.Marshaler, client ExposerClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := client.GetPrivKey(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_Exposer_GetPrivKey_0(ctx context.Context, marshaler runtime.Marshaler, server ExposerServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := server.GetPrivKey(ctx, &protoReq)
	return msg, metadata, err
}

func request_Exposer_GetPrivValidatorState_0(ctx context.Context, marshaler runtime.Marshaler, client ExposerClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := client.GetPrivValidatorState(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_Exposer_GetPrivValidatorState_0(ctx context.Context, marshaler runtime.Marshaler, server ExposerServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := server.GetPrivValidatorState(ctx, &protoReq)
	return msg, metadata, err
}

func request_Exposer_GetNodeKey_0(ctx context.Context, marshaler runtime.Marshaler, client ExposerClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := client.GetNodeKey(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_Exposer_GetNodeKey_0(ctx context.Context, marshaler runtime.Marshaler, server ExposerServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := server.GetNodeKey(ctx, &protoReq)
	return msg, metadata, err
}

func request_Exposer_CreateChannel_0(ctx context.Context, marshaler runtime.Marshaler, client ExposerClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq RequestCreateChannel
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.CreateChannel(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_Exposer_CreateChannel_0(ctx context.Context, marshaler runtime.Marshaler, server ExposerServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq RequestCreateChannel
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := server.CreateChannel(ctx, &protoReq)
	return msg, metadata, err
}

// RegisterExposerHandlerServer registers the http handlers for service Exposer to "mux".
// UnaryRPC     :call ExposerServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterExposerHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterExposerHandlerServer(ctx context.Context, mux *runtime.ServeMux, server ExposerServer) error {
	mux.Handle(http.MethodGet, pattern_Exposer_GetNodeID_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/exposer.Exposer/GetNodeID", runtime.WithHTTPPathPattern("/node_id"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Exposer_GetNodeID_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Exposer_GetNodeID_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_Exposer_GetPubKey_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/exposer.Exposer/GetPubKey", runtime.WithHTTPPathPattern("/pub_key"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Exposer_GetPubKey_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Exposer_GetPubKey_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_Exposer_GetGenesisFile_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/exposer.Exposer/GetGenesisFile", runtime.WithHTTPPathPattern("/genesis"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Exposer_GetGenesisFile_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Exposer_GetGenesisFile_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_Exposer_GetGenesisSSZ_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/exposer.Exposer/GetGenesisSSZ", runtime.WithHTTPPathPattern("/genesis_ssz"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Exposer_GetGenesisSSZ_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Exposer_GetGenesisSSZ_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_Exposer_GetConfigFile_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/exposer.Exposer/GetConfigFile", runtime.WithHTTPPathPattern("/config"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Exposer_GetConfigFile_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Exposer_GetConfigFile_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_Exposer_GetKeys_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/exposer.Exposer/GetKeys", runtime.WithHTTPPathPattern("/keys"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Exposer_GetKeys_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Exposer_GetKeys_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_Exposer_GetPrivKey_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/exposer.Exposer/GetPrivKey", runtime.WithHTTPPathPattern("/priv_keys"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Exposer_GetPrivKey_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Exposer_GetPrivKey_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_Exposer_GetPrivValidatorState_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/exposer.Exposer/GetPrivValidatorState", runtime.WithHTTPPathPattern("/priv_validator_state"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Exposer_GetPrivValidatorState_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Exposer_GetPrivValidatorState_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_Exposer_GetNodeKey_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/exposer.Exposer/GetNodeKey", runtime.WithHTTPPathPattern("/node_key"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Exposer_GetNodeKey_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Exposer_GetNodeKey_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodPost, pattern_Exposer_CreateChannel_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/exposer.Exposer/CreateChannel", runtime.WithHTTPPathPattern("/create_channel"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Exposer_CreateChannel_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Exposer_CreateChannel_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})

	return nil
}

// RegisterExposerHandlerFromEndpoint is same as RegisterExposerHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterExposerHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()
	return RegisterExposerHandler(ctx, mux, conn)
}

// RegisterExposerHandler registers the http handlers for service Exposer to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterExposerHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterExposerHandlerClient(ctx, mux, NewExposerClient(conn))
}

// RegisterExposerHandlerClient registers the http handlers for service Exposer
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "ExposerClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "ExposerClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "ExposerClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterExposerHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ExposerClient) error {
	mux.Handle(http.MethodGet, pattern_Exposer_GetNodeID_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/exposer.Exposer/GetNodeID", runtime.WithHTTPPathPattern("/node_id"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Exposer_GetNodeID_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Exposer_GetNodeID_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_Exposer_GetPubKey_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/exposer.Exposer/GetPubKey", runtime.WithHTTPPathPattern("/pub_key"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Exposer_GetPubKey_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Exposer_GetPubKey_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_Exposer_GetGenesisFile_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/exposer.Exposer/GetGenesisFile", runtime.WithHTTPPathPattern("/genesis"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Exposer_GetGenesisFile_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Exposer_GetGenesisFile_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_Exposer_GetGenesisSSZ_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/exposer.Exposer/GetGenesisSSZ", runtime.WithHTTPPathPattern("/genesis_ssz"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Exposer_GetGenesisSSZ_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Exposer_GetGenesisSSZ_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_Exposer_GetConfigFile_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/exposer.Exposer/GetConfigFile", runtime.WithHTTPPathPattern("/config"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Exposer_GetConfigFile_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Exposer_GetConfigFile_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_Exposer_GetKeys_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/exposer.Exposer/GetKeys", runtime.WithHTTPPathPattern("/keys"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Exposer_GetKeys_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Exposer_GetKeys_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_Exposer_GetPrivKey_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/exposer.Exposer/GetPrivKey", runtime.WithHTTPPathPattern("/priv_keys"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Exposer_GetPrivKey_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Exposer_GetPrivKey_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_Exposer_GetPrivValidatorState_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/exposer.Exposer/GetPrivValidatorState", runtime.WithHTTPPathPattern("/priv_validator_state"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Exposer_GetPrivValidatorState_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Exposer_GetPrivValidatorState_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodGet, pattern_Exposer_GetNodeKey_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/exposer.Exposer/GetNodeKey", runtime.WithHTTPPathPattern("/node_key"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Exposer_GetNodeKey_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Exposer_GetNodeKey_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodPost, pattern_Exposer_CreateChannel_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/exposer.Exposer/CreateChannel", runtime.WithHTTPPathPattern("/create_channel"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Exposer_CreateChannel_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Exposer_CreateChannel_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	return nil
}

var (
	pattern_Exposer_GetNodeID_0             = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"node_id"}, ""))
	pattern_Exposer_GetPubKey_0             = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"pub_key"}, ""))
	pattern_Exposer_GetGenesisFile_0        = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"genesis"}, ""))
	pattern_Exposer_GetGenesisSSZ_0         = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"genesis_ssz"}, ""))
	pattern_Exposer_GetConfigFile_0         = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"config"}, ""))
	pattern_Exposer_GetKeys_0               = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"keys"}, ""))
	pattern_Exposer_GetPrivKey_0            = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"priv_keys"}, ""))
	pattern_Exposer_GetPrivValidatorState_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"priv_validator_state"}, ""))
	pattern_Exposer_GetNodeKey_0            = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"node_key"}, ""))
	pattern_Exposer_CreateChannel_0         = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"create_channel"}, ""))
)

var (
	forward_Exposer_GetNodeID_0             = runtime.ForwardResponseMessage
	forward_Exposer_GetPubKey_0             = runtime.ForwardResponseMessage
	forward_Exposer_GetGenesisFile_0        = runtime.ForwardResponseMessage
	forward_Exposer_GetGenesisSSZ_0         = runtime.ForwardResponseMessage
	forward_Exposer_GetConfigFile_0         = runtime.ForwardResponseMessage
	forward_Exposer_GetKeys_0               = runtime.ForwardResponseMessage
	forward_Exposer_GetPrivKey_0            = runtime.ForwardResponseMessage
	forward_Exposer_GetPrivValidatorState_0 = runtime.ForwardResponseMessage
	forward_Exposer_GetNodeKey_0            = runtime.ForwardResponseMessage
	forward_Exposer_CreateChannel_0         = runtime.ForwardResponseMessage
)
