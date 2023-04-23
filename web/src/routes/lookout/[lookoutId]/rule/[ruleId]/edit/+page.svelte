<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import Button from "$lib/components/button/Button.svelte";
  import FormInputDropdown from "$lib/components/form/FormInputDropdown.svelte";
  import FormInputField from "$lib/components/form/FormInputField.svelte";
  import PageHeader from "$lib/components/header/PageHeader.svelte";
  import { useRule } from "$lib/usecase/review/query/getRule";
  import { useRuleUpdateMutation } from "$lib/usecase/review/mutation/updateRule";
  import { useRuleDeleteMutation } from "$lib/usecase/review/mutation/deleteRule";
  import RuleInputs from "../../RuleInputs.svelte";
  import type { RuleType } from "../../ruleType";
  import LoadingSpinner from "$lib/components/loading/LoadingSpinner.svelte";
  import DialogButton from "$lib/components/dialog/DialogButton.svelte";

  let ruleType: RuleType | undefined;
  let columnType: string | undefined;
  let columnTypeInput: "text" | "number";

  const updateRule = useRuleUpdateMutation();
  const deleteRule = useRuleDeleteMutation();
  const lookoutId = +$page.params.lookoutId;
  const ruleId = +$page.params.ruleId;
  const rule = useRule(ruleId);

  // update input type on select change
  $: {
    if (!columnType) break $;
    if (["int", "float"].includes(columnType)) {
      columnTypeInput = "number";
    } else {
      columnTypeInput = "text";
    }
  }

  // prefill selects
  $: {
    if (columnType === undefined && $rule.data) {
      columnType = $rule.data.rule.columnType;
    }
    if (ruleType === undefined && $rule.data) {
      if ($rule.data.rule.exactValue !== null) {
        ruleType = "exact";
      } else if ($rule.data.rule.shouldBeNull === true) {
        ruleType = "null";
      } else if (
        $rule.data.rule.lessThan !== null &&
        $rule.data.rule.greaterThan !== null
      ) {
        ruleType = "between";
      } else if ($rule.data.rule.lessThan !== null) {
        ruleType = "less";
      } else if ($rule.data.rule.greaterThan !== null) {
        ruleType = "greater";
      }
    }
  }

  const onSubmit = async (event: SubmitEvent) => {
    const data = new FormData(event.target as HTMLFormElement);
    const columnName = data.get("columnName")?.toString();
    const columnType = data.get("columnType")?.toString();
    const rowIndex = +(data.get("rowIndex") ?? 0);
    const exactValue = data.get("exactValue")?.toString();
    const lessThan = data.get("lessThan")?.toString();
    const greaterThan = data.get("greaterThan")?.toString();
    const shouldBeNull = data.has("shouldBeNull");

    const result = await updateRule({
      variables: {
        id: ruleId,
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

  const onDelete = async () => {
    const result = await deleteRule({ variables: { id: ruleId } });
    if (result.errors) {
      alert("Something went wrong");
      return;
    }

    goto(`/lookout/${lookoutId}`, { state: { refetch: true } });
  };
</script>

<PageHeader
  title="Update Rule"
  backAction={() => goto(`/lookout/${lookoutId}`)}
/>

{#if $rule.loading}
  <div class="flex justify-center pt-9">
    <LoadingSpinner />
  </div>
{:else if $rule.data}
  <form
    method="POST"
    class="w-full max-w-lg"
    on:submit|preventDefault={onSubmit}
  >
    <div class="flex flex-wrap -mx-3 mb-6">
      <FormInputField
        name="columnName"
        value={$rule.data.rule.columnName}
        required
        label="Column name"
        placeholder="columnName"
      />
      <FormInputField
        name="rowIndex"
        value={$rule.data.rule.rowIndex}
        type="number"
        required
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

      <RuleInputs
        {ruleType}
        value={$rule.data.rule}
        inputType={columnTypeInput}
      />
      <div class="w-full" />
      <div class="mt-4 mx-3 w-full flex flex-row-reverse gap-2">
        <DialogButton
          buttonText="Delete"
          buttonIcon="trash"
          buttonType="red"
          dialogTitle="Delete Rule"
          dialogDescription="Do you realy want to delete this rule? This action cannot be undone"
          actionButtons={[
            {
              title: "Cancel",
              onClick() {},
            },
            {
              title: "Delete",
              buttonType: "red",
              onClick() {
                onDelete();
              },
            },
          ]}
        />
        <Button title="Update" type="submit" leadingIcon="check" />
      </div>
    </div>
  </form>
{/if}
