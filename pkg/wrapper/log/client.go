package log

import (
	"context"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"

	"github.com/kzmake/traefik-custom-forward-auth/pkg/logger"
)

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	logger.Debugf("ClientCaller: Service: %s, Method: %s, ctx: %v", req.Service(), req.Method(), md)
	return c.Client.Call(ctx, req, rsp, opts...)
}

func (c *clientWrapper) Publish(ctx context.Context, p client.Message, opts ...client.PublishOption) error {
	md, _ := metadata.FromContext(ctx)
	logger.Debugf("ClientPublisher: Topic: %s, ContentType: %s, Payload: %v, ctx: %v", p.Topic(), p.ContentType(), p.Payload(), md)
	return c.Client.Publish(ctx, p, opts...)
}
