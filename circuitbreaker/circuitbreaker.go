package circuitbreaker

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
)

/*
 * This post final hystirx implementation for error and response based return types to handle post request
 */

func HystrixPostAsynch(commandName string, url string, postParam *bytes.Buffer, headerMap map[string]string) (chan string, chan error) {
	log.Println("Call starts: HystrixPostAsynch")
	resultChan := make(chan string)
	errChan := make(chan error)
	hystrix.ConfigureCommand(commandName, returnCommandConfig(commandName))
	errChan = hystrix.Go(commandName, func() error {
		log.Println("Call starts: HystrixPostAsynch", url)

		req, err := http.NewRequest("POST", url, postParam)
		//Forward all the headers passed from header for outgoing request
		if len(headerMap) > 0 {
			for k, _ := range headerMap {
				req.Header.Set(k, headerMap[k])
			}
		}
		//req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			errChan <- err
			log.Fatal("Error: hystrix failed to read response", err)
			return nil
		}
		//reader must close
		defer resp.Body.Close()
		body, readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			errChan <- readErr
			log.Fatal("Error: hystrix failed to read response", err)
			return nil
		}
		resultChan <- string(body)
		return nil
	}, func(err error) error {
		errChan <- err
		log.Fatal("Error: hystrix failed to execute http request", err)
		return nil
	})
	log.Println("Call Ends: HystrixPostAsynch")
	return resultChan, errChan
}

func HystrixGetAsynch(commandName string, serviceUrl string, queryParams map[string][]string, headerMap map[string]string) (chan string, chan error) {
	log.Println("Call starts: HystrixGetAsynch")

	resultChan := make(chan string)
	errChan := make(chan error)

	hystrix.ConfigureCommand(commandName, returnCommandConfig(commandName))

	errChan = hystrix.Go(commandName, func() error {
		log.Println("Call starts: HystrixGetAsynch", queryParams)

		req, err := http.NewRequest("GET", serviceUrl, nil)

		q := req.URL.Query()
		q = queryParams
		req.URL.RawQuery = q.Encode()
		req.Header.Set("Content-Type", "application/json")
		//Forward all the headers from incoming requests
		if len(headerMap) > 0 {
			for k, _ := range headerMap {
				req.Header.Set(k, headerMap[k])
			}
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			errChan <- err
			log.Fatal("Error: hystrix failed to connect API", err)
			return nil
		}

		//reader must close
		defer resp.Body.Close()

		body, readErr := ioutil.ReadAll(resp.Body)

		if readErr != nil {
			errChan <- readErr
			log.Fatal("Error: hystrix failed to read response", err)
			return nil
		}

		resultChan <- string(body)
		return nil
	}, func(err error) error {
		errChan <- err
		log.Fatal("Error: hystrix failed to execute http request", err)
		return nil
	})

	log.Println("Call Ends: HystrixGetAsynch")
	return resultChan, errChan
}

func returnCommandConfig(commandName string) hystrix.CommandConfig {
	log.Println("Call starts: returnCommandConfig")

	var configComment hystrix.CommandConfig
	if commandName == "API" {
		configComment = hystrix.CommandConfig{
			Timeout: 1500,
		}
		return configComment
	} else if commandName == "somethingelse" {
		configComment = hystrix.CommandConfig{
			Timeout: 1000,
		}
		return configComment
	}
	return configComment
}
