-- tabela curriculos
DELETE FROM curriculos WHERE cod_curriculo = 0;

INSERT INTO curriculos (cod_curriculo, nome_curriculo)
VALUES (0, 'Engenharia de Computação 2023');

-- tabela de semestres
DELETE FROM semestres WHERE cod_curriculo = 0;

INSERT INTO semestres (cod_curriculo, cod_disciplina, semestre)
VALUES
    (0, 'ENG4010', 1),
    (0, 'EMP1310', 1),
    (0, 'ENG4025', 1),
    (0, 'CTC4001', 1),
    (0, 'CTC4002', 1),
    (0, 'CTC4003', 1),
    (0, 'CRE1200', 1),

    (0, 'ENG4021', 2),
    (0, 'EMP1320', 2),
    (0, 'INF1037', 2),
    (0, 'MAT4161', 3),
    (0, 'FIS4001', 2),
    (0, 'MAT4200', 2),
    (0, 'INF1012', 2),
    (0, 'INF1009', 2),

    (0, 'ENG4033', 3),
    (0, 'ENG4431', 3),
    (0, 'ENG4420', 3),
    (0, 'MAT4162', 3),
    (0, 'FIS4002', 3),
    (0, 'MAT4202', 3),
    (0, 'INF1383', 3),
    (0, 'ENG4007', 3),

    (0, 'ENG4040', 4),
    (0, 'ENG4502', 4),
    (0, 'ENG4501', 4),
    (0, 'ENG4402', 4),
    (0, 'MAT4174', 4),
    (0, 'INF1010', 4),
    (0, 'INF1018', 4),
    (0, 'ENG4011', 4),

    (0, 'ENG4051', 5),
    (0, 'ENG4013', 5),
    (0, 'ENG4421', 5),
    (0, 'MAT1320', 5),
    (0, 'FIS4003', 5),
    (0, 'INF1301', 5),

    (0, 'ENG4061', 6),
    (0, 'ENG4448', 6),
    (0, 'INF1022', 6),
    (0, 'INF1631', 6),
    (0, 'INF1636', 6),
    (0, 'INF1316', 6),
    (0, 'CRE0712', 6),

    (0, 'ENG4405', 7),
    (0, 'INF1721', 7),
    (0, 'INF1041', 7),
    (0, 'INF1029', 7),
    (0, 'INF1640', 7),
    (0, 'FIL0300', 7),
    (0, 'CRE1241', 7),

    (0, 'ENG4030', 8),
    (0, 'EMP1330', 8),
    (0, 'ENG4500', 8),
    (0, 'INF0312', 8),
    (0, 'INF0303', 8),
    (0, 'INF0305', 8),
    (0, 'JUR0205', 8),

    (0, 'ENG4134', 9),
    (0, 'CRE1275', 9),

    (0, 'ENG4135', 10),
    (0, 'ENG4153', 10)
;