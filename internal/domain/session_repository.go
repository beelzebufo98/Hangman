package domain

type SessionRepository interface {
	Save(id string, session Session) error
	Load(id string) (Session, error)
}
