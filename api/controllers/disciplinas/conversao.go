package disciplinas

import "github.com/Leinadium/puc-disciplinas/api/controllers"

// From converte os resultados das queries em uma RespostaDisciplina
func (r *RespostaDisciplina) From(infos []ResultInfo, turmas []ResultTurmas) {
	if len(infos) == 0 {
		return
	}
	// pega as informacoes basicas da disciplina
	r.CodDisciplina = infos[0].CodDisciplina
	r.NomeDisciplina = infos[0].NomeDisciplina
	r.Ementa = infos[0].Ementa
	r.Creditos = infos[0].Creditos
	r.GrauMedio = controllers.RoundedFloat(infos[0].MediaGrau / 100)
	r.QtdAlunos = infos[0].QtdAlunos
	r.AvaliacaoMedia = controllers.RoundedFloat(infos[0].MediaAvaliacao / 100)
	r.QtdAvaliacoes = infos[0].QtdAvaliacoes

	// pega as informacoes das turmas
	r.Turmas = createTurmas(turmas)

	// pega os pre-requisitos
	// cria um mapa de grupoId -> preReqs
	var preReqs = make(map[int][]string)
	for _, info := range infos {
		if info.GrupoPrereq.Valid && info.CodDiscDepen.Valid {
			preReqs[int(info.GrupoPrereq.Int64)] = append(preReqs[int(info.GrupoPrereq.Int64)], info.CodDiscDepen.String)
		}

	}
	// converte o mapa para um slice
	for grupoId, preReqs := range preReqs {
		var novoPreReq RespostaPreReqs
		novoPreReq.GrupoId = grupoId
		novoPreReq.PreReqs = preReqs
		r.PreRequisitos = append(r.PreRequisitos, novoPreReq)
	}
}

// createTurmas gera as turmas a partir das informações niveladas
func createTurmas(infos []ResultTurmas) []RespostaTurmas {
	// cria tres mapas:
	// um para as turmas em si, contendo os RespostaTurmas
	// e outros dois mapas agindo como sets, para saber se ja foi inserido ou nao
	var turmas = make(map[string]RespostaTurmas) // chave: cod_turma
	var horarios = make(map[string][]struct{})   // chave: cod_turma + dia + hora_inicio
	var alocacoes = make(map[string][]struct{})  // chave: cod_turma + destino

	// itera sobre as informacoes das turmas
	for _, info := range infos {
		// cria uma nova turma se nao existir
		if _, ok := turmas[info.CodTurma]; !ok {
			var novaTurma RespostaTurmas
			novaTurma.CodTurma = info.CodTurma
			novaTurma.NomeProfessor = info.NomeProfessor
			novaTurma.NotaProfessor = controllers.RoundedFloat(info.NotaProfessor / 100)
			novaTurma.Shf = info.Shf
			novaTurma.Horarios = []RespostaHorarios{}
			novaTurma.Alocacoes = []RespostaAlocacoes{}
			turmas[info.CodTurma] = novaTurma
		}

		// cria um novo horario se nao existir
		var chaveHorario = info.CodTurma + info.Dia + string(rune(info.HoraInicio))
		if _, ok := horarios[chaveHorario]; !ok {
			var novoHorario RespostaHorarios
			novoHorario.Dia = info.Dia
			novoHorario.HoraInicio = info.HoraInicio
			novoHorario.HoraFim = info.HoraFim
			// adiciona o horario na turma
			temp := turmas[info.CodTurma]
			temp.Horarios = append(temp.Horarios, novoHorario)
			turmas[info.CodTurma] = temp
			// salva no cache
			horarios[chaveHorario] = append(horarios[chaveHorario], struct{}{})
		}

		// cria uma nova alocacao se nao existir
		var chaveAlocacao = info.CodTurma + info.Destino
		if _, ok := alocacoes[chaveAlocacao]; !ok {
			var novaAlocacao RespostaAlocacoes
			novaAlocacao.Destino = info.Destino
			novaAlocacao.QtdVagas = info.Vagas
			// adiciona a alocacao na turma
			temp := turmas[info.CodTurma]
			temp.Alocacoes = append(temp.Alocacoes, novaAlocacao)
			turmas[info.CodTurma] = temp
			// salva no cache
			alocacoes[chaveAlocacao] = append(alocacoes[chaveAlocacao], struct{}{})
		}
	}

	// pega os resultados dos valores do mapa
	var resultado []RespostaTurmas
	for _, turma := range turmas {
		resultado = append(resultado, turma)
	}

	return resultado
}
