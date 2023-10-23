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

<div class="turma-turma">
    <div class="turma-row">
        <span class="turma-grande">{codigo}</span>   
        <span class="turma-dest overflow">{destino}</span>
    </div>
    <span class="turma-vagas">
        {vagas}
        <span class="turma-vagas-texto">vagas</span>
    </span>
    <span class="turma-medio overflow">{professor}</span>
    
    <span class="turma-horarios">{horariosFinal}</span>
    
</div>

<style>
    div {
        box-sizing: border-box;
        text-align: center;
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

        cursor: pointer;
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

    .turma-row {
        display: flex;
        flex-flow: row nowrap;
        justify-content: center;
        align-items: flex-end;
        gap: 3%;
    }
</style>