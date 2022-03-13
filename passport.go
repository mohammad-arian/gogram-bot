package gogram

import "errors"

type PasswordData struct {
	data        []EncryptedPassportElement
	credentials EncryptedCredentials
}

type EncryptedCredentials struct {
	// Base64-encoded encrypted JSON-serialized data with unique user's payload,
	// data hashes and secrets required for EncryptedPassportElement decryption and authentication
	Data string `json:"data"`
	// Base64-encoded data hash for data authentication
	Hash string `json:"hash"`
	// Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
	Secret string `json:"secret"`
}

type PassportFile struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FileDate     int    `json:"file_date"`
}

type EncryptedPassportElement struct {
	Type        string         `json:"type"`
	Data        string         `json:"data"`
	PhoneNumber string         `json:"phone_number"`
	Email       string         `json:"email"`
	Files       []PassportFile `json:"files"`
	FrontSide   PassportFile   `json:"front_side"`
	ReverseSide PassportFile   `json:"reverse_side"`
	Selfie      PassportFile   `json:"selfie"`
	Translation []PassportFile `json:"translation"`
	Hash        string         `json:"hash"`
}

type passport interface {
	checkPassport() error
}

type PassportBase struct {
	Source  string `json:"source"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

type PassportElementErrorDataField struct {
	PassportBase
	FieldName string `json:"field_name"`
	DataHash  string `json:"data_hash"`
}

func (p PassportElementErrorDataField) checkPassport() error {
	if p.Source != "data" {
		return errors.New("source must be `data`")
	}
	var types = map[string]bool{"personal_details": true, "passport": true, "driver_license": true, "identity_card": true,
		"internal_passport": true, "address": true}
	if _, ok := types[p.Type]; ok == false {
		return errors.New(p.Type + " is an unknown action, read the document")
	}
	return nil
}

type PassportElementErrorFrontSide struct {
	PassportBase
	FileHash string `json:"file_hash"`
}

func (p PassportElementErrorFrontSide) checkPassport() error {
	if p.Source != "front_side" {
		return errors.New("source must be `front_side`")
	}
	var types = map[string]bool{"passport": true, "driver_license": true, "identity_card": true,
		"internal_passport": true}
	if _, ok := types[p.Type]; ok == false {
		return errors.New(p.Type + " is an unknown action, read the document")
	}
	return nil
}

type PassportElementErrorReverseSide struct {
	PassportBase
	FileHash string `json:"file_hash"`
}

func (p PassportElementErrorReverseSide) checkPassport() error {
	if p.Source != "reverse_side" {
		return errors.New("source must be `reverse_side`")
	}
	var types = map[string]bool{"driver_license": true, "identity_card": true}
	if _, ok := types[p.Type]; ok == false {
		return errors.New(p.Type + " is an unknown action, read the document")
	}
	return nil
}

type PassportElementErrorSelfie struct {
	PassportBase
	FileHash string `json:"file_hash"`
}

func (p PassportElementErrorSelfie) checkPassport() error {
	if p.Source != "selfie" {
		return errors.New("source must be `selfie`")
	}
	var types = map[string]bool{"passport": true, "driver_license": true, "identity_card": true,
		"internal_passport": true}
	if _, ok := types[p.Type]; ok == false {
		return errors.New(p.Type + " is an unknown action, read the document")
	}
	return nil
}

type PassportElementErrorFile struct {
	PassportBase
	FileHash string `json:"file_hash"`
}

func (p PassportElementErrorFile) checkPassport() error {
	if p.Source != "file" {
		return errors.New("source must be `file`")
	}
	var types = map[string]bool{"utility_bill": true, "bank_statement": true, "rental_agreement": true,
		"passport_registration": true, "temporary_registration": true}
	if _, ok := types[p.Type]; ok == false {
		return errors.New(p.Type + " is an unknown action, read the document")
	}
	return nil
}

type PassportElementErrorFiles struct {
	PassportBase
	FileHashes []string `json:"file_hashes"`
}

func (p PassportElementErrorFiles) checkPassport() error {
	if p.Source != "files" {
		return errors.New("source must be `files`")
	}
	var types = map[string]bool{"utility_bill": true, "bank_statement": true, "rental_agreement": true,
		"passport_registration": true, "temporary_registration": true}
	if _, ok := types[p.Type]; ok == false {
		return errors.New(p.Type + " is an unknown action, read the document")
	}
	return nil
}

type PassportElementErrorTranslationFile struct {
	PassportBase
	FileHash string `json:"file_hash"`
}

func (p PassportElementErrorTranslationFile) checkPassport() error {
	if p.Source != "translation_file" {
		return errors.New("source must be `translation_file`")
	}
	var types = map[string]bool{"passport": true, "driver_license": true, "identity_card": true,
		"utility_bill": true, "bank_statement": true, "rental_agreement": true, "passport_registration": true,
		"temporary_registration": true, "internal_passport": true}
	if _, ok := types[p.Type]; ok == false {
		return errors.New(p.Type + " is an unknown action, read the document")
	}
	return nil
}

type PassportElementErrorTranslationFiles struct {
	PassportBase
	FileHashes []string `json:"file_hashes"`
}

func (p PassportElementErrorTranslationFiles) checkPassport() error {
	if p.Source != "translation_files" {
		return errors.New("source must be `translation_files`")
	}
	var types = map[string]bool{"passport": true, "driver_license": true, "identity_card": true,
		"utility_bill": true, "bank_statement": true, "rental_agreement": true, "passport_registration": true,
		"temporary_registration": true, "internal_passport": true}
	if _, ok := types[p.Type]; ok == false {
		return errors.New(p.Type + " is an unknown action, read the document")
	}
	return nil
}

type PassportElementErrorUnspecified struct {
	PassportBase
	ElementHash string `json:"element_hash"`
}

func (p PassportElementErrorUnspecified) checkPassport() error {
	if p.Source != "unspecified" {
		return errors.New("source must be `unspecified`")
	}
	return nil
}

type SetPassportDataErrors struct {
	// user identifier
	ChatId int `json:"user_id"`
	// an array describing the errors
	Errors []passport `json:"errors"`
}

func (s SetPassportDataErrors) Send(b Bot) (response Response, err error) {
	return Request("setPassportDataErrors", b, s, &ResponseImpl{})
}

func (s SetPassportDataErrors) Check() error {
	for _, j := range s.Errors {
		if err := j.checkPassport(); err != nil {
			return err
		}
	}
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": s.ChatId, "Errors": s.Errors})
}
