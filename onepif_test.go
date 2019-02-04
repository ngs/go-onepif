package onepif

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	onepif := &Onepif{
		LocationKey: "amazon.com",
		Title:       "AWS",
		Location:    "https://12345.signin.aws.amazon.com/console",
		SecureContents: []SecureContent{{
			URLs: []URL{{
				Label: "website",
				URL:   "https://12345.signin.aws.amazon.com/console",
			}},
			Fields: []FormField{
				{Value: "ngs", Name: "username", Type: "T", Designation: "username"},
			},
			Sections: []Section{{
				Title: "Credentials",
				Fields: []Field{
					{Kind: "string", Value: "Foo", Title: "Bar"},
				}},
			},
		}},
	}
	str, err := json.MarshalIndent(onepif, "", "  ")
	if err != nil {
		t.Error(err)
	}
	obj := Onepif{}
	err = json.Unmarshal([]byte(str), &obj)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(str))
}
