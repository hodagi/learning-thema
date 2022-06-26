package example

import (
	cuejson "cuelang.org/go/encoding/json"

	"fmt"
	"testing"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"github.com/grafana/thema"
)

var lib thema.Library = thema.NewLibrary(cuecontext.New())
var shiplin thema.Lineage

// Ship.cueの血統を読み込む
func init() {
	var err error
	if shiplin, err = ShipLineage(lib); err != nil {
		panic(err)
	}
}

// jsonデータをcuelangデータに変換する
func dataAsValue(lib thema.Library) cue.Value {
	// The first parameter gives the CUE evaluator a name that it will use for
	// any future errors involving the extracted data. Usually this is derived
	// from a file path, but our input is coming from an arbitrary string, so we
	// must name it ourselves.
	expr, _ := cuejson.Extract("input", input)

	// Load our data into a CUE context-universe and return a cue.Value reference to it.
	return lib.Context().BuildExpr(expr)
}

// 元データ
var input = []byte(`{
    "firstfield": "foo"
}`)

type Ship00 struct {
	Firstfield string `json:"firstfield`
}

type Ship10 struct {
	Firstfield  string `json:"firstfield`
	Secondfield int    `json:"secondfield`
}

var targetVersion = thema.SV(1, 0)

func TestSearchByValid(t *testing.T) {

	fmt.Printf("#######################################\n")

	// JSONデータをCUELANGデータに変換する
	cue00 := dataAsValue(lib)

	// CUELANGデータが、ship.cueで定義されたいずれかのスキーマバージョンだよねってことを確認する
	inst00 := shiplin.ValidateAny(cue00)
	if inst00 == nil {
		t.Fatal("expected input data to validate against schema 0.0")
	}

	var ship00 Ship00
	// スキーマバージョンは0.0バージョンらしい
	inst00.UnwrapCUE().Decode(&ship00)
	fmt.Printf("version"+inst00.Schema().Version().String()+":%+v\n", ship00) // "{Firstfield:foo}"

	// 1.0バージョンに変換してみる
	var ship10 Ship10
	inst10, lacunas := inst00.Translate(targetVersion)

	// 先ほどのバージョンにはなかったSecondfieldができていることがわかる
	inst10.UnwrapCUE().Decode(&ship10)
	fmt.Printf("version"+inst10.Schema().Version().String()+":%+v\n", ship10) // "{Firstfield:foo Secondfield:-1}"

	// バージョン0.0~>1.0に上げた際に"欠落"した情報を確認する
	fmt.Printf("Lacuna(欠落したフィールド情報)\n")
	for _, s := range lacunas.AsList() {
		fmt.Println(s) // {[] [{secondfield <nil>}] 0 -1 used as a placeholder value - replace with a real value before persisting!}
	}
	fmt.Printf("#######################################\n")
}
