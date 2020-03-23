
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

type TThumbBarButtonList struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewThumbBarButtonList() *TThumbBarButtonList {
    t := new(TThumbBarButtonList)
    t.instance = ThumbBarButtonList_Create()
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsThumbBarButtonList(obj interface{}) *TThumbBarButtonList {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TThumbBarButtonList{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsThumbBarButtonList.
func ThumbBarButtonListFromInst(inst uintptr) *TThumbBarButtonList {
    return AsThumbBarButtonList(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsThumbBarButtonList.
func ThumbBarButtonListFromObj(obj IObject) *TThumbBarButtonList {
    return AsThumbBarButtonList(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsThumbBarButtonList.
func ThumbBarButtonListFromUnsafePointer(ptr unsafe.Pointer) *TThumbBarButtonList {
    return AsThumbBarButtonList(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (t *TThumbBarButtonList) Free() {
    if t.instance != 0 {
        ThumbBarButtonList_Free(t.instance)
        t.instance, t.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TThumbBarButtonList) Instance() uintptr {
    return t.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TThumbBarButtonList) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TThumbBarButtonList) IsValid() bool {
    return t.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (t *TThumbBarButtonList) Is() TIs {
    return TIs(t.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (t *TThumbBarButtonList) As() TAs {
//    return TAs(t.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TThumbBarButtonListClass() TClass {
    return ThumbBarButtonList_StaticClassType()
}

func (t *TThumbBarButtonList) Add() *TThumbBarButton {
    return AsThumbBarButton(ThumbBarButtonList_Add(t.instance))
}

// CN: 组件所有者。
// EN: component owner.
func (t *TThumbBarButtonList) Owner() *TObject {
    return AsObject(ThumbBarButtonList_Owner(t.instance))
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TThumbBarButtonList) Assign(Source IObject) {
    ThumbBarButtonList_Assign(t.instance, CheckPtr(Source))
}

func (t *TThumbBarButtonList) BeginUpdate() {
    ThumbBarButtonList_BeginUpdate(t.instance)
}

// CN: 清除。
// EN: .
func (t *TThumbBarButtonList) Clear() {
    ThumbBarButtonList_Clear(t.instance)
}

func (t *TThumbBarButtonList) ClearAndResetID() {
    ThumbBarButtonList_ClearAndResetID(t.instance)
}

func (t *TThumbBarButtonList) Delete(Index int32) {
    ThumbBarButtonList_Delete(t.instance, Index)
}

func (t *TThumbBarButtonList) EndUpdate() {
    ThumbBarButtonList_EndUpdate(t.instance)
}

func (t *TThumbBarButtonList) FindItemID(ID int32) *TCollectionItem {
    return AsCollectionItem(ThumbBarButtonList_FindItemID(t.instance, ID))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TThumbBarButtonList) GetNamePath() string {
    return ThumbBarButtonList_GetNamePath(t.instance)
}

func (t *TThumbBarButtonList) Insert(Index int32) *TCollectionItem {
    return AsCollectionItem(ThumbBarButtonList_Insert(t.instance, Index))
}

// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TThumbBarButtonList) DisposeOf() {
    ThumbBarButtonList_DisposeOf(t.instance)
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TThumbBarButtonList) ClassType() TClass {
    return ThumbBarButtonList_ClassType(t.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TThumbBarButtonList) ClassName() string {
    return ThumbBarButtonList_ClassName(t.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TThumbBarButtonList) InstanceSize() int32 {
    return ThumbBarButtonList_InstanceSize(t.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TThumbBarButtonList) InheritsFrom(AClass TClass) bool {
    return ThumbBarButtonList_InheritsFrom(t.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TThumbBarButtonList) Equals(Obj IObject) bool {
    return ThumbBarButtonList_Equals(t.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TThumbBarButtonList) GetHashCode() int32 {
    return ThumbBarButtonList_GetHashCode(t.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (t *TThumbBarButtonList) ToString() string {
    return ThumbBarButtonList_ToString(t.instance)
}

// CN: 设置改变事件。
// EN: Set changed event.
func (t *TThumbBarButtonList) SetOnChange(fn TNotifyEvent) {
    ThumbBarButtonList_SetOnChange(t.instance, fn)
}

func (t *TThumbBarButtonList) Capacity() int32 {
    return ThumbBarButtonList_GetCapacity(t.instance)
}

func (t *TThumbBarButtonList) SetCapacity(value int32) {
    ThumbBarButtonList_SetCapacity(t.instance, value)
}

func (t *TThumbBarButtonList) Count() int32 {
    return ThumbBarButtonList_GetCount(t.instance)
}

func (t *TThumbBarButtonList) Items(Index int32) *TThumbBarButton {
    return AsThumbBarButton(ThumbBarButtonList_GetItems(t.instance, Index))
}

func (t *TThumbBarButtonList) SetItems(Index int32, value *TThumbBarButton) {
    ThumbBarButtonList_SetItems(t.instance, Index, CheckPtr(value))
}

