package queue

import "testing"

func TestAdd(t *testing.T) {
	q := New(5)
	for i := 0; i < 5; i++ {
		if len(q.items) != i {
			t.Errorf("incorrect Queue elemnt count : %v,want %v", len(q.items), i)
		}
		if !q.append(i) {
			t.Errorf("failed to append item %v to Queue", i)

		}

	}
	if q.append(6) {
		t.Errorf("shoud not be able to add to Queue")
	}

}
func TestNew(t *testing.T) {
	q := New(5)
	for i := 0; i < 5; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("shoud not get items from Queue")
		}
		if item != i {
			t.Errorf("got wrong order : %v , wawnt %v", item, i)
		}
	}
	// Queue is empty at this point
	item, ok := q.Next()
	if ok {
		t.Errorf("shoud not be any item in Queue got : %v", item)
	}
}
