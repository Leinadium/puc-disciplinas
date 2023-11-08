-- tabela curriculos
DELETE FROM curriculos WHERE cod_curriculo = 'CCPBCC20230';

INSERT INTO curriculos (cod_curriculo, nome_curriculo)
VALUES ('CCPBCC20230', 'Ciência da Computação (Bacharelado) - Currículo 23.0');

-- tabela de semestres
DELETE FROM semestres WHERE cod_curriculo = 'CCPBCC20230';

-- INSERT INTO semestres (cod_curriculo, cod_disciplina, semestre)
-- SELECT * FROM (VALUES
--     ('CCPBCC20230', 'CRE1200', 1),
--     ('CCPBCC20230', 'CTC4002', 1),
--     ('CCPBCC20230', 'INF1009', 1),
--     ('CCPBCC20230', 'FIS1034', 1),
--     ('CCPBCC20230', 'INF1012', 1),
--     ('CCPBCC20230', 'INF1034', 1),
--     ('CCPBCC20230', 'LET1920', 1),
--
--     ('CCPBCC20230', 'INF1037', 2),
--     ('CCPBCC20230', 'INF1039', 2),
--     ('CCPBCC20230', 'INF1403', 2),
--     ('CCPBCC20230', 'MAT4161', 2),
--     ('CCPBCC20230', 'MAT4200', 2),
--
--     ('CCPBCC20230', 'CRE0712', 3),
--     ('CCPBCC20230', 'INF1010', 3),
--     ('CCPBCC20230', 'INF1018', 3),
--     ('CCPBCC20230', 'INF1040', 3),
--     ('CCPBCC20230', 'INF1383', 3),
--     ('CCPBCC20230', 'MAT4162', 3),
--
--     ('CCPBCC20230', 'INF0381', 4),
--     ('CCPBCC20230', 'INF1027', 4),
--     ('CCPBCC20230', 'INF1041', 4),
--     ('CCPBCC20230', 'INF1316', 4),
--     ('CCPBCC20230', 'INF1636', 4),
--     ('CCPBCC20230', 'INF1320', 4),
--     ('CCPBCC20230', 'MAT4202', 4),
--
--     ('CCPBCC20230', 'CRE1241', 5),
--     ('CCPBCC20230', 'INF0377', 5),
--     ('CCPBCC20230', 'INF1022', 5),
--     ('CCPBCC20230', 'INF1029', 5),
--     ('CCPBCC20230', 'INF1631', 5),
--     ('CCPBCC20230', 'INF1640', 5),
--
--     ('CCPBCC20230', 'CRE1275', 6),
--     ('CCPBCC20230', 'INF0305', 6),
--     ('CCPBCC20230', 'INF0306', 6),
--     ('CCPBCC20230', 'INF1014', 6),
--     ('CCPBCC20230', 'INF1028', 6),
--     ('CCPBCC20230', 'INF1721', 6),
--     ('CCPBCC20230', 'INF1920', 6),
--
--     ('CCPBCC20230', 'INF0303', 7),
--     ('CCPBCC20230', 'INF0308', 7),
--     ('CCPBCC20230', 'INF1416', 7),
--     ('CCPBCC20230', 'INF1950', 7),
--
--     ('CCPBCC20230', 'FIL0300', 8),
--     ('CCPBCC20230', 'INF1951', 8)
-- ) AS i (cod_curriculo, cod_disciplina, semestre)
-- WHERE EXISTS (
--     SELECT cod_disciplina
--     FROM disciplinas
--     WHERE cod_disciplina = i.cod_disciplina
-- );

INSERT INTO semestres (cod_curriculo, cod_disciplina, semestre)
VALUES
    ('CCPBCC20230', 'CRE1200', 1),
    ('CCPBCC20230', 'CTC4002', 1),
    ('CCPBCC20230', 'INF1009', 1),
    ('CCPBCC20230', 'FIS1034', 1),
    ('CCPBCC20230', 'INF1012', 1),
    ('CCPBCC20230', 'INF1034', 1),
    ('CCPBCC20230', 'LET1920', 1),

    ('CCPBCC20230', 'INF1037', 2),
    ('CCPBCC20230', 'INF1039', 2),
    ('CCPBCC20230', 'INF1403', 2),
    ('CCPBCC20230', 'MAT4161', 2),
    ('CCPBCC20230', 'MAT4200', 2),

    ('CCPBCC20230', 'CRE0712', 3),
    ('CCPBCC20230', 'INF1010', 3),
    ('CCPBCC20230', 'INF1018', 3),
    ('CCPBCC20230', 'INF1040', 3),
    ('CCPBCC20230', 'INF1383', 3),
    ('CCPBCC20230', 'MAT4162', 3),

    ('CCPBCC20230', 'INF0381', 4),
    ('CCPBCC20230', 'INF1027', 4),
    ('CCPBCC20230', 'INF1041', 4),
    ('CCPBCC20230', 'INF1316', 4),
    ('CCPBCC20230', 'INF1636', 4),
    ('CCPBCC20230', 'INF1320', 4),
    ('CCPBCC20230', 'MAT4202', 4),

    ('CCPBCC20230', 'CRE1241', 5),
    ('CCPBCC20230', 'INF0377', 5),
    ('CCPBCC20230', 'INF1022', 5),
    ('CCPBCC20230', 'INF1029', 5),
    ('CCPBCC20230', 'INF1631', 5),
    ('CCPBCC20230', 'INF1640', 5),

    ('CCPBCC20230', 'CRE1275', 6),
    ('CCPBCC20230', 'INF0305', 6),
    ('CCPBCC20230', 'INF0306', 6),
    ('CCPBCC20230', 'INF1014', 6),
    ('CCPBCC20230', 'INF1028', 6),
    ('CCPBCC20230', 'INF1721', 6),
    ('CCPBCC20230', 'INF1920', 6),

    ('CCPBCC20230', 'INF0303', 7),
    ('CCPBCC20230', 'INF0308', 7),
    ('CCPBCC20230', 'INF1416', 7),
    ('CCPBCC20230', 'INF1950', 7),

    ('CCPBCC20230', 'FIL0300', 8),
    ('CCPBCC20230', 'INF1951', 8)
;