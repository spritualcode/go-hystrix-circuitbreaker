# go-hystrix-circuitebreaker example

This is an example to demonstrate go-hystrix configuration for third party API calls. As mentioned, I use go-hystrix(https://github.com/afex/hystrix-go). But you can replcase package with anyother and it should work fine based on paramters defined....

In order to run this example...

1) Go to  ..handler/TestingHandler.go

Pass thirdParty URL at line number 21 of above function

Also make sure, you pass correct struct with post params, if API expect any post params.

2) Please make sure, you have glide installed

3) Run $ glide update

4) Run $ go build

5) $ ./go-hystrix-circuitebreaker

6) open browser localhost:8080/example

I use gorilla Mux for the routing.

In order to configure Post request based thirdparty calls, please configure following in your handler
func HystrixPostAsynch(commandName string, url string, postParam *bytes.Buffer, headerMap map[string]string)

In order to configure Get request based thirdparty calls, please configure following in your handler 
func HystrixPostAsynch(commandName string, url string, postParam *bytes.Buffer, headerMap map[string]string)

If you want to provide a specific configuration like timeut, concurrent calls to API, then create a command name and add it to following function
func returnCommandConfig(commandName string) 

If you have any question, feel free to contact me.....
