
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

type TComboExItems struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsComboExItems(obj interface{}) *TComboExItems {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TComboExItems{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsComboExItems.
func ComboExItemsFromInst(inst uintptr) *TComboExItems {
    return AsComboExItems(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsComboExItems.
func ComboExItemsFromObj(obj IObject) *TComboExItems {
    return AsComboExItems(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsComboExItems.
func ComboExItemsFromUnsafePointer(ptr unsafe.Pointer) *TComboExItems {
    return AsComboExItems(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (c *TComboExItems) Instance() uintptr {
    return c.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (c *TComboExItems) UnsafeAddr() unsafe.Pointer {
    return c.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (c *TComboExItems) IsValid() bool {
    return c.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (c *TComboExItems) Is() TIs {
    return TIs(c.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (c *TComboExItems) As() TAs {
//    return TAs(c.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TComboExItemsClass() TClass {
    return ComboExItems_StaticClassType()
}

func (c *TComboExItems) Add() *TComboExItem {
    return AsComboExItem(ComboExItems_Add(c.instance))
}

func (c *TComboExItems) AddItem(Caption string, ImageIndex int32, SelectedImageIndex int32, OverlayImageIndex int32, Indent int32, Data uintptr) *TComboExItem {
    return AsComboExItem(ComboExItems_AddItem(c.instance, Caption , ImageIndex , SelectedImageIndex , OverlayImageIndex , Indent , Data))
}

func (c *TComboExItems) Insert(Index int32) *TComboExItem {
    return AsComboExItem(ComboExItems_Insert(c.instance, Index))
}

// CN: 组件所有者。
// EN: component owner.
func (c *TComboExItems) Owner() *TObject {
    return AsObject(ComboExItems_Owner(c.instance))
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (c *TComboExItems) Assign(Source IObject) {
    ComboExItems_Assign(c.instance, CheckPtr(Source))
}

func (c *TComboExItems) BeginUpdate() {
    ComboExItems_BeginUpdate(c.instance)
}

// CN: 清除。
// EN: .
func (c *TComboExItems) Clear() {
    ComboExItems_Clear(c.instance)
}

func (c *TComboExItems) ClearAndResetID() {
    ComboExItems_ClearAndResetID(c.instance)
}

func (c *TComboExItems) Delete(Index int32) {
    ComboExItems_Delete(c.instance, Index)
}

func (c *TComboExItems) EndUpdate() {
    ComboExItems_EndUpdate(c.instance)
}

func (c *TComboExItems) FindItemID(ID int32) *TCollectionItem {
    return AsCollectionItem(ComboExItems_FindItemID(c.instance, ID))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (c *TComboExItems) GetNamePath() string {
    return ComboExItems_GetNamePath(c.instance)
}

// CN: 丢弃当前对象。
// EN: Discard the current object.
func (c *TComboExItems) DisposeOf() {
    ComboExItems_DisposeOf(c.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (c *TComboExItems) ClassType() TClass {
    return ComboExItems_ClassType(c.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (c *TComboExItems) ClassName() string {
    return ComboExItems_ClassName(c.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (c *TComboExItems) InstanceSize() int32 {
    return ComboExItems_InstanceSize(c.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (c *TComboExItems) InheritsFrom(AClass TClass) bool {
    return ComboExItems_InheritsFrom(c.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (c *TComboExItems) Equals(Obj IObject) bool {
    return ComboExItems_Equals(c.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (c *TComboExItems) GetHashCode() int32 {
    return ComboExItems_GetHashCode(c.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (c *TComboExItems) ToString() string {
    return ComboExItems_ToString(c.instance)
}

func (c *TComboExItems) Capacity() int32 {
    return ComboExItems_GetCapacity(c.instance)
}

func (c *TComboExItems) SetCapacity(value int32) {
    ComboExItems_SetCapacity(c.instance, value)
}

func (c *TComboExItems) Count() int32 {
    return ComboExItems_GetCount(c.instance)
}

func (c *TComboExItems) ComboItems(Index int32) *TComboExItem {
    return AsComboExItem(ComboExItems_GetComboItems(c.instance, Index))
}

