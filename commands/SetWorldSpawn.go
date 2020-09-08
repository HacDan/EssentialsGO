package commands
// Eren5960 <ahmederen123@gmail.com>
import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"github.com/eren5960/essentialsgo/commands/op"
	"github.com/eren5960/essentialsgo/global"
)

type SetWorldSpawnXYZ struct {
	X, Y, Z float64
}

func (t SetWorldSpawnXYZ) Run(source cmd.Source, output *cmd.Output) {
	if !op.IsOp(source){
		output.Error("You don't have permission for this command.")
		return
	}

	bp := world.BlockPos{int(t.X), int(t.Y), int(t.Z)}
	global.Server.World().SetSpawn(bp)
	output.Printf("Set the default world spawn point to (%d, %d, %d)", bp.X(), bp.Y(), bp.Z())
}

type SetWorldSpawn struct {
	Sub string `optional:""`
}

func (t SetWorldSpawn) Run(source cmd.Source, output *cmd.Output) {
	if t.Sub != "" {
		return
	}
	if !op.IsOp(source){
		output.Error("You don't have permission for this command.")
		return
	}
	if p, ok := source.(*player.Player); ok {
		pos := p.Position()
		bp := world.BlockPos{int(pos.X()), int(pos.Y()), int(pos.Z())}
		p.World().SetSpawn(bp)
		output.Printf("Set the world spawn point to (%d, %d, %d)", bp.X(), bp.Y(), bp.Z())
	} else {
		output.Error("Usage: /setworldspawn <X: float> <Y: float> <Z: float>")
	}
}