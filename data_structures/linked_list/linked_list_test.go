package linkedlist

// TODO list creation outsides the intern test function

//func TestNewEmptyLinkedList(t *testing.T) {
//	got := NewEmptyLinkedList[int]()
//	want := &linkedList[int]{}
//
//	if equal := reflect.DeepEqual(got, want); !equal {
//		t.Errorf("[Linked List] NewEmptyLinkedList, got %v want %v", got, want)
//	}
//}

//func TestSize(t *testing.T) {
//	tests := []struct {
//		name              string
//		numElementsToPush int64
//		realListSize      uint64
//	}{
//		{"empty size", 0, 0},
//		{"1 element", 1, 1},
//		{"9999 elements", 9999, 9999},
//		// TODO Test removing elements
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			list := NewEmptyLinkedList()
//
//			var i int64
//			for i = 0; i < tt.numElementsToPush; i++ {
//				list.PushBack(NewEmptyListElement())
//			}
//
//			if list.Size() != tt.realListSize {
//				t.Errorf("[Linked List] Size, got %d want %d", list.Size(), tt.realListSize)
//			}
//
//			var count uint64 = 0
//			var node *ListElement = list.Front()
//			for node != nil {
//				node = node.GetNext()
//				count++
//			}
//
//			if count != list.Size() {
//				t.Errorf("[Linked List] Size, got %d list.size %d want %d", count, list.Size(), tt.realListSize)
//			}
//		})
//	}
//}
//
//func TestEmpty(t *testing.T) {
//	tests := []struct {
//		name             string
//		numElementsToAdd int64
//		isEmpty          bool
//	}{
//		{"empty list", 0, true},
//		{"1 element list", 1, false},
//		{"100 element list", 100, false},
//		// TODO pop elements and check
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			list := NewEmptyLinkedList()
//			var i int64
//			for i = 0; i < tt.numElementsToAdd; i++ {
//				list.PushBack(NewEmptyListElement())
//			}
//
//			if list.Empty() != tt.isEmpty {
//				t.Errorf("[Linked List] Empty, got %t want %t", list.Empty(), tt.isEmpty)
//			}
//		})
//	}
//}
//
//func TestValueAt(t *testing.T) {
//	list := NewEmptyLinkedList()
//	listSize := 5
//	for i := 0; i < listSize; i++ {
//		list.PushBack(NewListElement(int64(i), nil))
//	}
//
//	tests := []struct {
//		name         string
//		index        uint64
//		valueAtIndex int64
//		wantError    bool
//	}{
//		{"head index", 0, list.Front().GetValue(), false},
//		{"tail index", list.Size() - 1, list.Back().GetValue(), false},
//		{"index 1", 1, list.Front().GetNext().GetValue(), false},
//		{"index 2", 2, list.Front().GetNext().GetNext().GetValue(), false},
//		{"index 3", 3, list.Front().GetNext().GetNext().GetNext().GetValue(), false},
//		{"out of index", list.Size() + 1, 0, true},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			value, err := list.ValueAt(tt.index)
//
//			if (tt.wantError && err == nil) || (!tt.wantError && err != nil) {
//				t.Errorf("[Linked List] ValueAt err, got %v wantError? %v", err, tt.wantError)
//			}
//
//			if value != tt.valueAtIndex {
//				t.Errorf("[Linked List] ValueAt value, got %d want %d", value, tt.valueAtIndex)
//			}
//		})
//	}
//}
//
//func TestPushBack(t *testing.T) {
//	tests := []struct {
//		name        string
//		numElements int
//		listSize    uint64
//	}{
//		{"nil PushBack", 0, 0},
//		{"1 PushBack", 1, 1},
//		{"100 PushBack", 100, 100},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			// new list
//			list := NewEmptyLinkedList()
//			// elements insertion
//			var lastElement *ListElement
//			for i := 0; i < tt.numElements; i++ {
//				lastElement = NewEmptyListElement()
//				list.PushBack(lastElement)
//			}
//
//			if list.Size() != tt.listSize {
//				t.Errorf("[Linked List] PushBack, List size %d want %d", list.Size(), tt.listSize)
//			}
//
//			if equal := reflect.DeepEqual(lastElement, list.Back()); !equal {
//				t.Errorf("[Linked List] PushBack, List Back %v want %v", list.Back(), lastElement)
//			}
//
//			var countElements uint64 = 0
//			node := list.Front()
//			for node != nil {
//				node = node.GetNext()
//				countElements++
//			}
//
//			if countElements != list.Size() {
//				t.Errorf("[Linked List] PushBack, Count elements %d want %d", countElements, tt.listSize)
//			}
//		})
//	}
//}
//
//func TestPushFront(t *testing.T) {
//	tests := []struct {
//		name        string
//		numElements int64
//		listSize    uint64
//	}{
//		{"nil PushFront", 0, 0},
//		{"1 PushFront", 1, 1},
//		{"100 PushFront", 100, 100},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			// new lists
//			list := NewEmptyLinkedList()
//			// elements insertion
//			var lastElementInserted *ListElement
//			var i int64
//			for i = 0; i < tt.numElements; i++ {
//				lastElementInserted = NewListElement(i, nil)
//				list.PushFront(lastElementInserted)
//			}
//
//			if list.Size() != tt.listSize {
//				t.Errorf("[Linked List] PushFront, List size %d want %d", list.Size(), tt.listSize)
//			}
//
//			if equal := reflect.DeepEqual(lastElementInserted, list.Front()); !equal {
//				t.Errorf("[Linked List] PushFront, List Front %v want %v", list.Front(), lastElementInserted)
//			}
//
//			var countElements uint64 = 0
//			node := list.Front()
//			for node != nil {
//				node = node.GetNext()
//				countElements++
//			}
//
//			if countElements != list.Size() {
//				t.Errorf("[Linked List] PushFront, Count elements %d want %d", countElements, tt.listSize)
//			}
//		})
//	}
//}
//
//func TestInsertAt(t *testing.T) {
//	list := NewEmptyLinkedList()
//	listSize := 5
//	for i := 0; i < listSize; i++ {
//		list.PushBack(NewListElement(int64(i), nil))
//	}
//
//	tests := []struct {
//		name            string
//		elementToInsert *ListElement
//		index           uint64
//		wantError       bool
//	}{
//		{"error, index > size", nil, list.Size() + 1, true},
//		{"insert at head", NewListElement(2222, nil), 0, false},
//		{"insert at tail", NewListElement(3333, nil), list.Size() - 1, false},
//		{"middle insertion", NewListElement(9999, nil), (list.Size() - 1) / 2, false},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			listTest := NewEmptyLinkedList()
//
//			err := listTest.InsertAt(tt.elementToInsert, tt.index)
//
//			if (tt.wantError && err == nil) || (!tt.wantError && err != nil) {
//				t.Errorf("[Linked List] InsertAt, got %v wantErr? %v", tt.wantError, err)
//				return
//			}
//
//			if equal := reflect.DeepEqual(list.ElementAt(tt.index), tt.elementToInsert); !equal {
//				t.Errorf("[Linked List] InsertAt, got %v want %v", listTest.ElementAt(tt.index), tt.elementToInsert)
//			}
//		})
//	}
//}
//
//func TestBack(t *testing.T) {
//	tests := []struct {
//		name    string
//		element *ListElement
//	}{
//		{"empty list", nil},
//		{"list size 1", &ListElement{}},
//		// TODO Test renoving elements
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			list := NewEmptyLinkedList()
//			list.PushBack(tt.element)
//			if equal := reflect.DeepEqual(list.Back(), tt.element); !equal {
//				t.Errorf("[Linked List] Back, %v.Back() = %v, want %v", list, list.Back(), tt.element)
//			}
//		})
//	}
//}
//
//func TestFront(t *testing.T) {
//	tests := []struct {
//		name              string
//		numElementsToPush int64
//		element           *ListElement
//	}{
//		{"empty list", 0, nil},
//		{"list size 1", 1, &ListElement{}},
//		{"list size 1", 10, &ListElement{}},
//		// TODO Test renoving elements
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			// new empty list
//			list := NewEmptyLinkedList()
//
//			list.PushBack(tt.element)
//
//			var i int64
//			for i = 0; i < tt.numElementsToPush; i++ {
//				list.PushBack(NewEmptyListElement())
//			}
//			if equal := reflect.DeepEqual(list.Front(), tt.element); !equal {
//				t.Errorf("[Linked List] Back, %v.Back() = %v, want %v", list, list.Back(), tt.element)
//			}
//		})
//	}
//}
