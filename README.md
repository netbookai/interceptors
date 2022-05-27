# Interceptors

<a href="https://pkg.go.dev/github.com/netbook-ai/interceptors"><img src="https://pkg.go.dev/badge/github.com/netbook-ai/interceptors.svg" alt="Go Reference"></a>

gRPC interceptors used in netbook services

It injects the following interceptor

* kit Interceptor - inject method name to the context
* loggingInterceptor - log request and response data, duration of the call
* recoveryInterceptor - recover from any API panics gracefully and logs error

## Installation

```
    go get -u github.com/netbook-ai/interceptors@v0.1.0
```

You may not be able to access the repo with netbook-devs path in GOPRIVATE,  update it as follows

```
export GOPRIVATE=gitlab.com/*
```

> can update it in your profile settings (.bashrc, .zshrc)

## Usage

Register interceptors when setting up gRPC server in application

### Using default interceptors

```
//get Interceptor interface
interceptor := interceptors.NewInterceptor("spawnerservice", logger)

//Calling Get will return ServerOptions
baseServer := grpc.NewServer(interceptor.Get())
```

### Using user defined interceptors
```
func InstrumentationInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		IncRequest(getMethod(info))
		return handler(ctx, req)
	}
}

//get Interceptor interface
interceptor := interceptors.NewInterceptor(
				"spawnerservice",
 				logger,
                  		interceptors.WithInterceptor(InstrumentationInterceptor()))

//Calling Get will return ServerOptions, which is passed to NewServer
baseServer := grpc.NewServer(interceptor.Get())
```

### skip sensitive methods from logging interceptor

You can skip sensitive API request and response from request-response logging interceptor.
```
//get Interceptor interface
var skipMethods = []string{"ReadCredential", "WriteCredential"}

interceptor := interceptors.NewInterceptor(
				"spawnerservice",
 				logger,
				interceptors.WithSkipMethod(skipMethods))

//Calling Get will return ServerOptions, which is passed to NewServer
baseServer := grpc.NewServer(interceptor.Get())
```

## Support
Raise a issue, we will get back to you.

## Roadmap

## Contributing

## Authors and acknowledgment

* Nishanth Shetty <nishanthspshetty@gmail.com>

## License

   GNU GENERAL PUBLIC LICENSE


## Project status
    
    
