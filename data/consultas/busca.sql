SELECT d.cod_disciplina, d.nome_disciplina, d.creditos, COUNT(t.cod_turma) qtdTurmas, SUM(t.vagas) qtdVagas
FROM disciplinas d
         LEFT JOIN (
    SELECT t.cod_turma, t.cod_disciplina, SUM(a.vagas) vagas
    FROM turmas t LEFT JOIN alocacoes a ON t.cod_turma = a.cod_turma AND t.cod_disciplina = a.cod_disciplina
    GROUP BY t.cod_turma, t.cod_disciplina
) t ON d.cod_disciplina = t.cod_disciplina
WHERE d.cod_disciplina || d.nome_disciplina ILIKE '%XXX%'
GROUP BY d.cod_disciplina, d.nome_disciplina, d.creditos
LIMIT 20;