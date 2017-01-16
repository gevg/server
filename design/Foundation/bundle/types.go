package design

import (
	. "github.com/goa-fhir/server/design/DataType/data_types"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Bundle is the user resource media type.
var Bundle = Type("Bundle", func() {
	Description("A user of the API")
	Required("type")

	Attribute("type", String, "document | message | transaction | transaction-response | batch | batch-response | history | searchset | collection. See http://hl7.org/fhir/ValueSet/bundle-type ", func() {
		Enum("ocument", "message", "transaction", "transaction-response", "batch", "batch-response", "history", "searchset", "collection")
	})
	Attribute("identifier", Identifier, "Persistent identifier for the bundle.", func() {
	})
	Attribute("total", Integer, "If search, the total number of matches.", func() {
	})
	Attribute("link", ArrayOf(BundleLink), "Link to another patient resource that concerns the same actual person", func() {
	})
	Attribute("entry", Entry, "Entry in the bundle - will have a resource, or information.", func() {
	})
	Attribute("signature", Signature, "Digital Signature.", func() {
	})

})

var Entry = Type("Entry", func() {
	Description("Entry in the bundle - will have a resource, or information")

	Attribute("link", BundleLink, "Persistent identifier for the bundle.", func() {
	})
	Attribute("fullUrl", String, "Absolute URL for resource (server address, or UUID/OID).", func() {
		//Type: uri
	})
	Attribute("resource", Resource, "A resource in the bundle.", func() {
	})
	Attribute("search", Search, ".", func() {
	})
	Attribute("request", BundleRequest, ".", func() {
	})
	Attribute("response", BundleResponse, ".", func() {
	})
})

var BundleLink = Type("BundleLink", func() {
	Description("Link to another patient resource that concerns the same actual patient.")
	Required("type")
	//Comments: There is no assumption that linked patient records have mutual links.
	//Reason for inclusion or contrainment:There are multiple usecases: * Duplicate patient records due to the clerical errors associated with the difficulties of identifying humans consistently, and * Distribution of patient information across multiple servers.
	Attribute("other", HL7Reference, "The other patient resource that the link refers to.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("type", String, "The type of link between this patient resource and another patient resource. See http://hl7.org/fhir/ValueSet/link-type", func() {
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("replace", "refer", "See also")
	})
})

var Search = Type("Search", func() {
	Description("Search related information.")
	Required("mode")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("mode", HL7Reference, "The other patient resource that the link refers to.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("match", "include", "outcome")
	})
	Attribute("score", Number, "Search ranking (between 0 and 1)", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
})

var BundleRequest = Type("Request", func() {
	Description("Transaction Related Information.")
	Required("method")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("method", String, "GET | POST | PUT | DELETE. See http://hl7.org/fhir/ValueSet/http-verb", func() {
		//Type: code
		//Comments:
		//Reason for inclusion or contrainment:
		Enum("GET", "POST", "PUT", "DELETE")
	})
	Attribute("url", String, "URL for HTTP equivalent of this entry.", func() {
		//Type: uri
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("ifNoneMatch", String, "For managing cache currency.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("ifModifiedSince", DateTime, "For managing update contention.", func() {
		//Type: instant
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("ifMatch", String, "For managing update contention.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("ifNoneExist", String, "For conditional creates.", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})

})

var BundleResponse = Type("BundleResponse", func() {
	Description("Transaction Related Information.")
	//Comments:
	//Reason for inclusion or contrainment:
	Attribute("status", String, "Status response code (text optional).", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("location", String, "The location, if the operation returns a location", func() {
		//Type: uri
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("etag", String, "The etag for the resource (if relevant).", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("status", String, "Status response code (text optional).", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("lastModified", DateTime, "Status response code (text optional).", func() {
		//Type: instant
		//Comments:
		//Reason for inclusion or contrainment:
	})
	Attribute("outcome", HL7Resource, "OperationOutcome with hints and warnings (for batch/transaction).", func() {
		//Comments:
		//Reason for inclusion or contrainment:
	})
})
