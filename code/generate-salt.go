const passwordSaltBytes = 16

func (uc *useCase) generateSalt() (salt []byte, err error) {
	salt = make([]byte, passwordSaltBytes)
	_, err = rand.Read(salt)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error generating salt: %s", err.Error()))
	}
	return salt, nil
}