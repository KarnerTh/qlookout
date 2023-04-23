<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import Button from "$lib/components/button/Button.svelte";
  import FormInputDropdown from "$lib/components/form/FormInputDropdown.svelte";
  import FormInputField from "$lib/components/form/FormInputField.svelte";
  import PageHeader from "$lib/components/header/PageHeader.svelte";
  import { useRuleCreateMutation } from "$lib/usecase/review/mutation/createRule";
  import type { RuleType } from "../ruleType";
  import RuleInputs from "../RuleInputs.svelte";

  let ruleType: RuleType;
  let columnType: string;
  let columnTypeInput: "text" | "number";
  const createRule = useRuleCreateMutation();

  // update input type on select change
  $: {
    if (["int", "float"].includes(columnType)) {
      columnTypeInput = "number";
    } else {
      columnTypeInput = "text";
    }
  }

  const onSubmit = async (event: SubmitEvent) => {
    const data = new FormData(event.target as HTMLFormElement);
    const lookoutId = +$page.params.lookoutId;
    const columnName = data.get("columnName")?.toString();
    const columnType = data.get("columnType")?.toString();
    const rowIndex = +(data.get("rowIndex") ?? 0);
    const exactValue = data.get("exactValue")?.toString();
    const lessThan = data.get("lessThan")?.toString();
    const greaterThan = data.get("greaterThan")?.toString();
    const shouldBeNull = data.has("shouldBeNull");

    if (!lookoutId || !columnName || !columnType) {
      alert("Something went wrong");
      return;
    }

    const result = await createRule({
      variables: {
        lookoutId,
        columnName,
        columnType,
        rowIndex,
        exactValue,
        lessThan,
        greaterThan,
        shouldBeNull,
      },
    });

    if (result.errors) {
      alert("Something went wrong");
      return;
    }

    goto(`/lookout/${lookoutId}`, { state: { refetch: true } });
  };
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
      type="number"
      label="Row index"
      placeholder="0"
    />

    <FormInputDropdown
      name="columnType"
      required
      label="Column Type"
      bind:value={columnType}
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

    <RuleInputs {ruleType} inputType={columnTypeInput} />
    <div class="w-full" />
    <div class="mt-4 mx-3 w-full flex flex-row-reverse">
      <Button title="Create" type="submit" leadingIcon="check" />
    </div>
  </div>
</form>
