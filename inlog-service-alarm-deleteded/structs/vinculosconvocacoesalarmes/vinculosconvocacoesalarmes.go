package vinculosconvocacoesalarmes

import (
	"database/sql"
	"log"

	extensions "../../infraestructure/extensions"
	repository "../../infraestructure/repository"
	"github.com/jmoiron/sqlx"
)

const tabelaVinculosConvocacoesAlarmes string = "VINCULOS_CONVOCACOES_ALARMES"

// VinculosConvocacoesAlarmes estrutura de vinculos entre convocação e alarmes
type VinculosConvocacoesAlarmes struct {
	Codigo                int                 `db:"CODIGO"`
	CodigoAlarmeEvento    int                 `db:"CODIGO_ALARME_EVENTO"`
	CodigoAlarme          int                 `db:"CODIGO_ALARME"`
	CodigoConvocacao      int                 `db:"CODIGO_CONVOCACAO"`
	IdentificadorVeiculo  string              `db:"IDENTIFICADOR_VEICULO"`
	DataAlarme            extensions.NullTime `db:"DATA_ALARME"`
	Ocorrencia            string              `db:"OCORRENCIA"`
	Velocidade            string              `db:"VELOCIDADE"`
	Reincidente           int                 `db:"REINCIDENTE"`
	Status                string              `db:"STATUS"`
	CodigoFCV             sql.NullInt64       `db:"CODIGO_FCV"`
	CodigoEmpresa         int                 `db:"CODIGO_EMPRESA"`
	CodigoVeiculo         int                 `db:"CODIGO_VEICULO"`
	Apagado               string              `db:"APAGADO"`
	DataExclusao          extensions.NullTime `db:"DATA_EXCLUSAO"`
	InformacoesAdicionais string              `db:"INFORMACOES_ADICIONAIS"`
	CodigoOperacaoVeiculo int                 `db:"CODIGO_OPERACAO_VEICULO"`
}

// GetAll busca todas as vinculações entre Convocações e alarmes
func (v VinculosConvocacoesAlarmes) GetAll() []VinculosConvocacoesAlarmes {
	r := repository.RepositoryFirebird{}

	str := `SELECT FIRST 2 
		CODIGO, CODIGO_ALARME_EVENTO, CODIGO_ALARME, CODIGO_CONVOCACAO, IDENTIFICADOR_VEICULO, DATA_ALARME, OCORRENCIA, VELOCIDADE, 
		REINCIDENTE, STATUS, CODIGO_FCV, CODIGO_EMPRESA, CODIGO_VEICULO, APAGADO, DATA_EXCLUSAO, INFORMACOES_ADICIONAIS, CODIGO_OPERACAO_VEICULO
	FROM ` + tabelaVinculosConvocacoesAlarmes

	list := []VinculosConvocacoesAlarmes{}
	err := r.Query(&list, str)

	if err != nil {
		log.Fatalln(err)
	}
	return list
}

// GetByCodigosAlarmes busca vinculações entre Convocações e alarmes por lista de códigos de alarmes
func (v VinculosConvocacoesAlarmes) GetByCodigosAlarmes(codigoAlarmes ...int) []VinculosConvocacoesAlarmes {
	r := repository.RepositoryFirebird{}

	str := `SELECT CODIGO, CODIGO_ALARME_EVENTO, CODIGO_ALARME, CODIGO_CONVOCACAO, IDENTIFICADOR_VEICULO, DATA_ALARME, OCORRENCIA, VELOCIDADE, REINCIDENTE, STATUS, CODIGO_FCV, CODIGO_EMPRESA, CODIGO_VEICULO, APAGADO, DATA_EXCLUSAO, INFORMACOES_ADICIONAIS, CODIGO_OPERACAO_VEICULO
	FROM ` + tabelaVinculosConvocacoesAlarmes + "WHERE CODIGO_ALARME IN (:collection)"

	list := []VinculosConvocacoesAlarmes{}

	arg := map[string]interface{}{
		"collection": codigoAlarmes,
	}

	query, args, err := sqlx.Named(str, arg)
	query, args, err = sqlx.In(query, args...)
	query = r.Conn.Rebind(query)
	r.Query(&list, query, args...)
	if err != nil {
		log.Fatalln(err)
	}
	return list
}
