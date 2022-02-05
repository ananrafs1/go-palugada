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
		stbs = append(stbs, (*((*st).(*[]stub)))[0])
		fmt.Println((*((*st).(*[]stub))))
		(*((*st).(*[]stub))) = (*((*st).(*[]stub)))[1:]
		return nil
	}
	wrap := RecurseTry(Klausa, func(subs *interface{}) bool { return len((*(*subs).(*[]stub))) < 1 }, 3, time.Duration(2*time.Second))
	var pr *interface{} = new(interface{})
	*pr = &stubs
	_ = wrap(pr)
	if len(stubs) > 0 {
		t.Error("Should return 0, got ", len(stubs))
	}

}

