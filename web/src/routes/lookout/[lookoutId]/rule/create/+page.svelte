<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import Button from "$lib/components/button/Button.svelte";
  import FormInputDropdown from "$lib/components/form/FormInputDropdown.svelte";
  import FormInputField from "$lib/components/form/FormInputField.svelte";
  import PageHeader from "$lib/components/header/PageHeader.svelte";
  import type { RuleType } from "../ruleType";
  import RuleInputs from "./RuleInputs.svelte";

  const onSubmit = async (event: SubmitEvent) => {};

  let ruleType: RuleType;
</script>

<PageHeader
  title="Create Rule"
  backAction={() => goto(`/lookout/${$page.params.lookoutId}`)}
/>

<form method="POST" class="w-full max-w-lg" on:submit|preventDefault={onSubmit}>
  <div class="flex flex-wrap -mx-3 mb-6">
    <FormInputField
      name="columnName"
      required
      label="Column name"
      placeholder="columnName"
    />
    <FormInputField
      name="rowIndex"
      required
      label="Row index"
      placeholder="0"
    />

    <FormInputDropdown
      name="columnType"
      required
      label="Column Type"
      options={[
        { name: "text", title: "text" },
        { name: "int", title: "int" },
        { name: "float", title: "float" },
      ]}
    />

    <FormInputDropdown
      name="ruleType"
      required
      bind:value={ruleType}
      label="Rule Type"
      options={[
        { name: "exact", title: "exact value" },
        { name: "null", title: "should be null" },
        { name: "less", title: "less than" },
        { name: "greater", title: "greater than" },
        { name: "between", title: "between" },
      ]}
    />

    <RuleInputs type={ruleType} />
    <div class="w-full" />
    <div class="mt-4 mx-3">
      <Button title="Create" type="submit" leadingIcon="check" />
    </div>
  </div>
</form>
