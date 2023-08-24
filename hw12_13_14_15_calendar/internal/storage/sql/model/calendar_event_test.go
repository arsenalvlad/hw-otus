// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

func testCalendarEventsUpsert(t *testing.T) {
	t.Parallel()

	if len(calendarEventAllColumns) == len(calendarEventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := CalendarEvent{}
	if err = randomize.Struct(seed, &o, calendarEventDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CalendarEvent: %s", err)
	}

	count, err := CalendarEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, calendarEventDBTypes, false, calendarEventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CalendarEvent: %s", err)
	}

	count, err = CalendarEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testCalendarEvents(t *testing.T) {
	t.Parallel()

	query := CalendarEvents()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCalendarEventsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CalendarEvent{}
	if err = randomize.Struct(seed, o, calendarEventDBTypes, true, calendarEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
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

	count, err := CalendarEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCalendarEventsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CalendarEvent{}
	if err = randomize.Struct(seed, o, calendarEventDBTypes, true, calendarEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := CalendarEvents().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CalendarEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCalendarEventsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CalendarEvent{}
	if err = randomize.Struct(seed, o, calendarEventDBTypes, true, calendarEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CalendarEventSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CalendarEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCalendarEventsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CalendarEvent{}
	if err = randomize.Struct(seed, o, calendarEventDBTypes, true, calendarEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CalendarEventExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if CalendarEvent exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CalendarEventExists to return true, but got false.")
	}
}

func testCalendarEventsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CalendarEvent{}
	if err = randomize.Struct(seed, o, calendarEventDBTypes, true, calendarEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	calendarEventFound, err := FindCalendarEvent(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if calendarEventFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCalendarEventsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CalendarEvent{}
	if err = randomize.Struct(seed, o, calendarEventDBTypes, true, calendarEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = CalendarEvents().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCalendarEventsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CalendarEvent{}
	if err = randomize.Struct(seed, o, calendarEventDBTypes, true, calendarEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := CalendarEvents().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCalendarEventsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	calendarEventOne := &CalendarEvent{}
	calendarEventTwo := &CalendarEvent{}
	if err = randomize.Struct(seed, calendarEventOne, calendarEventDBTypes, false, calendarEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
	}
	if err = randomize.Struct(seed, calendarEventTwo, calendarEventDBTypes, false, calendarEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = calendarEventOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = calendarEventTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CalendarEvents().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCalendarEventsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	calendarEventOne := &CalendarEvent{}
	calendarEventTwo := &CalendarEvent{}
	if err = randomize.Struct(seed, calendarEventOne, calendarEventDBTypes, false, calendarEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
	}
	if err = randomize.Struct(seed, calendarEventTwo, calendarEventDBTypes, false, calendarEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = calendarEventOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = calendarEventTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CalendarEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func calendarEventBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *CalendarEvent) error {
	*o = CalendarEvent{}
	return nil
}

func calendarEventAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *CalendarEvent) error {
	*o = CalendarEvent{}
	return nil
}

func calendarEventAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *CalendarEvent) error {
	*o = CalendarEvent{}
	return nil
}

func calendarEventBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *CalendarEvent) error {
	*o = CalendarEvent{}
	return nil
}

func calendarEventAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *CalendarEvent) error {
	*o = CalendarEvent{}
	return nil
}

func calendarEventBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *CalendarEvent) error {
	*o = CalendarEvent{}
	return nil
}

func calendarEventAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *CalendarEvent) error {
	*o = CalendarEvent{}
	return nil
}

func calendarEventBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *CalendarEvent) error {
	*o = CalendarEvent{}
	return nil
}

func calendarEventAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *CalendarEvent) error {
	*o = CalendarEvent{}
	return nil
}

func testCalendarEventsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &CalendarEvent{}
	o := &CalendarEvent{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, calendarEventDBTypes, false); err != nil {
		t.Errorf("Unable to randomize CalendarEvent object: %s", err)
	}

	AddCalendarEventHook(boil.BeforeInsertHook, calendarEventBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	calendarEventBeforeInsertHooks = []CalendarEventHook{}

	AddCalendarEventHook(boil.AfterInsertHook, calendarEventAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	calendarEventAfterInsertHooks = []CalendarEventHook{}

	AddCalendarEventHook(boil.AfterSelectHook, calendarEventAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	calendarEventAfterSelectHooks = []CalendarEventHook{}

	AddCalendarEventHook(boil.BeforeUpdateHook, calendarEventBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	calendarEventBeforeUpdateHooks = []CalendarEventHook{}

	AddCalendarEventHook(boil.AfterUpdateHook, calendarEventAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	calendarEventAfterUpdateHooks = []CalendarEventHook{}

	AddCalendarEventHook(boil.BeforeDeleteHook, calendarEventBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	calendarEventBeforeDeleteHooks = []CalendarEventHook{}

	AddCalendarEventHook(boil.AfterDeleteHook, calendarEventAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	calendarEventAfterDeleteHooks = []CalendarEventHook{}

	AddCalendarEventHook(boil.BeforeUpsertHook, calendarEventBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	calendarEventBeforeUpsertHooks = []CalendarEventHook{}

	AddCalendarEventHook(boil.AfterUpsertHook, calendarEventAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	calendarEventAfterUpsertHooks = []CalendarEventHook{}
}

func testCalendarEventsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CalendarEvent{}
	if err = randomize.Struct(seed, o, calendarEventDBTypes, true, calendarEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CalendarEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCalendarEventsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CalendarEvent{}
	if err = randomize.Struct(seed, o, calendarEventDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(calendarEventColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := CalendarEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCalendarEventsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CalendarEvent{}
	if err = randomize.Struct(seed, o, calendarEventDBTypes, true, calendarEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
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

func testCalendarEventsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CalendarEvent{}
	if err = randomize.Struct(seed, o, calendarEventDBTypes, true, calendarEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CalendarEventSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCalendarEventsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CalendarEvent{}
	if err = randomize.Struct(seed, o, calendarEventDBTypes, true, calendarEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CalendarEvents().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	calendarEventDBTypes = map[string]string{`ID`: `bigint`, `UserID`: `bigint`, `Title`: `character varying`, `EventTime`: `timestamp without time zone`, `Duration`: `bigint`, `Description`: `character varying`}
	_                    = bytes.MinRead
)

func testCalendarEventsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(calendarEventPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(calendarEventAllColumns) == len(calendarEventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CalendarEvent{}
	if err = randomize.Struct(seed, o, calendarEventDBTypes, true, calendarEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CalendarEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, calendarEventDBTypes, true, calendarEventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCalendarEventsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(calendarEventAllColumns) == len(calendarEventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CalendarEvent{}
	if err = randomize.Struct(seed, o, calendarEventDBTypes, true, calendarEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CalendarEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, calendarEventDBTypes, true, calendarEventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CalendarEvent struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(calendarEventAllColumns, calendarEventPrimaryKeyColumns) {
		fields = calendarEventAllColumns
	} else {
		fields = strmangle.SetComplement(
			calendarEventAllColumns,
			calendarEventPrimaryKeyColumns,
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

	slice := CalendarEventSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}