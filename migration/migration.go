package migration

func Migrate() {
	KabupatenMigrate()
	KotaMigrate()
	UserMigrate()
	PakaianAdatMigrate()
	ProvinsiMigrate()
	RoleMigrate()
}
