<script lang="ts">
  import IconButton from "$lib/components/button/IconButton.svelte";

  export let title: string;
  export let showBackButton: boolean = false;
</script>

<div class="flex items-center mb-5">
  {#if showBackButton}
    <span class="mr-2">
      <IconButton icon="arrow_left" on:click={() => history.back()} />
    </span>
  {/if}
  <h1 class="text-4xl font-extrabold tracking-tight">{title}</h1>
</div>
