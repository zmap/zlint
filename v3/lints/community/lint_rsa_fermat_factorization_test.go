package community

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCheckPrimeFactorsTooClose(t *testing.T) {
	data := []struct {
		p          *big.Int
		q          *big.Int
		n          *big.Int
		roundsLow  int
		roundsHigh int
	}{
		{
			p:          big.NewInt(101),
			q:          big.NewInt(59),
			n:          big.NewInt(5959),
			roundsLow:  2,
			roundsHigh: 3,
		},
		{
			p: bigIntOrDie("12451309173743450529024753538187635497858772172998414407116324997634262083672423797183640278969532658774374576700091736519352600717664126766443002156788367"),
			q: bigIntOrDie("12451309173743450529024753538187635497858772172998414407116324997634262083672423797183640278969532658774374576700091736519352600717664126766443002156788337"),
			n: big.NewInt(0).Mul(
				bigIntOrDie("12451309173743450529024753538187635497858772172998414407116324997634262083672423797183640278969532658774374576700091736519352600717664126766443002156788367"),
				bigIntOrDie("12451309173743450529024753538187635497858772172998414407116324997634262083672423797183640278969532658774374576700091736519352600717664126766443002156788337")),
			roundsLow:  0,
			roundsHigh: 1,
		},
		{
			p: bigIntOrDie("11779932606551869095289494662458707049283241949932278009554252037480401854504909149712949171865707598142483830639739537075502512627849249573564209082969463"),
			q: bigIntOrDie("11779932606551869095289494662458707049283241949932278009554252037480401854503793357623711855670284027157475142731886267090836872063809791989556295953329083"),
			n: big.NewInt(0).Mul(
				bigIntOrDie("11779932606551869095289494662458707049283241949932278009554252037480401854504909149712949171865707598142483830639739537075502512627849249573564209082969463"),
				bigIntOrDie("11779932606551869095289494662458707049283241949932278009554252037480401854503793357623711855670284027157475142731886267090836872063809791989556295953329083")),
			roundsLow:  13,
			roundsHigh: 14,
		},
	}
	for _, test := range data {
		test := test
		t.Run(test.n.String(), func(t *testing.T) {
			err := checkPrimeFactorsTooClose(test.n, test.roundsLow)
			if err != nil {
				t.Fatalf("factored n = %s in too few iterations, factored in %d", test.n, test.roundsLow)
			}
			err = checkPrimeFactorsTooClose(test.n, test.roundsHigh)
			if err == nil {
				t.Fatalf("failed to factor %s in %d rounds", test.n, test.roundsHigh)
			}
			errString := err.Error()
			wantP := fmt.Sprintf("p: %s", test.p)
			wantQ := fmt.Sprintf("q: %s", test.q)
			if !strings.Contains(errString, wantP) {
				t.Fatalf("unexpected p for n = %s, wanted '%s' but got %s", test.n, wantP, errString)
			}
			if !strings.Contains(errString, wantQ) {
				t.Fatalf("unexpected q for n = %s, wanted '%s' but got %s", test.n, wantQ, errString)
			}
		})
	}
}

func bigIntOrDie(from string) *big.Int {
	b, ok := big.NewInt(0).SetString(from, 10)
	if !ok {
		panic(fmt.Sprintf("failed to construct prime from string '%s'", from))
	}
	return b
}

func TestFailFermatFactorizationWithCert(t *testing.T) {
	// p: 12451309173743450529024753538187635497858772172998414407116324997634262083672423797183640278969532658774374576700091736519352600717664126766443002156788367
	// q: 12451309173743450529024753538187635497858772172998414407116324997634262083672423797183640278969532658774374576700091736519352600717664126766443002156788337
	inputPath := "rsaFermatFactorizationSusceptible.pem"
	expected := lint.Error
	out := test.TestLint("e_rsa_fermat_factorization", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestPassFermatFactorizationWithCert(t *testing.T) {
	// This is actually most useful as a benchmark to tune rounds.
	// Any RSA cert was randomly chosen.
	inputPath := "rsassapssWithSHA512.pem"
	expected := lint.Pass
	out := test.TestLint("e_rsa_fermat_factorization", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func BenchmarkFermatFactorization_Execute(b *testing.B) {
	//	cpu: Intel(R) Core(TM) i7-6700K CPU @ 4.00GHz
	//	BenchmarkFermatFactorization_Execute
	//	BenchmarkFermatFactorization_Execute/0
	//	BenchmarkFermatFactorization_Execute/0-8 	1000000000	         0.0005302 ns/o
	cert := test.ReadTestCert("rsassapssWithSHA512.pem")
	config, err := lint.NewConfigFromString(`
[e_rsa_fermat_factorization]
Rounds = 100
`)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		b.Run(strconv.FormatInt(int64(i), 10), func(b *testing.B) {
			test.TestLintCert("e_rsa_fermat_factorization", cert, config)
		})
	}
}
