package otpauth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var (
	intToBytes              Fixture
	zeroPaddingDigits6      Fixture
	zeroPaddingDigits8      Fixture
	decodeBase32            Fixture
	decodeBase32WithPadding Fixture
	encodeBase32            Fixture
	hmacHash                Fixture
	hotpValue               Fixture
)

type Fixture struct {
	TestCases []testCase `json:"test_cases"`
}

type testCase struct {
	Value  string
	Result string
}

func (tc *testCase) UnmarshalJSON(data []byte) error {
	var raw interface{}
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	if err := d.Decode(&raw); err != nil {
		return err
	}

	switch v := raw.(type) {
	case string:
		tc.Result = v
	case json.Number:
		tc.Result = v.String()
	case map[string]interface{}:
		res, ok := v["result"]
		if !ok {
			return errors.New("need result property")
		}
		switch r := res.(type) {
		case string:
			tc.Result = r
		case json.Number:
			tc.Result = r.String()
		default:
			return errors.New("invalid result property")
		}
		if val, ok := v["value"]; ok {
			switch v := val.(type) {
			case string:
				tc.Value = v
			case json.Number:
				tc.Value = v.String()
			default:
			}
		}
	default:
		tc = nil
	}
	return nil
}

func TestMain(m *testing.M) {
	log.Println("setup")

	var err error
	for _, f := range []struct {
		Fixture *Fixture
		Name    string
	}{
		{&intToBytes, "int_to_bytes"},
		{&zeroPaddingDigits6, "zero_padding"},
		{&zeroPaddingDigits8, "zero_padding_digits_8"},
		{&decodeBase32, "decode_base32"},
		{&decodeBase32WithPadding, "decode_base32_with_padding"},
		{&encodeBase32, "encode_base32"},
		{&hmacHash, "hmac_hash"},
		{&hotpValue, "hotp_value"},
	} {
		if *f.Fixture, err = readTestCaseJSON(f.Name); err != nil {
			log.Printf("Failed load fixture: testdata/%s.json, Error: %s\n", f.Name, err)
		}
	}

	code := m.Run()

	log.Println("teardown")

	os.Exit(code)
}

func readTestCaseJSON(name string) (f Fixture, err error) {
	b, err := ioutil.ReadFile(fmt.Sprintf("testdata/%s.json", name))
	if err != nil {
		return
	}

	if err = json.Unmarshal(b, &f); err != nil {
		return
	}

	return f, nil
}
