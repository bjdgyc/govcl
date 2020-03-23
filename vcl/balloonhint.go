
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

type TBalloonHint struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewBalloonHint(owner IComponent) *TBalloonHint {
    b := new(TBalloonHint)
    b.instance = BalloonHint_Create(CheckPtr(owner))
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsBalloonHint(obj interface{}) *TBalloonHint {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TBalloonHint{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsBalloonHint.
func BalloonHintFromInst(inst uintptr) *TBalloonHint {
    return AsBalloonHint(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsBalloonHint.
func BalloonHintFromObj(obj IObject) *TBalloonHint {
    return AsBalloonHint(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsBalloonHint.
func BalloonHintFromUnsafePointer(ptr unsafe.Pointer) *TBalloonHint {
    return AsBalloonHint(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (b *TBalloonHint) Free() {
    if b.instance != 0 {
        BalloonHint_Free(b.instance)
        b.instance, b.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (b *TBalloonHint) Instance() uintptr {
    return b.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (b *TBalloonHint) UnsafeAddr() unsafe.Pointer {
    return b.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (b *TBalloonHint) IsValid() bool {
    return b.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (b *TBalloonHint) Is() TIs {
    return TIs(b.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (b *TBalloonHint) As() TAs {
//    return TAs(b.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TBalloonHintClass() TClass {
    return BalloonHint_StaticClassType()
}

// CN: 显示鼠标悬停提示。
// EN: Show mouseover tips.
func (b *TBalloonHint) ShowHint() {
    BalloonHint_ShowHint(b.instance)
}

// CN: 隐藏鼠标悬停提示。
// EN: .
func (b *TBalloonHint) HideHint() {
    BalloonHint_HideHint(b.instance)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (b *TBalloonHint) FindComponent(AName string) *TComponent {
    return AsComponent(BalloonHint_FindComponent(b.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (b *TBalloonHint) GetNamePath() string {
    return BalloonHint_GetNamePath(b.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (b *TBalloonHint) HasParent() bool {
    return BalloonHint_HasParent(b.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (b *TBalloonHint) Assign(Source IObject) {
    BalloonHint_Assign(b.instance, CheckPtr(Source))
}

// CN: 丢弃当前对象。
// EN: Discard the current object.
func (b *TBalloonHint) DisposeOf() {
    BalloonHint_DisposeOf(b.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (b *TBalloonHint) ClassType() TClass {
    return BalloonHint_ClassType(b.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (b *TBalloonHint) ClassName() string {
    return BalloonHint_ClassName(b.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (b *TBalloonHint) InstanceSize() int32 {
    return BalloonHint_InstanceSize(b.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (b *TBalloonHint) InheritsFrom(AClass TClass) bool {
    return BalloonHint_InheritsFrom(b.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (b *TBalloonHint) Equals(Obj IObject) bool {
    return BalloonHint_Equals(b.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (b *TBalloonHint) GetHashCode() int32 {
    return BalloonHint_GetHashCode(b.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (b *TBalloonHint) ToString() string {
    return BalloonHint_ToString(b.instance)
}

func (b *TBalloonHint) ShowingHint() bool {
    return BalloonHint_GetShowingHint(b.instance)
}

func (b *TBalloonHint) Title() string {
    return BalloonHint_GetTitle(b.instance)
}

func (b *TBalloonHint) SetTitle(value string) {
    BalloonHint_SetTitle(b.instance, value)
}

func (b *TBalloonHint) Description() string {
    return BalloonHint_GetDescription(b.instance)
}

func (b *TBalloonHint) SetDescription(value string) {
    BalloonHint_SetDescription(b.instance, value)
}

// CN: 获取图像在images中的索引。
// EN: .
func (b *TBalloonHint) ImageIndex() int32 {
    return BalloonHint_GetImageIndex(b.instance)
}

// CN: 设置图像在images中的索引。
// EN: .
func (b *TBalloonHint) SetImageIndex(value int32) {
    BalloonHint_SetImageIndex(b.instance, value)
}

// CN: 获取图标索引列表对象。
// EN: .
func (b *TBalloonHint) Images() *TImageList {
    return AsImageList(BalloonHint_GetImages(b.instance))
}

// CN: 设置图标索引列表对象。
// EN: .
func (b *TBalloonHint) SetImages(value IComponent) {
    BalloonHint_SetImages(b.instance, CheckPtr(value))
}

func (b *TBalloonHint) Style() TBalloonHintStyle {
    return BalloonHint_GetStyle(b.instance)
}

func (b *TBalloonHint) SetStyle(value TBalloonHintStyle) {
    BalloonHint_SetStyle(b.instance, value)
}

func (b *TBalloonHint) Delay() uint32 {
    return BalloonHint_GetDelay(b.instance)
}

func (b *TBalloonHint) SetDelay(value uint32) {
    BalloonHint_SetDelay(b.instance, value)
}

func (b *TBalloonHint) HideAfter() int32 {
    return BalloonHint_GetHideAfter(b.instance)
}

func (b *TBalloonHint) SetHideAfter(value int32) {
    BalloonHint_SetHideAfter(b.instance, value)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (b *TBalloonHint) ComponentCount() int32 {
    return BalloonHint_GetComponentCount(b.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (b *TBalloonHint) ComponentIndex() int32 {
    return BalloonHint_GetComponentIndex(b.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (b *TBalloonHint) SetComponentIndex(value int32) {
    BalloonHint_SetComponentIndex(b.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (b *TBalloonHint) Owner() *TComponent {
    return AsComponent(BalloonHint_GetOwner(b.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (b *TBalloonHint) Name() string {
    return BalloonHint_GetName(b.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (b *TBalloonHint) SetName(value string) {
    BalloonHint_SetName(b.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (b *TBalloonHint) Tag() int {
    return BalloonHint_GetTag(b.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (b *TBalloonHint) SetTag(value int) {
    BalloonHint_SetTag(b.instance, value)
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (b *TBalloonHint) Components(AIndex int32) *TComponent {
    return AsComponent(BalloonHint_GetComponents(b.instance, AIndex))
}

