-- tabela curriculos
DELETE FROM curriculos WHERE cod_curriculo = 'CEGCCP20181';

INSERT INTO curriculos (cod_curriculo, nome_curriculo)
VALUES ('CEGCCP20181', 'Ciência da Computação (Bacharelado) - Currículo 18.1');

-- tabela de semestres
DELETE FROM semestres WHERE cod_curriculo = 'CEGCCP20181';

INSERT INTO semestres (cod_curriculo, cod_disciplina, semestre)
VALUES
    ('CEGCCP20181', 'INF1025', 1),
    ('CEGCCP20181', 'INF1012', 1),
    ('CEGCCP20181', 'INF1009', 1),
    ('CEGCCP20181', 'INF1031', 1),
    ('CEGCCP20181', 'MAT1157', 1),
    ('CEGCCP20181', 'MAT1320', 1),

    ('CEGCCP20181', 'CRE1100', 2),
    ('CEGCCP20181', 'INF0380', 2),
    ('CEGCCP20181', 'INF1403', 2),
    ('CEGCCP20181', 'INF1007', 2),
    ('CEGCCP20181', 'MAT1158', 2),
    ('CEGCCP20181', 'MAT1250', 2),

    ('CEGCCP20181', 'INF1383', 3),
    ('CEGCCP20181', 'INF1631', 3),
    ('CEGCCP20181', 'INF1010', 3),
    ('CEGCCP20181', 'INF1301', 3),
    ('CEGCCP20181', 'INF1018', 3),
    ('CEGCCP20181', 'MAT1162', 3),

    ('CEGCCP20181', 'CRE0710', 4),
    ('CEGCCP20181', 'INF1027', 4),
    ('CEGCCP20181', 'INF1029', 4),
    ('CEGCCP20181', 'INF1316', 4),
    ('CEGCCP20181', 'INF1629', 4),
    ('CEGCCP20181', 'INF1636', 4),

    ('CEGCCP20181', 'CRE1141', 5),
    ('CEGCCP20181', 'INF0377', 5),
    ('CEGCCP20181', 'INF1022', 5),
    ('CEGCCP20181', 'INF1028', 5),
    ('CEGCCP20181', 'INF1317', 5),
    ('CEGCCP20181', 'INF1721', 5),

    ('CEGCCP20181', 'INF1014', 6),
    ('CEGCCP20181', 'INF1033', 6),

    ('CEGCCP20181', 'CRE1175', 7),
    ('CEGCCP20181', 'FIL0300', 7),
    ('CEGCCP20181', 'INF1950', 7),

    ('CEGCCP20181', 'INF0381', 8),
    ('CEGCCP20181', 'INF1951', 8)
;