<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import Button from "$lib/components/button/Button.svelte";
  import Card from "$lib/components/card/Card.svelte";
  import Table from "$lib/components/table/Table.svelte";
  import {
    convertRuleModelToTableData,
    type RuleModel,
  } from "$lib/usecase/review/ruleModel";

  export let rules: RuleModel[];
  const onRowClicked = (id: number): void => {
    goto(`/lookout/${$page.params.lookoutId}/rule/${id}/edit`);
  };
</script>

<Card>
  <div class="flex mb-2 justify-between">
    <h5 class="mb-2 text-lg font-medium leading-tight text-neutral-800">
      Rules
    </h5>
    <Button
      title="Add rule"
      leadingIcon="plus"
      on:click={() => goto(`/lookout/${$page.params.lookoutId}/rule/create`)}
    />
  </div>
  <div class=" max-h-80 overflow-auto">
    <Table
      on:rowClicked={(event) => onRowClicked(event.detail.id)}
      columns={["Column name", "Column type", "Row index", "Rule"]}
      rows={rules.map((item) => convertRuleModelToTableData(item))}
    />
  </div>
</Card>
