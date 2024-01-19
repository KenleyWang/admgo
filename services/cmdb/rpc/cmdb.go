package main

import (
	"flag"
	"fmt"

	"github.com/admgo/admgo/services/cmdb/rpc/internal/config"
	attributetypeServer "github.com/admgo/admgo/services/cmdb/rpc/internal/server/attributetype"
	resourceServer "github.com/admgo/admgo/services/cmdb/rpc/internal/server/resource"
	resourceattributeServer "github.com/admgo/admgo/services/cmdb/rpc/internal/server/resourceattribute"
	resourcegroupServer "github.com/admgo/admgo/services/cmdb/rpc/internal/server/resourcegroup"
	resourcemodelServer "github.com/admgo/admgo/services/cmdb/rpc/internal/server/resourcemodel"
	resourcetypeServer "github.com/admgo/admgo/services/cmdb/rpc/internal/server/resourcetype"
	"github.com/admgo/admgo/services/cmdb/rpc/internal/svc"
	"github.com/admgo/admgo/services/cmdb/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/cmdb.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterResourceGroupServer(grpcServer, resourcegroupServer.NewResourceGroupServer(ctx))
		pb.RegisterResourceTypeServer(grpcServer, resourcetypeServer.NewResourceTypeServer(ctx))
		pb.RegisterResourceModelServer(grpcServer, resourcemodelServer.NewResourceModelServer(ctx))
		pb.RegisterResourceAttributeServer(grpcServer, resourceattributeServer.NewResourceAttributeServer(ctx))
		pb.RegisterAttributeTypeServer(grpcServer, attributetypeServer.NewAttributeTypeServer(ctx))
		pb.RegisterResourceServer(grpcServer, resourceServer.NewResourceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
