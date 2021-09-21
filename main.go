package main

import (
	"cart/domain/repository"
	service2 "cart/domain/service"
	"cart/handler"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/yann1989/common"
	"os"

	cart "cart/proto/cart"
	_ "github.com/go-sql-driver/mysql"
)

const QPS = 100 //每秒处理数量

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}
	//注册中心
	consulRegistry := common.GetConsulRegistry("127.0.0.1", 8500)



	// 链路追踪
	tracer, closer, err := common.NewTracer("micro.cart", "127.0.0.1:6831")
	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	//获取mysql配置, 前面配置中心设置了可以不带前缀, 这里不需要再带前缀了
	mysqlInfo, err := common.GetMysqlFromConsul(consulConfig, "mysql")
	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}
	//连接mysql
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlInfo.User,
		mysqlInfo.Pwd, mysqlInfo.Host, mysqlInfo.Port, mysqlInfo.Database)
	mysqlDb, err := gorm.Open("mysql", dns)
	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}
	defer mysqlDb.Close()
	mysqlDb.SingularTable(true)


	// New Service
	srv := micro.NewService(
		micro.Name("micro.cart"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:9996"),
		micro.Registry(consulRegistry),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		//添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)
	// 初始化服务
	srv.Init()

	// Register Handler
	productRepository := repository.NewCartRepository(mysqlDb)
	productRepository.InitTable()
	//注册handel
	err = cart.RegisterCartHandler(srv.Server(), handler.NewCart(service2.NewCartService(productRepository)))
	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
