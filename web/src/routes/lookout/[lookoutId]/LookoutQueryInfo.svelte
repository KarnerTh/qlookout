<script lang="ts">
  import Card from "$lib/components/card/Card.svelte";

  export let query: string;
</script>

<Card>
  <h5 class="mb-2 text-lg font-medium leading-tight text-neutral-800">Query</h5>
  <div class=" max-h-80 overflow-auto">
    <code class="text-sm text-neutral-600">
      {query}
    </code>
  </div>
</Card>
