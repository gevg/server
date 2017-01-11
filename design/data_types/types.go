package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var BackboneElement = Type("BackboneElement", func() {
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("element", Element)
	Attribute("modifierExtension", Extension)
})
var Element = Type("Element", func() {
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("id", String)
	Attribute("extension", Extension)
})
var Extension = Type("Extension", func() {
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("url", String)
	Attribute("ValueAddress", Address)
	Attribute("ValueAnnotation", Annotation)
	Attribute("ValueAttachment", Attachment)
	Attribute("ValueBase64Binary", String)
	Attribute("ValueBoolean", Boolean)
	Attribute("ValueCode", String)
	Attribute("ValueCodeableConcept", CodeableConcept)
	Attribute("ValueCoding", Coding)
	Attribute("ValueContactPoint", String) //ContactPoint
	Attribute("ValueDate", DateTime)
	Attribute("ValueDateTime", DateTime)
	Attribute("ValueDecimal", Number) //float64
	Attribute("ValueHumanName", HumanName)
	Attribute("ValueId", String)
	Attribute("ValueIdentifier", Identifier)
	Attribute("ValueInstant", DateTime)
	Attribute("ValueInteger", Integer)
	Attribute("ValueMarkdown", String)
	Attribute("ValueMeta", Meta)
	Attribute("ValueOid", String)
	Attribute("ValuePeriod", Period)
	Attribute("ValuePositiveInt", Number) //unit32
	Attribute("ValueQuantity", Quantity)
	Attribute("ValueRange", String)  //Range
	Attribute("ValueRatio", Integer) //Ratio
	Attribute("ValueReference", HL7Reference)
	Attribute("ValueSampledData", String) //SampleData
	Attribute("ValueSignature", String)   //Signature
	Attribute("ValueString", String)
	Attribute("ValueTime", DateTime)
	Attribute("ValueTiming", DateTime)    //Timing
	Attribute("ValueUnsignedInt", Number) //uint32
	Attribute("ValueUri", String)
})
var Attachment = Type("Attachment", func() {
	Description("For referring to data content defined in other formats.")
	//Comments: When providing a summary view (for example with Observation.value[x]) Attachment should be represented with a brief display text such as "Attachment".
	//Reason for inclusion or contrainment:
	Required("contentType")
	Attribute("contentType", String, `Identifies the type of the data in the attachment and allows a method to be chosen to interpret or render the data. 
				Includes mime type parameters such as charset where appropriate. See http://www.rfc-editor.org/bcp/bcp13.txt`, func() {
		//Comments:
		//Reason for inclusion or contrainment: Processors of the data need to be able to know how to interpret the data.
	})
	Attribute("data", String, "The actual data of the attachment - a sequence of bytes. In XML, represented using base64.", func() {
		//type: base64Binary
		//Comments: The base64-encoded data SHALL be expressed in the same character set as the base resource XML or JSON.
		//Reason for inclusion or contrainment: The data needs to able to be transmitted inline.
	})
	Attribute("size", Integer, "The number of bytes of data that make up this attachment.", func() {
		//type: unsignedInt
		//Comments: The number of bytes is redundant if the data is provided as a base64binary, but is useful if the data is provided as a url reference.
		//Reason for inclusion or contrainment: Representing the size allows applications to determine whether they should fetch the content automatically in advance, or
		//refuse to fetch it at all.
	})
	Attribute("hash", String, "The calculated hash of the data using SHA-1. Represented using base64.", func() {
		//type: base64Binary
		//Comments:
		//Reason for inclusion or contrainment: Included so that applications can verify that the contents of a location have not changed and so that a signature of the content can
		//implicitly sign the content of an image without having to include the data in the instance or reference the url in the signature.
	})
	Attribute("title", String, "A label or set of text to display in place of the data.", func() {
		//Comments:
		//Reason for inclusion or contrainment: Applications need a label to display to a human user in place of the actual data if the data cannot be rendered or perceived by the viewer.
	})
	Attribute("creation", DateTime, "The date that the attachment was first created.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
var Address = Type("Address", func() {
	Description("There is a variety of postal address formats defined around the world. This format defines a superset that is the basis for all addresses around the world.")
	//Comments: Note: address is for postal addresses, not physical locations.
	//Reason for inclusion or contrainment:
	Required("use", "type")
	Attribute("use", String, "The use of an address", func() {
		//Comments: This is labeled as "Is Modifier" because applications should not mistake a temporary or old address etc.for a current/permanent one.
		//Applications can assume that an address is current unless it explicitly says that it is temporary or old.
		//Reason for inclusion or contrainment: Allows an appropriate address to be chosen from a list of many.
		Enum("home", "work", "temp", "old - purpose of this address")
		Example("work")
	})
	Attribute("type", String, "Distinguishes between physical addresses (those you can visit) and mailing addresses (e.g. PO Boxes and care-of addresses). Most addresses are both.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("postal", "physical", "both")
		Example("postal")
	})
	Attribute("text", String, "Text representation of the address.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("line", ArrayOf(String), "This component contains the house number, apartment number, street name, street direction,  P.O. Box number, delivery hints, and similar address information.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("city", String, "The name of the city, town, village or other community or delivery center.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("distinct", String, "District is sometimes known as county, but in some regions 'county' in used in place of city (municipality), so county name should be conveyed in city instead.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("state", String, "Sub-unit of a country with limited sovereignty in a federally organized country. A code may be used if codes are in common use (i.e. US 2 letter state codes).", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("postalCode", String, "A postal code designating a region defined by the postal service.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("country", String, "Country (can be ISO 3166 3 letter code)", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("period", Period, "Time period when address was/is in use.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
var Annotation = Type("Annotation", func() {
	Description("A  text note which also  contains information about who made the statement and when.")
	//Comments: For systems that do not have structured annotations, they can simply communicate a single annotation with no author or time.
	//This element may need to be included in narrative because of the potential for modifying information.  *Annotations SHOULD NOT* be used to communicate "modifying"
	//information that could be computable. (This is a SHOULD because enforcing user behavior is nearly impossible).
	//Reason for inclusion or contrainment:
	Attribute("authorReference", HL7Reference, "The individual responsible for making the annotation.", func() {
		//Type: Referece to Practioner, Patient, RelatedPerson; String
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("authorString", String, "The individual responsible for making the annotation.", func() {
		//Type: Referece to Practioner, Patient, RelatedPerson; String
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("time", DateTime, "Indicates when this particular annotation was made.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
var Coding = Type("Coding", func() {
	Description("A reference to a code defined by a terminology system.")
	//Comments: Codes may be defined very casually in enumerations or code lists, up to very formal definitions such as SNOMED CT - See the HL7 v3 Core Principles for more information.
	//Reason for inclusion or contrainment:
	Attribute("system", String, "The identification of the code system that defines the meaning of the symbol in the code.Identity of the terminology system", func() {
		//Comments: The URI may be an OID (urn:oid:...) or a UUID (urn:uuid:...).  OIDs and UUIDs SHALL be references to the HL7 OID registry. Otherwise,
		//the URI should come from HL7's list of FHIR defined special URIs or it should de-reference to some definition that establish the system clearly and unambiguously.
		//Reason for inclusion or contrainment: Need to be unambiguous about the source of the definition of the symbol.
	})
	Attribute("version", String, `The version of the code system which was used when choosing this code. Note that a well-maintained code system does not need the 
		version reported, because the meaning of codes is consistent across versions. However this cannot consistently be assured. And when the meaning is not
		guaranteed to be consistent, the version SHOULD be exchanged.`, func() {
		//Comments: Where the terminology does not clearly define what string should be used to identify code system versions, the recommendation is to use the date
		//(expressed in FHIR date format) on which that version was officially published as the version date.
		//Reason for inclusion or contrainment:
	})
	Attribute("code", String, "A symbol in syntax defined by the system. The symbol may be a predefined code or an expression in a syntax defined by the coding system (e.g. post-coordination).", func() {
		//type: code, but that is currently (12/30/16) not a data type. Intervention Engine is using a String here too.
		//Comments:
		//Reason for inclusion or contrainment: "Need to refer to a particular code in the system."
	})
	Attribute("display", String, "A representation of the meaning of the code in the system, following the rules of the system.", func() {
		//Comments:
		//Reason for inclusion or contrainment: Need to be able to carry a human-readable meaning of the code for readers that do not know  the system.
	})
	Attribute("userSelected", Boolean, "Indicates that this coding was chosen by a user directly - i.e. off a pick list of available items (codes or displays).", func() {
		//Comments: Amongst a set of alternatives, a directly chosen code is the most appropriate starting point for new translations.
		//There is some ambiguity about what exactly 'directly chosen' implies, and trading partner agreement may be needed to clarify the use of this element
		//and its consequences more completely.
		//Reason for inclusion or contrainment: This has been identified as a clinical safety criterium - that this exact system/code pair was chosen explicitly,
		//rather than inferred by the system based on some rules or language processing.
	})
})
var CodeableConcept = Type("CodeableConcept", func() {
	Description("A concept that may be defined by a formal reference to a terminology or ontology or may be provided by text.")
	//Comments:Not all terminology uses fit this general pattern. In some cases, models should not use CodeableConcept and use Coding directly and provide their own structure for managing text,
	//codings, translations and the relationship between elements and pre- and post-coordination.
	//Reason for inclusion or contrainment:
	Attribute("coding", ArrayOf(Coding), "A concept that may be defined by a formal reference to a terminology or ontology or may be provided by text.", func() {
		//Comments: Codes may be defined very casually in enumerations, or code lists, up to very formal definitions such as SNOMED CT - See the HL7 v3 Core Principles for more information.
		//Ordering of codings is undefined and SHALL NOT be used to infer meaning. Generally, at most only one of the coding values will be labelled as UserSelected = true.
		//Reason for inclusion or contrainment: Allows for translations and alternate encodings within a code system.  Also supports communication of the same
		//instance to systems requiring different encodings.
	})
	Attribute("text", String, `The version of the code system which was used when choosing this code. Note that a well-maintained code system does not need the 
	version reported, because the meaning of codes is consistent across versions. However this cannot consistently be assured. And when the meaning is not
	guaranteed to be consistent, the version SHOULD be exchanged.`, func() {
		//Comments: Where the terminology does not clearly define what string should be used to identify code system versions, the recommendation is to use the date
		//(expressed in FHIR date format) on which that version was officially published as the version date.
		//Reason for inclusion or contrainment:
	})
})
var ContactPoint = Type("ContactPoint", func() {
	Description("Details for all kinds of technology mediated contact points for a person or organization, including telephone, email, etc.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("system", String, "Telecommunications form for contact point - what communications system is required to make use of the contact.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("phone", "fax", "email", "pager", "other")
	})
	Attribute("value", String, "The actual contact point details, in a form that is meaningful to the designated communication system (i.e. phone number or email address).", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("use", String, "Identifies the purpose for the contact point.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("home", "work", "temp", "old", "mobile")
	})
	Attribute("rank", Integer, "Specifies a preferred order in which to use a set of contacts. Contacts are ranked with lower values coming before higher values.", func() {
		//Type: postiveInt
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
var DomainResource = Type("DomainResource", func() {
	Description("")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("resource", HL7Resource, "", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("text", Narrative, "", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("contained", ArrayOf(Any), "", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("extension", Extension, "", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("modifierExtension", Extension, "", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
var HL7Reference = Type("HL7Reference", func() {
	Description("A reference from one resource to another.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("reference", String, `A reference to a location at which the other resource is found. The reference may be a relative reference, in which case it is relative to 
		the service base URL, or an absolute URL that resolves to the location where the resource is found. The reference may be version specific or not. If the reference is 
		not to a FHIR RESTful server, then it should be assumed to be version specific. Internal fragment references (start with '#') refer to contained resources.`, func() {
		//Comments: Using absolute URLs provides a stable scalable approach suitable for a cloud/web context, while using relative/logical references provides a
		//flexible approach suitable for use when trading across closed eco-system boundaries.   Absolute URLs do not need to point to a FHIR RESTful server, though
		//this is the preferred approach. If the URL conforms to the structure "/[type]/[id]" then it should be assumed that the reference is to a FHIR RESTful server.
		//Reason for inclusion or contrainment:
	})
	Attribute("display", String, "Plain text narrative that identifies the resource in addition to the resource reference.", func() {
		//Comments: This is generally not the same as the Resource.text of the referenced resource.  The purpose is to identify what's being referenced, not to fully describe it.
		//Reason for inclusion or contrainment:
	})
})
var HL7Resource = Type("HL7Resource", func() {
	Description("")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("resourceType", String, "", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("display", String, "", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("tyoe", Meta, "", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("ReferenceId", String, "", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("External", Boolean, "", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
var HumanName = Type("HumanName", func() {
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
	Attribute("given", ArrayOf(String), "Given names (not always 'first'). Includes middle names", func() {
		//Comments:
		//Reason for inclusion or contrainment: If only initials are recorded, they may be used in place of the full name.  Not called "first name" since given names do not always come first.
	})
	Attribute("prefix", ArrayOf(String), "Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the start of the name.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("suffix", ArrayOf(String), "Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the end of the name.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("period", Period, "Indicates the period of time when this name was valid for the named person.", func() {
		//Comments:
		//Reason for inclusion or contrainment: Allows names to be placed in historical context.
	})
})
var Identifier = Type("Identifier", func() {
	Description("A technical identifier - identifies some entity uniquely and unambiguously.")
	//Comments:
	//Reason for inclusion or contrainment:
	Required("use")
	Attribute("use", String, "The purpose of this identifier. Use http://hl7.org/fhir/ValueSet/identifier-use", func() {
		//This is labeled as "Is Modifier" because applications should not mistake a temporary id for a permanent one.
		//Applications can assume that an identifier is permanent unless it explicitly says that it is temporary.
		Enum("usual", "official", "temp", "secondary (If known)")
	})
	Attribute("type", CodeableConcept, "A coded type for the identifier that can be used to determine which identifier to use for a specific purpose. See http://hl7.org/fhir/ValueSet/identifier-type", func() {
		//Comments: This element deals only with general categories of identifiers.  It SHOULD not be used for codes that correspond 1..1 with the Identifier.system.
		//Some identifiers may fall into multiple categories due to common usage.   Where the system is known, a type is unnecessary because the type is always part
		//of the system definition. However systems often need to handle identifiers where the system is not known. There is not a 1:1 relationship between type and system,
		//since many different systems have the same type.
		//Reason for inclusion or contrainment: Allows users to make use of identifiers when the identifier system is not known.
	})
	Attribute("system", String, "Establishes the namespace in which set of possible id values is unique.", func() {
		//Type: uri
		//Comments:
		//Reason for inclusion or contrainment:
		Format("uri")
	})
	Attribute("value", String, "The portion of the identifier typically displayed to the user and which is unique within the context of the system.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("period", Period, "Time period during which identifier is/was valid for use.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("assigner", HL7Reference, "Organization that issued/manages the identifier.", func() {
		//Referece: Organization
		//Comments: The reference may be just a text description of the assigner.
		//Reason for inclusion or contrainment:
	})
})
var Meta = Type("Meta", func() {
	Description("The metadata about a resource. This is content in the resource that is maintained by the infrastructure. Changes to the content may not always be associated with version changes to the resource.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("versionId", String, "The version specific identifier, as it appears in the version portion of the URL. This values changes when the resource is created, updated, or deleted.", func() {
		//Type: id
		//Comments: The server assigns this value, and ignores what the client specifies, except in the case that the server is imposing version integrity on updates/deletes.
		//Reason for inclusion or contrainment:
	})
	Attribute("lastUpdated", DateTime, "When the resource last changed - e.g. when the version changed.", func() {
		//Type: instant
		//Comments: This value is always populated except when the resource is first being created. The server / resource manager sets this value; what a client provides is irrelevant.
		//Reason for inclusion or contrainment:
	})
	Attribute("profile", String, "A list of profiles [[[StructureDefinition]]]s that this resource claims to conform to. The URL is a reference to [[[StructureDefinition.url]]].", func() {
		//Type: uri
		//Comments: It is up to the server and/or other infrastructure of policy to determine whether/how these claims are verified and/or updated over time.  The list of profile URLs is a set.
		//Reason for inclusion or contrainment:
		Format("uri")
	})
	Attribute("security", Coding, "Security labels applied to this resource. These tags connect specific resources to the overall security policy and infrastructure.", func() {
		//Comments: The security labels can be updated without changing the stated version of the resource  The list of security labels is a set. Uniqueness is based the system/code, and version and display are ignored.
		//Reason for inclusion or contrainment:
	})
	Attribute("tag", Coding, "Tags applied to this resource. Tags are intended to be used to identify and relate resources to process and workflow, and applications are not required to consider the tags when interpreting the meaning of a resource.", func() {
		//Comments:The tags can be updated without changing the stated version of the resource.  The list of tags is a set. Uniqueness is based the system/code, and version and display are ignored.
		//Reason for inclusion or contrainment:
	})
})
var Narrative = Type("Narrative", func() {
	Description("A time period defined by a start and end date and optionally time.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("status", String, `The status of the narrative - whether it's entirely generated (from just the defined data or the extensions too),
	or whether a human authored it and it may contain additional data. See http://hl7.org/fhir/ValueSet/narrative-status.`, func() {
		//Type: code
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("generated", "extensions", "additional", "empty")
	})
	Attribute("div", String, "The actual narrative content, a stripped down version of XHTML.", func() {
		//Type: XHTML
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
var Period = Type("Period", func() {
	Description("A time period defined by a start and end date and optionally time.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("start", DateTime, "Starting time with inclusive boundary", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("end", DateTime, "End time with inclusive boundary, if not ongoing", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
var Quantity = Type("Quantity", func() {
	Description(`A measured amount (or an amount that can potentially be measured). Note that measured amounts include amounts that are not precisely quantified, 
	including amounts involving arbitrary units and floating currencies.`)
	//Comments: The context of use may frequently define what kind of quantity this is and therefore what kind of units can be used.
	//The context of use may also restrict the values for the comparator.
	//Reason for inclusion or contrainment:
	Required("comparator")
	Attribute("value", Number, "Numerical value (with implicit precision)", func() {
		//Type: decimal
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("comparator", String, `How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues;
		e.g. if the comparator is "<" , then the real value is < stated value. See http://hl7.org/fhir/ValueSet/quantity-comparator`, func() {
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("<", "<=", ">=", ">")
	})
	Attribute("unit", String, "A human-readable form of the unit.", func() {
		//Type:uri
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("system", String, "The identification of the system that provides the coded form of the unit.", func() {
		//Type:uri
		//Comments:
		//Reason for inclusion or contrainment:
		Format("uri")

	})
	Attribute("code", String, "A computer processable form of the unit in some unit representation system.", func() {
		//Type: code
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
var Range = Type("Range", func() {
	Description("A set of ordered Quantities defined by a low and high limit.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("low", Quantity, "The low limit. The boundary is inclusive.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("high", Quantity, "The high limit. The boundary is inclusive.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
var Ratio = Type("Ratio", func() {
	Description("A relationship of two Quantity values - expressed as a numerator and a denominator.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("numerator", Quantity, "The value of the numerator.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("denominator", Quantity, "The value of the denominator.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
var Repeat = Type("Repeat", func() {
	Description("Many timing schedules are determined by regular repetitions.")
	//Comments:
	//Reason for inclusion or contrainment: Many timing schedules are determined by regular repetitions.
	Required("periodUnits", "when", "durationUnits")

	Attribute("bounds", Any, "Length/Range of lengths, or (Start and/or end) limits", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("count", Integer, "A total count of the desired number of repetitions.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("duration", Number, "How long this thing happens for when it happens.", func() {
		//Type: decimal
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("durationMax", Number, "The upper limit of how long this thing happens for when it happens.", func() {
		//Type: decimal
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("durationUnits", String, "The units of time for the duration, in UCUM units. See http://hl7.org/fhir/ValueSet/units-of-time", func() {
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("s", "min", "h", "d", "wk", "mo", "a")
	})
	Attribute("frequency", Integer, "The number of times to repeat the action within the specified period / period range (i.e. both period and periodMax provided).", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("frequencyMax", Integer, "If present, indicates that the frequency is a range - so repeat between [frequency] and [frequencyMax] times within the period or period range.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("period", Number, "Event occurs frequency times per period", func() {
		//Type: decimal
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("periodMax", String, "If present, indicates that the period is a range from [period] to [periodMax], allowing expressing concepts such as 'do this once every 3-5 days.'", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("periodUnits", String, "The units of time for the period in UCUM units. See http://hl7.org/fhir/ValueSet/units-of-time", func() {
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("s", "min", "h", "d", "wk", "mo", "a")
	})
	Attribute("when", String, "A real world event that the occurrence of the event should be tied to. See http://hl7.org/fhir/ValueSet/event-timing", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
})

var SampleData = Type("SampleData", func() {
	Description("A series of measurements taken by a device, with upper and lower limits. There may be more than one dimension in the data.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("origin", Quantity, "The base quantity that a measured value of zero represents. In addition, this provides the units of the entire measurement series.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("period", Number, "The length of time between sampling times, measured in milliseconds.", func() {
		//Type: decimal
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("factor", Number, "A correction factor that is applied to the sampled data points before they are added to the origin.", func() {
		//Type: decimal
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("lowerLimit", Number, "The lower limit of detection of the measured points. This is needed if any of the data points have the value 'L' (lower than detection limit).", func() {
		//Type: decimal
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("upperLimit", Number, "The upper limit of detection of the measured points. This is needed if any of the data points have the value 'U' (higher than detection limit).", func() {
		//Type: decimal
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("dimensions", Integer, "The upper limit of detection of the measured points. This is needed if any of the data points have the value 'U' (higher than detection limit).", func() {
		//Type: positiveInt
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("data", String, "A series of data points which are decimal values separated by a single space (character u20). The special values 'E' (error), 'L' (below detection limit) and 'U' (above detection limit) can also be used in place of a decimal value.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
var Timing = Type("Timing", func() {
	Description(`Specifies an event that may occur multiple times. Timing schedules are used to record when things are expected 
	or requested to occur. The most common usage is in dosage instructions for medications. They are also used when planning care of various kinds.`)
	//Comments: A timing schedule can be either a list of events - intervals on which the event occurs, or a single event with repeating criteria or
	//just repeating criteria with no actual event.  When both event and a repeating specification are provided, the list of events should be
	//understood as an interpretation of the information in the repeat structure.
	//Reason for inclusion or contrainment:
	Attribute("event", DateTime, "Identifies specific times when the event occurs.", func() {
		//Comments: In an MAR, for instance, you need to take a general specification, and turn it into a precise specification.
		//Reason for inclusion or contrainment:
	})
	Attribute("repeat", Element, "A set of rules that describe when the event should occur.", func() {
		//Comments:
		//Reason for inclusion or contrainment: Many timing schedules are determined by regular repetitions.
	})
	Attribute("code", String, "A code for the timing pattern. Some codes such as BID are ubiquitous, but many institutions define their own additional codes. See http://hl7.org/fhir/ValueSet/timing-abbreviation", func() {
		//Type: CodeableConcept
		//Comments: A repeat should always be defined except for the common codes BID, TID, QID, AM and PM, which all systems are required to understand.
		//Reason for inclusion or contrainment:
		Enum("QD", "QOD", "Q4H", "Q6H", "BID", "TID", "QID", "AM", "PM +")
	})
})
