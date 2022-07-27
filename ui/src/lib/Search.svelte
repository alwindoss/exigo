<script>
    let searchText = "";
    export let searchFunc;

    async function localSearch(params) {
        await searchFunc(searchText);
    }
</script>

<input type="text" bind:value={searchText} />
<button on:click={localSearch}>Search</button>
