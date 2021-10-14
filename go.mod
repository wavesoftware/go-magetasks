module github.com/wavesoftware/go-magetasks

go 1.16

require (
	github.com/BurntSushi/toml v0.4.1 // indirect
	github.com/docker/cli v20.10.9+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.6.4 // indirect
	github.com/fatih/color v1.13.0
	github.com/go-logr/logr v1.1.0 // indirect
	github.com/google/go-containerregistry v0.6.0
	github.com/google/ko v0.9.3
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1
	github.com/joho/godotenv v1.4.0
	github.com/magefile/mage v1.11.0
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/wavesoftware/go-ensure v1.0.0
	golang.org/x/mod v0.5.1
	golang.org/x/net v0.0.0-20211013171255-e13a2654a71e // indirect
	golang.org/x/sys v0.0.0-20211013075003-97ac67df715c // indirect
	google.golang.org/genproto v0.0.0-20211013025323-ce878158c4d4 // indirect
	gotest.tools/v3 v3.0.3
	k8s.io/klog/v2 v2.20.0 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

// FIXME: remove after https://github.com/google/ko/issues/476
replace github.com/google/ko v0.9.3 => github.com/cardil/ko v0.9.4-0.20211013122324-2e666a856ec8
