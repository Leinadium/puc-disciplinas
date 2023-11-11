<div id="texto-abaixo">
    <p>
        <br>
        Este é um projeto de conclusão de curso de 
        Engenharia de Computação da PUC-Rio.
        <br><br>
        O objetivo é oferecer um sistema que integre tudo o que seja necessário
        para planejar seu semestre, e também recomendar as melhores opções
        de disciplinas, tanto obrigatórias ou eletivas.
        <br><br>
        Para isso, o sistema se utiliza se diversas fontes de informação,
        como o microhorário, o catálogo de ementas, currículos dos cursos,
        e históricos de alunos.
    </p>

    <div class="links">
        <a href="https://github.com/Leinadium/puc-disciplinas" target="_blank">Código fonte</a>
        <a href="https://github.com/Leinadium/" target="_blank">Sobre o autor</a>
        <a href="https://github.com/Leinadium/puc-tcc-disciplinas" target="_blank">Relatório (TCC)</a>
    </div>
    
</div>

<style>
    #texto-abaixo {
        width: 80%;
        height: 100%;

        color: var(--color-whiteff);

        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: center;
    }

    #texto-abaixo p {
        width: 60%;
        font-size: 1.2rem;
        font-style: italic;
        margin-bottom: 3%;
        text-align: center;
    }

    .links {
        width: 50%;

        display: flex;
        flex-flow: row nowrap;
        justify-content: space-around;
        align-items: center;
    }

    a {
        text-decoration: none;
        color: var(--color-whiteff);
        padding: 0.5rem;
        border-radius: var(--border-radius);
        background: var(--color-anal-2);
    }
</style>