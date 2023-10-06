import bs4
from typing import NamedTuple


class Historia(NamedTuple):
    periodo: int
    disciplina: str
    grau: int | None


def parse_historico(html: bytes) -> tuple[list[Historia], str | None]:
    soup = bs4.BeautifulSoup(html, features='html.parser')
    ret: list[Historia] = []
    cod_curriculo: str = ""

    # pega os dados do curso
    tabela: bs4.Tag = soup.find(id='tblDadosCursoAluno')
    body: bs4.Tag = tabela.find('tbody')

    # tenta montar o código do curriculo
    try:
        curso = body.find('td', id='Curso').text.strip()
        habilitacao = body.find('td', id='Habilitacao').text.strip()
        ano = body.find('td', id='AnoCurriculo').text.strip()
        codigo = body.find('td', id='CodigoCurriculo').text.strip()
        cod_curriculo = (curso + habilitacao + ano + codigo).upper()
    except AttributeError:
        pass

    # pega a tabela de disciplinas
    tabela: bs4.Tag = soup.find(id='grvEspelhoHistorico')
    body: bs4.Tag = tabela.find('tbody')

    ultimo_periodo: str = ""
    i_periodo: int = 0

    todas_disciplinas: set[str] = set()

    # pega as linhas da tabela
    # cada linha é um <tr> com class "tab_linha"
    for linha in body.find_all('tr'):
        periodo_tag: bs4.Tag = linha.find('span', class_='tab_col_valor_Periodo')
        disciplina_tag: bs4.Tag = linha.find('span', class_='tab_col_valor_Disciplina')
        grau_tag: bs4.Tag = linha.find('span', class_='tab_col_valor_Grau')
        situacao_tag: bs4.Tag = linha.find('span', class_='tab_col_valor_Situacao')
        grupo_tag: bs4.Tag = linha.find('span', class_='tab_col_valor_Grupo')

        if None not in (periodo_tag, disciplina_tag, grau_tag, situacao_tag):
            # converte o nome da disciplina para o codigo
            disciplina_str = disciplina_tag.text.strip()[:7]

            # ignora se a disciplina ja foi verificada
            if disciplina_str in todas_disciplinas:
                continue

            # verifica se é uma situação valida
            situacao = situacao_tag.text.strip()
            if situacao not in ('AP', 'CP', 'RF', 'RM'):
                print(f"ignorando disciplina {disciplina_tag.text} com situacao {situacao}")
                continue

            # faz uma conversao simples, jogando fora o valor
            # mas adicionando +1 se for um novo periodo
            periodo = periodo_tag.text.strip()
            if periodo != ultimo_periodo:
                ultimo_periodo = periodo
                i_periodo += 1
            # converte o grau para um numero (0 - 100)
            grau = grau_tag.text.strip()
            try:
                grau_int = int(float(grau.replace(',', '.')) * 10)
            except ValueError:
                grau_int = None

            # adiciona na lista
            ret.append(Historia(i_periodo, disciplina_str, grau_int))

            # se grupo nao for nulo, adiciona o grupo como se fosse uma disciplina
            # pois o grupo pode estar no curriculo
            if grupo_tag is not None:
                grupo_str = grupo_tag.text.strip()[:7]
                if grupo_str not in todas_disciplinas:
                    ret.append(Historia(i_periodo, grupo_str, grau_int))
                    todas_disciplinas.add(grupo_str)

    return ret, cod_curriculo
