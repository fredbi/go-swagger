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

package generator

import (
	"strings"
	"sync"
)

// resolvedSymbols builds the dictionary of enum symbols to be generated
var resolvedSymbols *genSymbolsDict

func init() {
	resolvedSymbols = newGenSymbolDict()
	resolvedSymbols.newDict = func(holder *GenEnum, referer *schemaGenContext) *dictEntry {
		// builds dictionary entries specifically for enums
		return &dictEntry{
			holder:  holder,
			referer: referer,
		}
	}
}

type dictEntry struct {
	holder  *GenEnum          // holder is the object holding the symbol
	referer *schemaGenContext // referer is the object responsible for creating this holder
}

// genSymbolsDict constructs a dictionary of generated const and vars definitions ("symbols")
// in the generated package "models".
// This ensures no duplicate symbols are generated.
//
// NOTE:  this structure could be extended to be more generic and handle more things than GenEnum pointers,
// like suggested by issue#1098.
//
type genSymbolsDict struct {
	symbolMutex *sync.Mutex
	// dictionary of generated enum value constants or variables
	symbolDict map[string]*dictEntry
	// utility to make dictEntries
	newDict func(holder *GenEnum, referer *schemaGenContext) *dictEntry
}

func newGenSymbolDict() *genSymbolsDict {
	return &genSymbolsDict{
		symbolMutex: &sync.Mutex{},
		symbolDict:  make(map[string]*dictEntry, 100),
	}
}

func (d *genSymbolsDict) lock() {
	d.symbolMutex.Lock()
}

func (d *genSymbolsDict) unlock() {
	d.symbolMutex.Unlock()
}

func (d *genSymbolsDict) reset() {
	d.lock()
	d.symbolDict = make(map[string]*dictEntry, 100)
	d.unlock()
}

func (d *genSymbolsDict) hasSymbol(name string) bool {
	d.lock()
	_, found := d.symbolDict[strings.ToLower(name)]
	d.unlock()
	return found
}

func (d *genSymbolsDict) addSymbol(p *dictEntry, name string) {
	d.lock()
	d.symbolDict[strings.ToLower(name)] = p
	d.unlock()
}

func (d *genSymbolsDict) remove(holder *GenEnum) {
	// removeHolder cleans a dictionary from the references previously created by holder
	d.lock()
	for k, v := range d.symbolDict {
		if v.holder == holder {
			delete(d.symbolDict, k)
		}
	}
	d.unlock()
}

func (d *genSymbolsDict) isOffender(p *dictEntry, name string) bool {
	// isOffender determines if the caller in a conflict is the offender or if the current symbol holder must revise its value
	d.lock()
	entry, found := d.symbolDict[strings.ToLower(name)]
	d.unlock()
	if !found || entry.holder == p.holder {
		// no conflict
		return false
	}
	return true
}

func checkEnumBeenThere(sc *GenSchema) {
	// checkEnumBeenThere removes allocated symbols when the caller rebuilds an element
	if sc.GenEnum != nil {
		resolvedSymbols.remove(sc.GenEnum)
	}
}
