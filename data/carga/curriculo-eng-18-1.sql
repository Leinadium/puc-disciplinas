-- tabela curriculos
DELETE FROM curriculos WHERE cod_curriculo = 'CEGBCO20181';

INSERT INTO curriculos (cod_curriculo, nome_curriculo)
VALUES ('CEGBCO20181', 'Eng. de Computação (Bacharelado) - Currículo 18.1');

-- tabela de semestres
DELETE FROM semestres WHERE cod_curriculo = 'CEGBCO20181';

INSERT INTO semestres (cod_curriculo, cod_disciplina, semestre)
SELECT * FROM (VALUES
    ('CEGBCO20181', 'ENG1000', 1),
    ('CEGBCO20181', 'FIS1025', 1),
    ('CEGBCO20181', 'INF1025', 1),
    ('CEGBCO20181', 'MAT1157', 1),
    ('CEGBCO20181', 'QUI1709', 1),
    ('CEGBCO20181', 'QUI1971', 1),

    ('CEGBCO20181', 'CRE1100', 2),
    ('CEGBCO20181', 'ENG1419', 2),
    ('CEGBCO20181', 'FIS1026', 2),
    ('CEGBCO20181', 'FIS1027', 3),
    ('CEGBCO20181', 'MAT1250', 2),
    ('CEGBCO20181', 'MAT1158', 2),
    ('CEGBCO20181', 'QUI1972', 2),

    ('CEGBCO20181', 'ENG1003', 3),
    ('CEGBCO20181', 'ENG1015', 3),
    ('CEGBCO20181', 'FIS1041', 3),
    ('CEGBCO20181', 'FIS1042', 3),
    ('CEGBCO20181', 'INF1007', 3),
    ('CEGBCO20181', 'INF1009', 3),
    ('CEGBCO20181', 'INF1012', 3),
    ('CEGBCO20181', 'MAT1162', 3),

    ('CEGBCO20181', 'CRE0710', 4),
    ('CEGBCO20181', 'FIS1051', 4),
    ('CEGBCO20181', 'FIS1052', 4),
    ('CEGBCO20181', 'INF1018', 4),
    ('CEGBCO20181', 'INF1383', 4),
    ('CEGBCO20181', 'MAT1154', 4),
    ('CEGBCO20181', 'MAT1320', 4),

    ('CEGBCO20181', 'ENG1007', 5),
    ('CEGBCO20181', 'ENG1011', 5),
    ('CEGBCO20181', 'ENG1400', 5),
    ('CEGBCO20181', 'ENG1420', 5),
    ('CEGBCO20181', 'INF1010', 5),
    ('CEGBCO20181', 'INF1022', 5),
    ('CEGBCO20181', 'INF1301', 5),

    ('CEGBCO20181', 'CRE1141', 6),
    ('CEGBCO20181', 'ENG1404', 6),
    ('CEGBCO20181', 'ENG1421', 6),
    ('CEGBCO20181', 'INF1019', 6),
    ('CEGBCO20181', 'INF1608', 6),
    ('CEGBCO20181', 'INF1631', 6),
    ('CEGBCO20181', 'INF1636', 6),

    ('CEGBCO20181', 'ENG1029', 7),
    ('CEGBCO20181', 'ENG1413', 7),
    ('CEGBCO20181', 'ENG1414', 7),
    ('CEGBCO20181', 'FIL0300', 7),
    ('CEGBCO20181', 'INF1640', 7),
    ('CEGBCO20181', 'INF1721', 7),

    ('CEGBCO20181', 'ENG1021', 8),
    ('CEGBCO20181', 'ENG1025', 8),
    ('CEGBCO20181', 'ENG1153', 8),
    ('CEGBCO20181', 'ENG1451', 8),
    ('CEGBCO20181', 'INF0303', 8),
    ('CEGBCO20181', 'JUR1016', 8),

    ('CEGBCO20181', 'CRE1175', 9),
    ('CEGBCO20181', 'ENG1023', 9),
    ('CEGBCO20181', 'ENG1132', 9),
    ('CEGBCO20181', 'ENG1448', 9),
    ('CEGBCO20181', 'INF1014', 9),

    ('CEGBCO20181', 'ENG1133', 10)
) AS i (cod_curriculo, cod_disciplina, semestre)
WHERE EXISTS (
    SELECT cod_disciplina
    FROM disciplinas
    WHERE cod_disciplina = i.cod_disciplina
);