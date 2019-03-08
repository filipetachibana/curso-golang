package main

import (
	"fmt"

	configuration "./configuration"
	repository "./infraestructure/repository"
	AlarmeBusiness "./structs/alarme"
	ConvocacoesAlarmeBusiness "./structs/vinculosconvocacoesalarmes"
)

func init() {
	app := configuration.GetConfiguration()
	repository.SetDbFirebird(app.ConnectionString)
}

func main() {
	getConvocacoesAlarmes()
}

func getAlarmes() {
	var alarme = AlarmeBusiness.Alarme{}
	lista := alarme.GetAll()

	for _, a := range lista {
		fmt.Println(a)
	}
}

func getConvocacoesAlarmes() {
	listaCodigosAlarmes := []int{2155445, 2155446, 2155447, 2155448, 2155452, 2155453, 2155456, 2155457, 2155458, 2155459, 2155460, 2155454, 2155461}

	var convovacoesAlarmes = ConvocacoesAlarmeBusiness.VinculosConvocacoesAlarmes{}

	listaConvocacoesAlarmes := convovacoesAlarmes.GetByCodigosAlarmes(listaCodigosAlarmes...)

	for _, c := range listaConvocacoesAlarmes {
		fmt.Println(c)
	}
}

func getAllConvocacoesAlarmes() {
	var convovacoesAlarmes = ConvocacoesAlarmeBusiness.VinculosConvocacoesAlarmes{}

	listaConvocacoesAlarmes := convovacoesAlarmes.GetAll()

	for _, c := range listaConvocacoesAlarmes {
		fmt.Println(c)
	}

}
