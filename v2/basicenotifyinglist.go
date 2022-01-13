package ecore

type eNotifyingListInternal[T any] interface {
	ENotifyingList[T]

	inverseAdd(object T, chain ENotificationChain) ENotificationChain

	inverseRemove(object T, chain ENotificationChain) ENotificationChain
}

// BasicENotifyingList ...
type BasicENotifyingList[T comparable] struct {
	basicEList[T]
}

// NewBasicENotifyingList ...
func NewBasicENotifyingList[T comparable]() *BasicENotifyingList[T] {
	l := new(BasicENotifyingList[T])
	l.interfaces = l
	l.data = []T{}
	l.isUnique = true
	return l
}

func newBasicENotifyingListFromData[T comparable](data []T) *BasicENotifyingList[T] {
	l := new(BasicENotifyingList[T])
	l.interfaces = l
	l.data = data
	l.isUnique = true
	return l
}

// GetNotifier ...
func (list *BasicENotifyingList[T]) GetNotifier() ENotifier {
	return nil
}

// GetFeature ...
func (list *BasicENotifyingList[T]) GetFeature() EStructuralFeature {
	return nil
}

// GetFeatureID ...
func (list *BasicENotifyingList[T]) GetFeatureID() int {
	return -1
}

type notifyingListNotification[T comparable] struct {
	AbstractNotification
	list *BasicENotifyingList[T]
}

func (notif *notifyingListNotification[T]) GetNotifier() ENotifier {
	return notif.list.interfaces.(ENotifyingList[T]).GetNotifier()
}

func (notif *notifyingListNotification[T]) GetFeature() EStructuralFeature {
	return notif.list.interfaces.(ENotifyingList[T]).GetFeature()
}

func (notif *notifyingListNotification[T]) GetFeatureID() int {
	return notif.list.interfaces.(ENotifyingList[T]).GetFeatureID()
}

func (list *BasicENotifyingList[T]) createNotification(eventType EventType, oldValue any, newValue any, position int) ENotification {
	n := new(notifyingListNotification[T])
	n.Initialize(n, eventType, oldValue, newValue, position)
	n.list = list
	return n
}

func (list *BasicENotifyingList[T]) isNotificationRequired() bool {
	notifier := list.interfaces.(ENotifyingList[T]).GetNotifier()
	return notifier != nil && notifier.EDeliver() && !notifier.EAdapters().Empty()
}

func (list *BasicENotifyingList[T]) createAndAddNotification(ns ENotificationChain, eventType EventType, oldValue any, newValue any, position int) ENotificationChain {
	notifications := ns
	if list.isNotificationRequired() {
		notification := list.createNotification(eventType, oldValue, newValue, position)
		if notifications != nil {
			notifications.Add(notification)
		} else {
			notifications = notification.(ENotificationChain)
		}
	}
	return notifications
}

func (list *BasicENotifyingList[T]) createAndDispatchNotification(notifications ENotificationChain, eventType EventType, oldValue any, newValue any, position int) {
	list.createAndDispatchNotificationFn(notifications, func() ENotification {
		return list.createNotification(eventType, oldValue, newValue, position)
	})
}

func (list *BasicENotifyingList[T]) createAndDispatchNotificationFn(notifications ENotificationChain, createNotification func() ENotification) {
	if list.isNotificationRequired() {
		notification := createNotification()
		if notifications != nil {
			notifications.Add(notification)
			notifications.Dispatch()
		} else {
			notifier := list.interfaces.(ENotifyingList[T]).GetNotifier()
			if notifier != nil {
				notifier.ENotify(notification)
			}
		}
	} else {
		if notifications != nil {
			notifications.Dispatch()
		}
	}
}

func (list *BasicENotifyingList[T]) inverseAdd(object T, notifications ENotificationChain) ENotificationChain {
	return notifications
}

func (list *BasicENotifyingList[T]) inverseRemove(object T, notifications ENotificationChain) ENotificationChain {
	return notifications
}

// AddWithNotification ...
func (list *BasicENotifyingList[T]) AddWithNotification(object T, notifications ENotificationChain) ENotificationChain {
	index := list.Size()
	list.basicEList.doAdd(object)
	return list.createAndAddNotification(notifications, ADD, nil, object, index)
}

// RemoveWithNotification ...
func (list *BasicENotifyingList[T]) RemoveWithNotification(object T, notifications ENotificationChain) ENotificationChain {
	index := list.interfaces.(EList[T]).IndexOf(object)
	if index != -1 {
		oldObject := list.basicEList.doRemove(index)
		return list.createAndAddNotification(notifications, REMOVE, oldObject, nil, index)
	}
	return notifications
}

// SetWithNotification ...
func (list *BasicENotifyingList[T]) SetWithNotification(index int, object T, notifications ENotificationChain) ENotificationChain {
	oldObject := list.basicEList.doSet(index, object)
	return list.createAndAddNotification(notifications, SET, oldObject, object, index)
}

func (list *BasicENotifyingList[T]) doAdd(object T) {
	index := list.interfaces.(EList[T]).Size()
	list.basicEList.doAdd(object)
	notifications := list.interfaces.(eNotifyingListInternal[T]).inverseAdd(object, nil)
	list.createAndDispatchNotification(notifications, ADD, nil, object, index)
}

func (list *BasicENotifyingList[T]) doAddAll(c ECollection[T]) bool {
	return list.interfaces.(abstractEList[T]).doInsertAll(list.Size(), c)
}

func (list *BasicENotifyingList[T]) doInsert(index int, object T) {
	list.basicEList.doInsert(index, object)
	notifications := list.interfaces.(eNotifyingListInternal[T]).inverseAdd(object, nil)
	list.createAndDispatchNotification(notifications, ADD, nil, object, index)
}

func (list *BasicENotifyingList[T]) doInsertAll(index int, c ECollection[T]) bool {
	if c.Empty() {
		return false
	}

	result := list.basicEList.doInsertAll(index, c)
	var notifications ENotificationChain = NewNotificationChain()
	for it := c.Iterator(); it.HasNext(); {
		notifications = list.interfaces.(eNotifyingListInternal[T]).inverseAdd(it.Next(), notifications)
	}
	list.createAndDispatchNotificationFn(notifications, func() ENotification {
		if c.Size() == 1 {
			return list.createNotification(ADD, nil, c.Iterator().Next(), index)
		} else {
			return list.createNotification(ADD_MANY, nil, ToAnyArray[T](c), index)
		}
	})
	return result
}

func (list *BasicENotifyingList[T]) doSet(index int, newObject T) T {
	oldObject := list.basicEList.doSet(index, newObject)
	if newObject != oldObject {
		var notifications ENotificationChain
		notifications = list.interfaces.(eNotifyingListInternal[T]).inverseRemove(oldObject, notifications)
		notifications = list.interfaces.(eNotifyingListInternal[T]).inverseAdd(newObject, notifications)
		list.createAndDispatchNotification(notifications, SET, oldObject, newObject, index)
	}
	return oldObject
}

func (list *BasicENotifyingList[T]) doClear() []T {
	oldData := list.basicEList.doClear()
	if len(oldData) == 0 {
		list.createAndDispatchNotification(nil, REMOVE_MANY, []T{}, nil, -1)
	} else {
		var notifications ENotificationChain = NewNotificationChain()
		for _, e := range oldData {
			notifications = list.interfaces.(eNotifyingListInternal[T]).inverseRemove(e, notifications)
		}

		list.createAndDispatchNotificationFn(notifications,
			func() ENotification {
				if len(oldData) == 1 {
					return list.createNotification(REMOVE, oldData[0], nil, 0)
				} else {
					return list.createNotification(REMOVE_MANY, oldData, nil, -1)
				}
			})
	}
	return oldData
}

func (list *BasicENotifyingList[T]) doMove(oldIndex, newIndex int) T {
	oldObject := list.basicEList.doMove(oldIndex, newIndex)
	list.createAndDispatchNotification(nil, MOVE, oldIndex, oldObject, newIndex)
	return oldObject
}

// RemoveAt ...
func (list *BasicENotifyingList[T]) doRemove(index int) T {
	oldObject := list.basicEList.doRemove(index)
	var notifications ENotificationChain
	notifications = list.interfaces.(eNotifyingListInternal[T]).inverseRemove(oldObject, notifications)
	list.createAndDispatchNotification(notifications, REMOVE, oldObject, nil, index)
	return oldObject
}
