package retry

import(
	"time"
	"testing"
	"fmt"
)

type stub struct {
	alpha string
}

func Test_RecurseTry_ShouldSucces(t *testing.T){
	stubs := []stub{
		stub{
			alpha:"aaaw",
		},
		stub{
			alpha:"aab",
		},
		stub{
			alpha:"add",
		},
	}
	stbs := make([]stub, 0)
	Klausa := func(st *interface{}) error {
		stubPointer := (*st).(*[]stub) //grab pointer
		fmt.Println(*stubPointer) //dereference to get value
		stbs = append(stbs, (*stubPointer)[0]) 
		(*stubPointer) = (*stubPointer)[1:]
		return nil
	}
	wrap := RecurseTry(Klausa, func(subs *interface{}) bool { 
		stubPointer := (*subs).(*[]stub)
		return len((*stubPointer)) < 1 
		}, 3, time.Duration(2*time.Second))
	pr := new(interface{})
	*pr = &stubs
	_ = wrap(pr)
	if len(stubs) > 0 {
		t.Error("Should return 0, got ", len(stubs))
	}

}

