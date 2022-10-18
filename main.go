package main

import (
	"log"
	"time"
	"github.com/intel/intel-device-plugins-for-kubernetes/pkg/deviceplugin"
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

const (
	name = "myname"
	devType = "mydevice"
	namespace = "mydeviceplugin.net"
	reconcilePeriod = time.Second * 5
)

type myDevicePlugin struct {
	name string
}

func newMyDevicePlugin(name string) *myDevicePlugin {
	log.Printf("New device plugin: %s\n", name)
	return &myDevicePlugin{
		name: name,
	}
}

func (dp *myDevicePlugin) Scan(notifier deviceplugin.Notifier) error {
	log.Printf("Scan device plugin: %s\n", dp.name)
	for {
		devTree := deviceplugin.NewDeviceTree()
		nodes := []pluginapi.DeviceSpec{}
		mounts := []pluginapi.Mount{}
		envs := map[string]string{}
		annotations := map[string]string{}
		devTree.AddDevice(devType, "0",
			deviceplugin.NewDeviceInfo(pluginapi.Healthy, nodes, mounts, envs, annotations))
		notifier.Notify(devTree)
		time.Sleep(reconcilePeriod)
	}
	return nil
}

func (dp *myDevicePlugin) PostAllocate(response *pluginapi.AllocateResponse) error {
	log.Printf("PostAllocate device plugin: %s\n", dp.name)
	if len(dp.name) > 0 {
		for _, containerResponse := range response.GetContainerResponses() {
			containerResponse.Annotations = map[string]string{
				"Name": dp.name,
			}
		}
	}
	return nil
}

func main() {
	log.Println("Starting...")
	dp := newMyDevicePlugin(name)
	manager := deviceplugin.NewManager(namespace, dp)
	manager.Run()
}
