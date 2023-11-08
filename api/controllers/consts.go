package controllers

import "os"

// nome da chave da sessao do login
const keyLoginSession = "usuario"
const keyNomeSession = "nome"

func getEnvDefault(envName, envDefault string) string {
	var v string
	v, ok := os.LookupEnv(envName)
	if !ok {
		v = envDefault
	}
	return v
}

var urlMicroLogin = getEnvDefault("URL_MICRO_LOGIN", "http://localhost:5000")
var urlMicroHistorico = getEnvDefault("URL_MICRO_HISTORICO", "http://localhost:5001")
