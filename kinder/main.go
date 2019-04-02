/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This package is a stub main wrapping cmd/kinder.Main()
package main

import (
	kinder "k8s.io/kubeadm/kinder/cmd/kinder"

	// forces kinder actions to register
	_ "k8s.io/kubeadm/kinder/pkg/actions"
)

func main() {
	kinder.Main()
}