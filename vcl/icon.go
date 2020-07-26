
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

type TIcon struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewIcon() *TIcon {
    i := new(TIcon)
    i.instance = Icon_Create()
    i.ptr = unsafe.Pointer(i.instance)
    // 不是TComponent应该是可以考虑加上的
    // runtime.SetFinalizer(i, (*TIcon).Free)
    return i
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsIcon(obj interface{}) *TIcon {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TIcon{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsIcon.
func IconFromInst(inst uintptr) *TIcon {
    return AsIcon(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsIcon.
func IconFromObj(obj IObject) *TIcon {
    return AsIcon(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsIcon.
func IconFromUnsafePointer(ptr unsafe.Pointer) *TIcon {
    return AsIcon(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (i *TIcon) Free() {
    if i.instance != 0 {
        Icon_Free(i.instance)
        i.instance, i.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (i *TIcon) Instance() uintptr {
    return i.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (i *TIcon) UnsafeAddr() unsafe.Pointer {
    return i.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (i *TIcon) IsValid() bool {
    return i.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (i *TIcon) Is() TIs {
    return TIs(i.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (i *TIcon) As() TAs {
//    return TAs(i.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TIconClass() TClass {
    return Icon_StaticClassType()
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (i *TIcon) Assign(Source IObject) {
    Icon_Assign(i.instance, CheckPtr(Source))
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (i *TIcon) HandleAllocated() bool {
    return Icon_HandleAllocated(i.instance)
}

// 文件流加载。
func (i *TIcon) LoadFromStream(Stream IStream) {
    Icon_LoadFromStream(i.instance, CheckPtr(Stream))
}

// 保存至流。
func (i *TIcon) SaveToStream(Stream IStream) {
    Icon_SaveToStream(i.instance, CheckPtr(Stream))
}

func (i *TIcon) SetSize(AWidth int32, AHeight int32) {
    Icon_SetSize(i.instance, AWidth , AHeight)
}

func (i *TIcon) LoadFromResourceName(Instance uintptr, ResName string) {
    Icon_LoadFromResourceName(i.instance, Instance , ResName)
}

func (i *TIcon) LoadFromResourceID(Instance uintptr, ResID int32) {
    Icon_LoadFromResourceID(i.instance, Instance , ResID)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (i *TIcon) Equals(Obj IObject) bool {
    return Icon_Equals(i.instance, CheckPtr(Obj))
}

// 从文件加载。
func (i *TIcon) LoadFromFile(Filename string) {
    Icon_LoadFromFile(i.instance, Filename)
}

// 保存至文件。
func (i *TIcon) SaveToFile(Filename string) {
    Icon_SaveToFile(i.instance, Filename)
}

// 获取类名路径。
//
// Get the class name path.
func (i *TIcon) GetNamePath() string {
    return Icon_GetNamePath(i.instance)
}

// 获取类的类型信息。
//
// Get class type information.
func (i *TIcon) ClassType() TClass {
    return Icon_ClassType(i.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (i *TIcon) ClassName() string {
    return Icon_ClassName(i.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (i *TIcon) InstanceSize() int32 {
    return Icon_InstanceSize(i.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (i *TIcon) InheritsFrom(AClass TClass) bool {
    return Icon_InheritsFrom(i.instance, AClass)
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (i *TIcon) GetHashCode() int32 {
    return Icon_GetHashCode(i.instance)
}

// 文本类信息。
//
// Text information.
func (i *TIcon) ToString() string {
    return Icon_ToString(i.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (i *TIcon) Handle() HICON {
    return Icon_GetHandle(i.instance)
}

// 设置控件句柄。
//
// Set Control handle.
func (i *TIcon) SetHandle(value HICON) {
    Icon_SetHandle(i.instance, value)
}

func (i *TIcon) Empty() bool {
    return Icon_GetEmpty(i.instance)
}

// 获取高度。
//
// Get height.
func (i *TIcon) Height() int32 {
    return Icon_GetHeight(i.instance)
}

// 设置高度。
//
// Set height.
func (i *TIcon) SetHeight(value int32) {
    Icon_SetHeight(i.instance, value)
}

// 获取修改。
//
// Get modified.
func (i *TIcon) Modified() bool {
    return Icon_GetModified(i.instance)
}

// 设置修改。
//
// Set modified.
func (i *TIcon) SetModified(value bool) {
    Icon_SetModified(i.instance, value)
}

func (i *TIcon) Palette() HPALETTE {
    return Icon_GetPalette(i.instance)
}

func (i *TIcon) SetPalette(value HPALETTE) {
    Icon_SetPalette(i.instance, value)
}

func (i *TIcon) PaletteModified() bool {
    return Icon_GetPaletteModified(i.instance)
}

func (i *TIcon) SetPaletteModified(value bool) {
    Icon_SetPaletteModified(i.instance, value)
}

// 获取透明。
//
// Get transparent.
func (i *TIcon) Transparent() bool {
    return Icon_GetTransparent(i.instance)
}

// 设置透明。
//
// Set transparent.
func (i *TIcon) SetTransparent(value bool) {
    Icon_SetTransparent(i.instance, value)
}

// 获取宽度。
//
// Get width.
func (i *TIcon) Width() int32 {
    return Icon_GetWidth(i.instance)
}

// 设置宽度。
//
// Set width.
func (i *TIcon) SetWidth(value int32) {
    Icon_SetWidth(i.instance, value)
}

// 设置改变事件。
//
// Set changed event.
func (i *TIcon) SetOnChange(fn TNotifyEvent) {
    Icon_SetOnChange(i.instance, fn)
}

