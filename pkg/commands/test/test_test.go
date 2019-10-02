package test

import (
	"context"
	"testing"

	"github.com/instrumenta/conftest/pkg/parser/yaml"
	"github.com/spf13/viper"
)

func TestWarnQuery(t *testing.T) {

	tests := []struct {
		in  string
		exp bool
	}{
		{"", false},
		{"warn", true},
		{"warnXYZ", false},
		{"warn_", false},
		{"warn_x", true},
		{"warn_x_y_z", true},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			res := WarnQ.MatchString(tt.in)

			if tt.exp != res {
				t.Fatalf("%s recognized as `warn` query - expected: %v actual: %v", tt.in, tt.exp, res)
			}
		})
	}
}

func TestCombineConfig(t *testing.T) {
	viper.Set("namespace", "main")
	testTable := []struct {
		name              string
		combineConfigFlag bool
		policyPath        string
		fileList          []string
	}{
		{
			name:              "valid policy with combine-config=true should namespace the configs into a map (single file)",
			combineConfigFlag: true,
			policyPath:        "testdata/policy/test_policy_multifile.rego",
			fileList:          []string{"testdata/deployment.yaml"},
		},
		{
			name:              "config combine-config=false no namespacing, individual evaluation (single file)",
			combineConfigFlag: false,
			policyPath:        "testdata/policy/test_policy.rego",
			fileList:          []string{"testdata/deployment.yaml"},
		},
		{
			name:              "config combine-config=false no namespacing, individual evaluation (multi-file)",
			combineConfigFlag: false,
			policyPath:        "testdata/policy/test_policy.rego",
			fileList:          []string{"testdata/deployment+service.yaml", "testdata/deployment.yaml"},
		},
		{
			name:              "valid policy with combine-config=true should namespace the configs into a map (multi-file)",
			combineConfigFlag: true,
			policyPath:        "testdata/policy/test_policy_multifile.rego",
			fileList:          []string{"testdata/deployment+service.yaml", "testdata/deployment.yaml"},
		},
	}

	for _, testunit := range testTable {
		t.Run(testunit.name, func(t *testing.T) {
			viper.Set(CombineConfigFlagName, testunit.combineConfigFlag)
			viper.Set("policy", testunit.policyPath)
			errorExitCodeFromCall := 0
			var outputPrinter *FakeOutputManager
			cmd := NewTestCommand(func(int) {
				errorExitCodeFromCall += 1
			}, func() OutputManager {
				outputPrinter = new(FakeOutputManager)
				return outputPrinter
			})
			cmd.Run(cmd, testunit.fileList)
			if outputPrinter.PutCallCount() != len(testunit.fileList) && !testunit.combineConfigFlag {
				t.Errorf(
					"Output manager when combine-config is false should print output for each file: expected %v calls but got %v",
					len(testunit.fileList),
					outputPrinter.PutCallCount(),
				)
			}
			if errorExitCodeFromCall == 0 && testunit.combineConfigFlag {
				t.Errorf(
					"Output manager when combine-config is true should have failed but it exited with a zero code: %v",
					errorExitCodeFromCall,
				)
			}
		})
	}

	t.Run("combine-config flag exists", func(t *testing.T) {
		callCount := 0
		cmd := NewTestCommand(func(int) {
			callCount += 1
		}, func() OutputManager {
			return new(FakeOutputManager)
		})
		if cmd.Flag("combine-config") == nil {
			t.Errorf("combine-config flag should exist")
		}
	})
}

func TestInputFlag(t *testing.T) {
	testTable := []struct {
		name       string
		fileList   []string
		input      string
		shouldFail bool
	}{
		{
			name:       "when flag exists it should use the flag value",
			input:      "tf",
			fileList:   []string{"testdata/deployment.yaml"},
			shouldFail: true,
		},
		{
			name:       "when flag doesnt exist it should use the file extension",
			input:      "",
			fileList:   []string{"testdata/deployment.yaml"},
			shouldFail: false,
		},
	}

	for _, testUnit := range testTable {
		t.Run(testUnit.name, func(t *testing.T) {
			viper.Set("policy", "testdata/policy/test_policy.rego")
			viper.Set("input", testUnit.input)
			exitCallCount := 0
			cmd := NewTestCommand(func(int) {
				exitCallCount += 1
			}, func() OutputManager {
				return new(FakeOutputManager)
			})
			cmd.Run(cmd, testUnit.fileList)

			if testUnit.shouldFail && exitCallCount == 0 {
				t.Error("we expected to fail but did not")
			}

			if testUnit.shouldFail == false && exitCallCount >= 1 {
				t.Error("we did not expect to fail here, yet we did")
			}
		})
	}
}
func TestFailQuery(t *testing.T) {

	tests := []struct {
		in  string
		exp bool
	}{
		{"", false},
		{"deny", true},
		{"violation", true},
		{"denyXYZ", false},
		{"violationXYZ", false},
		{"deny_", false},
		{"violation_", false},
		{"deny_x", true},
		{"violation_x", true},
		{"deny_x_y_z", true},
		{"violation_x_y_z", true},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			res := DenyQ.MatchString(tt.in)

			if tt.exp != res {
				t.Fatalf("%s recognized as `fail` query - expected: %v actual: %v", tt.in, tt.exp, res)
			}
		})
	}
}

func TestMultifileYaml(t *testing.T) {
	ctx := context.Background()

	config := `apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello-kubernetes
---
apiVersion: v1
kind: Service
metadata:
  name: hello-kubernetes`

	yaml := yaml.Parser{}

	var jsonConfig interface{}
	err := yaml.Unmarshal([]byte(config), &jsonConfig)
	if err != nil {
		t.Fatalf("Could not unmarshal yaml")
	}

	compiler, err := buildCompiler("testdata/policy/test_policy_multifile.rego")
	if err != nil {
		t.Fatalf("Could not build rego compiler")
	}

	results, err := processData(ctx, jsonConfig, compiler)
	if err != nil {
		t.Fatalf("Could not process policy file")
	}

	const expected = 2
	actual := len(results.Failures)
	if actual != expected {
		t.Logf("Multifile yaml test failure. Got %v failures, expected %v", actual, expected)
	}
}
