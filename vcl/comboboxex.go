
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

type TComboBoxEx struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewComboBoxEx(owner IComponent) *TComboBoxEx {
    c := new(TComboBoxEx)
    c.instance = ComboBoxEx_Create(CheckPtr(owner))
    c.ptr = unsafe.Pointer(c.instance)
    return c
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsComboBoxEx(obj interface{}) *TComboBoxEx {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TComboBoxEx{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsComboBoxEx.
func ComboBoxExFromInst(inst uintptr) *TComboBoxEx {
    return AsComboBoxEx(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsComboBoxEx.
func ComboBoxExFromObj(obj IObject) *TComboBoxEx {
    return AsComboBoxEx(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsComboBoxEx.
func ComboBoxExFromUnsafePointer(ptr unsafe.Pointer) *TComboBoxEx {
    return AsComboBoxEx(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (c *TComboBoxEx) Free() {
    if c.instance != 0 {
        ComboBoxEx_Free(c.instance)
        c.instance, c.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (c *TComboBoxEx) Instance() uintptr {
    return c.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (c *TComboBoxEx) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (c *TComboBoxEx) IsValid() bool {
    return c.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (c *TComboBoxEx) Is() TIs {
    return TIs(c.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (c *TComboBoxEx) As() TAs {
//    return TAs(c.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TComboBoxExClass() TClass {
    return ComboBoxEx_StaticClassType()
}

// 返回是否获取焦点。
//
// Return to get focus.
func (c *TComboBoxEx) Focused() bool {
    return ComboBoxEx_Focused(c.instance)
}

func (c *TComboBoxEx) AddItem(Item string, AObject IObject) {
    ComboBoxEx_AddItem(c.instance, Item , CheckPtr(AObject))
}

// 清除。
func (c *TComboBoxEx) Clear() {
    ComboBoxEx_Clear(c.instance)
}

// 清除选择。
func (c *TComboBoxEx) ClearSelection() {
    ComboBoxEx_ClearSelection(c.instance)
}

// 删除选择的。
func (c *TComboBoxEx) DeleteSelected() {
    ComboBoxEx_DeleteSelected(c.instance)
}

// 全选。
func (c *TComboBoxEx) SelectAll() {
    ComboBoxEx_SelectAll(c.instance)
}

// 是否可以获得焦点。
func (c *TComboBoxEx) CanFocus() bool {
    return ComboBoxEx_CanFocus(c.instance)
}

// 返回是否包含指定控件。
//
// it's contain a specified control.
func (c *TComboBoxEx) ContainsControl(Control IControl) bool {
    return ComboBoxEx_ContainsControl(c.instance, CheckPtr(Control))
}

// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (c *TComboBoxEx) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(ComboBoxEx_ControlAtPos(c.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// 禁用控件的对齐。
//
// Disable control alignment.
func (c *TComboBoxEx) DisableAlign() {
    ComboBoxEx_DisableAlign(c.instance)
}

// 启用控件对齐。
//
// Enabled control alignment.
func (c *TComboBoxEx) EnableAlign() {
    ComboBoxEx_EnableAlign(c.instance)
}

// 查找子控件。
//
// Find sub controls.
func (c *TComboBoxEx) FindChildControl(ControlName string) *TControl {
    return AsControl(ComboBoxEx_FindChildControl(c.instance, ControlName))
}

func (c *TComboBoxEx) FlipChildren(AllLevels bool) {
    ComboBoxEx_FlipChildren(c.instance, AllLevels)
}

// 句柄是否已经分配。
//
// Is the handle already allocated.
func (c *TComboBoxEx) HandleAllocated() bool {
    return ComboBoxEx_HandleAllocated(c.instance)
}

// 插入一个控件。
//
// Insert a control.
func (c *TComboBoxEx) InsertControl(AControl IControl) {
    ComboBoxEx_InsertControl(c.instance, CheckPtr(AControl))
}

// 要求重绘。
//
// Redraw.
func (c *TComboBoxEx) Invalidate() {
    ComboBoxEx_Invalidate(c.instance)
}

// 绘画至指定DC。
//
// Painting to the specified DC.
func (c *TComboBoxEx) PaintTo(DC HDC, X int32, Y int32) {
    ComboBoxEx_PaintTo(c.instance, DC , X , Y)
}

// 移除一个控件。
//
// Remove a control.
func (c *TComboBoxEx) RemoveControl(AControl IControl) {
    ComboBoxEx_RemoveControl(c.instance, CheckPtr(AControl))
}

// 重新对齐。
//
// Realign.
func (c *TComboBoxEx) Realign() {
    ComboBoxEx_Realign(c.instance)
}

// 重绘。
//
// Repaint.
func (c *TComboBoxEx) Repaint() {
    ComboBoxEx_Repaint(c.instance)
}

// 按比例缩放。
//
// Scale by.
func (c *TComboBoxEx) ScaleBy(M int32, D int32) {
    ComboBoxEx_ScaleBy(c.instance, M , D)
}

// 滚动至指定位置。
//
// Scroll by.
func (c *TComboBoxEx) ScrollBy(DeltaX int32, DeltaY int32) {
    ComboBoxEx_ScrollBy(c.instance, DeltaX , DeltaY)
}

// 设置组件边界。
//
// Set component boundaries.
func (c *TComboBoxEx) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ComboBoxEx_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight)
}

// 设置控件焦点。
//
// Set control focus.
func (c *TComboBoxEx) SetFocus() {
    ComboBoxEx_SetFocus(c.instance)
}

// 控件更新。
//
// Update.
func (c *TComboBoxEx) Update() {
    ComboBoxEx_Update(c.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (c *TComboBoxEx) BringToFront() {
    ComboBoxEx_BringToFront(c.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (c *TComboBoxEx) ClientToScreen(Point TPoint) TPoint {
    return ComboBoxEx_ClientToScreen(c.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (c *TComboBoxEx) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ComboBoxEx_ClientToParent(c.instance, Point , CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (c *TComboBoxEx) Dragging() bool {
    return ComboBoxEx_Dragging(c.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (c *TComboBoxEx) HasParent() bool {
    return ComboBoxEx_HasParent(c.instance)
}

// 隐藏控件。
//
// Hidden control.
func (c *TComboBoxEx) Hide() {
    ComboBoxEx_Hide(c.instance)
}

// 发送一个消息。
//
// Send a message.
func (c *TComboBoxEx) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ComboBoxEx_Perform(c.instance, Msg , WParam , LParam)
}

// 刷新控件。
//
// Refresh control.
func (c *TComboBoxEx) Refresh() {
    ComboBoxEx_Refresh(c.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (c *TComboBoxEx) ScreenToClient(Point TPoint) TPoint {
    return ComboBoxEx_ScreenToClient(c.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (c *TComboBoxEx) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ComboBoxEx_ParentToClient(c.instance, Point , CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (c *TComboBoxEx) SendToBack() {
    ComboBoxEx_SendToBack(c.instance)
}

// 显示控件。
//
// Show control.
func (c *TComboBoxEx) Show() {
    ComboBoxEx_Show(c.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (c *TComboBoxEx) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return ComboBoxEx_GetTextBuf(c.instance, Buffer , BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (c *TComboBoxEx) GetTextLen() int32 {
    return ComboBoxEx_GetTextLen(c.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (c *TComboBoxEx) SetTextBuf(Buffer string) {
    ComboBoxEx_SetTextBuf(c.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (c *TComboBoxEx) FindComponent(AName string) *TComponent {
    return AsComponent(ComboBoxEx_FindComponent(c.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (c *TComboBoxEx) GetNamePath() string {
    return ComboBoxEx_GetNamePath(c.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (c *TComboBoxEx) Assign(Source IObject) {
    ComboBoxEx_Assign(c.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (c *TComboBoxEx) ClassType() TClass {
    return ComboBoxEx_ClassType(c.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (c *TComboBoxEx) ClassName() string {
    return ComboBoxEx_ClassName(c.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (c *TComboBoxEx) InstanceSize() int32 {
    return ComboBoxEx_InstanceSize(c.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (c *TComboBoxEx) InheritsFrom(AClass TClass) bool {
    return ComboBoxEx_InheritsFrom(c.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (c *TComboBoxEx) Equals(Obj IObject) bool {
    return ComboBoxEx_Equals(c.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (c *TComboBoxEx) GetHashCode() int32 {
    return ComboBoxEx_GetHashCode(c.instance)
}

// 文本类信息。
//
// Text information.
func (c *TComboBoxEx) ToString() string {
    return ComboBoxEx_ToString(c.instance)
}

func (c *TComboBoxEx) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    ComboBoxEx_AnchorToNeighbour(c.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (c *TComboBoxEx) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    ComboBoxEx_AnchorParallel(c.instance, ASide , ASpace , CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (c *TComboBoxEx) AnchorHorizontalCenterTo(ASibling IControl) {
    ComboBoxEx_AnchorHorizontalCenterTo(c.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (c *TComboBoxEx) AnchorVerticalCenterTo(ASibling IControl) {
    ComboBoxEx_AnchorVerticalCenterTo(c.instance, CheckPtr(ASibling))
}

func (c *TComboBoxEx) AnchorSame(ASide TAnchorKind, ASibling IControl) {
    ComboBoxEx_AnchorSame(c.instance, ASide , CheckPtr(ASibling))
}

func (c *TComboBoxEx) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    ComboBoxEx_AnchorAsAlign(c.instance, ATheAlign , ASpace)
}

func (c *TComboBoxEx) AnchorClient(ASpace int32) {
    ComboBoxEx_AnchorClient(c.instance, ASpace)
}

func (c *TComboBoxEx) ScaleDesignToForm(ASize int32) int32 {
    return ComboBoxEx_ScaleDesignToForm(c.instance, ASize)
}

func (c *TComboBoxEx) ScaleFormToDesign(ASize int32) int32 {
    return ComboBoxEx_ScaleFormToDesign(c.instance, ASize)
}

func (c *TComboBoxEx) Scale96ToForm(ASize int32) int32 {
    return ComboBoxEx_Scale96ToForm(c.instance, ASize)
}

func (c *TComboBoxEx) ScaleFormTo96(ASize int32) int32 {
    return ComboBoxEx_ScaleFormTo96(c.instance, ASize)
}

func (c *TComboBoxEx) Scale96ToFont(ASize int32) int32 {
    return ComboBoxEx_Scale96ToFont(c.instance, ASize)
}

func (c *TComboBoxEx) ScaleFontTo96(ASize int32) int32 {
    return ComboBoxEx_ScaleFontTo96(c.instance, ASize)
}

func (c *TComboBoxEx) ScaleScreenToFont(ASize int32) int32 {
    return ComboBoxEx_ScaleScreenToFont(c.instance, ASize)
}

func (c *TComboBoxEx) ScaleFontToScreen(ASize int32) int32 {
    return ComboBoxEx_ScaleFontToScreen(c.instance, ASize)
}

func (c *TComboBoxEx) Scale96ToScreen(ASize int32) int32 {
    return ComboBoxEx_Scale96ToScreen(c.instance, ASize)
}

func (c *TComboBoxEx) ScaleScreenTo96(ASize int32) int32 {
    return ComboBoxEx_ScaleScreenTo96(c.instance, ASize)
}

func (c *TComboBoxEx) AutoAdjustLayout(AMode TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
    ComboBoxEx_AutoAdjustLayout(c.instance, AMode , AFromPPI , AToPPI , AOldFormWidth , ANewFormWidth)
}

func (c *TComboBoxEx) FixDesignFontsPPI(ADesignTimePPI int32) {
    ComboBoxEx_FixDesignFontsPPI(c.instance, ADesignTimePPI)
}

func (c *TComboBoxEx) ScaleFontsPPI(AToPPI int32, AProportion float64) {
    ComboBoxEx_ScaleFontsPPI(c.instance, AToPPI , AProportion)
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (c *TComboBoxEx) Align() TAlign {
    return ComboBoxEx_GetAlign(c.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (c *TComboBoxEx) SetAlign(value TAlign) {
    ComboBoxEx_SetAlign(c.instance, value)
}

func (c *TComboBoxEx) AutoCompleteOptions() TAutoCompleteOptions {
    return ComboBoxEx_GetAutoCompleteOptions(c.instance)
}

func (c *TComboBoxEx) SetAutoCompleteOptions(value TAutoCompleteOptions) {
    ComboBoxEx_SetAutoCompleteOptions(c.instance, value)
}

func (c *TComboBoxEx) ItemsEx() *TComboExItems {
    return AsComboExItems(ComboBoxEx_GetItemsEx(c.instance))
}

func (c *TComboBoxEx) SetItemsEx(value *TComboExItems) {
    ComboBoxEx_SetItemsEx(c.instance, CheckPtr(value))
}

func (c *TComboBoxEx) Style() TComboBoxExStyle {
    return ComboBoxEx_GetStyle(c.instance)
}

func (c *TComboBoxEx) SetStyle(value TComboBoxExStyle) {
    ComboBoxEx_SetStyle(c.instance, value)
}

func (c *TComboBoxEx) StyleEx() TComboBoxExStyles {
    return ComboBoxEx_GetStyleEx(c.instance)
}

func (c *TComboBoxEx) SetStyleEx(value TComboBoxExStyles) {
    ComboBoxEx_SetStyleEx(c.instance, value)
}

func (c *TComboBoxEx) Action() *TAction {
    return AsAction(ComboBoxEx_GetAction(c.instance))
}

func (c *TComboBoxEx) SetAction(value IComponent) {
    ComboBoxEx_SetAction(c.instance, CheckPtr(value))
}

// 获取四个角位置的锚点。
func (c *TComboBoxEx) Anchors() TAnchors {
    return ComboBoxEx_GetAnchors(c.instance)
}

// 设置四个角位置的锚点。
func (c *TComboBoxEx) SetAnchors(value TAnchors) {
    ComboBoxEx_SetAnchors(c.instance, value)
}

func (c *TComboBoxEx) BiDiMode() TBiDiMode {
    return ComboBoxEx_GetBiDiMode(c.instance)
}

func (c *TComboBoxEx) SetBiDiMode(value TBiDiMode) {
    ComboBoxEx_SetBiDiMode(c.instance, value)
}

// 获取颜色。
//
// Get color.
func (c *TComboBoxEx) Color() TColor {
    return ComboBoxEx_GetColor(c.instance)
}

// 设置颜色。
//
// Set color.
func (c *TComboBoxEx) SetColor(value TColor) {
    ComboBoxEx_SetColor(c.instance, value)
}

// 获取约束控件大小。
func (c *TComboBoxEx) Constraints() *TSizeConstraints {
    return AsSizeConstraints(ComboBoxEx_GetConstraints(c.instance))
}

// 设置约束控件大小。
func (c *TComboBoxEx) SetConstraints(value *TSizeConstraints) {
    ComboBoxEx_SetConstraints(c.instance, CheckPtr(value))
}

// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (c *TComboBoxEx) DoubleBuffered() bool {
    return ComboBoxEx_GetDoubleBuffered(c.instance)
}

// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (c *TComboBoxEx) SetDoubleBuffered(value bool) {
    ComboBoxEx_SetDoubleBuffered(c.instance, value)
}

// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (c *TComboBoxEx) DragCursor() TCursor {
    return ComboBoxEx_GetDragCursor(c.instance)
}

// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (c *TComboBoxEx) SetDragCursor(value TCursor) {
    ComboBoxEx_SetDragCursor(c.instance, value)
}

// 获取拖拽方式。
//
// Get Drag and drop.
func (c *TComboBoxEx) DragKind() TDragKind {
    return ComboBoxEx_GetDragKind(c.instance)
}

// 设置拖拽方式。
//
// Set Drag and drop.
func (c *TComboBoxEx) SetDragKind(value TDragKind) {
    ComboBoxEx_SetDragKind(c.instance, value)
}

// 获取拖拽模式。
//
// Get Drag mode.
func (c *TComboBoxEx) DragMode() TDragMode {
    return ComboBoxEx_GetDragMode(c.instance)
}

// 设置拖拽模式。
//
// Set Drag mode.
func (c *TComboBoxEx) SetDragMode(value TDragMode) {
    ComboBoxEx_SetDragMode(c.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (c *TComboBoxEx) Enabled() bool {
    return ComboBoxEx_GetEnabled(c.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (c *TComboBoxEx) SetEnabled(value bool) {
    ComboBoxEx_SetEnabled(c.instance, value)
}

// 获取字体。
//
// Get Font.
func (c *TComboBoxEx) Font() *TFont {
    return AsFont(ComboBoxEx_GetFont(c.instance))
}

// 设置字体。
//
// Set Font.
func (c *TComboBoxEx) SetFont(value *TFont) {
    ComboBoxEx_SetFont(c.instance, CheckPtr(value))
}

func (c *TComboBoxEx) ItemHeight() int32 {
    return ComboBoxEx_GetItemHeight(c.instance)
}

func (c *TComboBoxEx) SetItemHeight(value int32) {
    ComboBoxEx_SetItemHeight(c.instance, value)
}

// 获取最大长度。
func (c *TComboBoxEx) MaxLength() int32 {
    return ComboBoxEx_GetMaxLength(c.instance)
}

// 设置最大长度。
func (c *TComboBoxEx) SetMaxLength(value int32) {
    ComboBoxEx_SetMaxLength(c.instance, value)
}

// 获取使用父容器颜色。
//
// Get parent color.
func (c *TComboBoxEx) ParentColor() bool {
    return ComboBoxEx_GetParentColor(c.instance)
}

// 设置使用父容器颜色。
//
// Set parent color.
func (c *TComboBoxEx) SetParentColor(value bool) {
    ComboBoxEx_SetParentColor(c.instance, value)
}

// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (c *TComboBoxEx) ParentDoubleBuffered() bool {
    return ComboBoxEx_GetParentDoubleBuffered(c.instance)
}

// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (c *TComboBoxEx) SetParentDoubleBuffered(value bool) {
    ComboBoxEx_SetParentDoubleBuffered(c.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (c *TComboBoxEx) ParentFont() bool {
    return ComboBoxEx_GetParentFont(c.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (c *TComboBoxEx) SetParentFont(value bool) {
    ComboBoxEx_SetParentFont(c.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (c *TComboBoxEx) ParentShowHint() bool {
    return ComboBoxEx_GetParentShowHint(c.instance)
}

// 设置以父容器的ShowHint属性为准。
func (c *TComboBoxEx) SetParentShowHint(value bool) {
    ComboBoxEx_SetParentShowHint(c.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (c *TComboBoxEx) PopupMenu() *TPopupMenu {
    return AsPopupMenu(ComboBoxEx_GetPopupMenu(c.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (c *TComboBoxEx) SetPopupMenu(value IComponent) {
    ComboBoxEx_SetPopupMenu(c.instance, CheckPtr(value))
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (c *TComboBoxEx) ShowHint() bool {
    return ComboBoxEx_GetShowHint(c.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (c *TComboBoxEx) SetShowHint(value bool) {
    ComboBoxEx_SetShowHint(c.instance, value)
}

// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (c *TComboBoxEx) TabOrder() TTabOrder {
    return ComboBoxEx_GetTabOrder(c.instance)
}

// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (c *TComboBoxEx) SetTabOrder(value TTabOrder) {
    ComboBoxEx_SetTabOrder(c.instance, value)
}

// 获取Tab可停留。
//
// Get Tab can stay.
func (c *TComboBoxEx) TabStop() bool {
    return ComboBoxEx_GetTabStop(c.instance)
}

// 设置Tab可停留。
//
// Set Tab can stay.
func (c *TComboBoxEx) SetTabStop(value bool) {
    ComboBoxEx_SetTabStop(c.instance, value)
}

// 获取文本。
func (c *TComboBoxEx) Text() string {
    return getControlBufferText(c.GetTextLen, c.GetTextBuf)
}

// 设置文本。
func (c *TComboBoxEx) SetText(value string) {
    ComboBoxEx_SetText(c.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (c *TComboBoxEx) Visible() bool {
    return ComboBoxEx_GetVisible(c.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (c *TComboBoxEx) SetVisible(value bool) {
    ComboBoxEx_SetVisible(c.instance, value)
}

// 设置改变事件。
//
// Set changed event.
func (c *TComboBoxEx) SetOnChange(fn TNotifyEvent) {
    ComboBoxEx_SetOnChange(c.instance, fn)
}

// 设置控件单击事件。
//
// Set control click event.
func (c *TComboBoxEx) SetOnClick(fn TNotifyEvent) {
    ComboBoxEx_SetOnClick(c.instance, fn)
}

// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (c *TComboBoxEx) SetOnContextPopup(fn TContextPopupEvent) {
    ComboBoxEx_SetOnContextPopup(c.instance, fn)
}

// 设置双击事件。
func (c *TComboBoxEx) SetOnDblClick(fn TNotifyEvent) {
    ComboBoxEx_SetOnDblClick(c.instance, fn)
}

// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (c *TComboBoxEx) SetOnDragDrop(fn TDragDropEvent) {
    ComboBoxEx_SetOnDragDrop(c.instance, fn)
}

// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (c *TComboBoxEx) SetOnDragOver(fn TDragOverEvent) {
    ComboBoxEx_SetOnDragOver(c.instance, fn)
}

func (c *TComboBoxEx) SetOnDropDown(fn TNotifyEvent) {
    ComboBoxEx_SetOnDropDown(c.instance, fn)
}

// 设置停靠结束事件。
//
// Set Dock end event.
func (c *TComboBoxEx) SetOnEndDock(fn TEndDragEvent) {
    ComboBoxEx_SetOnEndDock(c.instance, fn)
}

// 设置拖拽结束。
//
// Set End of drag.
func (c *TComboBoxEx) SetOnEndDrag(fn TEndDragEvent) {
    ComboBoxEx_SetOnEndDrag(c.instance, fn)
}

// 设置焦点进入。
//
// Set Focus entry.
func (c *TComboBoxEx) SetOnEnter(fn TNotifyEvent) {
    ComboBoxEx_SetOnEnter(c.instance, fn)
}

// 设置焦点退出。
//
// Set Focus exit.
func (c *TComboBoxEx) SetOnExit(fn TNotifyEvent) {
    ComboBoxEx_SetOnExit(c.instance, fn)
}

// 设置键盘按键按下事件。
//
// Set Keyboard button press event.
func (c *TComboBoxEx) SetOnKeyDown(fn TKeyEvent) {
    ComboBoxEx_SetOnKeyDown(c.instance, fn)
}

// 设置键键下事件。
func (c *TComboBoxEx) SetOnKeyPress(fn TKeyPressEvent) {
    ComboBoxEx_SetOnKeyPress(c.instance, fn)
}

// 设置键盘按键抬起事件。
//
// Set Keyboard button lift event.
func (c *TComboBoxEx) SetOnKeyUp(fn TKeyEvent) {
    ComboBoxEx_SetOnKeyUp(c.instance, fn)
}

// 设置鼠标移动事件。
func (c *TComboBoxEx) SetOnMouseMove(fn TMouseMoveEvent) {
    ComboBoxEx_SetOnMouseMove(c.instance, fn)
}

func (c *TComboBoxEx) SetOnSelect(fn TNotifyEvent) {
    ComboBoxEx_SetOnSelect(c.instance, fn)
}

// 设置启动停靠。
func (c *TComboBoxEx) SetOnStartDock(fn TStartDockEvent) {
    ComboBoxEx_SetOnStartDock(c.instance, fn)
}

// 获取图标索引列表对象。
func (c *TComboBoxEx) Images() *TImageList {
    return AsImageList(ComboBoxEx_GetImages(c.instance))
}

// 设置图标索引列表对象。
func (c *TComboBoxEx) SetImages(value IComponent) {
    ComboBoxEx_SetImages(c.instance, CheckPtr(value))
}

func (c *TComboBoxEx) DropDownCount() int32 {
    return ComboBoxEx_GetDropDownCount(c.instance)
}

func (c *TComboBoxEx) SetDropDownCount(value int32) {
    ComboBoxEx_SetDropDownCount(c.instance, value)
}

// 获取选择的文本。
func (c *TComboBoxEx) SelText() string {
    return ComboBoxEx_GetSelText(c.instance)
}

// 设置选择的文本。
func (c *TComboBoxEx) SetSelText(value string) {
    ComboBoxEx_SetSelText(c.instance, value)
}

// 获取画布。
func (c *TComboBoxEx) Canvas() *TCanvas {
    return AsCanvas(ComboBoxEx_GetCanvas(c.instance))
}

func (c *TComboBoxEx) DroppedDown() bool {
    return ComboBoxEx_GetDroppedDown(c.instance)
}

func (c *TComboBoxEx) SetDroppedDown(value bool) {
    ComboBoxEx_SetDroppedDown(c.instance, value)
}

func (c *TComboBoxEx) Items() *TStrings {
    return AsStrings(ComboBoxEx_GetItems(c.instance))
}

func (c *TComboBoxEx) SetItems(value IStrings) {
    ComboBoxEx_SetItems(c.instance, CheckPtr(value))
}

// 获取选择的长度。
func (c *TComboBoxEx) SelLength() int32 {
    return ComboBoxEx_GetSelLength(c.instance)
}

// 设置选择的长度。
func (c *TComboBoxEx) SetSelLength(value int32) {
    ComboBoxEx_SetSelLength(c.instance, value)
}

// 获取选择的启始位置。
func (c *TComboBoxEx) SelStart() int32 {
    return ComboBoxEx_GetSelStart(c.instance)
}

// 设置选择的启始位置。
func (c *TComboBoxEx) SetSelStart(value int32) {
    ComboBoxEx_SetSelStart(c.instance, value)
}

func (c *TComboBoxEx) ItemIndex() int32 {
    return ComboBoxEx_GetItemIndex(c.instance)
}

func (c *TComboBoxEx) SetItemIndex(value int32) {
    ComboBoxEx_SetItemIndex(c.instance, value)
}

// 获取依靠客户端总数。
func (c *TComboBoxEx) DockClientCount() int32 {
    return ComboBoxEx_GetDockClientCount(c.instance)
}

// 获取停靠站点。
//
// Get Docking site.
func (c *TComboBoxEx) DockSite() bool {
    return ComboBoxEx_GetDockSite(c.instance)
}

// 设置停靠站点。
//
// Set Docking site.
func (c *TComboBoxEx) SetDockSite(value bool) {
    ComboBoxEx_SetDockSite(c.instance, value)
}

// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (c *TComboBoxEx) MouseInClient() bool {
    return ComboBoxEx_GetMouseInClient(c.instance)
}

// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (c *TComboBoxEx) VisibleDockClientCount() int32 {
    return ComboBoxEx_GetVisibleDockClientCount(c.instance)
}

// 获取画刷对象。
//
// Get Brush.
func (c *TComboBoxEx) Brush() *TBrush {
    return AsBrush(ComboBoxEx_GetBrush(c.instance))
}

// 获取子控件数。
//
// Get Number of child controls.
func (c *TComboBoxEx) ControlCount() int32 {
    return ComboBoxEx_GetControlCount(c.instance)
}

// 获取控件句柄。
//
// Get Control handle.
func (c *TComboBoxEx) Handle() HWND {
    return ComboBoxEx_GetHandle(c.instance)
}

// 获取父容器句柄。
//
// Get Parent container handle.
func (c *TComboBoxEx) ParentWindow() HWND {
    return ComboBoxEx_GetParentWindow(c.instance)
}

// 设置父容器句柄。
//
// Set Parent container handle.
func (c *TComboBoxEx) SetParentWindow(value HWND) {
    ComboBoxEx_SetParentWindow(c.instance, value)
}

func (c *TComboBoxEx) Showing() bool {
    return ComboBoxEx_GetShowing(c.instance)
}

// 获取使用停靠管理。
func (c *TComboBoxEx) UseDockManager() bool {
    return ComboBoxEx_GetUseDockManager(c.instance)
}

// 设置使用停靠管理。
func (c *TComboBoxEx) SetUseDockManager(value bool) {
    ComboBoxEx_SetUseDockManager(c.instance, value)
}

func (c *TComboBoxEx) BoundsRect() TRect {
    return ComboBoxEx_GetBoundsRect(c.instance)
}

func (c *TComboBoxEx) SetBoundsRect(value TRect) {
    ComboBoxEx_SetBoundsRect(c.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (c *TComboBoxEx) ClientHeight() int32 {
    return ComboBoxEx_GetClientHeight(c.instance)
}

// 设置客户区高度。
//
// Set client height.
func (c *TComboBoxEx) SetClientHeight(value int32) {
    ComboBoxEx_SetClientHeight(c.instance, value)
}

func (c *TComboBoxEx) ClientOrigin() TPoint {
    return ComboBoxEx_GetClientOrigin(c.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (c *TComboBoxEx) ClientRect() TRect {
    return ComboBoxEx_GetClientRect(c.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (c *TComboBoxEx) ClientWidth() int32 {
    return ComboBoxEx_GetClientWidth(c.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (c *TComboBoxEx) SetClientWidth(value int32) {
    ComboBoxEx_SetClientWidth(c.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (c *TComboBoxEx) ControlState() TControlState {
    return ComboBoxEx_GetControlState(c.instance)
}

// 设置控件状态。
//
// Set control state.
func (c *TComboBoxEx) SetControlState(value TControlState) {
    ComboBoxEx_SetControlState(c.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (c *TComboBoxEx) ControlStyle() TControlStyle {
    return ComboBoxEx_GetControlStyle(c.instance)
}

// 设置控件样式。
//
// Set control style.
func (c *TComboBoxEx) SetControlStyle(value TControlStyle) {
    ComboBoxEx_SetControlStyle(c.instance, value)
}

func (c *TComboBoxEx) Floating() bool {
    return ComboBoxEx_GetFloating(c.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (c *TComboBoxEx) Parent() *TWinControl {
    return AsWinControl(ComboBoxEx_GetParent(c.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (c *TComboBoxEx) SetParent(value IWinControl) {
    ComboBoxEx_SetParent(c.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (c *TComboBoxEx) Left() int32 {
    return ComboBoxEx_GetLeft(c.instance)
}

// 设置左边位置。
//
// Set Left position.
func (c *TComboBoxEx) SetLeft(value int32) {
    ComboBoxEx_SetLeft(c.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (c *TComboBoxEx) Top() int32 {
    return ComboBoxEx_GetTop(c.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (c *TComboBoxEx) SetTop(value int32) {
    ComboBoxEx_SetTop(c.instance, value)
}

// 获取宽度。
//
// Get width.
func (c *TComboBoxEx) Width() int32 {
    return ComboBoxEx_GetWidth(c.instance)
}

// 设置宽度。
//
// Set width.
func (c *TComboBoxEx) SetWidth(value int32) {
    ComboBoxEx_SetWidth(c.instance, value)
}

// 获取高度。
//
// Get height.
func (c *TComboBoxEx) Height() int32 {
    return ComboBoxEx_GetHeight(c.instance)
}

// 设置高度。
//
// Set height.
func (c *TComboBoxEx) SetHeight(value int32) {
    ComboBoxEx_SetHeight(c.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (c *TComboBoxEx) Cursor() TCursor {
    return ComboBoxEx_GetCursor(c.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (c *TComboBoxEx) SetCursor(value TCursor) {
    ComboBoxEx_SetCursor(c.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (c *TComboBoxEx) Hint() string {
    return ComboBoxEx_GetHint(c.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (c *TComboBoxEx) SetHint(value string) {
    ComboBoxEx_SetHint(c.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (c *TComboBoxEx) ComponentCount() int32 {
    return ComboBoxEx_GetComponentCount(c.instance)
}

// 获取组件索引。
//
// Get component index.
func (c *TComboBoxEx) ComponentIndex() int32 {
    return ComboBoxEx_GetComponentIndex(c.instance)
}

// 设置组件索引。
//
// Set component index.
func (c *TComboBoxEx) SetComponentIndex(value int32) {
    ComboBoxEx_SetComponentIndex(c.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (c *TComboBoxEx) Owner() *TComponent {
    return AsComponent(ComboBoxEx_GetOwner(c.instance))
}

// 获取组件名称。
//
// Get the component name.
func (c *TComboBoxEx) Name() string {
    return ComboBoxEx_GetName(c.instance)
}

// 设置组件名称。
//
// Set the component name.
func (c *TComboBoxEx) SetName(value string) {
    ComboBoxEx_SetName(c.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (c *TComboBoxEx) Tag() int {
    return ComboBoxEx_GetTag(c.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (c *TComboBoxEx) SetTag(value int) {
    ComboBoxEx_SetTag(c.instance, value)
}

// 获取左边锚点。
func (c *TComboBoxEx) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(ComboBoxEx_GetAnchorSideLeft(c.instance))
}

// 设置左边锚点。
func (c *TComboBoxEx) SetAnchorSideLeft(value *TAnchorSide) {
    ComboBoxEx_SetAnchorSideLeft(c.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (c *TComboBoxEx) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(ComboBoxEx_GetAnchorSideTop(c.instance))
}

// 设置顶边锚点。
func (c *TComboBoxEx) SetAnchorSideTop(value *TAnchorSide) {
    ComboBoxEx_SetAnchorSideTop(c.instance, CheckPtr(value))
}

// 获取右边锚点。
func (c *TComboBoxEx) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(ComboBoxEx_GetAnchorSideRight(c.instance))
}

// 设置右边锚点。
func (c *TComboBoxEx) SetAnchorSideRight(value *TAnchorSide) {
    ComboBoxEx_SetAnchorSideRight(c.instance, CheckPtr(value))
}

// 获取底边锚点。
func (c *TComboBoxEx) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(ComboBoxEx_GetAnchorSideBottom(c.instance))
}

// 设置底边锚点。
func (c *TComboBoxEx) SetAnchorSideBottom(value *TAnchorSide) {
    ComboBoxEx_SetAnchorSideBottom(c.instance, CheckPtr(value))
}

func (c *TComboBoxEx) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(ComboBoxEx_GetChildSizing(c.instance))
}

func (c *TComboBoxEx) SetChildSizing(value *TControlChildSizing) {
    ComboBoxEx_SetChildSizing(c.instance, CheckPtr(value))
}

// 获取边框间距。
func (c *TComboBoxEx) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(ComboBoxEx_GetBorderSpacing(c.instance))
}

// 设置边框间距。
func (c *TComboBoxEx) SetBorderSpacing(value *TControlBorderSpacing) {
    ComboBoxEx_SetBorderSpacing(c.instance, CheckPtr(value))
}

// 获取指定索引停靠客户端。
func (c *TComboBoxEx) DockClients(Index int32) *TControl {
    return AsControl(ComboBoxEx_GetDockClients(c.instance, Index))
}

// 获取指定索引子控件。
func (c *TComboBoxEx) Controls(Index int32) *TControl {
    return AsControl(ComboBoxEx_GetControls(c.instance, Index))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (c *TComboBoxEx) Components(AIndex int32) *TComponent {
    return AsComponent(ComboBoxEx_GetComponents(c.instance, AIndex))
}

// 获取锚侧面。
func (c *TComboBoxEx) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(ComboBoxEx_GetAnchorSide(c.instance, AKind))
}

