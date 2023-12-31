-- insere curriculo se ja nao existir
INSERT INTO curriculos (cod_curriculo, nome_curriculo)
SELECT * FROM (
    VALUES ('CEGBCO20231', 'Eng. de Computação (Bacharelado) - Currículo 23.1')
) AS i (cod_curriculo, nome_curriculo)
WHERE NOT EXISTS (
    SELECT cod_curriculo
    FROM curriculos
    WHERE cod_curriculo = i.cod_curriculo
);

-- procedure de inserir os curriculos se existir
CREATE OR REPLACE PROCEDURE inserir_curriculo_eng_23_0()
LANGUAGE SQL
BEGIN ATOMIC
    DELETE FROM semestres WHERE cod_curriculo = 'CEGBCO20231';

    INSERT INTO semestres (cod_curriculo, cod_disciplina, semestre)
    SELECT * FROM (VALUES
        ('CEGBCO20231', 'ENG4010', 1),
        ('CEGBCO20231', 'EMP1310', 1),
        ('CEGBCO20231', 'ENG4025', 1),
        ('CEGBCO20231', 'CTC4001', 1),
        ('CEGBCO20231', 'CTC4002', 1),
        ('CEGBCO20231', 'CTC4003', 1),
        ('CEGBCO20231', 'CRE1200', 1),

        ('CEGBCO20231', 'ENG4021', 2),
        ('CEGBCO20231', 'EMP1320', 2),
        ('CEGBCO20231', 'INF1037', 2),
        ('CEGBCO20231', 'MAT4161', 3),
        ('CEGBCO20231', 'FIS4001', 2),
        ('CEGBCO20231', 'MAT4200', 2),
        ('CEGBCO20231', 'INF1012', 2),
        ('CEGBCO20231', 'INF1009', 2),

        ('CEGBCO20231', 'ENG4033', 3),
        ('CEGBCO20231', 'ENG4431', 3),
        ('CEGBCO20231', 'ENG4420', 3),
        ('CEGBCO20231', 'MAT4162', 3),
        ('CEGBCO20231', 'FIS4002', 3),
        ('CEGBCO20231', 'MAT4202', 3),
        ('CEGBCO20231', 'INF1383', 3),
        ('CEGBCO20231', 'ENG4007', 3),

        ('CEGBCO20231', 'ENG4040', 4),
        ('CEGBCO20231', 'ENG4502', 4),
        ('CEGBCO20231', 'ENG4501', 4),
        ('CEGBCO20231', 'ENG4402', 4),
        ('CEGBCO20231', 'MAT4174', 4),
        ('CEGBCO20231', 'INF1010', 4),
        ('CEGBCO20231', 'INF1018', 4),
        ('CEGBCO20231', 'ENG4011', 4),

        ('CEGBCO20231', 'ENG4051', 5),
        ('CEGBCO20231', 'ENG4013', 5),
        ('CEGBCO20231', 'ENG4421', 5),
        ('CEGBCO20231', 'MAT1320', 5),
        ('CEGBCO20231', 'FIS4003', 5),
        ('CEGBCO20231', 'INF1301', 5),

        ('CEGBCO20231', 'ENG4061', 6),
        ('CEGBCO20231', 'ENG4448', 6),
        ('CEGBCO20231', 'INF1022', 6),
        ('CEGBCO20231', 'INF1631', 6),
        ('CEGBCO20231', 'INF1636', 6),
        ('CEGBCO20231', 'INF1316', 6),
        ('CEGBCO20231', 'CRE0712', 6),

        ('CEGBCO20231', 'ENG4405', 7),
        ('CEGBCO20231', 'INF1721', 7),
        ('CEGBCO20231', 'INF1041', 7),
        ('CEGBCO20231', 'INF1029', 7),
        ('CEGBCO20231', 'INF1640', 7),
        ('CEGBCO20231', 'FIL0300', 7),
        ('CEGBCO20231', 'CRE1241', 7),

        ('CEGBCO20231', 'ENG4030', 8),
        ('CEGBCO20231', 'EMP1330', 8),
        ('CEGBCO20231', 'ENG4500', 8),
        ('CEGBCO20231', 'INF0312', 8),
        ('CEGBCO20231', 'INF0303', 8),
        ('CEGBCO20231', 'INF0305', 8),
        ('CEGBCO20231', 'JUR0205', 8),

        ('CEGBCO20231', 'ENG4134', 9),
        ('CEGBCO20231', 'CRE1275', 9),

        ('CEGBCO20231', 'ENG4135', 10),
        ('CEGBCO20231', 'ENG4153', 10)
    ) AS i (cod_curriculo, cod_disciplina, semestre)
    WHERE EXISTS (
        SELECT cod_disciplina
        FROM disciplinas
        WHERE cod_disciplina = i.cod_disciplina
    );
END