package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

//	. "github.com/jamesallain/goa-fhir/design/data_types"

// PatientPayload defines the data structure used in the create vital request body.
// It is also the base type for the vital media type used to render bottles.
var HumanNamePayload = Type("HumanNamePayload", func() {
	Description("A human's name with the ability to identify parts and usage.")
	//Comments:Names may be changed, or repudiated, or people may have different names in different contexts. Names may be divided into parts of different
	//type that have variable significance depending on context, though the division into parts does not always matter. With personal names, the different
	//parts may or may not be imbued with some implicit meaning; various cultures associate different importance with the name parts and the degree to which
	//systems must care about name parts around the world varies widely.
	//Reason for inclusion or contrainment:
	//Attributes(func() {
	Required("use")
	Attribute("use", String, "Identifies the purpose for this name.", func() {
		//Comments:
		//Reason for inclusion or contrainment: Allows the appropriate name for a particular context of use to be selected from among a set of names.
		Enum("usual", "official", "temp", "nickname", "anonymous", "old", "maiden")
	})
	Attribute("family", ArrayOf(String), "The part of a name that links to the genealogy. In some cultures (e.g. Eritrea) the family name of a son is the first name of his father.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("given", String, "Given names (not always 'first'). Includes middle names", func() {
		//Comments:
		//Reason for inclusion or contrainment: If only initials are recorded, they may be used in place of the full name.  Not called "first name" since given names do not always come first.
	})
	Attribute("prefix", String, "Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the start of the name.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("suffix", String, "Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the end of the name.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("period", Period, "Indicates the period of time when this name was valid for the named person.", func() {
		//Comments:
		//Reason for inclusion or contrainment: Allows names to be placed in historical context.
	})
	//})
})
