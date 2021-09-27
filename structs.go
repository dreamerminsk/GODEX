package godex

type uleb128 uint64

type DexFile struct {
	Header        *HeaderItem
	StringIds     *[]StringIdItem
	TypeIds       *[]TypeIdItem
	ProtoIds      *[]ProtoIdItem
	FieldIds      *[]FieldIdItem
	MethodIds     *[]MethodIdItem
	ClassDefs     *[]ClassDefItem
	CallSiteIds   *[]CallSiteIdItem
	MethodHandles *[]MethodHandleItem
	Data          *[]uint8
	LinkData      *[]uint8
}

type HeaderItem struct {
	Magic         [8]uint8  //magic value
	Checksum      uint32    //adler32 checksum of the rest of the file (everything but Magic and this field)
	Signature     [20]uint8 //SHA-1 signature of the rest of the file (everything but Magic, Checksum, and this field)
	FileSize      uint32    //size of the entire file (including the header), in bytes
	HeaderSize    uint32    //size of the header (this entire section), in bytes = 0x70
	EndianTag     uint32    //endianness tag
	LinkSize      uint32    //size of the link section, or 0 if this file isn't statically linked
	LinkOff       uint32    //offset from the start of the file to the link section, or 0 if LinkSize == 0
	MapOff        uint32    //offset from the start of the file to the map item
	StringIdsSize uint32    //count of strings in the string ids list
	StringIdsOff  uint32    //offset from the start of the file to the string ids list, or 0 if StringIdsSize == 0
	TypeIdsSize   uint32    //count of elements in the type ids list, at most 65535
	TypeIdsOff    uint32    //offset from the start of the file to the type ids list, or 0 if TypeIdsSize == 0
	ProtoIdsSize  uint32    //count of elements in the prototype ids list, at most 65535
	ProtoIdsOff   uint32    //offset from the start of the file to the prototype ids list, or 0 if ProtoIdsSize == 0
	FieldIdsSize  uint32    //count of elements in the field ids list
	FieldIdsOff   uint32    //offset from the start of the file to the field ids list, or 0 if FieldIdsSize == 0
	MethodIdsSize uint32    //count of elements in the method ids list
	MethodIdsOff  uint32    //offset from the start of the file to the method identifiers list, or 0 if MethodIdsSize == 0
	ClassDefsSize uint32    //count of elements in the class defs list
	ClassDefsOff  uint32    //offset from the start of the file to the class defs list, or 0 if ClassDefsSize == 0
	DataSize      uint32    //Size of data section in bytes. Must be an even multiple of sizeof(uint).
	DataOff       uint32    //offset from the start of the file to the start of the data section.
}

type StringIdItem struct {
	StringDataOff uint32 //offset from the start of the file to the string data for this item. The offset should be to a location in the data section, and the data should be in the format specified by "string_data_item" below. There is no alignment requirement for the offset.
}

type StringDataItem struct {
	Utf16Size uleb128 //size of this string, in UTF-16 code units
	Data      []uint8 //a series of MUTF-8 code units followed by a byte of value 0
}

type TypeIdItem struct {
	DescriptorIdx uint32 //index into the StringIds list for the descriptor string of this type
}

type ProtoIdItem struct {
	ShortyIdx     uint32 //index into the StringIds list for the short-form descriptor string of this prototype
	ReturnTypeIdx uint32 //index into the TypeIds list for the return type of this prototype
	ParametersOff uint32 //offset from the start of the file to the list of parameter types for this prototype, or 0 if this prototype has no parameters
}

type FieldIdItem struct {
	ClassIdx uint16 //index into the TypeIdx list for the definer of this field
	TypeIdx  uint16 //index into the TypeIdx list for the type of this field
	NameIdx  uint32 //index into the StringIdx list for the name of this field
}

type MethodIdItem struct {
	ClassIdx uint16 //index into the TypeIdx list for the definer of this method
	ProtoIdx uint16 //index into the ProtoIds list for the prototype of this method
	NameIdx  uint32 //index into the StringIds list for the name of this method
}

type ClassDefItem struct {
	ClassIdx        uint32 //index into the TypeIds list for this class
	AccessFlags     uint32 //access flags for the class (public, final, etc.)
	SuperClassIdx   uint32 //index into the TypeIds list for the superclass, or the constant value NO_INDEX if this class has no superclass
	InterfacesOff   uint32 //offset from the start of the file to the list of interfaces, or 0 if there are none
	SourceFileIdx   uint32 //index into the StringIds list for the name of the file containing the original source for (at least most of) this class, or the special value NO_INDEX to represent a lack of this information.
	AnnotationsOff  uint32 //offset from the start of the file to the annotations structure for this class
	ClassDataOff    uint32 //offset from the start of the file to the associated class data for this item
	StaticValuesOff uint32 //offset from the start of the file to the list of initial values for static fields
}

type CallSiteIdItem struct {
	CallSiteOff uint32 //offset from the start of the file to call site definition
}

type CallSiteItem struct {
}

type MethodHandleItem struct {
	MethodHandleType uint16 //type of the method handle
	Unused1          uint16 //(unused)
	FieldOrMethodId  uint16 //Field or method id depending on whether the method handle type is an accessor or a method invoker
	Unused2          uint16 //(unused)
}

type MapList struct {
	Size uint32    //size of the list, in entries
	List []MapItem //elements of the list
}

type MapItem struct {
	Type   uint16 //type of the items; see table below
	Unused uint16 //(unused)
	Size   uint32 //count of the number of items to be found at the indicated offset
	Offset uint32 //offset from the start of the file to the items in question
}

type AccessFlags uint64

const (
	ACC_PUBLIC       AccessFlags = 1 << iota
	ACC_PRIVATE      AccessFlags = 1 << iota
	ACC_PROTECTED    AccessFlags = 1 << iota
	ACC_STATIC       AccessFlags = 1 << iota
	ACC_FINAL        AccessFlags = 1 << iota
	ACC_SYNCHRONIZED AccessFlags = 1 << iota
	ACC_VOLATILE     AccessFlags = 1 << iota
	ACC_TRANSIENT    AccessFlags = 1 << iota
	ACC_NATIVE       AccessFlags = 1 << iota
	ACC_INTERFACE    AccessFlags = 1 << iota
	ACC_ABSTRACT     AccessFlags = 1 << iota
	ACC_STRICT       AccessFlags = 1 << iota
	ACC_SYNTHETIC    AccessFlags = 1 << iota
	ACC_ANNOTATION   AccessFlags = 1 << iota
	ACC_ENUM         AccessFlags = 1 << iota
	_
	ACC_CONSTRUCTOR           AccessFlags = 1 << iota
	ACC_DECLARED_SYNCHRONIZED AccessFlags = 1 << iota
)

const (
	ACC_BRIDGE  AccessFlags = ACC_VOLATILE
	ACC_VARARGS AccessFlags = ACC_TRANSIENT
)
