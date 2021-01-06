# soa-platform-go/log

This package provides global logging functions for the service.

`log`'s implementation is request-centric and context-aware. It uses `context.Context` to store request-scoped metadata and print it later.

### What is 'request-centric' logger?
As service processes a request, it inevitably creates some metadata that should be printed with every log entry related to the request.

Metadata can be any info related to a single request - trace ID, user ID, request parameter, etc. 

Once you add any metadata to a context, `log` will print it with every entry. Example:

``` go
func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
    // first log
	log.Info(ctx,"serving handler","key","value")

	span := spans.NewChild(r)
    // embed metadata
    ctx = log.With(r.Context,"trace-id",span.TraceID)
    
    log.Info(ctx,"span extracted")

	userID := getUserID(r)
    // embed more metadata
    ctx = log.With(ctx,"user-id",userID)

	param := r.URL.Query().Get("param1")
	foo(ctx,userID,param)
}

func foo(ctx context.Context,param string) {
	go doSomething(ctx)
}

func doSomething(ctx context.Context) {
	log.Info(ctx,"doing stuff","foo","bar")
}
```

This code will yield following logs with TraceID=ABC and UserID=123:
``` 
serving handler, key=value
span extracted, trace-id=ABC
doing stuff, trace-id=ABC user-id=123 foo=bar
```

Note that code in `doSomething` does not know anything about the metadata.
This way it's easier to make code SOLID while tying up related log entries together.
