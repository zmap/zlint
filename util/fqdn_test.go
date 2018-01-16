/*
 * ZLint Copyright 2018 Regents of the University of Michigan
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

package util

import (
	"testing"
)

func TestIsFQDNCorrectFQDN(t *testing.T) {
	domain := "google.com"
	expected := true
	actual := IsFQDN(domain)
	if expected != actual {
		t.Error(
			"For", domain,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestIsFQDNQuestionMarkFQDN(t *testing.T) {
	domain := "?.?.abc.com"
	expected := true
	actual := IsFQDN(domain)
	if expected != actual {
		t.Error(
			"For", domain,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestIsFQDNQuestionMarkIncorrectPlaceFQDN(t *testing.T) {
	domain := "?.?.abc?.com"
	expected := false
	actual := IsFQDN(domain)
	if expected != actual {
		t.Error(
			"For", domain,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestIsFQDNManyQuestionMarksFQDN(t *testing.T) {
	domain := "?.?.?.?.?.?.?.abc.com"
	expected := true
	actual := IsFQDN(domain)
	if expected != actual {
		t.Error(
			"For", domain,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestIsFQDNWildcardFQDN(t *testing.T) {
	domain := "*.abc.com"
	expected := true
	actual := IsFQDN(domain)
	if expected != actual {
		t.Error(
			"For", domain,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestIsFQDNNotFQDN(t *testing.T) {
	domain := "abc"
	expected := false
	actual := IsFQDN(domain)
	if expected != actual {
		t.Error(
			"For", domain,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityBadURI(t *testing.T) {
	uri := "not//a/valid/uri"
	expected := ""
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostBadURI(t *testing.T) {
	uri := "not//a/valid/uri"
	expected := ""
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityRootless(t *testing.T) {
	uri := "sip:user@host.com"
	expected := ""
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostRootless(t *testing.T) {
	uri := "sip:user@host.com"
	expected := ""
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityNoUserinfoNoPortNoAbsolutePathNoQueryNoFragment(t *testing.T) {
	uri := "scheme://host.com"
	expected := "host.com"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostNoUserinfoNoPortNoAbsolutePathNoQueryNoFragment(t *testing.T) {
	uri := "scheme://host.com"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityWithUserinfoNoPortNoAbsolutePathNoQueryNoFragment(t *testing.T) {
	uri := "scheme://user@host.com"
	expected := "user@host.com"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostWithUserinfoNoPortNoAbsolutePathNoQueryNoFragment(t *testing.T) {
	uri := "scheme://user@host.com"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityNoUserinfoWithPortNoAbsolutePathNoQueryNoFragment(t *testing.T) {
	uri := "scheme://host.com:123"
	expected := "host.com:123"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostNoUserinfoWithPortNoAbsolutePathNoQueryNoFragment(t *testing.T) {
	uri := "scheme://host.com:123"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityWithUserinfoWithPortNoAbsolutePathNoQueryNoFragment(t *testing.T) {
	uri := "scheme://user@host.com:123"
	expected := "user@host.com:123"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostWithUserinfoWithPortNoAbsolutePathNoQueryNoFragment(t *testing.T) {
	uri := "scheme://user@host.com:123"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityNoUserinfoNoPortWithAbsolutePathNoQueryNoFragment(t *testing.T) {
	uri := "scheme://host.com/path/to/something"
	expected := "host.com"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostNoUserinfoNoPortWithAbsolutePathNoQueryNoFragment(t *testing.T) {
	uri := "scheme://host.com/path/to/something"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityWithUserinfoNoPortWithAbsolutePathNoQueryNoFragment(t *testing.T) {
	uri := "scheme://user@host.com/path/to/something"
	expected := "user@host.com"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostWithUserinfoNoPortWithAbsolutePathNoQueryNoFragment(t *testing.T) {
	uri := "scheme://user@host.com/path/to/something"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityNoUserinfoWithPortWithAbsolutePathNoQueryNoFragment(t *testing.T) {
	uri := "scheme://host.com:123/path/to/something"
	expected := "host.com:123"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostNoUserinfoWithPortWithAbsolutePathNoQueryNoFragment(t *testing.T) {
	uri := "scheme://host.com:123/path/to/something"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityWithUserinfoWithPortWithAbsolutePathNoQueryNoFragment(t *testing.T) {
	uri := "scheme://user@host.com:123/path/to/something"
	expected := "user@host.com:123"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostWithUserinfoWithPortWithAbsolutePathNoQueryNoFragment(t *testing.T) {
	uri := "scheme://user@host.com:123/path/to/something"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityNoUserinfoNoPortNoAbsolutePathWithQueryNoFragment(t *testing.T) {
	uri := "scheme://host.com?query=something"
	expected := "host.com"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostNoUserinfoNoPortNoAbsolutePathWithQueryNoFragment(t *testing.T) {
	uri := "scheme://host.com?query=something"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityWithUserinfoNoPortNoAbsolutePathWithQueryNoFragment(t *testing.T) {
	uri := "scheme://user@host.com?query=something"
	expected := "user@host.com"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostWithUserinfoNoPortNoAbsolutePathWithQueryNoFragment(t *testing.T) {
	uri := "scheme://user@host.com?query=something"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityNoUserinfoWithPortNoAbsolutePathWithQueryNoFragment(t *testing.T) {
	uri := "scheme://host.com:123?query=something"
	expected := "host.com:123"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostNoUserinfoWithPortNoAbsolutePathWithQueryNoFragment(t *testing.T) {
	uri := "scheme://host.com:123?query=something"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityWithUserinfoWithPortNoAbsolutePathWithQueryNoFragment(t *testing.T) {
	uri := "scheme://user@host.com:123?query=something"
	expected := "user@host.com:123"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostWithUserinfoWithPortNoAbsolutePathWithQueryNoFragment(t *testing.T) {
	uri := "scheme://user@host.com:123?query=something"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityNoUserinfoNoPortWithAbsolutePathWithQueryNoFragment(t *testing.T) {
	uri := "scheme://host.com/path/to/something?query=something"
	expected := "host.com"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostNoUserinfoNoPortWithAbsolutePathWithQueryNoFragment(t *testing.T) {
	uri := "scheme://host.com/path/to/something?query=something"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityWithUserinfoNoPortWithAbsolutePathWithQueryNoFragment(t *testing.T) {
	uri := "scheme://user@host.com/path/to/something?query=something"
	expected := "user@host.com"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostWithUserinfoNoPortWithAbsolutePathWithQueryNoFragment(t *testing.T) {
	uri := "scheme://user@host.com/path/to/something?query=something"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityNoUserinfoWithPortWithAbsolutePathWithQueryNoFragment(t *testing.T) {
	uri := "scheme://host.com:123/path/to/something?query=something"
	expected := "host.com:123"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostNoUserinfoWithPortWithAbsolutePathWithQueryNoFragment(t *testing.T) {
	uri := "scheme://host.com:123/path/to/something?query=something"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityWithUserinfoWithPortWithAbsolutePathWithQueryNoFragment(t *testing.T) {
	uri := "scheme://user@host.com:123/path/to/something?query=something"
	expected := "user@host.com:123"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostWithUserinfoWithPortWithAbsolutePathWithQueryNoFragment(t *testing.T) {
	uri := "scheme://user@host.com:123/path/to/something?query=something"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityNoUserinfoNoPortNoAbsolutePathNoQueryWithFragment(t *testing.T) {
	uri := "scheme://host.com#fragment"
	expected := "host.com"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostNoUserinfoNoPortNoAbsolutePathNoQueryWithFragment(t *testing.T) {
	uri := "scheme://host.com#fragment"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityWithUserinfoNoPortNoAbsolutePathNoQueryWithFragment(t *testing.T) {
	uri := "scheme://user@host.com#fragment"
	expected := "user@host.com"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostWithUserinfoNoPortNoAbsolutePathNoQueryWithFragment(t *testing.T) {
	uri := "scheme://user@host.com#fragment"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityNoUserinfoWithPortNoAbsolutePathNoQueryWithFragment(t *testing.T) {
	uri := "scheme://host.com:123#fragment"
	expected := "host.com:123"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostNoUserinfoWithPortNoAbsolutePathNoQueryWithFragment(t *testing.T) {
	uri := "scheme://host.com:123#fragment"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityWithUserinfoWithPortNoAbsolutePathNoQueryWithFragment(t *testing.T) {
	uri := "scheme://user@host.com:123#fragment"
	expected := "user@host.com:123"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostWithUserinfoWithPortNoAbsolutePathNoQueryWithFragment(t *testing.T) {
	uri := "scheme://user@host.com:123#fragment"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityNoUserinfoNoPortWithAbsolutePathNoQueryWithFragment(t *testing.T) {
	uri := "scheme://host.com/path/to/something#fragment"
	expected := "host.com"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostNoUserinfoNoPortWithAbsolutePathNoQueryWithFragment(t *testing.T) {
	uri := "scheme://host.com/path/to/something#fragment"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityWithUserinfoNoPortWithAbsolutePathNoQueryWithFragment(t *testing.T) {
	uri := "scheme://user@host.com/path/to/something#fragment"
	expected := "user@host.com"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostWithUserinfoNoPortWithAbsolutePathNoQueryWithFragment(t *testing.T) {
	uri := "scheme://user@host.com/path/to/something#fragment"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityNoUserinfoWithPortWithAbsolutePathNoQueryWithFragment(t *testing.T) {
	uri := "scheme://host.com:123/path/to/something#fragment"
	expected := "host.com:123"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostNoUserinfoWithPortWithAbsolutePathNoQueryWithFragment(t *testing.T) {
	uri := "scheme://host.com:123/path/to/something#fragment"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityWithUserinfoWithPortWithAbsolutePathNoQueryWithFragment(t *testing.T) {
	uri := "scheme://user@host.com:123/path/to/something#fragment"
	expected := "user@host.com:123"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostWithUserinfoWithPortWithAbsolutePathNoQueryWithFragment(t *testing.T) {
	uri := "scheme://user@host.com:123/path/to/something#fragment"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityNoUserinfoNoPortNoAbsolutePathWithQueryWithFragment(t *testing.T) {
	uri := "scheme://host.com?query=something#fragment"
	expected := "host.com"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostNoUserinfoNoPortNoAbsolutePathWithQueryWithFragment(t *testing.T) {
	uri := "scheme://host.com?query=something#fragment"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityWithUserinfoNoPortNoAbsolutePathWithQueryWithFragment(t *testing.T) {
	uri := "scheme://user@host.com?query=something#fragment"
	expected := "user@host.com"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostWithUserinfoNoPortNoAbsolutePathWithQueryWithFragment(t *testing.T) {
	uri := "scheme://user@host.com?query=something#fragment"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityNoUserinfoWithPortNoAbsolutePathWithQueryWithFragment(t *testing.T) {
	uri := "scheme://host.com:123?query=something#fragment"
	expected := "host.com:123"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostNoUserinfoWithPortNoAbsolutePathWithQueryWithFragment(t *testing.T) {
	uri := "scheme://host.com:123?query=something#fragment"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityWithUserinfoWithPortNoAbsolutePathWithQueryWithFragment(t *testing.T) {
	uri := "scheme://user@host.com:123?query=something#fragment"
	expected := "user@host.com:123"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostWithUserinfoWithPortNoAbsolutePathWithQueryWithFragment(t *testing.T) {
	uri := "scheme://user@host.com:123?query=something#fragment"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityNoUserinfoNoPortWithAbsolutePathWithQueryWithFragment(t *testing.T) {
	uri := "scheme://host.com/path/to/something?query=something#fragment"
	expected := "host.com"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostNoUserinfoNoPortWithAbsolutePathWithQueryWithFragment(t *testing.T) {
	uri := "scheme://host.com/path/to/something?query=something#fragment"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityWithUserinfoNoPortWithAbsolutePathWithQueryWithFragment(t *testing.T) {
	uri := "scheme://user@host.com/path/to/something?query=something#fragment"
	expected := "user@host.com"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostWithUserinfoNoPortWithAbsolutePathWithQueryWithFragment(t *testing.T) {
	uri := "scheme://user@host.com/path/to/something?query=something#fragment"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityNoUserinfoWithPortWithAbsolutePathWithQueryWithFragment(t *testing.T) {
	uri := "scheme://host.com:123/path/to/something?query=something#fragment"
	expected := "host.com:123"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostNoUserinfoWithPortWithAbsolutePathWithQueryWithFragment(t *testing.T) {
	uri := "scheme://host.com:123/path/to/something?query=something#fragment"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetAuthorityWithUserinfoWithPortWithAbsolutePathWithQueryWithFragment(t *testing.T) {
	uri := "scheme://user@host.com:123/path/to/something?query=something#fragment"
	expected := "user@host.com:123"
	actual := GetAuthority(uri)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestGetHostWithUserinfoWithPortWithAbsolutePathWithQueryWithFragment(t *testing.T) {
	uri := "scheme://user@host.com:123/path/to/something?query=something#fragment"
	expected := "host.com"
	authority := GetAuthority(uri)
	actual := GetHost(authority)
	if expected != actual {
		t.Error(
			"For", uri,
			"expected", expected,
			"got", actual,
		)
	}
}
