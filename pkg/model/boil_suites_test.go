// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("ChallengeRecords", testChallengeRecords)
	t.Run("ChallengeThemes", testChallengeThemes)
	t.Run("Favorites", testFavorites)
	t.Run("Tags", testTags)
	t.Run("ThemeTagRelations", testThemeTagRelations)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("ChallengeRecords", testChallengeRecordsDelete)
	t.Run("ChallengeThemes", testChallengeThemesDelete)
	t.Run("Favorites", testFavoritesDelete)
	t.Run("Tags", testTagsDelete)
	t.Run("ThemeTagRelations", testThemeTagRelationsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("ChallengeRecords", testChallengeRecordsQueryDeleteAll)
	t.Run("ChallengeThemes", testChallengeThemesQueryDeleteAll)
	t.Run("Favorites", testFavoritesQueryDeleteAll)
	t.Run("Tags", testTagsQueryDeleteAll)
	t.Run("ThemeTagRelations", testThemeTagRelationsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("ChallengeRecords", testChallengeRecordsSliceDeleteAll)
	t.Run("ChallengeThemes", testChallengeThemesSliceDeleteAll)
	t.Run("Favorites", testFavoritesSliceDeleteAll)
	t.Run("Tags", testTagsSliceDeleteAll)
	t.Run("ThemeTagRelations", testThemeTagRelationsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("ChallengeRecords", testChallengeRecordsExists)
	t.Run("ChallengeThemes", testChallengeThemesExists)
	t.Run("Favorites", testFavoritesExists)
	t.Run("Tags", testTagsExists)
	t.Run("ThemeTagRelations", testThemeTagRelationsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("ChallengeRecords", testChallengeRecordsFind)
	t.Run("ChallengeThemes", testChallengeThemesFind)
	t.Run("Favorites", testFavoritesFind)
	t.Run("Tags", testTagsFind)
	t.Run("ThemeTagRelations", testThemeTagRelationsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("ChallengeRecords", testChallengeRecordsBind)
	t.Run("ChallengeThemes", testChallengeThemesBind)
	t.Run("Favorites", testFavoritesBind)
	t.Run("Tags", testTagsBind)
	t.Run("ThemeTagRelations", testThemeTagRelationsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("ChallengeRecords", testChallengeRecordsOne)
	t.Run("ChallengeThemes", testChallengeThemesOne)
	t.Run("Favorites", testFavoritesOne)
	t.Run("Tags", testTagsOne)
	t.Run("ThemeTagRelations", testThemeTagRelationsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("ChallengeRecords", testChallengeRecordsAll)
	t.Run("ChallengeThemes", testChallengeThemesAll)
	t.Run("Favorites", testFavoritesAll)
	t.Run("Tags", testTagsAll)
	t.Run("ThemeTagRelations", testThemeTagRelationsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("ChallengeRecords", testChallengeRecordsCount)
	t.Run("ChallengeThemes", testChallengeThemesCount)
	t.Run("Favorites", testFavoritesCount)
	t.Run("Tags", testTagsCount)
	t.Run("ThemeTagRelations", testThemeTagRelationsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("ChallengeRecords", testChallengeRecordsHooks)
	t.Run("ChallengeThemes", testChallengeThemesHooks)
	t.Run("Favorites", testFavoritesHooks)
	t.Run("Tags", testTagsHooks)
	t.Run("ThemeTagRelations", testThemeTagRelationsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("ChallengeRecords", testChallengeRecordsInsert)
	t.Run("ChallengeRecords", testChallengeRecordsInsertWhitelist)
	t.Run("ChallengeThemes", testChallengeThemesInsert)
	t.Run("ChallengeThemes", testChallengeThemesInsertWhitelist)
	t.Run("Favorites", testFavoritesInsert)
	t.Run("Favorites", testFavoritesInsertWhitelist)
	t.Run("Tags", testTagsInsert)
	t.Run("Tags", testTagsInsertWhitelist)
	t.Run("ThemeTagRelations", testThemeTagRelationsInsert)
	t.Run("ThemeTagRelations", testThemeTagRelationsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("ChallengeRecordToChallengeThemeUsingChallengeTheme", testChallengeRecordToOneChallengeThemeUsingChallengeTheme)
	t.Run("ChallengeRecordToUserUsingUser", testChallengeRecordToOneUserUsingUser)
	t.Run("ChallengeThemeToUserUsingUser", testChallengeThemeToOneUserUsingUser)
	t.Run("FavoriteToChallengeRecordUsingChallengeRecord", testFavoriteToOneChallengeRecordUsingChallengeRecord)
	t.Run("FavoriteToUserUsingUser", testFavoriteToOneUserUsingUser)
	t.Run("ThemeTagRelationToChallengeThemeUsingChallengeTheme", testThemeTagRelationToOneChallengeThemeUsingChallengeTheme)
	t.Run("ThemeTagRelationToTagUsingTag", testThemeTagRelationToOneTagUsingTag)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("ChallengeRecordToFavorites", testChallengeRecordToManyFavorites)
	t.Run("ChallengeThemeToChallengeRecords", testChallengeThemeToManyChallengeRecords)
	t.Run("ChallengeThemeToThemeTagRelations", testChallengeThemeToManyThemeTagRelations)
	t.Run("TagToThemeTagRelations", testTagToManyThemeTagRelations)
	t.Run("UserToChallengeRecords", testUserToManyChallengeRecords)
	t.Run("UserToChallengeThemes", testUserToManyChallengeThemes)
	t.Run("UserToFavorites", testUserToManyFavorites)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("ChallengeRecordToChallengeThemeUsingChallengeRecords", testChallengeRecordToOneSetOpChallengeThemeUsingChallengeTheme)
	t.Run("ChallengeRecordToUserUsingChallengeRecords", testChallengeRecordToOneSetOpUserUsingUser)
	t.Run("ChallengeThemeToUserUsingChallengeThemes", testChallengeThemeToOneSetOpUserUsingUser)
	t.Run("FavoriteToChallengeRecordUsingFavorites", testFavoriteToOneSetOpChallengeRecordUsingChallengeRecord)
	t.Run("FavoriteToUserUsingFavorites", testFavoriteToOneSetOpUserUsingUser)
	t.Run("ThemeTagRelationToChallengeThemeUsingThemeTagRelations", testThemeTagRelationToOneSetOpChallengeThemeUsingChallengeTheme)
	t.Run("ThemeTagRelationToTagUsingThemeTagRelations", testThemeTagRelationToOneSetOpTagUsingTag)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("ChallengeRecordToFavorites", testChallengeRecordToManyAddOpFavorites)
	t.Run("ChallengeThemeToChallengeRecords", testChallengeThemeToManyAddOpChallengeRecords)
	t.Run("ChallengeThemeToThemeTagRelations", testChallengeThemeToManyAddOpThemeTagRelations)
	t.Run("TagToThemeTagRelations", testTagToManyAddOpThemeTagRelations)
	t.Run("UserToChallengeRecords", testUserToManyAddOpChallengeRecords)
	t.Run("UserToChallengeThemes", testUserToManyAddOpChallengeThemes)
	t.Run("UserToFavorites", testUserToManyAddOpFavorites)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("ChallengeRecords", testChallengeRecordsReload)
	t.Run("ChallengeThemes", testChallengeThemesReload)
	t.Run("Favorites", testFavoritesReload)
	t.Run("Tags", testTagsReload)
	t.Run("ThemeTagRelations", testThemeTagRelationsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("ChallengeRecords", testChallengeRecordsReloadAll)
	t.Run("ChallengeThemes", testChallengeThemesReloadAll)
	t.Run("Favorites", testFavoritesReloadAll)
	t.Run("Tags", testTagsReloadAll)
	t.Run("ThemeTagRelations", testThemeTagRelationsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("ChallengeRecords", testChallengeRecordsSelect)
	t.Run("ChallengeThemes", testChallengeThemesSelect)
	t.Run("Favorites", testFavoritesSelect)
	t.Run("Tags", testTagsSelect)
	t.Run("ThemeTagRelations", testThemeTagRelationsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("ChallengeRecords", testChallengeRecordsUpdate)
	t.Run("ChallengeThemes", testChallengeThemesUpdate)
	t.Run("Favorites", testFavoritesUpdate)
	t.Run("Tags", testTagsUpdate)
	t.Run("ThemeTagRelations", testThemeTagRelationsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("ChallengeRecords", testChallengeRecordsSliceUpdateAll)
	t.Run("ChallengeThemes", testChallengeThemesSliceUpdateAll)
	t.Run("Favorites", testFavoritesSliceUpdateAll)
	t.Run("Tags", testTagsSliceUpdateAll)
	t.Run("ThemeTagRelations", testThemeTagRelationsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
