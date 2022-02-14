# DEVELOP ENVIRONMENT

- wsl2 Ubuntu 18.04
- go 1.16
- kubebuilder
- minikube

## Install go 1.16

```
## install Go
$ wget https://dl.google.com/go/go1.17.6.linux-amd64.tar.gz
$ sudo tar -xvf go1.17.6.linux-amd64.tar.gz
$ sudo mv go /usr/local

## setting
$ vi ~/.bashrc

...
export GOROOT=/usr/local/go
export GOPATHG=~/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

$ source ~/.bashrc
$ mkdir ~/go

## show go env list
$ go env
~~
```

## Install kubebuilder

refer https://book.kubebuilder.io/quick-start.html#installation
```
$ curl -L -o kubebuilder https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)
$ chmod +x kubebuilder && mv kubebuilder /usr/local/bin/
```

## Install docker

refer https://docs.docker.com/engine/install/ubuntu/
```
$ sudo apt-get update
$ sudo apt-get install \
    ca-certificates \
    curl \
    gnupg \
    lsb-release
$ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
$ echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
$ sudo apt-get update
$ sudo apt-get install docker-ce docker-ce-cli containerd.io
$ sudo service docker start
```

add docker privileges to user
```
sudo groupadd docker
sudo usermod -aG docker $USER
newgrp docker
```

## Install minikube
refer https://minikube.sigs.k8s.io/docs/start/
```
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
minikube start
```

## Install kubectl
```
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
sudo apt-get update
sudo apt-get install kubectl
```