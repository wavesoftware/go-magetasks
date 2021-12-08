module github.com/wavesoftware/go-magetasks

go 1.16

require (
	github.com/blang/semver/v4 v4.0.0
	github.com/fatih/color v1.13.0
	github.com/google/go-containerregistry v0.7.0
	github.com/google/ko v0.9.4-0.20211208134726-54cddccd1cef
	github.com/hashicorp/go-multierror v1.1.1
	github.com/joho/godotenv v1.4.0
	github.com/magefile/mage v1.11.0
	github.com/wavesoftware/go-ensure v1.0.0
	golang.org/x/mod v0.5.1
	gotest.tools/v3 v3.0.3
)

// indirect dependencies
require (
	github.com/BurntSushi/toml v0.4.1 // indirect
	github.com/containerd/containerd v1.5.8 //indirect
	github.com/evanphx/json-patch/v5 v5.6.0 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	k8s.io/klog/v2 v2.30.0 // indirect
)
