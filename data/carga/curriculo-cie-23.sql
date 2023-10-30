-- tabela curriculos
DELETE FROM curriculos WHERE cod_curriculo = 'CEGCCP20230';

INSERT INTO curriculos (cod_curriculo, nome_curriculo)
VALUES ('CEGCCP20230', 'Ciência da Computação (Bacharelado) - Currículo 23.0');

-- tabela de semestres
DELETE FROM semestres WHERE cod_curriculo = 'CEGCCP20230';

INSERT INTO semestres (cod_curriculo, cod_disciplina, semestre)
VALUES
    ('CEGCCP20230', 'CRE1200', 1),
    ('CEGCCP20230', 'CTC4002', 1),
    ('CEGCCP20230', 'INF1009', 1),
    ('CEGCCP20230', 'FIS1034', 1),
    ('CEGCCP20230', 'INF1012', 1),
    ('CEGCCP20230', 'INF1034', 1),
    ('CEGCCP20230', 'LET1920', 1),

    ('CEGCCP20230', 'INF1037', 2),
    ('CEGCCP20230', 'INF1039', 2),
    ('CEGCCP20230', 'INF1403', 2),
    ('CEGCCP20230', 'MAT4161', 2),
    ('CEGCCP20230', 'MAT4200', 2),

    ('CEGCCP20230', 'CRE0712', 3),
    ('CEGCCP20230', 'INF1010', 3),
    ('CEGCCP20230', 'INF1018', 3),
    ('CEGCCP20230', 'INF1040', 3),
    ('CEGCCP20230', 'INF1383', 3),
    ('CEGCCP20230', 'MAT4162', 3),

    ('CEGCCP20230', 'INF0381', 4),
    ('CEGCCP20230', 'INF1027', 4),
    ('CEGCCP20230', 'INF1041', 4),
    ('CEGCCP20230', 'INF1316', 4),
    ('CEGCCP20230', 'INF1636', 4),
    ('CEGCCP20230', 'INF1320', 4),
    ('CEGCCP20230', 'MAT4202', 4),

    ('CEGCCP20230', 'CRE1241', 5),
    ('CEGCCP20230', 'INF0377', 5),
    ('CEGCCP20230', 'INF1022', 5),
    ('CEGCCP20230', 'INF1029', 5),
    ('CEGCCP20230', 'INF1631', 5),
    ('CEGCCP20230', 'INF1640', 5),

    ('CEGCCP20230', 'CRE1275', 6),
    ('CEGCCP20230', 'INF0305', 6),
    ('CEGCCP20230', 'INF0306', 6),
    ('CEGCCP20230', 'INF1014', 6),
    ('CEGCCP20230', 'INF1028', 6),
    ('CEGCCP20230', 'INF1721', 6),
    ('CEGCCP20230', 'INF1920', 6),

    ('CEGCCP20230', 'INF0303', 7),
    ('CEGCCP20230', 'INF0308', 7),
    ('CEGCCP20230', 'INF1416', 7),
    ('CEGCCP20230', 'INF1950', 7),

    ('CEGCCP20230', 'FIL0300', 8),
    ('CEGCCP20230', 'INF1951', 8)
;