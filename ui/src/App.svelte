<script>
  import Result from "./lib/Result.svelte";
  import Search from "./lib/Search.svelte";

  let resultValue = "";

  async function search(searchText) {
    console.log("Search Text: " + searchText);
    // const resp = await fetch("https://jsonplaceholder.typicode.com/todos/1");
    const resp = await fetch(
      "http://localhost:3000/api/search?query=" + searchText
    );
    let result = await resp.json();
    resultValue = JSON.stringify(result.value);
    console.log("Result Value: " + resultValue);
  }
</script>

<main class="main">
  <Search searchFunc={search} />
  <Result valueText={resultValue} />
</main>

<style>
  .main {
    flex: auto;
  }
</style>
