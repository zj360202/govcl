
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

type TFlowPanel struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewFlowPanel(owner IComponent) *TFlowPanel {
    f := new(TFlowPanel)
    f.instance = FlowPanel_Create(CheckPtr(owner))
    f.ptr = unsafe.Pointer(f.instance)
    // 不是TComponent应该是可以考虑加上的
    // runtime.SetFinalizer(f, (*TFlowPanel).Free)
    return f
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsFlowPanel(obj interface{}) *TFlowPanel {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TFlowPanel{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsFlowPanel.
func FlowPanelFromInst(inst uintptr) *TFlowPanel {
    return AsFlowPanel(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsFlowPanel.
func FlowPanelFromObj(obj IObject) *TFlowPanel {
    return AsFlowPanel(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsFlowPanel.
func FlowPanelFromUnsafePointer(ptr unsafe.Pointer) *TFlowPanel {
    return AsFlowPanel(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (f *TFlowPanel) Free() {
    if f.instance != 0 {
        FlowPanel_Free(f.instance)
        f.instance, f.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (f *TFlowPanel) Instance() uintptr {
    return f.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (f *TFlowPanel) UnsafeAddr() unsafe.Pointer {
    return f.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (f *TFlowPanel) IsValid() bool {
    return f.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (f *TFlowPanel) Is() TIs {
    return TIs(f.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (f *TFlowPanel) As() TAs {
//    return TAs(f.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TFlowPanelClass() TClass {
    return FlowPanel_StaticClassType()
}

func (f *TFlowPanel) GetControlIndex(AControl IControl) int32 {
    return FlowPanel_GetControlIndex(f.instance, CheckPtr(AControl))
}

func (f *TFlowPanel) SetControlIndex(AControl IControl, Index int32) {
    FlowPanel_SetControlIndex(f.instance, CheckPtr(AControl), Index)
}

// 是否可以获得焦点。
func (f *TFlowPanel) CanFocus() bool {
    return FlowPanel_CanFocus(f.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (f *TFlowPanel) ContainsControl(Control IControl) bool {
    return FlowPanel_ContainsControl(f.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (f *TFlowPanel) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(FlowPanel_ControlAtPos(f.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (f *TFlowPanel) DisableAlign() {
    FlowPanel_DisableAlign(f.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (f *TFlowPanel) EnableAlign() {
    FlowPanel_EnableAlign(f.instance)
}

// 查找子控件。
//
// Find sub controls.
func (f *TFlowPanel) FindChildControl(ControlName string) *TControl {
    return AsControl(FlowPanel_FindChildControl(f.instance, ControlName))
}

func (f *TFlowPanel) FlipChildren(AllLevels bool) {
    FlowPanel_FlipChildren(f.instance, AllLevels)
}

// 返回是否获取焦点。
//
// Return to get focus.
func (f *TFlowPanel) Focused() bool {
    return FlowPanel_Focused(f.instance)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (f *TFlowPanel) HandleAllocated() bool {
    return FlowPanel_HandleAllocated(f.instance)
}

// 插入一个控件。
//
// Insert a control.
func (f *TFlowPanel) InsertControl(AControl IControl) {
    FlowPanel_InsertControl(f.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (f *TFlowPanel) Invalidate() {
    FlowPanel_Invalidate(f.instance)
}

// 移除一个控件。
//
// Remove a control.
func (f *TFlowPanel) RemoveControl(AControl IControl) {
    FlowPanel_RemoveControl(f.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (f *TFlowPanel) Realign() {
    FlowPanel_Realign(f.instance)
}

// 重绘。
//
// Repaint.
func (f *TFlowPanel) Repaint() {
    FlowPanel_Repaint(f.instance)
}

// 按比例缩放。
//
// Scale by.
func (f *TFlowPanel) ScaleBy(M int32, D int32) {
    FlowPanel_ScaleBy(f.instance, M , D)
}

// 滚动至指定位置。
//
// Scroll by.
func (f *TFlowPanel) ScrollBy(DeltaX int32, DeltaY int32) {
    FlowPanel_ScrollBy(f.instance, DeltaX , DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (f *TFlowPanel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    FlowPanel_SetBounds(f.instance, ALeft , ATop , AWidth , AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (f *TFlowPanel) SetFocus() {
    FlowPanel_SetFocus(f.instance)
}

// 控件更新。
//
// Update.
func (f *TFlowPanel) Update() {
    FlowPanel_Update(f.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (f *TFlowPanel) BringToFront() {
    FlowPanel_BringToFront(f.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (f *TFlowPanel) ClientToScreen(Point TPoint) TPoint {
    return FlowPanel_ClientToScreen(f.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (f *TFlowPanel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return FlowPanel_ClientToParent(f.instance, Point , CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (f *TFlowPanel) Dragging() bool {
    return FlowPanel_Dragging(f.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (f *TFlowPanel) HasParent() bool {
    return FlowPanel_HasParent(f.instance)
}

// 隐藏控件。
//
// Hidden control.
func (f *TFlowPanel) Hide() {
    FlowPanel_Hide(f.instance)
}

// 发送一个消息。
//
// Send a message.
func (f *TFlowPanel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return FlowPanel_Perform(f.instance, Msg , WParam , LParam)
}

// 刷新控件。
//
// Refresh control.
func (f *TFlowPanel) Refresh() {
    FlowPanel_Refresh(f.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (f *TFlowPanel) ScreenToClient(Point TPoint) TPoint {
    return FlowPanel_ScreenToClient(f.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (f *TFlowPanel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return FlowPanel_ParentToClient(f.instance, Point , CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (f *TFlowPanel) SendToBack() {
    FlowPanel_SendToBack(f.instance)
}

// 显示控件。
//
// Show control.
func (f *TFlowPanel) Show() {
    FlowPanel_Show(f.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (f *TFlowPanel) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return FlowPanel_GetTextBuf(f.instance, Buffer , BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (f *TFlowPanel) GetTextLen() int32 {
    return FlowPanel_GetTextLen(f.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (f *TFlowPanel) SetTextBuf(Buffer string) {
    FlowPanel_SetTextBuf(f.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (f *TFlowPanel) FindComponent(AName string) *TComponent {
    return AsComponent(FlowPanel_FindComponent(f.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (f *TFlowPanel) GetNamePath() string {
    return FlowPanel_GetNamePath(f.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (f *TFlowPanel) Assign(Source IObject) {
    FlowPanel_Assign(f.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (f *TFlowPanel) ClassType() TClass {
    return FlowPanel_ClassType(f.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (f *TFlowPanel) ClassName() string {
    return FlowPanel_ClassName(f.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (f *TFlowPanel) InstanceSize() int32 {
    return FlowPanel_InstanceSize(f.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (f *TFlowPanel) InheritsFrom(AClass TClass) bool {
    return FlowPanel_InheritsFrom(f.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (f *TFlowPanel) Equals(Obj IObject) bool {
    return FlowPanel_Equals(f.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (f *TFlowPanel) GetHashCode() int32 {
    return FlowPanel_GetHashCode(f.instance)
}

// 文本类信息。
//
// Text information.
func (f *TFlowPanel) ToString() string {
    return FlowPanel_ToString(f.instance)
}

func (f *TFlowPanel) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    FlowPanel_AnchorToNeighbour(f.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (f *TFlowPanel) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    FlowPanel_AnchorParallel(f.instance, ASide , ASpace , CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (f *TFlowPanel) AnchorHorizontalCenterTo(ASibling IControl) {
    FlowPanel_AnchorHorizontalCenterTo(f.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (f *TFlowPanel) AnchorVerticalCenterTo(ASibling IControl) {
    FlowPanel_AnchorVerticalCenterTo(f.instance, CheckPtr(ASibling))
}

func (f *TFlowPanel) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    FlowPanel_AnchorAsAlign(f.instance, ATheAlign , ASpace)
}

func (f *TFlowPanel) AnchorClient(ASpace int32) {
    FlowPanel_AnchorClient(f.instance, ASpace)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (f *TFlowPanel) Align() TAlign {
    return FlowPanel_GetAlign(f.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (f *TFlowPanel) SetAlign(value TAlign) {
    FlowPanel_SetAlign(f.instance, value)
}

// 获取文字对齐。
//
// Get Text alignment.
func (f *TFlowPanel) Alignment() TAlignment {
    return FlowPanel_GetAlignment(f.instance)
}

// 设置文字对齐。
//
// Set Text alignment.
func (f *TFlowPanel) SetAlignment(value TAlignment) {
    FlowPanel_SetAlignment(f.instance, value)
}

// 获取四个角位置的锚点。
func (f *TFlowPanel) Anchors() TAnchors {
    return FlowPanel_GetAnchors(f.instance)
}

// 设置四个角位置的锚点。
func (f *TFlowPanel) SetAnchors(value TAnchors) {
    FlowPanel_SetAnchors(f.instance, value)
}

// 获取自动调整大小。
func (f *TFlowPanel) AutoSize() bool {
    return FlowPanel_GetAutoSize(f.instance)
}

// 设置自动调整大小。
func (f *TFlowPanel) SetAutoSize(value bool) {
    FlowPanel_SetAutoSize(f.instance, value)
}

func (f *TFlowPanel) AutoWrap() bool {
    return FlowPanel_GetAutoWrap(f.instance)
}

func (f *TFlowPanel) SetAutoWrap(value bool) {
    FlowPanel_SetAutoWrap(f.instance, value)
}

func (f *TFlowPanel) BiDiMode() TBiDiMode {
    return FlowPanel_GetBiDiMode(f.instance)
}

func (f *TFlowPanel) SetBiDiMode(value TBiDiMode) {
    FlowPanel_SetBiDiMode(f.instance, value)
}

// 获取边框的宽度。
func (f *TFlowPanel) BorderWidth() int32 {
    return FlowPanel_GetBorderWidth(f.instance)
}

// 设置边框的宽度。
func (f *TFlowPanel) SetBorderWidth(value int32) {
    FlowPanel_SetBorderWidth(f.instance, value)
}

// 获取窗口边框样式。比如：无边框，单一边框等。
func (f *TFlowPanel) BorderStyle() TBorderStyle {
    return FlowPanel_GetBorderStyle(f.instance)
}

// 设置窗口边框样式。比如：无边框，单一边框等。
func (f *TFlowPanel) SetBorderStyle(value TBorderStyle) {
    FlowPanel_SetBorderStyle(f.instance, value)
}

// 获取控件标题。
//
// Get the control title.
func (f *TFlowPanel) Caption() string {
    return FlowPanel_GetCaption(f.instance)
}

// 设置控件标题。
//
// Set the control title.
func (f *TFlowPanel) SetCaption(value string) {
    FlowPanel_SetCaption(f.instance, value)
}

// 获取颜色。
//
// Get color.
func (f *TFlowPanel) Color() TColor {
    return FlowPanel_GetColor(f.instance)
}

// 设置颜色。
//
// Set color.
func (f *TFlowPanel) SetColor(value TColor) {
    FlowPanel_SetColor(f.instance, value)
}

// 获取约束控件大小。
func (f *TFlowPanel) Constraints() *TSizeConstraints {
    return AsSizeConstraints(FlowPanel_GetConstraints(f.instance))
}

// 设置约束控件大小。
func (f *TFlowPanel) SetConstraints(value *TSizeConstraints) {
    FlowPanel_SetConstraints(f.instance, CheckPtr(value))
}

// 获取使用停靠管理。
func (f *TFlowPanel) UseDockManager() bool {
    return FlowPanel_GetUseDockManager(f.instance)
}

// 设置使用停靠管理。
func (f *TFlowPanel) SetUseDockManager(value bool) {
    FlowPanel_SetUseDockManager(f.instance, value)
}

// 获取停靠站点。
//
// Get Docking site.
func (f *TFlowPanel) DockSite() bool {
    return FlowPanel_GetDockSite(f.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (f *TFlowPanel) SetDockSite(value bool) {
    FlowPanel_SetDockSite(f.instance, value)
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (f *TFlowPanel) DoubleBuffered() bool {
    return FlowPanel_GetDoubleBuffered(f.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (f *TFlowPanel) SetDoubleBuffered(value bool) {
    FlowPanel_SetDoubleBuffered(f.instance, value)
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (f *TFlowPanel) DragCursor() TCursor {
    return FlowPanel_GetDragCursor(f.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (f *TFlowPanel) SetDragCursor(value TCursor) {
    FlowPanel_SetDragCursor(f.instance, value)
}

// 获取拖拽方式。
//
// Get Drag and drop.
func (f *TFlowPanel) DragKind() TDragKind {
    return FlowPanel_GetDragKind(f.instance)
}

// 设置拖拽方式。
//
// Set Drag and drop.
func (f *TFlowPanel) SetDragKind(value TDragKind) {
    FlowPanel_SetDragKind(f.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (f *TFlowPanel) DragMode() TDragMode {
    return FlowPanel_GetDragMode(f.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (f *TFlowPanel) SetDragMode(value TDragMode) {
    FlowPanel_SetDragMode(f.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (f *TFlowPanel) Enabled() bool {
    return FlowPanel_GetEnabled(f.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (f *TFlowPanel) SetEnabled(value bool) {
    FlowPanel_SetEnabled(f.instance, value)
}

func (f *TFlowPanel) FlowStyle() TFlowStyle {
    return FlowPanel_GetFlowStyle(f.instance)
}

func (f *TFlowPanel) SetFlowStyle(value TFlowStyle) {
    FlowPanel_SetFlowStyle(f.instance, value)
}

func (f *TFlowPanel) FullRepaint() bool {
    return FlowPanel_GetFullRepaint(f.instance)
}

func (f *TFlowPanel) SetFullRepaint(value bool) {
    FlowPanel_SetFullRepaint(f.instance, value)
}

// 获取字体。
//
// Get Font.
func (f *TFlowPanel) Font() *TFont {
    return AsFont(FlowPanel_GetFont(f.instance))
}

// 设置字体。
//
// Set Font.
func (f *TFlowPanel) SetFont(value *TFont) {
    FlowPanel_SetFont(f.instance, CheckPtr(value))
}

func (f *TFlowPanel) ParentBackground() bool {
    return FlowPanel_GetParentBackground(f.instance)
}

func (f *TFlowPanel) SetParentBackground(value bool) {
    FlowPanel_SetParentBackground(f.instance, value)
}

// 获取使用父容器颜色。
//
// Get parent color.
func (f *TFlowPanel) ParentColor() bool {
    return FlowPanel_GetParentColor(f.instance)
}

// 设置使用父容器颜色。
//
// Set parent color.
func (f *TFlowPanel) SetParentColor(value bool) {
    FlowPanel_SetParentColor(f.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (f *TFlowPanel) ParentDoubleBuffered() bool {
    return FlowPanel_GetParentDoubleBuffered(f.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (f *TFlowPanel) SetParentDoubleBuffered(value bool) {
    FlowPanel_SetParentDoubleBuffered(f.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (f *TFlowPanel) ParentFont() bool {
    return FlowPanel_GetParentFont(f.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (f *TFlowPanel) SetParentFont(value bool) {
    FlowPanel_SetParentFont(f.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (f *TFlowPanel) ParentShowHint() bool {
    return FlowPanel_GetParentShowHint(f.instance)
}

// 设置以父容器的ShowHint属性为准。
func (f *TFlowPanel) SetParentShowHint(value bool) {
    FlowPanel_SetParentShowHint(f.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (f *TFlowPanel) PopupMenu() *TPopupMenu {
    return AsPopupMenu(FlowPanel_GetPopupMenu(f.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (f *TFlowPanel) SetPopupMenu(value IComponent) {
    FlowPanel_SetPopupMenu(f.instance, CheckPtr(value))
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (f *TFlowPanel) ShowHint() bool {
    return FlowPanel_GetShowHint(f.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (f *TFlowPanel) SetShowHint(value bool) {
    FlowPanel_SetShowHint(f.instance, value)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (f *TFlowPanel) TabOrder() TTabOrder {
    return FlowPanel_GetTabOrder(f.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (f *TFlowPanel) SetTabOrder(value TTabOrder) {
    FlowPanel_SetTabOrder(f.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (f *TFlowPanel) TabStop() bool {
    return FlowPanel_GetTabStop(f.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (f *TFlowPanel) SetTabStop(value bool) {
    FlowPanel_SetTabStop(f.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (f *TFlowPanel) Visible() bool {
    return FlowPanel_GetVisible(f.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (f *TFlowPanel) SetVisible(value bool) {
    FlowPanel_SetVisible(f.instance, value)
}

// 设置对齐位置事件，当Align为alCustom时Parent会收到这个消息。
func (f *TFlowPanel) SetOnAlignPosition(fn TAlignPositionEvent) {
    FlowPanel_SetOnAlignPosition(f.instance, fn)
}

// 设置控件单击事件。
//
// Set control click event.
func (f *TFlowPanel) SetOnClick(fn TNotifyEvent) {
    FlowPanel_SetOnClick(f.instance, fn)
}

// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (f *TFlowPanel) SetOnContextPopup(fn TContextPopupEvent) {
    FlowPanel_SetOnContextPopup(f.instance, fn)
}

func (f *TFlowPanel) SetOnDockDrop(fn TDockDropEvent) {
    FlowPanel_SetOnDockDrop(f.instance, fn)
}

// 设置双击事件。
func (f *TFlowPanel) SetOnDblClick(fn TNotifyEvent) {
    FlowPanel_SetOnDblClick(f.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (f *TFlowPanel) SetOnDragDrop(fn TDragDropEvent) {
    FlowPanel_SetOnDragDrop(f.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (f *TFlowPanel) SetOnDragOver(fn TDragOverEvent) {
    FlowPanel_SetOnDragOver(f.instance, fn)
}

// 设置停靠结束事件。
//
// Set Dock end event.
func (f *TFlowPanel) SetOnEndDock(fn TEndDragEvent) {
    FlowPanel_SetOnEndDock(f.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (f *TFlowPanel) SetOnEndDrag(fn TEndDragEvent) {
    FlowPanel_SetOnEndDrag(f.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (f *TFlowPanel) SetOnEnter(fn TNotifyEvent) {
    FlowPanel_SetOnEnter(f.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (f *TFlowPanel) SetOnExit(fn TNotifyEvent) {
    FlowPanel_SetOnExit(f.instance, fn)
}

func (f *TFlowPanel) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    FlowPanel_SetOnGetSiteInfo(f.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (f *TFlowPanel) SetOnMouseDown(fn TMouseEvent) {
    FlowPanel_SetOnMouseDown(f.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (f *TFlowPanel) SetOnMouseEnter(fn TNotifyEvent) {
    FlowPanel_SetOnMouseEnter(f.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (f *TFlowPanel) SetOnMouseLeave(fn TNotifyEvent) {
    FlowPanel_SetOnMouseLeave(f.instance, fn)
}

// 设置鼠标移动事件。
func (f *TFlowPanel) SetOnMouseMove(fn TMouseMoveEvent) {
    FlowPanel_SetOnMouseMove(f.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (f *TFlowPanel) SetOnMouseUp(fn TMouseEvent) {
    FlowPanel_SetOnMouseUp(f.instance, fn)
}

// 设置大小被改变事件。
func (f *TFlowPanel) SetOnResize(fn TNotifyEvent) {
    FlowPanel_SetOnResize(f.instance, fn)
}

// 设置启动停靠。
func (f *TFlowPanel) SetOnStartDock(fn TStartDockEvent) {
    FlowPanel_SetOnStartDock(f.instance, fn)
}

func (f *TFlowPanel) SetOnUnDock(fn TUnDockEvent) {
    FlowPanel_SetOnUnDock(f.instance, fn)
}

// 获取依靠客户端总数。
func (f *TFlowPanel) DockClientCount() int32 {
    return FlowPanel_GetDockClientCount(f.instance)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (f *TFlowPanel) MouseInClient() bool {
    return FlowPanel_GetMouseInClient(f.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (f *TFlowPanel) VisibleDockClientCount() int32 {
    return FlowPanel_GetVisibleDockClientCount(f.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (f *TFlowPanel) Brush() *TBrush {
    return AsBrush(FlowPanel_GetBrush(f.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (f *TFlowPanel) ControlCount() int32 {
    return FlowPanel_GetControlCount(f.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (f *TFlowPanel) Handle() HWND {
    return FlowPanel_GetHandle(f.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (f *TFlowPanel) ParentWindow() HWND {
    return FlowPanel_GetParentWindow(f.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (f *TFlowPanel) SetParentWindow(value HWND) {
    FlowPanel_SetParentWindow(f.instance, value)
}

func (f *TFlowPanel) Showing() bool {
    return FlowPanel_GetShowing(f.instance)
}

func (f *TFlowPanel) Action() *TAction {
    return AsAction(FlowPanel_GetAction(f.instance))
}

func (f *TFlowPanel) SetAction(value IComponent) {
    FlowPanel_SetAction(f.instance, CheckPtr(value))
}

func (f *TFlowPanel) BoundsRect() TRect {
    return FlowPanel_GetBoundsRect(f.instance)
}

func (f *TFlowPanel) SetBoundsRect(value TRect) {
    FlowPanel_SetBoundsRect(f.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (f *TFlowPanel) ClientHeight() int32 {
    return FlowPanel_GetClientHeight(f.instance)
}

// 设置客户区高度。
//
// Set client height.
func (f *TFlowPanel) SetClientHeight(value int32) {
    FlowPanel_SetClientHeight(f.instance, value)
}

func (f *TFlowPanel) ClientOrigin() TPoint {
    return FlowPanel_GetClientOrigin(f.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (f *TFlowPanel) ClientRect() TRect {
    return FlowPanel_GetClientRect(f.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (f *TFlowPanel) ClientWidth() int32 {
    return FlowPanel_GetClientWidth(f.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (f *TFlowPanel) SetClientWidth(value int32) {
    FlowPanel_SetClientWidth(f.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (f *TFlowPanel) ControlState() TControlState {
    return FlowPanel_GetControlState(f.instance)
}

// 设置控件状态。
//
// Set control state.
func (f *TFlowPanel) SetControlState(value TControlState) {
    FlowPanel_SetControlState(f.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (f *TFlowPanel) ControlStyle() TControlStyle {
    return FlowPanel_GetControlStyle(f.instance)
}

// 设置控件样式。
//
// Set control style.
func (f *TFlowPanel) SetControlStyle(value TControlStyle) {
    FlowPanel_SetControlStyle(f.instance, value)
}

func (f *TFlowPanel) Floating() bool {
    return FlowPanel_GetFloating(f.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (f *TFlowPanel) Parent() *TWinControl {
    return AsWinControl(FlowPanel_GetParent(f.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (f *TFlowPanel) SetParent(value IWinControl) {
    FlowPanel_SetParent(f.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (f *TFlowPanel) Left() int32 {
    return FlowPanel_GetLeft(f.instance)
}

// 设置左边位置。
//
// Set Left position.
func (f *TFlowPanel) SetLeft(value int32) {
    FlowPanel_SetLeft(f.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (f *TFlowPanel) Top() int32 {
    return FlowPanel_GetTop(f.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (f *TFlowPanel) SetTop(value int32) {
    FlowPanel_SetTop(f.instance, value)
}

// 获取宽度。
//
// Get width.
func (f *TFlowPanel) Width() int32 {
    return FlowPanel_GetWidth(f.instance)
}

// 设置宽度。
//
// Set width.
func (f *TFlowPanel) SetWidth(value int32) {
    FlowPanel_SetWidth(f.instance, value)
}

// 获取高度。
//
// Get height.
func (f *TFlowPanel) Height() int32 {
    return FlowPanel_GetHeight(f.instance)
}

// 设置高度。
//
// Set height.
func (f *TFlowPanel) SetHeight(value int32) {
    FlowPanel_SetHeight(f.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (f *TFlowPanel) Cursor() TCursor {
    return FlowPanel_GetCursor(f.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (f *TFlowPanel) SetCursor(value TCursor) {
    FlowPanel_SetCursor(f.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (f *TFlowPanel) Hint() string {
    return FlowPanel_GetHint(f.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (f *TFlowPanel) SetHint(value string) {
    FlowPanel_SetHint(f.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (f *TFlowPanel) ComponentCount() int32 {
    return FlowPanel_GetComponentCount(f.instance)
}

// 获取组件索引。
//
// Get component index.
func (f *TFlowPanel) ComponentIndex() int32 {
    return FlowPanel_GetComponentIndex(f.instance)
}

// 设置组件索引。
//
// Set component index.
func (f *TFlowPanel) SetComponentIndex(value int32) {
    FlowPanel_SetComponentIndex(f.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (f *TFlowPanel) Owner() *TComponent {
    return AsComponent(FlowPanel_GetOwner(f.instance))
}

// 获取组件名称。
//
// Get the component name.
func (f *TFlowPanel) Name() string {
    return FlowPanel_GetName(f.instance)
}

// 设置组件名称。
//
// Set the component name.
func (f *TFlowPanel) SetName(value string) {
    FlowPanel_SetName(f.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (f *TFlowPanel) Tag() int {
    return FlowPanel_GetTag(f.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (f *TFlowPanel) SetTag(value int) {
    FlowPanel_SetTag(f.instance, value)
}

// 获取左边锚点。
func (f *TFlowPanel) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(FlowPanel_GetAnchorSideLeft(f.instance))
}

// 设置左边锚点。
func (f *TFlowPanel) SetAnchorSideLeft(value *TAnchorSide) {
    FlowPanel_SetAnchorSideLeft(f.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (f *TFlowPanel) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(FlowPanel_GetAnchorSideTop(f.instance))
}

// 设置顶边锚点。
func (f *TFlowPanel) SetAnchorSideTop(value *TAnchorSide) {
    FlowPanel_SetAnchorSideTop(f.instance, CheckPtr(value))
}

// 获取右边锚点。
func (f *TFlowPanel) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(FlowPanel_GetAnchorSideRight(f.instance))
}

// 设置右边锚点。
func (f *TFlowPanel) SetAnchorSideRight(value *TAnchorSide) {
    FlowPanel_SetAnchorSideRight(f.instance, CheckPtr(value))
}

// 获取底边锚点。
func (f *TFlowPanel) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(FlowPanel_GetAnchorSideBottom(f.instance))
}

// 设置底边锚点。
func (f *TFlowPanel) SetAnchorSideBottom(value *TAnchorSide) {
    FlowPanel_SetAnchorSideBottom(f.instance, CheckPtr(value))
}

func (f *TFlowPanel) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(FlowPanel_GetChildSizing(f.instance))
}

func (f *TFlowPanel) SetChildSizing(value *TControlChildSizing) {
    FlowPanel_SetChildSizing(f.instance, CheckPtr(value))
}

// 获取边框间距。
func (f *TFlowPanel) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(FlowPanel_GetBorderSpacing(f.instance))
}

// 设置边框间距。
func (f *TFlowPanel) SetBorderSpacing(value *TControlBorderSpacing) {
    FlowPanel_SetBorderSpacing(f.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (f *TFlowPanel) DockClients(Index int32) *TControl {
    return AsControl(FlowPanel_GetDockClients(f.instance, Index))
}

// 获取指定索引子控件。
func (f *TFlowPanel) Controls(Index int32) *TControl {
    return AsControl(FlowPanel_GetControls(f.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (f *TFlowPanel) Components(AIndex int32) *TComponent {
    return AsComponent(FlowPanel_GetComponents(f.instance, AIndex))
}

// 获取锚侧面。
func (f *TFlowPanel) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(FlowPanel_GetAnchorSide(f.instance, AKind))
}

