package entities

type Student struct {
	Id        uint `gorm:"primaryKey;autoIncrement:true;type:integer;default:nextval('seq_students')"`
	Name      string
	Matricula string
	Grades    []Grade
}
