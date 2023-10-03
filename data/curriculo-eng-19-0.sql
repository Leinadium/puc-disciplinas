-- tabela curriculos
DELETE FROM curriculos WHERE cod_curriculo = 1;

INSERT INTO curriculos (cod_curriculo, nome_curriculo)
VALUES (1, 'Engenharia de Computação 2019 - 0');

-- tabela de semestres
DELETE FROM semestres WHERE cod_curriculo = 1;

INSERT INTO semestres (cod_curriculo, cod_disciplina, semestre)
VALUES
    (1, 'ENG1000', 1),
    (1, 'ENG1003', 1),
    (1, 'FIS1033', 1),
    (1, 'FIS1034', 1),
    (1, 'INF1025', 1),
    (1, 'MAT1260', 1),
    (1, 'MAT1161', 1),

    (1, 'CRE1100', 2),
    (1, 'ENG1419', 2),
    (1, 'FIS1041', 2),
    (1, 'FIS1042', 3),
    (1, 'MAT1162', 2),
    (1, 'QUI1709', 2),
    (1, 'QUI1740', 2),

    (1, 'CRE0710', 3),
    (1, 'ENG1015', 3),
    (1, 'FIS1051', 3),
    (1, 'FIS1052', 3),
    (1, 'INF1007', 3),
    (1, 'INF1009', 3),
    (1, 'INF1012', 3),
    (1, 'MAT1154', 3),

    (1, 'ENG1007', 4),
    (1, 'ENG1011', 4),
    (1, 'ENG1400', 4),
    (1, 'ENG1420', 4),
    (1, 'INF1018', 4),
    (1, 'INF1383', 4),
    (1, 'MAT1320', 4),

    (1, 'ENG1404', 5),
    (1, 'ENG1421', 5),
    (1, 'INF1010', 5),
    (1, 'INF1019', 5),
    (1, 'INF1022', 5),
    (1, 'INF1301', 5),

    (1, 'CRE1141', 6),
    (1, 'ENG1413', 6),
    (1, 'ENG1414', 6),
    (1, 'INF1608', 6),
    (1, 'INF1631', 6),
    (1, 'INF1636', 6),

    (1, 'ENG1021', 7),
    (1, 'ENG1025', 7),
    (1, 'ENG1029', 7),
    (1, 'FIL0300', 7),
    (1, 'INF0303', 7),
    (1, 'INF1640', 7),
    (1, 'INF1721', 7),

    (1, 'CRE1175', 8),
    (1, 'ENG1153', 8),
    (1, 'ENG1448', 8),
    (1, 'ENG1451', 8),
    (1, 'JUR1016', 8),

    (1, 'ENG1023', 9),
    (1, 'ENG1132', 9),
    (1, 'INF1014', 9),

    (1, 'ENG1133', 10)
;