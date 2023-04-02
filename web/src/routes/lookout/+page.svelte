<script lang="ts">
  import type { PageData } from "./$types";

  export let data: PageData;
</script>

<h1>Lookout</h1>

<ul>
  {#each data.lookouts as { id, title }}
    <li><a href="/lookout/{id}">{title}</a></li>
  {/each}
</ul>
