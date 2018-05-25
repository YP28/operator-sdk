// Copyright 2018 The Operator-SDK Authors
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

package generator

import (
	"io"

	"gopkg.in/yaml.v2"
)

type Config struct {
	// APIVersion is the kubernetes apiVersion that has the format of $GROUP_NAME/$VERSION.
	APIVersion string `yaml:"apiVersion"`
	// Kind is the kubernetes resource kind.
	Kind string `yaml:"kind"`
	// ProjectName is name of the new operator application
	// and is also the name of the base directory.
	ProjectName string `yaml:"projectName"`
	// TmplCfg contains the configuration for the templates to use
	// to generate the Kubernetes configuration files.
	TmplCfg TemplateConfig `yaml:"templateConfig"`
}

type TemplateConfig struct {
	// CRDYamlTmpl is the path to the yaml file that contains
	// the template for the CustomResourceDefinition
	CRDYamlTmpl string `yaml:"crdYamlTmpl"`
	// CRYamlTmpl is the path to the yaml file that contains
	// the template for the CustomResource
	CRYamlTmpl string `yaml:"crYamlTmpl"`
	// OperatorYamlTmpl is the path to the yaml file that contains
	// the template for the Operator Deployment
	OperatorYamlTmpl string `yaml:"operatorYamlTmpl"`
	// RBACYamlTmpl is the path to the yaml file that contains
	// the template for the RBAC yaml file
	RBACYamlTmpl string `yaml:"rbacYamlTmpl"`
}

func renderConfigFile(w io.Writer, apiVersion, kind, projectName string) error {
	o, err := yaml.Marshal(&Config{
		APIVersion:  apiVersion,
		Kind:        kind,
		ProjectName: projectName,
	})
	if err != nil {
		return err
	}

	_, err = w.Write(o)
	return err
}
