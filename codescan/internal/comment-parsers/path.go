package parsers

import (
	"go/ast"
	"regexp"
	"strings"
)

type ParsedPathContent struct {
	Method, Path, ID string
	Tags             []string
	Remaining        *ast.CommentGroup
}

type PathParser struct{}

func NewPathParser() *PathParser {
	return &PathParser{}
}

func (p *PathParser) Route(lines *ast.CommentGroup) ParsedPathContent {
	if lines == nil {
		return ParsedPathContent{}
	}

	return p.parsePathAnnotation(rxRoute, lines.List)
}

func (p *PathParser) Operation(lines *ast.CommentGroup) ParsedPathContent {
	if lines == nil {
		return ParsedPathContent{}
	}

	return p.parsePathAnnotation(rxOperation, lines.List)
}

func (p *PathParser) parsePathAnnotation(annotation *regexp.Regexp, lines []*ast.Comment) (cnt ParsedPathContent) {
	var justMatched bool

	for _, cmt := range lines {
		txt := cmt.Text
		for _, line := range strings.Split(txt, "\n") {
			matches := annotation.FindStringSubmatch(line)
			if len(matches) > 3 {
				cnt.Method, cnt.Path, cnt.ID = matches[1], matches[2], matches[len(matches)-1]
				cnt.Tags = rxSpace.Split(matches[3], -1)
				if len(matches[3]) == 0 {
					cnt.Tags = nil
				}
				justMatched = true
			} else if cnt.Method != "" {
				if cnt.Remaining == nil {
					cnt.Remaining = new(ast.CommentGroup)
				}
				if !justMatched || strings.TrimSpace(rxStripComments.ReplaceAllString(line, "")) != "" {
					cc := new(ast.Comment)
					cc.Slash = cmt.Slash
					cc.Text = line
					cnt.Remaining.List = append(cnt.Remaining.List, cc)
					justMatched = false
				}
			}
		}
	}

	return
}
