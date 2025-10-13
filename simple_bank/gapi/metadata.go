package gapi

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpc-client"
	UserAgentHeader            = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

func (server *Server) extractMetada(ctx context.Context) *Metadata {
	mtdt := &Metadata{}
	if p, ok := peer.FromContext(ctx); ok {

		// log.Printf("Client IP: %s", p.Addr.String())
		mtdt.ClientIP = p.Addr.String()
	}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		// log.Printf("md: %+v\n", md)
		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
			log.Printf("md: %+v\n", mtdt.UserAgent)
		}
		if userAgents := md.Get(UserAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
			log.Printf("md: %+v\n", mtdt.UserAgent)
		}
		// if clientIPs := md.Get(xForwardedForHeader); len(clientIPs) > 0 {
		// 	mtdt.ClientIP = clientIPs[0]
		// }
		log.Printf("md: %+v\n", mtdt)
	}

	return mtdt
}
