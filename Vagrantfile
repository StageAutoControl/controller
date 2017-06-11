Vagrant.configure("2") do |config|
  config.vm.define "artnet" do |nodeconfig|
    nodeconfig.ssh.forward_agent = true
    nodeconfig.ssh.insert_key = false
    nodeconfig.vm.box = "bento/centos-7.3"
    nodeconfig.vm.hostname = "artnet"
    nodeconfig.vm.network "private_network", ip: "192.168.50.4", netmask: "24"

    nodeconfig.vm.provider :virtualbox do |v|
      v.name = "artnet"
      v.memory = 512
      v.cpus = 1
      v.gui = false
    end
  end
end
