<script lang="ts">
  import Card from "$lib/components/card/Card.svelte";
  import Table from "$lib/components/table/Table.svelte";
</script>

<Card>
  <h5 class="mb-2 text-xl font-medium leading-tight text-neutral-800">Rules</h5>
  <div class=" max-h-80 overflow-auto">
    <Table
      on:rowClicked={() => {}}
      on:checkBoxAllClicked={(_) => {}}
      on:checkBoxClicked={(_) => {}}
      selectedIds={[]}
      columns={["Column name", "Column type", "Row index", "Rule"]}
      rows={[]}
    />
  </div>
</Card>
