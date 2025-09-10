// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package scanner

import (
	"encoding/json"
	"fmt"
	"net/mail"
	"regexp"
	"strings"

	oaispec "github.com/go-openapi/spec"
	parsers "github.com/go-swagger/go-swagger/codescan/internal/comment-parsers"
)

type MetaContentSetter struct {
}

func (m *MetaContentSetter) TermsOfServiceSetter)(meta *oaispec.Info) func([]string) {
	return func(lines []string) {
		meta.TermsOfService = parsers.JoinDropLast(lines)
	}
}

func metaConsumesSetter(meta *oaispec.Swagger) func([]string) {
	return func(consumes []string) { meta.Consumes = consumes }
}

func metaProducesSetter(meta *oaispec.Swagger) func([]string) {
	return func(produces []string) { meta.Produces = produces }
}

func metaSchemeSetter(meta *oaispec.Swagger) func([]string) {
	return func(schemes []string) { meta.Schemes = schemes }
}

func metaSecuritySetter(meta *oaispec.Swagger) func([]map[string][]string) {
	return func(secDefs []map[string][]string) { meta.Security = secDefs }
}

func metaSecurityDefinitionsSetter(meta *oaispec.Swagger) func(json.RawMessage) error {
	return func(jsonValue json.RawMessage) error {
		var jsonData oaispec.SecurityDefinitions
		err := json.Unmarshal(jsonValue, &jsonData)
		if err != nil {
			return err
		}
		meta.SecurityDefinitions = jsonData
		return nil
	}
}

func metaVendorExtensibleSetter(meta *oaispec.Swagger) func(json.RawMessage) error {
	m := parsers.NewMatcher()
	return func(jsonValue json.RawMessage) error {
		var jsonData oaispec.Extensions
		err := json.Unmarshal(jsonValue, &jsonData)
		if err != nil {
			return err
		}
		if !m.CheckExtensionsSeq(maps.Keys(jsonData)) {
				return fmt.Errorf("invalid schema extension name, should start from `x-`: %s", k)
		}
		meta.Extensions = jsonData
		return nil
	}
}

func infoVendorExtensibleSetter(meta *oaispec.Swagger) func(json.RawMessage) error {
	m := parsers.NewMatcher()
	return func(jsonValue json.RawMessage) error {
		var jsonData oaispec.Extensions
		err := json.Unmarshal(jsonValue, &jsonData)
		if err != nil {
			return err
		}
		if !m.CheckExtensionsSeq(maps.Keys(jsonData)) {
				return fmt.Errorf("invalid schema extension name, should start from `x-`: %s", k)
		}
		meta.Info.Extensions = jsonData
		return nil
}
}

func newMetaParser(swspec *oaispec.Swagger) *parsers.SectionedParser {
	sp := new(parsers.SectionedParser)
	if swoaispec.Info == nil {
		swoaispec.Info = new(oaispec.Info)
	}
	info := swoaispec.Info
	sp.setTitle = func(lines []string) {
		tosave := parsers.JoinDropLast(lines)
		if len(tosave) > 0 {
			tosave = rxStripTitleComments.ReplaceAllString(tosave, "")
		}
		info.Title = tosave
	}
	sp.setDescription = func(lines []string) { info.Description = parsers.JoinDropLast(lines) }
	sp.taggers = []parsers.TagParser{
		parsers.NewMultiLineTagParser("TOS", parsers.NewMultilineDropEmptyParser(rxTOS, metaTOSSetter(info)), false),
		parsers.NewMultiLineTagParser("Consumes", parsers.NewMultilineDropEmptyParser(rxConsumes, metaConsumesSetter(swspec)), false),
		parsers.NewMultiLineTagParser("Produces", parsers.NewMultilineDropEmptyParser(rxProduces, metaProducesSetter(swspec)), false),
		parsers.NewSingleLineTagParser("Schemes", newSetSchemes(metaSchemeSetter(swspec))),
		parsers.NewMultiLineTagParser("Security", newSetSecurity(rxSecuritySchemes, metaSecuritySetter(swspec)), false),
		parsers.NewMultiLineTagParser("SecurityDefinitions", parsers.NewYAMLParser(rxSecurity, metaSecurityDefinitionsSetter(swspec)), true),
		parsers.NewSingleLineTagParser("Version", &setMetaSingle{swspec, rxVersion, setInfoVersion}),
		parsers.NewSingleLineTagParser("Host", &setMetaSingle{swspec, rxHost, setSwaggerHost}),
		parsers.NewSingleLineTagParser("BasePath", &setMetaSingle{swspec, rxBasePath, setSwaggerBasePath}),
		parsers.NewSingleLineTagParser("Contact", &setMetaSingle{swspec, rxContact, setInfoContact}),
		parsers.NewSingleLineTagParser("License", &setMetaSingle{swspec, rxLicense, setInfoLicense}),
		parsers.NewMultiLineTagParser("YAMLInfoExtensionsBlock", parsers.NewYAMLParser(rxInfoExtensions, infoVendorExtensibleSetter(swspec)), true),
		parsers.NewMultiLineTagParser("YAMLExtensionsBlock", parsers.NewYAMLParser(rxExtensions, metaVendorExtensibleSetter(swspec)), true),
	}
	return sp
}

type setMetaSingle struct {
	spec *oaispec.Swagger
	rx   *regexp.Regexp
	set  func(spec *oaispec.Swagger, lines []string) error
}

func (s *setMetaSingle) Matches(line string) bool {
	return s.rx.MatchString(line)
}

func (s *setMetaSingle) Parse(lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}
	matches := s.rx.FindStringSubmatch(lines[0])
	if len(matches) > 1 && len(matches[1]) > 0 {
		return s.set(s.spec, []string{matches[1]})
	}
	return nil
}

func setSwaggerHost(swspec *oaispec.Swagger, lines []string) error {
	lns := lines
	if len(lns) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		lns = []string{"localhost"}
	}
	swoaispec.Host = lns[0]
	return nil
}

func setSwaggerBasePath(swspec *oaispec.Swagger, lines []string) error {
	var ln string
	if len(lines) > 0 {
		ln = lines[0]
	}
	swoaispec.BasePath = ln
	return nil
}

func setInfoVersion(swspec *oaispec.Swagger, lines []string) error {
	if len(lines) == 0 {
		return nil
	}
	info := safeInfo(swspec)
	info.Version = strings.TrimSpace(lines[0])
	return nil
}

func setInfoContact(swspec *oaispec.Swagger, lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}
	contact, err := parseContactInfo(lines[0])
	if err != nil {
		return err
	}
	info := safeInfo(swspec)
	info.Contact = contact
	return nil
}

func parseContactInfo(line string) (*oaispec.ContactInfo, error) {
	nameEmail, url := splitURL(line)
	var name, email string
	if len(nameEmail) > 0 {
		addr, err := mail.ParseAddress(nameEmail)
		if err != nil {
			return nil, err
		}
		name, email = addr.Name, addr.Address
	}
	return &oaispec.ContactInfo{
		ContactInfoProps: oaispec.ContactInfoProps{
			URL:   url,
			Name:  name,
			Email: email,
		},
	}, nil
}

func setInfoLicense(swspec *oaispec.Swagger, lines []string) error {
	if len(lines) == 0 || (len(lines) == 1 && len(lines[0]) == 0) {
		return nil
	}
	info := safeInfo(swspec)
	line := lines[0]
	name, url := splitURL(line)
	info.License = &oaispec.License{
		LicenseProps: oaispec.LicenseProps{
			Name: name,
			URL:  url,
		},
	}
	return nil
}

func safeInfo(swspec *oaispec.Swagger) *oaispec.Info {
	if swoaispec.Info == nil {
		swoaispec.Info = new(oaispec.Info)
	}
	return swoaispec.Info
}

// httpFTPScheme matches http://, https://, ws://, wss://
var httpFTPScheme = regexp.MustCompile("(?:(?:ht|f)tp|ws)s?://")

func splitURL(line string) (notURL, url string) {
	str := strings.TrimSpace(line)
	parts := httpFTPScheme.FindStringIndex(str)
	if len(parts) == 0 {
		if len(str) > 0 {
			notURL = str
		}
		return
	}
	if len(parts) > 0 {
		notURL = strings.TrimSpace(str[:parts[0]])
		url = strings.TrimSpace(str[parts[0]:])
	}
	return
}
