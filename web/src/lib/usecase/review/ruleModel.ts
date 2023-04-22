import type { TableRow } from "$lib/components/table/table-data";

export interface RuleModel {
  id: number;
  lookoutId: number;
  columnName: string;
  columnType: string;
  rowIndex: number;
  exactValue: string | null;
  greaterThan: string | null;
  lessThan: string | null;
  shouldBeNull: boolean;
}

export const convertRuleModelToTableData = (model: RuleModel): TableRow => {
  let rule: string = "";
  if (model.exactValue) {
    rule = `=${model.exactValue}`;
  } else if (model.shouldBeNull) {
    rule = "null";
  } else if (model.lessThan && model.greaterThan) {
    rule = `${model.greaterThan} - ${model.lessThan}`;
  } else if (model.lessThan) {
    rule = `<${model.lessThan}`;
  } else if (model.greaterThan) {
    rule = `>${model.greaterThan}`;
  }

  return {
    id: model.id,
    data: [
      { type: "text", value: model.columnName },
      { type: "text", value: model.columnType },
      { type: "number", value: model.rowIndex },
      { type: "text", value: rule },
    ],
  };
};
