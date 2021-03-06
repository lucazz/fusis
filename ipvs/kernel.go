// Copyright 2012 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: jsing@google.com (Joel Sing)
package ipvs

import (
	"sync"

	log "github.com/golang/glog"
	gipvs "github.com/google/seesaw/ipvs"
)

var mt sync.Mutex

// initIPVS initialises the IPVS sub-component.
func initKernel() error {
	mt.Lock()
	defer mt.Unlock()
	log.Infof("Initialising gipvs...")
	if err := gipvs.Init(); err != nil {
		// TODO(jsing): modprobe ip_vs and try again.
		log.Fatalf("IPVS initialisation failed: %v", err)
		return err
	}
	return nil
}

// IPVSFlush flushes all services and destinations from the IPVS table.
func (k *IpvsKernel) Flush() error {
	mt.Lock()
	defer mt.Unlock()
	return gipvs.Flush()
}

// IPVSGetServices gets the currently configured services from the IPVS table.
func (k *IpvsKernel) GetServices() ([]*gipvs.Service, error) {
	mt.Lock()
	defer mt.Unlock()
	return gipvs.GetServices()
}

// IPVSGetService gets the currently configured service from the IPVS table,
// which matches the specified service.
// func (k *IpvsKernel) IPVSGetService(svc gipvs.Service) (error) {
// 	mt.Lock()
// 	defer mt.Unlock()
// 	so, err := gipvs.GetService(svc)
// 	if err != nil {
// 		return err
// 	}
// 	s.Services = []*gipvs.Service{so}
// 	return nil
// }
//
// // IPVSAddService adds the specified service to the IPVS table.
func (k *IpvsKernel) AddService(svc *gipvs.Service) error {
	mt.Lock()
	defer mt.Unlock()
	return gipvs.AddService(*svc)
}

//
// // IPVSUpdateService updates the specified service in the IPVS table.
// func (k *IpvsKernel) IPVSUpdateService(svc *gipvs.Service, out *int) error {
// 	mt.Lock()
// 	defer mt.Unlock()
// 	return gipvs.UpdateService(*svc)
// }
//
// IPVSDeleteService deletes the specified service from the IPVS table.
func (k *IpvsKernel) DeleteService(svc *gipvs.Service) error {
	mt.Lock()
	defer mt.Unlock()
	return gipvs.DeleteService(*svc)
}

//
// IPVSAddDestination adds the specified destination to the IPVS table.
func (k *IpvsKernel) AddDestination(svc gipvs.Service, dst gipvs.Destination) error {
	mt.Lock()
	defer mt.Unlock()
	return gipvs.AddDestination(svc, dst)
}

//
// // IPVSUpdateDestination updates the specified destination in the IPVS table.
// func (k *IpvsKernel) IPVSUpdateDestination(dst *ncctypes.IPVSDestination, out *int) error {
// 	mt.Lock()
// 	defer mt.Unlock()
// 	return gipvs.UpdateDestination(*dst.Service, *dst.Destination)
// }
//
// // IPVSDeleteDestination deletes the specified destination from the IPVS table.
func (k *IpvsKernel) DeleteDestination(svc gipvs.Service, dst gipvs.Destination) error {
	mt.Lock()
	defer mt.Unlock()
	return gipvs.DeleteDestination(svc, dst)
}
