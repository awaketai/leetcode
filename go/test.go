package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main()  {
	body := `{"idade":71,"nome":"CANDIDO GRACIANO FILHO","dataNascimento":"1949-12-07T00:00:00.000-02:00","signo":"Sagitário","nomeMae":"LICTICIA PIASSA CORREA DA COSTA","enderecos":{"endereco":[{"uf":"MG","cidade":"SANTA ROSA DA SERRA","idFinalidade":0,"logradouro":{"Numero":168,"Tipo":"R","CepNota":9,"Titulo":"DR","Nome":"ADOLFO PORTELA"},"cep":38805000},{"uf":"MG","cidade":"SANTA ROSA DA SERRA","idFinalidade":0,"logradouro":{"Numero":186,"Tipo":"R","CepNota":9,"Titulo":"DR","Nome":"ADOLFO PORTELA"},"cep":38805000},{"uf":"MG","cidade":"LUZ","idFinalidade":0,"logradouro":{"Numero":1141,"Tipo":"R","CepNota":9,"Titulo":"VIG","Nome":"PARREIRAS"},"bairro":"CENTRO","cep":35595000}]},"situacaoCadastral":{"situacao":"REGULAR","fontePesquisada":"HISTORICO","nome":"CANDIDO GRACIANO FILHO","dataConsulta":"2011-10-12T00:00:00.000-03:00","codigoControle":"0F96.DE7D.28B5.44AA","codigoSituacao":1},"finalidadeCbo":0,"cpf":17803217168,"inibir":0,"telefones":{"telefone":{"numero":34213258,"ddd":37,"finalidade":0}}}`
	t := "situacaoCadastral.situacao"

	tArr := strings.Split(t,".")
	var data map[string]interface{}
	err := json.Unmarshal([]byte(body),&data)
	//fmt.Println(data)
	if err != nil {
		fmt.Println("unmarshal:",err)
		return
	}

	l := len(tArr)
	//m := JSONToMap(body)
	for k,v := range tArr {
		//t := reflect.TypeOf(data[v])
		 _,ok := data[v]
		 if ok && k != l - 1 {
			 data = data[v].(map[string]interface{})
		 }else {
		 	//data = data[v].(string)
		 }
	}
	fmt.Println(data)
	//var x float32 = 3.4
	//fmt.Println("type:",reflect.TypeOf(x))
}

// json转Map ()
func JSONToMap(str string) map[string]interface{} {

	var tempMap = make(map[string]interface{})

	err := json.Unmarshal([]byte(str), &tempMap)

	if err != nil {
		panic(err)
	}

	return tempMap
}

func t()  {
	mainMapC := map[string]interface{}{}
	subMapC := make(map[string]string)
	subMapC["C_Key_1"] = "C_SubValue_1"
	subMapC["C_Key_2"] = "C_SubValue_2"
	mainMapC["MapC"] = subMapC
	fmt.Println("\nMultityMapC")
	for keyC, valC := range mainMapC {
		//此处必须实例化接口类型，即*.(map[string]string)
		//subMap := valC.(map[string]string)
		for subKeyC, subValC := range valC.(map[string]string) {
			fmt.Printf("mapName=%s Key=%s Value=%s\n", keyC, subKeyC, subValC)
		}
	}
}

