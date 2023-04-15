<script lang="ts">
  import IconButton from "$lib/components/button/IconButton.svelte";

  export let title: string;
  export let backAction: (() => void) | null = null;
</script>

<div class="flex items-center mb-5">
  {#if backAction}
    <span class="mr-2">
      <IconButton icon="arrow_left" on:click={backAction} />
    </span>
  {/if}
  <h1 class="text-4xl font-extrabold tracking-tight">{title}</h1>
</div>
