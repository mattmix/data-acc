package fake

import (
	"bytes"
	"fmt"
	"github.com/RSE-Cambridge/data-acc/internal/pkg/registry"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type HostInfo struct {
	MDTS []string       `yaml:"mdts,omitempty,flow"`
	OSTS map[string]int `yaml:"osts,omitempty,flow"`
}

type FSInfo struct {
	Hosts map[string]HostInfo `yaml:"hosts"`
	Vars  map[string]string   `yaml:"vars,flow"`
}

type FileSystems struct {
	Children map[string]FSInfo `yaml:"children"`
}

type Wrapper struct {
	All FileSystems
}

func printLustreInfo(volume registry.Volume, brickAllocations []registry.BrickAllocation) string {
	var mdt registry.BrickAllocation
	osts := make(map[string][]registry.BrickAllocation)
	for _, allocation := range brickAllocations {
		if allocation.AllocatedIndex == 0 {
			mdt = allocation
			current, success := osts[allocation.Hostname]
			if !success {
				// ensure hostname will be iterated through below to output mdt if required
				osts[allocation.Hostname] = current
			}
		} else {
			osts[allocation.Hostname] = append(osts[allocation.Hostname], allocation)
		}
	}

	hosts := make(map[string]HostInfo)
	for host, allocations := range osts {
		osts := make(map[string]int)
		for _, allocation := range allocations {
			osts[allocation.Device] = int(allocation.AllocatedIndex)
		}
		hostInfo := HostInfo{OSTS: osts}
		if mdt.Hostname == host {
			hostInfo.MDTS = []string{mdt.Device}
		}
		hosts[host] = hostInfo
	}
	fsinfo := FSInfo{
		Vars:  map[string]string{"mgsnode": mdt.Hostname},
		Hosts: hosts,
	}
	fsname := fmt.Sprintf("fs%s", volume.Name)
	data := Wrapper{All: FileSystems{Children: map[string]FSInfo{fsname: fsinfo}}}

	output, err := yaml.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}
	return string(output)
}

func printLustrePlaybook(volume registry.Volume) string {
	return fmt.Sprintf(`---
- name: Install Lustre
  hosts: fs%s
  become: yes
  gather_facts: no
  roles:
    - role: lustre`, volume.Name)
}

func executeTempAnsible(volume registry.Volume, brickAllocations []registry.BrickAllocation, teardown bool) error {
	dir, err := ioutil.TempDir("", fmt.Sprintf("fs%s_", volume.Name))
	if err != nil {
		return err
	}
	log.Println("Using ansible tempdir:", dir)
	if !teardown {
		// TODO... do this properly!
		defer os.RemoveAll(dir)
	}

	playbook := printLustrePlaybook(volume)
	tmpPlaybook := filepath.Join(dir, "dac.yml")
	if err := ioutil.WriteFile(tmpPlaybook, bytes.NewBufferString(playbook).Bytes(), 0666); err != nil {
		return err
	}
	log.Println(playbook)

	inventory := printLustreInfo(volume, brickAllocations)
	tmpInventory := filepath.Join(dir, "inventory")
	if err := ioutil.WriteFile(tmpInventory, bytes.NewBufferString(inventory).Bytes(), 0666); err != nil {
		return err
	}
	log.Println(inventory)

	cmd := exec.Command("cp", "-r",
		"/home/centos/go/src/github.com/JohnGarbutt/data-acc/fs-ansible/environment/roles", dir)
	output, err := cmd.CombinedOutput()
	log.Println(string(output))
	if err != nil {
		return err
	}
	cmd = exec.Command("cp", "-r",
		"/home/centos/go/src/github.com/JohnGarbutt/data-acc/fs-ansible/environment/.venv", dir)
	output, err = cmd.CombinedOutput()
	log.Println(string(output))
	if err != nil {
		return err
	}

	if !teardown {
		cmdStr := fmt.Sprintf(`cd %s; . .venv/bin/activate; ansible-playbook dac.yml -i inventory --tag format_mdtmgs --tag format_osts; ansible-playbook dac.yml -i inventory --tag start_osts --tag start_mgsdt --tag mount_fs`, dir)
		log.Println(cmdStr)
		cmd = exec.Command("bash", "-c", cmdStr)
		output, err := cmd.CombinedOutput()
		log.Println(string(output))
		return err
	} else {
		cmdStr := fmt.Sprintf(`cd %s; . .venv/bin/activate; ansible-playbook dac.yml -i inventory --tag stop_osts --tag stop_mgsdt --tag umount_fs; ansible-playbook dac.yml -i inventory --tag format_mdtmgs --tag format_osts`, dir)
		log.Println(cmdStr)
		cmd = exec.Command("bash", "-c", cmdStr)
		output, err := cmd.CombinedOutput()
		log.Println(string(output))
		return err
	}
}