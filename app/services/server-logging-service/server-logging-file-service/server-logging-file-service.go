package server_logging_file_service

import (
	"bufio"
	"container/list"
	"errors"
	"fmt"
	"os"
	"simple-hosting/controller/app/settings"
	"strconv"

	collection_utils "simple-hosting/go-commons/tools/sequence"
)

func GetLogsOnPage(request GetLogsOnPageRequest, config settings.ServiceSettings) (GetLogsOnPageResponse, error) {
	if request.Page < 0 {
		return GetLogsOnPageResponse{}, errors.New("Invalid page number, must be greater than 0 or equal, got: " + strconv.Itoa(int(request.Page)))
	}
	pageSize := config.App.Services.ServerLogging.PageSize
	file, err := os.Open(config.App.Services.ServerLogging.LogFile)
	if err != nil {
		fmt.Printf("Could not open log file: %v", err)
		return GetLogsOnPageResponse{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	startIndex := pageSize * request.Page
	buf := make([]string, 0)
	i := 0
	for scanner.Scan() {
		if i < startIndex {
			i++
			continue
		}
		buf = append(buf, scanner.Text())
		if i-startIndex >= pageSize {
			break
		}
	}

	return GetLogsOnPageResponse{
		Logs: collection_utils.MapperIt(buf, func(i int, r string) struct {
			Id     int
			Record string
		} {
			return struct {
				Id     int
				Record string
			}{i, r}
		}),
	}, nil
}

func GetLogsOnLastPage(config settings.ServiceSettings) (GetLogsOnPageResponse, error) {
	pageSize := config.App.Services.ServerLogging.PageSize
	file, err := os.Open(config.App.Services.ServerLogging.LogFile)
	if err != nil {
		fmt.Printf("Could not open log file: %v", err)
		return GetLogsOnPageResponse{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	list := list.New()
	for scanner.Scan() {
		// buf = append(buf, scanner.Text())
		list.PushBack(scanner.Text())
		if list.Len() > pageSize {
			list.Remove(list.Back())
		}
	}

	buf := make([]string, list.Len())
	for e, i := list.Front(), 0; e != nil; e = e.Next() {
		buf[i] = string(fmt.Sprintf("%v", e.Value))
		i++
	}

	return GetLogsOnPageResponse{
		Logs: collection_utils.MapperIt(buf, func(i int, r string) struct {
			Id     int
			Record string
		} {
			return struct {
				Id     int
				Record string
			}{i, r}
		}),
	}, nil
}
