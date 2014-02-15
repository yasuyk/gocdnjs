package main

func FixturePackage() *Package {
	return &Package{"960gs", "960.css", "0",
		"The 960 Grid System is an effort to streamline web development workflow by providing commonly used dimensions, based on a width of 960 pixels.",
		"http://960.gs",
		[]string{"960", "960gs", "grid system"}, "",
		[]Repositories{Repositories{"git", "https://github.com/nathansmith/960-Grid-System/blob/master/code/css/960.css"}},
		[]Assets{Assets{"0", []string{"960.css", "960.min.css"}}}}
}

func FixturePackageWithTwoAsserts() *Package {
	return &Package{"960gs", "960.css", "0",
		"The 960 Grid System is an effort to streamline web development workflow by providing commonly used dimensions, based on a width of 960 pixels.",
		"http://960.gs",
		[]string{"960", "960gs", "grid system"}, "",
		[]Repositories{Repositories{"git", "https://github.com/nathansmith/960-Grid-System/blob/master/code/css/960.css"}},
		[]Assets{
			Assets{"1.0.0", []string{"960.css", "960.min.css"}},
			Assets{"2.0.0", []string{"960.css", "960.min.css"}}}}
}
