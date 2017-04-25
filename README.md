# server
goa on FHIR Server [![Build Status](https://travis-ci.org/intervention-engine/fhir.svg?branch=master)](https://travis-ci.org/intervention-engine/fhir)[![GoDoc](https://godoc.org/github.com/intervention-engine/fhir?status.svg)](https://godoc.org/github.com/intervention-engine/fhir)
===================================================================================================================================================================

This project provides [HL7 FHIR DSTU2](http://hl7.org/fhir/DSTU2/index.html) models and server components implemented in Go and using Postgres as storage. This is a
library that can be embedded into other server applications. The library is a complete implementation of FHIR.

Development
-----------

This project uses Go 1.7, goa 1.1.0, gorma 1.0 and Postgres. To test the library, first, install all of the dependencies:

```
$ goagen bootstrap -d github.com/jamesallain/goa-fhir/design
$ goagen --design=github.com/jamesallain/goa-fhir/design gen --pkg-path=github.com/goadesign/gorma

$ go get -t ./...
```

Once the dependencies are installed, you should make sure that Postgres is also running. The test suite
will create a `fhir-test` database in a local instance of MongoDB to execute some tests. You can then
run the test suite with the following:

```
$ go test ./...
```

Usage
-----

Folder Structure
-------
	|-- Folder [FHIR Resource (eg Patient)]
		|-- File [Goa MediaType]
	    |-- File [Goa Type]
	|-- File [Goa Resource]

File Format
-------
-   MediaType
   -   example: var GoaMediaType = MediaType("application/vnd.goamediatype+json", func() {
   -   Description("insert FHIR Definition")
   -     //Comments: insert FHIR Comments 
   -     //Reason for inclusion or contrainment: insert FHIR Reason for inclusion or contrainment
   -  Attributes(func() {
		Required("list of required Attributes")
   -	Attribute("AttibuteName", AttributeType, `FHIR Definition`, func() {
   -       //Comments: insert FHIR Comments 
   -       //Reason for inclusion or contrainment: insert FHIR Reason for inclusion or contrainment
   -  })

