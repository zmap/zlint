package lint

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/pelletier/go-toml"
)

// Configuration is a ZLint configuration which serves as a target
// to hold the full TOML tree that is a physical ZLint configuration./
type Configuration struct {
	tree *toml.Tree
}

// Configure attempts to deserialize the provided namespace into the provided empty interface
//
// For example, let's say that the name if your lint is MyLint, then the configuration
// file might looks something like the following.
//
// ```
//	[MyLint]
//	A = 1
//	B = 2
// ```
//
// Given this, our target struct may look like the following.
//
// ```
//	type MytLint struct {
//		A int
//		B uint
//	}
// ```
//
// So deserializing into this struct would look like,
//
// ```
// configuration.Configure(&myLint, myLint.Name())
// ```
func (c Configuration) Configure(lint interface{}, namespace string) error {
	return c.deserializeConfigInto(lint, namespace)
}

// NewConfig attempts to instantiate a configuration by consuming the contents of the provided reader.
//
// The contents of the provided reader MUST be in a valid TOML format. The caller of this function
// is responsible for closing the reader, if appropriate.
func NewConfig(r io.Reader) (Configuration, error) {
	tree, err := toml.LoadReader(r)
	if err != nil {
		return Configuration{}, err
	}
	return Configuration{tree}, nil
}

// NewConfigFromFile attempts to instantiate a configuration from the provided filesystem path.
//
// The file pointed to by `path` MUST be valid TOML file.
func NewConfigFromFile(path string) (Configuration, error) {
	if path == "" {
		return NewEmptyConfig(), nil
	}
	f, err := os.Open(path)
	if err != nil {
		return Configuration{}, fmt.Errorf("failed to open the provided configuration at %s. Error: %s", path, err.Error())
	}
	defer f.Close()
	return NewConfig(f)
}

// NewConfigFromString attempts to instantiate a configuration from the provided string.
//
// The provided string MUST be in a valid TOML format.
func NewConfigFromString(config string) (Configuration, error) {
	return NewConfig(strings.NewReader(config))
}

// NewEmptyConfig returns a configuration that is backed by an entirely empty TOML tree.
//
// This is useful of no particular configuration is set at all by the user of ZLint as
// any attempt to resolve a namespace in `deserializeConfigInto` fails and thus results
// in all defaults for all lints being maintained.
func NewEmptyConfig() Configuration {
	ctx, _ := NewConfigFromString("")
	return ctx
}

// deserializeConfigInto deserializes the section labeled by the provided `namespace`
// into the provided target `interface{}`.
//
// For example, given the following configuration...
//
// ```
// [e_some_lint]
// field = 1
// flag = false
//
// [w_some_other_lint]
// is_web_pki = true
// ```
//
// And the following struct definition...
//
// ```
// type SomeOtherLint {
//		IsWebPKI bool `toml:"is_web_pki"`
// }
// ```
//
// Then the invocation of this function should be
//
// ```
// lint := &SomeOtherLink{}
// deserializeConfigInto(link, "w_some_other_lint")
// ```
//
// If there is no such namespace found in this configuration then provided the namespace specific data encoded
// within `target` is left unmodified. However, configuration of higher scoped fields will still be attempted.
func (c Configuration) deserializeConfigInto(target interface{}, namespace string) error {
	if tree := c.tree.Get(namespace); tree != nil {
		err := tree.(*toml.Tree).Unmarshal(target)
		if err != nil {
			return err
		}
	}
	return c.resolveHigherScopedReferences(target)
}

// resolveHigherScopeReferences takes in an interface{} value and attempts to
// find any field within its inner value that is either a struct or a pointer
// to a struct that is one of our global configurable types. If such a field
// exists then that higher scoped configuration will be deserialized, in-place,
// into the value held by the provided interface{}.
//
// This procedure is recursive.
func (c Configuration) resolveHigherScopedReferences(i interface{}) error {
	value := reflect.Indirect(reflect.ValueOf(i))
	if value.Kind() != reflect.Struct {
		// Our target higher scoped configurations are either structs
		// or are fields of structs. Any other Kind is simply cannot
		// be a target for deserialization here. For example, an interface
		// does not make sense since an interface cannot have fields nor
		// are any of our higher scoped configurations interfaces themselves.
		//
		// For a comprehensive list of Kinds you may review type.go in the
		// `reflect` package.
		return nil
	}
	// Iterate through every field within the struct held by the provided interface{}.
	// If the field is either one of our higher scoped configurations (or a pointer to one)
	// then deserialize that higher scoped configuration into that field. If the field
	// is not one of our higher scoped configurations then recursively pass it to this function
	// in an attempt to resolve it.
	for field := 0; field < value.NumField(); field++ {
		field := value.Field(field)
		if !field.CanInterface() {
			continue
		}
		if config, ok := field.Interface().(GlobalConfiguration); ok {
			if field.Kind() == reflect.Ptr && field.IsZero() {
				config = reflect.New(field.Type().Elem()).Interface().(GlobalConfiguration)
			}
			err := c.deserializeConfigInto(config, config.namespace())
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(config))
		} else {
			// In order to deserialize into a field it does indeed need to be addressable.
			if !field.CanAddr() {
				continue
			}
			err := c.resolveHigherScopedReferences(field.Addr().Interface())
			if err != nil {
				return err
			}
			continue
		}
	}
	return nil
}

// stripGlobalsFromExample takes in an interface{} and returns a mapping that is
// the provided struct but with all references to higher scoped configurations scrubbed.
//
// This is intended only for use when constructing an example configuration file via the
// `-exampleConfig` flag. This is to avoid visually redundant, and possibly incorrect,
// examples such as the following...
//
// ```
// [Global]
// something = false
// something_else = ""
//
// [e_some_lint]
// my_data = 0
// my_flag = false
// globals = { something = false, something_else = "" }
// ``` o
func stripGlobalsFromExample(i interface{}) interface{} {
	value := reflect.Indirect(reflect.ValueOf(i))
	if value.Kind() != reflect.Struct {
		return i
	}
	m := map[string]interface{}{}
	for field := 0; field < value.NumField(); field++ {
		name := value.Type().Field(field).Name
		field := value.Field(field)
		if !field.CanInterface() {
			continue
		}
		if _, ok := field.Interface().(GlobalConfiguration); ok {
			continue
		}
		if field.Kind() == reflect.Ptr && field.IsZero() {
			field = reflect.New(field.Type().Elem())
		}
		m[name] = stripGlobalsFromExample(field.Interface())
	}
	return m
}
