package data

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/yaoguangduan/reskeeper/pbgen"
	"google.golang.org/protobuf/proto"
	"os"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	users := &pbgen.UserTable{}
	bys := lo.Must(os.ReadFile("user.full.bin"))
	err := proto.Unmarshal(bys, users)
	if err != nil {
		panic(err)
	}
	fmt.Println(users.Users[0].Friends)
}
