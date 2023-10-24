<script lang="ts">
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();
    const close = () => { dispatch('close');}
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div id="popup">
    <div id="popup-background" on:click|preventDefault={close}></div>
    <div id="popup-content">
        <slot></slot>
    </div>
</div>

<style>
    #popup {
        position: fixed;
        top: 0;
        left: 0;
        z-index: 1000;
        width: 100vw;
        height: 100vh;
        background-color: rgba(0, 0, 0, 0.5);
        display: flex;
        justify-content: center;
        align-items: center;
    }

    #popup-background {
        position: absolute;
        width: 100%;
        height: 100%;
        z-index: 2;
    }

    #popup-content {
        z-index: 3;
    }
</style>
