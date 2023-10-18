-- Conteudo = C1 + C2, onde
--  C1 = {0.5: se a disciplina estiver no curriculo do aluno, 0: caso contrario}
--  C2 = 0.5 * #AlunosDoCurriculoQueFizeramADisciplina/#AlunosDoCurriculo
--
-- Horario = #TurmasComHorarioLivres/#TurmasDaDisciplina
--
-- Opiniao = avg(AvaliacaoDaDisciplina) / 5
--
-- Professor = avg(AvaliacaoDoProfessor) / 5
--
-- Avaliacao = avg(GrausDaDisciplina) / 100

-- CONTEUDO
-- CONTEUDO C1
SELECT cod_disciplina
FROM semestres
WHERE cod_curriculo =
      (SELECT cod_curriculo FROM usuarios where cod_usuario = '1910462');
-- CONTEUDO C2.1
SELECT s.cod_disciplina, count(u.cod_usuario) as qtd
FROM semestres s, usuarios u
WHERE s.cod_curriculo = u.cod_curriculo
AND u.cod_curriculo =
    (SELECT cod_curriculo FROM usuarios where cod_usuario = '1910462')
GROUP BY s.cod_disciplina;
-- CONTEUDO C2.2
SELECT count(cod_usuario)
FROM usuarios
WHERE cod_curriculo =
      (SELECT cod_curriculo FROM usuarios where cod_usuario = '1910462');

-- HORARIO
-- retorna somente as disciplinas que
SELECT h1.cod_disciplina, count(DISTINCT h1.cod_turma) as possiveis, todas
FROM horarios h1
         LEFT JOIN (
    SELECT cod_disciplina, count(cod_turma) as todas
    FROM turmas
    GROUP BY cod_disciplina
) AS h2 ON h1.cod_disciplina = h2.cod_disciplina
WHERE (dia, hora_inicio, hora_fim) NOT IN (
    SELECT dia, hora_inicio, hora_fim
    FROM horarios
    WHERE
        (cod_disciplina = 'INF1350' AND cod_turma = '3WA') OR
        (cod_disciplina = 'ENG4402' AND cod_turma = '3VA') OR
        (cod_disciplina = 'INF1383' AND cod_turma = '3WA') OR
        (cod_disciplina = 'INF1771' AND cod_turma = '3WA')
)
GROUP BY h1.cod_disciplina, todas;

-- OPINIAO
-- media de avaliacao das disciplinas
SELECT cod_disciplina, avg(nota_avaliacao) as media
FROM avaliacoes_disciplina
GROUP BY cod_disciplina;


-- PROFESSOR
-- media de avaliacao dos professores por disciplina
SELECT d.cod_disciplina, avg(a.nota_avaliacao) as media
FROM disciplinas d, avaliacoes_professores a, turmas t
WHERE d.cod_disciplina = t.cod_disciplina
AND t.nome_professor = a.nome_professor
GROUP BY d.cod_disciplina, d.nome_disciplina;

-- AVALIACAO
-- media das notas dos alunos
SELECT cod_disciplina, avg(grau)
FROM historicos
GROUP BY cod_disciplina;

-- FILTRO
SELECT d.cod_disciplina
FROM disciplinas d
WHERE cod_disciplina NOT IN (
    SELECT h.cod_disciplina
    FROM historicos h
    WHERE cod_usuario = '1910462'
);


------------------------------------------------
---- QUERY FINAL -------------------------------
WITH rec_c1 AS (
    SELECT cod_disciplina
    FROM semestres
    WHERE cod_curriculo =
          (SELECT cod_curriculo FROM usuarios where cod_usuario = '1910462')
), rec_c2_1 AS (
    SELECT s.cod_disciplina, count(u.cod_usuario) as qtd
    FROM semestres s, usuarios u
    WHERE s.cod_curriculo = u.cod_curriculo
      AND u.cod_curriculo =
          (SELECT cod_curriculo FROM usuarios where cod_usuario = '1910462')
    GROUP BY s.cod_disciplina
), rec_c2_2 AS (
    SELECT count(cod_usuario)
    FROM usuarios
    WHERE cod_curriculo =
          (SELECT cod_curriculo FROM usuarios where cod_usuario = '1910462')
), rec_h AS (
    SELECT h1.cod_disciplina, count(DISTINCT h1.cod_turma) as possiveis, todas
    FROM horarios h1
             LEFT JOIN (
        SELECT cod_disciplina, count(cod_turma) as todas
        FROM turmas
        GROUP BY cod_disciplina
    ) AS h2 ON h1.cod_disciplina = h2.cod_disciplina
    WHERE (dia, hora_inicio, hora_fim) NOT IN (
        SELECT dia, hora_inicio, hora_fim
        FROM horarios
        WHERE
            (cod_disciplina = 'INF1350' AND cod_turma = '3WA') OR
            (cod_disciplina = 'ENG4402' AND cod_turma = '3VA') OR
            (cod_disciplina = 'INF1383' AND cod_turma = '3WA') OR
            (cod_disciplina = 'INF1771' AND cod_turma = '3WA')
    )
    GROUP BY h1.cod_disciplina, todas
), rec_o AS (
    SELECT cod_disciplina, avg(nota_avaliacao) as media
    FROM avaliacoes_disciplina
    GROUP BY cod_disciplina
), rec_p AS (
    SELECT d.cod_disciplina, avg(a.nota_avaliacao) as media
    FROM disciplinas d, avaliacoes_professores a, turmas t
    WHERE d.cod_disciplina = t.cod_disciplina
      AND t.nome_professor = a.nome_professor
    GROUP BY d.cod_disciplina, d.nome_disciplina
), rec_a AS (
    SELECT cod_disciplina, avg(grau) as media
    FROM historicos
    GROUP BY cod_disciplina
), filtro AS (
    SELECT d.cod_disciplina
    FROM disciplinas d
    WHERE cod_disciplina NOT IN (
        SELECT h.cod_disciplina
        FROM historicos h
        WHERE cod_usuario = '1910462'
    )
)

SELECT f.cod_disciplina,
       c1.cod_disciplina IS NOT NULL conteudo1,
       c2_1.qtd conteudo21,
       c2_2.count conteudo22,
       h.possiveis horario1,
       h.todas horario2,
       o.media opiniao,
       p.media professor,
       a.media avaliacao
FROM filtro f
         LEFT JOIN rec_c1 c1 ON f.cod_disciplina = c1.cod_disciplina
         LEFT JOIN rec_c2_1 c2_1 ON f.cod_disciplina = c2_1.cod_disciplina
         LEFT JOIN rec_h h ON f.cod_disciplina = h.cod_disciplina
         LEFT JOIN rec_o o ON f.cod_disciplina = o.cod_disciplina
         LEFT JOIN rec_p p ON f.cod_disciplina = p.cod_disciplina
         LEFT JOIN rec_a a ON f.cod_disciplina = a.cod_disciplina,
     rec_c2_2 c2_2