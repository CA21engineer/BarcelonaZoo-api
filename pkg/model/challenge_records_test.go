// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testChallengeRecords(t *testing.T) {
	t.Parallel()

	query := ChallengeRecords()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testChallengeRecordsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChallengeRecord{}
	if err = randomize.Struct(seed, o, challengeRecordDBTypes, true, challengeRecordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ChallengeRecords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testChallengeRecordsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChallengeRecord{}
	if err = randomize.Struct(seed, o, challengeRecordDBTypes, true, challengeRecordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ChallengeRecords().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ChallengeRecords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testChallengeRecordsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChallengeRecord{}
	if err = randomize.Struct(seed, o, challengeRecordDBTypes, true, challengeRecordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ChallengeRecordSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ChallengeRecords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testChallengeRecordsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChallengeRecord{}
	if err = randomize.Struct(seed, o, challengeRecordDBTypes, true, challengeRecordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ChallengeRecordExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ChallengeRecord exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ChallengeRecordExists to return true, but got false.")
	}
}

func testChallengeRecordsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChallengeRecord{}
	if err = randomize.Struct(seed, o, challengeRecordDBTypes, true, challengeRecordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	challengeRecordFound, err := FindChallengeRecord(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if challengeRecordFound == nil {
		t.Error("want a record, got nil")
	}
}

func testChallengeRecordsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChallengeRecord{}
	if err = randomize.Struct(seed, o, challengeRecordDBTypes, true, challengeRecordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ChallengeRecords().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testChallengeRecordsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChallengeRecord{}
	if err = randomize.Struct(seed, o, challengeRecordDBTypes, true, challengeRecordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ChallengeRecords().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testChallengeRecordsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	challengeRecordOne := &ChallengeRecord{}
	challengeRecordTwo := &ChallengeRecord{}
	if err = randomize.Struct(seed, challengeRecordOne, challengeRecordDBTypes, false, challengeRecordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}
	if err = randomize.Struct(seed, challengeRecordTwo, challengeRecordDBTypes, false, challengeRecordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = challengeRecordOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = challengeRecordTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ChallengeRecords().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testChallengeRecordsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	challengeRecordOne := &ChallengeRecord{}
	challengeRecordTwo := &ChallengeRecord{}
	if err = randomize.Struct(seed, challengeRecordOne, challengeRecordDBTypes, false, challengeRecordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}
	if err = randomize.Struct(seed, challengeRecordTwo, challengeRecordDBTypes, false, challengeRecordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = challengeRecordOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = challengeRecordTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ChallengeRecords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func challengeRecordBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ChallengeRecord) error {
	*o = ChallengeRecord{}
	return nil
}

func challengeRecordAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ChallengeRecord) error {
	*o = ChallengeRecord{}
	return nil
}

func challengeRecordAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ChallengeRecord) error {
	*o = ChallengeRecord{}
	return nil
}

func challengeRecordBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ChallengeRecord) error {
	*o = ChallengeRecord{}
	return nil
}

func challengeRecordAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ChallengeRecord) error {
	*o = ChallengeRecord{}
	return nil
}

func challengeRecordBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ChallengeRecord) error {
	*o = ChallengeRecord{}
	return nil
}

func challengeRecordAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ChallengeRecord) error {
	*o = ChallengeRecord{}
	return nil
}

func challengeRecordBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ChallengeRecord) error {
	*o = ChallengeRecord{}
	return nil
}

func challengeRecordAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ChallengeRecord) error {
	*o = ChallengeRecord{}
	return nil
}

func testChallengeRecordsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ChallengeRecord{}
	o := &ChallengeRecord{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, challengeRecordDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord object: %s", err)
	}

	AddChallengeRecordHook(boil.BeforeInsertHook, challengeRecordBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	challengeRecordBeforeInsertHooks = []ChallengeRecordHook{}

	AddChallengeRecordHook(boil.AfterInsertHook, challengeRecordAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	challengeRecordAfterInsertHooks = []ChallengeRecordHook{}

	AddChallengeRecordHook(boil.AfterSelectHook, challengeRecordAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	challengeRecordAfterSelectHooks = []ChallengeRecordHook{}

	AddChallengeRecordHook(boil.BeforeUpdateHook, challengeRecordBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	challengeRecordBeforeUpdateHooks = []ChallengeRecordHook{}

	AddChallengeRecordHook(boil.AfterUpdateHook, challengeRecordAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	challengeRecordAfterUpdateHooks = []ChallengeRecordHook{}

	AddChallengeRecordHook(boil.BeforeDeleteHook, challengeRecordBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	challengeRecordBeforeDeleteHooks = []ChallengeRecordHook{}

	AddChallengeRecordHook(boil.AfterDeleteHook, challengeRecordAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	challengeRecordAfterDeleteHooks = []ChallengeRecordHook{}

	AddChallengeRecordHook(boil.BeforeUpsertHook, challengeRecordBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	challengeRecordBeforeUpsertHooks = []ChallengeRecordHook{}

	AddChallengeRecordHook(boil.AfterUpsertHook, challengeRecordAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	challengeRecordAfterUpsertHooks = []ChallengeRecordHook{}
}

func testChallengeRecordsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChallengeRecord{}
	if err = randomize.Struct(seed, o, challengeRecordDBTypes, true, challengeRecordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ChallengeRecords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testChallengeRecordsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChallengeRecord{}
	if err = randomize.Struct(seed, o, challengeRecordDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(challengeRecordColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ChallengeRecords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testChallengeRecordToManyFavorites(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ChallengeRecord
	var b, c Favorite

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, challengeRecordDBTypes, true, challengeRecordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, favoriteDBTypes, false, favoriteColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, favoriteDBTypes, false, favoriteColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ChallengeRecordID = a.ID
	c.ChallengeRecordID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Favorites().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ChallengeRecordID == b.ChallengeRecordID {
			bFound = true
		}
		if v.ChallengeRecordID == c.ChallengeRecordID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ChallengeRecordSlice{&a}
	if err = a.L.LoadFavorites(ctx, tx, false, (*[]*ChallengeRecord)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Favorites); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Favorites = nil
	if err = a.L.LoadFavorites(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Favorites); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testChallengeRecordToManyAddOpFavorites(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ChallengeRecord
	var b, c, d, e Favorite

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, challengeRecordDBTypes, false, strmangle.SetComplement(challengeRecordPrimaryKeyColumns, challengeRecordColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Favorite{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, favoriteDBTypes, false, strmangle.SetComplement(favoritePrimaryKeyColumns, favoriteColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Favorite{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddFavorites(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ChallengeRecordID {
			t.Error("foreign key was wrong value", a.ID, first.ChallengeRecordID)
		}
		if a.ID != second.ChallengeRecordID {
			t.Error("foreign key was wrong value", a.ID, second.ChallengeRecordID)
		}

		if first.R.ChallengeRecord != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.ChallengeRecord != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Favorites[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Favorites[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Favorites().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testChallengeRecordToOneChallengeThemeUsingChallengeTheme(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ChallengeRecord
	var foreign ChallengeTheme

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, challengeRecordDBTypes, false, challengeRecordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, challengeThemeDBTypes, false, challengeThemeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChallengeTheme struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ChallengeThemeID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ChallengeTheme().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ChallengeRecordSlice{&local}
	if err = local.L.LoadChallengeTheme(ctx, tx, false, (*[]*ChallengeRecord)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ChallengeTheme == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ChallengeTheme = nil
	if err = local.L.LoadChallengeTheme(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ChallengeTheme == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testChallengeRecordToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ChallengeRecord
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, challengeRecordDBTypes, false, challengeRecordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ChallengeRecordSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*ChallengeRecord)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testChallengeRecordToOneSetOpChallengeThemeUsingChallengeTheme(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ChallengeRecord
	var b, c ChallengeTheme

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, challengeRecordDBTypes, false, strmangle.SetComplement(challengeRecordPrimaryKeyColumns, challengeRecordColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, challengeThemeDBTypes, false, strmangle.SetComplement(challengeThemePrimaryKeyColumns, challengeThemeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, challengeThemeDBTypes, false, strmangle.SetComplement(challengeThemePrimaryKeyColumns, challengeThemeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ChallengeTheme{&b, &c} {
		err = a.SetChallengeTheme(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ChallengeTheme != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ChallengeRecords[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ChallengeThemeID != x.ID {
			t.Error("foreign key was wrong value", a.ChallengeThemeID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ChallengeThemeID))
		reflect.Indirect(reflect.ValueOf(&a.ChallengeThemeID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ChallengeThemeID != x.ID {
			t.Error("foreign key was wrong value", a.ChallengeThemeID, x.ID)
		}
	}
}
func testChallengeRecordToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ChallengeRecord
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, challengeRecordDBTypes, false, strmangle.SetComplement(challengeRecordPrimaryKeyColumns, challengeRecordColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ChallengeRecords[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testChallengeRecordsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChallengeRecord{}
	if err = randomize.Struct(seed, o, challengeRecordDBTypes, true, challengeRecordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testChallengeRecordsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChallengeRecord{}
	if err = randomize.Struct(seed, o, challengeRecordDBTypes, true, challengeRecordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ChallengeRecordSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testChallengeRecordsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ChallengeRecord{}
	if err = randomize.Struct(seed, o, challengeRecordDBTypes, true, challengeRecordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ChallengeRecords().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	challengeRecordDBTypes = map[string]string{`ID`: `int`, `Content`: `text`, `Comment`: `text`, `ChallengeThemeID`: `int`, `UserID`: `int`, `Record`: `float`, `CreatedAt`: `datetime`}
	_                      = bytes.MinRead
)

func testChallengeRecordsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(challengeRecordPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(challengeRecordAllColumns) == len(challengeRecordPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ChallengeRecord{}
	if err = randomize.Struct(seed, o, challengeRecordDBTypes, true, challengeRecordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ChallengeRecords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, challengeRecordDBTypes, true, challengeRecordPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testChallengeRecordsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(challengeRecordAllColumns) == len(challengeRecordPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ChallengeRecord{}
	if err = randomize.Struct(seed, o, challengeRecordDBTypes, true, challengeRecordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ChallengeRecords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, challengeRecordDBTypes, true, challengeRecordPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(challengeRecordAllColumns, challengeRecordPrimaryKeyColumns) {
		fields = challengeRecordAllColumns
	} else {
		fields = strmangle.SetComplement(
			challengeRecordAllColumns,
			challengeRecordPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ChallengeRecordSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testChallengeRecordsUpsert(t *testing.T) {
	t.Parallel()

	if len(challengeRecordAllColumns) == len(challengeRecordPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLChallengeRecordUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ChallengeRecord{}
	if err = randomize.Struct(seed, &o, challengeRecordDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ChallengeRecord: %s", err)
	}

	count, err := ChallengeRecords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, challengeRecordDBTypes, false, challengeRecordPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ChallengeRecord struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ChallengeRecord: %s", err)
	}

	count, err = ChallengeRecords().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
