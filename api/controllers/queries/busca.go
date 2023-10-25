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

const QUERY_PODE_CURSAR = `
SELECT DISTINCT p.cod_disc_orig AS cod_disciplina
FROM prerequisitos p
WHERE (p.cod_disc_orig, p.grupo_prereq) NOT IN (
    SELECT pp.cod_disc_orig, pp.grupo_prereq
    FROM prerequisitos pp
    WHERE pp.cod_disc_depen NOT IN (
        SELECT cod_disciplina
        FROM historicos h
        WHERE h.cod_usuario = @usuario
    )
)
UNION DISTINCT
SELECT DISTINCT cod_disciplina
FROM disciplinas
WHERE cod_disciplina NOT IN (
    SELECT cod_disc_orig
    FROM prerequisitos
) AND cod_disciplina NOT IN (
    SELECT cod_disciplina
    FROM historicos h
    WHERE h.cod_usuario = @usuario
);`

const QUERY_FALTA_CURSAR = `
SELECT d.cod_disciplina
FROM disciplinas d
         LEFT JOIN semestres s ON d.cod_disciplina = s.cod_disciplina
         LEFT JOIN usuarios u ON s.cod_curriculo = u.cod_curriculo
WHERE NOT EXISTS (SELECT cod_disciplina
                  FROM historicos
                  WHERE cod_usuario = @usuario
                    AND d.cod_disciplina = cod_disciplina)
  AND u.cod_usuario = @usuario;
`
