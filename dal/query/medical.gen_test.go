// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"testing"

	"box/dal/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&model.Medical{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.Medical{}) fail: %s", err)
	}
}

func Test_medicalQuery(t *testing.T) {
	medical := newMedical(db)
	medical = *medical.As(medical.TableName())
	_do := medical.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(medical.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <tblMedical> fail:", err)
		return
	}

	_, ok := medical.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from medical success")
	}

	err = _do.Create(&model.Medical{})
	if err != nil {
		t.Error("create item in table <tblMedical> fail:", err)
	}

	err = _do.Save(&model.Medical{})
	if err != nil {
		t.Error("create item in table <tblMedical> fail:", err)
	}

	err = _do.CreateInBatches([]*model.Medical{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <tblMedical> fail:", err)
	}

	_, err = _do.Select(medical.ALL).Take()
	if err != nil {
		t.Error("Take() on table <tblMedical> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <tblMedical> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <tblMedical> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <tblMedical> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.Medical{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <tblMedical> fail:", err)
	}

	_, err = _do.Select(medical.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <tblMedical> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <tblMedical> fail:", err)
	}

	_, err = _do.Select(medical.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <tblMedical> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <tblMedical> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <tblMedical> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <tblMedical> fail:", err)
	}

	_, err = _do.ScanByPage(&model.Medical{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <tblMedical> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <tblMedical> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <tblMedical> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <tblMedical> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <tblMedical> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <tblMedical> fail:", err)
	}
}
