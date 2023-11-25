package ipdataclient

import (
	"fmt"
	"reflect"

	"github.com/s-yakubovskiy/whereami/config"
	"github.com/s-yakubovskiy/whereami/internal/contracts"

	ipdata "github.com/ipdata/go"
)

type IPDataClient struct {
	url     string
	api_key string
	client  ipdata.Client
}

func NewIPDataClient(providerConfig config.ProviderConfig) (*IPDataClient, error) {
	ipd, err := ipdata.NewClient(providerConfig.APIKey)
	return &IPDataClient{
		url:     providerConfig.URL,
		api_key: providerConfig.APIKey,
		client:  ipd,
	}, err
}

func (l *IPDataClient) GetLocation(ip string) (*contracts.Location, error) {
	data, err := l.client.Lookup(ip)
	if err != nil {
		return nil, err
	}
	return ConvertIPToLocation(data)
}

func printStruct(s interface{}) {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	t := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := t.Field(i)
		value := val.Field(i)
		fmt.Printf("%s: %v\n", field.Name, value.Interface())
	}
}
