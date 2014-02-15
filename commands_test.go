package main

import (
	"fmt"
	"testing"
)

func ExampleInfoString() {
	e := FixturePackageWithTwoAsserts()
	fmt.Print(infoString(e, true))
	// Output:
	// (package "960gs" version "1.0.0" homepage "http://960.gs" description "The 960 Grid System is an effort to streamline web development workflow by providing commonly used dimensions, based on a width of 960 pixels." assets (("1.0.0" ("http://cdnjs.cloudflare.com/ajax/libs/960gs/1.0.0/960.css" "http://cdnjs.cloudflare.com/ajax/libs/960gs/1.0.0/960.min.css" ))("2.0.0" ("http://cdnjs.cloudflare.com/ajax/libs/960gs/2.0.0/960.css" "http://cdnjs.cloudflare.com/ajax/libs/960gs/2.0.0/960.min.css" ))))

}

func TestShowInfoWithoutPlistOption(t *testing.T) {
	e := `960gs

HOMEPAGE:
   http://960.gs

DESCRIPTION:
   The 960 Grid System is an effort to streamline web development workflow by providing commonly used dimensions, based on a width of 960 pixels.

ASSETS:
   http://cdnjs.cloudflare.com/ajax/libs/960gs/1.0.0/960.css
   http://cdnjs.cloudflare.com/ajax/libs/960gs/1.0.0/960.min.css

   http://cdnjs.cloudflare.com/ajax/libs/960gs/2.0.0/960.css
   http://cdnjs.cloudflare.com/ajax/libs/960gs/2.0.0/960.min.css

`

	if a := infoString(FixturePackageWithTwoAsserts(), false); a != e {
		t.Errorf("%s, want %s", a, e)
	}
}
