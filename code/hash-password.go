const (
	argonTime        = 1
	argonMemoryUsed  = 65536 // 64 * 1024
	argonThreadsUsed = 1
	argonKeyLen      = 32
)

func (uc *useCase) hashPassword(password string, salt []byte) (hashedPassword []byte) {
	key := argon2.IDKey([]byte(password), salt, argonTime, argonMemoryUsed, argonThreadsUsed, argonKeyLen)
	return key
}