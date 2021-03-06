package samb

// Project has your samb
// project directives.
type Project struct {
	Server          Server
	Require, Import []string
	Provider        []Global
	Package         string
	Author          string
	Packages        []string
	Routes          Routes
	Templates       Templates
	Global          []Global
}

type Global struct {
	Name, Type, Return, Comment string
}

type Server struct {
	Host, Key string
	Port      string
	Webroot   string
	Require   []string
	// Routes field will enable
	// route nesting.
	Routes Routes
	Start  Go
	// Functions to be invoked
	// on panic.
	Recover  Go
	Init     Go
	Shutdown Go
}

type Routes struct {
	Route   []Route
	Provide []string
}

type Route struct {
	Method, Path string
	Provide      []string
	Route        []Route
	Handler      string
	Go           Go
}

// Array of Go statements
// to be executed.
type Go struct {
	Do []string
}

type Templates struct {
	Template []Template
}

type Template struct {
	FilePath, Type, Name string
}
