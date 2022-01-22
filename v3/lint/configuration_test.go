/*
 * ZLint Copyright 2021 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package lint

import (
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"sync"
	"testing"

	"github.com/pelletier/go-toml"
)

func TestInt(t *testing.T) {
	type Test struct {
		A int
	}
	c, err := NewConfigFromString(`
[Test] 
A = 5`)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if test.A != 5 {
		t.Fatalf("wanted 5 got %d", test.A)
	}
}

func TestIntNegative(t *testing.T) {
	type Test struct {
		A int
	}
	c, err := NewConfigFromString(`
[Test] 
A = -5`)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if test.A != -5 {
		t.Fatalf("wanted -5 got %d", test.A)
	}
}

func TestUint(t *testing.T) {
	type Test struct {
		A uint
	}
	c, err := NewConfigFromString(`
[Test] 
A = 5`)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if test.A != 5 {
		t.Fatalf("wanted 5 got %d", test.A)
	}
}

func TestUintNegative(t *testing.T) {
	type Test struct {
		A uint
	}
	c, err := NewConfigFromString(`
[Test] 
A = -5`)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err == nil {
		t.Fatalf("expected an error when deserializing a negative number into a uint, got %v", test)
	}
}

func TestSmallInt(t *testing.T) {
	type Test struct {
		A uint8
	}
	c, err := NewConfigFromString(`
[Test] 
A = 300`)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err == nil {
		t.Fatalf("expected an error when deserializing a number too large to fit in a uint8, got %v", test)
	}
}

func TestByte(t *testing.T) {
	type Test struct {
		A byte
	}
	c, err := NewConfigFromString(`
[Test] 
A = 255`)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if test.A != 255 {
		t.Fatalf("wanted 255 got %d", test.A)
	}
}

func TestBool(t *testing.T) {
	type Test struct {
		A bool
	}
	c, err := NewConfigFromString(`
[Test] 
A = true`)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !test.A {
		t.Fatalf("wanted true got %v", test.A)
	}
}

func TestString(t *testing.T) {
	type Test struct {
		A string
	}
	c, err := NewConfigFromString(`
[Test] 
A = "the greatest song in the world"`)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if test.A != "the greatest song in the world" {
		t.Fatalf("wanted \"the greatest song in the world\" got %v", test.A)
	}
}

func TestArrayInt(t *testing.T) {
	type Test struct {
		A []int
	}
	c, err := NewConfigFromString(`
[Test] 
A = [1, 2, 3, 4]`)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test.A, []int{1, 2, 3, 4}) {
		t.Fatalf("wanted [1, 2, 3, 4] got %v", test.A)
	}
}

func TestArrayString(t *testing.T) {
	type Test struct {
		A []string
	}
	c, err := NewConfigFromString(`
[Test] 
A = ["1", "2", "3", "4"]`)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test.A, []string{"1", "2", "3", "4"}) {
		t.Fatalf("wanted [\"1\", \"2\", \"3\", \"4\"] got %v", test.A)
	}
}

func TestMapInt(t *testing.T) {
	type Test struct {
		A map[string]int
	}
	c, err := NewConfigFromString(`
[Test] 
A = { version = 42 }`)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test.A, map[string]int{"version": 42}) {
		t.Fatalf("wanted { \"version\": 42 } got %v", test.A)
	}
}

func TestMapString(t *testing.T) {
	type Test struct {
		A map[string]string
	}
	c, err := NewConfigFromString(`
[Test] 
A = { version = "1.2.3" }`)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test.A, map[string]string{"version": "1.2.3"}) {
		t.Fatalf("wanted { \"version\": \"1.2.3\" } got %v", test.A)
	}
}

func TestMapArray(t *testing.T) {
	type Test struct {
		A map[string][]int
	}
	c, err := NewConfigFromString(`
[Test] 
A = { version = [1, 2 ,3] }`)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test.A, map[string][]int{"version": {1, 2, 3}}) {
		t.Fatalf("wanted { \"version\": [1, 2 ,3] } got %v", test.A)
	}
}

func TestMapMap(t *testing.T) {
	type Test struct {
		A map[string]map[string]string
	}
	c, err := NewConfigFromString(`
[Test] 
A = { version = { commit = "29c848e565ebfa2a376767919bb0880be46b3c0f" } }`)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test.A, map[string]map[string]string{"version": {"commit": "29c848e565ebfa2a376767919bb0880be46b3c0f"}}) {
		t.Fatalf("wanted {\"versio\": { \"commit\": \"29c848e565ebfa2a376767919bb0880be46b3c0f\" } } got %v", test.A)
	}
}

func TestStruct(t *testing.T) {
	type Inner struct {
		B int
	}
	type Test struct {
		A Inner
	}
	c, err := NewConfigFromString(`
[Test] 
A = { B = 1 }`)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{Inner{1}}) {
		t.Fatalf("wanted {A {1}} got %v", test)
	}
}

func TestPointer(t *testing.T) {
	type Test struct {
		A *int
	}
	c, err := NewConfigFromString(`
[Test] 
A = 1`)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if test.A == nil {
		t.Fatal("wanted a pointer to 1, got nil")
	}
	if *test.A != 1 {
		t.Fatalf("wanted a pointer to 1, got a point to %d", *test.A)
	}
}

func TestInterface(t *testing.T) {
	type Test struct {
		A bool
		B io.Reader
	}
	c, err := NewConfigFromString(`
[Test] 
A = true`)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{true, nil}) {
		t.Fatalf("wanted {true nil} got %v", test)
	}
}

func TestSmokeExamplePrinting(t *testing.T) {
	type Inner struct {
		Things []int
	}
	type Test struct {
		A bool
		B io.Reader
		C *int
		D Inner
	}
	mapping := stripGlobalsFromExample(&Test{})
	rr, w := io.Pipe()
	var err error
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer w.Close()
		err = toml.NewEncoder(w).Indentation("").CompactComments(true).Encode(mapping)
	}()
	if err != nil {
		t.Fatal(err)
	}
	b, err := ioutil.ReadAll(rr)
	if err != nil {
		t.Fatal(err)
	}
	want := `A = false
C = 0

[D]
Things = []
`
	if want != string(b) {
		t.Fatalf("wanted `%s` got '%s'", want, string(b))
	}
}

func TestRecursiveStruct(t *testing.T) {
	type Test struct {
		A *Test
		B bool
	}
	c, err := NewConfigFromString(`
[Test]
A = { B = true }
B = true
`)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{&Test{nil, true}, true}) {
		t.Fatalf("wanted Test{&Test{nil, true}, true} got %v", test)
	}
}

func TestBadToml(t *testing.T) {
	_, err := NewConfigFromString(`(┛ಠ_ಠ)┛彡┻━┻`)
	if err == nil {
		t.Fatal("expected a parsing, however received a nil error")
	}
}

func TestPrivateMembers(t *testing.T) {
	type Test struct {
		private    string
		NotPrivate string
	}
	c, err := NewConfigFromString(`
[Test]
private = "this still should not show up"
NotPrivate = "just a string"
`)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if test.private != "" {
		t.Errorf("wanted '' got '%s'", test.private)
	}
	if test.NotPrivate != "just a string" {
		t.Errorf("wanted 'just a string' got '%s'", test.NotPrivate)
	}
}

func TestEmbedGlobal(t *testing.T) {
	type Test struct {
		Global        Global
		SomethingElse string
	}
	c, err := NewConfigFromString(`
    [Test]
    SomethingElse = "cool"
    `)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{Global: Global{}, SomethingElse: "cool"}) {
		t.Fatalf("wanted  Test{Global: Global{}, SomethingElse: \"cool\"}} got %v", test)
	}
}

func TestEmbedRFC5280Config(t *testing.T) {
	type Test struct {
		RFC5280Config RFC5280Config
		SomethingElse string
	}
	c, err := NewConfigFromString(`
    [Test]
    SomethingElse = "cool"
    `)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{RFC5280Config: RFC5280Config{}, SomethingElse: "cool"}) {
		t.Fatalf("wanted  Test{RFC5280Config: RFC5280Config{}, SomethingElse: \"cool\"}} got %v", test)
	}
}

func TestEmbedRFC5480Config(t *testing.T) {
	type Test struct {
		RFC5480Config RFC5480Config
		SomethingElse string
	}
	c, err := NewConfigFromString(`
    [Test]
    SomethingElse = "cool"
    `)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{RFC5480Config: RFC5480Config{}, SomethingElse: "cool"}) {
		t.Fatalf("wanted  Test{RFC5480Config: RFC5480Config{}, SomethingElse: \"cool\"}} got %v", test)
	}
}

func TestEmbedRFC5891Config(t *testing.T) {
	type Test struct {
		RFC5891Config RFC5891Config
		SomethingElse string
	}
	c, err := NewConfigFromString(`
    [Test]
    SomethingElse = "cool"
    `)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{RFC5891Config: RFC5891Config{}, SomethingElse: "cool"}) {
		t.Fatalf("wanted  Test{RFC5891Config: RFC5891Config{}, SomethingElse: \"cool\"}} got %v", test)
	}
}

func TestEmbedCABFBaselineRequirementsConfig(t *testing.T) {
	type Test struct {
		CABFBaselineRequirementsConfig CABFBaselineRequirementsConfig
		SomethingElse                  string
	}
	c, err := NewConfigFromString(`
    [Test]
    SomethingElse = "cool"
    `)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{CABFBaselineRequirementsConfig: CABFBaselineRequirementsConfig{}, SomethingElse: "cool"}) {
		t.Fatalf("wanted  Test{CABFBaselineRequirementsConfig: CABFBaselineRequirementsConfig{}, SomethingElse: \"cool\"}} got %v", test)
	}
}

func TestEmbedCABFEVGuidelinesConfig(t *testing.T) {
	type Test struct {
		CABFEVGuidelinesConfig CABFEVGuidelinesConfig
		SomethingElse          string
	}
	c, err := NewConfigFromString(`
    [Test]
    SomethingElse = "cool"
    `)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{CABFEVGuidelinesConfig: CABFEVGuidelinesConfig{}, SomethingElse: "cool"}) {
		t.Fatalf("wanted  Test{CABFEVGuidelinesConfig: CABFEVGuidelinesConfig{}, SomethingElse: \"cool\"}} got %v", test)
	}
}

func TestEmbedMozillaRootStorePolicyConfig(t *testing.T) {
	type Test struct {
		MozillaRootStorePolicyConfig MozillaRootStorePolicyConfig
		SomethingElse                string
	}
	c, err := NewConfigFromString(`
    [Test]
    SomethingElse = "cool"
    `)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{MozillaRootStorePolicyConfig: MozillaRootStorePolicyConfig{}, SomethingElse: "cool"}) {
		t.Fatalf("wanted  Test{MozillaRootStorePolicyConfig: MozillaRootStorePolicyConfig{}, SomethingElse: \"cool\"}} got %v", test)
	}
}

func TestEmbedAppleRootStorePolicyConfig(t *testing.T) {
	type Test struct {
		AppleRootStorePolicyConfig AppleRootStorePolicyConfig
		SomethingElse              string
	}
	c, err := NewConfigFromString(`
    [Test]
    SomethingElse = "cool"
    `)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{AppleRootStorePolicyConfig: AppleRootStorePolicyConfig{}, SomethingElse: "cool"}) {
		t.Fatalf("wanted  Test{AppleRootStorePolicyConfig: AppleRootStorePolicyConfig{}, SomethingElse: \"cool\"}} got %v", test)
	}
}

func TestEmbedCommunityConfig(t *testing.T) {
	type Test struct {
		CommunityConfig CommunityConfig
		SomethingElse   string
	}
	c, err := NewConfigFromString(`
    [Test]
    SomethingElse = "cool"
    `)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{CommunityConfig: CommunityConfig{}, SomethingElse: "cool"}) {
		t.Fatalf("wanted  Test{CommunityConfig: CommunityConfig{}, SomethingElse: \"cool\"}} got %v", test)
	}
}

func TestEmbedEtsiEsiConfig(t *testing.T) {
	type Test struct {
		EtsiEsiConfig EtsiEsiConfig
		SomethingElse string
	}
	c, err := NewConfigFromString(`
    [Test]
    SomethingElse = "cool"
    `)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{EtsiEsiConfig: EtsiEsiConfig{}, SomethingElse: "cool"}) {
		t.Fatalf("wanted  Test{EtsiEsiConfig: EtsiEsiConfig{}, SomethingElse: \"cool\"}} got %v", test)
	}
}

func TestEmbedPtrToGlobal(t *testing.T) {
	type Test struct {
		Global        *Global
		SomethingElse string
	}
	c, err := NewConfigFromString(`
    [Test]
    SomethingElse = "cool"
    `)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{Global: &Global{}, SomethingElse: "cool"}) {
		t.Fatalf("wanted  Test{Global: &Global{}, SomethingElse: \"cool\"}} got %v", test)
	}
}

func TestEmbedPtrToRFC5280Config(t *testing.T) {
	type Test struct {
		RFC5280Config *RFC5280Config
		SomethingElse string
	}
	c, err := NewConfigFromString(`
    [Test]
    SomethingElse = "cool"
    `)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{RFC5280Config: &RFC5280Config{}, SomethingElse: "cool"}) {
		t.Fatalf("wanted  Test{RFC5280Config: &RFC5280Config{}, SomethingElse: \"cool\"}} got %v", test)
	}
}

func TestEmbedPtrToRFC5480Config(t *testing.T) {
	type Test struct {
		RFC5480Config *RFC5480Config
		SomethingElse string
	}
	c, err := NewConfigFromString(`
    [Test]
    SomethingElse = "cool"
    `)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{RFC5480Config: &RFC5480Config{}, SomethingElse: "cool"}) {
		t.Fatalf("wanted  Test{RFC5480Config: &RFC5480Config{}, SomethingElse: \"cool\"}} got %v", test)
	}
}

func TestEmbedPtrToRFC5891Config(t *testing.T) {
	type Test struct {
		RFC5891Config *RFC5891Config
		SomethingElse string
	}
	c, err := NewConfigFromString(`
    [Test]
    SomethingElse = "cool"
    `)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{RFC5891Config: &RFC5891Config{}, SomethingElse: "cool"}) {
		t.Fatalf("wanted  Test{RFC5891Config: &RFC5891Config{}, SomethingElse: \"cool\"}} got %v", test)
	}
}

func TestEmbedPtrToCABFBaselineRequirementsConfig(t *testing.T) {
	type Test struct {
		CABFBaselineRequirementsConfig *CABFBaselineRequirementsConfig
		SomethingElse                  string
	}
	c, err := NewConfigFromString(`
    [Test]
    SomethingElse = "cool"
    `)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{CABFBaselineRequirementsConfig: &CABFBaselineRequirementsConfig{}, SomethingElse: "cool"}) {
		t.Fatalf("wanted  Test{CABFBaselineRequirementsConfig: &CABFBaselineRequirementsConfig{}, SomethingElse: \"cool\"}} got %v", test)
	}
}

func TestEmbedPtrToCABFEVGuidelinesConfig(t *testing.T) {
	type Test struct {
		CABFEVGuidelinesConfig *CABFEVGuidelinesConfig
		SomethingElse          string
	}
	c, err := NewConfigFromString(`
    [Test]
    SomethingElse = "cool"
    `)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{CABFEVGuidelinesConfig: &CABFEVGuidelinesConfig{}, SomethingElse: "cool"}) {
		t.Fatalf("wanted  Test{CABFEVGuidelinesConfig: &CABFEVGuidelinesConfig{}, SomethingElse: \"cool\"}} got %v", test)
	}
}

func TestEmbedPtrToMozillaRootStorePolicyConfig(t *testing.T) {
	type Test struct {
		MozillaRootStorePolicyConfig *MozillaRootStorePolicyConfig
		SomethingElse                string
	}
	c, err := NewConfigFromString(`
    [Test]
    SomethingElse = "cool"
    `)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{MozillaRootStorePolicyConfig: &MozillaRootStorePolicyConfig{}, SomethingElse: "cool"}) {
		t.Fatalf("wanted  Test{MozillaRootStorePolicyConfig: &MozillaRootStorePolicyConfig{}, SomethingElse: \"cool\"}} got %v", test)
	}
}

func TestEmbedPtrToAppleRootStorePolicyConfig(t *testing.T) {
	type Test struct {
		AppleRootStorePolicyConfig *AppleRootStorePolicyConfig
		SomethingElse              string
	}
	c, err := NewConfigFromString(`
    [Test]
    SomethingElse = "cool"
    `)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{AppleRootStorePolicyConfig: &AppleRootStorePolicyConfig{}, SomethingElse: "cool"}) {
		t.Fatalf("wanted  Test{AppleRootStorePolicyConfig: &AppleRootStorePolicyConfig{}, SomethingElse: \"cool\"}} got %v", test)
	}
}

func TestEmbedPtrToCommunityConfig(t *testing.T) {
	type Test struct {
		CommunityConfig *CommunityConfig
		SomethingElse   string
	}
	c, err := NewConfigFromString(`
    [Test]
    SomethingElse = "cool"
    `)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{CommunityConfig: &CommunityConfig{}, SomethingElse: "cool"}) {
		t.Fatalf("wanted  Test{CommunityConfig: &CommunityConfig{}, SomethingElse: \"cool\"}} got %v", test)
	}
}

func TestEmbedPtrToEtsiEsiConfig(t *testing.T) {
	type Test struct {
		EtsiEsiConfig *EtsiEsiConfig
		SomethingElse string
	}
	c, err := NewConfigFromString(`
    [Test]
    SomethingElse = "cool"
    `)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{EtsiEsiConfig: &EtsiEsiConfig{}, SomethingElse: "cool"}) {
		t.Fatalf("wanted  Test{EtsiEsiConfig: &EtsiEsiConfig{}, SomethingElse: \"cool\"}} got %v", test)
	}
}

func TestGlobalStripper(t *testing.T) {
	type Test struct {
		EtsiEsiConfig *EtsiEsiConfig
		SomethingElse string
	}
	c, err := NewConfigFromString(`
    [Test]
    SomethingElse = "cool"
    `)
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{EtsiEsiConfig: &EtsiEsiConfig{}, SomethingElse: "cool"}) {
		t.Fatalf("wanted  Test{EtsiEsiConfig: &EtsiEsiConfig{}, SomethingElse: \"cool\"}} got %v", test)
	}
}

func TestPrintConfiguration(t *testing.T) {
	gotBytes, err := NewRegistry().DefaultConfiguration()
	if err != nil {
		t.Fatal(err)
	}
	got := string(gotBytes)
	// I'm not a huge fan of this sort of test since it will have to be updated
	// on the slightest change, but it's better than not have a test for printing
	// out the configuration file.
	want := `
[AppleRootStorePolicyConfig]

[CABFBaselineRequirementsConfig]

[CABFEVGuidelinesConfig]

[CommunityConfig]

[MozillaRootStorePolicyConfig]

[RFC5280Config]

[RFC5480Config]

[RFC5891Config]
`
	if got != want {
		t.Fatalf("wanted '%s' but got '%s'", want, got)
	}
}

type TestGlobalConfigurable struct {
	A int
	B string
}

func (t *TestGlobalConfigurable) namespace() string {
	return "this_is_a_test"
}

func TestNewGlobal(t *testing.T) {
	type test struct {
		SomethingElse string `toml:"something_else"`
		T             *TestGlobalConfigurable
	}
	c, err := NewConfigFromString(`
[this_is_a_test]
A = 1
B = "the temples of syrinx"

[Test]
something_else = "fills our hallowed halls"
`)
	if err != nil {
		t.Fatal(err)
	}
	got := test{}
	err = c.Configure(&got, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if got.SomethingElse != "fills our hallowed halls" {
		t.Errorf("got '%s' want 'fills our hallowed halls", got.SomethingElse)
	}
	if got.T.A != 1 {
		t.Errorf("got %d want 1", got.T.A)
	}
	if got.T.B != "the temples of syrinx" {
		t.Errorf("got '%s' want 'the temples of syrinx", got.T.B)
	}
}

type TestGlobalConfigurableWithPrivates struct {
	A int
	B string
	c string
}

func (t *TestGlobalConfigurableWithPrivates) namespace() string {
	return "this_is_a_test"
}

func TestNewGlobalWithPrivateMembersDontGetPrinted(t *testing.T) {
	gotBytes, err := NewRegistry().defaultConfiguration([]GlobalConfiguration{&TestGlobalConfigurableWithPrivates{
		1, "2", "3",
	}})
	if err != nil {
		t.Fatal(err)
	}
	got := string(gotBytes)
	// I'm not a huge fan of this sort of test since it will have to be updated
	// on the slightest change, but it's better than not have a test for printing
	// out the configuration file.
	want := `
[this_is_a_test]
A = 1
B = "2"
`
	if got != want {
		t.Fatalf("wanted '%s' but got '%s'", want, got)
	}
}

func TestFailedGlobalDeser(t *testing.T) {
	type test struct {
		SomethingElse string `toml:"something_else"`
		T             *TestGlobalConfigurable
	}
	c, err := NewConfigFromString(`
[this_is_a_test]
A = "1" # It should be an int, not a string
B = "the temples of syrinx"

[Test]
something_else = "fills our hallowed halls"
`)
	if err != nil {
		t.Fatal(err)
	}
	got := test{}
	err = c.Configure(&got, "Test")
	if err == nil {
		t.Fatalf("expected error, but got %v", got)
	}
}

func TestFailedNestedGlobalDeser(t *testing.T) {
	type test struct {
		SomethingElse string `toml:"something_else"`
		Inner         struct {
			T *TestGlobalConfigurable
		}
	}
	c, err := NewConfigFromString(`
[this_is_a_test]
A = "1" # It should be an int, not a string
B = "the temples of syrinx"

[Test]
something_else = "fills our hallowed halls"
`)
	if err != nil {
		t.Fatal(err)
	}
	got := test{}
	err = c.Configure(&got, "Test")
	if err == nil {
		t.Fatalf("expected error, but got %v", got)
	}
}

func TestStripGlobalsFromStructWithPrivates(t *testing.T) {
	//nolint:staticheck
	type Test struct {
		A string
		B Global
		C int
		d int
	}
	test := Test{}
	got := stripGlobalsFromExample(&test).(map[string]interface{})
	want := map[string]interface{}{
		"A": "",
		"C": 0,
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wanted map[A: C:0], got %v", got)
	}
}

func TestNewEmptyConfig(t *testing.T) {
	c := NewEmptyConfig()
	got, err := c.tree.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	if got != nil {
		t.Fatalf("wanted nil byte slice, got %s", string(got))
	}
}

func TestConfigFromFile(t *testing.T) {
	type Test struct {
		A *Test
		B bool
	}
	f, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	_, err = f.WriteString(`
[Test]
A = { B = true }
B = true
`)
	if err != nil {
		f.Close()
		t.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		t.Fatal(err)
	}
	c, err := NewConfigFromFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	test := Test{}
	err = c.Configure(&test, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(test, Test{&Test{nil, true}, true}) {
		t.Fatalf("wanted Test{&Test{nil, true}, true} got %v", test)
	}
}

func TestBadConfigFromFile(t *testing.T) {
	f, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	_, err = f.WriteString(`
nope not gonna work
[Test]
A = { B = true }
B = true
`)
	if err != nil {
		f.Close()
		t.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		t.Fatal(err)
	}
	c, err := NewConfigFromFile(f.Name())
	if err == nil {
		t.Fatalf("expected error, got %v", c)
	}
}

func TestEmptyConfigFromEmptyPath(t *testing.T) {
	c, err := NewConfigFromFile("")
	if err != nil {
		t.Fatal(err)
	}
	got, err := c.tree.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	if got != nil {
		t.Fatalf("wanted nil byte slice, got %s", string(got))
	}
}

func TestFailedToOpenConfigFile(t *testing.T) {
	c, err := NewConfigFromFile("lol no not likely")
	if err == nil {
		t.Fatalf("expected an error got %v", c)
	}
}
