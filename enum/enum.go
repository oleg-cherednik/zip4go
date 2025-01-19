package enum

// ---------- basic ----------

type Enum struct {
	name string
}

func NewEnum(name string) *Enum {
	return &Enum{name}
}

func (t *Enum) GetName() string {
	return t.name
}

func (t *Enum) String() string {
	return t.GetName()
}

// ---------- code ----------

type CodeEnum struct {
	*Enum
	code int
}

func NewCodeEnum(name string, code int) *CodeEnum {
	p := NewEnum(name)
	return &CodeEnum{p, code}
}

func (t *CodeEnum) GetCode() int {
	return t.code
}

// ---------- title ----------

type TitleEnum struct {
	*Enum
	title string
}

func NewTitleEnum(name string, title string) *TitleEnum {
	p := NewEnum(name)
	return &TitleEnum{p, title}
}

func (t *TitleEnum) GetTitle() string {
	return t.title
}

// ---------- code & title ----------

type CodeTitleEnum struct {
	*Enum
	code  int
	title string
}

func NewCodeTitleEnum(name string, code int, title string) *CodeTitleEnum {
	p := NewEnum(name)
	return &CodeTitleEnum{p, code, title}
}

func (t *CodeTitleEnum) GetCode() int {
	return t.code
}
