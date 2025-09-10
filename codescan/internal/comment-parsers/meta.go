package parsers

import (
	"go/ast"

	oaispec "github.com/go-openapi/spec"
)

// MetaSection parses swagger spec metadata from a go comment.
type MetaSection struct {
	Comments *ast.CommentGroup
}

type MetaParser struct {
	*SectionedParser
}

func NewMetaParser(swspec *oaispec.Swagger) *MetaParser {
	sp := NewSectionedParser()
	if swspec.Info == nil {
		swspec.Info = new(oaispec.Info)
	}
	match := NewMatcher()
	info := swspec.Info

	sp.setTitle = func(lines []string) {
		tosave := match.JoinDropLast(lines)
		if len(tosave) > 0 {
			tosave = rxStripTitleComments.ReplaceAllString(tosave, "")
		}
		info.Title = tosave
	}

	sp.setDescription = func(lines []string) { info.Description = JoinDropLast(lines) }
	sp.taggers = []TagParser{
		NewMultiLineTagParser("TOS", NewMultilineDropEmptyParser(rxTOS, metaTOSSetter(info)), false),
		NewMultiLineTagParser("Consumes", NewMultilineDropEmptyParser(rxConsumes, metaConsumesSetter(swspec)), false),
		NewMultiLineTagParser("Produces", NewMultilineDropEmptyParser(rxProduces, metaProducesSetter(swspec)), false),
		NewSingleLineTagParser("Schemes", NewSetSchemesParser(metaSchemeSetter(swspec))),
		NewMultiLineTagParser("Security", NewSetSecurityParser(rxSecuritySchemes, metaSecuritySetter(swspec)), false),
		NewMultiLineTagParser("SecurityDefinitions", NewYAMLParser(rxSecurity, metaSecurityDefinitionsSetter(swspec)), true),
		NewSingleLineTagParser("Version", &setMetaSingle{swspec, rxVersion, setInfoVersion}),
		NewSingleLineTagParser("Host", &setMetaSingle{swspec, rxHost, setSwaggerHost}),
		NewSingleLineTagParser("BasePath", &setMetaSingle{swspec, rxBasePath, setSwaggerBasePath}),
		NewSingleLineTagParser("Contact", &setMetaSingle{swspec, rxContact, setInfoContact}),
		NewSingleLineTagParser("License", &setMetaSingle{swspec, rxLicense, setInfoLicense}),
		NewMultiLineTagParser("YAMLInfoExtensionsBlock", NewYAMLParser(rxInfoExtensions, infoVendorExtensibleSetter(swspec)), true),
		NewMultiLineTagParser("YAMLExtensionsBlock", NewYAMLParser(rxExtensions, metaVendorExtensibleSetter(swspec)), true),
	}

	return &MetaParser{
		SectionedParser: sp,
	}
}
