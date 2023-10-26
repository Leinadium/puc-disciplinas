package queries

const QUERY_AVALIACOES_DISCIPLINA = `
SELECT d.cod_disciplina, d.nome_disciplina, a.nota_avaliacao, m.media, m.quantidade
FROM disciplinas d 
    LEFT JOIN (
		SELECT cod_disciplina, nota_avaliacao
		FROM avaliacoes_disciplinas
		WHERE cod_usuario = @usuario
	) a ON d.cod_disciplina = a.cod_disciplina
	LEFT JOIN (
	    SELECT cod_disciplina, avg(nota_avaliacao) as media, count(cod_disciplina) as quantidade
	    FROM avaliacoes_disciplinas
	    GROUP BY cod_disciplina
	) m ON d.cod_disciplina = m.cod_disciplina
;
`

const QUERY_AVALIACOES_PROFESSOR = `
SELECT p.nome_professor, a.nota_avaliacao, m.media, m.quantidade
FROM professores p 
    LEFT JOIN (
		SELECT nome_professor, nota_avaliacao
		FROM avaliacoes_professores
		WHERE cod_usuario = @usuario
	) a on p.nome_professor = a.nome_professor
	LEFT JOIN (
	    SELECT nome_professor, avg(nota_avaliacao) as media, count(nome_professor) as quantidade
	    FROM avaliacoes_professores
	    GROUP BY nome_professor
	) m on p.nome_professor = m.nome_professor
;
`
