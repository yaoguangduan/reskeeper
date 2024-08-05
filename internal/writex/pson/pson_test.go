package pson

import (
	"fmt"
	"github.com/yaoguangduan/reskeeper/pbgen"
	"google.golang.org/protobuf/types/dynamicpb"
	"testing"
)

func TestPson(t *testing.T) {

	s := `{age:12,name:alen,addr:{name:BJ,x:43.12,y:98.41},assistants:[{name:AA,level:99},{name:CC,direction:health}],actonFlow:{1998:eat,2020:fly},papers:{9123987:{id:81723,desc:woigwieqw},97617232:{id:9872312,desc:owughqiweuyqw}}}`
	manager := pbgen.Manager{}
	m := dynamicpb.NewMessage(manager.ProtoReflect().Descriptor())
	Decode(m, s)
	fmt.Println(m)
}
