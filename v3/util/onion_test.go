package util

import "testing"

func TestIsOnionV3(t *testing.T) {
	data := []struct {
		in   string
		want bool
	}{
		{
			"*.facebookwkhpilnemxj7asaniu7vnjjbiltxjqhye3mhbshg7kx5tfyd.onion",
			true,
		},
		{
			"*.facebookwkhpilnemxj7asaniu7vnjjbiltxjqhye3mhbshg7kx5tfyd.com",
			false,
		},
		{
			// Tricky to spot, but different final byte (e instead of d)
			"*.facebookwkhpilnemxj7asaniu7vnjjbiltxjqhye3mhbshg7kx5tfye.onion",
			false,
		},
		{
			"pg6mmjiyjmcrsslvykfwnntlaru7p5svn6y2ymmju6nubxndf4pscryd.onion",
			true,
		},

		{
			"sp3k262uwy4r2k3ycr5awluarykdpag6a7y33jxop4cs2lu5uz5sseqd.onion",
			true,
		},

		{
			"xa4r2iadxm55fbnqgwwi5mymqdcofiu3w6rpbtqn7b2dyn7mgwj64jyd.onion",
			true,
		},
		{
			"facebook.onion",
			false,
		},
		{
			// Trigger bad base32 decoding with the leading #
			"#a4r2iadxm55fbnqgwwi5mymqdcofiu3w6rpbtqn7b2dyn7mgwj64jyd.onion",
			false,
		},
	}
	for _, test := range data {
		test := test
		t.Run(test.in, func(t *testing.T) {
			got := IsOnionV3(test.in)
			if got != test.want {
				t.Errorf("expected %v got %v", test.want, got)
			}
		})
	}
}

func TestAllAreOnionV3(t *testing.T) {
	data := []struct {
		in   []string
		want bool
	}{
		{
			[]string{"*.facebookwkhpilnemxj7asaniu7vnjjbiltxjqhye3mhbshg7kx5tfyd.onion"},
			true,
		},
		{
			[]string{},
			false,
		},
		{
			[]string{
				"pg6mmjiyjmcrsslvykfwnntlaru7p5svn6y2ymmju6nubxndf4pscryd.onion",
				"sp3k262uwy4r2k3ycr5awluarykdpag6a7y33jxop4cs2lu5uz5sseqd.onion",
				"xa4r2iadxm55fbnqgwwi5mymqdcofiu3w6rpbtqn7b2dyn7mgwj64jyd.onion",
			},
			true,
		},
		{
			[]string{
				"pg6mmjiyjmcrsslvykfwnntlaru7p5svn6y2ymmju6nubxndf4pscryd.onion",
				"facebook.com",
				"xa4r2iadxm55fbnqgwwi5mymqdcofiu3w6rpbtqn7b2dyn7mgwj64jyd.onion",
			},
			false,
		},
		{
			[]string{
				"facebook.com",
				"pg6mmjiyjmcrsslvykfwnntlaru7p5svn6y2ymmju6nubxndf4pscryd.onion",
				"xa4r2iadxm55fbnqgwwi5mymqdcofiu3w6rpbtqn7b2dyn7mgwj64jyd.onion",
			},
			false,
		},
		{
			[]string{
				"pg6mmjiyjmcrsslvykfwnntlaru7p5svn6y2ymmju6nubxndf4pscryd.onion",
				"xa4r2iadxm55fbnqgwwi5mymqdcofiu3w6rpbtqn7b2dyn7mgwj64jyd.onion",
				"facebook.com",
			},
			false,
		},
	}
	for _, test := range data {
		test := test
		var name string
		if len(test.in) == 0 {
			name = "empty"
		} else {
			name = test.in[0]
		}
		t.Run(name, func(t *testing.T) {
			got := AllAreOnionV3(test.in)
			if got != test.want {
				t.Errorf("expected %v got %v", test.want, got)
			}
		})
	}
}
