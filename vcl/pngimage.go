
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
    . "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TPngImage struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewPngImage() *TPngImage {
    p := new(TPngImage)
    p.instance = PngImage_Create()
    p.ptr = unsafe.Pointer(p.instance)
    // 不是TComponent应该是可以考虑加上的
    // runtime.SetFinalizer(p, (*TPngImage).Free)
    return p
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsPngImage(obj interface{}) *TPngImage {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TPngImage{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsPngImage.
func PngImageFromInst(inst uintptr) *TPngImage {
    return AsPngImage(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsPngImage.
func PngImageFromObj(obj IObject) *TPngImage {
    return AsPngImage(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsPngImage.
func PngImageFromUnsafePointer(ptr unsafe.Pointer) *TPngImage {
    return AsPngImage(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (p *TPngImage) Free() {
    if p.instance != 0 {
        PngImage_Free(p.instance)
        p.instance, p.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (p *TPngImage) Instance() uintptr {
    return p.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (p *TPngImage) UnsafeAddr() unsafe.Pointer {
    return p.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (p *TPngImage) IsValid() bool {
    return p.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (p *TPngImage) Is() TIs {
    return TIs(p.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (p *TPngImage) As() TAs {
//    return TAs(p.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TPngImageClass() TClass {
    return PngImage_StaticClassType()
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (p *TPngImage) Assign(Source IObject) {
    PngImage_Assign(p.instance, CheckPtr(Source))
}

// 文件流加载。
func (p *TPngImage) LoadFromStream(Stream IStream) {
    PngImage_LoadFromStream(p.instance, CheckPtr(Stream))
}

// 保存至流。
func (p *TPngImage) SaveToStream(Stream IStream) {
    PngImage_SaveToStream(p.instance, CheckPtr(Stream))
}

func (p *TPngImage) LoadFromResourceName(Instance uintptr, Name string) {
    PngImage_LoadFromResourceName(p.instance, Instance , Name)
}

func (p *TPngImage) LoadFromResourceID(Instance uintptr, ResID int32) {
    PngImage_LoadFromResourceID(p.instance, Instance , ResID)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (p *TPngImage) Equals(Obj IObject) bool {
    return PngImage_Equals(p.instance, CheckPtr(Obj))
}

// 从文件加载。
func (p *TPngImage) LoadFromFile(Filename string) {
    PngImage_LoadFromFile(p.instance, Filename)
}

// 保存至文件。
func (p *TPngImage) SaveToFile(Filename string) {
    PngImage_SaveToFile(p.instance, Filename)
}

func (p *TPngImage) SetSize(AWidth int32, AHeight int32) {
    PngImage_SetSize(p.instance, AWidth , AHeight)
}

// 获取类名路径。
//
// Get the class name path.
func (p *TPngImage) GetNamePath() string {
    return PngImage_GetNamePath(p.instance)
}

// 获取类的类型信息。
//
// Get class type information.
func (p *TPngImage) ClassType() TClass {
    return PngImage_ClassType(p.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (p *TPngImage) ClassName() string {
    return PngImage_ClassName(p.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (p *TPngImage) InstanceSize() int32 {
    return PngImage_InstanceSize(p.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (p *TPngImage) InheritsFrom(AClass TClass) bool {
    return PngImage_InheritsFrom(p.instance, AClass)
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (p *TPngImage) GetHashCode() int32 {
    return PngImage_GetHashCode(p.instance)
}

// 文本类信息。
//
// Text information.
func (p *TPngImage) ToString() string {
    return PngImage_ToString(p.instance)
}

// 获取画布。
func (p *TPngImage) Canvas() *TCanvas {
    return AsCanvas(PngImage_GetCanvas(p.instance))
}

// 获取宽度。
//
// Get width.
func (p *TPngImage) Width() int32 {
    return PngImage_GetWidth(p.instance)
}

// 获取高度。
//
// Get height.
func (p *TPngImage) Height() int32 {
    return PngImage_GetHeight(p.instance)
}

func (p *TPngImage) Empty() bool {
    return PngImage_GetEmpty(p.instance)
}

// 获取修改。
//
// Get modified.
func (p *TPngImage) Modified() bool {
    return PngImage_GetModified(p.instance)
}

// 设置修改。
//
// Set modified.
func (p *TPngImage) SetModified(value bool) {
    PngImage_SetModified(p.instance, value)
}

func (p *TPngImage) Palette() HPALETTE {
    return PngImage_GetPalette(p.instance)
}

func (p *TPngImage) SetPalette(value HPALETTE) {
    PngImage_SetPalette(p.instance, value)
}

func (p *TPngImage) PaletteModified() bool {
    return PngImage_GetPaletteModified(p.instance)
}

func (p *TPngImage) SetPaletteModified(value bool) {
    PngImage_SetPaletteModified(p.instance, value)
}

// 获取透明。
//
// Get transparent.
func (p *TPngImage) Transparent() bool {
    return PngImage_GetTransparent(p.instance)
}

// 设置透明。
//
// Set transparent.
func (p *TPngImage) SetTransparent(value bool) {
    PngImage_SetTransparent(p.instance, value)
}

// 设置改变事件。
//
// Set changed event.
func (p *TPngImage) SetOnChange(fn TNotifyEvent) {
    PngImage_SetOnChange(p.instance, fn)
}

