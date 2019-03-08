package alarme

import (
	"log"

	repository "../../infraestructure/repository"
)

const tabelaAlarme string = "ALARMES"

// Alarme strutura de alarme
type Alarme struct {
	Codigo           int    `db:"CODIGO"`
	Descricao        string `db:"DESCRICAO"`
	CodigoEmpresa    int    `db:"CODIGOEMPRESA"`
	Modelo           string `db:"MODELO"`
	Apagado          string `db:"APAGADO"`
	CodigoTecnologia int    `db:"CODIGOTECNOLOGIA"`
	CodigoOperacoes  int    `db:"CODIGOOPERACOES"`
}

// GetAll busca todos os alarmes
func (a Alarme) GetAll() []Alarme {
	r := repository.RepositoryFirebird{}

	str := "SELECT CODIGO, DESCRICAO, CODIGOEMPRESA, MODELO, APAGADO, CODIGOTECNOLOGIA, CODIGOOPERACOES FROM " + tabelaAlarme
	list := []Alarme{}
	err := r.Query(&list, str)

	if err != nil {
		log.Fatalln(err)
	}
	return list

}
