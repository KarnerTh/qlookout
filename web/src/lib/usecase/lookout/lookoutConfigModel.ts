import type { TableRow } from "$lib/components/table/table-data";
import type { RuleModel } from "../review/ruleModel";

export interface LookoutConfigModel {
  id: number;
  name: string;
  query: string;
  cron: string;
  notifyLocal: boolean;
  notifyMail: boolean;
}

export type LookoutConfigDetailModel = LookoutConfigModel & {
  rules: RuleModel[];
};

export const convertLookoutConfigModelToTableData = (
  model: LookoutConfigModel
): TableRow => {
  return {
    id: model.id,
    data: [
      { type: "text", value: model.name },
      { type: "text", value: model.cron },
      { type: "number", value: 0 },
      { type: "boolean", value: model.notifyLocal },
      { type: "boolean", value: model.notifyMail },
    ],
  };
};
