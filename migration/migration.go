package migration

func Migrate() {
	KabupatenMigrate()
	ProvinsiMigrate()
	MenuMigrate()
	PakaianAdatMigrate()
	ProvinsiMigrate()
	RoleMigrate()
	UserMigrate()
	UserRoleMigrate()
	RoleMenuMigrate()
	TokenMigrate()
	SessionMigrate()
}
