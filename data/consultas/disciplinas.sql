-- FALTA CURSAR ---------------------------------
-- v1 -------------------------------------------
SELECT d.cod_disciplina
FROM disciplinas d
    LEFT JOIN semestres s ON d.cod_disciplina = s.cod_disciplina
    LEFT JOIN usuarios u ON s.cod_curriculo = u.cod_curriculo
WHERE NOT EXISTS (SELECT cod_disciplina
                  FROM historicos
                  WHERE cod_usuario = '1910462'
                  AND d.cod_disciplina = cod_disciplina)
AND u.cod_usuario = '1910462';

-- v2 -------------------------------------------
SELECT d.cod_disciplina
FROM semestres s
    LEFT JOIN usuarios u ON u.cod_curriculo = s.cod_curriculo
    LEFT JOIN disciplinas d ON s.cod_disciplina = d.cod_disciplina
WHERE NOT EXISTS (SELECT cod_disciplina
                  FROM historicos
                  WHERE cod_usuario = '1910462'
                  AND d.cod_disciplina = cod_disciplina)
AND d.cod_disciplina IS NOT NULL
AND u.cod_usuario = '1910462';
------------------------------------------------
------------------------------------------------
-- LISTA LATERAL DAS DISCIPLINAS ---------------
SELECT d.cod_disciplina, d.nome_disciplina, d.creditos, COUNT(t.cod_turma) qtdTurmas, SUM(t.vagas) qtdVagas
FROM disciplinas d
    LEFT JOIN (
        SELECT t.cod_turma, t.cod_disciplina, SUM(a.vagas) vagas
        FROM turmas t LEFT JOIN alocacoes a ON t.cod_turma = a.cod_turma AND t.cod_disciplina = a.cod_disciplina
        GROUP BY t.cod_turma, t.cod_disciplina
    ) t ON d.cod_disciplina = t.cod_disciplina
GROUP BY d.cod_disciplina, d.nome_disciplina, d.creditos;
------------------------------------------------
------------------------------------------------
-- LISTA LATERAL DAS DISCIPLINAS COM USUARIO ---
SELECT d.cod_disciplina, d.nome_disciplina, d.creditos, COUNT(t.cod_turma) qtdTurmas, SUM(t.vagas) qtdVagas
FROM disciplinas d
         LEFT JOIN (
    SELECT t.cod_turma, t.cod_disciplina, SUM(a.vagas) vagas
    FROM turmas t LEFT JOIN alocacoes a on t.cod_turma = a.cod_turma and t.cod_disciplina = a.cod_disciplina
    GROUP BY t.cod_turma, t.cod_disciplina
) t ON d.cod_disciplina = t.cod_disciplina
WHERE d.cod_disciplina IN (
    SELECT d.cod_disciplina
    FROM disciplinas d
             LEFT JOIN semestres s ON d.cod_disciplina = s.cod_disciplina
             LEFT JOIN usuarios u ON s.cod_curriculo = u.cod_curriculo
    WHERE NOT EXISTS (SELECT cod_disciplina
                      FROM historicos
                      WHERE cod_usuario = '1910462'
                      AND d.cod_disciplina = cod_disciplina)
    AND u.cod_usuario = '1910462'
)
GROUP BY d.cod_disciplina, d.nome_disciplina, d.creditos
------------------------------------------------
------------------------------------------------