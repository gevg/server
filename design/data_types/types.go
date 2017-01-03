package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

//	. "github.com/goa-fhir/server/design/data_types"

// PatientPayload defines the data structure used in the create vital request body.
// It is also the base type for the vital media type used to render bottles.
var HumanNamePayload = Type("HumanNamePayload", func() {
	Description("A human's name with the ability to identify parts and usage.")
	//Comments:Names may be changed, or repudiated, or people may have different names in different contexts. Names may be divided into parts of different
	//type that have variable significance depending on context, though the division into parts does not always matter. With personal names, the different
	//parts may or may not be imbued with some implicit meaning; various cultures associate different importance with the name parts and the degree to which
	//systems must care about name parts around the world varies widely.
	//Reason for inclusion or contrainment:
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
})

var IdentifierPayload = Type("IdentifierPayload", func() {
	Description("A technical identifier - identifies some entity uniquely and unambiguously.")
	//Comments:
	//Reason for inclusion or contrainment:
	Required("use")
	Attribute("use", String, "Identifies the purpose for this identifier, if known.", func() {
		//This is labeled as "Is Modifier" because applications should not mistake a temporary id for a permanent one.
		//Applications can assume that an identifier is permanent unless it explicitly says that it is temporary.
		Enum("usual", "official", "temp", "secondary (If known)")
		Example("usual")
	})
	Attribute("CodeableConcept", CodeableConcept, "A coded type for the identifier that can be used to determine which identifier to use for a specific purpose.", func() {
		//Comments: This element deals only with general categories of identifiers.  It SHOULD not be used for codes that correspond 1..1 with the Identifier.system.
		//Some identifiers may fall into multiple categories due to common usage.   Where the system is known, a type is unnecessary because the type is always part
		//of the system definition. However systems often need to handle identifiers where the system is not known. There is not a 1:1 relationship between type and system,
		//since many different systems have the same type.
		//Reason for inclusion or contrainment: Allows users to make use of identifiers when the identifier system is not known.
	})
	Attribute("period", Period, "Time period during which identifier is/was valid for use.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("assigner", HL7Reference, "Organization that issued/manages the identifier.", func() {
		//Comments: The reference may be just a text description of the assigner.
		//Reason for inclusion or contrainment:
	})
})
