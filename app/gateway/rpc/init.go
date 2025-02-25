package rpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	"github.com/CocaineCong/tangseng/config"
	"github.com/CocaineCong/tangseng/idl/pb/favorite"
	"github.com/CocaineCong/tangseng/idl/pb/index_platform"
	"github.com/CocaineCong/tangseng/idl/pb/search_engine"
	"github.com/CocaineCong/tangseng/idl/pb/user"
	"github.com/CocaineCong/tangseng/pkg/discovery"
)

var (
	Register   *discovery.Resolver
	ctx        context.Context
	CancelFunc context.CancelFunc

	UserClient          user.UserServiceClient
	FavoriteClient      favorite.FavoritesServiceClient
	SearchEngineClient  search_engine.SearchEngineServiceClient
	IndexPlatformClient index_platform.IndexPlatformServiceClient
)

func Init() {
	Register = discovery.NewResolver([]string{config.Conf.Etcd.Address}, logrus.New())
	resolver.Register(Register)
	ctx, CancelFunc = context.WithTimeout(context.Background(), 3*time.Second)

	defer Register.Close()
	initClient(config.Conf.Domain["user"].Name, &UserClient)
	initClient(config.Conf.Domain["favorite"].Name, &FavoriteClient)
	initClient(config.Conf.Domain["search_engine"].Name, &SearchEngineClient)
	initClient(config.Conf.Domain["index_platform"].Name, &IndexPlatformClient)
}

func initClient(serviceName string, client interface{}) {
	conn, err := connectServer(serviceName)

	if err != nil {
		panic(err)
	}

	switch c := client.(type) {
	case *user.UserServiceClient:
		*c = user.NewUserServiceClient(conn)
	case *favorite.FavoritesServiceClient:
		*c = favorite.NewFavoritesServiceClient(conn)
	case *search_engine.SearchEngineServiceClient:
		*c = search_engine.NewSearchEngineServiceClient(conn)
	case *index_platform.IndexPlatformServiceClient:
		*c = index_platform.NewIndexPlatformServiceClient(conn)
	default:
		panic("unsupported woker type")
	}
}

func connectServer(serviceName string) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	addr := fmt.Sprintf("%s:///%s", Register.Scheme(), serviceName)

	// Load balance
	if config.Conf.Services[serviceName].LoadBalance {
		log.Printf("load balance enabled for %s\n", serviceName)
		opts = append(opts, grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, "round_robin")))
	}

	conn, err = grpc.DialContext(ctx, addr, opts...)
	return
}
