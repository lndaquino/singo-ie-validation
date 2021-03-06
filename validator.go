package singo

import (
	"regexp"
	"strings"

	"errors"

	"github.com/dilowagner/singo-ie-validation/validators"
)

// Constantes
const (
	Acre            = "AC" // AC Acre
	Alagoas         = "AL" // AL Alagoas
	Amapa           = "AP" // AP Amapá
	Amazonas        = "AM" // AM Amazonas
	Bahia           = "BA" // BA Bahia
	Ceara           = "CE" // CE Ceará
	DistritoFederal = "DF" // DF Distrito Federal
	EspiritoSanto   = "ES" // ES Espirito Santo
	Goais           = "GO" // GO Goias
	Maranhao        = "MA" // MA Maranhão
	MatoGrosso      = "MT" // MT Mato Grosso
	MatoGrossoSul   = "MS" // MS Mato Grosso do Sul
	MinasGerais     = "MG" // MG Minas Gerais
	Para            = "PA" // PA Pará
	Paraiba         = "PB" // PB Paraíba
	Parana          = "PR" // PR Paraná
	Pernambuco      = "PE" // PE Pernambuco
	Piaui           = "PI" // PI Piauí
	RioJaneiro      = "RJ" // RJ Rio de Janeiro
	RioGrandeNorte  = "RN" // RN Rio Grande do Norte
	RioGrandeSul    = "RS" // RS Rio Grande do Sul
	Rondonia        = "RO" // RO Rondônia
	Roraima         = "RR" // RR Roraima
	SantaCatarina   = "SC" // SC Santa Catarina
	SaoPaulo        = "SP" // SP São Paulo
	Sergipe         = "SE" // SE Sergipe
	Tocantins       = "TO" // TO Tocantins
)

// ValidatableManager interface
type ValidatableManager interface {
	Validate() (bool, error)
}

// IEValidator struct
type IEValidator struct {
	IE string
	UF string
}

// NewIEValidator create the instance for IEValidator
func NewIEValidator() *IEValidator {
	return &IEValidator{}
}

// Validate func
func (v IEValidator) Validate() (bool, error) {

	var validator validators.Validatable

	insc := v.filter(v.IE)
	if len(insc) == 0 {
		return false, errors.New("Inscrição estadual inválida")
	}
	uf := strings.ToUpper(v.UF)

	switch uf {
	case Acre:
		validator = validators.AC{}
	case Alagoas:
		validator = validators.AL{}
	case Amapa:
		validator = validators.AP{}
	case Amazonas:
		validator = validators.AM{}
	case Bahia:
		validator = validators.BA{}
	case Ceara:
		validator = validators.CE{}
	case DistritoFederal:
		validator = validators.DF{}
	case EspiritoSanto:
		validator = validators.ES{}
	case Goais:
		validator = validators.GO{}
	case Maranhao:
		validator = validators.MA{}
	case MatoGrosso:
		validator = validators.MT{}
	case MatoGrossoSul:
		validator = validators.MS{}
	case MinasGerais:
		validator = validators.MG{}
	case Para:
		validator = validators.PA{}
	case Paraiba:
		validator = validators.PB{}
	case Parana:
		validator = validators.PR{}
	case Pernambuco:
		validator = validators.PE{}
	case Piaui:
		validator = validators.PI{}
	case RioJaneiro:
		validator = validators.RJ{}
	case RioGrandeNorte:
		validator = validators.RN{}
	case RioGrandeSul:
		validator = validators.RS{}
	case Rondonia:
		validator = validators.RO{}
	case Roraima:
		validator = validators.RR{}
	case SantaCatarina:
		validator = validators.SC{}
	case SaoPaulo:
		validator = validators.SP{}
	case Sergipe:
		validator = validators.SE{}
	case Tocantins:
		validator = validators.TO{}
	default:
		return false, errors.New("UF inválida, verifique o estado passado por parâmetro!")
	}

	return validator.IsValid(insc), nil
}

func (v IEValidator) filter(ie string) string {

	regex, _ := regexp.Compile("[^P0-9]+")
	return regex.ReplaceAllString(strings.TrimSpace(strings.ToUpper(v.IE)), "")
}
