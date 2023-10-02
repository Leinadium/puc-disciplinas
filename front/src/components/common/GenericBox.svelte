<script lang="ts">
    import type { GenericColor } from "../../types/style";
    export let color: GenericColor = "default";
</script>

<div class="generic-box {color}">
    <slot></slot>
</div>

<style>
    .generic-box {
        box-sizing: border-box;
        width: max(100px, 10vw);
        height: 100px;
        border-radius: 10px;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        
        padding: 10px;
        margin: 10px;

        text-align: center;

        display: flex;
        flex-flow: column nowrap;
        justify-content: space-between;
    }

    .default {
        background-color: #aaa;
    }

    .yellow {
        background-color: #aa3;
    }

    .green {
        background-color: #5f5;
    }
</style>