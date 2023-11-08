package recomendacao

import "github.com/Leinadium/puc-disciplinas/api/controllers"

// executaAlgoritmo executa o algoritmo de recomendacao
// Pesos:
//
//	Conteudo = {
//		1: se a disciplina estiver no curriculo do aluno,
//		0.5 * #AlunosDoCurriculoQueFizeramADisciplina/#AlunosDoCurriculo: caso contrario
//	}
//
// Horario = #TurmasComHorarioLivres/#TurmasDaDisciplina
//
// Opiniao = avg(AvaliacaoDaDisciplina) / 5
//
// Professor = avg(AvaliacaoDoProfessor) / 5
//
// Avaliacao = avg(GrausDaDisciplina) / 100
//
// Valor final: (Conteudo * 41) + (Horario * 19) + (Opiniao * 8) + (Professor * 28) + (Avaliacao * 21)
func executaAlgoritmo(discs []resultQuery) []ResultAlg {
	var res []ResultAlg
	var nota float64
	for _, disc := range discs {
		nota = 0
		var notaC, notaH, notaO, notaP, notaA float64
		if disc.Existe {
			// se a disciplina existe, calcula os valores
			notaC = calculaConteudo(disc.Conteudo1, disc.Conteudo21, disc.Conteudo22)
			notaH = calculaHorario(disc.Horario1, disc.Horario2)
			notaO = calculaOpiniao(disc.Opiniao)
			notaP = calculaProfessor(disc.Professor)
			notaA = calculaAvaliacao(disc.Avaliacao)
			// faz a media ponderada
			// os pesos foram derivados das entrevistas
			nota = (notaC * 41.0) + (notaH * 19.0) + (notaO * 8.0) + (notaP * 28.0) + (notaA * 21.0)
			nota /= 41 + 19 + 8 + 28 + 21.0
		}

		// salva o resultado
		res = append(res, ResultAlg{
			CodDisciplina: disc.CodDisciplina,
			Valor:         controllers.RoundedFloat(nota),
			Pesos: ResultPesos{
				Conteudo:  controllers.RoundedFloat(notaC),
				Horario:   controllers.RoundedFloat(notaH),
				Opiniao:   controllers.RoundedFloat(notaO),
				Professor: controllers.RoundedFloat(notaP),
				Avaliacao: controllers.RoundedFloat(notaA),
			},
		})
	}
	return res
}

// calculaConteudo calcula o valor de conteudo
//
// C1 = {0.5: se a disciplina estiver no curriculo do aluno, 0: caso contrario}
//
// C2 = 0.5 * #AlunosDoCurriculoQueFizeramADisciplina / #AlunosDoCurriculo
func calculaConteudo(c1 bool, c21 int, c22 int) float64 {
	if c1 {
		return 1.0
	} else if c22 != 0 {
		return float64(c21) / float64(c22)
	}
	return 0.0
}

// calculaHorario calcula o valor de horario
//
// Horario = #TurmasComHorarioLivres / #TurmasDaDisciplina
func calculaHorario(h1, h2 int) float64 {
	if h2 != 0 {
		return float64(h1) / float64(h2)
	}
	return 0.0
}

// calculaOpiniao calcula o valor de opiniao
//
// Opiniao = avg(AvaliacaoDaDisciplina) / 5
func calculaOpiniao(o float64) float64 {
	return o / 5
}

// calculaProfessor calcula o valor de professor
//
// Professor = avg(AvaliacaoDoProfessor) / 5
func calculaProfessor(p float64) float64 {
	return p / 5
}

// calculaAvaliacao calcula o valor de avaliacao
//
// Avaliacao = avg(GrausDaDisciplina) / 100
func calculaAvaliacao(a float64) float64 {
	return a / 100
}
