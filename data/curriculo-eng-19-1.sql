-- tabela curriculos
DELETE FROM curriculos WHERE cod_curriculo = 2;

INSERT INTO curriculos (cod_curriculo, nome_curriculo)
VALUES (2, 'Engenharia de Computação 2019 - 1');

-- tabela de semestres
DELETE FROM semestres WHERE cod_curriculo = 2;

INSERT INTO semestres (cod_curriculo, cod_disciplina, semestre)
VALUES
    (2, 'ENG1000', 1),
    (2, 'FIS1025', 1),
    (2, 'INF1025', 1),
    (2, 'MAT1157', 1),
    (2, 'QUI1709', 1),
    (2, 'QUI1971', 1),

    (2, 'CRE1100', 2),
    (2, 'ENG1419', 2),
    (2, 'FIS1026', 2),
    (2, 'FIS1027', 3),
    (2, 'MAT1250', 2),
    (2, 'MAT1158', 2),
    (2, 'QUI1972', 2),

    (2, 'ENG1003', 3),
    (2, 'ENG1015', 3),
    (2, 'FIS1041', 3),
    (2, 'FIS1042', 3),
    (2, 'INF1007', 3),
    (2, 'INF1009', 3),
    (2, 'INF1012', 3),
    (2, 'MAT1162', 3),

    (2, 'CRE0710', 4),
    (2, 'FIS1051', 4),
    (2, 'FIS1052', 4),
    (2, 'INF1018', 4),
    (2, 'INF1383', 4),
    (2, 'MAT1154', 4),
    (2, 'MAT1320', 4),

    (2, 'ENG1007', 5),
    (2, 'ENG1011', 5),
    (2, 'ENG1400', 5),
    (2, 'ENG1420', 5),
    (2, 'INF1010', 5),
    (2, 'INF1022', 5),
    (2, 'INF1301', 5),

    (2, 'CRE1141', 6),
    (2, 'ENG1404', 6),
    (2, 'ENG1421', 6),
    (2, 'INF1019', 6),
    (2, 'INF1608', 6),
    (2, 'INF1631', 6),
    (2, 'INF1636', 6),

    (2, 'ENG1029', 7),
    (2, 'ENG1413', 7),
    (2, 'ENG1414', 7),
    (2, 'FIL0300', 7),
    (2, 'INF1640', 7),
    (2, 'INF1721', 7),

    (2, 'ENG1021', 8),
    (2, 'ENG1025', 8),
    (2, 'ENG1153', 8),
    (2, 'ENG1451', 8),
    (2, 'INF0303', 8),
    (2, 'JUR1016', 8),

    (2, 'CRE1175', 9),
    (2, 'ENG1023', 9),
    (2, 'ENG1132', 9),
    (2, 'ENG1448', 9),
    (2, 'INF1014', 9),

    (2, 'ENG1133', 10)
;