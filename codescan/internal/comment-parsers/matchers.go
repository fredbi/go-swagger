package parsers

import (
	"fmt"
	"go/ast"
	"regexp"
	"slices"
	"strings"
)

// Matcher identifies meaningful annotations in comments,
// such as "swagger:xxx" tags.
type Matcher struct{}

// NewMatcher yields a fresh code annotation matcher.
func NewMatcher() *Matcher {
	return &Matcher{}
}

// JoinDropLast builds a single string out of a multi-lines comment,
// with any trailing blank line removed.
func (m *Matcher) JoinDropLast(lines []string) string {
	l := len(lines)
	if l == 0 {
		return ""
	}

	if len(strings.TrimSpace(lines[l-1])) == 0 {
		lines = lines[:l-1]
	}

	return strings.Join(lines, "\n")
}

func (m *Matcher) AllOfMember(comments *ast.CommentGroup) bool {
	return matchInComments(rxAllOf)(comments)
}

func (m *Matcher) FileParam(comments *ast.CommentGroup) bool {
	return matchInComments(rxFileUpload)(comments)
}

func (m *Matcher) Ignored(comments *ast.CommentGroup) bool {
	return matchInComments(rxIgnoreOverride)(comments)
}

func (m *Matcher) AliasParam(comments *ast.CommentGroup) bool {
	return matchInComments(rxAlias)(comments)
}

func (m *Matcher) DefaultName(comments *ast.CommentGroup) (string, bool) {
	return submatchInComments(rxDefault)(comments)
}

func (m *Matcher) TypeName(comments *ast.CommentGroup) (string, bool) {
	return submatchInComments(rxType)(comments)
}

func (m *Matcher) StrfmtName(comments *ast.CommentGroup) (string, bool) {
	return submatchInComments(rxStrFmt)(comments)
}

func (m *Matcher) Model(comments *ast.CommentGroup) (string, bool) {
	return submatchInComments(rxModelOverride)(comments)
}

func (m *Matcher) Response(comments *ast.CommentGroup) (string, bool) {
	return submatchInComments(rxResponseOverride)(comments)
}

func (m *Matcher) Parameters(comments *ast.CommentGroup) ([]string, bool) {
	return submatchesInComments(rxParametersOverride)(comments)
}

func (m *Matcher) EnumName(comments *ast.CommentGroup) (string, bool) {
	return submatchInComments(rxEnum)(comments)
}

func (m *Matcher) Annotation(line string) (string, bool) {
	return submatchInLine(rxSwaggerAnnotation)(line)
}

func submatchesInComments(rx *regexp.Regexp) func(*ast.CommentGroup) ([]string, bool) {
	return func(comments *ast.CommentGroup) ([]string, bool) {
		if comments == nil {
			return nil, false
		}

		for _, cmt := range comments.List {
			for ln := range strings.SplitSeq(cmt.Text, "\n") {
				matches := rx.FindStringSubmatch(ln)
				if len(matches) <= 1 || len(strings.TrimSpace(matches[1])) == 0 {
					continue
				}

				result := make([]string, 0)
				for pt := range strings.SplitSeq(matches[1], " ") {
					tr := strings.TrimSpace(pt)
					if len(tr) == 0 {
						continue
					}

					result = append(result, tr)
				}

				return result, true
			}
		}

		return nil, false
	}
}

func removeEmptyLines(lines []string) (notEmpty []string) {
	for _, l := range lines {
		if len(strings.TrimSpace(l)) == 0 {
			continue
		}

		notEmpty = append(notEmpty, l)
	}

	return
}

func rxf(rxp, ar string) *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf(rxp, ar))
}

// matchInComments matches a single line pattern in a group of comments.
func matchInComments(rx *regexp.Regexp) func(*ast.CommentGroup) bool {
	return func(comments *ast.CommentGroup) bool {
		if comments == nil {
			return false
		}

		return slices.ContainsFunc(comments.List, func(cmt *ast.Comment) bool {
			lines := strings.Split(cmt.Text, "\n")
			return slices.ContainsFunc(lines, func(ln string) bool {
				return rx.MatchString(ln)
			})
		})
	}
}

// submatchInComments matches a single line pattern with a submatch in a group of comments
func submatchInComments(rx *regexp.Regexp) func(*ast.CommentGroup) (string, bool) {
	return func(comments *ast.CommentGroup) (string, bool) {
		if comments == nil {
			return "", false
		}

		for _, cmt := range comments.List {
			for ln := range strings.SplitSeq(cmt.Text, "\n") {
				matches := rx.FindStringSubmatch(ln)
				if len(matches) > 1 && len(strings.TrimSpace(matches[1])) > 0 {
					return strings.TrimSpace(matches[1]), true
				}
			}
		}

		return "", false
	}
}

func submatchInLine(rx *regexp.Regexp) func(string) (string, bool) {
	return func(line string) (string, bool) {
		if line == "" {
			return "", false
		}

		for ln := range strings.SplitSeq(line, "\n") {
			matches := rx.FindStringSubmatch(ln)
			if len(matches) > 1 && len(strings.TrimSpace(matches[1])) > 0 {
				return strings.TrimSpace(matches[1]), true
			}
		}

		return "", false
	}
}
