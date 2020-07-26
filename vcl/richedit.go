
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

type TRichEdit struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewRichEdit(owner IComponent) *TRichEdit {
    r := new(TRichEdit)
    r.instance = RichEdit_Create(CheckPtr(owner))
    r.ptr = unsafe.Pointer(r.instance)
    // 不是TComponent应该是可以考虑加上的
    // runtime.SetFinalizer(r, (*TRichEdit).Free)
    return r
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsRichEdit(obj interface{}) *TRichEdit {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TRichEdit{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsRichEdit.
func RichEditFromInst(inst uintptr) *TRichEdit {
    return AsRichEdit(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsRichEdit.
func RichEditFromObj(obj IObject) *TRichEdit {
    return AsRichEdit(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsRichEdit.
func RichEditFromUnsafePointer(ptr unsafe.Pointer) *TRichEdit {
    return AsRichEdit(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (r *TRichEdit) Free() {
    if r.instance != 0 {
        RichEdit_Free(r.instance)
        r.instance, r.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (r *TRichEdit) Instance() uintptr {
    return r.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (r *TRichEdit) UnsafeAddr() unsafe.Pointer {
    return r.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (r *TRichEdit) IsValid() bool {
    return r.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (r *TRichEdit) Is() TIs {
    return TIs(r.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (r *TRichEdit) As() TAs {
//    return TAs(r.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TRichEditClass() TClass {
    return RichEdit_StaticClassType()
}

// 清除。
func (r *TRichEdit) Clear() {
    RichEdit_Clear(r.instance)
}

func (r *TRichEdit) FindText(SearchStr string, StartPos int32, Length int32, Options TSearchTypes) int32 {
    return RichEdit_FindText(r.instance, SearchStr , StartPos , Length , Options)
}

// 清除选择。
func (r *TRichEdit) ClearSelection() {
    RichEdit_ClearSelection(r.instance)
}

// 复制到粘贴板。
func (r *TRichEdit) CopyToClipboard() {
    RichEdit_CopyToClipboard(r.instance)
}

// 剪切到粘贴板。
func (r *TRichEdit) CutToClipboard() {
    RichEdit_CutToClipboard(r.instance)
}

// 从剪切板粘贴。
func (r *TRichEdit) PasteFromClipboard() {
    RichEdit_PasteFromClipboard(r.instance)
}

// 撤销上一次操作。
func (r *TRichEdit) Undo() {
    RichEdit_Undo(r.instance)
}

// 全选。
func (r *TRichEdit) SelectAll() {
    RichEdit_SelectAll(r.instance)
}

// 是否可以获得焦点。
func (r *TRichEdit) CanFocus() bool {
    return RichEdit_CanFocus(r.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (r *TRichEdit) ContainsControl(Control IControl) bool {
    return RichEdit_ContainsControl(r.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (r *TRichEdit) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(RichEdit_ControlAtPos(r.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (r *TRichEdit) DisableAlign() {
    RichEdit_DisableAlign(r.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (r *TRichEdit) EnableAlign() {
    RichEdit_EnableAlign(r.instance)
}

// 查找子控件。
//
// Find sub controls.
func (r *TRichEdit) FindChildControl(ControlName string) *TControl {
    return AsControl(RichEdit_FindChildControl(r.instance, ControlName))
}

func (r *TRichEdit) FlipChildren(AllLevels bool) {
    RichEdit_FlipChildren(r.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (r *TRichEdit) Focused() bool {
    return RichEdit_Focused(r.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (r *TRichEdit) HandleAllocated() bool {
    return RichEdit_HandleAllocated(r.instance)
}

// 插入一个控件。
//
// Insert a control.
func (r *TRichEdit) InsertControl(AControl IControl) {
    RichEdit_InsertControl(r.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (r *TRichEdit) Invalidate() {
    RichEdit_Invalidate(r.instance)
}

// 移除一个控件。
//
// Remove a control.
func (r *TRichEdit) RemoveControl(AControl IControl) {
    RichEdit_RemoveControl(r.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (r *TRichEdit) Realign() {
    RichEdit_Realign(r.instance)
}

// 重绘。
//
// Repaint.
func (r *TRichEdit) Repaint() {
    RichEdit_Repaint(r.instance)
}

// 按比例缩放。
//
// Scale by.
func (r *TRichEdit) ScaleBy(M int32, D int32) {
    RichEdit_ScaleBy(r.instance, M , D)
}

// 滚动至指定位置。
//
// Scroll by.
func (r *TRichEdit) ScrollBy(DeltaX int32, DeltaY int32) {
    RichEdit_ScrollBy(r.instance, DeltaX , DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (r *TRichEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    RichEdit_SetBounds(r.instance, ALeft , ATop , AWidth , AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (r *TRichEdit) SetFocus() {
    RichEdit_SetFocus(r.instance)
}

// 控件更新。
//
// Update.
func (r *TRichEdit) Update() {
    RichEdit_Update(r.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (r *TRichEdit) BringToFront() {
    RichEdit_BringToFront(r.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (r *TRichEdit) ClientToScreen(Point TPoint) TPoint {
    return RichEdit_ClientToScreen(r.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (r *TRichEdit) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return RichEdit_ClientToParent(r.instance, Point , CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (r *TRichEdit) Dragging() bool {
    return RichEdit_Dragging(r.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (r *TRichEdit) HasParent() bool {
    return RichEdit_HasParent(r.instance)
}

// 隐藏控件。
//
// Hidden control.
func (r *TRichEdit) Hide() {
    RichEdit_Hide(r.instance)
}

// 发送一个消息。
//
// Send a message.
func (r *TRichEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return RichEdit_Perform(r.instance, Msg , WParam , LParam)
}

// 刷新控件。
//
// Refresh control.
func (r *TRichEdit) Refresh() {
    RichEdit_Refresh(r.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (r *TRichEdit) ScreenToClient(Point TPoint) TPoint {
    return RichEdit_ScreenToClient(r.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (r *TRichEdit) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return RichEdit_ParentToClient(r.instance, Point , CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (r *TRichEdit) SendToBack() {
    RichEdit_SendToBack(r.instance)
}

// 显示控件。
//
// Show control.
func (r *TRichEdit) Show() {
    RichEdit_Show(r.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (r *TRichEdit) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return RichEdit_GetTextBuf(r.instance, Buffer , BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (r *TRichEdit) GetTextLen() int32 {
    return RichEdit_GetTextLen(r.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (r *TRichEdit) SetTextBuf(Buffer string) {
    RichEdit_SetTextBuf(r.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (r *TRichEdit) FindComponent(AName string) *TComponent {
    return AsComponent(RichEdit_FindComponent(r.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (r *TRichEdit) GetNamePath() string {
    return RichEdit_GetNamePath(r.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (r *TRichEdit) Assign(Source IObject) {
    RichEdit_Assign(r.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (r *TRichEdit) ClassType() TClass {
    return RichEdit_ClassType(r.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (r *TRichEdit) ClassName() string {
    return RichEdit_ClassName(r.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (r *TRichEdit) InstanceSize() int32 {
    return RichEdit_InstanceSize(r.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (r *TRichEdit) InheritsFrom(AClass TClass) bool {
    return RichEdit_InheritsFrom(r.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (r *TRichEdit) Equals(Obj IObject) bool {
    return RichEdit_Equals(r.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (r *TRichEdit) GetHashCode() int32 {
    return RichEdit_GetHashCode(r.instance)
}

// 文本类信息。
//
// Text information.
func (r *TRichEdit) ToString() string {
    return RichEdit_ToString(r.instance)
}

func (r *TRichEdit) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    RichEdit_AnchorToNeighbour(r.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (r *TRichEdit) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    RichEdit_AnchorParallel(r.instance, ASide , ASpace , CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (r *TRichEdit) AnchorHorizontalCenterTo(ASibling IControl) {
    RichEdit_AnchorHorizontalCenterTo(r.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (r *TRichEdit) AnchorVerticalCenterTo(ASibling IControl) {
    RichEdit_AnchorVerticalCenterTo(r.instance, CheckPtr(ASibling))
}

func (r *TRichEdit) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    RichEdit_AnchorAsAlign(r.instance, ATheAlign , ASpace)
}

func (r *TRichEdit) AnchorClient(ASpace int32) {
    RichEdit_AnchorClient(r.instance, ASpace)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (r *TRichEdit) Align() TAlign {
    return RichEdit_GetAlign(r.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (r *TRichEdit) SetAlign(value TAlign) {
    RichEdit_SetAlign(r.instance, value)
}

// 获取文字对齐。
//
// Get Text alignment.
func (r *TRichEdit) Alignment() TAlignment {
    return RichEdit_GetAlignment(r.instance)
}

// 设置文字对齐。
//
// Set Text alignment.
func (r *TRichEdit) SetAlignment(value TAlignment) {
    RichEdit_SetAlignment(r.instance, value)
}

// 获取四个角位置的锚点。
func (r *TRichEdit) Anchors() TAnchors {
    return RichEdit_GetAnchors(r.instance)
}

// 设置四个角位置的锚点。
func (r *TRichEdit) SetAnchors(value TAnchors) {
    RichEdit_SetAnchors(r.instance, value)
}

func (r *TRichEdit) BiDiMode() TBiDiMode {
    return RichEdit_GetBiDiMode(r.instance)
}

func (r *TRichEdit) SetBiDiMode(value TBiDiMode) {
    RichEdit_SetBiDiMode(r.instance, value)
}

// 获取窗口边框样式。比如：无边框，单一边框等。
func (r *TRichEdit) BorderStyle() TBorderStyle {
    return RichEdit_GetBorderStyle(r.instance)
}

// 设置窗口边框样式。比如：无边框，单一边框等。
func (r *TRichEdit) SetBorderStyle(value TBorderStyle) {
    RichEdit_SetBorderStyle(r.instance, value)
}

// 获取边框的宽度。
func (r *TRichEdit) BorderWidth() int32 {
    return RichEdit_GetBorderWidth(r.instance)
}

// 设置边框的宽度。
func (r *TRichEdit) SetBorderWidth(value int32) {
    RichEdit_SetBorderWidth(r.instance, value)
}

// 获取颜色。
//
// Get color.
func (r *TRichEdit) Color() TColor {
    return RichEdit_GetColor(r.instance)
}

// 设置颜色。
//
// Set color.
func (r *TRichEdit) SetColor(value TColor) {
    RichEdit_SetColor(r.instance, value)
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (r *TRichEdit) DragCursor() TCursor {
    return RichEdit_GetDragCursor(r.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (r *TRichEdit) SetDragCursor(value TCursor) {
    RichEdit_SetDragCursor(r.instance, value)
}

// 获取拖拽方式。
//
// Get Drag and drop.
func (r *TRichEdit) DragKind() TDragKind {
    return RichEdit_GetDragKind(r.instance)
}

// 设置拖拽方式。
//
// Set Drag and drop.
func (r *TRichEdit) SetDragKind(value TDragKind) {
    RichEdit_SetDragKind(r.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (r *TRichEdit) DragMode() TDragMode {
    return RichEdit_GetDragMode(r.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (r *TRichEdit) SetDragMode(value TDragMode) {
    RichEdit_SetDragMode(r.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (r *TRichEdit) Enabled() bool {
    return RichEdit_GetEnabled(r.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (r *TRichEdit) SetEnabled(value bool) {
    RichEdit_SetEnabled(r.instance, value)
}

// 获取字体。
//
// Get Font.
func (r *TRichEdit) Font() *TFont {
    return AsFont(RichEdit_GetFont(r.instance))
}

// 设置字体。
//
// Set Font.
func (r *TRichEdit) SetFont(value *TFont) {
    RichEdit_SetFont(r.instance, CheckPtr(value))
}

// 获取隐藏选择。
func (r *TRichEdit) HideSelection() bool {
    return RichEdit_GetHideSelection(r.instance)
}

// 设置隐藏选择。
func (r *TRichEdit) SetHideSelection(value bool) {
    RichEdit_SetHideSelection(r.instance, value)
}

// 获取约束控件大小。
func (r *TRichEdit) Constraints() *TSizeConstraints {
    return AsSizeConstraints(RichEdit_GetConstraints(r.instance))
}

// 设置约束控件大小。
func (r *TRichEdit) SetConstraints(value *TSizeConstraints) {
    RichEdit_SetConstraints(r.instance, CheckPtr(value))
}

func (r *TRichEdit) Lines() *TStrings {
    return AsStrings(RichEdit_GetLines(r.instance))
}

func (r *TRichEdit) SetLines(value IStrings) {
    RichEdit_SetLines(r.instance, CheckPtr(value))
}

// 获取最大长度。
func (r *TRichEdit) MaxLength() int32 {
    return RichEdit_GetMaxLength(r.instance)
}

// 设置最大长度。
func (r *TRichEdit) SetMaxLength(value int32) {
    RichEdit_SetMaxLength(r.instance, value)
}

// 获取使用父容器颜色。
//
// Get parent color.
func (r *TRichEdit) ParentColor() bool {
    return RichEdit_GetParentColor(r.instance)
}

// 设置使用父容器颜色。
//
// Set parent color.
func (r *TRichEdit) SetParentColor(value bool) {
    RichEdit_SetParentColor(r.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (r *TRichEdit) ParentFont() bool {
    return RichEdit_GetParentFont(r.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (r *TRichEdit) SetParentFont(value bool) {
    RichEdit_SetParentFont(r.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (r *TRichEdit) ParentShowHint() bool {
    return RichEdit_GetParentShowHint(r.instance)
}

// 设置以父容器的ShowHint属性为准。
func (r *TRichEdit) SetParentShowHint(value bool) {
    RichEdit_SetParentShowHint(r.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (r *TRichEdit) PopupMenu() *TPopupMenu {
    return AsPopupMenu(RichEdit_GetPopupMenu(r.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (r *TRichEdit) SetPopupMenu(value IComponent) {
    RichEdit_SetPopupMenu(r.instance, CheckPtr(value))
}

// 获取只读。
func (r *TRichEdit) ReadOnly() bool {
    return RichEdit_GetReadOnly(r.instance)
}

// 设置只读。
func (r *TRichEdit) SetReadOnly(value bool) {
    RichEdit_SetReadOnly(r.instance, value)
}

func (r *TRichEdit) ScrollBars() TScrollStyle {
    return RichEdit_GetScrollBars(r.instance)
}

func (r *TRichEdit) SetScrollBars(value TScrollStyle) {
    RichEdit_SetScrollBars(r.instance, value)
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (r *TRichEdit) ShowHint() bool {
    return RichEdit_GetShowHint(r.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (r *TRichEdit) SetShowHint(value bool) {
    RichEdit_SetShowHint(r.instance, value)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (r *TRichEdit) TabOrder() TTabOrder {
    return RichEdit_GetTabOrder(r.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (r *TRichEdit) SetTabOrder(value TTabOrder) {
    RichEdit_SetTabOrder(r.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (r *TRichEdit) TabStop() bool {
    return RichEdit_GetTabStop(r.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (r *TRichEdit) SetTabStop(value bool) {
    RichEdit_SetTabStop(r.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (r *TRichEdit) Visible() bool {
    return RichEdit_GetVisible(r.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (r *TRichEdit) SetVisible(value bool) {
    RichEdit_SetVisible(r.instance, value)
}

func (r *TRichEdit) WantTabs() bool {
    return RichEdit_GetWantTabs(r.instance)
}

func (r *TRichEdit) SetWantTabs(value bool) {
    RichEdit_SetWantTabs(r.instance, value)
}

func (r *TRichEdit) WantReturns() bool {
    return RichEdit_GetWantReturns(r.instance)
}

func (r *TRichEdit) SetWantReturns(value bool) {
    RichEdit_SetWantReturns(r.instance, value)
}

// 获取自动换行。
//
// Get Automatic line break.
func (r *TRichEdit) WordWrap() bool {
    return RichEdit_GetWordWrap(r.instance)
}

// 设置自动换行。
//
// Set Automatic line break.
func (r *TRichEdit) SetWordWrap(value bool) {
    RichEdit_SetWordWrap(r.instance, value)
}

func (r *TRichEdit) Zoom() int32 {
    return RichEdit_GetZoom(r.instance)
}

func (r *TRichEdit) SetZoom(value int32) {
    RichEdit_SetZoom(r.instance, value)
}

// 设置改变事件。
//
// Set changed event.
func (r *TRichEdit) SetOnChange(fn TNotifyEvent) {
    RichEdit_SetOnChange(r.instance, fn)
}

// 设置控件单击事件。
//
// Set control click event.
func (r *TRichEdit) SetOnClick(fn TNotifyEvent) {
    RichEdit_SetOnClick(r.instance, fn)
}

// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (r *TRichEdit) SetOnContextPopup(fn TContextPopupEvent) {
    RichEdit_SetOnContextPopup(r.instance, fn)
}

// 设置双击事件。
func (r *TRichEdit) SetOnDblClick(fn TNotifyEvent) {
    RichEdit_SetOnDblClick(r.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (r *TRichEdit) SetOnDragDrop(fn TDragDropEvent) {
    RichEdit_SetOnDragDrop(r.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (r *TRichEdit) SetOnDragOver(fn TDragOverEvent) {
    RichEdit_SetOnDragOver(r.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (r *TRichEdit) SetOnEndDrag(fn TEndDragEvent) {
    RichEdit_SetOnEndDrag(r.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (r *TRichEdit) SetOnEnter(fn TNotifyEvent) {
    RichEdit_SetOnEnter(r.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (r *TRichEdit) SetOnExit(fn TNotifyEvent) {
    RichEdit_SetOnExit(r.instance, fn)
}

// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (r *TRichEdit) SetOnKeyDown(fn TKeyEvent) {
    RichEdit_SetOnKeyDown(r.instance, fn)
}

// 设置键键下事件。
func (r *TRichEdit) SetOnKeyPress(fn TKeyPressEvent) {
    RichEdit_SetOnKeyPress(r.instance, fn)
}

// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (r *TRichEdit) SetOnKeyUp(fn TKeyEvent) {
    RichEdit_SetOnKeyUp(r.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (r *TRichEdit) SetOnMouseDown(fn TMouseEvent) {
    RichEdit_SetOnMouseDown(r.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (r *TRichEdit) SetOnMouseEnter(fn TNotifyEvent) {
    RichEdit_SetOnMouseEnter(r.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (r *TRichEdit) SetOnMouseLeave(fn TNotifyEvent) {
    RichEdit_SetOnMouseLeave(r.instance, fn)
}

// 设置鼠标移动事件。
func (r *TRichEdit) SetOnMouseMove(fn TMouseMoveEvent) {
    RichEdit_SetOnMouseMove(r.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (r *TRichEdit) SetOnMouseUp(fn TMouseEvent) {
    RichEdit_SetOnMouseUp(r.instance, fn)
}

// 设置鼠标滚轮事件。
func (r *TRichEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
    RichEdit_SetOnMouseWheel(r.instance, fn)
}

// 设置鼠标滚轮按下事件。
func (r *TRichEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
    RichEdit_SetOnMouseWheelDown(r.instance, fn)
}

// 设置鼠标滚轮抬起事件。
func (r *TRichEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
    RichEdit_SetOnMouseWheelUp(r.instance, fn)
}

func (r *TRichEdit) DefAttributes() *TTextAttributes {
    return AsTextAttributes(RichEdit_GetDefAttributes(r.instance))
}

func (r *TRichEdit) SetDefAttributes(value *TTextAttributes) {
    RichEdit_SetDefAttributes(r.instance, CheckPtr(value))
}

func (r *TRichEdit) SelAttributes() *TTextAttributes {
    return AsTextAttributes(RichEdit_GetSelAttributes(r.instance))
}

func (r *TRichEdit) SetSelAttributes(value *TTextAttributes) {
    RichEdit_SetSelAttributes(r.instance, CheckPtr(value))
}

func (r *TRichEdit) Paragraph() *TParaAttributes {
    return AsParaAttributes(RichEdit_GetParagraph(r.instance))
}

func (r *TRichEdit) CaretPos() TPoint {
    return RichEdit_GetCaretPos(r.instance)
}

func (r *TRichEdit) SetCaretPos(value TPoint) {
    RichEdit_SetCaretPos(r.instance, value)
}

// 获取能否撤销。
func (r *TRichEdit) CanUndo() bool {
    return RichEdit_GetCanUndo(r.instance)
}

// 获取修改。
//
// Get modified.
func (r *TRichEdit) Modified() bool {
    return RichEdit_GetModified(r.instance)
}

// 设置修改。
//
// Set modified.
func (r *TRichEdit) SetModified(value bool) {
    RichEdit_SetModified(r.instance, value)
}

// 获取选择的长度。
func (r *TRichEdit) SelLength() int32 {
    return RichEdit_GetSelLength(r.instance)
}

// 设置选择的长度。
func (r *TRichEdit) SetSelLength(value int32) {
    RichEdit_SetSelLength(r.instance, value)
}

// 获取选择的启始位置。
func (r *TRichEdit) SelStart() int32 {
    return RichEdit_GetSelStart(r.instance)
}

// 设置选择的启始位置。
func (r *TRichEdit) SetSelStart(value int32) {
    RichEdit_SetSelStart(r.instance, value)
}

// 获取选择的文本。
func (r *TRichEdit) SelText() string {
    return RichEdit_GetSelText(r.instance)
}

// 设置选择的文本。
func (r *TRichEdit) SetSelText(value string) {
    RichEdit_SetSelText(r.instance, value)
}

// 获取文本。
func (r *TRichEdit) Text() string {
    strLen := r.GetTextLen()
    if strLen != 0 {
        var buffStr string
        r.GetTextBuf(&buffStr, strLen + 1)
        return buffStr
    }
    return ""
}

// 设置文本。
func (r *TRichEdit) SetText(value string) {
    RichEdit_SetText(r.instance, value)
}

// 获取提示文本。
func (r *TRichEdit) TextHint() string {
    return RichEdit_GetTextHint(r.instance)
}

// 设置提示文本。
func (r *TRichEdit) SetTextHint(value string) {
    RichEdit_SetTextHint(r.instance, value)
}

// 获取依靠客户端总数。
func (r *TRichEdit) DockClientCount() int32 {
    return RichEdit_GetDockClientCount(r.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (r *TRichEdit) DockSite() bool {
    return RichEdit_GetDockSite(r.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (r *TRichEdit) SetDockSite(value bool) {
    RichEdit_SetDockSite(r.instance, value)
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (r *TRichEdit) DoubleBuffered() bool {
    return RichEdit_GetDoubleBuffered(r.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (r *TRichEdit) SetDoubleBuffered(value bool) {
    RichEdit_SetDoubleBuffered(r.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (r *TRichEdit) MouseInClient() bool {
    return RichEdit_GetMouseInClient(r.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (r *TRichEdit) VisibleDockClientCount() int32 {
    return RichEdit_GetVisibleDockClientCount(r.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (r *TRichEdit) Brush() *TBrush {
    return AsBrush(RichEdit_GetBrush(r.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (r *TRichEdit) ControlCount() int32 {
    return RichEdit_GetControlCount(r.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (r *TRichEdit) Handle() HWND {
    return RichEdit_GetHandle(r.instance)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (r *TRichEdit) ParentDoubleBuffered() bool {
    return RichEdit_GetParentDoubleBuffered(r.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (r *TRichEdit) SetParentDoubleBuffered(value bool) {
    RichEdit_SetParentDoubleBuffered(r.instance, value)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (r *TRichEdit) ParentWindow() HWND {
    return RichEdit_GetParentWindow(r.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (r *TRichEdit) SetParentWindow(value HWND) {
    RichEdit_SetParentWindow(r.instance, value)
}

func (r *TRichEdit) Showing() bool {
    return RichEdit_GetShowing(r.instance)
}

// 获取使用停靠管理。
func (r *TRichEdit) UseDockManager() bool {
    return RichEdit_GetUseDockManager(r.instance)
}

// 设置使用停靠管理。
func (r *TRichEdit) SetUseDockManager(value bool) {
    RichEdit_SetUseDockManager(r.instance, value)
}

func (r *TRichEdit) Action() *TAction {
    return AsAction(RichEdit_GetAction(r.instance))
}

func (r *TRichEdit) SetAction(value IComponent) {
    RichEdit_SetAction(r.instance, CheckPtr(value))
}

func (r *TRichEdit) BoundsRect() TRect {
    return RichEdit_GetBoundsRect(r.instance)
}

func (r *TRichEdit) SetBoundsRect(value TRect) {
    RichEdit_SetBoundsRect(r.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (r *TRichEdit) ClientHeight() int32 {
    return RichEdit_GetClientHeight(r.instance)
}

// 设置客户区高度。
//
// Set client height.
func (r *TRichEdit) SetClientHeight(value int32) {
    RichEdit_SetClientHeight(r.instance, value)
}

func (r *TRichEdit) ClientOrigin() TPoint {
    return RichEdit_GetClientOrigin(r.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (r *TRichEdit) ClientRect() TRect {
    return RichEdit_GetClientRect(r.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (r *TRichEdit) ClientWidth() int32 {
    return RichEdit_GetClientWidth(r.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (r *TRichEdit) SetClientWidth(value int32) {
    RichEdit_SetClientWidth(r.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (r *TRichEdit) ControlState() TControlState {
    return RichEdit_GetControlState(r.instance)
}

// 设置控件状态。
//
// Set control state.
func (r *TRichEdit) SetControlState(value TControlState) {
    RichEdit_SetControlState(r.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (r *TRichEdit) ControlStyle() TControlStyle {
    return RichEdit_GetControlStyle(r.instance)
}

// 设置控件样式。
//
// Set control style.
func (r *TRichEdit) SetControlStyle(value TControlStyle) {
    RichEdit_SetControlStyle(r.instance, value)
}

func (r *TRichEdit) Floating() bool {
    return RichEdit_GetFloating(r.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (r *TRichEdit) Parent() *TWinControl {
    return AsWinControl(RichEdit_GetParent(r.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (r *TRichEdit) SetParent(value IWinControl) {
    RichEdit_SetParent(r.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (r *TRichEdit) Left() int32 {
    return RichEdit_GetLeft(r.instance)
}

// 设置左边位置。
//
// Set Left position.
func (r *TRichEdit) SetLeft(value int32) {
    RichEdit_SetLeft(r.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (r *TRichEdit) Top() int32 {
    return RichEdit_GetTop(r.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (r *TRichEdit) SetTop(value int32) {
    RichEdit_SetTop(r.instance, value)
}

// 获取宽度。
//
// Get width.
func (r *TRichEdit) Width() int32 {
    return RichEdit_GetWidth(r.instance)
}

// 设置宽度。
//
// Set width.
func (r *TRichEdit) SetWidth(value int32) {
    RichEdit_SetWidth(r.instance, value)
}

// 获取高度。
//
// Get height.
func (r *TRichEdit) Height() int32 {
    return RichEdit_GetHeight(r.instance)
}

// 设置高度。
//
// Set height.
func (r *TRichEdit) SetHeight(value int32) {
    RichEdit_SetHeight(r.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (r *TRichEdit) Cursor() TCursor {
    return RichEdit_GetCursor(r.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (r *TRichEdit) SetCursor(value TCursor) {
    RichEdit_SetCursor(r.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (r *TRichEdit) Hint() string {
    return RichEdit_GetHint(r.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (r *TRichEdit) SetHint(value string) {
    RichEdit_SetHint(r.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (r *TRichEdit) ComponentCount() int32 {
    return RichEdit_GetComponentCount(r.instance)
}

// 获取组件索引。
//
// Get component index.
func (r *TRichEdit) ComponentIndex() int32 {
    return RichEdit_GetComponentIndex(r.instance)
}

// 设置组件索引。
//
// Set component index.
func (r *TRichEdit) SetComponentIndex(value int32) {
    RichEdit_SetComponentIndex(r.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (r *TRichEdit) Owner() *TComponent {
    return AsComponent(RichEdit_GetOwner(r.instance))
}

// 获取组件名称。
//
// Get the component name.
func (r *TRichEdit) Name() string {
    return RichEdit_GetName(r.instance)
}

// 设置组件名称。
//
// Set the component name.
func (r *TRichEdit) SetName(value string) {
    RichEdit_SetName(r.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (r *TRichEdit) Tag() int {
    return RichEdit_GetTag(r.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (r *TRichEdit) SetTag(value int) {
    RichEdit_SetTag(r.instance, value)
}

// 获取左边锚点。
func (r *TRichEdit) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(RichEdit_GetAnchorSideLeft(r.instance))
}

// 设置左边锚点。
func (r *TRichEdit) SetAnchorSideLeft(value *TAnchorSide) {
    RichEdit_SetAnchorSideLeft(r.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (r *TRichEdit) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(RichEdit_GetAnchorSideTop(r.instance))
}

// 设置顶边锚点。
func (r *TRichEdit) SetAnchorSideTop(value *TAnchorSide) {
    RichEdit_SetAnchorSideTop(r.instance, CheckPtr(value))
}

// 获取右边锚点。
func (r *TRichEdit) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(RichEdit_GetAnchorSideRight(r.instance))
}

// 设置右边锚点。
func (r *TRichEdit) SetAnchorSideRight(value *TAnchorSide) {
    RichEdit_SetAnchorSideRight(r.instance, CheckPtr(value))
}

// 获取底边锚点。
func (r *TRichEdit) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(RichEdit_GetAnchorSideBottom(r.instance))
}

// 设置底边锚点。
func (r *TRichEdit) SetAnchorSideBottom(value *TAnchorSide) {
    RichEdit_SetAnchorSideBottom(r.instance, CheckPtr(value))
}

func (r *TRichEdit) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(RichEdit_GetChildSizing(r.instance))
}

func (r *TRichEdit) SetChildSizing(value *TControlChildSizing) {
    RichEdit_SetChildSizing(r.instance, CheckPtr(value))
}

// 获取边框间距。
func (r *TRichEdit) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(RichEdit_GetBorderSpacing(r.instance))
}

// 设置边框间距。
func (r *TRichEdit) SetBorderSpacing(value *TControlBorderSpacing) {
    RichEdit_SetBorderSpacing(r.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (r *TRichEdit) DockClients(Index int32) *TControl {
    return AsControl(RichEdit_GetDockClients(r.instance, Index))
}

// 获取指定索引子控件。
func (r *TRichEdit) Controls(Index int32) *TControl {
    return AsControl(RichEdit_GetControls(r.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (r *TRichEdit) Components(AIndex int32) *TComponent {
    return AsComponent(RichEdit_GetComponents(r.instance, AIndex))
}

// 获取锚侧面。
func (r *TRichEdit) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(RichEdit_GetAnchorSide(r.instance, AKind))
}

