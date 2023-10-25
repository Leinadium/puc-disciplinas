<script lang="ts">
    import { checkLogin } from "$lib/api";
	import { onMount } from "svelte";

    onMount(() => {
        checkLogin();
    })
</script>

<div id="check-login"></div>

<style>
    #check-login {
        display: none;
    }
</style>