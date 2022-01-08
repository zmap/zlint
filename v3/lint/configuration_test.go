package lint

import (
	"io"
	"io/ioutil"
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
