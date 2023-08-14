/*
Copyright © 2023 libraria-app
*/
package librariacli

import (
	"fmt"

	"github.com/libraria-app/cli/internal/domains"
)

var apiUrl = "http://18.194.168.63:8080/api" // overriden with flags on the build stage

type Lcli struct {
	service domains.Service
}

func New() (*Lcli, error) {
	if apiUrl == "" {
		return nil, fmt.Errorf("base api url is empty")
	}

	return &Lcli{}, nil
}

func (l *Lcli) GetService() domains.Service {
	if l.service == nil {
		l.service = domains.NewService(apiUrl)
	}

	return l.service
}
