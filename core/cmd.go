package core

import (
	"fmt"
	"io/ioutil"
	"net"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// prange is the range used to search for available ports
var prange = []int{1000, 65535}

// Cinit function uesd to handle initialize.
// Search for available porsts, generate network and
// .env file for docker
func Cinit(file string) string {
	sp, err := parseDCYML(file, `\$\{(\w+_PORT)\}`)
	if err != nil {
		return "[ERROR] ./docker-compose.yml is not exist, please spcify -cc flag"
	}

	var netname string
	ms := NewItems(nil)
	if len(sp) > 0 {
		ports, err := availablePorts(prange, len(sp))
		if err != nil {
			return "[ERROR] Available ports is not enough, please check system"
		}
		ms.AddTwoSlice(sp, ports)
		netname = initNetWork(ports)
	} else {
		netname = initNetWork(nil)
	}
	ms.AddTwoString("NETWORK_NAME", netname)

	c := ms.ToString("=")
	fileAbs, _ := filepath.Abs(filepath.Dir(file))
	envfile := filepath.Join(fileAbs, ".env")
	ioutil.WriteFile(envfile, []byte(c), 0644)
	return fmt.Sprintf("[SUCCESS] The path of .env file is  %s", envfile)
}

// initNetWork initialize network for docker,
// if ports is not nil, use the ports to generate the name of network, else
// does not generate
func initNetWork(ports []string) string {
	execCMD("docker network prune -f")

	var name string
	if ports != nil {
		name = "docker-" + strings.Join(ports, "_")
		execCMD(fmt.Sprintf("docker network create -d bridge %s", name))
	}

	return name
}

// availablePorts get available ports
func availablePorts(prange []int, count int) ([]string, error) {
	var ports []string
	num := count
	for i := prange[0]; i <= prange[1] && num > 0; i++ {
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", i))
		if err != nil {
			continue
		}

		ports = append(ports, fmt.Sprintf("%d", i))
		num--
		ln.Close()
	}

	var err error
	if len(ports) < count {
		err = fmt.Errorf("available ports not enough, only %d", len(ports))
	}

	return ports, err
}

// parseDCYML parse the docker-compose.yml file.
func parseDCYML(file string, pattern string) ([]string, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var rs []string
	r, _ := regexp.Compile(pattern)
	for _, sub := range r.FindAllStringSubmatch(string(b), -1) {
		rs = append(rs, sub[1])
	}

	return rs, nil
}

// execCMD execute the command, return result of execute and error
func execCMD(cmd string) (string, error) {
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
