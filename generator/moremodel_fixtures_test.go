package generator

func initFixture1042() {
	// testing ../fixtures/bugs/1487/fixture-1042.yaml with expand (--skip-flatten)

	/* when the specification incorrectly defines the allOf,
	generated unmarshalling is wrong.
	This fixture asserts that with correct spec, the generated models are correct.

	*/

	f := newModelFixture("../fixtures/bugs/1487/fixture-1042.yaml", "allOf marshalling")
	thisRun := f.AddRun(false)

	// load expectations for model: b.go
	thisRun.AddExpectations("b.go", []string{
		`type B struct {`,
		`	A`,
		`	BAllOf1`,
		`func (m *B) UnmarshalJSON(raw []byte) error {`,
		`	var aO0 A`,
		`	if err := swag.ReadJSON(raw, &aO0); err != nil {`,
		`	m.A = aO0`,
		`	var aO1 BAllOf1`,
		`	if err := swag.ReadJSON(raw, &aO1); err != nil {`,
		`	m.BAllOf1 = aO1`,
		`func (m B) MarshalJSON() ([]byte, error) {`,
		`	var _parts [][]byte`,
		`	aO0, err := swag.WriteJSON(m.A`,
		`	if err != nil {`,
		`		return nil, err`,
		`	_parts = append(_parts, aO0`,
		`	aO1, err := swag.WriteJSON(m.BAllOf1`,
		`	if err != nil {`,
		`		return nil, err`,
		`	_parts = append(_parts, aO1`,
		`	return swag.ConcatJSON(_parts...), nil`,
		`func (m *B) Validate(formats strfmt.Registry) error {`,
		`	if err := m.A.Validate(formats); err != nil {`,
		`	if err := m.BAllOf1.Validate(formats); err != nil {`,
		`		return errors.CompositeValidationError(res...`,
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

	// load expectations for model: b_all_of1.go
	thisRun.AddExpectations("b_all_of1.go", []string{
		`type BAllOf1 struct {`,
		"	F3 *string `json:\"f3\"`",
		"	F4 []string `json:\"f4\"`",
		`func (m *BAllOf1) Validate(formats strfmt.Registry) error {`,
		`	if err := m.validateF3(formats); err != nil {`,
		`	if err := m.validateF4(formats); err != nil {`,
		`		return errors.CompositeValidationError(res...`,
		`func (m *BAllOf1) validateF3(formats strfmt.Registry) error {`,
		`	if err := validate.Required("f3", "body", m.F3); err != nil {`,
		`func (m *BAllOf1) validateF4(formats strfmt.Registry) error {`,
		`	if err := validate.Required("f4", "body", m.F4); err != nil {`,
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

	// load expectations for model: a.go
	thisRun.AddExpectations("a.go", []string{
		`type A struct {`,
		"	F1 string `json:\"f1,omitempty\"`",
		"	F2 string `json:\"f2,omitempty\"`",
		// empty validation
		"func (m *A) Validate(formats strfmt.Registry) error {\n	return nil\n}",
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

}

func initFixture1042V2() {
	// testing ../fixtures/bugs/1487/fixture-1042-2.yaml with expand (--skip-flatten)

	/* when the specification incorrectly defines the allOf,
	generated unmarshalling is wrong.
	This fixture asserts that with correct spec, the generated models are correct.

	*/

	f := newModelFixture("../fixtures/bugs/1487/fixture-1042-2.yaml", "allOf marshalling")
	thisRun := f.AddRun(false)

	// load expectations for model: error_model.go
	thisRun.AddExpectations("error_model.go", []string{
		`type ErrorModel struct {`,
		"	Code *int64 `json:\"code\"`",
		"	Message *string `json:\"message\"`",
		`func (m *ErrorModel) Validate(formats strfmt.Registry) error {`,
		`	if err := m.validateCode(formats); err != nil {`,
		`	if err := m.validateMessage(formats); err != nil {`,
		`		return errors.CompositeValidationError(res...`,
		`func (m *ErrorModel) validateCode(formats strfmt.Registry) error {`,
		`	if err := validate.Required("code", "body", m.Code); err != nil {`,
		`	if err := validate.MinimumInt("code", "body", int64(*m.Code), 100, false); err != nil {`,
		`	if err := validate.MaximumInt("code", "body", int64(*m.Code), 600, false); err != nil {`,
		`func (m *ErrorModel) validateMessage(formats strfmt.Registry) error {`,
		`	if err := validate.Required("message", "body", m.Message); err != nil {`,
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

	// load expectations for model: extended_error_model.go
	thisRun.AddExpectations("extended_error_model.go", []string{
		`type ExtendedErrorModel struct {`,
		`	ErrorModel`,
		`	ExtendedErrorModelAllOf1`,
		`func (m *ExtendedErrorModel) UnmarshalJSON(raw []byte) error {`,
		`	var aO0 ErrorModel`,
		`	if err := swag.ReadJSON(raw, &aO0); err != nil {`,
		`	m.ErrorModel = aO0`,
		`	var aO1 ExtendedErrorModelAllOf1`,
		`	if err := swag.ReadJSON(raw, &aO1); err != nil {`,
		`	m.ExtendedErrorModelAllOf1 = aO1`,
		`func (m ExtendedErrorModel) MarshalJSON() ([]byte, error) {`,
		`	var _parts [][]byte`,
		`	aO0, err := swag.WriteJSON(m.ErrorModel`,
		`	if err != nil {`,
		`		return nil, err`,
		`	_parts = append(_parts, aO0`,
		`	aO1, err := swag.WriteJSON(m.ExtendedErrorModelAllOf1`,
		`	if err != nil {`,
		`		return nil, err`,
		`	_parts = append(_parts, aO1`,
		`	return swag.ConcatJSON(_parts...), nil`,
		`func (m *ExtendedErrorModel) Validate(formats strfmt.Registry) error {`,
		`	if err := m.ErrorModel.Validate(formats); err != nil {`,
		`	if err := m.ExtendedErrorModelAllOf1.Validate(formats); err != nil {`,
		`		return errors.CompositeValidationError(res...`,
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

	// load expectations for model: extended_error_model_all_of1.go
	thisRun.AddExpectations("extended_error_model_all_of1.go", []string{
		`type ExtendedErrorModelAllOf1 struct {`,
		"	RootCause *string `json:\"rootCause\"`",
		`func (m *ExtendedErrorModelAllOf1) Validate(formats strfmt.Registry) error {`,
		`	if err := m.validateRootCause(formats); err != nil {`,
		`		return errors.CompositeValidationError(res...`,
		`func (m *ExtendedErrorModelAllOf1) validateRootCause(formats strfmt.Registry) error {`,
		`	if err := validate.Required("rootCause", "body", m.RootCause); err != nil {`,
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

}

func initFixture979() {
	// testing ../fixtures/bugs/1487/fixture-979.yaml with expand (--skip-flatten)

	/* checking that properties is enough to figure out an object schema
	 */

	f := newModelFixture("../fixtures/bugs/1487/fixture-979.yaml", "allOf without the explicit type object")
	thisRun := f.AddRun(false)

	// load expectations for model: cluster.go
	thisRun.AddExpectations("cluster.go", []string{
		`type Cluster struct {`,
		`	NewCluster`,
		`	ClusterAllOf1`,
		`func (m *Cluster) UnmarshalJSON(raw []byte) error {`,
		`	var aO0 NewCluster`,
		`	if err := swag.ReadJSON(raw, &aO0); err != nil {`,
		`	m.NewCluster = aO0`,
		`	var aO1 ClusterAllOf1`,
		`	if err := swag.ReadJSON(raw, &aO1); err != nil {`,
		`	m.ClusterAllOf1 = aO1`,
		`func (m Cluster) MarshalJSON() ([]byte, error) {`,
		`	var _parts [][]byte`,
		`	aO0, err := swag.WriteJSON(m.NewCluster`,
		`	if err != nil {`,
		`		return nil, err`,
		`	_parts = append(_parts, aO0`,
		`	aO1, err := swag.WriteJSON(m.ClusterAllOf1`,
		`	if err != nil {`,
		`		return nil, err`,
		`	_parts = append(_parts, aO1`,
		`	return swag.ConcatJSON(_parts...), nil`,
		`func (m *Cluster) Validate(formats strfmt.Registry) error {`,
		`	if err := m.NewCluster.Validate(formats); err != nil {`,
		`	if err := m.ClusterAllOf1.Validate(formats); err != nil {`,
		`		return errors.CompositeValidationError(res...`,
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

	// load expectations for model: new_cluster.go
	thisRun.AddExpectations("new_cluster.go", []string{
		`type NewCluster struct {`,
		"	DummyProp1 int64 `json:\"dummyProp1,omitempty\"`",
		// empty validation
		"func (m *NewCluster) Validate(formats strfmt.Registry) error {\n	return nil\n}",
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

	// load expectations for model: cluster_all_of1.go
	thisRun.AddExpectations("cluster_all_of1.go", []string{
		`type ClusterAllOf1 struct {`,
		"	Result string `json:\"result,omitempty\"`",
		"	Status string `json:\"status,omitempty\"`",
		// empty validation
		"func (m *ClusterAllOf1) Validate(formats strfmt.Registry) error {\n	return nil\n}",
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

}

func initFixture842() {
	// testing ../fixtures/bugs/1487/fixture-842.yaml with expand (--skip-flatten)

	/* codegen fails to produce code that builds
	 */

	f := newModelFixture("../fixtures/bugs/1487/fixture-842.yaml", "polymorphic type containing an array of the base type")
	thisRun := f.AddRun(false)

	// load expectations for model: value_array_all_of1.go
	thisRun.AddExpectations("value_array_all_of1.go", []string{
		`type ValueArrayAllOf1 struct {`,
		`	valuesField []Value`,
		`func (m *ValueArrayAllOf1) Values() []Value {`,
		`	return m.valuesField`,
		`func (m *ValueArrayAllOf1) SetValues(val []Value) {`,
		`	m.valuesField = val`,
		`func (m *ValueArrayAllOf1) Validate(formats strfmt.Registry) error {`,
		`	if err := m.validateValues(formats); err != nil {`,
		`		return errors.CompositeValidationError(res...`,
		`func (m *ValueArrayAllOf1) validateValues(formats strfmt.Registry) error {`,
		`	if err := validate.Required("Values", "body", m.Values()); err != nil {`,
		`	for i := 0; i < len(m.Values()); i++ {`,
		`		if err := m.valuesField[i].Validate(formats); err != nil {`,
		`			if ve, ok := err.(*errors.Validation); ok {`,
		`				return ve.ValidateName("Values" + "." + strconv.Itoa(i)`,
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

	// load expectations for model: value_array.go
	thisRun.AddExpectations("value_array.go", []string{
		`type ValueArray struct {`,
		`	ValueArrayAllOf1`,
		`func (m *ValueArray) ValueType() string {`,
		`	return "ValueArray"`,
		`func (m *ValueArray) SetValueType(val string) {`,
		`func (m *ValueArray) Validate(formats strfmt.Registry) error {`,
		`	if err := m.ValueArrayAllOf1.Validate(formats); err != nil {`,
		`		return errors.CompositeValidationError(res...`,
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

	// load expectations for model: value.go
	thisRun.AddExpectations("value.go", []string{
		`type Value interface {`,
		`	runtime.Validatable`,
		`	ValueType() string`,
		`	SetValueType(string`,
		`type value struct {`,
		`	valueTypeField string`,
		`func (m *value) ValueType() string {`,
		`	return "Value"`,
		`func (m *value) SetValueType(val string) {`,
		`func (m *value) Validate(formats strfmt.Registry) error {`,
		`		return errors.CompositeValidationError(res...`,
		`func UnmarshalValueSlice(reader io.Reader, consumer runtime.Consumer) ([]Value, error) {`,
		`	var elements []json.RawMessage`,
		`	if err := consumer.Consume(reader, &elements); err != nil {`,
		`		return nil, err`,
		`	var result []Value`,
		`	for _, element := range elements {`,
		`		obj, err := unmarshalValue(element, consumer`,
		`		if err != nil {`,
		`			return nil, err`,
		`		result = append(result, obj`,
		`	return result, nil`,
		`func UnmarshalValue(reader io.Reader, consumer runtime.Consumer) (Value, error) {`,
		`	data, err := ioutil.ReadAll(reader`,
		`	if err != nil {`,
		`		return nil, err`,
		`	return unmarshalValue(data, consumer`,
		`func unmarshalValue(data []byte, consumer runtime.Consumer) (Value, error) {`,
		`	buf := bytes.NewBuffer(data`,
		`	buf2 := bytes.NewBuffer(data`,
		`	var getType struct {`,
		"		ValueType string `json:\"ValueType\"`",
		`	if err := consumer.Consume(buf, &getType); err != nil {`,
		`		return nil, err`,
		`	if err := validate.RequiredString("ValueType", "body", getType.ValueType); err != nil {`,
		`		return nil, err`,
		`	switch getType.ValueType {`,
		`	case "Value":`,
		`		var result value`,
		`		if err := consumer.Consume(buf2, &result); err != nil {`,
		`			return nil, err`,
		`		return &result, nil`,
		`	case "ValueArray":`,
		`		var result ValueArray`,
		`		if err := consumer.Consume(buf2, &result); err != nil {`,
		`			return nil, err`,
		`		return &result, nil`,
		`	return nil, errors.New(422, "invalid ValueType value: %q", getType.ValueType`,
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

}

func initFixture607() {
	// testing ../fixtures/bugs/1487/fixture-607.yaml with expand (--skip-flatten)

	/* broken code produced on polymorphic type
	 */

	f := newModelFixture("../fixtures/bugs/1487/fixture-607.yaml", "broken code when using array of polymorphic type")
	thisRun := f.AddRun(false)

	// load expectations for model: range_filter_all_of1.go
	thisRun.AddExpectations("range_filter_all_of1.go", []string{
		`type RangeFilterAllOf1 struct {`,
		"	Config *RangeFilterAllOf1Config `json:\"config\"`",
		`func (m *RangeFilterAllOf1) Validate(formats strfmt.Registry) error {`,
		`	if err := m.validateConfig(formats); err != nil {`,
		`		return errors.CompositeValidationError(res...`,
		`func (m *RangeFilterAllOf1) validateConfig(formats strfmt.Registry) error {`,
		`	if err := validate.Required("config", "body", m.Config); err != nil {`,
		`	if m.Config != nil {`,
		`		if err := m.Config.Validate(formats); err != nil {`,
		`			if ve, ok := err.(*errors.Validation); ok {`,
		`				return ve.ValidateName("config"`,
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

	// load expectations for model: filter.go
	thisRun.AddExpectations("filter.go", []string{
		`type Filter interface {`,
		`	runtime.Validatable`,
		`	Type() string`,
		`	SetType(string`,
		`type filter struct {`,
		`	typeField string`,
		`func (m *filter) Type() string {`,
		`	return "Filter"`,
		`func (m *filter) SetType(val string) {`,
		`func (m *filter) Validate(formats strfmt.Registry) error {`,
		`		return errors.CompositeValidationError(res...`,
		`func UnmarshalFilterSlice(reader io.Reader, consumer runtime.Consumer) ([]Filter, error) {`,
		`	var elements []json.RawMessage`,
		`	if err := consumer.Consume(reader, &elements); err != nil {`,
		`		return nil, err`,
		`	var result []Filter`,
		`	for _, element := range elements {`,
		`		obj, err := unmarshalFilter(element, consumer`,
		`		if err != nil {`,
		`			return nil, err`,
		`		result = append(result, obj`,
		`	return result, nil`,
		`func UnmarshalFilter(reader io.Reader, consumer runtime.Consumer) (Filter, error) {`,
		`	data, err := ioutil.ReadAll(reader`,
		`	if err != nil {`,
		`		return nil, err`,
		`	return unmarshalFilter(data, consumer`,
		`func unmarshalFilter(data []byte, consumer runtime.Consumer) (Filter, error) {`,
		`	buf := bytes.NewBuffer(data`,
		`	buf2 := bytes.NewBuffer(data`,
		`	var getType struct {`,
		"		Type string `json:\"type\"`",
		`	if err := consumer.Consume(buf, &getType); err != nil {`,
		`		return nil, err`,
		`	if err := validate.RequiredString("type", "body", getType.Type); err != nil {`,
		`		return nil, err`,
		`	switch getType.Type {`,
		`	case "AndFilter":`,
		`		var result AndFilter`,
		`		if err := consumer.Consume(buf2, &result); err != nil {`,
		`			return nil, err`,
		`		return &result, nil`,
		`	case "Filter":`,
		`		var result filter`,
		`		if err := consumer.Consume(buf2, &result); err != nil {`,
		`			return nil, err`,
		`		return &result, nil`,
		`	case "RangeFilter":`,
		`		var result RangeFilter`,
		`		if err := consumer.Consume(buf2, &result); err != nil {`,
		`			return nil, err`,
		`		return &result, nil`,
		`	return nil, errors.New(422, "invalid type value: %q", getType.Type`,
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

	// load expectations for model: and_filter_all_of1.go
	thisRun.AddExpectations("and_filter_all_of1.go", []string{
		`type AndFilterAllOf1 struct {`,
		`	configField []Filter`,
		`func (m *AndFilterAllOf1) Config() []Filter {`,
		`	return m.configField`,
		`func (m *AndFilterAllOf1) SetConfig(val []Filter) {`,
		`	m.configField = val`,
		`func (m *AndFilterAllOf1) UnmarshalJSON(raw []byte) error {`,
		`	var data struct {`,
		"		Config json.RawMessage `json:\"config\"`",
		`	buf := bytes.NewBuffer(raw`,
		`	dec := json.NewDecoder(buf`,
		`	dec.UseNumber(`,
		`	if err := dec.Decode(&data); err != nil {`,
		`	config, err := UnmarshalFilterSlice(bytes.NewBuffer(data.Config), runtime.JSONConsumer()`,
		`	if err != nil && err != io.EOF {`,
		`	var result AndFilterAllOf1`,
		`	result.configField = config`,
		`	*m = result`,
		`func (m AndFilterAllOf1) MarshalJSON() ([]byte, error) {`,
		`	var b1, b2, b3 []byte`,
		`	var err error`,
		`	b1, err = json.Marshal(struct {`,
		`	}{},`,
		`	if err != nil {`,
		`		return nil, err`,
		`	b2, err = json.Marshal(struct {`,
		"		Config []Filter `json:\"config\"`",
		`	}{`,
		`		Config: m.configField,`,
		`	},`,
		`	if err != nil {`,
		`		return nil, err`,
		`	return swag.ConcatJSON(b1, b2, b3), nil`,
		`func (m *AndFilterAllOf1) Validate(formats strfmt.Registry) error {`,
		`	if err := m.validateConfig(formats); err != nil {`,
		`		return errors.CompositeValidationError(res...`,
		`func (m *AndFilterAllOf1) validateConfig(formats strfmt.Registry) error {`,
		`	if err := validate.Required("config", "body", m.Config()); err != nil {`,
		`	for i := 0; i < len(m.Config()); i++ {`,
		`		if err := m.configField[i].Validate(formats); err != nil {`,
		`			if ve, ok := err.(*errors.Validation); ok {`,
		`				return ve.ValidateName("config" + "." + strconv.Itoa(i)`,
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

	// load expectations for model: and_filter.go
	thisRun.AddExpectations("and_filter.go", []string{
		`type AndFilter struct {`,
		`	AndFilterAllOf1`,
		`func (m *AndFilter) Type() string {`,
		`	return "AndFilter"`,
		`func (m *AndFilter) SetType(val string) {`,
		`func (m *AndFilter) UnmarshalJSON(raw []byte) error {`,
		`	var data struct {`,
		`		AndFilterAllOf1`,
		`	buf := bytes.NewBuffer(raw`,
		`	dec := json.NewDecoder(buf`,
		`	dec.UseNumber(`,
		`	if err := dec.Decode(&data); err != nil {`,
		`	var base struct {`,
		"		Type string `json:\"type\"`",
		`	buf = bytes.NewBuffer(raw`,
		`	dec = json.NewDecoder(buf`,
		`	dec.UseNumber(`,
		`	if err := dec.Decode(&base); err != nil {`,
		`	var result AndFilter`,
		`	if base.Type != result.Type() {`,
		`		return errors.New(422, "invalid type value: %q", base.Type`,
		`	result.AndFilterAllOf1 = data.AndFilterAllOf1`,
		`	*m = result`,
		`func (m AndFilter) MarshalJSON() ([]byte, error) {`,
		`	var b1, b2, b3 []byte`,
		`	var err error`,
		`	b1, err = json.Marshal(struct {`,
		`		AndFilterAllOf1`,
		`	}{`,
		`		AndFilterAllOf1: m.AndFilterAllOf1,`,
		`	},`,
		`	if err != nil {`,
		`		return nil, err`,
		`	b2, err = json.Marshal(struct {`,
		"		Type string `json:\"type\"`",
		`	}{`,
		`		Type: m.Type(),`,
		`	},`,
		`	if err != nil {`,
		`		return nil, err`,
		`	return swag.ConcatJSON(b1, b2, b3), nil`,
		`func (m *AndFilter) Validate(formats strfmt.Registry) error {`,
		`	if err := m.AndFilterAllOf1.Validate(formats); err != nil {`,
		`		return errors.CompositeValidationError(res...`,
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

	// load expectations for model: range_filter.go
	thisRun.AddExpectations("range_filter.go", []string{
		`type RangeFilter struct {`,
		`	RangeFilterAllOf1`,
		`func (m *RangeFilter) Type() string {`,
		`	return "RangeFilter"`,
		`func (m *RangeFilter) SetType(val string) {`,
		`func (m *RangeFilter) UnmarshalJSON(raw []byte) error {`,
		`	var data struct {`,
		`		RangeFilterAllOf1`,
		`	buf := bytes.NewBuffer(raw`,
		`	dec := json.NewDecoder(buf`,
		`	dec.UseNumber(`,
		`	if err := dec.Decode(&data); err != nil {`,
		`	var base struct {`,
		"		Type string `json:\"type\"`",
		`	buf = bytes.NewBuffer(raw`,
		`	dec = json.NewDecoder(buf`,
		`	dec.UseNumber(`,
		`	if err := dec.Decode(&base); err != nil {`,
		`	var result RangeFilter`,
		`	if base.Type != result.Type() {`,
		`		return errors.New(422, "invalid type value: %q", base.Type`,
		`	result.RangeFilterAllOf1 = data.RangeFilterAllOf1`,
		`	*m = result`,
		`func (m RangeFilter) MarshalJSON() ([]byte, error) {`,
		`	var b1, b2, b3 []byte`,
		`	var err error`,
		`	b1, err = json.Marshal(struct {`,
		`		RangeFilterAllOf1`,
		`	}{`,
		`		RangeFilterAllOf1: m.RangeFilterAllOf1,`,
		`	},`,
		`	if err != nil {`,
		`		return nil, err`,
		`	b2, err = json.Marshal(struct {`,
		"		Type string `json:\"type\"`",
		`	}{`,
		`		Type: m.Type(),`,
		`	},`,
		`	if err != nil {`,
		`		return nil, err`,
		`	return swag.ConcatJSON(b1, b2, b3), nil`,
		`func (m *RangeFilter) Validate(formats strfmt.Registry) error {`,
		`	if err := m.RangeFilterAllOf1.Validate(formats); err != nil {`,
		`		return errors.CompositeValidationError(res...`,
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

	// load expectations for model: range_filter_all_of1_config.go
	thisRun.AddExpectations("range_filter_all_of1_config.go", []string{
		`type RangeFilterAllOf1Config struct {`,
		"	Gt float64 `json:\"gt,omitempty\"`",
		"	Gte float64 `json:\"gte,omitempty\"`",
		"	Lt float64 `json:\"lt,omitempty\"`",
		"	Lte float64 `json:\"lte,omitempty\"`",
		// empty validation
		"func (m *RangeFilterAllOf1Config) Validate(formats strfmt.Registry) error {\n	return nil\n}",
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

}

func initFixture1336() {
	// testing ../fixtures/bugs/1487/fixture-1336.yaml with expand (--skip-flatten)

	/* broken code produced on polymorphic type
	 */

	f := newModelFixture("../fixtures/bugs/1487/fixture-1336.yaml", "broken code when using array of polymorphic type")
	thisRun := f.AddRun(false)

	// load expectations for model: node.go
	thisRun.AddExpectations("node.go", []string{
		`type Node interface {`,
		`	runtime.Validatable`,
		`	NodeType() string`,
		`	SetNodeType(string`,
		`type node struct {`,
		`	nodeTypeField string`,
		`func (m *node) NodeType() string {`,
		`	return "Node"`,
		`func (m *node) SetNodeType(val string) {`,
		`func (m *node) Validate(formats strfmt.Registry) error {`,
		`		return errors.CompositeValidationError(res...`,
		`func UnmarshalNodeSlice(reader io.Reader, consumer runtime.Consumer) ([]Node, error) {`,
		`	var elements []json.RawMessage`,
		`	if err := consumer.Consume(reader, &elements); err != nil {`,
		`		return nil, err`,
		`	var result []Node`,
		`	for _, element := range elements {`,
		`		obj, err := unmarshalNode(element, consumer`,
		`		if err != nil {`,
		`			return nil, err`,
		`		result = append(result, obj`,
		`	return result, nil`,
		`func UnmarshalNode(reader io.Reader, consumer runtime.Consumer) (Node, error) {`,
		`	data, err := ioutil.ReadAll(reader`,
		`	if err != nil {`,
		`		return nil, err`,
		`	return unmarshalNode(data, consumer`,
		`func unmarshalNode(data []byte, consumer runtime.Consumer) (Node, error) {`,
		`	buf := bytes.NewBuffer(data`,
		`	buf2 := bytes.NewBuffer(data`,
		`	var getType struct {`,
		"		NodeType string `json:\"NodeType\"`",
		`	if err := consumer.Consume(buf, &getType); err != nil {`,
		`		return nil, err`,
		`	if err := validate.RequiredString("NodeType", "body", getType.NodeType); err != nil {`,
		`		return nil, err`,
		`	switch getType.NodeType {`,
		`	case "CodeBlockNode":`,
		`		var result CodeBlockNode`,
		`		if err := consumer.Consume(buf2, &result); err != nil {`,
		`			return nil, err`,
		`		return &result, nil`,
		`	case "DocBlockNode":`,
		`		var result DocBlockNode`,
		`		if err := consumer.Consume(buf2, &result); err != nil {`,
		`			return nil, err`,
		`		return &result, nil`,
		`	case "Node":`,
		`		var result node`,
		`		if err := consumer.Consume(buf2, &result); err != nil {`,
		`			return nil, err`,
		`		return &result, nil`,
		`	return nil, errors.New(422, "invalid NodeType value: %q", getType.NodeType`,
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

	// load expectations for model: code_block_node_all_of1.go
	thisRun.AddExpectations("code_block_node_all_of1.go", []string{
		`type CodeBlockNodeAllOf1 struct {`,
		"	Code string `json:\"Code,omitempty\"`",
		// empty validation
		"func (m *CodeBlockNodeAllOf1) Validate(formats strfmt.Registry) error {\n	return nil\n}",
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

	// load expectations for model: graph.go
	thisRun.AddExpectations("graph.go", []string{
		`type Graph struct {`,
		`	nodesField []Node`,
		`func (m *Graph) Nodes() []Node {`,
		`	return m.nodesField`,
		`func (m *Graph) SetNodes(val []Node) {`,
		`	m.nodesField = val`,
		`func (m *Graph) UnmarshalJSON(raw []byte) error {`,
		`	var data struct {`,
		"		Nodes json.RawMessage `json:\"Nodes,omitempty\"`",
		`	buf := bytes.NewBuffer(raw`,
		`	dec := json.NewDecoder(buf`,
		`	dec.UseNumber(`,
		`	if err := dec.Decode(&data); err != nil {`,
		`	nodes, err := UnmarshalNodeSlice(bytes.NewBuffer(data.Nodes), runtime.JSONConsumer()`,
		`	if err != nil && err != io.EOF {`,
		`	var result Graph`,
		`	result.nodesField = nodes`,
		`	*m = result`,
		`func (m Graph) MarshalJSON() ([]byte, error) {`,
		`	var b1, b2, b3 []byte`,
		`	var err error`,
		`	b1, err = json.Marshal(struct {`,
		`	}{},`,
		`	if err != nil {`,
		`		return nil, err`,
		`	b2, err = json.Marshal(struct {`,
		"		Nodes []Node `json:\"Nodes,omitempty\"`",
		`	}{`,
		`		Nodes: m.nodesField,`,
		`	},`,
		`	if err != nil {`,
		`		return nil, err`,
		`	return swag.ConcatJSON(b1, b2, b3), nil`,
		`func (m *Graph) Validate(formats strfmt.Registry) error {`,
		`		return errors.CompositeValidationError(res...`,
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

	// load expectations for model: doc_block_node_all_of1.go
	thisRun.AddExpectations("doc_block_node_all_of1.go", []string{
		`type DocBlockNodeAllOf1 struct {`,
		"	Doc string `json:\"Doc,omitempty\"`",
		// empty validation
		"func (m *DocBlockNodeAllOf1) Validate(formats strfmt.Registry) error {\n	return nil\n}",
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

	// load expectations for model: doc_block_node.go
	thisRun.AddExpectations("doc_block_node.go", []string{
		`type DocBlockNode struct {`,
		`	DocBlockNodeAllOf1`,
		`func (m *DocBlockNode) NodeType() string {`,
		`	return "DocBlockNode"`,
		`func (m *DocBlockNode) SetNodeType(val string) {`,
		`func (m *DocBlockNode) UnmarshalJSON(raw []byte) error {`,
		`	var data struct {`,
		`		DocBlockNodeAllOf1`,
		`	buf := bytes.NewBuffer(raw`,
		`	dec := json.NewDecoder(buf`,
		`	dec.UseNumber(`,
		`	if err := dec.Decode(&data); err != nil {`,
		`	var base struct {`,
		"		NodeType string `json:\"NodeType\"`",
		`	buf = bytes.NewBuffer(raw`,
		`	dec = json.NewDecoder(buf`,
		`	dec.UseNumber(`,
		`	if err := dec.Decode(&base); err != nil {`,
		`	var result DocBlockNode`,
		`	if base.NodeType != result.NodeType() {`,
		`		return errors.New(422, "invalid NodeType value: %q", base.NodeType`,
		`	result.DocBlockNodeAllOf1 = data.DocBlockNodeAllOf1`,
		`	*m = result`,
		`func (m DocBlockNode) MarshalJSON() ([]byte, error) {`,
		`	var b1, b2, b3 []byte`,
		`	var err error`,
		`	b1, err = json.Marshal(struct {`,
		`		DocBlockNodeAllOf1`,
		`	}{`,
		`		DocBlockNodeAllOf1: m.DocBlockNodeAllOf1,`,
		`	},`,
		`	if err != nil {`,
		`		return nil, err`,
		`	b2, err = json.Marshal(struct {`,
		"		NodeType string `json:\"NodeType\"`",
		`	}{`,
		`		NodeType: m.NodeType(),`,
		`	},`,
		`	if err != nil {`,
		`		return nil, err`,
		`	return swag.ConcatJSON(b1, b2, b3), nil`,
		`func (m *DocBlockNode) Validate(formats strfmt.Registry) error {`,
		`	if err := m.DocBlockNodeAllOf1.Validate(formats); err != nil {`,
		`		return errors.CompositeValidationError(res...`,
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

	// load expectations for model: code_block_node.go
	thisRun.AddExpectations("code_block_node.go", []string{
		`type CodeBlockNode struct {`,
		`	CodeBlockNodeAllOf1`,
		`func (m *CodeBlockNode) NodeType() string {`,
		`	return "CodeBlockNode"`,
		`func (m *CodeBlockNode) SetNodeType(val string) {`,
		`func (m *CodeBlockNode) UnmarshalJSON(raw []byte) error {`,
		`	var data struct {`,
		`		CodeBlockNodeAllOf1`,
		`	buf := bytes.NewBuffer(raw`,
		`	dec := json.NewDecoder(buf`,
		`	dec.UseNumber(`,
		`	if err := dec.Decode(&data); err != nil {`,
		`	var base struct {`,
		"		NodeType string `json:\"NodeType\"`",
		`	buf = bytes.NewBuffer(raw`,
		`	dec = json.NewDecoder(buf`,
		`	dec.UseNumber(`,
		`	if err := dec.Decode(&base); err != nil {`,
		`	var result CodeBlockNode`,
		`	if base.NodeType != result.NodeType() {`,
		`		return errors.New(422, "invalid NodeType value: %q", base.NodeType`,
		`	result.CodeBlockNodeAllOf1 = data.CodeBlockNodeAllOf1`,
		`	*m = result`,
		`func (m CodeBlockNode) MarshalJSON() ([]byte, error) {`,
		`	var b1, b2, b3 []byte`,
		`	var err error`,
		`	b1, err = json.Marshal(struct {`,
		`		CodeBlockNodeAllOf1`,
		`	}{`,
		`		CodeBlockNodeAllOf1: m.CodeBlockNodeAllOf1,`,
		`	},`,
		`	if err != nil {`,
		`		return nil, err`,
		`	b2, err = json.Marshal(struct {`,
		"		NodeType string `json:\"NodeType\"`",
		`	}{`,
		`		NodeType: m.NodeType(),`,
		`	},`,
		`	if err != nil {`,
		`		return nil, err`,
		`	return swag.ConcatJSON(b1, b2, b3), nil`,
		`func (m *CodeBlockNode) Validate(formats strfmt.Registry) error {`,
		`	if err := m.CodeBlockNodeAllOf1.Validate(formats); err != nil {`,
		`		return errors.CompositeValidationError(res...`,
	},
		// not expected
		todo,
		// output in log
		noLines,
		noLines)

}
