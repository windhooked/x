package integration_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/theplant/gofixtures"

	"github.com/sunfmin/bran/ui"

	"github.com/sunfmin/bran/presets/gormop"
)

type TestVariant struct {
	ProductCode string
	ColorCode   string
	Name        string
}

var emptyData = gofixtures.Data(gofixtures.Sql(``, []string{"test_variants"}))

func (tv *TestVariant) PrimarySlug() string {
	return fmt.Sprintf("%s_%s", tv.ProductCode, tv.ColorCode)
}

func (tv *TestVariant) PrimaryColumnValuesBySlug(slug string) [][]string {
	segs := strings.Split(slug, "_")
	if len(segs) != 2 {
		panic("wrong slug")
	}

	return [][]string{
		{"product_code", segs[0]},
		{"color_code", segs[1]},
	}
}

func TestPrimarySlugger(t *testing.T) {
	db := ConnectDB()
	db.AutoMigrate(&TestVariant{})
	emptyData.TruncatePut(db)
	op := gormop.DataOperator(db)
	ctx := new(ui.EventContext)
	err := op.Save(&TestVariant{ProductCode: "P01", ColorCode: "C01", Name: "Product 1"}, "", ctx)
	if err != nil {
		panic(err)
	}

	err = op.Save(&TestVariant{Name: "Product 2"}, "P01_C01", ctx)
	if err != nil {
		panic(err)
	}

	tv, err := op.Fetch(&TestVariant{}, "P01_C01", ctx)
	if err != nil {
		panic(err)
	}

	if tv.(*TestVariant).Name != "Product 2" {
		t.Error("didn't update product 2", tv)
	}

	err = op.Delete(&TestVariant{}, "P01_C01", ctx)
	if err != nil {
		panic(err)
	}

	tv, err = op.Fetch(&TestVariant{}, "P01_C01", ctx)
	if err != gorm.ErrRecordNotFound {
		t.Error("didn't return not found after delete", tv, err)
	}

}
