-- tabela curriculos
DELETE FROM curriculos WHERE cod_curriculo = 'CEGBCO20180';

INSERT INTO curriculos (cod_curriculo, nome_curriculo)
VALUES ('CEGBCO20180', 'Eng. de Computação (Bacharelado) - Currículo 18.0');

-- tabela de semestres
DELETE FROM semestres WHERE cod_curriculo = 'CEGBCO20180';

INSERT INTO semestres (cod_curriculo, cod_disciplina, semestre)
SELECT * FROM (VALUES
    ('CEGBCO20180', 'ENG1000', 1),
    ('CEGBCO20180', 'ENG1003', 1),
    ('CEGBCO20180', 'FIS1033', 1),
    ('CEGBCO20180', 'FIS1034', 1),
    ('CEGBCO20180', 'INF1025', 1),
    ('CEGBCO20180', 'MAT1260', 1),
    ('CEGBCO20180', 'MAT1161', 1),

    ('CEGBCO20180', 'CRE1100', 2),
    ('CEGBCO20180', 'ENG1419', 2),
    ('CEGBCO20180', 'FIS1041', 2),
    ('CEGBCO20180', 'FIS1042', 3),
    ('CEGBCO20180', 'MAT1162', 2),
    ('CEGBCO20180', 'QUI1709', 2),
    ('CEGBCO20180', 'QUI1740', 2),

    ('CEGBCO20180', 'CRE0710', 3),
    ('CEGBCO20180', 'ENG1015', 3),
    ('CEGBCO20180', 'FIS1051', 3),
    ('CEGBCO20180', 'FIS1052', 3),
    ('CEGBCO20180', 'INF1007', 3),
    ('CEGBCO20180', 'INF1009', 3),
    ('CEGBCO20180', 'INF1012', 3),
    ('CEGBCO20180', 'MAT1154', 3),

    ('CEGBCO20180', 'ENG1007', 4),
    ('CEGBCO20180', 'ENG1011', 4),
    ('CEGBCO20180', 'ENG1400', 4),
    ('CEGBCO20180', 'ENG1420', 4),
    ('CEGBCO20180', 'INF1018', 4),
    ('CEGBCO20180', 'INF1383', 4),
    ('CEGBCO20180', 'MAT1320', 4),

    ('CEGBCO20180', 'ENG1404', 5),
    ('CEGBCO20180', 'ENG1421', 5),
    ('CEGBCO20180', 'INF1010', 5),
    ('CEGBCO20180', 'INF1019', 5),
    ('CEGBCO20180', 'INF1022', 5),
    ('CEGBCO20180', 'INF1301', 5),

    ('CEGBCO20180', 'CRE1141', 6),
    ('CEGBCO20180', 'ENG1413', 6),
    ('CEGBCO20180', 'ENG1414', 6),
    ('CEGBCO20180', 'INF1608', 6),
    ('CEGBCO20180', 'INF1631', 6),
    ('CEGBCO20180', 'INF1636', 6),

    ('CEGBCO20180', 'ENG1021', 7),
    ('CEGBCO20180', 'ENG1025', 7),
    ('CEGBCO20180', 'ENG1029', 7),
    ('CEGBCO20180', 'FIL0300', 7),
    ('CEGBCO20180', 'INF0303', 7),
    ('CEGBCO20180', 'INF1640', 7),
    ('CEGBCO20180', 'INF1721', 7),

    ('CEGBCO20180', 'CRE1175', 8),
    ('CEGBCO20180', 'ENG1153', 8),
    ('CEGBCO20180', 'ENG1448', 8),
    ('CEGBCO20180', 'ENG1451', 8),
    ('CEGBCO20180', 'JUR1016', 8),

    ('CEGBCO20180', 'ENG1023', 9),
    ('CEGBCO20180', 'ENG1132', 9),
    ('CEGBCO20180', 'INF1014', 9),

    ('CEGBCO20180', 'ENG1133', 10)
) AS i (cod_curriculo, cod_disciplina, semestre)
WHERE EXISTS (
    SELECT cod_disciplina
    FROM disciplinas
    WHERE cod_disciplina = i.cod_disciplina
);