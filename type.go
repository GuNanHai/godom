package godom

// Selector : CSS选择器类
// Type的值： "ID","CLASS","ELEMENT"
type Selector struct {
	Value     string
	ExtraInfo string
	Type      string
}

// Element : 存储HTML的元素
type Element struct {
	Raw   string
	Attrs []Attr
	Text  string
}

// Attr : 存储html元素的属性
type Attr struct {
	Name  string
	Value string
}

// ElementHalfLoc :	存储HTML元素的openning Tag(<)或者closing Tag(</)的index,用Sign值0表示openning Tag,1表示closing Tag
type ElementHalfLoc struct {
	Loc  []int
	Sign int
}

const (
	// ID :
	ID = "ID"
	// CLASS :
	CLASS = "CLASS"
	// ELEMENT :
	ELEMENT = "ELEMENT"
)
