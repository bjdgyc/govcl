
//----------------------------------------
// The code is automatically generated by the GenlibVcl tool.
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

type TScrollBox struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewScrollBox(owner IComponent) *TScrollBox {
    s := new(TScrollBox)
    s.instance = ScrollBox_Create(CheckPtr(owner))
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsScrollBox(obj interface{}) *TScrollBox {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TScrollBox{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsScrollBox.
func ScrollBoxFromInst(inst uintptr) *TScrollBox {
    return AsScrollBox(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsScrollBox.
func ScrollBoxFromObj(obj IObject) *TScrollBox {
    return AsScrollBox(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsScrollBox.
func ScrollBoxFromUnsafePointer(ptr unsafe.Pointer) *TScrollBox {
    return AsScrollBox(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (s *TScrollBox) Free() {
    if s.instance != 0 {
        ScrollBox_Free(s.instance)
        s.instance, s.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TScrollBox) Instance() uintptr {
    return s.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TScrollBox) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TScrollBox) IsValid() bool {
    return s.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (s *TScrollBox) Is() TIs {
    return TIs(s.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (s *TScrollBox) As() TAs {
//    return TAs(s.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TScrollBoxClass() TClass {
    return ScrollBox_StaticClassType()
}

func (s *TScrollBox) DisableAutoRange() {
    ScrollBox_DisableAutoRange(s.instance)
}

func (s *TScrollBox) EnableAutoRange() {
    ScrollBox_EnableAutoRange(s.instance)
}

func (s *TScrollBox) ScrollInView(AControl IControl) {
    ScrollBox_ScrollInView(s.instance, CheckPtr(AControl))
}

// CN: 是否可以获得焦点。
// EN: .
func (s *TScrollBox) CanFocus() bool {
    return ScrollBox_CanFocus(s.instance)
}

// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (s *TScrollBox) ContainsControl(Control IControl) bool {
    return ScrollBox_ContainsControl(s.instance, CheckPtr(Control))
}

// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (s *TScrollBox) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(ScrollBox_ControlAtPos(s.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (s *TScrollBox) DisableAlign() {
    ScrollBox_DisableAlign(s.instance)
}

// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (s *TScrollBox) EnableAlign() {
    ScrollBox_EnableAlign(s.instance)
}

// CN: 查找子控件。
// EN: Find sub controls.
func (s *TScrollBox) FindChildControl(ControlName string) *TControl {
    return AsControl(ScrollBox_FindChildControl(s.instance, ControlName))
}

func (s *TScrollBox) FlipChildren(AllLevels bool) {
    ScrollBox_FlipChildren(s.instance, AllLevels)
}

// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (s *TScrollBox) Focused() bool {
    return ScrollBox_Focused(s.instance)
}

// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (s *TScrollBox) HandleAllocated() bool {
    return ScrollBox_HandleAllocated(s.instance)
}

// CN: 插入一个控件。
// EN: Insert a control.
func (s *TScrollBox) InsertControl(AControl IControl) {
    ScrollBox_InsertControl(s.instance, CheckPtr(AControl))
}

// CN: 要求重绘。
// EN: Redraw.
func (s *TScrollBox) Invalidate() {
    ScrollBox_Invalidate(s.instance)
}

// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (s *TScrollBox) PaintTo(DC HDC, X int32, Y int32) {
    ScrollBox_PaintTo(s.instance, DC , X , Y)
}

// CN: 移除一个控件。
// EN: Remove a control.
func (s *TScrollBox) RemoveControl(AControl IControl) {
    ScrollBox_RemoveControl(s.instance, CheckPtr(AControl))
}

// CN: 重新对齐。
// EN: Realign.
func (s *TScrollBox) Realign() {
    ScrollBox_Realign(s.instance)
}

// CN: 重绘。
// EN: Repaint.
func (s *TScrollBox) Repaint() {
    ScrollBox_Repaint(s.instance)
}

// CN: 按比例缩放。
// EN: Scale by.
func (s *TScrollBox) ScaleBy(M int32, D int32) {
    ScrollBox_ScaleBy(s.instance, M , D)
}

// CN: 滚动至指定位置。
// EN: Scroll by.
func (s *TScrollBox) ScrollBy(DeltaX int32, DeltaY int32) {
    ScrollBox_ScrollBy(s.instance, DeltaX , DeltaY)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (s *TScrollBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ScrollBox_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

// CN: 设置控件焦点。
// EN: Set control focus.
func (s *TScrollBox) SetFocus() {
    ScrollBox_SetFocus(s.instance)
}

// CN: 控件更新。
// EN: Update.
func (s *TScrollBox) Update() {
    ScrollBox_Update(s.instance)
}

// CN: 更新控件状态。
// EN: Update control status.
func (s *TScrollBox) UpdateControlState() {
    ScrollBox_UpdateControlState(s.instance)
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (s *TScrollBox) BringToFront() {
    ScrollBox_BringToFront(s.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (s *TScrollBox) ClientToScreen(Point TPoint) TPoint {
    return ScrollBox_ClientToScreen(s.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (s *TScrollBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ScrollBox_ClientToParent(s.instance, Point , CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (s *TScrollBox) Dragging() bool {
    return ScrollBox_Dragging(s.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (s *TScrollBox) HasParent() bool {
    return ScrollBox_HasParent(s.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (s *TScrollBox) Hide() {
    ScrollBox_Hide(s.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (s *TScrollBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ScrollBox_Perform(s.instance, Msg , WParam , LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (s *TScrollBox) Refresh() {
    ScrollBox_Refresh(s.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (s *TScrollBox) ScreenToClient(Point TPoint) TPoint {
    return ScrollBox_ScreenToClient(s.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (s *TScrollBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ScrollBox_ParentToClient(s.instance, Point , CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (s *TScrollBox) SendToBack() {
    ScrollBox_SendToBack(s.instance)
}

// CN: 显示控件。
// EN: Show control.
func (s *TScrollBox) Show() {
    ScrollBox_Show(s.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (s *TScrollBox) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return ScrollBox_GetTextBuf(s.instance, Buffer , BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (s *TScrollBox) GetTextLen() int32 {
    return ScrollBox_GetTextLen(s.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (s *TScrollBox) SetTextBuf(Buffer string) {
    ScrollBox_SetTextBuf(s.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (s *TScrollBox) FindComponent(AName string) *TComponent {
    return AsComponent(ScrollBox_FindComponent(s.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (s *TScrollBox) GetNamePath() string {
    return ScrollBox_GetNamePath(s.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (s *TScrollBox) Assign(Source IObject) {
    ScrollBox_Assign(s.instance, CheckPtr(Source))
}

// CN: 丢弃当前对象。
// EN: Discard the current object.
func (s *TScrollBox) DisposeOf() {
    ScrollBox_DisposeOf(s.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TScrollBox) ClassType() TClass {
    return ScrollBox_ClassType(s.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TScrollBox) ClassName() string {
    return ScrollBox_ClassName(s.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TScrollBox) InstanceSize() int32 {
    return ScrollBox_InstanceSize(s.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TScrollBox) InheritsFrom(AClass TClass) bool {
    return ScrollBox_InheritsFrom(s.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TScrollBox) Equals(Obj IObject) bool {
    return ScrollBox_Equals(s.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TScrollBox) GetHashCode() int32 {
    return ScrollBox_GetHashCode(s.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (s *TScrollBox) ToString() string {
    return ScrollBox_ToString(s.instance)
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (s *TScrollBox) Align() TAlign {
    return ScrollBox_GetAlign(s.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (s *TScrollBox) SetAlign(value TAlign) {
    ScrollBox_SetAlign(s.instance, value)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (s *TScrollBox) Anchors() TAnchors {
    return ScrollBox_GetAnchors(s.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (s *TScrollBox) SetAnchors(value TAnchors) {
    ScrollBox_SetAnchors(s.instance, value)
}

func (s *TScrollBox) AutoScroll() bool {
    return ScrollBox_GetAutoScroll(s.instance)
}

func (s *TScrollBox) SetAutoScroll(value bool) {
    ScrollBox_SetAutoScroll(s.instance, value)
}

// CN: 获取自动调整大小。
// EN: .
func (s *TScrollBox) AutoSize() bool {
    return ScrollBox_GetAutoSize(s.instance)
}

// CN: 设置自动调整大小。
// EN: .
func (s *TScrollBox) SetAutoSize(value bool) {
    ScrollBox_SetAutoSize(s.instance, value)
}

func (s *TScrollBox) BevelEdges() TBevelEdges {
    return ScrollBox_GetBevelEdges(s.instance)
}

func (s *TScrollBox) SetBevelEdges(value TBevelEdges) {
    ScrollBox_SetBevelEdges(s.instance, value)
}

func (s *TScrollBox) BevelInner() TBevelCut {
    return ScrollBox_GetBevelInner(s.instance)
}

func (s *TScrollBox) SetBevelInner(value TBevelCut) {
    ScrollBox_SetBevelInner(s.instance, value)
}

func (s *TScrollBox) BevelOuter() TBevelCut {
    return ScrollBox_GetBevelOuter(s.instance)
}

func (s *TScrollBox) SetBevelOuter(value TBevelCut) {
    ScrollBox_SetBevelOuter(s.instance, value)
}

func (s *TScrollBox) BevelKind() TBevelKind {
    return ScrollBox_GetBevelKind(s.instance)
}

func (s *TScrollBox) SetBevelKind(value TBevelKind) {
    ScrollBox_SetBevelKind(s.instance, value)
}

func (s *TScrollBox) BiDiMode() TBiDiMode {
    return ScrollBox_GetBiDiMode(s.instance)
}

func (s *TScrollBox) SetBiDiMode(value TBiDiMode) {
    ScrollBox_SetBiDiMode(s.instance, value)
}

// CN: 获取窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (s *TScrollBox) BorderStyle() TBorderStyle {
    return ScrollBox_GetBorderStyle(s.instance)
}

// CN: 设置窗口边框样式。比如：无边框，单一边框等。
// EN: .
func (s *TScrollBox) SetBorderStyle(value TBorderStyle) {
    ScrollBox_SetBorderStyle(s.instance, value)
}

func (s *TScrollBox) Constraints() *TSizeConstraints {
    return AsSizeConstraints(ScrollBox_GetConstraints(s.instance))
}

func (s *TScrollBox) SetConstraints(value *TSizeConstraints) {
    ScrollBox_SetConstraints(s.instance, CheckPtr(value))
}

// CN: 获取停靠站点。
// EN: Get Docking site.
func (s *TScrollBox) DockSite() bool {
    return ScrollBox_GetDockSite(s.instance)
}

// CN: 设置停靠站点。
// EN: Set Docking site.
func (s *TScrollBox) SetDockSite(value bool) {
    ScrollBox_SetDockSite(s.instance, value)
}

// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (s *TScrollBox) DoubleBuffered() bool {
    return ScrollBox_GetDoubleBuffered(s.instance)
}

// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (s *TScrollBox) SetDoubleBuffered(value bool) {
    ScrollBox_SetDoubleBuffered(s.instance, value)
}

// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (s *TScrollBox) DragCursor() TCursor {
    return ScrollBox_GetDragCursor(s.instance)
}

// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (s *TScrollBox) SetDragCursor(value TCursor) {
    ScrollBox_SetDragCursor(s.instance, value)
}

// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (s *TScrollBox) DragKind() TDragKind {
    return ScrollBox_GetDragKind(s.instance)
}

// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (s *TScrollBox) SetDragKind(value TDragKind) {
    ScrollBox_SetDragKind(s.instance, value)
}

// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (s *TScrollBox) DragMode() TDragMode {
    return ScrollBox_GetDragMode(s.instance)
}

// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (s *TScrollBox) SetDragMode(value TDragMode) {
    ScrollBox_SetDragMode(s.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (s *TScrollBox) Enabled() bool {
    return ScrollBox_GetEnabled(s.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (s *TScrollBox) SetEnabled(value bool) {
    ScrollBox_SetEnabled(s.instance, value)
}

// CN: 获取颜色。
// EN: Get color.
func (s *TScrollBox) Color() TColor {
    return ScrollBox_GetColor(s.instance)
}

// CN: 设置颜色。
// EN: Set color.
func (s *TScrollBox) SetColor(value TColor) {
    ScrollBox_SetColor(s.instance, value)
}

func (s *TScrollBox) Ctl3D() bool {
    return ScrollBox_GetCtl3D(s.instance)
}

func (s *TScrollBox) SetCtl3D(value bool) {
    ScrollBox_SetCtl3D(s.instance, value)
}

// CN: 获取字体。
// EN: Get Font.
func (s *TScrollBox) Font() *TFont {
    return AsFont(ScrollBox_GetFont(s.instance))
}

// CN: 设置字体。
// EN: Set Font.
func (s *TScrollBox) SetFont(value *TFont) {
    ScrollBox_SetFont(s.instance, CheckPtr(value))
}

func (s *TScrollBox) ParentBackground() bool {
    return ScrollBox_GetParentBackground(s.instance)
}

func (s *TScrollBox) SetParentBackground(value bool) {
    ScrollBox_SetParentBackground(s.instance, value)
}

// CN: 获取父容器颜色。
// EN: Get parent color.
func (s *TScrollBox) ParentColor() bool {
    return ScrollBox_GetParentColor(s.instance)
}

// CN: 设置父容器颜色。
// EN: Set parent color.
func (s *TScrollBox) SetParentColor(value bool) {
    ScrollBox_SetParentColor(s.instance, value)
}

func (s *TScrollBox) ParentCtl3D() bool {
    return ScrollBox_GetParentCtl3D(s.instance)
}

func (s *TScrollBox) SetParentCtl3D(value bool) {
    ScrollBox_SetParentCtl3D(s.instance, value)
}

// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (s *TScrollBox) ParentDoubleBuffered() bool {
    return ScrollBox_GetParentDoubleBuffered(s.instance)
}

// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (s *TScrollBox) SetParentDoubleBuffered(value bool) {
    ScrollBox_SetParentDoubleBuffered(s.instance, value)
}

// CN: 获取父容器字体。
// EN: Get Parent container font.
func (s *TScrollBox) ParentFont() bool {
    return ScrollBox_GetParentFont(s.instance)
}

// CN: 设置父容器字体。
// EN: Set Parent container font.
func (s *TScrollBox) SetParentFont(value bool) {
    ScrollBox_SetParentFont(s.instance, value)
}

func (s *TScrollBox) ParentShowHint() bool {
    return ScrollBox_GetParentShowHint(s.instance)
}

func (s *TScrollBox) SetParentShowHint(value bool) {
    ScrollBox_SetParentShowHint(s.instance, value)
}

// CN: 获取右键菜单。
// EN: Get Right click menu.
func (s *TScrollBox) PopupMenu() *TPopupMenu {
    return AsPopupMenu(ScrollBox_GetPopupMenu(s.instance))
}

// CN: 设置右键菜单。
// EN: Set Right click menu.
func (s *TScrollBox) SetPopupMenu(value IComponent) {
    ScrollBox_SetPopupMenu(s.instance, CheckPtr(value))
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (s *TScrollBox) ShowHint() bool {
    return ScrollBox_GetShowHint(s.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (s *TScrollBox) SetShowHint(value bool) {
    ScrollBox_SetShowHint(s.instance, value)
}

// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (s *TScrollBox) TabOrder() TTabOrder {
    return ScrollBox_GetTabOrder(s.instance)
}

// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (s *TScrollBox) SetTabOrder(value TTabOrder) {
    ScrollBox_SetTabOrder(s.instance, value)
}

// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (s *TScrollBox) TabStop() bool {
    return ScrollBox_GetTabStop(s.instance)
}

// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (s *TScrollBox) SetTabStop(value bool) {
    ScrollBox_SetTabStop(s.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (s *TScrollBox) Visible() bool {
    return ScrollBox_GetVisible(s.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (s *TScrollBox) SetVisible(value bool) {
    ScrollBox_SetVisible(s.instance, value)
}

// CN: 获取样式元素。
// EN: Get Style element.
func (s *TScrollBox) StyleElements() TStyleElements {
    return ScrollBox_GetStyleElements(s.instance)
}

// CN: 设置样式元素。
// EN: Set Style element.
func (s *TScrollBox) SetStyleElements(value TStyleElements) {
    ScrollBox_SetStyleElements(s.instance, value)
}

// CN: 设置控件单击事件。
// EN: Set control click event.
func (s *TScrollBox) SetOnClick(fn TNotifyEvent) {
    ScrollBox_SetOnClick(s.instance, fn)
}

// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (s *TScrollBox) SetOnContextPopup(fn TContextPopupEvent) {
    ScrollBox_SetOnContextPopup(s.instance, fn)
}

// CN: 设置双击事件。
// EN: .
func (s *TScrollBox) SetOnDblClick(fn TNotifyEvent) {
    ScrollBox_SetOnDblClick(s.instance, fn)
}

func (s *TScrollBox) SetOnDockDrop(fn TDockDropEvent) {
    ScrollBox_SetOnDockDrop(s.instance, fn)
}

// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (s *TScrollBox) SetOnDragDrop(fn TDragDropEvent) {
    ScrollBox_SetOnDragDrop(s.instance, fn)
}

// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (s *TScrollBox) SetOnDragOver(fn TDragOverEvent) {
    ScrollBox_SetOnDragOver(s.instance, fn)
}

// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (s *TScrollBox) SetOnEndDock(fn TEndDragEvent) {
    ScrollBox_SetOnEndDock(s.instance, fn)
}

// CN: 设置拖拽结束。
// EN: Set End of drag.
func (s *TScrollBox) SetOnEndDrag(fn TEndDragEvent) {
    ScrollBox_SetOnEndDrag(s.instance, fn)
}

// CN: 设置焦点进入。
// EN: Set Focus entry.
func (s *TScrollBox) SetOnEnter(fn TNotifyEvent) {
    ScrollBox_SetOnEnter(s.instance, fn)
}

// CN: 设置焦点退出。
// EN: Set Focus exit.
func (s *TScrollBox) SetOnExit(fn TNotifyEvent) {
    ScrollBox_SetOnExit(s.instance, fn)
}

func (s *TScrollBox) SetOnGesture(fn TGestureEvent) {
    ScrollBox_SetOnGesture(s.instance, fn)
}

func (s *TScrollBox) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    ScrollBox_SetOnGetSiteInfo(s.instance, fn)
}

func (s *TScrollBox) SetOnMouseActivate(fn TMouseActivateEvent) {
    ScrollBox_SetOnMouseActivate(s.instance, fn)
}

// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (s *TScrollBox) SetOnMouseDown(fn TMouseEvent) {
    ScrollBox_SetOnMouseDown(s.instance, fn)
}

// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (s *TScrollBox) SetOnMouseEnter(fn TNotifyEvent) {
    ScrollBox_SetOnMouseEnter(s.instance, fn)
}

// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (s *TScrollBox) SetOnMouseLeave(fn TNotifyEvent) {
    ScrollBox_SetOnMouseLeave(s.instance, fn)
}

// CN: 设置鼠标移动事件。
// EN: .
func (s *TScrollBox) SetOnMouseMove(fn TMouseMoveEvent) {
    ScrollBox_SetOnMouseMove(s.instance, fn)
}

// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (s *TScrollBox) SetOnMouseUp(fn TMouseEvent) {
    ScrollBox_SetOnMouseUp(s.instance, fn)
}

// CN: 设置鼠标滚轮事件。
// EN: .
func (s *TScrollBox) SetOnMouseWheel(fn TMouseWheelEvent) {
    ScrollBox_SetOnMouseWheel(s.instance, fn)
}

// CN: 设置鼠标滚轮按下事件。
// EN: .
func (s *TScrollBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
    ScrollBox_SetOnMouseWheelDown(s.instance, fn)
}

// CN: 设置鼠标滚轮抬起事件。
// EN: .
func (s *TScrollBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
    ScrollBox_SetOnMouseWheelUp(s.instance, fn)
}

// CN: 设置大小被改变事件。
// EN: .
func (s *TScrollBox) SetOnResize(fn TNotifyEvent) {
    ScrollBox_SetOnResize(s.instance, fn)
}

// CN: 设置启动停靠。
// EN: .
func (s *TScrollBox) SetOnStartDock(fn TStartDockEvent) {
    ScrollBox_SetOnStartDock(s.instance, fn)
}

func (s *TScrollBox) SetOnUnDock(fn TUnDockEvent) {
    ScrollBox_SetOnUnDock(s.instance, fn)
}

func (s *TScrollBox) SetOnAlignPosition(fn TAlignPositionEvent) {
    ScrollBox_SetOnAlignPosition(s.instance, fn)
}

func (s *TScrollBox) HorzScrollBar() *TControlScrollBar {
    return AsControlScrollBar(ScrollBox_GetHorzScrollBar(s.instance))
}

func (s *TScrollBox) SetHorzScrollBar(value *TControlScrollBar) {
    ScrollBox_SetHorzScrollBar(s.instance, CheckPtr(value))
}

func (s *TScrollBox) VertScrollBar() *TControlScrollBar {
    return AsControlScrollBar(ScrollBox_GetVertScrollBar(s.instance))
}

func (s *TScrollBox) SetVertScrollBar(value *TControlScrollBar) {
    ScrollBox_SetVertScrollBar(s.instance, CheckPtr(value))
}

// CN: 获取依靠客户端总数。
// EN: .
func (s *TScrollBox) DockClientCount() int32 {
    return ScrollBox_GetDockClientCount(s.instance)
}

// CN: 获取禁用对齐。
// EN: .
func (s *TScrollBox) AlignDisabled() bool {
    return ScrollBox_GetAlignDisabled(s.instance)
}

// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (s *TScrollBox) MouseInClient() bool {
    return ScrollBox_GetMouseInClient(s.instance)
}

// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (s *TScrollBox) VisibleDockClientCount() int32 {
    return ScrollBox_GetVisibleDockClientCount(s.instance)
}

// CN: 获取画刷对象。
// EN: Get Brush.
func (s *TScrollBox) Brush() *TBrush {
    return AsBrush(ScrollBox_GetBrush(s.instance))
}

// CN: 获取子控件数。
// EN: Get Number of child controls.
func (s *TScrollBox) ControlCount() int32 {
    return ScrollBox_GetControlCount(s.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (s *TScrollBox) Handle() HWND {
    return ScrollBox_GetHandle(s.instance)
}

// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (s *TScrollBox) ParentWindow() HWND {
    return ScrollBox_GetParentWindow(s.instance)
}

// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (s *TScrollBox) SetParentWindow(value HWND) {
    ScrollBox_SetParentWindow(s.instance, value)
}

// CN: 获取使用停靠管理。
// EN: .
func (s *TScrollBox) UseDockManager() bool {
    return ScrollBox_GetUseDockManager(s.instance)
}

// CN: 设置使用停靠管理。
// EN: .
func (s *TScrollBox) SetUseDockManager(value bool) {
    ScrollBox_SetUseDockManager(s.instance, value)
}

func (s *TScrollBox) Action() *TAction {
    return AsAction(ScrollBox_GetAction(s.instance))
}

func (s *TScrollBox) SetAction(value IComponent) {
    ScrollBox_SetAction(s.instance, CheckPtr(value))
}

func (s *TScrollBox) BoundsRect() TRect {
    return ScrollBox_GetBoundsRect(s.instance)
}

func (s *TScrollBox) SetBoundsRect(value TRect) {
    ScrollBox_SetBoundsRect(s.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (s *TScrollBox) ClientHeight() int32 {
    return ScrollBox_GetClientHeight(s.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (s *TScrollBox) SetClientHeight(value int32) {
    ScrollBox_SetClientHeight(s.instance, value)
}

func (s *TScrollBox) ClientOrigin() TPoint {
    return ScrollBox_GetClientOrigin(s.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (s *TScrollBox) ClientRect() TRect {
    return ScrollBox_GetClientRect(s.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (s *TScrollBox) ClientWidth() int32 {
    return ScrollBox_GetClientWidth(s.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (s *TScrollBox) SetClientWidth(value int32) {
    ScrollBox_SetClientWidth(s.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (s *TScrollBox) ControlState() TControlState {
    return ScrollBox_GetControlState(s.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (s *TScrollBox) SetControlState(value TControlState) {
    ScrollBox_SetControlState(s.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (s *TScrollBox) ControlStyle() TControlStyle {
    return ScrollBox_GetControlStyle(s.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (s *TScrollBox) SetControlStyle(value TControlStyle) {
    ScrollBox_SetControlStyle(s.instance, value)
}

func (s *TScrollBox) ExplicitLeft() int32 {
    return ScrollBox_GetExplicitLeft(s.instance)
}

func (s *TScrollBox) ExplicitTop() int32 {
    return ScrollBox_GetExplicitTop(s.instance)
}

func (s *TScrollBox) ExplicitWidth() int32 {
    return ScrollBox_GetExplicitWidth(s.instance)
}

func (s *TScrollBox) ExplicitHeight() int32 {
    return ScrollBox_GetExplicitHeight(s.instance)
}

func (s *TScrollBox) Floating() bool {
    return ScrollBox_GetFloating(s.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (s *TScrollBox) Parent() *TWinControl {
    return AsWinControl(ScrollBox_GetParent(s.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (s *TScrollBox) SetParent(value IWinControl) {
    ScrollBox_SetParent(s.instance, CheckPtr(value))
}

// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (s *TScrollBox) AlignWithMargins() bool {
    return ScrollBox_GetAlignWithMargins(s.instance)
}

// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (s *TScrollBox) SetAlignWithMargins(value bool) {
    ScrollBox_SetAlignWithMargins(s.instance, value)
}

// CN: 获取左边位置。
// EN: Get Left position.
func (s *TScrollBox) Left() int32 {
    return ScrollBox_GetLeft(s.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (s *TScrollBox) SetLeft(value int32) {
    ScrollBox_SetLeft(s.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (s *TScrollBox) Top() int32 {
    return ScrollBox_GetTop(s.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (s *TScrollBox) SetTop(value int32) {
    ScrollBox_SetTop(s.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (s *TScrollBox) Width() int32 {
    return ScrollBox_GetWidth(s.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (s *TScrollBox) SetWidth(value int32) {
    ScrollBox_SetWidth(s.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (s *TScrollBox) Height() int32 {
    return ScrollBox_GetHeight(s.instance)
}

// CN: 设置高度。
// EN: Set height.
func (s *TScrollBox) SetHeight(value int32) {
    ScrollBox_SetHeight(s.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (s *TScrollBox) Cursor() TCursor {
    return ScrollBox_GetCursor(s.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (s *TScrollBox) SetCursor(value TCursor) {
    ScrollBox_SetCursor(s.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (s *TScrollBox) Hint() string {
    return ScrollBox_GetHint(s.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (s *TScrollBox) SetHint(value string) {
    ScrollBox_SetHint(s.instance, value)
}

// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (s *TScrollBox) Margins() *TMargins {
    return AsMargins(ScrollBox_GetMargins(s.instance))
}

// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (s *TScrollBox) SetMargins(value *TMargins) {
    ScrollBox_SetMargins(s.instance, CheckPtr(value))
}

// CN: 获取自定义提示。
// EN: Get custom hint.
func (s *TScrollBox) CustomHint() *TCustomHint {
    return AsCustomHint(ScrollBox_GetCustomHint(s.instance))
}

// CN: 设置自定义提示。
// EN: Set custom hint.
func (s *TScrollBox) SetCustomHint(value IComponent) {
    ScrollBox_SetCustomHint(s.instance, CheckPtr(value))
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (s *TScrollBox) ComponentCount() int32 {
    return ScrollBox_GetComponentCount(s.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (s *TScrollBox) ComponentIndex() int32 {
    return ScrollBox_GetComponentIndex(s.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (s *TScrollBox) SetComponentIndex(value int32) {
    ScrollBox_SetComponentIndex(s.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (s *TScrollBox) Owner() *TComponent {
    return AsComponent(ScrollBox_GetOwner(s.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (s *TScrollBox) Name() string {
    return ScrollBox_GetName(s.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (s *TScrollBox) SetName(value string) {
    ScrollBox_SetName(s.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (s *TScrollBox) Tag() int {
    return ScrollBox_GetTag(s.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (s *TScrollBox) SetTag(value int) {
    ScrollBox_SetTag(s.instance, value)
}

// CN: 获取指定索引停靠客户端。
// EN: .
func (s *TScrollBox) DockClients(Index int32) *TControl {
    return AsControl(ScrollBox_GetDockClients(s.instance, Index))
}

// CN: 获取指定索引子控件。
// EN: .
func (s *TScrollBox) Controls(Index int32) *TControl {
    return AsControl(ScrollBox_GetControls(s.instance, Index))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (s *TScrollBox) Components(AIndex int32) *TComponent {
    return AsComponent(ScrollBox_GetComponents(s.instance, AIndex))
}

