package queries

const QUERY_DISCIPLINA_INFO = `
WITH medias_grau AS (
    SELECT
        cod_disciplina,
        AVG(grau) AS media_grau,
        COUNT(cod_disciplina) AS qtd_alunos
    FROM historicos
    GROUP BY cod_disciplina
), medias_avaliacao AS (
    SELECT
        cod_disciplina,
        AVG(nota_avaliacao) AS media_avaliacao,
        COUNT(cod_disciplina) AS qtd_avaliacoes
    FROM avaliacoes_disciplinas
    GROUP BY cod_disciplina
)
SELECT
    d.cod_disciplina,
    d.nome_disciplina,
    d.creditos,
    d.ementa,
    p.grupo_prereq,
    p.cod_disc_depen,
    g.media_grau,
    g.qtd_alunos,
    a.media_avaliacao,
    a.qtd_avaliacoes
FROM disciplinas d
    LEFT JOIN prerequisitos p ON d.cod_disciplina = p.cod_disc_orig
    LEFT JOIN medias_grau g ON d.cod_disciplina = g.cod_disciplina
    LEFT JOIN medias_avaliacao a ON d.cod_disciplina = a.cod_disciplina
WHERE d.cod_disciplina = @disciplina;
`

const QUERY_TURMAS_INFO = `
SELECT
    t.cod_turma,
    t.nome_professor,
    t.shf,
    h.dia,
    h.hora_inicio,
    h.hora_fim,
    a.destino,
    a.vagas
FROM turmas t
    LEFT JOIN horarios h ON
        t.cod_disciplina = h.cod_disciplina AND
        t.cod_turma = h.cod_turma
    LEFT JOIN alocacoes a ON
        h.cod_disciplina = a.cod_disciplina AND
        h.cod_turma = a.cod_turma
WHERE t.cod_disciplina = @disciplina;
`
