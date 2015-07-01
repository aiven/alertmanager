// Copyright 2015 Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/prometheus/common/route"
	// "gopkg.in/yaml.v2"

	"github.com/prometheus/alertmanager/manager"
)

var (
	configFile = flag.String("config.file", "config.yml", "The configuration file")
)

func main() {
	conf, err := manager.LoadFile(*configFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conf)

	for _, r := range conf.Routes {
		fmt.Println(r)
	}

	state := manager.NewMemState()

	router := route.New()

	manager.NewAPI(router.WithPrefix("/api"), state)

	http.ListenAndServe(":9091", router)
}