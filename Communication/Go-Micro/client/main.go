package  main

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	proto "github.com/deividroger/advanced-cloud-native-go/Communication/Go-Micro/proto"
	micro "github.com/micro/go-micro"
	breaker "github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"net"
	"net/http"
	"time"
)

func main()  {
		service := micro.NewService(
			micro.Name("greeter"),
			micro.Version("latest"),
			micro.Metadata(map[string]string{
				"type": "helloworld",
			}),
		)
		service.Init(
				micro.WrapClient(breaker.NewClientWrapper()),
			)
		hystrix.DefaultVolumeThreshold= 3
		hystrix.DefaultErrorPercentThreshold =75
		hystrix.DefaultTimeout = 500
		hystrix.DefaultSleepWindow = 3500

		hysterixStreamHandler  := hystrix.NewStreamHandler()
		hysterixStreamHandler.Start()
		go http.ListenAndServe(net.JoinHostPort("","8081"),hysterixStreamHandler)

		greeter := proto.NewGreeterClient("greeter",service.Client())

		callEvery(3*time.Second,greeter,hello)
}

func hello(t time.Time, greeter proto.GreeterClient)  {
	rsp, err := greeter.Hello(context.TODO(),&proto.HelloRequest{Name: "Deivid calling at +" + t.String()})

	if err != nil{
		if err.Error() == "hysterix: timeout" {
			fmt.Printf("%s. Insert fallback logic here.\n",err.Error())
		}else {
			fmt.Println(err.Error())
		}
	}

	fmt.Printf("%s\n",rsp.Greeting)

}

func callEvery(d time.Duration, greeter proto.GreeterClient, f func(time2 time.Time,client proto.GreeterClient))  {
	for x := range time.Tick(d) {
		f(x,greeter)
	}
}