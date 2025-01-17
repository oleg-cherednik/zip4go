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
	*TitleEnum
	code uint
}

func NewCodeTitleEnum(name string, code uint, title string) *CodeTitleEnum {
	p := NewTitleEnum(name, title)
	return &CodeTitleEnum{p, code}
}

func (t *CodeTitleEnum) GetCode() uint {
	return t.code
}
