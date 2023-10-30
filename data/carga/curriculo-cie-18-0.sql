-- tabela curriculos
DELETE FROM curriculos WHERE cod_curriculo = 'CEGCCP20180';

INSERT INTO curriculos (cod_curriculo, nome_curriculo)
VALUES ('CEGCCP20180', 'Ciência da Computação (Bacharelado) - Currículo 18.0');

-- tabela de semestres
DELETE FROM semestres WHERE cod_curriculo = 'CEGCCP20180';

INSERT INTO semestres (cod_curriculo, cod_disciplina, semestre)
VALUES
    ('CEGCCP20180', 'INF1025', 1),
    ('CEGCCP20180', 'INF1012', 1),
    ('CEGCCP20180', 'INF1009', 1),
    ('CEGCCP20180', 'INF1031', 1),
    ('CEGCCP20180', 'MAT1161', 1),
    ('CEGCCP20180', 'MAT1320', 1),

    ('CEGCCP20180', 'CRE1100', 2),
    ('CEGCCP20180', 'INF0380', 2),
    ('CEGCCP20180', 'INF1403', 2),
    ('CEGCCP20180', 'INF1007', 2),
    ('CEGCCP20180', 'MAT0211', 2),

    ('CEGCCP20180', 'INF1383', 3),
    ('CEGCCP20180', 'INF1631', 3),
    ('CEGCCP20180', 'INF1010', 3),
    ('CEGCCP20180', 'INF1301', 3),
    ('CEGCCP20180', 'INF1018', 3),
    ('CEGCCP20180', 'MAT1162', 3),

    ('CEGCCP20180', 'CRE0710', 4),
    ('CEGCCP20180', 'INF1027', 4),
    ('CEGCCP20180', 'INF1029', 4),
    ('CEGCCP20180', 'INF1316', 4),
    ('CEGCCP20180', 'INF1629', 4),
    ('CEGCCP20180', 'INF1636', 4),

    ('CEGCCP20180', 'CRE1141', 5),
    ('CEGCCP20180', 'INF0377', 5),
    ('CEGCCP20180', 'INF1022', 5),
    ('CEGCCP20180', 'INF1028', 5),
    ('CEGCCP20180', 'INF1317', 5),
    ('CEGCCP20180', 'INF1721', 5),

    ('CEGCCP20180', 'INF1014', 6),
    ('CEGCCP20180', 'INF1033', 6),

    ('CEGCCP20180', 'CRE1175', 7),
    ('CEGCCP20180', 'FIL0300', 7),
    ('CEGCCP20180', 'INF1950', 7),

    ('CEGCCP20180', 'INF0381', 8),
    ('CEGCCP20180', 'INF1951', 8)
;