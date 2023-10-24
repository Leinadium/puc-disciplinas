<script lang="ts">
    import type { DisciplinaHorariosInfo } from "../../../types/data";

    export let codigo: string;
    export let professor: string;
    export let destino: string;
    export let vagas: number;
    export let horarios: DisciplinaHorariosInfo[];
    export let shf: number;

    // TODO: talvez deixar essa função de formatar
    //       mais bonita, como agrupar dias cujos
    //       horarios são iguais
    $: horariosMapeado = horarios.map(h => {
        if (h.dia.toUpperCase() === "XXX")
            return "Sem horário fixo"
        else
            return `${h.dia} ${h.inicio}-${h.fim}`
    });

    $: horariosFinal = horariosMapeado.join(", ") + (shf > 0 ? ` (${shf}h)`: '');
    
</script>

<a href="/#" class="turma-turma" on:click|preventDefault>
    <span class="turma-grande">{codigo}</span>   
    <span class="turma-dest overflow">{destino}</span>

    <span class="turma-vagas">
        {vagas}
        <span class="turma-vagas-texto">vagas</span>
    </span>
    <span class="turma-medio overflow">{professor}</span>
    
    <span class="turma-horarios">{horariosFinal}</span>
    
</a>

<style>
    a {
        box-sizing: border-box;
        text-align: center;
    }

    a {
        text-decoration: none;
        color: #000;
    }

    .overflow {
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }

    .turma-turma {
        height: 100%;
        max-width: 25%;
        padding: 5px;

        display: flex;
        flex-flow: column nowrap;
        justify-content: center;
        align-items: stretch;
        gap: 1%;

        background: #eee;
        border-radius: 15px;
    }

    .turma-turma:hover {
        background: #ddd;
    }

    .turma-grande {
        font-size: 1.3rem;
        font-weight: bold;
    }

    .turma-dest {
        font-size: 0.6rem;
        font-style: italic;
        margin-bottom: 0.1rem;
    }

    .turma-medio {
        font-size: 0.9rem;
    }

    .turma-vagas {
        font-size: 1.2rem;
        font-weight: bold;
        font-style: italic;
    }

    .turma-vagas-texto {
        font-size: 0.6rem;
        font-weight: normal;
    }

    .turma-horarios {
        font-size: 0.8rem;
        font-style: italic;
        white-space: nowrap;
    }

    /* .turma-row {
        display: flex;
        flex-flow: row nowrap;
        justify-content: center;
        align-items: flex-end;
        gap: 3%;
    } */
</style>