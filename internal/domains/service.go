/*
Copyright © 2023 libraria-app
*/
package domains

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Service interface {
	ExportTerms(apiKey, projectId, language, format string) ([]byte, error)
}

type service struct {
	apiUrl string
}

func NewService(apiUrl string) Service {
	return &service{
		apiUrl: apiUrl,
	}
}

func (s *service) ExportTerms(apiKey, projectId, language, format string) ([]byte, error) {
	query := url.Values{
		"lang":   {language},
		"format": {format},
	}
	uri := fmt.Sprintf("%s/cli/projects/%s/terms?%s", s.apiUrl, projectId, query.Encode())
	if _, err := url.Parse(uri); err != nil {
		return nil, fmt.Errorf("generate export terms url: %w", err)
	}

	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, uri, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("create export terms request: %w", err)
	}
	request.Header.Set("Authorization", apiKey)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("send export terms request: %w", err)
	}
	defer response.Body.Close()
	resBody, _ := io.ReadAll(response.Body)
	if response.StatusCode == http.StatusOK {
		return resBody, nil
	}
	responseObj := struct {
		Message string
	}{}
	if err := json.Unmarshal(resBody, &responseObj); err != nil {
		return nil, fmt.Errorf("parse service response: %w", err)
	}

	return nil, errors.New(fmt.Sprintf("fetch terms from request response: %s - %s", response.Status, responseObj.Message))
}
