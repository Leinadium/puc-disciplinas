package queries

const QUERY_PESQUISA_LATERAL = `
SELECT d.cod_disciplina, d.nome_disciplina, d.creditos, COUNT(t.cod_turma) qtdTurmas, SUM(t.vagas) qtdVagas
FROM disciplinas d
         LEFT JOIN (
    SELECT t.cod_turma, t.cod_disciplina, SUM(a.vagas) vagas
    FROM turmas t LEFT JOIN alocacoes a ON t.cod_turma = a.cod_turma AND t.cod_disciplina = a.cod_disciplina
    GROUP BY t.cod_turma, t.cod_disciplina
) t ON d.cod_disciplina = t.cod_disciplina
WHERE d.cod_disciplina || d.nome_disciplina ILIKE @query
GROUP BY d.cod_disciplina, d.nome_disciplina, d.creditos
`

const QUERY_INFORMACOES = `
SELECT d.cod_disciplina, d.nome_disciplina, d.creditos, COUNT(t.cod_turma) qtdTurmas, SUM(t.vagas) qtdVagas
FROM disciplinas d
    LEFT JOIN (
        SELECT t.cod_turma, t.cod_disciplina, SUM(a.vagas) vagas
        FROM turmas t LEFT JOIN alocacoes a ON t.cod_turma = a.cod_turma AND t.cod_disciplina = a.cod_disciplina
        GROUP BY t.cod_turma, t.cod_disciplina
    ) t ON d.cod_disciplina = t.cod_disciplina
GROUP BY d.cod_disciplina, d.nome_disciplina, d.creditos
`

//const QUERY_PODE_CURSAR = `
//SELECT d.cod_disciplina as codigo, (h.cod_disciplina IS NOT NULL) as cursada
//FROM disciplinas d
//    LEFT JOIN (
//        SELECT h.cod_disciplina
//        FROM historicos h
//        WHERE h.cod_usuario = @name
//    ) h ON d.cod_disciplina = h.cod_disciplina;
//`

const QUERY_FALTA_CURSAR = `
SELECT d.cod_disciplina
FROM disciplinas d
         LEFT JOIN semestres s ON d.cod_disciplina = s.cod_disciplina
         LEFT JOIN usuarios u ON s.cod_curriculo = u.cod_curriculo
WHERE NOT EXISTS (SELECT cod_disciplina
                  FROM historicos
                  WHERE cod_usuario = @name
                    AND d.cod_disciplina = cod_disciplina)
  AND u.cod_usuario = @name;
`
