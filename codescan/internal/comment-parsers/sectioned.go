package parsers

import (
	"go/ast"
	"strings"
)

// aggregates lines in header until it sees a tag.
type SectionedParser struct {
	header     []string
	matched    map[string]TagParser
	annotation valueParser

	seenTag        bool
	skipHeader     bool
	setTitle       func([]string)
	setDescription func([]string)
	workedOutTitle bool
	taggers        []TagParser
	currentTagger  *TagParser
	title          []string
	ignored        bool
}

func (st *SectionedParser) Title() []string {
	st.collectTitleDescription()

	return st.title
}

func (st *SectionedParser) Description() []string {
	st.collectTitleDescription()

	return st.header
}

func (st *SectionedParser) Parse(doc *ast.CommentGroup) error {
	if doc == nil {
		return nil
	}

COMMENTS:
	for _, c := range doc.List {
		for line := range strings.SplitSeq(c.Text, "\n") {
			if rxSwaggerAnnotation.MatchString(line) {
				if rxIgnoreOverride.MatchString(line) {
					st.ignored = true
					break COMMENTS // an explicit ignore terminates this parser
				}

				if st.annotation == nil || !st.annotation.Matches(line) {
					break COMMENTS // a new swagger: annotation terminates this parser
				}

				_ = st.annotation.Parse([]string{line})
				if len(st.header) > 0 {
					st.seenTag = true
				}

				continue
			}

			var matched bool
			for _, tg := range st.taggers {
				tagger := tg
				if tagger.Matches(line) {
					st.seenTag = true
					st.currentTagger = &tagger
					matched = true
					break
				}
			}

			if st.currentTagger == nil {
				if st.skipHeader || st.seenTag {
					// didn't match a tag, moving on
					continue
				}

				st.header = append(st.header, line)
				continue
			}

			if st.currentTagger.MultiLine && matched {
				// the first line of a multiline tagger doesn't count
				continue
			}

			ts, ok := st.matched[st.currentTagger.Name]
			if !ok {
				ts = *st.currentTagger
			}
			ts.Lines = append(ts.Lines, line)
			if st.matched == nil {
				st.matched = make(map[string]TagParser)
			}
			st.matched[st.currentTagger.Name] = ts

			if !st.currentTagger.MultiLine {
				st.currentTagger = nil
			}
		}
	}

	if st.setTitle != nil {
		st.setTitle(st.Title())
	}
	if st.setDescription != nil {
		st.setDescription(st.Description())
	}

	for _, mt := range st.matched {
		if !mt.SkipCleanUp {
			mt.Lines = cleanupScannerLines(mt.Lines, rxUncommentHeaders, nil)
		}
		if err := mt.Parse(mt.Lines); err != nil {
			return err
		}
	}

	return nil
}

func (st *SectionedParser) collectTitleDescription() {
	if st.workedOutTitle {
		return
	}
	if st.setTitle == nil {
		st.header = cleanupScannerLines(st.header, rxUncommentHeaders, nil)
		return
	}

	st.workedOutTitle = true
	st.title, st.header = collectScannerTitleDescription(st.header)
}
